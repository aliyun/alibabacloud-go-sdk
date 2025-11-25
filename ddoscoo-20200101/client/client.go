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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ddoscoo"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AddAutoCcBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAutoCcBlacklistResponse
func (client *Client) AddAutoCcBlacklistWithOptions(request *AddAutoCcBlacklistRequest, runtime *dara.RuntimeOptions) (_result *AddAutoCcBlacklistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Blacklist) {
		query["Blacklist"] = request.Blacklist
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAutoCcBlacklist"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAutoCcBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddAutoCcBlacklistRequest
//
// @return AddAutoCcBlacklistResponse
func (client *Client) AddAutoCcBlacklist(request *AddAutoCcBlacklistRequest) (_result *AddAutoCcBlacklistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAutoCcBlacklistResponse{}
	_body, _err := client.AddAutoCcBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds IP addresses to the IP address whitelist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// Description:
//
// You can call the AddAutoCcWhitelist operation to add IP addresses to the whitelist of an Anti-DDoS Pro or Anti-DDoS Premium instance. This way, the Anti-DDoS Pro or Anti-DDoS Premium instance allows traffic from the IP addresses.
//
// By default, the traffic from the IP addresses that you add to the whitelist is always allowed. If you no longer use the whitelist, you can call the [EmptyAutoCcWhitelist](https://help.aliyun.com/document_detail/157505.html) operation to remove the IP addresses from the whitelist.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddAutoCcWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAutoCcWhitelistResponse
func (client *Client) AddAutoCcWhitelistWithOptions(request *AddAutoCcWhitelistRequest, runtime *dara.RuntimeOptions) (_result *AddAutoCcWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Whitelist) {
		query["Whitelist"] = request.Whitelist
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAutoCcWhitelist"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAutoCcWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds IP addresses to the IP address whitelist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// Description:
//
// You can call the AddAutoCcWhitelist operation to add IP addresses to the whitelist of an Anti-DDoS Pro or Anti-DDoS Premium instance. This way, the Anti-DDoS Pro or Anti-DDoS Premium instance allows traffic from the IP addresses.
//
// By default, the traffic from the IP addresses that you add to the whitelist is always allowed. If you no longer use the whitelist, you can call the [EmptyAutoCcWhitelist](https://help.aliyun.com/document_detail/157505.html) operation to remove the IP addresses from the whitelist.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - AddAutoCcWhitelistRequest
//
// @return AddAutoCcWhitelistResponse
func (client *Client) AddAutoCcWhitelist(request *AddAutoCcWhitelistRequest) (_result *AddAutoCcWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAutoCcWhitelistResponse{}
	_body, _err := client.AddAutoCcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates an SSL certificate with the forwarding rule of a website.
//
// @param request - AssociateWebCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateWebCertResponse
func (client *Client) AssociateWebCertWithOptions(request *AssociateWebCertRequest, runtime *dara.RuntimeOptions) (_result *AssociateWebCertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Cert) {
		body["Cert"] = request.Cert
	}

	if !dara.IsNil(request.CertId) {
		body["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertIdentifier) {
		body["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.CertName) {
		body["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertRegion) {
		body["CertRegion"] = request.CertRegion
	}

	if !dara.IsNil(request.Domain) {
		body["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Key) {
		body["Key"] = request.Key
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateWebCert"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateWebCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an SSL certificate with the forwarding rule of a website.
//
// @param request - AssociateWebCertRequest
//
// @return AssociateWebCertResponse
func (client *Client) AssociateWebCert(request *AssociateWebCertRequest) (_result *AssociateWebCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateWebCertResponse{}
	_body, _err := client.AssociateWebCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an object to a scenario-specific custom policy for protection.
//
// @param request - AttachSceneDefenseObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachSceneDefenseObjectResponse
func (client *Client) AttachSceneDefenseObjectWithOptions(request *AttachSceneDefenseObjectRequest, runtime *dara.RuntimeOptions) (_result *AttachSceneDefenseObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.Objects) {
		query["Objects"] = request.Objects
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachSceneDefenseObject"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachSceneDefenseObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an object to a scenario-specific custom policy for protection.
//
// @param request - AttachSceneDefenseObjectRequest
//
// @return AttachSceneDefenseObjectResponse
func (client *Client) AttachSceneDefenseObject(request *AttachSceneDefenseObjectRequest) (_result *AttachSceneDefenseObjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachSceneDefenseObjectResponse{}
	_body, _err := client.AttachSceneDefenseObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the global mitigation policy feature, including the feature status and settings.
//
// @param request - ConfigDomainSecurityProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigDomainSecurityProfileResponse
func (client *Client) ConfigDomainSecurityProfileWithOptions(request *ConfigDomainSecurityProfileRequest, runtime *dara.RuntimeOptions) (_result *ConfigDomainSecurityProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cluster) {
		query["Cluster"] = request.Cluster
	}

	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigDomainSecurityProfile"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigDomainSecurityProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the global mitigation policy feature, including the feature status and settings.
//
// @param request - ConfigDomainSecurityProfileRequest
//
// @return ConfigDomainSecurityProfileResponse
func (client *Client) ConfigDomainSecurityProfile(request *ConfigDomainSecurityProfileRequest) (_result *ConfigDomainSecurityProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigDomainSecurityProfileResponse{}
	_body, _err := client.ConfigDomainSecurityProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置全局模板规则
//
// @param request - ConfigL7GlobalRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigL7GlobalRuleResponse
func (client *Client) ConfigL7GlobalRuleWithOptions(request *ConfigL7GlobalRuleRequest, runtime *dara.RuntimeOptions) (_result *ConfigL7GlobalRuleResponse, _err error) {
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

	if !dara.IsNil(request.RuleAttr) {
		query["RuleAttr"] = request.RuleAttr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigL7GlobalRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigL7GlobalRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置全局模板规则
//
// @param request - ConfigL7GlobalRuleRequest
//
// @return ConfigL7GlobalRuleResponse
func (client *Client) ConfigL7GlobalRule(request *ConfigL7GlobalRuleRequest) (_result *ConfigL7GlobalRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigL7GlobalRuleResponse{}
	_body, _err := client.ConfigL7GlobalRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a back-to-origin policy for the forwarding rule of a website.
//
// Description:
//
// If multiple origin servers are configured for a website that is added to Anti-DDoS Pro or Anti-DDoS Premium, you can modify the load balancing algorithms for back-to-origin traffic based on back-to-origin policies. The IP hash algorithm is used by default. You can change the algorithm to the round-robin or least response time algorithm. For more information, see the description of the **Policy*	- parameter in the "Request parameters" section of this topic.
//
// @param request - ConfigL7RsPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigL7RsPolicyResponse
func (client *Client) ConfigL7RsPolicyWithOptions(request *ConfigL7RsPolicyRequest, runtime *dara.RuntimeOptions) (_result *ConfigL7RsPolicyResponse, _err error) {
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

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.UpstreamRetry) {
		query["UpstreamRetry"] = request.UpstreamRetry
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigL7RsPolicy"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigL7RsPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a back-to-origin policy for the forwarding rule of a website.
//
// Description:
//
// If multiple origin servers are configured for a website that is added to Anti-DDoS Pro or Anti-DDoS Premium, you can modify the load balancing algorithms for back-to-origin traffic based on back-to-origin policies. The IP hash algorithm is used by default. You can change the algorithm to the round-robin or least response time algorithm. For more information, see the description of the **Policy*	- parameter in the "Request parameters" section of this topic.
//
// @param request - ConfigL7RsPolicyRequest
//
// @return ConfigL7RsPolicyResponse
func (client *Client) ConfigL7RsPolicy(request *ConfigL7RsPolicyRequest) (_result *ConfigL7RsPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigL7RsPolicyResponse{}
	_body, _err := client.ConfigL7RsPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the settings for back-to-origin persistent connections for a domain name.
//
// @param request - ConfigL7UsKeepaliveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigL7UsKeepaliveResponse
func (client *Client) ConfigL7UsKeepaliveWithOptions(request *ConfigL7UsKeepaliveRequest, runtime *dara.RuntimeOptions) (_result *ConfigL7UsKeepaliveResponse, _err error) {
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

	if !dara.IsNil(request.DownstreamKeepalive) {
		query["DownstreamKeepalive"] = request.DownstreamKeepalive
	}

	if !dara.IsNil(request.UpstreamKeepalive) {
		query["UpstreamKeepalive"] = request.UpstreamKeepalive
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigL7UsKeepalive"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigL7UsKeepaliveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the settings for back-to-origin persistent connections for a domain name.
//
// @param request - ConfigL7UsKeepaliveRequest
//
// @return ConfigL7UsKeepaliveResponse
func (client *Client) ConfigL7UsKeepalive(request *ConfigL7UsKeepaliveRequest) (_result *ConfigL7UsKeepaliveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigL7UsKeepaliveResponse{}
	_body, _err := client.ConfigL7UsKeepaliveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies a threshold for the clean bandwidth of an Anti-DDoS Pro or Anti-DDoS premium instance. If the threshold is reached, rate limiting is triggered.
//
// @param request - ConfigLayer4RealLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer4RealLimitResponse
func (client *Client) ConfigLayer4RealLimitWithOptions(request *ConfigLayer4RealLimitRequest, runtime *dara.RuntimeOptions) (_result *ConfigLayer4RealLimitResponse, _err error) {
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

	if !dara.IsNil(request.LimitValue) {
		query["LimitValue"] = request.LimitValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigLayer4RealLimit"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigLayer4RealLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies a threshold for the clean bandwidth of an Anti-DDoS Pro or Anti-DDoS premium instance. If the threshold is reached, rate limiting is triggered.
//
// @param request - ConfigLayer4RealLimitRequest
//
// @return ConfigLayer4RealLimitResponse
func (client *Client) ConfigLayer4RealLimit(request *ConfigLayer4RealLimitRequest) (_result *ConfigLayer4RealLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigLayer4RealLimitResponse{}
	_body, _err := client.ConfigLayer4RealLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a description to a port forwarding rule.
//
// @param request - ConfigLayer4RemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer4RemarkResponse
func (client *Client) ConfigLayer4RemarkWithOptions(request *ConfigLayer4RemarkRequest, runtime *dara.RuntimeOptions) (_result *ConfigLayer4RemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Listeners) {
		query["Listeners"] = request.Listeners
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigLayer4Remark"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigLayer4RemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a description to a port forwarding rule.
//
// @param request - ConfigLayer4RemarkRequest
//
// @return ConfigLayer4RemarkResponse
func (client *Client) ConfigLayer4Remark(request *ConfigLayer4RemarkRequest) (_result *ConfigLayer4RemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigLayer4RemarkResponse{}
	_body, _err := client.ConfigLayer4RemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the origin redundancy mode for a port forwarding rule.
//
// @param request - ConfigLayer4RuleBakModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer4RuleBakModeResponse
func (client *Client) ConfigLayer4RuleBakModeWithOptions(request *ConfigLayer4RuleBakModeRequest, runtime *dara.RuntimeOptions) (_result *ConfigLayer4RuleBakModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BakMode) {
		query["BakMode"] = request.BakMode
	}

	if !dara.IsNil(request.Listeners) {
		query["Listeners"] = request.Listeners
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigLayer4RuleBakMode"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigLayer4RuleBakModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the origin redundancy mode for a port forwarding rule.
//
// @param request - ConfigLayer4RuleBakModeRequest
//
// @return ConfigLayer4RuleBakModeResponse
func (client *Client) ConfigLayer4RuleBakMode(request *ConfigLayer4RuleBakModeRequest) (_result *ConfigLayer4RuleBakModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigLayer4RuleBakModeResponse{}
	_body, _err := client.ConfigLayer4RuleBakModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the IP addresses of the primary and secondary origin servers for a port forwarding rule.
//
// @param request - ConfigLayer4RulePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigLayer4RulePolicyResponse
func (client *Client) ConfigLayer4RulePolicyWithOptions(request *ConfigLayer4RulePolicyRequest, runtime *dara.RuntimeOptions) (_result *ConfigLayer4RulePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Listeners) {
		query["Listeners"] = request.Listeners
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigLayer4RulePolicy"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigLayer4RulePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the IP addresses of the primary and secondary origin servers for a port forwarding rule.
//
// @param request - ConfigLayer4RulePolicyRequest
//
// @return ConfigLayer4RulePolicyResponse
func (client *Client) ConfigLayer4RulePolicy(request *ConfigLayer4RulePolicyRequest) (_result *ConfigLayer4RulePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigLayer4RulePolicyResponse{}
	_body, _err := client.ConfigLayer4RulePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures blocked locations for an Anti-DDoS Proxy instance.
//
// @param request - ConfigNetworkRegionBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigNetworkRegionBlockResponse
func (client *Client) ConfigNetworkRegionBlockWithOptions(request *ConfigNetworkRegionBlockRequest, runtime *dara.RuntimeOptions) (_result *ConfigNetworkRegionBlockResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigNetworkRegionBlock"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigNetworkRegionBlockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures blocked locations for an Anti-DDoS Proxy instance.
//
// @param request - ConfigNetworkRegionBlockRequest
//
// @return ConfigNetworkRegionBlockResponse
func (client *Client) ConfigNetworkRegionBlock(request *ConfigNetworkRegionBlockRequest) (_result *ConfigNetworkRegionBlockResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigNetworkRegionBlockResponse{}
	_body, _err := client.ConfigNetworkRegionBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP addresses of the origin server that is configured in a port forwarding rule.
//
// @param request - ConfigNetworkRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigNetworkRulesResponse
func (client *Client) ConfigNetworkRulesWithOptions(request *ConfigNetworkRulesRequest, runtime *dara.RuntimeOptions) (_result *ConfigNetworkRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkRules) {
		query["NetworkRules"] = request.NetworkRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigNetworkRules"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigNetworkRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP addresses of the origin server that is configured in a port forwarding rule.
//
// @param request - ConfigNetworkRulesRequest
//
// @return ConfigNetworkRulesResponse
func (client *Client) ConfigNetworkRules(request *ConfigNetworkRulesRequest) (_result *ConfigNetworkRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigNetworkRulesResponse{}
	_body, _err := client.ConfigNetworkRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds the filtering policies for UDP reflection attacks on an Anti-DDoS Pro or Anti-DDoS Premium instance to filter out the source ports of UDP traffic.
//
// Description:
//
// You can call this operation to configure filtering policies to filter out UDP traffic from specific ports. This helps defend against UDP reflection attacks.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ConfigUdpReflectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigUdpReflectResponse
func (client *Client) ConfigUdpReflectWithOptions(request *ConfigUdpReflectRequest, runtime *dara.RuntimeOptions) (_result *ConfigUdpReflectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigUdpReflect"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigUdpReflectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds the filtering policies for UDP reflection attacks on an Anti-DDoS Pro or Anti-DDoS Premium instance to filter out the source ports of UDP traffic.
//
// Description:
//
// You can call this operation to configure filtering policies to filter out UDP traffic from specific ports. This helps defend against UDP reflection attacks.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ConfigUdpReflectRequest
//
// @return ConfigUdpReflectResponse
func (client *Client) ConfigUdpReflect(request *ConfigUdpReflectRequest) (_result *ConfigUdpReflectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigUdpReflectResponse{}
	_body, _err := client.ConfigUdpReflectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置新版基于匹配条件的cc规则
//
// @param request - ConfigWebCCRuleV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigWebCCRuleV2Response
func (client *Client) ConfigWebCCRuleV2WithOptions(request *ConfigWebCCRuleV2Request, runtime *dara.RuntimeOptions) (_result *ConfigWebCCRuleV2Response, _err error) {
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

	if !dara.IsNil(request.Expires) {
		query["Expires"] = request.Expires
	}

	if !dara.IsNil(request.RuleList) {
		query["RuleList"] = request.RuleList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigWebCCRuleV2"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigWebCCRuleV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置新版基于匹配条件的cc规则
//
// @param request - ConfigWebCCRuleV2Request
//
// @return ConfigWebCCRuleV2Response
func (client *Client) ConfigWebCCRuleV2(request *ConfigWebCCRuleV2Request) (_result *ConfigWebCCRuleV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigWebCCRuleV2Response{}
	_body, _err := client.ConfigWebCCRuleV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the mode of the Frequency Control policy for a website.
//
// @param request - ConfigWebCCTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigWebCCTemplateResponse
func (client *Client) ConfigWebCCTemplateWithOptions(request *ConfigWebCCTemplateRequest, runtime *dara.RuntimeOptions) (_result *ConfigWebCCTemplateResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigWebCCTemplate"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigWebCCTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the mode of the Frequency Control policy for a website.
//
// @param request - ConfigWebCCTemplateRequest
//
// @return ConfigWebCCTemplateResponse
func (client *Client) ConfigWebCCTemplate(request *ConfigWebCCTemplateRequest) (_result *ConfigWebCCTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigWebCCTemplateResponse{}
	_body, _err := client.ConfigWebCCTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the IP address whitelist and blacklist for a website.
//
// @param request - ConfigWebIpSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigWebIpSetResponse
func (client *Client) ConfigWebIpSetWithOptions(request *ConfigWebIpSetRequest, runtime *dara.RuntimeOptions) (_result *ConfigWebIpSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BlackList) {
		query["BlackList"] = request.BlackList
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.WhiteList) {
		query["WhiteList"] = request.WhiteList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigWebIpSet"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigWebIpSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the IP address whitelist and blacklist for a website.
//
// @param request - ConfigWebIpSetRequest
//
// @return ConfigWebIpSetResponse
func (client *Client) ConfigWebIpSet(request *ConfigWebIpSetRequest) (_result *ConfigWebIpSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigWebIpSetResponse{}
	_body, _err := client.ConfigWebIpSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an asynchronous export task to export forwarding rules for websites, port forwarding rules, session persistence and health check settings, DDoS mitigation policies, the IP address blacklist, or the IP address whitelist.
//
// @param request - CreateAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsyncTaskResponse
func (client *Client) CreateAsyncTaskWithOptions(request *CreateAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TaskParams) {
		query["TaskParams"] = request.TaskParams
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAsyncTask"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an asynchronous export task to export forwarding rules for websites, port forwarding rules, session persistence and health check settings, DDoS mitigation policies, the IP address blacklist, or the IP address whitelist.
//
// @param request - CreateAsyncTaskRequest
//
// @return CreateAsyncTaskResponse
func (client *Client) CreateAsyncTask(request *CreateAsyncTaskRequest) (_result *CreateAsyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.CreateAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a forwarding rule for a website.
//
// @param request - CreateDomainResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResourceResponse
func (client *Client) CreateDomainResourceWithOptions(request *CreateDomainResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainResourceResponse, _err error) {
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

	if !dara.IsNil(request.HttpsExt) {
		query["HttpsExt"] = request.HttpsExt
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ProxyTypes) {
		query["ProxyTypes"] = request.ProxyTypes
	}

	if !dara.IsNil(request.RealServers) {
		query["RealServers"] = request.RealServers
	}

	if !dara.IsNil(request.RsType) {
		query["RsType"] = request.RsType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomainResource"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a forwarding rule for a website.
//
// @param request - CreateDomainResourceRequest
//
// @return CreateDomainResourceResponse
func (client *Client) CreateDomainResource(request *CreateDomainResourceRequest) (_result *CreateDomainResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDomainResourceResponse{}
	_body, _err := client.CreateDomainResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a port forwarding rule.
//
// @param request - CreateNetworkRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkRulesResponse
func (client *Client) CreateNetworkRulesWithOptions(request *CreateNetworkRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkRules) {
		query["NetworkRules"] = request.NetworkRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkRules"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a port forwarding rule.
//
// @param request - CreateNetworkRulesRequest
//
// @return CreateNetworkRulesResponse
func (client *Client) CreateNetworkRules(request *CreateNetworkRulesRequest) (_result *CreateNetworkRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkRulesResponse{}
	_body, _err := client.CreateNetworkRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a port forwarding rule.
//
// Description:
//
// You can call this operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// @param request - CreatePortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePortResponse
func (client *Client) CreatePortWithOptions(request *CreatePortRequest, runtime *dara.RuntimeOptions) (_result *CreatePortResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendPort) {
		query["BackendPort"] = request.BackendPort
	}

	if !dara.IsNil(request.FrontendPort) {
		query["FrontendPort"] = request.FrontendPort
	}

	if !dara.IsNil(request.FrontendProtocol) {
		query["FrontendProtocol"] = request.FrontendProtocol
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProxyEnable) {
		query["ProxyEnable"] = request.ProxyEnable
	}

	if !dara.IsNil(request.RealServers) {
		query["RealServers"] = request.RealServers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePort"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a port forwarding rule.
//
// Description:
//
// You can call this operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// @param request - CreatePortRequest
//
// @return CreatePortResponse
func (client *Client) CreatePort(request *CreatePortRequest) (_result *CreatePortResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePortResponse{}
	_body, _err := client.CreatePortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scenario-specific custom policy.
//
// @param request - CreateSceneDefensePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSceneDefensePolicyResponse
func (client *Client) CreateSceneDefensePolicyWithOptions(request *CreateSceneDefensePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateSceneDefensePolicyResponse, _err error) {
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
		Action:      dara.String("CreateSceneDefensePolicy"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSceneDefensePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scenario-specific custom policy.
//
// @param request - CreateSceneDefensePolicyRequest
//
// @return CreateSceneDefensePolicyResponse
func (client *Client) CreateSceneDefensePolicy(request *CreateSceneDefensePolicyRequest) (_result *CreateSceneDefensePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSceneDefensePolicyResponse{}
	_body, _err := client.CreateSceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scheduling rule for Sec-Traffic Manager.
//
// @param request - CreateSchedulerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchedulerRuleResponse
func (client *Client) CreateSchedulerRuleWithOptions(request *CreateSchedulerRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateSchedulerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSchedulerRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSchedulerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduling rule for Sec-Traffic Manager.
//
// @param request - CreateSchedulerRuleRequest
//
// @return CreateSchedulerRuleResponse
func (client *Client) CreateSchedulerRule(request *CreateSchedulerRuleRequest) (_result *CreateSchedulerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSchedulerRuleResponse{}
	_body, _err := client.CreateSchedulerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to multiple Anti-DDoS Proxy instances at a time.
//
// Description:
//
// You can call the CreateTagResources operation to add tags to multiple Anti-DDoS Proxy instances at a time.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTagResourcesResponse
func (client *Client) CreateTagResourcesWithOptions(request *CreateTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *CreateTagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTagResources"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to multiple Anti-DDoS Proxy instances at a time.
//
// Description:
//
// You can call the CreateTagResources operation to add tags to multiple Anti-DDoS Proxy instances at a time.
//
// ### [](#qps-)Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateTagResourcesRequest
//
// @return CreateTagResourcesResponse
func (client *Client) CreateTagResources(request *CreateTagResourcesRequest) (_result *CreateTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTagResourcesResponse{}
	_body, _err := client.CreateTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateWebCCRule is deprecated, please use ddoscoo::2020-01-01::ConfigWebCCRuleV2 instead.
//
// Summary:
//
// Creates a custom frequency control rule for a website.
//
// @param request - CreateWebCCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWebCCRuleResponse
func (client *Client) CreateWebCCRuleWithOptions(request *CreateWebCCRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateWebCCRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Act) {
		query["Act"] = request.Act
	}

	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Uri) {
		query["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWebCCRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWebCCRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateWebCCRule is deprecated, please use ddoscoo::2020-01-01::ConfigWebCCRuleV2 instead.
//
// Summary:
//
// Creates a custom frequency control rule for a website.
//
// @param request - CreateWebCCRuleRequest
//
// @return CreateWebCCRuleResponse
// Deprecated
func (client *Client) CreateWebCCRule(request *CreateWebCCRuleRequest) (_result *CreateWebCCRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWebCCRuleResponse{}
	_body, _err := client.CreateWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a forwarding rule for a website.
//
// @param request - CreateWebRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWebRuleResponse
func (client *Client) CreateWebRuleWithOptions(request *CreateWebRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateWebRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseId) {
		query["DefenseId"] = request.DefenseId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.HttpsExt) {
		query["HttpsExt"] = request.HttpsExt
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RsType) {
		query["RsType"] = request.RsType
	}

	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWebRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWebRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a forwarding rule for a website.
//
// @param request - CreateWebRuleRequest
//
// @return CreateWebRuleResponse
func (client *Client) CreateWebRule(request *CreateWebRuleRequest) (_result *CreateWebRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWebRuleResponse{}
	_body, _err := client.CreateWebRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an asynchronous export task.
//
// @param request - DeleteAsyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAsyncTaskResponse
func (client *Client) DeleteAsyncTaskWithOptions(request *DeleteAsyncTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAsyncTask"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an asynchronous export task.
//
// @param request - DeleteAsyncTaskRequest
//
// @return DeleteAsyncTaskResponse
func (client *Client) DeleteAsyncTask(request *DeleteAsyncTaskRequest) (_result *DeleteAsyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAsyncTaskResponse{}
	_body, _err := client.DeleteAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes IP addresses from the IP address blacklist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DeleteAutoCcBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoCcBlacklistResponse
func (client *Client) DeleteAutoCcBlacklistWithOptions(request *DeleteAutoCcBlacklistRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoCcBlacklistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Blacklist) {
		query["Blacklist"] = request.Blacklist
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoCcBlacklist"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoCcBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes IP addresses from the IP address blacklist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DeleteAutoCcBlacklistRequest
//
// @return DeleteAutoCcBlacklistResponse
func (client *Client) DeleteAutoCcBlacklist(request *DeleteAutoCcBlacklistRequest) (_result *DeleteAutoCcBlacklistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAutoCcBlacklistResponse{}
	_body, _err := client.DeleteAutoCcBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes IP addresses from the IP address whitelist of an Anti-DDoS Proxy instance.
//
// @param request - DeleteAutoCcWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoCcWhitelistResponse
func (client *Client) DeleteAutoCcWhitelistWithOptions(request *DeleteAutoCcWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoCcWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.Whitelist) {
		query["Whitelist"] = request.Whitelist
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoCcWhitelist"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoCcWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes IP addresses from the IP address whitelist of an Anti-DDoS Proxy instance.
//
// @param request - DeleteAutoCcWhitelistRequest
//
// @return DeleteAutoCcWhitelistResponse
func (client *Client) DeleteAutoCcWhitelist(request *DeleteAutoCcWhitelistRequest) (_result *DeleteAutoCcWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAutoCcWhitelistResponse{}
	_body, _err := client.DeleteAutoCcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified forwarding rule of a website.
//
// @param request - DeleteDomainResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResourceResponse
func (client *Client) DeleteDomainResourceWithOptions(request *DeleteDomainResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainResourceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainResource"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified forwarding rule of a website.
//
// @param request - DeleteDomainResourceRequest
//
// @return DeleteDomainResourceResponse
func (client *Client) DeleteDomainResource(request *DeleteDomainResourceRequest) (_result *DeleteDomainResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainResourceResponse{}
	_body, _err := client.DeleteDomainResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a port forwarding rule. You can delete only one port forwarding rule at a time.
//
// @param request - DeleteNetworkRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkRuleResponse
func (client *Client) DeleteNetworkRuleWithOptions(request *DeleteNetworkRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkRule) {
		query["NetworkRule"] = request.NetworkRule
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a port forwarding rule. You can delete only one port forwarding rule at a time.
//
// @param request - DeleteNetworkRuleRequest
//
// @return DeleteNetworkRuleResponse
func (client *Client) DeleteNetworkRule(request *DeleteNetworkRuleRequest) (_result *DeleteNetworkRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkRuleResponse{}
	_body, _err := client.DeleteNetworkRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the specified port forwarding rule.
//
// Description:
//
// After you delete a port forwarding rule, the Anti-DDoS Pro or Anti-DDoS Premium instance no longer forwards service traffic on the Layer 4 port. Before you delete a specific port forwarding rule, make sure that the service traffic destined for the Layer 4 port is redirected to the origin server. This can prevent negative impacts on your services.
//
// > You can call this operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// @param request - DeletePortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePortResponse
func (client *Client) DeletePortWithOptions(request *DeletePortRequest, runtime *dara.RuntimeOptions) (_result *DeletePortResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendPort) {
		query["BackendPort"] = request.BackendPort
	}

	if !dara.IsNil(request.FrontendPort) {
		query["FrontendPort"] = request.FrontendPort
	}

	if !dara.IsNil(request.FrontendProtocol) {
		query["FrontendProtocol"] = request.FrontendProtocol
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RealServers) {
		query["RealServers"] = request.RealServers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePort"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified port forwarding rule.
//
// Description:
//
// After you delete a port forwarding rule, the Anti-DDoS Pro or Anti-DDoS Premium instance no longer forwards service traffic on the Layer 4 port. Before you delete a specific port forwarding rule, make sure that the service traffic destined for the Layer 4 port is redirected to the origin server. This can prevent negative impacts on your services.
//
// > You can call this operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// @param request - DeletePortRequest
//
// @return DeletePortResponse
func (client *Client) DeletePort(request *DeletePortRequest) (_result *DeletePortResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePortResponse{}
	_body, _err := client.DeletePortWithOptions(request, runtime)
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
// @param request - DeleteSceneDefensePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSceneDefensePolicyResponse
func (client *Client) DeleteSceneDefensePolicyWithOptions(request *DeleteSceneDefensePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteSceneDefensePolicyResponse, _err error) {
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
		Action:      dara.String("DeleteSceneDefensePolicy"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSceneDefensePolicyResponse{}
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
// @param request - DeleteSceneDefensePolicyRequest
//
// @return DeleteSceneDefensePolicyResponse
func (client *Client) DeleteSceneDefensePolicy(request *DeleteSceneDefensePolicyRequest) (_result *DeleteSceneDefensePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSceneDefensePolicyResponse{}
	_body, _err := client.DeleteSceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scheduling rule of Sec-Traffic Manager.
//
// @param request - DeleteSchedulerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSchedulerRuleResponse
func (client *Client) DeleteSchedulerRuleWithOptions(request *DeleteSchedulerRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteSchedulerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSchedulerRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSchedulerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scheduling rule of Sec-Traffic Manager.
//
// @param request - DeleteSchedulerRuleRequest
//
// @return DeleteSchedulerRuleResponse
func (client *Client) DeleteSchedulerRule(request *DeleteSchedulerRuleRequest) (_result *DeleteSchedulerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSchedulerRuleResponse{}
	_body, _err := client.DeleteSchedulerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from Anti-DDoS Proxy (Chinese Mainland) instances.
//
// Description:
//
// You can call the DeleteTagResources operation to remove tags from Anti-DDoS Proxy (Chinese Mainland) instances.
//
// >  Only Anti-DDoS Proxy (Chinese Mainland) supports tags.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTagResourcesResponse
func (client *Client) DeleteTagResourcesWithOptions(request *DeleteTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *DeleteTagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
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
		Action:      dara.String("DeleteTagResources"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from Anti-DDoS Proxy (Chinese Mainland) instances.
//
// Description:
//
// You can call the DeleteTagResources operation to remove tags from Anti-DDoS Proxy (Chinese Mainland) instances.
//
// >  Only Anti-DDoS Proxy (Chinese Mainland) supports tags.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteTagResourcesRequest
//
// @return DeleteTagResourcesResponse
func (client *Client) DeleteTagResources(request *DeleteTagResourcesRequest) (_result *DeleteTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTagResourcesResponse{}
	_body, _err := client.DeleteTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteWebCCRule is deprecated, please use ddoscoo::2020-01-01::DeleteWebCCRuleV2 instead.
//
// Summary:
//
// Deletes a custom frequency control rule of a website.
//
// @param request - DeleteWebCCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebCCRuleResponse
func (client *Client) DeleteWebCCRuleWithOptions(request *DeleteWebCCRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteWebCCRuleResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebCCRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebCCRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteWebCCRule is deprecated, please use ddoscoo::2020-01-01::DeleteWebCCRuleV2 instead.
//
// Summary:
//
// Deletes a custom frequency control rule of a website.
//
// @param request - DeleteWebCCRuleRequest
//
// @return DeleteWebCCRuleResponse
// Deprecated
func (client *Client) DeleteWebCCRule(request *DeleteWebCCRuleRequest) (_result *DeleteWebCCRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWebCCRuleResponse{}
	_body, _err := client.DeleteWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes custom frequency control rules of a website.
//
// @param request - DeleteWebCCRuleV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebCCRuleV2Response
func (client *Client) DeleteWebCCRuleV2WithOptions(request *DeleteWebCCRuleV2Request, runtime *dara.RuntimeOptions) (_result *DeleteWebCCRuleV2Response, _err error) {
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

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.RuleNames) {
		query["RuleNames"] = request.RuleNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebCCRuleV2"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebCCRuleV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes custom frequency control rules of a website.
//
// @param request - DeleteWebCCRuleV2Request
//
// @return DeleteWebCCRuleV2Response
func (client *Client) DeleteWebCCRuleV2(request *DeleteWebCCRuleV2Request) (_result *DeleteWebCCRuleV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWebCCRuleV2Response{}
	_body, _err := client.DeleteWebCCRuleV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the custom rules of the Static Page Caching policy for a website.
//
// Description:
//
// You can call the DeleteWebCacheCustomRule operation to delete the custom rules of the Static Page Caching policy for a website.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteWebCacheCustomRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebCacheCustomRuleResponse
func (client *Client) DeleteWebCacheCustomRuleWithOptions(request *DeleteWebCacheCustomRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteWebCacheCustomRuleResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RuleNames) {
		query["RuleNames"] = request.RuleNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebCacheCustomRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebCacheCustomRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the custom rules of the Static Page Caching policy for a website.
//
// Description:
//
// You can call the DeleteWebCacheCustomRule operation to delete the custom rules of the Static Page Caching policy for a website.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteWebCacheCustomRuleRequest
//
// @return DeleteWebCacheCustomRuleResponse
func (client *Client) DeleteWebCacheCustomRule(request *DeleteWebCacheCustomRuleRequest) (_result *DeleteWebCacheCustomRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWebCacheCustomRuleResponse{}
	_body, _err := client.DeleteWebCacheCustomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the accurate access control rules that are created for a website.
//
// @param request - DeleteWebPreciseAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebPreciseAccessRuleResponse
func (client *Client) DeleteWebPreciseAccessRuleWithOptions(request *DeleteWebPreciseAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteWebPreciseAccessRuleResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RuleNames) {
		query["RuleNames"] = request.RuleNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebPreciseAccessRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebPreciseAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the accurate access control rules that are created for a website.
//
// @param request - DeleteWebPreciseAccessRuleRequest
//
// @return DeleteWebPreciseAccessRuleResponse
func (client *Client) DeleteWebPreciseAccessRule(request *DeleteWebPreciseAccessRuleRequest) (_result *DeleteWebPreciseAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWebPreciseAccessRuleResponse{}
	_body, _err := client.DeleteWebPreciseAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a forwarding rule of a website.
//
// @param request - DeleteWebRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWebRuleResponse
func (client *Client) DeleteWebRuleWithOptions(request *DeleteWebRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteWebRuleResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWebRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWebRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a forwarding rule of a website.
//
// @param request - DeleteWebRuleRequest
//
// @return DeleteWebRuleResponse
func (client *Client) DeleteWebRule(request *DeleteWebRuleRequest) (_result *DeleteWebRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWebRuleResponse{}
	_body, _err := client.DeleteWebRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of asynchronous export tasks, such as the IDs, start time, end time, status, parameters, and results.
//
// Description:
//
// You can call the DescribeAsyncTasks operation to query the details of asynchronous export tasks, such as the IDs, start time, end time, status, parameters, and results.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeAsyncTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAsyncTasksResponse
func (client *Client) DescribeAsyncTasksWithOptions(request *DescribeAsyncTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAsyncTasksResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAsyncTasks"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAsyncTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of asynchronous export tasks, such as the IDs, start time, end time, status, parameters, and results.
//
// Description:
//
// You can call the DescribeAsyncTasks operation to query the details of asynchronous export tasks, such as the IDs, start time, end time, status, parameters, and results.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeAsyncTasksRequest
//
// @return DescribeAsyncTasksResponse
func (client *Client) DescribeAsyncTasks(request *DescribeAsyncTasksRequest) (_result *DescribeAsyncTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAsyncTasksResponse{}
	_body, _err := client.DescribeAsyncTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the peak QPS of DDoS attacks within the specific period of time.
//
// @param request - DescribeAttackAnalysisMaxQpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAttackAnalysisMaxQpsResponse
func (client *Client) DescribeAttackAnalysisMaxQpsWithOptions(request *DescribeAttackAnalysisMaxQpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAttackAnalysisMaxQpsResponse, _err error) {
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

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAttackAnalysisMaxQps"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAttackAnalysisMaxQpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the peak QPS of DDoS attacks within the specific period of time.
//
// @param request - DescribeAttackAnalysisMaxQpsRequest
//
// @return DescribeAttackAnalysisMaxQpsResponse
func (client *Client) DescribeAttackAnalysisMaxQps(request *DescribeAttackAnalysisMaxQpsRequest) (_result *DescribeAttackAnalysisMaxQpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAttackAnalysisMaxQpsResponse{}
	_body, _err := client.DescribeAttackAnalysisMaxQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IP addresses in the IP address blacklist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DescribeAutoCcBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoCcBlacklistResponse
func (client *Client) DescribeAutoCcBlacklistWithOptions(request *DescribeAutoCcBlacklistRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoCcBlacklistResponse, _err error) {
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

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoCcBlacklist"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoCcBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IP addresses in the IP address blacklist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DescribeAutoCcBlacklistRequest
//
// @return DescribeAutoCcBlacklistResponse
func (client *Client) DescribeAutoCcBlacklist(request *DescribeAutoCcBlacklistRequest) (_result *DescribeAutoCcBlacklistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoCcBlacklistResponse{}
	_body, _err := client.DescribeAutoCcBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the numbers of IP addresses in the IP address whitelist and IP address blacklist of an Anti-DDoS Proxy instance.
//
// @param request - DescribeAutoCcListCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoCcListCountResponse
func (client *Client) DescribeAutoCcListCountWithOptions(request *DescribeAutoCcListCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoCcListCountResponse, _err error) {
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

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoCcListCount"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoCcListCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the numbers of IP addresses in the IP address whitelist and IP address blacklist of an Anti-DDoS Proxy instance.
//
// @param request - DescribeAutoCcListCountRequest
//
// @return DescribeAutoCcListCountResponse
func (client *Client) DescribeAutoCcListCount(request *DescribeAutoCcListCountRequest) (_result *DescribeAutoCcListCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoCcListCountResponse{}
	_body, _err := client.DescribeAutoCcListCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IP addresses in the IP address whitelist of an Anti-DDoS Proxy instance.
//
// @param request - DescribeAutoCcWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoCcWhitelistResponse
func (client *Client) DescribeAutoCcWhitelistWithOptions(request *DescribeAutoCcWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoCcWhitelistResponse, _err error) {
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

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
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
		Action:      dara.String("DescribeAutoCcWhitelist"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoCcWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IP addresses in the IP address whitelist of an Anti-DDoS Proxy instance.
//
// @param request - DescribeAutoCcWhitelistRequest
//
// @return DescribeAutoCcWhitelistResponse
func (client *Client) DescribeAutoCcWhitelist(request *DescribeAutoCcWhitelistRequest) (_result *DescribeAutoCcWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoCcWhitelistResponse{}
	_body, _err := client.DescribeAutoCcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin CIDR blocks of Anti-DDoS Proxy.
//
// @param request - DescribeBackSourceCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackSourceCidrResponse
func (client *Client) DescribeBackSourceCidrWithOptions(request *DescribeBackSourceCidrRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackSourceCidrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Line) {
		query["Line"] = request.Line
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackSourceCidr"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackSourceCidrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin CIDR blocks of Anti-DDoS Proxy.
//
// @param request - DescribeBackSourceCidrRequest
//
// @return DescribeBackSourceCidrResponse
func (client *Client) DescribeBackSourceCidr(request *DescribeBackSourceCidrRequest) (_result *DescribeBackSourceCidrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackSourceCidrResponse{}
	_body, _err := client.DescribeBackSourceCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the blackhole filtering status of one or more Anti-DDoS Pro or Anti-DDoS Premium instances.
//
// @param request - DescribeBlackholeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBlackholeStatusResponse
func (client *Client) DescribeBlackholeStatusWithOptions(request *DescribeBlackholeStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeBlackholeStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBlackholeStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBlackholeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the blackhole filtering status of one or more Anti-DDoS Pro or Anti-DDoS Premium instances.
//
// @param request - DescribeBlackholeStatusRequest
//
// @return DescribeBlackholeStatusResponse
func (client *Client) DescribeBlackholeStatus(request *DescribeBlackholeStatusRequest) (_result *DescribeBlackholeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBlackholeStatusResponse{}
	_body, _err := client.DescribeBlackholeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Diversion from Origin Server configurations of one or more Anti-DDoS Proxy (Chinese Mainland) instances.
//
// Description:
//
// You can call this operation to query the Diversion from Origin Server configurations of one or more Anti-DDoS Proxy (Chinese Mainland) instances.
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeBlockStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBlockStatusResponse
func (client *Client) DescribeBlockStatusWithOptions(request *DescribeBlockStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeBlockStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBlockStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBlockStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Diversion from Origin Server configurations of one or more Anti-DDoS Proxy (Chinese Mainland) instances.
//
// Description:
//
// You can call this operation to query the Diversion from Origin Server configurations of one or more Anti-DDoS Proxy (Chinese Mainland) instances.
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeBlockStatusRequest
//
// @return DescribeBlockStatusResponse
func (client *Client) DescribeBlockStatus(request *DescribeBlockStatusRequest) (_result *DescribeBlockStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBlockStatusResponse{}
	_body, _err := client.DescribeBlockStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeCdnLinkageRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnLinkageRulesResponse
func (client *Client) DescribeCdnLinkageRulesWithOptions(request *DescribeCdnLinkageRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnLinkageRulesResponse, _err error) {
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
		Action:      dara.String("DescribeCdnLinkageRules"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnLinkageRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeCdnLinkageRulesRequest
//
// @return DescribeCdnLinkageRulesResponse
func (client *Client) DescribeCdnLinkageRules(request *DescribeCdnLinkageRulesRequest) (_result *DescribeCdnLinkageRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnLinkageRulesResponse{}
	_body, _err := client.DescribeCdnLinkageRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all certificates that can be associated with the current domain name instead of the certificate currently in use.
//
// Description:
//
// This operation is used to query all applicable certificates of a domain name that you want to add to Anti-DDoS Proxy. Multiple certificates may be queried for a domain name. You can use an exact domain name to query exact-domain certificates or wildcard-domain certificates.
//
// >  If you want to query the certificate that is in use for the current domain name, you can call the [DescribeWebRules](https://help.aliyun.com/document_detail/473610.html) operation to obtain the values of the CertName and CertRegion parameters. Then, you can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/411733.html) operation of Certificate Management Service to query the ID and other details of the certificate by using the value of the CertName parameter.
//
// @param request - DescribeCertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertsResponse
func (client *Client) DescribeCertsWithOptions(request *DescribeCertsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCertsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCerts"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all certificates that can be associated with the current domain name instead of the certificate currently in use.
//
// Description:
//
// This operation is used to query all applicable certificates of a domain name that you want to add to Anti-DDoS Proxy. Multiple certificates may be queried for a domain name. You can use an exact domain name to query exact-domain certificates or wildcard-domain certificates.
//
// >  If you want to query the certificate that is in use for the current domain name, you can call the [DescribeWebRules](https://help.aliyun.com/document_detail/473610.html) operation to obtain the values of the CertName and CertRegion parameters. Then, you can call the [ListUserCertificateOrder](https://help.aliyun.com/document_detail/411733.html) operation of Certificate Management Service to query the ID and other details of the certificate by using the value of the CertName parameter.
//
// @param request - DescribeCertsRequest
//
// @return DescribeCertsResponse
func (client *Client) DescribeCerts(request *DescribeCertsRequest) (_result *DescribeCertsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCertsResponse{}
	_body, _err := client.DescribeCertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeCnameReusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCnameReusesResponse
func (client *Client) DescribeCnameReusesWithOptions(request *DescribeCnameReusesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCnameReusesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCnameReuses"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCnameReusesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeCnameReusesRequest
//
// @return DescribeCnameReusesResponse
func (client *Client) DescribeCnameReuses(request *DescribeCnameReusesRequest) (_result *DescribeCnameReusesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCnameReusesResponse{}
	_body, _err := client.DescribeCnameReusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attack events launched against one or more Anti-DDoS Proxy instances.
//
// @param request - DescribeDDoSEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDoSEventsResponse
func (client *Client) DescribeDDoSEventsWithOptions(request *DescribeDDoSEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDoSEventsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDDoSEvents"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDoSEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attack events launched against one or more Anti-DDoS Proxy instances.
//
// @param request - DescribeDDoSEventsRequest
//
// @return DescribeDDoSEventsResponse
func (client *Client) DescribeDDoSEvents(request *DescribeDDoSEventsRequest) (_result *DescribeDDoSEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDoSEventsResponse{}
	_body, _err := client.DescribeDDoSEventsWithOptions(request, runtime)
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
// Description:
//
// You can call the DescribeDDosAllEventList operation to query DDoS attack events within a specific time range by page. The information about a DDoS attack event includes the start time and end time of the attack, attack event type, attacked object, peak bandwidth of attack traffic, and peak packet forwarding rate.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDDosAllEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDosAllEventListResponse
func (client *Client) DescribeDDosAllEventListWithOptions(request *DescribeDDosAllEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDosAllEventListResponse, _err error) {
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

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
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
		Action:      dara.String("DescribeDDosAllEventList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDosAllEventListResponse{}
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
// Description:
//
// You can call the DescribeDDosAllEventList operation to query DDoS attack events within a specific time range by page. The information about a DDoS attack event includes the start time and end time of the attack, attack event type, attacked object, peak bandwidth of attack traffic, and peak packet forwarding rate.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDDosAllEventListRequest
//
// @return DescribeDDosAllEventListResponse
func (client *Client) DescribeDDosAllEventList(request *DescribeDDosAllEventListRequest) (_result *DescribeDDosAllEventListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDosAllEventListResponse{}
	_body, _err := client.DescribeDDosAllEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the source region from which a volumetric attack is initiated.
//
// Description:
//
// > This operation is suitable only for volumetric attacks.
//
// @param request - DescribeDDosEventAreaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDosEventAreaResponse
func (client *Client) DescribeDDosEventAreaWithOptions(request *DescribeDDosEventAreaRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDosEventAreaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Range) {
		query["Range"] = request.Range
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDDosEventArea"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDosEventAreaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the source region from which a volumetric attack is initiated.
//
// Description:
//
// > This operation is suitable only for volumetric attacks.
//
// @param request - DescribeDDosEventAreaRequest
//
// @return DescribeDDosEventAreaResponse
func (client *Client) DescribeDDosEventArea(request *DescribeDDosEventAreaRequest) (_result *DescribeDDosEventAreaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDosEventAreaResponse{}
	_body, _err := client.DescribeDDosEventAreaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attack type details of a volumetric attack.
//
// Description:
//
// > This operation is suitable only for volumetric attacks.
//
// @param request - DescribeDDosEventAttackTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDosEventAttackTypeResponse
func (client *Client) DescribeDDosEventAttackTypeWithOptions(request *DescribeDDosEventAttackTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDosEventAttackTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDDosEventAttackType"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDosEventAttackTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attack type details of a volumetric attack.
//
// Description:
//
// > This operation is suitable only for volumetric attacks.
//
// @param request - DescribeDDosEventAttackTypeRequest
//
// @return DescribeDDosEventAttackTypeResponse
func (client *Client) DescribeDDosEventAttackType(request *DescribeDDosEventAttackTypeRequest) (_result *DescribeDDosEventAttackTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDosEventAttackTypeResponse{}
	_body, _err := client.DescribeDDosEventAttackTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Internet service provider (ISP) information about a volumetric attack.
//
// Description:
//
// > This operation is suitable only for volumetric attacks.
//
// @param request - DescribeDDosEventIspRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDosEventIspResponse
func (client *Client) DescribeDDosEventIspWithOptions(request *DescribeDDosEventIspRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDosEventIspResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Range) {
		query["Range"] = request.Range
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDDosEventIsp"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDosEventIspResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Internet service provider (ISP) information about a volumetric attack.
//
// Description:
//
// > This operation is suitable only for volumetric attacks.
//
// @param request - DescribeDDosEventIspRequest
//
// @return DescribeDDosEventIspResponse
func (client *Client) DescribeDDosEventIsp(request *DescribeDDosEventIspRequest) (_result *DescribeDDosEventIspResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDosEventIspResponse{}
	_body, _err := client.DescribeDDosEventIspWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the peaks of volumetric attacks (bit/s), connection flood attacks (CPS), and resource exhaustion attacks on websites (QPS).
//
// @param request - DescribeDDosEventMaxRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDosEventMaxResponse
func (client *Client) DescribeDDosEventMaxWithOptions(request *DescribeDDosEventMaxRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDosEventMaxResponse, _err error) {
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

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDDosEventMax"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDosEventMaxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the peaks of volumetric attacks (bit/s), connection flood attacks (CPS), and resource exhaustion attacks on websites (QPS).
//
// @param request - DescribeDDosEventMaxRequest
//
// @return DescribeDDosEventMaxResponse
func (client *Client) DescribeDDosEventMax(request *DescribeDDosEventMaxRequest) (_result *DescribeDDosEventMaxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDosEventMaxResponse{}
	_body, _err := client.DescribeDDosEventMaxWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the source IP address from which a volumetric attack is initiated.
//
// Description:
//
// > This operation is suitable only for volumetric attacks.
//
// @param request - DescribeDDosEventSrcIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDosEventSrcIpResponse
func (client *Client) DescribeDDosEventSrcIpWithOptions(request *DescribeDDosEventSrcIpRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDosEventSrcIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Range) {
		query["Range"] = request.Range
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDDosEventSrcIp"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDosEventSrcIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the source IP address from which a volumetric attack is initiated.
//
// Description:
//
// > This operation is suitable only for volumetric attacks.
//
// @param request - DescribeDDosEventSrcIpRequest
//
// @return DescribeDDosEventSrcIpResponse
func (client *Client) DescribeDDosEventSrcIp(request *DescribeDDosEventSrcIpRequest) (_result *DescribeDDosEventSrcIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDosEventSrcIpResponse{}
	_body, _err := client.DescribeDDosEventSrcIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on advanced mitigation sessions of an Anti-DDoS Proxy (Outside Chinese Mainland) instance.
//
// Description:
//
// You can call the DescribeDefenseCountStatistics operation to query the information about advanced mitigation sessions of an Anti-DDoS Proxy (Outside Chinese Mainland) instance. For example, you can query the number of advanced mitigation sessions that are used within the current calendar month and the number of remaining advanced mitigation sessions.
//
// >  This operation is suitable only for Anti-DDoS Proxy (Outside Chinese Mainland).
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDefenseCountStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseCountStatisticsResponse
func (client *Client) DescribeDefenseCountStatisticsWithOptions(request *DescribeDefenseCountStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseCountStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseCountStatistics"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseCountStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on advanced mitigation sessions of an Anti-DDoS Proxy (Outside Chinese Mainland) instance.
//
// Description:
//
// You can call the DescribeDefenseCountStatistics operation to query the information about advanced mitigation sessions of an Anti-DDoS Proxy (Outside Chinese Mainland) instance. For example, you can query the number of advanced mitigation sessions that are used within the current calendar month and the number of remaining advanced mitigation sessions.
//
// >  This operation is suitable only for Anti-DDoS Proxy (Outside Chinese Mainland).
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDefenseCountStatisticsRequest
//
// @return DescribeDefenseCountStatisticsResponse
func (client *Client) DescribeDefenseCountStatistics(request *DescribeDefenseCountStatisticsRequest) (_result *DescribeDefenseCountStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseCountStatisticsResponse{}
	_body, _err := client.DescribeDefenseCountStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the advanced mitigation logs of an Anti-DDoS Proxy (Outside Chinese Mainland) instance.
//
// Description:
//
// > This operation is suitable only for Anti-DDoS Premium.
//
// @param request - DescribeDefenseRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseRecordsResponse
func (client *Client) DescribeDefenseRecordsWithOptions(request *DescribeDefenseRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseRecordsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseRecords"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the advanced mitigation logs of an Anti-DDoS Proxy (Outside Chinese Mainland) instance.
//
// Description:
//
// > This operation is suitable only for Anti-DDoS Premium.
//
// @param request - DescribeDefenseRecordsRequest
//
// @return DescribeDefenseRecordsResponse
func (client *Client) DescribeDefenseRecords(request *DescribeDefenseRecordsRequest) (_result *DescribeDefenseRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseRecordsResponse{}
	_body, _err := client.DescribeDefenseRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of request packets received by the destination ports of the attacked IP address that is protected by Anti-DDoS Proxy.
//
// @param request - DescribeDestinationPortEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDestinationPortEventResponse
func (client *Client) DescribeDestinationPortEventWithOptions(request *DescribeDestinationPortEventRequest, runtime *dara.RuntimeOptions) (_result *DescribeDestinationPortEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Range) {
		query["Range"] = request.Range
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDestinationPortEvent"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDestinationPortEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of request packets received by the destination ports of the attacked IP address that is protected by Anti-DDoS Proxy.
//
// @param request - DescribeDestinationPortEventRequest
//
// @return DescribeDestinationPortEventResponse
func (client *Client) DescribeDestinationPortEvent(request *DescribeDestinationPortEventRequest) (_result *DescribeDestinationPortEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDestinationPortEventResponse{}
	_body, _err := client.DescribeDestinationPortEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attack events launched against a website.
//
// @param request - DescribeDomainAttackEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainAttackEventsResponse
func (client *Client) DescribeDomainAttackEventsWithOptions(request *DescribeDomainAttackEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainAttackEventsResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainAttackEvents"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainAttackEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attack events launched against a website.
//
// @param request - DescribeDomainAttackEventsRequest
//
// @return DescribeDomainAttackEventsResponse
func (client *Client) DescribeDomainAttackEvents(request *DescribeDomainAttackEventsRequest) (_result *DescribeDomainAttackEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainAttackEventsResponse{}
	_body, _err := client.DescribeDomainAttackEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bandwidths of a website.
//
// @param request - DescribeDomainBpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainBpsResponse
func (client *Client) DescribeDomainBpsWithOptions(request *DescribeDomainBpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainBpsResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainBps"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainBpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidths of a website.
//
// @param request - DescribeDomainBpsRequest
//
// @return DescribeDomainBpsResponse
func (client *Client) DescribeDomainBps(request *DescribeDomainBpsRequest) (_result *DescribeDomainBpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainBpsResponse{}
	_body, _err := client.DescribeDomainBpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about HTTP/2 fingerprints of a website.
//
// @param request - DescribeDomainH2FingerprintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainH2FingerprintResponse
func (client *Client) DescribeDomainH2FingerprintWithOptions(request *DescribeDomainH2FingerprintRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainH2FingerprintResponse, _err error) {
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
		Action:      dara.String("DescribeDomainH2Fingerprint"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainH2FingerprintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about HTTP/2 fingerprints of a website.
//
// @param request - DescribeDomainH2FingerprintRequest
//
// @return DescribeDomainH2FingerprintResponse
func (client *Client) DescribeDomainH2Fingerprint(request *DescribeDomainH2FingerprintRequest) (_result *DescribeDomainH2FingerprintResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainH2FingerprintResponse{}
	_body, _err := client.DescribeDomainH2FingerprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attack overview of a website, such as the peak HTTP and HTTPS traffic.
//
// @param request - DescribeDomainOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainOverviewResponse
func (client *Client) DescribeDomainOverviewWithOptions(request *DescribeDomainOverviewRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainOverviewResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainOverview"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attack overview of a website, such as the peak HTTP and HTTPS traffic.
//
// @param request - DescribeDomainOverviewRequest
//
// @return DescribeDomainOverviewResponse
func (client *Client) DescribeDomainOverview(request *DescribeDomainOverviewRequest) (_result *DescribeDomainOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainOverviewResponse{}
	_body, _err := client.DescribeDomainOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on the queries per second (QPS) of a website.
//
// @param request - DescribeDomainQPSListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainQPSListResponse
func (client *Client) DescribeDomainQPSListWithOptions(request *DescribeDomainQPSListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainQPSListResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainQPSList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainQPSListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the queries per second (QPS) of a website.
//
// @param request - DescribeDomainQPSListRequest
//
// @return DescribeDomainQPSListResponse
func (client *Client) DescribeDomainQPSList(request *DescribeDomainQPSListRequest) (_result *DescribeDomainQPSListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainQPSListResponse{}
	_body, _err := client.DescribeDomainQPSListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a forwarding rule.
//
// Description:
//
// You can call the DescribeDomainResource operation to query the configurations of the forwarding rules that you create for a website by page. The configurations include the domain name-related configurations, protocol-related configurations, HTTPS-related configurations, and configurations that are used to mitigate HTTP flood attacks.
//
// You can call this operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// ### Limits
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDomainResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainResourceResponse
func (client *Client) DescribeDomainResourceWithOptions(request *DescribeDomainResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainResourceResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryDomainPattern) {
		query["QueryDomainPattern"] = request.QueryDomainPattern
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainResource"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a forwarding rule.
//
// Description:
//
// You can call the DescribeDomainResource operation to query the configurations of the forwarding rules that you create for a website by page. The configurations include the domain name-related configurations, protocol-related configurations, HTTPS-related configurations, and configurations that are used to mitigate HTTP flood attacks.
//
// You can call this operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// ### Limits
//
// You can call this operation up to 50 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeDomainResourceRequest
//
// @return DescribeDomainResourceResponse
func (client *Client) DescribeDomainResource(request *DescribeDomainResourceRequest) (_result *DescribeDomainResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainResourceResponse{}
	_body, _err := client.DescribeDomainResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the global mitigation policy for a domain name.
//
// @param request - DescribeDomainSecurityProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainSecurityProfileResponse
func (client *Client) DescribeDomainSecurityProfileWithOptions(request *DescribeDomainSecurityProfileRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainSecurityProfileResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainSecurityProfile"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainSecurityProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global mitigation policy for a domain name.
//
// @param request - DescribeDomainSecurityProfileRequest
//
// @return DescribeDomainSecurityProfileResponse
func (client *Client) DescribeDomainSecurityProfile(request *DescribeDomainSecurityProfileRequest) (_result *DescribeDomainSecurityProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainSecurityProfileResponse{}
	_body, _err := client.DescribeDomainSecurityProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on HTTP status codes of a website within a specified period of time.
//
// @param request - DescribeDomainStatusCodeCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainStatusCodeCountResponse
func (client *Client) DescribeDomainStatusCodeCountWithOptions(request *DescribeDomainStatusCodeCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainStatusCodeCountResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainStatusCodeCount"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainStatusCodeCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on HTTP status codes of a website within a specified period of time.
//
// @param request - DescribeDomainStatusCodeCountRequest
//
// @return DescribeDomainStatusCodeCountResponse
func (client *Client) DescribeDomainStatusCodeCount(request *DescribeDomainStatusCodeCountRequest) (_result *DescribeDomainStatusCodeCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainStatusCodeCountResponse{}
	_body, _err := client.DescribeDomainStatusCodeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on HTTP status codes of a website.
//
// @param request - DescribeDomainStatusCodeListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainStatusCodeListResponse
func (client *Client) DescribeDomainStatusCodeListWithOptions(request *DescribeDomainStatusCodeListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainStatusCodeListResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainStatusCodeList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainStatusCodeListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on HTTP status codes of a website.
//
// @param request - DescribeDomainStatusCodeListRequest
//
// @return DescribeDomainStatusCodeListResponse
func (client *Client) DescribeDomainStatusCodeList(request *DescribeDomainStatusCodeListRequest) (_result *DescribeDomainStatusCodeListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainStatusCodeListResponse{}
	_body, _err := client.DescribeDomainStatusCodeListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the peak queries per second (QPS) information about a website, such as the attack QPS and total QPS, within a specific period of time.
//
// @param request - DescribeDomainTopAttackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainTopAttackListResponse
func (client *Client) DescribeDomainTopAttackListWithOptions(request *DescribeDomainTopAttackListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainTopAttackListResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainTopAttackList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainTopAttackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the peak queries per second (QPS) information about a website, such as the attack QPS and total QPS, within a specific period of time.
//
// @param request - DescribeDomainTopAttackListRequest
//
// @return DescribeDomainTopAttackListResponse
func (client *Client) DescribeDomainTopAttackList(request *DescribeDomainTopAttackListRequest) (_result *DescribeDomainTopAttackListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainTopAttackListResponse{}
	_body, _err := client.DescribeDomainTopAttackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the fingerprints of top N clients that access a website.
//
// @param request - DescribeDomainTopFingerprintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainTopFingerprintResponse
func (client *Client) DescribeDomainTopFingerprintWithOptions(request *DescribeDomainTopFingerprintRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainTopFingerprintResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainTopFingerprint"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainTopFingerprintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the fingerprints of top N clients that access a website.
//
// @param request - DescribeDomainTopFingerprintRequest
//
// @return DescribeDomainTopFingerprintResponse
func (client *Client) DescribeDomainTopFingerprint(request *DescribeDomainTopFingerprintRequest) (_result *DescribeDomainTopFingerprintResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainTopFingerprintResponse{}
	_body, _err := client.DescribeDomainTopFingerprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about top N HTTP methods of a website.
//
// @param request - DescribeDomainTopHttpMethodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainTopHttpMethodResponse
func (client *Client) DescribeDomainTopHttpMethodWithOptions(request *DescribeDomainTopHttpMethodRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainTopHttpMethodResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainTopHttpMethod"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainTopHttpMethodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about top N HTTP methods of a website.
//
// @param request - DescribeDomainTopHttpMethodRequest
//
// @return DescribeDomainTopHttpMethodResponse
func (client *Client) DescribeDomainTopHttpMethod(request *DescribeDomainTopHttpMethodRequest) (_result *DescribeDomainTopHttpMethodResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainTopHttpMethodResponse{}
	_body, _err := client.DescribeDomainTopHttpMethodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about top N referers of a website.
//
// @param request - DescribeDomainTopRefererRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainTopRefererResponse
func (client *Client) DescribeDomainTopRefererWithOptions(request *DescribeDomainTopRefererRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainTopRefererResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainTopReferer"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainTopRefererResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about top N referers of a website.
//
// @param request - DescribeDomainTopRefererRequest
//
// @return DescribeDomainTopRefererResponse
func (client *Client) DescribeDomainTopReferer(request *DescribeDomainTopRefererRequest) (_result *DescribeDomainTopRefererResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainTopRefererResponse{}
	_body, _err := client.DescribeDomainTopRefererWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about top user agents of a website.
//
// @param request - DescribeDomainTopUserAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainTopUserAgentResponse
func (client *Client) DescribeDomainTopUserAgentWithOptions(request *DescribeDomainTopUserAgentRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainTopUserAgentResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainTopUserAgent"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainTopUserAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about top user agents of a website.
//
// @param request - DescribeDomainTopUserAgentRequest
//
// @return DescribeDomainTopUserAgentResponse
func (client *Client) DescribeDomainTopUserAgent(request *DescribeDomainTopUserAgentRequest) (_result *DescribeDomainTopUserAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainTopUserAgentResponse{}
	_body, _err := client.DescribeDomainTopUserAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the areas and countries from which requests are sent to a website within a specified period of time.
//
// @param request - DescribeDomainViewSourceCountriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainViewSourceCountriesResponse
func (client *Client) DescribeDomainViewSourceCountriesWithOptions(request *DescribeDomainViewSourceCountriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainViewSourceCountriesResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainViewSourceCountries"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainViewSourceCountriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the areas and countries from which requests are sent to a website within a specified period of time.
//
// @param request - DescribeDomainViewSourceCountriesRequest
//
// @return DescribeDomainViewSourceCountriesResponse
func (client *Client) DescribeDomainViewSourceCountries(request *DescribeDomainViewSourceCountriesRequest) (_result *DescribeDomainViewSourceCountriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainViewSourceCountriesResponse{}
	_body, _err := client.DescribeDomainViewSourceCountriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the administrative regions in China from which requests are sent to a website within a specified period of time.
//
// @param request - DescribeDomainViewSourceProvincesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainViewSourceProvincesResponse
func (client *Client) DescribeDomainViewSourceProvincesWithOptions(request *DescribeDomainViewSourceProvincesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainViewSourceProvincesResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainViewSourceProvinces"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainViewSourceProvincesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the administrative regions in China from which requests are sent to a website within a specified period of time.
//
// @param request - DescribeDomainViewSourceProvincesRequest
//
// @return DescribeDomainViewSourceProvincesResponse
func (client *Client) DescribeDomainViewSourceProvinces(request *DescribeDomainViewSourceProvincesRequest) (_result *DescribeDomainViewSourceProvincesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainViewSourceProvincesResponse{}
	_body, _err := client.DescribeDomainViewSourceProvincesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top N URLs that require the longest time to respond to requests within a specified period of time.
//
// @param request - DescribeDomainViewTopCostTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainViewTopCostTimeResponse
func (client *Client) DescribeDomainViewTopCostTimeWithOptions(request *DescribeDomainViewTopCostTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainViewTopCostTimeResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Top) {
		query["Top"] = request.Top
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainViewTopCostTime"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainViewTopCostTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top N URLs that require the longest time to respond to requests within a specified period of time.
//
// @param request - DescribeDomainViewTopCostTimeRequest
//
// @return DescribeDomainViewTopCostTimeResponse
func (client *Client) DescribeDomainViewTopCostTime(request *DescribeDomainViewTopCostTimeRequest) (_result *DescribeDomainViewTopCostTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainViewTopCostTimeResponse{}
	_body, _err := client.DescribeDomainViewTopCostTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top N URLs that receive the most requests within a specified period of time.
//
// @param request - DescribeDomainViewTopUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainViewTopUrlResponse
func (client *Client) DescribeDomainViewTopUrlWithOptions(request *DescribeDomainViewTopUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainViewTopUrlResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Inerval) {
		query["Inerval"] = request.Inerval
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Top) {
		query["Top"] = request.Top
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainViewTopUrl"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainViewTopUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top N URLs that receive the most requests within a specified period of time.
//
// @param request - DescribeDomainViewTopUrlRequest
//
// @return DescribeDomainViewTopUrlResponse
func (client *Client) DescribeDomainViewTopUrl(request *DescribeDomainViewTopUrlRequest) (_result *DescribeDomainViewTopUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainViewTopUrlResponse{}
	_body, _err := client.DescribeDomainViewTopUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain names for which forwarding rules are created.
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
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomains"),
		Version:     dara.String("2020-01-01"),
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
// Queries domain names for which forwarding rules are created.
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
// Queries the available burstable protection bandwidths of an Anti-DDoS Proxy (Chinese Mainland) instance.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - DescribeElasticBandwidthSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticBandwidthSpecResponse
func (client *Client) DescribeElasticBandwidthSpecWithOptions(request *DescribeElasticBandwidthSpecRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticBandwidthSpecResponse, _err error) {
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
		Action:      dara.String("DescribeElasticBandwidthSpec"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticBandwidthSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available burstable protection bandwidths of an Anti-DDoS Proxy (Chinese Mainland) instance.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - DescribeElasticBandwidthSpecRequest
//
// @return DescribeElasticBandwidthSpecResponse
func (client *Client) DescribeElasticBandwidthSpec(request *DescribeElasticBandwidthSpecRequest) (_result *DescribeElasticBandwidthSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeElasticBandwidthSpecResponse{}
	_body, _err := client.DescribeElasticBandwidthSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the line chart of the bills for the burstable QPS of an Anti-DDoS Proxy instance.
//
// @param request - DescribeElasticQpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticQpsResponse
func (client *Client) DescribeElasticQpsWithOptions(request *DescribeElasticQpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticQpsResponse, _err error) {
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

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeElasticQps"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticQpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the line chart of the bills for the burstable QPS of an Anti-DDoS Proxy instance.
//
// @param request - DescribeElasticQpsRequest
//
// @return DescribeElasticQpsResponse
func (client *Client) DescribeElasticQps(request *DescribeElasticQpsRequest) (_result *DescribeElasticQpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeElasticQpsResponse{}
	_body, _err := client.DescribeElasticQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the burstable QPS details of an Anti-DDoS Proxy instance.
//
// @param request - DescribeElasticQpsRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticQpsRecordResponse
func (client *Client) DescribeElasticQpsRecordWithOptions(request *DescribeElasticQpsRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticQpsRecordResponse, _err error) {
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

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeElasticQpsRecord"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticQpsRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the burstable QPS details of an Anti-DDoS Proxy instance.
//
// @param request - DescribeElasticQpsRecordRequest
//
// @return DescribeElasticQpsRecordResponse
func (client *Client) DescribeElasticQpsRecord(request *DescribeElasticQpsRecordRequest) (_result *DescribeElasticQpsRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeElasticQpsRecordResponse{}
	_body, _err := client.DescribeElasticQpsRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the custom header that is specified for a domain name.
//
// @param request - DescribeHeadersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHeadersResponse
func (client *Client) DescribeHeadersWithOptions(request *DescribeHeadersRequest, runtime *dara.RuntimeOptions) (_result *DescribeHeadersResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHeaders"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHeadersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom header that is specified for a domain name.
//
// @param request - DescribeHeadersRequest
//
// @return DescribeHeadersResponse
func (client *Client) DescribeHeaders(request *DescribeHeadersRequest) (_result *DescribeHeadersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHeadersResponse{}
	_body, _err := client.DescribeHeadersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Layer 4 or Layer 7 health check configurations of a port forwarding rule.
//
// @param request - DescribeHealthCheckListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHealthCheckListResponse
func (client *Client) DescribeHealthCheckListWithOptions(request *DescribeHealthCheckListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHealthCheckListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkRules) {
		query["NetworkRules"] = request.NetworkRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHealthCheckList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHealthCheckListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Layer 4 or Layer 7 health check configurations of a port forwarding rule.
//
// @param request - DescribeHealthCheckListRequest
//
// @return DescribeHealthCheckListResponse
func (client *Client) DescribeHealthCheckList(request *DescribeHealthCheckListRequest) (_result *DescribeHealthCheckListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHealthCheckListResponse{}
	_body, _err := client.DescribeHealthCheckListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the health status of an origin server.
//
// @param request - DescribeHealthCheckStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHealthCheckStatusResponse
func (client *Client) DescribeHealthCheckStatusWithOptions(request *DescribeHealthCheckStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeHealthCheckStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkRules) {
		query["NetworkRules"] = request.NetworkRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHealthCheckStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHealthCheckStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health status of an origin server.
//
// @param request - DescribeHealthCheckStatusRequest
//
// @return DescribeHealthCheckStatusResponse
func (client *Client) DescribeHealthCheckStatus(request *DescribeHealthCheckStatusRequest) (_result *DescribeHealthCheckStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHealthCheckStatusResponse{}
	_body, _err := client.DescribeHealthCheckStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP addresses and Internet service provider (ISP) lines of Anti-DDoS Pro or Anti-DDoS Premium instances.
//
// Description:
//
// You can call the DescribeInstanceDetails operation to query the information about the IP addresses and ISP lines of the instances. The information includes the IP address, status, and protection line.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDetailsResponse
func (client *Client) DescribeInstanceDetailsWithOptions(request *DescribeInstanceDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceDetails"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses and Internet service provider (ISP) lines of Anti-DDoS Pro or Anti-DDoS Premium instances.
//
// Description:
//
// You can call the DescribeInstanceDetails operation to query the information about the IP addresses and ISP lines of the instances. The information includes the IP address, status, and protection line.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceDetailsRequest
//
// @return DescribeInstanceDetailsResponse
func (client *Client) DescribeInstanceDetails(request *DescribeInstanceDetailsRequest) (_result *DescribeInstanceDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceDetailsResponse{}
	_body, _err := client.DescribeInstanceDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about Anti-DDoS Pro and Anti-DDoS Premium instances.
//
// @param request - DescribeInstanceExtRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceExtResponse
func (client *Client) DescribeInstanceExtWithOptions(request *DescribeInstanceExtRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceExtResponse, _err error) {
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
		Action:      dara.String("DescribeInstanceExt"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceExtResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Anti-DDoS Pro and Anti-DDoS Premium instances.
//
// @param request - DescribeInstanceExtRequest
//
// @return DescribeInstanceExtResponse
func (client *Client) DescribeInstanceExt(request *DescribeInstanceExtRequest) (_result *DescribeInstanceExtResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceExtResponse{}
	_body, _err := client.DescribeInstanceExtWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The description of the instance.
//
// @param request - DescribeInstanceIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceIdsResponse
func (client *Client) DescribeInstanceIdsWithOptions(request *DescribeInstanceIdsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceIdsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceIds"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The description of the instance.
//
// @param request - DescribeInstanceIdsRequest
//
// @return DescribeInstanceIdsResponse
func (client *Client) DescribeInstanceIds(request *DescribeInstanceIdsRequest) (_result *DescribeInstanceIdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceIdsResponse{}
	_body, _err := client.DescribeInstanceIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the specifications of Anti-DDoS Pro or Anti-DDoS Premium instances.
//
// Description:
//
// You can call the DescribeInstanceSpecs operation to query the specifications of multiple Anti-DDoS Pro or Anti-DDoS Premium instances at a time. The specifications include the clean bandwidth, protection bandwidth, function plan, and the numbers of domain names and ports that can be protected.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceSpecsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSpecsResponse
func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSpecs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specifications of Anti-DDoS Pro or Anti-DDoS Premium instances.
//
// Description:
//
// You can call the DescribeInstanceSpecs operation to query the specifications of multiple Anti-DDoS Pro or Anti-DDoS Premium instances at a time. The specifications include the clean bandwidth, protection bandwidth, function plan, and the numbers of domain names and ports that can be protected.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeInstanceSpecsRequest
//
// @return DescribeInstanceSpecsResponse
func (client *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (_result *DescribeInstanceSpecsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DescribeInstanceSpecsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on one or more Anti-DDoS Proxy instances, such as the numbers of protected domain names and ports.
//
// @param request - DescribeInstanceStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceStatisticsResponse
func (client *Client) DescribeInstanceStatisticsWithOptions(request *DescribeInstanceStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceStatistics"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on one or more Anti-DDoS Proxy instances, such as the numbers of protected domain names and ports.
//
// @param request - DescribeInstanceStatisticsRequest
//
// @return DescribeInstanceStatisticsResponse
func (client *Client) DescribeInstanceStatistics(request *DescribeInstanceStatisticsRequest) (_result *DescribeInstanceStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceStatisticsResponse{}
	_body, _err := client.DescribeInstanceStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a specified Anti-DDoS Proxy instance.
//
// @param request - DescribeInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceStatusResponse
func (client *Client) DescribeInstanceStatusWithOptions(request *DescribeInstanceStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceStatusResponse, _err error) {
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

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a specified Anti-DDoS Proxy instance.
//
// @param request - DescribeInstanceStatusRequest
//
// @return DescribeInstanceStatusResponse
func (client *Client) DescribeInstanceStatus(request *DescribeInstanceStatusRequest) (_result *DescribeInstanceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceStatusResponse{}
	_body, _err := client.DescribeInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例列表
//
// Description:
//
// You can call the DescribeInstances operation to query the details of Anti-DDoS Pro or Anti-DDoS Premium instances within the Alibaba Cloud account by page. The details include the ID, mitigation plan, expiration time, and forwarding status.
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ExpireEndTime) {
		query["ExpireEndTime"] = request.ExpireEndTime
	}

	if !dara.IsNil(request.ExpireStartTime) {
		query["ExpireStartTime"] = request.ExpireStartTime
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例列表
//
// Description:
//
// You can call the DescribeInstances operation to query the details of Anti-DDoS Pro or Anti-DDoS Premium instances within the Alibaba Cloud account by page. The details include the ID, mitigation plan, expiration time, and forwarding status.
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示全局模板规则
//
// @param request - DescribeL7GlobalRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeL7GlobalRuleResponse
func (client *Client) DescribeL7GlobalRuleWithOptions(request *DescribeL7GlobalRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeL7GlobalRuleResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeL7GlobalRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeL7GlobalRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示全局模板规则
//
// @param request - DescribeL7GlobalRuleRequest
//
// @return DescribeL7GlobalRuleResponse
func (client *Client) DescribeL7GlobalRule(request *DescribeL7GlobalRuleRequest) (_result *DescribeL7GlobalRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeL7GlobalRuleResponse{}
	_body, _err := client.DescribeL7GlobalRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin policies for the forwarding rule of a website.
//
// @param request - DescribeL7RsPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeL7RsPolicyResponse
func (client *Client) DescribeL7RsPolicyWithOptions(request *DescribeL7RsPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeL7RsPolicyResponse, _err error) {
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

	if !dara.IsNil(request.RealServers) {
		query["RealServers"] = request.RealServers
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeL7RsPolicy"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeL7RsPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin policies for the forwarding rule of a website.
//
// @param request - DescribeL7RsPolicyRequest
//
// @return DescribeL7RsPolicyResponse
func (client *Client) DescribeL7RsPolicy(request *DescribeL7RsPolicyRequest) (_result *DescribeL7RsPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeL7RsPolicyResponse{}
	_body, _err := client.DescribeL7RsPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of back-to-origin persistent connections of a domain name.
//
// @param request - DescribeL7UsKeepaliveRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeL7UsKeepaliveResponse
func (client *Client) DescribeL7UsKeepaliveWithOptions(request *DescribeL7UsKeepaliveRequest, runtime *dara.RuntimeOptions) (_result *DescribeL7UsKeepaliveResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeL7UsKeepalive"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeL7UsKeepaliveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of back-to-origin persistent connections of a domain name.
//
// @param request - DescribeL7UsKeepaliveRequest
//
// @return DescribeL7UsKeepaliveResponse
func (client *Client) DescribeL7UsKeepalive(request *DescribeL7UsKeepaliveRequest) (_result *DescribeL7UsKeepaliveResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeL7UsKeepaliveResponse{}
	_body, _err := client.DescribeL7UsKeepaliveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin settings of a port forwarding rule.
//
// @param request - DescribeLayer4RulePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLayer4RulePolicyResponse
func (client *Client) DescribeLayer4RulePolicyWithOptions(request *DescribeLayer4RulePolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeLayer4RulePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Listeners) {
		query["Listeners"] = request.Listeners
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLayer4RulePolicy"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLayer4RulePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin settings of a port forwarding rule.
//
// @param request - DescribeLayer4RulePolicyRequest
//
// @return DescribeLayer4RulePolicyResponse
func (client *Client) DescribeLayer4RulePolicy(request *DescribeLayer4RulePolicyRequest) (_result *DescribeLayer4RulePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLayer4RulePolicyResponse{}
	_body, _err := client.DescribeLayer4RulePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a Logstore is created for Anti-DDoS Proxy.
//
// @param request - DescribeLogStoreExistStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogStoreExistStatusResponse
func (client *Client) DescribeLogStoreExistStatusWithOptions(request *DescribeLogStoreExistStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogStoreExistStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogStoreExistStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogStoreExistStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a Logstore is created for Anti-DDoS Proxy.
//
// @param request - DescribeLogStoreExistStatusRequest
//
// @return DescribeLogStoreExistStatusResponse
func (client *Client) DescribeLogStoreExistStatus(request *DescribeLogStoreExistStatusRequest) (_result *DescribeLogStoreExistStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogStoreExistStatusResponse{}
	_body, _err := client.DescribeLogStoreExistStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the blocked locations that are configured for an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DescribeNetworkRegionBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkRegionBlockResponse
func (client *Client) DescribeNetworkRegionBlockWithOptions(request *DescribeNetworkRegionBlockRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkRegionBlockResponse, _err error) {
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
		Action:      dara.String("DescribeNetworkRegionBlock"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkRegionBlockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the blocked locations that are configured for an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DescribeNetworkRegionBlockRequest
//
// @return DescribeNetworkRegionBlockResponse
func (client *Client) DescribeNetworkRegionBlock(request *DescribeNetworkRegionBlockRequest) (_result *DescribeNetworkRegionBlockResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkRegionBlockResponse{}
	_body, _err := client.DescribeNetworkRegionBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the mitigation settings of the port forwarding rule for a non-website service. The mitigation settings include session persistence and DDoS mitigation policies.
//
// @param request - DescribeNetworkRuleAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkRuleAttributesResponse
func (client *Client) DescribeNetworkRuleAttributesWithOptions(request *DescribeNetworkRuleAttributesRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkRuleAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkRules) {
		query["NetworkRules"] = request.NetworkRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkRuleAttributes"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkRuleAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mitigation settings of the port forwarding rule for a non-website service. The mitigation settings include session persistence and DDoS mitigation policies.
//
// @param request - DescribeNetworkRuleAttributesRequest
//
// @return DescribeNetworkRuleAttributesResponse
func (client *Client) DescribeNetworkRuleAttributes(request *DescribeNetworkRuleAttributesRequest) (_result *DescribeNetworkRuleAttributesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkRuleAttributesResponse{}
	_body, _err := client.DescribeNetworkRuleAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries port forwarding rules.
//
// @param request - DescribeNetworkRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkRulesResponse
func (client *Client) DescribeNetworkRulesWithOptions(request *DescribeNetworkRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForwardProtocol) {
		query["ForwardProtocol"] = request.ForwardProtocol
	}

	if !dara.IsNil(request.FrontendPort) {
		query["FrontendPort"] = request.FrontendPort
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkRules"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries port forwarding rules.
//
// @param request - DescribeNetworkRulesRequest
//
// @return DescribeNetworkRulesResponse
func (client *Client) DescribeNetworkRules(request *DescribeNetworkRulesRequest) (_result *DescribeNetworkRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkRulesResponse{}
	_body, _err := client.DescribeNetworkRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs of Anti-DDoS Proxy (Chinese Mainland).
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// You can query operations performed on Anti-DDoS Proxy (Chinese Mainland), such as configuring the burstable protection bandwidth, deactivating blackhole filtering, configuring the near-origin traffic diversion feature, using Anti-DDoS plans, changing the IP addresses of Elastic Compute Service (ECS) instances, and clearing all logs.
//
// @param request - DescribeOpEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpEntitiesResponse
func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
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

	if !dara.IsNil(request.EntityObject) {
		query["EntityObject"] = request.EntityObject
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOpEntities"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs of Anti-DDoS Proxy (Chinese Mainland).
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// You can query operations performed on Anti-DDoS Proxy (Chinese Mainland), such as configuring the burstable protection bandwidth, deactivating blackhole filtering, configuring the near-origin traffic diversion feature, using Anti-DDoS plans, changing the IP addresses of Elastic Compute Service (ECS) instances, and clearing all logs.
//
// @param request - DescribeOpEntitiesRequest
//
// @return DescribeOpEntitiesResponse
func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the port forwarding rules that are created for an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// Description:
//
// You can call this operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// @param request - DescribePortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortResponse
func (client *Client) DescribePortWithOptions(request *DescribePortRequest, runtime *dara.RuntimeOptions) (_result *DescribePortResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FrontendPort) {
		query["FrontendPort"] = request.FrontendPort
	}

	if !dara.IsNil(request.FrontendProtocol) {
		query["FrontendProtocol"] = request.FrontendProtocol
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePort"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the port forwarding rules that are created for an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// Description:
//
// You can call this operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// @param request - DescribePortRequest
//
// @return DescribePortResponse
func (client *Client) DescribePort(request *DescribePortRequest) (_result *DescribePortResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortResponse{}
	_body, _err := client.DescribePortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the peak attack traffic bandwidth and peak attack traffic packet rates of one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specified period of time.
//
// Description:
//
// You can call this operation to query the peak bandwidth and peak packet rate of attack traffic on one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specific period of time.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePortAttackMaxFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortAttackMaxFlowResponse
func (client *Client) DescribePortAttackMaxFlowWithOptions(request *DescribePortAttackMaxFlowRequest, runtime *dara.RuntimeOptions) (_result *DescribePortAttackMaxFlowResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortAttackMaxFlow"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortAttackMaxFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the peak attack traffic bandwidth and peak attack traffic packet rates of one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specified period of time.
//
// Description:
//
// You can call this operation to query the peak bandwidth and peak packet rate of attack traffic on one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specific period of time.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePortAttackMaxFlowRequest
//
// @return DescribePortAttackMaxFlowResponse
func (client *Client) DescribePortAttackMaxFlow(request *DescribePortAttackMaxFlowRequest) (_result *DescribePortAttackMaxFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortAttackMaxFlowResponse{}
	_body, _err := client.DescribePortAttackMaxFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of the Intelligent Protection policy for non-website services.
//
// @param request - DescribePortAutoCcStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortAutoCcStatusResponse
func (client *Client) DescribePortAutoCcStatusWithOptions(request *DescribePortAutoCcStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePortAutoCcStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortAutoCcStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortAutoCcStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of the Intelligent Protection policy for non-website services.
//
// @param request - DescribePortAutoCcStatusRequest
//
// @return DescribePortAutoCcStatusResponse
func (client *Client) DescribePortAutoCcStatus(request *DescribePortAutoCcStatusRequest) (_result *DescribePortAutoCcStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortAutoCcStatusResponse{}
	_body, _err := client.DescribePortAutoCcStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top source IP addresses of the volumetric attack events for the Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DescribePortCcAttackTopIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortCcAttackTopIPResponse
func (client *Client) DescribePortCcAttackTopIPWithOptions(request *DescribePortCcAttackTopIPRequest, runtime *dara.RuntimeOptions) (_result *DescribePortCcAttackTopIPResponse, _err error) {
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

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortCcAttackTopIP"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortCcAttackTopIPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top source IP addresses of the volumetric attack events for the Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DescribePortCcAttackTopIPRequest
//
// @return DescribePortCcAttackTopIPResponse
func (client *Client) DescribePortCcAttackTopIP(request *DescribePortCcAttackTopIPRequest) (_result *DescribePortCcAttackTopIPResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortCcAttackTopIPResponse{}
	_body, _err := client.DescribePortCcAttackTopIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The statistics on the connections established over the ports of one or more Anti-DDoS Pro or Anti-DDoS Premium instances are queried.
//
// @param request - DescribePortConnsCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortConnsCountResponse
func (client *Client) DescribePortConnsCountWithOptions(request *DescribePortConnsCountRequest, runtime *dara.RuntimeOptions) (_result *DescribePortConnsCountResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortConnsCount"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortConnsCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The statistics on the connections established over the ports of one or more Anti-DDoS Pro or Anti-DDoS Premium instances are queried.
//
// @param request - DescribePortConnsCountRequest
//
// @return DescribePortConnsCountResponse
func (client *Client) DescribePortConnsCount(request *DescribePortConnsCountRequest) (_result *DescribePortConnsCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortConnsCountResponse{}
	_body, _err := client.DescribePortConnsCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the connections established over the ports of one or more Anti-DDoS Proxy instances.
//
// @param request - DescribePortConnsListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortConnsListResponse
func (client *Client) DescribePortConnsListWithOptions(request *DescribePortConnsListRequest, runtime *dara.RuntimeOptions) (_result *DescribePortConnsListResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortConnsList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortConnsListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the connections established over the ports of one or more Anti-DDoS Proxy instances.
//
// @param request - DescribePortConnsListRequest
//
// @return DescribePortConnsListResponse
func (client *Client) DescribePortConnsList(request *DescribePortConnsListRequest) (_result *DescribePortConnsListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortConnsListResponse{}
	_body, _err := client.DescribePortConnsListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the traffic data of one or more Anti-DDoS Pro or Anti-DDoS Premium instances.
//
// @param request - DescribePortFlowListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortFlowListResponse
func (client *Client) DescribePortFlowListWithOptions(request *DescribePortFlowListRequest, runtime *dara.RuntimeOptions) (_result *DescribePortFlowListResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortFlowList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortFlowListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic data of one or more Anti-DDoS Pro or Anti-DDoS Premium instances.
//
// @param request - DescribePortFlowListRequest
//
// @return DescribePortFlowListResponse
func (client *Client) DescribePortFlowList(request *DescribePortFlowListRequest) (_result *DescribePortFlowListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortFlowListResponse{}
	_body, _err := client.DescribePortFlowListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the maximum number of connections that can be established over the ports of one or more Anti-DDoS Proxy instances.
//
// @param request - DescribePortMaxConnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortMaxConnsResponse
func (client *Client) DescribePortMaxConnsWithOptions(request *DescribePortMaxConnsRequest, runtime *dara.RuntimeOptions) (_result *DescribePortMaxConnsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortMaxConns"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortMaxConnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum number of connections that can be established over the ports of one or more Anti-DDoS Proxy instances.
//
// @param request - DescribePortMaxConnsRequest
//
// @return DescribePortMaxConnsResponse
func (client *Client) DescribePortMaxConns(request *DescribePortMaxConnsRequest) (_result *DescribePortMaxConnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortMaxConnsResponse{}
	_body, _err := client.DescribePortMaxConnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the areas and countries from which requests are sent to one or more Anti-DDoS Pro or Anti-DDoS Premium instances within the specified period of time.
//
// @param request - DescribePortViewSourceCountriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortViewSourceCountriesResponse
func (client *Client) DescribePortViewSourceCountriesWithOptions(request *DescribePortViewSourceCountriesRequest, runtime *dara.RuntimeOptions) (_result *DescribePortViewSourceCountriesResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortViewSourceCountries"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortViewSourceCountriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the areas and countries from which requests are sent to one or more Anti-DDoS Pro or Anti-DDoS Premium instances within the specified period of time.
//
// @param request - DescribePortViewSourceCountriesRequest
//
// @return DescribePortViewSourceCountriesResponse
func (client *Client) DescribePortViewSourceCountries(request *DescribePortViewSourceCountriesRequest) (_result *DescribePortViewSourceCountriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortViewSourceCountriesResponse{}
	_body, _err := client.DescribePortViewSourceCountriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Internet service providers (ISPs) from which requests are sent to one or more Anti-DDoS Pro or Anti-DDoS Premium instances within the specified period of time.
//
// Description:
//
// You can call the DescribePortViewSourceIsps operation to query the ISPs from which requests are sent to one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specific period of time.
//
// > The data returned for this operation cannot reflect the actual traffic volume because Layer 4 identity authentication algorithms are updated for Anti-DDoS Pro and Anti-DDoS Premium. You can call this operation to calculate only the proportion of requests sent from different ISPs. If you want to query the request traffic volume, we recommend that you call the [DescribePortFlowList](https://help.aliyun.com/document_detail/157460.html) operation.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePortViewSourceIspsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortViewSourceIspsResponse
func (client *Client) DescribePortViewSourceIspsWithOptions(request *DescribePortViewSourceIspsRequest, runtime *dara.RuntimeOptions) (_result *DescribePortViewSourceIspsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortViewSourceIsps"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortViewSourceIspsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Internet service providers (ISPs) from which requests are sent to one or more Anti-DDoS Pro or Anti-DDoS Premium instances within the specified period of time.
//
// Description:
//
// You can call the DescribePortViewSourceIsps operation to query the ISPs from which requests are sent to one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specific period of time.
//
// > The data returned for this operation cannot reflect the actual traffic volume because Layer 4 identity authentication algorithms are updated for Anti-DDoS Pro and Anti-DDoS Premium. You can call this operation to calculate only the proportion of requests sent from different ISPs. If you want to query the request traffic volume, we recommend that you call the [DescribePortFlowList](https://help.aliyun.com/document_detail/157460.html) operation.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribePortViewSourceIspsRequest
//
// @return DescribePortViewSourceIspsResponse
func (client *Client) DescribePortViewSourceIsps(request *DescribePortViewSourceIspsRequest) (_result *DescribePortViewSourceIspsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortViewSourceIspsResponse{}
	_body, _err := client.DescribePortViewSourceIspsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the administrative regions in China from which requests are sent to one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specified period of time.
//
// @param request - DescribePortViewSourceProvincesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePortViewSourceProvincesResponse
func (client *Client) DescribePortViewSourceProvincesWithOptions(request *DescribePortViewSourceProvincesRequest, runtime *dara.RuntimeOptions) (_result *DescribePortViewSourceProvincesResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePortViewSourceProvinces"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePortViewSourceProvincesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the administrative regions in China from which requests are sent to one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specified period of time.
//
// @param request - DescribePortViewSourceProvincesRequest
//
// @return DescribePortViewSourceProvincesResponse
func (client *Client) DescribePortViewSourceProvinces(request *DescribePortViewSourceProvincesRequest) (_result *DescribePortViewSourceProvincesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePortViewSourceProvincesResponse{}
	_body, _err := client.DescribePortViewSourceProvincesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the protected objects of a scenario-specific custom policy.
//
// Description:
//
// You can call the DescribeSceneDefenseObjects operation to query the protected objects of a scenario-specific custom policy.
//
// Before you call this operation, make sure that you have created a scenario-specific custom policy by calling the [CreateSceneDefensePolicy](https://help.aliyun.com/document_detail/159779.html) operation.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSceneDefenseObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSceneDefenseObjectsResponse
func (client *Client) DescribeSceneDefenseObjectsWithOptions(request *DescribeSceneDefenseObjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSceneDefenseObjectsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSceneDefenseObjects"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSceneDefenseObjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the protected objects of a scenario-specific custom policy.
//
// Description:
//
// You can call the DescribeSceneDefenseObjects operation to query the protected objects of a scenario-specific custom policy.
//
// Before you call this operation, make sure that you have created a scenario-specific custom policy by calling the [CreateSceneDefensePolicy](https://help.aliyun.com/document_detail/159779.html) operation.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSceneDefenseObjectsRequest
//
// @return DescribeSceneDefenseObjectsResponse
func (client *Client) DescribeSceneDefenseObjects(request *DescribeSceneDefenseObjectsRequest) (_result *DescribeSceneDefenseObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSceneDefenseObjectsResponse{}
	_body, _err := client.DescribeSceneDefenseObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a scenario-specific custom policy.
//
// Description:
//
// You can call the DescribeSceneDefensePolicies operation to query the configurations of a scenario-specific custom policy that is created. For example, you can query the status, protected objects, and protection rules of the policy.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSceneDefensePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSceneDefensePoliciesResponse
func (client *Client) DescribeSceneDefensePoliciesWithOptions(request *DescribeSceneDefensePoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSceneDefensePoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Template) {
		query["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSceneDefensePolicies"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSceneDefensePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a scenario-specific custom policy.
//
// Description:
//
// You can call the DescribeSceneDefensePolicies operation to query the configurations of a scenario-specific custom policy that is created. For example, you can query the status, protected objects, and protection rules of the policy.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSceneDefensePoliciesRequest
//
// @return DescribeSceneDefensePoliciesResponse
func (client *Client) DescribeSceneDefensePolicies(request *DescribeSceneDefensePoliciesRequest) (_result *DescribeSceneDefensePoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSceneDefensePoliciesResponse{}
	_body, _err := client.DescribeSceneDefensePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeSchedulerRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSchedulerRulesResponse
func (client *Client) DescribeSchedulerRulesWithOptions(request *DescribeSchedulerRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSchedulerRulesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSchedulerRules"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSchedulerRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeSchedulerRulesRequest
//
// @return DescribeSchedulerRulesResponse
func (client *Client) DescribeSchedulerRules(request *DescribeSchedulerRulesRequest) (_result *DescribeSchedulerRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSchedulerRulesResponse{}
	_body, _err := client.DescribeSchedulerRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the destination rate limit events.
//
// @param request - DescribeSlaEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlaEventListResponse
func (client *Client) DescribeSlaEventListWithOptions(request *DescribeSlaEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlaEventListResponse, _err error) {
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

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlaEventList"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlaEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the destination rate limit events.
//
// @param request - DescribeSlaEventListRequest
//
// @return DescribeSlaEventListResponse
func (client *Client) DescribeSlaEventList(request *DescribeSlaEventListRequest) (_result *DescribeSlaEventListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlaEventListResponse{}
	_body, _err := client.DescribeSlaEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Proxy is authorized to access Simple Log Service.
//
// @param request - DescribeSlsAuthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsAuthStatusResponse
func (client *Client) DescribeSlsAuthStatusWithOptions(request *DescribeSlsAuthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsAuthStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlsAuthStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Proxy is authorized to access Simple Log Service.
//
// @param request - DescribeSlsAuthStatusRequest
//
// @return DescribeSlsAuthStatusResponse
func (client *Client) DescribeSlsAuthStatus(request *DescribeSlsAuthStatusRequest) (_result *DescribeSlsAuthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.DescribeSlsAuthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the Logstore of the Anti-DDoS Proxy instance, such as the log storage capacity and log storage duration.
//
// @param request - DescribeSlsLogstoreInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsLogstoreInfoResponse
func (client *Client) DescribeSlsLogstoreInfoWithOptions(request *DescribeSlsLogstoreInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsLogstoreInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlsLogstoreInfo"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlsLogstoreInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the Logstore of the Anti-DDoS Proxy instance, such as the log storage capacity and log storage duration.
//
// @param request - DescribeSlsLogstoreInfoRequest
//
// @return DescribeSlsLogstoreInfoResponse
func (client *Client) DescribeSlsLogstoreInfo(request *DescribeSlsLogstoreInfoRequest) (_result *DescribeSlsLogstoreInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlsLogstoreInfoResponse{}
	_body, _err := client.DescribeSlsLogstoreInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether Simple Log Service is activated.
//
// @param request - DescribeSlsOpenStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsOpenStatusResponse
func (client *Client) DescribeSlsOpenStatusWithOptions(request *DescribeSlsOpenStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsOpenStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlsOpenStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlsOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether Simple Log Service is activated.
//
// @param request - DescribeSlsOpenStatusRequest
//
// @return DescribeSlsOpenStatusResponse
func (client *Client) DescribeSlsOpenStatus(request *DescribeSlsOpenStatusRequest) (_result *DescribeSlsOpenStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlsOpenStatusResponse{}
	_body, _err := client.DescribeSlsOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Pro or Anti-DDoS Premium is authorized to access other cloud services.
//
// Description:
//
// You can call the DescribeStsGrantStatus operation to query whether Anti-DDoS Pro or Anti-DDoS Premium of the current Alibaba Cloud account is authorized to access other cloud services.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeStsGrantStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStsGrantStatusResponse
func (client *Client) DescribeStsGrantStatusWithOptions(request *DescribeStsGrantStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeStsGrantStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStsGrantStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStsGrantStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Pro or Anti-DDoS Premium is authorized to access other cloud services.
//
// Description:
//
// You can call the DescribeStsGrantStatus operation to query whether Anti-DDoS Pro or Anti-DDoS Premium of the current Alibaba Cloud account is authorized to access other cloud services.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeStsGrantStatusRequest
//
// @return DescribeStsGrantStatusResponse
func (client *Client) DescribeStsGrantStatus(request *DescribeStsGrantStatusRequest) (_result *DescribeStsGrantStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStsGrantStatusResponse{}
	_body, _err := client.DescribeStsGrantStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the bills for the burstable clean bandwidth.
//
// Description:
//
// You can call the DescribeSystemLog operation to query the system logs of Anti-DDoS Pro or Anti-DDoS Premium. The system logs contain only billing logs for the burstable clean bandwidth.
//
// If you have enabled the burstable clean bandwidth feature, you can call this operation to query the details of the bills of the burstable clean bandwidth.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSystemLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemLogResponse
func (client *Client) DescribeSystemLogWithOptions(request *DescribeSystemLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemLogResponse, _err error) {
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

	if !dara.IsNil(request.EntityObject) {
		query["EntityObject"] = request.EntityObject
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
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
		Action:      dara.String("DescribeSystemLog"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the bills for the burstable clean bandwidth.
//
// Description:
//
// You can call the DescribeSystemLog operation to query the system logs of Anti-DDoS Pro or Anti-DDoS Premium. The system logs contain only billing logs for the burstable clean bandwidth.
//
// If you have enabled the burstable clean bandwidth feature, you can call this operation to query the details of the bills of the burstable clean bandwidth.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeSystemLogRequest
//
// @return DescribeSystemLogResponse
func (client *Client) DescribeSystemLog(request *DescribeSystemLogRequest) (_result *DescribeSystemLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSystemLogResponse{}
	_body, _err := client.DescribeSystemLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all tag keys and the number of Anti-DDoS Proxy (Chinese Mainland) instances to which each tag key is added.
//
// Description:
//
// You can call this operation to query all tag keys and the number of Anti-DDoS Proxy (Chinese Mainland) instances to which each tag key is added by page.
//
// >  Only Anti-DDoS Proxy (Chinese Mainland) supports tags.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagKeysResponse
func (client *Client) DescribeTagKeysWithOptions(request *DescribeTagKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagKeysResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagKeys"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all tag keys and the number of Anti-DDoS Proxy (Chinese Mainland) instances to which each tag key is added.
//
// Description:
//
// You can call this operation to query all tag keys and the number of Anti-DDoS Proxy (Chinese Mainland) instances to which each tag key is added by page.
//
// >  Only Anti-DDoS Proxy (Chinese Mainland) supports tags.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeTagKeysRequest
//
// @return DescribeTagKeysResponse
func (client *Client) DescribeTagKeys(request *DescribeTagKeysRequest) (_result *DescribeTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagKeysResponse{}
	_body, _err := client.DescribeTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the tags that are added to an Anti-DDoS Proxy (Chinese Mainland) instance.
//
// Description:
//
// You can call the DescribeTagResources operation to query the information about the tags that are added to an Anti-DDoS Proxy (Chinese Mainland) instance.
//
// >  Only Anti-DDoS Proxy (Chinese Mainland) supports tags.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagResourcesResponse
func (client *Client) DescribeTagResourcesWithOptions(request *DescribeTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagResources"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the tags that are added to an Anti-DDoS Proxy (Chinese Mainland) instance.
//
// Description:
//
// You can call the DescribeTagResources operation to query the information about the tags that are added to an Anti-DDoS Proxy (Chinese Mainland) instance.
//
// >  Only Anti-DDoS Proxy (Chinese Mainland) supports tags.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeTagResourcesRequest
//
// @return DescribeTagResourcesResponse
func (client *Client) DescribeTagResources(request *DescribeTagResourcesRequest) (_result *DescribeTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagResourcesResponse{}
	_body, _err := client.DescribeTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the peak bandwidth and peak packet rates of attack traffic on one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specific period of time.
//
// @param request - DescribeTotalAttackMaxFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTotalAttackMaxFlowResponse
func (client *Client) DescribeTotalAttackMaxFlowWithOptions(request *DescribeTotalAttackMaxFlowRequest, runtime *dara.RuntimeOptions) (_result *DescribeTotalAttackMaxFlowResponse, _err error) {
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

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTotalAttackMaxFlow"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTotalAttackMaxFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the peak bandwidth and peak packet rates of attack traffic on one or more Anti-DDoS Pro or Anti-DDoS Premium instances within a specific period of time.
//
// @param request - DescribeTotalAttackMaxFlowRequest
//
// @return DescribeTotalAttackMaxFlowResponse
func (client *Client) DescribeTotalAttackMaxFlow(request *DescribeTotalAttackMaxFlowRequest) (_result *DescribeTotalAttackMaxFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTotalAttackMaxFlowResponse{}
	_body, _err := client.DescribeTotalAttackMaxFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the source ports of UDP traffic that are filtered out by the filtering policies for UDP reflection attacks on an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DescribeUdpReflectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUdpReflectResponse
func (client *Client) DescribeUdpReflectWithOptions(request *DescribeUdpReflectRequest, runtime *dara.RuntimeOptions) (_result *DescribeUdpReflectResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUdpReflect"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUdpReflectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the source ports of UDP traffic that are filtered out by the filtering policies for UDP reflection attacks on an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - DescribeUdpReflectRequest
//
// @return DescribeUdpReflectResponse
func (client *Client) DescribeUdpReflect(request *DescribeUdpReflectRequest) (_result *DescribeUdpReflectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUdpReflectResponse{}
	_body, _err := client.DescribeUdpReflectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total quota and remaining quota for blackhole filtering deactivation.
//
// @param request - DescribeUnBlackholeCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUnBlackholeCountResponse
func (client *Client) DescribeUnBlackholeCountWithOptions(request *DescribeUnBlackholeCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeUnBlackholeCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUnBlackholeCount"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUnBlackholeCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total quota and remaining quota for blackhole filtering deactivation.
//
// @param request - DescribeUnBlackholeCountRequest
//
// @return DescribeUnBlackholeCountResponse
func (client *Client) DescribeUnBlackholeCount(request *DescribeUnBlackholeCountRequest) (_result *DescribeUnBlackholeCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUnBlackholeCountResponse{}
	_body, _err := client.DescribeUnBlackholeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number and the remaining number of times that you can enable the near-origin traffic diversion feature.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - DescribeUnBlockCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUnBlockCountResponse
func (client *Client) DescribeUnBlockCountWithOptions(request *DescribeUnBlockCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeUnBlockCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUnBlockCount"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUnBlockCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number and the remaining number of times that you can enable the near-origin traffic diversion feature.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - DescribeUnBlockCountRequest
//
// @return DescribeUnBlockCountResponse
func (client *Client) DescribeUnBlockCount(request *DescribeUnBlockCountRequest) (_result *DescribeUnBlockCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUnBlockCountResponse{}
	_body, _err := client.DescribeUnBlockCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the log analysis feature is enabled for all domain names.
//
// Description:
//
// You can call the DescribeWebAccessLogDispatchStatus operation to check whether the log analysis feature is enabled for all domain names that are added to your Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeWebAccessLogDispatchStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebAccessLogDispatchStatusResponse
func (client *Client) DescribeWebAccessLogDispatchStatusWithOptions(request *DescribeWebAccessLogDispatchStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebAccessLogDispatchStatusResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebAccessLogDispatchStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebAccessLogDispatchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the log analysis feature is enabled for all domain names.
//
// Description:
//
// You can call the DescribeWebAccessLogDispatchStatus operation to check whether the log analysis feature is enabled for all domain names that are added to your Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeWebAccessLogDispatchStatusRequest
//
// @return DescribeWebAccessLogDispatchStatusResponse
func (client *Client) DescribeWebAccessLogDispatchStatus(request *DescribeWebAccessLogDispatchStatusRequest) (_result *DescribeWebAccessLogDispatchStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebAccessLogDispatchStatusResponse{}
	_body, _err := client.DescribeWebAccessLogDispatchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the remaining quota that allows you to clear the Logstore.
//
// @param request - DescribeWebAccessLogEmptyCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebAccessLogEmptyCountResponse
func (client *Client) DescribeWebAccessLogEmptyCountWithOptions(request *DescribeWebAccessLogEmptyCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebAccessLogEmptyCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebAccessLogEmptyCount"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebAccessLogEmptyCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the remaining quota that allows you to clear the Logstore.
//
// @param request - DescribeWebAccessLogEmptyCountRequest
//
// @return DescribeWebAccessLogEmptyCountResponse
func (client *Client) DescribeWebAccessLogEmptyCount(request *DescribeWebAccessLogEmptyCountRequest) (_result *DescribeWebAccessLogEmptyCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebAccessLogEmptyCountResponse{}
	_body, _err := client.DescribeWebAccessLogEmptyCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the log analysis feature for a website, such as the feature status and the Simple Log Service project and Logstore that are used.
//
// @param request - DescribeWebAccessLogStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebAccessLogStatusResponse
func (client *Client) DescribeWebAccessLogStatusWithOptions(request *DescribeWebAccessLogStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebAccessLogStatusResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebAccessLogStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebAccessLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the log analysis feature for a website, such as the feature status and the Simple Log Service project and Logstore that are used.
//
// @param request - DescribeWebAccessLogStatusRequest
//
// @return DescribeWebAccessLogStatusResponse
func (client *Client) DescribeWebAccessLogStatus(request *DescribeWebAccessLogStatusRequest) (_result *DescribeWebAccessLogStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebAccessLogStatusResponse{}
	_body, _err := client.DescribeWebAccessLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the mode in which a website service is added to Anti-DDoS Pro or Anti-DDoS Premium.
//
// @param request - DescribeWebAccessModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebAccessModeResponse
func (client *Client) DescribeWebAccessModeWithOptions(request *DescribeWebAccessModeRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebAccessModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebAccessMode"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebAccessModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mode in which a website service is added to Anti-DDoS Pro or Anti-DDoS Premium.
//
// @param request - DescribeWebAccessModeRequest
//
// @return DescribeWebAccessModeResponse
func (client *Client) DescribeWebAccessMode(request *DescribeWebAccessModeRequest) (_result *DescribeWebAccessModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebAccessModeResponse{}
	_body, _err := client.DescribeWebAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Location Blacklist (Domain Names) configurations for websites.
//
// @param request - DescribeWebAreaBlockConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebAreaBlockConfigsResponse
func (client *Client) DescribeWebAreaBlockConfigsWithOptions(request *DescribeWebAreaBlockConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebAreaBlockConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebAreaBlockConfigs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebAreaBlockConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Location Blacklist (Domain Names) configurations for websites.
//
// @param request - DescribeWebAreaBlockConfigsRequest
//
// @return DescribeWebAreaBlockConfigsResponse
func (client *Client) DescribeWebAreaBlockConfigs(request *DescribeWebAreaBlockConfigsRequest) (_result *DescribeWebAreaBlockConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebAreaBlockConfigsResponse{}
	_body, _err := client.DescribeWebAreaBlockConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeWebCCRules is deprecated, please use ddoscoo::2020-01-01::ConfigWebCCRuleV2 instead.
//
// Summary:
//
// Queries the custom frequency control rules that are created for a website.
//
// @param request - DescribeWebCCRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebCCRulesResponse
func (client *Client) DescribeWebCCRulesWithOptions(request *DescribeWebCCRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebCCRulesResponse, _err error) {
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
		Action:      dara.String("DescribeWebCCRules"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebCCRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeWebCCRules is deprecated, please use ddoscoo::2020-01-01::ConfigWebCCRuleV2 instead.
//
// Summary:
//
// Queries the custom frequency control rules that are created for a website.
//
// @param request - DescribeWebCCRulesRequest
//
// @return DescribeWebCCRulesResponse
// Deprecated
func (client *Client) DescribeWebCCRules(request *DescribeWebCCRulesRequest) (_result *DescribeWebCCRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebCCRulesResponse{}
	_body, _err := client.DescribeWebCCRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the custom frequency control rules that are created for a website.
//
// @param request - DescribeWebCCRulesV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebCCRulesV2Response
func (client *Client) DescribeWebCCRulesV2WithOptions(request *DescribeWebCCRulesV2Request, runtime *dara.RuntimeOptions) (_result *DescribeWebCCRulesV2Response, _err error) {
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

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebCCRulesV2"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebCCRulesV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom frequency control rules that are created for a website.
//
// @param request - DescribeWebCCRulesV2Request
//
// @return DescribeWebCCRulesV2Response
func (client *Client) DescribeWebCCRulesV2(request *DescribeWebCCRulesV2Request) (_result *DescribeWebCCRulesV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebCCRulesV2Response{}
	_body, _err := client.DescribeWebCCRulesV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Static Page Caching configuration of websites.
//
// Description:
//
// You can call the DescribeWebCacheConfigs operation to query the Static Page Caching configurations of websites. The configurations include cache modes and custom caching rules.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeWebCacheConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebCacheConfigsResponse
func (client *Client) DescribeWebCacheConfigsWithOptions(request *DescribeWebCacheConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebCacheConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebCacheConfigs"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebCacheConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Static Page Caching configuration of websites.
//
// Description:
//
// You can call the DescribeWebCacheConfigs operation to query the Static Page Caching configurations of websites. The configurations include cache modes and custom caching rules.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeWebCacheConfigsRequest
//
// @return DescribeWebCacheConfigsResponse
func (client *Client) DescribeWebCacheConfigs(request *DescribeWebCacheConfigsRequest) (_result *DescribeWebCacheConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebCacheConfigsResponse{}
	_body, _err := client.DescribeWebCacheConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of each mitigation policy for a website.
//
// @param request - DescribeWebCcProtectSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebCcProtectSwitchResponse
func (client *Client) DescribeWebCcProtectSwitchWithOptions(request *DescribeWebCcProtectSwitchRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebCcProtectSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebCcProtectSwitch"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebCcProtectSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of each mitigation policy for a website.
//
// @param request - DescribeWebCcProtectSwitchRequest
//
// @return DescribeWebCcProtectSwitchResponse
func (client *Client) DescribeWebCcProtectSwitch(request *DescribeWebCcProtectSwitchRequest) (_result *DescribeWebCcProtectSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebCcProtectSwitchResponse{}
	_body, _err := client.DescribeWebCcProtectSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the supported custom ports of a website.
//
// @param request - DescribeWebCustomPortsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebCustomPortsResponse
func (client *Client) DescribeWebCustomPortsWithOptions(request *DescribeWebCustomPortsRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebCustomPortsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebCustomPorts"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebCustomPortsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the supported custom ports of a website.
//
// @param request - DescribeWebCustomPortsRequest
//
// @return DescribeWebCustomPortsResponse
func (client *Client) DescribeWebCustomPorts(request *DescribeWebCustomPortsRequest) (_result *DescribeWebCustomPortsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebCustomPortsResponse{}
	_body, _err := client.DescribeWebCustomPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about Anti-DDoS Pro or Anti-DDoS Premium instances to which a website service is added.
//
// @param request - DescribeWebInstanceRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebInstanceRelationsResponse
func (client *Client) DescribeWebInstanceRelationsWithOptions(request *DescribeWebInstanceRelationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebInstanceRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebInstanceRelations"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebInstanceRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Anti-DDoS Pro or Anti-DDoS Premium instances to which a website service is added.
//
// @param request - DescribeWebInstanceRelationsRequest
//
// @return DescribeWebInstanceRelationsResponse
func (client *Client) DescribeWebInstanceRelations(request *DescribeWebInstanceRelationsRequest) (_result *DescribeWebInstanceRelationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebInstanceRelationsResponse{}
	_body, _err := client.DescribeWebInstanceRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the accurate access control rules that are created for websites.
//
// @param request - DescribeWebPreciseAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebPreciseAccessRuleResponse
func (client *Client) DescribeWebPreciseAccessRuleWithOptions(request *DescribeWebPreciseAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebPreciseAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebPreciseAccessRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebPreciseAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accurate access control rules that are created for websites.
//
// @param request - DescribeWebPreciseAccessRuleRequest
//
// @return DescribeWebPreciseAccessRuleResponse
func (client *Client) DescribeWebPreciseAccessRule(request *DescribeWebPreciseAccessRuleRequest) (_result *DescribeWebPreciseAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebPreciseAccessRuleResponse{}
	_body, _err := client.DescribeWebPreciseAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top source IP addresses of the web resource exhaustion attacks for the Anti-DDoS Proxy instance.
//
// @param request - DescribeWebReportTopIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebReportTopIpResponse
func (client *Client) DescribeWebReportTopIpWithOptions(request *DescribeWebReportTopIpRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebReportTopIpResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Top) {
		query["Top"] = request.Top
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebReportTopIp"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebReportTopIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top source IP addresses of the web resource exhaustion attacks for the Anti-DDoS Proxy instance.
//
// @param request - DescribeWebReportTopIpRequest
//
// @return DescribeWebReportTopIpResponse
func (client *Client) DescribeWebReportTopIp(request *DescribeWebReportTopIpRequest) (_result *DescribeWebReportTopIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebReportTopIpResponse{}
	_body, _err := client.DescribeWebReportTopIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query Configuration of Website Business Forwarding Rules.
//
// Description:
//
// This interface is used for paginated querying of the configurations of website business forwarding rules you have created, such as forwarding protocol types, source server addresses, HTTPS configurations, IP blacklist configurations, and more.
//
// Before calling this interface, you must have already called [CreateWebRule](~~CreateWebRule~~) to create website business forwarding rules.
//
// ### QPS Limit
//
// The per-user QPS limit for this interface is 50 times/second. Exceeding this limit will result in API calls being throttled, which may impact your business; please use it reasonably.
//
// @param request - DescribeWebRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWebRulesResponse
func (client *Client) DescribeWebRulesWithOptions(request *DescribeWebRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeWebRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cname) {
		query["Cname"] = request.Cname
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryDomainPattern) {
		query["QueryDomainPattern"] = request.QueryDomainPattern
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWebRules"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWebRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Configuration of Website Business Forwarding Rules.
//
// Description:
//
// This interface is used for paginated querying of the configurations of website business forwarding rules you have created, such as forwarding protocol types, source server addresses, HTTPS configurations, IP blacklist configurations, and more.
//
// Before calling this interface, you must have already called [CreateWebRule](~~CreateWebRule~~) to create website business forwarding rules.
//
// ### QPS Limit
//
// The per-user QPS limit for this interface is 50 times/second. Exceeding this limit will result in API calls being throttled, which may impact your business; please use it reasonably.
//
// @param request - DescribeWebRulesRequest
//
// @return DescribeWebRulesResponse
func (client *Client) DescribeWebRules(request *DescribeWebRulesRequest) (_result *DescribeWebRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWebRulesResponse{}
	_body, _err := client.DescribeWebRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a protected object from a scenario-specific custom policy.
//
// @param request - DetachSceneDefenseObjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachSceneDefenseObjectResponse
func (client *Client) DetachSceneDefenseObjectWithOptions(request *DetachSceneDefenseObjectRequest, runtime *dara.RuntimeOptions) (_result *DetachSceneDefenseObjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.Objects) {
		query["Objects"] = request.Objects
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachSceneDefenseObject"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachSceneDefenseObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a protected object from a scenario-specific custom policy.
//
// @param request - DetachSceneDefenseObjectRequest
//
// @return DetachSceneDefenseObjectResponse
func (client *Client) DetachSceneDefenseObject(request *DetachSceneDefenseObjectRequest) (_result *DetachSceneDefenseObjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachSceneDefenseObjectResponse{}
	_body, _err := client.DetachSceneDefenseObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a scenario-specific custom policy.
//
// @param request - DisableSceneDefensePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSceneDefensePolicyResponse
func (client *Client) DisableSceneDefensePolicyWithOptions(request *DisableSceneDefensePolicyRequest, runtime *dara.RuntimeOptions) (_result *DisableSceneDefensePolicyResponse, _err error) {
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
		Action:      dara.String("DisableSceneDefensePolicy"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSceneDefensePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a scenario-specific custom policy.
//
// @param request - DisableSceneDefensePolicyRequest
//
// @return DisableSceneDefensePolicyResponse
func (client *Client) DisableSceneDefensePolicy(request *DisableSceneDefensePolicyRequest) (_result *DisableSceneDefensePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableSceneDefensePolicyResponse{}
	_body, _err := client.DisableSceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the log analysis feature for a website.
//
// @param request - DisableWebAccessLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableWebAccessLogConfigResponse
func (client *Client) DisableWebAccessLogConfigWithOptions(request *DisableWebAccessLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DisableWebAccessLogConfigResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableWebAccessLogConfig"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableWebAccessLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the log analysis feature for a website.
//
// @param request - DisableWebAccessLogConfigRequest
//
// @return DisableWebAccessLogConfigResponse
func (client *Client) DisableWebAccessLogConfig(request *DisableWebAccessLogConfigRequest) (_result *DisableWebAccessLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableWebAccessLogConfigResponse{}
	_body, _err := client.DisableWebAccessLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables frequency control for a website.
//
// @param request - DisableWebCCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableWebCCResponse
func (client *Client) DisableWebCCWithOptions(request *DisableWebCCRequest, runtime *dara.RuntimeOptions) (_result *DisableWebCCResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableWebCC"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableWebCCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables frequency control for a website.
//
// @param request - DisableWebCCRequest
//
// @return DisableWebCCResponse
func (client *Client) DisableWebCC(request *DisableWebCCRequest) (_result *DisableWebCCResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableWebCCResponse{}
	_body, _err := client.DisableWebCCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables custom frequency control rules for a website.
//
// @param request - DisableWebCCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableWebCCRuleResponse
func (client *Client) DisableWebCCRuleWithOptions(request *DisableWebCCRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableWebCCRuleResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableWebCCRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableWebCCRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables custom frequency control rules for a website.
//
// @param request - DisableWebCCRuleRequest
//
// @return DisableWebCCRuleResponse
func (client *Client) DisableWebCCRule(request *DisableWebCCRuleRequest) (_result *DisableWebCCRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableWebCCRuleResponse{}
	_body, _err := client.DisableWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears the IP address blacklist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - EmptyAutoCcBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EmptyAutoCcBlacklistResponse
func (client *Client) EmptyAutoCcBlacklistWithOptions(request *EmptyAutoCcBlacklistRequest, runtime *dara.RuntimeOptions) (_result *EmptyAutoCcBlacklistResponse, _err error) {
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
		Action:      dara.String("EmptyAutoCcBlacklist"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EmptyAutoCcBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the IP address blacklist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - EmptyAutoCcBlacklistRequest
//
// @return EmptyAutoCcBlacklistResponse
func (client *Client) EmptyAutoCcBlacklist(request *EmptyAutoCcBlacklistRequest) (_result *EmptyAutoCcBlacklistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EmptyAutoCcBlacklistResponse{}
	_body, _err := client.EmptyAutoCcBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears the IP address whitelist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - EmptyAutoCcWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EmptyAutoCcWhitelistResponse
func (client *Client) EmptyAutoCcWhitelistWithOptions(request *EmptyAutoCcWhitelistRequest, runtime *dara.RuntimeOptions) (_result *EmptyAutoCcWhitelistResponse, _err error) {
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
		Action:      dara.String("EmptyAutoCcWhitelist"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EmptyAutoCcWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the IP address whitelist of an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - EmptyAutoCcWhitelistRequest
//
// @return EmptyAutoCcWhitelistResponse
func (client *Client) EmptyAutoCcWhitelist(request *EmptyAutoCcWhitelistRequest) (_result *EmptyAutoCcWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EmptyAutoCcWhitelistResponse{}
	_body, _err := client.EmptyAutoCcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears the Logstore of Anti-DDoS Pro or Anti-DDoS Premium.
//
// @param request - EmptySlsLogstoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EmptySlsLogstoreResponse
func (client *Client) EmptySlsLogstoreWithOptions(request *EmptySlsLogstoreRequest, runtime *dara.RuntimeOptions) (_result *EmptySlsLogstoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EmptySlsLogstore"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EmptySlsLogstoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the Logstore of Anti-DDoS Pro or Anti-DDoS Premium.
//
// @param request - EmptySlsLogstoreRequest
//
// @return EmptySlsLogstoreResponse
func (client *Client) EmptySlsLogstore(request *EmptySlsLogstoreRequest) (_result *EmptySlsLogstoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EmptySlsLogstoreResponse{}
	_body, _err := client.EmptySlsLogstoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a scenario-specific custom policy.
//
// @param request - EnableSceneDefensePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSceneDefensePolicyResponse
func (client *Client) EnableSceneDefensePolicyWithOptions(request *EnableSceneDefensePolicyRequest, runtime *dara.RuntimeOptions) (_result *EnableSceneDefensePolicyResponse, _err error) {
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
		Action:      dara.String("EnableSceneDefensePolicy"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSceneDefensePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a scenario-specific custom policy.
//
// @param request - EnableSceneDefensePolicyRequest
//
// @return EnableSceneDefensePolicyResponse
func (client *Client) EnableSceneDefensePolicy(request *EnableSceneDefensePolicyRequest) (_result *EnableSceneDefensePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableSceneDefensePolicyResponse{}
	_body, _err := client.EnableSceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the log analysis feature for a website.
//
// @param request - EnableWebAccessLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableWebAccessLogConfigResponse
func (client *Client) EnableWebAccessLogConfigWithOptions(request *EnableWebAccessLogConfigRequest, runtime *dara.RuntimeOptions) (_result *EnableWebAccessLogConfigResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableWebAccessLogConfig"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableWebAccessLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the log analysis feature for a website.
//
// @param request - EnableWebAccessLogConfigRequest
//
// @return EnableWebAccessLogConfigResponse
func (client *Client) EnableWebAccessLogConfig(request *EnableWebAccessLogConfigRequest) (_result *EnableWebAccessLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableWebAccessLogConfigResponse{}
	_body, _err := client.EnableWebAccessLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the Frequency Control policy for a website.
//
// @param request - EnableWebCCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableWebCCResponse
func (client *Client) EnableWebCCWithOptions(request *EnableWebCCRequest, runtime *dara.RuntimeOptions) (_result *EnableWebCCResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableWebCC"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableWebCCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Frequency Control policy for a website.
//
// @param request - EnableWebCCRequest
//
// @return EnableWebCCResponse
func (client *Client) EnableWebCC(request *EnableWebCCRequest) (_result *EnableWebCCResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableWebCCResponse{}
	_body, _err := client.EnableWebCCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables custom frequency control rules for a website.
//
// @param request - EnableWebCCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableWebCCRuleResponse
func (client *Client) EnableWebCCRuleWithOptions(request *EnableWebCCRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableWebCCRuleResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableWebCCRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableWebCCRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables custom frequency control rules for a website.
//
// @param request - EnableWebCCRuleRequest
//
// @return EnableWebCCRuleResponse
func (client *Client) EnableWebCCRule(request *EnableWebCCRuleRequest) (_result *EnableWebCCRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableWebCCRuleResponse{}
	_body, _err := client.EnableWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches between the metering methods of the burstable clean bandwidth feature.
//
// Description:
//
// You can switch between the metering methods of the burstable clean bandwidth feature. The new metering method takes effect from 00:00 on the first day of the next month. You can change the metering method up to three times each calendar month. The most recent metering method that you select takes effect in the next month. You cannot change the metering method on the last day of each calendar month.
//
// @param request - ModifyBizBandWidthModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBizBandWidthModeResponse
func (client *Client) ModifyBizBandWidthModeWithOptions(request *ModifyBizBandWidthModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyBizBandWidthModeResponse, _err error) {
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

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBizBandWidthMode"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBizBandWidthModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches between the metering methods of the burstable clean bandwidth feature.
//
// Description:
//
// You can switch between the metering methods of the burstable clean bandwidth feature. The new metering method takes effect from 00:00 on the first day of the next month. You can change the metering method up to three times each calendar month. The most recent metering method that you select takes effect in the next month. You cannot change the metering method on the last day of each calendar month.
//
// @param request - ModifyBizBandWidthModeRequest
//
// @return ModifyBizBandWidthModeResponse
func (client *Client) ModifyBizBandWidthMode(request *ModifyBizBandWidthModeRequest) (_result *ModifyBizBandWidthModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBizBandWidthModeResponse{}
	_body, _err := client.ModifyBizBandWidthModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deactivates blackhole filtering that is triggered on an instance.
//
// @param request - ModifyBlackholeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBlackholeStatusResponse
func (client *Client) ModifyBlackholeStatusWithOptions(request *ModifyBlackholeStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyBlackholeStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BlackholeStatus) {
		query["BlackholeStatus"] = request.BlackholeStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBlackholeStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBlackholeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deactivates blackhole filtering that is triggered on an instance.
//
// @param request - ModifyBlackholeStatusRequest
//
// @return ModifyBlackholeStatusResponse
func (client *Client) ModifyBlackholeStatus(request *ModifyBlackholeStatusRequest) (_result *ModifyBlackholeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBlackholeStatusResponse{}
	_body, _err := client.ModifyBlackholeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the Diversion from Origin Server configuration of an Anti-DDoS Proxy (Chinese Mainland) instance.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - ModifyBlockStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBlockStatusResponse
func (client *Client) ModifyBlockStatusWithOptions(request *ModifyBlockStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyBlockStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lines) {
		query["Lines"] = request.Lines
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBlockStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBlockStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the Diversion from Origin Server configuration of an Anti-DDoS Proxy (Chinese Mainland) instance.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - ModifyBlockStatusRequest
//
// @return ModifyBlockStatusResponse
func (client *Client) ModifyBlockStatus(request *ModifyBlockStatusRequest) (_result *ModifyBlockStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBlockStatusResponse{}
	_body, _err := client.ModifyBlockStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables CNAME reuse for a website.
//
// Description:
//
// > This operation is suitable only for Anti-DDoS Premium.
//
// @param request - ModifyCnameReuseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCnameReuseResponse
func (client *Client) ModifyCnameReuseWithOptions(request *ModifyCnameReuseRequest, runtime *dara.RuntimeOptions) (_result *ModifyCnameReuseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cname) {
		query["Cname"] = request.Cname
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCnameReuse"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCnameReuseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables CNAME reuse for a website.
//
// Description:
//
// > This operation is suitable only for Anti-DDoS Premium.
//
// @param request - ModifyCnameReuseRequest
//
// @return ModifyCnameReuseResponse
func (client *Client) ModifyCnameReuse(request *ModifyCnameReuseRequest) (_result *ModifyCnameReuseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCnameReuseResponse{}
	_body, _err := client.ModifyCnameReuseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the forwarding rule of a website.
//
// @param request - ModifyDomainResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDomainResourceResponse
func (client *Client) ModifyDomainResourceWithOptions(request *ModifyDomainResourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyDomainResourceResponse, _err error) {
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

	if !dara.IsNil(request.HttpsExt) {
		query["HttpsExt"] = request.HttpsExt
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ProxyTypes) {
		query["ProxyTypes"] = request.ProxyTypes
	}

	if !dara.IsNil(request.RealServers) {
		query["RealServers"] = request.RealServers
	}

	if !dara.IsNil(request.RsType) {
		query["RsType"] = request.RsType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDomainResource"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDomainResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the forwarding rule of a website.
//
// @param request - ModifyDomainResourceRequest
//
// @return ModifyDomainResourceResponse
func (client *Client) ModifyDomainResource(request *ModifyDomainResourceRequest) (_result *ModifyDomainResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDomainResourceResponse{}
	_body, _err := client.ModifyDomainResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the burstable protection bandwidth of a specified Anti-DDoS Proxy (Chinese Mainland) instance.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - ModifyElasticBandWidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElasticBandWidthResponse
func (client *Client) ModifyElasticBandWidthWithOptions(request *ModifyElasticBandWidthRequest, runtime *dara.RuntimeOptions) (_result *ModifyElasticBandWidthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ElasticBandwidth) {
		query["ElasticBandwidth"] = request.ElasticBandwidth
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyElasticBandWidth"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyElasticBandWidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the burstable protection bandwidth of a specified Anti-DDoS Proxy (Chinese Mainland) instance.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - ModifyElasticBandWidthRequest
//
// @return ModifyElasticBandWidthResponse
func (client *Client) ModifyElasticBandWidth(request *ModifyElasticBandWidthRequest) (_result *ModifyElasticBandWidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyElasticBandWidthResponse{}
	_body, _err := client.ModifyElasticBandWidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the burstable clean bandwidth for an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// Description:
//
// Before you call this operation, make sure that you have fully understood the billing method and [pricing](https://help.aliyun.com/document_detail/283754.html) of the burstable clean bandwidth feature. After you call this operation for the first time, the modification immediately takes effect.
//
// @param request - ModifyElasticBizBandWidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElasticBizBandWidthResponse
func (client *Client) ModifyElasticBizBandWidthWithOptions(request *ModifyElasticBizBandWidthRequest, runtime *dara.RuntimeOptions) (_result *ModifyElasticBizBandWidthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ElasticBizBandwidth) {
		query["ElasticBizBandwidth"] = request.ElasticBizBandwidth
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyElasticBizBandWidth"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyElasticBizBandWidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the burstable clean bandwidth for an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// Description:
//
// Before you call this operation, make sure that you have fully understood the billing method and [pricing](https://help.aliyun.com/document_detail/283754.html) of the burstable clean bandwidth feature. After you call this operation for the first time, the modification immediately takes effect.
//
// @param request - ModifyElasticBizBandWidthRequest
//
// @return ModifyElasticBizBandWidthResponse
func (client *Client) ModifyElasticBizBandWidth(request *ModifyElasticBizBandWidthRequest) (_result *ModifyElasticBizBandWidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyElasticBizBandWidthResponse{}
	_body, _err := client.ModifyElasticBizBandWidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the burstable QPS and mode of an Anti-DDoS Proxy instance.
//
// Description:
//
// You can enable burstable QPS only for IPv4 instances.
//
// @param request - ModifyElasticBizQpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElasticBizQpsResponse
func (client *Client) ModifyElasticBizQpsWithOptions(request *ModifyElasticBizQpsRequest, runtime *dara.RuntimeOptions) (_result *ModifyElasticBizQpsResponse, _err error) {
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

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OpsElasticQps) {
		query["OpsElasticQps"] = request.OpsElasticQps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyElasticBizQps"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyElasticBizQpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the burstable QPS and mode of an Anti-DDoS Proxy instance.
//
// Description:
//
// You can enable burstable QPS only for IPv4 instances.
//
// @param request - ModifyElasticBizQpsRequest
//
// @return ModifyElasticBizQpsResponse
func (client *Client) ModifyElasticBizQps(request *ModifyElasticBizQpsRequest) (_result *ModifyElasticBizQpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyElasticBizQpsResponse{}
	_body, _err := client.ModifyElasticBizQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the log storage duration for Anti-DDoS Proxy.
//
// @param request - ModifyFullLogTtlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFullLogTtlResponse
func (client *Client) ModifyFullLogTtlWithOptions(request *ModifyFullLogTtlRequest, runtime *dara.RuntimeOptions) (_result *ModifyFullLogTtlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFullLogTtl"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFullLogTtlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the log storage duration for Anti-DDoS Proxy.
//
// @param request - ModifyFullLogTtlRequest
//
// @return ModifyFullLogTtlResponse
func (client *Client) ModifyFullLogTtl(request *ModifyFullLogTtlRequest) (_result *ModifyFullLogTtlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFullLogTtlResponse{}
	_body, _err := client.ModifyFullLogTtlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the custom header of a domain name that is added to an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - ModifyHeadersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHeadersResponse
func (client *Client) ModifyHeadersWithOptions(request *ModifyHeadersRequest, runtime *dara.RuntimeOptions) (_result *ModifyHeadersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomHeaders) {
		query["CustomHeaders"] = request.CustomHeaders
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHeaders"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHeadersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the custom header of a domain name that is added to an Anti-DDoS Pro or Anti-DDoS Premium instance.
//
// @param request - ModifyHeadersRequest
//
// @return ModifyHeadersResponse
func (client *Client) ModifyHeaders(request *ModifyHeadersRequest) (_result *ModifyHeadersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHeadersResponse{}
	_body, _err := client.ModifyHeadersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the Layer 4 or Layer 7 health check configuration of a port forwarding rule.
//
// @param request - ModifyHealthCheckConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHealthCheckConfigResponse
func (client *Client) ModifyHealthCheckConfigWithOptions(request *ModifyHealthCheckConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyHealthCheckConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForwardProtocol) {
		query["ForwardProtocol"] = request.ForwardProtocol
	}

	if !dara.IsNil(request.FrontendPort) {
		query["FrontendPort"] = request.FrontendPort
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHealthCheckConfig"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHealthCheckConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the Layer 4 or Layer 7 health check configuration of a port forwarding rule.
//
// @param request - ModifyHealthCheckConfigRequest
//
// @return ModifyHealthCheckConfigResponse
func (client *Client) ModifyHealthCheckConfig(request *ModifyHealthCheckConfigRequest) (_result *ModifyHealthCheckConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHealthCheckConfigResponse{}
	_body, _err := client.ModifyHealthCheckConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables HTTP/2 for the forwarding rule of a website.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - ModifyHttp2EnableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHttp2EnableResponse
func (client *Client) ModifyHttp2EnableWithOptions(request *ModifyHttp2EnableRequest, runtime *dara.RuntimeOptions) (_result *ModifyHttp2EnableResponse, _err error) {
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

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHttp2Enable"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHttp2EnableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables HTTP/2 for the forwarding rule of a website.
//
// Description:
//
// >  This operation is suitable only for Anti-DDoS Proxy (Chinese Mainland).
//
// @param request - ModifyHttp2EnableRequest
//
// @return ModifyHttp2EnableResponse
func (client *Client) ModifyHttp2Enable(request *ModifyHttp2EnableRequest) (_result *ModifyHttp2EnableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHttp2EnableResponse{}
	_body, _err := client.ModifyHttp2EnableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Instance adjustment, similar to BSS adjustment
//
// @param request - ModifyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstanceWithOptions(request *ModifyInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.BaseBandwidth) {
		query["BaseBandwidth"] = request.BaseBandwidth
	}

	if !dara.IsNil(request.DomainCount) {
		query["DomainCount"] = request.DomainCount
	}

	if !dara.IsNil(request.EditionSale) {
		query["EditionSale"] = request.EditionSale
	}

	if !dara.IsNil(request.FunctionVersion) {
		query["FunctionVersion"] = request.FunctionVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ModifyType) {
		query["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.NormalBandwidth) {
		query["NormalBandwidth"] = request.NormalBandwidth
	}

	if !dara.IsNil(request.NormalQps) {
		query["NormalQps"] = request.NormalQps
	}

	if !dara.IsNil(request.PortCount) {
		query["PortCount"] = request.PortCount
	}

	if !dara.IsNil(request.ProductPlan) {
		query["ProductPlan"] = request.ProductPlan
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ServiceBandwidth) {
		query["ServiceBandwidth"] = request.ServiceBandwidth
	}

	if !dara.IsNil(request.ServicePartner) {
		query["ServicePartner"] = request.ServicePartner
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstance"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Instance adjustment, similar to BSS adjustment
//
// @param request - ModifyInstanceRequest
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstance(request *ModifyInstanceRequest) (_result *ModifyInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.ModifyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an Anti-DDoS Proxy instance.
//
// @param request - ModifyInstanceRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceRemarkResponse
func (client *Client) ModifyInstanceRemarkWithOptions(request *ModifyInstanceRemarkRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceRemarkResponse, _err error) {
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

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceRemark"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an Anti-DDoS Proxy instance.
//
// @param request - ModifyInstanceRemarkRequest
//
// @return ModifyInstanceRemarkResponse
func (client *Client) ModifyInstanceRemark(request *ModifyInstanceRemarkRequest) (_result *ModifyInstanceRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceRemarkResponse{}
	_body, _err := client.ModifyInstanceRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the session persistence and DDoS mitigation policy settings of a port forwarding rule.
//
// @param request - ModifyNetworkRuleAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNetworkRuleAttributeResponse
func (client *Client) ModifyNetworkRuleAttributeWithOptions(request *ModifyNetworkRuleAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyNetworkRuleAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.ForwardProtocol) {
		query["ForwardProtocol"] = request.ForwardProtocol
	}

	if !dara.IsNil(request.FrontendPort) {
		query["FrontendPort"] = request.FrontendPort
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNetworkRuleAttribute"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNetworkRuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the session persistence and DDoS mitigation policy settings of a port forwarding rule.
//
// @param request - ModifyNetworkRuleAttributeRequest
//
// @return ModifyNetworkRuleAttributeResponse
func (client *Client) ModifyNetworkRuleAttribute(request *ModifyNetworkRuleAttributeRequest) (_result *ModifyNetworkRuleAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyNetworkRuleAttributeResponse{}
	_body, _err := client.ModifyNetworkRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies whether to enable the Online Certificate Status Protocol (OCSP) feature.
//
// Description:
//
// This feature is available only for a website that supports HTTPS. If HTTPS is selected for Protocol, we recommend that you enable this feature.
//
// @param request - ModifyOcspStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOcspStatusResponse
func (client *Client) ModifyOcspStatusWithOptions(request *ModifyOcspStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyOcspStatusResponse, _err error) {
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

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOcspStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOcspStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies whether to enable the Online Certificate Status Protocol (OCSP) feature.
//
// Description:
//
// This feature is available only for a website that supports HTTPS. If HTTPS is selected for Protocol, we recommend that you enable this feature.
//
// @param request - ModifyOcspStatusRequest
//
// @return ModifyOcspStatusResponse
func (client *Client) ModifyOcspStatus(request *ModifyOcspStatusRequest) (_result *ModifyOcspStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyOcspStatusResponse{}
	_body, _err := client.ModifyOcspStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a port forwarding rule.
//
// Description:
//
// You can call the ModifyPort operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// @param request - ModifyPortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPortResponse
func (client *Client) ModifyPortWithOptions(request *ModifyPortRequest, runtime *dara.RuntimeOptions) (_result *ModifyPortResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendPort) {
		query["BackendPort"] = request.BackendPort
	}

	if !dara.IsNil(request.FrontendPort) {
		query["FrontendPort"] = request.FrontendPort
	}

	if !dara.IsNil(request.FrontendProtocol) {
		query["FrontendProtocol"] = request.FrontendProtocol
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProxyEnable) {
		query["ProxyEnable"] = request.ProxyEnable
	}

	if !dara.IsNil(request.RealServers) {
		query["RealServers"] = request.RealServers
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPort"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a port forwarding rule.
//
// Description:
//
// You can call the ModifyPort operation by using Terraform. For more information about Terraform, see [What is Terraform?](https://help.aliyun.com/document_detail/95820.html).
//
// @param request - ModifyPortRequest
//
// @return ModifyPortResponse
func (client *Client) ModifyPort(request *ModifyPortRequest) (_result *ModifyPortResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPortResponse{}
	_body, _err := client.ModifyPortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the Intelligent Protection configuration of a non-website service.
//
// @param request - ModifyPortAutoCcStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPortAutoCcStatusResponse
func (client *Client) ModifyPortAutoCcStatusWithOptions(request *ModifyPortAutoCcStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyPortAutoCcStatusResponse, _err error) {
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

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Switch) {
		query["Switch"] = request.Switch
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPortAutoCcStatus"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPortAutoCcStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the Intelligent Protection configuration of a non-website service.
//
// @param request - ModifyPortAutoCcStatusRequest
//
// @return ModifyPortAutoCcStatusResponse
func (client *Client) ModifyPortAutoCcStatus(request *ModifyPortAutoCcStatusRequest) (_result *ModifyPortAutoCcStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPortAutoCcStatusResponse{}
	_body, _err := client.ModifyPortAutoCcStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches between the metering methods of the burstable clean bandwidth feature.
//
// @param request - ModifyQpsModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyQpsModeResponse
func (client *Client) ModifyQpsModeWithOptions(request *ModifyQpsModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyQpsModeResponse, _err error) {
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

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyQpsMode"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyQpsModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches between the metering methods of the burstable clean bandwidth feature.
//
// @param request - ModifyQpsModeRequest
//
// @return ModifyQpsModeResponse
func (client *Client) ModifyQpsMode(request *ModifyQpsModeRequest) (_result *ModifyQpsModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyQpsModeResponse{}
	_body, _err := client.ModifyQpsModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a scenario-specific custom policy.
//
// @param request - ModifySceneDefensePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySceneDefensePolicyResponse
func (client *Client) ModifySceneDefensePolicyWithOptions(request *ModifySceneDefensePolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifySceneDefensePolicyResponse, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
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
		Action:      dara.String("ModifySceneDefensePolicy"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySceneDefensePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a scenario-specific custom policy.
//
// @param request - ModifySceneDefensePolicyRequest
//
// @return ModifySceneDefensePolicyResponse
func (client *Client) ModifySceneDefensePolicy(request *ModifySceneDefensePolicyRequest) (_result *ModifySceneDefensePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySceneDefensePolicyResponse{}
	_body, _err := client.ModifySceneDefensePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the scheduling rule of Sec-Traffic Manager.
//
// @param request - ModifySchedulerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySchedulerRuleResponse
func (client *Client) ModifySchedulerRuleWithOptions(request *ModifySchedulerRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifySchedulerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySchedulerRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySchedulerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the scheduling rule of Sec-Traffic Manager.
//
// @param request - ModifySchedulerRuleRequest
//
// @return ModifySchedulerRuleResponse
func (client *Client) ModifySchedulerRule(request *ModifySchedulerRuleRequest) (_result *ModifySchedulerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySchedulerRuleResponse{}
	_body, _err := client.ModifySchedulerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the Transport Layer Security (TLS) policy configuration for the forwarding rule of a website.
//
// @param request - ModifyTlsConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTlsConfigResponse
func (client *Client) ModifyTlsConfigWithOptions(request *ModifyTlsConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyTlsConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTlsConfig"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTlsConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the Transport Layer Security (TLS) policy configuration for the forwarding rule of a website.
//
// @param request - ModifyTlsConfigRequest
//
// @return ModifyTlsConfigResponse
func (client *Client) ModifyTlsConfig(request *ModifyTlsConfigRequest) (_result *ModifyTlsConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTlsConfigResponse{}
	_body, _err := client.ModifyTlsConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the mode of the intelligent protection feature for a website.
//
// @param request - ModifyWebAIProtectModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebAIProtectModeResponse
func (client *Client) ModifyWebAIProtectModeWithOptions(request *ModifyWebAIProtectModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebAIProtectModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebAIProtectMode"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebAIProtectModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the mode of the intelligent protection feature for a website.
//
// @param request - ModifyWebAIProtectModeRequest
//
// @return ModifyWebAIProtectModeResponse
func (client *Client) ModifyWebAIProtectMode(request *ModifyWebAIProtectModeRequest) (_result *ModifyWebAIProtectModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebAIProtectModeResponse{}
	_body, _err := client.ModifyWebAIProtectModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the Intelligent Protection policy for a website.
//
// @param request - ModifyWebAIProtectSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebAIProtectSwitchResponse
func (client *Client) ModifyWebAIProtectSwitchWithOptions(request *ModifyWebAIProtectSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebAIProtectSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebAIProtectSwitch"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebAIProtectSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the Intelligent Protection policy for a website.
//
// @param request - ModifyWebAIProtectSwitchRequest
//
// @return ModifyWebAIProtectSwitchResponse
func (client *Client) ModifyWebAIProtectSwitch(request *ModifyWebAIProtectSwitchRequest) (_result *ModifyWebAIProtectSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebAIProtectSwitchResponse{}
	_body, _err := client.ModifyWebAIProtectSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the mode in which a website service is added to Anti-DDoS Pro or Anti-DDoS Premium.
//
// @param request - ModifyWebAccessModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebAccessModeResponse
func (client *Client) ModifyWebAccessModeWithOptions(request *ModifyWebAccessModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebAccessModeResponse, _err error) {
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

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebAccessMode"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebAccessModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the mode in which a website service is added to Anti-DDoS Pro or Anti-DDoS Premium.
//
// @param request - ModifyWebAccessModeRequest
//
// @return ModifyWebAccessModeResponse
func (client *Client) ModifyWebAccessMode(request *ModifyWebAccessModeRequest) (_result *ModifyWebAccessModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebAccessModeResponse{}
	_body, _err := client.ModifyWebAccessModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the blocked locations that are configured in the Location Blacklist (Domain Names) policy for a website.
//
// @param request - ModifyWebAreaBlockRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebAreaBlockResponse
func (client *Client) ModifyWebAreaBlockWithOptions(request *ModifyWebAreaBlockRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebAreaBlockResponse, _err error) {
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

	if !dara.IsNil(request.Regions) {
		query["Regions"] = request.Regions
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebAreaBlock"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebAreaBlockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the blocked locations that are configured in the Location Blacklist (Domain Names) policy for a website.
//
// @param request - ModifyWebAreaBlockRequest
//
// @return ModifyWebAreaBlockResponse
func (client *Client) ModifyWebAreaBlock(request *ModifyWebAreaBlockRequest) (_result *ModifyWebAreaBlockResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebAreaBlockResponse{}
	_body, _err := client.ModifyWebAreaBlockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the Location Blacklist (Domain Names) policy for a domain name.
//
// Description:
//
// You can call the ModifyWebAreaBlockSwitch operation to enable or disable the Location Blacklist (Domain Names) policy for a domain name.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyWebAreaBlockSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebAreaBlockSwitchResponse
func (client *Client) ModifyWebAreaBlockSwitchWithOptions(request *ModifyWebAreaBlockSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebAreaBlockSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebAreaBlockSwitch"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebAreaBlockSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the Location Blacklist (Domain Names) policy for a domain name.
//
// Description:
//
// You can call the ModifyWebAreaBlockSwitch operation to enable or disable the Location Blacklist (Domain Names) policy for a domain name.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyWebAreaBlockSwitchRequest
//
// @return ModifyWebAreaBlockSwitchResponse
func (client *Client) ModifyWebAreaBlockSwitch(request *ModifyWebAreaBlockSwitchRequest) (_result *ModifyWebAreaBlockSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebAreaBlockSwitchResponse{}
	_body, _err := client.ModifyWebAreaBlockSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the HTTP flood mitigation feature for a website.
//
// @param request - ModifyWebCCGlobalSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebCCGlobalSwitchResponse
func (client *Client) ModifyWebCCGlobalSwitchWithOptions(request *ModifyWebCCGlobalSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebCCGlobalSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcGlobalSwitch) {
		query["CcGlobalSwitch"] = request.CcGlobalSwitch
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebCCGlobalSwitch"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebCCGlobalSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the HTTP flood mitigation feature for a website.
//
// @param request - ModifyWebCCGlobalSwitchRequest
//
// @return ModifyWebCCGlobalSwitchResponse
func (client *Client) ModifyWebCCGlobalSwitch(request *ModifyWebCCGlobalSwitchRequest) (_result *ModifyWebCCGlobalSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebCCGlobalSwitchResponse{}
	_body, _err := client.ModifyWebCCGlobalSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifyWebCCRule is deprecated, please use ddoscoo::2020-01-01::ConfigWebCCRuleV2 instead.
//
// Summary:
//
// Modifies the custom frequency control rule of a website.
//
// @param request - ModifyWebCCRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebCCRuleResponse
func (client *Client) ModifyWebCCRuleWithOptions(request *ModifyWebCCRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebCCRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Act) {
		query["Act"] = request.Act
	}

	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Ttl) {
		query["Ttl"] = request.Ttl
	}

	if !dara.IsNil(request.Uri) {
		query["Uri"] = request.Uri
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebCCRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebCCRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyWebCCRule is deprecated, please use ddoscoo::2020-01-01::ConfigWebCCRuleV2 instead.
//
// Summary:
//
// Modifies the custom frequency control rule of a website.
//
// @param request - ModifyWebCCRuleRequest
//
// @return ModifyWebCCRuleResponse
// Deprecated
func (client *Client) ModifyWebCCRule(request *ModifyWebCCRuleRequest) (_result *ModifyWebCCRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebCCRuleResponse{}
	_body, _err := client.ModifyWebCCRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the custom rule of the Static Page Caching policy for a website.
//
// @param request - ModifyWebCacheCustomRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebCacheCustomRuleResponse
func (client *Client) ModifyWebCacheCustomRuleWithOptions(request *ModifyWebCacheCustomRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebCacheCustomRuleResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebCacheCustomRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebCacheCustomRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the custom rule of the Static Page Caching policy for a website.
//
// @param request - ModifyWebCacheCustomRuleRequest
//
// @return ModifyWebCacheCustomRuleResponse
func (client *Client) ModifyWebCacheCustomRule(request *ModifyWebCacheCustomRuleRequest) (_result *ModifyWebCacheCustomRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebCacheCustomRuleResponse{}
	_body, _err := client.ModifyWebCacheCustomRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the cache mode of the Static Page Caching policy for a website.
//
// @param request - ModifyWebCacheModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebCacheModeResponse
func (client *Client) ModifyWebCacheModeWithOptions(request *ModifyWebCacheModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebCacheModeResponse, _err error) {
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

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebCacheMode"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebCacheModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the cache mode of the Static Page Caching policy for a website.
//
// @param request - ModifyWebCacheModeRequest
//
// @return ModifyWebCacheModeResponse
func (client *Client) ModifyWebCacheMode(request *ModifyWebCacheModeRequest) (_result *ModifyWebCacheModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebCacheModeResponse{}
	_body, _err := client.ModifyWebCacheModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the Static Page Caching policy for a website.
//
// Description:
//
// You can call the ModifyWebCacheSwitch operation to enable or disable the Static Page Caching policy for a website.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyWebCacheSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebCacheSwitchResponse
func (client *Client) ModifyWebCacheSwitchWithOptions(request *ModifyWebCacheSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebCacheSwitchResponse, _err error) {
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

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebCacheSwitch"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebCacheSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the Static Page Caching policy for a website.
//
// Description:
//
// You can call the ModifyWebCacheSwitch operation to enable or disable the Static Page Caching policy for a website.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - ModifyWebCacheSwitchRequest
//
// @return ModifyWebCacheSwitchResponse
func (client *Client) ModifyWebCacheSwitch(request *ModifyWebCacheSwitchRequest) (_result *ModifyWebCacheSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebCacheSwitchResponse{}
	_body, _err := client.ModifyWebCacheSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the Black Lists and White Lists (Domain Names) policy for a domain name.
//
// @param request - ModifyWebIpSetSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebIpSetSwitchResponse
func (client *Client) ModifyWebIpSetSwitchWithOptions(request *ModifyWebIpSetSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebIpSetSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebIpSetSwitch"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebIpSetSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the Black Lists and White Lists (Domain Names) policy for a domain name.
//
// @param request - ModifyWebIpSetSwitchRequest
//
// @return ModifyWebIpSetSwitchResponse
func (client *Client) ModifyWebIpSetSwitch(request *ModifyWebIpSetSwitchRequest) (_result *ModifyWebIpSetSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebIpSetSwitchResponse{}
	_body, _err := client.ModifyWebIpSetSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies an accurate access control rule of a website.
//
// @param request - ModifyWebPreciseAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebPreciseAccessRuleResponse
func (client *Client) ModifyWebPreciseAccessRuleWithOptions(request *ModifyWebPreciseAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebPreciseAccessRuleResponse, _err error) {
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

	if !dara.IsNil(request.Expires) {
		query["Expires"] = request.Expires
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebPreciseAccessRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebPreciseAccessRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an accurate access control rule of a website.
//
// @param request - ModifyWebPreciseAccessRuleRequest
//
// @return ModifyWebPreciseAccessRuleResponse
func (client *Client) ModifyWebPreciseAccessRule(request *ModifyWebPreciseAccessRuleRequest) (_result *ModifyWebPreciseAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebPreciseAccessRuleResponse{}
	_body, _err := client.ModifyWebPreciseAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables accurate access control for a website.
//
// @param request - ModifyWebPreciseAccessSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebPreciseAccessSwitchResponse
func (client *Client) ModifyWebPreciseAccessSwitchWithOptions(request *ModifyWebPreciseAccessSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebPreciseAccessSwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebPreciseAccessSwitch"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebPreciseAccessSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables accurate access control for a website.
//
// @param request - ModifyWebPreciseAccessSwitchRequest
//
// @return ModifyWebPreciseAccessSwitchResponse
func (client *Client) ModifyWebPreciseAccessSwitch(request *ModifyWebPreciseAccessSwitchRequest) (_result *ModifyWebPreciseAccessSwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebPreciseAccessSwitchResponse{}
	_body, _err := client.ModifyWebPreciseAccessSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the forwarding rule of a website.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=ddoscoo\\&api=ModifyWebRule\\&type=RPC\\&version=2020-01-01)
//
// @param request - ModifyWebRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWebRuleResponse
func (client *Client) ModifyWebRuleWithOptions(request *ModifyWebRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyWebRuleResponse, _err error) {
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

	if !dara.IsNil(request.HttpsExt) {
		query["HttpsExt"] = request.HttpsExt
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ProxyTypes) {
		query["ProxyTypes"] = request.ProxyTypes
	}

	if !dara.IsNil(request.RealServers) {
		query["RealServers"] = request.RealServers
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RsType) {
		query["RsType"] = request.RsType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWebRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWebRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the forwarding rule of a website.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=ddoscoo\\&api=ModifyWebRule\\&type=RPC\\&version=2020-01-01)
//
// @param request - ModifyWebRuleRequest
//
// @return ModifyWebRuleResponse
func (client *Client) ModifyWebRule(request *ModifyWebRuleRequest) (_result *ModifyWebRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyWebRuleResponse{}
	_body, _err := client.ModifyWebRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the instance that you want to release.
//
// > You can release only expired instances. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/91478.html) operation to query the IDs and expiration status of all instances.
//
// Description:
//
// The ID of the request, which is used to locate and troubleshoot issues.
//
// @param request - ReleaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
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
		Action:      dara.String("ReleaseInstance"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the instance that you want to release.
//
// > You can release only expired instances. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/91478.html) operation to query the IDs and expiration status of all instances.
//
// Description:
//
// The ID of the request, which is used to locate and troubleshoot issues.
//
// @param request - ReleaseInstanceRequest
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches service traffic to an Anti-DDoS Pro or Anti-DDoS Premium instance for scrubbing or switches service traffic back to the associated cloud resources.
//
// Description:
//
// You can call the SwitchSchedulerRule operation to modify the resources to which service traffic is switched for a scheduling rule. For example, you can switch service traffic to an Anti-DDoS Pro or Anti-DDoS Premium instance for scrubbing or switch the service traffic back to the associated cloud resources.
//
// Before you call this operation, you must have created a scheduling rule by calling the [CreateSchedulerRule](https://help.aliyun.com/document_detail/157479.html) operation.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SwitchSchedulerRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchSchedulerRuleResponse
func (client *Client) SwitchSchedulerRuleWithOptions(request *SwitchSchedulerRuleRequest, runtime *dara.RuntimeOptions) (_result *SwitchSchedulerRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.SwitchData) {
		query["SwitchData"] = request.SwitchData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchSchedulerRule"),
		Version:     dara.String("2020-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchSchedulerRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches service traffic to an Anti-DDoS Pro or Anti-DDoS Premium instance for scrubbing or switches service traffic back to the associated cloud resources.
//
// Description:
//
// You can call the SwitchSchedulerRule operation to modify the resources to which service traffic is switched for a scheduling rule. For example, you can switch service traffic to an Anti-DDoS Pro or Anti-DDoS Premium instance for scrubbing or switch the service traffic back to the associated cloud resources.
//
// Before you call this operation, you must have created a scheduling rule by calling the [CreateSchedulerRule](https://help.aliyun.com/document_detail/157479.html) operation.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - SwitchSchedulerRuleRequest
//
// @return SwitchSchedulerRuleResponse
func (client *Client) SwitchSchedulerRule(request *SwitchSchedulerRuleRequest) (_result *SwitchSchedulerRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchSchedulerRuleResponse{}
	_body, _err := client.SwitchSchedulerRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
