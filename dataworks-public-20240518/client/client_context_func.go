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
// Calls this operation to change the status of a specified deployment process to terminated. The deployment process is not deleted and can still be queried through query operations.
//
// Description:
//
//	Notice: This operation may not be available in earlier versions of the SDK. In this case, use the AbolishDeployment operation. The parameters are the same as those described in this document.
//
// @param request - AbolishPipelineRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbolishPipelineRunResponse
func (client *Client) AbolishPipelineRunWithContext(ctx context.Context, request *AbolishPipelineRunRequest, runtime *dara.RuntimeOptions) (_result *AbolishPipelineRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbolishPipelineRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbolishPipelineRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a specified entity object to a Data Map collection. The collection object can be a Data Map category or a data album. The entity object currently supports only the data table type. To add an entity to a data album, the caller must have the AliyunDataWorksFullAccess permission, or be the creator or administrator of the album.
//
// Description:
//
// 1. DataWorks Professional Edition or a more advanced edition is required.
//
// @param request - AddEntityIntoMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEntityIntoMetaCollectionResponse
func (client *Client) AddEntityIntoMetaCollectionWithContext(ctx context.Context, request *AddEntityIntoMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *AddEntityIntoMetaCollectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MetaCollectionId) {
		query["MetaCollectionId"] = request.MetaCollectionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEntityIntoMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEntityIntoMetaCollectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an application for access permissions on a specific resource.
//
// Description:
//
// ## Request Description
//
// - **Reason**: The reason for the application. This parameter is required.
//
// - **ApplyContents**: Contains multiple resource permission application contents, each including the resource description (Resource), grantee description (Grantee), permission types (AccessTypes), and permission expiration time (ExpirationTime). The maximum limit per request is 400 entries.
//
// - **Resource**: The resource description. You need to specify the ResourceSchema.name and version that the resource parsing depends on, as well as the resource metadata MetaData.
//
// - **Grantee**: The grantee description. You need to specify the grantee type (PrincipalType) and the principal ID (PrincipalId).
//
// - **AccessTypes**: The list of permission types. Multiple permission combinations are supported.
//
// - **ExpirationTime**: The permission expiration time, provided as a milliseconds timestamp.
//
// - **AuthMethod**: An optional parameter that specifies the authorization method. The system uses the built-in default authorization method if not specified.
//
// - **ClientToken**: The client token used to prevent duplicate requests. This parameter is optional.
//
// Ensure all required fields are filled in correctly and comply with the corresponding constraints. For example, `DefVersion` and `MetaData` in `Resource` should match the selected `DefSchema`.
//
// @param tmpReq - ApplyResourceAccessPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyResourceAccessPermissionResponse
func (client *Client) ApplyResourceAccessPermissionWithContext(ctx context.Context, tmpReq *ApplyResourceAccessPermissionRequest, runtime *dara.RuntimeOptions) (_result *ApplyResourceAccessPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ApplyResourceAccessPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplyContents) {
		request.ApplyContentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplyContents, dara.String("ApplyContents"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyContentsShrink) {
		body["ApplyContents"] = request.ApplyContentsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Reason) {
		body["Reason"] = request.Reason
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyResourceAccessPermission"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyResourceAccessPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Approves or rejects a specified approval process instance.
//
// Description:
//
// ## Request description
//
// - This operation allows you to approve or reject a specified approval process instance by passing in the ProcessInstanceId and approval information (including ApprovalComment and ApprovalAction).
//
// - ApprovalAction can be `Agree` or `Deny`, indicating approval or rejection respectively.
//
// - ApprovalComment is required and records the specific approval opinion.
//
// @param request - ApproveProcessInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveProcessInstanceResponse
func (client *Client) ApproveProcessInstanceWithContext(ctx context.Context, request *ApproveProcessInstanceRequest, runtime *dara.RuntimeOptions) (_result *ApproveProcessInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalAction) {
		body["ApprovalAction"] = request.ApprovalAction
	}

	if !dara.IsNil(request.ApprovalComment) {
		body["ApprovalComment"] = request.ApprovalComment
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.NewExpiration) {
		body["NewExpiration"] = request.NewExpiration
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApproveProcessInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApproveProcessInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an image with a workspace.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or later to use this operation.
//
// 2. **Ensure the AliyunServiceRoleForDataWorks service-linked role is created before you call this operation.**
//
// @param request - AssociateProjectToImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateProjectToImageResponse
func (client *Client) AssociateProjectToImageWithContext(ctx context.Context, request *AssociateProjectToImageRequest, runtime *dara.RuntimeOptions) (_result *AssociateProjectToImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateProjectToImage"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateProjectToImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a resource group with a workspace.
//
// Description:
//
// 1. This operation requires DataWorks Basic Edition or a more advanced edition.
//
// 2. You must have one of the following roles in the DataWorks workspace:
//
// - tenant owner, workspace administrator, project owner, or operator
//
// @param request - AssociateProjectToResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateProjectToResourceGroupResponse
func (client *Client) AssociateProjectToResourceGroupWithContext(ctx context.Context, request *AssociateProjectToResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *AssociateProjectToResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateProjectToResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateProjectToResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AttachDataQualityRulesToEvaluationTask is deprecated, please use dataworks-public::2024-05-18::UpdateDataQualityScan instead.
//
// Summary:
//
// Associates data quality rules with a data quality check task.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - AttachDataQualityRulesToEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDataQualityRulesToEvaluationTaskResponse
func (client *Client) AttachDataQualityRulesToEvaluationTaskWithContext(ctx context.Context, tmpReq *AttachDataQualityRulesToEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *AttachDataQualityRulesToEvaluationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachDataQualityRulesToEvaluationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataQualityRuleIds) {
		request.DataQualityRuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataQualityRuleIds, dara.String("DataQualityRuleIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityEvaluationTaskId) {
		body["DataQualityEvaluationTaskId"] = request.DataQualityEvaluationTaskId
	}

	if !dara.IsNil(request.DataQualityRuleIdsShrink) {
		body["DataQualityRuleIds"] = request.DataQualityRuleIdsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachDataQualityRulesToEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachDataQualityRulesToEvaluationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates metadata entities in batches. All entities in the same batch must be of the same type. Currently, only custom entity types and extension table types (corresponding to Database/Table) are supported.
//
// Description:
//
// DataWorks Professional Edition or a higher edition is required.
//
// @param tmpReq - BatchCreateMetaEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateMetaEntitiesResponse
func (client *Client) BatchCreateMetaEntitiesWithContext(ctx context.Context, tmpReq *BatchCreateMetaEntitiesRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateMetaEntitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchCreateMetaEntitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Entities) {
		request.EntitiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Entities, dara.String("Entities"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EntitiesShrink) {
		body["Entities"] = request.EntitiesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateMetaEntities"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateMetaEntitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes metadata entity objects in batches. Supports deleting custom entities and extension table type objects (Database/Table). Does not support deleting columns individually. You can delete associated Column objects by deleting the Table.
//
// Description:
//
// Requires DataWorks Professional Edition or a higher edition.
//
// @param tmpReq - BatchDeleteMetaEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteMetaEntitiesResponse
func (client *Client) BatchDeleteMetaEntitiesWithContext(ctx context.Context, tmpReq *BatchDeleteMetaEntitiesRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteMetaEntitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchDeleteMetaEntitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteMetaEntities"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteMetaEntitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates multiple nodes in batches by using incremental updates.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param tmpReq - BatchUpdateTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateTasksResponse
func (client *Client) BatchUpdateTasksWithContext(ctx context.Context, tmpReq *BatchUpdateTasksRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tasks) {
		request.TasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tasks, dara.String("Tasks"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.TasksShrink) {
		body["Tasks"] = request.TasksShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Interrupts the Agent call for a specified session, supporting interruption during streaming responses.
//
// Description:
//
// ## Request description
//
// - This operation is mainly used to actively interrupt an ongoing session, especially when the session is in a streaming response state.
//
// - `sessionId` is a required parameter that identifies the specific session to cancel.
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
		Version:     dara.String("2024-05-18"),
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
// Clones a new data source based on an existing data source.
//
// Description:
//
// 1. You must have purchased DataWorks Basic Edition or a higher edition.
//
// 2. You must have at least one of the following roles in the DataWorks workspace:
//
// - Tenant Owner, Storage Management Administrator, Project Owner, or O&M Engineer
//
// @param request - CloneDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneDataSourceResponse
func (client *Client) CloneDataSourceWithContext(ctx context.Context, request *CloneDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CloneDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloneDataSourceName) {
		query["CloneDataSourceName"] = request.CloneDataSourceName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneDataSource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Agent.
//
// Description:
//
// ## Operation description
//
// - **Agent name**: Must be unique within the current account.
//
// - **Model configuration**: An optional parameter that specifies the model used by the Agent and its related settings.
//
// - **Visibility level**: Defines who can access the Agent. Supported levels include account-wide, project-specific, or user-specific visibility.
//
// - **Visibility scope**: When you set the visibility level to `PROJECT` or `USER`, you must specify the list of project IDs or user IDs.
//
// - **Other parameters**: Parameters such as display name and description are optional. Set them as needed.
//
// @param tmpReq - CreateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentResponse
func (client *Client) CreateAgentWithContext(ctx context.Context, tmpReq *CreateAgentRequest, runtime *dara.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallableAgents) {
		request.CallableAgentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallableAgents, dara.String("CallableAgents"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Metadata) {
		request.MetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Metadata, dara.String("Metadata"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Model) {
		request.ModelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Model, dara.String("Model"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Skills) {
		request.SkillsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Skills, dara.String("Skills"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tools) {
		request.ToolsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tools, dara.String("Tools"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VisibilityScope) {
		request.VisibilityScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VisibilityScope, dara.String("VisibilityScope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallableAgentsShrink) {
		body["CallableAgents"] = request.CallableAgentsShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.MetadataShrink) {
		body["Metadata"] = request.MetadataShrink
	}

	if !dara.IsNil(request.ModelShrink) {
		body["Model"] = request.ModelShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SkillsShrink) {
		body["Skills"] = request.SkillsShrink
	}

	if !dara.IsNil(request.SystemPrompt) {
		body["SystemPrompt"] = request.SystemPrompt
	}

	if !dara.IsNil(request.ToolsShrink) {
		body["Tools"] = request.ToolsShrink
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.VisibilityScopeShrink) {
		body["VisibilityScope"] = request.VisibilityScopeShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgent"),
		Version:     dara.String("2024-05-18"),
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
// Creates a new agent session and returns the session ID.
//
// Description:
//
// ## Request description
//
// - This operation creates a new agent session.
//
// - Use `_meta.agent.agentName` to specify the bound agent name. This parameter is required.
//
//   - dataworks_data_agent: DataWorks built-in agent — Data Agent, which provides intelligent data development AI capabilities covering the entire workflow of data integration, development, O&M, governance, and analytics.
//
//   - dataworks_chatbi_agent: DataWorks built-in agent — ChatBI, which uses natural language processing and intelligent analytics technologies to automate the entire analysis workflow from requirement parsing, data extraction, and automatic code generation to visualization report output through conversational interaction.
//
//   - dataworks_ai_assistant_agent: DataWorks built-in agent — AI Assistant Service, which is a DataWorks enterprise-grade dedicated AI assistant built on open source frameworks such as OpenClaw and Hermes Agent.
//
// - Use `_meta.config.sessionSource` to pass through a session source identifier for subsequent retrieval by source.
//
// - Use `_meta.config.sessionTags[].sessionTagCode` to pass in session tags.
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
		Version:     dara.String("2024-05-18"),
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
// Creates a custom monitoring alert rule.
//
// @param tmpReq - CreateAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertRuleResponse
func (client *Client) CreateAlertRuleWithContext(ctx context.Context, tmpReq *CreateAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAlertRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerCondition) {
		request.TriggerConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerCondition, dara.String("TriggerCondition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.TriggerConditionShrink) {
		query["TriggerCondition"] = request.TriggerConditionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow in DataStudio.
//
// @param request - CreateBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBusinessResponse
func (client *Client) CreateBusinessWithContext(ctx context.Context, request *CreateBusinessRequest, runtime *dara.RuntimeOptions) (_result *CreateBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.UseType) {
		body["UseType"] = request.UseType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBusinessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a component.
//
// Description:
//
//	Notice: This operation does not support batch operations. If you specify multiple publish entities in the parameters, all entities except the first one are ignored.
//
// @param request - CreateComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComponentResponse
func (client *Client) CreateComponentWithContext(ctx context.Context, request *CreateComponentRequest, runtime *dara.RuntimeOptions) (_result *CreateComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComponent"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a compute resource in a specified workspace. The compute resource can be in the development environment or production environment.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// You must have at least one of the following roles in the DataWorks workspace:
//
// Tenant Owner, Workspace Administrator, Project Owner, or O&M.
//
// @param request - CreateComputeResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComputeResourceResponse
func (client *Client) CreateComputeResourceWithContext(ctx context.Context, request *CreateComputeResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateComputeResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionProperties) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !dara.IsNil(request.ConnectionPropertiesMode) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComputeResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComputeResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a metadata crawler and configures the data source, collection scope, resource group, and scheduling method.
//
// Description:
//
// ## Scenarios
//
// Creates a metadata crawler for a specified data source and configures the collection scope, resource group, scheduling method, and extended configurations.
//
// ## Recommended workflow
//
// 1. Call `GetCrawlerTypeCapabilities` to query the crawler types and their configuration capabilities supported in the current region.
//
// 2. Create a crawler by using a data source that matches the `Type` value. Before creating a crawler, ensure that the data source and the selected resource group pass the connectivity test by calling the `TestDataSourceConnectivity` API to avoid creating an invalid crawler.
//
// 3. After the crawler is created, call `RunCrawler` to manually run it, or configure periodic scheduling for automatic execution.
//
// ## Edition requirements
//
// DataWorks Basic Edition or a higher edition is required.
//
// ## Precautions
//
// A successful creation only indicates that the crawler configuration has been generated. Metadata collection is not immediately executed.
//
// @param tmpReq - CreateCrawlerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCrawlerResponse
func (client *Client) CreateCrawlerWithContext(ctx context.Context, tmpReq *CreateCrawlerRequest, runtime *dara.RuntimeOptions) (_result *CreateCrawlerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCrawlerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("Options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Scope) {
		request.ScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scope, dara.String("Scope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.EnableAiComment) {
		body["EnableAiComment"] = request.EnableAiComment
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OptionsShrink) {
		body["Options"] = request.OptionsShrink
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		body["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.ScopeShrink) {
		body["Scope"] = request.ScopeShrink
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCrawler"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCrawlerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom attribute definition.
//
// @param tmpReq - CreateCustomAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomAttributeResponse
func (client *Client) CreateCustomAttributeWithContext(ctx context.Context, tmpReq *CreateCustomAttributeRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCustomAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EntityTypes) {
		request.EntityTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityTypes, dara.String("EntityTypes"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ValueEnums) {
		request.ValueEnumsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ValueEnums, dara.String("ValueEnums"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DisplayEnabled) {
		body["DisplayEnabled"] = request.DisplayEnabled
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EntityTypesShrink) {
		body["EntityTypes"] = request.EntityTypesShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.SearchFilterEnabled) {
		body["SearchFilterEnabled"] = request.SearchFilterEnabled
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.ValueEnumsShrink) {
		body["ValueEnums"] = request.ValueEnumsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomAttribute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert rule for a data integration task.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - CreateDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDIAlarmRuleResponse
func (client *Client) CreateDIAlarmRuleWithContext(ctx context.Context, tmpReq *CreateDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDIAlarmRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDIAlarmRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotificationSettings) {
		request.NotificationSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationSettings, dara.String("NotificationSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerConditions) {
		request.TriggerConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConditions, dara.String("TriggerConditions"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDIAlarmRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDIAlarmRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data integration task.
//
// Description:
//
// - You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// - This operation creates a data integration synchronization task. Parameters include the source configuration SourceDataSourceSettings and the destination configuration DestinationDataSourceSettings, the supported synchronization type MigrationType, transformation rules defined through TransformationRules for mapping operations such as adding columns and renaming tables, specific tables to synchronize and the mapping rules to apply defined in TableMappings, and task-level settings such as column mappings and scheduling configurations defined in JobSettings.
//
// @param tmpReq - CreateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDIJobResponse
func (client *Client) CreateDIJobWithContext(ctx context.Context, tmpReq *CreateDIJobRequest, runtime *dara.RuntimeOptions) (_result *CreateDIJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestinationDataSourceSettings) {
		request.DestinationDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationDataSourceSettings, dara.String("DestinationDataSourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.JobSettings) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, dara.String("JobSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSettings) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, dara.String("ResourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceDataSourceSettings) {
		request.SourceDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceDataSourceSettings, dara.String("SourceDataSourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TableMappings) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, dara.String("TableMappings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TransformationRules) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, dara.String("TransformationRules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationDataSourceType) {
		query["DestinationDataSourceType"] = request.DestinationDataSourceType
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.MigrationType) {
		query["MigrationType"] = request.MigrationType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SourceDataSourceType) {
		query["SourceDataSourceType"] = request.SourceDataSourceType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationDataSourceSettingsShrink) {
		body["DestinationDataSourceSettings"] = request.DestinationDataSourceSettingsShrink
	}

	if !dara.IsNil(request.FileSpec) {
		body["FileSpec"] = request.FileSpec
	}

	if !dara.IsNil(request.JobSettingsShrink) {
		body["JobSettings"] = request.JobSettingsShrink
	}

	if !dara.IsNil(request.ResourceSettingsShrink) {
		body["ResourceSettings"] = request.ResourceSettingsShrink
	}

	if !dara.IsNil(request.SourceDataSourceSettingsShrink) {
		body["SourceDataSourceSettings"] = request.SourceDataSourceSettingsShrink
	}

	if !dara.IsNil(request.TableMappingsShrink) {
		body["TableMappings"] = request.TableMappingsShrink
	}

	if !dara.IsNil(request.TransformationRulesShrink) {
		body["TransformationRules"] = request.TransformationRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDIJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tag.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - CreateDataAssetTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataAssetTagResponse
func (client *Client) CreateDataAssetTagWithContext(ctx context.Context, tmpReq *CreateDataAssetTagRequest, runtime *dara.RuntimeOptions) (_result *CreateDataAssetTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataAssetTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Managers) {
		request.ManagersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Managers, dara.String("Managers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Values) {
		request.ValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Values, dara.String("Values"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.ManagersShrink) {
		query["Managers"] = request.ManagersShrink
	}

	if !dara.IsNil(request.ValueType) {
		query["ValueType"] = request.ValueType
	}

	if !dara.IsNil(request.ValuesShrink) {
		query["Values"] = request.ValuesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataAssetTag"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataAssetTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring alert rule in a specified project.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - CreateDataQualityAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityAlertRuleResponse
func (client *Client) CreateDataQualityAlertRuleWithContext(ctx context.Context, tmpReq *CreateDataQualityAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityAlertRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Condition) {
		body["Condition"] = request.Condition
	}

	if !dara.IsNil(request.NotificationShrink) {
		body["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateDataQualityEvaluationTask is deprecated, please use dataworks-public::2024-05-18::CreateDataQualityScan instead.
//
// Summary:
//
// Creates a DataWorks data quality monitor.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - CreateDataQualityEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityEvaluationTaskResponse
func (client *Client) CreateDataQualityEvaluationTaskWithContext(ctx context.Context, tmpReq *CreateDataQualityEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityEvaluationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityEvaluationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataQualityRules) {
		request.DataQualityRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataQualityRules, dara.String("DataQualityRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hooks) {
		request.HooksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hooks, dara.String("Hooks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notifications) {
		request.NotificationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notifications, dara.String("Notifications"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityRulesShrink) {
		body["DataQualityRules"] = request.DataQualityRulesShrink
	}

	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HooksShrink) {
		body["Hooks"] = request.HooksShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationsShrink) {
		body["Notifications"] = request.NotificationsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeConf) {
		body["RuntimeConf"] = request.RuntimeConf
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityEvaluationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateDataQualityEvaluationTaskInstance is deprecated, please use dataworks-public::2024-05-18::CreateDataQualityScanRun instead.
//
// Summary:
//
// Creates a data quality check task instance.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - CreateDataQualityEvaluationTaskInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityEvaluationTaskInstanceResponse
func (client *Client) CreateDataQualityEvaluationTaskInstanceWithContext(ctx context.Context, tmpReq *CreateDataQualityEvaluationTaskInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityEvaluationTaskInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityEvaluationTaskInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityEvaluationTaskId) {
		body["DataQualityEvaluationTaskId"] = request.DataQualityEvaluationTaskId
	}

	if !dara.IsNil(request.Parameters) {
		body["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityEvaluationTaskInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityEvaluationTaskInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateDataQualityRule is deprecated, please use dataworks-public::2024-05-18::CreateDataQualityScan instead.
//
// Summary:
//
// Creates a data quality rule.
//
// Description:
//
// You must purchase DataWorks Basic Edition or higher to use this feature.
//
// @param tmpReq - CreateDataQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityRuleResponse
func (client *Client) CreateDataQualityRuleWithContext(ctx context.Context, tmpReq *CreateDataQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckingConfig) {
		request.CheckingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckingConfig, dara.String("CheckingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ErrorHandlers) {
		request.ErrorHandlersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ErrorHandlers, dara.String("ErrorHandlers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SamplingConfig) {
		request.SamplingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SamplingConfig, dara.String("SamplingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckingConfigShrink) {
		body["CheckingConfig"] = request.CheckingConfigShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ErrorHandlersShrink) {
		body["ErrorHandlers"] = request.ErrorHandlersShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SamplingConfigShrink) {
		body["SamplingConfig"] = request.SamplingConfigShrink
	}

	if !dara.IsNil(request.Severity) {
		body["Severity"] = request.Severity
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateDataQualityRuleTemplate is deprecated, please use dataworks-public::2024-05-18::CreateDataQualityTemplate instead.
//
// Summary:
//
// Creates a rule template.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - CreateDataQualityRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityRuleTemplateResponse
func (client *Client) CreateDataQualityRuleTemplateWithContext(ctx context.Context, tmpReq *CreateDataQualityRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityRuleTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityRuleTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckingConfig) {
		request.CheckingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckingConfig, dara.String("CheckingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SamplingConfig) {
		request.SamplingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SamplingConfig, dara.String("SamplingConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckingConfigShrink) {
		body["CheckingConfig"] = request.CheckingConfigShrink
	}

	if !dara.IsNil(request.DirectoryPath) {
		body["DirectoryPath"] = request.DirectoryPath
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SamplingConfigShrink) {
		body["SamplingConfig"] = request.SamplingConfigShrink
	}

	if !dara.IsNil(request.VisibleScope) {
		body["VisibleScope"] = request.VisibleScope
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityRuleTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityRuleTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality monitor.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param tmpReq - CreateDataQualityScanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityScanResponse
func (client *Client) CreateDataQualityScanWithContext(ctx context.Context, tmpReq *CreateDataQualityScanRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityScanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityScanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ComputeResource) {
		request.ComputeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComputeResource, dara.String("ComputeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hooks) {
		request.HooksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hooks, dara.String("Hooks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ComputeResourceShrink) {
		body["ComputeResource"] = request.ComputeResourceShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HooksShrink) {
		body["Hooks"] = request.HooksShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityScan"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityScanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers the execution of a specified data quality monitoring task and returns the run instance ID.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param tmpReq - CreateDataQualityScanRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityScanRunResponse
func (client *Client) CreateDataQualityScanRunWithContext(ctx context.Context, tmpReq *CreateDataQualityScanRunRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityScanRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityScanRunShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityScanId) {
		body["DataQualityScanId"] = request.DataQualityScanId
	}

	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityScanRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityScanRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality template.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - CreateDataQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityTemplateResponse
func (client *Client) CreateDataQualityTemplateWithContext(ctx context.Context, request *CreateDataQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data source in a specified project workspace for the development or production environment.
//
// Description:
//
// 1. You must have purchased DataWorks Basic Edition or a higher edition.
//
// 2. You must have at least one of the following roles in the DataWorks project workspace:
//
// - Tenant Owner, Storage Management Administrator, Project Owner, or O&M
//
// @param request - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithContext(ctx context.Context, request *CreateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionProperties) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !dara.IsNil(request.ConnectionPropertiesMode) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a sharing rule for a data source to share it with other workspaces or RAM users.
//
// Description:
//
// 1. This operation is available for all DataWorks editions.
//
// 2. To share a data source from Workspace A to Workspace B, you must have the data source sharing permissions in both workspaces. You must have one of the following roles in DataWorks:
//
// - Tenant Owner, Tenant Administrator, Workspace Administrator, and Workspace Owner
//
// @param request - CreateDataSourceSharedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceSharedRuleResponse
func (client *Client) CreateDataSourceSharedRuleWithContext(ctx context.Context, request *CreateDataSourceSharedRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceSharedRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.SharedUser) {
		query["SharedUser"] = request.SharedUser
	}

	if !dara.IsNil(request.TargetProjectId) {
		query["TargetProjectId"] = request.TargetProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSourceSharedRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSourceSharedRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataset. You can create a dataset only in a workspace that you have joined. Only DataWorks datasets are supported. The maximum number of datasets per tenant is 2000.
//
// @param tmpReq - CreateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithContext(ctx context.Context, tmpReq *CreateDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InitVersion) {
		request.InitVersionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InitVersion, dara.String("InitVersion"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
	}

	if !dara.IsNil(request.InitVersionShrink) {
		body["InitVersion"] = request.InitVersionShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Origin) {
		body["Origin"] = request.Origin
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StorageType) {
		body["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dataset version. Currently, only DataWorks datasets are supported. The maximum number of versions is 20.
//
// @param tmpReq - CreateDatasetVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetVersionResponse
func (client *Client) CreateDatasetVersionWithContext(ctx context.Context, tmpReq *CreateDatasetVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDatasetVersionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImportInfo) {
		request.ImportInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImportInfo, dara.String("ImportInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.ImportInfoShrink) {
		body["ImportInfo"] = request.ImportInfoShrink
	}

	if !dara.IsNil(request.MountPath) {
		body["MountPath"] = request.MountPath
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatasetVersion"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file in DataStudio. This operation does not support creating Data Integration nodes.
//
// @param request - CreateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileResponse
func (client *Client) CreateFileWithContext(ctx context.Context, request *CreateFileRequest, runtime *dara.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedSettings) {
		body["AdvancedSettings"] = request.AdvancedSettings
	}

	if !dara.IsNil(request.ApplyScheduleImmediately) {
		body["ApplyScheduleImmediately"] = request.ApplyScheduleImmediately
	}

	if !dara.IsNil(request.AutoParsing) {
		body["AutoParsing"] = request.AutoParsing
	}

	if !dara.IsNil(request.AutoRerunIntervalMillis) {
		body["AutoRerunIntervalMillis"] = request.AutoRerunIntervalMillis
	}

	if !dara.IsNil(request.AutoRerunTimes) {
		body["AutoRerunTimes"] = request.AutoRerunTimes
	}

	if !dara.IsNil(request.ConnectionName) {
		body["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.CreateFolderIfNotExists) {
		body["CreateFolderIfNotExists"] = request.CreateFolderIfNotExists
	}

	if !dara.IsNil(request.CronExpress) {
		body["CronExpress"] = request.CronExpress
	}

	if !dara.IsNil(request.CycleType) {
		body["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.DependentNodeIdList) {
		body["DependentNodeIdList"] = request.DependentNodeIdList
	}

	if !dara.IsNil(request.DependentType) {
		body["DependentType"] = request.DependentType
	}

	if !dara.IsNil(request.EndEffectDate) {
		body["EndEffectDate"] = request.EndEffectDate
	}

	if !dara.IsNil(request.FileDescription) {
		body["FileDescription"] = request.FileDescription
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.IgnoreParentSkipRunningProperty) {
		body["IgnoreParentSkipRunningProperty"] = request.IgnoreParentSkipRunningProperty
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InputList) {
		body["InputList"] = request.InputList
	}

	if !dara.IsNil(request.InputParameters) {
		body["InputParameters"] = request.InputParameters
	}

	if !dara.IsNil(request.OutputList) {
		body["OutputList"] = request.OutputList
	}

	if !dara.IsNil(request.OutputParameters) {
		body["OutputParameters"] = request.OutputParameters
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParaValue) {
		body["ParaValue"] = request.ParaValue
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.RerunMode) {
		body["RerunMode"] = request.RerunMode
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceGroupIdentifier) {
		body["ResourceGroupIdentifier"] = request.ResourceGroupIdentifier
	}

	if !dara.IsNil(request.SchedulerType) {
		body["SchedulerType"] = request.SchedulerType
	}

	if !dara.IsNil(request.StartEffectDate) {
		body["StartEffectDate"] = request.StartEffectDate
	}

	if !dara.IsNil(request.StartImmediately) {
		body["StartImmediately"] = request.StartImmediately
	}

	if !dara.IsNil(request.Stop) {
		body["Stop"] = request.Stop
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a folder.
//
// @param request - CreateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithContext(ctx context.Context, request *CreateFolderRequest, runtime *dara.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderPath) {
		body["FolderPath"] = request.FolderPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFolder"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a UDF function in DataStudio. The UDF function information is described in FlowSpec format.
//
// Description:
//
//	Notice: This operation does not support batch operations. If more than one UDF function is defined in the FlowSpec, all functions after the first one are ignored.
//
// @param request - CreateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionResponse
func (client *Client) CreateFunctionWithContext(ctx context.Context, request *CreateFunctionRequest, runtime *dara.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an identity credential.
//
// Description:
//
//	Notice:
//
// This operation does not support batch processing. If you specify multiple entities in the request parameters, only the first entity is processed and the rest are ignored.
//
// @param tmpReq - CreateIdentifyCredentialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentifyCredentialResponse
func (client *Client) CreateIdentifyCredentialWithContext(ctx context.Context, tmpReq *CreateIdentifyCredentialRequest, runtime *dara.RuntimeOptions) (_result *CreateIdentifyCredentialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateIdentifyCredentialShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IdentifyCredential) {
		request.IdentifyCredentialShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IdentifyCredential, dara.String("IdentifyCredential"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdentifyCredentialShrink) {
		body["IdentifyCredential"] = request.IdentifyCredentialShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIdentifyCredential"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdentifyCredentialResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a data lineage relationship in DataWorks Data Map. You can use this operation to establish lineage relationships between metadata entities managed by DataWorks, including table-to-table, column-to-column, table-to-column, and dataset-to-table scenarios. You can also establish lineage relationships between managed entities and custom entity objects registered by users. This operation is compatible with non-managed custom objects, but this approach is no longer recommended. Before calling this operation, make sure that the managed entities involved in the lineage registration already exist on the DataWorks platform.
//
// Description:
//
// 1. DataWorks Professional Edition or a higher edition is required.
//
// @param tmpReq - CreateLineageRelationshipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLineageRelationshipResponse
func (client *Client) CreateLineageRelationshipWithContext(ctx context.Context, tmpReq *CreateLineageRelationshipRequest, runtime *dara.RuntimeOptions) (_result *CreateLineageRelationshipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLineageRelationshipShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DstEntity) {
		request.DstEntityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DstEntity, dara.String("DstEntity"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SrcEntity) {
		request.SrcEntityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SrcEntity, dara.String("SrcEntity"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Task) {
		request.TaskShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Task, dara.String("Task"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DstEntityShrink) {
		query["DstEntity"] = request.DstEntityShrink
	}

	if !dara.IsNil(request.SrcEntityShrink) {
		query["SrcEntity"] = request.SrcEntityShrink
	}

	if !dara.IsNil(request.TaskShrink) {
		query["Task"] = request.TaskShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLineageRelationship"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLineageRelationshipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an MCP Server.
//
// Description:
//
// ## Operation description
//
// - After submission, the system verifies the availability of the MCP Server based on the provided connection information.
//
// - If the MCP Server connection is unavailable, the operation returns the corresponding error message.
//
// - The Name field must start with a lowercase letter and can contain only lowercase letters, digits, underscores (_), and hyphens (-). The name must be unique within the current account.
//
// - The Visibility field defines the visibility level of the MCP Server. Valid values: `TENANT` (visible within the account), `PROJECT` (visible to specified projects), and `USER` (visible to specified users). Depending on the selected value, provide the corresponding `VisibilityScope` parameter to further specify the visibility scope.
//
// @param tmpReq - CreateMcpServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcpServerResponse
func (client *Client) CreateMcpServerWithContext(ctx context.Context, tmpReq *CreateMcpServerRequest, runtime *dara.RuntimeOptions) (_result *CreateMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMcpServerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VisibilityScope) {
		request.VisibilityScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VisibilityScope, dara.String("VisibilityScope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigShrink) {
		body["Config"] = request.ConfigShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.VisibilityScopeShrink) {
		body["VisibilityScope"] = request.VisibilityScopeShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcpServer"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcpServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a collection in Data Map. Collections include categories, subcategories, data albums, and categories that are created in the data albums.
//
// Description:
//
// 1. DataWorks Professional Edition or a higher edition is required.
//
// @param request - CreateMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetaCollectionResponse
func (client *Client) CreateMetaCollectionWithContext(ctx context.Context, request *CreateMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *CreateMetaCollectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetaCollectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a metadata entity definition, including custom types and extended table types.
//
// Description:
//
// DataWorks Professional Edition or a more advanced edition is required.
//
// @param tmpReq - CreateMetaEntityDefRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetaEntityDefResponse
func (client *Client) CreateMetaEntityDefWithContext(ctx context.Context, tmpReq *CreateMetaEntityDefRequest, runtime *dara.RuntimeOptions) (_result *CreateMetaEntityDefResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMetaEntityDefShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AttributeDefs) {
		request.AttributeDefsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AttributeDefs, dara.String("AttributeDefs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttributeDefsShrink) {
		body["AttributeDefs"] = request.AttributeDefsShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Extend) {
		body["Extend"] = request.Extend
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetaEntityDef"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetaEntityDefResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network and associates the network with a general resource group.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkResponse
func (client *Client) CreateNetworkWithContext(ctx context.Context, request *CreateNetworkRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchId) {
		body["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetwork"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data development node in the new version of DataStudio.
//
// Description:
//
//	Notice: This operation does not support batch operations. If more than one node is defined in FlowSpec, all nodes after the first one are ignored.
//
// @param request - CreateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeResponse
func (client *Client) CreateNodeWithContext(ctx context.Context, request *CreateNodeRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContainerId) {
		body["ContainerId"] = request.ContainerId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a parameter.
//
// Description:
//
// This operation requires DataWorks Professional Edition or a later edition.
//
// @param tmpReq - CreateParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParameterResponse
func (client *Client) CreateParameterWithContext(ctx context.Context, tmpReq *CreateParameterRequest, runtime *dara.RuntimeOptions) (_result *CreateParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Properties) {
		request.PropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Properties, dara.String("Properties"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.PropertiesShrink) {
		body["Properties"] = request.PropertiesShrink
	}

	if !dara.IsNil(request.Scope) {
		body["Scope"] = request.Scope
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateParameter"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateParameterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a publish process for an entity in the new-version DataStudio.
//
// Description:
//
//	Notice: This operation does not support batch operations. If you specify multiple publish entities in the parameters, all entities except the first one are ignored.
//
//	Notice: This operation may not be available in earlier versions of the SDK. In this case, use the CreateDeployment operation. The parameters are the same as those described in this topic.
//
// @param tmpReq - CreatePipelineRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineRunResponse
func (client *Client) CreatePipelineRunWithContext(ctx context.Context, tmpReq *CreatePipelineRunRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePipelineRunShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectIds) {
		request.ObjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectIds, dara.String("ObjectIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoRunUntilStage) {
		body["AutoRunUntilStage"] = request.AutoRunUntilStage
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ObjectIdsShrink) {
		body["ObjectIds"] = request.ObjectIdsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RunMode) {
		body["RunMode"] = request.RunMode
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipelineRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new approval process definition, supporting custom configuration of approval rules and notification services.
//
// Description:
//
// ## Usage notes
//
// - This API operation allows you to create a new approval process definition, including setting basic information such as the approval policy name, description, type, and subtype.
//
// - You can define a list of condition rules (RuleConditions) to specify the conditions under which the approval process is triggered.
//
// - Multiple notification services (NotificationServices) can be configured to send notifications to relevant personnel at different stages of the approval process.
//
// - The approval node list (ApprovalNodes) defines the nodes that must be traversed during the approval process and the approver information for each node.
//
// - You can choose whether to immediately enable the newly created approval process definition.
//
// - Note: Certain fields such as Type have specific value constraints. Refer to the constraint descriptions in the documentation.
//
// @param tmpReq - CreateProcessDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProcessDefinitionResponse
func (client *Client) CreateProcessDefinitionWithContext(ctx context.Context, tmpReq *CreateProcessDefinitionRequest, runtime *dara.RuntimeOptions) (_result *CreateProcessDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProcessDefinitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApprovalNodes) {
		request.ApprovalNodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovalNodes, dara.String("ApprovalNodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NotificationServices) {
		request.NotificationServicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationServices, dara.String("NotificationServices"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuleConditions) {
		request.RuleConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuleConditions, dara.String("RuleConditions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalNodesShrink) {
		body["ApprovalNodes"] = request.ApprovalNodesShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationServicesShrink) {
		body["NotificationServices"] = request.NotificationServicesShrink
	}

	if !dara.IsNil(request.RuleConditionsShrink) {
		body["RuleConditions"] = request.RuleConditionsShrink
	}

	if !dara.IsNil(request.SubType) {
		body["SubType"] = request.SubType
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProcessDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProcessDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param tmpReq - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithContext(ctx context.Context, tmpReq *CreateProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AliyunResourceTags) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, dara.String("AliyunResourceTags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunResourceGroupId) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !dara.IsNil(request.AliyunResourceTagsShrink) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DevEnvironmentEnabled) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !dara.IsNil(request.DevRoleDisabled) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PaiTaskEnabled) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a workspace member and grants workspace roles to the member.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param tmpReq - CreateProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectMemberResponse
func (client *Client) CreateProjectMemberWithContext(ctx context.Context, tmpReq *CreateProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("RoleCodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProjectMember"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom role for a workspace.
//
// @param tmpReq - CreateProjectRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectRoleResponse
func (client *Client) CreateProjectRoleWithContext(ctx context.Context, tmpReq *CreateProjectRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProjectRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ModulePermissions) {
		request.ModulePermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModulePermissions, dara.String("ModulePermissions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ModulePermissionsShrink) {
		query["ModulePermissions"] = request.ModulePermissionsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProjectRole"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file resource for data development. The file resource information is defined in FlowSpec format.
//
// Description:
//
//	Notice: This operation does not support batch operations. If more than one resource file is defined in the FlowSpec, all resource files after the first one are ignored.
//
// @param request - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithContext(ctx context.Context, request *CreateResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceFile) {
		body["ResourceFile"] = request.ResourceFile
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Supports users in specifying their own files (such as JAR, PY, archive, or file) to create Data Development resource files.
//
// @param request - CreateResourceFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceFileResponse
func (client *Client) CreateResourceFileWithContext(ctx context.Context, request *CreateResourceFileRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceFileResponse, _err error) {
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

	if !dara.IsNil(request.FileDescription) {
		body["FileDescription"] = request.FileDescription
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.OriginResourceName) {
		body["OriginResourceName"] = request.OriginResourceName
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegisterToCalcEngine) {
		body["RegisterToCalcEngine"] = request.RegisterToCalcEngine
	}

	if !dara.IsNil(request.ResourceFile) {
		body["ResourceFile"] = request.ResourceFile
	}

	if !dara.IsNil(request.StorageURL) {
		body["StorageURL"] = request.StorageURL
	}

	if !dara.IsNil(request.UploadMode) {
		body["UploadMode"] = request.UploadMode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource group.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// 2. **Before you call this operation, make sure that you fully understand the billing of DataWorks common resource groups and the [pricing](https://help.aliyun.com/document_detail/2680173.html).**
//
// 3. **Before you call this operation, make sure that you have created the service-linked role AliyunServiceRoleForDataWorks.**
//
// @param tmpReq - CreateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithContext(ctx context.Context, tmpReq *CreateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AliyunResourceTags) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, dara.String("AliyunResourceTags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunResourceGroupId) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !dara.IsNil(request.AliyunResourceTagsShrink) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !dara.IsNil(request.AutoRenewEnabled) {
		body["AutoRenewEnabled"] = request.AutoRenewEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PaymentDuration) {
		body["PaymentDuration"] = request.PaymentDuration
	}

	if !dara.IsNil(request.PaymentDurationUnit) {
		body["PaymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !dara.IsNil(request.PaymentType) {
		body["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchId) {
		body["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route for a network.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteResponse
func (client *Client) CreateRouteWithContext(ctx context.Context, request *CreateRouteRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidr) {
		body["DestinationCidr"] = request.DestinationCidr
	}

	if !dara.IsNil(request.NetworkId) {
		body["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new security control policy to configure various modules and submodules. Requires both DataWorks tenant administrator and security administrator permissions.
//
// Description:
//
// ## Request
//
// - **SchemaName**: Select a schema that fits your business needs.
//
// - **ControlModule*	- and **ControlSubModule**: Specify the module and submodule for the policy, ensuring they match the selected schema.
//
// - **ControlDwScope**: Set the policy scope to either the tenant or workspace level.
//
// - **Workspaces**: If `ControlDwScope` is set to `Workspace`, provide the corresponding workspace IDs.
//
// - **Content.Controllers**: The controllers must match the definitions in the selected schema.
//
// - This operation cannot create system default policies.
//
// @param tmpReq - CreateSecurityStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecurityStrategyResponse
func (client *Client) CreateSecurityStrategyWithContext(ctx context.Context, tmpReq *CreateSecurityStrategyRequest, runtime *dara.RuntimeOptions) (_result *CreateSecurityStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSecurityStrategyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Workspaces) {
		request.WorkspacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Workspaces, dara.String("Workspaces"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ContentShrink) {
		body["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.ControlDwScope) {
		body["ControlDwScope"] = request.ControlDwScope
	}

	if !dara.IsNil(request.ControlModule) {
		body["ControlModule"] = request.ControlModule
	}

	if !dara.IsNil(request.ControlSubModule) {
		body["ControlSubModule"] = request.ControlSubModule
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SchemaName) {
		body["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.WorkspacesShrink) {
		body["Workspaces"] = request.WorkspacesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecurityStrategy"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecurityStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves a reusable semantic job definition. If you use a single-file source, apply for and complete the attachment upload first. After creation, call RunSemanticJob with the returned Name.
//
// Description:
//
// ## Scenarios
//
// Creates and saves a reusable semantic job definition. This operation only saves the data source, resource group, and reference file configurations without immediately executing the job.
//
// ## Recommended workflow
//
// 1. When `Source.type=singleTableFile`, call `UploadSemanticFile` first, use the returned `Data.UploadUrl` to complete the PUT upload, and then specify `Data.FileId` in `ReferenceFileIds`. Alternatively, you can provide a single accessible URI.
//
// 2. Configure `Source`, `ProjectId`, and `ResourceGroupId`, and then call this operation to save the job.
//
// 3. Use `Data.Name` from the response to call `RunSemanticJob`. After the job is complete, use `DownloadSemanticResults` to retrieve the output.
//
// ## Before you begin
//
// `Name` must be unique within the current tenant. The reference file quantity rules differ between single-file sources and other sources. For details, refer to the descriptions of the `ReferenceFileIds` and `ReferenceFileUris` fields.
//
// @param tmpReq - CreateSemanticJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSemanticJobResponse
func (client *Client) CreateSemanticJobWithContext(ctx context.Context, tmpReq *CreateSemanticJobRequest, runtime *dara.RuntimeOptions) (_result *CreateSemanticJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSemanticJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ReferenceFileIds) {
		request.ReferenceFileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceFileIds, dara.String("ReferenceFileIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ReferenceFileUris) {
		request.ReferenceFileUrisShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceFileUris, dara.String("ReferenceFileUris"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Source) {
		request.SourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Source, dara.String("Source"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ReferenceFileIdsShrink) {
		body["ReferenceFileIds"] = request.ReferenceFileIdsShrink
	}

	if !dara.IsNil(request.ReferenceFileUrisShrink) {
		body["ReferenceFileUris"] = request.ReferenceFileUrisShrink
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceShrink) {
		body["Source"] = request.SourceShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSemanticJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSemanticJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new Skill in DataWorks.
//
// Description:
//
// ## Request description
//
// - You must provide either SkillMdOverride or BundleUrl. One of the two parameters is required.
//
// - Visibility can be set to `TENANT`, `PROJECT`, or `USER`, which indicate visibility within the account, visibility to specified projects, or visibility to specified users, respectively.
//
// - When Visibility is set to `PROJECT`, specify the list of visible project IDs by using VisibilityScope.ProjectIds. When Visibility is set to `USER`, specify the list of visible user IDs by using VisibilityScope.UserIds.
//
// @param tmpReq - CreateSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillResponse
func (client *Client) CreateSkillWithContext(ctx context.Context, tmpReq *CreateSkillRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VisibilityScope) {
		request.VisibilityScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VisibilityScope, dara.String("VisibilityScope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BundleUrl) {
		body["BundleUrl"] = request.BundleUrl
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["Extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SkillMdOverride) {
		body["SkillMdOverride"] = request.SkillMdOverride
	}

	if !dara.IsNil(request.VersionNote) {
		body["VersionNote"] = request.VersionNote
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.VisibilityScopeShrink) {
		body["VisibilityScope"] = request.VisibilityScopeShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkill"),
		Version:     dara.String("2024-05-18"),
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
// Creates a file for a function in DataStudio.
//
// @param request - CreateUdfFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfFileResponse
func (client *Client) CreateUdfFileWithContext(ctx context.Context, request *CreateUdfFileRequest, runtime *dara.RuntimeOptions) (_result *CreateUdfFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		body["ClassName"] = request.ClassName
	}

	if !dara.IsNil(request.CmdDescription) {
		body["CmdDescription"] = request.CmdDescription
	}

	if !dara.IsNil(request.CreateFolderIfNotExists) {
		body["CreateFolderIfNotExists"] = request.CreateFolderIfNotExists
	}

	if !dara.IsNil(request.Example) {
		body["Example"] = request.Example
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FunctionType) {
		body["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.ParameterDescription) {
		body["ParameterDescription"] = request.ParameterDescription
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.Resources) {
		body["Resources"] = request.Resources
	}

	if !dara.IsNil(request.ReturnValue) {
		body["ReturnValue"] = request.ReturnValue
	}

	if !dara.IsNil(request.UdfDescription) {
		body["UdfDescription"] = request.UdfDescription
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUdfFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUdfFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow in a specified folder in DataStudio.
//
// Description:
//
//	Notice: This operation does not support batch operations. If more than one workflow is defined in FlowSpec, all workflows except the first one are ignored. In addition, nodes defined within the workflow are also ignored. Call the CreateNode operation to create internal nodes one by one.
//
// @param request - CreateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowDefinitionResponse
func (client *Client) CreateWorkflowDefinitionWithContext(ctx context.Context, request *CreateWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkflowDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkflowDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow instance, such as a data backfill workflow instance, based on configurations.
//
// Description:
//
// DataWorks Basic Edition or higher is required.
//
// @param tmpReq - CreateWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowInstancesResponse
func (client *Client) CreateWorkflowInstancesWithContext(ctx context.Context, tmpReq *CreateWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DefaultRunProperties) {
		request.DefaultRunPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultRunProperties, dara.String("DefaultRunProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Periods) {
		request.PeriodsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Periods, dara.String("Periods"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoStartEnabled) {
		body["AutoStartEnabled"] = request.AutoStartEnabled
	}

	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DefaultRunPropertiesShrink) {
		body["DefaultRunProperties"] = request.DefaultRunPropertiesShrink
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PeriodsShrink) {
		body["Periods"] = request.PeriodsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TagCreationPolicy) {
		body["TagCreationPolicy"] = request.TagCreationPolicy
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskParameters) {
		body["TaskParameters"] = request.TaskParameters
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkflowParameters) {
		body["WorkflowParameters"] = request.WorkflowParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkflowInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Agent.
//
// Description:
//
// ## Operation description
//
// This API operation deletes an Agent with the specified name from DataWorks. When calling this operation, you must provide the name of the Agent to delete.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgent"),
		Version:     dara.String("2024-05-18"),
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
// Deletes a custom monitoring alert rule.
//
// @param request - DeleteAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertRuleResponse
func (client *Client) DeleteAlertRuleWithContext(ctx context.Context, request *DeleteAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workflow.
//
// @param request - DeleteBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBusinessResponse
func (client *Client) DeleteBusinessWithContext(ctx context.Context, request *DeleteBusinessRequest, runtime *dara.RuntimeOptions) (_result *DeleteBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBusinessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a certificate file.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// 2. You must have at least one of the following roles in the DataWorks workspace: tenant owner, storage management administrator, project owner, or O&M engineer.
//
// @param request - DeleteCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCertificateResponse
func (client *Client) DeleteCertificateWithContext(ctx context.Context, request *DeleteCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCertificate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a component.
//
// Description:
//
//	Notice: After a UDF function is published, it cannot be deleted. You must offline the function before deleting it.
//
// @param request - DeleteComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComponentResponse
func (client *Client) DeleteComponentWithContext(ctx context.Context, request *DeleteComponentRequest, runtime *dara.RuntimeOptions) (_result *DeleteComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentId) {
		body["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComponent"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified computing resource based on the computing resource ID.
//
// Description:
//
// 1. DataWorks Basic Edition or a more advanced edition is required.
//
// 2. You must have at least one of the following roles in the DataWorks workspace:
//
// 3. Tenant Owner, Workspace Administrator, Project Owner, O\\&M
//
// @param request - DeleteComputeResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComputeResourceResponse
func (client *Client) DeleteComputeResourceWithContext(ctx context.Context, request *DeleteComputeResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteComputeResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComputeResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComputeResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified metadata crawler.
//
// Description:
//
// ## Scenarios
//
// Delete metadata crawlers that are no longer in use.
//
// ## Recommended procedure
//
// 1. Call `ListCrawlers` to query the crawler ID.
//
// 2. After confirming that the crawler is no longer needed, call this operation.
//
// ## Edition requirements
//
// DataWorks Basic Edition or higher is required.
//
// ## Precautions
//
// After the crawler is deleted, it cannot be queried, updated, or run. The collected metadata is cleaned up by the system, and the cleanup result may be delayed.
//
// @param request - DeleteCrawlerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCrawlerResponse
func (client *Client) DeleteCrawlerWithContext(ctx context.Context, request *DeleteCrawlerRequest, runtime *dara.RuntimeOptions) (_result *DeleteCrawlerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCrawler"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCrawlerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom attribute definition.
//
// @param request - DeleteCustomAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomAttributeResponse
func (client *Client) DeleteCustomAttributeWithContext(ctx context.Context, request *DeleteCustomAttributeRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomAttribute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert rule configured for a synchronization task.
//
// @param request - DeleteDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDIAlarmRuleResponse
func (client *Client) DeleteDIAlarmRuleWithContext(ctx context.Context, request *DeleteDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDIAlarmRuleResponse, _err error) {
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
		Action:      dara.String("DeleteDIAlarmRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDIAlarmRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a new-version synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDIJobResponse
func (client *Client) DeleteDIJobWithContext(ctx context.Context, request *DeleteDIJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteDIJobResponse, _err error) {
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
		Action:      dara.String("DeleteDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDIJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a tag.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - DeleteDataAssetTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataAssetTagResponse
func (client *Client) DeleteDataAssetTagWithContext(ctx context.Context, tmpReq *DeleteDataAssetTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataAssetTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDataAssetTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Values) {
		request.ValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Values, dara.String("Values"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.ValuesShrink) {
		query["Values"] = request.ValuesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataAssetTag"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataAssetTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitoring alert rule based on the specified ID.
//
// Description:
//
// DataWorks Basic Edition or a more advanced edition is required.
//
// @param request - DeleteDataQualityAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityAlertRuleResponse
func (client *Client) DeleteDataQualityAlertRuleWithContext(ctx context.Context, request *DeleteDataQualityAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteDataQualityEvaluationTask is deprecated, please use dataworks-public::2024-05-18::DeleteDataQualityScan instead.
//
// Summary:
//
// Deletes a data quality monitoring task.
//
// Description:
//
// You must purchase DataWorks Basic Edition or higher to use this feature.
//
// @param request - DeleteDataQualityEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityEvaluationTaskResponse
func (client *Client) DeleteDataQualityEvaluationTaskWithContext(ctx context.Context, request *DeleteDataQualityEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityEvaluationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityEvaluationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitoring rule.
//
// Description:
//
// 需要购买DataWorks基础版及以上版本才能使用
//
// @param request - DeleteDataQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityRuleResponse
func (client *Client) DeleteDataQualityRuleWithContext(ctx context.Context, request *DeleteDataQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteDataQualityRuleTemplate is deprecated, please use dataworks-public::2024-05-18::DeleteDataQualityTemplate instead.
//
// Summary:
//
// Deletes a custom data quality rule template.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param request - DeleteDataQualityRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityRuleTemplateResponse
func (client *Client) DeleteDataQualityRuleTemplateWithContext(ctx context.Context, request *DeleteDataQualityRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityRuleTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityRuleTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityRuleTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitor.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - DeleteDataQualityScanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityScanResponse
func (client *Client) DeleteDataQualityScanWithContext(ctx context.Context, request *DeleteDataQualityScanRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityScanResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityScan"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityScanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality rule template by ID.
//
// Description:
//
// ## Request description
//
// - **Id**: The unique identifier of a custom rule template, in the format of `USER_DEFINED:<template_id>`.
//
// - **ProjectId**: The ID of the DataWorks workspace to which the rule template belongs.
//
// This operation removes a data quality rule template that is no longer needed. Make sure that the `Id` and `ProjectId` values are correct. Otherwise, the deletion may fail or cause unexpected data loss. Exercise caution when performing this operation and verify the template information before proceeding.
//
// @param request - DeleteDataQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityTemplateResponse
func (client *Client) DeleteDataQualityTemplateWithContext(ctx context.Context, request *DeleteDataQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data source by data source ID.
//
// Description:
//
// 1. This operation is available for all DataWorks editions.
//
// 2. To call this operation, you must have one of the following roles in DataWorks:
//
// - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
//
// @param request - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithContext(ctx context.Context, request *DeleteDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
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
		Action:      dara.String("DeleteDataSource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data source sharing rule by rule ID.
//
// Description:
//
// 1. This operation is available for all DataWorks editions.
//
// 2. To delete a sharing rule of a data source from Workspace A to Workspace B, you must have the data source sharing permissions in Workspace A or Workspace B. You must have one of the following roles in DataWorks:
//
// - Tenant Owner, Tenant Administrator, Workspace Administrator, and Workspace Owner
//
// @param request - DeleteDataSourceSharedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceSharedRuleResponse
func (client *Client) DeleteDataSourceSharedRuleWithContext(ctx context.Context, request *DeleteDataSourceSharedRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceSharedRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSourceSharedRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceSharedRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataset. Only DataWorks datasets are supported. The corresponding dataset versions are cascade deleted. The operator must be the creator of the dataset or an administrator of the workspace to which the dataset belongs.
//
// @param request - DeleteDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithContext(ctx context.Context, request *DeleteDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataset"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataset version. Only non-v1 DataWorks datasets are supported. To delete v1 datasets, use the DeleteDataset operation. Requires dataset creator or workspace administrator permissions.
//
// @param request - DeleteDatasetVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetVersionResponse
func (client *Client) DeleteDatasetVersionWithContext(ctx context.Context, request *DeleteDatasetVersionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetVersion"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file from DataStudio. If the file has been committed, an asynchronous process is triggered to delete the file in the scheduling system. The value of the DeploymentId parameter returned is used to call the GetDeployment operation to poll the status of the asynchronous process.
//
// @param request - DeleteFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithContext(ctx context.Context, request *DeleteFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke DeleteFolder to delete a folder on the Data Development page.
//
// @param request - DeleteFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithContext(ctx context.Context, request *DeleteFolderRequest, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFolder"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function (UDF) in DataStudio.
//
// Description:
//
//	Notice:
//
// After a UDF is published, it cannot be deleted. You must unpublish the UDF before you can delete it.
//
// @param request - DeleteFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunctionWithContext(ctx context.Context, request *DeleteFunctionRequest, runtime *dara.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified data lineage relationship from DataWorks Data Map.
//
// Description:
//
// 1. You must purchase DataWorks Professional Edition or a higher edition to use this feature.
//
// @param request - DeleteLineageRelationshipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLineageRelationshipResponse
func (client *Client) DeleteLineageRelationshipWithContext(ctx context.Context, request *DeleteLineageRelationshipRequest, runtime *dara.RuntimeOptions) (_result *DeleteLineageRelationshipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLineageRelationship"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLineageRelationshipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete MCP Server
//
// Description:
//
// ## Request Description
//
// This API allows you to delete the corresponding MCP Server instance based on the provided MCP Server name. Make sure you have the appropriate permissions and verify that the MCP Server name to be deleted is correct before calling.
//
// ### Notes
//
// - The deletion operation is irreversible. Proceed with caution.
//
// - Ensure that you have sufficient permissions (`dataworks:DeleteMcpServer`) to perform this operation.
//
// @param request - DeleteMcpServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpServerResponse
func (client *Client) DeleteMcpServerWithContext(ctx context.Context, request *DeleteMcpServerRequest, runtime *dara.RuntimeOptions) (_result *DeleteMcpServerResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcpServer"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcpServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a collection object in the specified Data Map, including Data Map categories and data albums.
//
// When deleting a data album, the caller must have the AliyunDataWorksFullAccess permission or be the creator or administrator of the album.
//
// Description:
//
// 1. You must purchase DataWorks Professional Edition or a higher edition to use this feature.
//
// @param request - DeleteMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetaCollectionResponse
func (client *Client) DeleteMetaCollectionWithContext(ctx context.Context, request *DeleteMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetaCollectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetaCollectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a metadata entity definition, including custom entity types and extension table types.
//
// Description:
//
// DataWorks Professional Edition or a more advanced edition is required.
//
// @param request - DeleteMetaEntityDefRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetaEntityDefResponse
func (client *Client) DeleteMetaEntityDefWithContext(ctx context.Context, request *DeleteMetaEntityDefRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetaEntityDefResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityType) {
		body["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetaEntityDef"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetaEntityDefResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates and deletes a network resource from a general-purpose resource group.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - DeleteNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkResponse
func (client *Client) DeleteNetworkWithContext(ctx context.Context, request *DeleteNetworkRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetwork"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a node from DataStudio.
//
// Description:
//
//	Notice:
//
// After a node is published, it cannot be deleted. You must unpublish the node before you can delete it.
//
// @param request - DeleteNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeResponse
func (client *Client) DeleteNodeWithContext(ctx context.Context, request *DeleteNodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified parameter.
//
// Description:
//
// This operation is available only in DataWorks professional edition and later versions.
//
// @param request - DeleteParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParameterResponse
func (client *Client) DeleteParameterWithContext(ctx context.Context, request *DeleteParameterRequest, runtime *dara.RuntimeOptions) (_result *DeleteParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteParameter"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteParameterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a process definition by the specified ID.
//
// Description:
//
// ## Description
//
// - This API deletes a process definition by its ID.
//
// - This operation is irreversible. Proceed with caution.
//
// - Before calling this API, back up relevant data or confirm that the process definition is no longer required.
//
// @param request - DeleteProcessDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProcessDefinitionResponse
func (client *Client) DeleteProcessDefinitionWithContext(ctx context.Context, request *DeleteProcessDefinitionRequest, runtime *dara.RuntimeOptions) (_result *DeleteProcessDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProcessDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProcessDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a DataWorks workspace.
//
// Description:
//
// To call this API, you must purchase DataWorks Basic Edition or a higher edition.
//
// Note: When you delete a workspace, the system moves it to the Recycle Bin. After a 14-day retention period, the system permanently purges the workspace. During this time, you cannot create a new workspace with the same name. You can find the deleted workspace in the Recycle Bin on the Workspace page in the console.
//
// @param request - DeleteProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithContext(ctx context.Context, request *DeleteProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a workspace member and the workspace roles granted to the member.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - DeleteProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectMemberResponse
func (client *Client) DeleteProjectMemberWithContext(ctx context.Context, request *DeleteProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProjectMember"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom role from a workspace.
//
// @param request - DeleteProjectRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectRoleResponse
func (client *Client) DeleteProjectRoleWithContext(ctx context.Context, request *DeleteProjectRoleRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProjectRole"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file resource from DataStudio.
//
// Description:
//
//	Notice:
//
// After a file resource is published, it cannot be deleted. You must unpublish the file resource before you can delete it.
//
// @param request - DeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithContext(ctx context.Context, request *DeleteResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource group.
//
// Description:
//
// 1. This operation requires DataWorks Basic Edition or a later version.
//
// 2. **Before you use this operation, ensure you understand the billing method and [pricing](https://help.aliyun.com/document_detail/2680173.html) for DataWorks resource groups.**
//
// 3. **Before you use this operation, ensure you have created the Service-Linked Role AliyunServiceRoleForDataWorks.**
//
// @param request - DeleteResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithContext(ctx context.Context, request *DeleteResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a route from a network resource.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteResponse
func (client *Client) DeleteRouteWithContext(ctx context.Context, request *DeleteRouteRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRoute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a security strategy by its ID. To call this operation, you must have DataWorks tenant administrator and security administrator permissions.
//
// Description:
//
// ## Usage notes
//
// - You can delete a security strategy by providing its ID.
//
// - You cannot delete a system strategy.
//
// @param request - DeleteSecurityStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityStrategyResponse
func (client *Client) DeleteSecurityStrategyWithContext(ctx context.Context, request *DeleteSecurityStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityStrategy"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a semantic job definition by the Name of a created job. If the job is currently running, call KillSemanticJob with the ExecutorJobId of that run and confirm the stop before deletion.
//
// Description:
//
// ## Scenarios
//
// Archives and deletes a saved semantic job definition so that it no longer appears in the list of available jobs.
//
// ## Call flow
//
// 1. Obtain the job name from `CreateSemanticJob.Data.Name` or `ListSemanticJobs.Data.SemanticJobs[].Name`.
//
// 2. To check whether any active runs exist, call `ListSemanticJobRuns` first. If necessary, stop the execution by calling `KillSemanticJob`.
//
// 3. Call this operation to delete the job definition.
//
// ## Result description
//
// A successful response indicates that the deletion request is complete. After deletion, you can no longer use the name to call `RunSemanticJob`.
//
// @param request - DeleteSemanticJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSemanticJobResponse
func (client *Client) DeleteSemanticJobWithContext(ctx context.Context, request *DeleteSemanticJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteSemanticJobResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSemanticJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSemanticJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Skill
//
// Description:
//
// ## Request Description
//
// This API is used to delete a Skill with the specified name from DataWorks. The exact name of the Skill to delete must be provided when invoking this API.
//
// ### Notes
//
// - Ensure that you have sufficient permissions to perform the delete operation.
//
// - The delete operation is irreversible. Use it with caution.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSkill"),
		Version:     dara.String("2024-05-18"),
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
// Deletes a task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskResponse
func (client *Client) DeleteTaskWithContext(ctx context.Context, request *DeleteTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectEnv) {
		query["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workflow.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowResponse
func (client *Client) DeleteWorkflowWithContext(ctx context.Context, request *DeleteWorkflowRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientUniqueCode) {
		body["ClientUniqueCode"] = request.ClientUniqueCode
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkflow"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified workflow in data development.
//
// Description:
//
//	Notice: After a workflow is published, it cannot be deleted. You must offline the workflow before deleting it.
//
// @param request - DeleteWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowDefinitionResponse
func (client *Client) DeleteWorkflowDefinitionWithContext(ctx context.Context, request *DeleteWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkflowDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkflowDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys a file to the production environment.
//
// @param request - DeployFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployFileResponse
func (client *Client) DeployFileWithContext(ctx context.Context, request *DeployFileRequest, runtime *dara.RuntimeOptions) (_result *DeployFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DetachDataQualityRulesFromEvaluationTask is deprecated, please use dataworks-public::2024-05-18::UpdateDataQualityScan instead.
//
// Summary:
//
// Removes the association between a data quality rule and a data quality monitoring task.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - DetachDataQualityRulesFromEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDataQualityRulesFromEvaluationTaskResponse
func (client *Client) DetachDataQualityRulesFromEvaluationTaskWithContext(ctx context.Context, tmpReq *DetachDataQualityRulesFromEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *DetachDataQualityRulesFromEvaluationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetachDataQualityRulesFromEvaluationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataQualityRuleIds) {
		request.DataQualityRuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataQualityRuleIds, dara.String("DataQualityRuleIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityEvaluationTaskId) {
		body["DataQualityEvaluationTaskId"] = request.DataQualityEvaluationTaskId
	}

	if !dara.IsNil(request.DataQualityRuleIdsShrink) {
		body["DataQualityRuleIds"] = request.DataQualityRuleIdsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachDataQualityRulesFromEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachDataQualityRulesFromEvaluationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the specified approval process definition.
//
// Description:
//
// ## Request
//
// - This API disables the specified approval process definition.
//
// - A disabled approval process definition remains inactive until it is re-enabled.
//
// - You must provide a valid process definition ID as a path parameter.
//
// @param request - DisableProcessDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableProcessDefinitionResponse
func (client *Client) DisableProcessDefinitionWithContext(ctx context.Context, request *DisableProcessDefinitionRequest, runtime *dara.RuntimeOptions) (_result *DisableProcessDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableProcessDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableProcessDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an image from a workspace.
//
// Description:
//
// 1. This operation requires DataWorks Basic Edition or a later version.
//
// 2. **Before calling this operation, ensure you have created the AliyunServiceRoleForDataWorks service-linked role.**
//
// @param request - DissociateProjectFromImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateProjectFromImageResponse
func (client *Client) DissociateProjectFromImageWithContext(ctx context.Context, request *DissociateProjectFromImageRequest, runtime *dara.RuntimeOptions) (_result *DissociateProjectFromImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateProjectFromImage"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateProjectFromImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a resource group from a workspace.
//
// Description:
//
// 1. This operation requires a subscription to DataWorks Basic Edition or a higher edition.
//
// 2. You must have one of the following roles in the DataWorks workspace:
//
// - tenant owner, workspace administrator, project owner, or operator
//
// @param request - DissociateProjectFromResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateProjectFromResourceGroupResponse
func (client *Client) DissociateProjectFromResourceGroupWithContext(ctx context.Context, request *DissociateProjectFromResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DissociateProjectFromResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateProjectFromResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateProjectFromResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the download URL of artifacts by job name and an optional run ID after a run completes. If JobRunId is not specified, the results of the most recent run of the job are returned.
//
// Description:
//
// ## Scenarios
//
// Retrieves temporary download URLs for result files of a submitted semantic job run, such as semantic model YAML artifacts. This operation returns download URLs and does not directly return file content.
//
// ## Procedure
//
// 1. Use the job name `JobName` to locate the job.
//
// 2. To retrieve artifacts of a specific run, specify the `JobRunId` from the `RunSemanticJob.Data.JobRunId` or `ListSemanticJobRuns` response. If you do not specify this parameter, the artifacts of the most recent run are returned.
//
// 3. Download the corresponding files from `Data.Results[].DownloadUrl`.
//
// ## Before you begin
//
// The download URL is a temporary credential. Use it only briefly on the client side. Do not write it to logs or store it for long-term use.
//
// @param request - DownloadSemanticResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadSemanticResultsResponse
func (client *Client) DownloadSemanticResultsWithContext(ctx context.Context, request *DownloadSemanticResultsRequest, runtime *dara.RuntimeOptions) (_result *DownloadSemanticResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobName) {
		body["JobName"] = request.JobName
	}

	if !dara.IsNil(request.JobRunId) {
		body["JobRunId"] = request.JobRunId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadSemanticResults"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadSemanticResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a process definition with the specified ID.
//
// Description:
//
// ## Request
//
// This API enables an existing process definition. You must provide the process definition ID as a path parameter.
//
// @param request - EnableProcessDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableProcessDefinitionResponse
func (client *Client) EnableProcessDefinitionWithContext(ctx context.Context, request *EnableProcessDefinitionRequest, runtime *dara.RuntimeOptions) (_result *EnableProcessDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableProcessDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableProcessDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a table to a workflow. The call to this API operation is equivalent to performing the following operations: Go to the DataStudio page, find the desired workflow, and then click the workflow name. Right-click Table under the desired folder and select Import Table.
//
// @param request - EstablishRelationTableToBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstablishRelationTableToBusinessResponse
func (client *Client) EstablishRelationTableToBusinessWithContext(ctx context.Context, request *EstablishRelationTableToBusinessRequest, runtime *dara.RuntimeOptions) (_result *EstablishRelationTableToBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.TableGuid) {
		body["TableGuid"] = request.TableGuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EstablishRelationTableToBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EstablishRelationTableToBusinessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a specified stage of a publish flow.
//
// Description:
//
//	Notice: The stages of a publish flow are sequential. For more information, see the response of GetPipelineRun. You cannot skip or repeat a stage.
//
//	Notice: The execution is asynchronous. The response only indicates that the stage is triggered, not that the stage is executed. Check the response of GetPipelineRun for the execution result.
//
//	Notice: This operation may not be available in earlier SDK versions. In this case, use the ExecDeploymentStage operation. The parameters are the same as those described in this document.
//
// @param request - ExecPipelineRunStageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecPipelineRunStageResponse
func (client *Client) ExecPipelineRunStageWithContext(ctx context.Context, request *ExecPipelineRunStageRequest, runtime *dara.RuntimeOptions) (_result *ExecPipelineRunStageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecPipelineRunStage"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecPipelineRunStageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a temporary workflow instance based on the specified configuration.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param tmpReq - ExecuteAdhocWorkflowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAdhocWorkflowInstanceResponse
func (client *Client) ExecuteAdhocWorkflowInstanceWithContext(ctx context.Context, tmpReq *ExecuteAdhocWorkflowInstanceRequest, runtime *dara.RuntimeOptions) (_result *ExecuteAdhocWorkflowInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExecuteAdhocWorkflowInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tasks) {
		request.TasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tasks, dara.String("Tasks"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		body["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TasksShrink) {
		body["Tasks"] = request.TasksShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAdhocWorkflowInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAdhocWorkflowInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Find the security policy that best matches the given conditions.
//
// Description:
//
// ## Request Description
//
// This API is used to find the most suitable security policy based on the provided control module, sub-module, and workspace ID. If a workspace ID is provided, the policy at the specified workspace level is matched first; otherwise, the tenant-level policy is returned. Note that system policies cannot be deleted or modified.
//
// @param request - FindBestMatchSecurityStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindBestMatchSecurityStrategyResponse
func (client *Client) FindBestMatchSecurityStrategyWithContext(ctx context.Context, request *FindBestMatchSecurityStrategyRequest, runtime *dara.RuntimeOptions) (_result *FindBestMatchSecurityStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ControlModule) {
		query["ControlModule"] = request.ControlModule
	}

	if !dara.IsNil(request.ControlSubModule) {
		query["ControlSubModule"] = request.ControlSubModule
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindBestMatchSecurityStrategy"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindBestMatchSecurityStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves agent details by name.
//
// Description:
//
// ## Request
//
// This API uses an agent\\"s name, provided as a parameter, to retrieve its detailed configuration, including the model configuration, system prompt, and tool list.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgent"),
		Version:     dara.String("2024-05-18"),
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
// Retrieves the metadata and content of an artifact based on a session ID and artifact path.
//
// Description:
//
// ## Description
//
// - This operation retrieves the metadata and content of a single artifact based on `SessionId` and `ArtifactPath`.
//
// - `SessionId` and `ArtifactPath` are required.
//
// @param tmpReq - GetAgentSessionArtifactMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentSessionArtifactMetaResponse
func (client *Client) GetAgentSessionArtifactMetaWithContext(ctx context.Context, tmpReq *GetAgentSessionArtifactMetaRequest, runtime *dara.RuntimeOptions) (_result *GetAgentSessionArtifactMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAgentSessionArtifactMetaShrinkRequest{}
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
		Action:      dara.String("GetAgentSessionArtifactMeta"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentSessionArtifactMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the cumulative AI token usage for a specified session.
//
// Description:
//
// ## Description
//
// - This operation retrieves usage statistics for AI tokens in a specified session. It provides a breakdown of tokens for prompts, completions, and thoughts, as well as the total token count and the number of cache-hit tokens.
//
// @param tmpReq - GetAgentSessionTokenUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentSessionTokenUsageResponse
func (client *Client) GetAgentSessionTokenUsageWithContext(ctx context.Context, tmpReq *GetAgentSessionTokenUsageRequest, runtime *dara.RuntimeOptions) (_result *GetAgentSessionTokenUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetAgentSessionTokenUsageShrinkRequest{}
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
		Action:      dara.String("GetAgentSessionTokenUsage"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentSessionTokenUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a custom monitoring alert rule.
//
// @param request - GetAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertRuleResponse
func (client *Client) GetAlertRuleWithContext(ctx context.Context, request *GetAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAlertRuleResponse, _err error) {
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
		Action:      dara.String("GetAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content of resource permission applications under the specified process instance ID.
//
// Description:
//
// ## Request Description
//
// - This API is used to query the details of resource permission applications based on the provided `ProcessInstanceId`.
//
// - A valid `ProcessInstanceId` parameter must be provided in the request.
//
// - The response includes the basic information, status, and the list of specific application contents.
//
// - Each application content includes detailed resource information, the grantee, the requested operation permissions, and more.
//
// @param request - GetApplicationContentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationContentsResponse
func (client *Client) GetApplicationContentsWithContext(ctx context.Context, request *GetApplicationContentsRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationContentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProcessInstanceId) {
		query["ProcessInstanceId"] = request.ProcessInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationContents"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationContentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a business process by calling GetBusiness.
//
// @param request - GetBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBusinessResponse
func (client *Client) GetBusinessWithContext(ctx context.Context, request *GetBusinessRequest, runtime *dara.RuntimeOptions) (_result *GetBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBusinessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified data catalog in DataWorks Data Map. Currently, DLF and StarRocks catalog types are supported.
//
// Description:
//
// 1. You must have DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - GetCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogResponse
func (client *Client) GetCatalogWithContext(ctx context.Context, request *GetCatalogRequest, runtime *dara.RuntimeOptions) (_result *GetCatalogResponse, _err error) {
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
		Action:      dara.String("GetCatalog"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can view authentication files.
//
// Description:
//
// 1. This feature is available only in DataWorks Basic Edition and later versions.
//
// 2. You must have at least one of the following roles in the DataWorks project: Tenant Owner, Space Administrator, Deployment, Developer, Project Owner, or O\\&M.
//
// @param request - GetCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertificateResponse
func (client *Client) GetCertificateWithContext(ctx context.Context, request *GetCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetCertificateResponse, _err error) {
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
		Action:      dara.String("GetCertificate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified column in a Data Map table.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - GetColumnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetColumnResponse
func (client *Client) GetColumnWithContext(ctx context.Context, request *GetColumnRequest, runtime *dara.RuntimeOptions) (_result *GetColumnResponse, _err error) {
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
		Action:      dara.String("GetColumn"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetColumnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets component information.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - GetComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComponentResponse
func (client *Client) GetComponentWithContext(ctx context.Context, request *GetComponentRequest, runtime *dara.RuntimeOptions) (_result *GetComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComponent"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specified computing resource based on the computing resource ID.
//
// Description:
//
// 1. DataWorks Basic Edition or a more advanced edition is required.
//
// 2. You must have at least one of the following roles in the DataWorks workspace:
//
// - Tenant Owner, Workspace Administrator, Deploy, Developer, Project Owner, O\\&M
//
// @param request - GetComputeResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComputeResourceResponse
func (client *Client) GetComputeResourceWithContext(ctx context.Context, request *GetComputeResourceRequest, runtime *dara.RuntimeOptions) (_result *GetComputeResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComputeResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComputeResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration, status, and latest run information of a specified metadata crawler.
//
// Description:
//
// ## Scenarios
//
// Queries the configuration, availability status, and latest run information of a specified metadata crawler.
//
// ## Recommended workflow
//
// 1. Call `ListCrawlers` to query crawler IDs.
//
// 2. Call this operation to retrieve crawler details.
//
// 3. To query the complete run history, call `ListCrawlerRuns`.
//
// ## Edition requirements
//
// DataWorks Basic Edition or higher is required.
//
// ## Precautions
//
// If the crawler has not been run, the latest run status and task instance ID may be empty.
//
// @param request - GetCrawlerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrawlerResponse
func (client *Client) GetCrawlerWithContext(ctx context.Context, request *GetCrawlerRequest, runtime *dara.RuntimeOptions) (_result *GetCrawlerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCrawler"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCrawlerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the crawler types supported in the current region, along with their collection scope, scheduling, and extension configuration capabilities.
//
// Description:
//
// ## Scenarios
//
// Queries the crawler types that can be created in the current region, as well as the data sources, collection scope, resource groups, scheduling, AI metadata description, and extension configuration capabilities supported by each type.
//
// ## Recommended workflow
//
// 1. Call this operation before creating or updating a crawler.
//
// 2. Construct a `CreateCrawler` or `UpdateCrawler` request based on the returned capability information.
//
// ## Edition requirements
//
// DataWorks Basic Edition or a higher edition is required.
//
// ## Precautions
//
// Capabilities may vary by region and crawler type. Use the actual response of this operation as the reference.
//
// @param request - GetCrawlerTypeCapabilitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCrawlerTypeCapabilitiesResponse
func (client *Client) GetCrawlerTypeCapabilitiesWithContext(ctx context.Context, request *GetCrawlerTypeCapabilitiesRequest, runtime *dara.RuntimeOptions) (_result *GetCrawlerTypeCapabilitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetCrawlerTypeCapabilities"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCrawlerTypeCapabilitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of asynchronously creating a workflow instance.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - GetCreateWorkflowInstancesResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateWorkflowInstancesResultResponse
func (client *Client) GetCreateWorkflowInstancesResultWithContext(ctx context.Context, request *GetCreateWorkflowInstancesResultRequest, runtime *dara.RuntimeOptions) (_result *GetCreateWorkflowInstancesResultResponse, _err error) {
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
		Action:      dara.String("GetCreateWorkflowInstancesResult"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCreateWorkflowInstancesResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a custom attribute definition.
//
// @param request - GetCustomAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomAttributeResponse
func (client *Client) GetCustomAttributeWithContext(ctx context.Context, request *GetCustomAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetCustomAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomAttribute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View data integration tasks.
//
// Description:
//
// This operation requires DataWorks Basic Edition or later.
//
// @param request - GetDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIJobResponse
func (client *Client) GetDIJobWithContext(ctx context.Context, request *GetDIJobRequest, runtime *dara.RuntimeOptions) (_result *GetDIJobResponse, _err error) {
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
		Action:      dara.String("GetDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDIJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the task logs of a data integration node.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param request - GetDIJobLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIJobLogResponse
func (client *Client) GetDIJobLogWithContext(ctx context.Context, request *GetDIJobLogRequest, runtime *dara.RuntimeOptions) (_result *GetDIJobLogResponse, _err error) {
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
		Action:      dara.String("GetDIJobLog"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDIJobLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data quality alert rule by rule ID.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - GetDataQualityAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityAlertRuleResponse
func (client *Client) GetDataQualityAlertRuleWithContext(ctx context.Context, request *GetDataQualityAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataQualityAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetDataQualityEvaluationTask is deprecated, please use dataworks-public::2024-05-18::CreateDataQualityScan instead.
//
// Summary:
//
// Query the details of a data quality validation task.
//
// Description:
//
// Available only with DataWorks Basic Edition or higher.
//
// @param request - GetDataQualityEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityEvaluationTaskResponse
func (client *Client) GetDataQualityEvaluationTaskWithContext(ctx context.Context, request *GetDataQualityEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityEvaluationTaskResponse, _err error) {
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
		Action:      dara.String("GetDataQualityEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityEvaluationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetDataQualityEvaluationTaskInstance is deprecated, please use dataworks-public::2024-05-18::GetDataQualityScanRun instead.
//
// Summary:
//
// Retrieves the details of a data quality check task instance.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required to use this operation.
//
// @param request - GetDataQualityEvaluationTaskInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityEvaluationTaskInstanceResponse
func (client *Client) GetDataQualityEvaluationTaskInstanceWithContext(ctx context.Context, request *GetDataQualityEvaluationTaskInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityEvaluationTaskInstanceResponse, _err error) {
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
		Action:      dara.String("GetDataQualityEvaluationTaskInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityEvaluationTaskInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetDataQualityRule is deprecated, please use dataworks-public::2024-05-18::GetDataQualityScan instead.
//
// Summary:
//
// Queries the details of a data quality rule.
//
// Description:
//
// You must purchase DataWorks Basic Edition or above to use this feature.
//
// @param request - GetDataQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityRuleResponse
func (client *Client) GetDataQualityRuleWithContext(ctx context.Context, request *GetDataQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityRuleResponse, _err error) {
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
		Action:      dara.String("GetDataQualityRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetDataQualityRuleTemplate is deprecated, please use dataworks-public::2024-05-18::GetDataQualityTemplate instead.
//
// Summary:
//
// Queries the details of a data quality rule template.
//
// Description:
//
// You can call this operation only if you have purchased DataWorks Basic Edition or a more advanced edition.
//
// @param request - GetDataQualityRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityRuleTemplateResponse
func (client *Client) GetDataQualityRuleTemplateWithContext(ctx context.Context, request *GetDataQualityRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityRuleTemplateResponse, _err error) {
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
		Action:      dara.String("GetDataQualityRuleTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityRuleTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets data quality monitoring details.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - GetDataQualityScanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityScanResponse
func (client *Client) GetDataQualityScanWithContext(ctx context.Context, request *GetDataQualityScanRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityScanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataQualityScan"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityScanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring run instance.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - GetDataQualityScanRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityScanRunResponse
func (client *Client) GetDataQualityScanRunWithContext(ctx context.Context, request *GetDataQualityScanRunRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityScanRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataQualityScanRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityScanRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log of a specific task instance that monitors data quality.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - GetDataQualityScanRunLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityScanRunLogResponse
func (client *Client) GetDataQualityScanRunLogWithContext(ctx context.Context, request *GetDataQualityScanRunLogRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityScanRunLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataQualityScanRunLog"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityScanRunLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data quality rule template by template ID.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - GetDataQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityTemplateResponse
func (client *Client) GetDataQualityTemplateWithContext(ctx context.Context, request *GetDataQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataQualityTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specified data source by data source ID.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// 2. You must have at least one of the following roles in the DataWorks workspace:
//
// - Tenant Owner, Workspace Administrator, Deployment, Developer, Project Owner, or O&M Engineer
//
// @param request - GetDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceResponse
func (client *Client) GetDataSourceWithContext(ctx context.Context, request *GetDataSourceRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceResponse, _err error) {
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
		Action:      dara.String("GetDataSource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a specific database in Data Map.
//
// Description:
//
// 1. DataWorks Basic Edition or a higher edition is required.
//
// @param request - GetDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseResponse
func (client *Client) GetDatabaseWithContext(ctx context.Context, request *GetDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseResponse, _err error) {
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
		Action:      dara.String("GetDatabase"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a dataset.
//
// @param request - GetDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
func (client *Client) GetDatasetWithContext(ctx context.Context, request *GetDatasetRequest, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataset"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets information for a given dataset version.
//
// @param request - GetDatasetVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetVersionResponse
func (client *Client) GetDatasetVersionWithContext(ctx context.Context, request *GetDatasetVersionRequest, runtime *dara.RuntimeOptions) (_result *GetDatasetVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasetVersion"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a deployment package.
//
// @param request - GetDeploymentPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentPackageResponse
func (client *Client) GetDeploymentPackageWithContext(ctx context.Context, request *GetDeploymentPackageRequest, runtime *dara.RuntimeOptions) (_result *GetDeploymentPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		body["DeploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeploymentPackage"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a file.
//
// @param request - GetFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileResponse
func (client *Client) GetFileWithContext(ctx context.Context, request *GetFileRequest, runtime *dara.RuntimeOptions) (_result *GetFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the version details of a file.
//
// @param request - GetFileVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileVersionResponse
func (client *Client) GetFileVersionWithContext(ctx context.Context, request *GetFileVersionRequest, runtime *dara.RuntimeOptions) (_result *GetFileVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileVersion) {
		body["FileVersion"] = request.FileVersion
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileVersion"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a folder.
//
// @param request - GetFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFolderResponse
func (client *Client) GetFolderWithContext(ctx context.Context, request *GetFolderRequest, runtime *dara.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.FolderPath) {
		body["FolderPath"] = request.FolderPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFolder"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a user-defined function (UDF) in DataStudio.
//
// @param request - GetFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionResponse
func (client *Client) GetFunctionWithContext(ctx context.Context, request *GetFunctionRequest, runtime *dara.RuntimeOptions) (_result *GetFunctionResponse, _err error) {
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
		Action:      dara.String("GetFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data snapshot of an extension point based on the ID of a message in DataWorks OpenEvent when the related extension point event is triggered.
//
// @param request - GetIDEEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIDEEventDetailResponse
func (client *Client) GetIDEEventDetailWithContext(ctx context.Context, request *GetIDEEventDetailRequest, runtime *dara.RuntimeOptions) (_result *GetIDEEventDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIDEEventDetail"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIDEEventDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified image by image ID.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// 2. **Before calling this operation, make sure that the service-linked role AliyunServiceRoleForDataWorks has been created.**
//
// @param request - GetImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageResponse
func (client *Client) GetImageWithContext(ctx context.Context, request *GetImageRequest, runtime *dara.RuntimeOptions) (_result *GetImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ImageVersion) {
		query["ImageVersion"] = request.ImageVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImage"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the status of an asynchronous task. After calling an asynchronous API, poll this API to obtain the success status.
//
// @param request - GetJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobStatusResponse
func (client *Client) GetJobStatusWithContext(ctx context.Context, request *GetJobStatusRequest, runtime *dara.RuntimeOptions) (_result *GetJobStatusResponse, _err error) {
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
		Action:      dara.String("GetJobStatus"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified lineage relationship in DataWorks Data Map.
//
// Description:
//
// 1. You must have DataWorks Standard Edition or a higher edition to use this operation.
//
// @param request - GetLineageRelationshipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLineageRelationshipResponse
func (client *Client) GetLineageRelationshipWithContext(ctx context.Context, request *GetLineageRelationshipRequest, runtime *dara.RuntimeOptions) (_result *GetLineageRelationshipResponse, _err error) {
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
		Action:      dara.String("GetLineageRelationship"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLineageRelationshipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the details of an MCP Server.
//
// Description:
//
// ## Description
//
// This API returns the detailed configuration of a specific MCP Server by name. The response includes the creator ID, modifier ID, service address, and transport protocol. You must provide the exact name of the MCP Server in the request.
//
// ### Notes
//
// - Ensure you have the required permissions to call this API.
//
// - The MCP Server name is case-sensitive.
//
// @param request - GetMcpServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpServerResponse
func (client *Client) GetMcpServerWithContext(ctx context.Context, request *GetMcpServerRequest, runtime *dara.RuntimeOptions) (_result *GetMcpServerResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpServer"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a collection in Data Map. Collections include categories and data albums.
//
// Description:
//
// 1. DataWorks Professional Edition or a higher edition is required.
//
// @param request - GetMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaCollectionResponse
func (client *Client) GetMetaCollectionWithContext(ctx context.Context, request *GetMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *GetMetaCollectionResponse, _err error) {
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
		Action:      dara.String("GetMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaCollectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a custom entity.
//
// @param request - GetMetaEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaEntityResponse
func (client *Client) GetMetaEntityWithContext(ctx context.Context, request *GetMetaEntityRequest, runtime *dara.RuntimeOptions) (_result *GetMetaEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaEntity"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a custom entity definition.
//
// @param request - GetMetaEntityDefRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaEntityDefResponse
func (client *Client) GetMetaEntityDefWithContext(ctx context.Context, request *GetMetaEntityDefRequest, runtime *dara.RuntimeOptions) (_result *GetMetaEntityDefResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityType) {
		body["EntityType"] = request.EntityType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaEntityDef"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaEntityDefResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a network resource.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkResponse
func (client *Client) GetNetworkWithContext(ctx context.Context, request *GetNetworkRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkResponse, _err error) {
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
		Action:      dara.String("GetNetwork"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a node in DataStudio.
//
// @param request - GetNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeResponse
func (client *Client) GetNodeWithContext(ctx context.Context, request *GetNodeRequest, runtime *dara.RuntimeOptions) (_result *GetNodeResponse, _err error) {
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
		Action:      dara.String("GetNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a parameter by its ID.
//
// Description:
//
// This operation is available only in DataWorks Professional Edition or later.
//
// @param request - GetParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParameterResponse
func (client *Client) GetParameterWithContext(ctx context.Context, request *GetParameterRequest, runtime *dara.RuntimeOptions) (_result *GetParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParameter"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParameterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves partition details for a data map table. Currently supports MaxCompute and HMS (EMR cluster) types only.
//
// Description:
//
// 1. DataWorks Basic Edition or a higher edition is required.
//
// 2. Only MaxCompute and HMS (EMR cluster) table types are supported.
//
// @param request - GetPartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPartitionResponse
func (client *Client) GetPartitionWithContext(ctx context.Context, request *GetPartitionRequest, runtime *dara.RuntimeOptions) (_result *GetPartitionResponse, _err error) {
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
		Action:      dara.String("GetPartition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPartitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a deployment process.
//
// Description:
//
//	Notice: This operation may not be available in earlier SDK versions. In this case, use the GetDeployment operation instead. The parameters are the same as those described in this document.
//
// @param request - GetPipelineRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineRunResponse
func (client *Client) GetPipelineRunWithContext(ctx context.Context, request *GetPipelineRunRequest, runtime *dara.RuntimeOptions) (_result *GetPipelineRunResponse, _err error) {
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
		Action:      dara.String("GetPipelineRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineRunResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the approval policy for a given process definition ID.
//
// Description:
//
// ## Request
//
// - This API retrieves the details of a specific approval process definition using the `ID` parameter.
//
// - The `ID` parameter is required and must be a valid process definition ID.
//
// - The response includes the basic properties of the approval process definition, rule conditions, notification service configurations, and approval nodes.
//
// - A successful request returns the complete process definition object. A failed request returns an error code and message for troubleshooting.
//
// @param request - GetProcessDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProcessDefinitionResponse
func (client *Client) GetProcessDefinitionWithContext(ctx context.Context, request *GetProcessDefinitionRequest, runtime *dara.RuntimeOptions) (_result *GetProcessDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProcessDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProcessDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves details for a specified approval process instance, including its approval nodes and task list.
//
// Description:
//
// ## Request
//
// This API is used to monitor and manage the status of an approval process. By providing the approval process instance ID, you can query for related information, such as the approval process definition, current approval nodes, and the tasks on each node.
//
// @param request - GetProcessInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProcessInstanceResponse
func (client *Client) GetProcessInstanceWithContext(ctx context.Context, request *GetProcessInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetProcessInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProcessInstanceId) {
		query["ProcessInstanceId"] = request.ProcessInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProcessInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProcessInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a DataWorks workspace.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithContext(ctx context.Context, request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
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
		Action:      dara.String("GetProject"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific member in a Workspace.
//
// Description:
//
// This operation is available only in DataWorks Basic Edition and later.
//
// @param request - GetProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectMemberResponse
func (client *Client) GetProjectMemberWithContext(ctx context.Context, request *GetProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *GetProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectMember"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a workspace role.
//
// Description:
//
// You can call this operation only if you have purchased DataWorks Basic Edition or a later edition.
//
// @param request - GetProjectRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectRoleResponse
func (client *Client) GetProjectRoleWithContext(ctx context.Context, request *GetProjectRoleRequest, runtime *dara.RuntimeOptions) (_result *GetProjectRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectRole"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the result of an asynchronous rerun workflow instance
//
// @param request - GetRerunWorkflowInstancesResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRerunWorkflowInstancesResultResponse
func (client *Client) GetRerunWorkflowInstancesResultWithContext(ctx context.Context, request *GetRerunWorkflowInstancesResultRequest, runtime *dara.RuntimeOptions) (_result *GetRerunWorkflowInstancesResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRerunWorkflowInstancesResult"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRerunWorkflowInstancesResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a file resource.
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithContext(ctx context.Context, request *GetResourceRequest, runtime *dara.RuntimeOptions) (_result *GetResourceResponse, _err error) {
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
		Action:      dara.String("GetResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the information of a specified resource group by ID.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - GetResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroupWithContext(ctx context.Context, request *GetResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
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
		Action:      dara.String("GetResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a route based on its ID.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRouteResponse
func (client *Client) GetRouteWithContext(ctx context.Context, request *GetRouteRequest, runtime *dara.RuntimeOptions) (_result *GetRouteResponse, _err error) {
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
		Action:      dara.String("GetRoute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the schema details of a specified table in Data Map. Only MaxCompute and Hologres schemas are supported.
//
// Description:
//
// 1. DataWorks Basic Edition or a higher edition is required.
//
// 2. Only MaxCompute and Hologres types are supported.
//
// @param request - GetSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSchemaResponse
func (client *Client) GetSchemaWithContext(ctx context.Context, request *GetSchemaRequest, runtime *dara.RuntimeOptions) (_result *GetSchemaResponse, _err error) {
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
		Action:      dara.String("GetSchema"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSchemaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a security policy by its ID. This operation requires both DataWorks tenant administrator and security administrator permissions.
//
// Description:
//
// ## Request
//
// - This API retrieves the complete configuration information for a security policy by its ID.
//
// - The API returns an error message if the provided `Id` is invalid or does not exist.
//
// - The response includes basic policy information, such as its name and description, and policy details, such as control items and their settings.
//
// - Note: Some fields in a system default policy cannot be modified or deleted.
//
// @param request - GetSecurityStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityStrategyResponse
func (client *Client) GetSecurityStrategyWithContext(ctx context.Context, request *GetSecurityStrategyRequest, runtime *dara.RuntimeOptions) (_result *GetSecurityStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecurityStrategy"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecurityStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the executor status and runtime configuration by using the ExecutorJobId returned by RunSemanticJob or ListSemanticJobRuns.
//
// Description:
//
// ## Scenarios
//
// Queries the detailed status and runtime information of a semantic job run on the executor side. This is used to poll execution progress or troubleshoot run failures.
//
// ## Procedure
//
// 1. Call `RunSemanticJob` or `ListSemanticJobRuns` to obtain the `ExecutorJobId`.
//
// 2. Use the `ProjectId` returned by the job definition as the `ProjectId` for this operation.
//
// 3. Determine the current status based on the executor details in `Data`. If the job is still running, continue polling this operation.
//
// ## Related operations
//
// To retrieve logs, call `GetSemanticJobLog`. To stop a run, call `KillSemanticJob`.
//
// @param request - GetSemanticJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSemanticJobDetailResponse
func (client *Client) GetSemanticJobDetailWithContext(ctx context.Context, request *GetSemanticJobDetailRequest, runtime *dara.RuntimeOptions) (_result *GetSemanticJobDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutorJobId) {
		query["ExecutorJobId"] = request.ExecutorJobId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSemanticJobDetail"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSemanticJobDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries execution logs by using the ExecutorJobId returned by RunSemanticJob or ListSemanticJobRuns. The current POP contract exposes only the task and workspace identifiers and returns default log segments.
//
// Description:
//
// ## Scenarios
//
// Reads the execution logs of a semantic job run to observe the execution process and identify failure causes.
//
// ## Procedure
//
// 1. Specify the run by using `RunSemanticJob.Data.ExecutorJobId` or `ListSemanticJobRuns[].ExecutorJobId`.
//
// 2. Call this operation with the `ProjectId` of the corresponding task.
//
// 3. Analyze the log segments in `Data` together with the run status returned by `GetSemanticJobDetail`.
//
// ## Before you begin
//
// Logs are used for diagnostics and do not represent the final result files. Obtain result artifacts by calling `DownloadSemanticResults`.
//
// @param request - GetSemanticJobLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSemanticJobLogResponse
func (client *Client) GetSemanticJobLogWithContext(ctx context.Context, request *GetSemanticJobLogRequest, runtime *dara.RuntimeOptions) (_result *GetSemanticJobLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecutorJobId) {
		query["ExecutorJobId"] = request.ExecutorJobId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSemanticJobLog"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSemanticJobLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified Skill by name, including the body of the SKILL.md file and the bundle\\"s download link.
//
// Description:
//
// ## Overview
//
// - **request parameters**: The name of the target Skill.
//
// - **response parameters**: The details of the Skill, including its name, description, creator ID, modifier ID, visibility level, visibility scope, the body of the SKILL.md file, a temporary download link for bundle.zip (which requires no authentication and will expire), the creation time, and the last modified time.
//
// - **Note**: The `BundleUrl` is a temporary download link. Once the link expires, you must call this operation again to get a new one.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSkill"),
		Version:     dara.String("2024-05-18"),
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
// Retrieves the details of a version snapshot.
//
// @param request - GetSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSnapshotResponse
func (client *Client) GetSnapshotWithContext(ctx context.Context, request *GetSnapshotRequest, runtime *dara.RuntimeOptions) (_result *GetSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSnapshot"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified table in Data Map. You can choose whether to return business metadata.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - GetTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableResponse
func (client *Client) GetTableWithContext(ctx context.Context, request *GetTableRequest, runtime *dara.RuntimeOptions) (_result *GetTableResponse, _err error) {
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
		Action:      dara.String("GetTable"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a task.
//
// @param request - GetTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithContext(ctx context.Context, request *GetTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
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
		Action:      dara.String("GetTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an instance.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - GetTaskInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskInstanceResponse
func (client *Client) GetTaskInstanceWithContext(ctx context.Context, request *GetTaskInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetTaskInstanceResponse, _err error) {
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
		Action:      dara.String("GetTaskInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the run log generated during a specific run of an instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetTaskInstanceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskInstanceLogResponse
func (client *Client) GetTaskInstanceLogWithContext(ctx context.Context, request *GetTaskInstanceLogRequest, runtime *dara.RuntimeOptions) (_result *GetTaskInstanceLogResponse, _err error) {
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
		Action:      dara.String("GetTaskInstanceLog"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskInstanceLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowResponse
func (client *Client) GetWorkflowWithContext(ctx context.Context, request *GetWorkflowRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowResponse, _err error) {
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
		Action:      dara.String("GetWorkflow"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow.
//
// @param request - GetWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowDefinitionResponse
func (client *Client) GetWorkflowDefinitionWithContext(ctx context.Context, request *GetWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowDefinitionResponse, _err error) {
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
		Action:      dara.String("GetWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow instance.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - GetWorkflowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowInstanceResponse
func (client *Client) GetWorkflowInstanceWithContext(ctx context.Context, request *GetWorkflowInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowInstanceResponse, _err error) {
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
		Action:      dara.String("GetWorkflowInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns roles to members in a workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - GrantMemberProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantMemberProjectRolesResponse
func (client *Client) GrantMemberProjectRolesWithContext(ctx context.Context, tmpReq *GrantMemberProjectRolesRequest, runtime *dara.RuntimeOptions) (_result *GrantMemberProjectRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GrantMemberProjectRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("RoleCodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantMemberProjectRoles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantMemberProjectRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a certificate file.
//
// Description:
//
// 1. This feature requires DataWorks Basic Edition or a later version.
//
// 2. You must be assigned one of the following roles in the DataWorks project: Tenant Owner, Space Administrator, Project Owner, or O\\&M.
//
// @param request - ImportCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportCertificateResponse
func (client *Client) ImportCertificateWithContext(ctx context.Context, request *ImportCertificateRequest, runtime *dara.RuntimeOptions) (_result *ImportCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateFile) {
		query["CertificateFile"] = request.CertificateFile
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportCertificate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportCertificateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a workflow node defined by FlowSpec and its internal child nodes into DataStudio.
//
// Description:
//
//	Notice:
//
// - This operation does not support importing multiple workflows. If more than one workflow is defined in the FlowSpec, all workflows after the first one are ignored.
//
// - This is an asynchronous operation. Calling this operation returns an asynchronous task object. To query the execution status of the task, call GetJobStatus.
//
// @param request - ImportWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportWorkflowDefinitionResponse
func (client *Client) ImportWorkflowDefinitionWithContext(ctx context.Context, request *ImportWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *ImportWorkflowDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportWorkflowDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a specified run by using the ExecutorJobId returned by RunSemanticJob or ListSemanticJobRuns. A successful call indicates only that the stop request has been accepted. Query the final status by calling GetSemanticJobDetail.
//
// Description:
//
// ## Scenarios
//
// Sends a stop request to the executor for a specified semantic job run. This is applicable to scenarios where a job runs for an extended period, requires manual termination, or needs resource reclamation.
//
// ## Procedure
//
// 1. Obtain the `ExecutorJobId` from `RunSemanticJob` or `ListSemanticJobRuns`, and use the `ProjectId` of the job.
//
// 2. Optionally specify `RetryTimes`.
//
// 3. After the call, poll the final status by calling `GetSemanticJobDetail`. If necessary, call `GetSemanticJobLog` for diagnostics.
//
// ## Precautions
//
// A successful response indicates only that the stop request has been processed. It does not mean that the job has reached a desired state.
//
// @param request - KillSemanticJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillSemanticJobResponse
func (client *Client) KillSemanticJobWithContext(ctx context.Context, request *KillSemanticJobRequest, runtime *dara.RuntimeOptions) (_result *KillSemanticJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExecutorJobId) {
		body["ExecutorJobId"] = request.ExecutorJobId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RetryTimes) {
		body["RetryTimes"] = request.RetryTimes
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KillSemanticJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KillSemanticJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of artifact files produced by a specified session.
//
// Description:
//
// ## Operation description
//
// - This operation queries all artifact files generated in a specific session. You can use the `Params.RequestId` parameter to filter files produced by a single request.
//
// - The `NextToken` parameter is used to retrieve more results in a paginated manner. You do not need to provide this value for the first call.
//
// - By default, a maximum of 50 records are returned per page. You can adjust this value by using the `MaxResults` parameter.
//
// @param tmpReq - ListAgentSessionArtifactsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentSessionArtifactsResponse
func (client *Client) ListAgentSessionArtifactsWithContext(ctx context.Context, tmpReq *ListAgentSessionArtifactsRequest, runtime *dara.RuntimeOptions) (_result *ListAgentSessionArtifactsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAgentSessionArtifactsShrinkRequest{}
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
		Action:      dara.String("ListAgentSessionArtifacts"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentSessionArtifactsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the conversation history for the agent session.
//
// Description:
//
// ## Request
//
// - Specify at least one of `agentName` or `sessionSourceList`.
//
// - You can use the `tagList`, `sessionId`, and `sessionTitle` parameters for combined filtering.
//
// - The response follows the Alibaba Cloud OpenAPI pagination specification and includes the `totalCount`, `maxResults`, `nextToken`, and `sessionList` fields.
//
// - If you provide an invalid string for `nextToken`, its value defaults to `1`.
//
// - By default, this operation returns 50 records per page. You can use the `maxResults` parameter to adjust this number.
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
		Version:     dara.String("2024-05-18"),
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
// Queries the list of agents available to the current tenant, with support for filtering by name and pagination.
//
// Description:
//
// ## Request description
//
// - This operation queries all available agents under the current tenant.
//
// - Supports exact match filtering by using the `agentName` parameter.
//
// @param tmpReq - ListAgentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentsResponse
func (client *Client) ListAgentsWithContext(ctx context.Context, tmpReq *ListAgentsRequest, runtime *dara.RuntimeOptions) (_result *ListAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAgentsShrinkRequest{}
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
		Action:      dara.String("ListAgents"),
		Version:     dara.String("2024-05-18"),
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
// Queries a list of custom monitoring alert rules.
//
// @param tmpReq - ListAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertRulesResponse
func (client *Client) ListAlertRulesWithContext(ctx context.Context, tmpReq *ListAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAlertRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAlertRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Receiver) {
		query["Receiver"] = request.Receiver
	}

	if !dara.IsNil(request.TaskIdsShrink) {
		query["TaskIds"] = request.TaskIdsShrink
	}

	if !dara.IsNil(request.TypesShrink) {
		query["Types"] = request.TypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workflows.
//
// @param request - ListBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBusinessResponse
func (client *Client) ListBusinessWithContext(ctx context.Context, request *ListBusinessRequest, runtime *dara.RuntimeOptions) (_result *ListBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBusinessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of catalogs in Data Map. Only catalogs of the Data Lake Formation (DLF) and StarRocks metadata crawler types are supported. For the DLF metadata crawler type, all supported data catalogs are returned. For the StarRocks metadata crawler type, data catalogs in a specific instance are returned.
//
// Description:
//
// 1. DataWorks Basic Edition or a higher edition is required.
//
// @param tmpReq - ListCatalogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCatalogsResponse
func (client *Client) ListCatalogsWithContext(ctx context.Context, tmpReq *ListCatalogsRequest, runtime *dara.RuntimeOptions) (_result *ListCatalogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCatalogsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCatalogs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCatalogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of certificate files.
//
// Description:
//
// 1. This API operation is available for all DataWorks editions.
//
// 2. You can call this operation only if you are assigned one of the following roles in DataWorks: Tenant Owner, Workspace Administrator, Deploy, Develop, Visitor, Workspace Owner, O\\&M, Model Designer, Security Administrator, Data Analyst, OpenPlatform Administrator, and Data Governance Administrator.
//
// @param request - ListCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertificatesResponse
func (client *Client) ListCertificatesWithContext(ctx context.Context, request *ListCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListCertificatesResponse, _err error) {
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
		Action:      dara.String("ListCertificates"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCertificatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the column list of a specified table in DataWorks Data Map.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - ListColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListColumnsResponse
func (client *Client) ListColumnsWithContext(ctx context.Context, request *ListColumnsRequest, runtime *dara.RuntimeOptions) (_result *ListColumnsResponse, _err error) {
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
		Action:      dara.String("ListColumns"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListColumnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of components.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - ListComponentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentsResponse
func (client *Client) ListComponentsWithContext(ctx context.Context, request *ListComponentsRequest, runtime *dara.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponents"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of computing resources that meet the specified business information.
//
// Description:
//
// 1. DataWorks Basic Edition or a more advanced edition is required.
//
// 2. You must have at least one of the following roles in the DataWorks workspace:
//
// 3. Tenant Owner, Workspace Administrator, Deploy, Developer, Visitor, Project Owner, O\\&M, Model Designer, Security Administrator, Data Analyst, Development Platform Administrator, Data Governance Administrator
//
// @param tmpReq - ListComputeResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComputeResourcesResponse
func (client *Client) ListComputeResourcesWithContext(ctx context.Context, tmpReq *ListComputeResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListComputeResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListComputeResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TypesShrink) {
		query["Types"] = request.TypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComputeResources"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComputeResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the run records of a specified metadata crawler by paging, with optional filtering by time range and run status.
//
// Description:
//
// ## Scenarios
//
// Queries the run records of a specified metadata crawler within the last 30 days by paging, with optional filtering by run start time and status.
//
// ## Recommended workflow
//
// 1. Invoke `ListCrawlers` to obtain the crawler ID.
//
// 2. Invoke this operation to query run records and node instance IDs.
//
// 3. For asynchronous operations such as running or stopping, use the final status returned by this operation as the source of truth.
//
// ## Edition requirements
//
// DataWorks Basic Edition or a higher edition is required.
//
// ## Precautions
//
// If no time range is specified, the system queries records from the last 30 days by default.
//
// @param request - ListCrawlerRunsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrawlerRunsResponse
func (client *Client) ListCrawlerRunsWithContext(ctx context.Context, request *ListCrawlerRunsRequest, runtime *dara.RuntimeOptions) (_result *ListCrawlerRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimeFrom) {
		body["StartTimeFrom"] = request.StartTimeFrom
	}

	if !dara.IsNil(request.StartTimeTo) {
		body["StartTimeTo"] = request.StartTimeTo
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCrawlerRuns"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCrawlerRunsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metadata crawlers by paged query, and supports filtering by workspace, data source, type, environment, owner, and name.
//
// Description:
//
// ## Scenarios
//
// Performs a paged query of metadata crawlers that you have access to. Supports filtering by workspace, data source, crawler type, environment, owner, and name.
//
// ## Recommended flow
//
// 1. Combine filter conditions as needed to perform a conditional query of the crawler list.
//
// 2. Use the returned crawler IDs to invoke the get details, update, run, stop, run records, or delete operations.
//
// ## Version requirements
//
// DataWorks Basic Edition or higher is required.
//
// ## Precautions
//
// When multiple filter conditions are provided at the same time, they take effect in combination. The name field supports fuzzy match.
//
// @param tmpReq - ListCrawlersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrawlersResponse
func (client *Client) ListCrawlersWithContext(ctx context.Context, tmpReq *ListCrawlersRequest, runtime *dara.RuntimeOptions) (_result *ListCrawlersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCrawlersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSourceIds) {
		request.DataSourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSourceIds, dara.String("DataSourceIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceIdsShrink) {
		body["DataSourceIds"] = request.DataSourceIdsShrink
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCrawlers"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCrawlersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of custom agents.
//
// Description:
//
// ## Request
//
// - **Search keyword**: Use the `Q` parameter to perform a fuzzy search by agent name.
//
// - **Visibility level filtering**: Use the `Visibility` parameter to filter results by visibility level, such as `TENANT`, `PROJECT`, or `USER`.
//
// - **Paging information**: Use the `MaxResults` and `NextToken` parameters to implement paginated queries. `NextToken` retrieves the next page of results.
//
// @param tmpReq - ListCustomAgentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomAgentsResponse
func (client *Client) ListCustomAgentsWithContext(ctx context.Context, tmpReq *ListCustomAgentsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCustomAgentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Visibility) {
		request.VisibilityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Visibility, dara.String("Visibility"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Q) {
		body["Q"] = request.Q
	}

	if !dara.IsNil(request.VisibilityShrink) {
		body["Visibility"] = request.VisibilityShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomAgents"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomAgentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of custom attribute definitions.
//
// @param request - ListCustomAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomAttributesResponse
func (client *Client) ListCustomAttributesWithContext(ctx context.Context, request *ListCustomAttributesRequest, runtime *dara.RuntimeOptions) (_result *ListCustomAttributesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EntityTypes) {
		query["EntityTypes"] = request.EntityTypes
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomAttributes"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomAttributesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Views alert rules configured for a synchronization task.
//
// @param request - ListDIAlarmRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIAlarmRulesResponse
func (client *Client) ListDIAlarmRulesWithContext(ctx context.Context, request *ListDIAlarmRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDIAlarmRulesResponse, _err error) {
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
		Action:      dara.String("ListDIAlarmRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIAlarmRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries events for a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDIJobEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobEventsResponse
func (client *Client) ListDIJobEventsWithContext(ctx context.Context, request *ListDIJobEventsRequest, runtime *dara.RuntimeOptions) (_result *ListDIJobEventsResponse, _err error) {
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
		Action:      dara.String("ListDIJobEvents"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIJobEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metrics for a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ListDIJobMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobMetricsResponse
func (client *Client) ListDIJobMetricsWithContext(ctx context.Context, tmpReq *ListDIJobMetricsRequest, runtime *dara.RuntimeOptions) (_result *ListDIJobMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDIJobMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MetricName) {
		request.MetricNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricName, dara.String("MetricName"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDIJobMetrics"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIJobMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the running information about a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDIJobRunDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobRunDetailsResponse
func (client *Client) ListDIJobRunDetailsWithContext(ctx context.Context, request *ListDIJobRunDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListDIJobRunDetailsResponse, _err error) {
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
		Action:      dara.String("ListDIJobRunDetails"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIJobRunDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists Data Integration jobs.
//
// Description:
//
// This operation requires DataWorks Basic Edition or a later edition.
//
// @param request - ListDIJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobsResponse
func (client *Client) ListDIJobsWithContext(ctx context.Context, request *ListDIJobsRequest, runtime *dara.RuntimeOptions) (_result *ListDIJobsResponse, _err error) {
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
		Action:      dara.String("ListDIJobs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tags.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param request - ListDataAssetTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataAssetTagsResponse
func (client *Client) ListDataAssetTagsWithContext(ctx context.Context, request *ListDataAssetTagsRequest, runtime *dara.RuntimeOptions) (_result *ListDataAssetTagsResponse, _err error) {
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
		Action:      dara.String("ListDataAssetTags"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataAssetTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries DataWorks data assets that are associated with tags by paging.
//
// Description:
//
// You must purchase DataWorks Enterprise Edition or a higher edition to use this feature.
//
// @param tmpReq - ListDataAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataAssetsResponse
func (client *Client) ListDataAssetsWithContext(ctx context.Context, tmpReq *ListDataAssetsRequest, runtime *dara.RuntimeOptions) (_result *ListDataAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataAssetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataAssetIds) {
		request.DataAssetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataAssetIds, dara.String("DataAssetIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataAssets"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data quality alert rules in a project.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - ListDataQualityAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityAlertRulesResponse
func (client *Client) ListDataQualityAlertRulesWithContext(ctx context.Context, request *ListDataQualityAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityAlertRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityScanId) {
		query["DataQualityScanId"] = request.DataQualityScanId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataQualityAlertRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityAlertRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListDataQualityEvaluationTaskInstances is deprecated, please use dataworks-public::2024-05-18::ListDataQualityScanRuns instead.
//
// Summary:
//
// Performs a paginated query of the quality monitoring task instance list.
//
// Description:
//
// You must purchase DataWorks Basic Edition or above to use this feature.
//
// @param request - ListDataQualityEvaluationTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityEvaluationTaskInstancesResponse
func (client *Client) ListDataQualityEvaluationTaskInstancesWithContext(ctx context.Context, request *ListDataQualityEvaluationTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityEvaluationTaskInstancesResponse, _err error) {
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
		Action:      dara.String("ListDataQualityEvaluationTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityEvaluationTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListDataQualityEvaluationTasks is deprecated, please use dataworks-public::2024-05-18::ListDataQualityScans instead.
//
// Summary:
//
// Lists quality monitoring nodes by paging query.
//
// Description:
//
// 需要购买DataWorks基础版及以上版本才能使用
//
// @param request - ListDataQualityEvaluationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityEvaluationTasksResponse
func (client *Client) ListDataQualityEvaluationTasksWithContext(ctx context.Context, request *ListDataQualityEvaluationTasksRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityEvaluationTasksResponse, _err error) {
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
		Action:      dara.String("ListDataQualityEvaluationTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityEvaluationTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListDataQualityResults is deprecated, please use dataworks-public::2024-05-18::ListDataQualityScanRuns instead.
//
// Summary:
//
// Queries a list of data quality results by using paging.
//
// Description:
//
// 需要购买DataWorks基础版及以上版本才能使用
//
// @param request - ListDataQualityResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityResultsResponse
func (client *Client) ListDataQualityResultsWithContext(ctx context.Context, request *ListDataQualityResultsRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityResultsResponse, _err error) {
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
		Action:      dara.String("ListDataQualityResults"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListDataQualityRuleTemplates is deprecated, please use dataworks-public::2024-05-18::ListDataQualityTemplates instead.
//
// Summary:
//
// Queries a list of data quality rule templates.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - ListDataQualityRuleTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityRuleTemplatesResponse
func (client *Client) ListDataQualityRuleTemplatesWithContext(ctx context.Context, request *ListDataQualityRuleTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityRuleTemplatesResponse, _err error) {
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
		Action:      dara.String("ListDataQualityRuleTemplates"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityRuleTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListDataQualityRules is deprecated, please use dataworks-public::2024-05-18::ListDataQualityScans instead.
//
// Summary:
//
// Queries quality monitoring rules by paging.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - ListDataQualityRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityRulesResponse
func (client *Client) ListDataQualityRulesWithContext(ctx context.Context, request *ListDataQualityRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityRulesResponse, _err error) {
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
		Action:      dara.String("ListDataQualityRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the run history of data quality scans in a specified project.
//
// Description:
//
// This feature requires DataWorks basic edition or higher.
//
// @param tmpReq - ListDataQualityScanRunsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityScanRunsResponse
func (client *Client) ListDataQualityScanRunsWithContext(ctx context.Context, tmpReq *ListDataQualityScanRunsRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityScanRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataQualityScanRunsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeFrom) {
		query["CreateTimeFrom"] = request.CreateTimeFrom
	}

	if !dara.IsNil(request.CreateTimeTo) {
		query["CreateTimeTo"] = request.CreateTimeTo
	}

	if !dara.IsNil(request.DataQualityScanId) {
		query["DataQualityScanId"] = request.DataQualityScanId
	}

	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataQualityScanRuns"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityScanRunsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data quality scan tasks in a project.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - ListDataQualityScansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityScansResponse
func (client *Client) ListDataQualityScansWithContext(ctx context.Context, request *ListDataQualityScansRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityScansResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Table) {
		query["Table"] = request.Table
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataQualityScans"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityScansResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data quality rule templates in a specified project.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - ListDataQualityTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityTemplatesResponse
func (client *Client) ListDataQualityTemplatesWithContext(ctx context.Context, request *ListDataQualityTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Catalog) {
		query["Catalog"] = request.Catalog
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataQualityTemplates"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of sharing rules for a data source.
//
// Description:
//
// 1. This operation is available for all DataWorks editions.
//
// 2. To query the sharing rules of a data source associated with a workspace, you must have the data source sharing permissions in that workspace. You must have one of the following roles in DataWorks:
//
// - Tenant Owner, Tenant Administrator, Workspace Administrator, and Workspace Owner
//
// @param request - ListDataSourceSharedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceSharedRulesResponse
func (client *Client) ListDataSourceSharedRulesWithContext(ctx context.Context, request *ListDataSourceSharedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceSharedRulesResponse, _err error) {
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
		Action:      dara.String("ListDataSourceSharedRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceSharedRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data sources that match the specified filter conditions.
//
// Description:
//
// 1. This operation is available for all DataWorks editions.
//
// 2. To call this operation, you must have one of the following roles in DataWorks:
//
// - Tenant Owner, Workspace Administrator, Deploy, Develop, Visitor, Workspace Owner, O\\&M, Model Designer, Security Administrator, Data Analyst, OpenPlatform Administrator, and Data Governance Administrator
//
// @param tmpReq - ListDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithContext(ctx context.Context, tmpReq *ListDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSources"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of databases in an instance, cluster, or data catalog in Data Map. For DLF or StarRocks data sources, you can call this API operation to query databases in a data catalog. For StarRocks data sources, you can call this API operation to query databases in internal catalogs. For other types of data sources, you can call this API operation to query databases in an instance or cluster.
//
// Description:
//
// 1. DataWorks Basic Edition or a higher edition is required.
//
// 2. For the StarRocks type, only the Internal catalog is supported.
//
// @param request - ListDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithContext(ctx context.Context, request *ListDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
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
		Action:      dara.String("ListDatabases"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the version list for a given dataset.
//
// @param request - ListDatasetVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetVersionsResponse
func (client *Client) ListDatasetVersionsWithContext(ctx context.Context, request *ListDatasetVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListDatasetVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreatorId) {
		body["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.DatasetId) {
		body["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasetVersions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of datasets. Currently, DataWorks datasets and PAI datasets are supported.
//
// @param tmpReq - ListDatasetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
func (client *Client) ListDatasetsWithContext(ctx context.Context, tmpReq *ListDatasetsRequest, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDatasetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataTypeList) {
		request.DataTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataTypeList, dara.String("DataTypeList"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.StorageTypeList) {
		request.StorageTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StorageTypeList, dara.String("StorageTypeList"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CreatorId) {
		body["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.DataTypeListShrink) {
		body["DataTypeList"] = request.DataTypeListShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.Origin) {
		body["Origin"] = request.Origin
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StorageTypeListShrink) {
		body["StorageTypeList"] = request.StorageTypeListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatasets"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of file versions pending deployment.
//
// @param tmpReq - ListDeploymentPackageFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentPackageFilesResponse
func (client *Client) ListDeploymentPackageFilesWithContext(ctx context.Context, tmpReq *ListDeploymentPackageFilesRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentPackageFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDeploymentPackageFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileIds) {
		request.FileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileIds, dara.String("FileIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		query["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.ChangeType) {
		query["ChangeType"] = request.ChangeType
	}

	if !dara.IsNil(request.CommitFrom) {
		query["CommitFrom"] = request.CommitFrom
	}

	if !dara.IsNil(request.CommitTo) {
		query["CommitTo"] = request.CommitTo
	}

	if !dara.IsNil(request.CommitUserId) {
		query["CommitUserId"] = request.CommitUserId
	}

	if !dara.IsNil(request.FileIdsShrink) {
		query["FileIds"] = request.FileIdsShrink
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SolutionId) {
		query["SolutionId"] = request.SolutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentPackageFiles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentPackageFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of deployment packages.
//
// @param request - ListDeploymentPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentPackagesResponse
func (client *Client) ListDeploymentPackagesWithContext(ctx context.Context, request *ListDeploymentPackagesRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentPackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		body["Creator"] = request.Creator
	}

	if !dara.IsNil(request.EndCreateTime) {
		body["EndCreateTime"] = request.EndCreateTime
	}

	if !dara.IsNil(request.EndExecuteTime) {
		body["EndExecuteTime"] = request.EndExecuteTime
	}

	if !dara.IsNil(request.Executor) {
		body["Executor"] = request.Executor
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentPackages"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentPackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of downstream instances for a specified instance.
//
// Description:
//
// DataWorks Basic Edition or a more advanced edition is required.
//
// @param request - ListDownstreamTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownstreamTaskInstancesResponse
func (client *Client) ListDownstreamTaskInstancesWithContext(ctx context.Context, request *ListDownstreamTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListDownstreamTaskInstancesResponse, _err error) {
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
		Action:      dara.String("ListDownstreamTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDownstreamTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of descendant tasks of a task by page.
//
// @param request - ListDownstreamTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownstreamTasksResponse
func (client *Client) ListDownstreamTasksWithContext(ctx context.Context, request *ListDownstreamTasksRequest, runtime *dara.RuntimeOptions) (_result *ListDownstreamTasksResponse, _err error) {
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
		Action:      dara.String("ListDownstreamTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDownstreamTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of entities in a collection in Data Map. Collections include categories and data albums. Entities can only be tables.
//
// Description:
//
// 1. DataWorks Professional Edition or a higher edition is required.
//
// @param request - ListEntitiesInMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEntitiesInMetaCollectionResponse
func (client *Client) ListEntitiesInMetaCollectionWithContext(ctx context.Context, request *ListEntitiesInMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *ListEntitiesInMetaCollectionResponse, _err error) {
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
		Action:      dara.String("ListEntitiesInMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEntitiesInMetaCollectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the version list of a file.
//
// @param request - ListFileVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileVersionsResponse
func (client *Client) ListFileVersionsWithContext(ctx context.Context, request *ListFileVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListFileVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFileVersions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFileVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of files by calling ListFiles.
//
// @param request - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithContext(ctx context.Context, request *ListFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CommitStatus) {
		body["CommitStatus"] = request.CommitStatus
	}

	if !dara.IsNil(request.ExactFileName) {
		body["ExactFileName"] = request.ExactFileName
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileIdIn) {
		body["FileIdIn"] = request.FileIdIn
	}

	if !dara.IsNil(request.FileTypes) {
		body["FileTypes"] = request.FileTypes
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.LastEditUser) {
		body["LastEditUser"] = request.LastEditUser
	}

	if !dara.IsNil(request.NeedAbsoluteFolderPath) {
		body["NeedAbsoluteFolderPath"] = request.NeedAbsoluteFolderPath
	}

	if !dara.IsNil(request.NeedContent) {
		body["NeedContent"] = request.NeedContent
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.UseType) {
		body["UseType"] = request.UseType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFiles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of folders.
//
// @param request - ListFoldersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFoldersResponse
func (client *Client) ListFoldersWithContext(ctx context.Context, request *ListFoldersRequest, runtime *dara.RuntimeOptions) (_result *ListFoldersResponse, _err error) {
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

	if !dara.IsNil(request.ParentFolderPath) {
		body["ParentFolderPath"] = request.ParentFolderPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFolders"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFoldersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of UDF functions in DataStudio. You can also use filter conditions to filter UDF functions.
//
// @param request - ListFunctionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionsResponse
func (client *Client) ListFunctionsWithContext(ctx context.Context, request *ListFunctionsRequest, runtime *dara.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
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
		Action:      dara.String("ListFunctions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the workspaces associated with an image.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or higher to call this operation.
//
// 2. **Before you call this operation, ensure that the AliyunServiceRoleForDataWorks service-linked role is created.**
//
// @param request - ListImageAssociatedProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageAssociatedProjectsResponse
func (client *Client) ListImageAssociatedProjectsWithContext(ctx context.Context, request *ListImageAssociatedProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListImageAssociatedProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImageAssociatedProjects"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageAssociatedProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a specified image version.
//
// Description:
//
// 1. To use this API, you must purchase DataWorks Basic Edition or a later edition.
//
// 2. **Ensure you create the service-linked role AliyunServiceRoleForDataWorks before you call this API.**
//
// @param request - ListImageVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageVersionsResponse
func (client *Client) ListImageVersionsWithContext(ctx context.Context, request *ListImageVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListImageVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
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
		Action:      dara.String("ListImageVersions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of images.
//
// Description:
//
// 1. You must have DataWorks Basic Edition or a later version to use this API.
//
// 2. **Before you use this API, make sure that the service-linked role AliyunServiceRoleForDataWorks is created.**
//
// @param tmpReq - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithContext(ctx context.Context, tmpReq *ListImagesRequest, runtime *dara.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListImagesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ProjectIds) {
		request.ProjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProjectIds, dara.String("ProjectIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProviderTypes) {
		request.ProviderTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProviderTypes, dara.String("ProviderTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Stages) {
		request.StagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Stages, dara.String("Stages"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Statuses) {
		request.StatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Statuses, dara.String("Statuses"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SupportedModules) {
		request.SupportedModulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupportedModules, dara.String("SupportedModules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SupportedTaskTypes) {
		request.SupportedTaskTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupportedTaskTypes, dara.String("SupportedTaskTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Official) {
		query["Official"] = request.Official
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectIdsShrink) {
		query["ProjectIds"] = request.ProjectIdsShrink
	}

	if !dara.IsNil(request.ProviderTypesShrink) {
		query["ProviderTypes"] = request.ProviderTypesShrink
	}

	if !dara.IsNil(request.SearchAll) {
		query["SearchAll"] = request.SearchAll
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StagesShrink) {
		query["Stages"] = request.StagesShrink
	}

	if !dara.IsNil(request.StatusesShrink) {
		query["Statuses"] = request.StatusesShrink
	}

	if !dara.IsNil(request.SupportedModulesShrink) {
		query["SupportedModules"] = request.SupportedModulesShrink
	}

	if !dara.IsNil(request.SupportedTaskTypesShrink) {
		query["SupportedTaskTypes"] = request.SupportedTaskTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImages"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data lineage relationships between two specified entities (tables, fields, OSS files, etc.) in DataWorks Data Map.
//
// Description:
//
// 1. You must purchase DataWorks Standard Edition or a higher edition to use this feature.
//
// @param request - ListLineageRelationshipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLineageRelationshipsResponse
func (client *Client) ListLineageRelationshipsWithContext(ctx context.Context, request *ListLineageRelationshipsRequest, runtime *dara.RuntimeOptions) (_result *ListLineageRelationshipsResponse, _err error) {
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
		Action:      dara.String("ListLineageRelationships"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLineageRelationshipsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the upstream and downstream lineage entities for a specified entity in DataWorks Data Map, with an option to include detailed lineage relationship information.
//
// Description:
//
// 1. You must purchase DataWorks Standard Edition or a higher edition to use this operation.
//
// 2. This operation queries the upstream and downstream entities of the current entity, as well as the lineage relationships between entities.
//
// @param request - ListLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLineagesResponse
func (client *Client) ListLineagesWithContext(ctx context.Context, request *ListLineagesRequest, runtime *dara.RuntimeOptions) (_result *ListLineagesResponse, _err error) {
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
		Action:      dara.String("ListLineages"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLineagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of MCP servers.
//
// Description:
//
// ## Request
//
// This operation retrieves a paginated list of all MCP Servers within your account. You can filter the list by search keyword and visibility level, and control pagination by specifying the maximum number of results and a next page token.
//
// - **Q**: Optional. The search keyword for a fuzzy search on MCP Server names.
//
// - **Visibility**: Optional. The visibility level for filtering the results.
//
// - **MaxResults**: Optional. The maximum number of results to return per page. By default, no limit is applied.
//
// - **NextToken**: Optional. The next page token from a previous response. Use this parameter to retrieve the next page of results.
//
// @param tmpReq - ListMcpServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpServersResponse
func (client *Client) ListMcpServersWithContext(ctx context.Context, tmpReq *ListMcpServersRequest, runtime *dara.RuntimeOptions) (_result *ListMcpServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMcpServersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Visibility) {
		request.VisibilityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Visibility, dara.String("Visibility"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Q) {
		body["Q"] = request.Q
	}

	if !dara.IsNil(request.VisibilityShrink) {
		body["Visibility"] = request.VisibilityShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcpServers"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Data Map collections. Supports querying both Data Map categories and data albums.
//
// Description:
//
// 1. You must purchase DataWorks Professional Edition or a higher edition to use this feature.
//
// @param request - ListMetaCollectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetaCollectionsResponse
func (client *Client) ListMetaCollectionsWithContext(ctx context.Context, request *ListMetaCollectionsRequest, runtime *dara.RuntimeOptions) (_result *ListMetaCollectionsResponse, _err error) {
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
		Action:      dara.String("ListMetaCollections"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetaCollectionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of metadata entities. Currently, only custom entity types are supported.
//
// @param tmpReq - ListMetaEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetaEntitiesResponse
func (client *Client) ListMetaEntitiesWithContext(ctx context.Context, tmpReq *ListMetaEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListMetaEntitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMetaEntitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AttributeFilters) {
		request.AttributeFiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AttributeFilters, dara.String("AttributeFilters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomAttributeFilters) {
		request.CustomAttributeFiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomAttributeFilters, dara.String("CustomAttributeFilters"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttributeFiltersShrink) {
		body["AttributeFilters"] = request.AttributeFiltersShrink
	}

	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.CustomAttributeFiltersShrink) {
		body["CustomAttributeFilters"] = request.CustomAttributeFiltersShrink
	}

	if !dara.IsNil(request.EntityType) {
		body["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMetaEntities"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetaEntitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of custom entity definitions, including custom entity types and extension table types.
//
// @param request - ListMetaEntityDefsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetaEntityDefsResponse
func (client *Client) ListMetaEntityDefsWithContext(ctx context.Context, request *ListMetaEntityDefsRequest, runtime *dara.RuntimeOptions) (_result *ListMetaEntityDefsResponse, _err error) {
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

	if !dara.IsNil(request.Extend) {
		body["Extend"] = request.Extend
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMetaEntityDefs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetaEntityDefsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all resource access permission application orders initiated by the current user.
//
// Description:
//
// ## Request description
//
// - This API supports paginated queries, controlled by the `NextToken` and `PageSize` parameters.
//
// - `DefSchema` is a required parameter that specifies the resource type.
//
// - The `ResourceType` list can contain multiple resource types for more precise filtering of application orders.
//
// - You can set `StartTime` and `EndTime` to limit the time range of the query.
//
// - `Statuses` allows you to filter application orders by specific statuses, such as pending approval and authorized.
//
// - If you need to filter by specific resources or authorization targets, you can provide detailed information through the `Resource` and `Grantee` fields.
//
// @param tmpReq - ListMyApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMyApplicationsResponse
func (client *Client) ListMyApplicationsWithContext(ctx context.Context, tmpReq *ListMyApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListMyApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMyApplicationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Resource) {
		request.ResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resource, dara.String("Resource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceType) {
		request.ResourceTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceType, dara.String("ResourceType"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Statuses) {
		request.StatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Statuses, dara.String("Statuses"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DefSchema) {
		body["DefSchema"] = request.DefSchema
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceShrink) {
		body["Resource"] = request.ResourceShrink
	}

	if !dara.IsNil(request.ResourceTypeShrink) {
		body["ResourceType"] = request.ResourceTypeShrink
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatusesShrink) {
		body["Statuses"] = request.StatusesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMyApplications"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMyApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query all permission application orders where the current user is an approver or stakeholder.
//
// Description:
//
// ## Request Description
//
// - This API is used to retrieve all permission application orders where the current user is an approver, including pending and processed application orders.
//
// - You can use the `Statuses` parameter to filter application orders by specific status.
//
// - `NextToken` is used for paginated requests. It can be omitted or set to `null` for the first request. For subsequent requests, pass the `NextToken` value from the previous response.
//
// - `PageSize` defaults to 10, with a maximum of 200.
//
// - `DefSchema` and `ResourceType` are required fields. Other parameters can be filled in as needed.
//
// @param tmpReq - ListMyRelatedApprovalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMyRelatedApprovalsResponse
func (client *Client) ListMyRelatedApprovalsWithContext(ctx context.Context, tmpReq *ListMyRelatedApprovalsRequest, runtime *dara.RuntimeOptions) (_result *ListMyRelatedApprovalsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMyRelatedApprovalsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccessTypes) {
		request.AccessTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccessTypes, dara.String("AccessTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Grantee) {
		request.GranteeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Grantee, dara.String("Grantee"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Resource) {
		request.ResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resource, dara.String("Resource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceType) {
		request.ResourceTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceType, dara.String("ResourceType"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Statuses) {
		request.StatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Statuses, dara.String("Statuses"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessTypesShrink) {
		body["AccessTypes"] = request.AccessTypesShrink
	}

	if !dara.IsNil(request.DefSchema) {
		body["DefSchema"] = request.DefSchema
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GranteeShrink) {
		body["Grantee"] = request.GranteeShrink
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceShrink) {
		body["Resource"] = request.ResourceShrink
	}

	if !dara.IsNil(request.ResourceTypeShrink) {
		body["ResourceType"] = request.ResourceTypeShrink
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatusesShrink) {
		body["Statuses"] = request.StatusesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMyRelatedApprovals"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMyRelatedApprovalsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of network resources for a serverless resource group.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListNetworksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworksResponse
func (client *Client) ListNetworksWithContext(ctx context.Context, request *ListNetworksRequest, runtime *dara.RuntimeOptions) (_result *ListNetworksResponse, _err error) {
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
		Action:      dara.String("ListNetworks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets a paginated list of dependent nodes for a specified data development node.
//
// @param request - ListNodeDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDependenciesResponse
func (client *Client) ListNodeDependenciesWithContext(ctx context.Context, request *ListNodeDependenciesRequest, runtime *dara.RuntimeOptions) (_result *ListNodeDependenciesResponse, _err error) {
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
		Action:      dara.String("ListNodeDependencies"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeDependenciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data development nodes with paging, and supports filtered query by specified conditions.
//
// @param request - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithContext(ctx context.Context, request *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
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
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists parameter versions.
//
// Description:
//
// This feature is available in DataWorks Professional Edition and higher editions.
//
// @param request - ListParameterVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParameterVersionsResponse
func (client *Client) ListParameterVersionsWithContext(ctx context.Context, request *ListParameterVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListParameterVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListParameterVersions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListParameterVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query parameters.
//
// Description:
//
// This feature is available in DataWorks Professional Edition or higher.
//
// @param tmpReq - ListParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParametersResponse
func (client *Client) ListParametersWithContext(ctx context.Context, tmpReq *ListParametersRequest, runtime *dara.RuntimeOptions) (_result *ListParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListParametersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Names) {
		request.NamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Names, dara.String("Names"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.NamesShrink) {
		body["Names"] = request.NamesShrink
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Scope) {
		body["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListParameters"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListParametersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the partition list of a specified table in DataWorks Data Map. Currently supports MaxCompute and HMS (EMR cluster) types.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// 2. Only MaxCompute and HMS (EMR cluster) table types are supported.
//
// 3. Before calling this API, call ListCrawlers to obtain the MetaEntityId of the metadata crawler, then call ListDatabases to obtain the database ID. For MaxCompute projects with Schema enabled, call ListSchemas to obtain the schema ID. Then call ListTables to obtain the TableId, and pass the returned table ID to this API.
//
// @param request - ListPartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPartitionsResponse
func (client *Client) ListPartitionsWithContext(ctx context.Context, request *ListPartitionsRequest, runtime *dara.RuntimeOptions) (_result *ListPartitionsResponse, _err error) {
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
		Action:      dara.String("ListPartitions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPartitionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of pending permission request orders for which the current user is the approver.
//
// Description:
//
// ## Request Description
//
// This API is used to query all pending permission request orders for which the current logged-in user is the approver. It supports filtering by multiple conditions, including resource type, time range, and approval status, to more precisely locate specific request orders.
//
// - **ResourceType**: Specifies the resource type (such as table), which can have multiple values.
//
// - **Resource**: Provides specific resource search conditions, such as the project, database, or table name.
//
// - **StartTime and EndTime**: Define the time range during which the request was submitted.
//
// - **Statuses**: Allows filtering results by approval status, for example, to view only pending approval requests.
//
// - **Grantee**: Filters request orders based on the authorization principal information.
//
// - **AccessTypes**: Filters based on the specific permission types requested (such as read or update).
//
// - **PageSize and NextToken**: Used for pagination control, specifying the amount of data returned per request and the cursor needed to retrieve the next page of data.
//
// Notes:
//
// - If no filtering conditions are provided, all matching records are returned by default.
//
// - The `NextToken` parameter can be empty or omitted for the first call. Subsequent page requests must use the `NextToken` value provided in the previous response.
//
// - The default value of `PageSize` is 10, and the maximum value is 200. If the specified value exceeds the maximum limit, the maximum value is used.
//
// - When there is no more data to return, the `HasMore` field is set to `false`, and `NextToken` will be empty or absent.
//
// ## Response Description
//
// After a successful call to this API, the response body contains paginated results and detailed information for each request order, such as the application time, resource description, authorization principal, requested permissions, and more. Additionally, the approval process status and other related metadata are provided.
//
// - **Data**: The paginated result set, including page size (`PageSize`), cursor (`NextToken`), and whether more data is available (`HasMore`).
//
// - **ApplicationQueryResponse**: The specific content of each request order, including the reason for the request, submission time, status, and a detailed list of request content.
//
// @param tmpReq - ListPendingApprovalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPendingApprovalsResponse
func (client *Client) ListPendingApprovalsWithContext(ctx context.Context, tmpReq *ListPendingApprovalsRequest, runtime *dara.RuntimeOptions) (_result *ListPendingApprovalsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPendingApprovalsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccessTypes) {
		request.AccessTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccessTypes, dara.String("AccessTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Grantee) {
		request.GranteeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Grantee, dara.String("Grantee"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Resource) {
		request.ResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resource, dara.String("Resource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceType) {
		request.ResourceTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceType, dara.String("ResourceType"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccessTypesShrink) {
		body["AccessTypes"] = request.AccessTypesShrink
	}

	if !dara.IsNil(request.DefSchema) {
		body["DefSchema"] = request.DefSchema
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GranteeShrink) {
		body["Grantee"] = request.GranteeShrink
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceShrink) {
		body["Resource"] = request.ResourceShrink
	}

	if !dara.IsNil(request.ResourceTypeShrink) {
		body["ResourceType"] = request.ResourceTypeShrink
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPendingApprovals"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPendingApprovalsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about deployment objects by deployment process ID.
//
// @param request - ListPipelineRunItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineRunItemsResponse
func (client *Client) ListPipelineRunItemsWithContext(ctx context.Context, request *ListPipelineRunItemsRequest, runtime *dara.RuntimeOptions) (_result *ListPipelineRunItemsResponse, _err error) {
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
		Action:      dara.String("ListPipelineRunItems"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineRunItemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve a paginated list of deployment processes. You can also filter this list based on specific criteria.
//
// Description:
//
//	Notice:
//
// Earlier SDK versions may not include this interface. If so, use the ListDeployments interface. It accepts the same parameters.
//
// @param request - ListPipelineRunsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineRunsResponse
func (client *Client) ListPipelineRunsWithContext(ctx context.Context, request *ListPipelineRunsRequest, runtime *dara.RuntimeOptions) (_result *ListPipelineRunsResponse, _err error) {
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
		Action:      dara.String("ListPipelineRuns"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries process definitions of a specified type.
//
// Description:
//
// ## Description
//
// - Queries process definitions of a specified policy type.
//
// - This operation supports paginated queries. You can use the`PageSize` and`PageNumber` parameters to control the page size and page number.
//
// - You can also use the`NextToken` and`MaxResults` parameters to page through large result sets.
//
// - The response includes the total count, page size, current page number, and a list of process definitions.
//
// - Each process definition includes key attributes, such as its ID, enabled status, and priority.
//
// @param request - ListProcessDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProcessDefinitionsResponse
func (client *Client) ListProcessDefinitionsWithContext(ctx context.Context, request *ListProcessDefinitionsRequest, runtime *dara.RuntimeOptions) (_result *ListProcessDefinitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProcessDefinitions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProcessDefinitionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets a paginated list of workspace member details.
//
// Description:
//
// This feature is available in DataWorks Basic Edition and higher.
//
// @param tmpReq - ListProjectMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembersWithContext(ctx context.Context, tmpReq *ListProjectMembersRequest, runtime *dara.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("RoleCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("UserIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserIdsShrink) {
		body["UserIds"] = request.UserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectMembers"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a paginated list of roles in a workspace.
//
// Description:
//
// This feature is available in DataWorks Basic Edition and higher.
//
// @param tmpReq - ListProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectRolesResponse
func (client *Client) ListProjectRolesWithContext(ctx context.Context, tmpReq *ListProjectRolesRequest, runtime *dara.RuntimeOptions) (_result *ListProjectRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Codes) {
		request.CodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Codes, dara.String("Codes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Names) {
		request.NamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Names, dara.String("Names"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CodesShrink) {
		body["Codes"] = request.CodesShrink
	}

	if !dara.IsNil(request.NamesShrink) {
		body["Names"] = request.NamesShrink
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectRoles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of DataWorks workspaces of the tenant to which your account belongs.
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithContext(ctx context.Context, tmpReq *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AliyunResourceTags) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, dara.String("AliyunResourceTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Names) {
		request.NamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Names, dara.String("Names"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunResourceGroupId) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !dara.IsNil(request.AliyunResourceTagsShrink) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !dara.IsNil(request.DevEnvironmentEnabled) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !dara.IsNil(request.DevRoleDisabled) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.NamesShrink) {
		body["Names"] = request.NamesShrink
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PaiTaskEnabled) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of workspaces with which a resource group is associated
//
// Description:
//
// 1. DataWorks Basic Edition or a more advanced edition is required to use this feature.
//
// 2. **Make sure that the service-linked role AliyunServiceRoleForDataWorks has been created before you call this operation.**
//
// 3. This operation returns only the workspaces that the current caller has access to. Unauthorized workspaces are not included in the response.
//
// @param request - ListResourceGroupAssociateProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupAssociateProjectsResponse
func (client *Client) ListResourceGroupAssociateProjectsWithContext(ctx context.Context, request *ListResourceGroupAssociateProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupAssociateProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroupAssociateProjects"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupAssociateProjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metric data of a resource group.
//
// Description:
//
// 1. DataWorks Basic Edition or a more advanced edition is required to use this feature.
//
// 2. **Make sure that the service-linked role AliyunServiceRoleForDataWorks has been created before you call this operation.**
//
// 3. This operation applies only to serverless resource groups.
//
// @param request - ListResourceGroupMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupMetricDataResponse
func (client *Client) ListResourceGroupMetricDataWithContext(ctx context.Context, request *ListResourceGroupMetricDataRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		body["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Length) {
		body["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		body["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroupMetricData"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupMetricDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of resource groups.
//
// Description:
//
// 1. This operation requires DataWorks Basic Edition or higher.
//
// 2. **Before you call this operation, make sure that you have created the service-linked role AliyunServiceRoleForDataWorks.**
//
// @param tmpReq - ListResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithContext(ctx context.Context, tmpReq *ListResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListResourceGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AliyunResourceTags) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, dara.String("AliyunResourceTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceGroupTypes) {
		request.ResourceGroupTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceGroupTypes, dara.String("ResourceGroupTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Statuses) {
		request.StatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Statuses, dara.String("Statuses"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroups"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource files with pagination and filtering support.
//
// @param request - ListResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithContext(ctx context.Context, request *ListResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
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
		Action:      dara.String("ListResources"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the route list of a network resource.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param request - ListRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoutesResponse
func (client *Client) ListRoutesWithContext(ctx context.Context, request *ListRoutesRequest, runtime *dara.RuntimeOptions) (_result *ListRoutesResponse, _err error) {
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
		Action:      dara.String("ListRoutes"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRoutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of schemas under a specified database or MaxCompute project in Data Map. Currently supports MaxCompute and Holo types.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or higher to use this feature.
//
// @param tmpReq - ListSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemasResponse
func (client *Client) ListSchemasWithContext(ctx context.Context, tmpReq *ListSchemasRequest, runtime *dara.RuntimeOptions) (_result *ListSchemasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSchemasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSchemas"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSchemasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of security policies based on specified conditions. This operation requires DataWorks tenant administrator or security administrator permissions.
//
// Description:
//
// ## Request
//
// - This API retrieves a paginated list of configured security policies.
//
// - The `ControlModule` and `ControlSubModule` parameters filter policies by a specific module or submodule.
//
// - The `PageNum` and `PageSize` parameters control pagination. `PageNum` specifies the page number to retrieve (default: 1), and `PageSize` specifies the number of policies to return per page (default: 20).
//
// - Use the `MaxResults` and `NextToken` private parameters for advanced pagination.
//
// @param request - ListSecurityStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecurityStrategiesResponse
func (client *Client) ListSecurityStrategiesWithContext(ctx context.Context, request *ListSecurityStrategiesRequest, runtime *dara.RuntimeOptions) (_result *ListSecurityStrategiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ControlModule) {
		body["ControlModule"] = request.ControlModule
	}

	if !dara.IsNil(request.ControlSubModule) {
		body["ControlSubModule"] = request.ControlSubModule
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
		Action:      dara.String("ListSecurityStrategies"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecurityStrategiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the run records of a created node by Name with paging. The JobRunId in each record is active for retrieving the results of a specific run, and the ExecutorJobId is active for querying details, retrieving logs, or stopping the run.
//
// Description:
//
// ## Scenarios
//
// View the historical run records of a semantic job with pagination to obtain the run ID, executor job ID, status, and time information for each submission.
//
// ## Procedure
//
// 1. Use the job name from `CreateSemanticJob.Data.Name` or `ListSemanticJobs` as the `JobName`.
//
// 2. Use `PageNumber` and `PageSize` to read records page by page.
//
// 3. Use the `JobRunId` from a record to call `DownloadSemanticResults`, and use the `ExecutorJobId` to call the detail, log, or stop operations.
//
// ## Before you begin
//
// Pagination starts from page 1 by default. Each page contains a maximum of 200 records.
//
// @param request - ListSemanticJobRunsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSemanticJobRunsResponse
func (client *Client) ListSemanticJobRunsWithContext(ctx context.Context, request *ListSemanticJobRunsRequest, runtime *dara.RuntimeOptions) (_result *ListSemanticJobRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobName) {
		body["JobName"] = request.JobName
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
		Action:      dara.String("ListSemanticJobRuns"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSemanticJobRunsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the semantic node definitions of the current tenant with paging. The Name, ProjectId, and Source fields in the list items can be used for running/deleting nodes, querying run details, and verifying input scope, respectively.
//
// Description:
//
// ## Scenarios
//
// Queries the saved semantic node definitions of the current tenant with paging. Use this operation to display the node list, select a node to run, or obtain the workspace to which a node belongs.
//
// ## Invoke flow
//
// 1. Use `PageNumber` and `PageSize` to read `Data.SemanticJobs` with paging.
//
// 2. Use the `Name` field of a list item to invoke `RunSemanticJob`, `DeleteSemanticJob`, or `ListSemanticJobRuns`.
//
// 3. Use the `ProjectId` field of a list item together with `ExecutorJobId` to invoke the details, log, and stop operations.
//
// ## Notes
//
// This operation returns node definitions, not real-time run statuses. To query run statuses, invoke `ListSemanticJobRuns`.
//
// @param request - ListSemanticJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSemanticJobsResponse
func (client *Client) ListSemanticJobsWithContext(ctx context.Context, request *ListSemanticJobsRequest, runtime *dara.RuntimeOptions) (_result *ListSemanticJobsResponse, _err error) {
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
		Action:      dara.String("ListSemanticJobs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSemanticJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the Skills in your account.
//
// Description:
//
// ## Request
//
// This operation lists the Skills in your account. You can filter the results by criteria such as a search keyword and visibility level.
//
// - **Q**: An optional search keyword for a fuzzy match on Skill names.
//
// - **Visibility**: An optional parameter to filter Skills by their visibility level. You can specify multiple values.
//
// - **MaxResults**: An optional parameter that specifies the maximum number of results to return per page.
//
// - **NextToken**: An optional pagination token for retrieving the next page of results. Omit this parameter for the first request. For subsequent requests, pass the `NextToken` value from the previous response to fetch the next page.
//
// @param tmpReq - ListSkillsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSkillsResponse
func (client *Client) ListSkillsWithContext(ctx context.Context, tmpReq *ListSkillsRequest, runtime *dara.RuntimeOptions) (_result *ListSkillsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSkillsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Visibility) {
		request.VisibilityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Visibility, dara.String("Visibility"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Q) {
		body["Q"] = request.Q
	}

	if !dara.IsNil(request.VisibilityShrink) {
		body["Visibility"] = request.VisibilityShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSkills"),
		Version:     dara.String("2024-05-18"),
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
// Retrieves a list of version snapshots.
//
// @param request - ListSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshotsWithContext(ctx context.Context, request *ListSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSnapshots"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data tables in DataWorks Data Map. For types that do not support the schema level, you can query data tables under a specified database. For types that support the schema level, you can query data tables under a specified database, MaxCompute project, or schema. The response contains only basic table information and does not include technical metadata or business metadata.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - ListTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithContext(ctx context.Context, tmpReq *ListTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTablesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableTypes) {
		request.TableTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableTypes, dara.String("TableTypes"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of operation logs for a task instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// Only operation logs generated within the previous 31 days can be queried.
//
// @param request - ListTaskInstanceOperationLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskInstanceOperationLogsResponse
func (client *Client) ListTaskInstanceOperationLogsWithContext(ctx context.Context, request *ListTaskInstanceOperationLogsRequest, runtime *dara.RuntimeOptions) (_result *ListTaskInstanceOperationLogsResponse, _err error) {
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
		Action:      dara.String("ListTaskInstanceOperationLogs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskInstanceOperationLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances. You can also specify filter conditions to query specific instances.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - ListTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskInstancesResponse
func (client *Client) ListTaskInstancesWithContext(ctx context.Context, tmpReq *ListTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResource) {
		body["RuntimeResource"] = request.RuntimeResource
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskIdsShrink) {
		body["TaskIds"] = request.TaskIdsShrink
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TriggerRecurrence) {
		body["TriggerRecurrence"] = request.TriggerRecurrence
	}

	if !dara.IsNil(request.TriggerType) {
		body["TriggerType"] = request.TriggerType
	}

	if !dara.IsNil(request.UnifiedWorkflowInstanceId) {
		body["UnifiedWorkflowInstanceId"] = request.UnifiedWorkflowInstanceId
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkflowInstanceId) {
		body["WorkflowInstanceId"] = request.WorkflowInstanceId
	}

	if !dara.IsNil(request.WorkflowInstanceType) {
		body["WorkflowInstanceType"] = request.WorkflowInstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of operation logs for a task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// Only operation logs generated within the previous 31 days can be queried.
//
// @param request - ListTaskOperationLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskOperationLogsResponse
func (client *Client) ListTaskOperationLogsWithContext(ctx context.Context, request *ListTaskOperationLogsRequest, runtime *dara.RuntimeOptions) (_result *ListTaskOperationLogsResponse, _err error) {
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
		Action:      dara.String("ListTaskOperationLogs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskOperationLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tasks by page. You can also specify filter conditions to query tasks.
//
// Description:
//
// DataWorks Basic Edition or higher is required.
//
// @param tmpReq - ListTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithContext(ctx context.Context, tmpReq *ListTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResource) {
		body["RuntimeResource"] = request.RuntimeResource
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TriggerRecurrence) {
		body["TriggerRecurrence"] = request.TriggerRecurrence
	}

	if !dara.IsNil(request.TriggerType) {
		body["TriggerType"] = request.TriggerType
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor instances of an instance by page.
//
// @param request - ListUpstreamTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUpstreamTaskInstancesResponse
func (client *Client) ListUpstreamTaskInstancesWithContext(ctx context.Context, request *ListUpstreamTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListUpstreamTaskInstancesResponse, _err error) {
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
		Action:      dara.String("ListUpstreamTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUpstreamTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor tasks of a task by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListUpstreamTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUpstreamTasksResponse
func (client *Client) ListUpstreamTasksWithContext(ctx context.Context, request *ListUpstreamTasksRequest, runtime *dara.RuntimeOptions) (_result *ListUpstreamTasksResponse, _err error) {
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
		Action:      dara.String("ListUpstreamTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUpstreamTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workflows in DataStudio. You can also specify filter conditions to query specific workflows.
//
// @param request - ListWorkflowDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowDefinitionsResponse
func (client *Client) ListWorkflowDefinitionsWithContext(ctx context.Context, request *ListWorkflowDefinitionsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowDefinitionsResponse, _err error) {
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
		Action:      dara.String("ListWorkflowDefinitions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowDefinitionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paging list of workflow instances, with optional filtered query by conditions.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param tmpReq - ListWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowInstancesResponse
func (client *Client) ListWorkflowInstancesWithContext(ctx context.Context, tmpReq *ListWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		body["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.UnifiedWorkflowInstanceId) {
		body["UnifiedWorkflowInstanceId"] = request.UnifiedWorkflowInstanceId
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workflows by page. You can also specify filter conditions to query workflows.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - ListWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowsResponse
func (client *Client) ListWorkflowsWithContext(ctx context.Context, tmpReq *ListWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWorkflowsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TriggerType) {
		body["TriggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflows"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Loads the message history of an existing session and returns it as an SSE stream. If the session does not exist, the server sends a JSONRPCResponse.error with a code of 400 through the SSE stream. The Content-Type is text/event-stream. Use this operation to restore session context.
//
// Description:
//
// ## Request
//
// - This operation retrieves session details and streams the Agent response using Server-Sent Events (SSE).
//
// - If the target session does not exist, the operation returns an error frame with an error code of 400.
//
// - The response includes information about the Agent\\"s request processing, such as message chunks, thought processes, and tool call status updates.
//
// - The `stopReason` field indicates why the Agent stops the current turn. Possible values include reaching the maximum turn limit or being canceled.
//
// - The returned content conforms to the Agent Client Protocol (ACP). For more information, see https\\://agentclientprotocol.com.
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
// Loads the message history of an existing session and returns it as an SSE stream. If the session does not exist, the server sends a JSONRPCResponse.error with a code of 400 through the SSE stream. The Content-Type is text/event-stream. Use this operation to restore session context.
//
// Description:
//
// ## Request
//
// - This operation retrieves session details and streams the Agent response using Server-Sent Events (SSE).
//
// - If the target session does not exist, the operation returns an error frame with an error code of 400.
//
// - The response includes information about the Agent\\"s request processing, such as message chunks, thought processes, and tool call status updates.
//
// - The `stopReason` field indicates why the Agent stops the current turn. Possible values include reaching the maximum turn limit or being canceled.
//
// - The returned content conforms to the Agent Client Protocol (ACP). For more information, see https\\://agentclientprotocol.com.
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
		Version:     dara.String("2024-05-18"),
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
// Moves a user-defined function (UDF) to a path in DataStudio.
//
// @param request - MoveFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveFunctionResponse
func (client *Client) MoveFunctionWithContext(ctx context.Context, request *MoveFunctionRequest, runtime *dara.RuntimeOptions) (_result *MoveFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a node to a path in DataStudio.
//
// @param request - MoveNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveNodeResponse
func (client *Client) MoveNodeWithContext(ctx context.Context, request *MoveNodeRequest, runtime *dara.RuntimeOptions) (_result *MoveNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a file resource to a path in DataStudio.
//
// @param request - MoveResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceResponse
func (client *Client) MoveResourceWithContext(ctx context.Context, request *MoveResourceRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a workflow to a path in DataStudio.
//
// @param request - MoveWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveWorkflowDefinitionResponse
func (client *Client) MoveWorkflowDefinitionWithContext(ctx context.Context, request *MoveWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *MoveWorkflowDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveWorkflowDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Previews the content of a specified dataset version. Currently only text file preview is supported for OSS-type datasets. Supported MIME types 1. application/json 2. application/xml 3. text/html 4. text/plain 5. application/octet-stream
//
// @param request - PreviewDatasetVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewDatasetVersionResponse
func (client *Client) PreviewDatasetVersionWithContext(ctx context.Context, request *PreviewDatasetVersionRequest, runtime *dara.RuntimeOptions) (_result *PreviewDatasetVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewDatasetVersion"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewDatasetVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a user prompt to an existing session and returns the Agent response in streaming mode.
//
// Description:
//
// ## Operation description
//
// - This API sends a user prompt to a specified session ID and accepts the Agent response in SSE (Server-Sent Events) streaming mode.
//
// - The response may include message fragments, thinking procedures, tool calling status updates, and other information.
//
// - If the specified session does not exist, a 400 fault is returned through an SSE error frame.
//
// - The `stopReason` field indicates why the Agent stopped the current conversation turn.
//
// - Multiple types of content blocks are supported as prompt input, such as text and OSS file download links.
//
// - You can optionally provide additional meta information `Meta` to pass more context to the server.
//
// - The returned content conforms to the open-source Agent Client Protocol (ACP) specification. For more information, visit: https://agentclientprotocol.com
//
// - **Before invoking this API, make sure you fully understand the billing methods and pricing of the Data Agent product**: https://www.alibabacloud.com/help/en/dataworks/dataworks-data-agent-agent-billing
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
// Sends a user prompt to an existing session and returns the Agent response in streaming mode.
//
// Description:
//
// ## Operation description
//
// - This API sends a user prompt to a specified session ID and accepts the Agent response in SSE (Server-Sent Events) streaming mode.
//
// - The response may include message fragments, thinking procedures, tool calling status updates, and other information.
//
// - If the specified session does not exist, a 400 fault is returned through an SSE error frame.
//
// - The `stopReason` field indicates why the Agent stopped the current conversation turn.
//
// - Multiple types of content blocks are supported as prompt input, such as text and OSS file download links.
//
// - You can optionally provide additional meta information `Meta` to pass more context to the server.
//
// - The returned content conforms to the open-source Agent Client Protocol (ACP) specification. For more information, visit: https://agentclientprotocol.com
//
// - **Before invoking this API, make sure you fully understand the billing methods and pricing of the Data Agent product**: https://www.alibabacloud.com/help/en/dataworks/dataworks-data-agent-agent-billing
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
		Version:     dara.String("2024-05-18"),
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
// Remove an entity object from a Data Map collection. The collection supports Data Map categories and data albums, and the entity currently supports only the Data Table type.
//
// When removing an entity from a data album, the caller must have the AliyunDataWorksFullAccess permission or be the creator or administrator of the album.
//
// Description:
//
// 1. You must purchase DataWorks Professional Edition or a higher version to use this feature.
//
// @param request - RemoveEntityFromMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveEntityFromMetaCollectionResponse
func (client *Client) RemoveEntityFromMetaCollectionWithContext(ctx context.Context, request *RemoveEntityFromMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *RemoveEntityFromMetaCollectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MetaCollectionId) {
		query["MetaCollectionId"] = request.MetaCollectionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveEntityFromMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveEntityFromMetaCollectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes multiple upstream dependencies of an instance at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - RemoveTaskInstanceDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTaskInstanceDependenciesResponse
func (client *Client) RemoveTaskInstanceDependenciesWithContext(ctx context.Context, tmpReq *RemoveTaskInstanceDependenciesRequest, runtime *dara.RuntimeOptions) (_result *RemoveTaskInstanceDependenciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveTaskInstanceDependenciesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpstreamTaskInstanceIds) {
		request.UpstreamTaskInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpstreamTaskInstanceIds, dara.String("UpstreamTaskInstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.UpstreamTaskInstanceIdsShrink) {
		body["UpstreamTaskInstanceIds"] = request.UpstreamTaskInstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTaskInstanceDependencies"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTaskInstanceDependenciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a user-defined function (UDF) in DataStudio.
//
// @param request - RenameFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameFunctionResponse
func (client *Client) RenameFunctionWithContext(ctx context.Context, request *RenameFunctionRequest, runtime *dara.RuntimeOptions) (_result *RenameFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a node in DataStudio.
//
// @param request - RenameNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameNodeResponse
func (client *Client) RenameNodeWithContext(ctx context.Context, request *RenameNodeRequest, runtime *dara.RuntimeOptions) (_result *RenameNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a file resource in DataStudio.
//
// @param request - RenameResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameResourceResponse
func (client *Client) RenameResourceWithContext(ctx context.Context, request *RenameResourceRequest, runtime *dara.RuntimeOptions) (_result *RenameResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a workflow in DataStudio.
//
// @param request - RenameWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameWorkflowDefinitionResponse
func (client *Client) RenameWorkflowDefinitionWithContext(ctx context.Context, request *RenameWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *RenameWorkflowDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameWorkflowDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns multiple node instances in a batch.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a more advanced edition to use this operation.
//
// @param tmpReq - RerunTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RerunTaskInstancesResponse
func (client *Client) RerunTaskInstancesWithContext(ctx context.Context, tmpReq *RerunTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *RerunTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RerunTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.UseLatestConfig) {
		body["UseLatestConfig"] = request.UseLatestConfig
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RerunTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RerunTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Rerun a workflow instance
//
// @param tmpReq - RerunWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RerunWorkflowInstancesResponse
func (client *Client) RerunWorkflowInstancesWithContext(ctx context.Context, tmpReq *RerunWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *RerunWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RerunWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.EndTriggerTime) {
		body["EndTriggerTime"] = request.EndTriggerTime
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.FilterShrink) {
		body["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StartTriggerTime) {
		body["StartTriggerTime"] = request.StartTriggerTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RerunWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RerunWorkflowInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes multiple suspended instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ResumeTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeTaskInstancesResponse
func (client *Client) ResumeTaskInstancesWithContext(ctx context.Context, tmpReq *ResumeTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ResumeTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ResumeTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes roles from a workspace member.
//
// Description:
//
// DataWorks Basic Edition or a more advanced edition is required to use this operation.
//
// @param tmpReq - RevokeMemberProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeMemberProjectRolesResponse
func (client *Client) RevokeMemberProjectRolesWithContext(ctx context.Context, tmpReq *RevokeMemberProjectRolesRequest, runtime *dara.RuntimeOptions) (_result *RevokeMemberProjectRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RevokeMemberProjectRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("RoleCodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeMemberProjectRoles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeMemberProjectRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back a specified parameter.
//
// Description:
//
// This operation is available only in DataWorks Professional Edition or a later version.
//
// @param request - RollbackParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackParameterResponse
func (client *Client) RollbackParameterWithContext(ctx context.Context, request *RollbackParameterRequest, runtime *dara.RuntimeOptions) (_result *RollbackParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.RollbackVersion) {
		body["RollbackVersion"] = request.RollbackVersion
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackParameter"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackParameterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a specified metadata crawler to run and returns the submit status and associated task instance information.
//
// Description:
//
// ## Scenarios
//
// Submits a run request for a specified metadata crawler.
//
// ## Recommended process
//
// 1. Call `ListCrawlers` to query the IDs of available crawlers.
//
// 2. Call this operation to submit a run request.
//
// 3. Call `ListCrawlerRuns` to query the final run status.
//
// ## Edition requirements
//
// DataWorks Basic Edition or higher is required.
//
// ## Billing description
//
// Running a collection task consumes compute resources and may incur fees. The actual fees depend on the resource group used and the DataWorks billing rules.
//
// If the crawler has the AI metadata description feature enabled (`EnableAiComment=true`), collecting metadata and generating AI descriptions consumes tokens. For information about the complimentary token quota and billing rules after the quota is exceeded, see [Data Agent billing](https://www.alibabacloud.com/help/en/dataworks/dataworks-data-agent-agent-billing).
//
// ## Precautions
//
// A successful response only indicates that the run request has been accepted. It does not indicate that the collection task is complete.
//
// @param request - RunCrawlerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCrawlerResponse
func (client *Client) RunCrawlerWithContext(ctx context.Context, request *RunCrawlerRequest, runtime *dara.RuntimeOptions) (_result *RunCrawlerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCrawler"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCrawlerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a saved semantic job for execution by name and returns the run identifier and executor job identifier. A successful call indicates that the job has been submitted, not that the semantic model results have been generated.
//
// Description:
//
// ## Description
//
// Loads a saved semantic job definition by `Name` and submits a new analysis run to the executor. This operation does not accept runtime `Source`, resource group, or reference file overrides. The execution always uses the configuration saved by `CreateSemanticJob`.
//
// ## Pre-execution validation
//
// The service validates the existence and access permissions of the job, and re-validates whether the associated files still exist. For files associated through `ReferenceFileIds`, the service resolves them to temporary addresses readable by the current run before submission. Deleting a file after upload or specifying an invalid file ID causes the submission to fail.
//
// ## Response and What to do next
//
// `Data.JobRunId` is the identity of the current semantics node run and is used by `DownloadSemanticResults` to download the exact output of this run. `Data.ExecutorJobId` is the identity of the executor node and is used by `GetSemanticJobDetail`, `GetSemanticJobLog`, and `KillSemanticJob`. A successful response indicates that the executor has accepted the submission, not that the model analysis or result files are complete.
//
// ## Billing
//
// **Before using this operation, make sure that you fully understand the billing method and pricing of the [model calls](https://www.alibabacloud.com/help/en/dataworks/dataworks-data-agent-agent-billing) used by semantic construction.**
//
// @param request - RunSemanticJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSemanticJobResponse
func (client *Client) RunSemanticJobWithContext(ctx context.Context, request *RunSemanticJobRequest, runtime *dara.RuntimeOptions) (_result *RunSemanticJobResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSemanticJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSemanticJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the statuses of multiple instances to successful at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - SetSuccessTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSuccessTaskInstancesResponse
func (client *Client) SetSuccessTaskInstancesWithContext(ctx context.Context, tmpReq *SetSuccessTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *SetSuccessTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetSuccessTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSuccessTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSuccessTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a new-version synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - StartDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDIJobResponse
func (client *Client) StartDIJobWithContext(ctx context.Context, tmpReq *StartDIJobRequest, runtime *dara.RuntimeOptions) (_result *StartDIJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RealtimeStartSettings) {
		request.RealtimeStartSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RealtimeStartSettings, dara.String("RealtimeStartSettings"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDIJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts multiple workflow instances in a batch.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param tmpReq - StartWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartWorkflowInstancesResponse
func (client *Client) StartWorkflowInstancesWithContext(ctx context.Context, tmpReq *StartWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *StartWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartWorkflowInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the currently running task of a specified metadata crawler.
//
// Description:
//
// ## Scenarios
//
// Stops the currently running task of a specified metadata crawler.
//
// ## Recommended workflow
//
// 1. Call `ListCrawlerRuns` to confirm that the crawler has a running task.
//
// 2. Call this operation to submit a stop request.
//
// 3. Call `ListCrawlerRuns` again to confirm the final run status.
//
// ## Edition requirements
//
// DataWorks Basic Edition or a more advanced edition is required.
//
// ## Precautions
//
// The call fails if the crawler has no running task. A successful response only indicates that the stop request has been accepted.
//
// @param request - StopCrawlerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCrawlerResponse
func (client *Client) StopCrawlerWithContext(ctx context.Context, request *StopCrawlerRequest, runtime *dara.RuntimeOptions) (_result *StopCrawlerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCrawler"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCrawlerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - StopDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDIJobResponse
func (client *Client) StopDIJobWithContext(ctx context.Context, request *StopDIJobRequest, runtime *dara.RuntimeOptions) (_result *StopDIJobResponse, _err error) {
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
		Action:      dara.String("StopDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDIJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Withdraws or terminates a specified process instance.
//
// Description:
//
// ## Description
//
// - Requesters can use this operation to withdraw an approval process they initiated.
//
// - Only the initiator of the approval process can call this operation.
//
// - After a successful call, the operation terminates the approval process and updates its status to withdrawn.
//
// @param request - StopProcessInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopProcessInstanceResponse
func (client *Client) StopProcessInstanceWithContext(ctx context.Context, request *StopProcessInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopProcessInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProcessInstanceId) {
		body["ProcessInstanceId"] = request.ProcessInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopProcessInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopProcessInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops multiple instances in a batch.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this operation.
//
// @param tmpReq - StopTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTaskInstancesResponse
func (client *Client) StopTaskInstancesWithContext(ctx context.Context, tmpReq *StopTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *StopTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops multiple workflow instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - StopWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopWorkflowInstancesResponse
func (client *Client) StopWorkflowInstancesWithContext(ctx context.Context, tmpReq *StopWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *StopWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopWorkflowInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Commits a file to the development environment of the scheduling system to generate a task.
//
// @param request - SubmitFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitFileResponse
func (client *Client) SubmitFileWithContext(ctx context.Context, request *SubmitFileRequest, runtime *dara.RuntimeOptions) (_result *SubmitFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.SkipAllDeployFileExtensions) {
		body["SkipAllDeployFileExtensions"] = request.SkipAllDeployFileExtensions
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends multiple instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - SuspendTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendTaskInstancesResponse
func (client *Client) SuspendTaskInstancesWithContext(ctx context.Context, tmpReq *SuspendTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *SuspendTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SuspendTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to data assets.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - TagDataAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagDataAssetsResponse
func (client *Client) TagDataAssetsWithContext(ctx context.Context, tmpReq *TagDataAssetsRequest, runtime *dara.RuntimeOptions) (_result *TagDataAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TagDataAssetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataAssetIds) {
		request.DataAssetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataAssetIds, dara.String("DataAssetIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoTraceEnabled) {
		query["AutoTraceEnabled"] = request.AutoTraceEnabled
	}

	if !dara.IsNil(request.DataAssetIdsShrink) {
		query["DataAssetIds"] = request.DataAssetIdsShrink
	}

	if !dara.IsNil(request.DataAssetType) {
		query["DataAssetType"] = request.DataAssetType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagDataAssets"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagDataAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Test the connectivity of a data source on a resource group.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// 2. You must have at least one of the following roles in the DataWorks project space:
//
//	Tenant Owner, Space Administrator, Deployment, Developer, Project Owner, or O\\&M.
//
// @param request - TestDataSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestDataSourceConnectivityResponse
func (client *Client) TestDataSourceConnectivityWithContext(ctx context.Context, request *TestDataSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *TestDataSourceConnectivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestDataSourceConnectivity"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestDataSourceConnectivityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a task to run by using an HTTP Trigger node at a specified time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - TriggerSchedulerTaskInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerSchedulerTaskInstanceResponse
func (client *Client) TriggerSchedulerTaskInstanceWithContext(ctx context.Context, request *TriggerSchedulerTaskInstanceRequest, runtime *dara.RuntimeOptions) (_result *TriggerSchedulerTaskInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TriggerTime) {
		body["TriggerTime"] = request.TriggerTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerSchedulerTaskInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerSchedulerTaskInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from data assets.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - UnTagDataAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagDataAssetsResponse
func (client *Client) UnTagDataAssetsWithContext(ctx context.Context, tmpReq *UnTagDataAssetsRequest, runtime *dara.RuntimeOptions) (_result *UnTagDataAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UnTagDataAssetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataAssetIds) {
		request.DataAssetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataAssetIds, dara.String("DataAssetIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataAssetIdsShrink) {
		query["DataAssetIds"] = request.DataAssetIdsShrink
	}

	if !dara.IsNil(request.DataAssetType) {
		query["DataAssetType"] = request.DataAssetType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnTagDataAssets"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnTagDataAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a custom alert monitoring rule.
//
// @param tmpReq - UpdateAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertRuleResponse
func (client *Client) UpdateAlertRuleWithContext(ctx context.Context, tmpReq *UpdateAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAlertRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerCondition) {
		request.TriggerConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerCondition, dara.String("TriggerCondition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.TriggerConditionShrink) {
		query["TriggerCondition"] = request.TriggerConditionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a workflow.
//
// @param request - UpdateBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBusinessResponse
func (client *Client) UpdateBusinessWithContext(ctx context.Context, request *UpdateBusinessRequest, runtime *dara.RuntimeOptions) (_result *UpdateBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBusinessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the business metadata for a column in a data map. This operation can only update the business description and custom attributes.
//
// Description:
//
// 1. This operation requires DataWorks Basic Edition or a later version.
//
// 2. This operation supports only MaxCompute, hms, and dlf tables.
//
// @param tmpReq - UpdateColumnBusinessMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateColumnBusinessMetadataResponse
func (client *Client) UpdateColumnBusinessMetadataWithContext(ctx context.Context, tmpReq *UpdateColumnBusinessMetadataRequest, runtime *dara.RuntimeOptions) (_result *UpdateColumnBusinessMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateColumnBusinessMetadataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomAttributes) {
		request.CustomAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomAttributes, dara.String("CustomAttributes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomAttributesShrink) {
		body["CustomAttributes"] = request.CustomAttributesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateColumnBusinessMetadata"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateColumnBusinessMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates components.
//
// Description:
//
// This API is currently in trial. To use this API, submit an application. After the administrator adds you to the trial list, you can call this API.
//
// @param request - UpdateComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComponentResponse
func (client *Client) UpdateComponentWithContext(ctx context.Context, request *UpdateComponentRequest, runtime *dara.RuntimeOptions) (_result *UpdateComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentId) {
		body["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComponent"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComponentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the specified computing resource based on the computing resource ID.
//
// Description:
//
// 1. You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// 2. You must have at least one of the following roles in the DataWorks project space:
//
// 3. Tenant Owner, tenant administrator, Space Administrator, Project Owner, or O\\&M
//
// @param request - UpdateComputeResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComputeResourceResponse
func (client *Client) UpdateComputeResourceWithContext(ctx context.Context, request *UpdateComputeResourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateComputeResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionProperties) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !dara.IsNil(request.ConnectionPropertiesMode) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComputeResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComputeResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the resource group, collection scope, scheduling, and extension configurations of a specified metadata crawler.
//
// Description:
//
// ## Scenarios
//
// Partially updates the resource group, collection scope, scheduling, AI metadata description, or extension configurations of a specified metadata crawler.
//
// ## Recommended workflow
//
// 1. Call `GetCrawler` to query the current configurations.
//
// 2. Call `GetCrawlerTypeCapabilities` to check the configuration capabilities supported by the crawler type.
//
// 3. Call this operation with only the fields that you want to update.
//
// ## Edition requirements
//
// DataWorks Basic Edition or a more advanced edition is required.
//
// ## Precautions
//
// At least one updatable field must be provided. Fields that are not provided remain unchanged.
//
// @param tmpReq - UpdateCrawlerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCrawlerResponse
func (client *Client) UpdateCrawlerWithContext(ctx context.Context, tmpReq *UpdateCrawlerRequest, runtime *dara.RuntimeOptions) (_result *UpdateCrawlerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCrawlerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("Options"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Scope) {
		request.ScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scope, dara.String("Scope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnableAiComment) {
		body["EnableAiComment"] = request.EnableAiComment
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.OptionsShrink) {
		body["Options"] = request.OptionsShrink
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		body["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.ScopeShrink) {
		body["Scope"] = request.ScopeShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCrawler"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCrawlerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom attribute.
//
// @param tmpReq - UpdateCustomAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomAttributeResponse
func (client *Client) UpdateCustomAttributeWithContext(ctx context.Context, tmpReq *UpdateCustomAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateCustomAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EntityTypes) {
		request.EntityTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityTypes, dara.String("EntityTypes"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ValueEnums) {
		request.ValueEnumsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ValueEnums, dara.String("ValueEnums"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DisplayEnabled) {
		body["DisplayEnabled"] = request.DisplayEnabled
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EntityTypesShrink) {
		body["EntityTypes"] = request.EntityTypesShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.SearchFilterEnabled) {
		body["SearchFilterEnabled"] = request.SearchFilterEnabled
	}

	if !dara.IsNil(request.ValueEnumsShrink) {
		body["ValueEnums"] = request.ValueEnumsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomAttribute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an alert rule configured for a synchronization task.
//
// @param tmpReq - UpdateDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIAlarmRuleResponse
func (client *Client) UpdateDIAlarmRuleWithContext(ctx context.Context, tmpReq *UpdateDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDIAlarmRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDIAlarmRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotificationSettings) {
		request.NotificationSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationSettings, dara.String("NotificationSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerConditions) {
		request.TriggerConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConditions, dara.String("TriggerConditions"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDIAlarmRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDIAlarmRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a Data Integration task.
//
// Description:
//
// This feature requires DataWorks Basic Edition or higher.
//
// @param tmpReq - UpdateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIJobResponse
func (client *Client) UpdateDIJobWithContext(ctx context.Context, tmpReq *UpdateDIJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateDIJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobSettings) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, dara.String("JobSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSettings) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, dara.String("ResourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TableMappings) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, dara.String("TableMappings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TransformationRules) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, dara.String("TransformationRules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DIJobId) {
		query["DIJobId"] = request.DIJobId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSpec) {
		body["FileSpec"] = request.FileSpec
	}

	if !dara.IsNil(request.JobSettingsShrink) {
		body["JobSettings"] = request.JobSettingsShrink
	}

	if !dara.IsNil(request.ResourceSettingsShrink) {
		body["ResourceSettings"] = request.ResourceSettingsShrink
	}

	if !dara.IsNil(request.TableMappingsShrink) {
		body["TableMappings"] = request.TableMappingsShrink
	}

	if !dara.IsNil(request.TransformationRulesShrink) {
		body["TransformationRules"] = request.TransformationRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDIJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a label.
//
// Description:
//
// You must purchase DataWorks Enterprise Edition or a higher edition to use this feature.
//
// @param tmpReq - UpdateDataAssetTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataAssetTagResponse
func (client *Client) UpdateDataAssetTagWithContext(ctx context.Context, tmpReq *UpdateDataAssetTagRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataAssetTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataAssetTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Managers) {
		request.ManagersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Managers, dara.String("Managers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Values) {
		request.ValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Values, dara.String("Values"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.ManagersShrink) {
		query["Managers"] = request.ManagersShrink
	}

	if !dara.IsNil(request.ValuesShrink) {
		query["Values"] = request.ValuesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataAssetTag"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataAssetTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified data quality monitoring alert rule.
//
// Description:
//
// DataWorks Basic Edition or a more advanced edition is required.
//
// @param tmpReq - UpdateDataQualityAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityAlertRuleResponse
func (client *Client) UpdateDataQualityAlertRuleWithContext(ctx context.Context, tmpReq *UpdateDataQualityAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityAlertRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Condition) {
		body["Condition"] = request.Condition
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.NotificationShrink) {
		body["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityAlertRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateDataQualityEvaluationTask is deprecated, please use dataworks-public::2024-05-18::UpdateDataQualityScan instead.
//
// Summary:
//
// Updates a data quality validation task.
//
// Description:
//
// DataWorks Basic Edition or above must be purchased to use this operation.
//
// @param tmpReq - UpdateDataQualityEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityEvaluationTaskResponse
func (client *Client) UpdateDataQualityEvaluationTaskWithContext(ctx context.Context, tmpReq *UpdateDataQualityEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityEvaluationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityEvaluationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataQualityRules) {
		request.DataQualityRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataQualityRules, dara.String("DataQualityRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hooks) {
		request.HooksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hooks, dara.String("Hooks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notifications) {
		request.NotificationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notifications, dara.String("Notifications"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityRulesShrink) {
		body["DataQualityRules"] = request.DataQualityRulesShrink
	}

	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HooksShrink) {
		body["Hooks"] = request.HooksShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationsShrink) {
		body["Notifications"] = request.NotificationsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeConf) {
		body["RuntimeConf"] = request.RuntimeConf
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityEvaluationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateDataQualityRule is deprecated, please use dataworks-public::2024-05-18::UpdateDataQualityScan instead.
//
// Summary:
//
// Updates a data quality rule.
//
// Description:
//
// You must purchase DataWorks Basic Edition or above to use this feature.
//
// @param tmpReq - UpdateDataQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityRuleResponse
func (client *Client) UpdateDataQualityRuleWithContext(ctx context.Context, tmpReq *UpdateDataQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckingConfig) {
		request.CheckingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckingConfig, dara.String("CheckingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ErrorHandlers) {
		request.ErrorHandlersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ErrorHandlers, dara.String("ErrorHandlers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SamplingConfig) {
		request.SamplingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SamplingConfig, dara.String("SamplingConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckingConfigShrink) {
		body["CheckingConfig"] = request.CheckingConfigShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ErrorHandlersShrink) {
		body["ErrorHandlers"] = request.ErrorHandlersShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SamplingConfigShrink) {
		body["SamplingConfig"] = request.SamplingConfigShrink
	}

	if !dara.IsNil(request.Severity) {
		body["Severity"] = request.Severity
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateDataQualityRuleTemplate is deprecated, please use dataworks-public::2024-05-18::UpdateDataQualityTemplate instead.
//
// Summary:
//
// Updates a data quality rule template.
//
// Description:
//
// You can call this operation only after you purchase DataWorks Basic Edition or a higher edition.
//
// @param tmpReq - UpdateDataQualityRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityRuleTemplateResponse
func (client *Client) UpdateDataQualityRuleTemplateWithContext(ctx context.Context, tmpReq *UpdateDataQualityRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityRuleTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityRuleTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckingConfig) {
		request.CheckingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckingConfig, dara.String("CheckingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SamplingConfig) {
		request.SamplingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SamplingConfig, dara.String("SamplingConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckingConfigShrink) {
		body["CheckingConfig"] = request.CheckingConfigShrink
	}

	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.DirectoryPath) {
		body["DirectoryPath"] = request.DirectoryPath
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SamplingConfigShrink) {
		body["SamplingConfig"] = request.SamplingConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityRuleTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityRuleTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data quality monitor.
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param tmpReq - UpdateDataQualityScanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityScanResponse
func (client *Client) UpdateDataQualityScanWithContext(ctx context.Context, tmpReq *UpdateDataQualityScanRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityScanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityScanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ComputeResource) {
		request.ComputeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComputeResource, dara.String("ComputeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hooks) {
		request.HooksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hooks, dara.String("Hooks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ComputeResourceShrink) {
		body["ComputeResource"] = request.ComputeResourceShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HooksShrink) {
		body["Hooks"] = request.HooksShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityScan"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityScanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Updates a data quality rule template in a project
//
// Description:
//
// DataWorks Basic Edition or a higher edition is required.
//
// @param request - UpdateDataQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityTemplateResponse
func (client *Client) UpdateDataQualityTemplateWithContext(ctx context.Context, request *UpdateDataQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a data source by ID.
//
// @param request - UpdateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSourceWithContext(ctx context.Context, request *UpdateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionProperties) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !dara.IsNil(request.ConnectionPropertiesMode) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates dataset information. Only DataWorks datasets are supported. The operator must be the creator of the dataset or the administrator of the workspace where the dataset is located.
//
// @param request - UpdateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDatasetWithContext(ctx context.Context, request *UpdateDatasetRequest, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Readme) {
		body["Readme"] = request.Readme
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataset"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates dataset version information. Only DataWorks datasets are supported. Requires dataset creator or workspace administrator permissions.
//
// @param request - UpdateDatasetVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetVersionResponse
func (client *Client) UpdateDatasetVersionWithContext(ctx context.Context, request *UpdateDatasetVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdateDatasetVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDatasetVersion"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a file.
//
// @param request - UpdateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileResponse
func (client *Client) UpdateFileWithContext(ctx context.Context, request *UpdateFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedSettings) {
		body["AdvancedSettings"] = request.AdvancedSettings
	}

	if !dara.IsNil(request.ApplyScheduleImmediately) {
		body["ApplyScheduleImmediately"] = request.ApplyScheduleImmediately
	}

	if !dara.IsNil(request.AutoParsing) {
		body["AutoParsing"] = request.AutoParsing
	}

	if !dara.IsNil(request.AutoRerunIntervalMillis) {
		body["AutoRerunIntervalMillis"] = request.AutoRerunIntervalMillis
	}

	if !dara.IsNil(request.AutoRerunTimes) {
		body["AutoRerunTimes"] = request.AutoRerunTimes
	}

	if !dara.IsNil(request.ConnectionName) {
		body["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.CronExpress) {
		body["CronExpress"] = request.CronExpress
	}

	if !dara.IsNil(request.CycleType) {
		body["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.DependentNodeIdList) {
		body["DependentNodeIdList"] = request.DependentNodeIdList
	}

	if !dara.IsNil(request.DependentType) {
		body["DependentType"] = request.DependentType
	}

	if !dara.IsNil(request.EndEffectDate) {
		body["EndEffectDate"] = request.EndEffectDate
	}

	if !dara.IsNil(request.FileDescription) {
		body["FileDescription"] = request.FileDescription
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.IgnoreParentSkipRunningProperty) {
		body["IgnoreParentSkipRunningProperty"] = request.IgnoreParentSkipRunningProperty
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InputList) {
		body["InputList"] = request.InputList
	}

	if !dara.IsNil(request.InputParameters) {
		body["InputParameters"] = request.InputParameters
	}

	if !dara.IsNil(request.OutputList) {
		body["OutputList"] = request.OutputList
	}

	if !dara.IsNil(request.OutputParameters) {
		body["OutputParameters"] = request.OutputParameters
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParaValue) {
		body["ParaValue"] = request.ParaValue
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.RerunMode) {
		body["RerunMode"] = request.RerunMode
	}

	if !dara.IsNil(request.ResourceGroupIdentifier) {
		body["ResourceGroupIdentifier"] = request.ResourceGroupIdentifier
	}

	if !dara.IsNil(request.SchedulerType) {
		body["SchedulerType"] = request.SchedulerType
	}

	if !dara.IsNil(request.StartEffectDate) {
		body["StartEffectDate"] = request.StartEffectDate
	}

	if !dara.IsNil(request.StartImmediately) {
		body["StartImmediately"] = request.StartImmediately
	}

	if !dara.IsNil(request.Stop) {
		body["Stop"] = request.Stop
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invoke UpdateFolder to update the folder information.
//
// @param request - UpdateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolderWithContext(ctx context.Context, request *UpdateFolderRequest, runtime *dara.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.FolderName) {
		body["FolderName"] = request.FolderName
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFolder"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a user-defined function (UDF) in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunctionWithContext(ctx context.Context, request *UpdateFunctionRequest, runtime *dara.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Recalls the check result of the message of an extension point event.
//
// @param request - UpdateIDEEventResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIDEEventResultResponse
func (client *Client) UpdateIDEEventResultWithContext(ctx context.Context, request *UpdateIDEEventResultRequest, runtime *dara.RuntimeOptions) (_result *UpdateIDEEventResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckResult) {
		body["CheckResult"] = request.CheckResult
	}

	if !dara.IsNil(request.CheckResultTip) {
		body["CheckResultTip"] = request.CheckResultTip
	}

	if !dara.IsNil(request.ExtensionCode) {
		body["ExtensionCode"] = request.ExtensionCode
	}

	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIDEEventResult"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIDEEventResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an MCP Server.
//
// Description:
//
// ## Operation description
//
// This API operation allows you to update the configuration of a specified MCP Server, including the service URL, transport protocol, custom request headers, and visibility settings. Fields that are not provided retain their existing values.
//
// **Note**: When you modify the `Visibility` parameter, selectively provide `ProjectIds` or `UserIds` in `VisibilityScope` based on the visibility setting to ensure that the correct access control scope is applied.
//
// @param tmpReq - UpdateMcpServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcpServerResponse
func (client *Client) UpdateMcpServerWithContext(ctx context.Context, tmpReq *UpdateMcpServerRequest, runtime *dara.RuntimeOptions) (_result *UpdateMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMcpServerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomHeaders) {
		request.CustomHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomHeaders, dara.String("CustomHeaders"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VisibilityScope) {
		request.VisibilityScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VisibilityScope, dara.String("VisibilityScope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomHeadersShrink) {
		body["CustomHeaders"] = request.CustomHeadersShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Transport) {
		body["Transport"] = request.Transport
	}

	if !dara.IsNil(request.Url) {
		body["Url"] = request.Url
	}

	if !dara.IsNil(request.Visibility) {
		body["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.VisibilityScopeShrink) {
		body["VisibilityScope"] = request.VisibilityScopeShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcpServer"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcpServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update Data Map collection objects, including Data Map categories and data albums. You can update the collection name, description, and administrator information.
//
// When updating a data album, the caller must have the AliyunDataWorksFullAccess permission or be the creator or an administrator of the album.
//
// Description:
//
// 1. You must purchase DataWorks Professional Edition or a higher edition to use this feature.
//
// @param tmpReq - UpdateMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaCollectionResponse
func (client *Client) UpdateMetaCollectionWithContext(ctx context.Context, tmpReq *UpdateMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetaCollectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMetaCollectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Administrators) {
		request.AdministratorsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Administrators, dara.String("Administrators"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdministratorsShrink) {
		query["Administrators"] = request.AdministratorsShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMetaCollectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a metadata entity. You can update custom entities or objects of the extended table type, such as databases, tables, and columns.
//
// Description:
//
// You must purchase DataWorks Professional Edition or a higher edition to use this operation.
//
// @param tmpReq - UpdateMetaEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaEntityResponse
func (client *Client) UpdateMetaEntityWithContext(ctx context.Context, tmpReq *UpdateMetaEntityRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetaEntityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMetaEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Attributes) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, dara.String("Attributes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CustomAttributes) {
		request.CustomAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomAttributes, dara.String("CustomAttributes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttributesShrink) {
		body["Attributes"] = request.AttributesShrink
	}

	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.CustomAttributesShrink) {
		body["CustomAttributes"] = request.CustomAttributesShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetaEntity"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMetaEntityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a metadata entity definition, including custom entity types and extension table types.
//
// Description:
//
// DataWorks Professional Edition or a more advanced edition is required.
//
// @param tmpReq - UpdateMetaEntityDefRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaEntityDefResponse
func (client *Client) UpdateMetaEntityDefWithContext(ctx context.Context, tmpReq *UpdateMetaEntityDefRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetaEntityDefResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMetaEntityDefShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NewAttributeDefs) {
		request.NewAttributeDefsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NewAttributeDefs, dara.String("NewAttributeDefs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UpdateAttributeDefs) {
		request.UpdateAttributeDefsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpdateAttributeDefs, dara.String("UpdateAttributeDefs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EntityType) {
		body["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.NewAttributeDefsShrink) {
		body["NewAttributeDefs"] = request.NewAttributeDefsShrink
	}

	if !dara.IsNil(request.UpdateAttributeDefsShrink) {
		body["UpdateAttributeDefs"] = request.UpdateAttributeDefsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetaEntityDef"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMetaEntityDefResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a node in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeResponse
func (client *Client) UpdateNodeWithContext(ctx context.Context, request *UpdateNodeRequest, runtime *dara.RuntimeOptions) (_result *UpdateNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a parameter. This operation performs an incremental update and modifies only the specified fields.
//
// Description:
//
// This operation is available only in DataWorks Professional Edition and later.
//
// @param tmpReq - UpdateParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateParameterResponse
func (client *Client) UpdateParameterWithContext(ctx context.Context, tmpReq *UpdateParameterRequest, runtime *dara.RuntimeOptions) (_result *UpdateParameterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Properties) {
		request.PropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Properties, dara.String("Properties"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PropertiesShrink) {
		body["Properties"] = request.PropertiesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateParameter"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateParameterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an existing approval process definition.
//
// Description:
//
// ## Request
//
// - Use this API to modify an existing approval process definition, including its name, description, rule conditions, notification service, and approval nodes.
//
// - The required `Id` parameter identifies the approval process definition to update.
//
// - To overwrite the existing configuration, set the `Overwrite` parameter to `true`.
//
// - The optional `ClientToken` parameter ensures request idempotency.
//
// @param tmpReq - UpdateProcessDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProcessDefinitionResponse
func (client *Client) UpdateProcessDefinitionWithContext(ctx context.Context, tmpReq *UpdateProcessDefinitionRequest, runtime *dara.RuntimeOptions) (_result *UpdateProcessDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProcessDefinitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApprovalNodes) {
		request.ApprovalNodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovalNodes, dara.String("ApprovalNodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NotificationServices) {
		request.NotificationServicesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationServices, dara.String("NotificationServices"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuleConditions) {
		request.RuleConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuleConditions, dara.String("RuleConditions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalNodesShrink) {
		body["ApprovalNodes"] = request.ApprovalNodesShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationServicesShrink) {
		body["NotificationServices"] = request.NotificationServicesShrink
	}

	if !dara.IsNil(request.RuleConditionsShrink) {
		body["RuleConditions"] = request.RuleConditionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProcessDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProcessDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation updates a DataWorks workspace.
//
// Description:
//
// This feature requires DataWorks Basic Edition or a later version.
//
// @param request - UpdateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithContext(ctx context.Context, request *UpdateProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
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

	if !dara.IsNil(request.DevEnvironmentEnabled) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !dara.IsNil(request.DevRoleDisabled) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.PaiTaskEnabled) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProject"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the permissions of a custom role in a workspace.
//
// @param tmpReq - UpdateProjectRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectRoleResponse
func (client *Client) UpdateProjectRoleWithContext(ctx context.Context, tmpReq *UpdateProjectRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProjectRoleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ModulePermissions) {
		request.ModulePermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ModulePermissions, dara.String("ModulePermissions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.ModulePermissionsShrink) {
		query["ModulePermissions"] = request.ModulePermissionsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProjectRole"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a file resource in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithContext(ctx context.Context, request *UpdateResourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceFile) {
		body["ResourceFile"] = request.ResourceFile
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates basic information about a resource group.
//
// Description:
//
// You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// @param request - UpdateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroupWithContext(ctx context.Context, request *UpdateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunResourceGroupId) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a route.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRouteResponse
func (client *Client) UpdateRouteWithContext(ctx context.Context, request *UpdateRouteRequest, runtime *dara.RuntimeOptions) (_result *UpdateRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidr) {
		body["DestinationCidr"] = request.DestinationCidr
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRoute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// To modify the configuration of an existing security policy, you must be both a DataWorks tenant administrator and a security administrator.
//
// Description:
//
// ## Usage
//
// - Use this API to update a specified security policy, including its name, description, associated workspace IDs, and policy content.
//
// - You cannot modify some properties of default system policies, such as the schema name and control module.
//
// - When `ControlDwScope` is set to `Workspace`, use the `Workspaces` parameter to associate the policy with specific workspaces.
//
// - When updating the policy content (`Content`), ensure that the provided controllers (`Controllers`) conform to the requirements of the selected schema.
//
// - The optional `ClientToken` parameter ensures request idempotence.
//
// @param tmpReq - UpdateSecurityStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecurityStrategyResponse
func (client *Client) UpdateSecurityStrategyWithContext(ctx context.Context, tmpReq *UpdateSecurityStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateSecurityStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSecurityStrategyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Workspaces) {
		request.WorkspacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Workspaces, dara.String("Workspaces"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ContentShrink) {
		body["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.WorkspacesShrink) {
		body["Workspaces"] = request.WorkspacesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecurityStrategy"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecurityStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified Skill and generates a new version.
//
// Description:
//
// ## Request description
//
// This API allows you to update an existing Skill and create a new version based on the current highest version or a specified version. Fields not provided in the request retain their original values. You can update the Skill content by providing either `SkillMdOverride` or `BundleUrl`. You can also set additional information such as the visibility scope.
//
// @param tmpReq - UpdateSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSkillResponse
func (client *Client) UpdateSkillWithContext(ctx context.Context, tmpReq *UpdateSkillRequest, runtime *dara.RuntimeOptions) (_result *UpdateSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSkillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Extra) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, dara.String("Extra"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VisibilityScope) {
		request.VisibilityScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VisibilityScope, dara.String("VisibilityScope"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BundleUrl) {
		body["BundleUrl"] = request.BundleUrl
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExpectedVersion) {
		body["ExpectedVersion"] = request.ExpectedVersion
	}

	if !dara.IsNil(request.ExtraShrink) {
		body["Extra"] = request.ExtraShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SkillMdOverride) {
		body["SkillMdOverride"] = request.SkillMdOverride
	}

	if !dara.IsNil(request.VersionNote) {
		body["VersionNote"] = request.VersionNote
	}

	if !dara.IsNil(request.VisibilityScopeShrink) {
		body["VisibilityScope"] = request.VisibilityScopeShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSkill"),
		Version:     dara.String("2024-05-18"),
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

// Summary:
//
// Updates the business metadata of a table in Data Map. Currently, only the table usage description and custom attributes can be updated.
//
// Description:
//
// 1. You must have DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - UpdateTableBusinessMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableBusinessMetadataResponse
func (client *Client) UpdateTableBusinessMetadataWithContext(ctx context.Context, tmpReq *UpdateTableBusinessMetadataRequest, runtime *dara.RuntimeOptions) (_result *UpdateTableBusinessMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTableBusinessMetadataShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomAttributes) {
		request.CustomAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomAttributes, dara.String("CustomAttributes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomAttributesShrink) {
		body["CustomAttributes"] = request.CustomAttributesShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Readme) {
		body["Readme"] = request.Readme
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTableBusinessMetadata"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableBusinessMetadataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified node. The modifications are synchronized to DataStudio, where a new saved version is created.
//
// @param tmpReq - UpdateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskResponse
func (client *Client) UpdateTaskWithContext(ctx context.Context, tmpReq *UpdateTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSource) {
		request.DataSourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSource, dara.String("DataSource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Dependencies) {
		request.DependenciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dependencies, dara.String("Dependencies"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Inputs) {
		request.InputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Inputs, dara.String("Inputs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Outputs) {
		request.OutputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Outputs, dara.String("Outputs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Script) {
		request.ScriptShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Script, dara.String("Script"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientUniqueCode) {
		body["ClientUniqueCode"] = request.ClientUniqueCode
	}

	if !dara.IsNil(request.DataSourceShrink) {
		body["DataSource"] = request.DataSourceShrink
	}

	if !dara.IsNil(request.DependenciesShrink) {
		body["Dependencies"] = request.DependenciesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InputsShrink) {
		body["Inputs"] = request.InputsShrink
	}

	if !dara.IsNil(request.InstanceMode) {
		body["InstanceMode"] = request.InstanceMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputsShrink) {
		body["Outputs"] = request.OutputsShrink
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.RerunInterval) {
		body["RerunInterval"] = request.RerunInterval
	}

	if !dara.IsNil(request.RerunMode) {
		body["RerunMode"] = request.RerunMode
	}

	if !dara.IsNil(request.RerunTimes) {
		body["RerunTimes"] = request.RerunTimes
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	if !dara.IsNil(request.ScriptShrink) {
		body["Script"] = request.ScriptShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the properties of multiple task instances in batch, including priority, resource group, data source, and more.
//
// Description:
//
// You must purchase DataWorks Basic Edition or a higher edition to use this feature.
//
// @param tmpReq - UpdateTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskInstancesResponse
func (client *Client) UpdateTaskInstancesWithContext(ctx context.Context, tmpReq *UpdateTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskInstances) {
		request.TaskInstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskInstances, dara.String("TaskInstances"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.TaskInstancesShrink) {
		body["TaskInstances"] = request.TaskInstancesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the file information about a function.
//
// @param request - UpdateUdfFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfFileResponse
func (client *Client) UpdateUdfFileWithContext(ctx context.Context, request *UpdateUdfFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateUdfFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		body["ClassName"] = request.ClassName
	}

	if !dara.IsNil(request.CmdDescription) {
		body["CmdDescription"] = request.CmdDescription
	}

	if !dara.IsNil(request.Example) {
		body["Example"] = request.Example
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FunctionType) {
		body["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.ParameterDescription) {
		body["ParameterDescription"] = request.ParameterDescription
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.Resources) {
		body["Resources"] = request.Resources
	}

	if !dara.IsNil(request.ReturnValue) {
		body["ReturnValue"] = request.ReturnValue
	}

	if !dara.IsNil(request.UdfDescription) {
		body["UdfDescription"] = request.UdfDescription
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUdfFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUdfFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified workflow by using the full update method. Fields that can be synchronously updated to DataStudio include: owner, data source, schedule resource group, description, and trigger run mode (Normal, Skip, or Pause).
//
// Description:
//
// DataWorks Basic Edition or a more advanced edition is required.
//
// @param tmpReq - UpdateWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflowWithContext(ctx context.Context, tmpReq *UpdateWorkflowRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkflowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Dependencies) {
		request.DependenciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dependencies, dara.String("Dependencies"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Outputs) {
		request.OutputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Outputs, dara.String("Outputs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tasks) {
		request.TasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tasks, dara.String("Tasks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientUniqueCode) {
		body["ClientUniqueCode"] = request.ClientUniqueCode
	}

	if !dara.IsNil(request.DependenciesShrink) {
		body["Dependencies"] = request.DependenciesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceMode) {
		body["InstanceMode"] = request.InstanceMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputsShrink) {
		body["Outputs"] = request.OutputsShrink
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Parameters) {
		body["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TasksShrink) {
		body["Tasks"] = request.TasksShrink
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflow"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a workflow in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// Description:
//
//	Notice:
//
// This API does not support batch operations. If you define more than one workflow definition in the FlowSpec, all workflow definitions except the first one are ignored. In addition, nodes defined within the workflow definition are also ignored. Call the UpdateNode API to update internal nodes one by one.
//
// @param request - UpdateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowDefinitionResponse
func (client *Client) UpdateWorkflowDefinitionWithContext(ctx context.Context, request *UpdateWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkflowDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Requests a temporary OSS PUT upload URL. Complete the PUT upload before the URL expires, then pass the returned FileId to the ReferenceFileIds parameter of CreateSemanticJob.
//
// Description:
//
// ## Scenarios
//
// Requests an upload slot for a reference file to prepare a file for the `singleTableFile` source of `CreateSemanticJob`.
//
// ## Procedure
//
// 1. Pass the file name, MIME type, and actual size to this operation to obtain `Data.UploadUrl` and `Data.FileId`.
//
// 2. Perform an HTTP PUT upload with the same `Content-Type` before the `UploadUrl` expires.
//
// 3. After the upload is complete, use `FileId` as the only element of `CreateSemanticJob.ReferenceFileIds`.
//
// ## Security considerations
//
// `UploadUrl` is a short-lived pre-signed PUT URL. The holder can write to the corresponding object. Do not log, share, or persist this URL.
//
// @param request - UploadSemanticFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadSemanticFileResponse
func (client *Client) UploadSemanticFileWithContext(ctx context.Context, request *UploadSemanticFileRequest, runtime *dara.RuntimeOptions) (_result *UploadSemanticFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		body["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.SizeBytes) {
		body["SizeBytes"] = request.SizeBytes
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadSemanticFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadSemanticFileResponse{}
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
		Version:     dara.String("2024-05-18"),
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
		Version:     dara.String("2024-05-18"),
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
