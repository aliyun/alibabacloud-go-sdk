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
		"cn-hangzhou": dara.String("httpdns-api.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("httpdns"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a domain name based on the specified parameters.
//
// @param request - AddDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainResponse
func (client *Client) AddDomainWithOptions(request *AddDomainRequest, runtime *dara.RuntimeOptions) (_result *AddDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomain"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name based on the specified parameters.
//
// @param request - AddDomainRequest
//
// @return AddDomainResponse
func (client *Client) AddDomain(request *AddDomainRequest) (_result *AddDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDomainResponse{}
	_body, _err := client.AddDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom resolution policy for a domain name connected to the current account. This operation supports the LINE mode and the IP_RANGE mode. IP address ranges can include IPv4 and IPv6 addresses. When a match occurs, A or AAAA records are returned based on the rules. By default, the policy is created in the disabled state. If Enable is set to true, the policy is enabled upon creation, but actual resolution still requires waiting for configuration propagation. The returned Id is used for subsequent queries, updates, enabling, disabling, and deletion. Create a disabled policy first, and then enable it after querying and verifying it. The policy type and the bound domain name cannot be modified after creation. If the response times out or is lost, query the list first to verify and avoid duplicate creation.
//
// @param tmpReq - CreateSdnsRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSdnsRuleCategoryResponse
func (client *Client) CreateSdnsRuleCategoryWithOptions(tmpReq *CreateSdnsRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *CreateSdnsRuleCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSdnsRuleCategoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClientIpRanges) {
		request.ClientIpRangesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIpRanges, dara.String("ClientIpRanges"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Paths) {
		request.PathsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Paths, dara.String("Paths"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientIpRangesShrink) {
		body["ClientIpRanges"] = request.ClientIpRangesShrink
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Enable) {
		body["Enable"] = request.Enable
	}

	if !dara.IsNil(request.LineId) {
		body["LineId"] = request.LineId
	}

	if !dara.IsNil(request.MatchType) {
		body["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.PathsShrink) {
		body["Paths"] = request.PathsShrink
	}

	if !dara.IsNil(request.RulesShrink) {
		body["Rules"] = request.RulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSdnsRuleCategory"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSdnsRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom resolution policy for a domain name connected to the current account. This operation supports the LINE mode and the IP_RANGE mode. IP address ranges can include IPv4 and IPv6 addresses. When a match occurs, A or AAAA records are returned based on the rules. By default, the policy is created in the disabled state. If Enable is set to true, the policy is enabled upon creation, but actual resolution still requires waiting for configuration propagation. The returned Id is used for subsequent queries, updates, enabling, disabling, and deletion. Create a disabled policy first, and then enable it after querying and verifying it. The policy type and the bound domain name cannot be modified after creation. If the response times out or is lost, query the list first to verify and avoid duplicate creation.
//
// @param request - CreateSdnsRuleCategoryRequest
//
// @return CreateSdnsRuleCategoryResponse
func (client *Client) CreateSdnsRuleCategory(request *CreateSdnsRuleCategoryRequest) (_result *CreateSdnsRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSdnsRuleCategoryResponse{}
	_body, _err := client.CreateSdnsRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a domain name based on the specified parameters.
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a domain name based on the specified parameters.
//
// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the domain names of a user based on the specified parameters.
//
// @param request - DescribeDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainsResponse
func (client *Client) DescribeDomainsWithOptions(request *DescribeDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
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
		Action:      dara.String("DescribeDomains"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the domain names of a user based on the specified parameters.
//
// @param request - DescribeDomainsRequest
//
// @return DescribeDomainsResponse
func (client *Client) DescribeDomains(request *DescribeDomainsRequest) (_result *DescribeDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DescribeDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the current HTTPDNS account information.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountInfoResponse
func (client *Client) GetAccountInfoWithOptions(runtime *dara.RuntimeOptions) (_result *GetAccountInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountInfo"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the current HTTPDNS account information.
//
// @return GetAccountInfoResponse
func (client *Client) GetAccountInfo() (_result *GetAccountInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountInfoResponse{}
	_body, _err := client.GetAccountInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves parse and billing statistics, corresponding to the "Billing" page feature in the console.
//
// @param request - GetResolveCountSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResolveCountSummaryResponse
func (client *Client) GetResolveCountSummaryWithOptions(request *GetResolveCountSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetResolveCountSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.TimeSpan) {
		query["TimeSpan"] = request.TimeSpan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResolveCountSummary"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResolveCountSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves parse and billing statistics, corresponding to the "Billing" page feature in the console.
//
// @param request - GetResolveCountSummaryRequest
//
// @return GetResolveCountSummaryResponse
func (client *Client) GetResolveCountSummary(request *GetResolveCountSummaryRequest) (_result *GetResolveCountSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResolveCountSummaryResponse{}
	_body, _err := client.GetResolveCountSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the statistics information of parse requests for a specific domain name.
//
// @param request - GetResolveStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResolveStatisticsResponse
func (client *Client) GetResolveStatisticsWithOptions(request *GetResolveStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetResolveStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.ProtocolName) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !dara.IsNil(request.TimeSpan) {
		query["TimeSpan"] = request.TimeSpan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResolveStatistics"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResolveStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the statistics information of parse requests for a specific domain name.
//
// @param request - GetResolveStatisticsRequest
//
// @return GetResolveStatisticsResponse
func (client *Client) GetResolveStatistics(request *GetResolveStatisticsRequest) (_result *GetResolveStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResolveStatisticsResponse{}
	_body, _err := client.GetResolveStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the domain names and resolution counts of a user based on the specified parameters.
//
// Description:
//
// # Operation description
//
// - You can specify the page number (PageNumber) and the number of entries per page (PageSize) for the query.
//
// - By default, the domain name list is sorted in descending order of the time when the domain names were added.
//
// @param request - ListDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
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

	if !dara.IsNil(request.Search) {
		query["Search"] = request.Search
	}

	if !dara.IsNil(request.WithoutMeteringData) {
		query["WithoutMeteringData"] = request.WithoutMeteringData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomains"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the domain names and resolution counts of a user based on the specified parameters.
//
// Description:
//
// # Operation description
//
// - You can specify the page number (PageNumber) and the number of entries per page (PageSize) for the query.
//
// - By default, the domain name list is sorted in descending order of the time when the domain names were added.
//
// @param request - ListDomainsRequest
//
// @return ListDomainsResponse
func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Refreshes the domain name cache.
//
// @param tmpReq - RefreshResolveCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshResolveCacheResponse
func (client *Client) RefreshResolveCacheWithOptions(tmpReq *RefreshResolveCacheRequest, runtime *dara.RuntimeOptions) (_result *RefreshResolveCacheResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RefreshResolveCacheShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Domains) {
		request.DomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Domains, dara.String("Domains"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainsShrink) {
		query["Domains"] = request.DomainsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshResolveCache"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshResolveCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes the domain name cache.
//
// @param request - RefreshResolveCacheRequest
//
// @return RefreshResolveCacheResponse
func (client *Client) RefreshResolveCache(request *RefreshResolveCacheRequest) (_result *RefreshResolveCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshResolveCacheResponse{}
	_body, _err := client.RefreshResolveCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom resolution policy of the current account by policy ID. A deleted policy cannot be recovered by using its original ID, and a new ID is generated when the policy is recreated. Disable the policy and verify the resolution results before deletion. After an address block policy is deleted, its scope is excluded from overlap checks with other policies. Actual resolution takes effect after configuration propagation. You can call SdnsListRuleCategories to verify the deletion.
//
// @param request - SdnsDeleteRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdnsDeleteRuleCategoryResponse
func (client *Client) SdnsDeleteRuleCategoryWithOptions(request *SdnsDeleteRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *SdnsDeleteRuleCategoryResponse, _err error) {
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
		Action:      dara.String("SdnsDeleteRuleCategory"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdnsDeleteRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom resolution policy of the current account by policy ID. A deleted policy cannot be recovered by using its original ID, and a new ID is generated when the policy is recreated. Disable the policy and verify the resolution results before deletion. After an address block policy is deleted, its scope is excluded from overlap checks with other policies. Actual resolution takes effect after configuration propagation. You can call SdnsListRuleCategories to verify the deletion.
//
// @param request - SdnsDeleteRuleCategoryRequest
//
// @return SdnsDeleteRuleCategoryResponse
func (client *Client) SdnsDeleteRuleCategory(request *SdnsDeleteRuleCategoryRequest) (_result *SdnsDeleteRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SdnsDeleteRuleCategoryResponse{}
	_body, _err := client.SdnsDeleteRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据ID获取规则策略
//
// @param request - SdnsGetRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdnsGetRuleCategoryResponse
func (client *Client) SdnsGetRuleCategoryWithOptions(request *SdnsGetRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *SdnsGetRuleCategoryResponse, _err error) {
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
		Action:      dara.String("SdnsGetRuleCategory"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdnsGetRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据ID获取规则策略
//
// @param request - SdnsGetRuleCategoryRequest
//
// @return SdnsGetRuleCategoryResponse
func (client *Client) SdnsGetRuleCategory(request *SdnsGetRuleCategoryRequest) (_result *SdnsGetRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SdnsGetRuleCategoryResponse{}
	_body, _err := client.SdnsGetRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the custom resolution policies for a specified domain name under the current account, including line policies, client CIDR block policies, and policies in enabled or disabled states. The returned IDs can be used to query details, update, enable, disable, or delete policies. This operation does not provide pagination parameters and does not guarantee the order of returned results. CIDR blocks are returned in the original input format.
//
// @param request - SdnsListRuleCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdnsListRuleCategoriesResponse
func (client *Client) SdnsListRuleCategoriesWithOptions(request *SdnsListRuleCategoriesRequest, runtime *dara.RuntimeOptions) (_result *SdnsListRuleCategoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SdnsListRuleCategories"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdnsListRuleCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom resolution policies for a specified domain name under the current account, including line policies, client CIDR block policies, and policies in enabled or disabled states. The returned IDs can be used to query details, update, enable, disable, or delete policies. This operation does not provide pagination parameters and does not guarantee the order of returned results. CIDR blocks are returned in the original input format.
//
// @param request - SdnsListRuleCategoriesRequest
//
// @return SdnsListRuleCategoriesResponse
func (client *Client) SdnsListRuleCategories(request *SdnsListRuleCategoriesRequest) (_result *SdnsListRuleCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SdnsListRuleCategoriesResponse{}
	_body, _err := client.SdnsListRuleCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用、禁用规则策略
//
// @param request - SdnsSwitchRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdnsSwitchRuleCategoryResponse
func (client *Client) SdnsSwitchRuleCategoryWithOptions(request *SdnsSwitchRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *SdnsSwitchRuleCategoryResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SdnsSwitchRuleCategory"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdnsSwitchRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用、禁用规则策略
//
// @param request - SdnsSwitchRuleCategoryRequest
//
// @return SdnsSwitchRuleCategoryResponse
func (client *Client) SdnsSwitchRuleCategory(request *SdnsSwitchRuleCategoryRequest) (_result *SdnsSwitchRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SdnsSwitchRuleCategoryResponse{}
	_body, _err := client.SdnsSwitchRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validates the submitted rules without saving or enabling any policies. In IP_RANGE mode, the Cip parameter is used to simulate a client IPv4 or IPv6 address. The system checks for CIDR block matches first and then returns the matched rules. If no custom parameters are used, pass an empty JSON object {} for the top-level Args parameter. If no rules are matched, the Data field is empty or omitted. A successful API call does not indicate a policy match. This operation does not check for CIDR block conflicts with existing policies, perform A/AAAA query type filtering or weight sampling, or simulate the complete resolution chain such as legacy CIDR block configurations, line fallback, or Function Compute. After saving and enabling the policies, verify them by using the actual resolution API.
//
// @param tmpReq - SdnsVerifyRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdnsVerifyRuleCategoryResponse
func (client *Client) SdnsVerifyRuleCategoryWithOptions(tmpReq *SdnsVerifyRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *SdnsVerifyRuleCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SdnsVerifyRuleCategoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Args) {
		request.ArgsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Args, dara.String("Args"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ClientIpRanges) {
		request.ClientIpRangesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIpRanges, dara.String("ClientIpRanges"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ArgsShrink) {
		body["Args"] = request.ArgsShrink
	}

	if !dara.IsNil(request.Cip) {
		body["Cip"] = request.Cip
	}

	if !dara.IsNil(request.ClientIpRangesShrink) {
		body["ClientIpRanges"] = request.ClientIpRangesShrink
	}

	if !dara.IsNil(request.MatchType) {
		body["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.RulesShrink) {
		body["Rules"] = request.RulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SdnsVerifyRuleCategory"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdnsVerifyRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates the submitted rules without saving or enabling any policies. In IP_RANGE mode, the Cip parameter is used to simulate a client IPv4 or IPv6 address. The system checks for CIDR block matches first and then returns the matched rules. If no custom parameters are used, pass an empty JSON object {} for the top-level Args parameter. If no rules are matched, the Data field is empty or omitted. A successful API call does not indicate a policy match. This operation does not check for CIDR block conflicts with existing policies, perform A/AAAA query type filtering or weight sampling, or simulate the complete resolution chain such as legacy CIDR block configurations, line fallback, or Function Compute. After saving and enabling the policies, verify them by using the actual resolution API.
//
// @param request - SdnsVerifyRuleCategoryRequest
//
// @return SdnsVerifyRuleCategoryResponse
func (client *Client) SdnsVerifyRuleCategory(request *SdnsVerifyRuleCategoryRequest) (_result *SdnsVerifyRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SdnsVerifyRuleCategoryResponse{}
	_body, _err := client.SdnsVerifyRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the complete rule list of an existing policy for the current account. For IP_RANGE policies, the client IP address ranges can be replaced as a whole at the same time. If ClientIpRanges is omitted, the original IP address ranges are retained. Explicitly passing an empty array is not allowed. The update does not change the bound domain names, match types, or enabled status of the policy. The new ranges are checked for conflicts only with other IP address range policies under the same domain name. If the validation fails, the update is not saved. After a successful update, you can call SdnsGetRuleCategory to verify the configuration. Actual DNS resolution requires waiting for configuration propagation.
//
// @param tmpReq - UpdateSdnsRuleCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSdnsRuleCategoryResponse
func (client *Client) UpdateSdnsRuleCategoryWithOptions(tmpReq *UpdateSdnsRuleCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateSdnsRuleCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSdnsRuleCategoryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientIpRanges) {
		body["ClientIpRanges"] = request.ClientIpRanges
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.RulesShrink) {
		body["Rules"] = request.RulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSdnsRuleCategory"),
		Version:     dara.String("2016-02-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSdnsRuleCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the complete rule list of an existing policy for the current account. For IP_RANGE policies, the client IP address ranges can be replaced as a whole at the same time. If ClientIpRanges is omitted, the original IP address ranges are retained. Explicitly passing an empty array is not allowed. The update does not change the bound domain names, match types, or enabled status of the policy. The new ranges are checked for conflicts only with other IP address range policies under the same domain name. If the validation fails, the update is not saved. After a successful update, you can call SdnsGetRuleCategory to verify the configuration. Actual DNS resolution requires waiting for configuration propagation.
//
// @param request - UpdateSdnsRuleCategoryRequest
//
// @return UpdateSdnsRuleCategoryResponse
func (client *Client) UpdateSdnsRuleCategory(request *UpdateSdnsRuleCategoryRequest) (_result *UpdateSdnsRuleCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSdnsRuleCategoryResponse{}
	_body, _err := client.UpdateSdnsRuleCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
