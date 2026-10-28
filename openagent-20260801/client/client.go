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
	client.Endpoint, _err = client.GetEndpoint(dara.String("openagent"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Cancels a run.
//
// Description:
//
// Cancels a specified run and returns the run status after the cancellation.
//
// @param request - CancelRunRequest
//
// @param headers - CancelRunHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRunResponse
func (client *Client) CancelRunWithOptions(sessionId *string, runId *string, request *CancelRunRequest, headers *CancelRunHeaders, runtime *dara.RuntimeOptions) (_result *CancelRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Reason) {
		body["reason"] = request.Reason
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XOpenAgentIdempotencyKey) {
		realHeaders["X-OpenAgent-Idempotency-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XOpenAgentIdempotencyKey)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelRun"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions/" + dara.PercentEncode(dara.StringValue(sessionId)) + "/runs/" + dara.PercentEncode(dara.StringValue(runId)) + "/cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a run.
//
// Description:
//
// Cancels a specified run and returns the run status after the cancellation.
//
// @param request - CancelRunRequest
//
// @return CancelRunResponse
func (client *Client) CancelRun(sessionId *string, runId *string, request *CancelRunRequest) (_result *CancelRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CancelRunHeaders{}
	_result = &CancelRunResponse{}
	_body, _err := client.CancelRunWithOptions(sessionId, runId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a session.
//
// Description:
//
// Creates a session and returns the session information. Use the returned sessionId for subsequent operations such as sending messages, querying records, and updating the session.
//
// @param request - CreateSessionRequest
//
// @param headers - CreateSessionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionResponse
func (client *Client) CreateSessionWithOptions(request *CreateSessionRequest, headers *CreateSessionHeaders, runtime *dara.RuntimeOptions) (_result *CreateSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XOpenAgentIdempotencyKey) {
		realHeaders["X-OpenAgent-Idempotency-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XOpenAgentIdempotencyKey)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSession"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a session.
//
// Description:
//
// Creates a session and returns the session information. Use the returned sessionId for subsequent operations such as sending messages, querying records, and updating the session.
//
// @param request - CreateSessionRequest
//
// @return CreateSessionResponse
func (client *Client) CreateSession(request *CreateSessionRequest) (_result *CreateSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &CreateSessionHeaders{}
	_result = &CreateSessionResponse{}
	_body, _err := client.CreateSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a session.
//
// Description:
//
// Deletes a specified session. Returns 404 Session.NotFound if the session does not exist or has been deleted.
//
// @param request - DeleteSessionRequest
//
// @param headers - DeleteSessionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSessionResponse
func (client *Client) DeleteSessionWithOptions(sessionId *string, request *DeleteSessionRequest, headers *DeleteSessionHeaders, runtime *dara.RuntimeOptions) (_result *DeleteSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XOpenAgentIdempotencyKey) {
		realHeaders["X-OpenAgent-Idempotency-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XOpenAgentIdempotencyKey)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSession"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions/" + dara.PercentEncode(dara.StringValue(sessionId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a session.
//
// Description:
//
// Deletes a specified session. Returns 404 Session.NotFound if the session does not exist or has been deleted.
//
// @param request - DeleteSessionRequest
//
// @return DeleteSessionResponse
func (client *Client) DeleteSession(sessionId *string, request *DeleteSessionRequest) (_result *DeleteSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &DeleteSessionHeaders{}
	_result = &DeleteSessionResponse{}
	_body, _err := client.DeleteSessionWithOptions(sessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the read-only mode.
//
// Description:
//
// Queries the read-only mode switch for the current account. The read-only mode is enabled by default if it is not configured.
//
// @param request - GetReadonlyModeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReadonlyModeResponse
func (client *Client) GetReadonlyModeWithOptions(request *GetReadonlyModeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReadonlyModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReadonlyMode"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/user/configs/readonly"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReadonlyModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the read-only mode.
//
// Description:
//
// Queries the read-only mode switch for the current account. The read-only mode is enabled by default if it is not configured.
//
// @param request - GetReadonlyModeRequest
//
// @return GetReadonlyModeResponse
func (client *Client) GetReadonlyMode(request *GetReadonlyModeRequest) (_result *GetReadonlyModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetReadonlyModeResponse{}
	_body, _err := client.GetReadonlyModeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a run.
//
// Description:
//
// Queries the status and interrupt information of a specified run.
//
// @param request - GetRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRunResponse
func (client *Client) GetRunWithOptions(sessionId *string, runId *string, request *GetRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRun"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions/" + dara.PercentEncode(dara.StringValue(sessionId)) + "/runs/" + dara.PercentEncode(dara.StringValue(runId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a run.
//
// Description:
//
// Queries the status and interrupt information of a specified run.
//
// @param request - GetRunRequest
//
// @return GetRunResponse
func (client *Client) GetRun(sessionId *string, runId *string, request *GetRunRequest) (_result *GetRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRunResponse{}
	_body, _err := client.GetRunWithOptions(sessionId, runId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a session.
//
// Description:
//
// Queries the title, status, pinning information, and the ID and status of the most recent run of a specified session.
//
// @param request - GetSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSessionResponse
func (client *Client) GetSessionWithOptions(sessionId *string, request *GetSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSession"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions/" + dara.PercentEncode(dara.StringValue(sessionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a session.
//
// Description:
//
// Queries the title, status, pinning information, and the ID and status of the most recent run of a specified session.
//
// @param request - GetSessionRequest
//
// @return GetSessionResponse
func (client *Client) GetSession(sessionId *string, request *GetSessionRequest) (_result *GetSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSessionResponse{}
	_body, _err := client.GetSessionWithOptions(sessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists messages.
//
// Description:
//
// Performs a paging query on the message timeline of a specified session. Use the returned nextToken to retrieve the next page.
//
// @param request - ListMessagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessagesResponse
func (client *Client) ListMessagesWithOptions(sessionId *string, request *ListMessagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RunId) {
		query["runId"] = request.RunId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessages"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions/" + dara.PercentEncode(dara.StringValue(sessionId)) + "/messages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists messages.
//
// Description:
//
// Performs a paging query on the message timeline of a specified session. Use the returned nextToken to retrieve the next page.
//
// @param request - ListMessagesRequest
//
// @return ListMessagesResponse
func (client *Client) ListMessages(sessionId *string, request *ListMessagesRequest) (_result *ListMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMessagesResponse{}
	_body, _err := client.ListMessagesWithOptions(sessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists sessions.
//
// Description:
//
// Uses paging to query the list of sessions accessible to the current identity used to invoke. Use the returned nextToken to retrieve the next page.
//
// @param request - ListSessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSessionsResponse
func (client *Client) ListSessionsWithOptions(request *ListSessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSessions"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists sessions.
//
// Description:
//
// Uses paging to query the list of sessions accessible to the current identity used to invoke. Use the returned nextToken to retrieve the next page.
//
// @param request - ListSessionsRequest
//
// @return ListSessionsResponse
func (client *Client) ListSessions(request *ListSessionsRequest) (_result *ListSessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSessionsResponse{}
	_body, _err := client.ListSessionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a message.
//
// Description:
//
// Sends a message to a specified session, submits interrupt handling results, or reconnects to a specified run. Returns the run progress and results through a Server-Sent Events (SSE) event stream.
//
// @param request - SendMessageRequest
//
// @param headers - SendMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithSSE(sessionId *string, request *SendMessageRequest, headers *SendMessageHeaders, runtime *dara.RuntimeOptions, _yield chan *SendMessageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.sendMessageWithSSE_opYieldFunc(_yield, _yieldErr, sessionId, request, headers, runtime)
	return
}

// Summary:
//
// Sends a message.
//
// Description:
//
// Sends a message to a specified session, submits interrupt handling results, or reconnects to a specified run. Returns the run progress and results through a Server-Sent Events (SSE) event stream.
//
// @param request - SendMessageRequest
//
// @param headers - SendMessageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageResponse
func (client *Client) SendMessageWithOptions(sessionId *string, request *SendMessageRequest, headers *SendMessageHeaders, runtime *dara.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Resume) {
		body["resume"] = request.Resume
	}

	if !dara.IsNil(request.RunId) {
		body["runId"] = request.RunId
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XOpenAgentIdempotencyKey) {
		realHeaders["X-OpenAgent-Idempotency-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XOpenAgentIdempotencyKey)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessage"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions/" + dara.PercentEncode(dara.StringValue(sessionId)) + "/messages"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a message.
//
// Description:
//
// Sends a message to a specified session, submits interrupt handling results, or reconnects to a specified run. Returns the run progress and results through a Server-Sent Events (SSE) event stream.
//
// @param request - SendMessageRequest
//
// @return SendMessageResponse
func (client *Client) SendMessage(sessionId *string, request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &SendMessageHeaders{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(sessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the read-only mode.
//
// Description:
//
// Enables or disables the read-only mode for the current account.
//
// @param request - UpdateReadonlyModeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateReadonlyModeResponse
func (client *Client) UpdateReadonlyModeWithOptions(request *UpdateReadonlyModeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateReadonlyModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Value) {
		body["value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateReadonlyMode"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/user/configs/readonly"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateReadonlyModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the read-only mode.
//
// Description:
//
// Enables or disables the read-only mode for the current account.
//
// @param request - UpdateReadonlyModeRequest
//
// @return UpdateReadonlyModeResponse
func (client *Client) UpdateReadonlyMode(request *UpdateReadonlyModeRequest) (_result *UpdateReadonlyModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateReadonlyModeResponse{}
	_body, _err := client.UpdateReadonlyModeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the title or pinned status of a session.
//
// Description:
//
// Updates the title or pinned status of a specified session. Specify at least one of the title and pinned fields. Unspecified fields remain unchanged.
//
// @param request - UpdateSessionRequest
//
// @param headers - UpdateSessionHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSessionResponse
func (client *Client) UpdateSessionWithOptions(sessionId *string, request *UpdateSessionRequest, headers *UpdateSessionHeaders, runtime *dara.RuntimeOptions) (_result *UpdateSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Pinned) {
		body["pinned"] = request.Pinned
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XOpenAgentIdempotencyKey) {
		realHeaders["X-OpenAgent-Idempotency-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XOpenAgentIdempotencyKey)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSession"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions/" + dara.PercentEncode(dara.StringValue(sessionId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the title or pinned status of a session.
//
// Description:
//
// Updates the title or pinned status of a specified session. Specify at least one of the title and pinned fields. Unspecified fields remain unchanged.
//
// @param request - UpdateSessionRequest
//
// @return UpdateSessionResponse
func (client *Client) UpdateSession(sessionId *string, request *UpdateSessionRequest) (_result *UpdateSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &UpdateSessionHeaders{}
	_result = &UpdateSessionResponse{}
	_body, _err := client.UpdateSessionWithOptions(sessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) sendMessageWithSSE_opYieldFunc(_yield chan *SendMessageResponse, _yieldErr chan error, sessionId *string, request *SendMessageRequest, headers *SendMessageHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Resume) {
		body["resume"] = request.Resume
	}

	if !dara.IsNil(request.RunId) {
		body["runId"] = request.RunId
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XOpenAgentIdempotencyKey) {
		realHeaders["X-OpenAgent-Idempotency-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XOpenAgentIdempotencyKey)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessage"),
		Version:     dara.String("2026-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/v1/agent/sessions/" + dara.PercentEncode(dara.StringValue(sessionId)) + "/messages"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.StringValue(resp.Event.Data)
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
