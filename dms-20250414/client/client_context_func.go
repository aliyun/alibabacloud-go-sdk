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
// Adds a user to a specified workspace.
//
// @param request - AddUserToDataAgentWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToDataAgentWorkspaceResponse
func (client *Client) AddUserToDataAgentWorkspaceWithContext(ctx context.Context, request *AddUserToDataAgentWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *AddUserToDataAgentWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddUserToDataAgentWorkspace"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddUserToDataAgentWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates partitions for a data lake table in a batch.
//
// @param tmpReq - BatchCreateDataLakePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateDataLakePartitionsResponse
func (client *Client) BatchCreateDataLakePartitionsWithContext(ctx context.Context, tmpReq *BatchCreateDataLakePartitionsRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateDataLakePartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchCreateDataLakePartitionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInputs) {
		request.PartitionInputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInputs, dara.String("PartitionInputs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfNotExists) {
		query["IfNotExists"] = request.IfNotExists
	}

	if !dara.IsNil(request.NeedResult) {
		query["NeedResult"] = request.NeedResult
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputsShrink) {
		body["PartitionInputs"] = request.PartitionInputsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateDataLakePartitions"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateDataLakePartitionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch deletes partitions from a data lake table.
//
// @param request - BatchDeleteDataLakePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDataLakePartitionsResponse
func (client *Client) BatchDeleteDataLakePartitionsWithContext(ctx context.Context, request *BatchDeleteDataLakePartitionsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDataLakePartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfExists) {
		query["IfExists"] = request.IfExists
	}

	if !dara.IsNil(request.PartitionValuesList) {
		query["PartitionValuesList"] = request.PartitionValuesList
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDataLakePartitions"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDataLakePartitionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update partitions in a data lake table in batch.
//
// @param tmpReq - BatchUpdateDataLakePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateDataLakePartitionsResponse
func (client *Client) BatchUpdateDataLakePartitionsWithContext(ctx context.Context, tmpReq *BatchUpdateDataLakePartitionsRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateDataLakePartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateDataLakePartitionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInputs) {
		request.PartitionInputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInputs, dara.String("PartitionInputs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputsShrink) {
		body["PartitionInputs"] = request.PartitionInputsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateDataLakePartitions"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateDataLakePartitionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update an Airflow instance\\"s custom configuration
//
// Description:
//
// Configure the airflow\\.cfg file for DMS Airflow.
//
// @param tmpReq - ConfigAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigAirflowResponse
func (client *Client) ConfigAirflowWithContext(ctx context.Context, tmpReq *ConfigAirflowRequest, runtime *dara.RuntimeOptions) (_result *ConfigAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ConfigAirflowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomAirflowCfg) {
		request.CustomAirflowCfgShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomAirflowCfg, dara.String("CustomAirflowCfg"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	if !dara.IsNil(request.CustomAirflowCfgShrink) {
		query["CustomAirflowCfg"] = request.CustomAirflowCfgShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigAirflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Airflow instance in a workspace.
//
// Description:
//
// Creates an Airflow instance in a workspace.
//
// @param tmpReq - CreateAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAirflowResponse
func (client *Client) CreateAirflowWithContext(ctx context.Context, tmpReq *CreateAirflowRequest, runtime *dara.RuntimeOptions) (_result *CreateAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAirflowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataMountInfoList) {
		request.DataMountInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataMountInfoList, dara.String("DataMountInfoList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowName) {
		query["AirflowName"] = request.AirflowName
	}

	if !dara.IsNil(request.AirflowVersion) {
		query["AirflowVersion"] = request.AirflowVersion
	}

	if !dara.IsNil(request.AppSpec) {
		query["AppSpec"] = request.AppSpec
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DagsDir) {
		query["DagsDir"] = request.DagsDir
	}

	if !dara.IsNil(request.DataMountInfoListShrink) {
		query["DataMountInfoList"] = request.DataMountInfoListShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableServerless) {
		query["EnableServerless"] = request.EnableServerless
	}

	if !dara.IsNil(request.GracefulShutdownTimeout) {
		query["GracefulShutdownTimeout"] = request.GracefulShutdownTimeout
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssPath) {
		query["OssPath"] = request.OssPath
	}

	if !dara.IsNil(request.PluginsDir) {
		query["PluginsDir"] = request.PluginsDir
	}

	if !dara.IsNil(request.RequirementFile) {
		query["RequirementFile"] = request.RequirementFile
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.StartupFile) {
		query["StartupFile"] = request.StartupFile
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WorkerServerlessReplicas) {
		query["WorkerServerlessReplicas"] = request.WorkerServerlessReplicas
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAirflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains logon credentials for an Airflow instance hosted by Data Management Service (DMS). Use the returned token and host endpoint to construct a logon URL for the Airflow web UI.
//
// @param request - CreateAirflowLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAirflowLoginTokenResponse
func (client *Client) CreateAirflowLoginTokenWithContext(ctx context.Context, request *CreateAirflowLoginTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateAirflowLoginTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAirflowLoginToken"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAirflowLoginTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a custom agent in your personal space or a workspace.
//
// @param tmpReq - CreateCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomAgentResponse
func (client *Client) CreateCustomAgentWithContext(ctx context.Context, tmpReq *CreateCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateCustomAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallbackConfig) {
		request.CallbackConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallbackConfig, dara.String("CallbackConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExecutionConfig) {
		request.ExecutionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutionConfig, dara.String("ExecutionConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KnowledgeConfigList) {
		request.KnowledgeConfigListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KnowledgeConfigList, dara.String("KnowledgeConfigList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleTaskConfig) {
		request.ScheduleTaskConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleTaskConfig, dara.String("ScheduleTaskConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackConfigShrink) {
		query["CallbackConfig"] = request.CallbackConfigShrink
	}

	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.DataJson) {
		query["DataJson"] = request.DataJson
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionConfigShrink) {
		query["ExecutionConfig"] = request.ExecutionConfigShrink
	}

	if !dara.IsNil(request.Instruction) {
		query["Instruction"] = request.Instruction
	}

	if !dara.IsNil(request.Knowledge) {
		query["Knowledge"] = request.Knowledge
	}

	if !dara.IsNil(request.KnowledgeConfigListShrink) {
		query["KnowledgeConfigList"] = request.KnowledgeConfigListShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RelatedSessionId) {
		query["RelatedSessionId"] = request.RelatedSessionId
	}

	if !dara.IsNil(request.ScheduleTaskConfigShrink) {
		query["ScheduleTaskConfig"] = request.ScheduleTaskConfigShrink
	}

	if !dara.IsNil(request.TextReportConfig) {
		query["TextReportConfig"] = request.TextReportConfig
	}

	if !dara.IsNil(request.WebReportConfig) {
		query["WebReportConfig"] = request.WebReportConfig
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomAgent"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a DataAgent knowledge base. The knowledge base creator has read and write permissions. Other workspace members have permission to use it.
//
// @param request - CreateDataAgentKnowledgeBaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataAgentKnowledgeBaseResponse
func (client *Client) CreateDataAgentKnowledgeBaseWithContext(ctx context.Context, request *CreateDataAgentKnowledgeBaseRequest, runtime *dara.RuntimeOptions) (_result *CreateDataAgentKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FromKbUuid) {
		query["FromKbUuid"] = request.FromKbUuid
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataAgentKnowledgeBase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataAgentKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a DataAgent session
//
// @param tmpReq - CreateDataAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataAgentSessionResponse
func (client *Client) CreateDataAgentSessionWithContext(ctx context.Context, tmpReq *CreateDataAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *CreateDataAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SessionConfig) {
		request.SessionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionConfig, dara.String("SessionConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.File) {
		query["File"] = request.File
	}

	if !dara.IsNil(request.SessionConfigShrink) {
		query["SessionConfig"] = request.SessionConfigShrink
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataAgentSession"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a DataAgent collaborative workspace.
//
// @param request - CreateDataAgentWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataAgentWorkspaceResponse
func (client *Client) CreateDataAgentWorkspaceWithContext(ctx context.Context, request *CreateDataAgentWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataAgentWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.IsSessionShareEnabled) {
		query["IsSessionShareEnabled"] = request.IsSessionShareEnabled
	}

	if !dara.IsNil(request.WorkspaceDesc) {
		query["WorkspaceDesc"] = request.WorkspaceDesc
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataAgentWorkspace"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataAgentWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data lake database.
//
// @param tmpReq - CreateDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakeDatabaseResponse
func (client *Client) CreateDataLakeDatabaseWithContext(ctx context.Context, tmpReq *CreateDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakeDatabaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakeDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user-defined function (UDF) for a data lake.
//
// @param tmpReq - CreateDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakeFunctionResponse
func (client *Client) CreateDataLakeFunctionWithContext(ctx context.Context, tmpReq *CreateDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakeFunctionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FunctionInput) {
		request.FunctionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FunctionInput, dara.String("FunctionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FunctionInputShrink) {
		body["FunctionInput"] = request.FunctionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakeFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a partition for a data lake table.
//
// @param tmpReq - CreateDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakePartitionResponse
func (client *Client) CreateDataLakePartitionWithContext(ctx context.Context, tmpReq *CreateDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInput) {
		request.PartitionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInput, dara.String("PartitionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfNotExists) {
		query["IfNotExists"] = request.IfNotExists
	}

	if !dara.IsNil(request.NeedResult) {
		query["NeedResult"] = request.NeedResult
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputShrink) {
		body["PartitionInput"] = request.PartitionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakePartitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data lake table in Data Management (DMS).
//
// @param tmpReq - CreateDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakeTableResponse
func (client *Client) CreateDataLakeTableWithContext(ctx context.Context, tmpReq *CreateDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakeTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableInput) {
		request.TableInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableInput, dara.String("TableInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableInputShrink) {
		body["TableInput"] = request.TableInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakeTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Airflow instance.
//
// @param request - DeleteAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAirflowResponse
func (client *Client) DeleteAirflowWithContext(ctx context.Context, request *DeleteAirflowRequest, runtime *dara.RuntimeOptions) (_result *DeleteAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAirflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a custom agent in your personal workspace or a shared workspace. Note: Only custom agents that are newly created or offline can be deleted.
//
// @param request - DeleteCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomAgentResponse
func (client *Client) DeleteCustomAgentWithContext(ctx context.Context, request *DeleteCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomAgentId) {
		query["CustomAgentId"] = request.CustomAgentId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomAgent"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a DataAgent knowledge base.
//
// @param request - DeleteDataAgentKnowledgeBaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataAgentKnowledgeBaseResponse
func (client *Client) DeleteDataAgentKnowledgeBaseWithContext(ctx context.Context, request *DeleteDataAgentKnowledgeBaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataAgentKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.KbUuid) {
		query["KbUuid"] = request.KbUuid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataAgentKnowledgeBase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataAgentKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a DataAgent workspace.
//
// @param request - DeleteDataAgentWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataAgentWorkspaceResponse
func (client *Client) DeleteDataAgentWorkspaceWithContext(ctx context.Context, request *DeleteDataAgentWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataAgentWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataAgentWorkspace"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataAgentWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data lake database.
//
// @param request - DeleteDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakeDatabaseResponse
func (client *Client) DeleteDataLakeDatabaseWithContext(ctx context.Context, request *DeleteDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakeDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function in a data lake.
//
// @param request - DeleteDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakeFunctionResponse
func (client *Client) DeleteDataLakeFunctionWithContext(ctx context.Context, request *DeleteDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakeFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a partition from a data lake table.
//
// @param tmpReq - DeleteDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakePartitionResponse
func (client *Client) DeleteDataLakePartitionWithContext(ctx context.Context, tmpReq *DeleteDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionValues) {
		request.PartitionValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionValues, dara.String("PartitionValues"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfExists) {
		query["IfExists"] = request.IfExists
	}

	if !dara.IsNil(request.PartitionValuesShrink) {
		query["PartitionValues"] = request.PartitionValuesShrink
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakePartitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data lake table.
//
// @param request - DeleteDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakeTableResponse
func (client *Client) DeleteDataLakeTableWithContext(ctx context.Context, request *DeleteDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakeTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a document from a specified knowledge base.
//
// @param request - DeleteDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocumentWithContext(ctx context.Context, request *DeleteDocumentRequest, runtime *dara.RuntimeOptions) (_result *DeleteDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentName) {
		body["DocumentName"] = request.DocumentName
	}

	if !dara.IsNil(request.KbUuid) {
		body["KbUuid"] = request.KbUuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDocument"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes chunks from a document in a knowledge base.
//
// @param tmpReq - DeleteDocumentChunksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentChunksResponse
func (client *Client) DeleteDocumentChunksWithContext(ctx context.Context, tmpReq *DeleteDocumentChunksRequest, runtime *dara.RuntimeOptions) (_result *DeleteDocumentChunksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDocumentChunksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ChunkIds) {
		request.ChunkIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChunkIds, dara.String("ChunkIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChunkIdsShrink) {
		body["ChunkIds"] = request.ChunkIdsShrink
	}

	if !dara.IsNil(request.DocumentName) {
		body["DocumentName"] = request.DocumentName
	}

	if !dara.IsNil(request.KbUuid) {
		body["KbUuid"] = request.KbUuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDocumentChunks"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocumentChunksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteFileUpload
//
// @param request - DeleteFileUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileUploadResponse
func (client *Client) DeleteFileUploadWithContext(ctx context.Context, request *DeleteFileUploadRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileUploadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallFrom) {
		query["CallFrom"] = request.CallFrom
	}

	if !dara.IsNil(request.DmsUnit) {
		query["DmsUnit"] = request.DmsUnit
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFileUpload"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileUploadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a code file or directory from a workspace.
//
// Description:
//
// This operation permanently removes a specified code file or directory.
//
// @param request - DeleteWorkspaceCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceCodeResponse
func (client *Client) DeleteWorkspaceCodeWithContext(ctx context.Context, request *DeleteWorkspaceCodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Repo) {
		query["Repo"] = request.Repo
	}

	if !dara.IsNil(request.Symlink) {
		query["Symlink"] = request.Symlink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkspaceCode"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// `DescribeCustomAgent` retrieves the details of a custom agent by its agent ID.
//
// @param request - DescribeCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomAgentResponse
func (client *Client) DescribeCustomAgentWithContext(ctx context.Context, request *DescribeCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomAgentId) {
		query["CustomAgentId"] = request.CustomAgentId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomAgent"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the description of a DataAgent session.
//
// @param request - DescribeDataAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataAgentSessionResponse
func (client *Client) DescribeDataAgentSessionWithContext(ctx context.Context, request *DescribeDataAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataAgentSession"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific document.
//
// @param request - DescribeDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocumentResponse
func (client *Client) DescribeDocumentWithContext(ctx context.Context, request *DescribeDocumentRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentName) {
		body["DocumentName"] = request.DocumentName
	}

	if !dara.IsNil(request.KbUuid) {
		body["KbUuid"] = request.KbUuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDocument"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeFileUploadSignature
//
// @param request - DescribeFileUploadSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileUploadSignatureResponse
func (client *Client) DescribeFileUploadSignatureWithContext(ctx context.Context, request *DescribeFileUploadSignatureRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileUploadSignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallFrom) {
		query["CallFrom"] = request.CallFrom
	}

	if !dara.IsNil(request.DmsUnit) {
		query["DmsUnit"] = request.DmsUnit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFileUploadSignature"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFileUploadSignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the statistics of a knowledge base.
//
// @param request - DescribeKnowledgeBaseStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKnowledgeBaseStatsResponse
func (client *Client) DescribeKnowledgeBaseStatsWithContext(ctx context.Context, request *DescribeKnowledgeBaseStatsRequest, runtime *dara.RuntimeOptions) (_result *DescribeKnowledgeBaseStatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KbUuid) {
		query["KbUuid"] = request.KbUuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKnowledgeBaseStats"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKnowledgeBaseStatsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a signature to upload a document to a knowledge base.
//
// @param request - DescribeKnowledgeBaseUploadSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKnowledgeBaseUploadSignatureResponse
func (client *Client) DescribeKnowledgeBaseUploadSignatureWithContext(ctx context.Context, request *DescribeKnowledgeBaseUploadSignatureRequest, runtime *dara.RuntimeOptions) (_result *DescribeKnowledgeBaseUploadSignatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KbUuid) {
		query["KbUuid"] = request.KbUuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeKnowledgeBaseUploadSignature"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKnowledgeBaseUploadSignatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # FileUploadCallback
//
// @param request - FileUploadCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileUploadCallbackResponse
func (client *Client) FileUploadCallbackWithContext(ctx context.Context, request *FileUploadCallbackRequest, runtime *dara.RuntimeOptions) (_result *FileUploadCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallFrom) {
		query["CallFrom"] = request.CallFrom
	}

	if !dara.IsNil(request.DmsUnit) {
		query["DmsUnit"] = request.DmsUnit
	}

	if !dara.IsNil(request.FileSize) {
		query["FileSize"] = request.FileSize
	}

	if !dara.IsNil(request.Filename) {
		query["Filename"] = request.Filename
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.UploadLocation) {
		query["UploadLocation"] = request.UploadLocation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FileUploadCallback"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FileUploadCallbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration and status of an Airflow instance.
//
// @param request - GetAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAirflowResponse
func (client *Client) GetAirflowWithContext(ctx context.Context, request *GetAirflowRequest, runtime *dara.RuntimeOptions) (_result *GetAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAirflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves chat content from a specific checkpoint by specifying the session ID and AgentId.
//
// Description:
//
// ## Request Description
//
// - The response is returned as an SSE stream, where each event follows the `SSEEvent` schema and contains meta-information such as the message level.
//
// - The `content` field in each SSE event may carry actual message text or a JSON object, depending on the value of `content_type`.
//
// @param request - GetChatContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatContentResponse
func (client *Client) GetChatContentWithSSECtx(ctx context.Context, request *GetChatContentRequest, runtime *dara.RuntimeOptions, _yield chan *GetChatContentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.getChatContentWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// Retrieves chat content from a specific checkpoint by specifying the session ID and AgentId.
//
// Description:
//
// ## Request Description
//
// - The response is returned as an SSE stream, where each event follows the `SSEEvent` schema and contains meta-information such as the message level.
//
// - The `content` field in each SSE event may carry actual message text or a JSON object, depending on the value of `content_type`.
//
// @param request - GetChatContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatContentResponse
func (client *Client) GetChatContentWithContext(ctx context.Context, request *GetChatContentRequest, runtime *dara.RuntimeOptions) (_result *GetChatContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.Checkpoint) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatContent"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a RAM user that belongs to an Alibaba Cloud account.
//
// @param request - GetDataAgentSubAccountInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataAgentSubAccountInfoResponse
func (client *Client) GetDataAgentSubAccountInfoWithContext(ctx context.Context, request *GetDataAgentSubAccountInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDataAgentSubAccountInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DmsUnit) {
		query["DmsUnit"] = request.DmsUnit
	}

	if !dara.IsNil(request.SubAccountId) {
		query["SubAccountId"] = request.SubAccountId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataAgentSubAccountInfo"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataAgentSubAccountInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a collaborative workspace.
//
// @param request - GetDataAgentWorkspaceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataAgentWorkspaceInfoResponse
func (client *Client) GetDataAgentWorkspaceInfoWithContext(ctx context.Context, request *GetDataAgentWorkspaceInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDataAgentWorkspaceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataAgentWorkspaceInfo"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataAgentWorkspaceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the data catalog of a data lake.
//
// @param request - GetDataLakeCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeCatalogResponse
func (client *Client) GetDataLakeCatalogWithContext(ctx context.Context, request *GetDataLakeCatalogRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeCatalog"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a database in a data lake.
//
// @param request - GetDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeDatabaseResponse
func (client *Client) GetDataLakeDatabaseWithContext(ctx context.Context, request *GetDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a user-defined function in a data lake.
//
// @param request - GetDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeFunctionResponse
func (client *Client) GetDataLakeFunctionWithContext(ctx context.Context, request *GetDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a data lakehouse table partition.
//
// @param tmpReq - GetDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakePartitionResponse
func (client *Client) GetDataLakePartitionWithContext(ctx context.Context, tmpReq *GetDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionValues) {
		request.PartitionValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionValues, dara.String("PartitionValues"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.PartitionValuesShrink) {
		query["PartitionValues"] = request.PartitionValuesShrink
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakePartitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a table in a data lake.
//
// @param request - GetDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeTableResponse
func (client *Client) GetDataLakeTableWithContext(ctx context.Context, request *GetDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a task to schedule and run a Notebook file.
//
// @param request - GetNotebookAndSubmitTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotebookAndSubmitTaskResponse
func (client *Client) GetNotebookAndSubmitTaskWithContext(ctx context.Context, request *GetNotebookAndSubmitTaskRequest, runtime *dara.RuntimeOptions) (_result *GetNotebookAndSubmitTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		body["Params"] = request.Params
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.Retry) {
		body["Retry"] = request.Retry
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotebookAndSubmitTask"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotebookAndSubmitTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the execution status of a Notebook task.
//
// @param request - GetNotebookTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNotebookTaskStatusResponse
func (client *Client) GetNotebookTaskStatusWithContext(ctx context.Context, request *GetNotebookTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetNotebookTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNotebookTaskStatus"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNotebookTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs of the SQL window.
//
// @param request - GetSqlConsoleOperationLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSqlConsoleOperationLogResponse
func (client *Client) GetSqlConsoleOperationLogWithContext(ctx context.Context, request *GetSqlConsoleOperationLogRequest, runtime *dara.RuntimeOptions) (_result *GetSqlConsoleOperationLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.SqlType) {
		query["SqlType"] = request.SqlType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSqlConsoleOperationLog"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSqlConsoleOperationLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reads the content of a code file in the workspace and returns the file content along with mtime (in the header).
//
// Description:
//
// Obtains the resource configuration limit information and the instance purchase status of the workspace.
//
// @param request - GetWorkspaceCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceCodeResponse
func (client *Client) GetWorkspaceCodeWithContext(ctx context.Context, request *GetWorkspaceCodeRequest, runtime *dara.RuntimeOptions) (_result *GetWorkspaceCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Iac) {
		query["Iac"] = request.Iac
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspaceCode"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the deployment configuration of a workspace.
//
// Description:
//
// This operation retrieves the deployment configuration of a workspace. The configuration includes details such as repository and branch information, and directories to exclude.
//
// @param request - GetWorkspaceCodePublishSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceCodePublishSettingResponse
func (client *Client) GetWorkspaceCodePublishSettingWithContext(ctx context.Context, request *GetWorkspaceCodePublishSettingRequest, runtime *dara.RuntimeOptions) (_result *GetWorkspaceCodePublishSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspaceCodePublishSetting"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceCodePublishSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the quota for a workspace.
//
// Description:
//
// Retrieves the resource quotas and instance status for a workspace.
//
// @param request - GetWorkspaceQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceQuotaResponse
func (client *Client) GetWorkspaceQuotaWithContext(ctx context.Context, request *GetWorkspaceQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetWorkspaceQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspaceQuota"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the available Airflow versions.
//
// Description:
//
// Lists the available Airflow versions.
//
// @param request - ListAirflowVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAirflowVersionsResponse
func (client *Client) ListAirflowVersionsWithContext(ctx context.Context, request *ListAirflowVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListAirflowVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAirflowVersions"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAirflowVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of Airflow instances in a workspace.
//
// @param request - ListAirflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAirflowsResponse
func (client *Client) ListAirflowsWithContext(ctx context.Context, request *ListAirflowsRequest, runtime *dara.RuntimeOptions) (_result *ListAirflowsResponse, _err error) {
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

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAirflows"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAirflowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ListCustomAgent returns a list of all custom agents from the personal space and workspaces.
//
// @param request - ListCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomAgentResponse
func (client *Client) ListCustomAgentWithContext(ctx context.Context, request *ListCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *ListCustomAgentResponse, _err error) {
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

	if !dara.IsNil(request.QueryAllReleased) {
		query["QueryAllReleased"] = request.QueryAllReleased
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomAgent"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of historical session descriptions for a Data Agent.
//
// @param request - ListDataAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataAgentSessionResponse
func (client *Client) ListDataAgentSessionWithContext(ctx context.Context, request *ListDataAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *ListDataAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateEndTime) {
		query["CreateEndTime"] = request.CreateEndTime
	}

	if !dara.IsNil(request.CreateStartTime) {
		query["CreateStartTime"] = request.CreateStartTime
	}

	if !dara.IsNil(request.CustomAgentId) {
		query["CustomAgentId"] = request.CustomAgentId
	}

	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.IsSaved) {
		query["IsSaved"] = request.IsSaved
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataAgentSession"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataAgentSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the collaborative workspaces under the primary account with pagination.
//
// @param request - ListDataAgentWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataAgentWorkspaceResponse
func (client *Client) ListDataAgentWorkspaceWithContext(ctx context.Context, request *ListDataAgentWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *ListDataAgentWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	if !dara.IsNil(request.WorkspaceType) {
		query["WorkspaceType"] = request.WorkspaceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataAgentWorkspace"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataAgentWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all members in a workspace.
//
// @param request - ListDataAgentWorkspaceMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataAgentWorkspaceMemberResponse
func (client *Client) ListDataAgentWorkspaceMemberWithContext(ctx context.Context, request *ListDataAgentWorkspaceMemberRequest, runtime *dara.RuntimeOptions) (_result *ListDataAgentWorkspaceMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchMemberId) {
		query["SearchMemberId"] = request.SearchMemberId
	}

	if !dara.IsNil(request.SearchRoleName) {
		query["SearchRoleName"] = request.SearchRoleName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataAgentWorkspaceMember"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataAgentWorkspaceMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists user-uploaded files in a data center, excluding databases.
//
// @param request - ListDataCenterDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataCenterDatabaseResponse
func (client *Client) ListDataCenterDatabaseWithContext(ctx context.Context, request *ListDataCenterDatabaseRequest, runtime *dara.RuntimeOptions) (_result *ListDataCenterDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallFrom) {
		query["CallFrom"] = request.CallFrom
	}

	if !dara.IsNil(request.DmsUnit) {
		query["DmsUnit"] = request.DmsUnit
	}

	if !dara.IsNil(request.ImportType) {
		query["ImportType"] = request.ImportType
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataCenterDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataCenterDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of user-uploaded tables from the data center, for file types only.
//
// @param request - ListDataCenterTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataCenterTableResponse
func (client *Client) ListDataCenterTableWithContext(ctx context.Context, request *ListDataCenterTableRequest, runtime *dara.RuntimeOptions) (_result *ListDataCenterTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallFrom) {
		query["CallFrom"] = request.CallFrom
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DmsUnit) {
		query["DmsUnit"] = request.DmsUnit
	}

	if !dara.IsNil(request.ImportType) {
		query["ImportType"] = request.ImportType
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataCenterTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataCenterTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data lake catalogs.
//
// @param request - ListDataLakeCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeCatalogResponse
func (client *Client) ListDataLakeCatalogWithContext(ctx context.Context, request *ListDataLakeCatalogRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeCatalog"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of databases in a data lake.
//
// @param request - ListDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeDatabaseResponse
func (client *Client) ListDataLakeDatabaseWithContext(ctx context.Context, request *ListDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data lake functions.
//
// @param request - ListDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeFunctionResponse
func (client *Client) ListDataLakeFunctionWithContext(ctx context.Context, request *ListDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionNamePattern) {
		query["FunctionNamePattern"] = request.FunctionNamePattern
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data lake function names.
//
// @param request - ListDataLakeFunctionNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeFunctionNameResponse
func (client *Client) ListDataLakeFunctionNameWithContext(ctx context.Context, request *ListDataLakeFunctionNameRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeFunctionNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionNamePattern) {
		query["FunctionNamePattern"] = request.FunctionNamePattern
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeFunctionName"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeFunctionNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of partitions of a data lake table.
//
// @param tmpReq - ListDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakePartitionResponse
func (client *Client) ListDataLakePartitionWithContext(ctx context.Context, tmpReq *ListDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartNames) {
		request.PartNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartNames, dara.String("PartNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartNamesShrink) {
		body["PartNames"] = request.PartNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakePartitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of partitions for a data lake table based on filter conditions.
//
// @param request - ListDataLakePartitionByFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakePartitionByFilterResponse
func (client *Client) ListDataLakePartitionByFilterWithContext(ctx context.Context, request *ListDataLakePartitionByFilterRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakePartitionByFilterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakePartitionByFilter"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakePartitionByFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of partition names for a data lake table.
//
// @param request - ListDataLakePartitionNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakePartitionNameResponse
func (client *Client) ListDataLakePartitionNameWithContext(ctx context.Context, request *ListDataLakePartitionNameRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakePartitionNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakePartitionName"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakePartitionNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data lake tables.
//
// @param request - ListDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeTableResponse
func (client *Client) ListDataLakeTableWithContext(ctx context.Context, request *ListDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableNamePattern) {
		query["TableNamePattern"] = request.TableNamePattern
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of data lake table names.
//
// @param request - ListDataLakeTableNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeTableNameResponse
func (client *Client) ListDataLakeTableNameWithContext(ctx context.Context, request *ListDataLakeTableNameRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeTableNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableNamePattern) {
		query["TableNamePattern"] = request.TableNamePattern
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeTableName"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeTableNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the basic information about tables in a data lake.
//
// @param request - ListDataLakeTablebaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeTablebaseInfoResponse
func (client *Client) ListDataLakeTablebaseInfoWithContext(ctx context.Context, request *ListDataLakeTablebaseInfoRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeTablebaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.Rows) {
		query["Rows"] = request.Rows
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeTablebaseInfo"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeTablebaseInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of chunks.
//
// @param request - ListDocumentChunksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDocumentChunksResponse
func (client *Client) ListDocumentChunksWithContext(ctx context.Context, request *ListDocumentChunksRequest, runtime *dara.RuntimeOptions) (_result *ListDocumentChunksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChunkTitlePattern) {
		body["ChunkTitlePattern"] = request.ChunkTitlePattern
	}

	if !dara.IsNil(request.DocumentName) {
		body["DocumentName"] = request.DocumentName
	}

	if !dara.IsNil(request.KbUuid) {
		body["KbUuid"] = request.KbUuid
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SortFieldName) {
		body["SortFieldName"] = request.SortFieldName
	}

	if !dara.IsNil(request.SortOrder) {
		body["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDocumentChunks"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDocumentChunksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the documents in a knowledge base.
//
// @param request - ListDocumentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDocumentsResponse
func (client *Client) ListDocumentsWithContext(ctx context.Context, request *ListDocumentsRequest, runtime *dara.RuntimeOptions) (_result *ListDocumentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		body["Filters"] = request.Filters
	}

	if !dara.IsNil(request.KbUuid) {
		body["KbUuid"] = request.KbUuid
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NamePattern) {
		body["NamePattern"] = request.NamePattern
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SortFieldName) {
		body["SortFieldName"] = request.SortFieldName
	}

	if !dara.IsNil(request.SortOrder) {
		body["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDocuments"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDocumentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListFileUpload
//
// @param request - ListFileUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileUploadResponse
func (client *Client) ListFileUploadWithContext(ctx context.Context, request *ListFileUploadRequest, runtime *dara.RuntimeOptions) (_result *ListFileUploadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallFrom) {
		query["CallFrom"] = request.CallFrom
	}

	if !dara.IsNil(request.DmsUnit) {
		query["DmsUnit"] = request.DmsUnit
	}

	if !dara.IsNil(request.DownloadLinkExpire) {
		query["DownloadLinkExpire"] = request.DownloadLinkExpire
	}

	if !dara.IsNil(request.FileCategory) {
		query["FileCategory"] = request.FileCategory
	}

	if !dara.IsNil(request.FileFrom) {
		query["FileFrom"] = request.FileFrom
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SortColumn) {
		query["SortColumn"] = request.SortColumn
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFileUpload"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFileUploadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a paginated list of knowledge bases.
//
// @param request - ListKnowledgeBasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKnowledgeBasesResponse
func (client *Client) ListKnowledgeBasesWithContext(ctx context.Context, request *ListKnowledgeBasesRequest, runtime *dara.RuntimeOptions) (_result *ListKnowledgeBasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		body["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NamePattern) {
		body["NamePattern"] = request.NamePattern
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SortFieldName) {
		body["SortFieldName"] = request.SortFieldName
	}

	if !dara.IsNil(request.SortOrder) {
		body["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKnowledgeBases"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKnowledgeBasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists code files and directories at a specified path in a workspace.
//
// Description:
//
// This operation lists the code files and directories at a specified path in a workspace.
//
// @param request - ListWorkspaceCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceCodeResponse
func (client *Client) ListWorkspaceCodeWithContext(ctx context.Context, request *ListWorkspaceCodeRequest, runtime *dara.RuntimeOptions) (_result *ListWorkspaceCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaceCode"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspaceCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify a custom agent in a personal space or workspace
//
// @param tmpReq - ModifyCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomAgentResponse
func (client *Client) ModifyCustomAgentWithContext(ctx context.Context, tmpReq *ModifyCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyCustomAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallbackConfig) {
		request.CallbackConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallbackConfig, dara.String("CallbackConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ExecutionConfig) {
		request.ExecutionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExecutionConfig, dara.String("ExecutionConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.KnowledgeConfigList) {
		request.KnowledgeConfigListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.KnowledgeConfigList, dara.String("KnowledgeConfigList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleTaskConfig) {
		request.ScheduleTaskConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleTaskConfig, dara.String("ScheduleTaskConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallbackConfigShrink) {
		query["CallbackConfig"] = request.CallbackConfigShrink
	}

	if !dara.IsNil(request.CustomAgentId) {
		query["CustomAgentId"] = request.CustomAgentId
	}

	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.DataJson) {
		query["DataJson"] = request.DataJson
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionConfigShrink) {
		query["ExecutionConfig"] = request.ExecutionConfigShrink
	}

	if !dara.IsNil(request.Instruction) {
		query["Instruction"] = request.Instruction
	}

	if !dara.IsNil(request.Knowledge) {
		query["Knowledge"] = request.Knowledge
	}

	if !dara.IsNil(request.KnowledgeConfigListShrink) {
		query["KnowledgeConfigList"] = request.KnowledgeConfigListShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RelatedSessionId) {
		query["RelatedSessionId"] = request.RelatedSessionId
	}

	if !dara.IsNil(request.ScheduleTaskConfigShrink) {
		query["ScheduleTaskConfig"] = request.ScheduleTaskConfigShrink
	}

	if !dara.IsNil(request.TextReportConfig) {
		query["TextReportConfig"] = request.TextReportConfig
	}

	if !dara.IsNil(request.WebReportConfig) {
		query["WebReportConfig"] = request.WebReportConfig
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomAgent"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Operate custom agents in personal spaces and workspaces.
//
// @param request - OperateCustomAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateCustomAgentResponse
func (client *Client) OperateCustomAgentWithContext(ctx context.Context, request *OperateCustomAgentRequest, runtime *dara.RuntimeOptions) (_result *OperateCustomAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomAgentId) {
		query["CustomAgentId"] = request.CustomAgentId
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateCustomAgent"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateCustomAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Redeploy an Airflow instance
//
// Description:
//
// Redeploys an Airflow instance.
//
// @param request - RedeployAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RedeployAirflowResponse
func (client *Client) RedeployAirflowWithContext(ctx context.Context, request *RedeployAirflowRequest, runtime *dara.RuntimeOptions) (_result *RedeployAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RedeployAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RedeployAirflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a user from a workspace.
//
// @param request - RemoveUserToDataAgentWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserToDataAgentWorkspaceResponse
func (client *Client) RemoveUserToDataAgentWorkspaceWithContext(ctx context.Context, request *RemoveUserToDataAgentWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserToDataAgentWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserToDataAgentWorkspace"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserToDataAgentWorkspaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query a knowledge base
//
// @param request - RetrieveKnowledgeBaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveKnowledgeBaseResponse
func (client *Client) RetrieveKnowledgeBaseWithContext(ctx context.Context, request *RetrieveKnowledgeBaseRequest, runtime *dara.RuntimeOptions) (_result *RetrieveKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.HybridSearch) {
		body["HybridSearch"] = request.HybridSearch
	}

	if !dara.IsNil(request.HybridSearchArgs) {
		body["HybridSearchArgs"] = request.HybridSearchArgs
	}

	if !dara.IsNil(request.IncludeMetadataFields) {
		body["IncludeMetadataFields"] = request.IncludeMetadataFields
	}

	if !dara.IsNil(request.IncludeVector) {
		body["IncludeVector"] = request.IncludeVector
	}

	if !dara.IsNil(request.KbUuid) {
		body["KbUuid"] = request.KbUuid
	}

	if !dara.IsNil(request.Metrics) {
		body["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Offset) {
		body["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OrderBy) {
		body["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.RecallWindow) {
		body["RecallWindow"] = request.RecallWindow
	}

	if !dara.IsNil(request.RerankFactor) {
		body["RerankFactor"] = request.RerankFactor
	}

	if !dara.IsNil(request.TopK) {
		body["TopK"] = request.TopK
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetrieveKnowledgeBase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetrieveKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves workspace code. If the file does not exist, a new file is automatically created.
//
// Description:
//
// 发布工作空间的代码
//
// @param request - SaveWorkspaceCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveWorkspaceCodeResponse
func (client *Client) SaveWorkspaceCodeWithContext(ctx context.Context, request *SaveWorkspaceCodeRequest, runtime *dara.RuntimeOptions) (_result *SaveWorkspaceCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	if !dara.IsNil(request.Iac) {
		body["Iac"] = request.Iac
	}

	if !dara.IsNil(request.Mtime) {
		body["Mtime"] = request.Mtime
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.Repo) {
		body["Repo"] = request.Repo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveWorkspaceCode"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveWorkspaceCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Send a user message to a specified session or cancel the session.
//
// Description:
//
// ## Request description
//
// - `agent_id` and `session_id` are required fields.
//
// - `message_type` defaults to `primary`. When you need to append information or cancel a session, set it to `additional` or `cancel`.
//
// - The `reply_to` field indicates which Agent message this message is responding to. The default value is `0`.
//
// - When `message_type` is `additional`, the `question` field is required.
//
// - `quoted_message` can be used to quote the content of the user\\"s previous message.
//
// - Fields such as `data_source`, `dms_user`, `db_metadata`, and `session_config` are all optional, but provide more detailed context information.
//
// @param tmpReq - SendChatMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendChatMessageResponse
func (client *Client) SendChatMessageWithContext(ctx context.Context, tmpReq *SendChatMessageRequest, runtime *dara.RuntimeOptions) (_result *SendChatMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendChatMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSource) {
		request.DataSourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSource, dara.String("DataSource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DataSources) {
		request.DataSourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSources, dara.String("DataSources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SessionConfig) {
		request.SessionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionConfig, dara.String("SessionConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskConfig) {
		request.TaskConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskConfig, dara.String("TaskConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.DataSourceShrink) {
		query["DataSource"] = request.DataSourceShrink
	}

	if !dara.IsNil(request.DataSourcesShrink) {
		query["DataSources"] = request.DataSourcesShrink
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.MessageType) {
		query["MessageType"] = request.MessageType
	}

	if !dara.IsNil(request.ParentSessionId) {
		query["ParentSessionId"] = request.ParentSessionId
	}

	if !dara.IsNil(request.Question) {
		query["Question"] = request.Question
	}

	if !dara.IsNil(request.QuotedMessage) {
		query["QuotedMessage"] = request.QuotedMessage
	}

	if !dara.IsNil(request.ReplyTo) {
		query["ReplyTo"] = request.ReplyTo
	}

	if !dara.IsNil(request.SessionConfigShrink) {
		query["SessionConfig"] = request.SessionConfigShrink
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TaskConfigShrink) {
		query["TaskConfig"] = request.TaskConfigShrink
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendChatMessage"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendChatMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the code deployment configuration for a workspace. This is an asynchronous operation that returns a key. Use this key to query the operation\\"s status by calling the WorkspaceActionStatus operation.
//
// Description:
//
// Sets the default code deployment configuration for a workspace. This configuration includes the Git repository branch and the directories to exclude from deployment.
//
// @param request - SetWorkspaceCodePublishSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWorkspaceCodePublishSettingResponse
func (client *Client) SetWorkspaceCodePublishSettingWithContext(ctx context.Context, request *SetWorkspaceCodePublishSettingRequest, runtime *dara.RuntimeOptions) (_result *SetWorkspaceCodePublishSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetWorkspaceCodePublishSetting"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetWorkspaceCodePublishSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the quota for a workspace.
//
// Description:
//
// Sets the quota for a specific workspace.
//
// @param request - SetWorkspaceQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWorkspaceQuotaResponse
func (client *Client) SetWorkspaceQuotaWithContext(ctx context.Context, request *SetWorkspaceQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetWorkspaceQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CuQuota) {
		query["CuQuota"] = request.CuQuota
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetWorkspaceQuota"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetWorkspaceQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the properties of an Airflow instance.
//
// @param tmpReq - UpdateAirflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAirflowResponse
func (client *Client) UpdateAirflowWithContext(ctx context.Context, tmpReq *UpdateAirflowRequest, runtime *dara.RuntimeOptions) (_result *UpdateAirflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAirflowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataMountInfoList) {
		request.DataMountInfoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataMountInfoList, dara.String("DataMountInfoList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AirflowId) {
		query["AirflowId"] = request.AirflowId
	}

	if !dara.IsNil(request.AirflowName) {
		query["AirflowName"] = request.AirflowName
	}

	if !dara.IsNil(request.AppSpec) {
		query["AppSpec"] = request.AppSpec
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DagsDir) {
		query["DagsDir"] = request.DagsDir
	}

	if !dara.IsNil(request.DataMountInfoListShrink) {
		query["DataMountInfoList"] = request.DataMountInfoListShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableServerless) {
		query["EnableServerless"] = request.EnableServerless
	}

	if !dara.IsNil(request.GracefulShutdownTimeout) {
		query["GracefulShutdownTimeout"] = request.GracefulShutdownTimeout
	}

	if !dara.IsNil(request.PluginsDir) {
		query["PluginsDir"] = request.PluginsDir
	}

	if !dara.IsNil(request.RequirementFile) {
		query["RequirementFile"] = request.RequirementFile
	}

	if !dara.IsNil(request.StartupFile) {
		query["StartupFile"] = request.StartupFile
	}

	if !dara.IsNil(request.WorkerServerlessReplicas) {
		query["WorkerServerlessReplicas"] = request.WorkerServerlessReplicas
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAirflow"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAirflowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates workspace information.
//
// @param request - UpdateDataAgentSpaceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataAgentSpaceInfoResponse
func (client *Client) UpdateDataAgentSpaceInfoWithContext(ctx context.Context, request *UpdateDataAgentSpaceInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataAgentSpaceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.IsSessionShareEnabled) {
		query["IsSessionShareEnabled"] = request.IsSessionShareEnabled
	}

	if !dara.IsNil(request.WorkspaceDesc) {
		query["WorkspaceDesc"] = request.WorkspaceDesc
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataAgentSpaceInfo"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataAgentSpaceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the role of a workspace member.
//
// @param request - UpdateDataAgentWorkspaceMemberRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataAgentWorkspaceMemberRoleResponse
func (client *Client) UpdateDataAgentWorkspaceMemberRoleWithContext(ctx context.Context, request *UpdateDataAgentWorkspaceMemberRoleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataAgentWorkspaceMemberRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.MemberId) {
		query["MemberId"] = request.MemberId
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataAgentWorkspaceMemberRole"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataAgentWorkspaceMemberRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a data lakehouse database.
//
// @param tmpReq - UpdateDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakeDatabaseResponse
func (client *Client) UpdateDataLakeDatabaseWithContext(ctx context.Context, tmpReq *UpdateDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakeDatabaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakeDatabase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakeDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a user-defined function in a data lake.
//
// @param tmpReq - UpdateDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakeFunctionResponse
func (client *Client) UpdateDataLakeFunctionWithContext(ctx context.Context, tmpReq *UpdateDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakeFunctionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FunctionInput) {
		request.FunctionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FunctionInput, dara.String("FunctionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FunctionInputShrink) {
		body["FunctionInput"] = request.FunctionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakeFunction"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakeFunctionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the partition information of a data lake table.
//
// @param tmpReq - UpdateDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakePartitionResponse
func (client *Client) UpdateDataLakePartitionWithContext(ctx context.Context, tmpReq *UpdateDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInput) {
		request.PartitionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInput, dara.String("PartitionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputShrink) {
		body["PartitionInput"] = request.PartitionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakePartition"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakePartitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of a data lake table.
//
// @param tmpReq - UpdateDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakeTableResponse
func (client *Client) UpdateDataLakeTableWithContext(ctx context.Context, tmpReq *UpdateDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakeTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableInput) {
		request.TableInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableInput, dara.String("TableInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableInputShrink) {
		body["TableInput"] = request.TableInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakeTable"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakeTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a document.
//
// @param request - UpdateDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocumentResponse
func (client *Client) UpdateDocumentWithContext(ctx context.Context, request *UpdateDocumentRequest, runtime *dara.RuntimeOptions) (_result *UpdateDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentName) {
		body["DocumentName"] = request.DocumentName
	}

	if !dara.IsNil(request.KbUuid) {
		body["KbUuid"] = request.KbUuid
	}

	if !dara.IsNil(request.NewDescription) {
		body["NewDescription"] = request.NewDescription
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDocument"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a knowledge base.
//
// @param request - UpdateKnowledgeBaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKnowledgeBaseResponse
func (client *Client) UpdateKnowledgeBaseWithContext(ctx context.Context, request *UpdateKnowledgeBaseRequest, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseResponse, _err error) {
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

	if !dara.IsNil(request.KbUuid) {
		query["KbUuid"] = request.KbUuid
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKnowledgeBase"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a document to a knowledge base.
//
// @param tmpReq - UploadDocumentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDocumentResponse
func (client *Client) UploadDocumentWithContext(ctx context.Context, tmpReq *UploadDocumentRequest, runtime *dara.RuntimeOptions) (_result *UploadDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UploadDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Separators) {
		request.SeparatorsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Separators, dara.String("Separators"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChunkOverlap) {
		body["ChunkOverlap"] = request.ChunkOverlap
	}

	if !dara.IsNil(request.ChunkSize) {
		body["ChunkSize"] = request.ChunkSize
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DocumentLoaderName) {
		body["DocumentLoaderName"] = request.DocumentLoaderName
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.KbUuid) {
		body["KbUuid"] = request.KbUuid
	}

	if !dara.IsNil(request.Location) {
		body["Location"] = request.Location
	}

	if !dara.IsNil(request.SeparatorsShrink) {
		body["Separators"] = request.SeparatorsShrink
	}

	if !dara.IsNil(request.SplitterModel) {
		body["SplitterModel"] = request.SplitterModel
	}

	if !dara.IsNil(request.TextSplitterName) {
		body["TextSplitterName"] = request.TextSplitterName
	}

	if !dara.IsNil(request.VlEnhance) {
		body["VlEnhance"] = request.VlEnhance
	}

	if !dara.IsNil(request.ZhTitleEnhance) {
		body["ZhTitleEnhance"] = request.ZhTitleEnhance
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDocument"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upserts document chunks into a knowledge base.
//
// @param request - UpsertDocumentChunksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpsertDocumentChunksResponse
func (client *Client) UpsertDocumentChunksWithContext(ctx context.Context, request *UpsertDocumentChunksRequest, runtime *dara.RuntimeOptions) (_result *UpsertDocumentChunksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Chunks) {
		body["Chunks"] = request.Chunks
	}

	if !dara.IsNil(request.DocumentName) {
		body["DocumentName"] = request.DocumentName
	}

	if !dara.IsNil(request.KbUuid) {
		body["KbUuid"] = request.KbUuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpsertDocumentChunks"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpsertDocumentChunksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log of an asynchronous task in a workspace.
//
// Description:
//
// Pass the `key` to view the execution log of the corresponding asynchronous task. Use this API for troubleshooting.
//
// @param request - WorkspaceActionLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WorkspaceActionLogResponse
func (client *Client) WorkspaceActionLogWithContext(ctx context.Context, request *WorkspaceActionLogRequest, runtime *dara.RuntimeOptions) (_result *WorkspaceActionLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WorkspaceActionLog"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WorkspaceActionLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an asynchronous task in a workspace. Operations such as a deployment return a key. Call this operation with the key to retrieve the task\\"s status.
//
// Description:
//
// Provide the key returned by an asynchronous action, such as a deployment, to retrieve the task\\"s status.
//
// @param request - WorkspaceActionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WorkspaceActionStatusResponse
func (client *Client) WorkspaceActionStatusWithContext(ctx context.Context, request *WorkspaceActionStatusRequest, runtime *dara.RuntimeOptions) (_result *WorkspaceActionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WorkspaceActionStatus"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WorkspaceActionStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This asynchronous API returns a key. Use this key to query the WorkspaceActionStatus API for the code deployment status.
//
// Description:
//
// Deploys the code in a workspace.
//
// @param request - WorkspaceCodePublishRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WorkspaceCodePublishResponse
func (client *Client) WorkspaceCodePublishWithContext(ctx context.Context, request *WorkspaceCodePublishRequest, runtime *dara.RuntimeOptions) (_result *WorkspaceCodePublishResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WorkspaceCodePublish"),
		Version:     dara.String("2025-04-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WorkspaceCodePublishResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) getChatContentWithSSECtx_opYieldFunc(_yield chan *GetChatContentResponse, _yieldErr chan error, ctx context.Context, request *GetChatContentRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.Checkpoint) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !dara.IsNil(request.DMSUnit) {
		query["DMSUnit"] = request.DMSUnit
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatContent"),
		Version:     dara.String("2025-04-14"),
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
