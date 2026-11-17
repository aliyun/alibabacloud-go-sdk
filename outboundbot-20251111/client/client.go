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
	client.ProductId = dara.String("OutboundBot")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("outboundbot"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Stops an outbound campaign.
//
// Description:
//
// ***
//
// @param request - AbortCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortCampaignResponse
func (client *Client) AbortCampaignWithOptions(request *AbortCampaignRequest, runtime *dara.RuntimeOptions) (_result *AbortCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortCampaign"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an outbound campaign.
//
// Description:
//
// ***
//
// @param request - AbortCampaignRequest
//
// @return AbortCampaignResponse
func (client *Client) AbortCampaign(request *AbortCampaignRequest) (_result *AbortCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AbortCampaignResponse{}
	_body, _err := client.AbortCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates outbound call cases.
//
// Description:
//
// ***
//
// @param tmpReq - AbortCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortCasesResponse
func (client *Client) AbortCasesWithOptions(tmpReq *AbortCasesRequest, runtime *dara.RuntimeOptions) (_result *AbortCasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AbortCasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PhoneNumbers) {
		request.PhoneNumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneNumbers, dara.String("PhoneNumbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PhoneNumbersShrink) {
		query["PhoneNumbers"] = request.PhoneNumbersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortCases"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortCasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates outbound call cases.
//
// Description:
//
// ***
//
// @param request - AbortCasesRequest
//
// @return AbortCasesResponse
func (client *Client) AbortCases(request *AbortCasesRequest) (_result *AbortCasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AbortCasesResponse{}
	_body, _err := client.AbortCasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Appends outbound call cases.
//
// Description:
//
// ***
//
// @param tmpReq - AppendCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppendCasesResponse
func (client *Client) AppendCasesWithOptions(tmpReq *AppendCasesRequest, runtime *dara.RuntimeOptions) (_result *AppendCasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AppendCasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Cases) {
		request.CasesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cases, dara.String("Cases"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CasesShrink) {
		body["Cases"] = request.CasesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AppendCases"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AppendCasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Appends outbound call cases.
//
// Description:
//
// ***
//
// @param request - AppendCasesRequest
//
// @return AppendCasesResponse
func (client *Client) AppendCases(request *AppendCasesRequest) (_result *AppendCasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AppendCasesResponse{}
	_body, _err := client.AppendCasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a text conversation session for a specified instance and scenario, and streams the initial reply from the robot.
//
// Description:
//
// Creates a text conversation session for a specified instance and scenario, and streams the initial reply from the robot. An available scenario version must exist before you call this operation. Use DraftVersion to select the draft or published version, and use VendorParams to pass the accompanying variables required by the scenario. The response uses Server-Sent Events (SSE), and each data event carries a business response object. Save the returned SessionId for subsequent Dialogue and EndSession calls. A StreamEnd value of true indicates that the current response stream has ended, but the session remains active. The resource identifiers, phone numbers, and authentication credentials in the examples are masked or placeholder values. Use actual resources and service return values in production calls.
//
// @param request - BeginSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeginSessionResponse
func (client *Client) BeginSessionWithSSE(request *BeginSessionRequest, runtime *dara.RuntimeOptions, _yield chan *BeginSessionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.beginSessionWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// Creates a text conversation session for a specified instance and scenario, and streams the initial reply from the robot.
//
// Description:
//
// Creates a text conversation session for a specified instance and scenario, and streams the initial reply from the robot. An available scenario version must exist before you call this operation. Use DraftVersion to select the draft or published version, and use VendorParams to pass the accompanying variables required by the scenario. The response uses Server-Sent Events (SSE), and each data event carries a business response object. Save the returned SessionId for subsequent Dialogue and EndSession calls. A StreamEnd value of true indicates that the current response stream has ended, but the session remains active. The resource identifiers, phone numbers, and authentication credentials in the examples are masked or placeholder values. Use actual resources and service return values in production calls.
//
// @param request - BeginSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeginSessionResponse
func (client *Client) BeginSessionWithOptions(request *BeginSessionRequest, runtime *dara.RuntimeOptions) (_result *BeginSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DraftVersion) {
		query["DraftVersion"] = request.DraftVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.VendorParams) {
		query["VendorParams"] = request.VendorParams
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BeginSession"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BeginSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a text conversation session for a specified instance and scenario, and streams the initial reply from the robot.
//
// Description:
//
// Creates a text conversation session for a specified instance and scenario, and streams the initial reply from the robot. An available scenario version must exist before you call this operation. Use DraftVersion to select the draft or published version, and use VendorParams to pass the accompanying variables required by the scenario. The response uses Server-Sent Events (SSE), and each data event carries a business response object. Save the returned SessionId for subsequent Dialogue and EndSession calls. A StreamEnd value of true indicates that the current response stream has ended, but the session remains active. The resource identifiers, phone numbers, and authentication credentials in the examples are masked or placeholder values. Use actual resources and service return values in production calls.
//
// @param request - BeginSessionRequest
//
// @return BeginSessionResponse
func (client *Client) BeginSession(request *BeginSessionRequest) (_result *BeginSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BeginSessionResponse{}
	_body, _err := client.BeginSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a trial session.
//
// Description:
//
// Creates a trial text session in a specified instance and scenario by using the VisitorToken from a trial link, and returns the first-round reply in Server-Sent Events (SSE) format. The service verifies the channel and expiration time corresponding to the token, and then starts the session based on the tenant to which the channel belongs. The token does not replace the business parameters of the scenario and instance. Save the returned SessionId for subsequent calls to TrialDialogue to send user input, and call EndTrialSession to end the session. StreamEnd=true only indicates that the streaming response for the current round has ended. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - BeginTrialSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeginTrialSessionResponse
func (client *Client) BeginTrialSessionWithSSE(request *BeginTrialSessionRequest, runtime *dara.RuntimeOptions, _yield chan *BeginTrialSessionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.beginTrialSessionWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// Initiates a trial session.
//
// Description:
//
// Creates a trial text session in a specified instance and scenario by using the VisitorToken from a trial link, and returns the first-round reply in Server-Sent Events (SSE) format. The service verifies the channel and expiration time corresponding to the token, and then starts the session based on the tenant to which the channel belongs. The token does not replace the business parameters of the scenario and instance. Save the returned SessionId for subsequent calls to TrialDialogue to send user input, and call EndTrialSession to end the session. StreamEnd=true only indicates that the streaming response for the current round has ended. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - BeginTrialSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BeginTrialSessionResponse
func (client *Client) BeginTrialSessionWithOptions(request *BeginTrialSessionRequest, runtime *dara.RuntimeOptions) (_result *BeginTrialSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DraftVersion) {
		query["DraftVersion"] = request.DraftVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.VendorParams) {
		query["VendorParams"] = request.VendorParams
	}

	if !dara.IsNil(request.VisitorToken) {
		query["VisitorToken"] = request.VisitorToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BeginTrialSession"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BeginTrialSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a trial session.
//
// Description:
//
// Creates a trial text session in a specified instance and scenario by using the VisitorToken from a trial link, and returns the first-round reply in Server-Sent Events (SSE) format. The service verifies the channel and expiration time corresponding to the token, and then starts the session based on the tenant to which the channel belongs. The token does not replace the business parameters of the scenario and instance. Save the returned SessionId for subsequent calls to TrialDialogue to send user input, and call EndTrialSession to end the session. StreamEnd=true only indicates that the streaming response for the current round has ended. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - BeginTrialSessionRequest
//
// @return BeginTrialSessionResponse
func (client *Client) BeginTrialSession(request *BeginTrialSessionRequest) (_result *BeginTrialSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BeginTrialSessionResponse{}
	_body, _err := client.BeginTrialSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a trial call.
//
// Description:
//
// Initiates a trial web voice call for a specified scenario by using the VisitorToken of a trial link. The service verifies the token, the associated channel, and the validity period, and then delegates the web call service to bridge the web voice channel and the robot channel. The DeviceId must be consistent with the data channel device identifier of the trial client. After the call is initiated, use the returned ChannelId, ServerUrl, and Token to connect to the voice channel. If the data event channel of the trial client is required, you can call GetTrialDataChannelCredential with the same DeviceId. The web voice connection token and the data channel username and password are not the same set of credentials. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - BridgeTrialWebCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BridgeTrialWebCallResponse
func (client *Client) BridgeTrialWebCallWithOptions(request *BridgeTrialWebCallRequest, runtime *dara.RuntimeOptions) (_result *BridgeTrialWebCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DraftVersion) {
		query["DraftVersion"] = request.DraftVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SampleRate) {
		query["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	if !dara.IsNil(request.VisitorToken) {
		query["VisitorToken"] = request.VisitorToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BridgeTrialWebCall"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BridgeTrialWebCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a trial call.
//
// Description:
//
// Initiates a trial web voice call for a specified scenario by using the VisitorToken of a trial link. The service verifies the token, the associated channel, and the validity period, and then delegates the web call service to bridge the web voice channel and the robot channel. The DeviceId must be consistent with the data channel device identifier of the trial client. After the call is initiated, use the returned ChannelId, ServerUrl, and Token to connect to the voice channel. If the data event channel of the trial client is required, you can call GetTrialDataChannelCredential with the same DeviceId. The web voice connection token and the data channel username and password are not the same set of credentials. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - BridgeTrialWebCallRequest
//
// @return BridgeTrialWebCallResponse
func (client *Client) BridgeTrialWebCall(request *BridgeTrialWebCallRequest) (_result *BridgeTrialWebCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BridgeTrialWebCallResponse{}
	_body, _err := client.BridgeTrialWebCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates a call.
//
// Description:
//
// Initiates a web voice call for a specified instance and scenario by bridging the web voice channel and the robot SIP channel through an RTC gateway. Pass information such as DeviceId, audio sampling rate, and scenario version in the request. TimeoutSeconds specifies the timeout for unanswered calls, defaults to 60 seconds, and does not indicate the maximum duration of an answered call. Upon success, returns the connection context such as SessionId, ChannelId, ServerUrl, and Token. The client must use the returned address and valid credentials to establish the web voice connection. This operation initiates a web interactive voice response instead of dialing a real phone number specified by Number. Resource identity, numbers, and authentication credentials in the examples are masked or placeholder values. Use actual resources and service return values when you invoke the operation.
//
// @param request - BridgeWebCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BridgeWebCallResponse
func (client *Client) BridgeWebCallWithOptions(request *BridgeWebCallRequest, runtime *dara.RuntimeOptions) (_result *BridgeWebCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DraftVersion) {
		query["DraftVersion"] = request.DraftVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SampleRate) {
		query["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		query["TimeoutSeconds"] = request.TimeoutSeconds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BridgeWebCall"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BridgeWebCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a call.
//
// Description:
//
// Initiates a web voice call for a specified instance and scenario by bridging the web voice channel and the robot SIP channel through an RTC gateway. Pass information such as DeviceId, audio sampling rate, and scenario version in the request. TimeoutSeconds specifies the timeout for unanswered calls, defaults to 60 seconds, and does not indicate the maximum duration of an answered call. Upon success, returns the connection context such as SessionId, ChannelId, ServerUrl, and Token. The client must use the returned address and valid credentials to establish the web voice connection. This operation initiates a web interactive voice response instead of dialing a real phone number specified by Number. Resource identity, numbers, and authentication credentials in the examples are masked or placeholder values. Use actual resources and service return values when you invoke the operation.
//
// @param request - BridgeWebCallRequest
//
// @return BridgeWebCallResponse
func (client *Client) BridgeWebCall(request *BridgeWebCallRequest) (_result *BridgeWebCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BridgeWebCallResponse{}
	_body, _err := client.BridgeWebCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an outbound campaign.
//
// Description:
//
// ***
//
// @param tmpReq - CreateCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCampaignResponse
func (client *Client) CreateCampaignWithOptions(tmpReq *CreateCampaignRequest, runtime *dara.RuntimeOptions) (_result *CreateCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCampaignShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Cases) {
		request.CasesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cases, dara.String("Cases"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttemptOrder) {
		query["AttemptOrder"] = request.AttemptOrder
	}

	if !dara.IsNil(request.CallableTime) {
		query["CallableTime"] = request.CallableTime
	}

	if !dara.IsNil(request.CaseFileKey) {
		query["CaseFileKey"] = request.CaseFileKey
	}

	if !dara.IsNil(request.DialingTimeoutSeconds) {
		query["DialingTimeoutSeconds"] = request.DialingTimeoutSeconds
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FixedQuota) {
		query["FixedQuota"] = request.FixedQuota
	}

	if !dara.IsNil(request.FlashSmsParameters) {
		query["FlashSmsParameters"] = request.FlashSmsParameters
	}

	if !dara.IsNil(request.HolidayRestricted) {
		query["HolidayRestricted"] = request.HolidayRestricted
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxAttemptCount) {
		query["MaxAttemptCount"] = request.MaxAttemptCount
	}

	if !dara.IsNil(request.MinAttemptInterval) {
		query["MinAttemptInterval"] = request.MinAttemptInterval
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
	}

	if !dara.IsNil(request.RedialRestrictions) {
		query["RedialRestrictions"] = request.RedialRestrictions
	}

	if !dara.IsNil(request.RunUntilEndTime) {
		query["RunUntilEndTime"] = request.RunUntilEndTime
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CasesShrink) {
		body["Cases"] = request.CasesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCampaign"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an outbound campaign.
//
// Description:
//
// ***
//
// @param request - CreateCampaignRequest
//
// @return CreateCampaignResponse
func (client *Client) CreateCampaign(request *CreateCampaignRequest) (_result *CreateCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCampaignResponse{}
	_body, _err := client.CreateCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instance.
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
// Creates an instance.
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
// Creates a flash message configuration.
//
// @param tmpReq - CreateFlashSmsAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlashSmsAccessProfileResponse
func (client *Client) CreateFlashSmsAccessProfileWithOptions(tmpReq *CreateFlashSmsAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateFlashSmsAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFlashSmsAccessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccessProfile) {
		request.AccessProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccessProfile, dara.String("AccessProfile"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessProfileShrink) {
		body["AccessProfile"] = request.AccessProfileShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProviderId) {
		body["ProviderId"] = request.ProviderId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlashSmsAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlashSmsAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a flash message configuration.
//
// @param request - CreateFlashSmsAccessProfileRequest
//
// @return CreateFlashSmsAccessProfileResponse
func (client *Client) CreateFlashSmsAccessProfile(request *CreateFlashSmsAccessProfileRequest) (_result *CreateFlashSmsAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFlashSmsAccessProfileResponse{}
	_body, _err := client.CreateFlashSmsAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instance.
//
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.Concurrency) {
		body["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ServiceMode) {
		body["ServiceMode"] = request.ServiceMode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2025-11-11"),
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

// Summary:
//
// Creates an instance.
//
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
// Creates a language model access configuration.
//
// Description:
//
// Creates a language model access configuration for a specified instance.
//
// @param tmpReq - CreateLlmAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLlmAccessProfileResponse
func (client *Client) CreateLlmAccessProfileWithOptions(tmpReq *CreateLlmAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateLlmAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLlmAccessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Profile) {
		request.ProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Profile, dara.String("Profile"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProfileShrink) {
		body["Profile"] = request.ProfileShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLlmAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLlmAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a language model access configuration.
//
// Description:
//
// Creates a language model access configuration for a specified instance.
//
// @param request - CreateLlmAccessProfileRequest
//
// @return CreateLlmAccessProfileResponse
func (client *Client) CreateLlmAccessProfile(request *CreateLlmAccessProfileRequest) (_result *CreateLlmAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLlmAccessProfileResponse{}
	_body, _err := client.CreateLlmAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instance.
//
// @param tmpReq - CreateOutboundCallRestrictionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOutboundCallRestrictionResponse
func (client *Client) CreateOutboundCallRestrictionWithOptions(tmpReq *CreateOutboundCallRestrictionRequest, runtime *dara.RuntimeOptions) (_result *CreateOutboundCallRestrictionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOutboundCallRestrictionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OutboundCallRestriction) {
		request.OutboundCallRestrictionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutboundCallRestriction, dara.String("OutboundCallRestriction"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OutboundCallRestrictionShrink) {
		body["OutboundCallRestriction"] = request.OutboundCallRestrictionShrink
	}

	if !dara.IsNil(request.Policy) {
		body["Policy"] = request.Policy
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOutboundCallRestriction"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOutboundCallRestrictionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance.
//
// @param request - CreateOutboundCallRestrictionRequest
//
// @return CreateOutboundCallRestrictionResponse
func (client *Client) CreateOutboundCallRestriction(request *CreateOutboundCallRestrictionRequest) (_result *CreateOutboundCallRestrictionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOutboundCallRestrictionResponse{}
	_body, _err := client.CreateOutboundCallRestrictionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instance.
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
// Creates an instance.
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
// Creates a script configuration.
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

	if !dara.IsNil(tmpReq.LabelConfigs) {
		request.LabelConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelConfigs, dara.String("LabelConfigs"), dara.String("json"))
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

	if !dara.IsNil(request.LabelConfigsShrink) {
		body["LabelConfigs"] = request.LabelConfigsShrink
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
// Creates a script configuration.
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
// Creates an instance.
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
// Creates an instance.
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
// Deletes a scene.
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
// Deletes a scene.
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
// Deletes a tenant number.
//
// Description:
//
// Deletes a tenant number.
//
// @param request - DeleteCorpNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCorpNumberResponse
func (client *Client) DeleteCorpNumberWithOptions(request *DeleteCorpNumberRequest, runtime *dara.RuntimeOptions) (_result *DeleteCorpNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	if !dara.IsNil(request.Number) {
		body["Number"] = request.Number
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCorpNumber"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCorpNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a tenant number.
//
// Description:
//
// Deletes a tenant number.
//
// @param request - DeleteCorpNumberRequest
//
// @return DeleteCorpNumberResponse
func (client *Client) DeleteCorpNumber(request *DeleteCorpNumberRequest) (_result *DeleteCorpNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCorpNumberResponse{}
	_body, _err := client.DeleteCorpNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a flash message configuration.
//
// @param request - DeleteFlashSmsAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlashSmsAccessProfileResponse
func (client *Client) DeleteFlashSmsAccessProfileWithOptions(request *DeleteFlashSmsAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlashSmsAccessProfileResponse, _err error) {
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
		Action:      dara.String("DeleteFlashSmsAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlashSmsAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a flash message configuration.
//
// @param request - DeleteFlashSmsAccessProfileRequest
//
// @return DeleteFlashSmsAccessProfileResponse
func (client *Client) DeleteFlashSmsAccessProfile(request *DeleteFlashSmsAccessProfileRequest) (_result *DeleteFlashSmsAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFlashSmsAccessProfileResponse{}
	_body, _err := client.DeleteFlashSmsAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves instance details.
//
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2025-11-11"),
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

// Summary:
//
// Retrieves instance details.
//
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

// Summary:
//
// Deletes the language model access configuration.
//
// Description:
//
// Deletes the language model access configuration for a specified instance.
//
// @param request - DeleteLlmAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLlmAccessProfileResponse
func (client *Client) DeleteLlmAccessProfileWithOptions(request *DeleteLlmAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *DeleteLlmAccessProfileResponse, _err error) {
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
		Action:      dara.String("DeleteLlmAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLlmAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the language model access configuration.
//
// Description:
//
// Deletes the language model access configuration for a specified instance.
//
// @param request - DeleteLlmAccessProfileRequest
//
// @return DeleteLlmAccessProfileResponse
func (client *Client) DeleteLlmAccessProfile(request *DeleteLlmAccessProfileRequest) (_result *DeleteLlmAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLlmAccessProfileResponse{}
	_body, _err := client.DeleteLlmAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the numbers of an instance.
//
// Description:
//
// Deletes the numbers of a specified instance.
//
// @param tmpReq - DeleteNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNumbersResponse
func (client *Client) DeleteNumbersWithOptions(tmpReq *DeleteNumbersRequest, runtime *dara.RuntimeOptions) (_result *DeleteNumbersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteNumbersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NumbersShrink) {
		body["Numbers"] = request.NumbersShrink
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNumbers"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the numbers of an instance.
//
// Description:
//
// Deletes the numbers of a specified instance.
//
// @param request - DeleteNumbersRequest
//
// @return DeleteNumbersResponse
func (client *Client) DeleteNumbers(request *DeleteNumbersRequest) (_result *DeleteNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNumbersResponse{}
	_body, _err := client.DeleteNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes blacklists and whitelists.
//
// @param tmpReq - DeleteOutboundCallRestrictionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOutboundCallRestrictionResponse
func (client *Client) DeleteOutboundCallRestrictionWithOptions(tmpReq *DeleteOutboundCallRestrictionRequest, runtime *dara.RuntimeOptions) (_result *DeleteOutboundCallRestrictionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteOutboundCallRestrictionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RestrictionIdList) {
		request.RestrictionIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RestrictionIdList, dara.String("RestrictionIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RestrictionIdListShrink) {
		body["RestrictionIdList"] = request.RestrictionIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOutboundCallRestriction"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOutboundCallRestrictionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes blacklists and whitelists.
//
// @param request - DeleteOutboundCallRestrictionRequest
//
// @return DeleteOutboundCallRestrictionResponse
func (client *Client) DeleteOutboundCallRestriction(request *DeleteOutboundCallRestrictionRequest) (_result *DeleteOutboundCallRestrictionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOutboundCallRestrictionResponse{}
	_body, _err := client.DeleteOutboundCallRestrictionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scenario.
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
// Deletes a scenario.
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
// Deletes a third-party voice configuration.
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
// Deletes a third-party voice configuration.
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
// Initiates a session.
//
// Description:
//
// Sends a round of user input to a text session created by BeginSession and returns the robot reply and control instructions in Server-Sent Events (SSE) format. The InstanceId, ScriptId, and SessionId must be consistent with those used when the session was initiated. The SessionId must use the value returned by the session initiation operation. The conversation cannot continue if the session context does not exist. Utterance is used for plain text input, and Extras is used to pass structured events such as key presses and silence. You can call this operation multiple times in succession. StreamId identifies the reply round, and StreamEnd=true indicates that the output of the current round has ended. Call EndSession after the entire interaction is completed. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - DialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DialogueResponse
func (client *Client) DialogueWithSSE(request *DialogueRequest, runtime *dara.RuntimeOptions, _yield chan *DialogueResponse, _yieldErr chan error) {
	defer close(_yield)
	client.dialogueWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// Initiates a session.
//
// Description:
//
// Sends a round of user input to a text session created by BeginSession and returns the robot reply and control instructions in Server-Sent Events (SSE) format. The InstanceId, ScriptId, and SessionId must be consistent with those used when the session was initiated. The SessionId must use the value returned by the session initiation operation. The conversation cannot continue if the session context does not exist. Utterance is used for plain text input, and Extras is used to pass structured events such as key presses and silence. You can call this operation multiple times in succession. StreamId identifies the reply round, and StreamEnd=true indicates that the output of the current round has ended. Call EndSession after the entire interaction is completed. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
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
	if !dara.IsNil(request.Extras) {
		query["Extras"] = request.Extras
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Dialogue"),
		Version:     dara.String("2025-11-11"),
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

// Summary:
//
// Initiates a session.
//
// Description:
//
// Sends a round of user input to a text session created by BeginSession and returns the robot reply and control instructions in Server-Sent Events (SSE) format. The InstanceId, ScriptId, and SessionId must be consistent with those used when the session was initiated. The SessionId must use the value returned by the session initiation operation. The conversation cannot continue if the session context does not exist. Utterance is used for plain text input, and Extras is used to pass structured events such as key presses and silence. You can call this operation multiple times in succession. StreamId identifies the reply round, and StreamEnd=true indicates that the output of the current round has ended. Call EndSession after the entire interaction is completed. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
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

// Summary:
//
// Disables message subscription.
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
// Disables message subscription.
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
// Ends a session.
//
// Description:
//
// Ends a text session created by BeginSession, notifies the dialog engine to end the interaction, and clears the text session state. The InstanceId, ScriptId, and SessionId must match those used when the session was started. The Data field in the response returns the SessionId of the ended session. If the text session no longer exists due to prolonged inactivity or cache expiration, the current implementation still returns the SessionId without reporting a session not found error. This operation is used to conclude a text session and is not intended for hanging up web voice or actual phone calls. The resource identifiers, phone numbers, and authentication credentials in the examples are masked or placeholder values. Use actual resources and service return values for real calls.
//
// @param request - EndSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndSessionResponse
func (client *Client) EndSessionWithOptions(request *EndSessionRequest, runtime *dara.RuntimeOptions) (_result *EndSessionResponse, _err error) {
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

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EndSession"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EndSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ends a session.
//
// Description:
//
// Ends a text session created by BeginSession, notifies the dialog engine to end the interaction, and clears the text session state. The InstanceId, ScriptId, and SessionId must match those used when the session was started. The Data field in the response returns the SessionId of the ended session. If the text session no longer exists due to prolonged inactivity or cache expiration, the current implementation still returns the SessionId without reporting a session not found error. This operation is used to conclude a text session and is not intended for hanging up web voice or actual phone calls. The resource identifiers, phone numbers, and authentication credentials in the examples are masked or placeholder values. Use actual resources and service return values for real calls.
//
// @param request - EndSessionRequest
//
// @return EndSessionResponse
func (client *Client) EndSession(request *EndSessionRequest) (_result *EndSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EndSessionResponse{}
	_body, _err := client.EndSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Ends a trial session.
//
// Description:
//
// Ends a trial text session created by BeginTrialSession by using a valid VisitorToken from a trial link. The service first validates the token and then delegates the request to the text session termination service. The InstanceId, ScriptId, and SessionId must maintain session persistence consistent with those of the trial session. The returned Data is the SessionId of the ended session. If the text session context has expired, the termination service can still return the SessionId, but the validation will be marked as failed if the trial token itself is invalid or expired. This operation is not used to hang up web voice calls. The resource identity, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values when you invoke the operation.
//
// @param request - EndTrialSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndTrialSessionResponse
func (client *Client) EndTrialSessionWithOptions(request *EndTrialSessionRequest, runtime *dara.RuntimeOptions) (_result *EndTrialSessionResponse, _err error) {
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

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.VisitorToken) {
		query["VisitorToken"] = request.VisitorToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EndTrialSession"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EndTrialSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ends a trial session.
//
// Description:
//
// Ends a trial text session created by BeginTrialSession by using a valid VisitorToken from a trial link. The service first validates the token and then delegates the request to the text session termination service. The InstanceId, ScriptId, and SessionId must maintain session persistence consistent with those of the trial session. The returned Data is the SessionId of the ended session. If the text session context has expired, the termination service can still return the SessionId, but the validation will be marked as failed if the trial token itself is invalid or expired. This operation is not used to hang up web voice calls. The resource identity, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values when you invoke the operation.
//
// @param request - EndTrialSessionRequest
//
// @return EndTrialSessionResponse
func (client *Client) EndTrialSession(request *EndTrialSessionRequest) (_result *EndTrialSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EndTrialSessionResponse{}
	_body, _err := client.EndTrialSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the call details.
//
// Description:
//
// ***
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
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
// Retrieves the call details.
//
// Description:
//
// ***
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
// Retrieves the details of an outbound campaign.
//
// @param request - GetCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCampaignResponse
func (client *Client) GetCampaignWithOptions(request *GetCampaignRequest, runtime *dara.RuntimeOptions) (_result *GetCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCampaign"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an outbound campaign.
//
// @param request - GetCampaignRequest
//
// @return GetCampaignResponse
func (client *Client) GetCampaign(request *GetCampaignRequest) (_result *GetCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCampaignResponse{}
	_body, _err := client.GetCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a case.
//
// Description:
//
// ***
//
// @param request - GetCaseDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCaseDetailResponse
func (client *Client) GetCaseDetailWithOptions(request *GetCaseDetailRequest, runtime *dara.RuntimeOptions) (_result *GetCaseDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaseId) {
		query["CaseId"] = request.CaseId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCaseDetail"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCaseDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a case.
//
// Description:
//
// ***
//
// @param request - GetCaseDetailRequest
//
// @return GetCaseDetailResponse
func (client *Client) GetCaseDetail(request *GetCaseDetailRequest) (_result *GetCaseDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCaseDetailResponse{}
	_body, _err := client.GetCaseDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves voice conversation parameters.
//
// Description:
//
// Retrieves temporary connection credentials for the client data channel of a specified InstanceId and DeviceId. Requests are verified based on the current POP authentication and instance permission configurations. DeviceId cannot be empty. Unlike GetTrialDataChannelCredential, this operation does not use the VisitorToken for trial links. The client must use the actual Topic, ClientId, Endpoint, UserName, and Password returned to connect to and subscribe to the data channel, and must not infer the address, topic, or authentication information. The same client must keep the DeviceId consistent when initiating BridgeWebCall or MakePilotCall. This operation only retrieves data channel credentials and does not create text sessions, initiate phone calls, or establish web voice channels. The username and password for the data channel are not the same set of credentials as the token returned for web calls. The protocol, the unit of ExpirationTime, and field mappings are subject to the actual RTC gateway deployment version and return conventions. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - GetDataChannelCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataChannelCredentialResponse
func (client *Client) GetDataChannelCredentialWithOptions(request *GetDataChannelCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetDataChannelCredentialResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataChannelCredential"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataChannelCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves voice conversation parameters.
//
// Description:
//
// Retrieves temporary connection credentials for the client data channel of a specified InstanceId and DeviceId. Requests are verified based on the current POP authentication and instance permission configurations. DeviceId cannot be empty. Unlike GetTrialDataChannelCredential, this operation does not use the VisitorToken for trial links. The client must use the actual Topic, ClientId, Endpoint, UserName, and Password returned to connect to and subscribe to the data channel, and must not infer the address, topic, or authentication information. The same client must keep the DeviceId consistent when initiating BridgeWebCall or MakePilotCall. This operation only retrieves data channel credentials and does not create text sessions, initiate phone calls, or establish web voice channels. The username and password for the data channel are not the same set of credentials as the token returned for web calls. The protocol, the unit of ExpirationTime, and field mappings are subject to the actual RTC gateway deployment version and return conventions. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - GetDataChannelCredentialRequest
//
// @return GetDataChannelCredentialResponse
func (client *Client) GetDataChannelCredential(request *GetDataChannelCredentialRequest) (_result *GetDataChannelCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataChannelCredentialResponse{}
	_body, _err := client.GetDataChannelCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an instance.
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
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
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an instance.
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a call recording.
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
// Retrieves a call recording.
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
// Retrieves a prompt scenario template.
//
// @param request - GetScriptProfileTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScriptProfileTemplateResponse
func (client *Client) GetScriptProfileTemplateWithOptions(request *GetScriptProfileTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetScriptProfileTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScriptProfileTemplate"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScriptProfileTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a prompt scenario template.
//
// @param request - GetScriptProfileTemplateRequest
//
// @return GetScriptProfileTemplateResponse
func (client *Client) GetScriptProfileTemplate(request *GetScriptProfileTemplateRequest) (_result *GetScriptProfileTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetScriptProfileTemplateResponse{}
	_body, _err := client.GetScriptProfileTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the MQ configuration.
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
// Retrieves the MQ configuration.
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
// Retrieves the trial channel.
//
// Description:
//
// Retrieves the corresponding trial access channel and the list of available trial scenarios based on the VisitorToken in a trial link. The service verifies whether the token exists, whether it can locate the channel, and whether it has expired. Upon success, it returns information such as the instance to which the channel belongs, the channel type, display configurations, validity period, and scenario configurations. This operation can be used as an initialization step for the trial page: first read the returned instance and scenarios, and then call the corresponding trial text or web voice operation. This operation does not create a text session or initiate a voice call. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - GetTrialAccessChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrialAccessChannelResponse
func (client *Client) GetTrialAccessChannelWithOptions(request *GetTrialAccessChannelRequest, runtime *dara.RuntimeOptions) (_result *GetTrialAccessChannelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.VisitorToken) {
		query["VisitorToken"] = request.VisitorToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrialAccessChannel"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrialAccessChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the trial channel.
//
// Description:
//
// Retrieves the corresponding trial access channel and the list of available trial scenarios based on the VisitorToken in a trial link. The service verifies whether the token exists, whether it can locate the channel, and whether it has expired. Upon success, it returns information such as the instance to which the channel belongs, the channel type, display configurations, validity period, and scenario configurations. This operation can be used as an initialization step for the trial page: first read the returned instance and scenarios, and then call the corresponding trial text or web voice operation. This operation does not create a text session or initiate a voice call. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - GetTrialAccessChannelRequest
//
// @return GetTrialAccessChannelResponse
func (client *Client) GetTrialAccessChannel(request *GetTrialAccessChannelRequest) (_result *GetTrialAccessChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTrialAccessChannelResponse{}
	_body, _err := client.GetTrialAccessChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the parameters for a trial voice conversation.
//
// Description:
//
// Retrieves temporary connection credentials for the client data channel for a specified InstanceId and DeviceId by using a valid VisitorToken from a trial link. The response includes information such as Topic, ClientId, Endpoint, UserName, Password, and expiration time. The client must use the actual values returned by the service to connect to and subscribe to the data channel, and must not concatenate addresses, topics, or authentication information. The DeviceId must be consistent with the BridgeTrialWebCall request from the same client. This operation only retrieves data channel credentials and does not initiate text sessions or establish web voice calls. The credential protocol and expiration time unit are subject to the conventions returned by the actual RTC gateway. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - GetTrialDataChannelCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrialDataChannelCredentialResponse
func (client *Client) GetTrialDataChannelCredentialWithOptions(request *GetTrialDataChannelCredentialRequest, runtime *dara.RuntimeOptions) (_result *GetTrialDataChannelCredentialResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.VisitorToken) {
		query["VisitorToken"] = request.VisitorToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrialDataChannelCredential"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrialDataChannelCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the parameters for a trial voice conversation.
//
// Description:
//
// Retrieves temporary connection credentials for the client data channel for a specified InstanceId and DeviceId by using a valid VisitorToken from a trial link. The response includes information such as Topic, ClientId, Endpoint, UserName, Password, and expiration time. The client must use the actual values returned by the service to connect to and subscribe to the data channel, and must not concatenate addresses, topics, or authentication information. The DeviceId must be consistent with the BridgeTrialWebCall request from the same client. This operation only retrieves data channel credentials and does not initiate text sessions or establish web voice calls. The credential protocol and expiration time unit are subject to the conventions returned by the actual RTC gateway. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - GetTrialDataChannelCredentialRequest
//
// @return GetTrialDataChannelCredentialResponse
func (client *Client) GetTrialDataChannelCredential(request *GetTrialDataChannelCredentialRequest) (_result *GetTrialDataChannelCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTrialDataChannelCredentialResponse{}
	_body, _err := client.GetTrialDataChannelCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of ASR models.
//
// @param request - ListAsrModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsrModelsResponse
func (client *Client) ListAsrModelsWithOptions(request *ListAsrModelsRequest, runtime *dara.RuntimeOptions) (_result *ListAsrModelsResponse, _err error) {
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
		Action:      dara.String("ListAsrModels"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAsrModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of ASR models.
//
// @param request - ListAsrModelsRequest
//
// @return ListAsrModelsResponse
func (client *Client) ListAsrModels(request *ListAsrModelsRequest) (_result *ListAsrModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAsrModelsResponse{}
	_body, _err := client.ListAsrModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of outbound campaigns.
//
// Description:
//
// ***
//
// @param request - ListCampaignsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCampaignsResponse
func (client *Client) ListCampaignsWithOptions(request *ListCampaignsRequest, runtime *dara.RuntimeOptions) (_result *ListCampaignsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActualStartTimeFrom) {
		query["ActualStartTimeFrom"] = request.ActualStartTimeFrom
	}

	if !dara.IsNil(request.ActualStartTimeTo) {
		query["ActualStartTimeTo"] = request.ActualStartTimeTo
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlannedStartTimeFrom) {
		query["PlannedStartTimeFrom"] = request.PlannedStartTimeFrom
	}

	if !dara.IsNil(request.PlannedStartTimeTo) {
		query["PlannedStartTimeTo"] = request.PlannedStartTimeTo
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCampaigns"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCampaignsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of outbound campaigns.
//
// Description:
//
// ***
//
// @param request - ListCampaignsRequest
//
// @return ListCampaignsResponse
func (client *Client) ListCampaigns(request *ListCampaignsRequest) (_result *ListCampaignsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCampaignsResponse{}
	_body, _err := client.ListCampaignsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of cases.
//
// Description:
//
// ***
//
// @param tmpReq - ListCasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCasesResponse
func (client *Client) ListCasesWithOptions(tmpReq *ListCasesRequest, runtime *dara.RuntimeOptions) (_result *ListCasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CaseIds) {
		request.CaseIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CaseIds, dara.String("CaseIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DispositionCodes) {
		request.DispositionCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DispositionCodes, dara.String("DispositionCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DispositionReasons) {
		request.DispositionReasonsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DispositionReasons, dara.String("DispositionReasons"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LabelSearch) {
		request.LabelSearchShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelSearch, dara.String("LabelSearch"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.States) {
		request.StatesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.States, dara.String("States"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessChannelId) {
		query["AccessChannelId"] = request.AccessChannelId
	}

	if !dara.IsNil(request.AccessChannelType) {
		query["AccessChannelType"] = request.AccessChannelType
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.CaseCompleted) {
		query["CaseCompleted"] = request.CaseCompleted
	}

	if !dara.IsNil(request.CaseIdsShrink) {
		query["CaseIds"] = request.CaseIdsShrink
	}

	if !dara.IsNil(request.DispositionCodesShrink) {
		query["DispositionCodes"] = request.DispositionCodesShrink
	}

	if !dara.IsNil(request.DispositionReasonsShrink) {
		query["DispositionReasons"] = request.DispositionReasonsShrink
	}

	if !dara.IsNil(request.DraftVersion) {
		query["DraftVersion"] = request.DraftVersion
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LabelSearchShrink) {
		query["LabelSearch"] = request.LabelSearchShrink
	}

	if !dara.IsNil(request.MaxRingingDuration) {
		query["MaxRingingDuration"] = request.MaxRingingDuration
	}

	if !dara.IsNil(request.MaxTalkTime) {
		query["MaxTalkTime"] = request.MaxTalkTime
	}

	if !dara.IsNil(request.MaxTalkTurns) {
		query["MaxTalkTurns"] = request.MaxTalkTurns
	}

	if !dara.IsNil(request.MinRingingDuration) {
		query["MinRingingDuration"] = request.MinRingingDuration
	}

	if !dara.IsNil(request.MinTalkTime) {
		query["MinTalkTime"] = request.MinTalkTime
	}

	if !dara.IsNil(request.MinTalkTurns) {
		query["MinTalkTurns"] = request.MinTalkTurns
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatesShrink) {
		query["States"] = request.StatesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCases"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of cases.
//
// Description:
//
// ***
//
// @param request - ListCasesRequest
//
// @return ListCasesResponse
func (client *Client) ListCases(request *ListCasesRequest) (_result *ListCasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCasesResponse{}
	_body, _err := client.ListCasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of available models for voice cloning.
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
// Retrieves the list of available models for voice cloning.
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
// Retrieves the list of cloned voices.
//
// @param request - ListCloneVoicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCloneVoicesResponse
func (client *Client) ListCloneVoicesWithOptions(request *ListCloneVoicesRequest, runtime *dara.RuntimeOptions) (_result *ListCloneVoicesResponse, _err error) {
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
		Action:      dara.String("ListCloneVoices"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCloneVoicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of cloned voices.
//
// @param request - ListCloneVoicesRequest
//
// @return ListCloneVoicesResponse
func (client *Client) ListCloneVoices(request *ListCloneVoicesRequest) (_result *ListCloneVoicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCloneVoicesResponse{}
	_body, _err := client.ListCloneVoicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of tenant numbers.
//
// Description:
//
// Queries the list of tenant numbers.
//
// @param request - ListCorpNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCorpNumbersResponse
func (client *Client) ListCorpNumbersWithOptions(request *ListCorpNumbersRequest, runtime *dara.RuntimeOptions) (_result *ListCorpNumbersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Number) {
		body["Number"] = request.Number
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SipTrunkId) {
		body["SipTrunkId"] = request.SipTrunkId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCorpNumbers"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCorpNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of tenant numbers.
//
// Description:
//
// Queries the list of tenant numbers.
//
// @param request - ListCorpNumbersRequest
//
// @return ListCorpNumbersResponse
func (client *Client) ListCorpNumbers(request *ListCorpNumbersRequest) (_result *ListCorpNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCorpNumbersResponse{}
	_body, _err := client.ListCorpNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of flash SMS configurations.
//
// @param request - ListFlashSmsAccessProfilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsAccessProfilesResponse
func (client *Client) ListFlashSmsAccessProfilesWithOptions(request *ListFlashSmsAccessProfilesRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsAccessProfilesResponse, _err error) {
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
		Action:      dara.String("ListFlashSmsAccessProfiles"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsAccessProfilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of flash SMS configurations.
//
// @param request - ListFlashSmsAccessProfilesRequest
//
// @return ListFlashSmsAccessProfilesResponse
func (client *Client) ListFlashSmsAccessProfiles(request *ListFlashSmsAccessProfilesRequest) (_result *ListFlashSmsAccessProfilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlashSmsAccessProfilesResponse{}
	_body, _err := client.ListFlashSmsAccessProfilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of flash message providers.
//
// @param request - ListFlashSmsProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsProvidersResponse
func (client *Client) ListFlashSmsProvidersWithOptions(request *ListFlashSmsProvidersRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsProvidersResponse, _err error) {
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
		Action:      dara.String("ListFlashSmsProviders"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of flash message providers.
//
// @param request - ListFlashSmsProvidersRequest
//
// @return ListFlashSmsProvidersResponse
func (client *Client) ListFlashSmsProviders(request *ListFlashSmsProvidersRequest) (_result *ListFlashSmsProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlashSmsProvidersResponse{}
	_body, _err := client.ListFlashSmsProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves flash SMS templates.
//
// @param request - ListFlashSmsTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsTemplatesResponse
func (client *Client) ListFlashSmsTemplatesWithOptions(request *ListFlashSmsTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.ProviderId) {
		body["ProviderId"] = request.ProviderId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlashSmsTemplates"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves flash SMS templates.
//
// @param request - ListFlashSmsTemplatesRequest
//
// @return ListFlashSmsTemplatesResponse
func (client *Client) ListFlashSmsTemplates(request *ListFlashSmsTemplatesRequest) (_result *ListFlashSmsTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFlashSmsTemplatesResponse{}
	_body, _err := client.ListFlashSmsTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the instances bound to a number.
//
// Description:
//
// Queries the binding information between a number and instances.
//
// @param request - ListInstanceBindingsOfNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceBindingsOfNumberResponse
func (client *Client) ListInstanceBindingsOfNumberWithOptions(request *ListInstanceBindingsOfNumberRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceBindingsOfNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Number) {
		body["Number"] = request.Number
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceBindingsOfNumber"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceBindingsOfNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instances bound to a number.
//
// Description:
//
// Queries the binding information between a number and instances.
//
// @param request - ListInstanceBindingsOfNumberRequest
//
// @return ListInstanceBindingsOfNumberResponse
func (client *Client) ListInstanceBindingsOfNumber(request *ListInstanceBindingsOfNumberRequest) (_result *ListInstanceBindingsOfNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceBindingsOfNumberResponse{}
	_body, _err := client.ListInstanceBindingsOfNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves instance details.
//
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
	body := map[string]interface{}{}
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
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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

// Summary:
//
// Retrieves instance details.
//
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
// Queries the language model access configurations.
//
// Description:
//
// Queries the list of language model access configurations of a specified instance by paging.
//
// @param request - ListLlmAccessProfilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLlmAccessProfilesResponse
func (client *Client) ListLlmAccessProfilesWithOptions(request *ListLlmAccessProfilesRequest, runtime *dara.RuntimeOptions) (_result *ListLlmAccessProfilesResponse, _err error) {
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
		Action:      dara.String("ListLlmAccessProfiles"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLlmAccessProfilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the language model access configurations.
//
// Description:
//
// Queries the list of language model access configurations of a specified instance by paging.
//
// @param request - ListLlmAccessProfilesRequest
//
// @return ListLlmAccessProfilesResponse
func (client *Client) ListLlmAccessProfiles(request *ListLlmAccessProfilesRequest) (_result *ListLlmAccessProfilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLlmAccessProfilesResponse{}
	_body, _err := client.ListLlmAccessProfilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of numbers bound to an instance.
//
// Description:
//
// Queries the binding information between an instance and numbers.
//
// @param request - ListNumberBindingsOfInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNumberBindingsOfInstanceResponse
func (client *Client) ListNumberBindingsOfInstanceWithOptions(request *ListNumberBindingsOfInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListNumberBindingsOfInstanceResponse, _err error) {
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

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNumberBindingsOfInstance"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNumberBindingsOfInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of numbers bound to an instance.
//
// Description:
//
// Queries the binding information between an instance and numbers.
//
// @param request - ListNumberBindingsOfInstanceRequest
//
// @return ListNumberBindingsOfInstanceResponse
func (client *Client) ListNumberBindingsOfInstance(request *ListNumberBindingsOfInstanceRequest) (_result *ListNumberBindingsOfInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNumberBindingsOfInstanceResponse{}
	_body, _err := client.ListNumberBindingsOfInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the phone numbers of an instance.
//
// Description:
//
// Queries the phone numbers of a specified instance.
//
// @param request - ListNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNumbersResponse
func (client *Client) ListNumbersWithOptions(request *ListNumbersRequest, runtime *dara.RuntimeOptions) (_result *ListNumbersResponse, _err error) {
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

	if !dara.IsNil(request.Number) {
		body["Number"] = request.Number
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SipTrunkId) {
		body["SipTrunkId"] = request.SipTrunkId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNumbers"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the phone numbers of an instance.
//
// Description:
//
// Queries the phone numbers of a specified instance.
//
// @param request - ListNumbersRequest
//
// @return ListNumbersResponse
func (client *Client) ListNumbers(request *ListNumbersRequest) (_result *ListNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNumbersResponse{}
	_body, _err := client.ListNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the blacklists and whitelists.
//
// @param request - ListOutboundCallRestrictionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOutboundCallRestrictionsResponse
func (client *Client) ListOutboundCallRestrictionsWithOptions(request *ListOutboundCallRestrictionsRequest, runtime *dara.RuntimeOptions) (_result *ListOutboundCallRestrictionsResponse, _err error) {
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

	if !dara.IsNil(request.Policy) {
		body["Policy"] = request.Policy
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOutboundCallRestrictions"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOutboundCallRestrictionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the blacklists and whitelists.
//
// @param request - ListOutboundCallRestrictionsRequest
//
// @return ListOutboundCallRestrictionsResponse
func (client *Client) ListOutboundCallRestrictions(request *ListOutboundCallRestrictionsRequest) (_result *ListOutboundCallRestrictionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOutboundCallRestrictionsResponse{}
	_body, _err := client.ListOutboundCallRestrictionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of scenario configuration templates.
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

	if !dara.IsNil(request.NluEngine) {
		body["NluEngine"] = request.NluEngine
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
// Retrieves the list of scenario configuration templates.
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
// Retrieves the details of an instance.
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
	if !dara.IsNil(request.BuilderType) {
		body["BuilderType"] = request.BuilderType
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

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PublishOnly) {
		body["PublishOnly"] = request.PublishOnly
	}

	if !dara.IsNil(request.ScriptIdsShrink) {
		body["ScriptIds"] = request.ScriptIdsShrink
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
// Retrieves the details of an instance.
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
// Retrieves a list of scenarios based on the flow ID.
//
// @param request - ListScriptsByFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptsByFlowResponse
func (client *Client) ListScriptsByFlowWithOptions(request *ListScriptsByFlowRequest, runtime *dara.RuntimeOptions) (_result *ListScriptsByFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		body["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScriptsByFlow"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptsByFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of scenarios based on the flow ID.
//
// @param request - ListScriptsByFlowRequest
//
// @return ListScriptsByFlowResponse
func (client *Client) ListScriptsByFlow(request *ListScriptsByFlowRequest) (_result *ListScriptsByFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScriptsByFlowResponse{}
	_body, _err := client.ListScriptsByFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of system configurations.
//
// @param request - ListSystemConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSystemConfigsResponse
func (client *Client) ListSystemConfigsWithOptions(request *ListSystemConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListSystemConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		body["ObjectType"] = request.ObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSystemConfigs"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSystemConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of system configurations.
//
// @param request - ListSystemConfigsRequest
//
// @return ListSystemConfigsResponse
func (client *Client) ListSystemConfigs(request *ListSystemConfigsRequest) (_result *ListSystemConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSystemConfigsResponse{}
	_body, _err := client.ListSystemConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of voice access configurations.
//
// @param request - ListVoiceAccessProfilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVoiceAccessProfilesResponse
func (client *Client) ListVoiceAccessProfilesWithOptions(request *ListVoiceAccessProfilesRequest, runtime *dara.RuntimeOptions) (_result *ListVoiceAccessProfilesResponse, _err error) {
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
		Action:      dara.String("ListVoiceAccessProfiles"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVoiceAccessProfilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of voice access configurations.
//
// @param request - ListVoiceAccessProfilesRequest
//
// @return ListVoiceAccessProfilesResponse
func (client *Client) ListVoiceAccessProfiles(request *ListVoiceAccessProfilesRequest) (_result *ListVoiceAccessProfilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVoiceAccessProfilesResponse{}
	_body, _err := client.ListVoiceAccessProfilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates an outbound call.
//
// Description:
//
// Initiates an outbound call test to the real phone number corresponding to Number by using the specified instance and scenario. Before calling this operation, you must obtain an SMS verification code for the OutboundCallTest scenario for the called number and pass it in VerificationCode. The service verifies the code and initiates the call by using the outbound caller ID configured on the platform. DeviceId is used to bind the client data stream to the current call. CallVariables passes call variables in JSON object format, and Tags passes attached data. Upon success, a SessionId with the test- prefix is returned to associate with the current call. A successful API response does not indicate that the called party has answered or that the conversation has ended. Test only authorized numbers. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use real resources and service return values for actual calls.
//
// @param request - MakePilotCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MakePilotCallResponse
func (client *Client) MakePilotCallWithOptions(request *MakePilotCallRequest, runtime *dara.RuntimeOptions) (_result *MakePilotCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallVariables) {
		query["CallVariables"] = request.CallVariables
	}

	if !dara.IsNil(request.DeviceId) {
		query["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.DraftVersion) {
		query["DraftVersion"] = request.DraftVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.VerificationCode) {
		query["VerificationCode"] = request.VerificationCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MakePilotCall"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MakePilotCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an outbound call.
//
// Description:
//
// Initiates an outbound call test to the real phone number corresponding to Number by using the specified instance and scenario. Before calling this operation, you must obtain an SMS verification code for the OutboundCallTest scenario for the called number and pass it in VerificationCode. The service verifies the code and initiates the call by using the outbound caller ID configured on the platform. DeviceId is used to bind the client data stream to the current call. CallVariables passes call variables in JSON object format, and Tags passes attached data. Upon success, a SessionId with the test- prefix is returned to associate with the current call. A successful API response does not indicate that the called party has answered or that the conversation has ended. Test only authorized numbers. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use real resources and service return values for actual calls.
//
// @param request - MakePilotCallRequest
//
// @return MakePilotCallResponse
func (client *Client) MakePilotCall(request *MakePilotCallRequest) (_result *MakePilotCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MakePilotCallResponse{}
	_body, _err := client.MakePilotCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses an outbound campaign.
//
// @param request - PauseCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseCampaignResponse
func (client *Client) PauseCampaignWithOptions(request *PauseCampaignRequest, runtime *dara.RuntimeOptions) (_result *PauseCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseCampaign"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses an outbound campaign.
//
// @param request - PauseCampaignRequest
//
// @return PauseCampaignResponse
func (client *Client) PauseCampaign(request *PauseCampaignRequest) (_result *PauseCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PauseCampaignResponse{}
	_body, _err := client.PauseCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an instance.
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
// Updates an instance.
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
// Resumes an outbound campaign.
//
// Description:
//
// ***
//
// @param request - ResumeCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeCampaignResponse
func (client *Client) ResumeCampaignWithOptions(request *ResumeCampaignRequest, runtime *dara.RuntimeOptions) (_result *ResumeCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeCampaign"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes an outbound campaign.
//
// Description:
//
// ***
//
// @param request - ResumeCampaignRequest
//
// @return ResumeCampaignResponse
func (client *Client) ResumeCampaign(request *ResumeCampaignRequest) (_result *ResumeCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeCampaignResponse{}
	_body, _err := client.ResumeCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an outbound campaign.
//
// Description:
//
// ***
//
// @param request - SubmitCampaignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCampaignResponse
func (client *Client) SubmitCampaignWithOptions(request *SubmitCampaignRequest, runtime *dara.RuntimeOptions) (_result *SubmitCampaignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CampaignId) {
		query["CampaignId"] = request.CampaignId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCampaign"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an outbound campaign.
//
// Description:
//
// ***
//
// @param request - SubmitCampaignRequest
//
// @return SubmitCampaignResponse
func (client *Client) SubmitCampaign(request *SubmitCampaignRequest) (_result *SubmitCampaignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitCampaignResponse{}
	_body, _err := client.SubmitCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a user input to a trial text conversation.
//
// Description:
//
// Sends a round of user input to a trial text conversation created by BeginTrialSession and returns the robot reply in Server-Sent Events (SSE) format. Use the same InstanceId, ScriptId, SessionId, and experience link VisitorToken as those used to start the trial conversation. The service validates the token and its expiration time on each call. Utterance is plain text input, and Extras is a JSON array string of structured events such as key presses and silence. StreamEnd=true only ends the response stream for the current turn. Call EndTrialSession when the entire experience ends. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - TrialDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrialDialogueResponse
func (client *Client) TrialDialogueWithSSE(request *TrialDialogueRequest, runtime *dara.RuntimeOptions, _yield chan *TrialDialogueResponse, _yieldErr chan error) {
	defer close(_yield)
	client.trialDialogueWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// Sends a user input to a trial text conversation.
//
// Description:
//
// Sends a round of user input to a trial text conversation created by BeginTrialSession and returns the robot reply in Server-Sent Events (SSE) format. Use the same InstanceId, ScriptId, SessionId, and experience link VisitorToken as those used to start the trial conversation. The service validates the token and its expiration time on each call. Utterance is plain text input, and Extras is a JSON array string of structured events such as key presses and silence. StreamEnd=true only ends the response stream for the current turn. Call EndTrialSession when the entire experience ends. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - TrialDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrialDialogueResponse
func (client *Client) TrialDialogueWithOptions(request *TrialDialogueRequest, runtime *dara.RuntimeOptions) (_result *TrialDialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Extras) {
		query["Extras"] = request.Extras
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	if !dara.IsNil(request.VisitorToken) {
		query["VisitorToken"] = request.VisitorToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrialDialogue"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TrialDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a user input to a trial text conversation.
//
// Description:
//
// Sends a round of user input to a trial text conversation created by BeginTrialSession and returns the robot reply in Server-Sent Events (SSE) format. Use the same InstanceId, ScriptId, SessionId, and experience link VisitorToken as those used to start the trial conversation. The service validates the token and its expiration time on each call. Utterance is plain text input, and Extras is a JSON array string of structured events such as key presses and silence. StreamEnd=true only ends the response stream for the current turn. Call EndTrialSession when the entire experience ends. The resource identifiers, numbers, and authentication credentials in the examples are desensitized or placeholder values. Use actual resources and service return values for actual calls.
//
// @param request - TrialDialogueRequest
//
// @return TrialDialogueResponse
func (client *Client) TrialDialogue(request *TrialDialogueRequest) (_result *TrialDialogueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TrialDialogueResponse{}
	_body, _err := client.TrialDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an instance.
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
// Updates an instance.
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
// Updates the flash message configuration.
//
// @param tmpReq - UpdateFlashSmsAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlashSmsAccessProfileResponse
func (client *Client) UpdateFlashSmsAccessProfileWithOptions(tmpReq *UpdateFlashSmsAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateFlashSmsAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateFlashSmsAccessProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccessProfile) {
		request.AccessProfileShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccessProfile, dara.String("AccessProfile"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessProfileShrink) {
		body["AccessProfile"] = request.AccessProfileShrink
	}

	if !dara.IsNil(request.AccessProfileId) {
		body["AccessProfileId"] = request.AccessProfileId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProviderId) {
		body["ProviderId"] = request.ProviderId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFlashSmsAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlashSmsAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the flash message configuration.
//
// @param request - UpdateFlashSmsAccessProfileRequest
//
// @return UpdateFlashSmsAccessProfileResponse
func (client *Client) UpdateFlashSmsAccessProfile(request *UpdateFlashSmsAccessProfileRequest) (_result *UpdateFlashSmsAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFlashSmsAccessProfileResponse{}
	_body, _err := client.UpdateFlashSmsAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an instance.
//
// @param request - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an instance.
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the instances bound to a number.
//
// Description:
//
// Updates the binding between a number and instances.
//
// @param tmpReq - UpdateInstanceBindingsOfNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceBindingsOfNumberResponse
func (client *Client) UpdateInstanceBindingsOfNumberWithOptions(tmpReq *UpdateInstanceBindingsOfNumberRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceBindingsOfNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateInstanceBindingsOfNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		body["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.Number) {
		body["Number"] = request.Number
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceBindingsOfNumber"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceBindingsOfNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the instances bound to a number.
//
// Description:
//
// Updates the binding between a number and instances.
//
// @param request - UpdateInstanceBindingsOfNumberRequest
//
// @return UpdateInstanceBindingsOfNumberResponse
func (client *Client) UpdateInstanceBindingsOfNumber(request *UpdateInstanceBindingsOfNumberRequest) (_result *UpdateInstanceBindingsOfNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceBindingsOfNumberResponse{}
	_body, _err := client.UpdateInstanceBindingsOfNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the language model access configuration.
//
// Description:
//
// Updates the language model access configuration for a specified instance.
//
// @param tmpReq - UpdateLlmAccessProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLlmAccessProfileResponse
func (client *Client) UpdateLlmAccessProfileWithOptions(tmpReq *UpdateLlmAccessProfileRequest, runtime *dara.RuntimeOptions) (_result *UpdateLlmAccessProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateLlmAccessProfileShrinkRequest{}
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

	if !dara.IsNil(request.ProfileShrink) {
		body["Profile"] = request.ProfileShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLlmAccessProfile"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLlmAccessProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the language model access configuration.
//
// Description:
//
// Updates the language model access configuration for a specified instance.
//
// @param request - UpdateLlmAccessProfileRequest
//
// @return UpdateLlmAccessProfileResponse
func (client *Client) UpdateLlmAccessProfile(request *UpdateLlmAccessProfileRequest) (_result *UpdateLlmAccessProfileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLlmAccessProfileResponse{}
	_body, _err := client.UpdateLlmAccessProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the phone numbers bound to an instance.
//
// Description:
//
// Updates the binding relationship between an instance and phone numbers.
//
// @param tmpReq - UpdateNumberBindingsOfInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNumberBindingsOfInstanceResponse
func (client *Client) UpdateNumberBindingsOfInstanceWithOptions(tmpReq *UpdateNumberBindingsOfInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateNumberBindingsOfInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateNumberBindingsOfInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NumbersShrink) {
		body["Numbers"] = request.NumbersShrink
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNumberBindingsOfInstance"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNumberBindingsOfInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the phone numbers bound to an instance.
//
// Description:
//
// Updates the binding relationship between an instance and phone numbers.
//
// @param request - UpdateNumberBindingsOfInstanceRequest
//
// @return UpdateNumberBindingsOfInstanceResponse
func (client *Client) UpdateNumberBindingsOfInstance(request *UpdateNumberBindingsOfInstanceRequest) (_result *UpdateNumberBindingsOfInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNumberBindingsOfInstanceResponse{}
	_body, _err := client.UpdateNumberBindingsOfInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an instance.
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
// Updates an instance.
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
// Updates the scenario configuration.
//
// Description:
//
// ***
//
// @param tmpReq - UpdateScriptVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScriptVersionResponse
func (client *Client) UpdateScriptVersionWithOptions(tmpReq *UpdateScriptVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdateScriptVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateScriptVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InteractionConfig) {
		request.InteractionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InteractionConfig, dara.String("InteractionConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.LabelConfigs) {
		request.LabelConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelConfigs, dara.String("LabelConfigs"), dara.String("json"))
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

	if !dara.IsNil(request.LabelConfigsShrink) {
		body["LabelConfigs"] = request.LabelConfigsShrink
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.ScriptProfileShrink) {
		body["ScriptProfile"] = request.ScriptProfileShrink
	}

	if !dara.IsNil(request.SynthesizerConfigShrink) {
		body["SynthesizerConfig"] = request.SynthesizerConfigShrink
	}

	if !dara.IsNil(request.TranscriberConfigShrink) {
		body["TranscriberConfig"] = request.TranscriberConfigShrink
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScriptVersion"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScriptVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the scenario configuration.
//
// Description:
//
// ***
//
// @param request - UpdateScriptVersionRequest
//
// @return UpdateScriptVersionResponse
func (client *Client) UpdateScriptVersion(request *UpdateScriptVersionRequest) (_result *UpdateScriptVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateScriptVersionResponse{}
	_body, _err := client.UpdateScriptVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or updates a message queue (MQ) configuration.
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
// Creates or updates a message queue (MQ) configuration.
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
// Updates system configurations.
//
// @param tmpReq - UpdateSystemConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSystemConfigsResponse
func (client *Client) UpdateSystemConfigsWithOptions(tmpReq *UpdateSystemConfigsRequest, runtime *dara.RuntimeOptions) (_result *UpdateSystemConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSystemConfigsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Configs) {
		request.ConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configs, dara.String("Configs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigsShrink) {
		body["Configs"] = request.ConfigsShrink
	}

	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		body["ObjectType"] = request.ObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSystemConfigs"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSystemConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates system configurations.
//
// @param request - UpdateSystemConfigsRequest
//
// @return UpdateSystemConfigsResponse
func (client *Client) UpdateSystemConfigs(request *UpdateSystemConfigsRequest) (_result *UpdateSystemConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSystemConfigsResponse{}
	_body, _err := client.UpdateSystemConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an instance.
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
// Updates an instance.
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

func (client *Client) beginSessionWithSSE_opYieldFunc(_yield chan *BeginSessionResponse, _yieldErr chan error, request *BeginSessionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DraftVersion) {
		query["DraftVersion"] = request.DraftVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.VendorParams) {
		query["VendorParams"] = request.VendorParams
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BeginSession"),
		Version:     dara.String("2025-11-11"),
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

func (client *Client) beginTrialSessionWithSSE_opYieldFunc(_yield chan *BeginTrialSessionResponse, _yieldErr chan error, request *BeginTrialSessionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DraftVersion) {
		query["DraftVersion"] = request.DraftVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.VendorParams) {
		query["VendorParams"] = request.VendorParams
	}

	if !dara.IsNil(request.VisitorToken) {
		query["VisitorToken"] = request.VisitorToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BeginTrialSession"),
		Version:     dara.String("2025-11-11"),
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

func (client *Client) dialogueWithSSE_opYieldFunc(_yield chan *DialogueResponse, _yieldErr chan error, request *DialogueRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Extras) {
		query["Extras"] = request.Extras
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Dialogue"),
		Version:     dara.String("2025-11-11"),
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

func (client *Client) trialDialogueWithSSE_opYieldFunc(_yield chan *TrialDialogueResponse, _yieldErr chan error, request *TrialDialogueRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Extras) {
		query["Extras"] = request.Extras
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	if !dara.IsNil(request.VisitorToken) {
		query["VisitorToken"] = request.VisitorToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TrialDialogue"),
		Version:     dara.String("2025-11-11"),
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
