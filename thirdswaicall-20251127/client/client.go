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
	client.Endpoint, _err = client.GetEndpoint(dara.String("thirdswaicall"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 外呼任务通话列表查询
//
// @param tmpReq - ReadOutboundTaskCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadOutboundTaskCallListResponse
func (client *Client) ReadOutboundTaskCallListWithOptions(tmpReq *ReadOutboundTaskCallListRequest, runtime *dara.RuntimeOptions) (_result *ReadOutboundTaskCallListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReadOutboundTaskCallListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DisplayStatusList) {
		request.DisplayStatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DisplayStatusList, dara.String("DisplayStatusList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LabelTags) {
		request.LabelTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelTags, dara.String("LabelTags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.CustomerNameOrPhone) {
		body["CustomerNameOrPhone"] = request.CustomerNameOrPhone
	}

	if !dara.IsNil(request.DisplayStatusListShrink) {
		body["DisplayStatusList"] = request.DisplayStatusListShrink
	}

	if !dara.IsNil(request.LabelTagsShrink) {
		body["LabelTags"] = request.LabelTagsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadOutboundTaskCallList"),
		Version:     dara.String("2025-11-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadOutboundTaskCallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 外呼任务通话列表查询
//
// @param request - ReadOutboundTaskCallListRequest
//
// @return ReadOutboundTaskCallListResponse
func (client *Client) ReadOutboundTaskCallList(request *ReadOutboundTaskCallListRequest) (_result *ReadOutboundTaskCallListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadOutboundTaskCallListResponse{}
	_body, _err := client.ReadOutboundTaskCallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
