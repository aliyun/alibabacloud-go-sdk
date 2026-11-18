// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Stops a session execution or cancels a specified turn.
//
// Description:
//
// If `Params._meta.promptId` is provided, only that turn is canceled. If omitted, the current execution of the session is canceled. To avoid affecting other turns, use the actual `promptId` obtained from the response stream instead of the request `Id`. A cancellation receipt does not mean the turn has ended, nor does it revoke external operations that have already occurred. After calling this operation, continue reading the original response stream or resume reading through `LoadAgentSession` to confirm the desired state of the target turn. Disconnecting the stream connection cannot replace a cancellation request.
//
// @param tmpReq - CancelAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAgentSessionResponse
func (client *Client) CancelAgentSessionWithContext(ctx context.Context, tmpReq *CancelAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *CancelAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CancelAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAgentSession"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an agent.
//
// Description:
//
// - Name is the unique identifier of the agent within the current tenant. It can contain only letters, digits, underscores, and hyphens, and must be 1 to 128 characters in length. It cannot be modified after creation.
//
// - Each agent can be bound to only one knowledge base.
//
// @param request - CreateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentResponse
func (client *Client) CreateAgentWithContext(ctx context.Context, request *CreateAgentRequest, runtime *dara.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
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

	if !dara.IsNil(request.KnowledgeBases) {
		body["KnowledgeBases"] = request.KnowledgeBases
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Skills) {
		body["Skills"] = request.Skills
	}

	if !dara.IsNil(request.SystemPrompt) {
		body["SystemPrompt"] = request.SystemPrompt
	}

	if !dara.IsNil(request.Tools) {
		body["Tools"] = request.Tools
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgent"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a text session for a specified agent and obtains the session ID required for subsequent calls.
//
// Description:
//
// Before calling this operation, obtain the name of an agent accessible to the current identity. After the session is created, save `JsonRpcResponse.result.sessionId` and use it as `Params.sessionId` for subsequent operations such as `PromptAgentSession` and `LoadAgentSession`.
//
// You can set the approval mode during creation. If omitted, the mode defaults to `default`. The availability of `yolo` depends on the target agent. If it is not supported, an error is returned. This mode does not change the caller\\"s permissions or replace business Q&A. To retry the same creation request, reuse the same `clientToken` and the same business parameters.
//
// @param tmpReq - CreateAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentSessionResponse
func (client *Client) CreateAgentSessionWithContext(ctx context.Context, tmpReq *CreateAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *CreateAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentSession"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a skill.
//
// Description:
//
// - Creation rules:
//
//   - `Name` must be unique within the current tenant and cannot be the same as an official skill.
//
//   - `Description` and `Metadata` are required. `Metadata` must contain one and only one valid content source. Different sources cannot be mixed.
//
//   - After the skill is created, you can modify it by calling `UpdateSkill`.
//
// - Content sources:
//
//   - `skillMd`: Directly enter the Markdown body. No YAML front matter is required. Use `Name` and `Description` for the name and description.
//
//   - `transitId`: Upload and confirm a ZIP file through Transit. The call sequence is as follows:
//
//     1. Call `CreateTransitUploadPolicy` and pass in `FileShowName` to obtain `TransitId`, `FilePath`, and `PolicyInfo`.
//
//     2. Use `PolicyInfo` and `FilePath` to upload the ZIP file to Object Storage Service (OSS).
//
//     3. Call `ConfirmTransitUpload` and pass in `TransitId`. Proceed with creation only if `Confirmed=true` is returned.
//
//     4. Call `CreateSkill` and pass the confirmed `TransitId` in `Metadata.transitId`.
//
//   - `bundleUrl`: Pass in a public HTTPS ZIP direct link. The platform downloads the file and saves it as an artifact. The original URL is not saved or echoed.
//
// @param request - CreateSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillResponse
func (client *Client) CreateSkillWithContext(ctx context.Context, request *CreateSkillRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillResponse, _err error) {
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

	if !dara.IsNil(request.Metadata) {
		body["Metadata"] = request.Metadata
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkill"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates credentials for direct file upload. Callers can use the returned upload policy to upload files directly to object storage.
//
// Description:
//
// - The Network parameter can be set to public or internal to generate a public network or same-region internal network upload address, respectively. The default value is public.
//
// - The maximum size of a single file is 50 MiB.
//
// - PolicyInfo contains short-term upload authorization information and is only used for the current file upload. Do not log, permanently store, or forward it to other users.
//
// - ExpireMs specifies the validity period of the upload policy and the transit record in milliseconds, not an absolute timestamp. The default and maximum values are both 604800000 (7 days), and the minimum value is 1000 (1 second). The validity period is rounded down to the nearest whole second. For example, 1500 milliseconds takes effect as 1 second.
//
// - You can call GetTransitMeta to read ExpireAt and obtain the expiration time of the transit record. Confirming the upload and querying the record do not extend its validity period.
//
// @param request - CreateTransitUploadPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitUploadPolicyResponse
func (client *Client) CreateTransitUploadPolicyWithContext(ctx context.Context, request *CreateTransitUploadPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitUploadPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireMs) {
		body["ExpireMs"] = request.ExpireMs
	}

	if !dara.IsNil(request.FileShowName) {
		body["FileShowName"] = request.FileShowName
	}

	if !dara.IsNil(request.Network) {
		body["Network"] = request.Network
	}

	if !dara.IsNil(request.PathPrefix) {
		body["PathPrefix"] = request.PathPrefix
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransitUploadPolicy"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransitUploadPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified agent.
//
// Description:
//
// - Creators can delete the agents they created.
//
// - After deletion, the agent can no longer be queried, updated, or run.
//
// @param request - DeleteAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgentWithContext(ctx context.Context, request *DeleteAgentRequest, runtime *dara.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgent"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an idle session that is no longer needed.
//
// Description:
//
// A session cannot be directly deleted while it is being executed or waiting for interaction. You must first call `CancelAgentSession` and continue reading the original response stream, or resume reading through `LoadAgentSession`. Delete the session only after confirming that the target turn has ended.
//
// A successful deletion indicates that the session deletion operation is complete, but it does not mean that external operations that have already occurred are rolled back. Deleting a session does not delete the associated agent or assistant.
//
// @param tmpReq - DeleteAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentSessionResponse
func (client *Client) DeleteAgentSessionWithContext(ctx context.Context, tmpReq *DeleteAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *DeleteAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgentSession"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Soft-deletes a skill that the current caller has permission to modify.
//
// Description:
//
// - Deletion permissions:
//
//   - You can delete only custom skills that you have permission to modify.
//
//   - You must be the creator of the skill or the root account of the tenant that has permission to manage the tenant-level skill. Official skills cannot be deleted.
//
// - Deletion results:
//
//   - This operation performs a soft delete. After the skill is deleted, the `GetSkill` and `ListSkills` operations no longer return the skill, and you can create a new skill with the same name.
//
//   - Recovery is not supported. You cannot delete a skill multiple times or delete a non-existent skill.
//
// @param request - DeleteSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSkillResponse
func (client *Client) DeleteSkillWithContext(ctx context.Context, request *DeleteSkillRequest, runtime *dara.RuntimeOptions) (_result *DeleteSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSkill"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries agents visible to the current identity by name.
//
// Description:
//
// - You can retrieve agents created by yourself, agents visible within the current tenant, and official agents provided by the platform.
//
// @param request - GetAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithContext(ctx context.Context, request *GetAgentRequest, runtime *dara.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgent"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries session details such as the title, associated agent, and core timestamps based on the session ID.
//
// Description:
//
// For `Params.sessionId`, use the session ID returned by `CreateAgentSession` or `ListAgentSessions`. The business result is located in `JsonRpcResponse.result`, which is of type `any`. The fields within the object retain their actual casing.
//
// This operation provides a session summary and does not return the complete message history, pending interactions, or the final state of real-time turns. For history and disconnection recovery, use `LoadAgentSession`. For approvals and Q&A, read from the streaming notifications of `PromptAgentSession` or `LoadAgentSession`.
//
// @param tmpReq - GetAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentSessionResponse
func (client *Client) GetAgentSessionWithContext(ctx context.Context, tmpReq *GetAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *GetAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		query["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		query["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentSession"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the model and thinking level configurations currently used by the session.
//
// Description:
//
// - Use the `sessionId` returned by CreateAgentSession to query the configurations. This operation does not send messages or replay session history.
//
// - Read the current model from `JsonRpcResponse.result.modelId`. To modify it, call SetAgentSessionModel.
//
// - The thinking level is an optional advanced configuration. The supported levels and current values are subject to the actual response of this operation, which may vary across different agents, models, and sessions.
//
// @param tmpReq - GetAgentSessionModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentSessionModelResponse
func (client *Client) GetAgentSessionModelWithContext(ctx context.Context, tmpReq *GetAgentSessionModelRequest, runtime *dara.RuntimeOptions) (_result *GetAgentSessionModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAgentSessionModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		query["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		query["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentSessionModel"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentSessionModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the skills visible to the current caller by name.
//
// Description:
//
// - Query target:
//
//   - Queries by `Name`. Querying by `SkillId` is not supported.
//
//   - Prioritizes returning skills with the same name that are visible to the caller under the current tenant. If no visible records are found, it queries official skills with the same name.
//
//   - If `SkillVersion` is omitted, the current skill is returned. This parameter is omitted by default.
//
// - Download URL:
//
//   - `Network` supports `public` and `internal`. If omitted, no download URL is generated.
//
//   - If the skill has an accessible artifact, a temporary `DownloadUrl` and the corresponding `DownloadUrlNetwork` are returned.
//
//   - If the artifact does not exist, access is denied, or URL generation fails, the skill query still succeeds, but the download URL-related fields may not be returned.
//
//   - The original `bundleUrl` used during creation is not saved and will not be echoed back through this API.
//
// @param request - GetSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSkillResponse
func (client *Client) GetSkillWithContext(ctx context.Context, request *GetSkillRequest, runtime *dara.RuntimeOptions) (_result *GetSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Network) {
		query["Network"] = request.Network
	}

	if !dara.IsNil(request.SkillVersion) {
		query["SkillVersion"] = request.SkillVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkill"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata of a transit file, such as the upload status, file size, and expiration time, and retrieves a temporary download URL on demand.
//
// Description:
//
// - Specify at least one of TransitId and FilePath. If both are provided, TransitId takes precedence.
//
// - Use the TransitId returned by CreateTransitUploadPolicy. TransitId is a temporary capability identifier in the file upload process. Do not forward it to unrelated users.
//
// - FilePath is the opaque object path returned by CreateTransitUploadPolicy. Callers must use it as is. Do not parse, modify, or concatenate it.
//
// - When you query by FilePath, an error is returned if the record does not exist or is inaccessible to the current caller.
//
// - ExpireMs specifies the validity period of the download URL in milliseconds. The default value is 900000, which indicates 15 minutes. The validity period is rounded down to the nearest second. For example, 1500 milliseconds takes effect as 1 second.
//
// - ExpireAt is the expiration time of the transit record, not the expiration time of the download URL. Querying, generating a download URL, and confirming the upload do not extend the validity period of the record.
//
// @param request - GetTransitMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransitMetaResponse
func (client *Client) GetTransitMetaWithContext(ctx context.Context, request *GetTransitMetaRequest, runtime *dara.RuntimeOptions) (_result *GetTransitMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpireMs) {
		query["ExpireMs"] = request.ExpireMs
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.Network) {
		query["Network"] = request.Network
	}

	if !dara.IsNil(request.TransitId) {
		query["TransitId"] = request.TransitId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTransitMeta"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTransitMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global workflow configurations.
//
// Description:
//
// Queries the global workflow configurations for a tenant and a region. The content is sensitive information. In addition to the permissions required for this operation, the ada:DownloadFile permission is also required.
//
// @param request - GetWorkflowGlobalEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowGlobalEnvResponse
func (client *Client) GetWorkflowGlobalEnvWithContext(ctx context.Context, request *GetWorkflowGlobalEnvRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowGlobalEnvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkflowGlobalEnv"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowGlobalEnvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the session summary by paging for the caller to select or resume existing sessions.
//
// Description:
//
// You can filter sessions by agent name. Omit Params.NextToken on the first page. When querying subsequent pages, keep the filter conditions and page size unchanged, and pass the JsonRpcResponse.Result.NextToken returned from the previous page as is to Params.NextToken. Pagination ends when the token is not returned or is empty.
//
// The list returns session summaries such as session ID, title, and time, but does not provide the final state of turn execution. After selecting a session, you can use GetAgentSession to query details or LoadAgentSession to load history.
//
// Params in the request, and JsonRpcResponse and Result in the response are explicit objects. The pagination field MaxResults is of type integer<int32>, and NextToken is of type string. SessionList retains the any type. For session summary fields, see the supplementary description.
//
// @param tmpReq - ListAgentSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentSessionsResponse
func (client *Client) ListAgentSessionsWithContext(ctx context.Context, tmpReq *ListAgentSessionsRequest, runtime *dara.RuntimeOptions) (_result *ListAgentSessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAgentSessionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgentSessions"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentSessionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the agents visible to the current identity by using paging.
//
// Description:
//
// - By default, the API returns the official agents provided by the platform and the tenant agents visible to the current identity.
//
// - If Scope is set to SYSTEM, only official agents are queried. If Scope is set to CUSTOM, only custom agents are queried.
//
// - Supports both cursor-based pagination and page number-based pagination. When using cursor-based pagination, pass the NextToken value exactly as returned in the next request, and keep the calling identity, filter conditions, and MaxResults unchanged.
//
// @param request - ListAgentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentsResponse
func (client *Client) ListAgentsWithContext(ctx context.Context, request *ListAgentsRequest, runtime *dara.RuntimeOptions) (_result *ListAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreatorId) {
		query["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Q) {
		query["Q"] = request.Q
	}

	if !dara.IsNil(request.RequiredRuntime) {
		query["RequiredRuntime"] = request.RequiredRuntime
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgents"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the text models visible to the current caller for selection when setting a session model.
//
// Description:
//
// - For text sessions, pass `ModelType=text` and use `MaxResults` and `NextToken` for paging. Obtain the model identifier from `Models[].Id` and pass it as `Params.modelId` to SetAgentSessionModel.
//
// - Visibility in the model catalog does not mean the target session supports the model. After setting the model, refer to the return value of SetAgentSessionModel or the query result of GetAgentSessionModel.
//
// - `ThinkingLevels` and `ThinkingLevel` are used for optional thinking level configurations and do not affect the integration flow for selecting a model only.
//
// @param request - ListModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelsResponse
func (client *Client) ListModelsWithContext(ctx context.Context, request *ListModelsRequest, runtime *dara.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.ModelType) {
		query["ModelType"] = request.ModelType
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

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
		Action:      dara.String("ListModels"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query for skills visible to the current invoker.
//
// Description:
//
// - Query and filter:
//
//   - `Scope=SYSTEM` queries only official skills, and `Scope=CUSTOM` queries only custom skills. If omitted, both are queried.
//
//   - Custom skills with `user` visibility are visible only to the creator, while skills with `tenant` visibility are visible to the current tenant.
//
//   - Supports filtering by `CreatorId`, `Q`, and `Visibility`. `Q` performs a fuzzy match on the skill name or description.
//
//   - Results are sorted by update time in descending order by default.
//
// - Pagination:
//
//   - Cursor-based pagination uses `MaxResults` and `NextToken`. Do not pass `NextToken` for the first query. For subsequent pages, use the token from the previous response. When using `NextToken` for subsequent pages, `CreatorId`, `Q`, `Visibility`, `Scope`, and `MaxResults` must remain the same as the previous page. If you modify the query conditions, you must restart the query from the first page.
//
//   - Page number-based pagination uses `PageNumber` and `PageSize`. If `MaxResults` is explicitly passed, cursor-based pagination takes precedence. If `NextToken` is passed, `PageNumber` is ignored.
//
// - Returned content:
//
//   - This operation returns only skill summaries and does not generate bundle download URLs. To obtain a download URL, call `GetSkill` and specify `Network`.
//
// @param request - ListSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillsResponse
func (client *Client) ListSkillsWithContext(ctx context.Context, request *ListSkillsRequest, runtime *dara.RuntimeOptions) (_result *ListSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreatorId) {
		query["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Q) {
		query["Q"] = request.Q
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkills"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSkillsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Loads session history, receives subsequent output, or resumes Server-Sent Events (SSE) from a saved read position.
//
// Description:
//
// - A standard load requires only the session ID. It returns the configuration, history, and output to be continuously received. The response is SSE, not a standard JSON message list.
//
// - When resuming from a stream interruption, pass back the actually received cursor, event epoch, and turn ID. History and incremental data may be returned simultaneously. Existing pages should be rebuilt or deduplicated based on the confirmed cursor. Do not directly append all returned data.
//
// - During the load process, you may receive pending approvals or questions again. Use the interaction notification that is still valid for this session to call ReplyAgentSession.
//
// - For an idle session, `stopReason=end_turn` may only indicate the end of the current load. It cannot be used to determine the success of a previously interrupted turn. You must associate it with the actual desired state of the target turn. If it cannot be determined, keep the result for verification.
//
// @param tmpReq - LoadAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoadAgentSessionResponse
func (client *Client) LoadAgentSessionWithSSECtx(ctx context.Context, tmpReq *LoadAgentSessionRequest, runtime *dara.RuntimeOptions, _yield chan *LoadAgentSessionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.loadAgentSessionWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// Loads session history, receives subsequent output, or resumes Server-Sent Events (SSE) from a saved read position.
//
// Description:
//
// - A standard load requires only the session ID. It returns the configuration, history, and output to be continuously received. The response is SSE, not a standard JSON message list.
//
// - When resuming from a stream interruption, pass back the actually received cursor, event epoch, and turn ID. History and incremental data may be returned simultaneously. Existing pages should be rebuilt or deduplicated based on the confirmed cursor. Do not directly append all returned data.
//
// - During the load process, you may receive pending approvals or questions again. Use the interaction notification that is still valid for this session to call ReplyAgentSession.
//
// - For an idle session, `stopReason=end_turn` may only indicate the end of the current load. It cannot be used to determine the success of a previously interrupted turn. You must associate it with the actual desired state of the target turn. If it cannot be determined, keep the result for verification.
//
// @param tmpReq - LoadAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoadAgentSessionResponse
func (client *Client) LoadAgentSessionWithContext(ctx context.Context, tmpReq *LoadAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *LoadAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &LoadAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LoadAgentSession"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoadAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends text to an existing session and continuously receives agent output over Server-Sent Events (SSE).
//
// Description:
//
// - Pass the session ID returned by CreateAgentSession to Params.sessionId and place the user text in Params.prompt.
//
// - When an execution or pending interaction is already in progress, the default call path rejects new prompts. After receiving an approval or Q&A notification, submit your choice or answer through ReplyAgentSession before continuing to read the original stream.
//
// - A connection interruption does not cancel the execution. If the result is unknown, call LoadAgentSession to resume reading. Submitting duplicate prompts may generate new messages, and duplicate Id values are not deduplicated.
//
// @param tmpReq - PromptAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PromptAgentSessionResponse
func (client *Client) PromptAgentSessionWithSSECtx(ctx context.Context, tmpReq *PromptAgentSessionRequest, runtime *dara.RuntimeOptions, _yield chan *PromptAgentSessionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.promptAgentSessionWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// Sends text to an existing session and continuously receives agent output over Server-Sent Events (SSE).
//
// Description:
//
// - Pass the session ID returned by CreateAgentSession to Params.sessionId and place the user text in Params.prompt.
//
// - When an execution or pending interaction is already in progress, the default call path rejects new prompts. After receiving an approval or Q&A notification, submit your choice or answer through ReplyAgentSession before continuing to read the original stream.
//
// - A connection interruption does not cancel the execution. If the result is unknown, call LoadAgentSession to resume reading. Submitting duplicate prompts may generate new messages, and duplicate Id values are not deduplicated.
//
// @param tmpReq - PromptAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PromptAgentSessionResponse
func (client *Client) PromptAgentSessionWithContext(ctx context.Context, tmpReq *PromptAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *PromptAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PromptAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PromptAgentSession"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PromptAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits the approval choice or question answer for the current pending interaction.
//
// Description:
//
// - Read `data.sessionId`, `data.requestId`, and `data.options` from the `permission_request` notification of PromptAgentSession or LoadAgentSession. This operation uses the interaction ID, not the tool invocation ID or the server-side RequestId.
//
// - When selecting an action, use the `optionId` actually provided in the notification. For Q&A scenarios, you must also write the question key and answer text into `answers`. Do not pass this field for regular tool approvals.
//
// - `accepted=true` only indicates that the reply has been accepted. Subsequent output is still read through Prompt/Load. `false` indicates that the reply is not adopted this time. You should reload the current interaction and not automatically change to approve other requests.
//
// @param tmpReq - ReplyAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplyAgentSessionResponse
func (client *Client) ReplyAgentSessionWithContext(ctx context.Context, tmpReq *ReplyAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *ReplyAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReplyAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplyAgentSession"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplyAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves the global workflow configuration.
//
// Description:
//
// Saves or queries the global workflow configuration by tenant and region. The content is sensitive information.
//
// @param request - SaveWorkflowGlobalEnvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveWorkflowGlobalEnvResponse
func (client *Client) SaveWorkflowGlobalEnvWithContext(ctx context.Context, request *SaveWorkflowGlobalEnvRequest, runtime *dara.RuntimeOptions) (_result *SaveWorkflowGlobalEnvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveWorkflowGlobalEnv"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveWorkflowGlobalEnvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the model used by a session and optionally configures the thinking level.
//
// Description:
//
// - Pass the `sessionId` returned by CreateAgentSession and the `Models[].Id` returned by ListModels as `Params.sessionId` and `Params.modelId`, respectively. The visibility of a model in the model catalog does not mean that the target session supports the model. The final validation is performed by the target session.
//
// - `modelId` is required. When you only modify the model, you can omit `thinkingLevel`. For more information about how to configure the thinking level, see the supplementary description of request parameters.
//
// - This operation only modifies the target session and does not modify the global default model of the agent or account. The session may deny configuration modifications during execution, and switching during runtime is not guaranteed to be supported.
//
// - The model and thinking level settings are applied step by step. If the call fails, the model may have already been changed, and an overall rollback is not guaranteed. You should call GetAgentSessionModel to query the current configuration before deciding on subsequent operations. After a successful call, you can read the actual effective values from the returned result.
//
// @param tmpReq - SetAgentSessionModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAgentSessionModelResponse
func (client *Client) SetAgentSessionModelWithContext(ctx context.Context, tmpReq *SetAgentSessionModelRequest, runtime *dara.RuntimeOptions) (_result *SetAgentSessionModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetAgentSessionModelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAgentSessionModel"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAgentSessionModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of an agent.
//
// Description:
//
// - `Name` is used only to locate the agent and cannot be modified after creation.
//
// - Each agent can be bound to only one knowledge base.
//
// - `Tools` is updated according to the rules below. If `Skills` or `KnowledgeBases` is not passed, the original value is retained. If a non-empty array is passed, it replaces the original value entirely. If an empty array is passed, the corresponding binding is removed. Other optional fields retain their original values if not passed.
//
// @param request - UpdateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentResponse
func (client *Client) UpdateAgentWithContext(ctx context.Context, request *UpdateAgentRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.ExpectedVersion) {
		body["ExpectedVersion"] = request.ExpectedVersion
	}

	if !dara.IsNil(request.KnowledgeBases) {
		body["KnowledgeBases"] = request.KnowledgeBases
	}

	if !dara.IsNil(request.Skills) {
		body["Skills"] = request.Skills
	}

	if !dara.IsNil(request.SystemPrompt) {
		body["SystemPrompt"] = request.SystemPrompt
	}

	if !dara.IsNil(request.Tools) {
		body["Tools"] = request.Tools
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgent"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a skill.
//
// Description:
//
// - Update rules:
//
//   - `Name` is used only to locate the skill and cannot be modified.
//
//   - The caller must be the creator of the skill or the root account of the tenant that has permissions to manage the tenant-level skill.
//
// - Replacement rules:
//
//   - If `Description` and `Visibility` are omitted, their original values are retained. Passing `null` is treated as omission and cannot be used to clear the original values.
//
//   - `Metadata` is replaced as a whole, not merged incrementally. If `Metadata` is omitted, the original content is retained. If it is passed, old fields not included in the new object are deleted.
//
//   - When modifying only `Description` or `Visibility`, do not pass `Metadata` or an empty object `{}`. An empty object replaces the original metadata with an empty object.
//
// - Content sources:
//
//   - When replacing the body or bundle, you must choose one and only one of `skillMd`, `transitId`, or `bundleUrl`.
//
//   - For the field formats of the three sources, the Transit upload confirmation process, and the `bundleUrl` restrictions, refer to CreateSkill. Pass the selected source in `UpdateSkill.Metadata`.
//
// @param request - UpdateSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillResponse
func (client *Client) UpdateSkillWithContext(ctx context.Context, request *UpdateSkillRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExpectedVersion) {
		body["ExpectedVersion"] = request.ExpectedVersion
	}

	if !dara.IsNil(request.Metadata) {
		body["Metadata"] = request.Metadata
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkill"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSkillResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) loadAgentSessionWithSSECtx_opYieldFunc(_yield chan *LoadAgentSessionResponse, _yieldErr chan error, ctx context.Context, tmpReq *LoadAgentSessionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &LoadAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LoadAgentSession"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) promptAgentSessionWithSSECtx_opYieldFunc(_yield chan *PromptAgentSessionResponse, _yieldErr chan error, ctx context.Context, tmpReq *PromptAgentSessionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &PromptAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PromptAgentSession"),
		Version:     dara.String("2026-07-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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
