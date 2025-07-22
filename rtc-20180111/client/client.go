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
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("rtc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AddRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecordTemplateResponse
func (client *Client) AddRecordTemplateWithOptions(request *AddRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddRecordTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundColor) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.DelayStopTime) {
		query["DelayStopTime"] = request.DelayStopTime
	}

	if !dara.IsNil(request.EnableM3u8DateTime) {
		query["EnableM3u8DateTime"] = request.EnableM3u8DateTime
	}

	if !dara.IsNil(request.FileSplitInterval) {
		query["FileSplitInterval"] = request.FileSplitInterval
	}

	if !dara.IsNil(request.Formats) {
		query["Formats"] = request.Formats
	}

	if !dara.IsNil(request.HttpCallbackUrl) {
		query["HttpCallbackUrl"] = request.HttpCallbackUrl
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MnsQueue) {
		query["MnsQueue"] = request.MnsQueue
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssFilePrefix) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskProfile) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !dara.IsNil(request.Watermarks) {
		query["Watermarks"] = request.Watermarks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddRecordTemplateRequest
//
// @return AddRecordTemplateResponse
func (client *Client) AddRecordTemplate(request *AddRecordTemplateRequest) (_result *AddRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRecordTemplateResponse{}
	_body, _err := client.AddRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用智能体模版
//
// @param tmpReq - CreateAppAgentTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppAgentTemplateResponse
func (client *Client) CreateAppAgentTemplateWithOptions(tmpReq *CreateAppAgentTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppAgentTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppAgentTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AsrConfig) {
		request.AsrConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AsrConfig, dara.String("AsrConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LlmConfig) {
		request.LlmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LlmConfig, dara.String("LlmConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TtsConfig) {
		request.TtsConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TtsConfig, dara.String("TtsConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AsrConfigShrink) {
		query["AsrConfig"] = request.AsrConfigShrink
	}

	if !dara.IsNil(request.ChatMode) {
		query["ChatMode"] = request.ChatMode
	}

	if !dara.IsNil(request.Greeting) {
		query["Greeting"] = request.Greeting
	}

	if !dara.IsNil(request.InterruptMode) {
		query["InterruptMode"] = request.InterruptMode
	}

	if !dara.IsNil(request.LlmConfigShrink) {
		query["LlmConfig"] = request.LlmConfigShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TtsConfigShrink) {
		query["TtsConfig"] = request.TtsConfigShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppAgentTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppAgentTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用智能体模版
//
// @param request - CreateAppAgentTemplateRequest
//
// @return CreateAppAgentTemplateResponse
func (client *Client) CreateAppAgentTemplate(request *CreateAppAgentTemplateRequest) (_result *CreateAppAgentTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppAgentTemplateResponse{}
	_body, _err := client.CreateAppAgentTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增app自定义布局
//
// @param tmpReq - CreateAppLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppLayoutResponse
func (client *Client) CreateAppLayoutWithOptions(tmpReq *CreateAppLayoutRequest, runtime *dara.RuntimeOptions) (_result *CreateAppLayoutResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppLayoutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Layout) {
		request.LayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Layout, dara.String("Layout"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.LayoutShrink) {
		query["Layout"] = request.LayoutShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppLayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppLayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增app自定义布局
//
// @param request - CreateAppLayoutRequest
//
// @return CreateAppLayoutResponse
func (client *Client) CreateAppLayout(request *CreateAppLayoutRequest) (_result *CreateAppLayoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppLayoutResponse{}
	_body, _err := client.CreateAppLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增加应用录制模版
//
// @param tmpReq - CreateAppRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppRecordTemplateResponse
func (client *Client) CreateAppRecordTemplateWithOptions(tmpReq *CreateAppRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppRecordTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppRecordTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordTemplate) {
		request.RecordTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordTemplate, dara.String("RecordTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RecordTemplateShrink) {
		query["RecordTemplate"] = request.RecordTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加应用录制模版
//
// @param request - CreateAppRecordTemplateRequest
//
// @return CreateAppRecordTemplateResponse
func (client *Client) CreateAppRecordTemplate(request *CreateAppRecordTemplateRequest) (_result *CreateAppRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppRecordTemplateResponse{}
	_body, _err := client.CreateAppRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建应用推流模版
//
// @param tmpReq - CreateAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppStreamingOutTemplateResponse
func (client *Client) CreateAppStreamingOutTemplateWithOptions(tmpReq *CreateAppStreamingOutTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppStreamingOutTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppStreamingOutTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamingOutTemplate) {
		request.StreamingOutTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamingOutTemplate, dara.String("StreamingOutTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.StreamingOutTemplateShrink) {
		query["StreamingOutTemplate"] = request.StreamingOutTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppStreamingOutTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppStreamingOutTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用推流模版
//
// @param request - CreateAppStreamingOutTemplateRequest
//
// @return CreateAppStreamingOutTemplateResponse
func (client *Client) CreateAppStreamingOutTemplate(request *CreateAppStreamingOutTemplateRequest) (_result *CreateAppStreamingOutTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppStreamingOutTemplateResponse{}
	_body, _err := client.CreateAppStreamingOutTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoLiveStreamRuleResponse
func (client *Client) CreateAutoLiveStreamRuleWithOptions(request *CreateAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoLiveStreamRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallBack) {
		query["CallBack"] = request.CallBack
	}

	if !dara.IsNil(request.ChannelIdPrefixes) {
		query["ChannelIdPrefixes"] = request.ChannelIdPrefixes
	}

	if !dara.IsNil(request.ChannelIds) {
		query["ChannelIds"] = request.ChannelIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAutoLiveStreamRuleRequest
//
// @return CreateAutoLiveStreamRuleResponse
func (client *Client) CreateAutoLiveStreamRule(request *CreateAutoLiveStreamRuleRequest) (_result *CreateAutoLiveStreamRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAutoLiveStreamRuleResponse{}
	_body, _err := client.CreateAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增加纪要热词表
//
// @param tmpReq - CreateCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudNotePhrasesResponse
func (client *Client) CreateCloudNotePhrasesWithOptions(tmpReq *CreateCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudNotePhrasesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateCloudNotePhrasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Phrase) {
		request.PhraseShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Phrase, dara.String("Phrase"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PhraseShrink) {
		query["Phrase"] = request.PhraseShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudNotePhrases"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudNotePhrasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加纪要热词表
//
// @param request - CreateCloudNotePhrasesRequest
//
// @return CreateCloudNotePhrasesResponse
func (client *Client) CreateCloudNotePhrases(request *CreateCloudNotePhrasesRequest) (_result *CreateCloudNotePhrasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudNotePhrasesResponse{}
	_body, _err := client.CreateCloudNotePhrasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateEventSubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventSubscribeResponse
func (client *Client) CreateEventSubscribeWithOptions(request *CreateEventSubscribeRequest, runtime *dara.RuntimeOptions) (_result *CreateEventSubscribeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Events) {
		query["Events"] = request.Events
	}

	if !dara.IsNil(request.NeedCallbackAuth) {
		query["NeedCallbackAuth"] = request.NeedCallbackAuth
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.Users) {
		query["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEventSubscribe"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEventSubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateEventSubscribeRequest
//
// @return CreateEventSubscribeResponse
func (client *Client) CreateEventSubscribe(request *CreateEventSubscribeRequest) (_result *CreateEventSubscribeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEventSubscribeResponse{}
	_body, _err := client.CreateEventSubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMPULayoutResponse
func (client *Client) CreateMPULayoutWithOptions(request *CreateMPULayoutRequest, runtime *dara.RuntimeOptions) (_result *CreateMPULayoutResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AudioMixCount) {
		query["AudioMixCount"] = request.AudioMixCount
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMPULayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMPULayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMPULayoutRequest
//
// @return CreateMPULayoutResponse
func (client *Client) CreateMPULayout(request *CreateMPULayoutRequest) (_result *CreateMPULayoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMPULayoutResponse{}
	_body, _err := client.CreateMPULayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除应用智能体模版
//
// @param request - DeleteAppAgentTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppAgentTemplateResponse
func (client *Client) DeleteAppAgentTemplateWithOptions(request *DeleteAppAgentTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppAgentTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppAgentTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppAgentTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用智能体模版
//
// @param request - DeleteAppAgentTemplateRequest
//
// @return DeleteAppAgentTemplateResponse
func (client *Client) DeleteAppAgentTemplate(request *DeleteAppAgentTemplateRequest) (_result *DeleteAppAgentTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppAgentTemplateResponse{}
	_body, _err := client.DeleteAppAgentTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除app自定义布局
//
// @param tmpReq - DeleteAppLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppLayoutResponse
func (client *Client) DeleteAppLayoutWithOptions(tmpReq *DeleteAppLayoutRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppLayoutResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteAppLayoutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Layout) {
		request.LayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Layout, dara.String("Layout"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.LayoutShrink) {
		query["Layout"] = request.LayoutShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppLayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppLayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除app自定义布局
//
// @param request - DeleteAppLayoutRequest
//
// @return DeleteAppLayoutResponse
func (client *Client) DeleteAppLayout(request *DeleteAppLayoutRequest) (_result *DeleteAppLayoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppLayoutResponse{}
	_body, _err := client.DeleteAppLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除应用录制模版
//
// @param tmpReq - DeleteAppRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppRecordTemplateResponse
func (client *Client) DeleteAppRecordTemplateWithOptions(tmpReq *DeleteAppRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppRecordTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteAppRecordTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用录制模版
//
// @param request - DeleteAppRecordTemplateRequest
//
// @return DeleteAppRecordTemplateResponse
func (client *Client) DeleteAppRecordTemplate(request *DeleteAppRecordTemplateRequest) (_result *DeleteAppRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppRecordTemplateResponse{}
	_body, _err := client.DeleteAppRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除应用推流模版
//
// @param tmpReq - DeleteAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppStreamingOutTemplateResponse
func (client *Client) DeleteAppStreamingOutTemplateWithOptions(tmpReq *DeleteAppStreamingOutTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppStreamingOutTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteAppStreamingOutTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamingOutTemplate) {
		request.StreamingOutTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamingOutTemplate, dara.String("StreamingOutTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.StreamingOutTemplateShrink) {
		query["StreamingOutTemplate"] = request.StreamingOutTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppStreamingOutTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppStreamingOutTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用推流模版
//
// @param request - DeleteAppStreamingOutTemplateRequest
//
// @return DeleteAppStreamingOutTemplateResponse
func (client *Client) DeleteAppStreamingOutTemplate(request *DeleteAppStreamingOutTemplateRequest) (_result *DeleteAppStreamingOutTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppStreamingOutTemplateResponse{}
	_body, _err := client.DeleteAppStreamingOutTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoLiveStreamRuleResponse
func (client *Client) DeleteAutoLiveStreamRuleWithOptions(request *DeleteAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoLiveStreamRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAutoLiveStreamRuleRequest
//
// @return DeleteAutoLiveStreamRuleResponse
func (client *Client) DeleteAutoLiveStreamRule(request *DeleteAutoLiveStreamRuleRequest) (_result *DeleteAutoLiveStreamRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAutoLiveStreamRuleResponse{}
	_body, _err := client.DeleteAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChannelResponse
func (client *Client) DeleteChannelWithOptions(request *DeleteChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChannel"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteChannelRequest
//
// @return DeleteChannelResponse
func (client *Client) DeleteChannel(request *DeleteChannelRequest) (_result *DeleteChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChannelResponse{}
	_body, _err := client.DeleteChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除纪要热词表
//
// @param tmpReq - DeleteCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudNotePhrasesResponse
func (client *Client) DeleteCloudNotePhrasesWithOptions(tmpReq *DeleteCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudNotePhrasesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteCloudNotePhrasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Phrase) {
		request.PhraseShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Phrase, dara.String("Phrase"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PhraseShrink) {
		query["Phrase"] = request.PhraseShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudNotePhrases"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudNotePhrasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除纪要热词表
//
// @param request - DeleteCloudNotePhrasesRequest
//
// @return DeleteCloudNotePhrasesResponse
func (client *Client) DeleteCloudNotePhrases(request *DeleteCloudNotePhrasesRequest) (_result *DeleteCloudNotePhrasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudNotePhrasesResponse{}
	_body, _err := client.DeleteCloudNotePhrasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteEventSubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventSubscribeResponse
func (client *Client) DeleteEventSubscribeWithOptions(request *DeleteEventSubscribeRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventSubscribeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SubscribeId) {
		query["SubscribeId"] = request.SubscribeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventSubscribe"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventSubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteEventSubscribeRequest
//
// @return DeleteEventSubscribeResponse
func (client *Client) DeleteEventSubscribe(request *DeleteEventSubscribeRequest) (_result *DeleteEventSubscribeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEventSubscribeResponse{}
	_body, _err := client.DeleteEventSubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMPULayoutResponse
func (client *Client) DeleteMPULayoutWithOptions(request *DeleteMPULayoutRequest, runtime *dara.RuntimeOptions) (_result *DeleteMPULayoutResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMPULayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMPULayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMPULayoutRequest
//
// @return DeleteMPULayoutResponse
func (client *Client) DeleteMPULayout(request *DeleteMPULayoutRequest) (_result *DeleteMPULayoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMPULayoutResponse{}
	_body, _err := client.DeleteMPULayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecordTemplateResponse
func (client *Client) DeleteRecordTemplateWithOptions(request *DeleteRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteRecordTemplateRequest
//
// @return DeleteRecordTemplateResponse
func (client *Client) DeleteRecordTemplate(request *DeleteRecordTemplateRequest) (_result *DeleteRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecordTemplateResponse{}
	_body, _err := client.DeleteRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出系统支持的事件回调
//
// @param request - DescribeAllCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllCallbackResponse
func (client *Client) DescribeAllCallbackWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeAllCallbackResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAllCallback"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAllCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出系统支持的事件回调
//
// @return DescribeAllCallbackResponse
func (client *Client) DescribeAllCallback() (_result *DescribeAllCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAllCallbackResponse{}
	_body, _err := client.DescribeAllCallbackWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用智能体开关
//
// @param request - DescribeAppAgentFunctionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppAgentFunctionStatusResponse
func (client *Client) DescribeAppAgentFunctionStatusWithOptions(request *DescribeAppAgentFunctionStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppAgentFunctionStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppAgentFunctionStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppAgentFunctionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用智能体开关
//
// @param request - DescribeAppAgentFunctionStatusRequest
//
// @return DescribeAppAgentFunctionStatusResponse
func (client *Client) DescribeAppAgentFunctionStatus(request *DescribeAppAgentFunctionStatusRequest) (_result *DescribeAppAgentFunctionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppAgentFunctionStatusResponse{}
	_body, _err := client.DescribeAppAgentFunctionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用智能体模版列表
//
// @param request - DescribeAppAgentTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppAgentTemplatesResponse
func (client *Client) DescribeAppAgentTemplatesWithOptions(request *DescribeAppAgentTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppAgentTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppAgentTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppAgentTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用智能体模版列表
//
// @param request - DescribeAppAgentTemplatesRequest
//
// @return DescribeAppAgentTemplatesResponse
func (client *Client) DescribeAppAgentTemplates(request *DescribeAppAgentTemplatesRequest) (_result *DescribeAppAgentTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppAgentTemplatesResponse{}
	_body, _err := client.DescribeAppAgentTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看app回调开关
//
// @param request - DescribeAppCallStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppCallStatusResponse
func (client *Client) DescribeAppCallStatusWithOptions(request *DescribeAppCallStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppCallStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppCallStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppCallStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看app回调开关
//
// @param request - DescribeAppCallStatusRequest
//
// @return DescribeAppCallStatusResponse
func (client *Client) DescribeAppCallStatus(request *DescribeAppCallStatusRequest) (_result *DescribeAppCallStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppCallStatusResponse{}
	_body, _err := client.DescribeAppCallStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取app回调密钥
//
// @param request - DescribeAppCallbackSecretKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppCallbackSecretKeyResponse
func (client *Client) DescribeAppCallbackSecretKeyWithOptions(request *DescribeAppCallbackSecretKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppCallbackSecretKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppCallbackSecretKey"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppCallbackSecretKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取app回调密钥
//
// @param request - DescribeAppCallbackSecretKeyRequest
//
// @return DescribeAppCallbackSecretKeyResponse
func (client *Client) DescribeAppCallbackSecretKey(request *DescribeAppCallbackSecretKeyRequest) (_result *DescribeAppCallbackSecretKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppCallbackSecretKeyResponse{}
	_body, _err := client.DescribeAppCallbackSecretKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看AppKey
//
// @param request - DescribeAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppKeyResponse
func (client *Client) DescribeAppKeyWithOptions(request *DescribeAppKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppKey"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看AppKey
//
// @param request - DescribeAppKeyRequest
//
// @return DescribeAppKeyResponse
func (client *Client) DescribeAppKey(request *DescribeAppKeyRequest) (_result *DescribeAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppKeyResponse{}
	_body, _err := client.DescribeAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询app自定义布局
//
// @param tmpReq - DescribeAppLayoutsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppLayoutsResponse
func (client *Client) DescribeAppLayoutsWithOptions(tmpReq *DescribeAppLayoutsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppLayoutsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeAppLayoutsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Condition) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, dara.String("Condition"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppLayouts"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppLayoutsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询app自定义布局
//
// @param request - DescribeAppLayoutsRequest
//
// @return DescribeAppLayoutsResponse
func (client *Client) DescribeAppLayouts(request *DescribeAppLayoutsRequest) (_result *DescribeAppLayoutsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppLayoutsResponse{}
	_body, _err := client.DescribeAppLayoutsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看应用旁路开关
//
// @param request - DescribeAppLiveStreamStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppLiveStreamStatusResponse
func (client *Client) DescribeAppLiveStreamStatusWithOptions(request *DescribeAppLiveStreamStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppLiveStreamStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppLiveStreamStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppLiveStreamStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看应用旁路开关
//
// @param request - DescribeAppLiveStreamStatusRequest
//
// @return DescribeAppLiveStreamStatusResponse
func (client *Client) DescribeAppLiveStreamStatus(request *DescribeAppLiveStreamStatusRequest) (_result *DescribeAppLiveStreamStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppLiveStreamStatusResponse{}
	_body, _err := client.DescribeAppLiveStreamStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用录制开关
//
// @param request - DescribeAppRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppRecordStatusResponse
func (client *Client) DescribeAppRecordStatusWithOptions(request *DescribeAppRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppRecordStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppRecordStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppRecordStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用录制开关
//
// @param request - DescribeAppRecordStatusRequest
//
// @return DescribeAppRecordStatusResponse
func (client *Client) DescribeAppRecordStatus(request *DescribeAppRecordStatusRequest) (_result *DescribeAppRecordStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppRecordStatusResponse{}
	_body, _err := client.DescribeAppRecordStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用录制模版列表
//
// @param tmpReq - DescribeAppRecordTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppRecordTemplatesResponse
func (client *Client) DescribeAppRecordTemplatesWithOptions(tmpReq *DescribeAppRecordTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppRecordTemplatesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeAppRecordTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Condition) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, dara.String("Condition"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppRecordTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppRecordTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用录制模版列表
//
// @param request - DescribeAppRecordTemplatesRequest
//
// @return DescribeAppRecordTemplatesResponse
func (client *Client) DescribeAppRecordTemplates(request *DescribeAppRecordTemplatesRequest) (_result *DescribeAppRecordTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppRecordTemplatesResponse{}
	_body, _err := client.DescribeAppRecordTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询录制列表
//
// @param tmpReq - DescribeAppRecordingFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppRecordingFilesResponse
func (client *Client) DescribeAppRecordingFilesWithOptions(tmpReq *DescribeAppRecordingFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppRecordingFilesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeAppRecordingFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppRecordingFiles"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppRecordingFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询录制列表
//
// @param request - DescribeAppRecordingFilesRequest
//
// @return DescribeAppRecordingFilesResponse
func (client *Client) DescribeAppRecordingFiles(request *DescribeAppRecordingFilesRequest) (_result *DescribeAppRecordingFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppRecordingFilesResponse{}
	_body, _err := client.DescribeAppRecordingFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用推流模版列表
//
// @param tmpReq - DescribeAppStreamingOutTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppStreamingOutTemplatesResponse
func (client *Client) DescribeAppStreamingOutTemplatesWithOptions(tmpReq *DescribeAppStreamingOutTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppStreamingOutTemplatesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeAppStreamingOutTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Condition) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, dara.String("Condition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ConditionShrink) {
		query["Condition"] = request.ConditionShrink
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppStreamingOutTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppStreamingOutTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用推流模版列表
//
// @param request - DescribeAppStreamingOutTemplatesRequest
//
// @return DescribeAppStreamingOutTemplatesResponse
func (client *Client) DescribeAppStreamingOutTemplates(request *DescribeAppStreamingOutTemplatesRequest) (_result *DescribeAppStreamingOutTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppStreamingOutTemplatesResponse{}
	_body, _err := client.DescribeAppStreamingOutTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # App列表
//
// @param request - DescribeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppVersion) {
		query["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
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
		Action:      dara.String("DescribeApps"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # App列表
//
// @param request - DescribeAppsRequest
//
// @return DescribeAppsResponse
func (client *Client) DescribeApps(request *DescribeAppsRequest) (_result *DescribeAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoLiveStreamRuleResponse
func (client *Client) DescribeAutoLiveStreamRuleWithOptions(request *DescribeAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoLiveStreamRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAutoLiveStreamRuleRequest
//
// @return DescribeAutoLiveStreamRuleResponse
func (client *Client) DescribeAutoLiveStreamRule(request *DescribeAutoLiveStreamRuleRequest) (_result *DescribeAutoLiveStreamRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoLiveStreamRuleResponse{}
	_body, _err := client.DescribeAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeCall获取单次通信详情。
//
// @param request - DescribeCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallResponse
func (client *Client) DescribeCallWithOptions(request *DescribeCallRequest, runtime *dara.RuntimeOptions) (_result *DescribeCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.ExtDataType) {
		query["ExtDataType"] = request.ExtDataType
	}

	if !dara.IsNil(request.QueryExpInfo) {
		query["QueryExpInfo"] = request.QueryExpInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCall"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeCall获取单次通信详情。
//
// @param request - DescribeCallRequest
//
// @return DescribeCallResponse
func (client *Client) DescribeCall(request *DescribeCallRequest) (_result *DescribeCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCallResponse{}
	_body, _err := client.DescribeCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeCallList分页查询时间范围内创建的通信信息。
//
// @param request - DescribeCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallListResponse
func (client *Client) DescribeCallListWithOptions(request *DescribeCallListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCallListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallStatus) {
		query["CallStatus"] = request.CallStatus
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryMode) {
		query["QueryMode"] = request.QueryMode
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCallList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeCallList分页查询时间范围内创建的通信信息。
//
// @param request - DescribeCallListRequest
//
// @return DescribeCallListResponse
func (client *Client) DescribeCallList(request *DescribeCallListRequest) (_result *DescribeCallListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCallListResponse{}
	_body, _err := client.DescribeCallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// app事件回调列表
//
// @param request - DescribeCallbacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallbacksResponse
func (client *Client) DescribeCallbacksWithOptions(request *DescribeCallbacksRequest, runtime *dara.RuntimeOptions) (_result *DescribeCallbacksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCallbacks"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCallbacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// app事件回调列表
//
// @param request - DescribeCallbacksRequest
//
// @return DescribeCallbacksResponse
func (client *Client) DescribeCallbacks(request *DescribeCallbacksRequest) (_result *DescribeCallbacksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCallbacksResponse{}
	_body, _err := client.DescribeCallbacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeChannel
//
// @param request - DescribeChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelResponse
func (client *Client) DescribeChannelWithOptions(request *DescribeChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannel"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeChannel
//
// @param request - DescribeChannelRequest
//
// @return DescribeChannelResponse
func (client *Client) DescribeChannel(request *DescribeChannelRequest) (_result *DescribeChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelResponse{}
	_body, _err := client.DescribeChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询频道的所有参会者
//
// @param request - DescribeChannelAllUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAllUsersResponse
func (client *Client) DescribeChannelAllUsersWithOptions(request *DescribeChannelAllUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelAllUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelAllUsers"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelAllUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询频道的所有参会者
//
// @param request - DescribeChannelAllUsersRequest
//
// @return DescribeChannelAllUsersResponse
func (client *Client) DescribeChannelAllUsers(request *DescribeChannelAllUsersRequest) (_result *DescribeChannelAllUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelAllUsersResponse{}
	_body, _err := client.DescribeChannelAllUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelAreaDistributionStatData获取频道地区分布统计数据。
//
// @param request - DescribeChannelAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAreaDistributionStatDataResponse
func (client *Client) DescribeChannelAreaDistributionStatDataWithOptions(request *DescribeChannelAreaDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelAreaDistributionStatDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.ParentArea) {
		query["ParentArea"] = request.ParentArea
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelAreaDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelAreaDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelAreaDistributionStatData获取频道地区分布统计数据。
//
// @param request - DescribeChannelAreaDistributionStatDataRequest
//
// @return DescribeChannelAreaDistributionStatDataResponse
func (client *Client) DescribeChannelAreaDistributionStatData(request *DescribeChannelAreaDistributionStatDataRequest) (_result *DescribeChannelAreaDistributionStatDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelAreaDistributionStatDataResponse{}
	_body, _err := client.DescribeChannelAreaDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelDistributionStatData获取频道分布统计数据。
//
// @param request - DescribeChannelDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelDistributionStatDataResponse
func (client *Client) DescribeChannelDistributionStatDataWithOptions(request *DescribeChannelDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelDistributionStatDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.StatDim) {
		query["StatDim"] = request.StatDim
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelDistributionStatData获取频道分布统计数据。
//
// @param request - DescribeChannelDistributionStatDataRequest
//
// @return DescribeChannelDistributionStatDataResponse
func (client *Client) DescribeChannelDistributionStatData(request *DescribeChannelDistributionStatDataRequest) (_result *DescribeChannelDistributionStatDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelDistributionStatDataResponse{}
	_body, _err := client.DescribeChannelDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelOverallData查询频道概览数据。
//
// @param request - DescribeChannelOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelOverallDataResponse
func (client *Client) DescribeChannelOverallDataWithOptions(request *DescribeChannelOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelOverallDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelOverallData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelOverallData查询频道概览数据。
//
// @param request - DescribeChannelOverallDataRequest
//
// @return DescribeChannelOverallDataResponse
func (client *Client) DescribeChannelOverallData(request *DescribeChannelOverallDataRequest) (_result *DescribeChannelOverallDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelOverallDataResponse{}
	_body, _err := client.DescribeChannelOverallDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeChannelParticipantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelParticipantsResponse
func (client *Client) DescribeChannelParticipantsWithOptions(request *DescribeChannelParticipantsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelParticipantsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelParticipants"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelParticipantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeChannelParticipantsRequest
//
// @return DescribeChannelParticipantsResponse
func (client *Client) DescribeChannelParticipants(request *DescribeChannelParticipantsRequest) (_result *DescribeChannelParticipantsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelParticipantsResponse{}
	_body, _err := client.DescribeChannelParticipantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelTopPubUserList获取频道内发布端的用户列表（按用户在线时长降序）。
//
// @param request - DescribeChannelTopPubUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelTopPubUserListResponse
func (client *Client) DescribeChannelTopPubUserListWithOptions(request *DescribeChannelTopPubUserListRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelTopPubUserListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelTopPubUserList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelTopPubUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelTopPubUserList获取频道内发布端的用户列表（按用户在线时长降序）。
//
// @param request - DescribeChannelTopPubUserListRequest
//
// @return DescribeChannelTopPubUserListResponse
func (client *Client) DescribeChannelTopPubUserList(request *DescribeChannelTopPubUserListRequest) (_result *DescribeChannelTopPubUserListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelTopPubUserListResponse{}
	_body, _err := client.DescribeChannelTopPubUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeChannelUser
//
// @param request - DescribeChannelUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUserResponse
func (client *Client) DescribeChannelUserWithOptions(request *DescribeChannelUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelUser"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeChannelUser
//
// @param request - DescribeChannelUserRequest
//
// @return DescribeChannelUserResponse
func (client *Client) DescribeChannelUser(request *DescribeChannelUserRequest) (_result *DescribeChannelUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelUserResponse{}
	_body, _err := client.DescribeChannelUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelUserMetrics查询频道总览中的用户数据。
//
// @param request - DescribeChannelUserMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUserMetricsResponse
func (client *Client) DescribeChannelUserMetricsWithOptions(request *DescribeChannelUserMetricsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUserMetricsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelUserMetrics"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelUserMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelUserMetrics查询频道总览中的用户数据。
//
// @param request - DescribeChannelUserMetricsRequest
//
// @return DescribeChannelUserMetricsResponse
func (client *Client) DescribeChannelUserMetrics(request *DescribeChannelUserMetricsRequest) (_result *DescribeChannelUserMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelUserMetricsResponse{}
	_body, _err := client.DescribeChannelUserMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeChannelUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUsersResponse
func (client *Client) DescribeChannelUsersWithOptions(request *DescribeChannelUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannelUsers"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeChannelUsersRequest
//
// @return DescribeChannelUsersResponse
func (client *Client) DescribeChannelUsers(request *DescribeChannelUsersRequest) (_result *DescribeChannelUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelUsersResponse{}
	_body, _err := client.DescribeChannelUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询在线频道列表
//
// @param request - DescribeChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelsResponse
func (client *Client) DescribeChannelsWithOptions(request *DescribeChannelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeChannels"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询在线频道列表
//
// @param request - DescribeChannelsRequest
//
// @return DescribeChannelsResponse
func (client *Client) DescribeChannels(request *DescribeChannelsRequest) (_result *DescribeChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChannelsResponse{}
	_body, _err := client.DescribeChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 纪要热词列表
//
// @param tmpReq - DescribeCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudNotePhrasesResponse
func (client *Client) DescribeCloudNotePhrasesWithOptions(tmpReq *DescribeCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudNotePhrasesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCloudNotePhrasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Condition) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, dara.String("Condition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ConditionShrink) {
		query["Condition"] = request.ConditionShrink
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudNotePhrases"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudNotePhrasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 纪要热词列表
//
// @param request - DescribeCloudNotePhrasesRequest
//
// @return DescribeCloudNotePhrasesResponse
func (client *Client) DescribeCloudNotePhrases(request *DescribeCloudNotePhrasesRequest) (_result *DescribeCloudNotePhrasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudNotePhrasesResponse{}
	_body, _err := client.DescribeCloudNotePhrasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 纪要列表
//
// @param tmpReq - DescribeCloudNotesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudNotesResponse
func (client *Client) DescribeCloudNotesWithOptions(tmpReq *DescribeCloudNotesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudNotesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCloudNotesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudNotes"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudNotesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 纪要列表
//
// @param request - DescribeCloudNotesRequest
//
// @return DescribeCloudNotesResponse
func (client *Client) DescribeCloudNotes(request *DescribeCloudNotesRequest) (_result *DescribeCloudNotesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudNotesResponse{}
	_body, _err := client.DescribeCloudNotesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询录制任务状态
//
// @param request - DescribeCloudRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudRecordStatusResponse
func (client *Client) DescribeCloudRecordStatusWithOptions(request *DescribeCloudRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudRecordStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudRecordStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudRecordStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询录制任务状态
//
// @param request - DescribeCloudRecordStatusRequest
//
// @return DescribeCloudRecordStatusResponse
func (client *Client) DescribeCloudRecordStatus(request *DescribeCloudRecordStatusRequest) (_result *DescribeCloudRecordStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudRecordStatusResponse{}
	_body, _err := client.DescribeCloudRecordStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeEndPointEventList获取端对端用户事件列表。
//
// @param request - DescribeEndPointEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndPointEventListResponse
func (client *Client) DescribeEndPointEventListWithOptions(request *DescribeEndPointEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndPointEventListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.UserIdList) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEndPointEventList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEndPointEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeEndPointEventList获取端对端用户事件列表。
//
// @param request - DescribeEndPointEventListRequest
//
// @return DescribeEndPointEventListResponse
func (client *Client) DescribeEndPointEventList(request *DescribeEndPointEventListRequest) (_result *DescribeEndPointEventListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEndPointEventListResponse{}
	_body, _err := client.DescribeEndPointEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeEndPointMetricData获取端对端指标数据。
//
// @param request - DescribeEndPointMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndPointMetricDataResponse
func (client *Client) DescribeEndPointMetricDataWithOptions(request *DescribeEndPointMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndPointMetricDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.PubCallIdList) {
		query["PubCallIdList"] = request.PubCallIdList
	}

	if !dara.IsNil(request.PubUserId) {
		query["PubUserId"] = request.PubUserId
	}

	if !dara.IsNil(request.SubUserId) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEndPointMetricData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEndPointMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeEndPointMetricData获取端对端指标数据。
//
// @param request - DescribeEndPointMetricDataRequest
//
// @return DescribeEndPointMetricDataResponse
func (client *Client) DescribeEndPointMetricData(request *DescribeEndPointMetricDataRequest) (_result *DescribeEndPointMetricDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEndPointMetricDataResponse{}
	_body, _err := client.DescribeEndPointMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异常诊断影响因素分布
//
// @param request - DescribeFaultDiagnosisFactorDistributionStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisFactorDistributionStatResponse
func (client *Client) DescribeFaultDiagnosisFactorDistributionStatWithOptions(request *DescribeFaultDiagnosisFactorDistributionStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisFactorDistributionStatResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaultDiagnosisFactorDistributionStat"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaultDiagnosisFactorDistributionStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断影响因素分布
//
// @param request - DescribeFaultDiagnosisFactorDistributionStatRequest
//
// @return DescribeFaultDiagnosisFactorDistributionStatResponse
func (client *Client) DescribeFaultDiagnosisFactorDistributionStat(request *DescribeFaultDiagnosisFactorDistributionStatRequest) (_result *DescribeFaultDiagnosisFactorDistributionStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFaultDiagnosisFactorDistributionStatResponse{}
	_body, _err := client.DescribeFaultDiagnosisFactorDistributionStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异常诊断总览数据
//
// @param request - DescribeFaultDiagnosisOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisOverallDataResponse
func (client *Client) DescribeFaultDiagnosisOverallDataWithOptions(request *DescribeFaultDiagnosisOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisOverallDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	if !dara.IsNil(request.StatDim) {
		query["StatDim"] = request.StatDim
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaultDiagnosisOverallData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaultDiagnosisOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断总览数据
//
// @param request - DescribeFaultDiagnosisOverallDataRequest
//
// @return DescribeFaultDiagnosisOverallDataResponse
func (client *Client) DescribeFaultDiagnosisOverallData(request *DescribeFaultDiagnosisOverallDataRequest) (_result *DescribeFaultDiagnosisOverallDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFaultDiagnosisOverallDataResponse{}
	_body, _err := client.DescribeFaultDiagnosisOverallDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户详情
//
// @param request - DescribeFaultDiagnosisUserDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisUserDetailResponse
func (client *Client) DescribeFaultDiagnosisUserDetailWithOptions(request *DescribeFaultDiagnosisUserDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisUserDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.FaultType) {
		query["FaultType"] = request.FaultType
	}

	if !dara.IsNil(request.QueryCallUserInfo) {
		query["QueryCallUserInfo"] = request.QueryCallUserInfo
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaultDiagnosisUserDetail"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaultDiagnosisUserDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户详情
//
// @param request - DescribeFaultDiagnosisUserDetailRequest
//
// @return DescribeFaultDiagnosisUserDetailResponse
func (client *Client) DescribeFaultDiagnosisUserDetail(request *DescribeFaultDiagnosisUserDetailRequest) (_result *DescribeFaultDiagnosisUserDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFaultDiagnosisUserDetailResponse{}
	_body, _err := client.DescribeFaultDiagnosisUserDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户明细列表
//
// @param request - DescribeFaultDiagnosisUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisUserListResponse
func (client *Client) DescribeFaultDiagnosisUserListWithOptions(request *DescribeFaultDiagnosisUserListRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisUserListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.FaultTypes) {
		query["FaultTypes"] = request.FaultTypes
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFaultDiagnosisUserList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFaultDiagnosisUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户明细列表
//
// @param request - DescribeFaultDiagnosisUserListRequest
//
// @return DescribeFaultDiagnosisUserListResponse
func (client *Client) DescribeFaultDiagnosisUserList(request *DescribeFaultDiagnosisUserListRequest) (_result *DescribeFaultDiagnosisUserListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFaultDiagnosisUserListResponse{}
	_body, _err := client.DescribeFaultDiagnosisUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeMPULayoutInfoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMPULayoutInfoListResponse
func (client *Client) DescribeMPULayoutInfoListWithOptions(request *DescribeMPULayoutInfoListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMPULayoutInfoListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMPULayoutInfoList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMPULayoutInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMPULayoutInfoListRequest
//
// @return DescribeMPULayoutInfoListResponse
func (client *Client) DescribeMPULayoutInfoList(request *DescribeMPULayoutInfoListRequest) (_result *DescribeMPULayoutInfoListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMPULayoutInfoListResponse{}
	_body, _err := client.DescribeMPULayoutInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribePubUserListBySubUser根据订阅端获取通信中发布端用户列表。
//
// @param request - DescribePubUserListBySubUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePubUserListBySubUserResponse
func (client *Client) DescribePubUserListBySubUserWithOptions(request *DescribePubUserListBySubUserRequest, runtime *dara.RuntimeOptions) (_result *DescribePubUserListBySubUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.SubUserId) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePubUserListBySubUser"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePubUserListBySubUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribePubUserListBySubUser根据订阅端获取通信中发布端用户列表。
//
// @param request - DescribePubUserListBySubUserRequest
//
// @return DescribePubUserListBySubUserResponse
func (client *Client) DescribePubUserListBySubUser(request *DescribePubUserListBySubUserRequest) (_result *DescribePubUserListBySubUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePubUserListBySubUserResponse{}
	_body, _err := client.DescribePubUserListBySubUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeQoeMetricData获取单次通信中用户的下行体验质量指标。
//
// @param request - DescribeQoeMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQoeMetricDataResponse
func (client *Client) DescribeQoeMetricDataWithOptions(request *DescribeQoeMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQoeMetricDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CreatedTs) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !dara.IsNil(request.DestroyedTs) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQoeMetricData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQoeMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeQoeMetricData获取单次通信中用户的下行体验质量指标。
//
// @param request - DescribeQoeMetricDataRequest
//
// @return DescribeQoeMetricDataResponse
func (client *Client) DescribeQoeMetricData(request *DescribeQoeMetricDataRequest) (_result *DescribeQoeMetricDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeQoeMetricDataResponse{}
	_body, _err := client.DescribeQoeMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质量统计区域分布统计数据
//
// @param request - DescribeQualityAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityAreaDistributionStatDataResponse
func (client *Client) DescribeQualityAreaDistributionStatDataWithOptions(request *DescribeQualityAreaDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityAreaDistributionStatDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ParentArea) {
		query["ParentArea"] = request.ParentArea
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQualityAreaDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQualityAreaDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计区域分布统计数据
//
// @param request - DescribeQualityAreaDistributionStatDataRequest
//
// @return DescribeQualityAreaDistributionStatDataResponse
func (client *Client) DescribeQualityAreaDistributionStatData(request *DescribeQualityAreaDistributionStatDataRequest) (_result *DescribeQualityAreaDistributionStatDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeQualityAreaDistributionStatDataResponse{}
	_body, _err := client.DescribeQualityAreaDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质量统计分布数据
//
// @param request - DescribeQualityDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityDistributionStatDataResponse
func (client *Client) DescribeQualityDistributionStatDataWithOptions(request *DescribeQualityDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityDistributionStatDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatDim) {
		query["StatDim"] = request.StatDim
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQualityDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQualityDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计分布数据
//
// @param request - DescribeQualityDistributionStatDataRequest
//
// @return DescribeQualityDistributionStatDataResponse
func (client *Client) DescribeQualityDistributionStatData(request *DescribeQualityDistributionStatDataRequest) (_result *DescribeQualityDistributionStatDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeQualityDistributionStatDataResponse{}
	_body, _err := client.DescribeQualityDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeQualityOsSdkVersionDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeQualityOsSdkVersionDistributionStatDataWithOptions(request *DescribeQualityOsSdkVersionDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityOsSdkVersionDistributionStatDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQualityOsSdkVersionDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQualityOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeQualityOsSdkVersionDistributionStatDataRequest
//
// @return DescribeQualityOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeQualityOsSdkVersionDistributionStatData(request *DescribeQualityOsSdkVersionDistributionStatDataRequest) (_result *DescribeQualityOsSdkVersionDistributionStatDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeQualityOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.DescribeQualityOsSdkVersionDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质量统计概览数据
//
// @param request - DescribeQualityOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityOverallDataResponse
func (client *Client) DescribeQualityOverallDataWithOptions(request *DescribeQualityOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityOverallDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Types) {
		query["Types"] = request.Types
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQualityOverallData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQualityOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计概览数据
//
// @param request - DescribeQualityOverallDataRequest
//
// @return DescribeQualityOverallDataResponse
func (client *Client) DescribeQualityOverallData(request *DescribeQualityOverallDataRequest) (_result *DescribeQualityOverallDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeQualityOverallDataResponse{}
	_body, _err := client.DescribeQualityOverallDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRecordFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordFilesResponse
func (client *Client) DescribeRecordFilesWithOptions(request *DescribeRecordFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordFiles"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRecordFilesRequest
//
// @return DescribeRecordFilesResponse
func (client *Client) DescribeRecordFiles(request *DescribeRecordFilesRequest) (_result *DescribeRecordFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordFilesResponse{}
	_body, _err := client.DescribeRecordFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRecordTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordTemplatesResponse
func (client *Client) DescribeRecordTemplatesWithOptions(request *DescribeRecordTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecordTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRecordTemplatesRequest
//
// @return DescribeRecordTemplatesResponse
func (client *Client) DescribeRecordTemplates(request *DescribeRecordTemplatesRequest) (_result *DescribeRecordTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordTemplatesResponse{}
	_body, _err := client.DescribeRecordTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRtcChannelListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcChannelListResponse
func (client *Client) DescribeRtcChannelListWithOptions(request *DescribeRtcChannelListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcChannelListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceArea) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	if !dara.IsNil(request.TimePoint) {
		query["TimePoint"] = request.TimePoint
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcChannelList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcChannelListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcChannelListRequest
//
// @return DescribeRtcChannelListResponse
func (client *Client) DescribeRtcChannelList(request *DescribeRtcChannelListRequest) (_result *DescribeRtcChannelListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRtcChannelListResponse{}
	_body, _err := client.DescribeRtcChannelListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRtcChannelMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcChannelMetricResponse
func (client *Client) DescribeRtcChannelMetricWithOptions(request *DescribeRtcChannelMetricRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcChannelMetricResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TimePoint) {
		query["TimePoint"] = request.TimePoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcChannelMetric"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcChannelMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcChannelMetricRequest
//
// @return DescribeRtcChannelMetricResponse
func (client *Client) DescribeRtcChannelMetric(request *DescribeRtcChannelMetricRequest) (_result *DescribeRtcChannelMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRtcChannelMetricResponse{}
	_body, _err := client.DescribeRtcChannelMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRtcDurationDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcDurationDataResponse
func (client *Client) DescribeRtcDurationDataWithOptions(request *DescribeRtcDurationDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcDurationDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ServiceArea) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcDurationData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcDurationDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcDurationDataRequest
//
// @return DescribeRtcDurationDataResponse
func (client *Client) DescribeRtcDurationData(request *DescribeRtcDurationDataRequest) (_result *DescribeRtcDurationDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRtcDurationDataResponse{}
	_body, _err := client.DescribeRtcDurationDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRtcPeakChannelCntDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcPeakChannelCntDataResponse
func (client *Client) DescribeRtcPeakChannelCntDataWithOptions(request *DescribeRtcPeakChannelCntDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcPeakChannelCntDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ServiceArea) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcPeakChannelCntData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcPeakChannelCntDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcPeakChannelCntDataRequest
//
// @return DescribeRtcPeakChannelCntDataResponse
func (client *Client) DescribeRtcPeakChannelCntData(request *DescribeRtcPeakChannelCntDataRequest) (_result *DescribeRtcPeakChannelCntDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRtcPeakChannelCntDataResponse{}
	_body, _err := client.DescribeRtcPeakChannelCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRtcUserCntDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcUserCntDataResponse
func (client *Client) DescribeRtcUserCntDataWithOptions(request *DescribeRtcUserCntDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcUserCntDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ServiceArea) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcUserCntData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcUserCntDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcUserCntDataRequest
//
// @return DescribeRtcUserCntDataResponse
func (client *Client) DescribeRtcUserCntData(request *DescribeRtcUserCntDataRequest) (_result *DescribeRtcUserCntDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRtcUserCntDataResponse{}
	_body, _err := client.DescribeRtcUserCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询旁路推流状态
//
// @param request - DescribeStreamingOutStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamingOutStatusResponse
func (client *Client) DescribeStreamingOutStatusWithOptions(request *DescribeStreamingOutStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamingOutStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStreamingOutStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStreamingOutStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询旁路推流状态
//
// @param request - DescribeStreamingOutStatusRequest
//
// @return DescribeStreamingOutStatusResponse
func (client *Client) DescribeStreamingOutStatus(request *DescribeStreamingOutStatusRequest) (_result *DescribeStreamingOutStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStreamingOutStatusResponse{}
	_body, _err := client.DescribeStreamingOutStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 系统内置布局
//
// @param request - DescribeSystemLayoutListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemLayoutListResponse
func (client *Client) DescribeSystemLayoutListWithOptions(request *DescribeSystemLayoutListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemLayoutListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemLayoutList"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemLayoutListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 系统内置布局
//
// @param request - DescribeSystemLayoutListRequest
//
// @return DescribeSystemLayoutListResponse
func (client *Client) DescribeSystemLayoutList(request *DescribeSystemLayoutListRequest) (_result *DescribeSystemLayoutListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSystemLayoutListResponse{}
	_body, _err := client.DescribeSystemLayoutListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用量统计地域分布数据
//
// @param request - DescribeUsageAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageAreaDistributionStatDataResponse
func (client *Client) DescribeUsageAreaDistributionStatDataWithOptions(request *DescribeUsageAreaDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageAreaDistributionStatDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ParentArea) {
		query["ParentArea"] = request.ParentArea
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsageAreaDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsageAreaDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计地域分布数据
//
// @param request - DescribeUsageAreaDistributionStatDataRequest
//
// @return DescribeUsageAreaDistributionStatDataResponse
func (client *Client) DescribeUsageAreaDistributionStatData(request *DescribeUsageAreaDistributionStatDataRequest) (_result *DescribeUsageAreaDistributionStatDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUsageAreaDistributionStatDataResponse{}
	_body, _err := client.DescribeUsageAreaDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用量统计分布数据
//
// @param request - DescribeUsageDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageDistributionStatDataResponse
func (client *Client) DescribeUsageDistributionStatDataWithOptions(request *DescribeUsageDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageDistributionStatDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatDim) {
		query["StatDim"] = request.StatDim
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsageDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsageDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计分布数据
//
// @param request - DescribeUsageDistributionStatDataRequest
//
// @return DescribeUsageDistributionStatDataResponse
func (client *Client) DescribeUsageDistributionStatData(request *DescribeUsageDistributionStatDataRequest) (_result *DescribeUsageDistributionStatDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUsageDistributionStatDataResponse{}
	_body, _err := client.DescribeUsageDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeUsageOsSdkVersionDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeUsageOsSdkVersionDistributionStatDataWithOptions(request *DescribeUsageOsSdkVersionDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageOsSdkVersionDistributionStatDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsageOsSdkVersionDistributionStatData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsageOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeUsageOsSdkVersionDistributionStatDataRequest
//
// @return DescribeUsageOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeUsageOsSdkVersionDistributionStatData(request *DescribeUsageOsSdkVersionDistributionStatDataRequest) (_result *DescribeUsageOsSdkVersionDistributionStatDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUsageOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.DescribeUsageOsSdkVersionDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用量统计概览数据
//
// @param request - DescribeUsageOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageOverallDataResponse
func (client *Client) DescribeUsageOverallDataWithOptions(request *DescribeUsageOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageOverallDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.Types) {
		query["Types"] = request.Types
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUsageOverallData"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUsageOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计概览数据
//
// @param request - DescribeUsageOverallDataRequest
//
// @return DescribeUsageOverallDataResponse
func (client *Client) DescribeUsageOverallData(request *DescribeUsageOverallDataRequest) (_result *DescribeUsageOverallDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUsageOverallDataResponse{}
	_body, _err := client.DescribeUsageOverallDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeUserInfoInChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserInfoInChannelResponse
func (client *Client) DescribeUserInfoInChannelWithOptions(request *DescribeUserInfoInChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserInfoInChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserInfoInChannel"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserInfoInChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeUserInfoInChannelRequest
//
// @return DescribeUserInfoInChannelResponse
func (client *Client) DescribeUserInfoInChannel(request *DescribeUserInfoInChannelRequest) (_result *DescribeUserInfoInChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserInfoInChannelResponse{}
	_body, _err := client.DescribeUserInfoInChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAutoLiveStreamRuleResponse
func (client *Client) DisableAutoLiveStreamRuleWithOptions(request *DisableAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableAutoLiveStreamRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableAutoLiveStreamRuleRequest
//
// @return DisableAutoLiveStreamRuleResponse
func (client *Client) DisableAutoLiveStreamRule(request *DisableAutoLiveStreamRuleRequest) (_result *DisableAutoLiveStreamRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableAutoLiveStreamRuleResponse{}
	_body, _err := client.DisableAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAutoLiveStreamRuleResponse
func (client *Client) EnableAutoLiveStreamRuleWithOptions(request *EnableAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableAutoLiveStreamRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableAutoLiveStreamRuleRequest
//
// @return EnableAutoLiveStreamRuleResponse
func (client *Client) EnableAutoLiveStreamRule(request *EnableAutoLiveStreamRuleRequest) (_result *EnableAutoLiveStreamRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableAutoLiveStreamRuleResponse{}
	_body, _err := client.EnableAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// GetAgent。
//
// @param request - GetAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithOptions(request *GetAgentRequest, runtime *dara.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// GetAgent。
//
// @param request - GetAgentRequest
//
// @return GetAgentResponse
func (client *Client) GetAgent(request *GetAgentRequest) (_result *GetAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAgentResponse{}
	_body, _err := client.GetAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetMPUTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMPUTaskStatusResponse
func (client *Client) GetMPUTaskStatusWithOptions(request *GetMPUTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetMPUTaskStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMPUTaskStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMPUTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMPUTaskStatusRequest
//
// @return GetMPUTaskStatusResponse
func (client *Client) GetMPUTaskStatus(request *GetMPUTaskStatusRequest) (_result *GetMPUTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMPUTaskStatusResponse{}
	_body, _err := client.GetMPUTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改App信息
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApp"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改App信息
//
// @param request - ModifyAppRequest
//
// @return ModifyAppResponse
func (client *Client) ModifyApp(request *ModifyAppRequest) (_result *ModifyAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用智能体开关
//
// @param request - ModifyAppAgentFunctionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppAgentFunctionStatusResponse
func (client *Client) ModifyAppAgentFunctionStatusWithOptions(request *ModifyAppAgentFunctionStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppAgentFunctionStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppAgentFunctionStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppAgentFunctionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用智能体开关
//
// @param request - ModifyAppAgentFunctionStatusRequest
//
// @return ModifyAppAgentFunctionStatusResponse
func (client *Client) ModifyAppAgentFunctionStatus(request *ModifyAppAgentFunctionStatusRequest) (_result *ModifyAppAgentFunctionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppAgentFunctionStatusResponse{}
	_body, _err := client.ModifyAppAgentFunctionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新应用智能体模版
//
// @param tmpReq - ModifyAppAgentTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppAgentTemplateResponse
func (client *Client) ModifyAppAgentTemplateWithOptions(tmpReq *ModifyAppAgentTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppAgentTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAppAgentTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AsrConfig) {
		request.AsrConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AsrConfig, dara.String("AsrConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LlmConfig) {
		request.LlmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LlmConfig, dara.String("LlmConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TtsConfig) {
		request.TtsConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TtsConfig, dara.String("TtsConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AsrConfigShrink) {
		query["AsrConfig"] = request.AsrConfigShrink
	}

	if !dara.IsNil(request.ChatMode) {
		query["ChatMode"] = request.ChatMode
	}

	if !dara.IsNil(request.Greeting) {
		query["Greeting"] = request.Greeting
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InterruptMode) {
		query["InterruptMode"] = request.InterruptMode
	}

	if !dara.IsNil(request.LlmConfigShrink) {
		query["LlmConfig"] = request.LlmConfigShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TtsConfigShrink) {
		query["TtsConfig"] = request.TtsConfigShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppAgentTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppAgentTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用智能体模版
//
// @param request - ModifyAppAgentTemplateRequest
//
// @return ModifyAppAgentTemplateResponse
func (client *Client) ModifyAppAgentTemplate(request *ModifyAppAgentTemplateRequest) (_result *ModifyAppAgentTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppAgentTemplateResponse{}
	_body, _err := client.ModifyAppAgentTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新app回调事件开关
//
// @param request - ModifyAppCallbackStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppCallbackStatusResponse
func (client *Client) ModifyAppCallbackStatusWithOptions(request *ModifyAppCallbackStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppCallbackStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppCallbackStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppCallbackStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新app回调事件开关
//
// @param request - ModifyAppCallbackStatusRequest
//
// @return ModifyAppCallbackStatusResponse
func (client *Client) ModifyAppCallbackStatus(request *ModifyAppCallbackStatusRequest) (_result *ModifyAppCallbackStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppCallbackStatusResponse{}
	_body, _err := client.ModifyAppCallbackStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改app自定义布局
//
// @param tmpReq - ModifyAppLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppLayoutResponse
func (client *Client) ModifyAppLayoutWithOptions(tmpReq *ModifyAppLayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppLayoutResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAppLayoutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Layout) {
		request.LayoutShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Layout, dara.String("Layout"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.LayoutShrink) {
		query["Layout"] = request.LayoutShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppLayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppLayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改app自定义布局
//
// @param request - ModifyAppLayoutRequest
//
// @return ModifyAppLayoutResponse
func (client *Client) ModifyAppLayout(request *ModifyAppLayoutRequest) (_result *ModifyAppLayoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppLayoutResponse{}
	_body, _err := client.ModifyAppLayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用旁路开关
//
// @param request - ModifyAppLiveStreamStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppLiveStreamStatusResponse
func (client *Client) ModifyAppLiveStreamStatusWithOptions(request *ModifyAppLiveStreamStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppLiveStreamStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppLiveStreamStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppLiveStreamStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用旁路开关
//
// @param request - ModifyAppLiveStreamStatusRequest
//
// @return ModifyAppLiveStreamStatusResponse
func (client *Client) ModifyAppLiveStreamStatus(request *ModifyAppLiveStreamStatusRequest) (_result *ModifyAppLiveStreamStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppLiveStreamStatusResponse{}
	_body, _err := client.ModifyAppLiveStreamStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用录制开关
//
// @param request - ModifyAppRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppRecordStatusResponse
func (client *Client) ModifyAppRecordStatusWithOptions(request *ModifyAppRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppRecordStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppRecordStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppRecordStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用录制开关
//
// @param request - ModifyAppRecordStatusRequest
//
// @return ModifyAppRecordStatusResponse
func (client *Client) ModifyAppRecordStatus(request *ModifyAppRecordStatusRequest) (_result *ModifyAppRecordStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppRecordStatusResponse{}
	_body, _err := client.ModifyAppRecordStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改应用录制模版
//
// @param tmpReq - ModifyAppRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppRecordTemplateResponse
func (client *Client) ModifyAppRecordTemplateWithOptions(tmpReq *ModifyAppRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppRecordTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAppRecordTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordTemplate) {
		request.RecordTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordTemplate, dara.String("RecordTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RecordTemplateShrink) {
		query["RecordTemplate"] = request.RecordTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改应用录制模版
//
// @param request - ModifyAppRecordTemplateRequest
//
// @return ModifyAppRecordTemplateResponse
func (client *Client) ModifyAppRecordTemplate(request *ModifyAppRecordTemplateRequest) (_result *ModifyAppRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppRecordTemplateResponse{}
	_body, _err := client.ModifyAppRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新应用推流模版
//
// @param tmpReq - ModifyAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppStreamingOutTemplateResponse
func (client *Client) ModifyAppStreamingOutTemplateWithOptions(tmpReq *ModifyAppStreamingOutTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppStreamingOutTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAppStreamingOutTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamingOutTemplate) {
		request.StreamingOutTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamingOutTemplate, dara.String("StreamingOutTemplate"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.StreamingOutTemplateShrink) {
		query["StreamingOutTemplate"] = request.StreamingOutTemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppStreamingOutTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppStreamingOutTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用推流模版
//
// @param request - ModifyAppStreamingOutTemplateRequest
//
// @return ModifyAppStreamingOutTemplateResponse
func (client *Client) ModifyAppStreamingOutTemplate(request *ModifyAppStreamingOutTemplateRequest) (_result *ModifyAppStreamingOutTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppStreamingOutTemplateResponse{}
	_body, _err := client.ModifyAppStreamingOutTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新app回调
//
// @param tmpReq - ModifyCallbackMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCallbackMetaResponse
func (client *Client) ModifyCallbackMetaWithOptions(tmpReq *ModifyCallbackMetaRequest, runtime *dara.RuntimeOptions) (_result *ModifyCallbackMetaResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyCallbackMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Callback) {
		request.CallbackShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Callback, dara.String("Callback"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackShrink) {
		query["Callback"] = request.CallbackShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCallbackMeta"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCallbackMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新app回调
//
// @param request - ModifyCallbackMetaRequest
//
// @return ModifyCallbackMetaResponse
func (client *Client) ModifyCallbackMeta(request *ModifyCallbackMetaRequest) (_result *ModifyCallbackMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCallbackMetaResponse{}
	_body, _err := client.ModifyCallbackMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新纪要热词表
//
// @param tmpReq - ModifyCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudNotePhrasesResponse
func (client *Client) ModifyCloudNotePhrasesWithOptions(tmpReq *ModifyCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudNotePhrasesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyCloudNotePhrasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Phrase) {
		request.PhraseShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Phrase, dara.String("Phrase"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PhraseShrink) {
		query["Phrase"] = request.PhraseShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCloudNotePhrases"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCloudNotePhrasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新纪要热词表
//
// @param request - ModifyCloudNotePhrasesRequest
//
// @return ModifyCloudNotePhrasesResponse
func (client *Client) ModifyCloudNotePhrases(request *ModifyCloudNotePhrasesRequest) (_result *ModifyCloudNotePhrasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCloudNotePhrasesResponse{}
	_body, _err := client.ModifyCloudNotePhrasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMPULayoutResponse
func (client *Client) ModifyMPULayoutWithOptions(request *ModifyMPULayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyMPULayoutResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AudioMixCount) {
		query["AudioMixCount"] = request.AudioMixCount
	}

	if !dara.IsNil(request.LayoutId) {
		query["LayoutId"] = request.LayoutId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMPULayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMPULayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyMPULayoutRequest
//
// @return ModifyMPULayoutResponse
func (client *Client) ModifyMPULayout(request *ModifyMPULayoutRequest) (_result *ModifyMPULayoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMPULayoutResponse{}
	_body, _err := client.ModifyMPULayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # NotifyAgent
//
// @param request - NotifyAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NotifyAgentResponse
func (client *Client) NotifyAgentWithOptions(request *NotifyAgentRequest, runtime *dara.RuntimeOptions) (_result *NotifyAgentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CustomAttribute) {
		query["CustomAttribute"] = request.CustomAttribute
	}

	if !dara.IsNil(request.Interruptable) {
		query["Interruptable"] = request.Interruptable
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("NotifyAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &NotifyAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # NotifyAgent
//
// @param request - NotifyAgentRequest
//
// @return NotifyAgentResponse
func (client *Client) NotifyAgent(request *NotifyAgentRequest) (_result *NotifyAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &NotifyAgentResponse{}
	_body, _err := client.NotifyAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemoveTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTerminalsResponse
func (client *Client) RemoveTerminalsWithOptions(request *RemoveTerminalsRequest, runtime *dara.RuntimeOptions) (_result *RemoveTerminalsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TerminalIds) {
		query["TerminalIds"] = request.TerminalIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTerminals"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveTerminalsRequest
//
// @return RemoveTerminalsResponse
func (client *Client) RemoveTerminals(request *RemoveTerminalsRequest) (_result *RemoveTerminalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTerminalsResponse{}
	_body, _err := client.RemoveTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # RemoveUsers
//
// @param request - RemoveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsersWithOptions(request *RemoveUsersRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Users) {
		query["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUsers"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RemoveUsers
//
// @param request - RemoveUsersRequest
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsers(request *RemoveUsersRequest) (_result *RemoveUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUsersResponse{}
	_body, _err := client.RemoveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动AI Agent
//
// @param tmpReq - StartAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAgentResponse
func (client *Client) StartAgentWithOptions(tmpReq *StartAgentRequest, runtime *dara.RuntimeOptions) (_result *StartAgentResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RtcConfig) {
		request.RtcConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RtcConfig, dara.String("RtcConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VoiceChatConfig) {
		request.VoiceChatConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VoiceChatConfig, dara.String("VoiceChatConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.RtcConfigShrink) {
		query["RtcConfig"] = request.RtcConfigShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.VoiceChatConfigShrink) {
		query["VoiceChatConfig"] = request.VoiceChatConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动AI Agent
//
// @param request - StartAgentRequest
//
// @return StartAgentResponse
func (client *Client) StartAgent(request *StartAgentRequest) (_result *StartAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartAgentResponse{}
	_body, _err := client.StartAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启某个事件回调
//
// @param tmpReq - StartCategoryCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCategoryCallbackResponse
func (client *Client) StartCategoryCallbackWithOptions(tmpReq *StartCategoryCallbackRequest, runtime *dara.RuntimeOptions) (_result *StartCategoryCallbackResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartCategoryCallbackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Callback) {
		request.CallbackShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Callback, dara.String("Callback"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackShrink) {
		query["Callback"] = request.CallbackShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCategoryCallback"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCategoryCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启某个事件回调
//
// @param request - StartCategoryCallbackRequest
//
// @return StartCategoryCallbackResponse
func (client *Client) StartCategoryCallback(request *StartCategoryCallbackRequest) (_result *StartCategoryCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartCategoryCallbackResponse{}
	_body, _err := client.StartCategoryCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启智能纪要
//
// @param tmpReq - StartCloudNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudNoteResponse
func (client *Client) StartCloudNoteWithOptions(tmpReq *StartCloudNoteRequest, runtime *dara.RuntimeOptions) (_result *StartCloudNoteResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartCloudNoteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AutoChapters) {
		request.AutoChaptersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AutoChapters, dara.String("AutoChapters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomPrompt) {
		request.CustomPromptShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomPrompt, dara.String("CustomPrompt"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MeetingAssistance) {
		request.MeetingAssistanceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MeetingAssistance, dara.String("MeetingAssistance"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RealtimeSubtitle) {
		request.RealtimeSubtitleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RealtimeSubtitle, dara.String("RealtimeSubtitle"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ServiceInspection) {
		request.ServiceInspectionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceInspection, dara.String("ServiceInspection"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Summarization) {
		request.SummarizationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Summarization, dara.String("Summarization"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TextPolish) {
		request.TextPolishShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextPolish, dara.String("TextPolish"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Transcription) {
		request.TranscriptionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Transcription, dara.String("Transcription"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AutoChaptersShrink) {
		query["AutoChapters"] = request.AutoChaptersShrink
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CustomPromptShrink) {
		query["CustomPrompt"] = request.CustomPromptShrink
	}

	if !dara.IsNil(request.LanguageHints) {
		query["LanguageHints"] = request.LanguageHints
	}

	if !dara.IsNil(request.MeetingAssistanceShrink) {
		query["MeetingAssistance"] = request.MeetingAssistanceShrink
	}

	if !dara.IsNil(request.RealtimeSubtitleShrink) {
		query["RealtimeSubtitle"] = request.RealtimeSubtitleShrink
	}

	if !dara.IsNil(request.ServiceInspectionShrink) {
		query["ServiceInspection"] = request.ServiceInspectionShrink
	}

	if !dara.IsNil(request.SourceLanguage) {
		query["SourceLanguage"] = request.SourceLanguage
	}

	if !dara.IsNil(request.StorageConfig) {
		query["StorageConfig"] = request.StorageConfig
	}

	if !dara.IsNil(request.SummarizationShrink) {
		query["Summarization"] = request.SummarizationShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TextPolishShrink) {
		query["TextPolish"] = request.TextPolishShrink
	}

	if !dara.IsNil(request.TranscriptionShrink) {
		query["Transcription"] = request.TranscriptionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCloudNote"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCloudNoteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启智能纪要
//
// @param request - StartCloudNoteRequest
//
// @return StartCloudNoteResponse
func (client *Client) StartCloudNote(request *StartCloudNoteRequest) (_result *StartCloudNoteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartCloudNoteResponse{}
	_body, _err := client.StartCloudNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # StartCloudRecord
//
// @param tmpReq - StartCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudRecordResponse
func (client *Client) StartCloudRecordWithOptions(tmpReq *StartCloudRecordRequest, runtime *dara.RuntimeOptions) (_result *StartCloudRecordResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartCloudRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LayoutSpecifiedUsers) {
		request.LayoutSpecifiedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LayoutSpecifiedUsers, dara.String("LayoutSpecifiedUsers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SingleStreamingRecord) {
		request.SingleStreamingRecordShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SingleStreamingRecord, dara.String("SingleStreamingRecord"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Annotation) {
		query["Annotation"] = request.Annotation
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.LayoutSpecifiedUsersShrink) {
		query["LayoutSpecifiedUsers"] = request.LayoutSpecifiedUsersShrink
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	if !dara.IsNil(request.RecordMode) {
		query["RecordMode"] = request.RecordMode
	}

	if !dara.IsNil(request.RegionColor) {
		query["RegionColor"] = request.RegionColor
	}

	if !dara.IsNil(request.ReservePaneForNoCameraUser) {
		query["ReservePaneForNoCameraUser"] = request.ReservePaneForNoCameraUser
	}

	if !dara.IsNil(request.ShowDefaultBackgroundOnMute) {
		query["ShowDefaultBackgroundOnMute"] = request.ShowDefaultBackgroundOnMute
	}

	if !dara.IsNil(request.SingleStreamingRecordShrink) {
		query["SingleStreamingRecord"] = request.SingleStreamingRecordShrink
	}

	if !dara.IsNil(request.StartWithoutChannel) {
		query["StartWithoutChannel"] = request.StartWithoutChannel
	}

	if !dara.IsNil(request.StartWithoutChannelWaitTime) {
		query["StartWithoutChannelWaitTime"] = request.StartWithoutChannelWaitTime
	}

	if !dara.IsNil(request.StorageConfig) {
		query["StorageConfig"] = request.StorageConfig
	}

	if !dara.IsNil(request.SubHighResolutionStream) {
		query["SubHighResolutionStream"] = request.SubHighResolutionStream
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Texts) {
		query["Texts"] = request.Texts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCloudRecord"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StartCloudRecord
//
// @param request - StartCloudRecordRequest
//
// @return StartCloudRecordResponse
func (client *Client) StartCloudRecord(request *StartCloudRecordRequest) (_result *StartCloudRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.StartCloudRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMPUTaskResponse
func (client *Client) StartMPUTaskWithOptions(request *StartMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *StartMPUTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundColor) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MixMode) {
		query["MixMode"] = request.MixMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayloadType) {
		query["PayloadType"] = request.PayloadType
	}

	if !dara.IsNil(request.ReportVad) {
		query["ReportVad"] = request.ReportVad
	}

	if !dara.IsNil(request.RtpExtInfo) {
		query["RtpExtInfo"] = request.RtpExtInfo
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.StreamURL) {
		query["StreamURL"] = request.StreamURL
	}

	if !dara.IsNil(request.SubSpecAudioUsers) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !dara.IsNil(request.SubSpecCameraUsers) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !dara.IsNil(request.SubSpecShareScreenUsers) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !dara.IsNil(request.SubSpecUsers) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TimeStampRef) {
		query["TimeStampRef"] = request.TimeStampRef
	}

	if !dara.IsNil(request.UnsubSpecAudioUsers) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !dara.IsNil(request.UnsubSpecCameraUsers) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !dara.IsNil(request.UnsubSpecShareScreenUsers) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !dara.IsNil(request.UserPanes) {
		query["UserPanes"] = request.UserPanes
	}

	if !dara.IsNil(request.VadInterval) {
		query["VadInterval"] = request.VadInterval
	}

	if !dara.IsNil(request.Watermarks) {
		query["Watermarks"] = request.Watermarks
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.EnhancedParam) {
		bodyFlat["EnhancedParam"] = request.EnhancedParam
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartMPUTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartMPUTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartMPUTaskRequest
//
// @return StartMPUTaskResponse
func (client *Client) StartMPUTask(request *StartMPUTaskRequest) (_result *StartMPUTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartMPUTaskResponse{}
	_body, _err := client.StartMPUTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRecordTaskResponse
func (client *Client) StartRecordTaskWithOptions(request *StartRecordTaskRequest, runtime *dara.RuntimeOptions) (_result *StartRecordTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MixMode) {
		query["MixMode"] = request.MixMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.SubSpecAudioUsers) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !dara.IsNil(request.SubSpecCameraUsers) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !dara.IsNil(request.SubSpecShareScreenUsers) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !dara.IsNil(request.SubSpecUsers) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskProfile) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UnsubSpecAudioUsers) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !dara.IsNil(request.UnsubSpecCameraUsers) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !dara.IsNil(request.UnsubSpecShareScreenUsers) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !dara.IsNil(request.UserPanes) {
		query["UserPanes"] = request.UserPanes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRecordTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRecordTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartRecordTaskRequest
//
// @return StartRecordTaskResponse
func (client *Client) StartRecordTask(request *StartRecordTaskRequest) (_result *StartRecordTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartRecordTaskResponse{}
	_body, _err := client.StartRecordTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # StartStreamingOut
//
// @param tmpReq - StartStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartStreamingOutResponse
func (client *Client) StartStreamingOutWithOptions(tmpReq *StartStreamingOutRequest, runtime *dara.RuntimeOptions) (_result *StartStreamingOutResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartStreamingOutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LayoutSpecifiedUsers) {
		request.LayoutSpecifiedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LayoutSpecifiedUsers, dara.String("LayoutSpecifiedUsers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Annotation) {
		query["Annotation"] = request.Annotation
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.LayoutSpecifiedUsersShrink) {
		query["LayoutSpecifiedUsers"] = request.LayoutSpecifiedUsersShrink
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	if !dara.IsNil(request.RegionColor) {
		query["RegionColor"] = request.RegionColor
	}

	if !dara.IsNil(request.ReservePaneForNoCameraUser) {
		query["ReservePaneForNoCameraUser"] = request.ReservePaneForNoCameraUser
	}

	if !dara.IsNil(request.ShowDefaultBackgroundOnMute) {
		query["ShowDefaultBackgroundOnMute"] = request.ShowDefaultBackgroundOnMute
	}

	if !dara.IsNil(request.SpecMixedUserList) {
		query["SpecMixedUserList"] = request.SpecMixedUserList
	}

	if !dara.IsNil(request.StartWithoutChannel) {
		query["StartWithoutChannel"] = request.StartWithoutChannel
	}

	if !dara.IsNil(request.StartWithoutChannelWaitTime) {
		query["StartWithoutChannelWaitTime"] = request.StartWithoutChannelWaitTime
	}

	if !dara.IsNil(request.SubHighResolutionStream) {
		query["SubHighResolutionStream"] = request.SubHighResolutionStream
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Texts) {
		query["Texts"] = request.Texts
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartStreamingOut"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartStreamingOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StartStreamingOut
//
// @param request - StartStreamingOutRequest
//
// @return StartStreamingOutResponse
func (client *Client) StartStreamingOut(request *StartStreamingOutRequest) (_result *StartStreamingOutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartStreamingOutResponse{}
	_body, _err := client.StartStreamingOutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止AI Agent
//
// @param request - StopAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAgentResponse
func (client *Client) StopAgentWithOptions(request *StopAgentRequest, runtime *dara.RuntimeOptions) (_result *StopAgentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止AI Agent
//
// @param request - StopAgentRequest
//
// @return StopAgentResponse
func (client *Client) StopAgent(request *StopAgentRequest) (_result *StopAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopAgentResponse{}
	_body, _err := client.StopAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭某个事件回调
//
// @param tmpReq - StopCategoryCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCategoryCallbackResponse
func (client *Client) StopCategoryCallbackWithOptions(tmpReq *StopCategoryCallbackRequest, runtime *dara.RuntimeOptions) (_result *StopCategoryCallbackResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StopCategoryCallbackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Callback) {
		request.CallbackShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Callback, dara.String("Callback"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallbackShrink) {
		query["Callback"] = request.CallbackShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCategoryCallback"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCategoryCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭某个事件回调
//
// @param request - StopCategoryCallbackRequest
//
// @return StopCategoryCallbackResponse
func (client *Client) StopCategoryCallback(request *StopCategoryCallbackRequest) (_result *StopCategoryCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopCategoryCallbackResponse{}
	_body, _err := client.StopCategoryCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除频道
//
// @param request - StopChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopChannelResponse
func (client *Client) StopChannelWithOptions(request *StopChannelRequest, runtime *dara.RuntimeOptions) (_result *StopChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopChannel"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除频道
//
// @param request - StopChannelRequest
//
// @return StopChannelResponse
func (client *Client) StopChannel(request *StopChannelRequest) (_result *StopChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopChannelResponse{}
	_body, _err := client.StopChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止智能纪要
//
// @param request - StopCloudNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudNoteResponse
func (client *Client) StopCloudNoteWithOptions(request *StopCloudNoteRequest, runtime *dara.RuntimeOptions) (_result *StopCloudNoteResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCloudNote"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCloudNoteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止智能纪要
//
// @param request - StopCloudNoteRequest
//
// @return StopCloudNoteResponse
func (client *Client) StopCloudNote(request *StopCloudNoteRequest) (_result *StopCloudNoteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopCloudNoteResponse{}
	_body, _err := client.StopCloudNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # StopCloudRecord
//
// @param request - StopCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudRecordResponse
func (client *Client) StopCloudRecordWithOptions(request *StopCloudRecordRequest, runtime *dara.RuntimeOptions) (_result *StopCloudRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCloudRecord"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StopCloudRecord
//
// @param request - StopCloudRecordRequest
//
// @return StopCloudRecordResponse
func (client *Client) StopCloudRecord(request *StopCloudRecordRequest) (_result *StopCloudRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.StopCloudRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StopMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMPUTaskResponse
func (client *Client) StopMPUTaskWithOptions(request *StopMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *StopMPUTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopMPUTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopMPUTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopMPUTaskRequest
//
// @return StopMPUTaskResponse
func (client *Client) StopMPUTask(request *StopMPUTaskRequest) (_result *StopMPUTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopMPUTaskResponse{}
	_body, _err := client.StopMPUTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StopRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRecordTaskResponse
func (client *Client) StopRecordTaskWithOptions(request *StopRecordTaskRequest, runtime *dara.RuntimeOptions) (_result *StopRecordTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRecordTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRecordTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopRecordTaskRequest
//
// @return StopRecordTaskResponse
func (client *Client) StopRecordTask(request *StopRecordTaskRequest) (_result *StopRecordTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopRecordTaskResponse{}
	_body, _err := client.StopRecordTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # StopStreamingOut
//
// @param request - StopStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopStreamingOutResponse
func (client *Client) StopStreamingOutWithOptions(request *StopStreamingOutRequest, runtime *dara.RuntimeOptions) (_result *StopStreamingOutResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopStreamingOut"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopStreamingOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StopStreamingOut
//
// @param request - StopStreamingOutRequest
//
// @return StopStreamingOutResponse
func (client *Client) StopStreamingOut(request *StopStreamingOutRequest) (_result *StopStreamingOutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopStreamingOutResponse{}
	_body, _err := client.StopStreamingOutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新AI Agent
//
// @param tmpReq - UpdateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentResponse
func (client *Client) UpdateAgentWithOptions(tmpReq *UpdateAgentRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VoiceChatConfig) {
		request.VoiceChatConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VoiceChatConfig, dara.String("VoiceChatConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.VoiceChatConfigShrink) {
		query["VoiceChatConfig"] = request.VoiceChatConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgent"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AI Agent
//
// @param request - UpdateAgentRequest
//
// @return UpdateAgentResponse
func (client *Client) UpdateAgent(request *UpdateAgentRequest) (_result *UpdateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAgentResponse{}
	_body, _err := client.UpdateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoLiveStreamRuleResponse
func (client *Client) UpdateAutoLiveStreamRuleWithOptions(request *UpdateAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoLiveStreamRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CallBack) {
		query["CallBack"] = request.CallBack
	}

	if !dara.IsNil(request.ChannelIdPrefixes) {
		query["ChannelIdPrefixes"] = request.ChannelIdPrefixes
	}

	if !dara.IsNil(request.ChannelIds) {
		query["ChannelIds"] = request.ChannelIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayDomain) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAutoLiveStreamRule"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAutoLiveStreamRuleRequest
//
// @return UpdateAutoLiveStreamRuleResponse
func (client *Client) UpdateAutoLiveStreamRule(request *UpdateAutoLiveStreamRuleRequest) (_result *UpdateAutoLiveStreamRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAutoLiveStreamRuleResponse{}
	_body, _err := client.UpdateAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新云端录制任务
//
// @param tmpReq - UpdateCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudRecordResponse
func (client *Client) UpdateCloudRecordWithOptions(tmpReq *UpdateCloudRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudRecordResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateCloudRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LayoutSpecifiedUsers) {
		request.LayoutSpecifiedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LayoutSpecifiedUsers, dara.String("LayoutSpecifiedUsers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.LayoutSpecifiedUsersShrink) {
		query["LayoutSpecifiedUsers"] = request.LayoutSpecifiedUsersShrink
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Texts) {
		query["Texts"] = request.Texts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloudRecord"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloudRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新云端录制任务
//
// @param request - UpdateCloudRecordRequest
//
// @return UpdateCloudRecordResponse
func (client *Client) UpdateCloudRecord(request *UpdateCloudRecordRequest) (_result *UpdateCloudRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloudRecordResponse{}
	_body, _err := client.UpdateCloudRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMPUTaskResponse
func (client *Client) UpdateMPUTaskWithOptions(request *UpdateMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateMPUTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundColor) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MixMode) {
		query["MixMode"] = request.MixMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StreamType) {
		query["StreamType"] = request.StreamType
	}

	if !dara.IsNil(request.SubSpecAudioUsers) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !dara.IsNil(request.SubSpecCameraUsers) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !dara.IsNil(request.SubSpecShareScreenUsers) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !dara.IsNil(request.SubSpecUsers) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.UnsubSpecAudioUsers) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !dara.IsNil(request.UnsubSpecCameraUsers) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !dara.IsNil(request.UnsubSpecShareScreenUsers) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !dara.IsNil(request.UserPanes) {
		query["UserPanes"] = request.UserPanes
	}

	if !dara.IsNil(request.Watermarks) {
		query["Watermarks"] = request.Watermarks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMPUTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMPUTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateMPUTaskRequest
//
// @return UpdateMPUTaskResponse
func (client *Client) UpdateMPUTask(request *UpdateMPUTaskRequest) (_result *UpdateMPUTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMPUTaskResponse{}
	_body, _err := client.UpdateMPUTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordTaskResponse
func (client *Client) UpdateRecordTaskWithOptions(request *UpdateRecordTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SubSpecAudioUsers) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !dara.IsNil(request.SubSpecCameraUsers) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !dara.IsNil(request.SubSpecShareScreenUsers) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !dara.IsNil(request.SubSpecUsers) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskProfile) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UnsubSpecAudioUsers) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !dara.IsNil(request.UnsubSpecCameraUsers) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !dara.IsNil(request.UnsubSpecShareScreenUsers) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !dara.IsNil(request.UserPanes) {
		query["UserPanes"] = request.UserPanes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecordTask"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecordTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateRecordTaskRequest
//
// @return UpdateRecordTaskResponse
func (client *Client) UpdateRecordTask(request *UpdateRecordTaskRequest) (_result *UpdateRecordTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecordTaskResponse{}
	_body, _err := client.UpdateRecordTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordTemplateResponse
func (client *Client) UpdateRecordTemplateWithOptions(request *UpdateRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundColor) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.DelayStopTime) {
		query["DelayStopTime"] = request.DelayStopTime
	}

	if !dara.IsNil(request.EnableM3u8DateTime) {
		query["EnableM3u8DateTime"] = request.EnableM3u8DateTime
	}

	if !dara.IsNil(request.FileSplitInterval) {
		query["FileSplitInterval"] = request.FileSplitInterval
	}

	if !dara.IsNil(request.Formats) {
		query["Formats"] = request.Formats
	}

	if !dara.IsNil(request.HttpCallbackUrl) {
		query["HttpCallbackUrl"] = request.HttpCallbackUrl
	}

	if !dara.IsNil(request.LayoutIds) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !dara.IsNil(request.MediaEncode) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !dara.IsNil(request.MnsQueue) {
		query["MnsQueue"] = request.MnsQueue
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OssFilePrefix) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskProfile) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Watermarks) {
		query["Watermarks"] = request.Watermarks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRecordTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateRecordTemplateRequest
//
// @return UpdateRecordTemplateResponse
func (client *Client) UpdateRecordTemplate(request *UpdateRecordTemplateRequest) (_result *UpdateRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRecordTemplateResponse{}
	_body, _err := client.UpdateRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新旁路推流任务
//
// @param tmpReq - UpdateStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStreamingOutResponse
func (client *Client) UpdateStreamingOutWithOptions(tmpReq *UpdateStreamingOutRequest, runtime *dara.RuntimeOptions) (_result *UpdateStreamingOutResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateStreamingOutShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LayoutSpecifiedUsers) {
		request.LayoutSpecifiedUsersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LayoutSpecifiedUsers, dara.String("LayoutSpecifiedUsers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Backgrounds) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ClockWidgets) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.LayoutSpecifiedUsersShrink) {
		query["LayoutSpecifiedUsers"] = request.LayoutSpecifiedUsersShrink
	}

	if !dara.IsNil(request.Panes) {
		query["Panes"] = request.Panes
	}

	if !dara.IsNil(request.RegionColor) {
		query["RegionColor"] = request.RegionColor
	}

	if !dara.IsNil(request.SpecMixedUserList) {
		query["SpecMixedUserList"] = request.SpecMixedUserList
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Texts) {
		query["Texts"] = request.Texts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStreamingOut"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStreamingOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新旁路推流任务
//
// @param request - UpdateStreamingOutRequest
//
// @return UpdateStreamingOutResponse
func (client *Client) UpdateStreamingOut(request *UpdateStreamingOutRequest) (_result *UpdateStreamingOutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStreamingOutResponse{}
	_body, _err := client.UpdateStreamingOutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
