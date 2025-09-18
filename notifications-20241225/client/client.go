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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("notifications"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 方法描述：删除消息
//
// @param request - DelMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DelMessageResponse
func (client *Client) DelMessageWithOptions(request *DelMessageRequest, runtime *dara.RuntimeOptions) (_result *DelMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DelMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DelMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：删除消息
//
// @param request - DelMessageRequest
//
// @return DelMessageResponse
func (client *Client) DelMessage(request *DelMessageRequest) (_result *DelMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DelMessageResponse{}
	_body, _err := client.DelMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：站内信全部删除（逻辑删除）
//
// @param request - DeleteAllMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAllMessageResponse
func (client *Client) DeleteAllMessageWithOptions(request *DeleteAllMessageRequest, runtime *dara.RuntimeOptions) (_result *DeleteAllMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAllMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAllMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：站内信全部删除（逻辑删除）
//
// @param request - DeleteAllMessageRequest
//
// @return DeleteAllMessageResponse
func (client *Client) DeleteAllMessage(request *DeleteAllMessageRequest) (_result *DeleteAllMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAllMessageResponse{}
	_body, _err := client.DeleteAllMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：分类全部标记为已读，不填则全部标记
//
// @param request - ReadAllMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadAllMessageResponse
func (client *Client) ReadAllMessageWithOptions(request *ReadAllMessageRequest, runtime *dara.RuntimeOptions) (_result *ReadAllMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadAllMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadAllMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：分类全部标记为已读，不填则全部标记
//
// @param request - ReadAllMessageRequest
//
// @return ReadAllMessageResponse
func (client *Client) ReadAllMessage(request *ReadAllMessageRequest) (_result *ReadAllMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadAllMessageResponse{}
	_body, _err := client.ReadAllMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadClassNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadClassNameResponse
func (client *Client) ReadClassNameWithOptions(request *ReadClassNameRequest, runtime *dara.RuntimeOptions) (_result *ReadClassNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadClassName"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadClassNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadClassNameRequest
//
// @return ReadClassNameResponse
func (client *Client) ReadClassName(request *ReadClassNameRequest) (_result *ReadClassNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadClassNameResponse{}
	_body, _err := client.ReadClassNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：消息标记为已读
//
// @param request - ReadMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageResponse
func (client *Client) ReadMessageWithOptions(request *ReadMessageRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessage"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：消息标记为已读
//
// @param request - ReadMessageRequest
//
// @return ReadMessageResponse
func (client *Client) ReadMessage(request *ReadMessageRequest) (_result *ReadMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMessageResponse{}
	_body, _err := client.ReadMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息正文
//
// @param request - ReadMessageContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageContentResponse
func (client *Client) ReadMessageContentWithOptions(request *ReadMessageContentRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.History) {
		body["History"] = request.History
	}

	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageContent"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息正文
//
// @param request - ReadMessageContentRequest
//
// @return ReadMessageContentResponse
func (client *Client) ReadMessageContent(request *ReadMessageContentRequest) (_result *ReadMessageContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMessageContentResponse{}
	_body, _err := client.ReadMessageContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息列表
//
// @param request - ReadMessageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageListResponse
func (client *Client) ReadMessageListWithOptions(request *ReadMessageListRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClassId) {
		body["ClassId"] = request.ClassId
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.History) {
		body["History"] = request.History
	}

	if !dara.IsNil(request.Loc) {
		body["Loc"] = request.Loc
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Page) {
		body["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageList"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取消息列表
//
// @param request - ReadMessageListRequest
//
// @return ReadMessageListResponse
func (client *Client) ReadMessageList(request *ReadMessageListRequest) (_result *ReadMessageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMessageListResponse{}
	_body, _err := client.ReadMessageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取未读消息总数
//
// @param request - ReadMessageNewTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadMessageNewTotalResponse
func (client *Client) ReadMessageNewTotalWithOptions(request *ReadMessageNewTotalRequest, runtime *dara.RuntimeOptions) (_result *ReadMessageNewTotalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadMessageNewTotal"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadMessageNewTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取未读消息总数
//
// @param request - ReadMessageNewTotalRequest
//
// @return ReadMessageNewTotalResponse
func (client *Client) ReadMessageNewTotal(request *ReadMessageNewTotalRequest) (_result *ReadMessageNewTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadMessageNewTotalResponse{}
	_body, _err := client.ReadMessageNewTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadNumGroupByClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadNumGroupByClassResponse
func (client *Client) ReadNumGroupByClassWithOptions(request *ReadNumGroupByClassRequest, runtime *dara.RuntimeOptions) (_result *ReadNumGroupByClassResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadNumGroupByClass"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadNumGroupByClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取各分类已读消息数
//
// @param request - ReadNumGroupByClassRequest
//
// @return ReadNumGroupByClassResponse
func (client *Client) ReadNumGroupByClass(request *ReadNumGroupByClassRequest) (_result *ReadNumGroupByClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadNumGroupByClassResponse{}
	_body, _err := client.ReadNumGroupByClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 方法描述：获取所有分类下的信息
//
// @param request - ReadNumGroupTotalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadNumGroupTotalResponse
func (client *Client) ReadNumGroupTotalWithOptions(request *ReadNumGroupTotalRequest, runtime *dara.RuntimeOptions) (_result *ReadNumGroupTotalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.AppName) {
		body["AppName"] = request.AppName
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.CallerProtocol) {
		body["CallerProtocol"] = request.CallerProtocol
	}

	if !dara.IsNil(request.ClientSource) {
		body["ClientSource"] = request.ClientSource
	}

	if !dara.IsNil(request.Cookies) {
		body["Cookies"] = request.Cookies
	}

	if !dara.IsNil(request.SrcUrl) {
		body["SrcUrl"] = request.SrcUrl
	}

	if !dara.IsNil(request.TenantCode) {
		body["TenantCode"] = request.TenantCode
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UidType) {
		body["UidType"] = request.UidType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadNumGroupTotal"),
		Version:     dara.String("2024-12-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadNumGroupTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 方法描述：获取所有分类下的信息
//
// @param request - ReadNumGroupTotalRequest
//
// @return ReadNumGroupTotalResponse
func (client *Client) ReadNumGroupTotal(request *ReadNumGroupTotalRequest) (_result *ReadNumGroupTotalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadNumGroupTotalResponse{}
	_body, _err := client.ReadNumGroupTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
