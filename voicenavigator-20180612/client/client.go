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
	client.Endpoint, _err = client.GetEndpoint(dara.String("voicenavigator"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AssociateChatbotInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateChatbotInstanceResponse
func (client *Client) AssociateChatbotInstanceWithOptions(request *AssociateChatbotInstanceRequest, runtime *dara.RuntimeOptions) (_result *AssociateChatbotInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatbotInstanceId) {
		query["ChatbotInstanceId"] = request.ChatbotInstanceId
	}

	if !dara.IsNil(request.ChatbotName) {
		query["ChatbotName"] = request.ChatbotName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NluServiceParamsJson) {
		query["NluServiceParamsJson"] = request.NluServiceParamsJson
	}

	if !dara.IsNil(request.NluServiceType) {
		query["NluServiceType"] = request.NluServiceType
	}

	if !dara.IsNil(request.UnionSource) {
		query["UnionSource"] = request.UnionSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateChatbotInstance"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateChatbotInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AssociateChatbotInstanceRequest
//
// @return AssociateChatbotInstanceResponse
func (client *Client) AssociateChatbotInstance(request *AssociateChatbotInstanceRequest) (_result *AssociateChatbotInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateChatbotInstanceResponse{}
	_body, _err := client.AssociateChatbotInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AuditTTSVoice
//
// @param request - AuditTTSVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuditTTSVoiceResponse
func (client *Client) AuditTTSVoiceWithOptions(request *AuditTTSVoiceRequest, runtime *dara.RuntimeOptions) (_result *AuditTTSVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessKey) {
		query["AccessKey"] = request.AccessKey
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PitchRate) {
		query["PitchRate"] = request.PitchRate
	}

	if !dara.IsNil(request.SecretKey) {
		query["SecretKey"] = request.SecretKey
	}

	if !dara.IsNil(request.SpeechRate) {
		query["SpeechRate"] = request.SpeechRate
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	if !dara.IsNil(request.Voice) {
		query["Voice"] = request.Voice
	}

	if !dara.IsNil(request.Volume) {
		query["Volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuditTTSVoice"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuditTTSVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AuditTTSVoice
//
// @param request - AuditTTSVoiceRequest
//
// @return AuditTTSVoiceResponse
func (client *Client) AuditTTSVoice(request *AuditTTSVoiceRequest) (_result *AuditTTSVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuditTTSVoiceResponse{}
	_body, _err := client.AuditTTSVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启会话
//
// @param request - BeginDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeginDialogueResponse
func (client *Client) BeginDialogueWithOptions(request *BeginDialogueRequest, runtime *dara.RuntimeOptions) (_result *BeginDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InitialContext) {
		query["InitialContext"] = request.InitialContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceOwnerId) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BeginDialogue"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BeginDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启会话
//
// @param request - BeginDialogueRequest
//
// @return BeginDialogueResponse
func (client *Client) BeginDialogue(request *BeginDialogueRequest) (_result *BeginDialogueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BeginDialogueResponse{}
	_body, _err := client.BeginDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CollectedNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CollectedNumberResponse
func (client *Client) CollectedNumberWithOptions(request *CollectedNumberRequest, runtime *dara.RuntimeOptions) (_result *CollectedNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalContext) {
		query["AdditionalContext"] = request.AdditionalContext
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceOwnerId) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CollectedNumber"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CollectedNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CollectedNumberRequest
//
// @return CollectedNumberResponse
func (client *Client) CollectedNumber(request *CollectedNumberRequest) (_result *CollectedNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CollectedNumberResponse{}
	_body, _err := client.CollectedNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # CreateDownloadUrl
//
// @param request - CreateDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDownloadUrlResponse
func (client *Client) CreateDownloadUrlWithOptions(request *CreateDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *CreateDownloadUrlResponse, _err error) {
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
		Action:      dara.String("CreateDownloadUrl"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateDownloadUrl
//
// @param request - CreateDownloadUrlRequest
//
// @return CreateDownloadUrlResponse
func (client *Client) CreateDownloadUrl(request *CreateDownloadUrlRequest) (_result *CreateDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDownloadUrlResponse{}
	_body, _err := client.CreateDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		query["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NluServiceParamsJson) {
		query["NluServiceParamsJson"] = request.NluServiceParamsJson
	}

	if !dara.IsNil(request.UnionInstanceId) {
		query["UnionInstanceId"] = request.UnionInstanceId
	}

	if !dara.IsNil(request.UnionSource) {
		query["UnionSource"] = request.UnionSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 测试窗开启文本对话
//
// @param request - DebugBeginDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugBeginDialogueResponse
func (client *Client) DebugBeginDialogueWithOptions(request *DebugBeginDialogueRequest, runtime *dara.RuntimeOptions) (_result *DebugBeginDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InitialContext) {
		query["InitialContext"] = request.InitialContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ShouldUseSandBox) {
		query["ShouldUseSandBox"] = request.ShouldUseSandBox
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DebugBeginDialogue"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DebugBeginDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 测试窗开启文本对话
//
// @param request - DebugBeginDialogueRequest
//
// @return DebugBeginDialogueResponse
func (client *Client) DebugBeginDialogue(request *DebugBeginDialogueRequest) (_result *DebugBeginDialogueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DebugBeginDialogueResponse{}
	_body, _err := client.DebugBeginDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DebugCollectedNumber
//
// @param request - DebugCollectedNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugCollectedNumberResponse
func (client *Client) DebugCollectedNumberWithOptions(request *DebugCollectedNumberRequest, runtime *dara.RuntimeOptions) (_result *DebugCollectedNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DebugCollectedNumber"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DebugCollectedNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DebugCollectedNumber
//
// @param request - DebugCollectedNumberRequest
//
// @return DebugCollectedNumberResponse
func (client *Client) DebugCollectedNumber(request *DebugCollectedNumberRequest) (_result *DebugCollectedNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DebugCollectedNumberResponse{}
	_body, _err := client.DebugCollectedNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DebugDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DebugDialogueResponse
func (client *Client) DebugDialogueWithOptions(request *DebugDialogueRequest, runtime *dara.RuntimeOptions) (_result *DebugDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalContext) {
		query["AdditionalContext"] = request.AdditionalContext
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DebugDialogue"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DebugDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DebugDialogueRequest
//
// @return DebugDialogueResponse
func (client *Client) DebugDialogue(request *DebugDialogueRequest) (_result *DebugDialogueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DebugDialogueResponse{}
	_body, _err := client.DebugDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeConversationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConversationResponse
func (client *Client) DescribeConversationWithOptions(request *DescribeConversationRequest, runtime *dara.RuntimeOptions) (_result *DescribeConversationResponse, _err error) {
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
		Action:      dara.String("DescribeConversation"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeConversationRequest
//
// @return DescribeConversationResponse
func (client *Client) DescribeConversation(request *DescribeConversationRequest) (_result *DescribeConversationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeConversationResponse{}
	_body, _err := client.DescribeConversationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeConversationContext
//
// @param request - DescribeConversationContextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConversationContextResponse
func (client *Client) DescribeConversationContextWithOptions(request *DescribeConversationContextRequest, runtime *dara.RuntimeOptions) (_result *DescribeConversationContextResponse, _err error) {
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
		Action:      dara.String("DescribeConversationContext"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConversationContextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeConversationContext
//
// @param request - DescribeConversationContextRequest
//
// @return DescribeConversationContextResponse
func (client *Client) DescribeConversationContext(request *DescribeConversationContextRequest) (_result *DescribeConversationContextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeConversationContextResponse{}
	_body, _err := client.DescribeConversationContextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeExportProgress
//
// @param request - DescribeExportProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExportProgressResponse
func (client *Client) DescribeExportProgressWithOptions(request *DescribeExportProgressRequest, runtime *dara.RuntimeOptions) (_result *DescribeExportProgressResponse, _err error) {
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
		Action:      dara.String("DescribeExportProgress"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExportProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeExportProgress
//
// @param request - DescribeExportProgressRequest
//
// @return DescribeExportProgressResponse
func (client *Client) DescribeExportProgress(request *DescribeExportProgressRequest) (_result *DescribeExportProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExportProgressResponse{}
	_body, _err := client.DescribeExportProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
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
		Action:      dara.String("DescribeInstance"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeInstanceRequest
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeNavigationConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNavigationConfigResponse
func (client *Client) DescribeNavigationConfigWithOptions(request *DescribeNavigationConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeNavigationConfigResponse, _err error) {
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
		Action:      dara.String("DescribeNavigationConfig"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNavigationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeNavigationConfigRequest
//
// @return DescribeNavigationConfigResponse
func (client *Client) DescribeNavigationConfig(request *DescribeNavigationConfigRequest) (_result *DescribeNavigationConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNavigationConfigResponse{}
	_body, _err := client.DescribeNavigationConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordingResponse
func (client *Client) DescribeRecordingWithOptions(request *DescribeRecordingRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecordingResponse, _err error) {
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
		Action:      dara.String("DescribeRecording"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecordingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRecordingRequest
//
// @return DescribeRecordingResponse
func (client *Client) DescribeRecording(request *DescribeRecordingRequest) (_result *DescribeRecordingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecordingResponse{}
	_body, _err := client.DescribeRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeStatisticalDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStatisticalDataResponse
func (client *Client) DescribeStatisticalDataWithOptions(request *DescribeStatisticalDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeStatisticalDataResponse, _err error) {
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
		Action:      dara.String("DescribeStatisticalData"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStatisticalDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeStatisticalDataRequest
//
// @return DescribeStatisticalDataResponse
func (client *Client) DescribeStatisticalData(request *DescribeStatisticalDataRequest) (_result *DescribeStatisticalDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStatisticalDataResponse{}
	_body, _err := client.DescribeStatisticalDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取TTS配置
//
// @param request - DescribeTTSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTTSConfigResponse
func (client *Client) DescribeTTSConfigWithOptions(request *DescribeTTSConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeTTSConfigResponse, _err error) {
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
		Action:      dara.String("DescribeTTSConfig"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTTSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取TTS配置
//
// @param request - DescribeTTSConfigRequest
//
// @return DescribeTTSConfigResponse
func (client *Client) DescribeTTSConfig(request *DescribeTTSConfigRequest) (_result *DescribeTTSConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTTSConfigResponse{}
	_body, _err := client.DescribeTTSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DialogueResponse
func (client *Client) DialogueWithOptions(request *DialogueRequest, runtime *dara.RuntimeOptions) (_result *DialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalContext) {
		query["AdditionalContext"] = request.AdditionalContext
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Emotion) {
		query["Emotion"] = request.Emotion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceOwnerId) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Dialogue"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DialogueRequest
//
// @return DialogueResponse
func (client *Client) Dialogue(request *DialogueRequest) (_result *DialogueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DialogueResponse{}
	_body, _err := client.DialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInstanceResponse
func (client *Client) DisableInstanceWithOptions(request *DisableInstanceRequest, runtime *dara.RuntimeOptions) (_result *DisableInstanceResponse, _err error) {
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
		Action:      dara.String("DisableInstance"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableInstanceRequest
//
// @return DisableInstanceResponse
func (client *Client) DisableInstance(request *DisableInstanceRequest) (_result *DisableInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableInstanceResponse{}
	_body, _err := client.DisableInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInstanceResponse
func (client *Client) EnableInstanceWithOptions(request *EnableInstanceRequest, runtime *dara.RuntimeOptions) (_result *EnableInstanceResponse, _err error) {
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
		Action:      dara.String("EnableInstance"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableInstanceRequest
//
// @return EnableInstanceResponse
func (client *Client) EnableInstance(request *EnableInstanceRequest) (_result *EnableInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableInstanceResponse{}
	_body, _err := client.EnableInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EndDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndDialogueResponse
func (client *Client) EndDialogueWithOptions(request *EndDialogueRequest, runtime *dara.RuntimeOptions) (_result *EndDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.HangUpParams) {
		query["HangUpParams"] = request.HangUpParams
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceOwnerId) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EndDialogue"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EndDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EndDialogueRequest
//
// @return EndDialogueResponse
func (client *Client) EndDialogue(request *EndDialogueRequest) (_result *EndDialogueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EndDialogueResponse{}
	_body, _err := client.EndDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ExportConversationDetails
//
// @param request - ExportConversationDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportConversationDetailsResponse
func (client *Client) ExportConversationDetailsWithOptions(request *ExportConversationDetailsRequest, runtime *dara.RuntimeOptions) (_result *ExportConversationDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTimeLeftRange) {
		query["BeginTimeLeftRange"] = request.BeginTimeLeftRange
	}

	if !dara.IsNil(request.BeginTimeRightRange) {
		query["BeginTimeRightRange"] = request.BeginTimeRightRange
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.DebugConversation) {
		query["DebugConversation"] = request.DebugConversation
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Options) {
		query["Options"] = request.Options
	}

	if !dara.IsNil(request.Result) {
		query["Result"] = request.Result
	}

	if !dara.IsNil(request.RoundsLeftRange) {
		query["RoundsLeftRange"] = request.RoundsLeftRange
	}

	if !dara.IsNil(request.RoundsRightRange) {
		query["RoundsRightRange"] = request.RoundsRightRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportConversationDetails"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportConversationDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ExportConversationDetails
//
// @param request - ExportConversationDetailsRequest
//
// @return ExportConversationDetailsResponse
func (client *Client) ExportConversationDetails(request *ExportConversationDetailsRequest) (_result *ExportConversationDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportConversationDetailsResponse{}
	_body, _err := client.ExportConversationDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ExportStatisticalDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportStatisticalDataResponse
func (client *Client) ExportStatisticalDataWithOptions(request *ExportStatisticalDataRequest, runtime *dara.RuntimeOptions) (_result *ExportStatisticalDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTimeLeftRange) {
		query["BeginTimeLeftRange"] = request.BeginTimeLeftRange
	}

	if !dara.IsNil(request.BeginTimeRightRange) {
		query["BeginTimeRightRange"] = request.BeginTimeRightRange
	}

	if !dara.IsNil(request.ExportType) {
		query["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TimeUnit) {
		query["TimeUnit"] = request.TimeUnit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportStatisticalData"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportStatisticalDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ExportStatisticalDataRequest
//
// @return ExportStatisticalDataResponse
func (client *Client) ExportStatisticalData(request *ExportStatisticalDataRequest) (_result *ExportStatisticalDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportStatisticalDataResponse{}
	_body, _err := client.ExportStatisticalDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Asr配置
//
// @param request - GetAsrConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsrConfigResponse
func (client *Client) GetAsrConfigWithOptions(request *GetAsrConfigRequest, runtime *dara.RuntimeOptions) (_result *GetAsrConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigLevel) {
		query["ConfigLevel"] = request.ConfigLevel
	}

	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsrConfig"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsrConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Asr配置
//
// @param request - GetAsrConfigRequest
//
// @return GetAsrConfigResponse
func (client *Client) GetAsrConfig(request *GetAsrConfigRequest) (_result *GetAsrConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAsrConfigResponse{}
	_body, _err := client.GetAsrConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetRealTimeConcurrency
//
// @param request - GetRealTimeConcurrencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealTimeConcurrencyResponse
func (client *Client) GetRealTimeConcurrencyWithOptions(request *GetRealTimeConcurrencyRequest, runtime *dara.RuntimeOptions) (_result *GetRealTimeConcurrencyResponse, _err error) {
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
		Action:      dara.String("GetRealTimeConcurrency"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRealTimeConcurrencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetRealTimeConcurrency
//
// @param request - GetRealTimeConcurrencyRequest
//
// @return GetRealTimeConcurrencyResponse
func (client *Client) GetRealTimeConcurrency(request *GetRealTimeConcurrencyRequest) (_result *GetRealTimeConcurrencyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRealTimeConcurrencyResponse{}
	_body, _err := client.GetRealTimeConcurrencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListChatbotInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatbotInstancesResponse
func (client *Client) ListChatbotInstancesWithOptions(request *ListChatbotInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListChatbotInstancesResponse, _err error) {
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
		Action:      dara.String("ListChatbotInstances"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatbotInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListChatbotInstancesRequest
//
// @return ListChatbotInstancesResponse
func (client *Client) ListChatbotInstances(request *ListChatbotInstancesRequest) (_result *ListChatbotInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChatbotInstancesResponse{}
	_body, _err := client.ListChatbotInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListConversationDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConversationDetailsResponse
func (client *Client) ListConversationDetailsWithOptions(request *ListConversationDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListConversationDetailsResponse, _err error) {
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
		Action:      dara.String("ListConversationDetails"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConversationDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListConversationDetailsRequest
//
// @return ListConversationDetailsResponse
func (client *Client) ListConversationDetails(request *ListConversationDetailsRequest) (_result *ListConversationDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConversationDetailsResponse{}
	_body, _err := client.ListConversationDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会话列表
//
// @param request - ListConversationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConversationsResponse
func (client *Client) ListConversationsWithOptions(request *ListConversationsRequest, runtime *dara.RuntimeOptions) (_result *ListConversationsResponse, _err error) {
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
		Action:      dara.String("ListConversations"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConversationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话列表
//
// @param request - ListConversationsRequest
//
// @return ListConversationsResponse
func (client *Client) ListConversations(request *ListConversationsRequest) (_result *ListConversationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConversationsResponse{}
	_body, _err := client.ListConversationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载列表
//
// @param request - ListDownloadTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownloadTasksResponse
func (client *Client) ListDownloadTasksWithOptions(request *ListDownloadTasksRequest, runtime *dara.RuntimeOptions) (_result *ListDownloadTasksResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDownloadTasks"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDownloadTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载列表
//
// @param request - ListDownloadTasksRequest
//
// @return ListDownloadTasksResponse
func (client *Client) ListDownloadTasks(request *ListDownloadTasksRequest) (_result *ListDownloadTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDownloadTasksResponse{}
	_body, _err := client.ListDownloadTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
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
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Asr配置
//
// @param request - ModifyAsrConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAsrConfigResponse
func (client *Client) ModifyAsrConfigWithOptions(request *ModifyAsrConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyAsrConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.AsrAcousticModelId) {
		query["AsrAcousticModelId"] = request.AsrAcousticModelId
	}

	if !dara.IsNil(request.AsrClassVocabularyId) {
		query["AsrClassVocabularyId"] = request.AsrClassVocabularyId
	}

	if !dara.IsNil(request.AsrCustomizationId) {
		query["AsrCustomizationId"] = request.AsrCustomizationId
	}

	if !dara.IsNil(request.AsrOverrides) {
		query["AsrOverrides"] = request.AsrOverrides
	}

	if !dara.IsNil(request.AsrVocabularyId) {
		query["AsrVocabularyId"] = request.AsrVocabularyId
	}

	if !dara.IsNil(request.ConfigLevel) {
		query["ConfigLevel"] = request.ConfigLevel
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.NlsServiceType) {
		query["NlsServiceType"] = request.NlsServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAsrConfig"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAsrConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改Asr配置
//
// @param request - ModifyAsrConfigRequest
//
// @return ModifyAsrConfigResponse
func (client *Client) ModifyAsrConfig(request *ModifyAsrConfigRequest) (_result *ModifyAsrConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAsrConfigResponse{}
	_body, _err := client.ModifyAsrConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyGreetingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGreetingConfigResponse
func (client *Client) ModifyGreetingConfigWithOptions(request *ModifyGreetingConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyGreetingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GreetingWords) {
		query["GreetingWords"] = request.GreetingWords
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentTrigger) {
		query["IntentTrigger"] = request.IntentTrigger
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGreetingConfig"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGreetingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyGreetingConfigRequest
//
// @return ModifyGreetingConfigResponse
func (client *Client) ModifyGreetingConfig(request *ModifyGreetingConfigRequest) (_result *ModifyGreetingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGreetingConfigResponse{}
	_body, _err := client.ModifyGreetingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
	if !dara.IsNil(request.Concurrency) {
		query["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstance"),
		Version:     dara.String("2018-06-12"),
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

// @param request - ModifySilenceTimeoutConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySilenceTimeoutConfigResponse
func (client *Client) ModifySilenceTimeoutConfigWithOptions(request *ModifySilenceTimeoutConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifySilenceTimeoutConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FinalAction) {
		query["FinalAction"] = request.FinalAction
	}

	if !dara.IsNil(request.FinalActionParams) {
		query["FinalActionParams"] = request.FinalActionParams
	}

	if !dara.IsNil(request.FinalPrompt) {
		query["FinalPrompt"] = request.FinalPrompt
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentTrigger) {
		query["IntentTrigger"] = request.IntentTrigger
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySilenceTimeoutConfig"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySilenceTimeoutConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifySilenceTimeoutConfigRequest
//
// @return ModifySilenceTimeoutConfigResponse
func (client *Client) ModifySilenceTimeoutConfig(request *ModifySilenceTimeoutConfigRequest) (_result *ModifySilenceTimeoutConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySilenceTimeoutConfigResponse{}
	_body, _err := client.ModifySilenceTimeoutConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改TTS配置
//
// @param request - ModifyTTSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTTSConfigResponse
func (client *Client) ModifyTTSConfigWithOptions(request *ModifyTTSConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyTTSConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliCustomizedVoice) {
		query["AliCustomizedVoice"] = request.AliCustomizedVoice
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineXunfei) {
		query["EngineXunfei"] = request.EngineXunfei
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NlsServiceType) {
		query["NlsServiceType"] = request.NlsServiceType
	}

	if !dara.IsNil(request.PitchRate) {
		query["PitchRate"] = request.PitchRate
	}

	if !dara.IsNil(request.SpeechRate) {
		query["SpeechRate"] = request.SpeechRate
	}

	if !dara.IsNil(request.TtsOverrides) {
		query["TtsOverrides"] = request.TtsOverrides
	}

	if !dara.IsNil(request.Voice) {
		query["Voice"] = request.Voice
	}

	if !dara.IsNil(request.Volume) {
		query["Volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTTSConfig"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTTSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改TTS配置
//
// @param request - ModifyTTSConfigRequest
//
// @return ModifyTTSConfigResponse
func (client *Client) ModifyTTSConfig(request *ModifyTTSConfigRequest) (_result *ModifyTTSConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTTSConfigResponse{}
	_body, _err := client.ModifyTTSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyUnrecognizingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUnrecognizingConfigResponse
func (client *Client) ModifyUnrecognizingConfigWithOptions(request *ModifyUnrecognizingConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyUnrecognizingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FinalAction) {
		query["FinalAction"] = request.FinalAction
	}

	if !dara.IsNil(request.FinalActionParams) {
		query["FinalActionParams"] = request.FinalActionParams
	}

	if !dara.IsNil(request.FinalPrompt) {
		query["FinalPrompt"] = request.FinalPrompt
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUnrecognizingConfig"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUnrecognizingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyUnrecognizingConfigRequest
//
// @return ModifyUnrecognizingConfigResponse
func (client *Client) ModifyUnrecognizingConfig(request *ModifyUnrecognizingConfigRequest) (_result *ModifyUnrecognizingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUnrecognizingConfigResponse{}
	_body, _err := client.ModifyUnrecognizingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryConversationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConversationsResponse
func (client *Client) QueryConversationsWithOptions(request *QueryConversationsRequest, runtime *dara.RuntimeOptions) (_result *QueryConversationsResponse, _err error) {
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
		Action:      dara.String("QueryConversations"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConversationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryConversationsRequest
//
// @return QueryConversationsResponse
func (client *Client) QueryConversations(request *QueryConversationsRequest) (_result *QueryConversationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryConversationsResponse{}
	_body, _err := client.QueryConversationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveRecordingResponse
func (client *Client) SaveRecordingWithOptions(request *SaveRecordingRequest, runtime *dara.RuntimeOptions) (_result *SaveRecordingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceOwnerId) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VoiceSliceRecordingList) {
		query["VoiceSliceRecordingList"] = request.VoiceSliceRecordingList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveRecording"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveRecordingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveRecordingRequest
//
// @return SaveRecordingResponse
func (client *Client) SaveRecording(request *SaveRecordingRequest) (_result *SaveRecordingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveRecordingResponse{}
	_body, _err := client.SaveRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SilenceTimeoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SilenceTimeoutResponse
func (client *Client) SilenceTimeoutWithOptions(request *SilenceTimeoutRequest, runtime *dara.RuntimeOptions) (_result *SilenceTimeoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConversationId) {
		query["ConversationId"] = request.ConversationId
	}

	if !dara.IsNil(request.InitialContext) {
		query["InitialContext"] = request.InitialContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceOwnerId) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SilenceTimeout"),
		Version:     dara.String("2018-06-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SilenceTimeoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SilenceTimeoutRequest
//
// @return SilenceTimeoutResponse
func (client *Client) SilenceTimeout(request *SilenceTimeoutRequest) (_result *SilenceTimeoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SilenceTimeoutResponse{}
	_body, _err := client.SilenceTimeoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
