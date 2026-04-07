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

// Summary:
//
// 创建实例
//
// @param request - CreateCloneVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloneVoiceResponse
func (client *Client) CreateCloneVoiceWithOptions(request *CreateCloneVoiceRequest, runtime *dara.RuntimeOptions) (_result *CreateCloneVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloneVoice"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloneVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param request - CreateCloneVoiceRequest
//
// @return CreateCloneVoiceResponse
func (client *Client) CreateCloneVoice(request *CreateCloneVoiceRequest) (_result *CreateCloneVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloneVoiceResponse{}
	_body, _err := client.CreateCloneVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param request - CreateScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScriptResponse
func (client *Client) CreateScriptWithOptions(request *CreateScriptRequest, runtime *dara.RuntimeOptions) (_result *CreateScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NluEngine) {
		body["NluEngine"] = request.NluEngine
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param request - CreateScriptRequest
//
// @return CreateScriptResponse
func (client *Client) CreateScript(request *CreateScriptRequest) (_result *CreateScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScriptResponse{}
	_body, _err := client.CreateScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建场景配置
//
// @param tmpReq - CreateScriptVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScriptVersionResponse
func (client *Client) CreateScriptVersionWithOptions(tmpReq *CreateScriptVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateScriptVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateScriptVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InteractionConfig) {
		request.InteractionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InteractionConfig, dara.String("InteractionConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LabelConfig) {
		request.LabelConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelConfig, dara.String("LabelConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScriptProfile) {
		request.ScriptProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScriptProfile, dara.String("ScriptProfile"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SynthesizerConfig) {
		request.SynthesizerConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SynthesizerConfig, dara.String("SynthesizerConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TranscriberConfig) {
		request.TranscriberConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TranscriberConfig, dara.String("TranscriberConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InteractionConfigShrink) {
		body["InteractionConfig"] = request.InteractionConfigShrink
	}

	if !dara.IsNil(request.LabelConfigShrink) {
		body["LabelConfig"] = request.LabelConfigShrink
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.ScriptProfileShrink) {
		body["ScriptProfile"] = request.ScriptProfileShrink
	}

	if !dara.IsNil(request.SourceVersionId) {
		body["SourceVersionId"] = request.SourceVersionId
	}

	if !dara.IsNil(request.SynthesizerConfigShrink) {
		body["SynthesizerConfig"] = request.SynthesizerConfigShrink
	}

	if !dara.IsNil(request.TranscriberConfigShrink) {
		body["TranscriberConfig"] = request.TranscriberConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScriptVersion"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScriptVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建场景配置
//
// @param request - CreateScriptVersionRequest
//
// @return CreateScriptVersionResponse
func (client *Client) CreateScriptVersion(request *CreateScriptVersionRequest) (_result *CreateScriptVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScriptVersionResponse{}
	_body, _err := client.CreateScriptVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建变量
//
// @param request - CreateVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVariableResponse
func (client *Client) CreateVariableWithOptions(request *CreateVariableRequest, runtime *dara.RuntimeOptions) (_result *CreateVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVariable"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建变量
//
// @param request - CreateVariableRequest
//
// @return CreateVariableResponse
func (client *Client) CreateVariable(request *CreateVariableRequest) (_result *CreateVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVariableResponse{}
	_body, _err := client.CreateVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param tmpReq - CreateVocabularyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVocabularyResponse
func (client *Client) CreateVocabularyWithOptions(tmpReq *CreateVocabularyRequest, runtime *dara.RuntimeOptions) (_result *CreateVocabularyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateVocabularyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Words) {
		request.WordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Words, dara.String("Words"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.WordsShrink) {
		body["Words"] = request.WordsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVocabulary"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVocabularyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param request - CreateVocabularyRequest
//
// @return CreateVocabularyResponse
func (client *Client) CreateVocabulary(request *CreateVocabularyRequest) (_result *CreateVocabularyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVocabularyResponse{}
	_body, _err := client.CreateVocabularyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param tmpReq - CreateVoiceAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVoiceAccessProfileResponse
func (client *Client) CreateVoiceAccessProfileWithOptions(tmpReq *CreateVoiceAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateVoiceAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateVoiceAccessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Profile) {
		request.ProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Profile, dara.String("Profile"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NlsEngine) {
		body["NlsEngine"] = request.NlsEngine
	}

	if !dara.IsNil(request.ProfileShrink) {
		body["Profile"] = request.ProfileShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVoiceAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVoiceAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实例
//
// @param request - CreateVoiceAccessProfileRequest
//
// @return CreateVoiceAccessProfileResponse
func (client *Client) CreateVoiceAccessProfile(request *CreateVoiceAccessProfileRequest) (_result *CreateVoiceAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVoiceAccessProfileResponse{}
	_body, _err := client.CreateVoiceAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteCloneVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloneVoiceResponse
func (client *Client) DeleteCloneVoiceWithOptions(request *DeleteCloneVoiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloneVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloneVoiceId) {
		body["CloneVoiceId"] = request.CloneVoiceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloneVoice"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloneVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteCloneVoiceRequest
//
// @return DeleteCloneVoiceResponse
func (client *Client) DeleteCloneVoice(request *DeleteCloneVoiceRequest) (_result *DeleteCloneVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloneVoiceResponse{}
	_body, _err := client.DeleteCloneVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScriptResponse
func (client *Client) DeleteScriptWithOptions(request *DeleteScriptRequest, runtime *dara.RuntimeOptions) (_result *DeleteScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteScriptRequest
//
// @return DeleteScriptResponse
func (client *Client) DeleteScript(request *DeleteScriptRequest) (_result *DeleteScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteScriptResponse{}
	_body, _err := client.DeleteScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除变量
//
// @param request - DeleteVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVariableResponse
func (client *Client) DeleteVariableWithOptions(request *DeleteVariableRequest, runtime *dara.RuntimeOptions) (_result *DeleteVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VariableId) {
		body["VariableId"] = request.VariableId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVariable"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除变量
//
// @param request - DeleteVariableRequest
//
// @return DeleteVariableResponse
func (client *Client) DeleteVariable(request *DeleteVariableRequest) (_result *DeleteVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVariableResponse{}
	_body, _err := client.DeleteVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteVocabularyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVocabularyResponse
func (client *Client) DeleteVocabularyWithOptions(request *DeleteVocabularyRequest, runtime *dara.RuntimeOptions) (_result *DeleteVocabularyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VocabularyId) {
		body["VocabularyId"] = request.VocabularyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVocabulary"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVocabularyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteVocabularyRequest
//
// @return DeleteVocabularyResponse
func (client *Client) DeleteVocabulary(request *DeleteVocabularyRequest) (_result *DeleteVocabularyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVocabularyResponse{}
	_body, _err := client.DeleteVocabularyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteVoiceAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVoiceAccessProfileResponse
func (client *Client) DeleteVoiceAccessProfileWithOptions(request *DeleteVoiceAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteVoiceAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessProfileId) {
		body["AccessProfileId"] = request.AccessProfileId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVoiceAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVoiceAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除场景
//
// @param request - DeleteVoiceAccessProfileRequest
//
// @return DeleteVoiceAccessProfileResponse
func (client *Client) DeleteVoiceAccessProfile(request *DeleteVoiceAccessProfileRequest) (_result *DeleteVoiceAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVoiceAccessProfileResponse{}
	_body, _err := client.DeleteVoiceAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用消息订阅
//
// @param request - DisableSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSubscriptionResponse
func (client *Client) DisableSubscriptionWithOptions(request *DisableSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *DisableSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSubscription"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用消息订阅
//
// @param request - DisableSubscriptionRequest
//
// @return DisableSubscriptionResponse
func (client *Client) DisableSubscription(request *DisableSubscriptionRequest) (_result *DisableSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableSubscriptionResponse{}
	_body, _err := client.DisableSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ExportScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportScriptResponse
func (client *Client) ExportScriptWithOptions(request *ExportScriptRequest, runtime *dara.RuntimeOptions) (_result *ExportScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ExportScriptRequest
//
// @return ExportScriptResponse
func (client *Client) ExportScript(request *ExportScriptRequest) (_result *ExportScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportScriptResponse{}
	_body, _err := client.ExportScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出热词
//
// @param tmpReq - ExportVocabularyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportVocabularyResponse
func (client *Client) ExportVocabularyWithOptions(tmpReq *ExportVocabularyRequest, runtime *dara.RuntimeOptions) (_result *ExportVocabularyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExportVocabularyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VocabularyIds) {
		request.VocabularyIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VocabularyIds, dara.String("VocabularyIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VocabularyIdsShrink) {
		body["VocabularyIds"] = request.VocabularyIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportVocabulary"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportVocabularyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出热词
//
// @param request - ExportVocabularyRequest
//
// @return ExportVocabularyResponse
func (client *Client) ExportVocabulary(request *ExportVocabularyRequest) (_result *ExportVocabularyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportVocabularyResponse{}
	_body, _err := client.ExportVocabularyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - GenerateFileUploadParamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateFileUploadParamsResponse
func (client *Client) GenerateFileUploadParamsWithOptions(request *GenerateFileUploadParamsRequest, runtime *dara.RuntimeOptions) (_result *GenerateFileUploadParamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		body["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateFileUploadParams"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateFileUploadParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - GenerateFileUploadParamsRequest
//
// @return GenerateFileUploadParamsResponse
func (client *Client) GenerateFileUploadParams(request *GenerateFileUploadParamsRequest) (_result *GenerateFileUploadParamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateFileUploadParamsResponse{}
	_body, _err := client.GenerateFileUploadParamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取通话详情
//
// @param request - GetCallDetailRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCallDetailRecordResponse
func (client *Client) GetCallDetailRecordWithOptions(request *GetCallDetailRecordRequest, runtime *dara.RuntimeOptions) (_result *GetCallDetailRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCallDetailRecord"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCallDetailRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取通话详情
//
// @param request - GetCallDetailRecordRequest
//
// @return GetCallDetailRecordResponse
func (client *Client) GetCallDetailRecord(request *GetCallDetailRecordRequest) (_result *GetCallDetailRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCallDetailRecordResponse{}
	_body, _err := client.GetCallDetailRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例实时指标
//
// @param request - GetRealtimeInstanceStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealtimeInstanceStatsResponse
func (client *Client) GetRealtimeInstanceStatsWithOptions(request *GetRealtimeInstanceStatsRequest, runtime *dara.RuntimeOptions) (_result *GetRealtimeInstanceStatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRealtimeInstanceStats"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRealtimeInstanceStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例实时指标
//
// @param request - GetRealtimeInstanceStatsRequest
//
// @return GetRealtimeInstanceStatsResponse
func (client *Client) GetRealtimeInstanceStats(request *GetRealtimeInstanceStatsRequest) (_result *GetRealtimeInstanceStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRealtimeInstanceStatsResponse{}
	_body, _err := client.GetRealtimeInstanceStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取录音
//
// @param request - GetRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecordingResponse
func (client *Client) GetRecordingWithOptions(request *GetRecordingRequest, runtime *dara.RuntimeOptions) (_result *GetRecordingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRecording"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecordingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取录音
//
// @param request - GetRecordingRequest
//
// @return GetRecordingResponse
func (client *Client) GetRecording(request *GetRecordingRequest) (_result *GetRecordingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRecordingResponse{}
	_body, _err := client.GetRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - GetScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScriptResponse
func (client *Client) GetScriptWithOptions(request *GetScriptRequest, runtime *dara.RuntimeOptions) (_result *GetScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - GetScriptRequest
//
// @return GetScriptResponse
func (client *Client) GetScript(request *GetScriptRequest) (_result *GetScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScriptResponse{}
	_body, _err := client.GetScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取MQ配置
//
// @param request - GetSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscriptionResponse
func (client *Client) GetSubscriptionWithOptions(request *GetSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *GetSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSubscription"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取MQ配置
//
// @param request - GetSubscriptionRequest
//
// @return GetSubscriptionResponse
func (client *Client) GetSubscription(request *GetSubscriptionRequest) (_result *GetSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSubscriptionResponse{}
	_body, _err := client.GetSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - GetVocabularyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVocabularyResponse
func (client *Client) GetVocabularyWithOptions(request *GetVocabularyRequest, runtime *dara.RuntimeOptions) (_result *GetVocabularyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VocabularyId) {
		body["VocabularyId"] = request.VocabularyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVocabulary"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVocabularyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - GetVocabularyRequest
//
// @return GetVocabularyResponse
func (client *Client) GetVocabulary(request *GetVocabularyRequest) (_result *GetVocabularyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVocabularyResponse{}
	_body, _err := client.GetVocabularyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入热词
//
// @param request - ImportVocabularyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportVocabularyResponse
func (client *Client) ImportVocabularyWithOptions(request *ImportVocabularyRequest, runtime *dara.RuntimeOptions) (_result *ImportVocabularyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileKey) {
		body["FileKey"] = request.FileKey
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportVocabulary"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportVocabularyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入热词
//
// @param request - ImportVocabularyRequest
//
// @return ImportVocabularyResponse
func (client *Client) ImportVocabulary(request *ImportVocabularyRequest) (_result *ImportVocabularyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportVocabularyResponse{}
	_body, _err := client.ImportVocabularyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取背景音列表
//
// @param request - ListBackgroundMusicsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBackgroundMusicsResponse
func (client *Client) ListBackgroundMusicsWithOptions(request *ListBackgroundMusicsRequest, runtime *dara.RuntimeOptions) (_result *ListBackgroundMusicsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBackgroundMusics"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBackgroundMusicsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取背景音列表
//
// @param request - ListBackgroundMusicsRequest
//
// @return ListBackgroundMusicsResponse
func (client *Client) ListBackgroundMusics(request *ListBackgroundMusicsRequest) (_result *ListBackgroundMusicsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBackgroundMusicsResponse{}
	_body, _err := client.ListBackgroundMusicsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param tmpReq - ListCallDetailRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCallDetailRecordsResponse
func (client *Client) ListCallDetailRecordsWithOptions(tmpReq *ListCallDetailRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListCallDetailRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCallDetailRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DispositionCodes) {
		request.DispositionCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DispositionCodes, dara.String("DispositionCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DispositionReasons) {
		request.DispositionReasonsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DispositionReasons, dara.String("DispositionReasons"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SessionIds) {
		request.SessionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionIds, dara.String("SessionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessChannelId) {
		query["AccessChannelId"] = request.AccessChannelId
	}

	if !dara.IsNil(request.AccessChannelType) {
		query["AccessChannelType"] = request.AccessChannelType
	}

	if !dara.IsNil(request.DraftVersion) {
		query["DraftVersion"] = request.DraftVersion
	}

	if !dara.IsNil(request.IssueResolved) {
		query["IssueResolved"] = request.IssueResolved
	}

	if !dara.IsNil(request.MaxTalkTurns) {
		query["MaxTalkTurns"] = request.MaxTalkTurns
	}

	if !dara.IsNil(request.MinTalkTurns) {
		query["MinTalkTurns"] = request.MinTalkTurns
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Callee) {
		body["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		body["Caller"] = request.Caller
	}

	if !dara.IsNil(request.DispositionCodesShrink) {
		body["DispositionCodes"] = request.DispositionCodesShrink
	}

	if !dara.IsNil(request.DispositionReasonsShrink) {
		body["DispositionReasons"] = request.DispositionReasonsShrink
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SessionIdsShrink) {
		body["SessionIds"] = request.SessionIdsShrink
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCallDetailRecords"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCallDetailRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListCallDetailRecordsRequest
//
// @return ListCallDetailRecordsResponse
func (client *Client) ListCallDetailRecords(request *ListCallDetailRecordsRequest) (_result *ListCallDetailRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCallDetailRecordsResponse{}
	_body, _err := client.ListCallDetailRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListCloneVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloneVoiceResponse
func (client *Client) ListCloneVoiceWithOptions(request *ListCloneVoiceRequest, runtime *dara.RuntimeOptions) (_result *ListCloneVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloneVoice"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloneVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListCloneVoiceRequest
//
// @return ListCloneVoiceResponse
func (client *Client) ListCloneVoice(request *ListCloneVoiceRequest) (_result *ListCloneVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloneVoiceResponse{}
	_body, _err := client.ListCloneVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取克隆音色可用模型列表
//
// @param request - ListCloneVoiceModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloneVoiceModelsResponse
func (client *Client) ListCloneVoiceModelsWithOptions(request *ListCloneVoiceModelsRequest, runtime *dara.RuntimeOptions) (_result *ListCloneVoiceModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCloneVoiceModels"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloneVoiceModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取克隆音色可用模型列表
//
// @param request - ListCloneVoiceModelsRequest
//
// @return ListCloneVoiceModelsResponse
func (client *Client) ListCloneVoiceModels(request *ListCloneVoiceModelsRequest) (_result *ListCloneVoiceModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloneVoiceModelsResponse{}
	_body, _err := client.ListCloneVoiceModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取对话模型列表
//
// @param request - ListNluModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNluModelsResponse
func (client *Client) ListNluModelsWithOptions(request *ListNluModelsRequest, runtime *dara.RuntimeOptions) (_result *ListNluModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNluModels"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNluModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取对话模型列表
//
// @param request - ListNluModelsRequest
//
// @return ListNluModelsResponse
func (client *Client) ListNluModels(request *ListNluModelsRequest) (_result *ListNluModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNluModelsResponse{}
	_body, _err := client.ListNluModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取场景配置模板列表
//
// @param request - ListScriptProfileTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptProfileTemplatesResponse
func (client *Client) ListScriptProfileTemplatesWithOptions(request *ListScriptProfileTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListScriptProfileTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScriptProfileTemplates"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptProfileTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取场景配置模板列表
//
// @param request - ListScriptProfileTemplatesRequest
//
// @return ListScriptProfileTemplatesResponse
func (client *Client) ListScriptProfileTemplates(request *ListScriptProfileTemplatesRequest) (_result *ListScriptProfileTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScriptProfileTemplatesResponse{}
	_body, _err := client.ListScriptProfileTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param tmpReq - ListScriptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptsResponse
func (client *Client) ListScriptsWithOptions(tmpReq *ListScriptsRequest, runtime *dara.RuntimeOptions) (_result *ListScriptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListScriptsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ScriptIds) {
		request.ScriptIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScriptIds, dara.String("ScriptIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Number) {
		body["Number"] = request.Number
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptIdsShrink) {
		body["ScriptIds"] = request.ScriptIdsShrink
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScripts"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListScriptsRequest
//
// @return ListScriptsResponse
func (client *Client) ListScripts(request *ListScriptsRequest) (_result *ListScriptsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScriptsResponse{}
	_body, _err := client.ListScriptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取变量列表
//
// @param request - ListVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVariableResponse
func (client *Client) ListVariableWithOptions(request *ListVariableRequest, runtime *dara.RuntimeOptions) (_result *ListVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchPattern) {
		body["SearchPattern"] = request.SearchPattern
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVariable"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取变量列表
//
// @param request - ListVariableRequest
//
// @return ListVariableResponse
func (client *Client) ListVariable(request *ListVariableRequest) (_result *ListVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVariableResponse{}
	_body, _err := client.ListVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListVocabularyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVocabularyResponse
func (client *Client) ListVocabularyWithOptions(request *ListVocabularyRequest, runtime *dara.RuntimeOptions) (_result *ListVocabularyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVocabulary"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVocabularyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListVocabularyRequest
//
// @return ListVocabularyResponse
func (client *Client) ListVocabulary(request *ListVocabularyRequest) (_result *ListVocabularyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVocabularyResponse{}
	_body, _err := client.ListVocabularyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListVoiceAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVoiceAccessProfileResponse
func (client *Client) ListVoiceAccessProfileWithOptions(request *ListVoiceAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *ListVoiceAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVoiceAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVoiceAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListVoiceAccessProfileRequest
//
// @return ListVoiceAccessProfileResponse
func (client *Client) ListVoiceAccessProfile(request *ListVoiceAccessProfileRequest) (_result *ListVoiceAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVoiceAccessProfileResponse{}
	_body, _err := client.ListVoiceAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取引擎列表
//
// @param request - ListVoiceEnginesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVoiceEnginesResponse
func (client *Client) ListVoiceEnginesWithOptions(request *ListVoiceEnginesRequest, runtime *dara.RuntimeOptions) (_result *ListVoiceEnginesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVoiceEngines"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVoiceEnginesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取引擎列表
//
// @param request - ListVoiceEnginesRequest
//
// @return ListVoiceEnginesResponse
func (client *Client) ListVoiceEngines(request *ListVoiceEnginesRequest) (_result *ListVoiceEnginesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVoiceEnginesResponse{}
	_body, _err := client.ListVoiceEnginesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListVoicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVoicesResponse
func (client *Client) ListVoicesWithOptions(request *ListVoicesRequest, runtime *dara.RuntimeOptions) (_result *ListVoicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NlsAccessType) {
		body["NlsAccessType"] = request.NlsAccessType
	}

	if !dara.IsNil(request.NlsEngine) {
		body["NlsEngine"] = request.NlsEngine
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVoices"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVoicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例详情
//
// @param request - ListVoicesRequest
//
// @return ListVoicesResponse
func (client *Client) ListVoices(request *ListVoicesRequest) (_result *ListVoicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVoicesResponse{}
	_body, _err := client.ListVoicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 试听
//
// @param tmpReq - PreviewVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewVoiceResponse
func (client *Client) PreviewVoiceWithOptions(tmpReq *PreviewVoiceRequest, runtime *dara.RuntimeOptions) (_result *PreviewVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PreviewVoiceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.NlsAccessType) {
		body["NlsAccessType"] = request.NlsAccessType
	}

	if !dara.IsNil(request.NlsEngine) {
		body["NlsEngine"] = request.NlsEngine
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	if !dara.IsNil(request.Text) {
		body["Text"] = request.Text
	}

	if !dara.IsNil(request.Voice) {
		body["Voice"] = request.Voice
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewVoice"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 试听
//
// @param request - PreviewVoiceRequest
//
// @return PreviewVoiceResponse
func (client *Client) PreviewVoice(request *PreviewVoiceRequest) (_result *PreviewVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreviewVoiceResponse{}
	_body, _err := client.PreviewVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - PublishScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishScriptResponse
func (client *Client) PublishScriptWithOptions(request *PublishScriptRequest, runtime *dara.RuntimeOptions) (_result *PublishScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - PublishScriptRequest
//
// @return PublishScriptResponse
func (client *Client) PublishScript(request *PublishScriptRequest) (_result *PublishScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishScriptResponse{}
	_body, _err := client.PublishScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - UpdateCloneVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloneVoiceResponse
func (client *Client) UpdateCloneVoiceWithOptions(request *UpdateCloneVoiceRequest, runtime *dara.RuntimeOptions) (_result *UpdateCloneVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CloneVoiceId) {
		body["CloneVoiceId"] = request.CloneVoiceId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCloneVoice"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCloneVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - UpdateCloneVoiceRequest
//
// @return UpdateCloneVoiceResponse
func (client *Client) UpdateCloneVoice(request *UpdateCloneVoiceRequest) (_result *UpdateCloneVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCloneVoiceResponse{}
	_body, _err := client.UpdateCloneVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - UpdateScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScriptResponse
func (client *Client) UpdateScriptWithOptions(request *UpdateScriptRequest, runtime *dara.RuntimeOptions) (_result *UpdateScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScript"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - UpdateScriptRequest
//
// @return UpdateScriptResponse
func (client *Client) UpdateScript(request *UpdateScriptRequest) (_result *UpdateScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateScriptResponse{}
	_body, _err := client.UpdateScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建或更新MQ配置
//
// @param tmpReq - UpdateSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscriptionWithOptions(tmpReq *UpdateSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *UpdateSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSubscriptionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EventSubscriptions) {
		request.EventSubscriptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventSubscriptions, dara.String("EventSubscriptions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Endpoint) {
		body["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.EventSubscriptionsShrink) {
		body["EventSubscriptions"] = request.EventSubscriptionsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MqInstanceId) {
		body["MqInstanceId"] = request.MqInstanceId
	}

	if !dara.IsNil(request.MqType) {
		body["MqType"] = request.MqType
	}

	if !dara.IsNil(request.Password) {
		body["Password"] = request.Password
	}

	if !dara.IsNil(request.ProducerId) {
		body["ProducerId"] = request.ProducerId
	}

	if !dara.IsNil(request.Topic) {
		body["Topic"] = request.Topic
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSubscription"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建或更新MQ配置
//
// @param request - UpdateSubscriptionRequest
//
// @return UpdateSubscriptionResponse
func (client *Client) UpdateSubscription(request *UpdateSubscriptionRequest) (_result *UpdateSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSubscriptionResponse{}
	_body, _err := client.UpdateSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新变量
//
// @param request - UpdateVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVariableResponse
func (client *Client) UpdateVariableWithOptions(request *UpdateVariableRequest, runtime *dara.RuntimeOptions) (_result *UpdateVariableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VariableId) {
		body["VariableId"] = request.VariableId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVariable"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新变量
//
// @param request - UpdateVariableRequest
//
// @return UpdateVariableResponse
func (client *Client) UpdateVariable(request *UpdateVariableRequest) (_result *UpdateVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVariableResponse{}
	_body, _err := client.UpdateVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param tmpReq - UpdateVocabularyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVocabularyResponse
func (client *Client) UpdateVocabularyWithOptions(tmpReq *UpdateVocabularyRequest, runtime *dara.RuntimeOptions) (_result *UpdateVocabularyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateVocabularyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Words) {
		request.WordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Words, dara.String("Words"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.VocabularyId) {
		body["VocabularyId"] = request.VocabularyId
	}

	if !dara.IsNil(request.WordsShrink) {
		body["Words"] = request.WordsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVocabulary"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVocabularyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - UpdateVocabularyRequest
//
// @return UpdateVocabularyResponse
func (client *Client) UpdateVocabulary(request *UpdateVocabularyRequest) (_result *UpdateVocabularyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVocabularyResponse{}
	_body, _err := client.UpdateVocabularyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param tmpReq - UpdateVoiceAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVoiceAccessProfileResponse
func (client *Client) UpdateVoiceAccessProfileWithOptions(tmpReq *UpdateVoiceAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateVoiceAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateVoiceAccessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Profile) {
		request.ProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Profile, dara.String("Profile"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessProfileId) {
		body["AccessProfileId"] = request.AccessProfileId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NlsEngine) {
		body["NlsEngine"] = request.NlsEngine
	}

	if !dara.IsNil(request.ProfileShrink) {
		body["Profile"] = request.ProfileShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVoiceAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVoiceAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例
//
// @param request - UpdateVoiceAccessProfileRequest
//
// @return UpdateVoiceAccessProfileResponse
func (client *Client) UpdateVoiceAccessProfile(request *UpdateVoiceAccessProfileRequest) (_result *UpdateVoiceAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVoiceAccessProfileResponse{}
	_body, _err := client.UpdateVoiceAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
