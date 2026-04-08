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
	client.Endpoint, _err = client.GetEndpoint(dara.String("appflow"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # Generate Login Session Token
//
// @param request - GenerateUserSessionTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUserSessionTokenResponse
func (client *Client) GenerateUserSessionTokenWithOptions(request *GenerateUserSessionTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateUserSessionTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatbotId) {
		query["ChatbotId"] = request.ChatbotId
	}

	if !dara.IsNil(request.ExpireSecond) {
		query["ExpireSecond"] = request.ExpireSecond
	}

	if !dara.IsNil(request.ExtraInfo) {
		query["ExtraInfo"] = request.ExtraInfo
	}

	if !dara.IsNil(request.IntegrateId) {
		query["IntegrateId"] = request.IntegrateId
	}

	if !dara.IsNil(request.UserAvatar) {
		query["UserAvatar"] = request.UserAvatar
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUserSessionToken"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUserSessionTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Generate Login Session Token
//
// @param request - GenerateUserSessionTokenRequest
//
// @return GenerateUserSessionTokenResponse
func (client *Client) GenerateUserSessionToken(request *GenerateUserSessionTokenRequest) (_result *GenerateUserSessionTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateUserSessionTokenResponse{}
	_body, _err := client.GenerateUserSessionTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运行连接器的执行动作
//
// @param tmpReq - InvokeActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeActionResponse
func (client *Client) InvokeActionWithSSE(tmpReq *InvokeActionRequest, runtime *dara.RuntimeOptions, _yield chan *InvokeActionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.invokeActionWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
	return
}

// Summary:
//
// 运行连接器的执行动作
//
// @param tmpReq - InvokeActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeActionResponse
func (client *Client) InvokeActionWithOptions(tmpReq *InvokeActionRequest, runtime *dara.RuntimeOptions) (_result *InvokeActionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InvokeActionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthConfig) {
		request.AuthConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConfig, dara.String("AuthConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("Body"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Headers) {
		request.HeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Headers, dara.String("Headers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Path) {
		request.PathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Path, dara.String("Path"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionId) {
		query["ActionId"] = request.ActionId
	}

	if !dara.IsNil(request.ActionVersion) {
		query["ActionVersion"] = request.ActionVersion
	}

	if !dara.IsNil(request.AuthConfigShrink) {
		query["AuthConfig"] = request.AuthConfigShrink
	}

	if !dara.IsNil(request.BodyShrink) {
		query["Body"] = request.BodyShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	if !dara.IsNil(request.HeadersShrink) {
		query["Headers"] = request.HeadersShrink
	}

	if !dara.IsNil(request.PathShrink) {
		query["Path"] = request.PathShrink
	}

	if !dara.IsNil(request.QueryShrink) {
		query["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeAction"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行连接器的执行动作
//
// @param request - InvokeActionRequest
//
// @return InvokeActionResponse
func (client *Client) InvokeAction(request *InvokeActionRequest) (_result *InvokeActionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InvokeActionResponse{}
	_body, _err := client.InvokeActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) invokeActionWithSSE_opYieldFunc(_yield chan *InvokeActionResponse, _yieldErr chan error, tmpReq *InvokeActionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &InvokeActionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AuthConfig) {
		request.AuthConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AuthConfig, dara.String("AuthConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Body) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, dara.String("Body"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Headers) {
		request.HeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Headers, dara.String("Headers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Path) {
		request.PathShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Path, dara.String("Path"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionId) {
		query["ActionId"] = request.ActionId
	}

	if !dara.IsNil(request.ActionVersion) {
		query["ActionVersion"] = request.ActionVersion
	}

	if !dara.IsNil(request.AuthConfigShrink) {
		query["AuthConfig"] = request.AuthConfigShrink
	}

	if !dara.IsNil(request.BodyShrink) {
		query["Body"] = request.BodyShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.ConnectorVersion) {
		query["ConnectorVersion"] = request.ConnectorVersion
	}

	if !dara.IsNil(request.HeadersShrink) {
		query["Headers"] = request.HeadersShrink
	}

	if !dara.IsNil(request.PathShrink) {
		query["Path"] = request.PathShrink
	}

	if !dara.IsNil(request.QueryShrink) {
		query["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.Stream) {
		query["Stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeAction"),
		Version:     dara.String("2023-09-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
