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
		"ap-northeast-1":              dara.String("dcdn.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("dcdn.aliyuncs.com"),
		"ap-south-1":                  dara.String("dcdn.aliyuncs.com"),
		"ap-southeast-1":              dara.String("dcdn.aliyuncs.com"),
		"ap-southeast-2":              dara.String("dcdn.aliyuncs.com"),
		"ap-southeast-3":              dara.String("dcdn.aliyuncs.com"),
		"ap-southeast-5":              dara.String("dcdn.aliyuncs.com"),
		"cn-beijing":                  dara.String("dcdn.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("dcdn.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("dcdn.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("dcdn.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("dcdn.aliyuncs.com"),
		"cn-chengdu":                  dara.String("dcdn.aliyuncs.com"),
		"cn-edge-1":                   dara.String("dcdn.aliyuncs.com"),
		"cn-fujian":                   dara.String("dcdn.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("dcdn.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("dcdn.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("dcdn.aliyuncs.com"),
		"cn-hongkong":                 dara.String("dcdn.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("dcdn.aliyuncs.com"),
		"cn-huhehaote":                dara.String("dcdn.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("dcdn.aliyuncs.com"),
		"cn-qingdao":                  dara.String("dcdn.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("dcdn.aliyuncs.com"),
		"cn-shanghai":                 dara.String("dcdn.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("dcdn.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("dcdn.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("dcdn.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("dcdn.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("dcdn.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("dcdn.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("dcdn.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("dcdn.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("dcdn.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("dcdn.aliyuncs.com"),
		"cn-wuhan":                    dara.String("dcdn.aliyuncs.com"),
		"cn-yushanfang":               dara.String("dcdn.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("dcdn.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("dcdn.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("dcdn.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("dcdn.aliyuncs.com"),
		"eu-central-1":                dara.String("dcdn.aliyuncs.com"),
		"eu-west-1":                   dara.String("dcdn.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("dcdn.aliyuncs.com"),
		"me-east-1":                   dara.String("dcdn.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("dcdn.aliyuncs.com"),
		"us-east-1":                   dara.String("dcdn.aliyuncs.com"),
		"us-west-1":                   dara.String("dcdn.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dcdn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}) (_result map[string]interface{}, _err error) {
	request_ := dara.NewRequest()
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
	response_, _err := dara.DoRequest(request_, nil)
	if _err != nil {
		return nil, _err
	}

	_result, _err = _postOSSObject_opResponse(response_)
	if _err != nil {
		return nil, _err
	}

	return _result, nil
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
// Adds a domain name to accelerate. You can specify only one domain name in each request.
//
// Description:
//
// > 	- Dynamic Content Delivery Network (DCDN) is activated.
//
// > 	- Internet content provider (ICP) filing is complete for the accelerated domain name.
//
// > 	- If the content of the origin server is not stored on Alibaba Cloud, the content must be reviewed. After you submit the request, the review is complete by the end of the following business day.
//
// > 	- You can call this operation up to 30 times per second per account.
//
// @param request - AddDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDcdnDomainResponse
func (client *Client) AddDcdnDomainWithOptions(request *AddDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *AddDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionType) {
		query["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name to accelerate. You can specify only one domain name in each request.
//
// Description:
//
// > 	- Dynamic Content Delivery Network (DCDN) is activated.
//
// > 	- Internet content provider (ICP) filing is complete for the accelerated domain name.
//
// > 	- If the content of the origin server is not stored on Alibaba Cloud, the content must be reviewed. After you submit the request, the review is complete by the end of the following business day.
//
// > 	- You can call this operation up to 30 times per second per account.
//
// @param request - AddDcdnDomainRequest
//
// @return AddDcdnDomainResponse
func (client *Client) AddDcdnDomain(request *AddDcdnDomainRequest) (_result *AddDcdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDcdnDomainResponse{}
	_body, _err := client.AddDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a domain name to IPA. You can specify only one domain name in each request.
//
// Description:
//
// >
//
//   - Make sure that the IPA service is activated before you add a domain name to accelerate.
//
//   - Make sure that the Internet content provider (ICP) filling is complete for the domain name to accelerate.
//
//   - If the content on the origin server is not stored on Alibaba Cloud, the content must be reviewed. The review is complete by the end of the next business day after you submit the request.
//
//   - You can call this operation up to 10 times per second per user.
//
// @param request - AddDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDcdnIpaDomainResponse
func (client *Client) AddDcdnIpaDomainWithOptions(request *AddDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *AddDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name to IPA. You can specify only one domain name in each request.
//
// Description:
//
// >
//
//   - Make sure that the IPA service is activated before you add a domain name to accelerate.
//
//   - Make sure that the Internet content provider (ICP) filling is complete for the domain name to accelerate.
//
//   - If the content on the origin server is not stored on Alibaba Cloud, the content must be reviewed. The review is complete by the end of the next business day after you submit the request.
//
//   - You can call this operation up to 10 times per second per user.
//
// @param request - AddDcdnIpaDomainRequest
//
// @return AddDcdnIpaDomainResponse
func (client *Client) AddDcdnIpaDomain(request *AddDcdnIpaDomainRequest) (_result *AddDcdnIpaDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDcdnIpaDomainResponse{}
	_body, _err := client.AddDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds one or more domain names to DCDN at a time.
//
// Description:
//
// *Prerequisites**:
//
//   - The [DCDN service is activated](https://help.aliyun.com/document_detail/64926.html).
//
//   - Internet content provider (ICP) filing is complete for the accelerated domain names.
//
// > 	- If the content of the origin server is not stored on Alibaba Cloud, the content must be reviewed. After you submit the request, the review is complete by the end of the following business day.
//
// >	- You can specify up to 50 domain names in each request.
//
// >	- You can call this operation up to 30 times per second per account.
//
// @param request - BatchAddDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddDcdnDomainResponse
func (client *Client) BatchAddDcdnDomainWithOptions(request *BatchAddDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchAddDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAddDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAddDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more domain names to DCDN at a time.
//
// Description:
//
// *Prerequisites**:
//
//   - The [DCDN service is activated](https://help.aliyun.com/document_detail/64926.html).
//
//   - Internet content provider (ICP) filing is complete for the accelerated domain names.
//
// > 	- If the content of the origin server is not stored on Alibaba Cloud, the content must be reviewed. After you submit the request, the review is complete by the end of the following business day.
//
// >	- You can specify up to 50 domain names in each request.
//
// >	- You can call this operation up to 30 times per second per account.
//
// @param request - BatchAddDcdnDomainRequest
//
// @return BatchAddDcdnDomainResponse
func (client *Client) BatchAddDcdnDomain(request *BatchAddDcdnDomainRequest) (_result *BatchAddDcdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchAddDcdnDomainResponse{}
	_body, _err := client.BatchAddDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates Web Application Firewall (WAF) protection rules.
//
// Description:
//
// >  You can call this operation up to 20 times per second per account.
//
// @param request - BatchCreateDcdnWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateDcdnWafRulesResponse
func (client *Client) BatchCreateDcdnWafRulesWithOptions(request *BatchCreateDcdnWafRulesRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateDcdnWafRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RuleConfigs) {
		body["RuleConfigs"] = request.RuleConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateDcdnWafRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateDcdnWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates Web Application Firewall (WAF) protection rules.
//
// Description:
//
// >  You can call this operation up to 20 times per second per account.
//
// @param request - BatchCreateDcdnWafRulesRequest
//
// @return BatchCreateDcdnWafRulesResponse
func (client *Client) BatchCreateDcdnWafRules(request *BatchCreateDcdnWafRulesRequest) (_result *BatchCreateDcdnWafRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCreateDcdnWafRulesResponse{}
	_body, _err := client.BatchCreateDcdnWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes configurations of multiple accelerated domain names at a time.
//
// Description:
//
// > - You can specify up to 50 domain names in each request.
//
// > - You can call this operation up to 30 times per second per account.
//
// @param request - BatchDeleteDcdnDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDcdnDomainConfigsResponse
func (client *Client) BatchDeleteDcdnDomainConfigsWithOptions(request *BatchDeleteDcdnDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDcdnDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

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
		Action:      dara.String("BatchDeleteDcdnDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDcdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes configurations of multiple accelerated domain names at a time.
//
// Description:
//
// > - You can specify up to 50 domain names in each request.
//
// > - You can call this operation up to 30 times per second per account.
//
// @param request - BatchDeleteDcdnDomainConfigsRequest
//
// @return BatchDeleteDcdnDomainConfigsResponse
func (client *Client) BatchDeleteDcdnDomainConfigs(request *BatchDeleteDcdnDomainConfigsRequest) (_result *BatchDeleteDcdnDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteDcdnDomainConfigsResponse{}
	_body, _err := client.BatchDeleteDcdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除kv数据，支持最大2M的请求体
//
// @param tmpReq - BatchDeleteDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDcdnKvResponse
func (client *Client) BatchDeleteDcdnKvWithOptions(tmpReq *BatchDeleteDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDcdnKvResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &BatchDeleteDcdnKvShrinkRequest{}
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
		Action:      dara.String("BatchDeleteDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDcdnKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除kv数据，支持最大2M的请求体
//
// @param request - BatchDeleteDcdnKvRequest
//
// @return BatchDeleteDcdnKvResponse
func (client *Client) BatchDeleteDcdnKv(request *BatchDeleteDcdnKvRequest) (_result *BatchDeleteDcdnKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteDcdnKvResponse{}
	_body, _err := client.BatchDeleteDcdnKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除kv数据，支持最大100M的请求体
//
// @param request - BatchDeleteDcdnKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDcdnKvWithHighCapacityResponse
func (client *Client) BatchDeleteDcdnKvWithHighCapacityWithOptions(request *BatchDeleteDcdnKvWithHighCapacityRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDcdnKvWithHighCapacityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("BatchDeleteDcdnKvWithHighCapacity"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDcdnKvWithHighCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除kv数据，支持最大100M的请求体
//
// @param request - BatchDeleteDcdnKvWithHighCapacityRequest
//
// @return BatchDeleteDcdnKvWithHighCapacityResponse
func (client *Client) BatchDeleteDcdnKvWithHighCapacity(request *BatchDeleteDcdnKvWithHighCapacityRequest) (_result *BatchDeleteDcdnKvWithHighCapacityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteDcdnKvWithHighCapacityResponse{}
	_body, _err := client.BatchDeleteDcdnKvWithHighCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteDcdnKvWithHighCapacityAdvance(request *BatchDeleteDcdnKvWithHighCapacityAdvanceRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDcdnKvWithHighCapacityResponse, _err error) {
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
		"Product":  dara.String("dcdn"),
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
	batchDeleteDcdnKvWithHighCapacityReq := &BatchDeleteDcdnKvWithHighCapacityRequest{}
	openapiutil.Convert(request, batchDeleteDcdnKvWithHighCapacityReq)
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
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		batchDeleteDcdnKvWithHighCapacityReq.Url = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	batchDeleteDcdnKvWithHighCapacityResp, _err := client.BatchDeleteDcdnKvWithHighCapacityWithOptions(batchDeleteDcdnKvWithHighCapacityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = batchDeleteDcdnKvWithHighCapacityResp
	return _result, _err
}

// Summary:
//
// Deletes multiple Web Application Firewall (WAF) protection rules at a time.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - BatchDeleteDcdnWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDcdnWafRulesResponse
func (client *Client) BatchDeleteDcdnWafRulesWithOptions(request *BatchDeleteDcdnWafRulesRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDcdnWafRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleIds) {
		body["RuleIds"] = request.RuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDcdnWafRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDcdnWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple Web Application Firewall (WAF) protection rules at a time.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - BatchDeleteDcdnWafRulesRequest
//
// @return BatchDeleteDcdnWafRulesResponse
func (client *Client) BatchDeleteDcdnWafRules(request *BatchDeleteDcdnWafRulesRequest) (_result *BatchDeleteDcdnWafRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteDcdnWafRulesResponse{}
	_body, _err := client.BatchDeleteDcdnWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies multiple Web Application Firewall (WAF) protection rules. Only Bot management rules can be modified.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - BatchModifyDcdnWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchModifyDcdnWafRulesResponse
func (client *Client) BatchModifyDcdnWafRulesWithOptions(request *BatchModifyDcdnWafRulesRequest, runtime *dara.RuntimeOptions) (_result *BatchModifyDcdnWafRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RuleConfigs) {
		body["RuleConfigs"] = request.RuleConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchModifyDcdnWafRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchModifyDcdnWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies multiple Web Application Firewall (WAF) protection rules. Only Bot management rules can be modified.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - BatchModifyDcdnWafRulesRequest
//
// @return BatchModifyDcdnWafRulesResponse
func (client *Client) BatchModifyDcdnWafRules(request *BatchModifyDcdnWafRulesRequest) (_result *BatchModifyDcdnWafRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchModifyDcdnWafRulesResponse{}
	_body, _err := client.BatchModifyDcdnWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures multiple key-value (KV) pairs for a namespace.
//
// @param tmpReq - BatchPutDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPutDcdnKvResponse
func (client *Client) BatchPutDcdnKvWithOptions(tmpReq *BatchPutDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *BatchPutDcdnKvResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &BatchPutDcdnKvShrinkRequest{}
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
		Action:      dara.String("BatchPutDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchPutDcdnKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures multiple key-value (KV) pairs for a namespace.
//
// @param request - BatchPutDcdnKvRequest
//
// @return BatchPutDcdnKvResponse
func (client *Client) BatchPutDcdnKv(request *BatchPutDcdnKvRequest) (_result *BatchPutDcdnKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchPutDcdnKvResponse{}
	_body, _err := client.BatchPutDcdnKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量写入kv数据，支持最大100M的请求体
//
// @param request - BatchPutDcdnKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchPutDcdnKvWithHighCapacityResponse
func (client *Client) BatchPutDcdnKvWithHighCapacityWithOptions(request *BatchPutDcdnKvWithHighCapacityRequest, runtime *dara.RuntimeOptions) (_result *BatchPutDcdnKvWithHighCapacityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("BatchPutDcdnKvWithHighCapacity"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchPutDcdnKvWithHighCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量写入kv数据，支持最大100M的请求体
//
// @param request - BatchPutDcdnKvWithHighCapacityRequest
//
// @return BatchPutDcdnKvWithHighCapacityResponse
func (client *Client) BatchPutDcdnKvWithHighCapacity(request *BatchPutDcdnKvWithHighCapacityRequest) (_result *BatchPutDcdnKvWithHighCapacityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchPutDcdnKvWithHighCapacityResponse{}
	_body, _err := client.BatchPutDcdnKvWithHighCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchPutDcdnKvWithHighCapacityAdvance(request *BatchPutDcdnKvWithHighCapacityAdvanceRequest, runtime *dara.RuntimeOptions) (_result *BatchPutDcdnKvWithHighCapacityResponse, _err error) {
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
		"Product":  dara.String("dcdn"),
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
	batchPutDcdnKvWithHighCapacityReq := &BatchPutDcdnKvWithHighCapacityRequest{}
	openapiutil.Convert(request, batchPutDcdnKvWithHighCapacityReq)
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
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		batchPutDcdnKvWithHighCapacityReq.Url = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	batchPutDcdnKvWithHighCapacityResp, _err := client.BatchPutDcdnKvWithHighCapacityWithOptions(batchPutDcdnKvWithHighCapacityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = batchPutDcdnKvWithHighCapacityResp
	return _result, _err
}

// Summary:
//
// Configures features for one or more domain names.
//
// Description:
//
//	  You can specify up to 50 domain names in each request.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - BatchSetDcdnDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetDcdnDomainConfigsResponse
func (client *Client) BatchSetDcdnDomainConfigsWithOptions(request *BatchSetDcdnDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetDcdnDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

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
		Action:      dara.String("BatchSetDcdnDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetDcdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures features for one or more domain names.
//
// Description:
//
//	  You can specify up to 50 domain names in each request.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - BatchSetDcdnDomainConfigsRequest
//
// @return BatchSetDcdnDomainConfigsResponse
func (client *Client) BatchSetDcdnDomainConfigs(request *BatchSetDcdnDomainConfigsRequest) (_result *BatchSetDcdnDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSetDcdnDomainConfigsResponse{}
	_body, _err := client.BatchSetDcdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures multiple domain names to be accelerated by IP Application Accelerator (IPA).
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - BatchSetDcdnIpaDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetDcdnIpaDomainConfigsResponse
func (client *Client) BatchSetDcdnIpaDomainConfigsWithOptions(request *BatchSetDcdnIpaDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetDcdnIpaDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

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
		Action:      dara.String("BatchSetDcdnIpaDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetDcdnIpaDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures multiple domain names to be accelerated by IP Application Accelerator (IPA).
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - BatchSetDcdnIpaDomainConfigsRequest
//
// @return BatchSetDcdnIpaDomainConfigsResponse
func (client *Client) BatchSetDcdnIpaDomainConfigs(request *BatchSetDcdnIpaDomainConfigsRequest) (_result *BatchSetDcdnIpaDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSetDcdnIpaDomainConfigsResponse{}
	_body, _err := client.BatchSetDcdnIpaDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the protection status of multiple domain names at a time.
//
// Description:
//
// #
//
//   - You can call this operation up to 20 times per second.
//
//   - Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - BatchSetDcdnWafDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetDcdnWafDomainConfigsResponse
func (client *Client) BatchSetDcdnWafDomainConfigsWithOptions(request *BatchSetDcdnWafDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *BatchSetDcdnWafDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientIpTag) {
		body["ClientIpTag"] = request.ClientIpTag
	}

	if !dara.IsNil(request.DefenseStatus) {
		body["DefenseStatus"] = request.DefenseStatus
	}

	if !dara.IsNil(request.DomainNames) {
		body["DomainNames"] = request.DomainNames
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetDcdnWafDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetDcdnWafDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the protection status of multiple domain names at a time.
//
// Description:
//
// #
//
//   - You can call this operation up to 20 times per second.
//
//   - Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - BatchSetDcdnWafDomainConfigsRequest
//
// @return BatchSetDcdnWafDomainConfigsResponse
func (client *Client) BatchSetDcdnWafDomainConfigs(request *BatchSetDcdnWafDomainConfigsRequest) (_result *BatchSetDcdnWafDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSetDcdnWafDomainConfigsResponse{}
	_body, _err := client.BatchSetDcdnWafDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables one or more accelerated domain names. After the accelerated domain names are enabled, the value of the DomainStatus parameter for the domain names changes to Online.
//
// Description:
//
// >
//
//   - If an accelerated domain name is in an invalid state or your account has an overdue payment, the accelerated domain name cannot be enabled.
//
//   - You can specify up to 50 domain names in each request.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - BatchStartDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartDcdnDomainResponse
func (client *Client) BatchStartDcdnDomainWithOptions(request *BatchStartDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStartDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

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
		Action:      dara.String("BatchStartDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStartDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables one or more accelerated domain names. After the accelerated domain names are enabled, the value of the DomainStatus parameter for the domain names changes to Online.
//
// Description:
//
// >
//
//   - If an accelerated domain name is in an invalid state or your account has an overdue payment, the accelerated domain name cannot be enabled.
//
//   - You can specify up to 50 domain names in each request.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - BatchStartDcdnDomainRequest
//
// @return BatchStartDcdnDomainResponse
func (client *Client) BatchStartDcdnDomain(request *BatchStartDcdnDomainRequest) (_result *BatchStartDcdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchStartDcdnDomainResponse{}
	_body, _err := client.BatchStartDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables one or more accelerated domain names. After an accelerated domain name is disabled, the value of the DomainStatus parameter for the domain name changes to Offline.
//
// Description:
//
// > 	- After an accelerated domain name is disabled, Dynamic Content Delivery Network (DCDN) retains the domain name information. The system automatically reroutes all requests that are destined for the accelerated domain name to the origin.
//
// >	- You can specify up to 50 domain names in each request.
//
// >	- You can call this operation up to 30 times per second per account.
//
// @param request - BatchStopDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopDcdnDomainResponse
func (client *Client) BatchStopDcdnDomainWithOptions(request *BatchStopDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStopDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

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
		Action:      dara.String("BatchStopDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStopDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables one or more accelerated domain names. After an accelerated domain name is disabled, the value of the DomainStatus parameter for the domain name changes to Offline.
//
// Description:
//
// > 	- After an accelerated domain name is disabled, Dynamic Content Delivery Network (DCDN) retains the domain name information. The system automatically reroutes all requests that are destined for the accelerated domain name to the origin.
//
// >	- You can specify up to 50 domain names in each request.
//
// >	- You can call this operation up to 30 times per second per account.
//
// @param request - BatchStopDcdnDomainRequest
//
// @return BatchStopDcdnDomainResponse
func (client *Client) BatchStopDcdnDomain(request *BatchStopDcdnDomainRequest) (_result *BatchStopDcdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchStopDcdnDomainResponse{}
	_body, _err := client.BatchStopDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a real-time log delivery project exists.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - CheckDcdnProjectExistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDcdnProjectExistResponse
func (client *Client) CheckDcdnProjectExistWithOptions(request *CheckDcdnProjectExistRequest, runtime *dara.RuntimeOptions) (_result *CheckDcdnProjectExistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDcdnProjectExist"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDcdnProjectExistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a real-time log delivery project exists.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - CheckDcdnProjectExistRequest
//
// @return CheckDcdnProjectExistResponse
func (client *Client) CheckDcdnProjectExist(request *CheckDcdnProjectExistRequest) (_result *CheckDcdnProjectExistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDcdnProjectExistResponse{}
	_body, _err := client.CheckDcdnProjectExistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates an official code version from unstable JavaScript code that is in the staging
//
//	environment. The version can be used in the canary release or production environment.
//
// Description:
//
// >  The call frequency of the API is no more than 100 queries per second.
//
// @param request - CommitStagingRoutineCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommitStagingRoutineCodeResponse
func (client *Client) CommitStagingRoutineCodeWithOptions(request *CommitStagingRoutineCodeRequest, runtime *dara.RuntimeOptions) (_result *CommitStagingRoutineCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("CommitStagingRoutineCode"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CommitStagingRoutineCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an official code version from unstable JavaScript code that is in the staging
//
//	environment. The version can be used in the canary release or production environment.
//
// Description:
//
// >  The call frequency of the API is no more than 100 queries per second.
//
// @param request - CommitStagingRoutineCodeRequest
//
// @return CommitStagingRoutineCodeResponse
func (client *Client) CommitStagingRoutineCode(request *CommitStagingRoutineCodeRequest) (_result *CommitStagingRoutineCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CommitStagingRoutineCodeResponse{}
	_body, _err := client.CommitStagingRoutineCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a certificate signing request (CSR) file.
//
// @param request - CreateDcdnCertificateSigningRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnCertificateSigningRequestResponse
func (client *Client) CreateDcdnCertificateSigningRequestWithOptions(request *CreateDcdnCertificateSigningRequestRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnCertificateSigningRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.CommonName) {
		query["CommonName"] = request.CommonName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.OrganizationUnit) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !dara.IsNil(request.SANs) {
		query["SANs"] = request.SANs
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnCertificateSigningRequest"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnCertificateSigningRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a certificate signing request (CSR) file.
//
// @param request - CreateDcdnCertificateSigningRequestRequest
//
// @return CreateDcdnCertificateSigningRequestResponse
func (client *Client) CreateDcdnCertificateSigningRequest(request *CreateDcdnCertificateSigningRequestRequest) (_result *CreateDcdnCertificateSigningRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDcdnCertificateSigningRequestResponse{}
	_body, _err := client.CreateDcdnCertificateSigningRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a tracking task. After you create a tracking task, the system periodically sends operations reports to you by email.
//
// Description:
//
// *
//
// **You can call this operation up to three times per second.
//
// @param request - CreateDcdnDeliverTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnDeliverTaskResponse
func (client *Client) CreateDcdnDeliverTaskWithOptions(request *CreateDcdnDeliverTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnDeliverTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Deliver) {
		body["Deliver"] = request.Deliver
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Reports) {
		body["Reports"] = request.Reports
	}

	if !dara.IsNil(request.Schedule) {
		body["Schedule"] = request.Schedule
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnDeliverTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnDeliverTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tracking task. After you create a tracking task, the system periodically sends operations reports to you by email.
//
// Description:
//
// *
//
// **You can call this operation up to three times per second.
//
// @param request - CreateDcdnDeliverTaskRequest
//
// @return CreateDcdnDeliverTaskResponse
func (client *Client) CreateDcdnDeliverTask(request *CreateDcdnDeliverTaskRequest) (_result *CreateDcdnDeliverTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDcdnDeliverTaskResponse{}
	_body, _err := client.CreateDcdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - CreateDcdnSLSRealTimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnSLSRealTimeLogDeliveryResponse
func (client *Client) CreateDcdnSLSRealTimeLogDeliveryWithOptions(request *CreateDcdnSLSRealTimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnSLSRealTimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.DataCenter) {
		body["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SLSLogStore) {
		body["SLSLogStore"] = request.SLSLogStore
	}

	if !dara.IsNil(request.SLSProject) {
		body["SLSProject"] = request.SLSProject
	}

	if !dara.IsNil(request.SLSRegion) {
		body["SLSRegion"] = request.SLSRegion
	}

	if !dara.IsNil(request.SamplingRate) {
		body["SamplingRate"] = request.SamplingRate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnSLSRealTimeLogDelivery"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnSLSRealTimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - CreateDcdnSLSRealTimeLogDeliveryRequest
//
// @return CreateDcdnSLSRealTimeLogDeliveryResponse
func (client *Client) CreateDcdnSLSRealTimeLogDelivery(request *CreateDcdnSLSRealTimeLogDeliveryRequest) (_result *CreateDcdnSLSRealTimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDcdnSLSRealTimeLogDeliveryResponse{}
	_body, _err := client.CreateDcdnSLSRealTimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom operations report.
//
// Description:
//
// > 	- This operation allows you to create a custom operations report for a specific domain name. You can view the statistics about the domain name in the report.
//
// > 	- You can call this operation up to three times per second per account.
//
// @param request - CreateDcdnSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnSubTaskResponse
func (client *Client) CreateDcdnSubTaskWithOptions(request *CreateDcdnSubTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnSubTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ReportIds) {
		body["ReportIds"] = request.ReportIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnSubTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom operations report.
//
// Description:
//
// > 	- This operation allows you to create a custom operations report for a specific domain name. You can view the statistics about the domain name in the report.
//
// > 	- You can call this operation up to three times per second per account.
//
// @param request - CreateDcdnSubTaskRequest
//
// @return CreateDcdnSubTaskResponse
func (client *Client) CreateDcdnSubTask(request *CreateDcdnSubTaskRequest) (_result *CreateDcdnSubTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDcdnSubTaskResponse{}
	_body, _err := client.CreateDcdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a custom WAF rule group.
//
// @param request - CreateDcdnWafGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnWafGroupResponse
func (client *Client) CreateDcdnWafGroupWithOptions(request *CreateDcdnWafGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnWafGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Subscribe) {
		body["Subscribe"] = request.Subscribe
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnWafGroup"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnWafGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a custom WAF rule group.
//
// @param request - CreateDcdnWafGroupRequest
//
// @return CreateDcdnWafGroupResponse
func (client *Client) CreateDcdnWafGroup(request *CreateDcdnWafGroupRequest) (_result *CreateDcdnWafGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDcdnWafGroupResponse{}
	_body, _err := client.CreateDcdnWafGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Web Application Firewall (WAF) protection policy.
//
// Description:
//
//	  You can call this operation up to 20 times per second per user.
//
//		- Alibaba Cloud Dynamic Route for CDN (DCDN) supports POST requests.
//
// @param request - CreateDcdnWafPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDcdnWafPolicyResponse
func (client *Client) CreateDcdnWafPolicyWithOptions(request *CreateDcdnWafPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateDcdnWafPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		body["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyStatus) {
		body["PolicyStatus"] = request.PolicyStatus
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDcdnWafPolicy"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDcdnWafPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Web Application Firewall (WAF) protection policy.
//
// Description:
//
//	  You can call this operation up to 20 times per second per user.
//
//		- Alibaba Cloud Dynamic Route for CDN (DCDN) supports POST requests.
//
// @param request - CreateDcdnWafPolicyRequest
//
// @return CreateDcdnWafPolicyResponse
func (client *Client) CreateDcdnWafPolicy(request *CreateDcdnWafPolicyRequest) (_result *CreateDcdnWafPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDcdnWafPolicyResponse{}
	_body, _err := client.CreateDcdnWafPolicyWithOptions(request, runtime)
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
// Description:
//
// > 	- The parameters must comply with the rules of EnvConf. The description of a routine cannot exceed 50 characters in length.
//
// >	- You can only specify the production and staging environments when you call this operation.
//
// >	- You can call this operation up to 100 times per second per account.
//
// @param tmpReq - CreateRoutineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoutineResponse
func (client *Client) CreateRoutineWithOptions(tmpReq *CreateRoutineRequest, runtime *dara.RuntimeOptions) (_result *CreateRoutineResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateRoutineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EnvConf) {
		request.EnvConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnvConf, dara.String("EnvConf"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvConfShrink) {
		body["EnvConf"] = request.EnvConfShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoutine"),
		Version:     dara.String("2018-01-15"),
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
// Description:
//
// > 	- The parameters must comply with the rules of EnvConf. The description of a routine cannot exceed 50 characters in length.
//
// >	- You can only specify the production and staging environments when you call this operation.
//
// >	- You can call this operation up to 100 times per second per account.
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
// Creates a service-linked role (SLR) and a Log Service project.
//
// Description:
//
// >  You can call this operation up to 100 times per second per account.
//
// @param request - CreateSlrAndSlsProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSlrAndSlsProjectResponse
func (client *Client) CreateSlrAndSlsProjectWithOptions(request *CreateSlrAndSlsProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateSlrAndSlsProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSlrAndSlsProject"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSlrAndSlsProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role (SLR) and a Log Service project.
//
// Description:
//
// >  You can call this operation up to 100 times per second per account.
//
// @param request - CreateSlrAndSlsProjectRequest
//
// @return CreateSlrAndSlsProjectResponse
func (client *Client) CreateSlrAndSlsProject(request *CreateSlrAndSlsProjectRequest) (_result *CreateSlrAndSlsProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSlrAndSlsProjectResponse{}
	_body, _err := client.CreateSlrAndSlsProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # A客户定制实时日志删除接口
//
// @param request - DeleteCustomDomainSampleRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomDomainSampleRateResponse
func (client *Client) DeleteCustomDomainSampleRateWithOptions(request *DeleteCustomDomainSampleRateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomDomainSampleRateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		body["DomainNames"] = request.DomainNames
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomDomainSampleRate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomDomainSampleRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A客户定制实时日志删除接口
//
// @param request - DeleteCustomDomainSampleRateRequest
//
// @return DeleteCustomDomainSampleRateResponse
func (client *Client) DeleteCustomDomainSampleRate(request *DeleteCustomDomainSampleRateRequest) (_result *DeleteCustomDomainSampleRateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomDomainSampleRateResponse{}
	_body, _err := client.DeleteCustomDomainSampleRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes tracking tasks by task ID.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 3.
//
// @param request - DeleteDcdnDeliverTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnDeliverTaskResponse
func (client *Client) DeleteDcdnDeliverTaskWithOptions(request *DeleteDcdnDeliverTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnDeliverTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliverId) {
		query["DeliverId"] = request.DeliverId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnDeliverTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnDeliverTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes tracking tasks by task ID.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 3.
//
// @param request - DeleteDcdnDeliverTaskRequest
//
// @return DeleteDcdnDeliverTaskResponse
func (client *Client) DeleteDcdnDeliverTask(request *DeleteDcdnDeliverTaskRequest) (_result *DeleteDcdnDeliverTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnDeliverTaskResponse{}
	_body, _err := client.DeleteDcdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified accelerated domain name.
//
// Description:
//
// > 	- Before you delete your domain name, you need to request the Domain Name System (DNS) provider to restore the A record of the domain name. Otherwise, the domain name may become inaccessible after you delete it.
//
// > 	- If you call the **DeleteDcdnDomain*	- operation, all the information about the accelerated domain name is deleted. If you want to disable an accelerated domain name, call the [StopDcdnDomain](https://help.aliyun.com/document_detail/130622.html) operation.
//
// > 	- You can call this operation up to 10 times per second per account.
//
// @param request - DeleteDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnDomainResponse
func (client *Client) DeleteDcdnDomainWithOptions(request *DeleteDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

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
		Action:      dara.String("DeleteDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified accelerated domain name.
//
// Description:
//
// > 	- Before you delete your domain name, you need to request the Domain Name System (DNS) provider to restore the A record of the domain name. Otherwise, the domain name may become inaccessible after you delete it.
//
// > 	- If you call the **DeleteDcdnDomain*	- operation, all the information about the accelerated domain name is deleted. If you want to disable an accelerated domain name, call the [StopDcdnDomain](https://help.aliyun.com/document_detail/130622.html) operation.
//
// > 	- You can call this operation up to 10 times per second per account.
//
// @param request - DeleteDcdnDomainRequest
//
// @return DeleteDcdnDomainResponse
func (client *Client) DeleteDcdnDomain(request *DeleteDcdnDomainRequest) (_result *DeleteDcdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnDomainResponse{}
	_body, _err := client.DeleteDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an accelerated domain name from IP Application Accelerator (IPA).
//
// Description:
//
// >
//
//   - Before you delete your domain name, we recommend that you request the Domain Name System (DNS) provider to restore the A record of the domain name. Otherwise, the domain name may become inaccessible after you delete it.
//
//   - This operation deletes all records of the specified accelerated domain name. If you want to temporarily disable an accelerated domain name, call the **StopDcdnIpaDomain*	- operation.****
//
//   - You can call this operation up to 10 times per second per account.
//
// @param request - DeleteDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnIpaDomainResponse
func (client *Client) DeleteDcdnIpaDomainWithOptions(request *DeleteDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

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
		Action:      dara.String("DeleteDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an accelerated domain name from IP Application Accelerator (IPA).
//
// Description:
//
// >
//
//   - Before you delete your domain name, we recommend that you request the Domain Name System (DNS) provider to restore the A record of the domain name. Otherwise, the domain name may become inaccessible after you delete it.
//
//   - This operation deletes all records of the specified accelerated domain name. If you want to temporarily disable an accelerated domain name, call the **StopDcdnIpaDomain*	- operation.****
//
//   - You can call this operation up to 10 times per second per account.
//
// @param request - DeleteDcdnIpaDomainRequest
//
// @return DeleteDcdnIpaDomainResponse
func (client *Client) DeleteDcdnIpaDomain(request *DeleteDcdnIpaDomainRequest) (_result *DeleteDcdnIpaDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnIpaDomainResponse{}
	_body, _err := client.DeleteDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes specific configurations of an accelerated domain name from IP Application Accelerator (IPA).
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DeleteDcdnIpaSpecificConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnIpaSpecificConfigResponse
func (client *Client) DeleteDcdnIpaSpecificConfigWithOptions(request *DeleteDcdnIpaSpecificConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnIpaSpecificConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("DeleteDcdnIpaSpecificConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnIpaSpecificConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes specific configurations of an accelerated domain name from IP Application Accelerator (IPA).
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DeleteDcdnIpaSpecificConfigRequest
//
// @return DeleteDcdnIpaSpecificConfigResponse
func (client *Client) DeleteDcdnIpaSpecificConfig(request *DeleteDcdnIpaSpecificConfigRequest) (_result *DeleteDcdnIpaSpecificConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnIpaSpecificConfigResponse{}
	_body, _err := client.DeleteDcdnIpaSpecificConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the key-value pairs in a namespace that you specify when you call the PutDcdnKvNamespace operation. EdgeKV provides a global key-value database for Dynamic Route for CDN (DCDN) points of presence (POPs).
//
// @param request - DeleteDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnKvResponse
func (client *Client) DeleteDcdnKvWithOptions(request *DeleteDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnKvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the key-value pairs in a namespace that you specify when you call the PutDcdnKvNamespace operation. EdgeKV provides a global key-value database for Dynamic Route for CDN (DCDN) points of presence (POPs).
//
// @param request - DeleteDcdnKvRequest
//
// @return DeleteDcdnKvResponse
func (client *Client) DeleteDcdnKv(request *DeleteDcdnKvRequest) (_result *DeleteDcdnKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnKvResponse{}
	_body, _err := client.DeleteDcdnKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a namespace that belongs to your account.
//
// @param request - DeleteDcdnKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnKvNamespaceResponse
func (client *Client) DeleteDcdnKvNamespaceWithOptions(request *DeleteDcdnKvNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnKvNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnKvNamespace"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnKvNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a namespace that belongs to your account.
//
// @param request - DeleteDcdnKvNamespaceRequest
//
// @return DeleteDcdnKvNamespaceResponse
func (client *Client) DeleteDcdnKvNamespace(request *DeleteDcdnKvNamespaceRequest) (_result *DeleteDcdnKvNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnKvNamespaceResponse{}
	_body, _err := client.DeleteDcdnKvNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// >  You can call this operation up to 100 times per second per account.
//
// @param request - DeleteDcdnRealTimeLogProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnRealTimeLogProjectResponse
func (client *Client) DeleteDcdnRealTimeLogProjectWithOptions(request *DeleteDcdnRealTimeLogProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnRealTimeLogProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnRealTimeLogProject"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnRealTimeLogProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// >  You can call this operation up to 100 times per second per account.
//
// @param request - DeleteDcdnRealTimeLogProjectRequest
//
// @return DeleteDcdnRealTimeLogProjectResponse
func (client *Client) DeleteDcdnRealTimeLogProject(request *DeleteDcdnRealTimeLogProjectRequest) (_result *DeleteDcdnRealTimeLogProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnRealTimeLogProjectResponse{}
	_body, _err := client.DeleteDcdnRealTimeLogProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes configurations of a domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DeleteDcdnSpecificConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnSpecificConfigResponse
func (client *Client) DeleteDcdnSpecificConfigWithOptions(request *DeleteDcdnSpecificConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnSpecificConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("DeleteDcdnSpecificConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnSpecificConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes configurations of a domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DeleteDcdnSpecificConfigRequest
//
// @return DeleteDcdnSpecificConfigResponse
func (client *Client) DeleteDcdnSpecificConfig(request *DeleteDcdnSpecificConfigRequest) (_result *DeleteDcdnSpecificConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnSpecificConfigResponse{}
	_body, _err := client.DeleteDcdnSpecificConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the configurations of an accelerated domain name in the canary release environment.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DeleteDcdnSpecificStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnSpecificStagingConfigResponse
func (client *Client) DeleteDcdnSpecificStagingConfigWithOptions(request *DeleteDcdnSpecificStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnSpecificStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("DeleteDcdnSpecificStagingConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnSpecificStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configurations of an accelerated domain name in the canary release environment.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DeleteDcdnSpecificStagingConfigRequest
//
// @return DeleteDcdnSpecificStagingConfigResponse
func (client *Client) DeleteDcdnSpecificStagingConfig(request *DeleteDcdnSpecificStagingConfigRequest) (_result *DeleteDcdnSpecificStagingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnSpecificStagingConfigResponse{}
	_body, _err := client.DeleteDcdnSpecificStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes all custom operations reports.
//
// Description:
//
// > You can call this operation up to 3 times per second per account.
//
// @param request - DeleteDcdnSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnSubTaskResponse
func (client *Client) DeleteDcdnSubTaskWithOptions(runtime *dara.RuntimeOptions) (_result *DeleteDcdnSubTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnSubTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes all custom operations reports.
//
// Description:
//
// > You can call this operation up to 3 times per second per account.
//
// @return DeleteDcdnSubTaskResponse
func (client *Client) DeleteDcdnSubTask() (_result *DeleteDcdnSubTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnSubTaskResponse{}
	_body, _err := client.DeleteDcdnSubTaskWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes feature configurations by user.
//
// @param request - DeleteDcdnUserConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnUserConfigResponse
func (client *Client) DeleteDcdnUserConfigWithOptions(request *DeleteDcdnUserConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnUserConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnUserConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes feature configurations by user.
//
// @param request - DeleteDcdnUserConfigRequest
//
// @return DeleteDcdnUserConfigResponse
func (client *Client) DeleteDcdnUserConfig(request *DeleteDcdnUserConfigRequest) (_result *DeleteDcdnUserConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnUserConfigResponse{}
	_body, _err := client.DeleteDcdnUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom WAF rule group.
//
// @param request - DeleteDcdnWafGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnWafGroupResponse
func (client *Client) DeleteDcdnWafGroupWithOptions(request *DeleteDcdnWafGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnWafGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnWafGroup"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnWafGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom WAF rule group.
//
// @param request - DeleteDcdnWafGroupRequest
//
// @return DeleteDcdnWafGroupResponse
func (client *Client) DeleteDcdnWafGroup(request *DeleteDcdnWafGroupRequest) (_result *DeleteDcdnWafGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnWafGroupResponse{}
	_body, _err := client.DeleteDcdnWafGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a protection policy.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - DeleteDcdnWafPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDcdnWafPolicyResponse
func (client *Client) DeleteDcdnWafPolicyWithOptions(request *DeleteDcdnWafPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteDcdnWafPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDcdnWafPolicy"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDcdnWafPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a protection policy.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - DeleteDcdnWafPolicyRequest
//
// @return DeleteDcdnWafPolicyResponse
func (client *Client) DeleteDcdnWafPolicy(request *DeleteDcdnWafPolicyRequest) (_result *DeleteDcdnWafPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDcdnWafPolicyResponse{}
	_body, _err := client.DeleteDcdnWafPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRoutineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineResponse
func (client *Client) DeleteRoutineWithOptions(request *DeleteRoutineRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2018-01-15"),
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
// Deletes a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
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
// Deletes the code of the specified version from a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRoutineCodeRevisionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineCodeRevisionResponse
func (client *Client) DeleteRoutineCodeRevisionWithOptions(request *DeleteRoutineCodeRevisionRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineCodeRevisionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SelectCodeRevision) {
		body["SelectCodeRevision"] = request.SelectCodeRevision
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoutineCodeRevision"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoutineCodeRevisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the code of the specified version from a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRoutineCodeRevisionRequest
//
// @return DeleteRoutineCodeRevisionResponse
func (client *Client) DeleteRoutineCodeRevision(request *DeleteRoutineCodeRevisionRequest) (_result *DeleteRoutineCodeRevisionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRoutineCodeRevisionResponse{}
	_body, _err := client.DeleteRoutineCodeRevisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes canary release environments from a routine.
//
// Description:
//
// >
//
//   - This operation deletes only custom preset canary release environments. You cannot delete production or staging environments.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param tmpReq - DeleteRoutineConfEnvsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoutineConfEnvsResponse
func (client *Client) DeleteRoutineConfEnvsWithOptions(tmpReq *DeleteRoutineConfEnvsRequest, runtime *dara.RuntimeOptions) (_result *DeleteRoutineConfEnvsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteRoutineConfEnvsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Envs) {
		request.EnvsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Envs, dara.String("Envs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvsShrink) {
		body["Envs"] = request.EnvsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoutineConfEnvs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRoutineConfEnvsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes canary release environments from a routine.
//
// Description:
//
// >
//
//   - This operation deletes only custom preset canary release environments. You cannot delete production or staging environments.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRoutineConfEnvsRequest
//
// @return DeleteRoutineConfEnvsResponse
func (client *Client) DeleteRoutineConfEnvs(request *DeleteRoutineConfEnvsRequest) (_result *DeleteRoutineConfEnvsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRoutineConfEnvsResponse{}
	_body, _err := client.DeleteRoutineConfEnvsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # A客户定制查询域名采样率
//
// @param request - DescribeCustomDomainSampleRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomDomainSampleRateResponse
func (client *Client) DescribeCustomDomainSampleRateWithOptions(request *DescribeCustomDomainSampleRateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomDomainSampleRateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
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
		Action:      dara.String("DescribeCustomDomainSampleRate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomDomainSampleRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A客户定制查询域名采样率
//
// @param request - DescribeCustomDomainSampleRateRequest
//
// @return DescribeCustomDomainSampleRateResponse
func (client *Client) DescribeCustomDomainSampleRate(request *DescribeCustomDomainSampleRateRequest) (_result *DescribeCustomDomainSampleRateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomDomainSampleRateResponse{}
	_body, _err := client.DescribeCustomDomainSampleRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries precise access control rules.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DescribeDcdnAclFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnAclFieldsResponse
func (client *Client) DescribeDcdnAclFieldsWithOptions(request *DescribeDcdnAclFieldsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnAclFieldsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnAclFields"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnAclFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries precise access control rules.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DescribeDcdnAclFieldsRequest
//
// @return DescribeDcdnAclFieldsResponse
func (client *Client) DescribeDcdnAclFields(request *DescribeDcdnAclFieldsRequest) (_result *DescribeDcdnAclFieldsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnAclFieldsResponse{}
	_body, _err := client.DescribeDcdnAclFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries bandwidth data for Border Gateway Protocol (BGP) accelerated domain names. Data is collected every 5 minutes.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range. You must set both parameters or leave both of them empty.
//
//		- If you specify multiple Internet service providers (ISPs), the data for the ISPs is aggregated.
//
//		- You can query data in the last 90 days.
//
//		- The maximum time range from the start time to the end time is 31 days. The start time is specified by the StartTime parameter and the end time is specified by the EndTime parameter.
//
//		- If the time range from the start time to the end time is 72 hours or shorter, you can specify the interval as 5 minutes. If the time range is longer than 72 hours, you must specify the interval as 1 hour.
//
//		- You can call this operation up to five times per second per account.
//
// @param request - DescribeDcdnBgpBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnBgpBpsDataResponse
func (client *Client) DescribeDcdnBgpBpsDataWithOptions(request *DescribeDcdnBgpBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnBgpBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.DevicePort) {
		query["DevicePort"] = request.DevicePort
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnBgpBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnBgpBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth data for Border Gateway Protocol (BGP) accelerated domain names. Data is collected every 5 minutes.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range. You must set both parameters or leave both of them empty.
//
//		- If you specify multiple Internet service providers (ISPs), the data for the ISPs is aggregated.
//
//		- You can query data in the last 90 days.
//
//		- The maximum time range from the start time to the end time is 31 days. The start time is specified by the StartTime parameter and the end time is specified by the EndTime parameter.
//
//		- If the time range from the start time to the end time is 72 hours or shorter, you can specify the interval as 5 minutes. If the time range is longer than 72 hours, you must specify the interval as 1 hour.
//
//		- You can call this operation up to five times per second per account.
//
// @param request - DescribeDcdnBgpBpsDataRequest
//
// @return DescribeDcdnBgpBpsDataResponse
func (client *Client) DescribeDcdnBgpBpsData(request *DescribeDcdnBgpBpsDataRequest) (_result *DescribeDcdnBgpBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnBgpBpsDataResponse{}
	_body, _err := client.DescribeDcdnBgpBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries traffic data for BGP accelerated domain names. Data is collected every 5 minutes.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range. You must set both parameters or leave both parameters empty.
//
//		- If you specify multiple Internet service providers (ISPs), the data for the ISPs is aggregated.
//
//		- You can query data in the last 90 days.
//
//		- The maximum time range that you can specify is 31 days. StartTime specifies the start time and EndTime specifies the end time of the time range.
//
//		- If the time range from the start time to the end time is 72 hours or shorter, you can specify the interval as 5 minutes. If the time range is longer than 72 hours, you must specify the interval as 1 hour.
//
//		- You can call this operation up to five times per second per account.
//
// @param request - DescribeDcdnBgpTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnBgpTrafficDataResponse
func (client *Client) DescribeDcdnBgpTrafficDataWithOptions(request *DescribeDcdnBgpTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnBgpTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnBgpTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnBgpTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traffic data for BGP accelerated domain names. Data is collected every 5 minutes.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range. You must set both parameters or leave both parameters empty.
//
//		- If you specify multiple Internet service providers (ISPs), the data for the ISPs is aggregated.
//
//		- You can query data in the last 90 days.
//
//		- The maximum time range that you can specify is 31 days. StartTime specifies the start time and EndTime specifies the end time of the time range.
//
//		- If the time range from the start time to the end time is 72 hours or shorter, you can specify the interval as 5 minutes. If the time range is longer than 72 hours, you must specify the interval as 1 hour.
//
//		- You can call this operation up to five times per second per account.
//
// @param request - DescribeDcdnBgpTrafficDataRequest
//
// @return DescribeDcdnBgpTrafficDataResponse
func (client *Client) DescribeDcdnBgpTrafficData(request *DescribeDcdnBgpTrafficDataRequest) (_result *DescribeDcdnBgpTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnBgpTrafficDataResponse{}
	_body, _err := client.DescribeDcdnBgpTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries countries and regions that can be added to the blacklist.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnBlockedRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnBlockedRegionsResponse
func (client *Client) DescribeDcdnBlockedRegionsWithOptions(request *DescribeDcdnBlockedRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnBlockedRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnBlockedRegions"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnBlockedRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries countries and regions that can be added to the blacklist.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnBlockedRegionsRequest
//
// @return DescribeDcdnBlockedRegionsResponse
func (client *Client) DescribeDcdnBlockedRegions(request *DescribeDcdnBlockedRegionsRequest) (_result *DescribeDcdnBlockedRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnBlockedRegionsResponse{}
	_body, _err := client.DescribeDcdnBlockedRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about a certificate.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnCertificateDetailResponse
func (client *Client) DescribeDcdnCertificateDetailWithOptions(request *DescribeDcdnCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnCertificateDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

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
		Action:      dara.String("DescribeDcdnCertificateDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about a certificate.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnCertificateDetailRequest
//
// @return DescribeDcdnCertificateDetailResponse
func (client *Client) DescribeDcdnCertificateDetail(request *DescribeDcdnCertificateDetailRequest) (_result *DescribeDcdnCertificateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnCertificateDetailResponse{}
	_body, _err := client.DescribeDcdnCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeDcdnCertificateList is deprecated, please use dcdn::2018-01-15::DescribeDcdnSSLCertificateList instead.
//
// Summary:
//
// Queries the certificates of one or more accelerated domain names.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnCertificateListResponse
func (client *Client) DescribeDcdnCertificateListWithOptions(request *DescribeDcdnCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("DescribeDcdnCertificateList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeDcdnCertificateList is deprecated, please use dcdn::2018-01-15::DescribeDcdnSSLCertificateList instead.
//
// Summary:
//
// Queries the certificates of one or more accelerated domain names.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnCertificateListRequest
//
// @return DescribeDcdnCertificateListResponse
// Deprecated
func (client *Client) DescribeDcdnCertificateList(request *DescribeDcdnCertificateListRequest) (_result *DescribeDcdnCertificateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnCertificateListResponse{}
	_body, _err := client.DescribeDcdnCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of DCDN DDoS mitigation.
//
// @param request - DescribeDcdnDdosServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDdosServiceResponse
func (client *Client) DescribeDcdnDdosServiceWithOptions(request *DescribeDcdnDdosServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDdosServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnDdosService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDdosServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of DCDN DDoS mitigation.
//
// @param request - DescribeDcdnDdosServiceRequest
//
// @return DescribeDcdnDdosServiceResponse
func (client *Client) DescribeDcdnDdosService(request *DescribeDcdnDdosServiceRequest) (_result *DescribeDcdnDdosServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDdosServiceResponse{}
	_body, _err := client.DescribeDcdnDdosServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the specifications of DCDN DDoS versions.
//
// @param request - DescribeDcdnDdosSpecInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDdosSpecInfoResponse
func (client *Client) DescribeDcdnDdosSpecInfoWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnDdosSpecInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDdosSpecInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDdosSpecInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specifications of DCDN DDoS versions.
//
// @return DescribeDcdnDdosSpecInfoResponse
func (client *Client) DescribeDcdnDdosSpecInfo() (_result *DescribeDcdnDdosSpecInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDdosSpecInfoResponse{}
	_body, _err := client.DescribeDcdnDdosSpecInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the domain names that are deleted from your Alibaba Cloud account.
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnDeletedDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDeletedDomainsResponse
func (client *Client) DescribeDcdnDeletedDomainsWithOptions(request *DescribeDcdnDeletedDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDeletedDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnDeletedDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDeletedDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names that are deleted from your Alibaba Cloud account.
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnDeletedDomainsRequest
//
// @return DescribeDcdnDeletedDomainsResponse
func (client *Client) DescribeDcdnDeletedDomains(request *DescribeDcdnDeletedDomainsRequest) (_result *DescribeDcdnDeletedDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDeletedDomainsResponse{}
	_body, _err := client.DescribeDcdnDeletedDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all tracking tasks of operations reports.
//
// Description:
//
// >You can call this operation up to three times per second.
//
// @param request - DescribeDcdnDeliverListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDeliverListResponse
func (client *Client) DescribeDcdnDeliverListWithOptions(request *DescribeDcdnDeliverListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDeliverListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliverId) {
		query["DeliverId"] = request.DeliverId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDeliverList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDeliverListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all tracking tasks of operations reports.
//
// Description:
//
// >You can call this operation up to three times per second.
//
// @param request - DescribeDcdnDeliverListRequest
//
// @return DescribeDcdnDeliverListResponse
func (client *Client) DescribeDcdnDeliverList(request *DescribeDcdnDeliverListRequest) (_result *DescribeDcdnDeliverListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDeliverListResponse{}
	_body, _err := client.DescribeDcdnDeliverListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of network bandwidth for one or more accelerated domain names. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainBpsDataResponse
func (client *Client) DescribeDcdnDomainBpsDataWithOptions(request *DescribeDcdnDomainBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of network bandwidth for one or more accelerated domain names. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainBpsDataRequest
//
// @return DescribeDcdnDomainBpsDataResponse
func (client *Client) DescribeDcdnDomainBpsData(request *DescribeDcdnDomainBpsDataRequest) (_result *DescribeDcdnDomainBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries bandwidth data of accelerated domain names.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set **StartTime*	- or **EndTime**, the request returns the data collected in the last 24 hours. If you set both **StartTime*	- and **EndTime**, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainBpsDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainBpsDataByLayerResponse
func (client *Client) DescribeDcdnDomainBpsDataByLayerWithOptions(request *DescribeDcdnDomainBpsDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainBpsDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainBpsDataByLayer"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainBpsDataByLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth data of accelerated domain names.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set **StartTime*	- or **EndTime**, the request returns the data collected in the last 24 hours. If you set both **StartTime*	- and **EndTime**, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainBpsDataByLayerRequest
//
// @return DescribeDcdnDomainBpsDataByLayerResponse
func (client *Client) DescribeDcdnDomainBpsDataByLayer(request *DescribeDcdnDomainBpsDataByLayerRequest) (_result *DescribeDcdnDomainBpsDataByLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainBpsDataByLayerResponse{}
	_body, _err := client.DescribeDcdnDomainBpsDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names by SSL certificate.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainByCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainByCertificateResponse
func (client *Client) DescribeDcdnDomainByCertificateWithOptions(request *DescribeDcdnDomainByCertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainByCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Exact) {
		query["Exact"] = request.Exact
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SSLStatus) {
		query["SSLStatus"] = request.SSLStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainByCertificate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainByCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names by SSL certificate.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainByCertificateRequest
//
// @return DescribeDcdnDomainByCertificateResponse
func (client *Client) DescribeDcdnDomainByCertificate(request *DescribeDcdnDomainByCertificateRequest) (_result *DescribeDcdnDomainByCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainByCertificateResponse{}
	_body, _err := client.DescribeDcdnDomainByCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries logs of rate limiting.
//
// Description:
//
// >
//
//   - If you do not configure the StartTime or EndTime parameter, data collected over the last 24 hours is queried. If you configure both the StartTime and EndTime parameters, data collected within the specified time range is queried.
//
//   - You can query data collected over the last 30 days.
//
//   - You can call the RefreshObjectCaches operation up to 50 times per second per account.
//
// @param request - DescribeDcdnDomainCcActivityLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainCcActivityLogResponse
func (client *Client) DescribeDcdnDomainCcActivityLogWithOptions(request *DescribeDcdnDomainCcActivityLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainCcActivityLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TriggerObject) {
		query["TriggerObject"] = request.TriggerObject
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainCcActivityLog"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainCcActivityLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries logs of rate limiting.
//
// Description:
//
// >
//
//   - If you do not configure the StartTime or EndTime parameter, data collected over the last 24 hours is queried. If you configure both the StartTime and EndTime parameters, data collected within the specified time range is queried.
//
//   - You can query data collected over the last 30 days.
//
//   - You can call the RefreshObjectCaches operation up to 50 times per second per account.
//
// @param request - DescribeDcdnDomainCcActivityLogRequest
//
// @return DescribeDcdnDomainCcActivityLogResponse
func (client *Client) DescribeDcdnDomainCcActivityLog(request *DescribeDcdnDomainCcActivityLogRequest) (_result *DescribeDcdnDomainCcActivityLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainCcActivityLogResponse{}
	_body, _err := client.DescribeDcdnDomainCcActivityLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificate information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainCertificateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainCertificateInfoResponse
func (client *Client) DescribeDcdnDomainCertificateInfoWithOptions(request *DescribeDcdnDomainCertificateInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainCertificateInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainCertificateInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainCertificateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainCertificateInfoRequest
//
// @return DescribeDcdnDomainCertificateInfoResponse
func (client *Client) DescribeDcdnDomainCertificateInfo(request *DescribeDcdnDomainCertificateInfoRequest) (_result *DescribeDcdnDomainCertificateInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainCertificateInfoResponse{}
	_body, _err := client.DescribeDcdnDomainCertificateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether CNAME records are configured for one or more accelerated domain names.
//
// Description:
//
// > You can call this operation up to 80 times per second per account.
//
// @param request - DescribeDcdnDomainCnameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainCnameResponse
func (client *Client) DescribeDcdnDomainCnameWithOptions(request *DescribeDcdnDomainCnameRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainCnameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainCname"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainCnameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether CNAME records are configured for one or more accelerated domain names.
//
// Description:
//
// > You can call this operation up to 80 times per second per account.
//
// @param request - DescribeDcdnDomainCnameRequest
//
// @return DescribeDcdnDomainCnameResponse
func (client *Client) DescribeDcdnDomainCname(request *DescribeDcdnDomainCnameRequest) (_result *DescribeDcdnDomainCnameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainCnameResponse{}
	_body, _err := client.DescribeDcdnDomainCnameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of an accelerated domain name.
//
// Description:
//
// > 	- You can query the configurations of one or more features in a request.
//
// > 	- You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainConfigsResponse
func (client *Client) DescribeDcdnDomainConfigsWithOptions(request *DescribeDcdnDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

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
		Action:      dara.String("DescribeDcdnDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an accelerated domain name.
//
// Description:
//
// > 	- You can query the configurations of one or more features in a request.
//
// > 	- You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnDomainConfigsRequest
//
// @return DescribeDcdnDomainConfigsResponse
func (client *Client) DescribeDcdnDomainConfigs(request *DescribeDcdnDomainConfigsRequest) (_result *DescribeDcdnDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainConfigsResponse{}
	_body, _err := client.DescribeDcdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic configuration information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainDetailResponse
func (client *Client) DescribeDcdnDomainDetailWithOptions(request *DescribeDcdnDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("DescribeDcdnDomainDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic configuration information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnDomainDetailRequest
//
// @return DescribeDcdnDomainDetailResponse
func (client *Client) DescribeDcdnDomainDetail(request *DescribeDcdnDomainDetailRequest) (_result *DescribeDcdnDomainDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainDetailResponse{}
	_body, _err := client.DescribeDcdnDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request hit ratios of one or more accelerated domain names. You can query data collected within the last 90 days.
//
// Description:
//
// #
//
//   - You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity*	- The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table. |Time granularity |Maximum time range per query |Historical data available |Data delay | -------------- | -------------- | ------ |5 minutes |3 days |93 days |15 minutes |1 hour |31 days |186 days |4 hours |1 day |366 days |366 days |04:00 on the next day
//
// @param request - DescribeDcdnDomainHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainHitRateDataResponse
func (client *Client) DescribeDcdnDomainHitRateDataWithOptions(request *DescribeDcdnDomainHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainHitRateData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request hit ratios of one or more accelerated domain names. You can query data collected within the last 90 days.
//
// Description:
//
// #
//
//   - You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity*	- The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table. |Time granularity |Maximum time range per query |Historical data available |Data delay | -------------- | -------------- | ------ |5 minutes |3 days |93 days |15 minutes |1 hour |31 days |186 days |4 hours |1 day |366 days |366 days |04:00 on the next day
//
// @param request - DescribeDcdnDomainHitRateDataRequest
//
// @return DescribeDcdnDomainHitRateDataResponse
func (client *Client) DescribeDcdnDomainHitRateData(request *DescribeDcdnDomainHitRateDataRequest) (_result *DescribeDcdnDomainHitRateDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainHitRateDataResponse{}
	_body, _err := client.DescribeDcdnDomainHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from one or more accelerated domain names. Data is collected every 5 minutes. You can query data in the last 90 days.
//
// Description:
//
// If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - You can call this operation up to 100 times per second per account.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainHttpCodeDataWithOptions(request *DescribeDcdnDomainHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainHttpCodeData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from one or more accelerated domain names. Data is collected every 5 minutes. You can query data in the last 90 days.
//
// Description:
//
// If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - You can call this operation up to 100 times per second per account.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainHttpCodeDataRequest
//
// @return DescribeDcdnDomainHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainHttpCodeData(request *DescribeDcdnDomainHttpCodeDataRequest) (_result *DescribeDcdnDomainHttpCodeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainHttpCodeDataResponse{}
	_body, _err := client.DescribeDcdnDomainHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the distribution of HTTP status codes by protocol.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - You cannot query the distribution of HTTP status codes by IP protocol.
//
//   - If you do not specify the **StartTime*	- or **EndTime*	- parameter, the data that is collected within the last 24 hours is collected. If you specify both the **StartTime*	- and **EndTime*	- parameters, the data that is collected within the time range that you specify is collected.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainHttpCodeDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainHttpCodeDataByLayerResponse
func (client *Client) DescribeDcdnDomainHttpCodeDataByLayerWithOptions(request *DescribeDcdnDomainHttpCodeDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainHttpCodeDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainHttpCodeDataByLayer"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainHttpCodeDataByLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution of HTTP status codes by protocol.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - You cannot query the distribution of HTTP status codes by IP protocol.
//
//   - If you do not specify the **StartTime*	- or **EndTime*	- parameter, the data that is collected within the last 24 hours is collected. If you specify both the **StartTime*	- and **EndTime*	- parameters, the data that is collected within the time range that you specify is collected.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainHttpCodeDataByLayerRequest
//
// @return DescribeDcdnDomainHttpCodeDataByLayerResponse
func (client *Client) DescribeDcdnDomainHttpCodeDataByLayer(request *DescribeDcdnDomainHttpCodeDataByLayerRequest) (_result *DescribeDcdnDomainHttpCodeDataByLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainHttpCodeDataByLayerResponse{}
	_body, _err := client.DescribeDcdnDomainHttpCodeDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries bandwidth of accelerated domain names for which Layer 4 acceleration is enabled. You can query the data that is collected over the last 90 days.
//
// Description:
//
// >
//
//   - The bandwidth is measured in bit/s.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainIpaBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainIpaBpsDataResponse
func (client *Client) DescribeDcdnDomainIpaBpsDataWithOptions(request *DescribeDcdnDomainIpaBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainIpaBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FixTimeGap) {
		query["FixTimeGap"] = request.FixTimeGap
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeMerge) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainIpaBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainIpaBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth of accelerated domain names for which Layer 4 acceleration is enabled. You can query the data that is collected over the last 90 days.
//
// Description:
//
// >
//
//   - The bandwidth is measured in bit/s.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainIpaBpsDataRequest
//
// @return DescribeDcdnDomainIpaBpsDataResponse
func (client *Client) DescribeDcdnDomainIpaBpsData(request *DescribeDcdnDomainIpaBpsDataRequest) (_result *DescribeDcdnDomainIpaBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainIpaBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainIpaBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of IPA user connections.
//
// Description:
//
//	  You can call this operation up to 10 times per second per user.
//
//		- If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//		- The minimum time granularity at which the data is queried is 5 minutes. The maximum time range for a single query is 31 days. The period within which historical data is available is 366 days. The data latency is no more than 10 minutes.
//
// @param request - DescribeDcdnDomainIpaConnDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainIpaConnDataResponse
func (client *Client) DescribeDcdnDomainIpaConnDataWithOptions(request *DescribeDcdnDomainIpaConnDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainIpaConnDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainIpaConnData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainIpaConnDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of IPA user connections.
//
// Description:
//
//	  You can call this operation up to 10 times per second per user.
//
//		- If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//		- The minimum time granularity at which the data is queried is 5 minutes. The maximum time range for a single query is 31 days. The period within which historical data is available is 366 days. The data latency is no more than 10 minutes.
//
// @param request - DescribeDcdnDomainIpaConnDataRequest
//
// @return DescribeDcdnDomainIpaConnDataResponse
func (client *Client) DescribeDcdnDomainIpaConnData(request *DescribeDcdnDomainIpaConnDataRequest) (_result *DescribeDcdnDomainIpaConnDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainIpaConnDataResponse{}
	_body, _err := client.DescribeDcdnDomainIpaConnDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries traffic of one or more accelerated domain names for which Layer 4 acceleration is enabled. You can query the data that is collected over the last 90 days.
//
// Description:
//
// >
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - Unit: bytes.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainIpaTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainIpaTrafficDataResponse
func (client *Client) DescribeDcdnDomainIpaTrafficDataWithOptions(request *DescribeDcdnDomainIpaTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainIpaTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FixTimeGap) {
		query["FixTimeGap"] = request.FixTimeGap
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeMerge) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainIpaTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainIpaTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traffic of one or more accelerated domain names for which Layer 4 acceleration is enabled. You can query the data that is collected over the last 90 days.
//
// Description:
//
// >
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - Unit: bytes.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainIpaTrafficDataRequest
//
// @return DescribeDcdnDomainIpaTrafficDataResponse
func (client *Client) DescribeDcdnDomainIpaTrafficData(request *DescribeDcdnDomainIpaTrafficDataRequest) (_result *DescribeDcdnDomainIpaTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainIpaTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainIpaTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the proportions of data usage of different Internet service providers (ISPs). You can query data within the last 90 days.
//
// Description:
//
// >
//
//   - You can call this operation up to 100 times per second per account.
//
//   - If **StartTime*	- is set but **EndTime*	- is not set, the data within the hour that starts from **StartTime*	- is queried.
//
//   - If **EndTime*	- is set but **StartTime*	- is not set, the data within the last hour that precedes **EndTime*	- is queried.
//
//   - You can query data of a domain name or all domain names that belong to your account.
//
//   - You can view data that is collected over the last seven days. The interval at which data is queried is based on the time range specified by **StartTime*	- and **EndTime**.
//
//   - **If the time range is shorter than or equal to one hour**, data is queried every minute.
//
//   - **If the time range is longer than 1 hour but shorter than or equal to three days**, data is queried every five minutes.
//
//   - **If the time range is longer than three days but shorter than or equal to seven days**, data is queried every hour.
//
// @param request - DescribeDcdnDomainIspDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainIspDataResponse
func (client *Client) DescribeDcdnDomainIspDataWithOptions(request *DescribeDcdnDomainIspDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainIspDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainIspData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainIspDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proportions of data usage of different Internet service providers (ISPs). You can query data within the last 90 days.
//
// Description:
//
// >
//
//   - You can call this operation up to 100 times per second per account.
//
//   - If **StartTime*	- is set but **EndTime*	- is not set, the data within the hour that starts from **StartTime*	- is queried.
//
//   - If **EndTime*	- is set but **StartTime*	- is not set, the data within the last hour that precedes **EndTime*	- is queried.
//
//   - You can query data of a domain name or all domain names that belong to your account.
//
//   - You can view data that is collected over the last seven days. The interval at which data is queried is based on the time range specified by **StartTime*	- and **EndTime**.
//
//   - **If the time range is shorter than or equal to one hour**, data is queried every minute.
//
//   - **If the time range is longer than 1 hour but shorter than or equal to three days**, data is queried every five minutes.
//
//   - **If the time range is longer than three days but shorter than or equal to seven days**, data is queried every hour.
//
// @param request - DescribeDcdnDomainIspDataRequest
//
// @return DescribeDcdnDomainIspDataResponse
func (client *Client) DescribeDcdnDomainIspData(request *DescribeDcdnDomainIspDataRequest) (_result *DescribeDcdnDomainIspDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainIspDataResponse{}
	_body, _err := client.DescribeDcdnDomainIspDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the address where you can download the log data of a domain name.
//
// Description:
//
// >
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.********
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainLogResponse
func (client *Client) DescribeDcdnDomainLogWithOptions(request *DescribeDcdnDomainLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainLog"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the address where you can download the log data of a domain name.
//
// Description:
//
// >
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.********
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainLogRequest
//
// @return DescribeDcdnDomainLogResponse
func (client *Client) DescribeDcdnDomainLog(request *DescribeDcdnDomainLogRequest) (_result *DescribeDcdnDomainLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainLogResponse{}
	_body, _err := client.DescribeDcdnDomainLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeDcdnDomainLogExTtlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainLogExTtlResponse
func (client *Client) DescribeDcdnDomainLogExTtlWithOptions(request *DescribeDcdnDomainLogExTtlRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainLogExTtlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainLogExTtl"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainLogExTtlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeDcdnDomainLogExTtlRequest
//
// @return DescribeDcdnDomainLogExTtlResponse
func (client *Client) DescribeDcdnDomainLogExTtl(request *DescribeDcdnDomainLogExTtlRequest) (_result *DescribeDcdnDomainLogExTtlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainLogExTtlResponse{}
	_body, _err := client.DescribeDcdnDomainLogExTtlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the billable items of accelerated domain names. The data is collected at least every 5 minutes. The billable items do not include non-request items.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, data within the last 10 minutes is queried. You can set both the StartTime and EndTime parameters to specify a time range.
//
//		- You can specify one or more accelerated domain names. Separate domain names with commas (,).
//
//		- You can query data within the last 90 days.
//
//		- The time range cannot exceed 1 hour.
//
// @param request - DescribeDcdnDomainMultiUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainMultiUsageDataResponse
func (client *Client) DescribeDcdnDomainMultiUsageDataWithOptions(request *DescribeDcdnDomainMultiUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainMultiUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainMultiUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainMultiUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the billable items of accelerated domain names. The data is collected at least every 5 minutes. The billable items do not include non-request items.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, data within the last 10 minutes is queried. You can set both the StartTime and EndTime parameters to specify a time range.
//
//		- You can specify one or more accelerated domain names. Separate domain names with commas (,).
//
//		- You can query data within the last 90 days.
//
//		- The time range cannot exceed 1 hour.
//
// @param request - DescribeDcdnDomainMultiUsageDataRequest
//
// @return DescribeDcdnDomainMultiUsageDataResponse
func (client *Client) DescribeDcdnDomainMultiUsageData(request *DescribeDcdnDomainMultiUsageDataRequest) (_result *DescribeDcdnDomainMultiUsageDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainMultiUsageDataResponse{}
	_body, _err := client.DescribeDcdnDomainMultiUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the origin bandwidth data for one or more accelerated domain names. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainOriginBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainOriginBpsDataResponse
func (client *Client) DescribeDcdnDomainOriginBpsDataWithOptions(request *DescribeDcdnDomainOriginBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainOriginBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainOriginBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainOriginBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the origin bandwidth data for one or more accelerated domain names. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainOriginBpsDataRequest
//
// @return DescribeDcdnDomainOriginBpsDataResponse
func (client *Client) DescribeDcdnDomainOriginBpsData(request *DescribeDcdnDomainOriginBpsDataRequest) (_result *DescribeDcdnDomainOriginBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainOriginBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainOriginBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin traffic of one or more accelerated domain names.
//
// Description:
//
// - You can call this operation up to 100 times per second per account.
//
// - If you do not set the **StartTime*	- or **EndTime*	- parameters, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter varies with the maximum time range per query. The following table describes the time period within which historical data is available and the data delay.
//
// | Time granularity | Maximum time range per query | Historical data available | Data delay |
//
// | ---------------- | ---------------------------- | ------------------------- | ---------- |
//
// | 5 minutes | 3 days | 93 days | 15 minutes |
//
// | 1 hour | 31 days | 186 days | 4 hours |
//
// | 1 day | 366 days | 366 days | 04:00 on the next day |
//
// @param request - DescribeDcdnDomainOriginTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainOriginTrafficDataResponse
func (client *Client) DescribeDcdnDomainOriginTrafficDataWithOptions(request *DescribeDcdnDomainOriginTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainOriginTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainOriginTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainOriginTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin traffic of one or more accelerated domain names.
//
// Description:
//
// - You can call this operation up to 100 times per second per account.
//
// - If you do not set the **StartTime*	- or **EndTime*	- parameters, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter varies with the maximum time range per query. The following table describes the time period within which historical data is available and the data delay.
//
// | Time granularity | Maximum time range per query | Historical data available | Data delay |
//
// | ---------------- | ---------------------------- | ------------------------- | ---------- |
//
// | 5 minutes | 3 days | 93 days | 15 minutes |
//
// | 1 hour | 31 days | 186 days | 4 hours |
//
// | 1 day | 366 days | 366 days | 04:00 on the next day |
//
// @param request - DescribeDcdnDomainOriginTrafficDataRequest
//
// @return DescribeDcdnDomainOriginTrafficDataResponse
func (client *Client) DescribeDcdnDomainOriginTrafficData(request *DescribeDcdnDomainOriginTrafficDataRequest) (_result *DescribeDcdnDomainOriginTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainOriginTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainOriginTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the protocol type of IP Application Accelerator (IPA).
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnDomainPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainPropertyResponse
func (client *Client) DescribeDcdnDomainPropertyWithOptions(request *DescribeDcdnDomainPropertyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainPropertyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainProperty"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the protocol type of IP Application Accelerator (IPA).
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnDomainPropertyRequest
//
// @return DescribeDcdnDomainPropertyResponse
func (client *Client) DescribeDcdnDomainProperty(request *DescribeDcdnDomainPropertyRequest) (_result *DescribeDcdnDomainPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainPropertyResponse{}
	_body, _err := client.DescribeDcdnDomainPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries page view (PV) data of an accelerated domain name. Data can be collected at minimum intervals of one hour.
//
// @param request - DescribeDcdnDomainPvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainPvDataResponse
func (client *Client) DescribeDcdnDomainPvDataWithOptions(request *DescribeDcdnDomainPvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainPvDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainPvData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainPvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries page view (PV) data of an accelerated domain name. Data can be collected at minimum intervals of one hour.
//
// @param request - DescribeDcdnDomainPvDataRequest
//
// @return DescribeDcdnDomainPvDataResponse
func (client *Client) DescribeDcdnDomainPvData(request *DescribeDcdnDomainPvDataRequest) (_result *DescribeDcdnDomainPvDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainPvDataResponse{}
	_body, _err := client.DescribeDcdnDomainPvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of requests to an accelerated domain name per second. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainQpsDataResponse
func (client *Client) DescribeDcdnDomainQpsDataWithOptions(request *DescribeDcdnDomainQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainQpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of requests to an accelerated domain name per second. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainQpsDataRequest
//
// @return DescribeDcdnDomainQpsDataResponse
func (client *Client) DescribeDcdnDomainQpsData(request *DescribeDcdnDomainQpsDataRequest) (_result *DescribeDcdnDomainQpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainQpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The number of queries per second in the Chinese mainland.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainQpsDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainQpsDataByLayerResponse
func (client *Client) DescribeDcdnDomainQpsDataByLayerWithOptions(request *DescribeDcdnDomainQpsDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainQpsDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainQpsDataByLayer"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainQpsDataByLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The number of queries per second in the Chinese mainland.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainQpsDataByLayerRequest
//
// @return DescribeDcdnDomainQpsDataByLayerResponse
func (client *Client) DescribeDcdnDomainQpsDataByLayer(request *DescribeDcdnDomainQpsDataByLayerRequest) (_result *DescribeDcdnDomainQpsDataByLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainQpsDataByLayerResponse{}
	_body, _err := client.DescribeDcdnDomainQpsDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time network bandwidth of a domain name.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify **StartTime*	- or **EndTime**, the request returns the data collected in the last hour by default. If you specify both parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|3 to 4 hours|
//
// @param request - DescribeDcdnDomainRealTimeBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeBpsDataResponse
func (client *Client) DescribeDcdnDomainRealTimeBpsDataWithOptions(request *DescribeDcdnDomainRealTimeBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time network bandwidth of a domain name.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify **StartTime*	- or **EndTime**, the request returns the data collected in the last hour by default. If you specify both parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|3 to 4 hours|
//
// @param request - DescribeDcdnDomainRealTimeBpsDataRequest
//
// @return DescribeDcdnDomainRealTimeBpsDataResponse
func (client *Client) DescribeDcdnDomainRealTimeBpsData(request *DescribeDcdnDomainRealTimeBpsDataRequest) (_result *DescribeDcdnDomainRealTimeBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries byte hit ratios at a time granularity of 1 minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the byte hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last hour. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// | Time granularity | Maximum time range per query | Historical data available | Data delay |
//
// |----|------|-----|--------|
//
// | 1 minute | 1 hour | 7 days | 5 minutes |
//
// | 5 minutes | 3 days | 93 days | 15 minutes |
//
// | 1 hour | 31 days | 186 days | 4 hours |
//
// @param request - DescribeDcdnDomainRealTimeByteHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeByteHitRateDataResponse
func (client *Client) DescribeDcdnDomainRealTimeByteHitRateDataWithOptions(request *DescribeDcdnDomainRealTimeByteHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeByteHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeByteHitRateData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries byte hit ratios at a time granularity of 1 minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the byte hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last hour. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// | Time granularity | Maximum time range per query | Historical data available | Data delay |
//
// |----|------|-----|--------|
//
// | 1 minute | 1 hour | 7 days | 5 minutes |
//
// | 5 minutes | 3 days | 93 days | 15 minutes |
//
// | 1 hour | 31 days | 186 days | 4 hours |
//
// @param request - DescribeDcdnDomainRealTimeByteHitRateDataRequest
//
// @return DescribeDcdnDomainRealTimeByteHitRateDataResponse
func (client *Client) DescribeDcdnDomainRealTimeByteHitRateData(request *DescribeDcdnDomainRealTimeByteHitRateDataRequest) (_result *DescribeDcdnDomainRealTimeByteHitRateDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeByteHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries traffic usage through each Internet service provider (ISP) and the number of visits in each region. The resolution of the data is one minute. The maximum time range to query for this operation is 10 minutes.
//
// Description:
//
// >
//
// > - You can call this operation up to 10 times per second per account.
//
// > - This operation is available only to users whose daily peak bandwidth value is higher than 1 Gbit/s. If you meet this requirement, you can [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl) to apply for permissions to use this operation.
//
// @param request - DescribeDcdnDomainRealTimeDetailDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeDetailDataResponse
func (client *Client) DescribeDcdnDomainRealTimeDetailDataWithOptions(request *DescribeDcdnDomainRealTimeDetailDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeDetailDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeDetailData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeDetailDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traffic usage through each Internet service provider (ISP) and the number of visits in each region. The resolution of the data is one minute. The maximum time range to query for this operation is 10 minutes.
//
// Description:
//
// >
//
// > - You can call this operation up to 10 times per second per account.
//
// > - This operation is available only to users whose daily peak bandwidth value is higher than 1 Gbit/s. If you meet this requirement, you can [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl) to apply for permissions to use this operation.
//
// @param request - DescribeDcdnDomainRealTimeDetailDataRequest
//
// @return DescribeDcdnDomainRealTimeDetailDataResponse
func (client *Client) DescribeDcdnDomainRealTimeDetailData(request *DescribeDcdnDomainRealTimeDetailDataRequest) (_result *DescribeDcdnDomainRealTimeDetailDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeDetailDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeDetailDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainRealTimeHttpCodeDataWithOptions(request *DescribeDcdnDomainRealTimeHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeHttpCodeData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeHttpCodeDataRequest
//
// @return DescribeDcdnDomainRealTimeHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainRealTimeHttpCodeData(request *DescribeDcdnDomainRealTimeHttpCodeDataRequest) (_result *DescribeDcdnDomainRealTimeHttpCodeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The number of QPS for one or more accelerated domain names is queried. Data is collected every minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeQpsDataResponse
func (client *Client) DescribeDcdnDomainRealTimeQpsDataWithOptions(request *DescribeDcdnDomainRealTimeQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeQpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The number of QPS for one or more accelerated domain names is queried. Data is collected every minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeQpsDataRequest
//
// @return DescribeDcdnDomainRealTimeQpsDataResponse
func (client *Client) DescribeDcdnDomainRealTimeQpsData(request *DescribeDcdnDomainRealTimeQpsDataRequest) (_result *DescribeDcdnDomainRealTimeQpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeQpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request hit rates with a time granularity of 1 minute.
//
// Description:
//
// You can call this operation up to 10 times per second per user.
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the byte hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last hour. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity*	- The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeReqHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeReqHitRateDataResponse
func (client *Client) DescribeDcdnDomainRealTimeReqHitRateDataWithOptions(request *DescribeDcdnDomainRealTimeReqHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeReqHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeReqHitRateData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request hit rates with a time granularity of 1 minute.
//
// Description:
//
// You can call this operation up to 10 times per second per user.
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the byte hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last hour. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity*	- The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeReqHitRateDataRequest
//
// @return DescribeDcdnDomainRealTimeReqHitRateDataResponse
func (client *Client) DescribeDcdnDomainRealTimeReqHitRateData(request *DescribeDcdnDomainRealTimeReqHitRateDataRequest) (_result *DescribeDcdnDomainRealTimeReqHitRateDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeReqHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data of back-to-origin requests. Data is collected every minute. You can query data collected in the last 7 days.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// | Time granularity | Maximum time range per query | Historical data available | Data delay |
//
// |-----|-----|-----|--------|
//
// | 1 minute | 1 hour | 7 days | 5 minutes |
//
// | 5 minutes | 3 days | 93 days | 15 minutes | | 1 hour | 31 days | 186 days | 4 hours |
//
// @param request - DescribeDcdnDomainRealTimeSrcBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeSrcBpsDataResponse
func (client *Client) DescribeDcdnDomainRealTimeSrcBpsDataWithOptions(request *DescribeDcdnDomainRealTimeSrcBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeSrcBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeSrcBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data of back-to-origin requests. Data is collected every minute. You can query data collected in the last 7 days.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// | Time granularity | Maximum time range per query | Historical data available | Data delay |
//
// |-----|-----|-----|--------|
//
// | 1 minute | 1 hour | 7 days | 5 minutes |
//
// | 5 minutes | 3 days | 93 days | 15 minutes | | 1 hour | 31 days | 186 days | 4 hours |
//
// @param request - DescribeDcdnDomainRealTimeSrcBpsDataRequest
//
// @return DescribeDcdnDomainRealTimeSrcBpsDataResponse
func (client *Client) DescribeDcdnDomainRealTimeSrcBpsData(request *DescribeDcdnDomainRealTimeSrcBpsDataRequest) (_result *DescribeDcdnDomainRealTimeSrcBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeSrcBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the proportions of HTTP status codes based on back-to-origin statistics with a time granularity of one minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainRealTimeSrcHttpCodeDataWithOptions(request *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeSrcHttpCodeData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proportions of HTTP status codes based on back-to-origin statistics with a time granularity of one minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest
//
// @return DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainRealTimeSrcHttpCodeData(request *DescribeDcdnDomainRealTimeSrcHttpCodeDataRequest) (_result *DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeSrcHttpCodeDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeSrcHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the origin traffic monitoring data for an accelerated domain name. Data is collected every minute. You can query data in the last 90 days.
//
// Description:
//
// If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeSrcTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeSrcTrafficDataResponse
func (client *Client) DescribeDcdnDomainRealTimeSrcTrafficDataWithOptions(request *DescribeDcdnDomainRealTimeSrcTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeSrcTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeSrcTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the origin traffic monitoring data for an accelerated domain name. Data is collected every minute. You can query data in the last 90 days.
//
// Description:
//
// If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeSrcTrafficDataRequest
//
// @return DescribeDcdnDomainRealTimeSrcTrafficDataResponse
func (client *Client) DescribeDcdnDomainRealTimeSrcTrafficData(request *DescribeDcdnDomainRealTimeSrcTrafficDataRequest) (_result *DescribeDcdnDomainRealTimeSrcTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeSrcTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the traffic monitoring data of an accelerated domain name. Data is collected every minute.
//
// Description:
//
// You can call this operation up to 50 times per second per user.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRealTimeTrafficDataResponse
func (client *Client) DescribeDcdnDomainRealTimeTrafficDataWithOptions(request *DescribeDcdnDomainRealTimeTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRealTimeTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRealTimeTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRealTimeTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic monitoring data of an accelerated domain name. Data is collected every minute.
//
// Description:
//
// You can call this operation up to 50 times per second per user.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDcdnDomainRealTimeTrafficDataRequest
//
// @return DescribeDcdnDomainRealTimeTrafficDataResponse
func (client *Client) DescribeDcdnDomainRealTimeTrafficData(request *DescribeDcdnDomainRealTimeTrafficDataRequest) (_result *DescribeDcdnDomainRealTimeTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRealTimeTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainRealTimeTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries regional distribution of users. Data is collected every day. You can query data within the last 90 days.
//
// Description:
//
// >
//
//   - If you do not specify the StartTime and EndTime parameters, the data within the last 24 hours is queried. If you specify the StartTime and EndTime parameters, the data within the specified time range is queried.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainRegionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainRegionDataResponse
func (client *Client) DescribeDcdnDomainRegionDataWithOptions(request *DescribeDcdnDomainRegionDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainRegionDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainRegionData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainRegionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regional distribution of users. Data is collected every day. You can query data within the last 90 days.
//
// Description:
//
// >
//
//   - If you do not specify the StartTime and EndTime parameters, the data within the last 24 hours is queried. If you specify the StartTime and EndTime parameters, the data within the specified time range is queried.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnDomainRegionDataRequest
//
// @return DescribeDcdnDomainRegionDataResponse
func (client *Client) DescribeDcdnDomainRegionData(request *DescribeDcdnDomainRegionDataRequest) (_result *DescribeDcdnDomainRegionDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainRegionDataResponse{}
	_body, _err := client.DescribeDcdnDomainRegionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the environment configuration in the canary release environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnDomainStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainStagingConfigResponse
func (client *Client) DescribeDcdnDomainStagingConfigWithOptions(request *DescribeDcdnDomainStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainStagingConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the environment configuration in the canary release environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnDomainStagingConfigRequest
//
// @return DescribeDcdnDomainStagingConfigResponse
func (client *Client) DescribeDcdnDomainStagingConfig(request *DescribeDcdnDomainStagingConfigRequest) (_result *DescribeDcdnDomainStagingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainStagingConfigResponse{}
	_body, _err := client.DescribeDcdnDomainStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries and sorts frequently requested web pages on a specified day. You can query data collected within the last 90 days.
//
// Description:
//
//	  If you do not set the StartTime parameter, the data on the previous day is queried.
//
//		- You can specify only one domain name.
//
// @param request - DescribeDcdnDomainTopReferVisitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainTopReferVisitResponse
func (client *Client) DescribeDcdnDomainTopReferVisitWithOptions(request *DescribeDcdnDomainTopReferVisitRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainTopReferVisitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainTopReferVisit"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainTopReferVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries and sorts frequently requested web pages on a specified day. You can query data collected within the last 90 days.
//
// Description:
//
//	  If you do not set the StartTime parameter, the data on the previous day is queried.
//
//		- You can specify only one domain name.
//
// @param request - DescribeDcdnDomainTopReferVisitRequest
//
// @return DescribeDcdnDomainTopReferVisitResponse
func (client *Client) DescribeDcdnDomainTopReferVisit(request *DescribeDcdnDomainTopReferVisitRequest) (_result *DescribeDcdnDomainTopReferVisitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainTopReferVisitResponse{}
	_body, _err := client.DescribeDcdnDomainTopReferVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries frequently requested URLs on a day.
//
// Description:
//
// > You can query data in the last seven days.
//
// @param request - DescribeDcdnDomainTopUrlVisitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainTopUrlVisitResponse
func (client *Client) DescribeDcdnDomainTopUrlVisitWithOptions(request *DescribeDcdnDomainTopUrlVisitRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainTopUrlVisitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainTopUrlVisit"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainTopUrlVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries frequently requested URLs on a day.
//
// Description:
//
// > You can query data in the last seven days.
//
// @param request - DescribeDcdnDomainTopUrlVisitRequest
//
// @return DescribeDcdnDomainTopUrlVisitResponse
func (client *Client) DescribeDcdnDomainTopUrlVisit(request *DescribeDcdnDomainTopUrlVisitRequest) (_result *DescribeDcdnDomainTopUrlVisitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainTopUrlVisitResponse{}
	_body, _err := client.DescribeDcdnDomainTopUrlVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network traffic of accelerated domain names. You can query data collected in the last 90 days.
//
// Description:
//
// If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - You can call this operation up to 100 times per second per account.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainTrafficDataResponse
func (client *Client) DescribeDcdnDomainTrafficDataWithOptions(request *DescribeDcdnDomainTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network traffic of accelerated domain names. You can query data collected in the last 90 days.
//
// Description:
//
// If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
//   - You can call this operation up to 100 times per second per account.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainTrafficDataRequest
//
// @return DescribeDcdnDomainTrafficDataResponse
func (client *Client) DescribeDcdnDomainTrafficData(request *DescribeDcdnDomainTrafficDataRequest) (_result *DescribeDcdnDomainTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resource usage about domain names in a billable region.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - Usage data includes traffic (measured in bytes), bandwidth values (measured in bit/s), and the number of requests.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainUsageDataResponse
func (client *Client) DescribeDcdnDomainUsageDataWithOptions(request *DescribeDcdnDomainUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.DataProtocol) {
		query["DataProtocol"] = request.DataProtocol
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Field) {
		query["Field"] = request.Field
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource usage about domain names in a billable region.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - Usage data includes traffic (measured in bytes), bandwidth values (measured in bit/s), and the number of requests.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainUsageDataRequest
//
// @return DescribeDcdnDomainUsageDataResponse
func (client *Client) DescribeDcdnDomainUsageData(request *DescribeDcdnDomainUsageDataRequest) (_result *DescribeDcdnDomainUsageDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainUsageDataResponse{}
	_body, _err := client.DescribeDcdnDomainUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of unique visitors (UVs) to an accelerated domain name. Data is collected every hour. You can query data within the last 90 days.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//		- You can specify only one accelerated domain name or all the accelerated domain names that belong to your Alibaba Cloud account.
//
// @param request - DescribeDcdnDomainUvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainUvDataResponse
func (client *Client) DescribeDcdnDomainUvDataWithOptions(request *DescribeDcdnDomainUvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainUvDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainUvData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainUvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of unique visitors (UVs) to an accelerated domain name. Data is collected every hour. You can query data within the last 90 days.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//		- You can specify only one accelerated domain name or all the accelerated domain names that belong to your Alibaba Cloud account.
//
// @param request - DescribeDcdnDomainUvDataRequest
//
// @return DescribeDcdnDomainUvDataResponse
func (client *Client) DescribeDcdnDomainUvData(request *DescribeDcdnDomainUvDataRequest) (_result *DescribeDcdnDomainUvDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainUvDataResponse{}
	_body, _err := client.DescribeDcdnDomainUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries bandwidth of one or more accelerated domain names for which WebSocket is enabled. You can query the data that is collected over the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainWebsocketBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainWebsocketBpsDataResponse
func (client *Client) DescribeDcdnDomainWebsocketBpsDataWithOptions(request *DescribeDcdnDomainWebsocketBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainWebsocketBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainWebsocketBpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainWebsocketBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth of one or more accelerated domain names for which WebSocket is enabled. You can query the data that is collected over the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainWebsocketBpsDataRequest
//
// @return DescribeDcdnDomainWebsocketBpsDataResponse
func (client *Client) DescribeDcdnDomainWebsocketBpsData(request *DescribeDcdnDomainWebsocketBpsDataRequest) (_result *DescribeDcdnDomainWebsocketBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainWebsocketBpsDataResponse{}
	_body, _err := client.DescribeDcdnDomainWebsocketBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The total number and proportions of HTTP status codes returned from one or more accelerated domain names for which WebSocket is enabled are queried. Data can be collected at minimum intervals of 5 minutes.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainWebsocketHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainWebsocketHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainWebsocketHttpCodeDataWithOptions(request *DescribeDcdnDomainWebsocketHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainWebsocketHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainWebsocketHttpCodeData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainWebsocketHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The total number and proportions of HTTP status codes returned from one or more accelerated domain names for which WebSocket is enabled are queried. Data can be collected at minimum intervals of 5 minutes.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainWebsocketHttpCodeDataRequest
//
// @return DescribeDcdnDomainWebsocketHttpCodeDataResponse
func (client *Client) DescribeDcdnDomainWebsocketHttpCodeData(request *DescribeDcdnDomainWebsocketHttpCodeDataRequest) (_result *DescribeDcdnDomainWebsocketHttpCodeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainWebsocketHttpCodeDataResponse{}
	_body, _err := client.DescribeDcdnDomainWebsocketHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the traffic monitoring data for an accelerated domain name with WebSocket enabled. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainWebsocketTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainWebsocketTrafficDataResponse
func (client *Client) DescribeDcdnDomainWebsocketTrafficDataWithOptions(request *DescribeDcdnDomainWebsocketTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainWebsocketTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainWebsocketTrafficData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainWebsocketTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic monitoring data for an accelerated domain name with WebSocket enabled. You can query data in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the last 24 hours. If you set both the **StartTime*	- and **EndTime*	- parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDcdnDomainWebsocketTrafficDataRequest
//
// @return DescribeDcdnDomainWebsocketTrafficDataResponse
func (client *Client) DescribeDcdnDomainWebsocketTrafficData(request *DescribeDcdnDomainWebsocketTrafficDataRequest) (_result *DescribeDcdnDomainWebsocketTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainWebsocketTrafficDataResponse{}
	_body, _err := client.DescribeDcdnDomainWebsocketTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries DCDN-accelerated domain names by origin server.
//
// @param request - DescribeDcdnDomainsBySourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnDomainsBySourceResponse
func (client *Client) DescribeDcdnDomainsBySourceWithOptions(request *DescribeDcdnDomainsBySourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnDomainsBySourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnDomainsBySource"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnDomainsBySourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries DCDN-accelerated domain names by origin server.
//
// @param request - DescribeDcdnDomainsBySourceRequest
//
// @return DescribeDcdnDomainsBySourceResponse
func (client *Client) DescribeDcdnDomainsBySource(request *DescribeDcdnDomainsBySourceRequest) (_result *DescribeDcdnDomainsBySourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnDomainsBySourceResponse{}
	_body, _err := client.DescribeDcdnDomainsBySourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of times that a routine is executed within a specified period of time.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- The minimum time granularity for a query is 1 hour. The maximum time span for a query is 24 hours. The time period within which historical data is available for a query is 366 days.
//
// @param request - DescribeDcdnErUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnErUsageDataResponse
func (client *Client) DescribeDcdnErUsageDataWithOptions(request *DescribeDcdnErUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnErUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RoutineID) {
		query["RoutineID"] = request.RoutineID
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnErUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnErUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of times that a routine is executed within a specified period of time.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- The minimum time granularity for a query is 1 hour. The maximum time span for a query is 24 hours. The time period within which historical data is available for a query is 366 days.
//
// @param request - DescribeDcdnErUsageDataRequest
//
// @return DescribeDcdnErUsageDataResponse
func (client *Client) DescribeDcdnErUsageData(request *DescribeDcdnErUsageDataRequest) (_result *DescribeDcdnErUsageDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnErUsageDataResponse{}
	_body, _err := client.DescribeDcdnErUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of blocked IP addresses.
//
// Description:
//
// > 	- If you specify IP addresses or CIDR blocks, IP addresses that are effective and the corresponding expiration time are returned. If you do not specify IP addresses or CIDR blocks, all effective IP addresses and the corresponding expiration time are returned.
//
// > 	- The results are written to OSS and returned as OSS URLs. The content in OSS objects is in the format of `IP address-Corresponding expiration time`. The expiration time is in the YYYY-MM-DD hh:mm:ss format.
//
// > 	- You can share OSS URLs with others. The shared URLs are valid for three days.
//
// @param request - DescribeDcdnFullDomainsBlockIPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnFullDomainsBlockIPConfigResponse
func (client *Client) DescribeDcdnFullDomainsBlockIPConfigWithOptions(request *DescribeDcdnFullDomainsBlockIPConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnFullDomainsBlockIPConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnFullDomainsBlockIPConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnFullDomainsBlockIPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of blocked IP addresses.
//
// Description:
//
// > 	- If you specify IP addresses or CIDR blocks, IP addresses that are effective and the corresponding expiration time are returned. If you do not specify IP addresses or CIDR blocks, all effective IP addresses and the corresponding expiration time are returned.
//
// > 	- The results are written to OSS and returned as OSS URLs. The content in OSS objects is in the format of `IP address-Corresponding expiration time`. The expiration time is in the YYYY-MM-DD hh:mm:ss format.
//
// > 	- You can share OSS URLs with others. The shared URLs are valid for three days.
//
// @param request - DescribeDcdnFullDomainsBlockIPConfigRequest
//
// @return DescribeDcdnFullDomainsBlockIPConfigResponse
func (client *Client) DescribeDcdnFullDomainsBlockIPConfig(request *DescribeDcdnFullDomainsBlockIPConfigRequest) (_result *DescribeDcdnFullDomainsBlockIPConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnFullDomainsBlockIPConfigResponse{}
	_body, _err := client.DescribeDcdnFullDomainsBlockIPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户海量封禁历史
//
// Description:
//
//	  For a specified IP addresses and time range, the time when the IP address was delivered to the edge and the corresponding result are returned.
//
//		- If a specified IP address or CIDR block has multiple blocking records in a specified time range, the records are sorted by delivery time in descending order.
//
//		- The maximum time range to query is 90 days.
//
//		- If no blocking record exists or delivery fails for the given IP address and time range, the delivery time is empty.
//
// @param request - DescribeDcdnFullDomainsBlockIPHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnFullDomainsBlockIPHistoryResponse
func (client *Client) DescribeDcdnFullDomainsBlockIPHistoryWithOptions(request *DescribeDcdnFullDomainsBlockIPHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnFullDomainsBlockIPHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IPList) {
		body["IPList"] = request.IPList
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnFullDomainsBlockIPHistory"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnFullDomainsBlockIPHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户海量封禁历史
//
// Description:
//
//	  For a specified IP addresses and time range, the time when the IP address was delivered to the edge and the corresponding result are returned.
//
//		- If a specified IP address or CIDR block has multiple blocking records in a specified time range, the records are sorted by delivery time in descending order.
//
//		- The maximum time range to query is 90 days.
//
//		- If no blocking record exists or delivery fails for the given IP address and time range, the delivery time is empty.
//
// @param request - DescribeDcdnFullDomainsBlockIPHistoryRequest
//
// @return DescribeDcdnFullDomainsBlockIPHistoryResponse
func (client *Client) DescribeDcdnFullDomainsBlockIPHistory(request *DescribeDcdnFullDomainsBlockIPHistoryRequest) (_result *DescribeDcdnFullDomainsBlockIPHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnFullDomainsBlockIPHistoryResponse{}
	_body, _err := client.DescribeDcdnFullDomainsBlockIPHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all certificates that belong to your account.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnHttpsDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnHttpsDomainListResponse
func (client *Client) DescribeDcdnHttpsDomainListWithOptions(request *DescribeDcdnHttpsDomainListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnHttpsDomainListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
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
		Action:      dara.String("DescribeDcdnHttpsDomainList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnHttpsDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all certificates that belong to your account.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnHttpsDomainListRequest
//
// @return DescribeDcdnHttpsDomainListResponse
func (client *Client) DescribeDcdnHttpsDomainList(request *DescribeDcdnHttpsDomainListRequest) (_result *DescribeDcdnHttpsDomainListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnHttpsDomainListResponse{}
	_body, _err := client.DescribeDcdnHttpsDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether an IP address belongs to a POP.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnIpInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpInfoResponse
func (client *Client) DescribeDcdnIpInfoWithOptions(request *DescribeDcdnIpInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnIpInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether an IP address belongs to a POP.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnIpInfoRequest
//
// @return DescribeDcdnIpInfoResponse
func (client *Client) DescribeDcdnIpInfo(request *DescribeDcdnIpInfoRequest) (_result *DescribeDcdnIpInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnIpInfoResponse{}
	_body, _err := client.DescribeDcdnIpInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin CIDR blocks of IPA-accelerated domain names. If you want to call this API operation, you must submit a ticket to apply for the required permissions.
//
// Description:
//
// >  This operation can be called globally up to 50 times per second. This operation can be called up to 10 times per second per account.
//
// @param request - DescribeDcdnIpaDomainCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaDomainCidrResponse
func (client *Client) DescribeDcdnIpaDomainCidrWithOptions(request *DescribeDcdnIpaDomainCidrRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaDomainCidrResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnIpaDomainCidr"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaDomainCidrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin CIDR blocks of IPA-accelerated domain names. If you want to call this API operation, you must submit a ticket to apply for the required permissions.
//
// Description:
//
// >  This operation can be called globally up to 50 times per second. This operation can be called up to 10 times per second per account.
//
// @param request - DescribeDcdnIpaDomainCidrRequest
//
// @return DescribeDcdnIpaDomainCidrResponse
func (client *Client) DescribeDcdnIpaDomainCidr(request *DescribeDcdnIpaDomainCidrRequest) (_result *DescribeDcdnIpaDomainCidrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnIpaDomainCidrResponse{}
	_body, _err := client.DescribeDcdnIpaDomainCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of an accelerated domain name. You can query the configurations of one or more features in each request.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnIpaDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaDomainConfigsResponse
func (client *Client) DescribeDcdnIpaDomainConfigsWithOptions(request *DescribeDcdnIpaDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

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
		Action:      dara.String("DescribeDcdnIpaDomainConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an accelerated domain name. You can query the configurations of one or more features in each request.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnIpaDomainConfigsRequest
//
// @return DescribeDcdnIpaDomainConfigsResponse
func (client *Client) DescribeDcdnIpaDomainConfigs(request *DescribeDcdnIpaDomainConfigsRequest) (_result *DescribeDcdnIpaDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnIpaDomainConfigsResponse{}
	_body, _err := client.DescribeDcdnIpaDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic configuration information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnIpaDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaDomainDetailResponse
func (client *Client) DescribeDcdnIpaDomainDetailWithOptions(request *DescribeDcdnIpaDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaDomainDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("DescribeDcdnIpaDomainDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic configuration information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnIpaDomainDetailRequest
//
// @return DescribeDcdnIpaDomainDetailResponse
func (client *Client) DescribeDcdnIpaDomainDetail(request *DescribeDcdnIpaDomainDetailRequest) (_result *DescribeDcdnIpaDomainDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnIpaDomainDetailResponse{}
	_body, _err := client.DescribeDcdnIpaDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of IPA. The information includes the time when the service was activated, the current service status, the current billing method, and the billing method of the next cycle.
//
// Description:
//
// *
//
// **The maximum number of times that each user can call this operation per second is 20.
//
// @param request - DescribeDcdnIpaServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaServiceResponse
func (client *Client) DescribeDcdnIpaServiceWithOptions(request *DescribeDcdnIpaServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnIpaService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of IPA. The information includes the time when the service was activated, the current service status, the current billing method, and the billing method of the next cycle.
//
// Description:
//
// *
//
// **The maximum number of times that each user can call this operation per second is 20.
//
// @param request - DescribeDcdnIpaServiceRequest
//
// @return DescribeDcdnIpaServiceResponse
func (client *Client) DescribeDcdnIpaService(request *DescribeDcdnIpaServiceRequest) (_result *DescribeDcdnIpaServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnIpaServiceResponse{}
	_body, _err := client.DescribeDcdnIpaServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all domain names that are accelerated by IP Application Accelerator (IPA) in your account. Fuzzy search and filtering by domain status are supported.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnIpaUserDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnIpaUserDomainsResponse
func (client *Client) DescribeDcdnIpaUserDomainsWithOptions(request *DescribeDcdnIpaUserDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnIpaUserDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckDomainShow) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSearchType) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !dara.IsNil(request.FuncFilter) {
		query["FuncFilter"] = request.FuncFilter
	}

	if !dara.IsNil(request.FuncId) {
		query["FuncId"] = request.FuncId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnIpaUserDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnIpaUserDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all domain names that are accelerated by IP Application Accelerator (IPA) in your account. Fuzzy search and filtering by domain status are supported.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnIpaUserDomainsRequest
//
// @return DescribeDcdnIpaUserDomainsResponse
func (client *Client) DescribeDcdnIpaUserDomains(request *DescribeDcdnIpaUserDomainsRequest) (_result *DescribeDcdnIpaUserDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnIpaUserDomainsResponse{}
	_body, _err := client.DescribeDcdnIpaUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the key-value pairs that belong to your account.
//
// @param request - DescribeDcdnKvAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnKvAccountResponse
func (client *Client) DescribeDcdnKvAccountWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnKvAccountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnKvAccount"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnKvAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the key-value pairs that belong to your account.
//
// @return DescribeDcdnKvAccountResponse
func (client *Client) DescribeDcdnKvAccount() (_result *DescribeDcdnKvAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnKvAccountResponse{}
	_body, _err := client.DescribeDcdnKvAccountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the KV status of an account.
//
// @param request - DescribeDcdnKvAccountStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnKvAccountStatusResponse
func (client *Client) DescribeDcdnKvAccountStatusWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnKvAccountStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnKvAccountStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnKvAccountStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the KV status of an account.
//
// @return DescribeDcdnKvAccountStatusResponse
func (client *Client) DescribeDcdnKvAccountStatus() (_result *DescribeDcdnKvAccountStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnKvAccountStatusResponse{}
	_body, _err := client.DescribeDcdnKvAccountStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace.
//
// @param request - DescribeDcdnKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnKvNamespaceResponse
func (client *Client) DescribeDcdnKvNamespaceWithOptions(request *DescribeDcdnKvNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnKvNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnKvNamespace"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnKvNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace.
//
// @param request - DescribeDcdnKvNamespaceRequest
//
// @return DescribeDcdnKvNamespaceResponse
func (client *Client) DescribeDcdnKvNamespace(request *DescribeDcdnKvNamespaceRequest) (_result *DescribeDcdnKvNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnKvNamespaceResponse{}
	_body, _err := client.DescribeDcdnKvNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries CIDR blocks of Dynamic Content Delivery Network (DCDN) points of presence (POPs).
//
// Description:
//
// > 	- To use this operation, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
//
// > 	- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnL2IpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnL2IpsResponse
func (client *Client) DescribeDcdnL2IpsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnL2IpsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnL2Ips"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnL2IpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries CIDR blocks of Dynamic Content Delivery Network (DCDN) points of presence (POPs).
//
// Description:
//
// > 	- To use this operation, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
//
// > 	- You can call this operation up to 100 times per second per account.
//
// @return DescribeDcdnL2IpsResponse
func (client *Client) DescribeDcdnL2Ips() (_result *DescribeDcdnL2IpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnL2IpsResponse{}
	_body, _err := client.DescribeDcdnL2IpsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the origin CIDR blocks by domain name. The CIDR blocks include IPv4 and IPv6 CIDR blocks.
//
// @param request - DescribeDcdnL2VipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnL2VipsResponse
func (client *Client) DescribeDcdnL2VipsWithOptions(request *DescribeDcdnL2VipsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnL2VipsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnL2Vips"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnL2VipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the origin CIDR blocks by domain name. The CIDR blocks include IPv4 and IPv6 CIDR blocks.
//
// @param request - DescribeDcdnL2VipsRequest
//
// @return DescribeDcdnL2VipsResponse
func (client *Client) DescribeDcdnL2Vips(request *DescribeDcdnL2VipsRequest) (_result *DescribeDcdnL2VipsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnL2VipsResponse{}
	_body, _err := client.DescribeDcdnL2VipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the origin server for a DCDN-accelerated domain name.
//
// @param request - DescribeDcdnOriginSiteHealthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnOriginSiteHealthStatusResponse
func (client *Client) DescribeDcdnOriginSiteHealthStatusWithOptions(request *DescribeDcdnOriginSiteHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnOriginSiteHealthStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnOriginSiteHealthStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnOriginSiteHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the origin server for a DCDN-accelerated domain name.
//
// @param request - DescribeDcdnOriginSiteHealthStatusRequest
//
// @return DescribeDcdnOriginSiteHealthStatusResponse
func (client *Client) DescribeDcdnOriginSiteHealthStatus(request *DescribeDcdnOriginSiteHealthStatusRequest) (_result *DescribeDcdnOriginSiteHealthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnOriginSiteHealthStatusResponse{}
	_body, _err := client.DescribeDcdnOriginSiteHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the fields in real-time log entries.
//
// Description:
//
// >  You can call this API operation up to 100 times per second per account.
//
// @param request - DescribeDcdnRealTimeDeliveryFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRealTimeDeliveryFieldResponse
func (client *Client) DescribeDcdnRealTimeDeliveryFieldWithOptions(request *DescribeDcdnRealTimeDeliveryFieldRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRealTimeDeliveryFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnRealTimeDeliveryField"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRealTimeDeliveryFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fields in real-time log entries.
//
// Description:
//
// >  You can call this API operation up to 100 times per second per account.
//
// @param request - DescribeDcdnRealTimeDeliveryFieldRequest
//
// @return DescribeDcdnRealTimeDeliveryFieldResponse
func (client *Client) DescribeDcdnRealTimeDeliveryField(request *DescribeDcdnRealTimeDeliveryFieldRequest) (_result *DescribeDcdnRealTimeDeliveryFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnRealTimeDeliveryFieldResponse{}
	_body, _err := client.DescribeDcdnRealTimeDeliveryFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the maximum number and the remaining number of URLs and directories that can be refreshed or the maximum number and the remaining number of URLs that can be prefetched per day.
//
// Description:
//
// >
//
//   - You can call the **RefreshDcdnObjectCaches*	- operation to refresh content and call the **PreloadDcdnObjectCaches*	- operation to prefetch content.
//
//   - You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnRefreshQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRefreshQuotaResponse
func (client *Client) DescribeDcdnRefreshQuotaWithOptions(request *DescribeDcdnRefreshQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRefreshQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnRefreshQuota"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRefreshQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum number and the remaining number of URLs and directories that can be refreshed or the maximum number and the remaining number of URLs that can be prefetched per day.
//
// Description:
//
// >
//
//   - You can call the **RefreshDcdnObjectCaches*	- operation to refresh content and call the **PreloadDcdnObjectCaches*	- operation to prefetch content.
//
//   - You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnRefreshQuotaRequest
//
// @return DescribeDcdnRefreshQuotaResponse
func (client *Client) DescribeDcdnRefreshQuota(request *DescribeDcdnRefreshQuotaRequest) (_result *DescribeDcdnRefreshQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnRefreshQuotaResponse{}
	_body, _err := client.DescribeDcdnRefreshQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of refresh or prefetch tasks by task ID.
//
// Description:
//
// >
//
//   - You can query data within the last three days.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnRefreshTaskByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRefreshTaskByIdResponse
func (client *Client) DescribeDcdnRefreshTaskByIdWithOptions(request *DescribeDcdnRefreshTaskByIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRefreshTaskByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnRefreshTaskById"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRefreshTaskByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of refresh or prefetch tasks by task ID.
//
// Description:
//
// >
//
//   - You can query data within the last three days.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnRefreshTaskByIdRequest
//
// @return DescribeDcdnRefreshTaskByIdResponse
func (client *Client) DescribeDcdnRefreshTaskById(request *DescribeDcdnRefreshTaskByIdRequest) (_result *DescribeDcdnRefreshTaskByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnRefreshTaskByIdResponse{}
	_body, _err := client.DescribeDcdnRefreshTaskByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the refresh or prefetch tasks. You can query the tasks in the last three days.
//
// Description:
//
//	  You can query the refresh or prefetch tasks by ID or URL.
//
//		- You can set both **TaskId*	- and **ObjectPath*	- in a request. If you do not set **TaskId*	- or **ObjectPath**, the data in the last 3 days on the first page is returned. By default, a maximum of 20 entries can be displayed on each page.
//
//		- If you specify **DomainName*	- or **Status**, you must also specify **ObjectType**.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnRefreshTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRefreshTasksResponse
func (client *Client) DescribeDcdnRefreshTasksWithOptions(request *DescribeDcdnRefreshTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRefreshTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnRefreshTasks"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRefreshTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the refresh or prefetch tasks. You can query the tasks in the last three days.
//
// Description:
//
//	  You can query the refresh or prefetch tasks by ID or URL.
//
//		- You can set both **TaskId*	- and **ObjectPath*	- in a request. If you do not set **TaskId*	- or **ObjectPath**, the data in the last 3 days on the first page is returned. By default, a maximum of 20 entries can be displayed on each page.
//
//		- If you specify **DomainName*	- or **Status**, you must also specify **ObjectType**.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnRefreshTasksRequest
//
// @return DescribeDcdnRefreshTasksResponse
func (client *Client) DescribeDcdnRefreshTasks(request *DescribeDcdnRefreshTasksRequest) (_result *DescribeDcdnRefreshTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnRefreshTasksResponse{}
	_body, _err := client.DescribeDcdnRefreshTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Internet service providers (ISPs) and regions.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnRegionAndIspRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnRegionAndIspResponse
func (client *Client) DescribeDcdnRegionAndIspWithOptions(request *DescribeDcdnRegionAndIspRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnRegionAndIspResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnRegionAndIsp"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnRegionAndIspResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Internet service providers (ISPs) and regions.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnRegionAndIspRequest
//
// @return DescribeDcdnRegionAndIspResponse
func (client *Client) DescribeDcdnRegionAndIsp(request *DescribeDcdnRegionAndIspRequest) (_result *DescribeDcdnRegionAndIspResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnRegionAndIspResponse{}
	_body, _err := client.DescribeDcdnRegionAndIspWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the content of an operations report.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DescribeDcdnReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnReportResponse
func (client *Client) DescribeDcdnReportWithOptions(request *DescribeDcdnReportRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.HttpCode) {
		query["HttpCode"] = request.HttpCode
	}

	if !dara.IsNil(request.IsOverseas) {
		query["IsOverseas"] = request.IsOverseas
	}

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnReport"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content of an operations report.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DescribeDcdnReportRequest
//
// @return DescribeDcdnReportResponse
func (client *Client) DescribeDcdnReport(request *DescribeDcdnReportRequest) (_result *DescribeDcdnReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnReportResponse{}
	_body, _err := client.DescribeDcdnReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom operations reports.
//
// Description:
//
// > 	- This operation queries the metadata of all operations reports. The statistics in the reports are not returned.
//
// > 	- You can call this operation up to three times per second per account.
//
// @param request - DescribeDcdnReportListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnReportListResponse
func (client *Client) DescribeDcdnReportListWithOptions(request *DescribeDcdnReportListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnReportListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnReportList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnReportListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom operations reports.
//
// Description:
//
// > 	- This operation queries the metadata of all operations reports. The statistics in the reports are not returned.
//
// > 	- You can call this operation up to three times per second per account.
//
// @param request - DescribeDcdnReportListRequest
//
// @return DescribeDcdnReportListResponse
func (client *Client) DescribeDcdnReportList(request *DescribeDcdnReportListRequest) (_result *DescribeDcdnReportListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnReportListResponse{}
	_body, _err := client.DescribeDcdnReportListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries supported types of real-time logs.
//
// @param request - DescribeDcdnSLSRealTimeLogTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSLSRealTimeLogTypeResponse
func (client *Client) DescribeDcdnSLSRealTimeLogTypeWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnSLSRealTimeLogTypeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSLSRealTimeLogType"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSLSRealTimeLogTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries supported types of real-time logs.
//
// @return DescribeDcdnSLSRealTimeLogTypeResponse
func (client *Client) DescribeDcdnSLSRealTimeLogType() (_result *DescribeDcdnSLSRealTimeLogTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnSLSRealTimeLogTypeResponse{}
	_body, _err := client.DescribeDcdnSLSRealTimeLogTypeWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnSLSRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSLSRealtimeLogDeliveryResponse
func (client *Client) DescribeDcdnSLSRealtimeLogDeliveryWithOptions(request *DescribeDcdnSLSRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSLSRealtimeLogDelivery"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnSLSRealtimeLogDeliveryRequest
//
// @return DescribeDcdnSLSRealtimeLogDeliveryResponse
func (client *Client) DescribeDcdnSLSRealtimeLogDelivery(request *DescribeDcdnSLSRealtimeLogDeliveryRequest) (_result *DescribeDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.DescribeDcdnSLSRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a ShangMi (SM) certificate.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnSMCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSMCertificateDetailResponse
func (client *Client) DescribeDcdnSMCertificateDetailWithOptions(request *DescribeDcdnSMCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSMCertificateDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

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
		Action:      dara.String("DescribeDcdnSMCertificateDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSMCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a ShangMi (SM) certificate.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnSMCertificateDetailRequest
//
// @return DescribeDcdnSMCertificateDetailResponse
func (client *Client) DescribeDcdnSMCertificateDetail(request *DescribeDcdnSMCertificateDetailRequest) (_result *DescribeDcdnSMCertificateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnSMCertificateDetailResponse{}
	_body, _err := client.DescribeDcdnSMCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ShangMi (SM) certificates of an accelerated domain name.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnSMCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSMCertificateListResponse
func (client *Client) DescribeDcdnSMCertificateListWithOptions(request *DescribeDcdnSMCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSMCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("DescribeDcdnSMCertificateList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSMCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ShangMi (SM) certificates of an accelerated domain name.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnSMCertificateListRequest
//
// @return DescribeDcdnSMCertificateListResponse
func (client *Client) DescribeDcdnSMCertificateList(request *DescribeDcdnSMCertificateListRequest) (_result *DescribeDcdnSMCertificateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnSMCertificateListResponse{}
	_body, _err := client.DescribeDcdnSMCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificates of accelerated domain names.
//
// @param request - DescribeDcdnSSLCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSSLCertificateListResponse
func (client *Client) DescribeDcdnSSLCertificateListWithOptions(request *DescribeDcdnSSLCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSSLCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKeyword) {
		query["SearchKeyword"] = request.SearchKeyword
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSSLCertificateList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSSLCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates of accelerated domain names.
//
// @param request - DescribeDcdnSSLCertificateListRequest
//
// @return DescribeDcdnSSLCertificateListResponse
func (client *Client) DescribeDcdnSSLCertificateList(request *DescribeDcdnSSLCertificateListRequest) (_result *DescribeDcdnSSLCertificateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnSSLCertificateListResponse{}
	_body, _err := client.DescribeDcdnSSLCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an edge security drop-down list in the Dynamic Content Delivery Network (DCDN) console.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnSecFuncInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSecFuncInfoResponse
func (client *Client) DescribeDcdnSecFuncInfoWithOptions(request *DescribeDcdnSecFuncInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnSecFuncInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SecFuncType) {
		query["SecFuncType"] = request.SecFuncType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSecFuncInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSecFuncInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an edge security drop-down list in the Dynamic Content Delivery Network (DCDN) console.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnSecFuncInfoRequest
//
// @return DescribeDcdnSecFuncInfoResponse
func (client *Client) DescribeDcdnSecFuncInfo(request *DescribeDcdnSecFuncInfoRequest) (_result *DescribeDcdnSecFuncInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnSecFuncInfoResponse{}
	_body, _err := client.DescribeDcdnSecFuncInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the version of secure Dynamic Route for CDN (DCDN) and the security rules.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnSecSpecInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSecSpecInfoResponse
func (client *Client) DescribeDcdnSecSpecInfoWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnSecSpecInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSecSpecInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSecSpecInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the version of secure Dynamic Route for CDN (DCDN) and the security rules.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @return DescribeDcdnSecSpecInfoResponse
func (client *Client) DescribeDcdnSecSpecInfo() (_result *DescribeDcdnSecSpecInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnSecSpecInfoResponse{}
	_body, _err := client.DescribeDcdnSecSpecInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the Dynamic Content Delivery Network (DCDN) service. The information includes the time when the service was activated, the current service status, the current billing method, and the billing method of the next cycle.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnServiceResponse
func (client *Client) DescribeDcdnServiceWithOptions(request *DescribeDcdnServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the Dynamic Content Delivery Network (DCDN) service. The information includes the time when the service was activated, the current service status, the current billing method, and the billing method of the next cycle.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnServiceRequest
//
// @return DescribeDcdnServiceResponse
func (client *Client) DescribeDcdnService(request *DescribeDcdnServiceRequest) (_result *DescribeDcdnServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnServiceResponse{}
	_body, _err := client.DescribeDcdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries valid virtual IP addresses (VIPs) in the staging environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnStagingIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnStagingIpResponse
func (client *Client) DescribeDcdnStagingIpWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnStagingIpResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnStagingIp"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnStagingIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries valid virtual IP addresses (VIPs) in the staging environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @return DescribeDcdnStagingIpResponse
func (client *Client) DescribeDcdnStagingIp() (_result *DescribeDcdnStagingIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnStagingIpResponse{}
	_body, _err := client.DescribeDcdnStagingIpWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom operations reports.
//
// Description:
//
// > 	- By default, this operation queries all custom operations reports. However, only one operations report can be displayed. Therefore, only one operations report is returned.
//
// > 	- You can call this API operation up to three times per second per account.
//
// @param request - DescribeDcdnSubListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnSubListResponse
func (client *Client) DescribeDcdnSubListWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnSubListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnSubList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnSubListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom operations reports.
//
// Description:
//
// > 	- By default, this operation queries all custom operations reports. However, only one operations report can be displayed. Therefore, only one operations report is returned.
//
// > 	- You can call this API operation up to three times per second per account.
//
// @return DescribeDcdnSubListResponse
func (client *Client) DescribeDcdnSubList() (_result *DescribeDcdnSubListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnSubListResponse{}
	_body, _err := client.DescribeDcdnSubListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of one or more resources.
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnTagResourcesResponse
func (client *Client) DescribeDcdnTagResourcesWithOptions(request *DescribeDcdnTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeDcdnTagResources"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of one or more resources.
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDcdnTagResourcesRequest
//
// @return DescribeDcdnTagResourcesResponse
func (client *Client) DescribeDcdnTagResources(request *DescribeDcdnTagResourcesRequest) (_result *DescribeDcdnTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnTagResourcesResponse{}
	_body, _err := client.DescribeDcdnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain names ranked by network traffic. You can query data within the last 90 days.
//
// Description:
//
// If you do not specify the StartTime and EndTime parameters, the data within the current month is queried. If you specify the StartTime and EndTime parameters, the data within the specified time range is queried.
//
// @param request - DescribeDcdnTopDomainsByFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnTopDomainsByFlowResponse
func (client *Client) DescribeDcdnTopDomainsByFlowWithOptions(request *DescribeDcdnTopDomainsByFlowRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnTopDomainsByFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnTopDomainsByFlow"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnTopDomainsByFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain names ranked by network traffic. You can query data within the last 90 days.
//
// Description:
//
// If you do not specify the StartTime and EndTime parameters, the data within the current month is queried. If you specify the StartTime and EndTime parameters, the data within the specified time range is queried.
//
// @param request - DescribeDcdnTopDomainsByFlowRequest
//
// @return DescribeDcdnTopDomainsByFlowResponse
func (client *Client) DescribeDcdnTopDomainsByFlow(request *DescribeDcdnTopDomainsByFlowRequest) (_result *DescribeDcdnTopDomainsByFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnTopDomainsByFlowResponse{}
	_body, _err := client.DescribeDcdnTopDomainsByFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the billing records of an Alibaba Cloud account. The maximum time range that you can specify is one month.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserBillHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserBillHistoryResponse
func (client *Client) DescribeDcdnUserBillHistoryWithOptions(request *DescribeDcdnUserBillHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserBillHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserBillHistory"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserBillHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the billing records of an Alibaba Cloud account. The maximum time range that you can specify is one month.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserBillHistoryRequest
//
// @return DescribeDcdnUserBillHistoryResponse
func (client *Client) DescribeDcdnUserBillHistory(request *DescribeDcdnUserBillHistoryRequest) (_result *DescribeDcdnUserBillHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserBillHistoryResponse{}
	_body, _err := client.DescribeDcdnUserBillHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metering method that is used in Dynamic Content Delivery Network (DCDN).
//
// @param request - DescribeDcdnUserBillTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserBillTypeResponse
func (client *Client) DescribeDcdnUserBillTypeWithOptions(request *DescribeDcdnUserBillTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserBillTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserBillType"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserBillTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metering method that is used in Dynamic Content Delivery Network (DCDN).
//
// @param request - DescribeDcdnUserBillTypeRequest
//
// @return DescribeDcdnUserBillTypeResponse
func (client *Client) DescribeDcdnUserBillType(request *DescribeDcdnUserBillTypeRequest) (_result *DescribeDcdnUserBillTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserBillTypeResponse{}
	_body, _err := client.DescribeDcdnUserBillTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of domain names whose SSL certificates are about to expire or have already expired.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserCertificateExpireCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserCertificateExpireCountResponse
func (client *Client) DescribeDcdnUserCertificateExpireCountWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserCertificateExpireCountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserCertificateExpireCount"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserCertificateExpireCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of domain names whose SSL certificates are about to expire or have already expired.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
// @return DescribeDcdnUserCertificateExpireCountResponse
func (client *Client) DescribeDcdnUserCertificateExpireCount() (_result *DescribeDcdnUserCertificateExpireCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserCertificateExpireCountResponse{}
	_body, _err := client.DescribeDcdnUserCertificateExpireCountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of security features.
//
// Description:
//
// You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnUserConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserConfigsResponse
func (client *Client) DescribeDcdnUserConfigsWithOptions(request *DescribeDcdnUserConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserConfigs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of security features.
//
// Description:
//
// You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnUserConfigsRequest
//
// @return DescribeDcdnUserConfigsResponse
func (client *Client) DescribeDcdnUserConfigs(request *DescribeDcdnUserConfigsRequest) (_result *DescribeDcdnUserConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserConfigsResponse{}
	_body, _err := client.DescribeDcdnUserConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the accelerated domain names that belong to your Alibaba Cloud account. You can filter domains by name or by status. Fuzzy match is supported when you filter domains by name.
//
// Description:
//
// > You can call this operation up to 80 times per second per account.
//
// @param request - DescribeDcdnUserDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserDomainsResponse
func (client *Client) DescribeDcdnUserDomainsWithOptions(request *DescribeDcdnUserDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeEndTime) {
		query["ChangeEndTime"] = request.ChangeEndTime
	}

	if !dara.IsNil(request.ChangeStartTime) {
		query["ChangeStartTime"] = request.ChangeStartTime
	}

	if !dara.IsNil(request.CheckDomainShow) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !dara.IsNil(request.Coverage) {
		query["Coverage"] = request.Coverage
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSearchType) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.WebSiteType) {
		query["WebSiteType"] = request.WebSiteType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accelerated domain names that belong to your Alibaba Cloud account. You can filter domains by name or by status. Fuzzy match is supported when you filter domains by name.
//
// Description:
//
// > You can call this operation up to 80 times per second per account.
//
// @param request - DescribeDcdnUserDomainsRequest
//
// @return DescribeDcdnUserDomainsResponse
func (client *Client) DescribeDcdnUserDomains(request *DescribeDcdnUserDomainsRequest) (_result *DescribeDcdnUserDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserDomainsResponse{}
	_body, _err := client.DescribeDcdnUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all accelerated domain names with specified features configured that belong to your Alibaba Cloud account based on the FuncId parameter.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserDomainsByFuncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserDomainsByFuncResponse
func (client *Client) DescribeDcdnUserDomainsByFuncWithOptions(request *DescribeDcdnUserDomainsByFuncRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserDomainsByFuncResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FuncFilter) {
		query["FuncFilter"] = request.FuncFilter
	}

	if !dara.IsNil(request.FuncId) {
		query["FuncId"] = request.FuncId
	}

	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserDomainsByFunc"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserDomainsByFuncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all accelerated domain names with specified features configured that belong to your Alibaba Cloud account based on the FuncId parameter.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserDomainsByFuncRequest
//
// @return DescribeDcdnUserDomainsByFuncResponse
func (client *Client) DescribeDcdnUserDomainsByFunc(request *DescribeDcdnUserDomainsByFuncRequest) (_result *DescribeDcdnUserDomainsByFuncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserDomainsByFuncResponse{}
	_body, _err := client.DescribeDcdnUserDomainsByFuncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource quotas and the used resources.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 30.
//
// @param request - DescribeDcdnUserQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserQuotaResponse
func (client *Client) DescribeDcdnUserQuotaWithOptions(request *DescribeDcdnUserQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnUserQuota"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource quotas and the used resources.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 30.
//
// @param request - DescribeDcdnUserQuotaRequest
//
// @return DescribeDcdnUserQuotaResponse
func (client *Client) DescribeDcdnUserQuota(request *DescribeDcdnUserQuotaRequest) (_result *DescribeDcdnUserQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserQuotaResponse{}
	_body, _err := client.DescribeDcdnUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the fields that are selected.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserRealTimeDeliveryFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserRealTimeDeliveryFieldResponse
func (client *Client) DescribeDcdnUserRealTimeDeliveryFieldWithOptions(request *DescribeDcdnUserRealTimeDeliveryFieldRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserRealTimeDeliveryField"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fields that are selected.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserRealTimeDeliveryFieldRequest
//
// @return DescribeDcdnUserRealTimeDeliveryFieldResponse
func (client *Client) DescribeDcdnUserRealTimeDeliveryField(request *DescribeDcdnUserRealTimeDeliveryFieldRequest) (_result *DescribeDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.DescribeDcdnUserRealTimeDeliveryFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the resource plans in your Alibaba Cloud account.
//
// Description:
//
// The maximum number of times that each user can call this operation per second is 30.
//
// @param request - DescribeDcdnUserResourcePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserResourcePackageResponse
func (client *Client) DescribeDcdnUserResourcePackageWithOptions(request *DescribeDcdnUserResourcePackageRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserResourcePackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserResourcePackage"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the resource plans in your Alibaba Cloud account.
//
// Description:
//
// The maximum number of times that each user can call this operation per second is 30.
//
// @param request - DescribeDcdnUserResourcePackageRequest
//
// @return DescribeDcdnUserResourcePackageResponse
func (client *Client) DescribeDcdnUserResourcePackage(request *DescribeDcdnUserResourcePackageRequest) (_result *DescribeDcdnUserResourcePackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserResourcePackageResponse{}
	_body, _err := client.DescribeDcdnUserResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of packets blocked by a specified security feature.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnUserSecDropRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserSecDropResponse
func (client *Client) DescribeDcdnUserSecDropWithOptions(request *DescribeDcdnUserSecDropRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserSecDropResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.SecFunc) {
		query["SecFunc"] = request.SecFunc
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserSecDrop"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserSecDropResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of packets blocked by a specified security feature.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnUserSecDropRequest
//
// @return DescribeDcdnUserSecDropResponse
func (client *Client) DescribeDcdnUserSecDrop(request *DescribeDcdnUserSecDropRequest) (_result *DescribeDcdnUserSecDropResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserSecDropResponse{}
	_body, _err := client.DescribeDcdnUserSecDropWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of packets that are blocked by security features at the application layer in a specific time range.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnUserSecDropByMinuteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserSecDropByMinuteResponse
func (client *Client) DescribeDcdnUserSecDropByMinuteWithOptions(request *DescribeDcdnUserSecDropByMinuteRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserSecDropByMinuteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Object) {
		query["Object"] = request.Object
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

	if !dara.IsNil(request.SecFunc) {
		query["SecFunc"] = request.SecFunc
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserSecDropByMinute"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserSecDropByMinuteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of packets that are blocked by security features at the application layer in a specific time range.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnUserSecDropByMinuteRequest
//
// @return DescribeDcdnUserSecDropByMinuteResponse
func (client *Client) DescribeDcdnUserSecDropByMinute(request *DescribeDcdnUserSecDropByMinuteRequest) (_result *DescribeDcdnUserSecDropByMinuteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserSecDropByMinuteResponse{}
	_body, _err := client.DescribeDcdnUserSecDropByMinuteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user tags.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnUserTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserTagsResponse
func (client *Client) DescribeDcdnUserTagsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserTagsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserTags"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user tags.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @return DescribeDcdnUserTagsResponse
func (client *Client) DescribeDcdnUserTags() (_result *DescribeDcdnUserTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserTagsResponse{}
	_body, _err := client.DescribeDcdnUserTagsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries virtual IP addresses of the POPs by domain name.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnUserVipsByDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnUserVipsByDomainResponse
func (client *Client) DescribeDcdnUserVipsByDomainWithOptions(request *DescribeDcdnUserVipsByDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnUserVipsByDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Available) {
		query["Available"] = request.Available
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnUserVipsByDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnUserVipsByDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries virtual IP addresses of the POPs by domain name.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDcdnUserVipsByDomainRequest
//
// @return DescribeDcdnUserVipsByDomainResponse
func (client *Client) DescribeDcdnUserVipsByDomain(request *DescribeDcdnUserVipsByDomainRequest) (_result *DescribeDcdnUserVipsByDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnUserVipsByDomainResponse{}
	_body, _err := client.DescribeDcdnUserVipsByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ownership verification content of a domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnVerifyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnVerifyContentResponse
func (client *Client) DescribeDcdnVerifyContentWithOptions(request *DescribeDcdnVerifyContentRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnVerifyContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnVerifyContent"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnVerifyContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ownership verification content of a domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnVerifyContentRequest
//
// @return DescribeDcdnVerifyContentResponse
func (client *Client) DescribeDcdnVerifyContent(request *DescribeDcdnVerifyContentRequest) (_result *DescribeDcdnVerifyContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnVerifyContentResponse{}
	_body, _err := client.DescribeDcdnVerifyContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SDK authentication key for the Alibaba Cloud account. You can also use the SDK authentication key to send SDK initialization requests. The key must be included in the integration code.
//
// @param request - DescribeDcdnWafBotAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafBotAppKeyResponse
func (client *Client) DescribeDcdnWafBotAppKeyWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafBotAppKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafBotAppKey"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafBotAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SDK authentication key for the Alibaba Cloud account. You can also use the SDK authentication key to send SDK initialization requests. The key must be included in the integration code.
//
// @return DescribeDcdnWafBotAppKeyResponse
func (client *Client) DescribeDcdnWafBotAppKey() (_result *DescribeDcdnWafBotAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafBotAppKeyResponse{}
	_body, _err := client.DescribeDcdnWafBotAppKeyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default configurations of a WAF rule.
//
// @param request - DescribeDcdnWafDefaultRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafDefaultRulesResponse
func (client *Client) DescribeDcdnWafDefaultRulesWithOptions(request *DescribeDcdnWafDefaultRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafDefaultRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafDefaultRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafDefaultRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default configurations of a WAF rule.
//
// @param request - DescribeDcdnWafDefaultRulesRequest
//
// @return DescribeDcdnWafDefaultRulesResponse
func (client *Client) DescribeDcdnWafDefaultRules(request *DescribeDcdnWafDefaultRulesRequest) (_result *DescribeDcdnWafDefaultRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafDefaultRulesResponse{}
	_body, _err := client.DescribeDcdnWafDefaultRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain names that have Web Application Firewall (WAF) enabled and the relevant information, including the status of the access control list (ACL), protection against HTTP flood attacks, domain name, and WAF.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnWafDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafDomainResponse
func (client *Client) DescribeDcdnWafDomainWithOptions(request *DescribeDcdnWafDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain names that have Web Application Firewall (WAF) enabled and the relevant information, including the status of the access control list (ACL), protection against HTTP flood attacks, domain name, and WAF.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDcdnWafDomainRequest
//
// @return DescribeDcdnWafDomainResponse
func (client *Client) DescribeDcdnWafDomain(request *DescribeDcdnWafDomainRequest) (_result *DescribeDcdnWafDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafDomainResponse{}
	_body, _err := client.DescribeDcdnWafDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the protection policy of a domain name.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafDomainDetailResponse
func (client *Client) DescribeDcdnWafDomainDetailWithOptions(request *DescribeDcdnWafDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafDomainDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafDomainDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the protection policy of a domain name.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafDomainDetailRequest
//
// @return DescribeDcdnWafDomainDetailResponse
func (client *Client) DescribeDcdnWafDomainDetail(request *DescribeDcdnWafDomainDetailRequest) (_result *DescribeDcdnWafDomainDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafDomainDetailResponse{}
	_body, _err := client.DescribeDcdnWafDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the accelerated domain names that are protected by Web Application Firewall (WAF). Fuzzy search is supported.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafDomainsResponse
func (client *Client) DescribeDcdnWafDomainsWithOptions(request *DescribeDcdnWafDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accelerated domain names that are protected by Web Application Firewall (WAF). Fuzzy search is supported.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafDomainsRequest
//
// @return DescribeDcdnWafDomainsResponse
func (client *Client) DescribeDcdnWafDomains(request *DescribeDcdnWafDomainsRequest) (_result *DescribeDcdnWafDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafDomainsResponse{}
	_body, _err := client.DescribeDcdnWafDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about match conditions in a custom protection rule, such as the match fields, logical characters, and match content.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafFilterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafFilterInfoResponse
func (client *Client) DescribeDcdnWafFilterInfoWithOptions(request *DescribeDcdnWafFilterInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafFilterInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScenes) {
		query["DefenseScenes"] = request.DefenseScenes
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafFilterInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafFilterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about match conditions in a custom protection rule, such as the match fields, logical characters, and match content.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafFilterInfoRequest
//
// @return DescribeDcdnWafFilterInfoResponse
func (client *Client) DescribeDcdnWafFilterInfo(request *DescribeDcdnWafFilterInfoRequest) (_result *DescribeDcdnWafFilterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafFilterInfoResponse{}
	_body, _err := client.DescribeDcdnWafFilterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the countries and regions that can be added to the blacklist of Web Application Firewall (WAF).
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafGeoInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafGeoInfoResponse
func (client *Client) DescribeDcdnWafGeoInfoWithOptions(request *DescribeDcdnWafGeoInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafGeoInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafGeoInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafGeoInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the countries and regions that can be added to the blacklist of Web Application Firewall (WAF).
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafGeoInfoRequest
//
// @return DescribeDcdnWafGeoInfoResponse
func (client *Client) DescribeDcdnWafGeoInfo(request *DescribeDcdnWafGeoInfoRequest) (_result *DescribeDcdnWafGeoInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafGeoInfoResponse{}
	_body, _err := client.DescribeDcdnWafGeoInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a custom WAF rule group by page.
//
// @param request - DescribeDcdnWafGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafGroupResponse
func (client *Client) DescribeDcdnWafGroupWithOptions(request *DescribeDcdnWafGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafGroup"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a custom WAF rule group by page.
//
// @param request - DescribeDcdnWafGroupRequest
//
// @return DescribeDcdnWafGroupResponse
func (client *Client) DescribeDcdnWafGroup(request *DescribeDcdnWafGroupRequest) (_result *DescribeDcdnWafGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafGroupResponse{}
	_body, _err := client.DescribeDcdnWafGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom Web Application Firewall (WAF) rule groups.
//
// @param request - DescribeDcdnWafGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafGroupsResponse
func (client *Client) DescribeDcdnWafGroupsWithOptions(request *DescribeDcdnWafGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafGroups"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom Web Application Firewall (WAF) rule groups.
//
// @param request - DescribeDcdnWafGroupsRequest
//
// @return DescribeDcdnWafGroupsResponse
func (client *Client) DescribeDcdnWafGroups(request *DescribeDcdnWafGroupsRequest) (_result *DescribeDcdnWafGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafGroupsResponse{}
	_body, _err := client.DescribeDcdnWafGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the address from which you can download the Web Application Firewall (WAF) logs of a domain name.
//
// Description:
//
// >
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//   - The log data is collected every hour.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnWafLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafLogsResponse
func (client *Client) DescribeDcdnWafLogsWithOptions(request *DescribeDcdnWafLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafLogs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the address from which you can download the Web Application Firewall (WAF) logs of a domain name.
//
// Description:
//
// >
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//   - The log data is collected every hour.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDcdnWafLogsRequest
//
// @return DescribeDcdnWafLogsResponse
func (client *Client) DescribeDcdnWafLogs(request *DescribeDcdnWafLogsRequest) (_result *DescribeDcdnWafLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafLogsResponse{}
	_body, _err := client.DescribeDcdnWafLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the Web Application Firewall (WAF) protection policies that you configured.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafPoliciesResponse
func (client *Client) DescribeDcdnWafPoliciesWithOptions(request *DescribeDcdnWafPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafPoliciesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafPolicies"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the Web Application Firewall (WAF) protection policies that you configured.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafPoliciesRequest
//
// @return DescribeDcdnWafPoliciesResponse
func (client *Client) DescribeDcdnWafPolicies(request *DescribeDcdnWafPoliciesRequest) (_result *DescribeDcdnWafPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafPoliciesResponse{}
	_body, _err := client.DescribeDcdnWafPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a protection policy.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafPolicyResponse
func (client *Client) DescribeDcdnWafPolicyWithOptions(request *DescribeDcdnWafPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafPolicy"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a protection policy.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafPolicyRequest
//
// @return DescribeDcdnWafPolicyResponse
func (client *Client) DescribeDcdnWafPolicy(request *DescribeDcdnWafPolicyRequest) (_result *DescribeDcdnWafPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafPolicyResponse{}
	_body, _err := client.DescribeDcdnWafPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the accelerated domain names that are protected by a specified Web Application Firewall (WAF) protection policy.
//
// Description:
//
// You can call this operation up to 20 times per second per user.
//
// @param request - DescribeDcdnWafPolicyDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafPolicyDomainsResponse
func (client *Client) DescribeDcdnWafPolicyDomainsWithOptions(request *DescribeDcdnWafPolicyDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafPolicyDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnWafPolicyDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafPolicyDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accelerated domain names that are protected by a specified Web Application Firewall (WAF) protection policy.
//
// Description:
//
// You can call this operation up to 20 times per second per user.
//
// @param request - DescribeDcdnWafPolicyDomainsRequest
//
// @return DescribeDcdnWafPolicyDomainsResponse
func (client *Client) DescribeDcdnWafPolicyDomains(request *DescribeDcdnWafPolicyDomainsRequest) (_result *DescribeDcdnWafPolicyDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafPolicyDomainsResponse{}
	_body, _err := client.DescribeDcdnWafPolicyDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the domain names that can be bound to a custom protection policy.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafPolicyValidDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafPolicyValidDomainsResponse
func (client *Client) DescribeDcdnWafPolicyValidDomainsWithOptions(request *DescribeDcdnWafPolicyValidDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafPolicyValidDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.DomainNameLike) {
		query["DomainNameLike"] = request.DomainNameLike
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
		Action:      dara.String("DescribeDcdnWafPolicyValidDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafPolicyValidDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names that can be bound to a custom protection policy.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafPolicyValidDomainsRequest
//
// @return DescribeDcdnWafPolicyValidDomainsResponse
func (client *Client) DescribeDcdnWafPolicyValidDomains(request *DescribeDcdnWafPolicyValidDomainsRequest) (_result *DescribeDcdnWafPolicyValidDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafPolicyValidDomainsResponse{}
	_body, _err := client.DescribeDcdnWafPolicyValidDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified protection rule.
//
// Description:
//
// #
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafRuleResponse
func (client *Client) DescribeDcdnWafRuleWithOptions(request *DescribeDcdnWafRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafRule"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified protection rule.
//
// Description:
//
// #
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafRuleRequest
//
// @return DescribeDcdnWafRuleResponse
func (client *Client) DescribeDcdnWafRule(request *DescribeDcdnWafRuleRequest) (_result *DescribeDcdnWafRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafRuleResponse{}
	_body, _err := client.DescribeDcdnWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the protection rules that you configured.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafRulesResponse
func (client *Client) DescribeDcdnWafRulesWithOptions(request *DescribeDcdnWafRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryArgs) {
		query["QueryArgs"] = request.QueryArgs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafRules"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the protection rules that you configured.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafRulesRequest
//
// @return DescribeDcdnWafRulesResponse
func (client *Client) DescribeDcdnWafRules(request *DescribeDcdnWafRulesRequest) (_result *DescribeDcdnWafRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafRulesResponse{}
	_body, _err := client.DescribeDcdnWafRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the type of the protection policy that you use.
//
// Description:
//
// You can call this operation up to 20 times per second per user.
//
// @param request - DescribeDcdnWafScenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafScenesResponse
func (client *Client) DescribeDcdnWafScenesWithOptions(request *DescribeDcdnWafScenesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafScenesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScenes) {
		query["DefenseScenes"] = request.DefenseScenes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafScenes"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the type of the protection policy that you use.
//
// Description:
//
// You can call this operation up to 20 times per second per user.
//
// @param request - DescribeDcdnWafScenesRequest
//
// @return DescribeDcdnWafScenesResponse
func (client *Client) DescribeDcdnWafScenes(request *DescribeDcdnWafScenesRequest) (_result *DescribeDcdnWafScenesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafScenesResponse{}
	_body, _err := client.DescribeDcdnWafScenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about Dynamic Content Delivery Network (DCDN) Web Application Firewall WAF), including the time when WAF is enabled, edition of WAF, current status of WAF, metering method for requests, and metering method for rules.
//
// Description:
//
// # Usage notes
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafServiceResponse
func (client *Client) DescribeDcdnWafServiceWithOptions(request *DescribeDcdnWafServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnWafService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Dynamic Content Delivery Network (DCDN) Web Application Firewall WAF), including the time when WAF is enabled, edition of WAF, current status of WAF, metering method for requests, and metering method for rules.
//
// Description:
//
// # Usage notes
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafServiceRequest
//
// @return DescribeDcdnWafServiceResponse
func (client *Client) DescribeDcdnWafService(request *DescribeDcdnWafServiceRequest) (_result *DescribeDcdnWafServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafServiceResponse{}
	_body, _err := client.DescribeDcdnWafServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the version of Web Application Firewall (WAF) used in Dynamic Content Delivery Network (DCDN).
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnWafSpecInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafSpecInfoResponse
func (client *Client) DescribeDcdnWafSpecInfoWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafSpecInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafSpecInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafSpecInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the version of Web Application Firewall (WAF) used in Dynamic Content Delivery Network (DCDN).
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @return DescribeDcdnWafSpecInfoResponse
func (client *Client) DescribeDcdnWafSpecInfo() (_result *DescribeDcdnWafSpecInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafSpecInfoResponse{}
	_body, _err := client.DescribeDcdnWafSpecInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The number of used SeCUs.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- The minimum time granularity for a query is 5 minutes. The maximum time span for a query is 31 days. The time period within which historical data is available for a query is 90 days.
//
// @param request - DescribeDcdnWafUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnWafUsageDataResponse
func (client *Client) DescribeDcdnWafUsageDataWithOptions(request *DescribeDcdnWafUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnWafUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDcdnWafUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnWafUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The number of used SeCUs.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- The minimum time granularity for a query is 5 minutes. The maximum time span for a query is 31 days. The time period within which historical data is available for a query is 90 days.
//
// @param request - DescribeDcdnWafUsageDataRequest
//
// @return DescribeDcdnWafUsageDataResponse
func (client *Client) DescribeDcdnWafUsageData(request *DescribeDcdnWafUsageDataRequest) (_result *DescribeDcdnWafUsageDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnWafUsageDataResponse{}
	_body, _err := client.DescribeDcdnWafUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about Dynamic Content Delivery Network (DCDN), such as the service activation time, the expiration time, and the current status.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnsecServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDcdnsecServiceResponse
func (client *Client) DescribeDcdnsecServiceWithOptions(request *DescribeDcdnsecServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDcdnsecServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeDcdnsecService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDcdnsecServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Dynamic Content Delivery Network (DCDN), such as the service activation time, the expiration time, and the current status.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDcdnsecServiceRequest
//
// @return DescribeDcdnsecServiceResponse
func (client *Client) DescribeDcdnsecService(request *DescribeDcdnsecServiceRequest) (_result *DescribeDcdnsecServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDcdnsecServiceResponse{}
	_body, _err := client.DescribeDcdnsecServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries attack events.
//
// @param request - DescribeDdosAllEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosAllEventListResponse
func (client *Client) DescribeDdosAllEventListWithOptions(request *DescribeDdosAllEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDdosAllEventListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDdosAllEventList"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDdosAllEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries attack events.
//
// @param request - DescribeDdosAllEventListRequest
//
// @return DescribeDdosAllEventListResponse
func (client *Client) DescribeDdosAllEventList(request *DescribeDdosAllEventListRequest) (_result *DescribeDdosAllEventListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDdosAllEventListResponse{}
	_body, _err := client.DescribeDdosAllEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the encrypted RoutineUid of a routine.
//
// @param request - DescribeEncryptRoutineUidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEncryptRoutineUidResponse
func (client *Client) DescribeEncryptRoutineUidWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeEncryptRoutineUidResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEncryptRoutineUid"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEncryptRoutineUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the encrypted RoutineUid of a routine.
//
// @return DescribeEncryptRoutineUidResponse
func (client *Client) DescribeEncryptRoutineUid() (_result *DescribeEncryptRoutineUidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEncryptRoutineUidResponse{}
	_body, _err := client.DescribeEncryptRoutineUidWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the highlighted data of attack details. You can query the reasons for which requests are blocked based on TraceIDs in logs of requests that are blocked by Basic Web Protection. The highlighted data matches the content blocked by the basic web protection module.
//
// @param request - DescribeHighlightInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHighlightInfoResponse
func (client *Client) DescribeHighlightInfoWithOptions(request *DescribeHighlightInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeHighlightInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TraceId) {
		query["TraceId"] = request.TraceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHighlightInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHighlightInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the highlighted data of attack details. You can query the reasons for which requests are blocked based on TraceIDs in logs of requests that are blocked by Basic Web Protection. The highlighted data matches the content blocked by the basic web protection module.
//
// @param request - DescribeHighlightInfoRequest
//
// @return DescribeHighlightInfoResponse
func (client *Client) DescribeHighlightInfo(request *DescribeHighlightInfoRequest) (_result *DescribeHighlightInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHighlightInfoResponse{}
	_body, _err := client.DescribeHighlightInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// kv存储实时Qps监控数据
//
// @param request - DescribeKvRealTimeQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKvRealTimeQpsDataResponse
func (client *Client) DescribeKvRealTimeQpsDataWithOptions(request *DescribeKvRealTimeQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeKvRealTimeQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKvRealTimeQpsData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKvRealTimeQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// kv存储实时Qps监控数据
//
// @param request - DescribeKvRealTimeQpsDataRequest
//
// @return DescribeKvRealTimeQpsDataResponse
func (client *Client) DescribeKvRealTimeQpsData(request *DescribeKvRealTimeQpsDataRequest) (_result *DescribeKvRealTimeQpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeKvRealTimeQpsDataResponse{}
	_body, _err := client.DescribeKvRealTimeQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage data of KV storage.
//
// Description:
//
// *Note**	- You can call this operation up to 5 times per second per account.
//
//   - The usage data indicates the number of requests.
//
// **Time granularity:*	- This operation supports only the time granularity of 1 hour.
//
// |Time granularity|Time range to query|Historical data available|Data latency|
//
// |---|---|---|---|
//
// |1 hour|31 days|90 days|3 to 4 hours|
//
// @param request - DescribeKvUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKvUsageDataResponse
func (client *Client) DescribeKvUsageDataWithOptions(request *DescribeKvUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeKvUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Field) {
		query["Field"] = request.Field
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.ResponseType) {
		query["ResponseType"] = request.ResponseType
	}

	if !dara.IsNil(request.SplitBy) {
		query["SplitBy"] = request.SplitBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKvUsageData"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKvUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage data of KV storage.
//
// Description:
//
// *Note**	- You can call this operation up to 5 times per second per account.
//
//   - The usage data indicates the number of requests.
//
// **Time granularity:*	- This operation supports only the time granularity of 1 hour.
//
// |Time granularity|Time range to query|Historical data available|Data latency|
//
// |---|---|---|---|
//
// |1 hour|31 days|90 days|3 to 4 hours|
//
// @param request - DescribeKvUsageDataRequest
//
// @return DescribeKvUsageDataResponse
func (client *Client) DescribeKvUsageData(request *DescribeKvUsageDataRequest) (_result *DescribeKvUsageDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeKvUsageDataResponse{}
	_body, _err := client.DescribeKvUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the feature configurations of an accelerated domain name in the resource directory.
//
// @param request - DescribeRDDomainConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRDDomainConfigResponse
func (client *Client) DescribeRDDomainConfigWithOptions(request *DescribeRDDomainConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeRDDomainConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRDDomainConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRDDomainConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the feature configurations of an accelerated domain name in the resource directory.
//
// @param request - DescribeRDDomainConfigRequest
//
// @return DescribeRDDomainConfigResponse
func (client *Client) DescribeRDDomainConfig(request *DescribeRDDomainConfigRequest) (_result *DescribeRDDomainConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRDDomainConfigResponse{}
	_body, _err := client.DescribeRDDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all domain names of Alibaba Cloud CDN and Dynamic Content Delivery Network (DCDN) in a Resource Directory (RD).
//
// Description:
//
// A domain name can be in one of the following states:
//
//   - online
//
//   - offline
//
//   - configuring
//
//   - configure_failed
//
//   - checking
//
//   - check_failed
//
// @param request - DescribeRDDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRDDomainsResponse
func (client *Client) DescribeRDDomainsWithOptions(request *DescribeRDDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRDDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeRDDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRDDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all domain names of Alibaba Cloud CDN and Dynamic Content Delivery Network (DCDN) in a Resource Directory (RD).
//
// Description:
//
// A domain name can be in one of the following states:
//
//   - online
//
//   - offline
//
//   - configuring
//
//   - configure_failed
//
//   - checking
//
//   - check_failed
//
// @param request - DescribeRDDomainsRequest
//
// @return DescribeRDDomainsResponse
func (client *Client) DescribeRDDomains(request *DescribeRDDomainsRequest) (_result *DescribeRDDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRDDomainsResponse{}
	_body, _err := client.DescribeRDDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metadata of a specified routine. The metadata includes the routine configuration, configuration version, and code version.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRoutineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoutineResponse
func (client *Client) DescribeRoutineWithOptions(request *DescribeRoutineRequest, runtime *dara.RuntimeOptions) (_result *DescribeRoutineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoutine"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoutineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata of a specified routine. The metadata includes the routine configuration, configuration version, and code version.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRoutineRequest
//
// @return DescribeRoutineResponse
func (client *Client) DescribeRoutine(request *DescribeRoutineRequest) (_result *DescribeRoutineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRoutineResponse{}
	_body, _err := client.DescribeRoutineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the canary release environments that are supported by a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRoutineCanaryEnvsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoutineCanaryEnvsResponse
func (client *Client) DescribeRoutineCanaryEnvsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeRoutineCanaryEnvsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoutineCanaryEnvs"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoutineCanaryEnvsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the canary release environments that are supported by a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @return DescribeRoutineCanaryEnvsResponse
func (client *Client) DescribeRoutineCanaryEnvs() (_result *DescribeRoutineCanaryEnvsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRoutineCanaryEnvsResponse{}
	_body, _err := client.DescribeRoutineCanaryEnvsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the JavaScript code version of a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRoutineCodeRevisionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoutineCodeRevisionResponse
func (client *Client) DescribeRoutineCodeRevisionWithOptions(request *DescribeRoutineCodeRevisionRequest, runtime *dara.RuntimeOptions) (_result *DescribeRoutineCodeRevisionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SelectCodeRevision) {
		body["SelectCodeRevision"] = request.SelectCodeRevision
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoutineCodeRevision"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoutineCodeRevisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the JavaScript code version of a routine.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRoutineCodeRevisionRequest
//
// @return DescribeRoutineCodeRevisionResponse
func (client *Client) DescribeRoutineCodeRevision(request *DescribeRoutineCodeRevisionRequest) (_result *DescribeRoutineCodeRevisionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRoutineCodeRevisionResponse{}
	_body, _err := client.DescribeRoutineCodeRevisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of domain names that are associated with a routine.
//
// @param request - DescribeRoutineRelatedDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoutineRelatedDomainsResponse
func (client *Client) DescribeRoutineRelatedDomainsWithOptions(request *DescribeRoutineRelatedDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRoutineRelatedDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoutineRelatedDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoutineRelatedDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of domain names that are associated with a routine.
//
// @param request - DescribeRoutineRelatedDomainsRequest
//
// @return DescribeRoutineRelatedDomainsResponse
func (client *Client) DescribeRoutineRelatedDomains(request *DescribeRoutineRelatedDomainsRequest) (_result *DescribeRoutineRelatedDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRoutineRelatedDomainsResponse{}
	_body, _err := client.DescribeRoutineRelatedDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the supported specifications for routines. The private preview supports the following CPU time slice specifications: 5 ms, 50 ms, and 100 ms.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRoutineSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoutineSpecResponse
func (client *Client) DescribeRoutineSpecWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeRoutineSpecResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoutineSpec"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoutineSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the supported specifications for routines. The private preview supports the following CPU time slice specifications: 5 ms, 50 ms, and 100 ms.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @return DescribeRoutineSpecResponse
func (client *Client) DescribeRoutineSpec() (_result *DescribeRoutineSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRoutineSpecResponse{}
	_body, _err := client.DescribeRoutineSpecWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the subdomains and routines that belong to your Alibaba Cloud account.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRoutineUserInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoutineUserInfoResponse
func (client *Client) DescribeRoutineUserInfoWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeRoutineUserInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRoutineUserInfo"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoutineUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subdomains and routines that belong to your Alibaba Cloud account.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @return DescribeRoutineUserInfoResponse
func (client *Client) DescribeRoutineUserInfo() (_result *DescribeRoutineUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRoutineUserInfoResponse{}
	_body, _err := client.DescribeRoutineUserInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Whether IPA is enabled and whether you have overdue payments for your IPA are queried.
//
// Description:
//
// *
//
// **The maximum number of times that each user can call this operation per second is 20.
//
// @param request - DescribeUserDcdnIpaStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserDcdnIpaStatusResponse
func (client *Client) DescribeUserDcdnIpaStatusWithOptions(request *DescribeUserDcdnIpaStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserDcdnIpaStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeUserDcdnIpaStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserDcdnIpaStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Whether IPA is enabled and whether you have overdue payments for your IPA are queried.
//
// Description:
//
// *
//
// **The maximum number of times that each user can call this operation per second is 20.
//
// @param request - DescribeUserDcdnIpaStatusRequest
//
// @return DescribeUserDcdnIpaStatusResponse
func (client *Client) DescribeUserDcdnIpaStatus(request *DescribeUserDcdnIpaStatusRequest) (_result *DescribeUserDcdnIpaStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserDcdnIpaStatusResponse{}
	_body, _err := client.DescribeUserDcdnIpaStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether DCDN is activated and whether your account has overdue payments.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeUserDcdnStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserDcdnStatusResponse
func (client *Client) DescribeUserDcdnStatusWithOptions(request *DescribeUserDcdnStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserDcdnStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserDcdnStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserDcdnStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether DCDN is activated and whether your account has overdue payments.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeUserDcdnStatusRequest
//
// @return DescribeUserDcdnStatusResponse
func (client *Client) DescribeUserDcdnStatus(request *DescribeUserDcdnStatusRequest) (_result *DescribeUserDcdnStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserDcdnStatusResponse{}
	_body, _err := client.DescribeUserDcdnStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether EdgeRoutine (ER) is activated or has an overdue payment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeUserErStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserErStatusResponse
func (client *Client) DescribeUserErStatusWithOptions(request *DescribeUserErStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserErStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserErStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserErStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether EdgeRoutine (ER) is activated or has an overdue payment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeUserErStatusRequest
//
// @return DescribeUserErStatusResponse
func (client *Client) DescribeUserErStatus(request *DescribeUserErStatusRequest) (_result *DescribeUserErStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserErStatusResponse{}
	_body, _err := client.DescribeUserErStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Log Service is activated and whether you have overdue payments for your Log Service.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeUserLogserviceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserLogserviceStatusResponse
func (client *Client) DescribeUserLogserviceStatusWithOptions(request *DescribeUserLogserviceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserLogserviceStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("DescribeUserLogserviceStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserLogserviceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Log Service is activated and whether you have overdue payments for your Log Service.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeUserLogserviceStatusRequest
//
// @return DescribeUserLogserviceStatusResponse
func (client *Client) DescribeUserLogserviceStatus(request *DescribeUserLogserviceStatusRequest) (_result *DescribeUserLogserviceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserLogserviceStatusResponse{}
	_body, _err := client.DescribeUserLogserviceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a routine.
//
// Description:
//
// >
//
//   - This operation modifies only the specified configurations. Other configurations remain unchanged.
//
//   - If you want to delete a setting, delete the parameter value.
//
//   - This operation can add canary release environments. Make sure that the environment names comply with the naming rules. Otherwise, you will fail to add the environments.
//
//   - Dynamic Route for CDN (DCDN) provides 35 canary release environments. Among these environments, 34 are deployed in China and 1 is deployed outside China. The canary release environments are:
//
//   - Outside China: presetCanaryOverseas.
//
//   - In China: The 34 canary release environments are named in the format of presetCanaryXX. For example, presetCanaryBeijing represents the canary release environment in Beijing. A canary release environment is in each of the following regions: Anhui, Beijing, Chongqing, Fujian, Gansu, Guangdong, Guangxi, Guizhou, Hainan, Hebei, Heilongjiang, Henan, Hong Kong, Hubei, Hunan, Jiangsu, Jiangxi, Jilin, Liaoning, Macao, Neimenggu, Ningxia, Qinghai, Shaanxi, Shandong, Shanghai, Shanxi, Sichuan, Taiwan, Tianjin, Xinjiang, Xizang, Yunan, and Zhejiang.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param tmpReq - EditRoutineConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditRoutineConfResponse
func (client *Client) EditRoutineConfWithOptions(tmpReq *EditRoutineConfRequest, runtime *dara.RuntimeOptions) (_result *EditRoutineConfResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &EditRoutineConfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EnvConf) {
		request.EnvConfShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EnvConf, dara.String("EnvConf"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvConfShrink) {
		body["EnvConf"] = request.EnvConfShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditRoutineConf"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditRoutineConfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a routine.
//
// Description:
//
// >
//
//   - This operation modifies only the specified configurations. Other configurations remain unchanged.
//
//   - If you want to delete a setting, delete the parameter value.
//
//   - This operation can add canary release environments. Make sure that the environment names comply with the naming rules. Otherwise, you will fail to add the environments.
//
//   - Dynamic Route for CDN (DCDN) provides 35 canary release environments. Among these environments, 34 are deployed in China and 1 is deployed outside China. The canary release environments are:
//
//   - Outside China: presetCanaryOverseas.
//
//   - In China: The 34 canary release environments are named in the format of presetCanaryXX. For example, presetCanaryBeijing represents the canary release environment in Beijing. A canary release environment is in each of the following regions: Anhui, Beijing, Chongqing, Fujian, Gansu, Guangdong, Guangxi, Guizhou, Hainan, Hebei, Heilongjiang, Henan, Hong Kong, Hubei, Hunan, Jiangsu, Jiangxi, Jilin, Liaoning, Macao, Neimenggu, Ningxia, Qinghai, Shaanxi, Shandong, Shanghai, Shanxi, Sichuan, Taiwan, Tianjin, Xinjiang, Xizang, Yunan, and Zhejiang.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - EditRoutineConfRequest
//
// @return EditRoutineConfResponse
func (client *Client) EditRoutineConf(request *EditRoutineConfRequest) (_result *EditRoutineConfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditRoutineConfResponse{}
	_body, _err := client.EditRoutineConfWithOptions(request, runtime)
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
// @param request - GetDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDcdnKvResponse
func (client *Client) GetDcdnKvWithOptions(request *GetDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *GetDcdnKvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDcdnKvResponse{}
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
// @param request - GetDcdnKvRequest
//
// @return GetDcdnKvResponse
func (client *Client) GetDcdnKv(request *GetDcdnKvRequest) (_result *GetDcdnKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDcdnKvResponse{}
	_body, _err := client.GetDcdnKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询KV对的值以及TTL信息
//
// @param request - GetDcdnKvDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDcdnKvDetailResponse
func (client *Client) GetDcdnKvDetailWithOptions(request *GetDcdnKvDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDcdnKvDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDcdnKvDetail"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDcdnKvDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询KV对的值以及TTL信息
//
// @param request - GetDcdnKvDetailRequest
//
// @return GetDcdnKvDetailResponse
func (client *Client) GetDcdnKvDetail(request *GetDcdnKvDetailRequest) (_result *GetDcdnKvDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDcdnKvDetailResponse{}
	_body, _err := client.GetDcdnKvDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the KV status by key value.
//
// @param request - GetDcdnKvStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDcdnKvStatusResponse
func (client *Client) GetDcdnKvStatusWithOptions(request *GetDcdnKvStatusRequest, runtime *dara.RuntimeOptions) (_result *GetDcdnKvStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDcdnKvStatus"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDcdnKvStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the KV status by key value.
//
// @param request - GetDcdnKvStatusRequest
//
// @return GetDcdnKvStatusResponse
func (client *Client) GetDcdnKvStatus(request *GetDcdnKvStatusRequest) (_result *GetDcdnKvStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDcdnKvStatusResponse{}
	_body, _err := client.GetDcdnKvStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Traverses the values of keys in a namespace.
//
// @param request - ListDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDcdnKvResponse
func (client *Client) ListDcdnKvWithOptions(request *ListDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *ListDcdnKvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDcdnKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Traverses the values of keys in a namespace.
//
// @param request - ListDcdnKvRequest
//
// @return ListDcdnKvResponse
func (client *Client) ListDcdnKv(request *ListDcdnKvRequest) (_result *ListDcdnKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDcdnKvResponse{}
	_body, _err := client.ListDcdnKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListDcdnRealTimeDeliveryProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDcdnRealTimeDeliveryProjectResponse
func (client *Client) ListDcdnRealTimeDeliveryProjectWithOptions(request *ListDcdnRealTimeDeliveryProjectRequest, runtime *dara.RuntimeOptions) (_result *ListDcdnRealTimeDeliveryProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
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
		Action:      dara.String("ListDcdnRealTimeDeliveryProject"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDcdnRealTimeDeliveryProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListDcdnRealTimeDeliveryProjectRequest
//
// @return ListDcdnRealTimeDeliveryProjectResponse
func (client *Client) ListDcdnRealTimeDeliveryProject(request *ListDcdnRealTimeDeliveryProjectRequest) (_result *ListDcdnRealTimeDeliveryProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDcdnRealTimeDeliveryProjectResponse{}
	_body, _err := client.ListDcdnRealTimeDeliveryProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # A客户定制新增修改域名采样率接口
//
// @param request - ModifyCustomDomainSampleRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomDomainSampleRateResponse
func (client *Client) ModifyCustomDomainSampleRateWithOptions(request *ModifyCustomDomainSampleRateRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomDomainSampleRateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseConfigID) {
		body["BaseConfigID"] = request.BaseConfigID
	}

	if !dara.IsNil(request.DomainNames) {
		body["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.SampleRate) {
		body["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SinkID) {
		body["SinkID"] = request.SinkID
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomDomainSampleRate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomDomainSampleRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A客户定制新增修改域名采样率接口
//
// @param request - ModifyCustomDomainSampleRateRequest
//
// @return ModifyCustomDomainSampleRateResponse
func (client *Client) ModifyCustomDomainSampleRate(request *ModifyCustomDomainSampleRateRequest) (_result *ModifyCustomDomainSampleRateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCustomDomainSampleRateResponse{}
	_body, _err := client.ModifyCustomDomainSampleRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the acceleration region.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ModifyDCdnDomainSchdmByPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDCdnDomainSchdmByPropertyResponse
func (client *Client) ModifyDCdnDomainSchdmByPropertyWithOptions(request *ModifyDCdnDomainSchdmByPropertyRequest, runtime *dara.RuntimeOptions) (_result *ModifyDCdnDomainSchdmByPropertyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Property) {
		query["Property"] = request.Property
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDCdnDomainSchdmByProperty"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDCdnDomainSchdmByPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the acceleration region.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ModifyDCdnDomainSchdmByPropertyRequest
//
// @return ModifyDCdnDomainSchdmByPropertyResponse
func (client *Client) ModifyDCdnDomainSchdmByProperty(request *ModifyDCdnDomainSchdmByPropertyRequest) (_result *ModifyDCdnDomainSchdmByPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDCdnDomainSchdmByPropertyResponse{}
	_body, _err := client.ModifyDCdnDomainSchdmByPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a custom Web Application Firewall (WAF) rule group.
//
// @param request - ModifyDcdnWafGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDcdnWafGroupResponse
func (client *Client) ModifyDcdnWafGroupWithOptions(request *ModifyDcdnWafGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDcdnWafGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Rules) {
		body["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDcdnWafGroup"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDcdnWafGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a custom Web Application Firewall (WAF) rule group.
//
// @param request - ModifyDcdnWafGroupRequest
//
// @return ModifyDcdnWafGroupResponse
func (client *Client) ModifyDcdnWafGroup(request *ModifyDcdnWafGroupRequest) (_result *ModifyDcdnWafGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDcdnWafGroupResponse{}
	_body, _err := client.ModifyDcdnWafGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name or the status of a protection policy.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - ModifyDcdnWafPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDcdnWafPolicyResponse
func (client *Client) ModifyDcdnWafPolicyWithOptions(request *ModifyDcdnWafPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyDcdnWafPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyStatus) {
		body["PolicyStatus"] = request.PolicyStatus
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDcdnWafPolicy"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDcdnWafPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name or the status of a protection policy.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
// @param request - ModifyDcdnWafPolicyRequest
//
// @return ModifyDcdnWafPolicyResponse
func (client *Client) ModifyDcdnWafPolicy(request *ModifyDcdnWafPolicyRequest) (_result *ModifyDcdnWafPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDcdnWafPolicyResponse{}
	_body, _err := client.ModifyDcdnWafPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the accelerated domain names that are bound to a protection policy.
//
// Description:
//
// # Usage notes
//
//   - You can call this operation up to 20 times per second per account.
//
//   - Alibaba Cloud Dynamic Route for CDN (DCDN) supports POST requests.
//
// @param request - ModifyDcdnWafPolicyDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDcdnWafPolicyDomainsResponse
func (client *Client) ModifyDcdnWafPolicyDomainsWithOptions(request *ModifyDcdnWafPolicyDomainsRequest, runtime *dara.RuntimeOptions) (_result *ModifyDcdnWafPolicyDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BindDomains) {
		body["BindDomains"] = request.BindDomains
	}

	if !dara.IsNil(request.Method) {
		body["Method"] = request.Method
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.UnbindDomains) {
		body["UnbindDomains"] = request.UnbindDomains
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDcdnWafPolicyDomains"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDcdnWafPolicyDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the accelerated domain names that are bound to a protection policy.
//
// Description:
//
// # Usage notes
//
//   - You can call this operation up to 20 times per second per account.
//
//   - Alibaba Cloud Dynamic Route for CDN (DCDN) supports POST requests.
//
// @param request - ModifyDcdnWafPolicyDomainsRequest
//
// @return ModifyDcdnWafPolicyDomainsResponse
func (client *Client) ModifyDcdnWafPolicyDomains(request *ModifyDcdnWafPolicyDomainsRequest) (_result *ModifyDcdnWafPolicyDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDcdnWafPolicyDomainsResponse{}
	_body, _err := client.ModifyDcdnWafPolicyDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name, status, or configurations of a protection rule.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
//		- You must configure at least one of the **RuleStatus**, **RuleName*	- and **RuleConfig*	- parameters.
//
// @param request - ModifyDcdnWafRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDcdnWafRuleResponse
func (client *Client) ModifyDcdnWafRuleWithOptions(request *ModifyDcdnWafRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyDcdnWafRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RuleConfig) {
		body["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		body["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleStatus) {
		body["RuleStatus"] = request.RuleStatus
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDcdnWafRule"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDcdnWafRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name, status, or configurations of a protection rule.
//
// Description:
//
//	  You can call this operation up to 20 times per second per account.
//
//		- Alibaba Cloud Dynamic Content Delivery Network (DCDN) supports POST requests.
//
//		- You must configure at least one of the **RuleStatus**, **RuleName*	- and **RuleConfig*	- parameters.
//
// @param request - ModifyDcdnWafRuleRequest
//
// @return ModifyDcdnWafRuleResponse
func (client *Client) ModifyDcdnWafRule(request *ModifyDcdnWafRuleRequest) (_result *ModifyDcdnWafRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDcdnWafRuleResponse{}
	_body, _err := client.ModifyDcdnWafRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates Dynamic Route for CDN (DCDN).
//
// Description:
//
// >
//
//   - DCDN can be activated only once per Alibaba Cloud account. The Alibaba Cloud account must pass real-name verification.
//
//   - You can call this operation up to five times per second per user.
//
// @param request - OpenDcdnServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDcdnServiceResponse
func (client *Client) OpenDcdnServiceWithOptions(request *OpenDcdnServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenDcdnServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillType) {
		query["BillType"] = request.BillType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WebsocketBillType) {
		query["WebsocketBillType"] = request.WebsocketBillType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenDcdnService"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenDcdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates Dynamic Route for CDN (DCDN).
//
// Description:
//
// >
//
//   - DCDN can be activated only once per Alibaba Cloud account. The Alibaba Cloud account must pass real-name verification.
//
//   - You can call this operation up to five times per second per user.
//
// @param request - OpenDcdnServiceRequest
//
// @return OpenDcdnServiceResponse
func (client *Client) OpenDcdnService(request *OpenDcdnServiceRequest) (_result *OpenDcdnServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenDcdnServiceResponse{}
	_body, _err := client.OpenDcdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Prefetches content from origin servers to points of presence (POPs). This reduces workloads on origin servers because users can hit cache upon their first visits.
//
// Description:
//
//	  You can call the [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html) operation to refresh content and call the [PreloadDcdnObjectCaches](https://help.aliyun.com/document_detail/130636.html) operation to prefetch content.
//
//		- Dynamic Content Delivery Network (DCDN) supports POST requests in which parameters are sent as a form.
//
//		- By default, each Alibaba Cloud account can submit up to 1,000 URLs per day. If the daily peak bandwidth value of your workloads exceeds 200 Mbit/s, you can [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to increase your daily quota. Alibaba Cloud reviews your application and then increases the quota accordingly.
//
//		- You can specify up to 100 URLs to prefetch.
//
//		- The prefetch queue of each Alibaba Cloud account can contain up to 100,000 URLs. DCDN executes prefetch tasks based on the time at which you submit the URLs.
//
//		- You can call this operation up to 15 times per second per account.
//
// ## Description
//
//   - After a refresh task is submitted and completed, the POPs immediately start to retrieve resources from the origin server. Therefore, a large number of refresh tasks cause a large number of concurrent download tasks. This increases the number of requests that are redirected to the origin server. The back-to-origin routing process consumes more bandwidth resources and the origin server may be overwhelmed.
//
//   - The time required for a prefetch task to complete is proportional to the size of the prefetched file. In actual practice, most prefetch tasks require 5 to 30 minutes to complete. A task with a smaller average file size requires less time.
//
//   - To allow RAM users to perform this operation, you need to first grant them the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/445051.html).
//
// @param request - PreloadDcdnObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreloadDcdnObjectCachesResponse
func (client *Client) PreloadDcdnObjectCachesWithOptions(request *PreloadDcdnObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *PreloadDcdnObjectCachesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.L2Preload) {
		query["L2Preload"] = request.L2Preload
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryHashkey) {
		query["QueryHashkey"] = request.QueryHashkey
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WithHeader) {
		query["WithHeader"] = request.WithHeader
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreloadDcdnObjectCaches"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreloadDcdnObjectCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prefetches content from origin servers to points of presence (POPs). This reduces workloads on origin servers because users can hit cache upon their first visits.
//
// Description:
//
//	  You can call the [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html) operation to refresh content and call the [PreloadDcdnObjectCaches](https://help.aliyun.com/document_detail/130636.html) operation to prefetch content.
//
//		- Dynamic Content Delivery Network (DCDN) supports POST requests in which parameters are sent as a form.
//
//		- By default, each Alibaba Cloud account can submit up to 1,000 URLs per day. If the daily peak bandwidth value of your workloads exceeds 200 Mbit/s, you can [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to increase your daily quota. Alibaba Cloud reviews your application and then increases the quota accordingly.
//
//		- You can specify up to 100 URLs to prefetch.
//
//		- The prefetch queue of each Alibaba Cloud account can contain up to 100,000 URLs. DCDN executes prefetch tasks based on the time at which you submit the URLs.
//
//		- You can call this operation up to 15 times per second per account.
//
// ## Description
//
//   - After a refresh task is submitted and completed, the POPs immediately start to retrieve resources from the origin server. Therefore, a large number of refresh tasks cause a large number of concurrent download tasks. This increases the number of requests that are redirected to the origin server. The back-to-origin routing process consumes more bandwidth resources and the origin server may be overwhelmed.
//
//   - The time required for a prefetch task to complete is proportional to the size of the prefetched file. In actual practice, most prefetch tasks require 5 to 30 minutes to complete. A task with a smaller average file size requires less time.
//
//   - To allow RAM users to perform this operation, you need to first grant them the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/445051.html).
//
// @param request - PreloadDcdnObjectCachesRequest
//
// @return PreloadDcdnObjectCachesResponse
func (client *Client) PreloadDcdnObjectCaches(request *PreloadDcdnObjectCachesRequest) (_result *PreloadDcdnObjectCachesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreloadDcdnObjectCachesResponse{}
	_body, _err := client.PreloadDcdnObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes the configurations of an accelerated domain name from the staging environment to the production environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - PublishDcdnStagingConfigToProductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishDcdnStagingConfigToProductionResponse
func (client *Client) PublishDcdnStagingConfigToProductionWithOptions(request *PublishDcdnStagingConfigToProductionRequest, runtime *dara.RuntimeOptions) (_result *PublishDcdnStagingConfigToProductionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishDcdnStagingConfigToProduction"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishDcdnStagingConfigToProductionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes the configurations of an accelerated domain name from the staging environment to the production environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - PublishDcdnStagingConfigToProductionRequest
//
// @return PublishDcdnStagingConfigToProductionResponse
func (client *Client) PublishDcdnStagingConfigToProduction(request *PublishDcdnStagingConfigToProductionRequest) (_result *PublishDcdnStagingConfigToProductionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishDcdnStagingConfigToProductionResponse{}
	_body, _err := client.PublishDcdnStagingConfigToProductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a specified version of routine code to an environment.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param tmpReq - PublishRoutineCodeRevisionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRoutineCodeRevisionResponse
func (client *Client) PublishRoutineCodeRevisionWithOptions(tmpReq *PublishRoutineCodeRevisionRequest, runtime *dara.RuntimeOptions) (_result *PublishRoutineCodeRevisionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &PublishRoutineCodeRevisionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Envs) {
		request.EnvsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Envs, dara.String("Envs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvsShrink) {
		body["Envs"] = request.EnvsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SelectCodeRevision) {
		body["SelectCodeRevision"] = request.SelectCodeRevision
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishRoutineCodeRevision"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishRoutineCodeRevisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a specified version of routine code to an environment.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - PublishRoutineCodeRevisionRequest
//
// @return PublishRoutineCodeRevisionResponse
func (client *Client) PublishRoutineCodeRevision(request *PublishRoutineCodeRevisionRequest) (_result *PublishRoutineCodeRevisionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishRoutineCodeRevisionResponse{}
	_body, _err := client.PublishRoutineCodeRevisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets key-value pairs in a namespace.
//
// @param request - PutDcdnKvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDcdnKvResponse
func (client *Client) PutDcdnKvWithOptions(request *PutDcdnKvRequest, runtime *dara.RuntimeOptions) (_result *PutDcdnKvResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("PutDcdnKv"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDcdnKvResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets key-value pairs in a namespace.
//
// @param request - PutDcdnKvRequest
//
// @return PutDcdnKvResponse
func (client *Client) PutDcdnKv(request *PutDcdnKvRequest) (_result *PutDcdnKvResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutDcdnKvResponse{}
	_body, _err := client.PutDcdnKvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds namespaces to your account.
//
// @param request - PutDcdnKvNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDcdnKvNamespaceResponse
func (client *Client) PutDcdnKvNamespaceWithOptions(request *PutDcdnKvNamespaceRequest, runtime *dara.RuntimeOptions) (_result *PutDcdnKvNamespaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("PutDcdnKvNamespace"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDcdnKvNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds namespaces to your account.
//
// @param request - PutDcdnKvNamespaceRequest
//
// @return PutDcdnKvNamespaceResponse
func (client *Client) PutDcdnKvNamespace(request *PutDcdnKvNamespaceRequest) (_result *PutDcdnKvNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutDcdnKvNamespaceResponse{}
	_body, _err := client.PutDcdnKvNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置Namespace的key-value对，支持最大25M的请求体
//
// @param request - PutDcdnKvWithHighCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDcdnKvWithHighCapacityResponse
func (client *Client) PutDcdnKvWithHighCapacityWithOptions(request *PutDcdnKvWithHighCapacityRequest, runtime *dara.RuntimeOptions) (_result *PutDcdnKvWithHighCapacityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("PutDcdnKvWithHighCapacity"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDcdnKvWithHighCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置Namespace的key-value对，支持最大25M的请求体
//
// @param request - PutDcdnKvWithHighCapacityRequest
//
// @return PutDcdnKvWithHighCapacityResponse
func (client *Client) PutDcdnKvWithHighCapacity(request *PutDcdnKvWithHighCapacityRequest) (_result *PutDcdnKvWithHighCapacityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutDcdnKvWithHighCapacityResponse{}
	_body, _err := client.PutDcdnKvWithHighCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定缓存tag刷新节点上的文件内容
//
// @param request - RefreshDcdnObjectCacheByCacheTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshDcdnObjectCacheByCacheTagResponse
func (client *Client) RefreshDcdnObjectCacheByCacheTagWithOptions(request *RefreshDcdnObjectCacheByCacheTagRequest, runtime *dara.RuntimeOptions) (_result *RefreshDcdnObjectCacheByCacheTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheTag) {
		query["CacheTag"] = request.CacheTag
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshDcdnObjectCacheByCacheTag"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshDcdnObjectCacheByCacheTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定缓存tag刷新节点上的文件内容
//
// @param request - RefreshDcdnObjectCacheByCacheTagRequest
//
// @return RefreshDcdnObjectCacheByCacheTagResponse
func (client *Client) RefreshDcdnObjectCacheByCacheTag(request *RefreshDcdnObjectCacheByCacheTagRequest) (_result *RefreshDcdnObjectCacheByCacheTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshDcdnObjectCacheByCacheTagResponse{}
	_body, _err := client.RefreshDcdnObjectCacheByCacheTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Refreshes specified objects on points of presence (POPs). The objects can be included in the content of files or URLs. You can refresh multiple URLs in each request.
//
// Description:
//
//	  Dynamic Content Delivery Network (DCDN) supports POST requests in which parameters are sent as a form.
//
//		- You can call the [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html) operation to purge content and call the [PreloadDcdnObjectCaches](https://help.aliyun.com/document_detail/130636.html) operation to prefetch content.
//
//		- By default, each Alibaba Cloud account can purge content from a maximum of 10,000 URLs and 100 directories including subdirectories per day. If the daily peak bandwidth of your Alibaba Cloud account exceeds 200 Mbit/s, [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to request a quota increase. Alibaba Cloud determines whether to approve your application based on your workloads.
//
//		- You can specify up to 1,000 URLs or 100 directories that you want to purge in each request.
//
//		- You can specify up to 1,000 URLs that you want to purge per minute for each domain name.
//
//		- You can call this operation up to 30 times per second per account.
//
// #### [](#)Precautions
//
//   - After a purge task is completed, your resources that are cached on points of presence (POPs) are removed. When a POP receives a request for your resources, the request is redirected to the origin server to retrieve the resources. Then, the resources are returned to the client and cached on POPs. If you frequently run purge tasks, more requests are redirected to the origin server for resources. This results in high bandwidth costs and more loads on the origin server.
//
//   - A purge task takes effect 5 to 6 minutes after being submitted. If the resource you want to purge has a TTL of less than 5 minutes, you wait for it to expire instead of manually running a purge task.
//
//   - To allow RAM users to perform this operation, you need to first grant them the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/445051.html).
//
// @param request - RefreshDcdnObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshDcdnObjectCachesResponse
func (client *Client) RefreshDcdnObjectCachesWithOptions(request *RefreshDcdnObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *RefreshDcdnObjectCachesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

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
		Action:      dara.String("RefreshDcdnObjectCaches"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshDcdnObjectCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes specified objects on points of presence (POPs). The objects can be included in the content of files or URLs. You can refresh multiple URLs in each request.
//
// Description:
//
//	  Dynamic Content Delivery Network (DCDN) supports POST requests in which parameters are sent as a form.
//
//		- You can call the [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html) operation to purge content and call the [PreloadDcdnObjectCaches](https://help.aliyun.com/document_detail/130636.html) operation to prefetch content.
//
//		- By default, each Alibaba Cloud account can purge content from a maximum of 10,000 URLs and 100 directories including subdirectories per day. If the daily peak bandwidth of your Alibaba Cloud account exceeds 200 Mbit/s, [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to request a quota increase. Alibaba Cloud determines whether to approve your application based on your workloads.
//
//		- You can specify up to 1,000 URLs or 100 directories that you want to purge in each request.
//
//		- You can specify up to 1,000 URLs that you want to purge per minute for each domain name.
//
//		- You can call this operation up to 30 times per second per account.
//
// #### [](#)Precautions
//
//   - After a purge task is completed, your resources that are cached on points of presence (POPs) are removed. When a POP receives a request for your resources, the request is redirected to the origin server to retrieve the resources. Then, the resources are returned to the client and cached on POPs. If you frequently run purge tasks, more requests are redirected to the origin server for resources. This results in high bandwidth costs and more loads on the origin server.
//
//   - A purge task takes effect 5 to 6 minutes after being submitted. If the resource you want to purge has a TTL of less than 5 minutes, you wait for it to expire instead of manually running a purge task.
//
//   - To allow RAM users to perform this operation, you need to first grant them the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/445051.html).
//
// @param request - RefreshDcdnObjectCachesRequest
//
// @return RefreshDcdnObjectCachesResponse
func (client *Client) RefreshDcdnObjectCaches(request *RefreshDcdnObjectCachesRequest) (_result *RefreshDcdnObjectCachesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshDcdnObjectCachesResponse{}
	_body, _err := client.RefreshDcdnObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Refreshes the cache that is written by calling the cache operation of EdgeRoutine. You can refresh multiple URLs in each request.
//
// Description:
//
// > 	- Dynamic Content Delivery Network (DCDN) supports POST requests in which parameters are sent as a form.
//
// > 	- Related operation: [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html).
//
// > 	- By default, each Alibaba Cloud account can purge content from a maxim> um of 10,000 URLs and 100 directories including subdirectories per day.
//
// > 	- You can specify up to 1,000 URLs or 100 directories that you want to purge in each request.
//
// > 	- You can specify up to 1,000 URLs that you want to purge per minute for each domain name.
//
// > 	- You can call this operation up to 30 times per second per account.
//
// @param request - RefreshErObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshErObjectCachesResponse
func (client *Client) RefreshErObjectCachesWithOptions(request *RefreshErObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *RefreshErObjectCachesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.MergeDomainName) {
		query["MergeDomainName"] = request.MergeDomainName
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.RoutineId) {
		query["RoutineId"] = request.RoutineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshErObjectCaches"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshErObjectCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes the cache that is written by calling the cache operation of EdgeRoutine. You can refresh multiple URLs in each request.
//
// Description:
//
// > 	- Dynamic Content Delivery Network (DCDN) supports POST requests in which parameters are sent as a form.
//
// > 	- Related operation: [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html).
//
// > 	- By default, each Alibaba Cloud account can purge content from a maxim> um of 10,000 URLs and 100 directories including subdirectories per day.
//
// > 	- You can specify up to 1,000 URLs or 100 directories that you want to purge in each request.
//
// > 	- You can specify up to 1,000 URLs that you want to purge per minute for each domain name.
//
// > 	- You can call this operation up to 30 times per second per account.
//
// @param request - RefreshErObjectCachesRequest
//
// @return RefreshErObjectCachesResponse
func (client *Client) RefreshErObjectCaches(request *RefreshErObjectCachesRequest) (_result *RefreshErObjectCachesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshErObjectCachesResponse{}
	_body, _err := client.RefreshErObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back the configurations of an accelerated domain name from the staging environment to the production environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - RollbackDcdnStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackDcdnStagingConfigResponse
func (client *Client) RollbackDcdnStagingConfigWithOptions(request *RollbackDcdnStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *RollbackDcdnStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackDcdnStagingConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackDcdnStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back the configurations of an accelerated domain name from the staging environment to the production environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - RollbackDcdnStagingConfigRequest
//
// @return RollbackDcdnStagingConfigResponse
func (client *Client) RollbackDcdnStagingConfig(request *RollbackDcdnStagingConfigRequest) (_result *RollbackDcdnStagingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackDcdnStagingConfigResponse{}
	_body, _err := client.RollbackDcdnStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures an SSL certificate for a specified domain name.
//
// @param request - SetDcdnDomainCSRCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnDomainCSRCertificateResponse
func (client *Client) SetDcdnDomainCSRCertificateWithOptions(request *SetDcdnDomainCSRCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnDomainCSRCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ServerCertificate) {
		query["ServerCertificate"] = request.ServerCertificate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnDomainCSRCertificate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnDomainCSRCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures an SSL certificate for a specified domain name.
//
// @param request - SetDcdnDomainCSRCertificateRequest
//
// @return SetDcdnDomainCSRCertificateResponse
func (client *Client) SetDcdnDomainCSRCertificate(request *SetDcdnDomainCSRCertificateRequest) (_result *SetDcdnDomainCSRCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDcdnDomainCSRCertificateResponse{}
	_body, _err := client.SetDcdnDomainCSRCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the ShangMi (SM) certificate for a domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetDcdnDomainSMCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnDomainSMCertificateResponse
func (client *Client) SetDcdnDomainSMCertificateWithOptions(request *SetDcdnDomainSMCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnDomainSMCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnDomainSMCertificate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnDomainSMCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the ShangMi (SM) certificate for a domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetDcdnDomainSMCertificateRequest
//
// @return SetDcdnDomainSMCertificateResponse
func (client *Client) SetDcdnDomainSMCertificate(request *SetDcdnDomainSMCertificateRequest) (_result *SetDcdnDomainSMCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDcdnDomainSMCertificateResponse{}
	_body, _err := client.SetDcdnDomainSMCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the SSL certificate for a domain name and updates certificate details.
//
// @param request - SetDcdnDomainSSLCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnDomainSSLCertificateResponse
func (client *Client) SetDcdnDomainSSLCertificateWithOptions(request *SetDcdnDomainSSLCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnDomainSSLCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertRegion) {
		query["CertRegion"] = request.CertRegion
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLPri) {
		query["SSLPri"] = request.SSLPri
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnDomainSSLCertificate"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnDomainSSLCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the SSL certificate for a domain name and updates certificate details.
//
// @param request - SetDcdnDomainSSLCertificateRequest
//
// @return SetDcdnDomainSSLCertificateResponse
func (client *Client) SetDcdnDomainSSLCertificate(request *SetDcdnDomainSSLCertificateRequest) (_result *SetDcdnDomainSSLCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDcdnDomainSSLCertificateResponse{}
	_body, _err := client.SetDcdnDomainSSLCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets or modifies the domain name configuration in the canary release environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetDcdnDomainStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnDomainStagingConfigResponse
func (client *Client) SetDcdnDomainStagingConfigWithOptions(request *SetDcdnDomainStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnDomainStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnDomainStagingConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnDomainStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets or modifies the domain name configuration in the canary release environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetDcdnDomainStagingConfigRequest
//
// @return SetDcdnDomainStagingConfigResponse
func (client *Client) SetDcdnDomainStagingConfig(request *SetDcdnDomainStagingConfigRequest) (_result *SetDcdnDomainStagingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDcdnDomainStagingConfigResponse{}
	_body, _err := client.SetDcdnDomainStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Blocks or unblocks IP addresses or CIDR blocks.
//
// Description:
//
// >  You can call this operation to block or unblock a large number of IP addresses or CIDR blocks. You can block or unblock up to 1,000 IP addresses or CIDR blocks in a request.
//
// @param request - SetDcdnFullDomainsBlockIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnFullDomainsBlockIPResponse
func (client *Client) SetDcdnFullDomainsBlockIPWithOptions(request *SetDcdnFullDomainsBlockIPRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnFullDomainsBlockIPResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BlockInterval) {
		body["BlockInterval"] = request.BlockInterval
	}

	if !dara.IsNil(request.IPList) {
		body["IPList"] = request.IPList
	}

	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.UpdateType) {
		body["UpdateType"] = request.UpdateType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDcdnFullDomainsBlockIP"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnFullDomainsBlockIPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Blocks or unblocks IP addresses or CIDR blocks.
//
// Description:
//
// >  You can call this operation to block or unblock a large number of IP addresses or CIDR blocks. You can block or unblock up to 1,000 IP addresses or CIDR blocks in a request.
//
// @param request - SetDcdnFullDomainsBlockIPRequest
//
// @return SetDcdnFullDomainsBlockIPResponse
func (client *Client) SetDcdnFullDomainsBlockIP(request *SetDcdnFullDomainsBlockIPRequest) (_result *SetDcdnFullDomainsBlockIPResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDcdnFullDomainsBlockIPResponse{}
	_body, _err := client.SetDcdnFullDomainsBlockIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures features for a user.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetDcdnUserConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDcdnUserConfigResponse
func (client *Client) SetDcdnUserConfigWithOptions(request *SetDcdnUserConfigRequest, runtime *dara.RuntimeOptions) (_result *SetDcdnUserConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Configs) {
		query["Configs"] = request.Configs
	}

	if !dara.IsNil(request.FunctionId) {
		query["FunctionId"] = request.FunctionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

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
		Action:      dara.String("SetDcdnUserConfig"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDcdnUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures features for a user.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetDcdnUserConfigRequest
//
// @return SetDcdnUserConfigResponse
func (client *Client) SetDcdnUserConfig(request *SetDcdnUserConfigRequest) (_result *SetDcdnUserConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDcdnUserConfigResponse{}
	_body, _err := client.SetDcdnUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a subdomain for a routine.
//
// Description:
//
// >
//
//   - Each subdomain is globally unique. Resource Access Management (RAM) users cannot create duplicate subdomains.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param tmpReq - SetRoutineSubdomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRoutineSubdomainResponse
func (client *Client) SetRoutineSubdomainWithOptions(tmpReq *SetRoutineSubdomainRequest, runtime *dara.RuntimeOptions) (_result *SetRoutineSubdomainResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetRoutineSubdomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Subdomains) {
		request.SubdomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Subdomains, dara.String("Subdomains"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SubdomainsShrink) {
		body["Subdomains"] = request.SubdomainsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetRoutineSubdomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetRoutineSubdomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a subdomain for a routine.
//
// Description:
//
// >
//
//   - Each subdomain is globally unique. Resource Access Management (RAM) users cannot create duplicate subdomains.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - SetRoutineSubdomainRequest
//
// @return SetRoutineSubdomainResponse
func (client *Client) SetRoutineSubdomain(request *SetRoutineSubdomainRequest) (_result *SetRoutineSubdomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetRoutineSubdomainResponse{}
	_body, _err := client.SetRoutineSubdomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a specified disabled accelerated domain. After the accelerated domain is enabled, the value of the DomainStatus parameter changes to Online for the domain.
//
// Description:
//
// >
//
//   - If an accelerated domain name is in invalid state or your account has an overdue payment, the accelerated domain name cannot be enabled.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - StartDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDcdnDomainResponse
func (client *Client) StartDcdnDomainWithOptions(request *StartDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *StartDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("StartDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a specified disabled accelerated domain. After the accelerated domain is enabled, the value of the DomainStatus parameter changes to Online for the domain.
//
// Description:
//
// >
//
//   - If an accelerated domain name is in invalid state or your account has an overdue payment, the accelerated domain name cannot be enabled.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - StartDcdnDomainRequest
//
// @return StartDcdnDomainResponse
func (client *Client) StartDcdnDomain(request *StartDcdnDomainRequest) (_result *StartDcdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDcdnDomainResponse{}
	_body, _err := client.StartDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables one or more accelerated domain names. After the accelerated domain names are enabled, the value of the DomainStatus parameter for the domain names changes to Online.
//
// Description:
//
//	  If an accelerated domain name is in invalid state or your account has an overdue payment, the accelerated domain name cannot be enabled.
//
//		- You can call this operation up to 20 times per second per account.
//
// @param request - StartDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDcdnIpaDomainResponse
func (client *Client) StartDcdnIpaDomainWithOptions(request *StartDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *StartDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("StartDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables one or more accelerated domain names. After the accelerated domain names are enabled, the value of the DomainStatus parameter for the domain names changes to Online.
//
// Description:
//
//	  If an accelerated domain name is in invalid state or your account has an overdue payment, the accelerated domain name cannot be enabled.
//
//		- You can call this operation up to 20 times per second per account.
//
// @param request - StartDcdnIpaDomainRequest
//
// @return StartDcdnIpaDomainResponse
func (client *Client) StartDcdnIpaDomain(request *StartDcdnIpaDomainRequest) (_result *StartDcdnIpaDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDcdnIpaDomainResponse{}
	_body, _err := client.StartDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a specified accelerated domain. After the accelerated domain is disabled,
//
//	the value of the DomainStatus parameter changes to Offline for the domain.
//
// Description:
//
// >
//
//   - After an accelerated domain is disabled, Dynamic Content Delivery Network (DCDN) retains its information and routes all the requests that are destined for the accelerated domain to the origin server.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - StopDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDcdnDomainResponse
func (client *Client) StopDcdnDomainWithOptions(request *StopDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *StopDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("StopDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a specified accelerated domain. After the accelerated domain is disabled,
//
//	the value of the DomainStatus parameter changes to Offline for the domain.
//
// Description:
//
// >
//
//   - After an accelerated domain is disabled, Dynamic Content Delivery Network (DCDN) retains its information and routes all the requests that are destined for the accelerated domain to the origin server.
//
//   - You can call this operation up to 30 times per second per account.
//
// @param request - StopDcdnDomainRequest
//
// @return StopDcdnDomainResponse
func (client *Client) StopDcdnDomain(request *StopDcdnDomainRequest) (_result *StopDcdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDcdnDomainResponse{}
	_body, _err := client.StopDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an accelerated domain name. After an accelerated domain name is disabled, the value of the DomainStatus parameter for the domain name changes to Offline.
//
// Description:
//
// >
//
//   - If you disable an accelerated domain, the configurations of the accelerated domain are still retained. The system automatically forwards all the requests that are destined for this domain to the origin.
//
//   - You can call this operation up to 20 times per second per account.
//
// @param request - StopDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDcdnIpaDomainResponse
func (client *Client) StopDcdnIpaDomainWithOptions(request *StopDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *StopDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

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
		Action:      dara.String("StopDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an accelerated domain name. After an accelerated domain name is disabled, the value of the DomainStatus parameter for the domain name changes to Offline.
//
// Description:
//
// >
//
//   - If you disable an accelerated domain, the configurations of the accelerated domain are still retained. The system automatically forwards all the requests that are destined for this domain to the origin.
//
//   - You can call this operation up to 20 times per second per account.
//
// @param request - StopDcdnIpaDomainRequest
//
// @return StopDcdnIpaDomainResponse
func (client *Client) StopDcdnIpaDomain(request *StopDcdnIpaDomainRequest) (_result *StopDcdnIpaDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDcdnIpaDomainResponse{}
	_body, _err := client.StopDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds one or more tags to a resource.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - TagDcdnResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagDcdnResourcesResponse
func (client *Client) TagDcdnResourcesWithOptions(request *TagDcdnResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagDcdnResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("TagDcdnResources"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagDcdnResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more tags to a resource.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - TagDcdnResourcesRequest
//
// @return TagDcdnResourcesResponse
func (client *Client) TagDcdnResources(request *TagDcdnResourcesRequest) (_result *TagDcdnResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagDcdnResourcesResponse{}
	_body, _err := client.TagDcdnResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes one or more tags from a resource.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UntagDcdnResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagDcdnResourcesResponse
func (client *Client) UntagDcdnResourcesWithOptions(request *UntagDcdnResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagDcdnResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
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
		Action:      dara.String("UntagDcdnResources"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagDcdnResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes one or more tags from a resource.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UntagDcdnResourcesRequest
//
// @return UntagDcdnResourcesResponse
func (client *Client) UntagDcdnResources(request *UntagDcdnResourcesRequest) (_result *UntagDcdnResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagDcdnResourcesResponse{}
	_body, _err := client.UntagDcdnResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a tracking task by task ID.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateDcdnDeliverTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnDeliverTaskResponse
func (client *Client) UpdateDcdnDeliverTaskWithOptions(request *UpdateDcdnDeliverTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnDeliverTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Deliver) {
		body["Deliver"] = request.Deliver
	}

	if !dara.IsNil(request.DeliverId) {
		body["DeliverId"] = request.DeliverId
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Reports) {
		body["Reports"] = request.Reports
	}

	if !dara.IsNil(request.Schedule) {
		body["Schedule"] = request.Schedule
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnDeliverTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnDeliverTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a tracking task by task ID.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateDcdnDeliverTaskRequest
//
// @return UpdateDcdnDeliverTaskResponse
func (client *Client) UpdateDcdnDeliverTask(request *UpdateDcdnDeliverTaskRequest) (_result *UpdateDcdnDeliverTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDcdnDeliverTaskResponse{}
	_body, _err := client.UpdateDcdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - UpdateDcdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnDomainResponse
func (client *Client) UpdateDcdnDomainWithOptions(request *UpdateDcdnDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - UpdateDcdnDomainRequest
//
// @return UpdateDcdnDomainResponse
func (client *Client) UpdateDcdnDomain(request *UpdateDcdnDomainRequest) (_result *UpdateDcdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDcdnDomainResponse{}
	_body, _err := client.UpdateDcdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a domain name that is accelerated by IP Application Accelerator (IPA).
//
// Description:
//
// >  You can call this operation up to 20 times per second per account.
//
// @param request - UpdateDcdnIpaDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnIpaDomainResponse
func (client *Client) UpdateDcdnIpaDomainWithOptions(request *UpdateDcdnIpaDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnIpaDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnIpaDomain"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnIpaDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a domain name that is accelerated by IP Application Accelerator (IPA).
//
// Description:
//
// >  You can call this operation up to 20 times per second per account.
//
// @param request - UpdateDcdnIpaDomainRequest
//
// @return UpdateDcdnIpaDomainResponse
func (client *Client) UpdateDcdnIpaDomain(request *UpdateDcdnIpaDomainRequest) (_result *UpdateDcdnIpaDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDcdnIpaDomainResponse{}
	_body, _err := client.UpdateDcdnIpaDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UpdateDcdnSLSRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnSLSRealtimeLogDeliveryResponse
func (client *Client) UpdateDcdnSLSRealtimeLogDeliveryWithOptions(request *UpdateDcdnSLSRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataCenter) {
		body["DataCenter"] = request.DataCenter
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SLSLogStore) {
		body["SLSLogStore"] = request.SLSLogStore
	}

	if !dara.IsNil(request.SLSProject) {
		body["SLSProject"] = request.SLSProject
	}

	if !dara.IsNil(request.SLSRegion) {
		body["SLSRegion"] = request.SLSRegion
	}

	if !dara.IsNil(request.SamplingRate) {
		body["SamplingRate"] = request.SamplingRate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnSLSRealtimeLogDelivery"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a real-time log delivery project.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UpdateDcdnSLSRealtimeLogDeliveryRequest
//
// @return UpdateDcdnSLSRealtimeLogDeliveryResponse
func (client *Client) UpdateDcdnSLSRealtimeLogDelivery(request *UpdateDcdnSLSRealtimeLogDeliveryRequest) (_result *UpdateDcdnSLSRealtimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDcdnSLSRealtimeLogDeliveryResponse{}
	_body, _err := client.UpdateDcdnSLSRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates one or more operations reports.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateDcdnSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnSubTaskResponse
func (client *Client) UpdateDcdnSubTaskWithOptions(request *UpdateDcdnSubTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnSubTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ReportIds) {
		body["ReportIds"] = request.ReportIds
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnSubTask"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates one or more operations reports.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateDcdnSubTaskRequest
//
// @return UpdateDcdnSubTaskResponse
func (client *Client) UpdateDcdnSubTask(request *UpdateDcdnSubTaskRequest) (_result *UpdateDcdnSubTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDcdnSubTaskResponse{}
	_body, _err := client.UpdateDcdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the fields in real-time log entries.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UpdateDcdnUserRealTimeDeliveryFieldRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDcdnUserRealTimeDeliveryFieldResponse
func (client *Client) UpdateDcdnUserRealTimeDeliveryFieldWithOptions(request *UpdateDcdnUserRealTimeDeliveryFieldRequest, runtime *dara.RuntimeOptions) (_result *UpdateDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.Fields) {
		query["Fields"] = request.Fields
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDcdnUserRealTimeDeliveryField"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the fields in real-time log entries.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UpdateDcdnUserRealTimeDeliveryFieldRequest
//
// @return UpdateDcdnUserRealTimeDeliveryFieldResponse
func (client *Client) UpdateDcdnUserRealTimeDeliveryField(request *UpdateDcdnUserRealTimeDeliveryFieldRequest) (_result *UpdateDcdnUserRealTimeDeliveryFieldResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDcdnUserRealTimeDeliveryFieldResponse{}
	_body, _err := client.UpdateDcdnUserRealTimeDeliveryFieldWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads code to EdgeRoutine (ER).
//
// Description:
//
// >
//
//   - Each time you submit code, a version of the code is generated. You can manage and publish code by version.
//
//   - Each routine can retain at most 10 versions. If the upper limit is reached, you must call the DeleteRoutineCodeRevision operation to manually delete versions that are no longer needed before new versions can be saved.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - UploadRoutineCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRoutineCodeResponse
func (client *Client) UploadRoutineCodeWithOptions(request *UploadRoutineCodeRequest, runtime *dara.RuntimeOptions) (_result *UploadRoutineCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("UploadRoutineCode"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadRoutineCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads code to EdgeRoutine (ER).
//
// Description:
//
// >
//
//   - Each time you submit code, a version of the code is generated. You can manage and publish code by version.
//
//   - Each routine can retain at most 10 versions. If the upper limit is reached, you must call the DeleteRoutineCodeRevision operation to manually delete versions that are no longer needed before new versions can be saved.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - UploadRoutineCodeRequest
//
// @return UploadRoutineCodeResponse
func (client *Client) UploadRoutineCode(request *UploadRoutineCodeRequest) (_result *UploadRoutineCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadRoutineCodeResponse{}
	_body, _err := client.UploadRoutineCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads code to a routine for testing.
//
// Description:
//
// >
//
//   - Each time you upload code to a routine, a version is generated. The number of versions is counted by CodeRev. The uploaded code is used only for testing.
//
//   - The code is automatically published to a staging environment.
//
//   - Each routine can retain at most 10 versions. If the upper limit is reached, you need to call the DeleteRoutineCodeRevision operation to manually delete versions that are no longer needed before new versions can be saved.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - UploadStagingRoutineCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadStagingRoutineCodeResponse
func (client *Client) UploadStagingRoutineCodeWithOptions(request *UploadStagingRoutineCodeRequest, runtime *dara.RuntimeOptions) (_result *UploadStagingRoutineCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Action:      dara.String("UploadStagingRoutineCode"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadStagingRoutineCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads code to a routine for testing.
//
// Description:
//
// >
//
//   - Each time you upload code to a routine, a version is generated. The number of versions is counted by CodeRev. The uploaded code is used only for testing.
//
//   - The code is automatically published to a staging environment.
//
//   - Each routine can retain at most 10 versions. If the upper limit is reached, you need to call the DeleteRoutineCodeRevision operation to manually delete versions that are no longer needed before new versions can be saved.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - UploadStagingRoutineCodeRequest
//
// @return UploadStagingRoutineCodeResponse
func (client *Client) UploadStagingRoutineCode(request *UploadStagingRoutineCodeRequest) (_result *UploadStagingRoutineCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadStagingRoutineCodeResponse{}
	_body, _err := client.UploadStagingRoutineCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the ownership of a domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - VerifyDcdnDomainOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyDcdnDomainOwnerResponse
func (client *Client) VerifyDcdnDomainOwnerWithOptions(request *VerifyDcdnDomainOwnerRequest, runtime *dara.RuntimeOptions) (_result *VerifyDcdnDomainOwnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyDcdnDomainOwner"),
		Version:     dara.String("2018-01-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyDcdnDomainOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the ownership of a domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - VerifyDcdnDomainOwnerRequest
//
// @return VerifyDcdnDomainOwnerResponse
func (client *Client) VerifyDcdnDomainOwner(request *VerifyDcdnDomainOwnerRequest) (_result *VerifyDcdnDomainOwnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyDcdnDomainOwnerResponse{}
	_body, _err := client.VerifyDcdnDomainOwnerWithOptions(request, runtime)
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
