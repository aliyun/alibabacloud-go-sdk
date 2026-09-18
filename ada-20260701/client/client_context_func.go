// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates an Agent.
//
// Description:
//
// - `Name` is the unique identifier of the Agent within the current tenant. It can contain only letters, digits, underscores, and hyphens, and must be 1 to 128 characters in length. The name cannot be modified after creation.
//
// - Each Agent can be associated with only one knowledge base.
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
// Creates a Skill.
//
// Description:
//
// - Creation rules:
//
//   - `Name` must be unique within the current tenant and cannot share a name with an official Skill.
//
//   - `Description` and `Metadata` are required. `Metadata` must contain exactly one valid content source. Different sources cannot be mixed.
//
//   - After the Skill is created, you can modify it by calling `UpdateSkill`.
//
// - Content sources:
//
//   - `skillMd`: Directly provide the Markdown body without a YAML header. Specify the name and description by using `Name` and `Description`.
//
//   - `transitId`: Upload and confirm a ZIP file through Transit. Call the operations in the following order:
//
//     1. Call `CreateTransitUploadPolicy` with `FileShowName` to obtain `TransitId`, `FilePath`, and `PolicyInfo`.
//
//     2. Upload the ZIP file to object storage by using `PolicyInfo` and `FilePath`.
//
//     3. Call `ConfirmTransitUpload` with `TransitId`. Proceed with creation only when the response returns `Confirmed=true`.
//
//     4. Call `CreateSkill` and pass the confirmed `TransitId` in `Metadata.transitId`.
//
//   - `bundleUrl`: Provide a public HTTPS direct link to a ZIP file. The platform downloads the file and saves it as an Artifact. The original URL is not persisted or returned in responses.
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
// Generates a credential for direct file upload. The caller can use the returned upload policy to upload a file directly to object storage.
//
// Description:
//
// - `Network` can be set to `public` or `internal` to generate a public or same-region internal upload address. Default value: `public`.
//
// - The maximum size of a single file is 50 MiB.
//
// - `PolicyInfo` contains short-term upload authorization information intended only for the current file upload. Do not log it, persist it long-term, or forward it to other users.
//
// - `ExpireMs` controls the validity period of the upload policy and the Transit record, in milliseconds. It is not an absolute timestamp. The default and maximum value is `604800000` (7 days), and the minimum value is `1000` (1 second). The validity period is rounded down to the nearest whole second. For example, 1500 milliseconds takes effect as 1 second.
//
// - Call `GetTransitMeta` and read `ExpireAt` to obtain the expiration time of the Transit record. Confirming the upload or querying the record does not extend the validity period.
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
// Deletes a specified Agent.
//
// Description:
//
// - Creators can delete Agents that they created.
//
// - After deletion, the Agent can no longer be queried, updated, or run.
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
// Soft-deletes a Skill that the current caller has permission to modify.
//
// Description:
//
// - Deletion permissions:
//
//   - You can only delete custom Skills that the current caller has permission to modify.
//
//   - The caller must be the Skill creator or the tenant root account that has permission to manage the tenant-level Skill. Official Skills cannot be deleted.
//
// - Deletion results:
//
//   - Deletion uses soft delete. After successful deletion, `GetSkill` and `ListSkills` no longer return the Skill, and you can create a new Skill with the same name.
//
//   - Recovery is not supported. You cannot delete a Skill that has already been deleted or does not exist.
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
// Queries an agent visible to the current identity by name.
//
// Description:
//
// - You can retrieve agents that you created, agents visible within the current tenant, and official agents provided by the platform.
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
// Queries a Skill by name that is visible to the current caller.
//
// Description:
//
// - Query target:
//
//   - Queries by `Name`. Querying by `SkillId` is not supported.
//
//   - Returns the Skill with the matching name that is visible to the current caller under the current tenant first. If no visible record exists, queries the official Skill with the same name.
//
//   - If `SkillVersion` is omitted, the current Skill is returned. This parameter is omitted by default.
//
// - Download URL:
//
//   - `Network` supports `public` and `internal`. If omitted, no download URL is generated.
//
//   - If the Skill has an accessible Artifact, a temporary `DownloadUrl` and the corresponding `DownloadUrlNetwork` are returned.
//
//   - If the Artifact does not exist, is inaccessible, or the URL generation fails, the Skill query still succeeds, but download URL-related fields may not be returned.
//
//   - The original `bundleUrl` used during creation is not stored and is not returned by this operation.
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
// Queries the metadata of a Transit file, including the upload status, file size, and expiration time, and optionally generates a temporary download URL.
//
// Description:
//
// - Specify at least one of `TransitId` and `FilePath`. If both are specified, `TransitId` takes precedence.
//
// - Use the `TransitId` returned by `CreateTransitUploadPolicy` to query the file. `TransitId` is a temporary capability identifier used during the file upload process. Do not share it with unauthorized users.
//
// - `FilePath` is an opaque object path returned by `CreateTransitUploadPolicy`. Use it as-is. Do not parse, modify, or construct it manually.
//
// - When you query by `FilePath`, an error is returned if the record does not exist or is not accessible to the caller.
//
// - `ExpireMs` specifies the validity period of the download URL in milliseconds. Default value: `900000` (15 minutes). The validity period is rounded down to the nearest whole second. For example, 1500 milliseconds is rounded down to 1 second.
//
// - `ExpireAt` is the expiration time of the Transit record, not the expiration time of the download URL. Querying, generating a download URL, and confirming the upload do not extend the record validity period.
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
// Queries agents visible to the current identity by paging.
//
// Description:
//
// - By default, returns official platform agents and tenant agents visible to the current identity.
//
// - Set `Scope=SYSTEM` to query only official agents, or `Scope=CUSTOM` to query only custom agents.
//
// - Both cursor-based pagination and page number-based pagination are supported. When using cursor-based pagination, pass the `NextToken` value from the previous response to the next request, and keep the caller identity, filter conditions, and `MaxResults` unchanged.
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
// Queries Skills visible to the current caller by paging.
//
// Description:
//
// - Query and filtering:
//
//   - Set `Scope=SYSTEM` to query only official Skills, or `Scope=CUSTOM` to query only custom Skills. If omitted, both official and custom Skills are queried.
//
//   - Custom Skills with `user` visibility are visible only to the creator. Skills with `tenant` visibility are visible to the current tenant.
//
//   - Filtering by `CreatorId`, `Q`, and `Visibility` is supported. `Q` performs a fuzzy match on the Skill name or description.
//
//   - Results are sorted by update time in descending order by default.
//
// - Pagination:
//
//   - For cursor-based pagination, use `MaxResults` and `NextToken`. Do not pass `NextToken` for the first query. For subsequent pages, use the token returned in the previous response. When using `NextToken` for subsequent pages, `CreatorId`, `Q`, `Visibility`, `Scope`, and `MaxResults` must remain the same as the previous page. If you change the query conditions, start over from the first page.
//
//   - For page-number-based pagination, use `PageNumber` and `PageSize`. If `MaxResults` is explicitly specified, cursor-based pagination takes precedence. If `NextToken` is specified, `PageNumber` is ignored.
//
// - Response content:
//
//   - This operation returns only Skill summaries and does not generate Bundle download URLs. To obtain download URLs, call `GetSkill` and specify `Network`.
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
// Updates the configuration of an Agent.
//
// Description:
//
// - `Name` is used only to locate the Agent and cannot be modified after creation.
//
// - Each Agent can be bindded to only one knowledge base.
//
// - `Tools` is updated according to the rules below. If `Skills` or `KnowledgeBases` is not specified, the existing value is retained. A non-empty array replaces the entire value. An empty array removes the corresponding binddings. Other optional fields retain their existing values if not specified.
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
// Updates a Skill.
//
// Description:
//
// - Update rules:
//
//   - `Name` is used only to locate the Skill and cannot be modified.
//
//   - The caller must be the Skill creator or the tenant root account that has permissions to manage the tenant-level Skill.
//
// - Replacement rules:
//
//   - `Description` and `Visibility` retain their original values when omitted. Passing `null` is treated the same as omitting the field and cannot be used to clear the original value.
//
//   - `Metadata` is replaced as a whole, not merged incrementally. Omitting `Metadata` preserves the original content. When provided, any old fields not included in the new object are deleted.
//
//   - When modifying only `Description` or `Visibility`, do not pass `Metadata` or an empty object `{}`. An empty object replaces the entire original Metadata with an empty value.
//
// - Content source:
//
//   - When replacing the body or bundle, specify exactly one of `skillMd`, `transitId`, or `bundleUrl`.
//
//   - For field formats, the Transit upload confirmation process, and `bundleUrl` restrictions of the three sources, refer to CreateSkill. Pass the selected source in `UpdateSkill.Metadata`.
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
