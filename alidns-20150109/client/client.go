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
		"public": dara.String("alidns.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("alidns"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a custom line for a domain name.
//
// Description:
//
// The end IP address of an IP address segment must be greater than or equal to its start IP address.
//
// The IP address ranges of segments cannot overlap across any custom lines for the domain name.
//
// @param request - AddCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomLineResponse
func (client *Client) AddCustomLineWithOptions(request *AddCustomLineRequest, runtime *dara.RuntimeOptions) (_result *AddCustomLineResponse, _err error) {
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

	if !dara.IsNil(request.IpSegment) {
		query["IpSegment"] = request.IpSegment
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineName) {
		query["LineName"] = request.LineName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCustomLine"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCustomLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a custom line for a domain name.
//
// Description:
//
// The end IP address of an IP address segment must be greater than or equal to its start IP address.
//
// The IP address ranges of segments cannot overlap across any custom lines for the domain name.
//
// @param request - AddCustomLineRequest
//
// @return AddCustomLineResponse
func (client *Client) AddCustomLine(request *AddCustomLineRequest) (_result *AddCustomLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCustomLineResponse{}
	_body, _err := client.AddCustomLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a domain name to the DNS authoritative proxy service.
//
// @param request - AddDnsCacheDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsCacheDomainResponse
func (client *Client) AddDnsCacheDomainWithOptions(request *AddDnsCacheDomainRequest, runtime *dara.RuntimeOptions) (_result *AddDnsCacheDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheTtlMax) {
		query["CacheTtlMax"] = request.CacheTtlMax
	}

	if !dara.IsNil(request.CacheTtlMin) {
		query["CacheTtlMin"] = request.CacheTtlMin
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SourceDnsServer) {
		query["SourceDnsServer"] = request.SourceDnsServer
	}

	if !dara.IsNil(request.SourceEdns) {
		query["SourceEdns"] = request.SourceEdns
	}

	if !dara.IsNil(request.SourceProtocol) {
		query["SourceProtocol"] = request.SourceProtocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnsCacheDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnsCacheDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name to the DNS authoritative proxy service.
//
// @param request - AddDnsCacheDomainRequest
//
// @return AddDnsCacheDomainResponse
func (client *Client) AddDnsCacheDomain(request *AddDnsCacheDomainRequest) (_result *AddDnsCacheDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDnsCacheDomainResponse{}
	_body, _err := client.AddDnsCacheDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an access strategy.
//
// @param request - AddDnsGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsGtmAccessStrategyResponse
func (client *Client) AddDnsGtmAccessStrategyWithOptions(request *AddDnsGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *AddDnsGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultAddrPool) {
		query["DefaultAddrPool"] = request.DefaultAddrPool
	}

	if !dara.IsNil(request.DefaultAddrPoolType) {
		query["DefaultAddrPoolType"] = request.DefaultAddrPoolType
	}

	if !dara.IsNil(request.DefaultLatencyOptimization) {
		query["DefaultLatencyOptimization"] = request.DefaultLatencyOptimization
	}

	if !dara.IsNil(request.DefaultLbaStrategy) {
		query["DefaultLbaStrategy"] = request.DefaultLbaStrategy
	}

	if !dara.IsNil(request.DefaultMaxReturnAddrNum) {
		query["DefaultMaxReturnAddrNum"] = request.DefaultMaxReturnAddrNum
	}

	if !dara.IsNil(request.DefaultMinAvailableAddrNum) {
		query["DefaultMinAvailableAddrNum"] = request.DefaultMinAvailableAddrNum
	}

	if !dara.IsNil(request.FailoverAddrPool) {
		query["FailoverAddrPool"] = request.FailoverAddrPool
	}

	if !dara.IsNil(request.FailoverAddrPoolType) {
		query["FailoverAddrPoolType"] = request.FailoverAddrPoolType
	}

	if !dara.IsNil(request.FailoverLatencyOptimization) {
		query["FailoverLatencyOptimization"] = request.FailoverLatencyOptimization
	}

	if !dara.IsNil(request.FailoverLbaStrategy) {
		query["FailoverLbaStrategy"] = request.FailoverLbaStrategy
	}

	if !dara.IsNil(request.FailoverMaxReturnAddrNum) {
		query["FailoverMaxReturnAddrNum"] = request.FailoverMaxReturnAddrNum
	}

	if !dara.IsNil(request.FailoverMinAvailableAddrNum) {
		query["FailoverMinAvailableAddrNum"] = request.FailoverMinAvailableAddrNum
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lines) {
		query["Lines"] = request.Lines
	}

	if !dara.IsNil(request.StrategyMode) {
		query["StrategyMode"] = request.StrategyMode
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnsGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnsGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an access strategy.
//
// @param request - AddDnsGtmAccessStrategyRequest
//
// @return AddDnsGtmAccessStrategyResponse
func (client *Client) AddDnsGtmAccessStrategy(request *AddDnsGtmAccessStrategyRequest) (_result *AddDnsGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDnsGtmAccessStrategyResponse{}
	_body, _err := client.AddDnsGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an address pool.
//
// @param request - AddDnsGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsGtmAddressPoolResponse
func (client *Client) AddDnsGtmAddressPoolWithOptions(request *AddDnsGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *AddDnsGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addr) {
		query["Addr"] = request.Addr
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LbaStrategy) {
		query["LbaStrategy"] = request.LbaStrategy
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.MonitorStatus) {
		query["MonitorStatus"] = request.MonitorStatus
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnsGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnsGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an address pool.
//
// @param request - AddDnsGtmAddressPoolRequest
//
// @return AddDnsGtmAddressPoolResponse
func (client *Client) AddDnsGtmAddressPool(request *AddDnsGtmAddressPoolRequest) (_result *AddDnsGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDnsGtmAddressPoolResponse{}
	_body, _err := client.AddDnsGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a health check.
//
// Description:
//
// **
//
// @param request - AddDnsGtmMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnsGtmMonitorResponse
func (client *Client) AddDnsGtmMonitorWithOptions(request *AddDnsGtmMonitorRequest, runtime *dara.RuntimeOptions) (_result *AddDnsGtmMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnsGtmMonitor"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnsGtmMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a health check.
//
// Description:
//
// **
//
// @param request - AddDnsGtmMonitorRequest
//
// @return AddDnsGtmMonitorResponse
func (client *Client) AddDnsGtmMonitor(request *AddDnsGtmMonitorRequest) (_result *AddDnsGtmMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDnsGtmMonitorResponse{}
	_body, _err := client.AddDnsGtmMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a domain name.
//
// Description:
//
// For more information, see <props="china">[Domain name validity](https://help.aliyun.com/document_detail/67788.html)<props="intl">[Domain name validity](https://www.alibabacloud.com/help/zh/doc-detail/67788.htm).
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
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomain"),
		Version:     dara.String("2015-01-09"),
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
// Adds a domain name.
//
// Description:
//
// For more information, see <props="china">[Domain name validity](https://help.aliyun.com/document_detail/67788.html)<props="intl">[Domain name validity](https://www.alibabacloud.com/help/zh/doc-detail/67788.htm).
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
// Creates a backup for a domain based on the specified domain name and backup cycle.
//
// @param request - AddDomainBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainBackupResponse
func (client *Client) AddDomainBackupWithOptions(request *AddDomainBackupRequest, runtime *dara.RuntimeOptions) (_result *AddDomainBackupResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PeriodType) {
		query["PeriodType"] = request.PeriodType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomainBackup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDomainBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup for a domain based on the specified domain name and backup cycle.
//
// @param request - AddDomainBackupRequest
//
// @return AddDomainBackupResponse
func (client *Client) AddDomainBackup(request *AddDomainBackupRequest) (_result *AddDomainBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDomainBackupResponse{}
	_body, _err := client.AddDomainBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a domain name group.
//
// @param request - AddDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainGroupResponse
func (client *Client) AddDomainGroupWithOptions(request *AddDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *AddDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomainGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a domain name group.
//
// @param request - AddDomainGroupRequest
//
// @return AddDomainGroupResponse
func (client *Client) AddDomainGroup(request *AddDomainGroupRequest) (_result *AddDomainGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDomainGroupResponse{}
	_body, _err := client.AddDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a DNS record.
//
// @param request - AddDomainRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDomainRecordResponse
func (client *Client) AddDomainRecordWithOptions(request *AddDomainRecordRequest, runtime *dara.RuntimeOptions) (_result *AddDomainRecordResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RR) {
		query["RR"] = request.RR
	}

	if !dara.IsNil(request.TTL) {
		query["TTL"] = request.TTL
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDomainRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDomainRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a DNS record.
//
// @param request - AddDomainRecordRequest
//
// @return AddDomainRecordResponse
func (client *Client) AddDomainRecord(request *AddDomainRecordRequest) (_result *AddDomainRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDomainRecordResponse{}
	_body, _err := client.AddDomainRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access strategy based on the specified parameters.
//
// @param request - AddGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGtmAccessStrategyResponse
func (client *Client) AddGtmAccessStrategyWithOptions(request *AddGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *AddGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessLines) {
		query["AccessLines"] = request.AccessLines
	}

	if !dara.IsNil(request.DefaultAddrPoolId) {
		query["DefaultAddrPoolId"] = request.DefaultAddrPoolId
	}

	if !dara.IsNil(request.FailoverAddrPoolId) {
		query["FailoverAddrPoolId"] = request.FailoverAddrPoolId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access strategy based on the specified parameters.
//
// @param request - AddGtmAccessStrategyRequest
//
// @return AddGtmAccessStrategyResponse
func (client *Client) AddGtmAccessStrategy(request *AddGtmAccessStrategyRequest) (_result *AddGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGtmAccessStrategyResponse{}
	_body, _err := client.AddGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an address pool.
//
// @param request - AddGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGtmAddressPoolResponse
func (client *Client) AddGtmAddressPoolWithOptions(request *AddGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *AddGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addr) {
		query["Addr"] = request.Addr
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MinAvailableAddrNum) {
		query["MinAvailableAddrNum"] = request.MinAvailableAddrNum
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.MonitorStatus) {
		query["MonitorStatus"] = request.MonitorStatus
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an address pool.
//
// @param request - AddGtmAddressPoolRequest
//
// @return AddGtmAddressPoolResponse
func (client *Client) AddGtmAddressPool(request *AddGtmAddressPoolRequest) (_result *AddGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGtmAddressPoolResponse{}
	_body, _err := client.AddGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a health check.
//
// @param request - AddGtmMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGtmMonitorResponse
func (client *Client) AddGtmMonitorWithOptions(request *AddGtmMonitorRequest, runtime *dara.RuntimeOptions) (_result *AddGtmMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGtmMonitor"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGtmMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a health check.
//
// @param request - AddGtmMonitorRequest
//
// @return AddGtmMonitorResponse
func (client *Client) AddGtmMonitor(request *AddGtmMonitorRequest) (_result *AddGtmMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGtmMonitorResponse{}
	_body, _err := client.AddGtmMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a disaster recovery plan.
//
// @param request - AddGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGtmRecoveryPlanResponse
func (client *Client) AddGtmRecoveryPlanWithOptions(request *AddGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *AddGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FaultAddrPool) {
		query["FaultAddrPool"] = request.FaultAddrPool
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a disaster recovery plan.
//
// @param request - AddGtmRecoveryPlanRequest
//
// @return AddGtmRecoveryPlanResponse
func (client *Client) AddGtmRecoveryPlan(request *AddGtmRecoveryPlanRequest) (_result *AddGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddGtmRecoveryPlanResponse{}
	_body, _err := client.AddGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an authoritative record for recursive resolution.
//
// Description:
//
// - You can specify a domain name (DomainName), page number (PageNumber), and page size (PageSize) to retrieve the list of DNS records for that domain name.
//
// - To find DNS records that contain a specific keyword, you can specify the keyword for the host record (RRKeyWord), record type (TypeKeyWord), or record value (ValueKeyWord).
//
// - By default, the list of DNS records is sorted from newest to oldest.
//
// - You can specify a domain group ID (GroupId). The \\`All Domains\\` group includes all domain names. The \\`Default Group\\` includes domain names that are not assigned to a group.
//
// @param request - AddRecursionRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecursionRecordResponse
func (client *Client) AddRecursionRecordWithOptions(request *AddRecursionRecordRequest, runtime *dara.RuntimeOptions) (_result *AddRecursionRecordResponse, _err error) {
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

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecursionRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecursionRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an authoritative record for recursive resolution.
//
// Description:
//
// - You can specify a domain name (DomainName), page number (PageNumber), and page size (PageSize) to retrieve the list of DNS records for that domain name.
//
// - To find DNS records that contain a specific keyword, you can specify the keyword for the host record (RRKeyWord), record type (TypeKeyWord), or record value (ValueKeyWord).
//
// - By default, the list of DNS records is sorted from newest to oldest.
//
// - You can specify a domain group ID (GroupId). The \\`All Domains\\` group includes all domain names. The \\`Default Group\\` includes domain names that are not assigned to a group.
//
// @param request - AddRecursionRecordRequest
//
// @return AddRecursionRecordResponse
func (client *Client) AddRecursionRecord(request *AddRecursionRecordRequest) (_result *AddRecursionRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRecursionRecordResponse{}
	_body, _err := client.AddRecursionRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a built-in authoritative domain name zone for recursive resolution.
//
// Description:
//
// The end IP address of each IP range must be greater than or equal to the start IP address.
//
// The IP address ranges of all IP ranges in all custom lines for a domain name cannot overlap.
//
// @param request - AddRecursionZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecursionZoneResponse
func (client *Client) AddRecursionZoneWithOptions(request *AddRecursionZoneRequest, runtime *dara.RuntimeOptions) (_result *AddRecursionZoneResponse, _err error) {
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

	if !dara.IsNil(request.ProxyPattern) {
		query["ProxyPattern"] = request.ProxyPattern
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecursionZone"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecursionZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a built-in authoritative domain name zone for recursive resolution.
//
// Description:
//
// The end IP address of each IP range must be greater than or equal to the start IP address.
//
// The IP address ranges of all IP ranges in all custom lines for a domain name cannot overlap.
//
// @param request - AddRecursionZoneRequest
//
// @return AddRecursionZoneResponse
func (client *Client) AddRecursionZone(request *AddRecursionZoneRequest) (_result *AddRecursionZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRecursionZoneResponse{}
	_body, _err := client.AddRecursionZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the serverHold status to a specified domain name.
//
// Description:
//
// ## Request description
//
// - This API adds the serverHold property to a specified domain name.
//
// @param request - AddRspDomainServerHoldStatusForGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRspDomainServerHoldStatusForGatewayResponse
func (client *Client) AddRspDomainServerHoldStatusForGatewayWithOptions(request *AddRspDomainServerHoldStatusForGatewayRequest, runtime *dara.RuntimeOptions) (_result *AddRspDomainServerHoldStatusForGatewayResponse, _err error) {
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

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.StatusMsg) {
		query["StatusMsg"] = request.StatusMsg
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRspDomainServerHoldStatusForGateway"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRspDomainServerHoldStatusForGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the serverHold status to a specified domain name.
//
// Description:
//
// ## Request description
//
// - This API adds the serverHold property to a specified domain name.
//
// @param request - AddRspDomainServerHoldStatusForGatewayRequest
//
// @return AddRspDomainServerHoldStatusForGatewayResponse
func (client *Client) AddRspDomainServerHoldStatusForGateway(request *AddRspDomainServerHoldStatusForGatewayRequest) (_result *AddRspDomainServerHoldStatusForGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRspDomainServerHoldStatusForGatewayResponse{}
	_body, _err := client.AddRspDomainServerHoldStatusForGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the serverHold status for a specified domain name.
//
// Description:
//
// ## Request description
//
// - Adds the serverHold status for a specified domain name.
//
// @param request - AddRspDomainServerHoldStatusForGatewayOteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRspDomainServerHoldStatusForGatewayOteResponse
func (client *Client) AddRspDomainServerHoldStatusForGatewayOteWithOptions(request *AddRspDomainServerHoldStatusForGatewayOteRequest, runtime *dara.RuntimeOptions) (_result *AddRspDomainServerHoldStatusForGatewayOteResponse, _err error) {
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

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.StatusMsg) {
		query["StatusMsg"] = request.StatusMsg
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRspDomainServerHoldStatusForGatewayOte"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRspDomainServerHoldStatusForGatewayOteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the serverHold status for a specified domain name.
//
// Description:
//
// ## Request description
//
// - Adds the serverHold status for a specified domain name.
//
// @param request - AddRspDomainServerHoldStatusForGatewayOteRequest
//
// @return AddRspDomainServerHoldStatusForGatewayOteResponse
func (client *Client) AddRspDomainServerHoldStatusForGatewayOte(request *AddRspDomainServerHoldStatusForGatewayOteRequest) (_result *AddRspDomainServerHoldStatusForGatewayOteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRspDomainServerHoldStatusForGatewayOteResponse{}
	_body, _err := client.AddRspDomainServerHoldStatusForGatewayOteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds paid domain names in Alibaba Cloud DNS to an instance ID.
//
// Description:
//
// An instance is considered a new instance if its ID starts with \\"dns-\\". New instances support multiple domain names. You can call this operation to bind domain names directly to the instance. An error occurs if the number of domain names exceeds the instance\\"s limit.
//
// An instance is considered a legacy instance if its ID does not start with \\"dns-\\". Legacy instances support only one domain name. If you call this operation on a legacy instance that already has a domain name, the existing domain name is replaced.
//
// @param request - BindInstanceDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindInstanceDomainsResponse
func (client *Client) BindInstanceDomainsWithOptions(request *BindInstanceDomainsRequest, runtime *dara.RuntimeOptions) (_result *BindInstanceDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindInstanceDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindInstanceDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds paid domain names in Alibaba Cloud DNS to an instance ID.
//
// Description:
//
// An instance is considered a new instance if its ID starts with \\"dns-\\". New instances support multiple domain names. You can call this operation to bind domain names directly to the instance. An error occurs if the number of domain names exceeds the instance\\"s limit.
//
// An instance is considered a legacy instance if its ID does not start with \\"dns-\\". Legacy instances support only one domain name. If you call this operation on a legacy instance that already has a domain name, the existing domain name is replaced.
//
// @param request - BindInstanceDomainsRequest
//
// @return BindInstanceDomainsResponse
func (client *Client) BindInstanceDomains(request *BindInstanceDomainsRequest) (_result *BindInstanceDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindInstanceDomainsResponse{}
	_body, _err := client.BindInstanceDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a domain name to a new group.
//
// Description:
//
// You can specify the ID of a domain name group (GroupId). The All Domains group contains all domain names, while the Default group contains domain names that are not assigned to any group.
//
// @param request - ChangeDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDomainGroupResponse
func (client *Client) ChangeDomainGroupWithOptions(request *ChangeDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeDomainGroupResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeDomainGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a domain name to a new group.
//
// Description:
//
// You can specify the ID of a domain name group (GroupId). The All Domains group contains all domain names, while the Default group contains domain names that are not assigned to any group.
//
// @param request - ChangeDomainGroupRequest
//
// @return ChangeDomainGroupResponse
func (client *Client) ChangeDomainGroup(request *ChangeDomainGroupRequest) (_result *ChangeDomainGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeDomainGroupResponse{}
	_body, _err := client.ChangeDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the domain name that is attached to a Cloud DNS product.
//
// Description:
//
// - **You can call this operation to change the domain name that is attached to a Cloud DNS product. To detach a domain name, call this operation and leave the NewDomain parameter empty.**
//
// - **This operation applies to instances of earlier versions. If you use a new edition, such as Personal Edition, Enterprise Standard Edition, or Enterprise Ultimate Edition, call the BindInstanceDomains operation instead.**
//
// @param request - ChangeDomainOfDnsProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDomainOfDnsProductResponse
func (client *Client) ChangeDomainOfDnsProductWithOptions(request *ChangeDomainOfDnsProductRequest, runtime *dara.RuntimeOptions) (_result *ChangeDomainOfDnsProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewDomain) {
		query["NewDomain"] = request.NewDomain
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeDomainOfDnsProduct"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDomainOfDnsProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the domain name that is attached to a Cloud DNS product.
//
// Description:
//
// - **You can call this operation to change the domain name that is attached to a Cloud DNS product. To detach a domain name, call this operation and leave the NewDomain parameter empty.**
//
// - **This operation applies to instances of earlier versions. If you use a new edition, such as Personal Edition, Enterprise Standard Edition, or Enterprise Ultimate Edition, call the BindInstanceDomains operation instead.**
//
// @param request - ChangeDomainOfDnsProductRequest
//
// @return ChangeDomainOfDnsProductResponse
func (client *Client) ChangeDomainOfDnsProduct(request *ChangeDomainOfDnsProductRequest) (_result *ChangeDomainOfDnsProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeDomainOfDnsProductResponse{}
	_body, _err := client.ChangeDomainOfDnsProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Copies a Global Traffic Manager (GTM) configuration.
//
// @param request - CopyGtmConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyGtmConfigResponse
func (client *Client) CopyGtmConfigWithOptions(request *CopyGtmConfigRequest, runtime *dara.RuntimeOptions) (_result *CopyGtmConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CopyType) {
		query["CopyType"] = request.CopyType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyGtmConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyGtmConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies a Global Traffic Manager (GTM) configuration.
//
// @param request - CopyGtmConfigRequest
//
// @return CopyGtmConfigResponse
func (client *Client) CopyGtmConfig(request *CopyGtmConfigRequest) (_result *CopyGtmConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyGtmConfigResponse{}
	_body, _err := client.CopyGtmConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers an agent.
//
// @param tmpReq - CreateAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAtiAgentRegisterInfoResponse
func (client *Client) CreateAtiAgentRegisterInfoWithOptions(tmpReq *CreateAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAtiAgentRegisterInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Endpoints) {
		request.EndpointsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Endpoints, dara.String("Endpoints"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentDescription) {
		query["AgentDescription"] = request.AgentDescription
	}

	if !dara.IsNil(request.AgentDisplayName) {
		query["AgentDisplayName"] = request.AgentDisplayName
	}

	if !dara.IsNil(request.AgentHost) {
		query["AgentHost"] = request.AgentHost
	}

	if !dara.IsNil(request.AgentVersion) {
		query["AgentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointsShrink) {
		query["Endpoints"] = request.EndpointsShrink
	}

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers an agent.
//
// @param request - CreateAtiAgentRegisterInfoRequest
//
// @return CreateAtiAgentRegisterInfoResponse
func (client *Client) CreateAtiAgentRegisterInfo(request *CreateAtiAgentRegisterInfoRequest) (_result *CreateAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAtiAgentRegisterInfoResponse{}
	_body, _err := client.CreateAtiAgentRegisterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers an agent - Step 2: Generates a DNS record for domain ownership verification.
//
// @param request - CreateAtiAgentRegisterInfoAcmeChallengeRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse
func (client *Client) CreateAtiAgentRegisterInfoAcmeChallengeRecordWithOptions(request *CreateAtiAgentRegisterInfoAcmeChallengeRecordRequest, runtime *dara.RuntimeOptions) (_result *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAtiAgentRegisterInfoAcmeChallengeRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers an agent - Step 2: Generates a DNS record for domain ownership verification.
//
// @param request - CreateAtiAgentRegisterInfoAcmeChallengeRecordRequest
//
// @return CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse
func (client *Client) CreateAtiAgentRegisterInfoAcmeChallengeRecord(request *CreateAtiAgentRegisterInfoAcmeChallengeRecordRequest) (_result *CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAtiAgentRegisterInfoAcmeChallengeRecordResponse{}
	_body, _err := client.CreateAtiAgentRegisterInfoAcmeChallengeRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a real-name registrant.
//
// @param request - CreateAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAtiRegistrantResponse
func (client *Client) CreateAtiRegistrantWithOptions(request *CreateAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *CreateAtiRegistrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cc) {
		query["Cc"] = request.Cc
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DocumentCode) {
		query["DocumentCode"] = request.DocumentCode
	}

	if !dara.IsNil(request.DocumentImage) {
		query["DocumentImage"] = request.DocumentImage
	}

	if !dara.IsNil(request.DocumentType) {
		query["DocumentType"] = request.DocumentType
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Street) {
		query["Street"] = request.Street
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a real-name registrant.
//
// @param request - CreateAtiRegistrantRequest
//
// @return CreateAtiRegistrantResponse
func (client *Client) CreateAtiRegistrant(request *CreateAtiRegistrantRequest) (_result *CreateAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAtiRegistrantResponse{}
	_body, _err := client.CreateAtiRegistrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an address in Global Traffic Manager (GTM) 3.0.
//
// @param tmpReq - CreateCloudGtmAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudGtmAddressResponse
func (client *Client) CreateCloudGtmAddressWithOptions(tmpReq *CreateCloudGtmAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudGtmAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCloudGtmAddressShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HealthTasks) {
		request.HealthTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HealthTasks, dara.String("HealthTasks"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AttributeInfo) {
		query["AttributeInfo"] = request.AttributeInfo
	}

	if !dara.IsNil(request.AvailableMode) {
		query["AvailableMode"] = request.AvailableMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthJudgement) {
		query["HealthJudgement"] = request.HealthJudgement
	}

	if !dara.IsNil(request.HealthTasksShrink) {
		query["HealthTasks"] = request.HealthTasksShrink
	}

	if !dara.IsNil(request.ManualAvailableStatus) {
		query["ManualAvailableStatus"] = request.ManualAvailableStatus
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudGtmAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudGtmAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an address in Global Traffic Manager (GTM) 3.0.
//
// @param request - CreateCloudGtmAddressRequest
//
// @return CreateCloudGtmAddressResponse
func (client *Client) CreateCloudGtmAddress(request *CreateCloudGtmAddressRequest) (_result *CreateCloudGtmAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudGtmAddressResponse{}
	_body, _err := client.CreateCloudGtmAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an address pool based on the specified parameters.
//
// @param request - CreateCloudGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudGtmAddressPoolResponse
func (client *Client) CreateCloudGtmAddressPoolWithOptions(request *CreateCloudGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolName) {
		query["AddressPoolName"] = request.AddressPoolName
	}

	if !dara.IsNil(request.AddressPoolType) {
		query["AddressPoolType"] = request.AddressPoolType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthJudgement) {
		query["HealthJudgement"] = request.HealthJudgement
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an address pool based on the specified parameters.
//
// @param request - CreateCloudGtmAddressPoolRequest
//
// @return CreateCloudGtmAddressPoolResponse
func (client *Client) CreateCloudGtmAddressPool(request *CreateCloudGtmAddressPoolRequest) (_result *CreateCloudGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudGtmAddressPoolResponse{}
	_body, _err := client.CreateCloudGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a configuration for a Global Traffic Manager (GTM) instance.
//
// @param request - CreateCloudGtmInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudGtmInstanceConfigResponse
func (client *Client) CreateCloudGtmInstanceConfigWithOptions(request *CreateCloudGtmInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudGtmInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ScheduleHostname) {
		query["ScheduleHostname"] = request.ScheduleHostname
	}

	if !dara.IsNil(request.ScheduleRrType) {
		query["ScheduleRrType"] = request.ScheduleRrType
	}

	if !dara.IsNil(request.ScheduleZoneMode) {
		query["ScheduleZoneMode"] = request.ScheduleZoneMode
	}

	if !dara.IsNil(request.ScheduleZoneName) {
		query["ScheduleZoneName"] = request.ScheduleZoneName
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudGtmInstanceConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudGtmInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a configuration for a Global Traffic Manager (GTM) instance.
//
// @param request - CreateCloudGtmInstanceConfigRequest
//
// @return CreateCloudGtmInstanceConfigResponse
func (client *Client) CreateCloudGtmInstanceConfig(request *CreateCloudGtmInstanceConfigRequest) (_result *CreateCloudGtmInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudGtmInstanceConfigResponse{}
	_body, _err := client.CreateCloudGtmInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a health check template in Global Traffic Manager (GTM) 3.0.
//
// @param tmpReq - CreateCloudGtmMonitorTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudGtmMonitorTemplateResponse
func (client *Client) CreateCloudGtmMonitorTemplateWithOptions(tmpReq *CreateCloudGtmMonitorTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudGtmMonitorTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCloudGtmMonitorTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IspCityNodes) {
		request.IspCityNodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IspCityNodes, dara.String("IspCityNodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.ExtendInfo) {
		query["ExtendInfo"] = request.ExtendInfo
	}

	if !dara.IsNil(request.FailureRate) {
		query["FailureRate"] = request.FailureRate
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.IspCityNodesShrink) {
		query["IspCityNodes"] = request.IspCityNodesShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudGtmMonitorTemplate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudGtmMonitorTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a health check template in Global Traffic Manager (GTM) 3.0.
//
// @param request - CreateCloudGtmMonitorTemplateRequest
//
// @return CreateCloudGtmMonitorTemplateResponse
func (client *Client) CreateCloudGtmMonitorTemplate(request *CreateCloudGtmMonitorTemplateRequest) (_result *CreateCloudGtmMonitorTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudGtmMonitorTemplateResponse{}
	_body, _err := client.CreateCloudGtmMonitorTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a public DNS AccessKey.
//
// @param request - CreatePdnsAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdnsAppKeyResponse
func (client *Client) CreatePdnsAppKeyWithOptions(request *CreatePdnsAppKeyRequest, runtime *dara.RuntimeOptions) (_result *CreatePdnsAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdnsAppKey"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdnsAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public DNS AccessKey.
//
// @param request - CreatePdnsAppKeyRequest
//
// @return CreatePdnsAppKeyResponse
func (client *Client) CreatePdnsAppKey(request *CreatePdnsAppKeyRequest) (_result *CreatePdnsAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePdnsAppKeyResponse{}
	_body, _err := client.CreatePdnsAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Public DNS UDP IP Address Segment
//
// @param request - CreatePdnsUdpIpSegmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdnsUdpIpSegmentResponse
func (client *Client) CreatePdnsUdpIpSegmentWithOptions(request *CreatePdnsUdpIpSegmentRequest, runtime *dara.RuntimeOptions) (_result *CreatePdnsUdpIpSegmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.IpToken) {
		query["IpToken"] = request.IpToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdnsUdpIpSegment"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdnsUdpIpSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Public DNS UDP IP Address Segment
//
// @param request - CreatePdnsUdpIpSegmentRequest
//
// @return CreatePdnsUdpIpSegmentResponse
func (client *Client) CreatePdnsUdpIpSegment(request *CreatePdnsUdpIpSegmentRequest) (_result *CreatePdnsUdpIpSegmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePdnsUdpIpSegmentResponse{}
	_body, _err := client.CreatePdnsUdpIpSegmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes Agent registration information.
//
// @param request - DeleteAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAtiAgentRegisterInfoResponse
func (client *Client) DeleteAtiAgentRegisterInfoWithOptions(request *DeleteAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *DeleteAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes Agent registration information.
//
// @param request - DeleteAtiAgentRegisterInfoRequest
//
// @return DeleteAtiAgentRegisterInfoResponse
func (client *Client) DeleteAtiAgentRegisterInfo(request *DeleteAtiAgentRegisterInfoRequest) (_result *DeleteAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAtiAgentRegisterInfoResponse{}
	_body, _err := client.DeleteAtiAgentRegisterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a verified registrant.
//
// @param request - DeleteAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAtiRegistrantResponse
func (client *Client) DeleteAtiRegistrantWithOptions(request *DeleteAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *DeleteAtiRegistrantResponse, _err error) {
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

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a verified registrant.
//
// @param request - DeleteAtiRegistrantRequest
//
// @return DeleteAtiRegistrantResponse
func (client *Client) DeleteAtiRegistrant(request *DeleteAtiRegistrantRequest) (_result *DeleteAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAtiRegistrantResponse{}
	_body, _err := client.DeleteAtiRegistrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an address.
//
// @param request - DeleteCloudGtmAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudGtmAddressResponse
func (client *Client) DeleteCloudGtmAddressWithOptions(request *DeleteCloudGtmAddressRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudGtmAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudGtmAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudGtmAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address.
//
// @param request - DeleteCloudGtmAddressRequest
//
// @return DeleteCloudGtmAddressResponse
func (client *Client) DeleteCloudGtmAddress(request *DeleteCloudGtmAddressRequest) (_result *DeleteCloudGtmAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudGtmAddressResponse{}
	_body, _err := client.DeleteCloudGtmAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an address pool.
//
// @param request - DeleteCloudGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudGtmAddressPoolResponse
func (client *Client) DeleteCloudGtmAddressPoolWithOptions(request *DeleteCloudGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address pool.
//
// @param request - DeleteCloudGtmAddressPoolRequest
//
// @return DeleteCloudGtmAddressPoolResponse
func (client *Client) DeleteCloudGtmAddressPool(request *DeleteCloudGtmAddressPoolRequest) (_result *DeleteCloudGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudGtmAddressPoolResponse{}
	_body, _err := client.DeleteCloudGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access domain name from the configuration of a Global Traffic Manager (GTM) 3.0 instance.
//
// @param request - DeleteCloudGtmInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudGtmInstanceConfigResponse
func (client *Client) DeleteCloudGtmInstanceConfigWithOptions(request *DeleteCloudGtmInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudGtmInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudGtmInstanceConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudGtmInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access domain name from the configuration of a Global Traffic Manager (GTM) 3.0 instance.
//
// @param request - DeleteCloudGtmInstanceConfigRequest
//
// @return DeleteCloudGtmInstanceConfigResponse
func (client *Client) DeleteCloudGtmInstanceConfig(request *DeleteCloudGtmInstanceConfigRequest) (_result *DeleteCloudGtmInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudGtmInstanceConfigResponse{}
	_body, _err := client.DeleteCloudGtmInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a health check template.
//
// @param request - DeleteCloudGtmMonitorTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudGtmMonitorTemplateResponse
func (client *Client) DeleteCloudGtmMonitorTemplateWithOptions(request *DeleteCloudGtmMonitorTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudGtmMonitorTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudGtmMonitorTemplate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudGtmMonitorTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a health check template.
//
// @param request - DeleteCloudGtmMonitorTemplateRequest
//
// @return DeleteCloudGtmMonitorTemplateResponse
func (client *Client) DeleteCloudGtmMonitorTemplate(request *DeleteCloudGtmMonitorTemplateRequest) (_result *DeleteCloudGtmMonitorTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudGtmMonitorTemplateResponse{}
	_body, _err := client.DeleteCloudGtmMonitorTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a batch of custom lines by specifying their unique IDs.
//
// @param request - DeleteCustomLinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomLinesResponse
func (client *Client) DeleteCustomLinesWithOptions(request *DeleteCustomLinesRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomLinesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineIds) {
		query["LineIds"] = request.LineIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomLines"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a batch of custom lines by specifying their unique IDs.
//
// @param request - DeleteCustomLinesRequest
//
// @return DeleteCustomLinesResponse
func (client *Client) DeleteCustomLines(request *DeleteCustomLinesRequest) (_result *DeleteCustomLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomLinesResponse{}
	_body, _err := client.DeleteCustomLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified domain name from the authoritative DNS proxy.
//
// @param request - DeleteDnsCacheDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDnsCacheDomainResponse
func (client *Client) DeleteDnsCacheDomainWithOptions(request *DeleteDnsCacheDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDnsCacheDomainResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDnsCacheDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDnsCacheDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified domain name from the authoritative DNS proxy.
//
// @param request - DeleteDnsCacheDomainRequest
//
// @return DeleteDnsCacheDomainResponse
func (client *Client) DeleteDnsCacheDomain(request *DeleteDnsCacheDomainRequest) (_result *DeleteDnsCacheDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDnsCacheDomainResponse{}
	_body, _err := client.DeleteDnsCacheDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access policy by its ID.
//
// @param request - DeleteDnsGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDnsGtmAccessStrategyResponse
func (client *Client) DeleteDnsGtmAccessStrategyWithOptions(request *DeleteDnsGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteDnsGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDnsGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDnsGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access policy by its ID.
//
// @param request - DeleteDnsGtmAccessStrategyRequest
//
// @return DeleteDnsGtmAccessStrategyResponse
func (client *Client) DeleteDnsGtmAccessStrategy(request *DeleteDnsGtmAccessStrategyRequest) (_result *DeleteDnsGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDnsGtmAccessStrategyResponse{}
	_body, _err := client.DeleteDnsGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an address pool by its ID.
//
// @param request - DeleteDnsGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDnsGtmAddressPoolResponse
func (client *Client) DeleteDnsGtmAddressPoolWithOptions(request *DeleteDnsGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteDnsGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDnsGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDnsGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address pool by its ID.
//
// @param request - DeleteDnsGtmAddressPoolRequest
//
// @return DeleteDnsGtmAddressPoolResponse
func (client *Client) DeleteDnsGtmAddressPool(request *DeleteDnsGtmAddressPoolRequest) (_result *DeleteDnsGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDnsGtmAddressPoolResponse{}
	_body, _err := client.DeleteDnsGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified domain name.
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
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2015-01-09"),
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
// Deletes the specified domain name.
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
// Deletes a domain name group and moves its domain names to the default group.
//
// Description:
//
// > The default group cannot be deleted.
//
// @param request - DeleteDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainGroupResponse
func (client *Client) DeleteDomainGroupWithOptions(request *DeleteDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a domain name group and moves its domain names to the default group.
//
// Description:
//
// > The default group cannot be deleted.
//
// @param request - DeleteDomainGroupRequest
//
// @return DeleteDomainGroupResponse
func (client *Client) DeleteDomainGroup(request *DeleteDomainGroupRequest) (_result *DeleteDomainGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainGroupResponse{}
	_body, _err := client.DeleteDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a DNS record based on the specified request parameters.
//
// @param request - DeleteDomainRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainRecordResponse
func (client *Client) DeleteDomainRecordWithOptions(request *DeleteDomainRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a DNS record based on the specified request parameters.
//
// @param request - DeleteDomainRecordRequest
//
// @return DeleteDomainRecordResponse
func (client *Client) DeleteDomainRecord(request *DeleteDomainRecordRequest) (_result *DeleteDomainRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainRecordResponse{}
	_body, _err := client.DeleteDomainRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access policy.
//
// @param request - DeleteGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGtmAccessStrategyResponse
func (client *Client) DeleteGtmAccessStrategyWithOptions(request *DeleteGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access policy.
//
// @param request - DeleteGtmAccessStrategyRequest
//
// @return DeleteGtmAccessStrategyResponse
func (client *Client) DeleteGtmAccessStrategy(request *DeleteGtmAccessStrategyRequest) (_result *DeleteGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGtmAccessStrategyResponse{}
	_body, _err := client.DeleteGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an address pool.
//
// @param request - DeleteGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGtmAddressPoolResponse
func (client *Client) DeleteGtmAddressPoolWithOptions(request *DeleteGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an address pool.
//
// @param request - DeleteGtmAddressPoolRequest
//
// @return DeleteGtmAddressPoolResponse
func (client *Client) DeleteGtmAddressPool(request *DeleteGtmAddressPoolRequest) (_result *DeleteGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGtmAddressPoolResponse{}
	_body, _err := client.DeleteGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a disaster recovery plan.
//
// @param request - DeleteGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGtmRecoveryPlanResponse
func (client *Client) DeleteGtmRecoveryPlanWithOptions(request *DeleteGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a disaster recovery plan.
//
// @param request - DeleteGtmRecoveryPlanRequest
//
// @return DeleteGtmRecoveryPlanResponse
func (client *Client) DeleteGtmRecoveryPlan(request *DeleteGtmRecoveryPlanRequest) (_result *DeleteGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGtmRecoveryPlanResponse{}
	_body, _err := client.DeleteGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a built-in authoritative DNS record used for recursive resolution.
//
// @param request - DeleteRecursionRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecursionRecordResponse
func (client *Client) DeleteRecursionRecordWithOptions(request *DeleteRecursionRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecursionRecordResponse, _err error) {
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

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecursionRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecursionRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a built-in authoritative DNS record used for recursive resolution.
//
// @param request - DeleteRecursionRecordRequest
//
// @return DeleteRecursionRecordResponse
func (client *Client) DeleteRecursionRecord(request *DeleteRecursionRecordRequest) (_result *DeleteRecursionRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecursionRecordResponse{}
	_body, _err := client.DeleteRecursionRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a built-in authoritative zone used for recursive resolution.
//
// Description:
//
// If a zone contains locked DNS records, this operation does not delete them.
//
// @param request - DeleteRecursionZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecursionZoneResponse
func (client *Client) DeleteRecursionZoneWithOptions(request *DeleteRecursionZoneRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecursionZoneResponse, _err error) {
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

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecursionZone"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecursionZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a built-in authoritative zone used for recursive resolution.
//
// Description:
//
// If a zone contains locked DNS records, this operation does not delete them.
//
// @param request - DeleteRecursionZoneRequest
//
// @return DeleteRecursionZoneResponse
func (client *Client) DeleteRecursionZone(request *DeleteRecursionZoneRequest) (_result *DeleteRecursionZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecursionZoneResponse{}
	_body, _err := client.DeleteRecursionZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the DNS records for a specified host record.
//
// Description:
//
// Locked DNS records will not be deleted.
//
// @param request - DeleteSubDomainRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubDomainRecordsResponse
func (client *Client) DeleteSubDomainRecordsWithOptions(request *DeleteSubDomainRecordsRequest, runtime *dara.RuntimeOptions) (_result *DeleteSubDomainRecordsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RR) {
		query["RR"] = request.RR
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSubDomainRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubDomainRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the DNS records for a specified host record.
//
// Description:
//
// Locked DNS records will not be deleted.
//
// @param request - DeleteSubDomainRecordsRequest
//
// @return DeleteSubDomainRecordsResponse
func (client *Client) DeleteSubDomainRecords(request *DeleteSubDomainRecordsRequest) (_result *DeleteSubDomainRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSubDomainRecordsResponse{}
	_body, _err := client.DeleteSubDomainRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an Agent registration.
//
// @param request - DescribeAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAtiAgentRegisterInfoResponse
func (client *Client) DescribeAtiAgentRegisterInfoWithOptions(request *DescribeAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an Agent registration.
//
// @param request - DescribeAtiAgentRegisterInfoRequest
//
// @return DescribeAtiAgentRegisterInfoResponse
func (client *Client) DescribeAtiAgentRegisterInfo(request *DescribeAtiAgentRegisterInfoRequest) (_result *DescribeAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAtiAgentRegisterInfoResponse{}
	_body, _err := client.DescribeAtiAgentRegisterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries alert settings.
//
// @param request - DescribeAtiAlertSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAtiAlertSettingsResponse
func (client *Client) DescribeAtiAlertSettingsWithOptions(request *DescribeAtiAlertSettingsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAtiAlertSettingsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAtiAlertSettings"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAtiAlertSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert settings.
//
// @param request - DescribeAtiAlertSettingsRequest
//
// @return DescribeAtiAlertSettingsResponse
func (client *Client) DescribeAtiAlertSettings(request *DescribeAtiAlertSettingsRequest) (_result *DescribeAtiAlertSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAtiAlertSettingsResponse{}
	_body, _err := client.DescribeAtiAlertSettingsWithOptions(request, runtime)
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
// @param request - DescribeAtiCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAtiCertificateResponse
func (client *Client) DescribeAtiCertificateWithOptions(request *DescribeAtiCertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeAtiCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentCertificateId) {
		query["AgentCertificateId"] = request.AgentCertificateId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAtiCertificate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAtiCertificateResponse{}
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
// @param request - DescribeAtiCertificateRequest
//
// @return DescribeAtiCertificateResponse
func (client *Client) DescribeAtiCertificate(request *DescribeAtiCertificateRequest) (_result *DescribeAtiCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAtiCertificateResponse{}
	_body, _err := client.DescribeAtiCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a real-name verified registrant.
//
// @param request - DescribeAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAtiRegistrantResponse
func (client *Client) DescribeAtiRegistrantWithOptions(request *DescribeAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *DescribeAtiRegistrantResponse, _err error) {
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

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a real-name verified registrant.
//
// @param request - DescribeAtiRegistrantRequest
//
// @return DescribeAtiRegistrantResponse
func (client *Client) DescribeAtiRegistrant(request *DescribeAtiRegistrantRequest) (_result *DescribeAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAtiRegistrantResponse{}
	_body, _err := client.DescribeAtiRegistrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution results of a batch operation task using a task ID. If you do not specify a task ID, the results of the most recent batch task are returned.
//
// @param request - DescribeBatchResultCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBatchResultCountResponse
func (client *Client) DescribeBatchResultCountWithOptions(request *DescribeBatchResultCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeBatchResultCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchType) {
		query["BatchType"] = request.BatchType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBatchResultCount"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBatchResultCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution results of a batch operation task using a task ID. If you do not specify a task ID, the results of the most recent batch task are returned.
//
// @param request - DescribeBatchResultCountRequest
//
// @return DescribeBatchResultCountResponse
func (client *Client) DescribeBatchResultCount(request *DescribeBatchResultCountRequest) (_result *DescribeBatchResultCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBatchResultCountResponse{}
	_body, _err := client.DescribeBatchResultCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a batch processing result.
//
// Description:
//
// *Prerequisite: You can call this operation after the batch task is complete.**
//
// @param request - DescribeBatchResultDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBatchResultDetailResponse
func (client *Client) DescribeBatchResultDetailWithOptions(request *DescribeBatchResultDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeBatchResultDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchType) {
		query["BatchType"] = request.BatchType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeBatchResultDetail"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBatchResultDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a batch processing result.
//
// Description:
//
// *Prerequisite: You can call this operation after the batch task is complete.**
//
// @param request - DescribeBatchResultDetailRequest
//
// @return DescribeBatchResultDetailResponse
func (client *Client) DescribeBatchResultDetail(request *DescribeBatchResultDetailRequest) (_result *DescribeBatchResultDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBatchResultDetailResponse{}
	_body, _err := client.DescribeBatchResultDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of an address based on the specified input parameters.
//
// @param request - DescribeCloudGtmAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmAddressResponse
func (client *Client) DescribeCloudGtmAddressWithOptions(request *DescribeCloudGtmAddressRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of an address based on the specified input parameters.
//
// @param request - DescribeCloudGtmAddressRequest
//
// @return DescribeCloudGtmAddressResponse
func (client *Client) DescribeCloudGtmAddress(request *DescribeCloudGtmAddressRequest) (_result *DescribeCloudGtmAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmAddressResponse{}
	_body, _err := client.DescribeCloudGtmAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a specified address pool.
//
// @param request - DescribeCloudGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmAddressPoolResponse
func (client *Client) DescribeCloudGtmAddressPoolWithOptions(request *DescribeCloudGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a specified address pool.
//
// @param request - DescribeCloudGtmAddressPoolRequest
//
// @return DescribeCloudGtmAddressPoolResponse
func (client *Client) DescribeCloudGtmAddressPool(request *DescribeCloudGtmAddressPoolRequest) (_result *DescribeCloudGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmAddressPoolResponse{}
	_body, _err := client.DescribeCloudGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about the instances that reference an address pool.
//
// @param request - DescribeCloudGtmAddressPoolReferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmAddressPoolReferenceResponse
func (client *Client) DescribeCloudGtmAddressPoolReferenceWithOptions(request *DescribeCloudGtmAddressPoolReferenceRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmAddressPoolReferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmAddressPoolReference"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmAddressPoolReferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about the instances that reference an address pool.
//
// @param request - DescribeCloudGtmAddressPoolReferenceRequest
//
// @return DescribeCloudGtmAddressPoolReferenceResponse
func (client *Client) DescribeCloudGtmAddressPoolReference(request *DescribeCloudGtmAddressPoolReferenceRequest) (_result *DescribeCloudGtmAddressPoolReferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmAddressPoolReferenceResponse{}
	_body, _err := client.DescribeCloudGtmAddressPoolReferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the address pools and Global Traffic Manager (GTM) 3.0 instances that reference a specified address.
//
// @param request - DescribeCloudGtmAddressReferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmAddressReferenceResponse
func (client *Client) DescribeCloudGtmAddressReferenceWithOptions(request *DescribeCloudGtmAddressReferenceRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmAddressReferenceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmAddressReference"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmAddressReferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the address pools and Global Traffic Manager (GTM) 3.0 instances that reference a specified address.
//
// @param request - DescribeCloudGtmAddressReferenceRequest
//
// @return DescribeCloudGtmAddressReferenceResponse
func (client *Client) DescribeCloudGtmAddressReference(request *DescribeCloudGtmAddressReferenceRequest) (_result *DescribeCloudGtmAddressReferenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmAddressReferenceResponse{}
	_body, _err := client.DescribeCloudGtmAddressReferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the global alert configuration for Global Traffic Manager (GTM).
//
// @param request - DescribeCloudGtmGlobalAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmGlobalAlertResponse
func (client *Client) DescribeCloudGtmGlobalAlertWithOptions(request *DescribeCloudGtmGlobalAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmGlobalAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmGlobalAlert"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmGlobalAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global alert configuration for Global Traffic Manager (GTM).
//
// @param request - DescribeCloudGtmGlobalAlertRequest
//
// @return DescribeCloudGtmGlobalAlertResponse
func (client *Client) DescribeCloudGtmGlobalAlert(request *DescribeCloudGtmGlobalAlertRequest) (_result *DescribeCloudGtmGlobalAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmGlobalAlertResponse{}
	_body, _err := client.DescribeCloudGtmGlobalAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alert configuration for an instance.
//
// @param request - DescribeCloudGtmInstanceConfigAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmInstanceConfigAlertResponse
func (client *Client) DescribeCloudGtmInstanceConfigAlertWithOptions(request *DescribeCloudGtmInstanceConfigAlertRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmInstanceConfigAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmInstanceConfigAlert"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmInstanceConfigAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert configuration for an instance.
//
// @param request - DescribeCloudGtmInstanceConfigAlertRequest
//
// @return DescribeCloudGtmInstanceConfigAlertResponse
func (client *Client) DescribeCloudGtmInstanceConfigAlert(request *DescribeCloudGtmInstanceConfigAlertRequest) (_result *DescribeCloudGtmInstanceConfigAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmInstanceConfigAlertResponse{}
	_body, _err := client.DescribeCloudGtmInstanceConfigAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the full configuration of a GTM 3.0 access domain name, including alert settings, address pools, and address details.
//
// @param request - DescribeCloudGtmInstanceConfigFullInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmInstanceConfigFullInfoResponse
func (client *Client) DescribeCloudGtmInstanceConfigFullInfoWithOptions(request *DescribeCloudGtmInstanceConfigFullInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmInstanceConfigFullInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmInstanceConfigFullInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmInstanceConfigFullInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the full configuration of a GTM 3.0 access domain name, including alert settings, address pools, and address details.
//
// @param request - DescribeCloudGtmInstanceConfigFullInfoRequest
//
// @return DescribeCloudGtmInstanceConfigFullInfoResponse
func (client *Client) DescribeCloudGtmInstanceConfigFullInfo(request *DescribeCloudGtmInstanceConfigFullInfoRequest) (_result *DescribeCloudGtmInstanceConfigFullInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmInstanceConfigFullInfoResponse{}
	_body, _err := client.DescribeCloudGtmInstanceConfigFullInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a specified health check template.
//
// @param request - DescribeCloudGtmMonitorTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmMonitorTemplateResponse
func (client *Client) DescribeCloudGtmMonitorTemplateWithOptions(request *DescribeCloudGtmMonitorTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmMonitorTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmMonitorTemplate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmMonitorTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a specified health check template.
//
// @param request - DescribeCloudGtmMonitorTemplateRequest
//
// @return DescribeCloudGtmMonitorTemplateResponse
func (client *Client) DescribeCloudGtmMonitorTemplate(request *DescribeCloudGtmMonitorTemplateRequest) (_result *DescribeCloudGtmMonitorTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmMonitorTemplateResponse{}
	_body, _err := client.DescribeCloudGtmMonitorTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeCloudGtmSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmSummaryResponse
func (client *Client) DescribeCloudGtmSummaryWithOptions(request *DescribeCloudGtmSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeCloudGtmSummaryRequest
//
// @return DescribeCloudGtmSummaryResponse
func (client *Client) DescribeCloudGtmSummary(request *DescribeCloudGtmSummaryRequest) (_result *DescribeCloudGtmSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmSummaryResponse{}
	_body, _err := client.DescribeCloudGtmSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the system lines supported by Global Traffic Manager (GTM).
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudGtmSystemLinesResponse
func (client *Client) DescribeCloudGtmSystemLinesWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeCloudGtmSystemLinesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudGtmSystemLines"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudGtmSystemLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the system lines supported by Global Traffic Manager (GTM).
//
// @return DescribeCloudGtmSystemLinesResponse
func (client *Client) DescribeCloudGtmSystemLines() (_result *DescribeCloudGtmSystemLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudGtmSystemLinesResponse{}
	_body, _err := client.DescribeCloudGtmSystemLinesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a custom line by its unique ID.
//
// @param request - DescribeCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomLineResponse
func (client *Client) DescribeCustomLineWithOptions(request *DescribeCustomLineRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineId) {
		query["LineId"] = request.LineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomLine"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a custom line by its unique ID.
//
// @param request - DescribeCustomLineRequest
//
// @return DescribeCustomLineResponse
func (client *Client) DescribeCustomLine(request *DescribeCustomLineRequest) (_result *DescribeCustomLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomLineResponse{}
	_body, _err := client.DescribeCustomLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the custom lines for a domain name.
//
// @param request - DescribeCustomLinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomLinesResponse
func (client *Client) DescribeCustomLinesWithOptions(request *DescribeCustomLinesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomLinesResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeCustomLines"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom lines for a domain name.
//
// @param request - DescribeCustomLinesRequest
//
// @return DescribeCustomLinesResponse
func (client *Client) DescribeCustomLines(request *DescribeCustomLinesRequest) (_result *DescribeCustomLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomLinesResponse{}
	_body, _err := client.DescribeCustomLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of subdomains that have weight configurations based on the specified parameters.
//
// @param request - DescribeDNSSLBSubDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDNSSLBSubDomainsResponse
func (client *Client) DescribeDNSSLBSubDomainsWithOptions(request *DescribeDNSSLBSubDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDNSSLBSubDomainsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDNSSLBSubDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDNSSLBSubDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of subdomains that have weight configurations based on the specified parameters.
//
// @param request - DescribeDNSSLBSubDomainsRequest
//
// @return DescribeDNSSLBSubDomainsResponse
func (client *Client) DescribeDNSSLBSubDomains(request *DescribeDNSSLBSubDomainsRequest) (_result *DescribeDNSSLBSubDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDNSSLBSubDomainsResponse{}
	_body, _err := client.DescribeDNSSLBSubDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries authoritative proxy domain names based on the specified parameters.
//
// @param request - DescribeDnsCacheDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsCacheDomainsResponse
func (client *Client) DescribeDnsCacheDomainsWithOptions(request *DescribeDnsCacheDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsCacheDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeDnsCacheDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsCacheDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries authoritative proxy domain names based on the specified parameters.
//
// @param request - DescribeDnsCacheDomainsRequest
//
// @return DescribeDnsCacheDomainsResponse
func (client *Client) DescribeDnsCacheDomains(request *DescribeDnsCacheDomainsRequest) (_result *DescribeDnsCacheDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsCacheDomainsResponse{}
	_body, _err := client.DescribeDnsCacheDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access strategies for a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmAccessStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAccessStrategiesResponse
func (client *Client) DescribeDnsGtmAccessStrategiesWithOptions(request *DescribeDnsGtmAccessStrategiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAccessStrategiesResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StrategyMode) {
		query["StrategyMode"] = request.StrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAccessStrategies"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAccessStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access strategies for a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmAccessStrategiesRequest
//
// @return DescribeDnsGtmAccessStrategiesResponse
func (client *Client) DescribeDnsGtmAccessStrategies(request *DescribeDnsGtmAccessStrategiesRequest) (_result *DescribeDnsGtmAccessStrategiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAccessStrategiesResponse{}
	_body, _err := client.DescribeDnsGtmAccessStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified access strategy.
//
// @param request - DescribeDnsGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAccessStrategyResponse
func (client *Client) DescribeDnsGtmAccessStrategyWithOptions(request *DescribeDnsGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified access strategy.
//
// @param request - DescribeDnsGtmAccessStrategyRequest
//
// @return DescribeDnsGtmAccessStrategyResponse
func (client *Client) DescribeDnsGtmAccessStrategy(request *DescribeDnsGtmAccessStrategyRequest) (_result *DescribeDnsGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAccessStrategyResponse{}
	_body, _err := client.DescribeDnsGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the available configurations for an access policy based on an instance ID.
//
// @param request - DescribeDnsGtmAccessStrategyAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAccessStrategyAvailableConfigResponse
func (client *Client) DescribeDnsGtmAccessStrategyAvailableConfigWithOptions(request *DescribeDnsGtmAccessStrategyAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAccessStrategyAvailableConfigResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyMode) {
		query["StrategyMode"] = request.StrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAccessStrategyAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the available configurations for an access policy based on an instance ID.
//
// @param request - DescribeDnsGtmAccessStrategyAvailableConfigRequest
//
// @return DescribeDnsGtmAccessStrategyAvailableConfigResponse
func (client *Client) DescribeDnsGtmAccessStrategyAvailableConfig(request *DescribeDnsGtmAccessStrategyAvailableConfigRequest) (_result *DescribeDnsGtmAccessStrategyAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.DescribeDnsGtmAccessStrategyAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the region where an address is located.
//
// @param request - DescribeDnsGtmAddrAttributeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAddrAttributeInfoResponse
func (client *Client) DescribeDnsGtmAddrAttributeInfoWithOptions(request *DescribeDnsGtmAddrAttributeInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAddrAttributeInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addrs) {
		query["Addrs"] = request.Addrs
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAddrAttributeInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAddrAttributeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the region where an address is located.
//
// @param request - DescribeDnsGtmAddrAttributeInfoRequest
//
// @return DescribeDnsGtmAddrAttributeInfoResponse
func (client *Client) DescribeDnsGtmAddrAttributeInfo(request *DescribeDnsGtmAddrAttributeInfoRequest) (_result *DescribeDnsGtmAddrAttributeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAddrAttributeInfoResponse{}
	_body, _err := client.DescribeDnsGtmAddrAttributeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available configurations for an address pool in a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmAddressPoolAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAddressPoolAvailableConfigResponse
func (client *Client) DescribeDnsGtmAddressPoolAvailableConfigWithOptions(request *DescribeDnsGtmAddressPoolAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAddressPoolAvailableConfigResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAddressPoolAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAddressPoolAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available configurations for an address pool in a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmAddressPoolAvailableConfigRequest
//
// @return DescribeDnsGtmAddressPoolAvailableConfigResponse
func (client *Client) DescribeDnsGtmAddressPoolAvailableConfig(request *DescribeDnsGtmAddressPoolAvailableConfigRequest) (_result *DescribeDnsGtmAddressPoolAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAddressPoolAvailableConfigResponse{}
	_body, _err := client.DescribeDnsGtmAddressPoolAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available alert contact groups.
//
// @param request - DescribeDnsGtmAvailableAlertGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmAvailableAlertGroupResponse
func (client *Client) DescribeDnsGtmAvailableAlertGroupWithOptions(request *DescribeDnsGtmAvailableAlertGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmAvailableAlertGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmAvailableAlertGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmAvailableAlertGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available alert contact groups.
//
// @param request - DescribeDnsGtmAvailableAlertGroupRequest
//
// @return DescribeDnsGtmAvailableAlertGroupResponse
func (client *Client) DescribeDnsGtmAvailableAlertGroup(request *DescribeDnsGtmAvailableAlertGroupRequest) (_result *DescribeDnsGtmAvailableAlertGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmAvailableAlertGroupResponse{}
	_body, _err := client.DescribeDnsGtmAvailableAlertGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an instance based on the specified instance ID.
//
// @param request - DescribeDnsGtmInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceResponse
func (client *Client) DescribeDnsGtmInstanceWithOptions(request *DescribeDnsGtmInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstance"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an instance based on the specified instance ID.
//
// @param request - DescribeDnsGtmInstanceRequest
//
// @return DescribeDnsGtmInstanceResponse
func (client *Client) DescribeDnsGtmInstance(request *DescribeDnsGtmInstanceRequest) (_result *DescribeDnsGtmInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceResponse{}
	_body, _err := client.DescribeDnsGtmInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an address pool by its ID.
//
// @param request - DescribeDnsGtmInstanceAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceAddressPoolResponse
func (client *Client) DescribeDnsGtmInstanceAddressPoolWithOptions(request *DescribeDnsGtmInstanceAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstanceAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an address pool by its ID.
//
// @param request - DescribeDnsGtmInstanceAddressPoolRequest
//
// @return DescribeDnsGtmInstanceAddressPoolResponse
func (client *Client) DescribeDnsGtmInstanceAddressPool(request *DescribeDnsGtmInstanceAddressPoolRequest) (_result *DescribeDnsGtmInstanceAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceAddressPoolResponse{}
	_body, _err := client.DescribeDnsGtmInstanceAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the address pools of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmInstanceAddressPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceAddressPoolsResponse
func (client *Client) DescribeDnsGtmInstanceAddressPoolsWithOptions(request *DescribeDnsGtmInstanceAddressPoolsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceAddressPoolsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeDnsGtmInstanceAddressPools"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceAddressPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the address pools of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeDnsGtmInstanceAddressPoolsRequest
//
// @return DescribeDnsGtmInstanceAddressPoolsResponse
func (client *Client) DescribeDnsGtmInstanceAddressPools(request *DescribeDnsGtmInstanceAddressPoolsRequest) (_result *DescribeDnsGtmInstanceAddressPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceAddressPoolsResponse{}
	_body, _err := client.DescribeDnsGtmInstanceAddressPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the status of an instance based on its ID.
//
// @param request - DescribeDnsGtmInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceStatusResponse
func (client *Client) DescribeDnsGtmInstanceStatusWithOptions(request *DescribeDnsGtmInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceStatusResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstanceStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the status of an instance based on its ID.
//
// @param request - DescribeDnsGtmInstanceStatusRequest
//
// @return DescribeDnsGtmInstanceStatusResponse
func (client *Client) DescribeDnsGtmInstanceStatus(request *DescribeDnsGtmInstanceStatusRequest) (_result *DescribeDnsGtmInstanceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceStatusResponse{}
	_body, _err := client.DescribeDnsGtmInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the system-assigned CNAME based on the specified instance ID.
//
// @param request - DescribeDnsGtmInstanceSystemCnameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstanceSystemCnameResponse
func (client *Client) DescribeDnsGtmInstanceSystemCnameWithOptions(request *DescribeDnsGtmInstanceSystemCnameRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstanceSystemCnameResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmInstanceSystemCname"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstanceSystemCnameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the system-assigned CNAME based on the specified instance ID.
//
// @param request - DescribeDnsGtmInstanceSystemCnameRequest
//
// @return DescribeDnsGtmInstanceSystemCnameResponse
func (client *Client) DescribeDnsGtmInstanceSystemCname(request *DescribeDnsGtmInstanceSystemCnameRequest) (_result *DescribeDnsGtmInstanceSystemCnameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstanceSystemCnameResponse{}
	_body, _err := client.DescribeDnsGtmInstanceSystemCnameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param request - DescribeDnsGtmInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmInstancesResponse
func (client *Client) DescribeDnsGtmInstancesWithOptions(request *DescribeDnsGtmInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeDnsGtmInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param request - DescribeDnsGtmInstancesRequest
//
// @return DescribeDnsGtmInstancesResponse
func (client *Client) DescribeDnsGtmInstances(request *DescribeDnsGtmInstancesRequest) (_result *DescribeDnsGtmInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmInstancesResponse{}
	_body, _err := client.DescribeDnsGtmInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs for an instance.
//
// @param request - DescribeDnsGtmLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmLogsResponse
func (client *Client) DescribeDnsGtmLogsWithOptions(request *DescribeDnsGtmLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs for an instance.
//
// @param request - DescribeDnsGtmLogsRequest
//
// @return DescribeDnsGtmLogsResponse
func (client *Client) DescribeDnsGtmLogs(request *DescribeDnsGtmLogsRequest) (_result *DescribeDnsGtmLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmLogsResponse{}
	_body, _err := client.DescribeDnsGtmLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available configurations for DNS health checks.
//
// @param request - DescribeDnsGtmMonitorAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmMonitorAvailableConfigResponse
func (client *Client) DescribeDnsGtmMonitorAvailableConfigWithOptions(request *DescribeDnsGtmMonitorAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmMonitorAvailableConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmMonitorAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmMonitorAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available configurations for DNS health checks.
//
// @param request - DescribeDnsGtmMonitorAvailableConfigRequest
//
// @return DescribeDnsGtmMonitorAvailableConfigResponse
func (client *Client) DescribeDnsGtmMonitorAvailableConfig(request *DescribeDnsGtmMonitorAvailableConfigRequest) (_result *DescribeDnsGtmMonitorAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmMonitorAvailableConfigResponse{}
	_body, _err := client.DescribeDnsGtmMonitorAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the health check configuration for an address pool.
//
// @param request - DescribeDnsGtmMonitorConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsGtmMonitorConfigResponse
func (client *Client) DescribeDnsGtmMonitorConfigWithOptions(request *DescribeDnsGtmMonitorConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsGtmMonitorConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsGtmMonitorConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsGtmMonitorConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health check configuration for an address pool.
//
// @param request - DescribeDnsGtmMonitorConfigRequest
//
// @return DescribeDnsGtmMonitorConfigResponse
func (client *Client) DescribeDnsGtmMonitorConfig(request *DescribeDnsGtmMonitorConfigRequest) (_result *DescribeDnsGtmMonitorConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsGtmMonitorConfigResponse{}
	_body, _err := client.DescribeDnsGtmMonitorConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a paid Alibaba Cloud DNS instance by its instance ID.
//
// @param request - DescribeDnsProductInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsProductInstanceResponse
func (client *Client) DescribeDnsProductInstanceWithOptions(request *DescribeDnsProductInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsProductInstanceResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsProductInstance"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsProductInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a paid Alibaba Cloud DNS instance by its instance ID.
//
// @param request - DescribeDnsProductInstanceRequest
//
// @return DescribeDnsProductInstanceResponse
func (client *Client) DescribeDnsProductInstance(request *DescribeDnsProductInstanceRequest) (_result *DescribeDnsProductInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsProductInstanceResponse{}
	_body, _err := client.DescribeDnsProductInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of paid DNS product instances that match the specified parameters.
//
// Description:
//
// > **If the response does not contain a domain name, the Alibaba Cloud DNS instance is not associated with any domain names.**
//
// @param request - DescribeDnsProductInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnsProductInstancesResponse
func (client *Client) DescribeDnsProductInstancesWithOptions(request *DescribeDnsProductInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnsProductInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.VersionCode) {
		query["VersionCode"] = request.VersionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnsProductInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnsProductInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of paid DNS product instances that match the specified parameters.
//
// Description:
//
// > **If the response does not contain a domain name, the Alibaba Cloud DNS instance is not associated with any domain names.**
//
// @param request - DescribeDnsProductInstancesRequest
//
// @return DescribeDnsProductInstancesResponse
func (client *Client) DescribeDnsProductInstances(request *DescribeDnsProductInstancesRequest) (_result *DescribeDnsProductInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDnsProductInstancesResponse{}
	_body, _err := client.DescribeDnsProductInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an overview of request statistics for a DNS over HTTPS (DoH) account.
//
// @param request - DescribeDohAccountStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohAccountStatisticsResponse
func (client *Client) DescribeDohAccountStatisticsWithOptions(request *DescribeDohAccountStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohAccountStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohAccountStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohAccountStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an overview of request statistics for a DNS over HTTPS (DoH) account.
//
// @param request - DescribeDohAccountStatisticsRequest
//
// @return DescribeDohAccountStatisticsResponse
func (client *Client) DescribeDohAccountStatistics(request *DescribeDohAccountStatisticsRequest) (_result *DescribeDohAccountStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohAccountStatisticsResponse{}
	_body, _err := client.DescribeDohAccountStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves an overview of statistics for DNS over HTTPS (DoH) requests for a domain name.
//
// @param request - DescribeDohDomainStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohDomainStatisticsResponse
func (client *Client) DescribeDohDomainStatisticsWithOptions(request *DescribeDohDomainStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohDomainStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohDomainStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohDomainStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an overview of statistics for DNS over HTTPS (DoH) requests for a domain name.
//
// @param request - DescribeDohDomainStatisticsRequest
//
// @return DescribeDohDomainStatisticsResponse
func (client *Client) DescribeDohDomainStatistics(request *DescribeDohDomainStatisticsRequest) (_result *DescribeDohDomainStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohDomainStatisticsResponse{}
	_body, _err := client.DescribeDohDomainStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries request statistics for DNS over HTTPS (DoH) domain names.
//
// @param request - DescribeDohDomainStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohDomainStatisticsSummaryResponse
func (client *Client) DescribeDohDomainStatisticsSummaryWithOptions(request *DescribeDohDomainStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohDomainStatisticsSummaryResponse, _err error) {
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

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
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
		Action:      dara.String("DescribeDohDomainStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohDomainStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries request statistics for DNS over HTTPS (DoH) domain names.
//
// @param request - DescribeDohDomainStatisticsSummaryRequest
//
// @return DescribeDohDomainStatisticsSummaryResponse
func (client *Client) DescribeDohDomainStatisticsSummary(request *DescribeDohDomainStatisticsSummaryRequest) (_result *DescribeDohDomainStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohDomainStatisticsSummaryResponse{}
	_body, _err := client.DescribeDohDomainStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries for statistics on DNS over HTTPS (DoH) requests for a subdomain.
//
// @param request - DescribeDohSubDomainStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohSubDomainStatisticsResponse
func (client *Client) DescribeDohSubDomainStatisticsWithOptions(request *DescribeDohSubDomainStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohSubDomainStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohSubDomainStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohSubDomainStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries for statistics on DNS over HTTPS (DoH) requests for a subdomain.
//
// @param request - DescribeDohSubDomainStatisticsRequest
//
// @return DescribeDohSubDomainStatisticsResponse
func (client *Client) DescribeDohSubDomainStatistics(request *DescribeDohSubDomainStatisticsRequest) (_result *DescribeDohSubDomainStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohSubDomainStatisticsResponse{}
	_body, _err := client.DescribeDohSubDomainStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a summary of request statistics for subdomains using DNS over HTTPS (DoH).
//
// @param request - DescribeDohSubDomainStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohSubDomainStatisticsSummaryResponse
func (client *Client) DescribeDohSubDomainStatisticsSummaryWithOptions(request *DescribeDohSubDomainStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohSubDomainStatisticsSummaryResponse, _err error) {
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

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohSubDomainStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohSubDomainStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a summary of request statistics for subdomains using DNS over HTTPS (DoH).
//
// @param request - DescribeDohSubDomainStatisticsSummaryRequest
//
// @return DescribeDohSubDomainStatisticsSummaryResponse
func (client *Client) DescribeDohSubDomainStatisticsSummary(request *DescribeDohSubDomainStatisticsSummaryRequest) (_result *DescribeDohSubDomainStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohSubDomainStatisticsSummaryResponse{}
	_body, _err := client.DescribeDohSubDomainStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the basic information of a DNS over HTTPS (DoH) user.
//
// @param request - DescribeDohUserInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDohUserInfoResponse
func (client *Client) DescribeDohUserInfoWithOptions(request *DescribeDohUserInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDohUserInfoResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDohUserInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDohUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the basic information of a DNS over HTTPS (DoH) user.
//
// @param request - DescribeDohUserInfoRequest
//
// @return DescribeDohUserInfoResponse
func (client *Client) DescribeDohUserInfo(request *DescribeDohUserInfoRequest) (_result *DescribeDohUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDohUserInfoResponse{}
	_body, _err := client.DescribeDohUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Domain Name System Security Extensions (DNSSEC) information for a specified domain name.
//
// @param request - DescribeDomainDnssecInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainDnssecInfoResponse
func (client *Client) DescribeDomainDnssecInfoWithOptions(request *DescribeDomainDnssecInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainDnssecInfoResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainDnssecInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainDnssecInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Domain Name System Security Extensions (DNSSEC) information for a specified domain name.
//
// @param request - DescribeDomainDnssecInfoRequest
//
// @return DescribeDomainDnssecInfoResponse
func (client *Client) DescribeDomainDnssecInfo(request *DescribeDomainDnssecInfoRequest) (_result *DescribeDomainDnssecInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainDnssecInfoResponse{}
	_body, _err := client.DescribeDomainDnssecInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain name groups.
//
// @param request - DescribeDomainGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainGroupsResponse
func (client *Client) DescribeDomainGroupsWithOptions(request *DescribeDomainGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeDomainGroups"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain name groups.
//
// @param request - DescribeDomainGroupsRequest
//
// @return DescribeDomainGroupsResponse
func (client *Client) DescribeDomainGroups(request *DescribeDomainGroupsRequest) (_result *DescribeDomainGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainGroupsResponse{}
	_body, _err := client.DescribeDomainGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a specified domain name.
//
// Description:
//
// In this example, the domain name is bound to an instance of Alibaba Cloud DNS Ultimate Edition. For more information about line enumeration, see the RecordLines response parameter.
//
// @param request - DescribeDomainInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainInfoResponse
func (client *Client) DescribeDomainInfoWithOptions(request *DescribeDomainInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainInfoResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NeedDetailAttributes) {
		query["NeedDetailAttributes"] = request.NeedDetailAttributes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specified domain name.
//
// Description:
//
// In this example, the domain name is bound to an instance of Alibaba Cloud DNS Ultimate Edition. For more information about line enumeration, see the RecordLines response parameter.
//
// @param request - DescribeDomainInfoRequest
//
// @return DescribeDomainInfoResponse
func (client *Client) DescribeDomainInfo(request *DescribeDomainInfoRequest) (_result *DescribeDomainInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainInfoResponse{}
	_body, _err := client.DescribeDomainInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs for a domain name based on the specified parameters.
//
// @param request - DescribeDomainLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainLogsResponse
func (client *Client) DescribeDomainLogsWithOptions(request *DescribeDomainLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs for a domain name based on the specified parameters.
//
// @param request - DescribeDomainLogsRequest
//
// @return DescribeDomainLogsResponse
func (client *Client) DescribeDomainLogs(request *DescribeDomainLogsRequest) (_result *DescribeDomainLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainLogsResponse{}
	_body, _err := client.DescribeDomainLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current list of name servers for a domain name and determines whether the servers are managed by Alibaba Cloud DNS.
//
// Description:
//
// > This operation directly queries the authoritative server of the domain name registry to retrieve the DNS server names for the domain name. An error may be returned if the domain name is inactive. For example, if the domain name has a serverHold or clientHold status, or has not passed identity verification.
//
// @param request - DescribeDomainNsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainNsResponse
func (client *Client) DescribeDomainNsWithOptions(request *DescribeDomainNsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainNsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainNs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainNsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current list of name servers for a domain name and determines whether the servers are managed by Alibaba Cloud DNS.
//
// Description:
//
// > This operation directly queries the authoritative server of the domain name registry to retrieve the DNS server names for the domain name. An error may be returned if the domain name is inactive. For example, if the domain name has a serverHold or clientHold status, or has not passed identity verification.
//
// @param request - DescribeDomainNsRequest
//
// @return DescribeDomainNsResponse
func (client *Client) DescribeDomainNs(request *DescribeDomainNsRequest) (_result *DescribeDomainNsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainNsResponse{}
	_body, _err := client.DescribeDomainNsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a DNS record by its ID.
//
// @param request - DescribeDomainRecordInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRecordInfoResponse
func (client *Client) DescribeDomainRecordInfoWithOptions(request *DescribeDomainRecordInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRecordInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRecordInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRecordInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a DNS record by its ID.
//
// @param request - DescribeDomainRecordInfoRequest
//
// @return DescribeDomainRecordInfoResponse
func (client *Client) DescribeDomainRecordInfo(request *DescribeDomainRecordInfoRequest) (_result *DescribeDomainRecordInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRecordInfoResponse{}
	_body, _err := client.DescribeDomainRecordInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the DNS records for a specified root domain based on the input parameters.
//
// Description:
//
// - You can specify the domain name (DomainName), page number (PageNumber), and page size (PageSize) to retrieve a list of DNS records.
//
// - You can specify a keyword for the host record (RRKeyWord), record type (TypeKeyWord), or record value (ValueKeyWord) to query DNS records that contain the keyword.
//
// - By default, DNS records are sorted in descending order by the time they were added.
//
// - You can specify a domain group ID (GroupId) to query the DNS records in a specific group.
//
// @param request - DescribeDomainRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRecordsResponse
func (client *Client) DescribeDomainRecordsWithOptions(request *DescribeDomainRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RRKeyWord) {
		query["RRKeyWord"] = request.RRKeyWord
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.TypeKeyWord) {
		query["TypeKeyWord"] = request.TypeKeyWord
	}

	if !dara.IsNil(request.ValueKeyWord) {
		query["ValueKeyWord"] = request.ValueKeyWord
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the DNS records for a specified root domain based on the input parameters.
//
// Description:
//
// - You can specify the domain name (DomainName), page number (PageNumber), and page size (PageSize) to retrieve a list of DNS records.
//
// - You can specify a keyword for the host record (RRKeyWord), record type (TypeKeyWord), or record value (ValueKeyWord) to query DNS records that contain the keyword.
//
// - By default, DNS records are sorted in descending order by the time they were added.
//
// - You can specify a domain group ID (GroupId) to query the DNS records in a specific group.
//
// @param request - DescribeDomainRecordsRequest
//
// @return DescribeDomainRecordsResponse
func (client *Client) DescribeDomainRecords(request *DescribeDomainRecordsRequest) (_result *DescribeDomainRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRecordsResponse{}
	_body, _err := client.DescribeDomainRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request volumes for all paid domain names in your account.
//
// @param request - DescribeDomainResolveStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainResolveStatisticsSummaryResponse
func (client *Client) DescribeDomainResolveStatisticsSummaryWithOptions(request *DescribeDomainResolveStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainResolveStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainResolveStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainResolveStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request volumes for all paid domain names in your account.
//
// @param request - DescribeDomainResolveStatisticsSummaryRequest
//
// @return DescribeDomainResolveStatisticsSummaryResponse
func (client *Client) DescribeDomainResolveStatisticsSummary(request *DescribeDomainResolveStatisticsSummaryRequest) (_result *DescribeDomainResolveStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainResolveStatisticsSummaryResponse{}
	_body, _err := client.DescribeDomainResolveStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of real-time requests for a specified primary domain name.
//
// Description:
//
// Real-time data is collected hourly.
//
// @param request - DescribeDomainStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainStatisticsResponse
func (client *Client) DescribeDomainStatisticsWithOptions(request *DescribeDomainStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of real-time requests for a specified primary domain name.
//
// Description:
//
// Real-time data is collected hourly.
//
// @param request - DescribeDomainStatisticsRequest
//
// @return DescribeDomainStatisticsResponse
func (client *Client) DescribeDomainStatistics(request *DescribeDomainStatisticsRequest) (_result *DescribeDomainStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainStatisticsResponse{}
	_body, _err := client.DescribeDomainStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of requests for all paid domain names in your account.
//
// @param request - DescribeDomainStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainStatisticsSummaryResponse
func (client *Client) DescribeDomainStatisticsSummaryWithOptions(request *DescribeDomainStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainStatisticsSummaryResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of requests for all paid domain names in your account.
//
// @param request - DescribeDomainStatisticsSummaryRequest
//
// @return DescribeDomainStatisticsSummaryResponse
func (client *Client) DescribeDomainStatisticsSummary(request *DescribeDomainStatisticsSummaryRequest) (_result *DescribeDomainStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainStatisticsSummaryResponse{}
	_body, _err := client.DescribeDomainStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of domain names based on specified parameters.
//
// Description:
//
// 1. You can specify a page number (PageNumber) and page size (PageSize) to retrieve a paginated list of domain names.
//
// 2. You can specify a keyword (KeyWord) to query for domain names that contain the specified keyword.
//
// 3. By default, domain names are sorted in descending order of their creation time.
//
// 4. You can specify a domain name group ID (GroupId) to query for domain names in a specific group. This lets you retrieve all domain names or only the domain names that are not assigned to a group.
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
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.Starmark) {
		query["Starmark"] = request.Starmark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomains"),
		Version:     dara.String("2015-01-09"),
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
// Queries a list of domain names based on specified parameters.
//
// Description:
//
// 1. You can specify a page number (PageNumber) and page size (PageSize) to retrieve a paginated list of domain names.
//
// 2. You can specify a keyword (KeyWord) to query for domain names that contain the specified keyword.
//
// 3. By default, domain names are sorted in descending order of their creation time.
//
// 4. You can specify a domain name group ID (GroupId) to query for domain names in a specific group. This lets you retrieve all domain names or only the domain names that are not assigned to a group.
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
// Queries the access policies for an instance.
//
// @param request - DescribeGtmAccessStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmAccessStrategiesResponse
func (client *Client) DescribeGtmAccessStrategiesWithOptions(request *DescribeGtmAccessStrategiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmAccessStrategiesResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeGtmAccessStrategies"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmAccessStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access policies for an instance.
//
// @param request - DescribeGtmAccessStrategiesRequest
//
// @return DescribeGtmAccessStrategiesResponse
func (client *Client) DescribeGtmAccessStrategies(request *DescribeGtmAccessStrategiesRequest) (_result *DescribeGtmAccessStrategiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmAccessStrategiesResponse{}
	_body, _err := client.DescribeGtmAccessStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an access policy based on the policy ID.
//
// @param request - DescribeGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmAccessStrategyResponse
func (client *Client) DescribeGtmAccessStrategyWithOptions(request *DescribeGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an access policy based on the policy ID.
//
// @param request - DescribeGtmAccessStrategyRequest
//
// @return DescribeGtmAccessStrategyResponse
func (client *Client) DescribeGtmAccessStrategy(request *DescribeGtmAccessStrategyRequest) (_result *DescribeGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmAccessStrategyResponse{}
	_body, _err := client.DescribeGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available configurations for an access policy.
//
// @param request - DescribeGtmAccessStrategyAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmAccessStrategyAvailableConfigResponse
func (client *Client) DescribeGtmAccessStrategyAvailableConfigWithOptions(request *DescribeGtmAccessStrategyAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmAccessStrategyAvailableConfigResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmAccessStrategyAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available configurations for an access policy.
//
// @param request - DescribeGtmAccessStrategyAvailableConfigRequest
//
// @return DescribeGtmAccessStrategyAvailableConfigResponse
func (client *Client) DescribeGtmAccessStrategyAvailableConfig(request *DescribeGtmAccessStrategyAvailableConfigRequest) (_result *DescribeGtmAccessStrategyAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmAccessStrategyAvailableConfigResponse{}
	_body, _err := client.DescribeGtmAccessStrategyAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of available alert contact groups for a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmAvailableAlertGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmAvailableAlertGroupResponse
func (client *Client) DescribeGtmAvailableAlertGroupWithOptions(request *DescribeGtmAvailableAlertGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmAvailableAlertGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmAvailableAlertGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmAvailableAlertGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of available alert contact groups for a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmAvailableAlertGroupRequest
//
// @return DescribeGtmAvailableAlertGroupResponse
func (client *Client) DescribeGtmAvailableAlertGroup(request *DescribeGtmAvailableAlertGroupRequest) (_result *DescribeGtmAvailableAlertGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmAvailableAlertGroupResponse{}
	_body, _err := client.DescribeGtmAvailableAlertGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceResponse
func (client *Client) DescribeGtmInstanceWithOptions(request *DescribeGtmInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NeedDetailAttributes) {
		query["NeedDetailAttributes"] = request.NeedDetailAttributes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstance"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmInstanceRequest
//
// @return DescribeGtmInstanceResponse
func (client *Client) DescribeGtmInstance(request *DescribeGtmInstanceRequest) (_result *DescribeGtmInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceResponse{}
	_body, _err := client.DescribeGtmInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the details of an address pool.
//
// @param request - DescribeGtmInstanceAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceAddressPoolResponse
func (client *Client) DescribeGtmInstanceAddressPoolWithOptions(request *DescribeGtmInstanceAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstanceAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the details of an address pool.
//
// @param request - DescribeGtmInstanceAddressPoolRequest
//
// @return DescribeGtmInstanceAddressPoolResponse
func (client *Client) DescribeGtmInstanceAddressPool(request *DescribeGtmInstanceAddressPoolRequest) (_result *DescribeGtmInstanceAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceAddressPoolResponse{}
	_body, _err := client.DescribeGtmInstanceAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the address pools of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmInstanceAddressPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceAddressPoolsResponse
func (client *Client) DescribeGtmInstanceAddressPoolsWithOptions(request *DescribeGtmInstanceAddressPoolsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceAddressPoolsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeGtmInstanceAddressPools"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceAddressPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the address pools of a Global Traffic Manager (GTM) instance.
//
// @param request - DescribeGtmInstanceAddressPoolsRequest
//
// @return DescribeGtmInstanceAddressPoolsResponse
func (client *Client) DescribeGtmInstanceAddressPools(request *DescribeGtmInstanceAddressPoolsRequest) (_result *DescribeGtmInstanceAddressPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceAddressPoolsResponse{}
	_body, _err := client.DescribeGtmInstanceAddressPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current status of an instance.
//
// @param request - DescribeGtmInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceStatusResponse
func (client *Client) DescribeGtmInstanceStatusWithOptions(request *DescribeGtmInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceStatusResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstanceStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current status of an instance.
//
// @param request - DescribeGtmInstanceStatusRequest
//
// @return DescribeGtmInstanceStatusResponse
func (client *Client) DescribeGtmInstanceStatus(request *DescribeGtmInstanceStatusRequest) (_result *DescribeGtmInstanceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceStatusResponse{}
	_body, _err := client.DescribeGtmInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the system-assigned CNAME domain name.
//
// @param request - DescribeGtmInstanceSystemCnameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstanceSystemCnameResponse
func (client *Client) DescribeGtmInstanceSystemCnameWithOptions(request *DescribeGtmInstanceSystemCnameRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstanceSystemCnameResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmInstanceSystemCname"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstanceSystemCnameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the system-assigned CNAME domain name.
//
// @param request - DescribeGtmInstanceSystemCnameRequest
//
// @return DescribeGtmInstanceSystemCnameResponse
func (client *Client) DescribeGtmInstanceSystemCname(request *DescribeGtmInstanceSystemCnameRequest) (_result *DescribeGtmInstanceSystemCnameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstanceSystemCnameResponse{}
	_body, _err := client.DescribeGtmInstanceSystemCnameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Global Traffic Manager (GTM) instances.
//
// @param request - DescribeGtmInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmInstancesResponse
func (client *Client) DescribeGtmInstancesWithOptions(request *DescribeGtmInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NeedDetailAttributes) {
		query["NeedDetailAttributes"] = request.NeedDetailAttributes
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
		Action:      dara.String("DescribeGtmInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Global Traffic Manager (GTM) instances.
//
// @param request - DescribeGtmInstancesRequest
//
// @return DescribeGtmInstancesResponse
func (client *Client) DescribeGtmInstances(request *DescribeGtmInstancesRequest) (_result *DescribeGtmInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmInstancesResponse{}
	_body, _err := client.DescribeGtmInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of logs.
//
// @param request - DescribeGtmLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmLogsResponse
func (client *Client) DescribeGtmLogsWithOptions(request *DescribeGtmLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of logs.
//
// @param request - DescribeGtmLogsRequest
//
// @return DescribeGtmLogsResponse
func (client *Client) DescribeGtmLogs(request *DescribeGtmLogsRequest) (_result *DescribeGtmLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmLogsResponse{}
	_body, _err := client.DescribeGtmLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the available configurations for health checks.
//
// @param request - DescribeGtmMonitorAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmMonitorAvailableConfigResponse
func (client *Client) DescribeGtmMonitorAvailableConfigWithOptions(request *DescribeGtmMonitorAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmMonitorAvailableConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmMonitorAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmMonitorAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the available configurations for health checks.
//
// @param request - DescribeGtmMonitorAvailableConfigRequest
//
// @return DescribeGtmMonitorAvailableConfigResponse
func (client *Client) DescribeGtmMonitorAvailableConfig(request *DescribeGtmMonitorAvailableConfigRequest) (_result *DescribeGtmMonitorAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmMonitorAvailableConfigResponse{}
	_body, _err := client.DescribeGtmMonitorAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the health check configuration for an address pool.
//
// @param request - DescribeGtmMonitorConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmMonitorConfigResponse
func (client *Client) DescribeGtmMonitorConfigWithOptions(request *DescribeGtmMonitorConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmMonitorConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmMonitorConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmMonitorConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the health check configuration for an address pool.
//
// @param request - DescribeGtmMonitorConfigRequest
//
// @return DescribeGtmMonitorConfigResponse
func (client *Client) DescribeGtmMonitorConfig(request *DescribeGtmMonitorConfigRequest) (_result *DescribeGtmMonitorConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmMonitorConfigResponse{}
	_body, _err := client.DescribeGtmMonitorConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a disaster recovery plan.
//
// @param request - DescribeGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmRecoveryPlanResponse
func (client *Client) DescribeGtmRecoveryPlanWithOptions(request *DescribeGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a disaster recovery plan.
//
// @param request - DescribeGtmRecoveryPlanRequest
//
// @return DescribeGtmRecoveryPlanResponse
func (client *Client) DescribeGtmRecoveryPlan(request *DescribeGtmRecoveryPlanRequest) (_result *DescribeGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmRecoveryPlanResponse{}
	_body, _err := client.DescribeGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the available configurations for a disaster recovery plan.
//
// @param request - DescribeGtmRecoveryPlanAvailableConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmRecoveryPlanAvailableConfigResponse
func (client *Client) DescribeGtmRecoveryPlanAvailableConfigWithOptions(request *DescribeGtmRecoveryPlanAvailableConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmRecoveryPlanAvailableConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGtmRecoveryPlanAvailableConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmRecoveryPlanAvailableConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the available configurations for a disaster recovery plan.
//
// @param request - DescribeGtmRecoveryPlanAvailableConfigRequest
//
// @return DescribeGtmRecoveryPlanAvailableConfigResponse
func (client *Client) DescribeGtmRecoveryPlanAvailableConfig(request *DescribeGtmRecoveryPlanAvailableConfigRequest) (_result *DescribeGtmRecoveryPlanAvailableConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmRecoveryPlanAvailableConfigResponse{}
	_body, _err := client.DescribeGtmRecoveryPlanAvailableConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of disaster recovery plans.
//
// @param request - DescribeGtmRecoveryPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGtmRecoveryPlansResponse
func (client *Client) DescribeGtmRecoveryPlansWithOptions(request *DescribeGtmRecoveryPlansRequest, runtime *dara.RuntimeOptions) (_result *DescribeGtmRecoveryPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeGtmRecoveryPlans"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGtmRecoveryPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of disaster recovery plans.
//
// @param request - DescribeGtmRecoveryPlansRequest
//
// @return DescribeGtmRecoveryPlansResponse
func (client *Client) DescribeGtmRecoveryPlans(request *DescribeGtmRecoveryPlansRequest) (_result *DescribeGtmRecoveryPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGtmRecoveryPlansResponse{}
	_body, _err := client.DescribeGtmRecoveryPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of domain names that are attached to an instance.
//
// @param request - DescribeInstanceDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDomainsResponse
func (client *Client) DescribeInstanceDomainsWithOptions(request *DescribeInstanceDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainKeywords) {
		query["DomainKeywords"] = request.DomainKeywords
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribeInstanceDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of domain names that are attached to an instance.
//
// @param request - DescribeInstanceDomainsRequest
//
// @return DescribeInstanceDomainsResponse
func (client *Client) DescribeInstanceDomains(request *DescribeInstanceDomainsRequest) (_result *DescribeInstanceDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceDomainsResponse{}
	_body, _err := client.DescribeInstanceDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a global overview of statistics for public authoritative DNS.
//
// Description:
//
// Real-time data is aggregated hourly.
//
// @param request - DescribeInterAuthStatisticsGlobalOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterAuthStatisticsGlobalOverviewResponse
func (client *Client) DescribeInterAuthStatisticsGlobalOverviewWithOptions(request *DescribeInterAuthStatisticsGlobalOverviewRequest, runtime *dara.RuntimeOptions) (_result *DescribeInterAuthStatisticsGlobalOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OverviewPeriod) {
		query["OverviewPeriod"] = request.OverviewPeriod
	}

	if !dara.IsNil(request.ServerRegion) {
		query["ServerRegion"] = request.ServerRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInterAuthStatisticsGlobalOverview"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInterAuthStatisticsGlobalOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a global overview of statistics for public authoritative DNS.
//
// Description:
//
// Real-time data is aggregated hourly.
//
// @param request - DescribeInterAuthStatisticsGlobalOverviewRequest
//
// @return DescribeInterAuthStatisticsGlobalOverviewResponse
func (client *Client) DescribeInterAuthStatisticsGlobalOverview(request *DescribeInterAuthStatisticsGlobalOverviewRequest) (_result *DescribeInterAuthStatisticsGlobalOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInterAuthStatisticsGlobalOverviewResponse{}
	_body, _err := client.DescribeInterAuthStatisticsGlobalOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Historical statistics for public authoritative DNS resolution
//
// Description:
//
// Real-time data statistics are aggregated hourly.
//
// @param request - DescribeInterAuthStatisticsHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterAuthStatisticsHistoryResponse
func (client *Client) DescribeInterAuthStatisticsHistoryWithOptions(request *DescribeInterAuthStatisticsHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeInterAuthStatisticsHistoryResponse, _err error) {
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

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.Rcode) {
		query["Rcode"] = request.Rcode
	}

	if !dara.IsNil(request.ServerRegion) {
		query["ServerRegion"] = request.ServerRegion
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.StatisticalType) {
		query["StatisticalType"] = request.StatisticalType
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInterAuthStatisticsHistory"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInterAuthStatisticsHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Historical statistics for public authoritative DNS resolution
//
// Description:
//
// Real-time data statistics are aggregated hourly.
//
// @param request - DescribeInterAuthStatisticsHistoryRequest
//
// @return DescribeInterAuthStatisticsHistoryResponse
func (client *Client) DescribeInterAuthStatisticsHistory(request *DescribeInterAuthStatisticsHistoryRequest) (_result *DescribeInterAuthStatisticsHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInterAuthStatisticsHistoryResponse{}
	_body, _err := client.DescribeInterAuthStatisticsHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the summary list of public authoritative parse statistics.
//
// Description:
//
// Real-time data is aggregated by hour.
//
// @param request - DescribeInterAuthStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterAuthStatisticsSummaryResponse
func (client *Client) DescribeInterAuthStatisticsSummaryWithOptions(request *DescribeInterAuthStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeInterAuthStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.GrowType) {
		query["GrowType"] = request.GrowType
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Rcode) {
		query["Rcode"] = request.Rcode
	}

	if !dara.IsNil(request.ServerRegion) {
		query["ServerRegion"] = request.ServerRegion
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.StatisticalType) {
		query["StatisticalType"] = request.StatisticalType
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInterAuthStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInterAuthStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the summary list of public authoritative parse statistics.
//
// Description:
//
// Real-time data is aggregated by hour.
//
// @param request - DescribeInterAuthStatisticsSummaryRequest
//
// @return DescribeInterAuthStatisticsSummaryResponse
func (client *Client) DescribeInterAuthStatisticsSummary(request *DescribeInterAuthStatisticsSummaryRequest) (_result *DescribeInterAuthStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInterAuthStatisticsSummaryResponse{}
	_body, _err := client.DescribeInterAuthStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a zone-level overview of statistics for public authoritative DNS.
//
// @param request - DescribeInterAuthStatisticsZoneOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterAuthStatisticsZoneOverviewResponse
func (client *Client) DescribeInterAuthStatisticsZoneOverviewWithOptions(request *DescribeInterAuthStatisticsZoneOverviewRequest, runtime *dara.RuntimeOptions) (_result *DescribeInterAuthStatisticsZoneOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OverviewPeriod) {
		query["OverviewPeriod"] = request.OverviewPeriod
	}

	if !dara.IsNil(request.ServerRegion) {
		query["ServerRegion"] = request.ServerRegion
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInterAuthStatisticsZoneOverview"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInterAuthStatisticsZoneOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a zone-level overview of statistics for public authoritative DNS.
//
// @param request - DescribeInterAuthStatisticsZoneOverviewRequest
//
// @return DescribeInterAuthStatisticsZoneOverviewResponse
func (client *Client) DescribeInterAuthStatisticsZoneOverview(request *DescribeInterAuthStatisticsZoneOverviewRequest) (_result *DescribeInterAuthStatisticsZoneOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInterAuthStatisticsZoneOverviewResponse{}
	_body, _err := client.DescribeInterAuthStatisticsZoneOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query parsing logs
//
// @param request - DescribeInternetDnsLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInternetDnsLogsResponse
func (client *Client) DescribeInternetDnsLogsWithOptions(request *DescribeInternetDnsLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInternetDnsLogsResponse, _err error) {
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

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Module) {
		query["Module"] = request.Module
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PreciseSort) {
		query["PreciseSort"] = request.PreciseSort
	}

	if !dara.IsNil(request.QueryCondition) {
		query["QueryCondition"] = request.QueryCondition
	}

	if !dara.IsNil(request.RecursionProtocolType) {
		query["RecursionProtocolType"] = request.RecursionProtocolType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInternetDnsLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInternetDnsLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query parsing logs
//
// @param request - DescribeInternetDnsLogsRequest
//
// @return DescribeInternetDnsLogsResponse
func (client *Client) DescribeInternetDnsLogs(request *DescribeInternetDnsLogsRequest) (_result *DescribeInternetDnsLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInternetDnsLogsResponse{}
	_body, _err := client.DescribeInternetDnsLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of cache refresh instances.
//
// @param request - DescribeIspFlushCacheInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIspFlushCacheInstancesResponse
func (client *Client) DescribeIspFlushCacheInstancesWithOptions(request *DescribeIspFlushCacheInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeIspFlushCacheInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIspFlushCacheInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIspFlushCacheInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of cache refresh instances.
//
// @param request - DescribeIspFlushCacheInstancesRequest
//
// @return DescribeIspFlushCacheInstancesResponse
func (client *Client) DescribeIspFlushCacheInstances(request *DescribeIspFlushCacheInstancesRequest) (_result *DescribeIspFlushCacheInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIspFlushCacheInstancesResponse{}
	_body, _err := client.DescribeIspFlushCacheInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the remaining number of cache refresh operations available.
//
// @param request - DescribeIspFlushCacheRemainQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIspFlushCacheRemainQuotaResponse
func (client *Client) DescribeIspFlushCacheRemainQuotaWithOptions(request *DescribeIspFlushCacheRemainQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeIspFlushCacheRemainQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIspFlushCacheRemainQuota"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIspFlushCacheRemainQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the remaining number of cache refresh operations available.
//
// @param request - DescribeIspFlushCacheRemainQuotaRequest
//
// @return DescribeIspFlushCacheRemainQuotaResponse
func (client *Client) DescribeIspFlushCacheRemainQuota(request *DescribeIspFlushCacheRemainQuotaRequest) (_result *DescribeIspFlushCacheRemainQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIspFlushCacheRemainQuotaResponse{}
	_body, _err := client.DescribeIspFlushCacheRemainQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a cache flush task.
//
// @param request - DescribeIspFlushCacheTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIspFlushCacheTaskResponse
func (client *Client) DescribeIspFlushCacheTaskWithOptions(request *DescribeIspFlushCacheTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeIspFlushCacheTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIspFlushCacheTask"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIspFlushCacheTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cache flush task.
//
// @param request - DescribeIspFlushCacheTaskRequest
//
// @return DescribeIspFlushCacheTaskResponse
func (client *Client) DescribeIspFlushCacheTask(request *DescribeIspFlushCacheTaskRequest) (_result *DescribeIspFlushCacheTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIspFlushCacheTaskResponse{}
	_body, _err := client.DescribeIspFlushCacheTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of cache refresh tasks.
//
// @param request - DescribeIspFlushCacheTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIspFlushCacheTasksResponse
func (client *Client) DescribeIspFlushCacheTasksWithOptions(request *DescribeIspFlushCacheTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeIspFlushCacheTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
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
		Action:      dara.String("DescribeIspFlushCacheTasks"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIspFlushCacheTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of cache refresh tasks.
//
// @param request - DescribeIspFlushCacheTasksRequest
//
// @return DescribeIspFlushCacheTasksResponse
func (client *Client) DescribeIspFlushCacheTasks(request *DescribeIspFlushCacheTasksRequest) (_result *DescribeIspFlushCacheTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIspFlushCacheTasksResponse{}
	_body, _err := client.DescribeIspFlushCacheTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the overview of Public DNS user data.
//
// @param request - DescribePdnsAccountSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsAccountSummaryResponse
func (client *Client) DescribePdnsAccountSummaryWithOptions(request *DescribePdnsAccountSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsAccountSummaryResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsAccountSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsAccountSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the overview of Public DNS user data.
//
// @param request - DescribePdnsAccountSummaryRequest
//
// @return DescribePdnsAccountSummaryResponse
func (client *Client) DescribePdnsAccountSummary(request *DescribePdnsAccountSummaryRequest) (_result *DescribePdnsAccountSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsAccountSummaryResponse{}
	_body, _err := client.DescribePdnsAccountSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the details of a public DNS AppKey
//
// @param request - DescribePdnsAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsAppKeyResponse
func (client *Client) DescribePdnsAppKeyWithOptions(request *DescribePdnsAppKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKeyId) {
		query["AppKeyId"] = request.AppKeyId
	}

	if !dara.IsNil(request.AuthCode) {
		query["AuthCode"] = request.AuthCode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsAppKey"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the details of a public DNS AppKey
//
// @param request - DescribePdnsAppKeyRequest
//
// @return DescribePdnsAppKeyResponse
func (client *Client) DescribePdnsAppKey(request *DescribePdnsAppKeyRequest) (_result *DescribePdnsAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsAppKeyResponse{}
	_body, _err := client.DescribePdnsAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Queries the list of AppKeys of Public DNS
//
// @param request - DescribePdnsAppKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsAppKeysResponse
func (client *Client) DescribePdnsAppKeysWithOptions(request *DescribePdnsAppKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsAppKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsAppKeys"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsAppKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries the list of AppKeys of Public DNS
//
// @param request - DescribePdnsAppKeysRequest
//
// @return DescribePdnsAppKeysResponse
func (client *Client) DescribePdnsAppKeys(request *DescribePdnsAppKeysRequest) (_result *DescribePdnsAppKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsAppKeysResponse{}
	_body, _err := client.DescribePdnsAppKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation log list of Public DNS.
//
// @param request - DescribePdnsOperateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsOperateLogsResponse
func (client *Client) DescribePdnsOperateLogsWithOptions(request *DescribePdnsOperateLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsOperateLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsOperateLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsOperateLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation log list of Public DNS.
//
// @param request - DescribePdnsOperateLogsRequest
//
// @return DescribePdnsOperateLogsResponse
func (client *Client) DescribePdnsOperateLogs(request *DescribePdnsOperateLogsRequest) (_result *DescribePdnsOperateLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsOperateLogsResponse{}
	_body, _err := client.DescribePdnsOperateLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request statistics for Public DNS.
//
// @param request - DescribePdnsRequestStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsRequestStatisticResponse
func (client *Client) DescribePdnsRequestStatisticWithOptions(request *DescribePdnsRequestStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsRequestStatisticResponse, _err error) {
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

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsRequestStatistic"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsRequestStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request statistics for Public DNS.
//
// @param request - DescribePdnsRequestStatisticRequest
//
// @return DescribePdnsRequestStatisticResponse
func (client *Client) DescribePdnsRequestStatistic(request *DescribePdnsRequestStatisticRequest) (_result *DescribePdnsRequestStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsRequestStatisticResponse{}
	_body, _err := client.DescribePdnsRequestStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request statistics for a specified subdomain.
//
// @param request - DescribePdnsRequestStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsRequestStatisticsResponse
func (client *Client) DescribePdnsRequestStatisticsWithOptions(request *DescribePdnsRequestStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsRequestStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsRequestStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsRequestStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request statistics for a specified subdomain.
//
// @param request - DescribePdnsRequestStatisticsRequest
//
// @return DescribePdnsRequestStatisticsResponse
func (client *Client) DescribePdnsRequestStatistics(request *DescribePdnsRequestStatisticsRequest) (_result *DescribePdnsRequestStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsRequestStatisticsResponse{}
	_body, _err := client.DescribePdnsRequestStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Public DNS threat logs.
//
// @param request - DescribePdnsThreatLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsThreatLogsResponse
func (client *Client) DescribePdnsThreatLogsWithOptions(request *DescribePdnsThreatLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsThreatLogsResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.ThreatLevel) {
		query["ThreatLevel"] = request.ThreatLevel
	}

	if !dara.IsNil(request.ThreatSourceIp) {
		query["ThreatSourceIp"] = request.ThreatSourceIp
	}

	if !dara.IsNil(request.ThreatType) {
		query["ThreatType"] = request.ThreatType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsThreatLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsThreatLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Public DNS threat logs.
//
// @param request - DescribePdnsThreatLogsRequest
//
// @return DescribePdnsThreatLogsResponse
func (client *Client) DescribePdnsThreatLogs(request *DescribePdnsThreatLogsRequest) (_result *DescribePdnsThreatLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsThreatLogsResponse{}
	_body, _err := client.DescribePdnsThreatLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves Public DNS threat statistics.
//
// @param request - DescribePdnsThreatStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsThreatStatisticResponse
func (client *Client) DescribePdnsThreatStatisticWithOptions(request *DescribePdnsThreatStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsThreatStatisticResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.ThreatSourceIp) {
		query["ThreatSourceIp"] = request.ThreatSourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsThreatStatistic"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsThreatStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves Public DNS threat statistics.
//
// @param request - DescribePdnsThreatStatisticRequest
//
// @return DescribePdnsThreatStatisticResponse
func (client *Client) DescribePdnsThreatStatistic(request *DescribePdnsThreatStatisticRequest) (_result *DescribePdnsThreatStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsThreatStatisticResponse{}
	_body, _err := client.DescribePdnsThreatStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the public DNS threat statistics list.
//
// @param request - DescribePdnsThreatStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsThreatStatisticsResponse
func (client *Client) DescribePdnsThreatStatisticsWithOptions(request *DescribePdnsThreatStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsThreatStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.ThreatLevel) {
		query["ThreatLevel"] = request.ThreatLevel
	}

	if !dara.IsNil(request.ThreatSourceIp) {
		query["ThreatSourceIp"] = request.ThreatSourceIp
	}

	if !dara.IsNil(request.ThreatType) {
		query["ThreatType"] = request.ThreatType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsThreatStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsThreatStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the public DNS threat statistics list.
//
// @param request - DescribePdnsThreatStatisticsRequest
//
// @return DescribePdnsThreatStatisticsResponse
func (client *Client) DescribePdnsThreatStatistics(request *DescribePdnsThreatStatisticsRequest) (_result *DescribePdnsThreatStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsThreatStatisticsResponse{}
	_body, _err := client.DescribePdnsThreatStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of Public DNS UDP IP address ranges.
//
// @param request - DescribePdnsUdpIpSegmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsUdpIpSegmentsResponse
func (client *Client) DescribePdnsUdpIpSegmentsWithOptions(request *DescribePdnsUdpIpSegmentsRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsUdpIpSegmentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Action:      dara.String("DescribePdnsUdpIpSegments"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsUdpIpSegmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of Public DNS UDP IP address ranges.
//
// @param request - DescribePdnsUdpIpSegmentsRequest
//
// @return DescribePdnsUdpIpSegmentsResponse
func (client *Client) DescribePdnsUdpIpSegments(request *DescribePdnsUdpIpSegmentsRequest) (_result *DescribePdnsUdpIpSegmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsUdpIpSegmentsResponse{}
	_body, _err := client.DescribePdnsUdpIpSegmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves user information for Public DNS.
//
// @param request - DescribePdnsUserInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePdnsUserInfoResponse
func (client *Client) DescribePdnsUserInfoWithOptions(request *DescribePdnsUserInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribePdnsUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePdnsUserInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePdnsUserInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves user information for Public DNS.
//
// @param request - DescribePdnsUserInfoRequest
//
// @return DescribePdnsUserInfoResponse
func (client *Client) DescribePdnsUserInfo(request *DescribePdnsUserInfoRequest) (_result *DescribePdnsUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePdnsUserInfoResponse{}
	_body, _err := client.DescribePdnsUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs for the DNS records of a domain name.
//
// @param request - DescribeRecordLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordLogsResponse
func (client *Client) DescribeRecordLogsWithOptions(request *DescribeRecordLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordLogsResponse, _err error) {
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

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs for the DNS records of a domain name.
//
// @param request - DescribeRecordLogsRequest
//
// @return DescribeRecordLogsResponse
func (client *Client) DescribeRecordLogs(request *DescribeRecordLogsRequest) (_result *DescribeRecordLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordLogsResponse{}
	_body, _err := client.DescribeRecordLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries statistics on the request volume for all subdomains of a specified domain name.
//
// @param request - DescribeRecordResolveStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordResolveStatisticsSummaryResponse
func (client *Client) DescribeRecordResolveStatisticsSummaryWithOptions(request *DescribeRecordResolveStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordResolveStatisticsSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordResolveStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordResolveStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics on the request volume for all subdomains of a specified domain name.
//
// @param request - DescribeRecordResolveStatisticsSummaryRequest
//
// @return DescribeRecordResolveStatisticsSummaryResponse
func (client *Client) DescribeRecordResolveStatisticsSummary(request *DescribeRecordResolveStatisticsSummaryRequest) (_result *DescribeRecordResolveStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordResolveStatisticsSummaryResponse{}
	_body, _err := client.DescribeRecordResolveStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries real-time data about DNS requests for a specified subdomain.
//
// Description:
//
// Real-time data is collected hourly.
//
// @param request - DescribeRecordStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordStatisticsResponse
func (client *Client) DescribeRecordStatisticsWithOptions(request *DescribeRecordStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordStatistics"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries real-time data about DNS requests for a specified subdomain.
//
// Description:
//
// Real-time data is collected hourly.
//
// @param request - DescribeRecordStatisticsRequest
//
// @return DescribeRecordStatisticsResponse
func (client *Client) DescribeRecordStatistics(request *DescribeRecordStatisticsRequest) (_result *DescribeRecordStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordStatisticsResponse{}
	_body, _err := client.DescribeRecordStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries statistics on the request volume for all subdomains of a specified domain name.
//
// @param request - DescribeRecordStatisticsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordStatisticsSummaryResponse
func (client *Client) DescribeRecordStatisticsSummaryWithOptions(request *DescribeRecordStatisticsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordStatisticsSummaryResponse, _err error) {
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

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMode) {
		query["SearchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordStatisticsSummary"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordStatisticsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics on the request volume for all subdomains of a specified domain name.
//
// @param request - DescribeRecordStatisticsSummaryRequest
//
// @return DescribeRecordStatisticsSummaryResponse
func (client *Client) DescribeRecordStatisticsSummary(request *DescribeRecordStatisticsSummaryRequest) (_result *DescribeRecordStatisticsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordStatisticsSummaryResponse{}
	_body, _err := client.DescribeRecordStatisticsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes a built-in authoritative DNS record used for recursive resolution.
//
// @param request - DescribeRecursionRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecursionRecordResponse
func (client *Client) DescribeRecursionRecordWithOptions(request *DescribeRecursionRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecursionRecordResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecursionRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecursionRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes a built-in authoritative DNS record used for recursive resolution.
//
// @param request - DescribeRecursionRecordRequest
//
// @return DescribeRecursionRecordResponse
func (client *Client) DescribeRecursionRecord(request *DescribeRecursionRecordRequest) (_result *DescribeRecursionRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecursionRecordResponse{}
	_body, _err := client.DescribeRecursionRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an authoritative zone for recursive resolution.
//
// Description:
//
// Real-time data is measured hourly.
//
// @param request - DescribeRecursionZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecursionZoneResponse
func (client *Client) DescribeRecursionZoneWithOptions(request *DescribeRecursionZoneRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecursionZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecursionZone"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecursionZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an authoritative zone for recursive resolution.
//
// Description:
//
// Real-time data is measured hourly.
//
// @param request - DescribeRecursionZoneRequest
//
// @return DescribeRecursionZoneResponse
func (client *Client) DescribeRecursionZone(request *DescribeRecursionZoneRequest) (_result *DescribeRecursionZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecursionZoneResponse{}
	_body, _err := client.DescribeRecursionZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all DNS records for a specific subdomain based on the specified parameters.
//
// @param request - DescribeSubDomainRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSubDomainRecordsResponse
func (client *Client) DescribeSubDomainRecordsWithOptions(request *DescribeSubDomainRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSubDomainRecordsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSubDomainRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSubDomainRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all DNS records for a specific subdomain based on the specified parameters.
//
// @param request - DescribeSubDomainRecordsRequest
//
// @return DescribeSubDomainRecordsResponse
func (client *Client) DescribeSubDomainRecords(request *DescribeSubDomainRecordsRequest) (_result *DescribeSubDomainRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSubDomainRecordsResponse{}
	_body, _err := client.DescribeSubDomainRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all lines supported by Cloud DNS.
//
// @param request - DescribeSupportLinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportLinesResponse
func (client *Client) DescribeSupportLinesWithOptions(request *DescribeSupportLinesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupportLinesResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupportLines"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupportLinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all lines supported by Cloud DNS.
//
// @param request - DescribeSupportLinesRequest
//
// @return DescribeSupportLinesResponse
func (client *Client) DescribeSupportLines(request *DescribeSupportLinesRequest) (_result *DescribeSupportLinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSupportLinesResponse{}
	_body, _err := client.DescribeSupportLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries existing tags.
//
// @param request - DescribeTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries existing tags.
//
// @param request - DescribeTagsRequest
//
// @return DescribeTagsResponse
func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the domain names transferred between accounts based on the specified input parameters.
//
// @param request - DescribeTransferDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransferDomainsResponse
func (client *Client) DescribeTransferDomainsWithOptions(request *DescribeTransferDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransferDomainsResponse, _err error) {
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

	if !dara.IsNil(request.FromUserId) {
		query["FromUserId"] = request.FromUserId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	if !dara.IsNil(request.TransferType) {
		query["TransferType"] = request.TransferType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTransferDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTransferDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names transferred between accounts based on the specified input parameters.
//
// @param request - DescribeTransferDomainsRequest
//
// @return DescribeTransferDomainsResponse
func (client *Client) DescribeTransferDomains(request *DescribeTransferDomainsRequest) (_result *DescribeTransferDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTransferDomainsResponse{}
	_body, _err := client.DescribeTransferDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a disaster recovery plan.
//
// @param request - ExecuteGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteGtmRecoveryPlanResponse
func (client *Client) ExecuteGtmRecoveryPlanWithOptions(request *ExecuteGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *ExecuteGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a disaster recovery plan.
//
// @param request - ExecuteGtmRecoveryPlanRequest
//
// @return ExecuteGtmRecoveryPlanResponse
func (client *Client) ExecuteGtmRecoveryPlan(request *ExecuteGtmRecoveryPlanRequest) (_result *ExecuteGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteGtmRecoveryPlanResponse{}
	_body, _err := client.ExecuteGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the primary domain name from a specified domain name.
//
// Description:
//
// # For more information about primary and subdomain name levels, see
//
// <props="china">[Domain name levels](https://help.aliyun.com/document_detail/39803.html?spm=a2c4g.2357293.0.0.211f41ffUR1cPb). For example, if you enter `www.abc.com`, the output is abc.com.
//
// <props="intl">[Domain name levels](https://www.alibabacloud.com/help/zh/faq-detail/39803.htm). For example, if you enter `www.abc.com`, the output is abc.com.
//
// @param request - GetMainDomainNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMainDomainNameResponse
func (client *Client) GetMainDomainNameWithOptions(request *GetMainDomainNameRequest, runtime *dara.RuntimeOptions) (_result *GetMainDomainNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InputString) {
		query["InputString"] = request.InputString
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMainDomainName"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMainDomainNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the primary domain name from a specified domain name.
//
// Description:
//
// # For more information about primary and subdomain name levels, see
//
// <props="china">[Domain name levels](https://help.aliyun.com/document_detail/39803.html?spm=a2c4g.2357293.0.0.211f41ffUR1cPb). For example, if you enter `www.abc.com`, the output is abc.com.
//
// <props="intl">[Domain name levels](https://www.alibabacloud.com/help/zh/faq-detail/39803.htm). For example, if you enter `www.abc.com`, the output is abc.com.
//
// @param request - GetMainDomainNameRequest
//
// @return GetMainDomainNameResponse
func (client *Client) GetMainDomainName(request *GetMainDomainNameRequest) (_result *GetMainDomainNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMainDomainNameResponse{}
	_body, _err := client.GetMainDomainNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a TXT record for domain and subdomain verification. This operation supports batch retrieval.
//
// @param request - GetTxtRecordForVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTxtRecordForVerifyResponse
func (client *Client) GetTxtRecordForVerifyWithOptions(request *GetTxtRecordForVerifyRequest, runtime *dara.RuntimeOptions) (_result *GetTxtRecordForVerifyResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTxtRecordForVerify"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTxtRecordForVerifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a TXT record for domain and subdomain verification. This operation supports batch retrieval.
//
// @param request - GetTxtRecordForVerifyRequest
//
// @return GetTxtRecordForVerifyResponse
func (client *Client) GetTxtRecordForVerify(request *GetTxtRecordForVerifyRequest) (_result *GetTxtRecordForVerifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTxtRecordForVerifyResponse{}
	_body, _err := client.GetTxtRecordForVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Agent registration information.
//
// @param request - ListAtiAgentRegisterInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAtiAgentRegisterInfosResponse
func (client *Client) ListAtiAgentRegisterInfosWithOptions(request *ListAtiAgentRegisterInfosRequest, runtime *dara.RuntimeOptions) (_result *ListAtiAgentRegisterInfosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentDisplayName) {
		query["AgentDisplayName"] = request.AgentDisplayName
	}

	if !dara.IsNil(request.AgentHost) {
		query["AgentHost"] = request.AgentHost
	}

	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentVersion) {
		query["AgentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAtiAgentRegisterInfos"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAtiAgentRegisterInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Agent registration information.
//
// @param request - ListAtiAgentRegisterInfosRequest
//
// @return ListAtiAgentRegisterInfosResponse
func (client *Client) ListAtiAgentRegisterInfos(request *ListAtiAgentRegisterInfosRequest) (_result *ListAtiAgentRegisterInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAtiAgentRegisterInfosResponse{}
	_body, _err := client.ListAtiAgentRegisterInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of real-name registered contacts.
//
// @param request - ListAtiChangeLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAtiChangeLogsResponse
func (client *Client) ListAtiChangeLogsWithOptions(request *ListAtiChangeLogsRequest, runtime *dara.RuntimeOptions) (_result *ListAtiChangeLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.OperatorAccount) {
		query["OperatorAccount"] = request.OperatorAccount
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAtiChangeLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAtiChangeLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of real-name registered contacts.
//
// @param request - ListAtiChangeLogsRequest
//
// @return ListAtiChangeLogsResponse
func (client *Client) ListAtiChangeLogs(request *ListAtiChangeLogsRequest) (_result *ListAtiChangeLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAtiChangeLogsResponse{}
	_body, _err := client.ListAtiChangeLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实名注册人列表
//
// @param request - ListAtiRegistrantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAtiRegistrantsResponse
func (client *Client) ListAtiRegistrantsWithOptions(request *ListAtiRegistrantsRequest, runtime *dara.RuntimeOptions) (_result *ListAtiRegistrantsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAtiRegistrants"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAtiRegistrantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实名注册人列表
//
// @param request - ListAtiRegistrantsRequest
//
// @return ListAtiRegistrantsResponse
func (client *Client) ListAtiRegistrants(request *ListAtiRegistrantsRequest) (_result *ListAtiRegistrantsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAtiRegistrantsResponse{}
	_body, _err := client.ListAtiRegistrantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of address pools in Global Traffic Manager (GTM) 3.0.
//
// @param request - ListCloudGtmAddressPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmAddressPoolsResponse
func (client *Client) ListCloudGtmAddressPoolsWithOptions(request *ListCloudGtmAddressPoolsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmAddressPoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolName) {
		query["AddressPoolName"] = request.AddressPoolName
	}

	if !dara.IsNil(request.AddressPoolType) {
		query["AddressPoolType"] = request.AddressPoolType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmAddressPools"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmAddressPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of address pools in Global Traffic Manager (GTM) 3.0.
//
// @param request - ListCloudGtmAddressPoolsRequest
//
// @return ListCloudGtmAddressPoolsResponse
func (client *Client) ListCloudGtmAddressPools(request *ListCloudGtmAddressPoolsRequest) (_result *ListCloudGtmAddressPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmAddressPoolsResponse{}
	_body, _err := client.ListCloudGtmAddressPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of addresses based on the specified parameters.
//
// @param request - ListCloudGtmAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmAddressesResponse
func (client *Client) ListCloudGtmAddressesWithOptions(request *ListCloudGtmAddressesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.MonitorTemplateId) {
		query["MonitorTemplateId"] = request.MonitorTemplateId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmAddresses"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmAddressesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of addresses based on the specified parameters.
//
// @param request - ListCloudGtmAddressesRequest
//
// @return ListCloudGtmAddressesResponse
func (client *Client) ListCloudGtmAddresses(request *ListCloudGtmAddressesRequest) (_result *ListCloudGtmAddressesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmAddressesResponse{}
	_body, _err := client.ListCloudGtmAddressesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of alert logs.
//
// @param request - ListCloudGtmAlertLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmAlertLogsResponse
func (client *Client) ListCloudGtmAlertLogsWithOptions(request *ListCloudGtmAlertLogsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmAlertLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmAlertLogs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmAlertLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of alert logs.
//
// @param request - ListCloudGtmAlertLogsRequest
//
// @return ListCloudGtmAlertLogsResponse
func (client *Client) ListCloudGtmAlertLogs(request *ListCloudGtmAlertLogsRequest) (_result *ListCloudGtmAlertLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmAlertLogsResponse{}
	_body, _err := client.ListCloudGtmAlertLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListCloudGtmAvailableAlertGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmAvailableAlertGroupsResponse
func (client *Client) ListCloudGtmAvailableAlertGroupsWithOptions(request *ListCloudGtmAvailableAlertGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmAvailableAlertGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmAvailableAlertGroups"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmAvailableAlertGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListCloudGtmAvailableAlertGroupsRequest
//
// @return ListCloudGtmAvailableAlertGroupsResponse
func (client *Client) ListCloudGtmAvailableAlertGroups(request *ListCloudGtmAvailableAlertGroupsRequest) (_result *ListCloudGtmAvailableAlertGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmAvailableAlertGroupsResponse{}
	_body, _err := client.ListCloudGtmAvailableAlertGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of instance configurations that contain access domain names and address pool information.
//
// @param request - ListCloudGtmInstanceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmInstanceConfigsResponse
func (client *Client) ListCloudGtmInstanceConfigsWithOptions(request *ListCloudGtmInstanceConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmInstanceConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ScheduleDomainName) {
		query["ScheduleDomainName"] = request.ScheduleDomainName
	}

	if !dara.IsNil(request.ScheduleZoneName) {
		query["ScheduleZoneName"] = request.ScheduleZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmInstanceConfigs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmInstanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of instance configurations that contain access domain names and address pool information.
//
// @param request - ListCloudGtmInstanceConfigsRequest
//
// @return ListCloudGtmInstanceConfigsResponse
func (client *Client) ListCloudGtmInstanceConfigs(request *ListCloudGtmInstanceConfigsRequest) (_result *ListCloudGtmInstanceConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmInstanceConfigsResponse{}
	_body, _err := client.ListCloudGtmInstanceConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of GTM 3.0 instances based on the specified parameters.
//
// @param request - ListCloudGtmInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmInstancesResponse
func (client *Client) ListCloudGtmInstancesWithOptions(request *ListCloudGtmInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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
		Action:      dara.String("ListCloudGtmInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of GTM 3.0 instances based on the specified parameters.
//
// @param request - ListCloudGtmInstancesRequest
//
// @return ListCloudGtmInstancesResponse
func (client *Client) ListCloudGtmInstances(request *ListCloudGtmInstancesRequest) (_result *ListCloudGtmInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmInstancesResponse{}
	_body, _err := client.ListCloudGtmInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the health check monitoring nodes based on the specified input parameters.
//
// @param request - ListCloudGtmMonitorNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmMonitorNodesResponse
func (client *Client) ListCloudGtmMonitorNodesWithOptions(request *ListCloudGtmMonitorNodesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmMonitorNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ShowDisabledNodes) {
		query["ShowDisabledNodes"] = request.ShowDisabledNodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmMonitorNodes"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmMonitorNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the health check monitoring nodes based on the specified input parameters.
//
// @param request - ListCloudGtmMonitorNodesRequest
//
// @return ListCloudGtmMonitorNodesResponse
func (client *Client) ListCloudGtmMonitorNodes(request *ListCloudGtmMonitorNodesRequest) (_result *ListCloudGtmMonitorNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmMonitorNodesResponse{}
	_body, _err := client.ListCloudGtmMonitorNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries health check templates based on the specified parameters.
//
// @param request - ListCloudGtmMonitorTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloudGtmMonitorTemplatesResponse
func (client *Client) ListCloudGtmMonitorTemplatesWithOptions(request *ListCloudGtmMonitorTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListCloudGtmMonitorTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloudGtmMonitorTemplates"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloudGtmMonitorTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries health check templates based on the specified parameters.
//
// @param request - ListCloudGtmMonitorTemplatesRequest
//
// @return ListCloudGtmMonitorTemplatesResponse
func (client *Client) ListCloudGtmMonitorTemplates(request *ListCloudGtmMonitorTemplatesRequest) (_result *ListCloudGtmMonitorTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloudGtmMonitorTemplatesResponse{}
	_body, _err := client.ListCloudGtmMonitorTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the built-in authoritative records for recursive resolution.
//
// @param request - ListRecursionRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecursionRecordsResponse
func (client *Client) ListRecursionRecordsWithOptions(request *ListRecursionRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListRecursionRecordsResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecursionRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecursionRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the built-in authoritative records for recursive resolution.
//
// @param request - ListRecursionRecordsRequest
//
// @return ListRecursionRecordsResponse
func (client *Client) ListRecursionRecords(request *ListRecursionRecordsRequest) (_result *ListRecursionRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecursionRecordsResponse{}
	_body, _err := client.ListRecursionRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the zones that are used for recursive resolution of built-in authoritative domain names.
//
// Description:
//
// - Specify at least ResourceId.N or Tag.N (Tag.N.Key and Tag.N.Value) in your request to identify the resources to retrieve.
//
// - Tag.N is a resource tag that consists of a key-value pair. If you specify only Tag.N.Key, all tag values associated with the tag key are returned. An error is returned if you specify only Tag.N.Value.
//
// - If you specify both Tag.N and ResourceId.N to filter resources, only the resources that are specified by ResourceId.N and match all the specified key-value pairs are returned.
//
// - If you specify multiple tag key-value pairs, the resources that match all of them are returned.
//
// @param request - ListRecursionZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecursionZonesResponse
func (client *Client) ListRecursionZonesWithOptions(request *ListRecursionZonesRequest, runtime *dara.RuntimeOptions) (_result *ListRecursionZonesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecursionZones"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecursionZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the zones that are used for recursive resolution of built-in authoritative domain names.
//
// Description:
//
// - Specify at least ResourceId.N or Tag.N (Tag.N.Key and Tag.N.Value) in your request to identify the resources to retrieve.
//
// - Tag.N is a resource tag that consists of a key-value pair. If you specify only Tag.N.Key, all tag values associated with the tag key are returned. An error is returned if you specify only Tag.N.Value.
//
// - If you specify both Tag.N and ResourceId.N to filter resources, only the resources that are specified by ResourceId.N and match all the specified key-value pairs are returned.
//
// - If you specify multiple tag key-value pairs, the resources that match all of them are returned.
//
// @param request - ListRecursionZonesRequest
//
// @return ListRecursionZonesResponse
func (client *Client) ListRecursionZones(request *ListRecursionZonesRequest) (_result *ListRecursionZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecursionZonesResponse{}
	_body, _err := client.ListRecursionZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resources by tag.
//
// Description:
//
// - To specify the objects to retrieve, your request must contain at least one of the following parameters: `ResourceId.N` or `Tag.N` (which consists of Tag.N.Key and Tag.N.Value).
//
// - Tag.N is a resource tag that consists of a key-value pair. If you specify only Tag.N.Key, all tag values associated with that tag key are returned. An error occurs if you specify only Tag.N.Value.
//
// - If you specify both Tag.N and ResourceId.N to filter resources, only the resources that are specified by ResourceId.N and match all the specified tag key-value pairs are returned.
//
// - If you specify multiple tag key-value pairs, only the resources that have all the specified key-value pairs are returned.
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
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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
		Version:     dara.String("2015-01-09"),
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
// Queries resources by tag.
//
// Description:
//
// - To specify the objects to retrieve, your request must contain at least one of the following parameters: `ResourceId.N` or `Tag.N` (which consists of Tag.N.Key and Tag.N.Value).
//
// - Tag.N is a resource tag that consists of a key-value pair. If you specify only Tag.N.Key, all tag values associated with that tag key are returned. An error occurs if you specify only Tag.N.Value.
//
// - If you specify both Tag.N and ResourceId.N to filter resources, only the resources that are specified by ResourceId.N and match all the specified tag key-value pairs are returned.
//
// - If you specify multiple tag key-value pairs, only the resources that have all the specified key-value pairs are returned.
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
// Changes the DNS servers for a domain name from a third-party provider to Alibaba Cloud DNS.
//
// Description:
//
// After the operation is successful, the DNS servers are changed to Alibaba Cloud DNS servers. The names of these new servers end with hichina.com.
//
// > **Prerequisite: This operation applies to domain names that are registered with Alibaba Cloud and currently use third-party DNS servers.**
//
// @param request - ModifyHichinaDomainDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHichinaDomainDNSResponse
func (client *Client) ModifyHichinaDomainDNSWithOptions(request *ModifyHichinaDomainDNSRequest, runtime *dara.RuntimeOptions) (_result *ModifyHichinaDomainDNSResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHichinaDomainDNS"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHichinaDomainDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the DNS servers for a domain name from a third-party provider to Alibaba Cloud DNS.
//
// Description:
//
// After the operation is successful, the DNS servers are changed to Alibaba Cloud DNS servers. The names of these new servers end with hichina.com.
//
// > **Prerequisite: This operation applies to domain names that are registered with Alibaba Cloud and currently use third-party DNS servers.**
//
// @param request - ModifyHichinaDomainDNSRequest
//
// @return ModifyHichinaDomainDNSResponse
func (client *Client) ModifyHichinaDomainDNS(request *ModifyHichinaDomainDNSRequest) (_result *ModifyHichinaDomainDNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHichinaDomainDNSResponse{}
	_body, _err := client.ModifyHichinaDomainDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a domain name to a different resource group.
//
// @param request - MoveDomainResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveDomainResourceGroupResponse
func (client *Client) MoveDomainResourceGroupWithOptions(request *MoveDomainResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveDomainResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveDomainResourceGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveDomainResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a domain name to a different resource group.
//
// @param request - MoveDomainResourceGroupRequest
//
// @return MoveDomainResourceGroupResponse
func (client *Client) MoveDomainResourceGroup(request *MoveDomainResourceGroupRequest) (_result *MoveDomainResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveDomainResourceGroupResponse{}
	_body, _err := client.MoveDomainResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a Global Traffic Manager (GTM) instance to a new resource group.
//
// @param request - MoveGtmResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveGtmResourceGroupResponse
func (client *Client) MoveGtmResourceGroupWithOptions(request *MoveGtmResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveGtmResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveGtmResourceGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveGtmResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a Global Traffic Manager (GTM) instance to a new resource group.
//
// @param request - MoveGtmResourceGroupRequest
//
// @return MoveGtmResourceGroupResponse
func (client *Client) MoveGtmResourceGroup(request *MoveGtmResourceGroupRequest) (_result *MoveGtmResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveGtmResourceGroupResponse{}
	_body, _err := client.MoveGtmResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or deletes domain names and DNS records in batches.
//
// Description:
//
// Use this operation for batch DNS tasks that do not require immediate execution.
//
// @param request - OperateBatchDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateBatchDomainResponse
func (client *Client) OperateBatchDomainWithOptions(request *OperateBatchDomainRequest, runtime *dara.RuntimeOptions) (_result *OperateBatchDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainRecordInfo) {
		query["DomainRecordInfo"] = request.DomainRecordInfo
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateBatchDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateBatchDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or deletes domain names and DNS records in batches.
//
// Description:
//
// Use this operation for batch DNS tasks that do not require immediate execution.
//
// @param request - OperateBatchDomainRequest
//
// @return OperateBatchDomainResponse
func (client *Client) OperateBatchDomain(request *OperateBatchDomainRequest) (_result *OperateBatchDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateBatchDomainResponse{}
	_body, _err := client.OperateBatchDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Pause Public DNS Service
//
// @param request - PausePdnsServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PausePdnsServiceResponse
func (client *Client) PausePdnsServiceWithOptions(request *PausePdnsServiceRequest, runtime *dara.RuntimeOptions) (_result *PausePdnsServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PausePdnsService"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PausePdnsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Pause Public DNS Service
//
// @param request - PausePdnsServiceRequest
//
// @return PausePdnsServiceResponse
func (client *Client) PausePdnsService(request *PausePdnsServiceRequest) (_result *PausePdnsServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PausePdnsServiceResponse{}
	_body, _err := client.PausePdnsServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Previews a disaster recovery plan.
//
// @param request - PreviewGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewGtmRecoveryPlanResponse
func (client *Client) PreviewGtmRecoveryPlanWithOptions(request *PreviewGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *PreviewGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Previews a disaster recovery plan.
//
// @param request - PreviewGtmRecoveryPlanRequest
//
// @return PreviewGtmRecoveryPlanResponse
func (client *Client) PreviewGtmRecoveryPlan(request *PreviewGtmRecoveryPlanRequest) (_result *PreviewGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreviewGtmRecoveryPlanResponse{}
	_body, _err := client.PreviewGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Public DNS AppKey
//
// @param request - RemovePdnsAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePdnsAppKeyResponse
func (client *Client) RemovePdnsAppKeyWithOptions(request *RemovePdnsAppKeyRequest, runtime *dara.RuntimeOptions) (_result *RemovePdnsAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKeyId) {
		query["AppKeyId"] = request.AppKeyId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePdnsAppKey"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePdnsAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Public DNS AppKey
//
// @param request - RemovePdnsAppKeyRequest
//
// @return RemovePdnsAppKeyResponse
func (client *Client) RemovePdnsAppKey(request *RemovePdnsAppKeyRequest) (_result *RemovePdnsAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePdnsAppKeyResponse{}
	_body, _err := client.RemovePdnsAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Remove Public DNS UDP IP Segment
//
// @param request - RemovePdnsUdpIpSegmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePdnsUdpIpSegmentResponse
func (client *Client) RemovePdnsUdpIpSegmentWithOptions(request *RemovePdnsUdpIpSegmentRequest, runtime *dara.RuntimeOptions) (_result *RemovePdnsUdpIpSegmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemovePdnsUdpIpSegment"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemovePdnsUdpIpSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Remove Public DNS UDP IP Segment
//
// @param request - RemovePdnsUdpIpSegmentRequest
//
// @return RemovePdnsUdpIpSegmentResponse
func (client *Client) RemovePdnsUdpIpSegment(request *RemovePdnsUdpIpSegmentRequest) (_result *RemovePdnsUdpIpSegmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemovePdnsUdpIpSegmentResponse{}
	_body, _err := client.RemovePdnsUdpIpSegmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the serverHold status of a specified domain name.
//
// Description:
//
// ## Description
//
// - This operation removes the serverHold status of a specified domain name.
//
// @param request - RemoveRspDomainServerHoldStatusForGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveRspDomainServerHoldStatusForGatewayResponse
func (client *Client) RemoveRspDomainServerHoldStatusForGatewayWithOptions(request *RemoveRspDomainServerHoldStatusForGatewayRequest, runtime *dara.RuntimeOptions) (_result *RemoveRspDomainServerHoldStatusForGatewayResponse, _err error) {
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

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.StatusMsg) {
		query["StatusMsg"] = request.StatusMsg
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveRspDomainServerHoldStatusForGateway"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveRspDomainServerHoldStatusForGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the serverHold status of a specified domain name.
//
// Description:
//
// ## Description
//
// - This operation removes the serverHold status of a specified domain name.
//
// @param request - RemoveRspDomainServerHoldStatusForGatewayRequest
//
// @return RemoveRspDomainServerHoldStatusForGatewayResponse
func (client *Client) RemoveRspDomainServerHoldStatusForGateway(request *RemoveRspDomainServerHoldStatusForGatewayRequest) (_result *RemoveRspDomainServerHoldStatusForGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveRspDomainServerHoldStatusForGatewayResponse{}
	_body, _err := client.RemoveRspDomainServerHoldStatusForGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the serverHold status for a specified domain name.
//
// Description:
//
// ## Description
//
// - This operation removes the serverHold status for a specified domain name.
//
// @param request - RemoveRspDomainServerHoldStatusForGatewayOteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveRspDomainServerHoldStatusForGatewayOteResponse
func (client *Client) RemoveRspDomainServerHoldStatusForGatewayOteWithOptions(request *RemoveRspDomainServerHoldStatusForGatewayOteRequest, runtime *dara.RuntimeOptions) (_result *RemoveRspDomainServerHoldStatusForGatewayOteResponse, _err error) {
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

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.StatusMsg) {
		query["StatusMsg"] = request.StatusMsg
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveRspDomainServerHoldStatusForGatewayOte"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveRspDomainServerHoldStatusForGatewayOteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the serverHold status for a specified domain name.
//
// Description:
//
// ## Description
//
// - This operation removes the serverHold status for a specified domain name.
//
// @param request - RemoveRspDomainServerHoldStatusForGatewayOteRequest
//
// @return RemoveRspDomainServerHoldStatusForGatewayOteResponse
func (client *Client) RemoveRspDomainServerHoldStatusForGatewayOte(request *RemoveRspDomainServerHoldStatusForGatewayOteRequest) (_result *RemoveRspDomainServerHoldStatusForGatewayOteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveRspDomainServerHoldStatusForGatewayOteResponse{}
	_body, _err := client.RemoveRspDomainServerHoldStatusForGatewayOteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replaces the addresses in a specified address pool.
//
// @param tmpReq - ReplaceCloudGtmAddressPoolAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceCloudGtmAddressPoolAddressResponse
func (client *Client) ReplaceCloudGtmAddressPoolAddressWithOptions(tmpReq *ReplaceCloudGtmAddressPoolAddressRequest, runtime *dara.RuntimeOptions) (_result *ReplaceCloudGtmAddressPoolAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReplaceCloudGtmAddressPoolAddressShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Addresses) {
		request.AddressesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Addresses, dara.String("Addresses"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.AddressesShrink) {
		query["Addresses"] = request.AddressesShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceCloudGtmAddressPoolAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceCloudGtmAddressPoolAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces the addresses in a specified address pool.
//
// @param request - ReplaceCloudGtmAddressPoolAddressRequest
//
// @return ReplaceCloudGtmAddressPoolAddressResponse
func (client *Client) ReplaceCloudGtmAddressPoolAddress(request *ReplaceCloudGtmAddressPoolAddressRequest) (_result *ReplaceCloudGtmAddressPoolAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReplaceCloudGtmAddressPoolAddressResponse{}
	_body, _err := client.ReplaceCloudGtmAddressPoolAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replaces the address pools associated with an instance.
//
// @param tmpReq - ReplaceCloudGtmInstanceConfigAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceCloudGtmInstanceConfigAddressPoolResponse
func (client *Client) ReplaceCloudGtmInstanceConfigAddressPoolWithOptions(tmpReq *ReplaceCloudGtmInstanceConfigAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *ReplaceCloudGtmInstanceConfigAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AddressPools) {
		request.AddressPoolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AddressPools, dara.String("AddressPools"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolsShrink) {
		query["AddressPools"] = request.AddressPoolsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceCloudGtmInstanceConfigAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceCloudGtmInstanceConfigAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces the address pools associated with an instance.
//
// @param request - ReplaceCloudGtmInstanceConfigAddressPoolRequest
//
// @return ReplaceCloudGtmInstanceConfigAddressPoolResponse
func (client *Client) ReplaceCloudGtmInstanceConfigAddressPool(request *ReplaceCloudGtmInstanceConfigAddressPoolRequest) (_result *ReplaceCloudGtmInstanceConfigAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReplaceCloudGtmInstanceConfigAddressPoolResponse{}
	_body, _err := client.ReplaceCloudGtmInstanceConfigAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Resume Public DNS Service
//
// @param request - ResumePdnsServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumePdnsServiceResponse
func (client *Client) ResumePdnsServiceWithOptions(request *ResumePdnsServiceRequest, runtime *dara.RuntimeOptions) (_result *ResumePdnsServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumePdnsService"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumePdnsServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Resume Public DNS Service
//
// @param request - ResumePdnsServiceRequest
//
// @return ResumePdnsServiceResponse
func (client *Client) ResumePdnsService(request *ResumePdnsServiceRequest) (_result *ResumePdnsServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumePdnsServiceResponse{}
	_body, _err := client.ResumePdnsServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a domain name.
//
// Description:
//
// Before you can retrieve a domain name, you must verify it using a TXT record. Use this operation in conjunction with the <props="china">[Generate a TXT record](https://help.aliyun.com/document_detail/145533.html) <props="intl">[Generate a TXT record](https://www.alibabacloud.com/help/zh/alibaba-cloud-dns/latest/generating-a-txt-record) operation.
//
// @param request - RetrieveDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveDomainResponse
func (client *Client) RetrieveDomainWithOptions(request *RetrieveDomainRequest, runtime *dara.RuntimeOptions) (_result *RetrieveDomainResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetrieveDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetrieveDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a domain name.
//
// Description:
//
// Before you can retrieve a domain name, you must verify it using a TXT record. Use this operation in conjunction with the <props="china">[Generate a TXT record](https://help.aliyun.com/document_detail/145533.html) <props="intl">[Generate a TXT record](https://www.alibabacloud.com/help/zh/alibaba-cloud-dns/latest/generating-a-txt-record) operation.
//
// @param request - RetrieveDomainRequest
//
// @return RetrieveDomainResponse
func (client *Client) RetrieveDomain(request *RetrieveDomainRequest) (_result *RetrieveDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetrieveDomainResponse{}
	_body, _err := client.RetrieveDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the registration information of an Agent.
//
// @param request - RevokeAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeAtiAgentRegisterInfoResponse
func (client *Client) RevokeAtiAgentRegisterInfoWithOptions(request *RevokeAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *RevokeAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	if !dara.IsNil(request.ReasonCode) {
		query["ReasonCode"] = request.ReasonCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the registration information of an Agent.
//
// @param request - RevokeAtiAgentRegisterInfoRequest
//
// @return RevokeAtiAgentRegisterInfoResponse
func (client *Client) RevokeAtiAgentRegisterInfo(request *RevokeAtiAgentRegisterInfoRequest) (_result *RevokeAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeAtiAgentRegisterInfoResponse{}
	_body, _err := client.RevokeAtiAgentRegisterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a real-name registrant.
//
// @param request - RevokeAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeAtiRegistrantResponse
func (client *Client) RevokeAtiRegistrantWithOptions(request *RevokeAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *RevokeAtiRegistrantResponse, _err error) {
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

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a real-name registrant.
//
// @param request - RevokeAtiRegistrantRequest
//
// @return RevokeAtiRegistrantResponse
func (client *Client) RevokeAtiRegistrant(request *RevokeAtiRegistrantRequest) (_result *RevokeAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeAtiRegistrantResponse{}
	_body, _err := client.RevokeAtiRegistrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back a disaster recovery plan.
//
// @param request - RollbackGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackGtmRecoveryPlanResponse
func (client *Client) RollbackGtmRecoveryPlanWithOptions(request *RollbackGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *RollbackGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back a disaster recovery plan.
//
// @param request - RollbackGtmRecoveryPlanRequest
//
// @return RollbackGtmRecoveryPlanResponse
func (client *Client) RollbackGtmRecoveryPlan(request *RollbackGtmRecoveryPlanRequest) (_result *RollbackGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackGtmRecoveryPlanResponse{}
	_body, _err := client.RollbackGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for agents in the Agent Marketplace.
//
// @param request - SearchAtiAgentRegisterInfoMarketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAtiAgentRegisterInfoMarketResponse
func (client *Client) SearchAtiAgentRegisterInfoMarketWithOptions(request *SearchAtiAgentRegisterInfoMarketRequest, runtime *dara.RuntimeOptions) (_result *SearchAtiAgentRegisterInfoMarketResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TrustLevel) {
		query["TrustLevel"] = request.TrustLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchAtiAgentRegisterInfoMarket"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchAtiAgentRegisterInfoMarketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for agents in the Agent Marketplace.
//
// @param request - SearchAtiAgentRegisterInfoMarketRequest
//
// @return SearchAtiAgentRegisterInfoMarketResponse
func (client *Client) SearchAtiAgentRegisterInfoMarket(request *SearchAtiAgentRegisterInfoMarketRequest) (_result *SearchAtiAgentRegisterInfoMarketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchAtiAgentRegisterInfoMarketResponse{}
	_body, _err := client.SearchAtiAgentRegisterInfoMarketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for address pools by name, remarks, or other criteria.
//
// @param request - SearchCloudGtmAddressPoolsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmAddressPoolsResponse
func (client *Client) SearchCloudGtmAddressPoolsWithOptions(request *SearchCloudGtmAddressPoolsRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmAddressPoolsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolName) {
		query["AddressPoolName"] = request.AddressPoolName
	}

	if !dara.IsNil(request.AddressPoolType) {
		query["AddressPoolType"] = request.AddressPoolType
	}

	if !dara.IsNil(request.AvailableStatus) {
		query["AvailableStatus"] = request.AvailableStatus
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCloudGtmAddressPools"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmAddressPoolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for address pools by name, remarks, or other criteria.
//
// @param request - SearchCloudGtmAddressPoolsRequest
//
// @return SearchCloudGtmAddressPoolsResponse
func (client *Client) SearchCloudGtmAddressPools(request *SearchCloudGtmAddressPoolsRequest) (_result *SearchCloudGtmAddressPoolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmAddressPoolsResponse{}
	_body, _err := client.SearchCloudGtmAddressPoolsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for addresses based on criteria such as address name, remarks, referenced health check template, or address ID.
//
// @param request - SearchCloudGtmAddressesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmAddressesResponse
func (client *Client) SearchCloudGtmAddressesWithOptions(request *SearchCloudGtmAddressesRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmAddressesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.AvailableStatus) {
		query["AvailableStatus"] = request.AvailableStatus
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.MonitorTemplateName) {
		query["MonitorTemplateName"] = request.MonitorTemplateName
	}

	if !dara.IsNil(request.NameSearchCondition) {
		query["NameSearchCondition"] = request.NameSearchCondition
	}

	if !dara.IsNil(request.Names) {
		query["Names"] = request.Names
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RemarkSearchCondition) {
		query["RemarkSearchCondition"] = request.RemarkSearchCondition
	}

	if !dara.IsNil(request.Remarks) {
		query["Remarks"] = request.Remarks
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCloudGtmAddresses"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmAddressesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for addresses based on criteria such as address name, remarks, referenced health check template, or address ID.
//
// @param request - SearchCloudGtmAddressesRequest
//
// @return SearchCloudGtmAddressesResponse
func (client *Client) SearchCloudGtmAddresses(request *SearchCloudGtmAddressesRequest) (_result *SearchCloudGtmAddressesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmAddressesResponse{}
	_body, _err := client.SearchCloudGtmAddressesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves instance configurations that match the specified parameters.
//
// @param request - SearchCloudGtmInstanceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmInstanceConfigsResponse
func (client *Client) SearchCloudGtmInstanceConfigsWithOptions(request *SearchCloudGtmInstanceConfigsRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmInstanceConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AvailableStatus) {
		query["AvailableStatus"] = request.AvailableStatus
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.HealthStatus) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ScheduleDomainName) {
		query["ScheduleDomainName"] = request.ScheduleDomainName
	}

	if !dara.IsNil(request.ScheduleZoneName) {
		query["ScheduleZoneName"] = request.ScheduleZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCloudGtmInstanceConfigs"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmInstanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves instance configurations that match the specified parameters.
//
// @param request - SearchCloudGtmInstanceConfigsRequest
//
// @return SearchCloudGtmInstanceConfigsResponse
func (client *Client) SearchCloudGtmInstanceConfigs(request *SearchCloudGtmInstanceConfigsRequest) (_result *SearchCloudGtmInstanceConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmInstanceConfigsResponse{}
	_body, _err := client.SearchCloudGtmInstanceConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation searches for instances based on specified parameters.
//
// @param request - SearchCloudGtmInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmInstancesResponse
func (client *Client) SearchCloudGtmInstancesWithOptions(request *SearchCloudGtmInstancesRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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
		Action:      dara.String("SearchCloudGtmInstances"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation searches for instances based on specified parameters.
//
// @param request - SearchCloudGtmInstancesRequest
//
// @return SearchCloudGtmInstancesResponse
func (client *Client) SearchCloudGtmInstances(request *SearchCloudGtmInstancesRequest) (_result *SearchCloudGtmInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmInstancesResponse{}
	_body, _err := client.SearchCloudGtmInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for health check templates.
//
// @param request - SearchCloudGtmMonitorTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchCloudGtmMonitorTemplatesResponse
func (client *Client) SearchCloudGtmMonitorTemplatesWithOptions(request *SearchCloudGtmMonitorTemplatesRequest, runtime *dara.RuntimeOptions) (_result *SearchCloudGtmMonitorTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchCloudGtmMonitorTemplates"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchCloudGtmMonitorTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for health check templates.
//
// @param request - SearchCloudGtmMonitorTemplatesRequest
//
// @return SearchCloudGtmMonitorTemplatesResponse
func (client *Client) SearchCloudGtmMonitorTemplates(request *SearchCloudGtmMonitorTemplatesRequest) (_result *SearchCloudGtmMonitorTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchCloudGtmMonitorTemplatesResponse{}
	_body, _err := client.SearchCloudGtmMonitorTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for built-in authoritative DNS records used for recursive resolution.
//
// @param request - SearchRecursionRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchRecursionRecordsResponse
func (client *Client) SearchRecursionRecordsWithOptions(request *SearchRecursionRecordsRequest, runtime *dara.RuntimeOptions) (_result *SearchRecursionRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchRecursionRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchRecursionRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for built-in authoritative DNS records used for recursive resolution.
//
// @param request - SearchRecursionRecordsRequest
//
// @return SearchRecursionRecordsResponse
func (client *Client) SearchRecursionRecords(request *SearchRecursionRecordsRequest) (_result *SearchRecursionRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchRecursionRecordsResponse{}
	_body, _err := client.SearchRecursionRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for zones of built-in authoritative domain names used for recursive resolution.
//
// Description:
//
// - To retrieve resources, you must specify at least `ResourceId.N` or `Tag.N` (`Tag.N.Key` and `Tag.N.Value`) in the request.
//
// - `Tag.N` is a resource tag that consists of a key-value pair. If you specify only `Tag.N.Key`, all tag values associated with that key are returned. If you specify only `Tag.N.Value`, an error is returned.
//
// - If you specify both `Tag.N` and `ResourceId.N`, the operation returns only the resources that are identified by `ResourceId.N` and match all the specified tag key-value pairs.
//
// - If you specify multiple tag key-value pairs, only resources that match all of them are returned.
//
// @param tmpReq - SearchRecursionZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchRecursionZonesResponse
func (client *Client) SearchRecursionZonesWithOptions(tmpReq *SearchRecursionZonesRequest, runtime *dara.RuntimeOptions) (_result *SearchRecursionZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchRecursionZonesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EffectiveScopes) {
		request.EffectiveScopesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EffectiveScopes, dara.String("EffectiveScopes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EffectiveScopesShrink) {
		query["EffectiveScopes"] = request.EffectiveScopesShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ZoneName) {
		query["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchRecursionZones"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchRecursionZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for zones of built-in authoritative domain names used for recursive resolution.
//
// Description:
//
// - To retrieve resources, you must specify at least `ResourceId.N` or `Tag.N` (`Tag.N.Key` and `Tag.N.Value`) in the request.
//
// - `Tag.N` is a resource tag that consists of a key-value pair. If you specify only `Tag.N.Key`, all tag values associated with that key are returned. If you specify only `Tag.N.Value`, an error is returned.
//
// - If you specify both `Tag.N` and `ResourceId.N`, the operation returns only the resources that are identified by `ResourceId.N` and match all the specified tag key-value pairs.
//
// - If you specify multiple tag key-value pairs, only resources that match all of them are returned.
//
// @param request - SearchRecursionZonesRequest
//
// @return SearchRecursionZonesResponse
func (client *Client) SearchRecursionZones(request *SearchRecursionZonesRequest) (_result *SearchRecursionZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchRecursionZonesResponse{}
	_body, _err := client.SearchRecursionZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the network traffic analysis feature for a Global Traffic Manager (GTM) instance. After this feature is enabled, you can view resolution logs and statistical reports for the domain name. You can also use the intelligent alerting feature based on abnormal metrics, such as resolution success rate and sudden changes in queries per second (QPS). This improves the observability and operations and maintenance (O&M) efficiency of the GTM instance.
//
// @param request - SetCloudGtmInstanceConfigLogSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCloudGtmInstanceConfigLogSwitchResponse
func (client *Client) SetCloudGtmInstanceConfigLogSwitchWithOptions(request *SetCloudGtmInstanceConfigLogSwitchRequest, runtime *dara.RuntimeOptions) (_result *SetCloudGtmInstanceConfigLogSwitchResponse, _err error) {
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

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCloudGtmInstanceConfigLogSwitch"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCloudGtmInstanceConfigLogSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the network traffic analysis feature for a Global Traffic Manager (GTM) instance. After this feature is enabled, you can view resolution logs and statistical reports for the domain name. You can also use the intelligent alerting feature based on abnormal metrics, such as resolution success rate and sudden changes in queries per second (QPS). This improves the observability and operations and maintenance (O&M) efficiency of the GTM instance.
//
// @param request - SetCloudGtmInstanceConfigLogSwitchRequest
//
// @return SetCloudGtmInstanceConfigLogSwitchResponse
func (client *Client) SetCloudGtmInstanceConfigLogSwitch(request *SetCloudGtmInstanceConfigLogSwitchRequest) (_result *SetCloudGtmInstanceConfigLogSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCloudGtmInstanceConfigLogSwitchResponse{}
	_body, _err := client.SetCloudGtmInstanceConfigLogSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the weight configuration.
//
// @param request - SetDNSSLBStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDNSSLBStatusResponse
func (client *Client) SetDNSSLBStatusWithOptions(request *SetDNSSLBStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDNSSLBStatusResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.Open) {
		query["Open"] = request.Open
	}

	if !dara.IsNil(request.SubDomain) {
		query["SubDomain"] = request.SubDomain
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDNSSLBStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDNSSLBStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the weight configuration.
//
// @param request - SetDNSSLBStatusRequest
//
// @return SetDNSSLBStatusResponse
func (client *Client) SetDNSSLBStatus(request *SetDNSSLBStatusRequest) (_result *SetDNSSLBStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDNSSLBStatusResponse{}
	_body, _err := client.SetDNSSLBStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the access mode for an access policy.
//
// Description:
//
// **
//
// @param request - SetDnsGtmAccessModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDnsGtmAccessModeResponse
func (client *Client) SetDnsGtmAccessModeWithOptions(request *SetDnsGtmAccessModeRequest, runtime *dara.RuntimeOptions) (_result *SetDnsGtmAccessModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessMode) {
		query["AccessMode"] = request.AccessMode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDnsGtmAccessMode"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDnsGtmAccessModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the access mode for an access policy.
//
// Description:
//
// **
//
// @param request - SetDnsGtmAccessModeRequest
//
// @return SetDnsGtmAccessModeResponse
func (client *Client) SetDnsGtmAccessMode(request *SetDnsGtmAccessModeRequest) (_result *SetDnsGtmAccessModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDnsGtmAccessModeResponse{}
	_body, _err := client.SetDnsGtmAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the health check status for an address pool.
//
// @param request - SetDnsGtmMonitorStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDnsGtmMonitorStatusResponse
func (client *Client) SetDnsGtmMonitorStatusWithOptions(request *SetDnsGtmMonitorStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDnsGtmMonitorStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDnsGtmMonitorStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDnsGtmMonitorStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the health check status for an address pool.
//
// @param request - SetDnsGtmMonitorStatusRequest
//
// @return SetDnsGtmMonitorStatusResponse
func (client *Client) SetDnsGtmMonitorStatus(request *SetDnsGtmMonitorStatusRequest) (_result *SetDnsGtmMonitorStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDnsGtmMonitorStatusResponse{}
	_body, _err := client.SetDnsGtmMonitorStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables Domain Name System Security Extensions (DNSSEC) for a domain name. This feature is available only to users of paid Alibaba Cloud DNS.
//
// @param request - SetDomainDnssecStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainDnssecStatusResponse
func (client *Client) SetDomainDnssecStatusWithOptions(request *SetDomainDnssecStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDomainDnssecStatusResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainDnssecStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainDnssecStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables Domain Name System Security Extensions (DNSSEC) for a domain name. This feature is available only to users of paid Alibaba Cloud DNS.
//
// @param request - SetDomainDnssecStatusRequest
//
// @return SetDomainDnssecStatusResponse
func (client *Client) SetDomainDnssecStatus(request *SetDomainDnssecStatusRequest) (_result *SetDomainDnssecStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDomainDnssecStatusResponse{}
	_body, _err := client.SetDomainDnssecStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the status of a DNS record.
//
// @param request - SetDomainRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainRecordStatusResponse
func (client *Client) SetDomainRecordStatusWithOptions(request *SetDomainRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDomainRecordStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainRecordStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainRecordStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the status of a DNS record.
//
// @param request - SetDomainRecordStatusRequest
//
// @return SetDomainRecordStatusResponse
func (client *Client) SetDomainRecordStatus(request *SetDomainRecordStatusRequest) (_result *SetDomainRecordStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDomainRecordStatusResponse{}
	_body, _err := client.SetDomainRecordStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the address pool switchover policy based on the request parameters.
//
// @param request - SetGtmAccessModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetGtmAccessModeResponse
func (client *Client) SetGtmAccessModeWithOptions(request *SetGtmAccessModeRequest, runtime *dara.RuntimeOptions) (_result *SetGtmAccessModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessMode) {
		query["AccessMode"] = request.AccessMode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetGtmAccessMode"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetGtmAccessModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the address pool switchover policy based on the request parameters.
//
// @param request - SetGtmAccessModeRequest
//
// @return SetGtmAccessModeResponse
func (client *Client) SetGtmAccessMode(request *SetGtmAccessModeRequest) (_result *SetGtmAccessModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetGtmAccessModeResponse{}
	_body, _err := client.SetGtmAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the health check status of an address pool.
//
// @param request - SetGtmMonitorStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetGtmMonitorStatusResponse
func (client *Client) SetGtmMonitorStatusWithOptions(request *SetGtmMonitorStatusRequest, runtime *dara.RuntimeOptions) (_result *SetGtmMonitorStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetGtmMonitorStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetGtmMonitorStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the health check status of an address pool.
//
// @param request - SetGtmMonitorStatusRequest
//
// @return SetGtmMonitorStatusResponse
func (client *Client) SetGtmMonitorStatus(request *SetGtmMonitorStatusRequest) (_result *SetGtmMonitorStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetGtmMonitorStatusResponse{}
	_body, _err := client.SetGtmMonitorStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits Agent registration information.
//
// @param request - SubmitAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAtiAgentRegisterInfoResponse
func (client *Client) SubmitAtiAgentRegisterInfoWithOptions(request *SubmitAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *SubmitAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IdentityCsr) {
		query["IdentityCsr"] = request.IdentityCsr
	}

	if !dara.IsNil(request.ServerCertPem) {
		query["ServerCertPem"] = request.ServerCertPem
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits Agent registration information.
//
// @param request - SubmitAtiAgentRegisterInfoRequest
//
// @return SubmitAtiAgentRegisterInfoResponse
func (client *Client) SubmitAtiAgentRegisterInfo(request *SubmitAtiAgentRegisterInfoRequest) (_result *SubmitAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAtiAgentRegisterInfoResponse{}
	_body, _err := client.SubmitAtiAgentRegisterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a cache refresh task.
//
// @param request - SubmitIspFlushCacheTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIspFlushCacheTaskResponse
func (client *Client) SubmitIspFlushCacheTaskWithOptions(request *SubmitIspFlushCacheTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitIspFlushCacheTaskResponse, _err error) {
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

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIspFlushCacheTask"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIspFlushCacheTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a cache refresh task.
//
// @param request - SubmitIspFlushCacheTaskRequest
//
// @return SubmitIspFlushCacheTaskResponse
func (client *Client) SubmitIspFlushCacheTask(request *SubmitIspFlushCacheTaskRequest) (_result *SubmitIspFlushCacheTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitIspFlushCacheTaskResponse{}
	_body, _err := client.SubmitIspFlushCacheTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches the access strategy mode of an instance.
//
// @param request - SwitchDnsGtmInstanceStrategyModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchDnsGtmInstanceStrategyModeResponse
func (client *Client) SwitchDnsGtmInstanceStrategyModeWithOptions(request *SwitchDnsGtmInstanceStrategyModeRequest, runtime *dara.RuntimeOptions) (_result *SwitchDnsGtmInstanceStrategyModeResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyMode) {
		query["StrategyMode"] = request.StrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchDnsGtmInstanceStrategyMode"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchDnsGtmInstanceStrategyModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the access strategy mode of an instance.
//
// @param request - SwitchDnsGtmInstanceStrategyModeRequest
//
// @return SwitchDnsGtmInstanceStrategyModeResponse
func (client *Client) SwitchDnsGtmInstanceStrategyMode(request *SwitchDnsGtmInstanceStrategyModeRequest) (_result *SwitchDnsGtmInstanceStrategyModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchDnsGtmInstanceStrategyModeResponse{}
	_body, _err := client.SwitchDnsGtmInstanceStrategyModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or modifies tags for resources.
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
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Version:     dara.String("2015-01-09"),
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
// Adds or modifies tags for resources.
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
// Batch transfers DNS permissions for multiple domain names to a specified execution account.
//
// @param request - TransferDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferDomainResponse
func (client *Client) TransferDomainWithOptions(request *TransferDomainRequest, runtime *dara.RuntimeOptions) (_result *TransferDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch transfers DNS permissions for multiple domain names to a specified execution account.
//
// @param request - TransferDomainRequest
//
// @return TransferDomainResponse
func (client *Client) TransferDomain(request *TransferDomainRequest) (_result *TransferDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferDomainResponse{}
	_body, _err := client.TransferDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches domain names from a paid Alibaba Cloud DNS instance.
//
// Description:
//
// An instance with an ID that starts with \\`dns-\\` is a new version instance. New version instances support attaching multiple domain names. You can call an API operation to attach domain names directly to an instance. An error is returned if the number of domain names exceeds the instance limit.
//
// An instance with an ID that does not start with \\`dns-\\` is a legacy instance. Legacy instances support only one domain name. Therefore, if you call this operation for an instance that already has a domain name attached, the domain name is replaced.
//
// @param request - UnbindInstanceDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindInstanceDomainsResponse
func (client *Client) UnbindInstanceDomainsWithOptions(request *UnbindInstanceDomainsRequest, runtime *dara.RuntimeOptions) (_result *UnbindInstanceDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindInstanceDomains"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindInstanceDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches domain names from a paid Alibaba Cloud DNS instance.
//
// Description:
//
// An instance with an ID that starts with \\`dns-\\` is a new version instance. New version instances support attaching multiple domain names. You can call an API operation to attach domain names directly to an instance. An error is returned if the number of domain names exceeds the instance limit.
//
// An instance with an ID that does not start with \\`dns-\\` is a legacy instance. Legacy instances support only one domain name. Therefore, if you call this operation for an instance that already has a domain name attached, the domain name is replaced.
//
// @param request - UnbindInstanceDomainsRequest
//
// @return UnbindInstanceDomainsResponse
func (client *Client) UnbindInstanceDomains(request *UnbindInstanceDomainsRequest) (_result *UnbindInstanceDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindInstanceDomainsResponse{}
	_body, _err := client.UnbindInstanceDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
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
		Version:     dara.String("2015-01-09"),
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
// Removes tags from resources.
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
// # Modify AppKey State
//
// @param request - UpdateAppKeyStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppKeyStateResponse
func (client *Client) UpdateAppKeyStateWithOptions(request *UpdateAppKeyStateRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppKeyStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKeyId) {
		query["AppKeyId"] = request.AppKeyId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppKeyState"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppKeyStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify AppKey State
//
// @param request - UpdateAppKeyStateRequest
//
// @return UpdateAppKeyStateResponse
func (client *Client) UpdateAppKeyState(request *UpdateAppKeyStateRequest) (_result *UpdateAppKeyStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppKeyStateResponse{}
	_body, _err := client.UpdateAppKeyStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the registration information of an Agent.
//
// @param tmpReq - UpdateAtiAgentRegisterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAtiAgentRegisterInfoResponse
func (client *Client) UpdateAtiAgentRegisterInfoWithOptions(tmpReq *UpdateAtiAgentRegisterInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateAtiAgentRegisterInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAtiAgentRegisterInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Endpoints) {
		request.EndpointsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Endpoints, dara.String("Endpoints"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentDescription) {
		query["AgentDescription"] = request.AgentDescription
	}

	if !dara.IsNil(request.AgentDisplayName) {
		query["AgentDisplayName"] = request.AgentDisplayName
	}

	if !dara.IsNil(request.AgentHost) {
		query["AgentHost"] = request.AgentHost
	}

	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.AgentVersion) {
		query["AgentVersion"] = request.AgentVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndpointsShrink) {
		query["Endpoints"] = request.EndpointsShrink
	}

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAtiAgentRegisterInfo"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAtiAgentRegisterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the registration information of an Agent.
//
// @param request - UpdateAtiAgentRegisterInfoRequest
//
// @return UpdateAtiAgentRegisterInfoResponse
func (client *Client) UpdateAtiAgentRegisterInfo(request *UpdateAtiAgentRegisterInfoRequest) (_result *UpdateAtiAgentRegisterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAtiAgentRegisterInfoResponse{}
	_body, _err := client.UpdateAtiAgentRegisterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新告警设置
//
// @param request - UpdateAtiAlertSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAtiAlertSettingsResponse
func (client *Client) UpdateAtiAlertSettingsWithOptions(request *UpdateAtiAlertSettingsRequest, runtime *dara.RuntimeOptions) (_result *UpdateAtiAlertSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.AlertGroup) {
		query["AlertGroup"] = request.AlertGroup
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAtiAlertSettings"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAtiAlertSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新告警设置
//
// @param request - UpdateAtiAlertSettingsRequest
//
// @return UpdateAtiAlertSettingsResponse
func (client *Client) UpdateAtiAlertSettings(request *UpdateAtiAlertSettingsRequest) (_result *UpdateAtiAlertSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAtiAlertSettingsResponse{}
	_body, _err := client.UpdateAtiAlertSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a registrant profile.
//
// @param request - UpdateAtiRegistrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAtiRegistrantResponse
func (client *Client) UpdateAtiRegistrantWithOptions(request *UpdateAtiRegistrantRequest, runtime *dara.RuntimeOptions) (_result *UpdateAtiRegistrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cc) {
		query["Cc"] = request.Cc
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DocumentCode) {
		query["DocumentCode"] = request.DocumentCode
	}

	if !dara.IsNil(request.DocumentImage) {
		query["DocumentImage"] = request.DocumentImage
	}

	if !dara.IsNil(request.DocumentType) {
		query["DocumentType"] = request.DocumentType
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Phone) {
		query["Phone"] = request.Phone
	}

	if !dara.IsNil(request.RegistrantId) {
		query["RegistrantId"] = request.RegistrantId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Street) {
		query["Street"] = request.Street
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAtiRegistrant"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAtiRegistrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a registrant profile.
//
// @param request - UpdateAtiRegistrantRequest
//
// @return UpdateAtiRegistrantResponse
func (client *Client) UpdateAtiRegistrant(request *UpdateAtiRegistrantRequest) (_result *UpdateAtiRegistrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAtiRegistrantResponse{}
	_body, _err := client.UpdateAtiRegistrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic configuration of a specified address, such as the address name, type, and value.
//
// @param tmpReq - UpdateCloudGtmAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressResponse
func (client *Client) UpdateCloudGtmAddressWithOptions(tmpReq *UpdateCloudGtmAddressRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudGtmAddressShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.HealthTasks) {
		request.HealthTasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HealthTasks, dara.String("HealthTasks"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.AttributeInfo) {
		query["AttributeInfo"] = request.AttributeInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.HealthJudgement) {
		query["HealthJudgement"] = request.HealthJudgement
	}

	if !dara.IsNil(request.HealthTasksShrink) {
		query["HealthTasks"] = request.HealthTasksShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddress"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic configuration of a specified address, such as the address name, type, and value.
//
// @param request - UpdateCloudGtmAddressRequest
//
// @return UpdateCloudGtmAddressResponse
func (client *Client) UpdateCloudGtmAddress(request *UpdateCloudGtmAddressRequest) (_result *UpdateCloudGtmAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressResponse{}
	_body, _err := client.UpdateCloudGtmAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enabled status of an address.
//
// Description:
//
// - The service status of an address is **active*	- if the address is **enabled*	- and its health check status is **Normal**.
//
// - The service status of an address is **unavailable*	- if the address is **disabled*	- or its health check status is **abnormal**.
//
// @param request - UpdateCloudGtmAddressEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressEnableStatusResponse
func (client *Client) UpdateCloudGtmAddressEnableStatusWithOptions(request *UpdateCloudGtmAddressEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressEnableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enabled status of an address.
//
// Description:
//
// - The service status of an address is **active*	- if the address is **enabled*	- and its health check status is **Normal**.
//
// - The service status of an address is **unavailable*	- if the address is **disabled*	- or its health check status is **abnormal**.
//
// @param request - UpdateCloudGtmAddressEnableStatusRequest
//
// @return UpdateCloudGtmAddressEnableStatusResponse
func (client *Client) UpdateCloudGtmAddressEnableStatus(request *UpdateCloudGtmAddressEnableStatusRequest) (_result *UpdateCloudGtmAddressEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressEnableStatusResponse{}
	_body, _err := client.UpdateCloudGtmAddressEnableStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the failover method for an address based on the specified parameters.
//
// @param request - UpdateCloudGtmAddressManualAvailableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressManualAvailableStatusResponse
func (client *Client) UpdateCloudGtmAddressManualAvailableStatusWithOptions(request *UpdateCloudGtmAddressManualAvailableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressManualAvailableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.AvailableMode) {
		query["AvailableMode"] = request.AvailableMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ManualAvailableStatus) {
		query["ManualAvailableStatus"] = request.ManualAvailableStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressManualAvailableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressManualAvailableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the failover method for an address based on the specified parameters.
//
// @param request - UpdateCloudGtmAddressManualAvailableStatusRequest
//
// @return UpdateCloudGtmAddressManualAvailableStatusResponse
func (client *Client) UpdateCloudGtmAddressManualAvailableStatus(request *UpdateCloudGtmAddressManualAvailableStatusRequest) (_result *UpdateCloudGtmAddressManualAvailableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressManualAvailableStatusResponse{}
	_body, _err := client.UpdateCloudGtmAddressManualAvailableStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic configuration of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolBasicConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressPoolBasicConfigResponse
func (client *Client) UpdateCloudGtmAddressPoolBasicConfigWithOptions(request *UpdateCloudGtmAddressPoolBasicConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressPoolBasicConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.AddressPoolName) {
		query["AddressPoolName"] = request.AddressPoolName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.HealthJudgement) {
		query["HealthJudgement"] = request.HealthJudgement
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressPoolBasicConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressPoolBasicConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic configuration of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolBasicConfigRequest
//
// @return UpdateCloudGtmAddressPoolBasicConfigResponse
func (client *Client) UpdateCloudGtmAddressPoolBasicConfig(request *UpdateCloudGtmAddressPoolBasicConfigRequest) (_result *UpdateCloudGtmAddressPoolBasicConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressPoolBasicConfigResponse{}
	_body, _err := client.UpdateCloudGtmAddressPoolBasicConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enabled status of an address pool.
//
// Description:
//
// - An address pool is considered **active*	- if it is **enabled*	- and its health check status is **Normal**.
//
// - An address pool is considered **unavailable*	- if it is **disabled*	- or its health check status is **abnormal**.
//
// @param request - UpdateCloudGtmAddressPoolEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressPoolEnableStatusResponse
func (client *Client) UpdateCloudGtmAddressPoolEnableStatusWithOptions(request *UpdateCloudGtmAddressPoolEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressPoolEnableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressPoolEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressPoolEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enabled status of an address pool.
//
// Description:
//
// - An address pool is considered **active*	- if it is **enabled*	- and its health check status is **Normal**.
//
// - An address pool is considered **unavailable*	- if it is **disabled*	- or its health check status is **abnormal**.
//
// @param request - UpdateCloudGtmAddressPoolEnableStatusRequest
//
// @return UpdateCloudGtmAddressPoolEnableStatusResponse
func (client *Client) UpdateCloudGtmAddressPoolEnableStatus(request *UpdateCloudGtmAddressPoolEnableStatusRequest) (_result *UpdateCloudGtmAddressPoolEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressPoolEnableStatusResponse{}
	_body, _err := client.UpdateCloudGtmAddressPoolEnableStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the load balancing policy of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolLbStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressPoolLbStrategyResponse
func (client *Client) UpdateCloudGtmAddressPoolLbStrategyWithOptions(request *UpdateCloudGtmAddressPoolLbStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressPoolLbStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressLbStrategy) {
		query["AddressLbStrategy"] = request.AddressLbStrategy
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.SequenceLbStrategyMode) {
		query["SequenceLbStrategyMode"] = request.SequenceLbStrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressPoolLbStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressPoolLbStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the load balancing policy of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolLbStrategyRequest
//
// @return UpdateCloudGtmAddressPoolLbStrategyResponse
func (client *Client) UpdateCloudGtmAddressPoolLbStrategy(request *UpdateCloudGtmAddressPoolLbStrategyRequest) (_result *UpdateCloudGtmAddressPoolLbStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressPoolLbStrategyResponse{}
	_body, _err := client.UpdateCloudGtmAddressPoolLbStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressPoolRemarkResponse
func (client *Client) UpdateCloudGtmAddressPoolRemarkWithOptions(request *UpdateCloudGtmAddressPoolRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressPoolRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolId) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressPoolRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressPoolRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks of an address pool.
//
// @param request - UpdateCloudGtmAddressPoolRemarkRequest
//
// @return UpdateCloudGtmAddressPoolRemarkResponse
func (client *Client) UpdateCloudGtmAddressPoolRemark(request *UpdateCloudGtmAddressPoolRemarkRequest) (_result *UpdateCloudGtmAddressPoolRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressPoolRemarkResponse{}
	_body, _err := client.UpdateCloudGtmAddressPoolRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks for an address in Global Traffic Manager (GTM) 3.0.
//
// @param request - UpdateCloudGtmAddressRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmAddressRemarkResponse
func (client *Client) UpdateCloudGtmAddressRemarkWithOptions(request *UpdateCloudGtmAddressRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmAddressRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressId) {
		query["AddressId"] = request.AddressId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmAddressRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmAddressRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks for an address in Global Traffic Manager (GTM) 3.0.
//
// @param request - UpdateCloudGtmAddressRemarkRequest
//
// @return UpdateCloudGtmAddressRemarkResponse
func (client *Client) UpdateCloudGtmAddressRemark(request *UpdateCloudGtmAddressRemarkRequest) (_result *UpdateCloudGtmAddressRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmAddressRemarkResponse{}
	_body, _err := client.UpdateCloudGtmAddressRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新全局流量管理告警设置
//
// @param tmpReq - UpdateCloudGtmGlobalAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmGlobalAlertResponse
func (client *Client) UpdateCloudGtmGlobalAlertWithOptions(tmpReq *UpdateCloudGtmGlobalAlertRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmGlobalAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudGtmGlobalAlertShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlertConfig) {
		request.AlertConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertConfig, dara.String("AlertConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AlertGroup) {
		request.AlertGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertGroup, dara.String("AlertGroup"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AlertConfigShrink) {
		query["AlertConfig"] = request.AlertConfigShrink
	}

	if !dara.IsNil(request.AlertGroupShrink) {
		query["AlertGroup"] = request.AlertGroupShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmGlobalAlert"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmGlobalAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新全局流量管理告警设置
//
// @param request - UpdateCloudGtmGlobalAlertRequest
//
// @return UpdateCloudGtmGlobalAlertResponse
func (client *Client) UpdateCloudGtmGlobalAlert(request *UpdateCloudGtmGlobalAlertRequest) (_result *UpdateCloudGtmGlobalAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmGlobalAlertResponse{}
	_body, _err := client.UpdateCloudGtmGlobalAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateCloudGtmInstanceConfigAlertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigAlertResponse
func (client *Client) UpdateCloudGtmInstanceConfigAlertWithOptions(tmpReq *UpdateCloudGtmInstanceConfigAlertRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigAlertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudGtmInstanceConfigAlertShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlertConfig) {
		request.AlertConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertConfig, dara.String("AlertConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AlertGroup) {
		request.AlertGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertGroup, dara.String("AlertGroup"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AlertConfigShrink) {
		query["AlertConfig"] = request.AlertConfigShrink
	}

	if !dara.IsNil(request.AlertGroupShrink) {
		query["AlertGroup"] = request.AlertGroupShrink
	}

	if !dara.IsNil(request.AlertMode) {
		query["AlertMode"] = request.AlertMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigAlert"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCloudGtmInstanceConfigAlertRequest
//
// @return UpdateCloudGtmInstanceConfigAlertResponse
func (client *Client) UpdateCloudGtmInstanceConfigAlert(request *UpdateCloudGtmInstanceConfigAlertRequest) (_result *UpdateCloudGtmInstanceConfigAlertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigAlertResponse{}
	_body, _err := client.UpdateCloudGtmInstanceConfigAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the global TTL configuration of a Global Traffic Manager (GTM) 3.0 instance based on the specified parameters.
//
// @param request - UpdateCloudGtmInstanceConfigBasicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigBasicResponse
func (client *Client) UpdateCloudGtmInstanceConfigBasicWithOptions(request *UpdateCloudGtmInstanceConfigBasicRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigBasicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScheduleHostname) {
		query["ScheduleHostname"] = request.ScheduleHostname
	}

	if !dara.IsNil(request.ScheduleZoneName) {
		query["ScheduleZoneName"] = request.ScheduleZoneName
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigBasic"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigBasicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the global TTL configuration of a Global Traffic Manager (GTM) 3.0 instance based on the specified parameters.
//
// @param request - UpdateCloudGtmInstanceConfigBasicRequest
//
// @return UpdateCloudGtmInstanceConfigBasicResponse
func (client *Client) UpdateCloudGtmInstanceConfigBasic(request *UpdateCloudGtmInstanceConfigBasicRequest) (_result *UpdateCloudGtmInstanceConfigBasicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigBasicResponse{}
	_body, _err := client.UpdateCloudGtmInstanceConfigBasicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enablement status of an instance configuration based on the input parameters.
//
// Description:
//
// - If a domain name is **enabled*	- and its health status is **Normal**, the service associated with the access domain name is **active**.
//
// - If a domain name is **disabled*	- or its health status is **abnormal**, the service associated with the access domain name is **unavailable**.
//
// @param request - UpdateCloudGtmInstanceConfigEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigEnableStatusResponse
func (client *Client) UpdateCloudGtmInstanceConfigEnableStatusWithOptions(request *UpdateCloudGtmInstanceConfigEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigEnableStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enablement status of an instance configuration based on the input parameters.
//
// Description:
//
// - If a domain name is **enabled*	- and its health status is **Normal**, the service associated with the access domain name is **active**.
//
// - If a domain name is **disabled*	- or its health status is **abnormal**, the service associated with the access domain name is **unavailable**.
//
// @param request - UpdateCloudGtmInstanceConfigEnableStatusRequest
//
// @return UpdateCloudGtmInstanceConfigEnableStatusResponse
func (client *Client) UpdateCloudGtmInstanceConfigEnableStatus(request *UpdateCloudGtmInstanceConfigEnableStatusRequest) (_result *UpdateCloudGtmInstanceConfigEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigEnableStatusResponse{}
	_body, _err := client.UpdateCloudGtmInstanceConfigEnableStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the load balancing policy for an instance configuration.
//
// @param request - UpdateCloudGtmInstanceConfigLbStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigLbStrategyResponse
func (client *Client) UpdateCloudGtmInstanceConfigLbStrategyWithOptions(request *UpdateCloudGtmInstanceConfigLbStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigLbStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AddressPoolLbStrategy) {
		query["AddressPoolLbStrategy"] = request.AddressPoolLbStrategy
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SequenceLbStrategyMode) {
		query["SequenceLbStrategyMode"] = request.SequenceLbStrategyMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigLbStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigLbStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the load balancing policy for an instance configuration.
//
// @param request - UpdateCloudGtmInstanceConfigLbStrategyRequest
//
// @return UpdateCloudGtmInstanceConfigLbStrategyResponse
func (client *Client) UpdateCloudGtmInstanceConfigLbStrategy(request *UpdateCloudGtmInstanceConfigLbStrategyRequest) (_result *UpdateCloudGtmInstanceConfigLbStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigLbStrategyResponse{}
	_body, _err := client.UpdateCloudGtmInstanceConfigLbStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks for an instance configuration.
//
// @param request - UpdateCloudGtmInstanceConfigRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceConfigRemarkResponse
func (client *Client) UpdateCloudGtmInstanceConfigRemarkWithOptions(request *UpdateCloudGtmInstanceConfigRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceConfigRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceConfigRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceConfigRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks for an instance configuration.
//
// @param request - UpdateCloudGtmInstanceConfigRemarkRequest
//
// @return UpdateCloudGtmInstanceConfigRemarkResponse
func (client *Client) UpdateCloudGtmInstanceConfigRemark(request *UpdateCloudGtmInstanceConfigRemarkRequest) (_result *UpdateCloudGtmInstanceConfigRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceConfigRemarkResponse{}
	_body, _err := client.UpdateCloudGtmInstanceConfigRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateCloudGtmInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmInstanceNameResponse
func (client *Client) UpdateCloudGtmInstanceNameWithOptions(request *UpdateCloudGtmInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmInstanceNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmInstanceName"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCloudGtmInstanceNameRequest
//
// @return UpdateCloudGtmInstanceNameResponse
func (client *Client) UpdateCloudGtmInstanceName(request *UpdateCloudGtmInstanceNameRequest) (_result *UpdateCloudGtmInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmInstanceNameResponse{}
	_body, _err := client.UpdateCloudGtmInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a health check template.
//
// @param tmpReq - UpdateCloudGtmMonitorTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmMonitorTemplateResponse
func (client *Client) UpdateCloudGtmMonitorTemplateWithOptions(tmpReq *UpdateCloudGtmMonitorTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmMonitorTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCloudGtmMonitorTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IspCityNodes) {
		request.IspCityNodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IspCityNodes, dara.String("IspCityNodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.ExtendInfo) {
		query["ExtendInfo"] = request.ExtendInfo
	}

	if !dara.IsNil(request.FailureRate) {
		query["FailureRate"] = request.FailureRate
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNodesShrink) {
		query["IspCityNodes"] = request.IspCityNodesShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmMonitorTemplate"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmMonitorTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a health check template.
//
// @param request - UpdateCloudGtmMonitorTemplateRequest
//
// @return UpdateCloudGtmMonitorTemplateResponse
func (client *Client) UpdateCloudGtmMonitorTemplate(request *UpdateCloudGtmMonitorTemplateRequest) (_result *UpdateCloudGtmMonitorTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmMonitorTemplateResponse{}
	_body, _err := client.UpdateCloudGtmMonitorTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateCloudGtmMonitorTemplateRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudGtmMonitorTemplateRemarkResponse
func (client *Client) UpdateCloudGtmMonitorTemplateRemarkWithOptions(request *UpdateCloudGtmMonitorTemplateRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudGtmMonitorTemplateRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudGtmMonitorTemplateRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudGtmMonitorTemplateRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateCloudGtmMonitorTemplateRemarkRequest
//
// @return UpdateCloudGtmMonitorTemplateRemarkResponse
func (client *Client) UpdateCloudGtmMonitorTemplateRemark(request *UpdateCloudGtmMonitorTemplateRemarkRequest) (_result *UpdateCloudGtmMonitorTemplateRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudGtmMonitorTemplateRemarkResponse{}
	_body, _err := client.UpdateCloudGtmMonitorTemplateRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom line specified by its unique ID.
//
// Description:
//
// For each IP range, the value of EndIp must be greater than or equal to the value of StartIp.
//
// The IP ranges of IP ranges cannot overlap across all custom lines for a domain name.
//
// @param request - UpdateCustomLineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomLineResponse
func (client *Client) UpdateCustomLineWithOptions(request *UpdateCustomLineRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomLineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpSegment) {
		query["IpSegment"] = request.IpSegment
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LineId) {
		query["LineId"] = request.LineId
	}

	if !dara.IsNil(request.LineName) {
		query["LineName"] = request.LineName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomLine"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomLineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom line specified by its unique ID.
//
// Description:
//
// For each IP range, the value of EndIp must be greater than or equal to the value of StartIp.
//
// The IP ranges of IP ranges cannot overlap across all custom lines for a domain name.
//
// @param request - UpdateCustomLineRequest
//
// @return UpdateCustomLineResponse
func (client *Client) UpdateCustomLine(request *UpdateCustomLineRequest) (_result *UpdateCustomLineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomLineResponse{}
	_body, _err := client.UpdateCustomLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the weight of a DNS record based on the specified parameters.
//
// @param request - UpdateDNSSLBWeightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDNSSLBWeightResponse
func (client *Client) UpdateDNSSLBWeightWithOptions(request *UpdateDNSSLBWeightRequest, runtime *dara.RuntimeOptions) (_result *UpdateDNSSLBWeightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDNSSLBWeight"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDNSSLBWeightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the weight of a DNS record based on the specified parameters.
//
// @param request - UpdateDNSSLBWeightRequest
//
// @return UpdateDNSSLBWeightResponse
func (client *Client) UpdateDNSSLBWeight(request *UpdateDNSSLBWeightRequest) (_result *UpdateDNSSLBWeightResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDNSSLBWeightResponse{}
	_body, _err := client.UpdateDNSSLBWeightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a DNS authoritative proxy domain.
//
// @param request - UpdateDnsCacheDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsCacheDomainResponse
func (client *Client) UpdateDnsCacheDomainWithOptions(request *UpdateDnsCacheDomainRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsCacheDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheTtlMax) {
		query["CacheTtlMax"] = request.CacheTtlMax
	}

	if !dara.IsNil(request.CacheTtlMin) {
		query["CacheTtlMin"] = request.CacheTtlMin
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceDnsServer) {
		query["SourceDnsServer"] = request.SourceDnsServer
	}

	if !dara.IsNil(request.SourceEdns) {
		query["SourceEdns"] = request.SourceEdns
	}

	if !dara.IsNil(request.SourceProtocol) {
		query["SourceProtocol"] = request.SourceProtocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsCacheDomain"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsCacheDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a DNS authoritative proxy domain.
//
// @param request - UpdateDnsCacheDomainRequest
//
// @return UpdateDnsCacheDomainResponse
func (client *Client) UpdateDnsCacheDomain(request *UpdateDnsCacheDomainRequest) (_result *UpdateDnsCacheDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsCacheDomainResponse{}
	_body, _err := client.UpdateDnsCacheDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remark for a domain name in the DNS cache.
//
// @param request - UpdateDnsCacheDomainRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsCacheDomainRemarkResponse
func (client *Client) UpdateDnsCacheDomainRemarkWithOptions(request *UpdateDnsCacheDomainRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsCacheDomainRemarkResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsCacheDomainRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsCacheDomainRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remark for a domain name in the DNS cache.
//
// @param request - UpdateDnsCacheDomainRemarkRequest
//
// @return UpdateDnsCacheDomainRemarkResponse
func (client *Client) UpdateDnsCacheDomainRemark(request *UpdateDnsCacheDomainRemarkRequest) (_result *UpdateDnsCacheDomainRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsCacheDomainRemarkResponse{}
	_body, _err := client.UpdateDnsCacheDomainRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an existing access policy.
//
// @param request - UpdateDnsGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsGtmAccessStrategyResponse
func (client *Client) UpdateDnsGtmAccessStrategyWithOptions(request *UpdateDnsGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessMode) {
		query["AccessMode"] = request.AccessMode
	}

	if !dara.IsNil(request.DefaultAddrPool) {
		query["DefaultAddrPool"] = request.DefaultAddrPool
	}

	if !dara.IsNil(request.DefaultAddrPoolType) {
		query["DefaultAddrPoolType"] = request.DefaultAddrPoolType
	}

	if !dara.IsNil(request.DefaultLatencyOptimization) {
		query["DefaultLatencyOptimization"] = request.DefaultLatencyOptimization
	}

	if !dara.IsNil(request.DefaultLbaStrategy) {
		query["DefaultLbaStrategy"] = request.DefaultLbaStrategy
	}

	if !dara.IsNil(request.DefaultMaxReturnAddrNum) {
		query["DefaultMaxReturnAddrNum"] = request.DefaultMaxReturnAddrNum
	}

	if !dara.IsNil(request.DefaultMinAvailableAddrNum) {
		query["DefaultMinAvailableAddrNum"] = request.DefaultMinAvailableAddrNum
	}

	if !dara.IsNil(request.FailoverAddrPool) {
		query["FailoverAddrPool"] = request.FailoverAddrPool
	}

	if !dara.IsNil(request.FailoverAddrPoolType) {
		query["FailoverAddrPoolType"] = request.FailoverAddrPoolType
	}

	if !dara.IsNil(request.FailoverLatencyOptimization) {
		query["FailoverLatencyOptimization"] = request.FailoverLatencyOptimization
	}

	if !dara.IsNil(request.FailoverLbaStrategy) {
		query["FailoverLbaStrategy"] = request.FailoverLbaStrategy
	}

	if !dara.IsNil(request.FailoverMaxReturnAddrNum) {
		query["FailoverMaxReturnAddrNum"] = request.FailoverMaxReturnAddrNum
	}

	if !dara.IsNil(request.FailoverMinAvailableAddrNum) {
		query["FailoverMinAvailableAddrNum"] = request.FailoverMinAvailableAddrNum
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Lines) {
		query["Lines"] = request.Lines
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an existing access policy.
//
// @param request - UpdateDnsGtmAccessStrategyRequest
//
// @return UpdateDnsGtmAccessStrategyResponse
func (client *Client) UpdateDnsGtmAccessStrategy(request *UpdateDnsGtmAccessStrategyRequest) (_result *UpdateDnsGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsGtmAccessStrategyResponse{}
	_body, _err := client.UpdateDnsGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of an address pool.
//
// @param request - UpdateDnsGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsGtmAddressPoolResponse
func (client *Client) UpdateDnsGtmAddressPoolWithOptions(request *UpdateDnsGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addr) {
		query["Addr"] = request.Addr
	}

	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LbaStrategy) {
		query["LbaStrategy"] = request.LbaStrategy
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of an address pool.
//
// @param request - UpdateDnsGtmAddressPoolRequest
//
// @return UpdateDnsGtmAddressPoolResponse
func (client *Client) UpdateDnsGtmAddressPool(request *UpdateDnsGtmAddressPoolRequest) (_result *UpdateDnsGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsGtmAddressPoolResponse{}
	_body, _err := client.UpdateDnsGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a Global Traffic Manager (GTM) instance.
//
// @param request - UpdateDnsGtmInstanceGlobalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsGtmInstanceGlobalConfigResponse
func (client *Client) UpdateDnsGtmInstanceGlobalConfigWithOptions(request *UpdateDnsGtmInstanceGlobalConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsGtmInstanceGlobalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.AlertGroup) {
		query["AlertGroup"] = request.AlertGroup
	}

	if !dara.IsNil(request.CnameType) {
		query["CnameType"] = request.CnameType
	}

	if !dara.IsNil(request.ForceUpdate) {
		query["ForceUpdate"] = request.ForceUpdate
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PublicCnameMode) {
		query["PublicCnameMode"] = request.PublicCnameMode
	}

	if !dara.IsNil(request.PublicRr) {
		query["PublicRr"] = request.PublicRr
	}

	if !dara.IsNil(request.PublicUserDomainName) {
		query["PublicUserDomainName"] = request.PublicUserDomainName
	}

	if !dara.IsNil(request.PublicZoneName) {
		query["PublicZoneName"] = request.PublicZoneName
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsGtmInstanceGlobalConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsGtmInstanceGlobalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a Global Traffic Manager (GTM) instance.
//
// @param request - UpdateDnsGtmInstanceGlobalConfigRequest
//
// @return UpdateDnsGtmInstanceGlobalConfigResponse
func (client *Client) UpdateDnsGtmInstanceGlobalConfig(request *UpdateDnsGtmInstanceGlobalConfigRequest) (_result *UpdateDnsGtmInstanceGlobalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsGtmInstanceGlobalConfigResponse{}
	_body, _err := client.UpdateDnsGtmInstanceGlobalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a health check configuration.
//
// @param request - UpdateDnsGtmMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDnsGtmMonitorResponse
func (client *Client) UpdateDnsGtmMonitorWithOptions(request *UpdateDnsGtmMonitorRequest, runtime *dara.RuntimeOptions) (_result *UpdateDnsGtmMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDnsGtmMonitor"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDnsGtmMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a health check configuration.
//
// @param request - UpdateDnsGtmMonitorRequest
//
// @return UpdateDnsGtmMonitorResponse
func (client *Client) UpdateDnsGtmMonitor(request *UpdateDnsGtmMonitorRequest) (_result *UpdateDnsGtmMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDnsGtmMonitorResponse{}
	_body, _err := client.UpdateDnsGtmMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of a domain name group.
//
// Description:
//
// This operation modifies the name of an existing domain name group.
//
// @param request - UpdateDomainGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainGroupResponse
func (client *Client) UpdateDomainGroupWithOptions(request *UpdateDomainGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainGroup"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a domain name group.
//
// Description:
//
// This operation modifies the name of an existing domain name group.
//
// @param request - UpdateDomainGroupRequest
//
// @return UpdateDomainGroupResponse
func (client *Client) UpdateDomainGroup(request *UpdateDomainGroupRequest) (_result *UpdateDomainGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainGroupResponse{}
	_body, _err := client.UpdateDomainGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a DNS record based on the specified parameters.
//
// @param request - UpdateDomainRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainRecordResponse
func (client *Client) UpdateDomainRecordWithOptions(request *UpdateDomainRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RR) {
		query["RR"] = request.RR
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.TTL) {
		query["TTL"] = request.TTL
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a DNS record based on the specified parameters.
//
// @param request - UpdateDomainRecordRequest
//
// @return UpdateDomainRecordResponse
func (client *Client) UpdateDomainRecord(request *UpdateDomainRecordRequest) (_result *UpdateDomainRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainRecordResponse{}
	_body, _err := client.UpdateDomainRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks of a DNS record.
//
// @param request - UpdateDomainRecordRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainRecordRemarkResponse
func (client *Client) UpdateDomainRecordRemarkWithOptions(request *UpdateDomainRecordRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainRecordRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.UserClientIp) {
		query["UserClientIp"] = request.UserClientIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainRecordRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainRecordRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks of a DNS record.
//
// @param request - UpdateDomainRecordRemarkRequest
//
// @return UpdateDomainRecordRemarkResponse
func (client *Client) UpdateDomainRecordRemark(request *UpdateDomainRecordRemarkRequest) (_result *UpdateDomainRecordRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainRecordRemarkResponse{}
	_body, _err := client.UpdateDomainRecordRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks for a domain name.
//
// @param request - UpdateDomainRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainRemarkResponse
func (client *Client) UpdateDomainRemarkWithOptions(request *UpdateDomainRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainRemarkResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks for a domain name.
//
// @param request - UpdateDomainRemarkRequest
//
// @return UpdateDomainRemarkResponse
func (client *Client) UpdateDomainRemark(request *UpdateDomainRemarkRequest) (_result *UpdateDomainRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainRemarkResponse{}
	_body, _err := client.UpdateDomainRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an access policy.
//
// @param request - UpdateGtmAccessStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmAccessStrategyResponse
func (client *Client) UpdateGtmAccessStrategyWithOptions(request *UpdateGtmAccessStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmAccessStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessLines) {
		query["AccessLines"] = request.AccessLines
	}

	if !dara.IsNil(request.DefaultAddrPoolId) {
		query["DefaultAddrPoolId"] = request.DefaultAddrPoolId
	}

	if !dara.IsNil(request.FailoverAddrPoolId) {
		query["FailoverAddrPoolId"] = request.FailoverAddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.StrategyId) {
		query["StrategyId"] = request.StrategyId
	}

	if !dara.IsNil(request.StrategyName) {
		query["StrategyName"] = request.StrategyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGtmAccessStrategy"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmAccessStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an access policy.
//
// @param request - UpdateGtmAccessStrategyRequest
//
// @return UpdateGtmAccessStrategyResponse
func (client *Client) UpdateGtmAccessStrategy(request *UpdateGtmAccessStrategyRequest) (_result *UpdateGtmAccessStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmAccessStrategyResponse{}
	_body, _err := client.UpdateGtmAccessStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an address pool.
//
// @param request - UpdateGtmAddressPoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmAddressPoolResponse
func (client *Client) UpdateGtmAddressPoolWithOptions(request *UpdateGtmAddressPoolRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmAddressPoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Addr) {
		query["Addr"] = request.Addr
	}

	if !dara.IsNil(request.AddrPoolId) {
		query["AddrPoolId"] = request.AddrPoolId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MinAvailableAddrNum) {
		query["MinAvailableAddrNum"] = request.MinAvailableAddrNum
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
		Action:      dara.String("UpdateGtmAddressPool"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an address pool.
//
// @param request - UpdateGtmAddressPoolRequest
//
// @return UpdateGtmAddressPoolResponse
func (client *Client) UpdateGtmAddressPool(request *UpdateGtmAddressPoolRequest) (_result *UpdateGtmAddressPoolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmAddressPoolResponse{}
	_body, _err := client.UpdateGtmAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the global configuration of a Global Traffic Manager (GTM) instance.
//
// @param request - UpdateGtmInstanceGlobalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmInstanceGlobalConfigResponse
func (client *Client) UpdateGtmInstanceGlobalConfigWithOptions(request *UpdateGtmInstanceGlobalConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmInstanceGlobalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertGroup) {
		query["AlertGroup"] = request.AlertGroup
	}

	if !dara.IsNil(request.CnameCustomDomainName) {
		query["CnameCustomDomainName"] = request.CnameCustomDomainName
	}

	if !dara.IsNil(request.CnameMode) {
		query["CnameMode"] = request.CnameMode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LbaStrategy) {
		query["LbaStrategy"] = request.LbaStrategy
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.UserDomainName) {
		query["UserDomainName"] = request.UserDomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGtmInstanceGlobalConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmInstanceGlobalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the global configuration of a Global Traffic Manager (GTM) instance.
//
// @param request - UpdateGtmInstanceGlobalConfigRequest
//
// @return UpdateGtmInstanceGlobalConfigResponse
func (client *Client) UpdateGtmInstanceGlobalConfig(request *UpdateGtmInstanceGlobalConfigRequest) (_result *UpdateGtmInstanceGlobalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmInstanceGlobalConfigResponse{}
	_body, _err := client.UpdateGtmInstanceGlobalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a health check configuration.
//
// @param request - UpdateGtmMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmMonitorResponse
func (client *Client) UpdateGtmMonitorWithOptions(request *UpdateGtmMonitorRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCityNode) {
		query["IspCityNode"] = request.IspCityNode
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MonitorConfigId) {
		query["MonitorConfigId"] = request.MonitorConfigId
	}

	if !dara.IsNil(request.MonitorExtendInfo) {
		query["MonitorExtendInfo"] = request.MonitorExtendInfo
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGtmMonitor"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a health check configuration.
//
// @param request - UpdateGtmMonitorRequest
//
// @return UpdateGtmMonitorResponse
func (client *Client) UpdateGtmMonitor(request *UpdateGtmMonitorRequest) (_result *UpdateGtmMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmMonitorResponse{}
	_body, _err := client.UpdateGtmMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a disaster recovery plan.
//
// @param request - UpdateGtmRecoveryPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGtmRecoveryPlanResponse
func (client *Client) UpdateGtmRecoveryPlanWithOptions(request *UpdateGtmRecoveryPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateGtmRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FaultAddrPool) {
		query["FaultAddrPool"] = request.FaultAddrPool
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RecoveryPlanId) {
		query["RecoveryPlanId"] = request.RecoveryPlanId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGtmRecoveryPlan"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGtmRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a disaster recovery plan.
//
// @param request - UpdateGtmRecoveryPlanRequest
//
// @return UpdateGtmRecoveryPlanResponse
func (client *Client) UpdateGtmRecoveryPlan(request *UpdateGtmRecoveryPlanRequest) (_result *UpdateGtmRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGtmRecoveryPlanResponse{}
	_body, _err := client.UpdateGtmRecoveryPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a cache refresh plan.
//
// @param request - UpdateIspFlushCacheInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIspFlushCacheInstanceConfigResponse
func (client *Client) UpdateIspFlushCacheInstanceConfigWithOptions(request *UpdateIspFlushCacheInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateIspFlushCacheInstanceConfigResponse, _err error) {
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

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIspFlushCacheInstanceConfig"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIspFlushCacheInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a cache refresh plan.
//
// @param request - UpdateIspFlushCacheInstanceConfigRequest
//
// @return UpdateIspFlushCacheInstanceConfigResponse
func (client *Client) UpdateIspFlushCacheInstanceConfig(request *UpdateIspFlushCacheInstanceConfigRequest) (_result *UpdateIspFlushCacheInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIspFlushCacheInstanceConfigResponse{}
	_body, _err := client.UpdateIspFlushCacheInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a built-in authoritative record for recursive resolution.
//
// Description:
//
// If a DNS record is locked, it cannot be deleted.
//
// @param request - UpdateRecursionRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordResponse
func (client *Client) UpdateRecursionRecordWithOptions(request *UpdateRecursionRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordResponse, _err error) {
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

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a built-in authoritative record for recursive resolution.
//
// Description:
//
// If a DNS record is locked, it cannot be deleted.
//
// @param request - UpdateRecursionRecordRequest
//
// @return UpdateRecursionRecordResponse
func (client *Client) UpdateRecursionRecord(request *UpdateRecursionRecordRequest) (_result *UpdateRecursionRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordResponse{}
	_body, _err := client.UpdateRecursionRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enable status of a recursion record.
//
// @param request - UpdateRecursionRecordEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordEnableStatusResponse
func (client *Client) UpdateRecursionRecordEnableStatusWithOptions(request *UpdateRecursionRecordEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordEnableStatusResponse, _err error) {
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

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecordEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enable status of a recursion record.
//
// @param request - UpdateRecursionRecordEnableStatusRequest
//
// @return UpdateRecursionRecordEnableStatusResponse
func (client *Client) UpdateRecursionRecordEnableStatus(request *UpdateRecursionRecordEnableStatusRequest) (_result *UpdateRecursionRecordEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordEnableStatusResponse{}
	_body, _err := client.UpdateRecursionRecordEnableStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the remarks for a built-in authoritative record in HTTPDNS.
//
// @param request - UpdateRecursionRecordRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordRemarkResponse
func (client *Client) UpdateRecursionRecordRemarkWithOptions(request *UpdateRecursionRecordRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordRemarkResponse, _err error) {
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

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecordRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the remarks for a built-in authoritative record in HTTPDNS.
//
// @param request - UpdateRecursionRecordRemarkRequest
//
// @return UpdateRecursionRecordRemarkResponse
func (client *Client) UpdateRecursionRecordRemark(request *UpdateRecursionRecordRemarkRequest) (_result *UpdateRecursionRecordRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordRemarkResponse{}
	_body, _err := client.UpdateRecursionRecordRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the weight of a DNS record for recursive resolution.
//
// @param request - UpdateRecursionRecordWeightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordWeightResponse
func (client *Client) UpdateRecursionRecordWeightWithOptions(request *UpdateRecursionRecordWeightRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordWeightResponse, _err error) {
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

	if !dara.IsNil(request.RecordId) {
		query["RecordId"] = request.RecordId
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecordWeight"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordWeightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the weight of a DNS record for recursive resolution.
//
// @param request - UpdateRecursionRecordWeightRequest
//
// @return UpdateRecursionRecordWeightResponse
func (client *Client) UpdateRecursionRecordWeight(request *UpdateRecursionRecordWeightRequest) (_result *UpdateRecursionRecordWeightResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordWeightResponse{}
	_body, _err := client.UpdateRecursionRecordWeightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the enabled status of the weight algorithm for a DNS record used for recursive resolution.
//
// @param request - UpdateRecursionRecordWeightEnableStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionRecordWeightEnableStatusResponse
func (client *Client) UpdateRecursionRecordWeightEnableStatusWithOptions(request *UpdateRecursionRecordWeightEnableStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionRecordWeightEnableStatusResponse, _err error) {
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

	if !dara.IsNil(request.EnableStatus) {
		query["EnableStatus"] = request.EnableStatus
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	if !dara.IsNil(request.Rr) {
		query["Rr"] = request.Rr
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionRecordWeightEnableStatus"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionRecordWeightEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the enabled status of the weight algorithm for a DNS record used for recursive resolution.
//
// @param request - UpdateRecursionRecordWeightEnableStatusRequest
//
// @return UpdateRecursionRecordWeightEnableStatusResponse
func (client *Client) UpdateRecursionRecordWeightEnableStatus(request *UpdateRecursionRecordWeightEnableStatusRequest) (_result *UpdateRecursionRecordWeightEnableStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionRecordWeightEnableStatusResponse{}
	_body, _err := client.UpdateRecursionRecordWeightEnableStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the effective scope of a built-in authoritative domain name zone in HTTPDNS.
//
// @param tmpReq - UpdateRecursionZoneEffectiveScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionZoneEffectiveScopeResponse
func (client *Client) UpdateRecursionZoneEffectiveScopeWithOptions(tmpReq *UpdateRecursionZoneEffectiveScopeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionZoneEffectiveScopeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRecursionZoneEffectiveScopeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EffectiveScopes) {
		request.EffectiveScopesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EffectiveScopes, dara.String("EffectiveScopes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EffectiveScopesShrink) {
		query["EffectiveScopes"] = request.EffectiveScopesShrink
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionZoneEffectiveScope"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionZoneEffectiveScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the effective scope of a built-in authoritative domain name zone in HTTPDNS.
//
// @param request - UpdateRecursionZoneEffectiveScopeRequest
//
// @return UpdateRecursionZoneEffectiveScopeResponse
func (client *Client) UpdateRecursionZoneEffectiveScope(request *UpdateRecursionZoneEffectiveScopeRequest) (_result *UpdateRecursionZoneEffectiveScopeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionZoneEffectiveScopeResponse{}
	_body, _err := client.UpdateRecursionZoneEffectiveScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the recursive proxy mode for a zone.
//
// Description:
//
// The end IP address of each IP segment must be greater than or equal to the start IP address.
//
// The IP address ranges of all IP segments for the domain name cannot overlap across custom lines.
//
// @param request - UpdateRecursionZoneProxyPatternRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionZoneProxyPatternResponse
func (client *Client) UpdateRecursionZoneProxyPatternWithOptions(request *UpdateRecursionZoneProxyPatternRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionZoneProxyPatternResponse, _err error) {
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

	if !dara.IsNil(request.ProxyPattern) {
		query["ProxyPattern"] = request.ProxyPattern
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionZoneProxyPattern"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionZoneProxyPatternResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the recursive proxy mode for a zone.
//
// Description:
//
// The end IP address of each IP segment must be greater than or equal to the start IP address.
//
// The IP address ranges of all IP segments for the domain name cannot overlap across custom lines.
//
// @param request - UpdateRecursionZoneProxyPatternRequest
//
// @return UpdateRecursionZoneProxyPatternResponse
func (client *Client) UpdateRecursionZoneProxyPattern(request *UpdateRecursionZoneProxyPatternRequest) (_result *UpdateRecursionZoneProxyPatternResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionZoneProxyPatternResponse{}
	_body, _err := client.UpdateRecursionZoneProxyPatternWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the remarks for a built-in authoritative domain name zone used for recursive resolution.
//
// Description:
//
// The end IP address of each IP segment must be greater than or equal to the start IP address.
//
// For a domain name, the IP address ranges of all IP segments in all custom lines cannot overlap.
//
// @param request - UpdateRecursionZoneRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecursionZoneRemarkResponse
func (client *Client) UpdateRecursionZoneRemarkWithOptions(request *UpdateRecursionZoneRemarkRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecursionZoneRemarkResponse, _err error) {
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

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecursionZoneRemark"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecursionZoneRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the remarks for a built-in authoritative domain name zone used for recursive resolution.
//
// Description:
//
// The end IP address of each IP segment must be greater than or equal to the start IP address.
//
// For a domain name, the IP address ranges of all IP segments in all custom lines cannot overlap.
//
// @param request - UpdateRecursionZoneRemarkRequest
//
// @return UpdateRecursionZoneRemarkResponse
func (client *Client) UpdateRecursionZoneRemark(request *UpdateRecursionZoneRemarkRequest) (_result *UpdateRecursionZoneRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecursionZoneRemarkResponse{}
	_body, _err := client.UpdateRecursionZoneRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the server-side status of a domain name.
//
// Description:
//
// ## Request description
//
// - This operation updates the server-side status of a domain name.
//
// @param request - UpdateRspDomainServerProhibitStatusForGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRspDomainServerProhibitStatusForGatewayResponse
func (client *Client) UpdateRspDomainServerProhibitStatusForGatewayWithOptions(request *UpdateRspDomainServerProhibitStatusForGatewayRequest, runtime *dara.RuntimeOptions) (_result *UpdateRspDomainServerProhibitStatusForGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddStatusList) {
		query["AddStatusList"] = request.AddStatusList
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeleteStatusList) {
		query["DeleteStatusList"] = request.DeleteStatusList
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRspDomainServerProhibitStatusForGateway"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRspDomainServerProhibitStatusForGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the server-side status of a domain name.
//
// Description:
//
// ## Request description
//
// - This operation updates the server-side status of a domain name.
//
// @param request - UpdateRspDomainServerProhibitStatusForGatewayRequest
//
// @return UpdateRspDomainServerProhibitStatusForGatewayResponse
func (client *Client) UpdateRspDomainServerProhibitStatusForGateway(request *UpdateRspDomainServerProhibitStatusForGatewayRequest) (_result *UpdateRspDomainServerProhibitStatusForGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRspDomainServerProhibitStatusForGatewayResponse{}
	_body, _err := client.UpdateRspDomainServerProhibitStatusForGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the server-side status of a domain name.
//
// Description:
//
// This operation updates the server-side status of a domain name.
//
// @param request - UpdateRspDomainServerProhibitStatusForGatewayOteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRspDomainServerProhibitStatusForGatewayOteResponse
func (client *Client) UpdateRspDomainServerProhibitStatusForGatewayOteWithOptions(request *UpdateRspDomainServerProhibitStatusForGatewayOteRequest, runtime *dara.RuntimeOptions) (_result *UpdateRspDomainServerProhibitStatusForGatewayOteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddStatusList) {
		query["AddStatusList"] = request.AddStatusList
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeleteStatusList) {
		query["DeleteStatusList"] = request.DeleteStatusList
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRspDomainServerProhibitStatusForGatewayOte"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRspDomainServerProhibitStatusForGatewayOteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the server-side status of a domain name.
//
// Description:
//
// This operation updates the server-side status of a domain name.
//
// @param request - UpdateRspDomainServerProhibitStatusForGatewayOteRequest
//
// @return UpdateRspDomainServerProhibitStatusForGatewayOteResponse
func (client *Client) UpdateRspDomainServerProhibitStatusForGatewayOte(request *UpdateRspDomainServerProhibitStatusForGatewayOteRequest) (_result *UpdateRspDomainServerProhibitStatusForGatewayOteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRspDomainServerProhibitStatusForGatewayOteResponse{}
	_body, _err := client.UpdateRspDomainServerProhibitStatusForGatewayOteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether an instance hostname can be added.
//
// @param request - ValidateDnsGtmCnameRrCanUseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateDnsGtmCnameRrCanUseResponse
func (client *Client) ValidateDnsGtmCnameRrCanUseWithOptions(request *ValidateDnsGtmCnameRrCanUseRequest, runtime *dara.RuntimeOptions) (_result *ValidateDnsGtmCnameRrCanUseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CnameMode) {
		query["CnameMode"] = request.CnameMode
	}

	if !dara.IsNil(request.CnameRr) {
		query["CnameRr"] = request.CnameRr
	}

	if !dara.IsNil(request.CnameType) {
		query["CnameType"] = request.CnameType
	}

	if !dara.IsNil(request.CnameZone) {
		query["CnameZone"] = request.CnameZone
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateDnsGtmCnameRrCanUse"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateDnsGtmCnameRrCanUseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether an instance hostname can be added.
//
// @param request - ValidateDnsGtmCnameRrCanUseRequest
//
// @return ValidateDnsGtmCnameRrCanUseResponse
func (client *Client) ValidateDnsGtmCnameRrCanUse(request *ValidateDnsGtmCnameRrCanUseRequest) (_result *ValidateDnsGtmCnameRrCanUseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateDnsGtmCnameRrCanUseResponse{}
	_body, _err := client.ValidateDnsGtmCnameRrCanUseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validates a UDP IP address segment for Public DNS.
//
// @param request - ValidatePdnsUdpIpSegmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidatePdnsUdpIpSegmentResponse
func (client *Client) ValidatePdnsUdpIpSegmentWithOptions(request *ValidatePdnsUdpIpSegmentRequest, runtime *dara.RuntimeOptions) (_result *ValidatePdnsUdpIpSegmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.IpToken) {
		query["IpToken"] = request.IpToken
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidatePdnsUdpIpSegment"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidatePdnsUdpIpSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates a UDP IP address segment for Public DNS.
//
// @param request - ValidatePdnsUdpIpSegmentRequest
//
// @return ValidatePdnsUdpIpSegmentResponse
func (client *Client) ValidatePdnsUdpIpSegment(request *ValidatePdnsUdpIpSegmentRequest) (_result *ValidatePdnsUdpIpSegmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidatePdnsUdpIpSegmentResponse{}
	_body, _err := client.ValidatePdnsUdpIpSegmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies DNS records.
//
// @param request - VerifyAtiAgentDnsRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyAtiAgentDnsRecordsResponse
func (client *Client) VerifyAtiAgentDnsRecordsWithOptions(request *VerifyAtiAgentDnsRecordsRequest, runtime *dara.RuntimeOptions) (_result *VerifyAtiAgentDnsRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyAtiAgentDnsRecords"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyAtiAgentDnsRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies DNS records.
//
// @param request - VerifyAtiAgentDnsRecordsRequest
//
// @return VerifyAtiAgentDnsRecordsResponse
func (client *Client) VerifyAtiAgentDnsRecords(request *VerifyAtiAgentDnsRecordsRequest) (_result *VerifyAtiAgentDnsRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyAtiAgentDnsRecordsResponse{}
	_body, _err := client.VerifyAtiAgentDnsRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 触发 ACME 预检
//
// @param request - VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse
func (client *Client) VerifyAtiAgentRegisterInfoAcmeChallengeRecordWithOptions(request *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest, runtime *dara.RuntimeOptions) (_result *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRegisterInfoId) {
		query["AgentRegisterInfoId"] = request.AgentRegisterInfoId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyAtiAgentRegisterInfoAcmeChallengeRecord"),
		Version:     dara.String("2015-01-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 触发 ACME 预检
//
// @param request - VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest
//
// @return VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse
func (client *Client) VerifyAtiAgentRegisterInfoAcmeChallengeRecord(request *VerifyAtiAgentRegisterInfoAcmeChallengeRecordRequest) (_result *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponse{}
	_body, _err := client.VerifyAtiAgentRegisterInfoAcmeChallengeRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
