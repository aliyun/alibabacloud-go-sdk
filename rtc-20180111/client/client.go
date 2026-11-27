// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("rtc")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
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

// Summary:
//
// Adds a recording configuration template.
//
// Description:
//
// # Usage notes
//
// This operation adds a recording configuration template. For example, when you add a recording configuration template, you can configure parameters such as recording specifications and encoding options.
//
// # QPS limits
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - AddRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecordTemplateResponse
func (client *Client) AddRecordTemplateWithOptions(request *AddRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Adds a recording configuration template.
//
// Description:
//
// # Usage notes
//
// This operation adds a recording configuration template. For example, when you add a recording configuration template, you can configure parameters such as recording specifications and encoding options.
//
// # QPS limits
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Creates an application agent template.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second for this API operation. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param tmpReq - CreateAppAgentTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppAgentTemplateResponse
func (client *Client) CreateAppAgentTemplateWithOptions(tmpReq *CreateAppAgentTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppAgentTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppAgentTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentSilenceConfig) {
		request.AgentSilenceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentSilenceConfig, dara.String("AgentSilenceConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AmbientSoundConfig) {
		request.AmbientSoundConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AmbientSoundConfig, dara.String("AmbientSoundConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AsrConfig) {
		request.AsrConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AsrConfig, dara.String("AsrConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BackChannelConfig) {
		request.BackChannelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackChannelConfig, dara.String("BackChannelConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InterruptConfig) {
		request.InterruptConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InterruptConfig, dara.String("InterruptConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LlmConfig) {
		request.LlmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LlmConfig, dara.String("LlmConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TtsConfig) {
		request.TtsConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TtsConfig, dara.String("TtsConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSilenceConfigShrink) {
		query["AgentSilenceConfig"] = request.AgentSilenceConfigShrink
	}

	if !dara.IsNil(request.AmbientSoundConfigShrink) {
		query["AmbientSoundConfig"] = request.AmbientSoundConfigShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AsrConfigShrink) {
		query["AsrConfig"] = request.AsrConfigShrink
	}

	if !dara.IsNil(request.BackChannelConfigShrink) {
		query["BackChannelConfig"] = request.BackChannelConfigShrink
	}

	if !dara.IsNil(request.ChatMode) {
		query["ChatMode"] = request.ChatMode
	}

	if !dara.IsNil(request.EnableVideoUnderstanding) {
		query["EnableVideoUnderstanding"] = request.EnableVideoUnderstanding
	}

	if !dara.IsNil(request.Greeting) {
		query["Greeting"] = request.Greeting
	}

	if !dara.IsNil(request.InterruptConfigShrink) {
		query["InterruptConfig"] = request.InterruptConfigShrink
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

	if !dara.IsNil(request.PreferVideo) {
		query["PreferVideo"] = request.PreferVideo
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
// Creates an application agent template.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second for this API operation. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
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
// Creates a custom layout for an application.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param tmpReq - CreateAppLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppLayoutResponse
func (client *Client) CreateAppLayoutWithOptions(tmpReq *CreateAppLayoutRequest, runtime *dara.RuntimeOptions) (_result *CreateAppLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a custom layout for an application.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
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
// Creates an application recording template.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable rate.
//
// @param tmpReq - CreateAppRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppRecordTemplateResponse
func (client *Client) CreateAppRecordTemplateWithOptions(tmpReq *CreateAppRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates an application recording template.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable rate.
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
// Creates a stream forwarding template for an RTC 3.0 application.
//
// Description:
//
// ### Usage notes
//
// This operation creates a stream forwarding template for an RTC 3.0 application.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
//
// @param tmpReq - CreateAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppStreamingOutTemplateResponse
func (client *Client) CreateAppStreamingOutTemplateWithOptions(tmpReq *CreateAppStreamingOutTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppStreamingOutTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a stream forwarding template for an RTC 3.0 application.
//
// Description:
//
// ### Usage notes
//
// This operation creates a stream forwarding template for an RTC 3.0 application.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
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

// Summary:
//
// Creates an application stream mixing template.
//
// Description:
//
// ### QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation as needed.
//
// @param tmpReq - CreateAppViewTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppViewTemplateResponse
func (client *Client) CreateAppViewTemplateWithOptions(tmpReq *CreateAppViewTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAppViewTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAppViewTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppViewTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppViewTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application stream mixing template.
//
// Description:
//
// ### QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation as needed.
//
// @param request - CreateAppViewTemplateRequest
//
// @return CreateAppViewTemplateResponse
func (client *Client) CreateAppViewTemplate(request *CreateAppViewTemplateRequest) (_result *CreateAppViewTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppViewTemplateResponse{}
	_body, _err := client.CreateAppViewTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls CreateAutoLiveStreamRule to create an automatic single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// This operation is used to create a relay rule.
//
// # QPS limit
//
// The per-user QPS limit of this operation is 1,000 calls per minute. If the limit is exceeded, API calls will be throttled, which may affect your business. Call the operation properly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - CreateAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoLiveStreamRuleResponse
func (client *Client) CreateAutoLiveStreamRuleWithOptions(request *CreateAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Calls CreateAutoLiveStreamRule to create an automatic single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// This operation is used to create a relay rule.
//
// # QPS limit
//
// The per-user QPS limit of this operation is 1,000 calls per minute. If the limit is exceeded, API calls will be throttled, which may affect your business. Call the operation properly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Creates a hotword table for meeting minutes.
//
// @param tmpReq - CreateCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudNotePhrasesResponse
func (client *Client) CreateCloudNotePhrasesWithOptions(tmpReq *CreateCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudNotePhrasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Creates a hotword table for meeting minutes.
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

// Summary:
//
// Calls CreateEventSubscribe to create a callback for subscribing to room messages.
//
// Description:
//
// # Usage notes
//
// This operation is used to create a callback for subscribing to room messages. For example, when you create a callback, you can configure parameters such as the callback URL and event type.
//
// # QPS limit
//
// The per-user QPS limit of this operation is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make sure that you call this operation at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - CreateEventSubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventSubscribeResponse
func (client *Client) CreateEventSubscribeWithOptions(request *CreateEventSubscribeRequest, runtime *dara.RuntimeOptions) (_result *CreateEventSubscribeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Calls CreateEventSubscribe to create a callback for subscribing to room messages.
//
// Description:
//
// # Usage notes
//
// This operation is used to create a callback for subscribing to room messages. For example, when you create a callback, you can configure parameters such as the callback URL and event type.
//
// # QPS limit
//
// The per-user QPS limit of this operation is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make sure that you call this operation at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Creates a bypass live streaming layout.
//
// Description:
//
// # Usage notes
//
// This operation creates a bypass live streaming layout. For example, you can configure parameters such as the number of audio mixes and the main pane.
//
// # QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - CreateMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMPULayoutResponse
func (client *Client) CreateMPULayoutWithOptions(request *CreateMPULayoutRequest, runtime *dara.RuntimeOptions) (_result *CreateMPULayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Creates a bypass live streaming layout.
//
// Description:
//
// # Usage notes
//
// This operation creates a bypass live streaming layout. For example, you can configure parameters such as the number of audio mixes and the main pane.
//
// # QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Deletes an application agent template.
//
// Description:
//
// ### QPS limits
//
// The QPS limit per user for this API operation is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - DeleteAppAgentTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppAgentTemplateResponse
func (client *Client) DeleteAppAgentTemplateWithOptions(request *DeleteAppAgentTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppAgentTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes an application agent template.
//
// Description:
//
// ### QPS limits
//
// The QPS limit per user for this API operation is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
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
// Deletes a custom layout for an application.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit per user for this API operation is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
//
// @param tmpReq - DeleteAppLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppLayoutResponse
func (client *Client) DeleteAppLayoutWithOptions(tmpReq *DeleteAppLayoutRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a custom layout for an application.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit per user for this API operation is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
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
// Deletes an application recording template.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param tmpReq - DeleteAppRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppRecordTemplateResponse
func (client *Client) DeleteAppRecordTemplateWithOptions(tmpReq *DeleteAppRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes an application recording template.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
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
// Deletes a stream ingest template for an RTC 3.0 application.
//
// Description:
//
// ### Usage notes
//
// This operation deletes a stream ingest template for an RTC 3.0 application.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param tmpReq - DeleteAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppStreamingOutTemplateResponse
func (client *Client) DeleteAppStreamingOutTemplateWithOptions(tmpReq *DeleteAppStreamingOutTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppStreamingOutTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a stream ingest template for an RTC 3.0 application.
//
// Description:
//
// ### Usage notes
//
// This operation deletes a stream ingest template for an RTC 3.0 application.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
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

// Summary:
//
// Deletes an application\\"s composite stream template.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for this API is 10 calls per second for each user. If you exceed this limit, rate limiting is triggered. This may affect your business. Call this API at a reasonable rate.
//
// @param tmpReq - DeleteAppViewTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppViewTemplateResponse
func (client *Client) DeleteAppViewTemplateWithOptions(tmpReq *DeleteAppViewTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppViewTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAppViewTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppViewTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppViewTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application\\"s composite stream template.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for this API is 10 calls per second for each user. If you exceed this limit, rate limiting is triggered. This may affect your business. Call this API at a reasonable rate.
//
// @param request - DeleteAppViewTemplateRequest
//
// @return DeleteAppViewTemplateResponse
func (client *Client) DeleteAppViewTemplate(request *DeleteAppViewTemplateRequest) (_result *DeleteAppViewTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppViewTemplateResponse{}
	_body, _err := client.DeleteAppViewTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an automatic bypass single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limits
//
// The QPS limit for a single user is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DeleteAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoLiveStreamRuleResponse
func (client *Client) DeleteAutoLiveStreamRuleWithOptions(request *DeleteAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Deletes an automatic bypass single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limits
//
// The QPS limit for a single user is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Deletes a channel.
//
// Description:
//
// # Usage notes
//
// This operation deletes a channel.
//
// # QPS limit
//
// This operation has no QPS limit per user. Call this operation as needed. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DeleteChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChannelResponse
func (client *Client) DeleteChannelWithOptions(request *DeleteChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Deletes a channel.
//
// Description:
//
// # Usage notes
//
// This operation deletes a channel.
//
// # QPS limit
//
// This operation has no QPS limit per user. Call this operation as needed. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Deletes a hot word table.
//
// @param tmpReq - DeleteCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudNotePhrasesResponse
func (client *Client) DeleteCloudNotePhrasesWithOptions(tmpReq *DeleteCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudNotePhrasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Deletes a hot word table.
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

// Summary:
//
// Deletes a callback for subscribing to room messages by calling DeleteEventSubscribe.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateEventSubscribe](https://help.aliyun.com/document_detail/274447.html) to create a callback for subscribing to room messages.
//
// # QPS limits
//
// The QPS limit of this operation is 1,000 calls per minute per user. If the limit is exceeded, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DeleteEventSubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventSubscribeResponse
func (client *Client) DeleteEventSubscribeWithOptions(request *DeleteEventSubscribeRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventSubscribeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Deletes a callback for subscribing to room messages by calling DeleteEventSubscribe.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateEventSubscribe](https://help.aliyun.com/document_detail/274447.html) to create a callback for subscribing to room messages.
//
// # QPS limits
//
// The QPS limit of this operation is 1,000 calls per minute per user. If the limit is exceeded, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Deletes a bypass live streaming layout.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateMPULayout](https://help.aliyun.com/document_detail/162192.html) to create a bypass live streaming layout.
//
// # QPS limits
//
// The per-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DeleteMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMPULayoutResponse
func (client *Client) DeleteMPULayoutWithOptions(request *DeleteMPULayoutRequest, runtime *dara.RuntimeOptions) (_result *DeleteMPULayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Deletes a bypass live streaming layout.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateMPULayout](https://help.aliyun.com/document_detail/162192.html) to create a bypass live streaming layout.
//
// # QPS limits
//
// The per-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Deletes a recording configuration template.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html) to add a recording configuration template.
//
// # QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DeleteRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecordTemplateResponse
func (client *Client) DeleteRecordTemplateWithOptions(request *DeleteRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Deletes a recording configuration template.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html) to add a recording configuration template.
//
// # QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Lists the event callbacks supported by the system.
//
// Description:
//
// ### QPS limit
//
// This API supports a maximum of 10 calls per second per user. If you exceed this limit, rate limiting is applied to your API calls, which may affect your business operations. Use this API judiciously.
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
// Lists the event callbacks supported by the system.
//
// Description:
//
// ### QPS limit
//
// This API supports a maximum of 10 calls per second per user. If you exceed this limit, rate limiting is applied to your API calls, which may affect your business operations. Use this API judiciously.
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
// Query the status of the application agent feature.
//
// Description:
//
// ### QPS limit
//
// This API operation supports a maximum of 10 queries per second per user. If you exceed this limit, rate limiting is triggered, which may affect your application. Call this API only as needed.
//
// @param request - DescribeAppAgentFunctionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppAgentFunctionStatusResponse
func (client *Client) DescribeAppAgentFunctionStatusWithOptions(request *DescribeAppAgentFunctionStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppAgentFunctionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Query the status of the application agent feature.
//
// Description:
//
// ### QPS limit
//
// This API operation supports a maximum of 10 queries per second per user. If you exceed this limit, rate limiting is triggered, which may affect your application. Call this API only as needed.
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
// Lists agent templates.
//
// Description:
//
// ### Rate limits
//
// The rate limit for this API operation is 10 queries per second (QPS) per user. If you exceed the limit, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeAppAgentTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppAgentTemplatesResponse
func (client *Client) DescribeAppAgentTemplatesWithOptions(request *DescribeAppAgentTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppAgentTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Lists agent templates.
//
// Description:
//
// ### Rate limits
//
// The rate limit for this API operation is 10 queries per second (QPS) per user. If you exceed the limit, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
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
// Queries the callback switch of an application.
//
// Description:
//
// ### QPS limits
//
// The QPS limit for a single user is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the API operation at a reasonable rate.
//
// @param request - DescribeAppCallStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppCallStatusResponse
func (client *Client) DescribeAppCallStatusWithOptions(request *DescribeAppCallStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppCallStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the callback switch of an application.
//
// Description:
//
// ### QPS limits
//
// The QPS limit for a single user is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the API operation at a reasonable rate.
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
// # Obtain the application callback key
//
// Description:
//
// ### Queries per second (QPS) limit
//
// The queries per second (QPS) limit for this API is 10 requests per second per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your business. Please invoke the API appropriately.
//
// @param request - DescribeAppCallbackSecretKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppCallbackSecretKeyResponse
func (client *Client) DescribeAppCallbackSecretKeyWithOptions(request *DescribeAppCallbackSecretKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppCallbackSecretKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// # Obtain the application callback key
//
// Description:
//
// ### Queries per second (QPS) limit
//
// The queries per second (QPS) limit for this API is 10 requests per second per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your business. Please invoke the API appropriately.
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
// Invoke DescribeAppKey to query the AppKey of a specified application.
//
// Description:
//
// ### How-To
//
// Before invoking this API, you must have already created an application in the RTC console.
//
// ### queries per second (QPS) limit
//
// The QPS limit for this API is 200 requests per minute per user. If the limit is exceeded, API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
//
// @param request - DescribeAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppKeyResponse
func (client *Client) DescribeAppKeyWithOptions(request *DescribeAppKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invoke DescribeAppKey to query the AppKey of a specified application.
//
// Description:
//
// ### How-To
//
// Before invoking this API, you must have already created an application in the RTC console.
//
// ### queries per second (QPS) limit
//
// The QPS limit for this API is 200 requests per minute per user. If the limit is exceeded, API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
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
// Query custom application layouts.
//
// Description:
//
// ### QPS limit
//
// This API supports up to 10 queries per second (QPS) per user. Exceeding this limit triggers rate limiting, which may affect your business operations. Use the API judiciously.
//
// @param tmpReq - DescribeAppLayoutsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppLayoutsResponse
func (client *Client) DescribeAppLayoutsWithOptions(tmpReq *DescribeAppLayoutsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppLayoutsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Query custom application layouts.
//
// Description:
//
// ### QPS limit
//
// This API supports up to 10 queries per second (QPS) per user. Exceeding this limit triggers rate limiting, which may affect your business operations. Use the API judiciously.
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
// View the application bypass switch.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for this API is 10 for each user. API calls that exceed this limit are subject to rate limiting, which may impact your business. We recommend that you call this API at a reasonable rate.
//
// @param request - DescribeAppLiveStreamStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppLiveStreamStatusResponse
func (client *Client) DescribeAppLiveStreamStatusWithOptions(request *DescribeAppLiveStreamStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppLiveStreamStatusResponse, _err error) {
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
// View the application bypass switch.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for this API is 10 for each user. API calls that exceed this limit are subject to rate limiting, which may impact your business. We recommend that you call this API at a reasonable rate.
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
// Queries the recording status of an application.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for this operation is 10 calls per second per user. Exceeding this limit triggers rate limiting, which may affect your business. Call this operation only as needed.
//
// @param request - DescribeAppRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppRecordStatusResponse
func (client *Client) DescribeAppRecordStatusWithOptions(request *DescribeAppRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppRecordStatusResponse, _err error) {
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
// Queries the recording status of an application.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for this operation is 10 calls per second per user. Exceeding this limit triggers rate limiting, which may affect your business. Call this operation only as needed.
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
// Queries the list of application recording templates.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user calling this API operation is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this API operation at a reasonable frequency.
//
// @param tmpReq - DescribeAppRecordTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppRecordTemplatesResponse
func (client *Client) DescribeAppRecordTemplatesWithOptions(tmpReq *DescribeAppRecordTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppRecordTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of application recording templates.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user calling this API operation is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this API operation at a reasonable frequency.
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
// Lists recording files.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for this API is 10 requests per second per user. Exceeding this limit triggers rate limiting, which may affect your service. Call this API only when necessary.
//
// @param tmpReq - DescribeAppRecordingFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppRecordingFilesResponse
func (client *Client) DescribeAppRecordingFilesWithOptions(tmpReq *DescribeAppRecordingFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppRecordingFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Lists recording files.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for this API is 10 requests per second per user. Exceeding this limit triggers rate limiting, which may affect your service. Call this API only when necessary.
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
// Query the list of stream ingest templates under an RTC 3.0 application.
//
// Description:
//
// ### How-To
//
// This API is used to view the list of stream ingest templates for an RTC 3.0 application.
//
// ### Queries per second (QPS) limit
//
// The QPS limit for this API is 10 queries per second per user. If this limit is exceeded, API calls will be rate-limited, which may impact your business. Please invoke the API appropriately.
//
// @param tmpReq - DescribeAppStreamingOutTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppStreamingOutTemplatesResponse
func (client *Client) DescribeAppStreamingOutTemplatesWithOptions(tmpReq *DescribeAppStreamingOutTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppStreamingOutTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Query the list of stream ingest templates under an RTC 3.0 application.
//
// Description:
//
// ### How-To
//
// This API is used to view the list of stream ingest templates for an RTC 3.0 application.
//
// ### Queries per second (QPS) limit
//
// The QPS limit for this API is 10 queries per second per user. If this limit is exceeded, API calls will be rate-limited, which may impact your business. Please invoke the API appropriately.
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
// # View the stream merging switch
//
// Description:
//
// ### QPS limits
//
// This API has a QPS limit of 10 queries per second (QPS) per user. If you exceed this limit, rate limiting is triggered, which can affect your business operations. Call the API at a reasonable rate.
//
// @param request - DescribeAppViewStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppViewStatusResponse
func (client *Client) DescribeAppViewStatusWithOptions(request *DescribeAppViewStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppViewStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppViewStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppViewStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View the stream merging switch
//
// Description:
//
// ### QPS limits
//
// This API has a QPS limit of 10 queries per second (QPS) per user. If you exceed this limit, rate limiting is triggered, which can affect your business operations. Call the API at a reasonable rate.
//
// @param request - DescribeAppViewStatusRequest
//
// @return DescribeAppViewStatusResponse
func (client *Client) DescribeAppViewStatus(request *DescribeAppViewStatusRequest) (_result *DescribeAppViewStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppViewStatusResponse{}
	_body, _err := client.DescribeAppViewStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists application compositing templates.
//
// Description:
//
// ### QPS limit
//
// The per-user QPS limit for this operation is 10 calls per second. Exceeding this limit triggers rate limiting, which may affect your business. You should call this operation only as needed.
//
// @param tmpReq - DescribeAppViewTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppViewTemplatesResponse
func (client *Client) DescribeAppViewTemplatesWithOptions(tmpReq *DescribeAppViewTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppViewTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeAppViewTemplatesShrinkRequest{}
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
		Action:      dara.String("DescribeAppViewTemplates"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppViewTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists application compositing templates.
//
// Description:
//
// ### QPS limit
//
// The per-user QPS limit for this operation is 10 calls per second. Exceeding this limit triggers rate limiting, which may affect your business. You should call this operation only as needed.
//
// @param request - DescribeAppViewTemplatesRequest
//
// @return DescribeAppViewTemplatesResponse
func (client *Client) DescribeAppViewTemplates(request *DescribeAppViewTemplatesRequest) (_result *DescribeAppViewTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppViewTemplatesResponse{}
	_body, _err := client.DescribeAppViewTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of applications.
//
// Description:
//
// ### Usage notes
//
// This operation queries the list of applications. For example, you can obtain data such as the application ID, application name, and application type from the query results.
//
// ### QPS limits
//
// The QPS limit for a single user is 200 calls per minute. If this limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
//
// @param request - DescribeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the list of applications.
//
// Description:
//
// ### Usage notes
//
// This operation queries the list of applications. For example, you can obtain data such as the application ID, application name, and application type from the query results.
//
// ### QPS limits
//
// The QPS limit for a single user is 200 calls per minute. If this limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
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

// Summary:
//
// Queries the automatic bypass single-stream relay rules.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limit
//
// The QPS limit for a single user is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DescribeAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoLiveStreamRuleResponse
func (client *Client) DescribeAutoLiveStreamRuleWithOptions(request *DescribeAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Queries the automatic bypass single-stream relay rules.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limit
//
// The QPS limit for a single user is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Queries the details of a single communication session. For example, you can query the communication status and user details.
//
// Description:
//
// ### Usage notes
//
// This operation queries the details of a single communication session. For example, you can query the communication status and user details.
//
// ### QPS limit
//
// The QPS limit per user for this operation is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency.
//
// @param request - DescribeCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallResponse
func (client *Client) DescribeCallWithOptions(request *DescribeCallRequest, runtime *dara.RuntimeOptions) (_result *DescribeCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the details of a single communication session. For example, you can query the communication status and user details.
//
// Description:
//
// ### Usage notes
//
// This operation queries the details of a single communication session. For example, you can query the communication status and user details.
//
// ### QPS limit
//
// The QPS limit per user for this operation is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency.
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
// Invoke DescribeCallList to perform a paged query of communication records created within a specified time range.
//
// Description:
//
// ### How-To
//
// This API is used to query communication records created within a specified time range. For example, you can query information such as call duration, status, and user count.
//
// ### Rate Limiting
//
// The queries per second (QPS) limit for this API is 10 calls per second per user. If this limit is exceeded, API calls will be rate-limited, which may impact your business. Please invoke the API appropriately.
//
// @param request - DescribeCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallListResponse
func (client *Client) DescribeCallListWithOptions(request *DescribeCallListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCallListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invoke DescribeCallList to perform a paged query of communication records created within a specified time range.
//
// Description:
//
// ### How-To
//
// This API is used to query communication records created within a specified time range. For example, you can query information such as call duration, status, and user count.
//
// ### Rate Limiting
//
// The queries per second (QPS) limit for this API is 10 calls per second per user. If this limit is exceeded, API calls will be rate-limited, which may impact your business. Please invoke the API appropriately.
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
// Lists the event callbacks configured for an application.
//
// Description:
//
// ### QPS limit
//
// This API supports up to 10 queries per second (QPS) per user. If you exceed this limit, the system applies rate limiting to your API calls, which may affect your business operations. Plan your API calls accordingly.
//
// @param request - DescribeCallbacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallbacksResponse
func (client *Client) DescribeCallbacksWithOptions(request *DescribeCallbacksRequest, runtime *dara.RuntimeOptions) (_result *DescribeCallbacksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Lists the event callbacks configured for an application.
//
// Description:
//
// ### QPS limit
//
// This API supports up to 10 queries per second (QPS) per user. If you exceed this limit, the system applies rate limiting to your API calls, which may affect your business operations. Plan your API calls accordingly.
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
// Description:
//
// ### How-To
//
// This API is used to query online channel information.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second per user. If the limit is exceeded, Rate Limiting will be applied to your API Call, which may Impact your business. Please invoke the API reasonably.
//
// @param request - DescribeChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelResponse
func (client *Client) DescribeChannelWithOptions(request *DescribeChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Description:
//
// ### How-To
//
// This API is used to query online channel information.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second per user. If the limit is exceeded, Rate Limiting will be applied to your API Call, which may Impact your business. Please invoke the API reasonably.
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
// This API is used to query the user list in a channel.
//
// Description:
//
// ### How-To
//
// This API is used to query the list of users in an online channel.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 requests per second per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your business. Invoke this API appropriately.
//
// @param request - DescribeChannelAllUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAllUsersResponse
func (client *Client) DescribeChannelAllUsersWithOptions(request *DescribeChannelAllUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelAllUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// This API is used to query the user list in a channel.
//
// Description:
//
// ### How-To
//
// This API is used to query the list of users in an online channel.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 requests per second per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your business. Invoke this API appropriately.
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
// Retrieves the regional distribution statistics of a channel.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the regional distribution statistics of a channel.
//
// ### QPS limit
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Plan your API calls accordingly.
//
// @param request - DescribeChannelAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAreaDistributionStatDataResponse
func (client *Client) DescribeChannelAreaDistributionStatDataWithOptions(request *DescribeChannelAreaDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelAreaDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the regional distribution statistics of a channel.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the regional distribution statistics of a channel.
//
// ### QPS limit
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Plan your API calls accordingly.
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
// Queries the channel distribution statistics.
//
// Description:
//
// ### Usage notes
//
// This operation queries the channel distribution statistics. You can query the statistics by dimension, such as operating system type or SDK version.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param request - DescribeChannelDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelDistributionStatDataResponse
func (client *Client) DescribeChannelDistributionStatDataWithOptions(request *DescribeChannelDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the channel distribution statistics.
//
// Description:
//
// ### Usage notes
//
// This operation queries the channel distribution statistics. You can query the statistics by dimension, such as operating system type or SDK version.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
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
// This API is used to query channel overview data.
//
// Description:
//
// ### How-To
//
// This API is used to query channel overview data. For example, you can query basic communication information, Trend Indicators, and other information.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second per User. If the limit is exceeded, Rate Limiting will be applied to API Calls, which may Impact your business. Invoke the API appropriately.
//
// @param request - DescribeChannelOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelOverallDataResponse
func (client *Client) DescribeChannelOverallDataWithOptions(request *DescribeChannelOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelOverallDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// This API is used to query channel overview data.
//
// Description:
//
// ### How-To
//
// This API is used to query channel overview data. For example, you can query basic communication information, Trend Indicators, and other information.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second per User. If the limit is exceeded, Rate Limiting will be applied to API Calls, which may Impact your business. Invoke the API appropriately.
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

// Summary:
//
// Queries the list of online users in a channel without detailed information.
//
// Description:
//
// # Usage notes
//
// This operation queries the list of online users in a channel, excluding detailed user information.
//
// # QPS limit
//
// The QPS limit for a single user is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls appropriately. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DescribeChannelParticipantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelParticipantsResponse
func (client *Client) DescribeChannelParticipantsWithOptions(request *DescribeChannelParticipantsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelParticipantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Queries the list of online users in a channel without detailed information.
//
// Description:
//
// # Usage notes
//
// This operation queries the list of online users in a channel, excluding detailed user information.
//
// # QPS limit
//
// The QPS limit for a single user is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls appropriately. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Retrieves the list of publishing users in a channel in descending order of online duration.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the list of publishing users in a channel in descending order of online duration. For example, you can query information such as the time when a user joined the call and the communication duration from the user list.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param request - DescribeChannelTopPubUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelTopPubUserListResponse
func (client *Client) DescribeChannelTopPubUserListWithOptions(request *DescribeChannelTopPubUserListRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelTopPubUserListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the list of publishing users in a channel in descending order of online duration.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the list of publishing users in a channel in descending order of online duration. For example, you can query information such as the time when a user joined the call and the communication duration from the user list.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
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
// Queries the information of a user in a channel.
//
// Description:
//
// ### Usage notes
//
// This operation queries the information of a user in an online channel.
//
// ### QPS limit
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable rate.
//
// @param request - DescribeChannelUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUserResponse
func (client *Client) DescribeChannelUserWithOptions(request *DescribeChannelUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the information of a user in a channel.
//
// Description:
//
// ### Usage notes
//
// This operation queries the information of a user in an online channel.
//
// ### QPS limit
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable rate.
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
// Queries the user data in the channel overview.
//
// Description:
//
// ### Usage notes
//
// This operation queries the user data in the channel overview. For example, you can query metric data details and overview data.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency.
//
// @param request - DescribeChannelUserMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUserMetricsResponse
func (client *Client) DescribeChannelUserMetricsWithOptions(request *DescribeChannelUserMetricsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUserMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the user data in the channel overview.
//
// Description:
//
// ### Usage notes
//
// This operation queries the user data in the channel overview. For example, you can query metric data details and overview data.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency.
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

// Summary:
//
// Queries the list of real-time online users in a channel.
//
// Description:
//
// # Usage notes
//
// This operation queries the details of online users in a channel. For example, you can query the total number of users in communication mode and live streaming mode.
//
// # QPS limit
//
// The queries per second (QPS) limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DescribeChannelUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUsersResponse
func (client *Client) DescribeChannelUsersWithOptions(request *DescribeChannelUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Queries the list of real-time online users in a channel.
//
// Description:
//
// # Usage notes
//
// This operation queries the details of online users in a channel. For example, you can query the total number of users in communication mode and live streaming mode.
//
// # QPS limit
//
// The queries per second (QPS) limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Queries all active channels under a single AppId.
//
// Description:
//
// ### Usage notes
//
// This operation queries all active channels under a single AppId.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
//
// @param request - DescribeChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelsResponse
func (client *Client) DescribeChannelsWithOptions(request *DescribeChannelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChannelsResponse, _err error) {
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
// Queries all active channels under a single AppId.
//
// Description:
//
// ### Usage notes
//
// This operation queries all active channels under a single AppId.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
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
// Queries the hotword lists for meeting minutes.
//
// @param tmpReq - DescribeCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudNotePhrasesResponse
func (client *Client) DescribeCloudNotePhrasesWithOptions(tmpReq *DescribeCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudNotePhrasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the hotword lists for meeting minutes.
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
// # Meeting Summary List
//
// Description:
//
// ### How-To
//
// This API is used to query the list of meeting summary job files that have generated summary results.
//
// ### Queries Per Second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second (QPS) per user. If this limit is exceeded, API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
//
// @param tmpReq - DescribeCloudNotesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudNotesResponse
func (client *Client) DescribeCloudNotesWithOptions(tmpReq *DescribeCloudNotesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudNotesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// # Meeting Summary List
//
// Description:
//
// ### How-To
//
// This API is used to query the list of meeting summary job files that have generated summary results.
//
// ### Queries Per Second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second (QPS) per user. If this limit is exceeded, API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
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
// Queries the status of a recording task.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit per user for this API operation is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param request - DescribeCloudRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudRecordStatusResponse
func (client *Client) DescribeCloudRecordStatusWithOptions(request *DescribeCloudRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudRecordStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the status of a recording task.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit per user for this API operation is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
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
// Retrieves a list of end-to-end user events.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves a list of end-to-end user events. For example, you can obtain information such as the event name and event type from the event list.
//
// ### QPS limits
//
// The QPS limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
//
// @param request - DescribeEndPointEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndPointEventListResponse
func (client *Client) DescribeEndPointEventListWithOptions(request *DescribeEndPointEventListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndPointEventListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves a list of end-to-end user events.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves a list of end-to-end user events. For example, you can obtain information such as the event name and event type from the event list.
//
// ### QPS limits
//
// The QPS limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
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
// Retrieves end-to-end metric data.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves end-to-end metric data. For example, you can retrieve metric information for subscribers and publishers.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the number of calls per second exceeds the limit, throttling is triggered, which may affect your business. Make API calls at a reasonable frequency.
//
// @param request - DescribeEndPointMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndPointMetricDataResponse
func (client *Client) DescribeEndPointMetricDataWithOptions(request *DescribeEndPointMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeEndPointMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves end-to-end metric data.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves end-to-end metric data. For example, you can retrieve metric information for subscribers and publishers.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the number of calls per second exceeds the limit, throttling is triggered, which may affect your business. Make API calls at a reasonable frequency.
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
// Invoke the DescribeFaultDiagnosisFactorDistributionStat API to obtain the distribution of factors impacting abnormal diagnosis.
//
// Description:
//
// ### How-To
//
// This API is used to retrieve the distribution of factors impacting abnormal diagnosis. When using this API, you can query data from the last 48 hours, with each query covering a time range of no more than 24 hours.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 per user. If this limit is exceeded, API calls will be subject to rate limiting, which may impact your service. Please call the API responsibly.
//
// @param request - DescribeFaultDiagnosisFactorDistributionStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisFactorDistributionStatResponse
func (client *Client) DescribeFaultDiagnosisFactorDistributionStatWithOptions(request *DescribeFaultDiagnosisFactorDistributionStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisFactorDistributionStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Invoke the DescribeFaultDiagnosisFactorDistributionStat API to obtain the distribution of factors impacting abnormal diagnosis.
//
// Description:
//
// ### How-To
//
// This API is used to retrieve the distribution of factors impacting abnormal diagnosis. When using this API, you can query data from the last 48 hours, with each query covering a time range of no more than 24 hours.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 per user. If this limit is exceeded, API calls will be subject to rate limiting, which may impact your service. Please call the API responsibly.
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
// Retrieves the overview data of exception diagnostics.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the overview data of exception diagnostics. When you call this operation, you can query data from the last 48 hours, and the query time range cannot exceed 24 hours.
//
// ### QPS limits
//
// The queries per second (QPS) limit on this operation is 10 per user. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation as needed.
//
// @param request - DescribeFaultDiagnosisOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisOverallDataResponse
func (client *Client) DescribeFaultDiagnosisOverallDataWithOptions(request *DescribeFaultDiagnosisOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisOverallDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the overview data of exception diagnostics.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the overview data of exception diagnostics. When you call this operation, you can query data from the last 48 hours, and the query time range cannot exceed 24 hours.
//
// ### QPS limits
//
// The queries per second (QPS) limit on this operation is 10 per user. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation as needed.
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
// [responses_200_schema_properties_UserDetail_properties_Duration_type]integer
//
// Description:
//
// [responses_200_schema_properties_UserDetail_properties_CreatedTs_description]Call creation time, represented as a UNIX timestamp in seconds.
//
// @param request - DescribeFaultDiagnosisUserDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisUserDetailResponse
func (client *Client) DescribeFaultDiagnosisUserDetailWithOptions(request *DescribeFaultDiagnosisUserDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisUserDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// [responses_200_schema_properties_UserDetail_properties_Duration_type]integer
//
// Description:
//
// [responses_200_schema_properties_UserDetail_properties_CreatedTs_description]Call creation time, represented as a UNIX timestamp in seconds.
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
// Queries the detailed user list for fault diagnosis.
//
// Description:
//
// ### Usage notes
//
// This operation is used to retrieve the detailed user list for fault diagnosis. When you call this operation, you can query data from the last 48 hours, and the query time range cannot exceed 24 hours.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
//
// @param request - DescribeFaultDiagnosisUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisUserListResponse
func (client *Client) DescribeFaultDiagnosisUserListWithOptions(request *DescribeFaultDiagnosisUserListRequest, runtime *dara.RuntimeOptions) (_result *DescribeFaultDiagnosisUserListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the detailed user list for fault diagnosis.
//
// Description:
//
// ### Usage notes
//
// This operation is used to retrieve the detailed user list for fault diagnosis. When you call this operation, you can query data from the last 48 hours, and the query time range cannot exceed 24 hours.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
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

// Summary:
//
// Call DescribeMPULayoutInfoList to retrieve layout parameters for a bypass live streaming task.
//
// Description:
//
// # Usage notes
//
// Before calling this operation, call [CreateMPULayout](https://help.aliyun.com/document_detail/162192.html) to create a bypass live streaming layout.
//
// # QPS limit
//
// The per-user QPS limit for this operation is 10 calls per second. Exceeding this limit triggers rate limiting, which may affect your business. Call this operation appropriately. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DescribeMPULayoutInfoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMPULayoutInfoListResponse
func (client *Client) DescribeMPULayoutInfoListWithOptions(request *DescribeMPULayoutInfoListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMPULayoutInfoListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Call DescribeMPULayoutInfoList to retrieve layout parameters for a bypass live streaming task.
//
// Description:
//
// # Usage notes
//
// Before calling this operation, call [CreateMPULayout](https://help.aliyun.com/document_detail/162192.html) to create a bypass live streaming layout.
//
// # QPS limit
//
// The per-user QPS limit for this operation is 10 calls per second. Exceeding this limit triggers rate limiting, which may affect your business. Call this operation appropriately. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// This API is used to obtain the list of publishing users in a communication session based on the subscriber.
//
// Description:
//
// ### How-To
//
// This API is used to obtain the list of publishing users in a communication session based on the subscriber. For example, you can retrieve detailed information about both the subscriber and the publishing users.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 requests per second per user. If this limit is exceeded, API calls will be subject to rate limiting, which may affect your business. Please invoke the API appropriately.
//
// @param request - DescribePubUserListBySubUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePubUserListBySubUserResponse
func (client *Client) DescribePubUserListBySubUserWithOptions(request *DescribePubUserListBySubUserRequest, runtime *dara.RuntimeOptions) (_result *DescribePubUserListBySubUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// This API is used to obtain the list of publishing users in a communication session based on the subscriber.
//
// Description:
//
// ### How-To
//
// This API is used to obtain the list of publishing users in a communication session based on the subscriber. For example, you can retrieve detailed information about both the subscriber and the publishing users.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 requests per second per user. If this limit is exceeded, API calls will be subject to rate limiting, which may affect your business. Please invoke the API appropriately.
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
// This API is used to obtain the downlink experience quality metrics of a user in a single communication session.
//
// Description:
//
// ### How-To
//
// This API is used to obtain the downlink experience quality metrics of a user in a single communication session. For example, you can retrieve information such as the metric types that impact communication experience and audio/video trend indicators.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 per user. If this limit is exceeded, API calls will be subject to rate limiting, which may affect your business. Please invoke the API appropriately.
//
// @param request - DescribeQoeMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQoeMetricDataResponse
func (client *Client) DescribeQoeMetricDataWithOptions(request *DescribeQoeMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQoeMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// This API is used to obtain the downlink experience quality metrics of a user in a single communication session.
//
// Description:
//
// ### How-To
//
// This API is used to obtain the downlink experience quality metrics of a user in a single communication session. For example, you can retrieve information such as the metric types that impact communication experience and audio/video trend indicators.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 per user. If this limit is exceeded, API calls will be subject to rate limiting, which may affect your business. Please invoke the API appropriately.
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
// Queries the regional distribution data of quality statistics.
//
// Description:
//
// ### Usage notes
//
// This operation queries the regional distribution data of quality statistics. When you call this operation, the following limits apply:
//
// - You can query data within any time range in the last 365 days, excluding the current day.
//
// - If the query time range is less than 24 hours, the data is aggregated by hour. If the query time range is 24 hours or more, the data is aggregated by day.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency.
//
// @param request - DescribeQualityAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityAreaDistributionStatDataResponse
func (client *Client) DescribeQualityAreaDistributionStatDataWithOptions(request *DescribeQualityAreaDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityAreaDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the regional distribution data of quality statistics.
//
// Description:
//
// ### Usage notes
//
// This operation queries the regional distribution data of quality statistics. When you call this operation, the following limits apply:
//
// - You can query data within any time range in the last 365 days, excluding the current day.
//
// - If the query time range is less than 24 hours, the data is aggregated by hour. If the query time range is 24 hours or more, the data is aggregated by day.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency.
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
// This API is used to obtain the distribution data of quality statistics.
//
// Description:
//
// ### How-To
//
// This API is used to obtain the distribution data of quality statistics. The following limits apply when using this API:
//
// - You can query data for any period within the last 365 days (excluding the current day).
//
// - If the query range is less than 24 hours, the data is aggregated by hour. If the query range is greater than or equal to 24 hours, the data is aggregated by day.
//
// ### Queries per second (QPS) limit
//
// The QPS limit for this API is 10 queries per second (QPS) per user. If you exceed this limit, your API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API reasonably.
//
// @param request - DescribeQualityDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityDistributionStatDataResponse
func (client *Client) DescribeQualityDistributionStatDataWithOptions(request *DescribeQualityDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// This API is used to obtain the distribution data of quality statistics.
//
// Description:
//
// ### How-To
//
// This API is used to obtain the distribution data of quality statistics. The following limits apply when using this API:
//
// - You can query data for any period within the last 365 days (excluding the current day).
//
// - If the query range is less than 24 hours, the data is aggregated by hour. If the query range is greater than or equal to 24 hours, the data is aggregated by day.
//
// ### Queries per second (QPS) limit
//
// The QPS limit for this API is 10 queries per second (QPS) per user. If you exceed this limit, your API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API reasonably.
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
// Retrieves the distribution data of operating systems and SDK versions in quality statistics.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the distribution data of operating systems and SDK versions in quality statistics. The following limits apply when you call this operation:
//
// - You can query data within the last 365 days, excluding the current day.
//
// - If the query time range is less than 24 hours, the data is aggregated by hour. If the query time range is 24 hours or more, the data is aggregated by day.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency.
//
// @param request - DescribeQualityOsSdkVersionDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeQualityOsSdkVersionDistributionStatDataWithOptions(request *DescribeQualityOsSdkVersionDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityOsSdkVersionDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Retrieves the distribution data of operating systems and SDK versions in quality statistics.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the distribution data of operating systems and SDK versions in quality statistics. The following limits apply when you call this operation:
//
// - You can query data within the last 365 days, excluding the current day.
//
// - If the query time range is less than 24 hours, the data is aggregated by hour. If the query time range is 24 hours or more, the data is aggregated by day.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency.
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
// This API is used to obtain overview data of quality statistics.
//
// Description:
//
// ### How-To
//
// This API is used to obtain overview data of quality statistics. The following restrictions apply when using this API:
//
// - You can query data for any time range within the last 365 days (excluding the current day).
//
// - If the query range is less than 24 hours, the data is aggregated by hour. If the query range is greater than or equal to 24 hours, the data is aggregated by day.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second per user. If you exceed this limit, your API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
//
// @param request - DescribeQualityOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityOverallDataResponse
func (client *Client) DescribeQualityOverallDataWithOptions(request *DescribeQualityOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeQualityOverallDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// This API is used to obtain overview data of quality statistics.
//
// Description:
//
// ### How-To
//
// This API is used to obtain overview data of quality statistics. The following restrictions apply when using this API:
//
// - You can query data for any time range within the last 365 days (excluding the current day).
//
// - If the query range is less than 24 hours, the data is aggregated by hour. If the query range is greater than or equal to 24 hours, the data is aggregated by day.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second per user. If you exceed this limit, your API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
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

// Summary:
//
// Calls DescribeRecordFiles to query the list of files generated by recording.
//
// Description:
//
// # Usage notes
//
// This operation is used to query the list of files generated by recording. For example, you can query parameters such as the creation time and recording duration of files generated by recording.
//
// # QPS limit
//
// The per-user QPS limit of this operation is 10 calls per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call this operation reasonably. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DescribeRecordFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordFilesResponse
func (client *Client) DescribeRecordFilesWithOptions(request *DescribeRecordFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Calls DescribeRecordFiles to query the list of files generated by recording.
//
// Description:
//
// # Usage notes
//
// This operation is used to query the list of files generated by recording. For example, you can query parameters such as the creation time and recording duration of files generated by recording.
//
// # QPS limit
//
// The per-user QPS limit of this operation is 10 calls per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Please call this operation reasonably. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Queries a list of recording template configurations.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html) to add a recording template.
//
// # QPS limit
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DescribeRecordTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordTemplatesResponse
func (client *Client) DescribeRecordTemplatesWithOptions(request *DescribeRecordTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Queries a list of recording template configurations.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html) to add a recording template.
//
// # QPS limit
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Queries the list of channel communication records.
//
// @param request - DescribeRtcChannelListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcChannelListResponse
func (client *Client) DescribeRtcChannelListWithOptions(request *DescribeRtcChannelListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcChannelListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Queries the list of channel communication records.
//
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

// Summary:
//
// Queries the details of communication records in a channel. You can query data within the last 180 days.
//
// @param request - DescribeRtcChannelMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcChannelMetricResponse
func (client *Client) DescribeRtcChannelMetricWithOptions(request *DescribeRtcChannelMetricRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcChannelMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Queries the details of communication records in a channel. You can query data within the last 180 days.
//
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

// Summary:
//
// Retrieves the total communication duration for an application within a specified time range. Data is aggregated by audio and video specifications.
//
// Description:
//
// > The maximum query time range is 30 days.
//
// @param request - DescribeRtcDurationDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcDurationDataResponse
func (client *Client) DescribeRtcDurationDataWithOptions(request *DescribeRtcDurationDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcDurationDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Retrieves the total communication duration for an application within a specified time range. Data is aggregated by audio and video specifications.
//
// Description:
//
// > The maximum query time range is 30 days.
//
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

// Summary:
//
// Queries the peak number of concurrent channels of an application in a specified time range.
//
// Description:
//
// > The maximum time range that can be queried is 30 days.
//
// @param request - DescribeRtcPeakChannelCntDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcPeakChannelCntDataResponse
func (client *Client) DescribeRtcPeakChannelCntDataWithOptions(request *DescribeRtcPeakChannelCntDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcPeakChannelCntDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Queries the peak number of concurrent channels of an application in a specified time range.
//
// Description:
//
// > The maximum time range that can be queried is 30 days.
//
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

// Summary:
//
// Queries the number of active users of an application within a specified time range, that is, the number of user terminals that engaged in communication.
//
// Description:
//
// > The maximum time span between the start time and the end time of a data query is 30 days.
//
// @param request - DescribeRtcUserCntDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcUserCntDataResponse
func (client *Client) DescribeRtcUserCntDataWithOptions(request *DescribeRtcUserCntDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcUserCntDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Queries the number of active users of an application within a specified time range, that is, the number of user terminals that engaged in communication.
//
// Description:
//
// > The maximum time span between the start time and the end time of a data query is 30 days.
//
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
// Queries the bypass streaming status.
//
// Description:
//
// ### QPS limit
//
// The QPS limit per user for this API operation is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - DescribeStreamingOutStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStreamingOutStatusResponse
func (client *Client) DescribeStreamingOutStatusWithOptions(request *DescribeStreamingOutStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeStreamingOutStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the bypass streaming status.
//
// Description:
//
// ### QPS limit
//
// The QPS limit per user for this API operation is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
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
// # Built-in system layouts
//
// Description:
//
// ### QPS limit
//
// The per-user QPS limit for this API is 10 calls per second. Exceeding this limit triggers rate limiting, which may affect your business. Call the API responsibly.
//
// @param request - DescribeSystemLayoutListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemLayoutListResponse
func (client *Client) DescribeSystemLayoutListWithOptions(request *DescribeSystemLayoutListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemLayoutListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// # Built-in system layouts
//
// Description:
//
// ### QPS limit
//
// The per-user QPS limit for this API is 10 calls per second. Exceeding this limit triggers rate limiting, which may affect your business. Call the API responsibly.
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
// This API is used to retrieve the regional distribution data of usage statistics.
//
// Description:
//
// ### How-To
//
// This API is used to retrieve the regional distribution data of usage statistics. The following restrictions apply when using this API:
//
// - You can query data for any period within the last 365 days (excluding the current day).
//
// - If the query range is less than 24 hours, the data is aggregated by hour. If the query range is greater than or equal to 24 hours, the data is aggregated by day.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 requests per second per user. If you exceed this limit, your API calls will be subject to rate limiting, which may impact your business. Please invoke the API appropriately.
//
// @param request - DescribeUsageAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageAreaDistributionStatDataResponse
func (client *Client) DescribeUsageAreaDistributionStatDataWithOptions(request *DescribeUsageAreaDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageAreaDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// This API is used to retrieve the regional distribution data of usage statistics.
//
// Description:
//
// ### How-To
//
// This API is used to retrieve the regional distribution data of usage statistics. The following restrictions apply when using this API:
//
// - You can query data for any period within the last 365 days (excluding the current day).
//
// - If the query range is less than 24 hours, the data is aggregated by hour. If the query range is greater than or equal to 24 hours, the data is aggregated by day.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 requests per second per user. If you exceed this limit, your API calls will be subject to rate limiting, which may impact your business. Please invoke the API appropriately.
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
// Queries the distribution data of usage statistics.
//
// Description:
//
// ### Usage notes
//
// This operation queries the distribution data of usage statistics. When you call this operation, the following limits apply:
//
// - You can query data within any time range in the last 365 days, excluding the current day.
//
// - If the query time range is less than 24 hours, the data is aggregated by hour. If the query time range is 24 hours or more, the data is aggregated by day.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate.
//
// @param request - DescribeUsageDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageDistributionStatDataResponse
func (client *Client) DescribeUsageDistributionStatDataWithOptions(request *DescribeUsageDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the distribution data of usage statistics.
//
// Description:
//
// ### Usage notes
//
// This operation queries the distribution data of usage statistics. When you call this operation, the following limits apply:
//
// - You can query data within any time range in the last 365 days, excluding the current day.
//
// - If the query time range is less than 24 hours, the data is aggregated by hour. If the query time range is 24 hours or more, the data is aggregated by day.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate.
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
// This API is used to obtain the distribution data of operating systems and SDK versions in usage statistics.
//
// Description:
//
// ### How-To
//
// This API is used to obtain the distribution data of operating systems and SDK versions in usage statistics. When using this API, the following restrictions apply:
//
// - You can query data for any period within the last 365 days (excluding the current day).
//
// - If the query range is less than 24 hours, the data is aggregated by hour. If the query range is greater than or equal to 24 hours, the data is aggregated by day.
//
// ### Rate Limiting
//
// The queries per second (QPS) limit for this API is 10 requests per second per user. If you exceed this limit, your API calls will be rate-limited, which may impact your business. Please invoke the API appropriately.
//
// @param request - DescribeUsageOsSdkVersionDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeUsageOsSdkVersionDistributionStatDataWithOptions(request *DescribeUsageOsSdkVersionDistributionStatDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageOsSdkVersionDistributionStatDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// This API is used to obtain the distribution data of operating systems and SDK versions in usage statistics.
//
// Description:
//
// ### How-To
//
// This API is used to obtain the distribution data of operating systems and SDK versions in usage statistics. When using this API, the following restrictions apply:
//
// - You can query data for any period within the last 365 days (excluding the current day).
//
// - If the query range is less than 24 hours, the data is aggregated by hour. If the query range is greater than or equal to 24 hours, the data is aggregated by day.
//
// ### Rate Limiting
//
// The queries per second (QPS) limit for this API is 10 requests per second per user. If you exceed this limit, your API calls will be rate-limited, which may impact your business. Please invoke the API appropriately.
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
// Queries the overview data of usage statistics.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the overview data of usage statistics. The following limits apply when you call this operation:
//
// - You can query data for any time range within the last 365 days, excluding the current day.
//
// - If the query time range is less than 24 hours, the data is aggregated by hour. If the query time range is 24 hours or more, the data is aggregated by day.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate.
//
// @param request - DescribeUsageOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageOverallDataResponse
func (client *Client) DescribeUsageOverallDataWithOptions(request *DescribeUsageOverallDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeUsageOverallDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Queries the overview data of usage statistics.
//
// Description:
//
// ### Usage notes
//
// This operation retrieves the overview data of usage statistics. The following limits apply when you call this operation:
//
// - You can query data for any time range within the last 365 days, excluding the current day.
//
// - If the query time range is less than 24 hours, the data is aggregated by hour. If the query time range is 24 hours or more, the data is aggregated by day.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate.
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

// Summary:
//
// Calls DescribeUserInfoInChannel to query the real-time information of a user in a channel.
//
// Description:
//
// # Usage notes
//
// This operation is used to query the real-time information of a user in a channel. For example, you can query information such as the user role type and the time of joining the meeting.
//
// # QPS limit
//
// The per-user QPS limit of this operation is 100 times per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Call this operation properly. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DescribeUserInfoInChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserInfoInChannelResponse
func (client *Client) DescribeUserInfoInChannelWithOptions(request *DescribeUserInfoInChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserInfoInChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Calls DescribeUserInfoInChannel to query the real-time information of a user in a channel.
//
// Description:
//
// # Usage notes
//
// This operation is used to query the real-time information of a user in a channel. For example, you can query information such as the user role type and the time of joining the meeting.
//
// # QPS limit
//
// The per-user QPS limit of this operation is 100 times per second. If the limit is exceeded, API calls will be throttled, which may affect your business. Call this operation properly. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Disables the automatic bypass single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limit
//
// The QPS limit per user is 1,000 requests per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - DisableAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAutoLiveStreamRuleResponse
func (client *Client) DisableAutoLiveStreamRuleWithOptions(request *DisableAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Disables the automatic bypass single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limit
//
// The QPS limit per user is 1,000 requests per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Enables the automatic bypass single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limit
//
// The QPS limit per user for this operation is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - EnableAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAutoLiveStreamRuleResponse
func (client *Client) EnableAutoLiveStreamRuleWithOptions(request *EnableAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Enables the automatic bypass single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limit
//
// The QPS limit per user for this operation is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// GetAgent.
//
// @param request - GetAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithOptions(request *GetAgentRequest, runtime *dara.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// GetAgent.
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

// Summary:
//
// Retrieves the status of a bypass transcoding task.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [StartMpuTask](https://help.aliyun.com/document_detail/93183.html) to start a bypass transcoding task.
//
// # QPS limit
//
// The QPS limit for a single user is 4,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - GetMPUTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMPUTaskStatusResponse
func (client *Client) GetMPUTaskStatusWithOptions(request *GetMPUTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetMPUTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Retrieves the status of a bypass transcoding task.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [StartMpuTask](https://help.aliyun.com/document_detail/93183.html) to start a bypass transcoding task.
//
// # QPS limit
//
// The QPS limit for a single user is 4,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Modifies the name of a specified application.
//
// Description:
//
// #### Usage notes
//
// This operation modifies the name of a specified application.
//
// #### QPS limit
//
// The queries per second (QPS) limit per user for this operation is 200 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modifies the name of a specified application.
//
// Description:
//
// #### Usage notes
//
// This operation modifies the name of a specified application.
//
// #### QPS limit
//
// The queries per second (QPS) limit per user for this operation is 200 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
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
// Modifies the agent feature switch for an application.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit per user for this API operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - ModifyAppAgentFunctionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppAgentFunctionStatusResponse
func (client *Client) ModifyAppAgentFunctionStatusWithOptions(request *ModifyAppAgentFunctionStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppAgentFunctionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modifies the agent feature switch for an application.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit per user for this API operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
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
// Updates an application agent template.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the API at a reasonable rate.
//
// @param tmpReq - ModifyAppAgentTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppAgentTemplateResponse
func (client *Client) ModifyAppAgentTemplateWithOptions(tmpReq *ModifyAppAgentTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppAgentTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppAgentTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentSilenceConfig) {
		request.AgentSilenceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentSilenceConfig, dara.String("AgentSilenceConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AmbientSoundConfig) {
		request.AmbientSoundConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AmbientSoundConfig, dara.String("AmbientSoundConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AsrConfig) {
		request.AsrConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AsrConfig, dara.String("AsrConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BackChannelConfig) {
		request.BackChannelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackChannelConfig, dara.String("BackChannelConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InterruptConfig) {
		request.InterruptConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InterruptConfig, dara.String("InterruptConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LlmConfig) {
		request.LlmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LlmConfig, dara.String("LlmConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TtsConfig) {
		request.TtsConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TtsConfig, dara.String("TtsConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSilenceConfigShrink) {
		query["AgentSilenceConfig"] = request.AgentSilenceConfigShrink
	}

	if !dara.IsNil(request.AmbientSoundConfigShrink) {
		query["AmbientSoundConfig"] = request.AmbientSoundConfigShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AsrConfigShrink) {
		query["AsrConfig"] = request.AsrConfigShrink
	}

	if !dara.IsNil(request.BackChannelConfigShrink) {
		query["BackChannelConfig"] = request.BackChannelConfigShrink
	}

	if !dara.IsNil(request.ChatMode) {
		query["ChatMode"] = request.ChatMode
	}

	if !dara.IsNil(request.EnableVideoUnderstanding) {
		query["EnableVideoUnderstanding"] = request.EnableVideoUnderstanding
	}

	if !dara.IsNil(request.Greeting) {
		query["Greeting"] = request.Greeting
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InterruptConfigShrink) {
		query["InterruptConfig"] = request.InterruptConfigShrink
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

	if !dara.IsNil(request.PreferVideo) {
		query["PreferVideo"] = request.PreferVideo
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
// Updates an application agent template.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the API at a reasonable rate.
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
// Updates the callback event switch for an application.
//
// Description:
//
// ### QPS limits
//
// The QPS limit per user for this API operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate.
//
// @param request - ModifyAppCallbackStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppCallbackStatusResponse
func (client *Client) ModifyAppCallbackStatusWithOptions(request *ModifyAppCallbackStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppCallbackStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates the callback event switch for an application.
//
// Description:
//
// ### QPS limits
//
// The QPS limit per user for this API operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate.
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
// Modify a custom application layout.
//
// Description:
//
// ### QPS limits
//
// This API supports a maximum of 10 calls per second per user. Exceeding this limit triggers rate limiting, which may impact your business operations. Use the API judiciously.
//
// @param tmpReq - ModifyAppLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppLayoutResponse
func (client *Client) ModifyAppLayoutWithOptions(tmpReq *ModifyAppLayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modify a custom application layout.
//
// Description:
//
// ### QPS limits
//
// This API supports a maximum of 10 calls per second per user. Exceeding this limit triggers rate limiting, which may impact your business operations. Use the API judiciously.
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
// # Modify the bypass toggle of an application
//
// Description:
//
// ### Queries per second (QPS) limit
//
// The queries per second (QPS) limit for this API is 10 per user per second. If the limit is exceeded, API calls will be subject to Rate Limiting, which may impact your business. Please invoke the API appropriately.
//
// @param request - ModifyAppLiveStreamStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppLiveStreamStatusResponse
func (client *Client) ModifyAppLiveStreamStatusWithOptions(request *ModifyAppLiveStreamStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppLiveStreamStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// # Modify the bypass toggle of an application
//
// Description:
//
// ### Queries per second (QPS) limit
//
// The queries per second (QPS) limit for this API is 10 per user per second. If the limit is exceeded, API calls will be subject to Rate Limiting, which may impact your business. Please invoke the API appropriately.
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
// Modifies the recording switch of an application.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit per user for this API operation is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable rate.
//
// @param request - ModifyAppRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppRecordStatusResponse
func (client *Client) ModifyAppRecordStatusWithOptions(request *ModifyAppRecordStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppRecordStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modifies the recording switch of an application.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit per user for this API operation is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable rate.
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
// Modify an application recording template.
//
// Description:
//
// ### QPS limit
//
// This API supports up to 10 calls per second per user. If you exceed this limit, rate limiting is triggered, which might affect your service. Call the API at a reasonable frequency.
//
// @param tmpReq - ModifyAppRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppRecordTemplateResponse
func (client *Client) ModifyAppRecordTemplateWithOptions(tmpReq *ModifyAppRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modify an application recording template.
//
// Description:
//
// ### QPS limit
//
// This API supports up to 10 calls per second per user. If you exceed this limit, rate limiting is triggered, which might affect your service. Call the API at a reasonable frequency.
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
// Modify the stream ingest template for an RTC 3.0 application.
//
// Description:
//
// ### How-To
//
// This API modifies the stream ingest template for an RTC 3.0 application.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 per user. If you exceed this limit, your API calls will be subject to rate limiting, which may impact your business. Please invoke the API appropriately.
//
// @param tmpReq - ModifyAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppStreamingOutTemplateResponse
func (client *Client) ModifyAppStreamingOutTemplateWithOptions(tmpReq *ModifyAppStreamingOutTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppStreamingOutTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Modify the stream ingest template for an RTC 3.0 application.
//
// Description:
//
// ### How-To
//
// This API modifies the stream ingest template for an RTC 3.0 application.
//
// ### QPS Limit
//
// The queries per second (QPS) limit for this API is 10 per user. If you exceed this limit, your API calls will be subject to rate limiting, which may impact your business. Please invoke the API appropriately.
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
// Modifies the merging switch.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param request - ModifyAppViewStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppViewStatusResponse
func (client *Client) ModifyAppViewStatusWithOptions(request *ModifyAppViewStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppViewStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppViewStatus"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppViewStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the merging switch.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param request - ModifyAppViewStatusRequest
//
// @return ModifyAppViewStatusResponse
func (client *Client) ModifyAppViewStatus(request *ModifyAppViewStatusRequest) (_result *ModifyAppViewStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppViewStatusResponse{}
	_body, _err := client.ModifyAppViewStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update an application view template.
//
// Description:
//
// ### QPS limits
//
// This API supports up to 10 calls per second per user. Exceeding this limit triggers rate limiting, which may affect your business. Use this API with caution.
//
// @param tmpReq - ModifyAppViewTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppViewTemplateResponse
func (client *Client) ModifyAppViewTemplateWithOptions(tmpReq *ModifyAppViewTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppViewTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAppViewTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Template) {
		request.TemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Template, dara.String("Template"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.TemplateShrink) {
		query["Template"] = request.TemplateShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAppViewTemplate"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppViewTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update an application view template.
//
// Description:
//
// ### QPS limits
//
// This API supports up to 10 calls per second per user. Exceeding this limit triggers rate limiting, which may affect your business. Use this API with caution.
//
// @param request - ModifyAppViewTemplateRequest
//
// @return ModifyAppViewTemplateResponse
func (client *Client) ModifyAppViewTemplate(request *ModifyAppViewTemplateRequest) (_result *ModifyAppViewTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppViewTemplateResponse{}
	_body, _err := client.ModifyAppViewTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the callback configuration of an application.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the API at a reasonable frequency.
//
// @param tmpReq - ModifyCallbackMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCallbackMetaResponse
func (client *Client) ModifyCallbackMetaWithOptions(tmpReq *ModifyCallbackMetaRequest, runtime *dara.RuntimeOptions) (_result *ModifyCallbackMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates the callback configuration of an application.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the API at a reasonable frequency.
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
// Updates the hotwords table for meeting minutes.
//
// @param tmpReq - ModifyCloudNotePhrasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudNotePhrasesResponse
func (client *Client) ModifyCloudNotePhrasesWithOptions(tmpReq *ModifyCloudNotePhrasesRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudNotePhrasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates the hotwords table for meeting minutes.
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

// Summary:
//
// Modifies a bypass live streaming layout.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateMPULayout](https://help.aliyun.com/document_detail/162192.html) to create a bypass live streaming layout.
//
// # QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - ModifyMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMPULayoutResponse
func (client *Client) ModifyMPULayoutWithOptions(request *ModifyMPULayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyMPULayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Modifies a bypass live streaming layout.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateMPULayout](https://help.aliyun.com/document_detail/162192.html) to create a bypass live streaming layout.
//
// # QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Sets the stream properties.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param tmpReq - ModifyStreamingPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyStreamingPropertyResponse
func (client *Client) ModifyStreamingPropertyWithOptions(tmpReq *ModifyStreamingPropertyRequest, runtime *dara.RuntimeOptions) (_result *ModifyStreamingPropertyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyStreamingPropertyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ViewSubscribers) {
		request.ViewSubscribersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ViewSubscribers, dara.String("ViewSubscribers"), dara.String("simple"))
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

	if !dara.IsNil(request.ViewContent) {
		query["ViewContent"] = request.ViewContent
	}

	if !dara.IsNil(request.ViewSubscribersShrink) {
		query["ViewSubscribers"] = request.ViewSubscribersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyStreamingProperty"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyStreamingPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the stream properties.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - ModifyStreamingPropertyRequest
//
// @return ModifyStreamingPropertyResponse
func (client *Client) ModifyStreamingProperty(request *ModifyStreamingPropertyRequest) (_result *ModifyStreamingPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyStreamingPropertyResponse{}
	_body, _err := client.ModifyStreamingPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the mixed stream layout.
//
// Description:
//
// ### QPS limit
//
// The QPS limit per user is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your services. Call the API at a reasonable rate.
//
// @param tmpReq - ModifyViewLayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyViewLayoutResponse
func (client *Client) ModifyViewLayoutWithOptions(tmpReq *ModifyViewLayoutRequest, runtime *dara.RuntimeOptions) (_result *ModifyViewLayoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyViewLayoutShrinkRequest{}
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
		Action:      dara.String("ModifyViewLayout"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyViewLayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the mixed stream layout.
//
// Description:
//
// ### QPS limit
//
// The QPS limit per user is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your services. Call the API at a reasonable rate.
//
// @param request - ModifyViewLayoutRequest
//
// @return ModifyViewLayoutResponse
func (client *Client) ModifyViewLayout(request *ModifyViewLayoutRequest) (_result *ModifyViewLayoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyViewLayoutResponse{}
	_body, _err := client.ModifyViewLayoutWithOptions(request, runtime)
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
// @param tmpReq - NotifyAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NotifyAgentResponse
func (client *Client) NotifyAgentWithOptions(tmpReq *NotifyAgentRequest, runtime *dara.RuntimeOptions) (_result *NotifyAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &NotifyAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BackgroundMusic) {
		request.BackgroundMusicShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BackgroundMusic, dara.String("BackgroundMusic"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BackgroundMusicShrink) {
		query["BackgroundMusic"] = request.BackgroundMusicShrink
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

// Summary:
//
// Call the RemoveTerminals operation to remove specified end users from a channel.
//
// Description:
//
// # Usage notes
//
// This API removes specified end users from a channel.
//
// # QPS limit
//
// The queries per second (QPS) limit for this API is 1,000 calls per minute per user. If you exceed the limit, API calls are subject to rate limiting, which may affect your business. Call this API at a reasonable rate. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - RemoveTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTerminalsResponse
func (client *Client) RemoveTerminalsWithOptions(request *RemoveTerminalsRequest, runtime *dara.RuntimeOptions) (_result *RemoveTerminalsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Call the RemoveTerminals operation to remove specified end users from a channel.
//
// Description:
//
// # Usage notes
//
// This API removes specified end users from a channel.
//
// # QPS limit
//
// The queries per second (QPS) limit for this API is 1,000 calls per minute per user. If you exceed the limit, API calls are subject to rate limiting, which may affect your business. Call this API at a reasonable rate. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
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
// Description:
//
// ### How-To
//
// This API is used to kick users out of an online channel.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 requests per second per user. If the limit is exceeded, API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
//
// @param request - RemoveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsersWithOptions(request *RemoveUsersRequest, runtime *dara.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Description:
//
// ### How-To
//
// This API is used to kick users out of an online channel.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 requests per second per user. If the limit is exceeded, API calls will be subject to Rate Limiting, which may Impact your business. Please invoke the API appropriately.
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
// Cancels a SIP invite.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for this API is 10 per user. If you exceed this limit, API calls are subject to rate limiting, which may affect your business. Call this API at a reasonable frequency.
//
// @param request - RtcCancelSipInviteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RtcCancelSipInviteResponse
func (client *Client) RtcCancelSipInviteWithOptions(request *RtcCancelSipInviteRequest, runtime *dara.RuntimeOptions) (_result *RtcCancelSipInviteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RtcCancelSipInvite"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RtcCancelSipInviteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a SIP invite.
//
// Description:
//
// ### QPS limit
//
// The queries per second (QPS) limit for this API is 10 per user. If you exceed this limit, API calls are subject to rate limiting, which may affect your business. Call this API at a reasonable frequency.
//
// @param request - RtcCancelSipInviteRequest
//
// @return RtcCancelSipInviteResponse
func (client *Client) RtcCancelSipInvite(request *RtcCancelSipInviteRequest) (_result *RtcCancelSipInviteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RtcCancelSipInviteResponse{}
	_body, _err := client.RtcCancelSipInviteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends DTMF key signals to a target OneRTC SIP call in a specified channel.
//
// @param request - RtcSendSipDtmfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RtcSendSipDtmfResponse
func (client *Client) RtcSendSipDtmfWithOptions(request *RtcSendSipDtmfRequest, runtime *dara.RuntimeOptions) (_result *RtcSendSipDtmfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ChannelUserId) {
		query["ChannelUserId"] = request.ChannelUserId
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.Digits) {
		query["Digits"] = request.Digits
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RtcSendSipDtmf"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RtcSendSipDtmfResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends DTMF key signals to a target OneRTC SIP call in a specified channel.
//
// @param request - RtcSendSipDtmfRequest
//
// @return RtcSendSipDtmfResponse
func (client *Client) RtcSendSipDtmf(request *RtcSendSipDtmfRequest) (_result *RtcSendSipDtmfResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RtcSendSipDtmfResponse{}
	_body, _err := client.RtcSendSipDtmfWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invites a SIP endpoint to join a channel.
//
// Description:
//
// ### Rate limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - RtcSipInviteMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RtcSipInviteMemberResponse
func (client *Client) RtcSipInviteMemberWithOptions(request *RtcSipInviteMemberRequest, runtime *dara.RuntimeOptions) (_result *RtcSipInviteMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppToken) {
		query["AppToken"] = request.AppToken
	}

	if !dara.IsNil(request.CallNumber) {
		query["CallNumber"] = request.CallNumber
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.Registered) {
		query["Registered"] = request.Registered
	}

	if !dara.IsNil(request.ServerAddress) {
		query["ServerAddress"] = request.ServerAddress
	}

	if !dara.IsNil(request.SipDisplayName) {
		query["SipDisplayName"] = request.SipDisplayName
	}

	if !dara.IsNil(request.SipRoomId) {
		query["SipRoomId"] = request.SipRoomId
	}

	if !dara.IsNil(request.SipUri) {
		query["SipUri"] = request.SipUri
	}

	if !dara.IsNil(request.SipUserAgent) {
		query["SipUserAgent"] = request.SipUserAgent
	}

	if !dara.IsNil(request.SipUserId) {
		query["SipUserId"] = request.SipUserId
	}

	if !dara.IsNil(request.SipUserPassword) {
		query["SipUserPassword"] = request.SipUserPassword
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RtcSipInviteMember"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RtcSipInviteMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invites a SIP endpoint to join a channel.
//
// Description:
//
// ### Rate limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - RtcSipInviteMemberRequest
//
// @return RtcSipInviteMemberResponse
func (client *Client) RtcSipInviteMember(request *RtcSipInviteMemberRequest) (_result *RtcSipInviteMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RtcSipInviteMemberResponse{}
	_body, _err := client.RtcSipInviteMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Mutes or unmutes the microphone for a SIP endpoint.
//
// Description:
//
// ### QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - RtcSipMuteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RtcSipMuteResponse
func (client *Client) RtcSipMuteWithOptions(request *RtcSipMuteRequest, runtime *dara.RuntimeOptions) (_result *RtcSipMuteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Operations) {
		query["Operations"] = request.Operations
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RtcSipMute"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RtcSipMuteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Mutes or unmutes the microphone for a SIP endpoint.
//
// Description:
//
// ### QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this API operation at a reasonable frequency.
//
// @param request - RtcSipMuteRequest
//
// @return RtcSipMuteResponse
func (client *Client) RtcSipMute(request *RtcSipMuteRequest) (_result *RtcSipMuteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RtcSipMuteResponse{}
	_body, _err := client.RtcSipMuteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an AI agent.
//
// @param tmpReq - StartAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAgentResponse
func (client *Client) StartAgentWithOptions(tmpReq *StartAgentRequest, runtime *dara.RuntimeOptions) (_result *StartAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Starts an AI agent.
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
// Starts an event callback.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit for this API is 10 calls per second per user. If you exceed the limit, API calls are subject to rate limiting, which may affect your business. Call the API at a reasonable rate.
//
// @param tmpReq - StartCategoryCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCategoryCallbackResponse
func (client *Client) StartCategoryCallbackWithOptions(tmpReq *StartCategoryCallbackRequest, runtime *dara.RuntimeOptions) (_result *StartCategoryCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Starts an event callback.
//
// Description:
//
// ### QPS limits
//
// The queries per second (QPS) limit for this API is 10 calls per second per user. If you exceed the limit, API calls are subject to rate limiting, which may affect your business. Call the API at a reasonable rate.
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
// Starts an intelligent meeting minutes task.
//
// Description:
//
// ### Usage notes
//
// This operation starts a meeting minutes task for an online channel. If you need meeting minutes and subtitles, enable all features.
//
// ### QPS limits
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable frequency.
//
// > Concurrency limits
//
// > - Only one meeting minutes task can run in the same channel at a time.
//
// > - You can start and stop meeting minutes tasks multiple times in the same channel, but the TaskId must be different each time.
//
// @param tmpReq - StartCloudNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudNoteResponse
func (client *Client) StartCloudNoteWithOptions(tmpReq *StartCloudNoteRequest, runtime *dara.RuntimeOptions) (_result *StartCloudNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Starts an intelligent meeting minutes task.
//
// Description:
//
// ### Usage notes
//
// This operation starts a meeting minutes task for an online channel. If you need meeting minutes and subtitles, enable all features.
//
// ### QPS limits
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable frequency.
//
// > Concurrency limits
//
// > - Only one meeting minutes task can run in the same channel at a time.
//
// > - You can start and stop meeting minutes tasks multiple times in the same channel, but the TaskId must be different each time.
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
// Starts a cloud recording task.
//
// Description:
//
// ### Usage notes
//
// This operation starts a cloud recording task for an online channel.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param tmpReq - StartCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudRecordResponse
func (client *Client) StartCloudRecordWithOptions(tmpReq *StartCloudRecordRequest, runtime *dara.RuntimeOptions) (_result *StartCloudRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Starts a cloud recording task.
//
// Description:
//
// ### Usage notes
//
// This operation starts a cloud recording task for an online channel.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
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

// Summary:
//
// Starts a bypass relaying task.
//
// Description:
//
// # Usage notes
//
// This operation is used to start a bypass relaying task. When you use the bypass relaying service, ApsaraVideo Live has the following limits:
//
// <props="china">- For the concurrency limit on tasks, see [Concurrency limits](https://help.aliyun.com/document_detail/71050.html#section-eaz-anq-a47). If you have special requirements, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex) two weeks in advance.
//
// <props="intl">- For the concurrency limit on tasks, see [Concurrency limits](https://www.alibabacloud.com/help/rtc-61399/latest/60e652#section-eaz-anq-a47). If you have special requirements, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex) two weeks in advance.
//
// - If the transcoding feature is disabled, each account can concurrently push up to 20 original-quality live streams per live streaming accelerated domain name by default. If you have special requirements, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex).
//
// - If the transcoding feature is enabled, each account can concurrently push up to 10 transcoded live streams per live streaming accelerated domain name by default. If you have special requirements, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex).
//
// # QPS limit
//
// The QPS limit for a single user calling this operation is 500 queries per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - StartMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMPUTaskResponse
func (client *Client) StartMPUTaskWithOptions(request *StartMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *StartMPUTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Starts a bypass relaying task.
//
// Description:
//
// # Usage notes
//
// This operation is used to start a bypass relaying task. When you use the bypass relaying service, ApsaraVideo Live has the following limits:
//
// <props="china">- For the concurrency limit on tasks, see [Concurrency limits](https://help.aliyun.com/document_detail/71050.html#section-eaz-anq-a47). If you have special requirements, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex) two weeks in advance.
//
// <props="intl">- For the concurrency limit on tasks, see [Concurrency limits](https://www.alibabacloud.com/help/rtc-61399/latest/60e652#section-eaz-anq-a47). If you have special requirements, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex) two weeks in advance.
//
// - If the transcoding feature is disabled, each account can concurrently push up to 20 original-quality live streams per live streaming accelerated domain name by default. If you have special requirements, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex).
//
// - If the transcoding feature is enabled, each account can concurrently push up to 10 transcoded live streams per live streaming accelerated domain name by default. If you have special requirements, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex).
//
// # QPS limit
//
// The QPS limit for a single user calling this operation is 500 queries per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Starts a video recording task.
//
// Description:
//
// # Usage notes
//
// - Before you call this operation, you must call [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html) to add a recording configuration template.
//
// <props="china">- By default, you can start a maximum of 100 recording tasks at the same time. If the limit is exceeded, new recording tasks become unavailable. If you have special requirements, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex).
//
// <props="intl">- By default, you can start a maximum of 100 recording tasks at the same time. If the limit is exceeded, new recording tasks become unavailable. If you have special requirements, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
// # QPS limit
//
// The QPS limit for a single user is 500 queries per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls properly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - StartRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRecordTaskResponse
func (client *Client) StartRecordTaskWithOptions(request *StartRecordTaskRequest, runtime *dara.RuntimeOptions) (_result *StartRecordTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Starts a video recording task.
//
// Description:
//
// # Usage notes
//
// - Before you call this operation, you must call [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html) to add a recording configuration template.
//
// <props="china">- By default, you can start a maximum of 100 recording tasks at the same time. If the limit is exceeded, new recording tasks become unavailable. If you have special requirements, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex).
//
// <props="intl">- By default, you can start a maximum of 100 recording tasks at the same time. If the limit is exceeded, new recording tasks become unavailable. If you have special requirements, [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
// # QPS limit
//
// The QPS limit for a single user is 500 queries per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls properly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Starts a bypass streaming task.
//
// Description:
//
// ### Usage notes
//
// This operation starts a bypass streaming task for an online channel.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
//
// @param tmpReq - StartStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartStreamingOutResponse
func (client *Client) StartStreamingOutWithOptions(tmpReq *StartStreamingOutRequest, runtime *dara.RuntimeOptions) (_result *StartStreamingOutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Starts a bypass streaming task.
//
// Description:
//
// ### Usage notes
//
// This operation starts a bypass streaming task for an online channel.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable frequency.
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
// Starts stream mixing.
//
// Description:
//
// ### QPS limits
//
// The QPS limit for a single user is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the API operation at a reasonable frequency.
//
// @param tmpReq - StartViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartViewResponse
func (client *Client) StartViewWithOptions(tmpReq *StartViewRequest, runtime *dara.RuntimeOptions) (_result *StartViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartViewShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ViewSubscribers) {
		request.ViewSubscribersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ViewSubscribers, dara.String("ViewSubscribers"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BgColor) {
		query["BgColor"] = request.BgColor
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.CropMode) {
		query["CropMode"] = request.CropMode
	}

	if !dara.IsNil(request.RegionColor) {
		query["RegionColor"] = request.RegionColor
	}

	if !dara.IsNil(request.StartWithoutChannel) {
		query["StartWithoutChannel"] = request.StartWithoutChannel
	}

	if !dara.IsNil(request.StartWithoutChannelWaitTime) {
		query["StartWithoutChannelWaitTime"] = request.StartWithoutChannelWaitTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.ViewContent) {
		query["ViewContent"] = request.ViewContent
	}

	if !dara.IsNil(request.ViewSubscribersShrink) {
		query["ViewSubscribers"] = request.ViewSubscribersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartView"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts stream mixing.
//
// Description:
//
// ### QPS limits
//
// The QPS limit for a single user is 10 queries per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the API operation at a reasonable frequency.
//
// @param request - StartViewRequest
//
// @return StartViewResponse
func (client *Client) StartView(request *StartViewRequest) (_result *StartViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartViewResponse{}
	_body, _err := client.StartViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an AI agent.
//
// @param request - StopAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAgentResponse
func (client *Client) StopAgentWithOptions(request *StopAgentRequest, runtime *dara.RuntimeOptions) (_result *StopAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Stops an AI agent.
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
// Disables an event callback.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
//
// @param tmpReq - StopCategoryCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCategoryCallbackResponse
func (client *Client) StopCategoryCallbackWithOptions(tmpReq *StopCategoryCallbackRequest, runtime *dara.RuntimeOptions) (_result *StopCategoryCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Disables an event callback.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
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
// This API is used by the server to actively shut down a channel.
//
// Description:
//
// ### How-To
//
// This API is used by the server to actively shut down an online channel.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your service. Invoke this API appropriately.
//
// @param request - StopChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopChannelResponse
func (client *Client) StopChannelWithOptions(request *StopChannelRequest, runtime *dara.RuntimeOptions) (_result *StopChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// This API is used by the server to actively shut down a channel.
//
// Description:
//
// ### How-To
//
// This API is used by the server to actively shut down an online channel.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your service. Invoke this API appropriately.
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
// Stops an intelligent meeting minutes task.
//
// Description:
//
// ### Usage notes
//
// This operation stops a meeting minutes task for an online channel.
//
// ### QPS limits
//
// The QPS limit per user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable rate.
//
// @param request - StopCloudNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudNoteResponse
func (client *Client) StopCloudNoteWithOptions(request *StopCloudNoteRequest, runtime *dara.RuntimeOptions) (_result *StopCloudNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Stops an intelligent meeting minutes task.
//
// Description:
//
// ### Usage notes
//
// This operation stops a meeting minutes task for an online channel.
//
// ### QPS limits
//
// The QPS limit per user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable rate.
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
// Description:
//
// ### How-To
//
// This API is used to stop an online channel cloud recording job.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second (QPS) per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your business. Please invoke the API appropriately.
//
// @param request - StopCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudRecordResponse
func (client *Client) StopCloudRecordWithOptions(request *StopCloudRecordRequest, runtime *dara.RuntimeOptions) (_result *StopCloudRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Description:
//
// ### How-To
//
// This API is used to stop an online channel cloud recording job.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second (QPS) per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your business. Please invoke the API appropriately.
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

// Summary:
//
// Calls the StopMPUTask operation to stop a stream relay task.
//
// Description:
//
// # Usage notes
//
// - Before you call this operation, you must have called [StartMpuTask](https://help.aliyun.com/document_detail/93183.html) to start a stream relay task.
//
// - If you need to stop a stream relay task and the stream relay task is abnormal (the StopMPUTask operation is not called to stop the task), the stream relay task is automatically stopped 2 minutes after the last user leaves the channel. After the task is stopped, if you need to resume stream relay, you need to call the [StartMpuTask](https://help.aliyun.com/document_detail/93183.html) operation again.
//
// # QPS limit
//
// The QPS limit per user for this operation is 500 times per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the operation properly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - StopMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMPUTaskResponse
func (client *Client) StopMPUTaskWithOptions(request *StopMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *StopMPUTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Calls the StopMPUTask operation to stop a stream relay task.
//
// Description:
//
// # Usage notes
//
// - Before you call this operation, you must have called [StartMpuTask](https://help.aliyun.com/document_detail/93183.html) to start a stream relay task.
//
// - If you need to stop a stream relay task and the stream relay task is abnormal (the StopMPUTask operation is not called to stop the task), the stream relay task is automatically stopped 2 minutes after the last user leaves the channel. After the task is stopped, if you need to resume stream relay, you need to call the [StartMpuTask](https://help.aliyun.com/document_detail/93183.html) operation again.
//
// # QPS limit
//
// The QPS limit per user for this operation is 500 times per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call the operation properly. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Calls the StopRecordTask operation to stop a video recording task.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call the [StartRecordTask](https://help.aliyun.com/document_detail/193528.html) operation to start a video recording task.
//
// # QPS limit
//
// You can call this operation up to 4,000 times per minute per user. If the number of calls per minute exceeds the limit, throttling is triggered, which may affect your business. Call this operation as required. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - StopRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRecordTaskResponse
func (client *Client) StopRecordTaskWithOptions(request *StopRecordTaskRequest, runtime *dara.RuntimeOptions) (_result *StopRecordTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Calls the StopRecordTask operation to stop a video recording task.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call the [StartRecordTask](https://help.aliyun.com/document_detail/193528.html) operation to start a video recording task.
//
// # QPS limit
//
// You can call this operation up to 4,000 times per minute per user. If the number of calls per minute exceeds the limit, throttling is triggered, which may affect your business. Call this operation as required. For more information, see [QPS limit](https://help.aliyun.com/document_detail/341645.html).
//
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
// Description:
//
// ### How-To
//
// This API is used to stop a bypass stream ingest job for an online channel.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second (QPS) per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your business. Invoke this API appropriately.
//
// @param request - StopStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopStreamingOutResponse
func (client *Client) StopStreamingOutWithOptions(request *StopStreamingOutRequest, runtime *dara.RuntimeOptions) (_result *StopStreamingOutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Description:
//
// ### How-To
//
// This API is used to stop a bypass stream ingest job for an online channel.
//
// ### queries per second (QPS) Limit
//
// The QPS limit for this API is 10 queries per second (QPS) per user. If the limit is exceeded, API calls will be subject to rate limiting, which may impact your business. Invoke this API appropriately.
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
// Stop stream mixing.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for this API is 10 requests per second per user. Exceeding this limit triggers rate limiting, which can affect your application. Therefore, call the API judiciously.
//
// @param request - StopViewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopViewResponse
func (client *Client) StopViewWithOptions(request *StopViewRequest, runtime *dara.RuntimeOptions) (_result *StopViewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
		Action:      dara.String("StopView"),
		Version:     dara.String("2018-01-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stop stream mixing.
//
// Description:
//
// ### QPS limit
//
// The QPS limit for this API is 10 requests per second per user. Exceeding this limit triggers rate limiting, which can affect your application. Therefore, call the API judiciously.
//
// @param request - StopViewRequest
//
// @return StopViewResponse
func (client *Client) StopView(request *StopViewRequest) (_result *StopViewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopViewResponse{}
	_body, _err := client.StopViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update an AI Agent
//
// @param tmpReq - UpdateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentResponse
func (client *Client) UpdateAgentWithOptions(tmpReq *UpdateAgentRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// # Update an AI Agent
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

// Summary:
//
// Modifies an automatic bypass single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limits
//
// The QPS limit for a single user is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - UpdateAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoLiveStreamRuleResponse
func (client *Client) UpdateAutoLiveStreamRuleWithOptions(request *UpdateAutoLiveStreamRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoLiveStreamRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Modifies an automatic bypass single-stream relay rule.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [CreateAutoLiveStreamRule](https://help.aliyun.com/document_detail/353179.html) to create a bypass rule.
//
// # QPS limits
//
// The QPS limit for a single user is 1,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Updates a cloud recording task.
//
// Description:
//
// ### Usage
//
// Updates a cloud recording task for an online channel.
//
// ### QPS limit
//
// This API supports up to 10 requests per second per user. Exceeding this limit triggers rate limiting, which may affect your service. Call the API responsibly.
//
// @param tmpReq - UpdateCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudRecordResponse
func (client *Client) UpdateCloudRecordWithOptions(tmpReq *UpdateCloudRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloudRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates a cloud recording task.
//
// Description:
//
// ### Usage
//
// Updates a cloud recording task for an online channel.
//
// ### QPS limit
//
// This API supports up to 10 requests per second per user. Exceeding this limit triggers rate limiting, which may affect your service. Call the API responsibly.
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

// Summary:
//
// Call UpdateMPUTask to update a bypass relay task.
//
// Description:
//
// # Usage notes
//
// - Before calling this API, you must have called [StartMpuTask](https://help.aliyun.com/document_detail/93183.html) to start a bypass relay task.
//
// - This API must be called while the task is running normally. Calls are invalid when the task has not started, has ended, or is in an abnormal state. You can call [GetMPUTaskStatus](https://help.aliyun.com/document_detail/93181.html) to obtain the task running status.
//
// # QPS limit
//
// The per-user QPS limit of this API is 4,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Please call reasonably. For more information, see [QPS Limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - UpdateMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMPUTaskResponse
func (client *Client) UpdateMPUTaskWithOptions(request *UpdateMPUTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateMPUTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Call UpdateMPUTask to update a bypass relay task.
//
// Description:
//
// # Usage notes
//
// - Before calling this API, you must have called [StartMpuTask](https://help.aliyun.com/document_detail/93183.html) to start a bypass relay task.
//
// - This API must be called while the task is running normally. Calls are invalid when the task has not started, has ended, or is in an abnormal state. You can call [GetMPUTaskStatus](https://help.aliyun.com/document_detail/93181.html) to obtain the task running status.
//
// # QPS limit
//
// The per-user QPS limit of this API is 4,000 calls per minute. If the limit is exceeded, API calls are throttled, which may affect your business. Please call reasonably. For more information, see [QPS Limits](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Updates a video recording task.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [StartRecordTask](https://help.aliyun.com/document_detail/193528.html) to start a video recording task.
//
// # QPS limits
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - UpdateRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordTaskResponse
func (client *Client) UpdateRecordTaskWithOptions(request *UpdateRecordTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Updates a video recording task.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [StartRecordTask](https://help.aliyun.com/document_detail/193528.html) to start a video recording task.
//
// # QPS limits
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable frequency. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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

// Summary:
//
// Updates a recording configuration template.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html) to add a recording configuration template.
//
// # QPS limits
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
// @param request - UpdateRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordTemplateResponse
func (client *Client) UpdateRecordTemplateWithOptions(request *UpdateRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecordTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

// Summary:
//
// Updates a recording configuration template.
//
// Description:
//
// # Usage notes
//
// Before you call this operation, you must call [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html) to add a recording configuration template.
//
// # QPS limits
//
// The QPS limit for a single user is 100 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Make API calls at a reasonable rate. For more information, see [QPS limits](https://help.aliyun.com/document_detail/341645.html).
//
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
// Updates a bypass live streaming task.
//
// Description:
//
// ### Usage notes
//
// This operation updates a bypass live streaming task for an online channel.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 requests per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable frequency.
//
// @param tmpReq - UpdateStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStreamingOutResponse
func (client *Client) UpdateStreamingOutWithOptions(tmpReq *UpdateStreamingOutRequest, runtime *dara.RuntimeOptions) (_result *UpdateStreamingOutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
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
// Updates a bypass live streaming task.
//
// Description:
//
// ### Usage notes
//
// This operation updates a bypass live streaming task for an online channel.
//
// ### QPS limit
//
// The QPS limit for a single user is 10 requests per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable frequency.
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
