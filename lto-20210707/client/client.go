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
	client.Endpoint, _err = client.GetEndpoint(dara.String("lto"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AddBaaSAntChainBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBaaSAntChainBizChainResponse
func (client *Client) AddBaaSAntChainBizChainWithOptions(request *AddBaaSAntChainBizChainRequest, runtime *dara.RuntimeOptions) (_result *AddBaaSAntChainBizChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaaSAntChainChainId) {
		query["BaaSAntChainChainId"] = request.BaaSAntChainChainId
	}

	if !dara.IsNil(request.BaaSAntChainConsortiumId) {
		query["BaaSAntChainConsortiumId"] = request.BaaSAntChainConsortiumId
	}

	if !dara.IsNil(request.CaCert) {
		query["CaCert"] = request.CaCert
	}

	if !dara.IsNil(request.CaCertPassword) {
		query["CaCertPassword"] = request.CaCertPassword
	}

	if !dara.IsNil(request.ClientCert) {
		query["ClientCert"] = request.ClientCert
	}

	if !dara.IsNil(request.ClientKey) {
		query["ClientKey"] = request.ClientKey
	}

	if !dara.IsNil(request.ClientKeyPassword) {
		query["ClientKeyPassword"] = request.ClientKeyPassword
	}

	if !dara.IsNil(request.ContractTemplateIdList) {
		query["ContractTemplateIdList"] = request.ContractTemplateIdList
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeNameList) {
		query["NodeNameList"] = request.NodeNameList
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.UserKey) {
		query["UserKey"] = request.UserKey
	}

	if !dara.IsNil(request.UserKeyPassword) {
		query["UserKeyPassword"] = request.UserKeyPassword
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBaaSAntChainBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBaaSAntChainBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddBaaSAntChainBizChainRequest
//
// @return AddBaaSAntChainBizChainResponse
func (client *Client) AddBaaSAntChainBizChain(request *AddBaaSAntChainBizChainRequest) (_result *AddBaaSAntChainBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddBaaSAntChainBizChainResponse{}
	_body, _err := client.AddBaaSAntChainBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddBaaSFabricBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBaaSFabricBizChainResponse
func (client *Client) AddBaaSFabricBizChainWithOptions(request *AddBaaSFabricBizChainRequest, runtime *dara.RuntimeOptions) (_result *AddBaaSFabricBizChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaaSFabricChannelId) {
		query["BaaSFabricChannelId"] = request.BaaSFabricChannelId
	}

	if !dara.IsNil(request.BaaSFabricConsortiumId) {
		query["BaaSFabricConsortiumId"] = request.BaaSFabricConsortiumId
	}

	if !dara.IsNil(request.BaaSFabricOrganizationId) {
		query["BaaSFabricOrganizationId"] = request.BaaSFabricOrganizationId
	}

	if !dara.IsNil(request.ContractTemplateIdList) {
		query["ContractTemplateIdList"] = request.ContractTemplateIdList
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBaaSFabricBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBaaSFabricBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddBaaSFabricBizChainRequest
//
// @return AddBaaSFabricBizChainResponse
func (client *Client) AddBaaSFabricBizChain(request *AddBaaSFabricBizChainRequest) (_result *AddBaaSFabricBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddBaaSFabricBizChainResponse{}
	_body, _err := client.AddBaaSFabricBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddBsnFabricBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBsnFabricBizChainResponse
func (client *Client) AddBsnFabricBizChainWithOptions(request *AddBsnFabricBizChainRequest, runtime *dara.RuntimeOptions) (_result *AddBsnFabricBizChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeList) {
		query["NodeList"] = request.NodeList
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.UserCode) {
		query["UserCode"] = request.UserCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBsnFabricBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBsnFabricBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddBsnFabricBizChainRequest
//
// @return AddBsnFabricBizChainResponse
func (client *Client) AddBsnFabricBizChain(request *AddBsnFabricBizChainRequest) (_result *AddBsnFabricBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddBsnFabricBizChainResponse{}
	_body, _err := client.AddBsnFabricBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加设备组
//
// @param request - AddDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDeviceGroupResponse
func (client *Client) AddDeviceGroupWithOptions(request *AddDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *AddDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizedCount) {
		query["AuthorizedCount"] = request.AuthorizedCount
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDeviceGroup"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加设备组
//
// @param request - AddDeviceGroupRequest
//
// @return AddDeviceGroupResponse
func (client *Client) AddDeviceGroup(request *AddDeviceGroupRequest) (_result *AddDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDeviceGroupResponse{}
	_body, _err := client.AddDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加成员
//
// @param request - AddMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMemberResponse
func (client *Client) AddMemberWithOptions(request *AddMemberRequest, runtime *dara.RuntimeOptions) (_result *AddMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizedCount) {
		query["AuthorizedCount"] = request.AuthorizedCount
	}

	if !dara.IsNil(request.AuthorizedDeviceCount) {
		query["AuthorizedDeviceCount"] = request.AuthorizedDeviceCount
	}

	if !dara.IsNil(request.Contactor) {
		query["Contactor"] = request.Contactor
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Telephony) {
		query["Telephony"] = request.Telephony
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMember"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加成员
//
// @param request - AddMemberRequest
//
// @return AddMemberResponse
func (client *Client) AddMember(request *AddMemberRequest) (_result *AddMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddMemberResponse{}
	_body, _err := client.AddMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddPrivacyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPrivacyRuleResponse
func (client *Client) AddPrivacyRuleWithOptions(request *AddPrivacyRuleRequest, runtime *dara.RuntimeOptions) (_result *AddPrivacyRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlgImpl) {
		query["AlgImpl"] = request.AlgImpl
	}

	if !dara.IsNil(request.AlgType) {
		query["AlgType"] = request.AlgType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddPrivacyRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddPrivacyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddPrivacyRuleRequest
//
// @return AddPrivacyRuleResponse
func (client *Client) AddPrivacyRule(request *AddPrivacyRuleRequest) (_result *AddPrivacyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddPrivacyRuleResponse{}
	_body, _err := client.AddPrivacyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddRouteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRouteRuleResponse
func (client *Client) AddRouteRuleWithOptions(request *AddRouteRuleRequest, runtime *dara.RuntimeOptions) (_result *AddRouteRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.ChainUpMode) {
		query["ChainUpMode"] = request.ChainUpMode
	}

	if !dara.IsNil(request.ContractName) {
		query["ContractName"] = request.ContractName
	}

	if !dara.IsNil(request.ContractTemplateId) {
		query["ContractTemplateId"] = request.ContractTemplateId
	}

	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.InvokeType) {
		query["InvokeType"] = request.InvokeType
	}

	if !dara.IsNil(request.PrivacyRuleId) {
		query["PrivacyRuleId"] = request.PrivacyRuleId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRouteRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddRouteRuleRequest
//
// @return AddRouteRuleResponse
func (client *Client) AddRouteRule(request *AddRouteRuleRequest) (_result *AddRouteRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRouteRuleResponse{}
	_body, _err := client.AddRouteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 成员同意接入
//
// @param request - AgreeMemberAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgreeMemberAccessResponse
func (client *Client) AgreeMemberAccessWithOptions(request *AgreeMemberAccessRequest, runtime *dara.RuntimeOptions) (_result *AgreeMemberAccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberAccountId) {
		query["MemberAccountId"] = request.MemberAccountId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AgreeMemberAccess"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AgreeMemberAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 成员同意接入
//
// @param request - AgreeMemberAccessRequest
//
// @return AgreeMemberAccessResponse
func (client *Client) AgreeMemberAccess(request *AgreeMemberAccessRequest) (_result *AgreeMemberAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AgreeMemberAccessResponse{}
	_body, _err := client.AgreeMemberAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AuthorizeBaaSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeBaaSResponse
func (client *Client) AuthorizeBaaSWithOptions(request *AuthorizeBaaSRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeBaaSResponse, _err error) {
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
		Action:      dara.String("AuthorizeBaaS"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeBaaSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuthorizeBaaSRequest
//
// @return AuthorizeBaaSResponse
func (client *Client) AuthorizeBaaS(request *AuthorizeBaaSRequest) (_result *AuthorizeBaaSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthorizeBaaSResponse{}
	_body, _err := client.AuthorizeBaaSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AuthorizeDeviceGroupBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeDeviceGroupBizChainResponse
func (client *Client) AuthorizeDeviceGroupBizChainWithOptions(request *AuthorizeDeviceGroupBizChainRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeDeviceGroupBizChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainIdList) {
		query["BizChainIdList"] = request.BizChainIdList
	}

	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeDeviceGroupBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeDeviceGroupBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuthorizeDeviceGroupBizChainRequest
//
// @return AuthorizeDeviceGroupBizChainResponse
func (client *Client) AuthorizeDeviceGroupBizChain(request *AuthorizeDeviceGroupBizChainRequest) (_result *AuthorizeDeviceGroupBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthorizeDeviceGroupBizChainResponse{}
	_body, _err := client.AuthorizeDeviceGroupBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AuthorizeMemberBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeMemberBizChainResponse
func (client *Client) AuthorizeMemberBizChainWithOptions(request *AuthorizeMemberBizChainRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeMemberBizChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainInfo) {
		query["BizChainInfo"] = request.BizChainInfo
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeMemberBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeMemberBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuthorizeMemberBizChainRequest
//
// @return AuthorizeMemberBizChainResponse
func (client *Client) AuthorizeMemberBizChain(request *AuthorizeMemberBizChainRequest) (_result *AuthorizeMemberBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthorizeMemberBizChainResponse{}
	_body, _err := client.AuthorizeMemberBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeletePrivacyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivacyRuleResponse
func (client *Client) DeletePrivacyRuleWithOptions(request *DeletePrivacyRuleRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivacyRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PrivacyRuleId) {
		query["PrivacyRuleId"] = request.PrivacyRuleId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivacyRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivacyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeletePrivacyRuleRequest
//
// @return DeletePrivacyRuleResponse
func (client *Client) DeletePrivacyRule(request *DeletePrivacyRuleRequest) (_result *DeletePrivacyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePrivacyRuleResponse{}
	_body, _err := client.DeletePrivacyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteRouteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteRuleResponse
func (client *Client) DeleteRouteRuleWithOptions(request *DeleteRouteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteRuleResponse, _err error) {
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

	if !dara.IsNil(request.RouteRuleId) {
		query["RouteRuleId"] = request.RouteRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRouteRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteRouteRuleRequest
//
// @return DeleteRouteRuleResponse
func (client *Client) DeleteRouteRule(request *DeleteRouteRuleRequest) (_result *DeleteRouteRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRouteRuleResponse{}
	_body, _err := client.DeleteRouteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 成员拒绝接入
//
// @param request - DeniedMemberAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeniedMemberAccessResponse
func (client *Client) DeniedMemberAccessWithOptions(request *DeniedMemberAccessRequest, runtime *dara.RuntimeOptions) (_result *DeniedMemberAccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberAccountId) {
		query["MemberAccountId"] = request.MemberAccountId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeniedMemberAccess"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeniedMemberAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 成员拒绝接入
//
// @param request - DeniedMemberAccessRequest
//
// @return DeniedMemberAccessResponse
func (client *Client) DeniedMemberAccess(request *DeniedMemberAccessRequest) (_result *DeniedMemberAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeniedMemberAccessResponse{}
	_body, _err := client.DeniedMemberAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAccountRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountRoleResponse
func (client *Client) DescribeAccountRoleWithOptions(request *DescribeAccountRoleRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountRoleResponse, _err error) {
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
		Action:      dara.String("DescribeAccountRole"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAccountRoleRequest
//
// @return DescribeAccountRoleResponse
func (client *Client) DescribeAccountRole(request *DescribeAccountRoleRequest) (_result *DescribeAccountRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountRoleResponse{}
	_body, _err := client.DescribeAccountRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询管理方信息
//
// @param request - DescribeAdminInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdminInfoResponse
func (client *Client) DescribeAdminInfoWithOptions(request *DescribeAdminInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdminInfoResponse, _err error) {
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
		Action:      dara.String("DescribeAdminInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdminInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询管理方信息
//
// @param request - DescribeAdminInfoRequest
//
// @return DescribeAdminInfoResponse
func (client *Client) DescribeAdminInfo(request *DescribeAdminInfoRequest) (_result *DescribeAdminInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdminInfoResponse{}
	_body, _err := client.DescribeAdminInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeBizChainStatInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBizChainStatInfoResponse
func (client *Client) DescribeBizChainStatInfoWithOptions(request *DescribeBizChainStatInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeBizChainStatInfoResponse, _err error) {
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
		Action:      dara.String("DescribeBizChainStatInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBizChainStatInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeBizChainStatInfoRequest
//
// @return DescribeBizChainStatInfoResponse
func (client *Client) DescribeBizChainStatInfo(request *DescribeBizChainStatInfoRequest) (_result *DescribeBizChainStatInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBizChainStatInfoResponse{}
	_body, _err := client.DescribeBizChainStatInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询概览API信息
//
// @param request - DescribeDashboardApiInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDashboardApiInfoResponse
func (client *Client) DescribeDashboardApiInfoWithOptions(request *DescribeDashboardApiInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDashboardApiInfoResponse, _err error) {
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
		Action:      dara.String("DescribeDashboardApiInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDashboardApiInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询概览API信息
//
// @param request - DescribeDashboardApiInfoRequest
//
// @return DescribeDashboardApiInfoResponse
func (client *Client) DescribeDashboardApiInfo(request *DescribeDashboardApiInfoRequest) (_result *DescribeDashboardApiInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDashboardApiInfoResponse{}
	_body, _err := client.DescribeDashboardApiInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询概览信息
//
// @param request - DescribeDashboardBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDashboardBaseInfoResponse
func (client *Client) DescribeDashboardBaseInfoWithOptions(request *DescribeDashboardBaseInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDashboardBaseInfoResponse, _err error) {
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
		Action:      dara.String("DescribeDashboardBaseInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDashboardBaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询概览信息
//
// @param request - DescribeDashboardBaseInfoRequest
//
// @return DescribeDashboardBaseInfoResponse
func (client *Client) DescribeDashboardBaseInfo(request *DescribeDashboardBaseInfoRequest) (_result *DescribeDashboardBaseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDashboardBaseInfoResponse{}
	_body, _err := client.DescribeDashboardBaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询概览设备信息
//
// @param request - DescribeDashboardDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDashboardDeviceInfoResponse
func (client *Client) DescribeDashboardDeviceInfoWithOptions(request *DescribeDashboardDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDashboardDeviceInfoResponse, _err error) {
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
		Action:      dara.String("DescribeDashboardDeviceInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDashboardDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询概览设备信息
//
// @param request - DescribeDashboardDeviceInfoRequest
//
// @return DescribeDashboardDeviceInfoResponse
func (client *Client) DescribeDashboardDeviceInfo(request *DescribeDashboardDeviceInfoRequest) (_result *DescribeDashboardDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDashboardDeviceInfoResponse{}
	_body, _err := client.DescribeDashboardDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询概览成员API信息
//
// @param request - DescribeDashboardMemberApiInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDashboardMemberApiInfoResponse
func (client *Client) DescribeDashboardMemberApiInfoWithOptions(request *DescribeDashboardMemberApiInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDashboardMemberApiInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDashboardMemberApiInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDashboardMemberApiInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询概览成员API信息
//
// @param request - DescribeDashboardMemberApiInfoRequest
//
// @return DescribeDashboardMemberApiInfoResponse
func (client *Client) DescribeDashboardMemberApiInfo(request *DescribeDashboardMemberApiInfoRequest) (_result *DescribeDashboardMemberApiInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDashboardMemberApiInfoResponse{}
	_body, _err := client.DescribeDashboardMemberApiInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询概览成员设备信息
//
// @param request - DescribeDashboardMemberDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDashboardMemberDeviceInfoResponse
func (client *Client) DescribeDashboardMemberDeviceInfoWithOptions(request *DescribeDashboardMemberDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDashboardMemberDeviceInfoResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDashboardMemberDeviceInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDashboardMemberDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询概览成员设备信息
//
// @param request - DescribeDashboardMemberDeviceInfoRequest
//
// @return DescribeDashboardMemberDeviceInfoResponse
func (client *Client) DescribeDashboardMemberDeviceInfo(request *DescribeDashboardMemberDeviceInfoRequest) (_result *DescribeDashboardMemberDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDashboardMemberDeviceInfoResponse{}
	_body, _err := client.DescribeDashboardMemberDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备信息
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
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceInfo"),
		Version:     dara.String("2021-07-07"),
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
// 查询设备信息
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
// 查询边缘一体机统计信息
//
// @param request - DescribeEdgeStatInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEdgeStatInfoResponse
func (client *Client) DescribeEdgeStatInfoWithOptions(request *DescribeEdgeStatInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeEdgeStatInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EdgeDn) {
		query["EdgeDn"] = request.EdgeDn
	}

	if !dara.IsNil(request.EdgePk) {
		query["EdgePk"] = request.EdgePk
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEdgeStatInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEdgeStatInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询边缘一体机统计信息
//
// @param request - DescribeEdgeStatInfoRequest
//
// @return DescribeEdgeStatInfoResponse
func (client *Client) DescribeEdgeStatInfo(request *DescribeEdgeStatInfoRequest) (_result *DescribeEdgeStatInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEdgeStatInfoResponse{}
	_body, _err := client.DescribeEdgeStatInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeMemberBizChainStatInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMemberBizChainStatInfoResponse
func (client *Client) DescribeMemberBizChainStatInfoWithOptions(request *DescribeMemberBizChainStatInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeMemberBizChainStatInfoResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMemberBizChainStatInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMemberBizChainStatInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMemberBizChainStatInfoRequest
//
// @return DescribeMemberBizChainStatInfoResponse
func (client *Client) DescribeMemberBizChainStatInfo(request *DescribeMemberBizChainStatInfoRequest) (_result *DescribeMemberBizChainStatInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMemberBizChainStatInfoResponse{}
	_body, _err := client.DescribeMemberBizChainStatInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询统计成员API信息
//
// @param request - DescribeMemberStatInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMemberStatInfoResponse
func (client *Client) DescribeMemberStatInfoWithOptions(request *DescribeMemberStatInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeMemberStatInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMemberStatInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMemberStatInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询统计成员API信息
//
// @param request - DescribeMemberStatInfoRequest
//
// @return DescribeMemberStatInfoResponse
func (client *Client) DescribeMemberStatInfo(request *DescribeMemberStatInfoRequest) (_result *DescribeMemberStatInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMemberStatInfoResponse{}
	_body, _err := client.DescribeMemberStatInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeMemberTotalStatInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMemberTotalStatInfoResponse
func (client *Client) DescribeMemberTotalStatInfoWithOptions(request *DescribeMemberTotalStatInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeMemberTotalStatInfoResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMemberTotalStatInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMemberTotalStatInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMemberTotalStatInfoRequest
//
// @return DescribeMemberTotalStatInfoResponse
func (client *Client) DescribeMemberTotalStatInfo(request *DescribeMemberTotalStatInfoRequest) (_result *DescribeMemberTotalStatInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMemberTotalStatInfoResponse{}
	_body, _err := client.DescribeMemberTotalStatInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribePackgeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackgeInfoResponse
func (client *Client) DescribePackgeInfoWithOptions(request *DescribePackgeInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribePackgeInfoResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePackgeInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePackgeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribePackgeInfoRequest
//
// @return DescribePackgeInfoResponse
func (client *Client) DescribePackgeInfo(request *DescribePackgeInfoRequest) (_result *DescribePackgeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePackgeInfoResponse{}
	_body, _err := client.DescribePackgeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询统计设备信息
//
// @param request - DescribeStatDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStatDeviceInfoResponse
func (client *Client) DescribeStatDeviceInfoWithOptions(request *DescribeStatDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeStatDeviceInfoResponse, _err error) {
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
		Action:      dara.String("DescribeStatDeviceInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStatDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询统计设备信息
//
// @param request - DescribeStatDeviceInfoRequest
//
// @return DescribeStatDeviceInfoResponse
func (client *Client) DescribeStatDeviceInfo(request *DescribeStatDeviceInfoRequest) (_result *DescribeStatDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStatDeviceInfoResponse{}
	_body, _err := client.DescribeStatDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询统计成员设备信息
//
// @param request - DescribeStatMemberDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStatMemberDeviceInfoResponse
func (client *Client) DescribeStatMemberDeviceInfoWithOptions(request *DescribeStatMemberDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeStatMemberDeviceInfoResponse, _err error) {
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
		Action:      dara.String("DescribeStatMemberDeviceInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStatMemberDeviceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询统计成员设备信息
//
// @param request - DescribeStatMemberDeviceInfoRequest
//
// @return DescribeStatMemberDeviceInfoResponse
func (client *Client) DescribeStatMemberDeviceInfo(request *DescribeStatMemberDeviceInfoRequest) (_result *DescribeStatMemberDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStatMemberDeviceInfoResponse{}
	_body, _err := client.DescribeStatMemberDeviceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeTotalStatInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTotalStatInfoResponse
func (client *Client) DescribeTotalStatInfoWithOptions(request *DescribeTotalStatInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeTotalStatInfoResponse, _err error) {
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
		Action:      dara.String("DescribeTotalStatInfo"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTotalStatInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeTotalStatInfoRequest
//
// @return DescribeTotalStatInfoResponse
func (client *Client) DescribeTotalStatInfo(request *DescribeTotalStatInfoRequest) (_result *DescribeTotalStatInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTotalStatInfoResponse{}
	_body, _err := client.DescribeTotalStatInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDeviceResponse
func (client *Client) DisableDeviceWithOptions(request *DisableDeviceRequest, runtime *dara.RuntimeOptions) (_result *DisableDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDevice"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableDeviceRequest
//
// @return DisableDeviceResponse
func (client *Client) DisableDevice(request *DisableDeviceRequest) (_result *DisableDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDeviceResponse{}
	_body, _err := client.DisableDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDeviceGroupResponse
func (client *Client) DisableDeviceGroupWithOptions(request *DisableDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *DisableDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDeviceGroup"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableDeviceGroupRequest
//
// @return DisableDeviceGroupResponse
func (client *Client) DisableDeviceGroup(request *DisableDeviceGroupRequest) (_result *DisableDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDeviceGroupResponse{}
	_body, _err := client.DisableDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DownloadPrivacyKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadPrivacyKeyResponse
func (client *Client) DownloadPrivacyKeyWithOptions(request *DownloadPrivacyKeyRequest, runtime *dara.RuntimeOptions) (_result *DownloadPrivacyKeyResponse, _err error) {
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
		Action:      dara.String("DownloadPrivacyKey"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadPrivacyKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DownloadPrivacyKeyRequest
//
// @return DownloadPrivacyKeyResponse
func (client *Client) DownloadPrivacyKey(request *DownloadPrivacyKeyRequest) (_result *DownloadPrivacyKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadPrivacyKeyResponse{}
	_body, _err := client.DownloadPrivacyKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDeviceResponse
func (client *Client) EnableDeviceWithOptions(request *EnableDeviceRequest, runtime *dara.RuntimeOptions) (_result *EnableDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDevice"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableDeviceRequest
//
// @return EnableDeviceResponse
func (client *Client) EnableDevice(request *EnableDeviceRequest) (_result *EnableDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDeviceResponse{}
	_body, _err := client.EnableDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDeviceGroupResponse
func (client *Client) EnableDeviceGroupWithOptions(request *EnableDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *EnableDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDeviceGroup"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableDeviceGroupRequest
//
// @return EnableDeviceGroupResponse
func (client *Client) EnableDeviceGroup(request *EnableDeviceGroupRequest) (_result *EnableDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDeviceGroupResponse{}
	_body, _err := client.EnableDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FreezeMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FreezeMemberResponse
func (client *Client) FreezeMemberWithOptions(request *FreezeMemberRequest, runtime *dara.RuntimeOptions) (_result *FreezeMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FreezeMember"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FreezeMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - FreezeMemberRequest
//
// @return FreezeMemberResponse
func (client *Client) FreezeMember(request *FreezeMemberRequest) (_result *FreezeMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FreezeMemberResponse{}
	_body, _err := client.FreezeMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetEdgeTotalDeviceCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEdgeTotalDeviceCountResponse
func (client *Client) GetEdgeTotalDeviceCountWithOptions(request *GetEdgeTotalDeviceCountRequest, runtime *dara.RuntimeOptions) (_result *GetEdgeTotalDeviceCountResponse, _err error) {
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
		Action:      dara.String("GetEdgeTotalDeviceCount"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEdgeTotalDeviceCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetEdgeTotalDeviceCountRequest
//
// @return GetEdgeTotalDeviceCountResponse
func (client *Client) GetEdgeTotalDeviceCount(request *GetEdgeTotalDeviceCountRequest) (_result *GetEdgeTotalDeviceCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEdgeTotalDeviceCountResponse{}
	_body, _err := client.GetEdgeTotalDeviceCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAllAdminRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllAdminResponse
func (client *Client) ListAllAdminWithOptions(request *ListAllAdminRequest, runtime *dara.RuntimeOptions) (_result *ListAllAdminResponse, _err error) {
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
		Action:      dara.String("ListAllAdmin"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAllAdminRequest
//
// @return ListAllAdminResponse
func (client *Client) ListAllAdmin(request *ListAllAdminRequest) (_result *ListAllAdminResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllAdminResponse{}
	_body, _err := client.ListAllAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAllBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllBizChainResponse
func (client *Client) ListAllBizChainWithOptions(request *ListAllBizChainRequest, runtime *dara.RuntimeOptions) (_result *ListAllBizChainResponse, _err error) {
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
		Action:      dara.String("ListAllBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAllBizChainRequest
//
// @return ListAllBizChainResponse
func (client *Client) ListAllBizChain(request *ListAllBizChainRequest) (_result *ListAllBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllBizChainResponse{}
	_body, _err := client.ListAllBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAllBizChainContractRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllBizChainContractResponse
func (client *Client) ListAllBizChainContractWithOptions(request *ListAllBizChainContractRequest, runtime *dara.RuntimeOptions) (_result *ListAllBizChainContractResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllBizChainContract"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllBizChainContractResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAllBizChainContractRequest
//
// @return ListAllBizChainContractResponse
func (client *Client) ListAllBizChainContract(request *ListAllBizChainContractRequest) (_result *ListAllBizChainContractResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllBizChainContractResponse{}
	_body, _err := client.ListAllBizChainContractWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAllDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllDeviceGroupResponse
func (client *Client) ListAllDeviceGroupWithOptions(request *ListAllDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *ListAllDeviceGroupResponse, _err error) {
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
		Action:      dara.String("ListAllDeviceGroup"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAllDeviceGroupRequest
//
// @return ListAllDeviceGroupResponse
func (client *Client) ListAllDeviceGroup(request *ListAllDeviceGroupRequest) (_result *ListAllDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllDeviceGroupResponse{}
	_body, _err := client.ListAllDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAllMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllMemberResponse
func (client *Client) ListAllMemberWithOptions(request *ListAllMemberRequest, runtime *dara.RuntimeOptions) (_result *ListAllMemberResponse, _err error) {
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
		Action:      dara.String("ListAllMember"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAllMemberRequest
//
// @return ListAllMemberResponse
func (client *Client) ListAllMember(request *ListAllMemberRequest) (_result *ListAllMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllMemberResponse{}
	_body, _err := client.ListAllMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAllPrivacyAlgorithmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllPrivacyAlgorithmResponse
func (client *Client) ListAllPrivacyAlgorithmWithOptions(request *ListAllPrivacyAlgorithmRequest, runtime *dara.RuntimeOptions) (_result *ListAllPrivacyAlgorithmResponse, _err error) {
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
		Action:      dara.String("ListAllPrivacyAlgorithm"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllPrivacyAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAllPrivacyAlgorithmRequest
//
// @return ListAllPrivacyAlgorithmResponse
func (client *Client) ListAllPrivacyAlgorithm(request *ListAllPrivacyAlgorithmRequest) (_result *ListAllPrivacyAlgorithmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllPrivacyAlgorithmResponse{}
	_body, _err := client.ListAllPrivacyAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAllPrivacyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllPrivacyRuleResponse
func (client *Client) ListAllPrivacyRuleWithOptions(request *ListAllPrivacyRuleRequest, runtime *dara.RuntimeOptions) (_result *ListAllPrivacyRuleResponse, _err error) {
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
		Action:      dara.String("ListAllPrivacyRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllPrivacyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAllPrivacyRuleRequest
//
// @return ListAllPrivacyRuleResponse
func (client *Client) ListAllPrivacyRule(request *ListAllPrivacyRuleRequest) (_result *ListAllPrivacyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllPrivacyRuleResponse{}
	_body, _err := client.ListAllPrivacyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAllProductKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllProductKeyResponse
func (client *Client) ListAllProductKeyWithOptions(request *ListAllProductKeyRequest, runtime *dara.RuntimeOptions) (_result *ListAllProductKeyResponse, _err error) {
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
		Action:      dara.String("ListAllProductKey"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllProductKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAllProductKeyRequest
//
// @return ListAllProductKeyResponse
func (client *Client) ListAllProductKey(request *ListAllProductKeyRequest) (_result *ListAllProductKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllProductKeyResponse{}
	_body, _err := client.ListAllProductKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAllSystemContractRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllSystemContractResponse
func (client *Client) ListAllSystemContractWithOptions(request *ListAllSystemContractRequest, runtime *dara.RuntimeOptions) (_result *ListAllSystemContractResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BlockChainType) {
		query["BlockChainType"] = request.BlockChainType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllSystemContract"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllSystemContractResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAllSystemContractRequest
//
// @return ListAllSystemContractResponse
func (client *Client) ListAllSystemContract(request *ListAllSystemContractRequest) (_result *ListAllSystemContractResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllSystemContractResponse{}
	_body, _err := client.ListAllSystemContractWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBaaSAntChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBaaSAntChainResponse
func (client *Client) ListBaaSAntChainWithOptions(request *ListBaaSAntChainRequest, runtime *dara.RuntimeOptions) (_result *ListBaaSAntChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaaSAntChainConsortiumId) {
		query["BaaSAntChainConsortiumId"] = request.BaaSAntChainConsortiumId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBaaSAntChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBaaSAntChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBaaSAntChainRequest
//
// @return ListBaaSAntChainResponse
func (client *Client) ListBaaSAntChain(request *ListBaaSAntChainRequest) (_result *ListBaaSAntChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBaaSAntChainResponse{}
	_body, _err := client.ListBaaSAntChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBaaSAntChainConsortiumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBaaSAntChainConsortiumResponse
func (client *Client) ListBaaSAntChainConsortiumWithOptions(request *ListBaaSAntChainConsortiumRequest, runtime *dara.RuntimeOptions) (_result *ListBaaSAntChainConsortiumResponse, _err error) {
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
		Action:      dara.String("ListBaaSAntChainConsortium"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBaaSAntChainConsortiumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBaaSAntChainConsortiumRequest
//
// @return ListBaaSAntChainConsortiumResponse
func (client *Client) ListBaaSAntChainConsortium(request *ListBaaSAntChainConsortiumRequest) (_result *ListBaaSAntChainConsortiumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBaaSAntChainConsortiumResponse{}
	_body, _err := client.ListBaaSAntChainConsortiumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBaaSAntChainPeerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBaaSAntChainPeerResponse
func (client *Client) ListBaaSAntChainPeerWithOptions(request *ListBaaSAntChainPeerRequest, runtime *dara.RuntimeOptions) (_result *ListBaaSAntChainPeerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaaSAntChainChainId) {
		query["BaaSAntChainChainId"] = request.BaaSAntChainChainId
	}

	if !dara.IsNil(request.BaaSAntChainConsortiumId) {
		query["BaaSAntChainConsortiumId"] = request.BaaSAntChainConsortiumId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBaaSAntChainPeer"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBaaSAntChainPeerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBaaSAntChainPeerRequest
//
// @return ListBaaSAntChainPeerResponse
func (client *Client) ListBaaSAntChainPeer(request *ListBaaSAntChainPeerRequest) (_result *ListBaaSAntChainPeerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBaaSAntChainPeerResponse{}
	_body, _err := client.ListBaaSAntChainPeerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBaaSFabricChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBaaSFabricChannelResponse
func (client *Client) ListBaaSFabricChannelWithOptions(request *ListBaaSFabricChannelRequest, runtime *dara.RuntimeOptions) (_result *ListBaaSFabricChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaaSFabricConsortiumId) {
		query["BaaSFabricConsortiumId"] = request.BaaSFabricConsortiumId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBaaSFabricChannel"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBaaSFabricChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBaaSFabricChannelRequest
//
// @return ListBaaSFabricChannelResponse
func (client *Client) ListBaaSFabricChannel(request *ListBaaSFabricChannelRequest) (_result *ListBaaSFabricChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBaaSFabricChannelResponse{}
	_body, _err := client.ListBaaSFabricChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBaaSFabricConsortiumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBaaSFabricConsortiumResponse
func (client *Client) ListBaaSFabricConsortiumWithOptions(request *ListBaaSFabricConsortiumRequest, runtime *dara.RuntimeOptions) (_result *ListBaaSFabricConsortiumResponse, _err error) {
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
		Action:      dara.String("ListBaaSFabricConsortium"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBaaSFabricConsortiumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBaaSFabricConsortiumRequest
//
// @return ListBaaSFabricConsortiumResponse
func (client *Client) ListBaaSFabricConsortium(request *ListBaaSFabricConsortiumRequest) (_result *ListBaaSFabricConsortiumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBaaSFabricConsortiumResponse{}
	_body, _err := client.ListBaaSFabricConsortiumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBaaSFabricOrganizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBaaSFabricOrganizationResponse
func (client *Client) ListBaaSFabricOrganizationWithOptions(request *ListBaaSFabricOrganizationRequest, runtime *dara.RuntimeOptions) (_result *ListBaaSFabricOrganizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaaSFabricChannelId) {
		query["BaaSFabricChannelId"] = request.BaaSFabricChannelId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBaaSFabricOrganization"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBaaSFabricOrganizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBaaSFabricOrganizationRequest
//
// @return ListBaaSFabricOrganizationResponse
func (client *Client) ListBaaSFabricOrganization(request *ListBaaSFabricOrganizationRequest) (_result *ListBaaSFabricOrganizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBaaSFabricOrganizationResponse{}
	_body, _err := client.ListBaaSFabricOrganizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizChainResponse
func (client *Client) ListBizChainWithOptions(request *ListBizChainRequest, runtime *dara.RuntimeOptions) (_result *ListBizChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBizChainRequest
//
// @return ListBizChainResponse
func (client *Client) ListBizChain(request *ListBizChainRequest) (_result *ListBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBizChainResponse{}
	_body, _err := client.ListBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBizChainDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBizChainDataResponse
func (client *Client) ListBizChainDataWithOptions(request *ListBizChainDataRequest, runtime *dara.RuntimeOptions) (_result *ListBizChainDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IoTDataDID) {
		query["IoTDataDID"] = request.IoTDataDID
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBizChainData"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBizChainDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBizChainDataRequest
//
// @return ListBizChainDataResponse
func (client *Client) ListBizChainData(request *ListBizChainDataRequest) (_result *ListBizChainDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBizChainDataResponse{}
	_body, _err := client.ListBizChainDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceResponse
func (client *Client) ListDeviceWithOptions(request *ListDeviceRequest, runtime *dara.RuntimeOptions) (_result *ListDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDevice"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDeviceRequest
//
// @return ListDeviceResponse
func (client *Client) ListDevice(request *ListDeviceRequest) (_result *ListDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeviceResponse{}
	_body, _err := client.ListDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备组列表
//
// @param request - ListDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceGroupResponse
func (client *Client) ListDeviceGroupWithOptions(request *ListDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *ListDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberName) {
		query["MemberName"] = request.MemberName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceGroup"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备组列表
//
// @param request - ListDeviceGroupRequest
//
// @return ListDeviceGroupResponse
func (client *Client) ListDeviceGroup(request *ListDeviceGroupRequest) (_result *ListDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeviceGroupResponse{}
	_body, _err := client.ListDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListDeviceGroupAuthorizedBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceGroupAuthorizedBizChainResponse
func (client *Client) ListDeviceGroupAuthorizedBizChainWithOptions(request *ListDeviceGroupAuthorizedBizChainRequest, runtime *dara.RuntimeOptions) (_result *ListDeviceGroupAuthorizedBizChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceGroupAuthorizedBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceGroupAuthorizedBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDeviceGroupAuthorizedBizChainRequest
//
// @return ListDeviceGroupAuthorizedBizChainResponse
func (client *Client) ListDeviceGroupAuthorizedBizChain(request *ListDeviceGroupAuthorizedBizChainRequest) (_result *ListDeviceGroupAuthorizedBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeviceGroupAuthorizedBizChainResponse{}
	_body, _err := client.ListDeviceGroupAuthorizedBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询边缘设备列表
//
// @param request - ListEdgeDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeDeviceResponse
func (client *Client) ListEdgeDeviceWithOptions(request *ListEdgeDeviceRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.ProductKey) {
		query["ProductKey"] = request.ProductKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEdgeDevice"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询边缘设备列表
//
// @param request - ListEdgeDeviceRequest
//
// @return ListEdgeDeviceResponse
func (client *Client) ListEdgeDevice(request *ListEdgeDeviceRequest) (_result *ListEdgeDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEdgeDeviceResponse{}
	_body, _err := client.ListEdgeDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询边缘设备组列表
//
// @param request - ListEdgeDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEdgeDeviceGroupResponse
func (client *Client) ListEdgeDeviceGroupWithOptions(request *ListEdgeDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *ListEdgeDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEdgeDeviceGroup"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEdgeDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询边缘设备组列表
//
// @param request - ListEdgeDeviceGroupRequest
//
// @return ListEdgeDeviceGroupResponse
func (client *Client) ListEdgeDeviceGroup(request *ListEdgeDeviceGroupRequest) (_result *ListEdgeDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEdgeDeviceGroupResponse{}
	_body, _err := client.ListEdgeDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询成员列表
//
// @param request - ListMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemberResponse
func (client *Client) ListMemberWithOptions(request *ListMemberRequest, runtime *dara.RuntimeOptions) (_result *ListMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Contactor) {
		query["Contactor"] = request.Contactor
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMember"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询成员列表
//
// @param request - ListMemberRequest
//
// @return ListMemberResponse
func (client *Client) ListMember(request *ListMemberRequest) (_result *ListMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMemberResponse{}
	_body, _err := client.ListMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询成员接入记录分页列表
//
// @param request - ListMemberAccessRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemberAccessRecordResponse
func (client *Client) ListMemberAccessRecordWithOptions(request *ListMemberAccessRecordRequest, runtime *dara.RuntimeOptions) (_result *ListMemberAccessRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessStatus) {
		query["AccessStatus"] = request.AccessStatus
	}

	if !dara.IsNil(request.Contactor) {
		query["Contactor"] = request.Contactor
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMemberAccessRecord"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemberAccessRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询成员接入记录分页列表
//
// @param request - ListMemberAccessRecordRequest
//
// @return ListMemberAccessRecordResponse
func (client *Client) ListMemberAccessRecord(request *ListMemberAccessRecordRequest) (_result *ListMemberAccessRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMemberAccessRecordResponse{}
	_body, _err := client.ListMemberAccessRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMemberAuthorizedBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemberAuthorizedBizChainResponse
func (client *Client) ListMemberAuthorizedBizChainWithOptions(request *ListMemberAuthorizedBizChainRequest, runtime *dara.RuntimeOptions) (_result *ListMemberAuthorizedBizChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMemberAuthorizedBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemberAuthorizedBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMemberAuthorizedBizChainRequest
//
// @return ListMemberAuthorizedBizChainResponse
func (client *Client) ListMemberAuthorizedBizChain(request *ListMemberAuthorizedBizChainRequest) (_result *ListMemberAuthorizedBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMemberAuthorizedBizChainResponse{}
	_body, _err := client.ListMemberAuthorizedBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPrivacyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivacyRuleResponse
func (client *Client) ListPrivacyRuleWithOptions(request *ListPrivacyRuleRequest, runtime *dara.RuntimeOptions) (_result *ListPrivacyRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivacyRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivacyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPrivacyRuleRequest
//
// @return ListPrivacyRuleResponse
func (client *Client) ListPrivacyRule(request *ListPrivacyRuleRequest) (_result *ListPrivacyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivacyRuleResponse{}
	_body, _err := client.ListPrivacyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListPrivacyRuleSharedMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivacyRuleSharedMemberResponse
func (client *Client) ListPrivacyRuleSharedMemberWithOptions(request *ListPrivacyRuleSharedMemberRequest, runtime *dara.RuntimeOptions) (_result *ListPrivacyRuleSharedMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PrivacyRuleId) {
		query["PrivacyRuleId"] = request.PrivacyRuleId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivacyRuleSharedMember"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivacyRuleSharedMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListPrivacyRuleSharedMemberRequest
//
// @return ListPrivacyRuleSharedMemberResponse
func (client *Client) ListPrivacyRuleSharedMember(request *ListPrivacyRuleSharedMemberRequest) (_result *ListPrivacyRuleSharedMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivacyRuleSharedMemberResponse{}
	_body, _err := client.ListPrivacyRuleSharedMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListRouteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRouteRuleResponse
func (client *Client) ListRouteRuleWithOptions(request *ListRouteRuleRequest, runtime *dara.RuntimeOptions) (_result *ListRouteRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainName) {
		query["BizChainName"] = request.BizChainName
	}

	if !dara.IsNil(request.ChainUpMode) {
		query["ChainUpMode"] = request.ChainUpMode
	}

	if !dara.IsNil(request.DeviceGroupName) {
		query["DeviceGroupName"] = request.DeviceGroupName
	}

	if !dara.IsNil(request.Num) {
		query["Num"] = request.Num
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRouteRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListRouteRuleRequest
//
// @return ListRouteRuleResponse
func (client *Client) ListRouteRule(request *ListRouteRuleRequest) (_result *ListRouteRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRouteRuleResponse{}
	_body, _err := client.ListRouteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryBlockchainDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBlockchainDataResponse
func (client *Client) QueryBlockchainDataWithOptions(request *QueryBlockchainDataRequest, runtime *dara.RuntimeOptions) (_result *QueryBlockchainDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.ContractName) {
		query["ContractName"] = request.ContractName
	}

	if !dara.IsNil(request.InvokeType) {
		query["InvokeType"] = request.InvokeType
	}

	if !dara.IsNil(request.IotDataDID) {
		query["IotDataDID"] = request.IotDataDID
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TransactionId) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBlockchainData"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBlockchainDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryBlockchainDataRequest
//
// @return QueryBlockchainDataResponse
func (client *Client) QueryBlockchainData(request *QueryBlockchainDataRequest) (_result *QueryBlockchainDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBlockchainDataResponse{}
	_body, _err := client.QueryBlockchainDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryBlockchainMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBlockchainMetadataResponse
func (client *Client) QueryBlockchainMetadataWithOptions(request *QueryBlockchainMetadataRequest, runtime *dara.RuntimeOptions) (_result *QueryBlockchainMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.ContractName) {
		query["ContractName"] = request.ContractName
	}

	if !dara.IsNil(request.InvokeType) {
		query["InvokeType"] = request.InvokeType
	}

	if !dara.IsNil(request.IotDataDID) {
		query["IotDataDID"] = request.IotDataDID
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TransactionId) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBlockchainMetadata"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBlockchainMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryBlockchainMetadataRequest
//
// @return QueryBlockchainMetadataResponse
func (client *Client) QueryBlockchainMetadata(request *QueryBlockchainMetadataRequest) (_result *QueryBlockchainMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBlockchainMetadataResponse{}
	_body, _err := client.QueryBlockchainMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SharePrivacyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SharePrivacyRuleResponse
func (client *Client) SharePrivacyRuleWithOptions(request *SharePrivacyRuleRequest, runtime *dara.RuntimeOptions) (_result *SharePrivacyRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberIdList) {
		query["MemberIdList"] = request.MemberIdList
	}

	if !dara.IsNil(request.PrivacyRuleId) {
		query["PrivacyRuleId"] = request.PrivacyRuleId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SharePrivacyRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SharePrivacyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SharePrivacyRuleRequest
//
// @return SharePrivacyRuleResponse
func (client *Client) SharePrivacyRule(request *SharePrivacyRuleRequest) (_result *SharePrivacyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SharePrivacyRuleResponse{}
	_body, _err := client.SharePrivacyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UnFreezeMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnFreezeMemberResponse
func (client *Client) UnFreezeMemberWithOptions(request *UnFreezeMemberRequest, runtime *dara.RuntimeOptions) (_result *UnFreezeMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnFreezeMember"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnFreezeMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UnFreezeMemberRequest
//
// @return UnFreezeMemberResponse
func (client *Client) UnFreezeMember(request *UnFreezeMemberRequest) (_result *UnFreezeMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnFreezeMemberResponse{}
	_body, _err := client.UnFreezeMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateBizChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBizChainResponse
func (client *Client) UpdateBizChainWithOptions(request *UpdateBizChainRequest, runtime *dara.RuntimeOptions) (_result *UpdateBizChainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBizChain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBizChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateBizChainRequest
//
// @return UpdateBizChainResponse
func (client *Client) UpdateBizChain(request *UpdateBizChainRequest) (_result *UpdateBizChainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBizChainResponse{}
	_body, _err := client.UpdateBizChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改成员信息
//
// @param request - UpdateMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemberResponse
func (client *Client) UpdateMemberWithOptions(request *UpdateMemberRequest, runtime *dara.RuntimeOptions) (_result *UpdateMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizedCount) {
		query["AuthorizedCount"] = request.AuthorizedCount
	}

	if !dara.IsNil(request.AuthorizedDeviceCount) {
		query["AuthorizedDeviceCount"] = request.AuthorizedDeviceCount
	}

	if !dara.IsNil(request.Contactor) {
		query["Contactor"] = request.Contactor
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Telephony) {
		query["Telephony"] = request.Telephony
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMember"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改成员信息
//
// @param request - UpdateMemberRequest
//
// @return UpdateMemberResponse
func (client *Client) UpdateMember(request *UpdateMemberRequest) (_result *UpdateMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMemberResponse{}
	_body, _err := client.UpdateMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdatePrivacyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivacyRuleResponse
func (client *Client) UpdatePrivacyRuleWithOptions(request *UpdatePrivacyRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrivacyRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlgImpl) {
		query["AlgImpl"] = request.AlgImpl
	}

	if !dara.IsNil(request.AlgType) {
		query["AlgType"] = request.AlgType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PrivacyRuleId) {
		query["PrivacyRuleId"] = request.PrivacyRuleId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivacyRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivacyRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdatePrivacyRuleRequest
//
// @return UpdatePrivacyRuleResponse
func (client *Client) UpdatePrivacyRule(request *UpdatePrivacyRuleRequest) (_result *UpdatePrivacyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePrivacyRuleResponse{}
	_body, _err := client.UpdatePrivacyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateRouteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRouteRuleResponse
func (client *Client) UpdateRouteRuleWithOptions(request *UpdateRouteRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateRouteRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizChainId) {
		query["BizChainId"] = request.BizChainId
	}

	if !dara.IsNil(request.ContractName) {
		query["ContractName"] = request.ContractName
	}

	if !dara.IsNil(request.ContractTemplateId) {
		query["ContractTemplateId"] = request.ContractTemplateId
	}

	if !dara.IsNil(request.InvokeType) {
		query["InvokeType"] = request.InvokeType
	}

	if !dara.IsNil(request.PrivacyRuleId) {
		query["PrivacyRuleId"] = request.PrivacyRuleId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.RouteRuleId) {
		query["RouteRuleId"] = request.RouteRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRouteRule"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRouteRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateRouteRuleRequest
//
// @return UpdateRouteRuleResponse
func (client *Client) UpdateRouteRule(request *UpdateRouteRuleRequest) (_result *UpdateRouteRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRouteRuleResponse{}
	_body, _err := client.UpdateRouteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UploadIoTDataToBlockchainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadIoTDataToBlockchainResponse
func (client *Client) UploadIoTDataToBlockchainWithOptions(request *UploadIoTDataToBlockchainRequest, runtime *dara.RuntimeOptions) (_result *UploadIoTDataToBlockchainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IotAuthType) {
		query["IotAuthType"] = request.IotAuthType
	}

	if !dara.IsNil(request.IotDataDID) {
		query["IotDataDID"] = request.IotDataDID
	}

	if !dara.IsNil(request.IotDataDigest) {
		query["IotDataDigest"] = request.IotDataDigest
	}

	if !dara.IsNil(request.IotDataToken) {
		query["IotDataToken"] = request.IotDataToken
	}

	if !dara.IsNil(request.IotId) {
		query["IotId"] = request.IotId
	}

	if !dara.IsNil(request.IotIdServiceProvider) {
		query["IotIdServiceProvider"] = request.IotIdServiceProvider
	}

	if !dara.IsNil(request.IotIdSource) {
		query["IotIdSource"] = request.IotIdSource
	}

	if !dara.IsNil(request.PlainData) {
		query["PlainData"] = request.PlainData
	}

	if !dara.IsNil(request.PrivacyData) {
		query["PrivacyData"] = request.PrivacyData
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadIoTDataToBlockchain"),
		Version:     dara.String("2021-07-07"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadIoTDataToBlockchainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UploadIoTDataToBlockchainRequest
//
// @return UploadIoTDataToBlockchainResponse
func (client *Client) UploadIoTDataToBlockchain(request *UploadIoTDataToBlockchainRequest) (_result *UploadIoTDataToBlockchainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadIoTDataToBlockchainResponse{}
	_body, _err := client.UploadIoTDataToBlockchainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
