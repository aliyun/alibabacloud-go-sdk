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
	client.Endpoint, _err = client.GetEndpoint(dara.String("ciomarketpop"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 全员营销
//
// @param request - GetEveryOneSellsFormListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEveryOneSellsFormListResponse
func (client *Client) GetEveryOneSellsFormListWithOptions(request *GetEveryOneSellsFormListRequest, runtime *dara.RuntimeOptions) (_result *GetEveryOneSellsFormListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEveryOneSellsFormList"),
		Version:     dara.String("2025-07-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("array"),
	}
	_result = &GetEveryOneSellsFormListResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 全员营销
//
// @param request - GetEveryOneSellsFormListRequest
//
// @return GetEveryOneSellsFormListResponse
func (client *Client) GetEveryOneSellsFormList(request *GetEveryOneSellsFormListRequest) (_result *GetEveryOneSellsFormListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEveryOneSellsFormListResponse{}
	_body, _err := client.GetEveryOneSellsFormListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送钉钉消息
//
// @param tmpReq - PushEveryOneSellMsgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushEveryOneSellMsgResponse
func (client *Client) PushEveryOneSellMsgWithOptions(tmpReq *PushEveryOneSellMsgRequest, runtime *dara.RuntimeOptions) (_result *PushEveryOneSellMsgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PushEveryOneSellMsgShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DingIdList) {
		request.DingIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DingIdList, dara.String("DingIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DingIdListShrink) {
		body["DingIdList"] = request.DingIdListShrink
	}

	if !dara.IsNil(request.PushMsg) {
		body["PushMsg"] = request.PushMsg
	}

	if !dara.IsNil(request.PushType) {
		body["PushType"] = request.PushType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushEveryOneSellMsg"),
		Version:     dara.String("2025-07-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("string"),
	}
	_result = &PushEveryOneSellMsgResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送钉钉消息
//
// @param request - PushEveryOneSellMsgRequest
//
// @return PushEveryOneSellMsgResponse
func (client *Client) PushEveryOneSellMsg(request *PushEveryOneSellMsgRequest) (_result *PushEveryOneSellMsgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushEveryOneSellMsgResponse{}
	_body, _err := client.PushEveryOneSellMsgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
