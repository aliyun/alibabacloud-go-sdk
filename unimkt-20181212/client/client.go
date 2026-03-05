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
		"cn-hangzhou": dara.String("cloudcode.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai": dara.String("cloudcode.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("unimkt"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # CreateProxyBrandUser
//
// @param request - CreateProxyBrandUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProxyBrandUserResponse
func (client *Client) CreateProxyBrandUserWithOptions(request *CreateProxyBrandUserRequest, runtime *dara.RuntimeOptions) (_result *CreateProxyBrandUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandUserNick) {
		query["BrandUserNick"] = request.BrandUserNick
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Company) {
		query["Company"] = request.Company
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.ContactPhone) {
		query["ContactPhone"] = request.ContactPhone
	}

	if !dara.IsNil(request.Industry) {
		query["Industry"] = request.Industry
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProxyBrandUser"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProxyBrandUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateProxyBrandUser
//
// @param request - CreateProxyBrandUserRequest
//
// @return CreateProxyBrandUserResponse
func (client *Client) CreateProxyBrandUser(request *CreateProxyBrandUserRequest) (_result *CreateProxyBrandUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProxyBrandUserResponse{}
	_body, _err := client.CreateProxyBrandUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # CreateUnionTask
//
// @param tmpReq - CreateUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUnionTaskResponse
func (client *Client) CreateUnionTaskWithOptions(tmpReq *CreateUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUnionTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MediaIdBlackList) {
		request.MediaIdBlackListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MediaIdBlackList, dara.String("MediaIdBlackList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MediaIdWhiteList) {
		request.MediaIdWhiteListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MediaIdWhiteList, dara.String("MediaIdWhiteList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AnchorId) {
		query["AnchorId"] = request.AnchorId
	}

	if !dara.IsNil(request.BrandUserId) {
		query["BrandUserId"] = request.BrandUserId
	}

	if !dara.IsNil(request.BrandUserNick) {
		query["BrandUserNick"] = request.BrandUserNick
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ChargePloy) {
		query["ChargePloy"] = request.ChargePloy
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ContentId) {
		query["ContentId"] = request.ContentId
	}

	if !dara.IsNil(request.ContentUrl) {
		query["ContentUrl"] = request.ContentUrl
	}

	if !dara.IsNil(request.CustomCreativeType) {
		query["CustomCreativeType"] = request.CustomCreativeType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IndustryLabelBagId) {
		query["IndustryLabelBagId"] = request.IndustryLabelBagId
	}

	if !dara.IsNil(request.MediaIdBlackListShrink) {
		query["MediaIdBlackList"] = request.MediaIdBlackListShrink
	}

	if !dara.IsNil(request.MediaIdWhiteListShrink) {
		query["MediaIdWhiteList"] = request.MediaIdWhiteListShrink
	}

	if !dara.IsNil(request.MediaIndustry) {
		query["MediaIndustry"] = request.MediaIndustry
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OptimizationSwitch) {
		query["OptimizationSwitch"] = request.OptimizationSwitch
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.Quota) {
		query["Quota"] = request.Quota
	}

	if !dara.IsNil(request.QuotaDay) {
		query["QuotaDay"] = request.QuotaDay
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskBizType) {
		query["TaskBizType"] = request.TaskBizType
	}

	if !dara.IsNil(request.TaskRuleType) {
		query["TaskRuleType"] = request.TaskRuleType
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUnionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateUnionTask
//
// @param request - CreateUnionTaskRequest
//
// @return CreateUnionTaskResponse
func (client *Client) CreateUnionTask(request *CreateUnionTaskRequest) (_result *CreateUnionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUnionTaskResponse{}
	_body, _err := client.CreateUnionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除聚合拉新品牌
//
// @param request - DeleteUnionBrandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUnionBrandResponse
func (client *Client) DeleteUnionBrandWithOptions(request *DeleteUnionBrandRequest, runtime *dara.RuntimeOptions) (_result *DeleteUnionBrandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandMainId) {
		query["BrandMainId"] = request.BrandMainId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUnionBrand"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUnionBrandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除聚合拉新品牌
//
// @param request - DeleteUnionBrandRequest
//
// @return DeleteUnionBrandResponse
func (client *Client) DeleteUnionBrand(request *DeleteUnionBrandRequest) (_result *DeleteUnionBrandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUnionBrandResponse{}
	_body, _err := client.DeleteUnionBrandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # EndUnionTask
//
// @param request - EndUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndUnionTaskResponse
func (client *Client) EndUnionTaskWithOptions(request *EndUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *EndUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EndUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EndUnionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # EndUnionTask
//
// @param request - EndUnionTaskRequest
//
// @return EndUnionTaskResponse
func (client *Client) EndUnionTask(request *EndUnionTaskRequest) (_result *EndUnionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EndUnionTaskResponse{}
	_body, _err := client.EndUnionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 聚合拉新开放接口查询媒体行业
//
// @param request - ListUnionMediaIndustryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUnionMediaIndustryResponse
func (client *Client) ListUnionMediaIndustryWithOptions(request *ListUnionMediaIndustryRequest, runtime *dara.RuntimeOptions) (_result *ListUnionMediaIndustryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUnionMediaIndustry"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUnionMediaIndustryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 聚合拉新开放接口查询媒体行业
//
// @param request - ListUnionMediaIndustryRequest
//
// @return ListUnionMediaIndustryResponse
func (client *Client) ListUnionMediaIndustry(request *ListUnionMediaIndustryRequest) (_result *ListUnionMediaIndustryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUnionMediaIndustryResponse{}
	_body, _err := client.ListUnionMediaIndustryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAvailableBalanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvailableBalanceResponse
func (client *Client) QueryAvailableBalanceWithOptions(request *QueryAvailableBalanceRequest, runtime *dara.RuntimeOptions) (_result *QueryAvailableBalanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAvailableBalance"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAvailableBalanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAvailableBalanceRequest
//
// @return QueryAvailableBalanceResponse
func (client *Client) QueryAvailableBalance(request *QueryAvailableBalanceRequest) (_result *QueryAvailableBalanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAvailableBalanceResponse{}
	_body, _err := client.QueryAvailableBalanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # QueryContentList
//
// @param request - QueryContentListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryContentListResponse
func (client *Client) QueryContentListWithOptions(request *QueryContentListRequest, runtime *dara.RuntimeOptions) (_result *QueryContentListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandUserId) {
		query["BrandUserId"] = request.BrandUserId
	}

	if !dara.IsNil(request.BrandUserNick) {
		query["BrandUserNick"] = request.BrandUserNick
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskBizType) {
		query["TaskBizType"] = request.TaskBizType
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryContentList"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryContentListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryContentList
//
// @param request - QueryContentListRequest
//
// @return QueryContentListResponse
func (client *Client) QueryContentList(request *QueryContentListRequest) (_result *QueryContentListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryContentListResponse{}
	_body, _err := client.QueryContentListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取行业标签包
//
// @param request - QueryIndustryLabelBagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIndustryLabelBagResponse
func (client *Client) QueryIndustryLabelBagWithOptions(request *QueryIndustryLabelBagRequest, runtime *dara.RuntimeOptions) (_result *QueryIndustryLabelBagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryIndustryLabelBag"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryIndustryLabelBagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取行业标签包
//
// @param request - QueryIndustryLabelBagRequest
//
// @return QueryIndustryLabelBagResponse
func (client *Client) QueryIndustryLabelBag(request *QueryIndustryLabelBagRequest) (_result *QueryIndustryLabelBagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryIndustryLabelBagResponse{}
	_body, _err := client.QueryIndustryLabelBagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryTaskBizTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskBizTypeResponse
func (client *Client) QueryTaskBizTypeWithOptions(request *QueryTaskBizTypeRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskBizTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskBizType"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskBizTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTaskBizTypeRequest
//
// @return QueryTaskBizTypeResponse
func (client *Client) QueryTaskBizType(request *QueryTaskBizTypeRequest) (_result *QueryTaskBizTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTaskBizTypeResponse{}
	_body, _err := client.QueryTaskBizTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # QueryTaskRuleLimit
//
// @param request - QueryTaskRuleLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskRuleLimitResponse
func (client *Client) QueryTaskRuleLimitWithOptions(request *QueryTaskRuleLimitRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskRuleLimitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskRuleLimit"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskRuleLimitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryTaskRuleLimit
//
// @param request - QueryTaskRuleLimitRequest
//
// @return QueryTaskRuleLimitResponse
func (client *Client) QueryTaskRuleLimit(request *QueryTaskRuleLimitRequest) (_result *QueryTaskRuleLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTaskRuleLimitResponse{}
	_body, _err := client.QueryTaskRuleLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryUnionChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnionChannelResponse
func (client *Client) QueryUnionChannelWithOptions(request *QueryUnionChannelRequest, runtime *dara.RuntimeOptions) (_result *QueryUnionChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUnionChannel"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUnionChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryUnionChannelRequest
//
// @return QueryUnionChannelResponse
func (client *Client) QueryUnionChannel(request *QueryUnionChannelRequest) (_result *QueryUnionChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUnionChannelResponse{}
	_body, _err := client.QueryUnionChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # QueryUnionContentInfo
//
// @param request - QueryUnionContentInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnionContentInfoResponse
func (client *Client) QueryUnionContentInfoWithOptions(request *QueryUnionContentInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryUnionContentInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ContentId) {
		query["ContentId"] = request.ContentId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUnionContentInfo"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUnionContentInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryUnionContentInfo
//
// @param request - QueryUnionContentInfoRequest
//
// @return QueryUnionContentInfoResponse
func (client *Client) QueryUnionContentInfo(request *QueryUnionContentInfoRequest) (_result *QueryUnionContentInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUnionContentInfoResponse{}
	_body, _err := client.QueryUnionContentInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryUnionTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnionTaskInfoResponse
func (client *Client) QueryUnionTaskInfoWithOptions(request *QueryUnionTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryUnionTaskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUnionTaskInfo"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUnionTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryUnionTaskInfoRequest
//
// @return QueryUnionTaskInfoResponse
func (client *Client) QueryUnionTaskInfo(request *QueryUnionTaskInfoRequest) (_result *QueryUnionTaskInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUnionTaskInfoResponse{}
	_body, _err := client.QueryUnionTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryUnionTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnionTaskListResponse
func (client *Client) QueryUnionTaskListWithOptions(request *QueryUnionTaskListRequest, runtime *dara.RuntimeOptions) (_result *QueryUnionTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrandUserId) {
		query["BrandUserId"] = request.BrandUserId
	}

	if !dara.IsNil(request.BrandUserNick) {
		query["BrandUserNick"] = request.BrandUserNick
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUnionTaskList"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUnionTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryUnionTaskListRequest
//
// @return QueryUnionTaskListResponse
func (client *Client) QueryUnionTaskList(request *QueryUnionTaskListRequest) (_result *QueryUnionTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUnionTaskListResponse{}
	_body, _err := client.QueryUnionTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启聚合拉新计划
//
// @param request - StartUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartUnionTaskResponse
func (client *Client) StartUnionTaskWithOptions(request *StartUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *StartUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartUnionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启聚合拉新计划
//
// @param request - StartUnionTaskRequest
//
// @return StartUnionTaskResponse
func (client *Client) StartUnionTask(request *StartUnionTaskRequest) (_result *StartUnionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartUnionTaskResponse{}
	_body, _err := client.StartUnionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暂停聚合拉新计划
//
// @param request - StopUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopUnionTaskResponse
func (client *Client) StopUnionTaskWithOptions(request *StopUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *StopUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopUnionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停聚合拉新计划
//
// @param request - StopUnionTaskRequest
//
// @return StopUnionTaskResponse
func (client *Client) StopUnionTask(request *StopUnionTaskRequest) (_result *StopUnionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopUnionTaskResponse{}
	_body, _err := client.StopUnionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 聚合拉新对外接口更新计划
//
// @param request - UpdateUnionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUnionTaskResponse
func (client *Client) UpdateUnionTaskWithOptions(request *UpdateUnionTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateUnionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ChargePloy) {
		query["ChargePloy"] = request.ChargePloy
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OptimizationSwitch) {
		query["OptimizationSwitch"] = request.OptimizationSwitch
	}

	if !dara.IsNil(request.ProxyUserId) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !dara.IsNil(request.Quota) {
		query["Quota"] = request.Quota
	}

	if !dara.IsNil(request.QuotaDay) {
		query["QuotaDay"] = request.QuotaDay
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUnionTask"),
		Version:     dara.String("2018-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUnionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 聚合拉新对外接口更新计划
//
// @param request - UpdateUnionTaskRequest
//
// @return UpdateUnionTaskResponse
func (client *Client) UpdateUnionTask(request *UpdateUnionTaskRequest) (_result *UpdateUnionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUnionTaskResponse{}
	_body, _err := client.UpdateUnionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
