// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Create outbound call jobs in batch.
//
// @param request - AssignJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignJobsResponse
func (client *Client) AssignJobsWithContext(ctx context.Context, request *AssignJobsRequest, runtime *dara.RuntimeOptions) (_result *AssignJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsAsynchrony) {
		query["IsAsynchrony"] = request.IsAsynchrony
	}

	if !dara.IsNil(request.JobDataParsingTaskId) {
		query["JobDataParsingTaskId"] = request.JobDataParsingTaskId
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.JobsJson) {
		query["JobsJson"] = request.JobsJson
	}

	if !dara.IsNil(request.RosterType) {
		query["RosterType"] = request.RosterType
	}

	if !dara.IsNil(request.StrategyJson) {
		query["StrategyJson"] = request.StrategyJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignJobs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously create an outbound call job.
//
// @param tmpReq - AssignJobsAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignJobsAsyncResponse
func (client *Client) AssignJobsAsyncWithContext(ctx context.Context, tmpReq *AssignJobsAsyncRequest, runtime *dara.RuntimeOptions) (_result *AssignJobsAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AssignJobsAsyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallingNumber) {
		request.CallingNumberShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallingNumber, dara.String("CallingNumber"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.JobsJson) {
		request.JobsJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobsJson, dara.String("JobsJson"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallingNumberShrink) {
		body["CallingNumber"] = request.CallingNumberShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupId) {
		body["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.JobsJsonShrink) {
		body["JobsJson"] = request.JobsJsonShrink
	}

	if !dara.IsNil(request.StrategyJson) {
		body["StrategyJson"] = request.StrategyJson
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssignJobsAsync"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssignJobsAsyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancel a job that is about to be executed or is currently executing (cancellation is invalid for jobs already in a call).
//
// @param request - CancelJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelJobsResponse
func (client *Client) CancelJobsWithContext(ctx context.Context, request *CancelJobsRequest, runtime *dara.RuntimeOptions) (_result *CancelJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobReferenceId) {
		query["JobReferenceId"] = request.JobReferenceId
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelJobs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Change the resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an intelligent configuration.
//
// @param request - CreateAgentProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentProfileResponse
func (client *Client) CreateAgentProfileWithContext(ctx context.Context, request *CreateAgentProfileRequest, runtime *dara.RuntimeOptions) (_result *CreateAgentProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentProfileTemplateId) {
		body["AgentProfileTemplateId"] = request.AgentProfileTemplateId
	}

	if !dara.IsNil(request.AppIp) {
		body["AppIp"] = request.AppIp
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FaqCategoryIds) {
		body["FaqCategoryIds"] = request.FaqCategoryIds
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstructionJson) {
		body["InstructionJson"] = request.InstructionJson
	}

	if !dara.IsNil(request.LabelsJson) {
		body["LabelsJson"] = request.LabelsJson
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.ModelConfig) {
		body["ModelConfig"] = request.ModelConfig
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptJson) {
		body["PromptJson"] = request.PromptJson
	}

	if !dara.IsNil(request.Scenario) {
		body["Scenario"] = request.Scenario
	}

	if !dara.IsNil(request.ScriptId) {
		body["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.VariablesJson) {
		body["VariablesJson"] = request.VariablesJson
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentProfile"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create an annotation job
//
// @param tmpReq - CreateAnnotationMissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnnotationMissionResponse
func (client *Client) CreateAnnotationMissionWithContext(ctx context.Context, tmpReq *CreateAnnotationMissionRequest, runtime *dara.RuntimeOptions) (_result *CreateAnnotationMissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAnnotationMissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AnnotationMissionDebugDataSourceList) {
		request.AnnotationMissionDebugDataSourceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AnnotationMissionDebugDataSourceList, dara.String("AnnotationMissionDebugDataSourceList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.AnnotationMissionDataSourceType) {
		query["AnnotationMissionDataSourceType"] = request.AnnotationMissionDataSourceType
	}

	if !dara.IsNil(request.AnnotationMissionDebugDataSourceListShrink) {
		query["AnnotationMissionDebugDataSourceList"] = request.AnnotationMissionDebugDataSourceListShrink
	}

	if !dara.IsNil(request.AnnotationMissionDebugDataSourceListJsonString) {
		query["AnnotationMissionDebugDataSourceListJsonString"] = request.AnnotationMissionDebugDataSourceListJsonString
	}

	if !dara.IsNil(request.AnnotationMissionName) {
		query["AnnotationMissionName"] = request.AnnotationMissionName
	}

	if !dara.IsNil(request.ChatbotId) {
		query["ChatbotId"] = request.ChatbotId
	}

	if !dara.IsNil(request.ConversationTimeEndFilter) {
		query["ConversationTimeEndFilter"] = request.ConversationTimeEndFilter
	}

	if !dara.IsNil(request.ConversationTimeStartFilter) {
		query["ConversationTimeStartFilter"] = request.ConversationTimeStartFilter
	}

	if !dara.IsNil(request.ExcludeOtherSession) {
		query["ExcludeOtherSession"] = request.ExcludeOtherSession
	}

	if !dara.IsNil(request.Finished) {
		query["Finished"] = request.Finished
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SamplingCount) {
		query["SamplingCount"] = request.SamplingCount
	}

	if !dara.IsNil(request.SamplingRate) {
		query["SamplingRate"] = request.SamplingRate
	}

	if !dara.IsNil(request.SamplingType) {
		query["SamplingType"] = request.SamplingType
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SessionEndReasonFilterList) {
		query["SessionEndReasonFilterList"] = request.SessionEndReasonFilterList
	}

	if !dara.IsNil(request.SessionEndReasonFilterListJsonString) {
		query["SessionEndReasonFilterListJsonString"] = request.SessionEndReasonFilterListJsonString
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAnnotationMission"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAnnotationMissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch create outbound call jobs. This operation is deprecated.
//
// @param request - CreateBatchJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchJobsResponse
func (client *Client) CreateBatchJobsWithContext(ctx context.Context, request *CreateBatchJobsRequest, runtime *dara.RuntimeOptions) (_result *CreateBatchJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchJobDescription) {
		query["BatchJobDescription"] = request.BatchJobDescription
	}

	if !dara.IsNil(request.BatchJobName) {
		query["BatchJobName"] = request.BatchJobName
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobFilePath) {
		query["JobFilePath"] = request.JobFilePath
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.StrategyJson) {
		query["StrategyJson"] = request.StrategyJson
	}

	if !dara.IsNil(request.Submitted) {
		query["Submitted"] = request.Submitted
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBatchJobs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBatchJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copy an existing task group to reinitiate an outbound call job.
//
// @param request - CreateBatchRepeatJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBatchRepeatJobResponse
func (client *Client) CreateBatchRepeatJobWithContext(ctx context.Context, request *CreateBatchRepeatJobRequest, runtime *dara.RuntimeOptions) (_result *CreateBatchRepeatJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FilterStatus) {
		query["FilterStatus"] = request.FilterStatus
	}

	if !dara.IsNil(request.FlashSmsExtras) {
		query["FlashSmsExtras"] = request.FlashSmsExtras
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MinConcurrency) {
		query["MinConcurrency"] = request.MinConcurrency
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RecallCallingNumber) {
		query["RecallCallingNumber"] = request.RecallCallingNumber
	}

	if !dara.IsNil(request.RecallStrategyJson) {
		query["RecallStrategyJson"] = request.RecallStrategyJson
	}

	if !dara.IsNil(request.RingingDuration) {
		query["RingingDuration"] = request.RingingDuration
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SourceGroupId) {
		query["SourceGroupId"] = request.SourceGroupId
	}

	if !dara.IsNil(request.StrategyJson) {
		query["StrategyJson"] = request.StrategyJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBatchRepeatJob"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBatchRepeatJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateBeebotIntent
//
// @param tmpReq - CreateBeebotIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBeebotIntentResponse
func (client *Client) CreateBeebotIntentWithContext(ctx context.Context, tmpReq *CreateBeebotIntentRequest, runtime *dara.RuntimeOptions) (_result *CreateBeebotIntentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBeebotIntentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IntentDefinition) {
		request.IntentDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentDefinition, dara.String("IntentDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentDefinitionShrink) {
		query["IntentDefinition"] = request.IntentDefinitionShrink
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBeebotIntent"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBeebotIntentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateBeebotIntentLgf
//
// @param tmpReq - CreateBeebotIntentLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBeebotIntentLgfResponse
func (client *Client) CreateBeebotIntentLgfWithContext(ctx context.Context, tmpReq *CreateBeebotIntentLgfRequest, runtime *dara.RuntimeOptions) (_result *CreateBeebotIntentLgfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBeebotIntentLgfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LgfDefinition) {
		request.LgfDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LgfDefinition, dara.String("LgfDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LgfDefinitionShrink) {
		query["LgfDefinition"] = request.LgfDefinitionShrink
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBeebotIntentLgf"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBeebotIntentLgfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateBeebotIntentUserSay
//
// @param tmpReq - CreateBeebotIntentUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBeebotIntentUserSayResponse
func (client *Client) CreateBeebotIntentUserSayWithContext(ctx context.Context, tmpReq *CreateBeebotIntentUserSayRequest, runtime *dara.RuntimeOptions) (_result *CreateBeebotIntentUserSayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBeebotIntentUserSayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserSayDefinition) {
		request.UserSayDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserSayDefinition, dara.String("UserSayDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.UserSayDefinitionShrink) {
		query["UserSayDefinition"] = request.UserSayDefinitionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBeebotIntentUserSay"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBeebotIntentUserSayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a conversation flow using the legacy canvas API.
//
// @param request - CreateDialogueFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDialogueFlowResponse
func (client *Client) CreateDialogueFlowWithContext(ctx context.Context, request *CreateDialogueFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateDialogueFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DialogueFlowType) {
		query["DialogueFlowType"] = request.DialogueFlowType
	}

	if !dara.IsNil(request.DialogueName) {
		query["DialogueName"] = request.DialogueName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDialogueFlow"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDialogueFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a URL link for a Download Hub job.
//
// @param request - CreateDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDownloadUrlResponse
func (client *Client) CreateDownloadUrlWithContext(ctx context.Context, request *CreateDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *CreateDownloadUrlResponse, _err error) {
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
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDownloadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a global question in the legacy canvas scenario.
//
// @param request - CreateGlobalQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGlobalQuestionResponse
func (client *Client) CreateGlobalQuestionWithContext(ctx context.Context, request *CreateGlobalQuestionRequest, runtime *dara.RuntimeOptions) (_result *CreateGlobalQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Answers) {
		query["Answers"] = request.Answers
	}

	if !dara.IsNil(request.GlobalQuestionName) {
		query["GlobalQuestionName"] = request.GlobalQuestionName
	}

	if !dara.IsNil(request.GlobalQuestionType) {
		query["GlobalQuestionType"] = request.GlobalQuestionType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Questions) {
		query["Questions"] = request.Questions
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGlobalQuestion"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGlobalQuestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an Intelligent Outbound Call business instance.
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.InstanceDescription) {
		query["InstanceDescription"] = request.InstanceDescription
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.MaxConcurrentConversation) {
		query["MaxConcurrentConversation"] = request.MaxConcurrentConversation
	}

	if !dara.IsNil(request.NluServiceType) {
		query["NluServiceType"] = request.NluServiceType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attach a number to a business instance.
//
// @param request - CreateInstanceBindNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceBindNumberResponse
func (client *Client) CreateInstanceBindNumberWithContext(ctx context.Context, request *CreateInstanceBindNumberRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceBindNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceList) {
		query["InstanceList"] = request.InstanceList
	}

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceBindNumber"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceBindNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a trigger intent in the legacy canvas scenario.
//
// @param request - CreateIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntentResponse
func (client *Client) CreateIntentWithContext(ctx context.Context, request *CreateIntentRequest, runtime *dara.RuntimeOptions) (_result *CreateIntentResponse, _err error) {
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

	if !dara.IsNil(request.IntentDescription) {
		query["IntentDescription"] = request.IntentDescription
	}

	if !dara.IsNil(request.IntentName) {
		query["IntentName"] = request.IntentName
	}

	if !dara.IsNil(request.Keywords) {
		query["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.Utterances) {
		query["Utterances"] = request.Utterances
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIntent"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIntentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a Data Parsing job.
//
// @param request - CreateJobDataParsingTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobDataParsingTaskResponse
func (client *Client) CreateJobDataParsingTaskWithContext(ctx context.Context, request *CreateJobDataParsingTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateJobDataParsingTaskResponse, _err error) {
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

	if !dara.IsNil(request.JobFilePath) {
		query["JobFilePath"] = request.JobFilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJobDataParsingTask"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobDataParsingTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a job group.
//
// @param request - CreateJobGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobGroupResponse
func (client *Client) CreateJobGroupWithContext(ctx context.Context, request *CreateJobGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateJobGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.FlashSmsExtras) {
		query["FlashSmsExtras"] = request.FlashSmsExtras
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupDescription) {
		query["JobGroupDescription"] = request.JobGroupDescription
	}

	if !dara.IsNil(request.JobGroupName) {
		query["JobGroupName"] = request.JobGroupName
	}

	if !dara.IsNil(request.MinConcurrency) {
		query["MinConcurrency"] = request.MinConcurrency
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RecallCallingNumber) {
		query["RecallCallingNumber"] = request.RecallCallingNumber
	}

	if !dara.IsNil(request.RecallStrategyJson) {
		query["RecallStrategyJson"] = request.RecallStrategyJson
	}

	if !dara.IsNil(request.RingingDuration) {
		query["RingingDuration"] = request.RingingDuration
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.StrategyJson) {
		query["StrategyJson"] = request.StrategyJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJobGroup"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an outbound job group export task.
//
// @param request - CreateJobGroupExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobGroupExportTaskResponse
func (client *Client) CreateJobGroupExportTaskWithContext(ctx context.Context, request *CreateJobGroupExportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateJobGroupExportTaskResponse, _err error) {
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

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJobGroupExportTask"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobGroupExportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a scenario.
//
// @param request - CreateScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScriptResponse
func (client *Client) CreateScriptWithContext(ctx context.Context, request *CreateScriptRequest, runtime *dara.RuntimeOptions) (_result *CreateScriptResponse, _err error) {
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

	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.AgentLlm) {
		query["AgentLlm"] = request.AgentLlm
	}

	if !dara.IsNil(request.AsrConfig) {
		query["AsrConfig"] = request.AsrConfig
	}

	if !dara.IsNil(request.ChatbotId) {
		query["ChatbotId"] = request.ChatbotId
	}

	if !dara.IsNil(request.EmotionEnable) {
		query["EmotionEnable"] = request.EmotionEnable
	}

	if !dara.IsNil(request.Industry) {
		query["Industry"] = request.Industry
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LongWaitEnable) {
		query["LongWaitEnable"] = request.LongWaitEnable
	}

	if !dara.IsNil(request.MiniPlaybackEnable) {
		query["MiniPlaybackEnable"] = request.MiniPlaybackEnable
	}

	if !dara.IsNil(request.NewBargeInEnable) {
		query["NewBargeInEnable"] = request.NewBargeInEnable
	}

	if !dara.IsNil(request.NluAccessType) {
		query["NluAccessType"] = request.NluAccessType
	}

	if !dara.IsNil(request.NluEngine) {
		query["NluEngine"] = request.NluEngine
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.ScriptContent) {
		query["ScriptContent"] = request.ScriptContent
	}

	if !dara.IsNil(request.ScriptDescription) {
		query["ScriptDescription"] = request.ScriptDescription
	}

	if !dara.IsNil(request.ScriptName) {
		query["ScriptName"] = request.ScriptName
	}

	if !dara.IsNil(request.ScriptNluProfileJsonString) {
		query["ScriptNluProfileJsonString"] = request.ScriptNluProfileJsonString
	}

	if !dara.IsNil(request.ScriptWaveform) {
		query["ScriptWaveform"] = request.ScriptWaveform
	}

	if !dara.IsNil(request.TtsConfig) {
		query["TtsConfig"] = request.TtsConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScript"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a recording of the script for voice broadcasting.
//
// Description:
//
// ***
//
// @param request - CreateScriptWaveformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScriptWaveformResponse
func (client *Client) CreateScriptWaveformWithContext(ctx context.Context, request *CreateScriptWaveformRequest, runtime *dara.RuntimeOptions) (_result *CreateScriptWaveformResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptContent) {
		query["ScriptContent"] = request.ScriptContent
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScriptWaveform"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScriptWaveformResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a label for use in the legacy canvas.
//
// @param request - CreateTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTagResponse
func (client *Client) CreateTagWithContext(ctx context.Context, request *CreateTagRequest, runtime *dara.RuntimeOptions) (_result *CreateTagResponse, _err error) {
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

	if !dara.IsNil(request.TagGroup) {
		query["TagGroup"] = request.TagGroup
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTag"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create an asynchronous task to export outbound call history.
//
// @param request - CreateTaskExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskExportTaskResponse
func (client *Client) CreateTaskExportTaskWithContext(ctx context.Context, request *CreateTaskExportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateTaskExportTaskResponse, _err error) {
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
		Action:      dara.String("CreateTaskExportTask"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskExportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete an LLM-based robot.
//
// @param tmpReq - DeleteAgentProfilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentProfilesResponse
func (client *Client) DeleteAgentProfilesWithContext(ctx context.Context, tmpReq *DeleteAgentProfilesRequest, runtime *dara.RuntimeOptions) (_result *DeleteAgentProfilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAgentProfilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentProfileIds) {
		request.AgentProfileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentProfileIds, dara.String("AgentProfileIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentProfileIdsShrink) {
		body["AgentProfileIds"] = request.AgentProfileIdsShrink
	}

	if !dara.IsNil(request.AppIp) {
		body["AppIp"] = request.AppIp
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgentProfiles"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentProfilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteBeebotIntent
//
// @param request - DeleteBeebotIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBeebotIntentResponse
func (client *Client) DeleteBeebotIntentWithContext(ctx context.Context, request *DeleteBeebotIntentRequest, runtime *dara.RuntimeOptions) (_result *DeleteBeebotIntentResponse, _err error) {
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

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBeebotIntent"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBeebotIntentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteBeebotIntentLgf
//
// @param request - DeleteBeebotIntentLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBeebotIntentLgfResponse
func (client *Client) DeleteBeebotIntentLgfWithContext(ctx context.Context, request *DeleteBeebotIntentLgfRequest, runtime *dara.RuntimeOptions) (_result *DeleteBeebotIntentLgfResponse, _err error) {
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

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.LgfId) {
		query["LgfId"] = request.LgfId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBeebotIntentLgf"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBeebotIntentLgfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteBeebotIntentUserSay
//
// @param request - DeleteBeebotIntentUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBeebotIntentUserSayResponse
func (client *Client) DeleteBeebotIntentUserSayWithContext(ctx context.Context, request *DeleteBeebotIntentUserSayRequest, runtime *dara.RuntimeOptions) (_result *DeleteBeebotIntentUserSayResponse, _err error) {
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

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.UserSayId) {
		query["UserSayId"] = request.UserSayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBeebotIntentUserSay"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBeebotIntentUserSayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a Do-Not-Call list.
//
// @param request - DeleteContactBlockListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactBlockListResponse
func (client *Client) DeleteContactBlockListWithContext(ctx context.Context, request *DeleteContactBlockListRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactBlockListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactBlockListId) {
		query["ContactBlockListId"] = request.ContactBlockListId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Operator) {
		query["Operator"] = request.Operator
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContactBlockList"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactBlockListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a global outbound call policy whitelist.
//
// @param request - DeleteContactWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactWhiteListResponse
func (client *Client) DeleteContactWhiteListWithContext(ctx context.Context, request *DeleteContactWhiteListRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactWhiteListId) {
		query["ContactWhiteListId"] = request.ContactWhiteListId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Operator) {
		query["Operator"] = request.Operator
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContactWhiteList"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactWhiteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a conversation flow using the legacy canvas API.
//
// @param request - DeleteDialogueFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDialogueFlowResponse
func (client *Client) DeleteDialogueFlowWithContext(ctx context.Context, request *DeleteDialogueFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteDialogueFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DialogueFlowId) {
		query["DialogueFlowId"] = request.DialogueFlowId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDialogueFlow"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDialogueFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a global question in a legacy canvas scenario.
//
// @param request - DeleteGlobalQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGlobalQuestionResponse
func (client *Client) DeleteGlobalQuestionWithContext(ctx context.Context, request *DeleteGlobalQuestionRequest, runtime *dara.RuntimeOptions) (_result *DeleteGlobalQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalQuestionId) {
		query["GlobalQuestionId"] = request.GlobalQuestionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGlobalQuestion"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGlobalQuestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete an Intelligent Outbound Call business instance.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete an intent using the legacy canvas API.
//
// @param request - DeleteIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIntentResponse
func (client *Client) DeleteIntentWithContext(ctx context.Context, request *DeleteIntentRequest, runtime *dara.RuntimeOptions) (_result *DeleteIntentResponse, _err error) {
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

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIntent"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIntentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a job.
//
// @param request - DeleteJobGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobGroupResponse
func (client *Client) DeleteJobGroupWithContext(ctx context.Context, request *DeleteJobGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteJobGroupResponse, _err error) {
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

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJobGroup"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJobGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an outbound call number. This API is deprecated.
//
// @param request - DeleteOutboundCallNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOutboundCallNumberResponse
func (client *Client) DeleteOutboundCallNumberWithContext(ctx context.Context, request *DeleteOutboundCallNumberRequest, runtime *dara.RuntimeOptions) (_result *DeleteOutboundCallNumberResponse, _err error) {
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

	if !dara.IsNil(request.OutboundCallNumberId) {
		query["OutboundCallNumberId"] = request.OutboundCallNumberId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOutboundCallNumber"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOutboundCallNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a script.
//
// @param request - DeleteScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScriptResponse
func (client *Client) DeleteScriptWithContext(ctx context.Context, request *DeleteScriptRequest, runtime *dara.RuntimeOptions) (_result *DeleteScriptResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScript"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete uploaded recordings in a small model scenario.
//
// @param request - DeleteScriptRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScriptRecordingResponse
func (client *Client) DeleteScriptRecordingWithContext(ctx context.Context, request *DeleteScriptRecordingRequest, runtime *dara.RuntimeOptions) (_result *DeleteScriptRecordingResponse, _err error) {
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

	if !dara.IsNil(request.UuidsJson) {
		query["UuidsJson"] = request.UuidsJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScriptRecording"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScriptRecordingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete script recording.
//
// @param request - DeleteScriptWaveformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScriptWaveformResponse
func (client *Client) DeleteScriptWaveformWithContext(ctx context.Context, request *DeleteScriptWaveformRequest, runtime *dara.RuntimeOptions) (_result *DeleteScriptWaveformResponse, _err error) {
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

	if !dara.IsNil(request.ScriptWaveformId) {
		query["ScriptWaveformId"] = request.ScriptWaveformId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScriptWaveform"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScriptWaveformResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeBeebotIntent
//
// Description:
//
// ***
//
// @param request - DescribeBeebotIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBeebotIntentResponse
func (client *Client) DescribeBeebotIntentWithContext(ctx context.Context, request *DescribeBeebotIntentRequest, runtime *dara.RuntimeOptions) (_result *DescribeBeebotIntentResponse, _err error) {
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

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBeebotIntent"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBeebotIntentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the statistics information of job labels.
//
// @param request - DescribeDialogueNodeStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDialogueNodeStatisticsResponse
func (client *Client) DescribeDialogueNodeStatisticsWithContext(ctx context.Context, request *DescribeDialogueNodeStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDialogueNodeStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDialogueNodeStatistics"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDialogueNodeStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeDsReports
//
// @param request - DescribeDsReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDsReportsResponse
func (client *Client) DescribeDsReportsWithContext(ctx context.Context, request *DescribeDsReportsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDsReportsResponse, _err error) {
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

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDsReports"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDsReportsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns data of a global question in the legacy canvas scenario.
//
// Description:
//
// ***
//
// @param request - DescribeGlobalQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalQuestionResponse
func (client *Client) DescribeGlobalQuestionWithContext(ctx context.Context, request *DescribeGlobalQuestionRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalQuestionId) {
		query["GlobalQuestionId"] = request.GlobalQuestionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGlobalQuestion"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalQuestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the task execution status within a task group.
//
// @param request - DescribeGroupExecutingInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupExecutingInfoResponse
func (client *Client) DescribeGroupExecutingInfoWithContext(ctx context.Context, request *DescribeGroupExecutingInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupExecutingInfoResponse, _err error) {
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

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupExecutingInfo"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupExecutingInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Intelligent Outbound Calling business instance information.
//
// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithContext(ctx context.Context, request *DescribeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
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
		Action:      dara.String("DescribeInstance"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve an intent using the legacy canvas API.
//
// @param request - DescribeIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIntentResponse
func (client *Client) DescribeIntentWithContext(ctx context.Context, request *DescribeIntentRequest, runtime *dara.RuntimeOptions) (_result *DescribeIntentResponse, _err error) {
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

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIntent"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIntentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the intent statistics information under a task group.
//
// @param request - DescribeIntentStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIntentStatisticsResponse
func (client *Client) DescribeIntentStatisticsWithContext(ctx context.Context, request *DescribeIntentStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeIntentStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIntentStatistics"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIntentStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the conversation data of a job.
//
// @param request - DescribeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobResponse
func (client *Client) DescribeJobWithContext(ctx context.Context, request *DescribeJobRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.WithScript) {
		query["WithScript"] = request.WithScript
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJob"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the parsing progress of the uploaded job file.
//
// @param request - DescribeJobDataParsingTaskProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobDataParsingTaskProgressResponse
func (client *Client) DescribeJobDataParsingTaskProgressWithContext(ctx context.Context, request *DescribeJobDataParsingTaskProgressRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobDataParsingTaskProgressResponse, _err error) {
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

	if !dara.IsNil(request.JobDataParsingTaskId) {
		query["JobDataParsingTaskId"] = request.JobDataParsingTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobDataParsingTaskProgress"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobDataParsingTaskProgressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Intelligent Outbound Call Job Data.
//
// @param request - DescribeJobGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobGroupResponse
func (client *Client) DescribeJobGroupWithContext(ctx context.Context, request *DescribeJobGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BriefTypes) {
		query["BriefTypes"] = request.BriefTypes
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobGroup"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the progress of a task group export job. This API is deprecated and can be replaced by ListDownloadTasks.
//
// @param request - DescribeJobGroupExportTaskProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobGroupExportTaskProgressResponse
func (client *Client) DescribeJobGroupExportTaskProgressWithContext(ctx context.Context, request *DescribeJobGroupExportTaskProgressRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobGroupExportTaskProgressResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobGroupExportTaskProgress"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobGroupExportTaskProgressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns information about a script.
//
// @param request - DescribeScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScriptResponse
func (client *Client) DescribeScriptWithContext(ctx context.Context, request *DescribeScriptRequest, runtime *dara.RuntimeOptions) (_result *DescribeScriptResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScript"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a scenario voice configuration in the legacy canvas scenario.
//
// @param request - DescribeScriptVoiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScriptVoiceConfigResponse
func (client *Client) DescribeScriptVoiceConfigWithContext(ctx context.Context, request *DescribeScriptVoiceConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeScriptVoiceConfigResponse, _err error) {
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

	if !dara.IsNil(request.ScriptVoiceConfigId) {
		query["ScriptVoiceConfigId"] = request.ScriptVoiceConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScriptVoiceConfig"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScriptVoiceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query TTS configuration information.
//
// @param request - DescribeTTSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTTSConfigResponse
func (client *Client) DescribeTTSConfigWithContext(ctx context.Context, request *DescribeTTSConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeTTSConfigResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTTSConfig"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTTSConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Time Tunnel试听。
//
// @param request - DescribeTTSDemoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTTSDemoResponse
func (client *Client) DescribeTTSDemoWithContext(ctx context.Context, request *DescribeTTSDemoRequest, runtime *dara.RuntimeOptions) (_result *DescribeTTSDemoResponse, _err error) {
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

	if !dara.IsNil(request.AliCustomizedVoice) {
		query["AliCustomizedVoice"] = request.AliCustomizedVoice
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

	if !dara.IsNil(request.NlsServiceType) {
		query["NlsServiceType"] = request.NlsServiceType
	}

	if !dara.IsNil(request.PitchRate) {
		query["PitchRate"] = request.PitchRate
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
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
		Action:      dara.String("DescribeTTSDemo"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTTSDemoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query tag statistics using the legacy canvas API.
//
// @param request - DescribeTagHitsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagHitsSummaryResponse
func (client *Client) DescribeTagHitsSummaryWithContext(ctx context.Context, request *DescribeTagHitsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagHitsSummaryResponse, _err error) {
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

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagHitsSummary"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagHitsSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List of number binding instances under the tenant.
//
// @param request - DescribeTenantBindNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTenantBindNumberResponse
func (client *Client) DescribeTenantBindNumberWithContext(ctx context.Context, request *DescribeTenantBindNumberRequest, runtime *dara.RuntimeOptions) (_result *DescribeTenantBindNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTenantBindNumber"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTenantBindNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The robot conversation API for CC callback outbound calls. (Deprecated.)
//
// @param request - DialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DialogueResponse
func (client *Client) DialogueWithContext(ctx context.Context, request *DialogueRequest, runtime *dara.RuntimeOptions) (_result *DialogueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionKey) {
		query["ActionKey"] = request.ActionKey
	}

	if !dara.IsNil(request.ActionParams) {
		query["ActionParams"] = request.ActionParams
	}

	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CallType) {
		query["CallType"] = request.CallType
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Utterance) {
		query["Utterance"] = request.Utterance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Dialogue"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DialogueResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Download call recordings.
//
// @param request - DownloadRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadRecordingResponse
func (client *Client) DownloadRecordingWithContext(ctx context.Context, request *DownloadRecordingRequest, runtime *dara.RuntimeOptions) (_result *DownloadRecordingResponse, _err error) {
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

	if !dara.IsNil(request.NeedVoiceSliceRecording) {
		query["NeedVoiceSliceRecording"] = request.NeedVoiceSliceRecording
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadRecording"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadRecordingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Download a recording (obtain the recording URL), dedicated to small model scenarios.
//
// @param request - DownloadScriptRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadScriptRecordingResponse
func (client *Client) DownloadScriptRecordingWithContext(ctx context.Context, request *DownloadScriptRecordingRequest, runtime *dara.RuntimeOptions) (_result *DownloadScriptRecordingResponse, _err error) {
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

	if !dara.IsNil(request.Uuid) {
		query["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadScriptRecording"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadScriptRecordingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copy a new script from an existing one (for legacy canvas only).
//
// @param request - DuplicateScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DuplicateScriptResponse
func (client *Client) DuplicateScriptWithContext(ctx context.Context, request *DuplicateScriptRequest, runtime *dara.RuntimeOptions) (_result *DuplicateScriptResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SourceScriptId) {
		query["SourceScriptId"] = request.SourceScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DuplicateScript"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DuplicateScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Export a scenario.
//
// @param request - ExportScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportScriptResponse
func (client *Client) ExportScriptWithContext(ctx context.Context, request *ExportScriptRequest, runtime *dara.RuntimeOptions) (_result *ExportScriptResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportScript"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upload a hotword list file.
//
// @param request - GenerateUploadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUploadUrlResponse
func (client *Client) GenerateUploadUrlWithContext(ctx context.Context, request *GenerateUploadUrlRequest, runtime *dara.RuntimeOptions) (_result *GenerateUploadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUploadUrl"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUploadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the delay time for audio playback after the call is answered.
//
// @param request - GetAfterAnswerDelayPlaybackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAfterAnswerDelayPlaybackResponse
func (client *Client) GetAfterAnswerDelayPlaybackWithContext(ctx context.Context, request *GetAfterAnswerDelayPlaybackRequest, runtime *dara.RuntimeOptions) (_result *GetAfterAnswerDelayPlaybackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAfterAnswerDelayPlayback"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAfterAnswerDelayPlaybackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the agent configuration.
//
// @param request - GetAgentProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentProfileResponse
func (client *Client) GetAgentProfileWithContext(ctx context.Context, request *GetAgentProfileRequest, runtime *dara.RuntimeOptions) (_result *GetAgentProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentProfileId) {
		body["AgentProfileId"] = request.AgentProfileId
	}

	if !dara.IsNil(request.AppIp) {
		body["AppIp"] = request.AppIp
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentProfile"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the agent configuration template.
//
// @param request - GetAgentProfileTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentProfileTemplateResponse
func (client *Client) GetAgentProfileTemplateWithContext(ctx context.Context, request *GetAgentProfileTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetAgentProfileTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentProfileTemplateId) {
		body["AgentProfileTemplateId"] = request.AgentProfileTemplateId
	}

	if !dara.IsNil(request.AppIp) {
		body["AppIp"] = request.AppIp
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentProfileTemplate"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentProfileTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// null
//
// @param request - GetAnnotationMissionSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAnnotationMissionSummaryResponse
func (client *Client) GetAnnotationMissionSummaryWithContext(ctx context.Context, request *GetAnnotationMissionSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetAnnotationMissionSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnnotationMissionId) {
		query["AnnotationMissionId"] = request.AnnotationMissionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAnnotationMissionSummary"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAnnotationMissionSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAnnotationMissionTagInfoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAnnotationMissionTagInfoListResponse
func (client *Client) GetAnnotationMissionTagInfoListWithContext(ctx context.Context, request *GetAnnotationMissionTagInfoListRequest, runtime *dara.RuntimeOptions) (_result *GetAnnotationMissionTagInfoListResponse, _err error) {
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

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAnnotationMissionTagInfoList"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAnnotationMissionTagInfoListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain information such as the ASR model list and hotword list. This API is deprecated.
//
// @param request - GetAsrServerInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsrServerInfoResponse
func (client *Client) GetAsrServerInfoWithContext(ctx context.Context, request *GetAsrServerInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAsrServerInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsrServerInfo"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsrServerInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the upload result of an asynchronous outbound call job.
//
// @param request - GetAssignJobsAsyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAssignJobsAsyncResultResponse
func (client *Client) GetAssignJobsAsyncResultWithContext(ctx context.Context, request *GetAssignJobsAsyncResultRequest, runtime *dara.RuntimeOptions) (_result *GetAssignJobsAsyncResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsyncTaskId) {
		query["AsyncTaskId"] = request.AsyncTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAssignJobsAsyncResult"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAssignJobsAsyncResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the global outbound calling time segments in the outbound calling system.
//
// @param request - GetBaseStrategyPeriodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBaseStrategyPeriodResponse
func (client *Client) GetBaseStrategyPeriodWithContext(ctx context.Context, request *GetBaseStrategyPeriodRequest, runtime *dara.RuntimeOptions) (_result *GetBaseStrategyPeriodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBaseStrategyPeriod"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBaseStrategyPeriodResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the outbound call blocklist.
//
// @param request - GetContactBlockListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContactBlockListResponse
func (client *Client) GetContactBlockListWithContext(ctx context.Context, request *GetContactBlockListRequest, runtime *dara.RuntimeOptions) (_result *GetContactBlockListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CountTotalRow) {
		query["CountTotalRow"] = request.CountTotalRow
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContactBlockList"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContactBlockListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the global outbound call policy whitelist.
//
// @param request - GetContactWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContactWhiteListResponse
func (client *Client) GetContactWhiteListWithContext(ctx context.Context, request *GetContactWhiteListRequest, runtime *dara.RuntimeOptions) (_result *GetContactWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CountTotalRow) {
		query["CountTotalRow"] = request.CountTotalRow
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContactWhiteList"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContactWhiteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the current concurrency of the instance.
//
// @param request - GetCurrentConcurrencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCurrentConcurrencyResponse
func (client *Client) GetCurrentConcurrencyWithContext(ctx context.Context, request *GetCurrentConcurrencyRequest, runtime *dara.RuntimeOptions) (_result *GetCurrentConcurrencyResponse, _err error) {
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
		Action:      dara.String("GetCurrentConcurrency"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCurrentConcurrencyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Toggle for the nonexistent number no-more-calls feature at the global dimension.
//
// @param request - GetEmptyNumberNoMoreCallsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmptyNumberNoMoreCallsInfoResponse
func (client *Client) GetEmptyNumberNoMoreCallsInfoWithContext(ctx context.Context, request *GetEmptyNumberNoMoreCallsInfoRequest, runtime *dara.RuntimeOptions) (_result *GetEmptyNumberNoMoreCallsInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmptyNumberNoMoreCallsInfo"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmptyNumberNoMoreCallsInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain job data upload parameters.
//
// @param request - GetJobDataUploadParamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDataUploadParamsResponse
func (client *Client) GetJobDataUploadParamsWithContext(ctx context.Context, request *GetJobDataUploadParamsRequest, runtime *dara.RuntimeOptions) (_result *GetJobDataUploadParamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusiType) {
		query["BusiType"] = request.BusiType
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobDataUploadParams"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobDataUploadParamsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the daily call limit for the called number.
//
// @param request - GetMaxAttemptsPerDayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMaxAttemptsPerDayResponse
func (client *Client) GetMaxAttemptsPerDayWithContext(ctx context.Context, request *GetMaxAttemptsPerDayRequest, runtime *dara.RuntimeOptions) (_result *GetMaxAttemptsPerDayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMaxAttemptsPerDay"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMaxAttemptsPerDayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetRealtimeConcurrencyReport
//
// @param request - GetRealtimeConcurrencyReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRealtimeConcurrencyReportResponse
func (client *Client) GetRealtimeConcurrencyReportWithContext(ctx context.Context, request *GetRealtimeConcurrencyReportRequest, runtime *dara.RuntimeOptions) (_result *GetRealtimeConcurrencyReportResponse, _err error) {
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
		Action:      dara.String("GetRealtimeConcurrencyReport"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRealtimeConcurrencyReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Overview of outbound call instances.
//
// @param request - GetSummaryInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSummaryInfoResponse
func (client *Client) GetSummaryInfoWithContext(ctx context.Context, request *GetSummaryInfoRequest, runtime *dara.RuntimeOptions) (_result *GetSummaryInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSummaryInfo"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSummaryInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve a job by UUID.
//
// @param request - GetTaskByUuidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskByUuidResponse
func (client *Client) GetTaskByUuidWithContext(ctx context.Context, request *GetTaskByUuidRequest, runtime *dara.RuntimeOptions) (_result *GetTaskByUuidResponse, _err error) {
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
		Action:      dara.String("GetTaskByUuid"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskByUuidResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Import a scenario.
//
// @param request - ImportScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportScriptResponse
func (client *Client) ImportScriptWithContext(ctx context.Context, request *ImportScriptRequest, runtime *dara.RuntimeOptions) (_result *ImportScriptResponse, _err error) {
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

	if !dara.IsNil(request.NluEngine) {
		query["NluEngine"] = request.NluEngine
	}

	if !dara.IsNil(request.SignatureUrl) {
		query["SignatureUrl"] = request.SignatureUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportScript"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # InflightTaskTimeout
//
// @param request - InflightTaskTimeoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InflightTaskTimeoutResponse
func (client *Client) InflightTaskTimeoutWithContext(ctx context.Context, request *InflightTaskTimeoutRequest, runtime *dara.RuntimeOptions) (_result *InflightTaskTimeoutResponse, _err error) {
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

	if !dara.IsNil(request.InstanceOwnerId) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InflightTaskTimeout"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InflightTaskTimeoutResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the historical publish list of agent configurations under this LLM scenario.
//
// @param request - ListAgentProfilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentProfilesResponse
func (client *Client) ListAgentProfilesWithContext(ctx context.Context, request *ListAgentProfilesRequest, runtime *dara.RuntimeOptions) (_result *ListAgentProfilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppIp) {
		body["AppIp"] = request.AppIp
	}

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
		Action:      dara.String("ListAgentProfiles"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentProfilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the annotation job list.
//
// @param request - ListAnnotationMissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnnotationMissionResponse
func (client *Client) ListAnnotationMissionWithContext(ctx context.Context, request *ListAnnotationMissionRequest, runtime *dara.RuntimeOptions) (_result *ListAnnotationMissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnnotationMissionId) {
		query["AnnotationMissionId"] = request.AnnotationMissionId
	}

	if !dara.IsNil(request.AnnotationMissionName) {
		query["AnnotationMissionName"] = request.AnnotationMissionName
	}

	if !dara.IsNil(request.AnnotationStatusListFilter) {
		query["AnnotationStatusListFilter"] = request.AnnotationStatusListFilter
	}

	if !dara.IsNil(request.AnnotationStatusListStringFilter) {
		query["AnnotationStatusListStringFilter"] = request.AnnotationStatusListStringFilter
	}

	if !dara.IsNil(request.CreateTimeEndFilter) {
		query["CreateTimeEndFilter"] = request.CreateTimeEndFilter
	}

	if !dara.IsNil(request.CreateTimeStartFilter) {
		query["CreateTimeStartFilter"] = request.CreateTimeStartFilter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnnotationMission"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnnotationMissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListAnnotationMissionSession
//
// @param request - ListAnnotationMissionSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnnotationMissionSessionResponse
func (client *Client) ListAnnotationMissionSessionWithContext(ctx context.Context, request *ListAnnotationMissionSessionRequest, runtime *dara.RuntimeOptions) (_result *ListAnnotationMissionSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnnotationMissionId) {
		query["AnnotationMissionId"] = request.AnnotationMissionId
	}

	if !dara.IsNil(request.AnnotationMissionSessionId) {
		query["AnnotationMissionSessionId"] = request.AnnotationMissionSessionId
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.IncludeStatusListJsonString) {
		query["IncludeStatusListJsonString"] = request.IncludeStatusListJsonString
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnnotationMissionSession"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnnotationMissionSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the API plugin list.
//
// @param request - ListApiPluginsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiPluginsResponse
func (client *Client) ListApiPluginsWithContext(ctx context.Context, request *ListApiPluginsRequest, runtime *dara.RuntimeOptions) (_result *ListApiPluginsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UuidsJson) {
		query["UuidsJson"] = request.UuidsJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiPlugins"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiPluginsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListBeebotIntent
//
// @param request - ListBeebotIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBeebotIntentResponse
func (client *Client) ListBeebotIntentWithContext(ctx context.Context, request *ListBeebotIntentRequest, runtime *dara.RuntimeOptions) (_result *ListBeebotIntentResponse, _err error) {
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

	if !dara.IsNil(request.IntentName) {
		query["IntentName"] = request.IntentName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBeebotIntent"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBeebotIntentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListBeebotIntentLgf
//
// @param request - ListBeebotIntentLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBeebotIntentLgfResponse
func (client *Client) ListBeebotIntentLgfWithContext(ctx context.Context, request *ListBeebotIntentLgfRequest, runtime *dara.RuntimeOptions) (_result *ListBeebotIntentLgfResponse, _err error) {
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

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.LgfText) {
		query["LgfText"] = request.LgfText
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBeebotIntentLgf"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBeebotIntentLgfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListBeebotIntentUserSay
//
// @param request - ListBeebotIntentUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBeebotIntentUserSayResponse
func (client *Client) ListBeebotIntentUserSayWithContext(ctx context.Context, request *ListBeebotIntentUserSayRequest, runtime *dara.RuntimeOptions) (_result *ListBeebotIntentUserSayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBeebotIntentUserSay"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBeebotIntentUserSayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of public cloud instances of Cloud XiaoMi bots.
//
// @param request - ListChatbotInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChatbotInstancesResponse
func (client *Client) ListChatbotInstancesWithContext(ctx context.Context, request *ListChatbotInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListChatbotInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
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
		Action:      dara.String("ListChatbotInstances"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChatbotInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of conversation flows under a script in the legacy canvas scenario.
//
// @param request - ListDialogueFlowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDialogueFlowsResponse
func (client *Client) ListDialogueFlowsWithContext(ctx context.Context, request *ListDialogueFlowsRequest, runtime *dara.RuntimeOptions) (_result *ListDialogueFlowsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDialogueFlows"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDialogueFlowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View the Download Hub list.
//
// @param request - ListDownloadTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownloadTasksResponse
func (client *Client) ListDownloadTasksWithContext(ctx context.Context, request *ListDownloadTasksRequest, runtime *dara.RuntimeOptions) (_result *ListDownloadTasksResponse, _err error) {
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
		Action:      dara.String("ListDownloadTasks"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDownloadTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use ListFlashSmsTemplates to retrieve flash SMS templates.
//
// @param request - ListFlashSmsTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlashSmsTemplatesResponse
func (client *Client) ListFlashSmsTemplatesWithContext(ctx context.Context, request *ListFlashSmsTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListFlashSmsTemplatesResponse, _err error) {
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
		Action:      dara.String("ListFlashSmsTemplates"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlashSmsTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of key questions. This is an old canvas API.
//
// @param request - ListGlobalQuestionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGlobalQuestionsResponse
func (client *Client) ListGlobalQuestionsWithContext(ctx context.Context, request *ListGlobalQuestionsRequest, runtime *dara.RuntimeOptions) (_result *ListGlobalQuestionsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGlobalQuestions"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGlobalQuestionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of Intelligent Outbound Calling business instances.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query intents in the annotation center.
//
// @param request - ListIntentionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntentionsResponse
func (client *Client) ListIntentionsWithContext(ctx context.Context, request *ListIntentionsRequest, runtime *dara.RuntimeOptions) (_result *ListIntentionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnnotationMissionDataSourceType) {
		query["AnnotationMissionDataSourceType"] = request.AnnotationMissionDataSourceType
	}

	if !dara.IsNil(request.BotId) {
		query["BotId"] = request.BotId
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.UserNick) {
		query["UserNick"] = request.UserNick
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntentions"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntentionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the intent list. This is a legacy canvas API.
//
// @param request - ListIntentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntentsResponse
func (client *Client) ListIntentsWithContext(ctx context.Context, request *ListIntentsRequest, runtime *dara.RuntimeOptions) (_result *ListIntentsResponse, _err error) {
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

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntents"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the outbound call job list.
//
// @param request - ListJobGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobGroupsResponse
func (client *Client) ListJobGroupsWithContext(ctx context.Context, request *ListJobGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListJobGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsyncQuery) {
		query["AsyncQuery"] = request.AsyncQuery
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupStatusFilter) {
		query["JobGroupStatusFilter"] = request.JobGroupStatusFilter
	}

	if !dara.IsNil(request.OnlyMinConcurrencyEnabled) {
		query["OnlyMinConcurrencyEnabled"] = request.OnlyMinConcurrencyEnabled
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchText) {
		query["SearchText"] = request.SearchText
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobGroups"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously retrieve the outbound call job list. When the Async parameter is specified in the ListJobGroup API request, use this API to obtain the asynchronous ListJobGroup result.
//
// @param request - ListJobGroupsAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobGroupsAsyncResponse
func (client *Client) ListJobGroupsAsyncWithContext(ctx context.Context, request *ListJobGroupsAsyncRequest, runtime *dara.RuntimeOptions) (_result *ListJobGroupsAsyncResponse, _err error) {
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
		Action:      dara.String("ListJobGroupsAsync"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobGroupsAsyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query multiple jobs.
//
// @param request - ListJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithContext(ctx context.Context, request *ListJobsRequest, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query jobs by job group.
//
// @param request - ListJobsByGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsByGroupResponse
func (client *Client) ListJobsByGroupWithContext(ctx context.Context, request *ListJobsByGroupRequest, runtime *dara.RuntimeOptions) (_result *ListJobsByGroupResponse, _err error) {
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

	if !dara.IsNil(request.JobFailureReason) {
		query["JobFailureReason"] = request.JobFailureReason
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.JobStatus) {
		query["JobStatus"] = request.JobStatus
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
		Action:      dara.String("ListJobsByGroup"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsByGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of outbound call caller numbers.
//
// @param request - ListOutboundCallNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOutboundCallNumbersResponse
func (client *Client) ListOutboundCallNumbersWithContext(ctx context.Context, request *ListOutboundCallNumbersRequest, runtime *dara.RuntimeOptions) (_result *ListOutboundCallNumbersResponse, _err error) {
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
		Action:      dara.String("ListOutboundCallNumbers"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOutboundCallNumbersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query tags of instance business resources.
//
// @param request - ListResourceTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTagsResponse
func (client *Client) ListResourceTagsWithContext(ctx context.Context, request *ListResourceTagsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceTagsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceTags"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of outbound call publishing history.
//
// @param request - ListScriptPublishHistoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptPublishHistoriesResponse
func (client *Client) ListScriptPublishHistoriesWithContext(ctx context.Context, request *ListScriptPublishHistoriesRequest, runtime *dara.RuntimeOptions) (_result *ListScriptPublishHistoriesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScriptPublishHistories"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptPublishHistoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of recording files for use in small model scenarios.
//
// @param request - ListScriptRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptRecordingResponse
func (client *Client) ListScriptRecordingWithContext(ctx context.Context, request *ListScriptRecordingRequest, runtime *dara.RuntimeOptions) (_result *ListScriptRecordingResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RefIdsJson) {
		query["RefIdsJson"] = request.RefIdsJson
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.Search) {
		query["Search"] = request.Search
	}

	if !dara.IsNil(request.StatesJson) {
		query["StatesJson"] = request.StatesJson
	}

	if !dara.IsNil(request.UuidsJson) {
		query["UuidsJson"] = request.UuidsJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScriptRecording"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptRecordingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of scenario voices in the old canvas scenario.
//
// Description:
//
// ***
//
// @param request - ListScriptVoiceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptVoiceConfigsResponse
func (client *Client) ListScriptVoiceConfigsWithContext(ctx context.Context, request *ListScriptVoiceConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListScriptVoiceConfigsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScriptVoiceConfigs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptVoiceConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the script list.
//
// @param request - ListScriptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptsResponse
func (client *Client) ListScriptsWithContext(ctx context.Context, request *ListScriptsRequest, runtime *dara.RuntimeOptions) (_result *ListScriptsResponse, _err error) {
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

	if !dara.IsNil(request.NluEngine) {
		query["NluEngine"] = request.NluEngine
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScriptName) {
		query["ScriptName"] = request.ScriptName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScripts"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the relationship between resources and tags.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query tags under a script, legacy canvas API.
//
// @param request - ListTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithContext(ctx context.Context, request *ListTagsRequest, runtime *dara.RuntimeOptions) (_result *ListTagsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTags"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the agent configuration.
//
// @param tmpReq - ModifyAgentProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAgentProfileResponse
func (client *Client) ModifyAgentProfileWithContext(ctx context.Context, tmpReq *ModifyAgentProfileRequest, runtime *dara.RuntimeOptions) (_result *ModifyAgentProfileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAgentProfileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FaqCategoryIds) {
		request.FaqCategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FaqCategoryIds, dara.String("FaqCategoryIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentProfileId) {
		body["AgentProfileId"] = request.AgentProfileId
	}

	if !dara.IsNil(request.ApiPluginJson) {
		body["ApiPluginJson"] = request.ApiPluginJson
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FaqCategoryIdsShrink) {
		body["FaqCategoryIds"] = request.FaqCategoryIdsShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstructionJson) {
		body["InstructionJson"] = request.InstructionJson
	}

	if !dara.IsNil(request.LabelsJson) {
		body["LabelsJson"] = request.LabelsJson
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.ModelConfig) {
		body["ModelConfig"] = request.ModelConfig
	}

	if !dara.IsNil(request.Prompt) {
		body["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.PromptJson) {
		body["PromptJson"] = request.PromptJson
	}

	if !dara.IsNil(request.Scenario) {
		body["Scenario"] = request.Scenario
	}

	if !dara.IsNil(request.VariablesJson) {
		body["VariablesJson"] = request.VariablesJson
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAgentProfile"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAgentProfileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the status and name of an annotation job.
//
// @param request - ModifyAnnotationMissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAnnotationMissionResponse
func (client *Client) ModifyAnnotationMissionWithContext(ctx context.Context, request *ModifyAnnotationMissionRequest, runtime *dara.RuntimeOptions) (_result *ModifyAnnotationMissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnnotationMissionId) {
		query["AnnotationMissionId"] = request.AnnotationMissionId
	}

	if !dara.IsNil(request.AnnotationMissionName) {
		query["AnnotationMissionName"] = request.AnnotationMissionName
	}

	if !dara.IsNil(request.AnnotationStatus) {
		query["AnnotationStatus"] = request.AnnotationStatus
	}

	if !dara.IsNil(request.Delete) {
		query["Delete"] = request.Delete
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAnnotationMission"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAnnotationMissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a job (including the job itself). This operation is deprecated.
//
// @param request - ModifyBatchJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBatchJobsResponse
func (client *Client) ModifyBatchJobsWithContext(ctx context.Context, request *ModifyBatchJobsRequest, runtime *dara.RuntimeOptions) (_result *ModifyBatchJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchJobName) {
		query["BatchJobName"] = request.BatchJobName
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobFilePath) {
		query["JobFilePath"] = request.JobFilePath
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.StrategyJson) {
		query["StrategyJson"] = request.StrategyJson
	}

	if !dara.IsNil(request.Submitted) {
		query["Submitted"] = request.Submitted
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBatchJobs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBatchJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModifyBeebotIntent
//
// @param tmpReq - ModifyBeebotIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBeebotIntentResponse
func (client *Client) ModifyBeebotIntentWithContext(ctx context.Context, tmpReq *ModifyBeebotIntentRequest, runtime *dara.RuntimeOptions) (_result *ModifyBeebotIntentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyBeebotIntentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IntentDefinition) {
		request.IntentDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentDefinition, dara.String("IntentDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IntentDefinitionShrink) {
		query["IntentDefinition"] = request.IntentDefinitionShrink
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBeebotIntent"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBeebotIntentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModifyBeebotIntentLgf
//
// @param tmpReq - ModifyBeebotIntentLgfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBeebotIntentLgfResponse
func (client *Client) ModifyBeebotIntentLgfWithContext(ctx context.Context, tmpReq *ModifyBeebotIntentLgfRequest, runtime *dara.RuntimeOptions) (_result *ModifyBeebotIntentLgfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyBeebotIntentLgfShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LgfDefinition) {
		request.LgfDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LgfDefinition, dara.String("LgfDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LgfDefinitionShrink) {
		query["LgfDefinition"] = request.LgfDefinitionShrink
	}

	if !dara.IsNil(request.LgfId) {
		query["LgfId"] = request.LgfId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBeebotIntentLgf"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBeebotIntentLgfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModifyBeebotIntentUserSay
//
// @param tmpReq - ModifyBeebotIntentUserSayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBeebotIntentUserSayResponse
func (client *Client) ModifyBeebotIntentUserSayWithContext(ctx context.Context, tmpReq *ModifyBeebotIntentUserSayRequest, runtime *dara.RuntimeOptions) (_result *ModifyBeebotIntentUserSayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyBeebotIntentUserSayShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserSayDefinition) {
		request.UserSayDefinitionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserSayDefinition, dara.String("UserSayDefinition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.UserSayDefinitionShrink) {
		query["UserSayDefinition"] = request.UserSayDefinitionShrink
	}

	if !dara.IsNil(request.UserSayId) {
		query["UserSayId"] = request.UserSayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBeebotIntentUserSay"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBeebotIntentUserSayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify a conversation flow using the legacy canvas API.
//
// @param request - ModifyDialogueFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDialogueFlowResponse
func (client *Client) ModifyDialogueFlowWithContext(ctx context.Context, request *ModifyDialogueFlowRequest, runtime *dara.RuntimeOptions) (_result *ModifyDialogueFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DialogueFlowDefinition) {
		query["DialogueFlowDefinition"] = request.DialogueFlowDefinition
	}

	if !dara.IsNil(request.DialogueFlowId) {
		query["DialogueFlowId"] = request.DialogueFlowId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsDrafted) {
		query["IsDrafted"] = request.IsDrafted
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDialogueFlow"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDialogueFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Save the toggle for the feature that stops outbound calls to nonexistent numbers at the global dimension.
//
// @param request - ModifyEmptyNumberNoMoreCallsInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEmptyNumberNoMoreCallsInfoResponse
func (client *Client) ModifyEmptyNumberNoMoreCallsInfoWithContext(ctx context.Context, request *ModifyEmptyNumberNoMoreCallsInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyEmptyNumberNoMoreCallsInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EmptyNumberNoMoreCalls) {
		query["EmptyNumberNoMoreCalls"] = request.EmptyNumberNoMoreCalls
	}

	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEmptyNumberNoMoreCallsInfo"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEmptyNumberNoMoreCallsInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify a global question in the legacy canvas API.
//
// @param request - ModifyGlobalQuestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalQuestionResponse
func (client *Client) ModifyGlobalQuestionWithContext(ctx context.Context, request *ModifyGlobalQuestionRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Answers) {
		query["Answers"] = request.Answers
	}

	if !dara.IsNil(request.GlobalQuestionId) {
		query["GlobalQuestionId"] = request.GlobalQuestionId
	}

	if !dara.IsNil(request.GlobalQuestionName) {
		query["GlobalQuestionName"] = request.GlobalQuestionName
	}

	if !dara.IsNil(request.GlobalQuestionType) {
		query["GlobalQuestionType"] = request.GlobalQuestionType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Questions) {
		query["Questions"] = request.Questions
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGlobalQuestion"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalQuestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify an Intelligent Outbound Calling business instance.
//
// @param request - ModifyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstanceWithContext(ctx context.Context, request *ModifyInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.InstanceDescription) {
		query["InstanceDescription"] = request.InstanceDescription
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.MaxConcurrentConversation) {
		query["MaxConcurrentConversation"] = request.MaxConcurrentConversation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstance"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify an intent in a legacy canvas scenario using the legacy canvas API.
//
// @param request - ModifyIntentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIntentResponse
func (client *Client) ModifyIntentWithContext(ctx context.Context, request *ModifyIntentRequest, runtime *dara.RuntimeOptions) (_result *ModifyIntentResponse, _err error) {
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

	if !dara.IsNil(request.IntentDescription) {
		query["IntentDescription"] = request.IntentDescription
	}

	if !dara.IsNil(request.IntentId) {
		query["IntentId"] = request.IntentId
	}

	if !dara.IsNil(request.IntentName) {
		query["IntentName"] = request.IntentName
	}

	if !dara.IsNil(request.Keywords) {
		query["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.Utterances) {
		query["Utterances"] = request.Utterances
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIntent"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIntentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify an Intelligent outbound calling job.
//
// @param request - ModifyJobGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyJobGroupResponse
func (client *Client) ModifyJobGroupWithContext(ctx context.Context, request *ModifyJobGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyJobGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FlashSmsExtras) {
		query["FlashSmsExtras"] = request.FlashSmsExtras
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.JobGroupStatus) {
		query["JobGroupStatus"] = request.JobGroupStatus
	}

	if !dara.IsNil(request.MinConcurrency) {
		query["MinConcurrency"] = request.MinConcurrency
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RecallCallingNumber) {
		query["RecallCallingNumber"] = request.RecallCallingNumber
	}

	if !dara.IsNil(request.RecallStrategyJson) {
		query["RecallStrategyJson"] = request.RecallStrategyJson
	}

	if !dara.IsNil(request.RingingDuration) {
		query["RingingDuration"] = request.RingingDuration
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.StrategyJson) {
		query["StrategyJson"] = request.StrategyJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyJobGroup"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyJobGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the outbound call number. This API is deprecated.
//
// @param request - ModifyOutboundCallNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOutboundCallNumberResponse
func (client *Client) ModifyOutboundCallNumberWithContext(ctx context.Context, request *ModifyOutboundCallNumberRequest, runtime *dara.RuntimeOptions) (_result *ModifyOutboundCallNumberResponse, _err error) {
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

	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.OutboundCallNumberId) {
		query["OutboundCallNumberId"] = request.OutboundCallNumberId
	}

	if !dara.IsNil(request.RateLimitCount) {
		query["RateLimitCount"] = request.RateLimitCount
	}

	if !dara.IsNil(request.RateLimitPeriod) {
		query["RateLimitPeriod"] = request.RateLimitPeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOutboundCallNumber"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOutboundCallNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify script information.
//
// @param request - ModifyScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScriptResponse
func (client *Client) ModifyScriptWithContext(ctx context.Context, request *ModifyScriptRequest, runtime *dara.RuntimeOptions) (_result *ModifyScriptResponse, _err error) {
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

	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.AgentLlm) {
		query["AgentLlm"] = request.AgentLlm
	}

	if !dara.IsNil(request.AsrConfig) {
		query["AsrConfig"] = request.AsrConfig
	}

	if !dara.IsNil(request.ChatConfig) {
		query["ChatConfig"] = request.ChatConfig
	}

	if !dara.IsNil(request.ChatbotId) {
		query["ChatbotId"] = request.ChatbotId
	}

	if !dara.IsNil(request.EmotionEnable) {
		query["EmotionEnable"] = request.EmotionEnable
	}

	if !dara.IsNil(request.Industry) {
		query["Industry"] = request.Industry
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LabelConfig) {
		query["LabelConfig"] = request.LabelConfig
	}

	if !dara.IsNil(request.LongWaitEnable) {
		query["LongWaitEnable"] = request.LongWaitEnable
	}

	if !dara.IsNil(request.MiniPlaybackConfigListJsonString) {
		query["MiniPlaybackConfigListJsonString"] = request.MiniPlaybackConfigListJsonString
	}

	if !dara.IsNil(request.MiniPlaybackEnable) {
		query["MiniPlaybackEnable"] = request.MiniPlaybackEnable
	}

	if !dara.IsNil(request.NewBargeInEnable) {
		query["NewBargeInEnable"] = request.NewBargeInEnable
	}

	if !dara.IsNil(request.NlsConfig) {
		query["NlsConfig"] = request.NlsConfig
	}

	if !dara.IsNil(request.NluAccessType) {
		query["NluAccessType"] = request.NluAccessType
	}

	if !dara.IsNil(request.NluEngine) {
		query["NluEngine"] = request.NluEngine
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.ScriptContent) {
		query["ScriptContent"] = request.ScriptContent
	}

	if !dara.IsNil(request.ScriptDescription) {
		query["ScriptDescription"] = request.ScriptDescription
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.ScriptName) {
		query["ScriptName"] = request.ScriptName
	}

	if !dara.IsNil(request.ScriptWaveform) {
		query["ScriptWaveform"] = request.ScriptWaveform
	}

	if !dara.IsNil(request.TtsConfig) {
		query["TtsConfig"] = request.TtsConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScript"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the scene audio configuration in an old canvas scenario.
//
// Description:
//
// ***
//
// @param request - ModifyScriptVoiceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScriptVoiceConfigResponse
func (client *Client) ModifyScriptVoiceConfigWithContext(ctx context.Context, request *ModifyScriptVoiceConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyScriptVoiceConfigResponse, _err error) {
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

	if !dara.IsNil(request.ScriptVoiceConfigId) {
		query["ScriptVoiceConfigId"] = request.ScriptVoiceConfigId
	}

	if !dara.IsNil(request.ScriptWaveformRelation) {
		query["ScriptWaveformRelation"] = request.ScriptWaveformRelation
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScriptVoiceConfig"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScriptVoiceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the TTS configuration for an old canvas scenario.
//
// @param request - ModifyTTSConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTTSConfigResponse
func (client *Client) ModifyTTSConfigWithContext(ctx context.Context, request *ModifyTTSConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyTTSConfigResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NlsServiceType) {
		query["NlsServiceType"] = request.NlsServiceType
	}

	if !dara.IsNil(request.PitchRate) {
		query["PitchRate"] = request.PitchRate
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.SpeechRate) {
		query["SpeechRate"] = request.SpeechRate
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
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTTSConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify a tag group using the legacy canvas API.
//
// @param request - ModifyTagGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTagGroupsResponse
func (client *Client) ModifyTagGroupsWithContext(ctx context.Context, request *ModifyTagGroupsRequest, runtime *dara.RuntimeOptions) (_result *ModifyTagGroupsResponse, _err error) {
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

	if !dara.IsNil(request.TagGroups) {
		query["TagGroups"] = request.TagGroups
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTagGroups"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTagGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publish a scenario.
//
// @param request - PublishScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishScriptResponse
func (client *Client) PublishScriptWithContext(ctx context.Context, request *PublishScriptRequest, runtime *dara.RuntimeOptions) (_result *PublishScriptResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishScript"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publish script (debug version).
//
// @param request - PublishScriptForDebugRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishScriptForDebugResponse
func (client *Client) PublishScriptForDebugWithContext(ctx context.Context, request *PublishScriptForDebugRequest, runtime *dara.RuntimeOptions) (_result *PublishScriptForDebugResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishScriptForDebug"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishScriptForDebugResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the job list.
//
// @param request - QueryJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryJobsResponse
func (client *Client) QueryJobsWithContext(ctx context.Context, request *QueryJobsRequest, runtime *dara.RuntimeOptions) (_result *QueryJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
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

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeAlignment) {
		query["TimeAlignment"] = request.TimeAlignment
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryJobs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query jobs by task result.
//
// @param request - QueryJobsWithResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryJobsWithResultResponse
func (client *Client) QueryJobsWithResultWithContext(ctx context.Context, request *QueryJobsWithResultRequest, runtime *dara.RuntimeOptions) (_result *QueryJobsWithResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndActualTimeFilter) {
		query["EndActualTimeFilter"] = request.EndActualTimeFilter
	}

	if !dara.IsNil(request.HasAnsweredFilter) {
		query["HasAnsweredFilter"] = request.HasAnsweredFilter
	}

	if !dara.IsNil(request.HasHangUpByRejectionFilter) {
		query["HasHangUpByRejectionFilter"] = request.HasHangUpByRejectionFilter
	}

	if !dara.IsNil(request.HasReachedEndOfFlowFilter) {
		query["HasReachedEndOfFlowFilter"] = request.HasReachedEndOfFlowFilter
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobFailureReasonsFilter) {
		query["JobFailureReasonsFilter"] = request.JobFailureReasonsFilter
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.JobStatusFilter) {
		query["JobStatusFilter"] = request.JobStatusFilter
	}

	if !dara.IsNil(request.LabelsJson) {
		query["LabelsJson"] = request.LabelsJson
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryText) {
		query["QueryText"] = request.QueryText
	}

	if !dara.IsNil(request.StartActualTimeFilter) {
		query["StartActualTimeFilter"] = request.StartActualTimeFilter
	}

	if !dara.IsNil(request.TaskStatusFilter) {
		query["TaskStatusFilter"] = request.TaskStatusFilter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryJobsWithResult"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryJobsWithResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Legacy canvas API for querying the list of recordings for script texts.
//
// @param request - QueryScriptWaveformsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScriptWaveformsResponse
func (client *Client) QueryScriptWaveformsWithContext(ctx context.Context, request *QueryScriptWaveformsRequest, runtime *dara.RuntimeOptions) (_result *QueryScriptWaveformsResponse, _err error) {
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

	if !dara.IsNil(request.ScriptContent) {
		query["ScriptContent"] = request.ScriptContent
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryScriptWaveforms"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryScriptWaveformsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query scenario information based on scenario status.
//
// @param request - QueryScriptsByStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryScriptsByStatusResponse
func (client *Client) QueryScriptsByStatusWithContext(ctx context.Context, request *QueryScriptsByStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryScriptsByStatusResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryScriptsByStatus"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryScriptsByStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CC submits the reason for call failure to the outbound calling Operational System. This API is deprecated.
//
// @param request - RecordFailureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecordFailureResponse
func (client *Client) RecordFailureWithContext(ctx context.Context, request *RecordFailureRequest, runtime *dara.RuntimeOptions) (_result *RecordFailureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActualTime) {
		query["ActualTime"] = request.ActualTime
	}

	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.DispositionCode) {
		query["DispositionCode"] = request.DispositionCode
	}

	if !dara.IsNil(request.ExceptionCodes) {
		query["ExceptionCodes"] = request.ExceptionCodes
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecordFailure"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecordFailureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restart a paused job.
//
// @param request - ResumeJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeJobsResponse
func (client *Client) ResumeJobsWithContext(ctx context.Context, request *ResumeJobsRequest, runtime *dara.RuntimeOptions) (_result *ResumeJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobReferenceId) {
		query["JobReferenceId"] = request.JobReferenceId
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeJobs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Roll back a published small model scenario.
//
// @param request - RollbackScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackScriptResponse
func (client *Client) RollbackScriptWithContext(ctx context.Context, request *RollbackScriptRequest, runtime *dara.RuntimeOptions) (_result *RollbackScriptResponse, _err error) {
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

	if !dara.IsNil(request.RollbackVersion) {
		query["RollbackVersion"] = request.RollbackVersion
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackScript"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Save the delay playback time after answering.
//
// @param request - SaveAfterAnswerDelayPlaybackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAfterAnswerDelayPlaybackResponse
func (client *Client) SaveAfterAnswerDelayPlaybackWithContext(ctx context.Context, request *SaveAfterAnswerDelayPlaybackRequest, runtime *dara.RuntimeOptions) (_result *SaveAfterAnswerDelayPlaybackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AfterAnswerDelayPlayback) {
		query["AfterAnswerDelayPlayback"] = request.AfterAnswerDelayPlayback
	}

	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveAfterAnswerDelayPlayback"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveAfterAnswerDelayPlaybackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// [parameters_AnnotationMissionSessionList_schema_items_properties_AnnotationMissionSessionId_type]string
//
// @param request - SaveAnnotationMissionSessionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAnnotationMissionSessionListResponse
func (client *Client) SaveAnnotationMissionSessionListWithContext(ctx context.Context, request *SaveAnnotationMissionSessionListRequest, runtime *dara.RuntimeOptions) (_result *SaveAnnotationMissionSessionListResponse, _err error) {
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

	if !dara.IsNil(request.AgentKey) {
		query["AgentKey"] = request.AgentKey
	}

	if !dara.IsNil(request.AnnotationMissionDataSourceType) {
		query["AnnotationMissionDataSourceType"] = request.AnnotationMissionDataSourceType
	}

	if !dara.IsNil(request.AnnotationMissionSessionList) {
		query["AnnotationMissionSessionList"] = request.AnnotationMissionSessionList
	}

	if !dara.IsNil(request.AnnotationMissionSessionListJsonString) {
		query["AnnotationMissionSessionListJsonString"] = request.AnnotationMissionSessionListJsonString
	}

	if !dara.IsNil(request.Environment) {
		query["Environment"] = request.Environment
	}

	if !dara.IsNil(request.UserNick) {
		query["UserNick"] = request.UserNick
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveAnnotationMissionSessionList"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveAnnotationMissionSessionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Save instance tags in batch.
//
// @param request - SaveAnnotationMissionTagInfoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAnnotationMissionTagInfoListResponse
func (client *Client) SaveAnnotationMissionTagInfoListWithContext(ctx context.Context, request *SaveAnnotationMissionTagInfoListRequest, runtime *dara.RuntimeOptions) (_result *SaveAnnotationMissionTagInfoListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnnotationMissionTagInfoList) {
		query["AnnotationMissionTagInfoList"] = request.AnnotationMissionTagInfoList
	}

	if !dara.IsNil(request.AnnotationMissionTagInfoListJsonString) {
		query["AnnotationMissionTagInfoListJsonString"] = request.AnnotationMissionTagInfoListJsonString
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Reset) {
		query["Reset"] = request.Reset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveAnnotationMissionTagInfoList"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveAnnotationMissionTagInfoListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edit the global callable time segments in system administration.
//
// @param request - SaveBaseStrategyPeriodRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveBaseStrategyPeriodResponse
func (client *Client) SaveBaseStrategyPeriodWithContext(ctx context.Context, request *SaveBaseStrategyPeriodRequest, runtime *dara.RuntimeOptions) (_result *SaveBaseStrategyPeriodResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.OnlyWeekdays) {
		query["OnlyWeekdays"] = request.OnlyWeekdays
	}

	if !dara.IsNil(request.OnlyWorkdays) {
		query["OnlyWorkdays"] = request.OnlyWorkdays
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	if !dara.IsNil(request.WorkingTime) {
		query["WorkingTime"] = request.WorkingTime
	}

	if !dara.IsNil(request.WorkingTimeFramesJson) {
		query["WorkingTimeFramesJson"] = request.WorkingTimeFramesJson
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveBaseStrategyPeriod"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveBaseStrategyPeriodResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a "Do Not Call List" for the instance.
//
// @param request - SaveContactBlockListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveContactBlockListResponse
func (client *Client) SaveContactBlockListWithContext(ctx context.Context, request *SaveContactBlockListRequest, runtime *dara.RuntimeOptions) (_result *SaveContactBlockListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactBlockListList) {
		query["ContactBlockListList"] = request.ContactBlockListList
	}

	if !dara.IsNil(request.ContactBlockListsJson) {
		query["ContactBlockListsJson"] = request.ContactBlockListsJson
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveContactBlockList"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveContactBlockListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Save the outbound call whitelist.
//
// @param request - SaveContactWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveContactWhiteListResponse
func (client *Client) SaveContactWhiteListWithContext(ctx context.Context, request *SaveContactWhiteListRequest, runtime *dara.RuntimeOptions) (_result *SaveContactWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactWhiteListList) {
		query["ContactWhiteListList"] = request.ContactWhiteListList
	}

	if !dara.IsNil(request.ContactWhiteListsJson) {
		query["ContactWhiteListsJson"] = request.ContactWhiteListsJson
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveContactWhiteList"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveContactWhiteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Save the validity period of the basic policy. This operation is deprecated.
//
// @param request - SaveEffectiveDaysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveEffectiveDaysResponse
func (client *Client) SaveEffectiveDaysWithContext(ctx context.Context, request *SaveEffectiveDaysRequest, runtime *dara.RuntimeOptions) (_result *SaveEffectiveDaysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveDays) {
		query["EffectiveDays"] = request.EffectiveDays
	}

	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveEffectiveDays"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveEffectiveDaysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Save the daily call limit for the called number.
//
// @param request - SaveMaxAttemptsPerDayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveMaxAttemptsPerDayResponse
func (client *Client) SaveMaxAttemptsPerDayWithContext(ctx context.Context, request *SaveMaxAttemptsPerDayRequest, runtime *dara.RuntimeOptions) (_result *SaveMaxAttemptsPerDayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntryId) {
		query["EntryId"] = request.EntryId
	}

	if !dara.IsNil(request.MaxAttemptsPerDay) {
		query["MaxAttemptsPerDay"] = request.MaxAttemptsPerDay
	}

	if !dara.IsNil(request.StrategyLevel) {
		query["StrategyLevel"] = request.StrategyLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveMaxAttemptsPerDay"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveMaxAttemptsPerDayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search for jobs.
//
// @param request - SearchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTaskResponse
func (client *Client) SearchTaskWithContext(ctx context.Context, request *SearchTaskRequest, runtime *dara.RuntimeOptions) (_result *SearchTaskResponse, _err error) {
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
		Action:      dara.String("SearchTask"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiate a call directly without going through the CDN mapping system. This API is available only to whitelist users.
//
// @param request - StartJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartJobResponse
func (client *Client) StartJobWithContext(ctx context.Context, request *StartJobRequest, runtime *dara.RuntimeOptions) (_result *StartJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallingNumber) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.JobJson) {
		query["JobJson"] = request.JobJson
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartJob"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submit a job group in Draft status for execution. This operation is deprecated.
//
// @param request - SubmitBatchJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitBatchJobsResponse
func (client *Client) SubmitBatchJobsWithContext(ctx context.Context, request *SubmitBatchJobsRequest, runtime *dara.RuntimeOptions) (_result *SubmitBatchJobsResponse, _err error) {
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

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitBatchJobs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitBatchJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submit call recordings to outbound calls. This API is deprecated.
//
// @param request - SubmitRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitRecordingResponse
func (client *Client) SubmitRecordingWithContext(ctx context.Context, request *SubmitRecordingRequest, runtime *dara.RuntimeOptions) (_result *SubmitRecordingResponse, _err error) {
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

	if !dara.IsNil(request.MergedRecording) {
		query["MergedRecording"] = request.MergedRecording
	}

	if !dara.IsNil(request.ResourceRecording) {
		query["ResourceRecording"] = request.ResourceRecording
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitRecording"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitRecordingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submit a scenario for publish review.
//
// @param request - SubmitScriptReviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitScriptReviewResponse
func (client *Client) SubmitScriptReviewWithContext(ctx context.Context, request *SubmitScriptReviewRequest, runtime *dara.RuntimeOptions) (_result *SubmitScriptReviewResponse, _err error) {
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

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitScriptReview"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitScriptReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stop calls based on the list data.
//
// @param request - SuspendCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendCallResponse
func (client *Client) SuspendCallWithContext(ctx context.Context, request *SuspendCallRequest, runtime *dara.RuntimeOptions) (_result *SuspendCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumbers) {
		query["CalledNumbers"] = request.CalledNumbers
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendCall"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stop calling based on file data.
//
// @param request - SuspendCallWithFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendCallWithFileResponse
func (client *Client) SuspendCallWithFileWithContext(ctx context.Context, request *SuspendCallWithFileRequest, runtime *dara.RuntimeOptions) (_result *SuspendCallWithFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendCallWithFile"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendCallWithFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pause a job.
//
// @param request - SuspendJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendJobsResponse
func (client *Client) SuspendJobsWithContext(ctx context.Context, request *SuspendJobsRequest, runtime *dara.RuntimeOptions) (_result *SuspendJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobGroupId) {
		query["JobGroupId"] = request.JobGroupId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobReferenceId) {
		query["JobReferenceId"] = request.JobReferenceId
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendJobs"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add tags to instance resources.
//
// Description:
//
// *
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prepare to execute the job.
//
// @param request - TaskPreparingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskPreparingResponse
func (client *Client) TaskPreparingWithContext(ctx context.Context, request *TaskPreparingRequest, runtime *dara.RuntimeOptions) (_result *TaskPreparingResponse, _err error) {
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

	if !dara.IsNil(request.InstanceOwnerId) {
		query["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TaskPreparing"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TaskPreparingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stop a call during the call procedure.
//
// @param request - TerminateCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateCallResponse
func (client *Client) TerminateCallWithContext(ctx context.Context, request *TerminateCallRequest, runtime *dara.RuntimeOptions) (_result *TerminateCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TerminateCall"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminateCallResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete resource tags.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upload a recording file.
//
// @param request - UploadScriptRecordingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadScriptRecordingResponse
func (client *Client) UploadScriptRecordingWithContext(ctx context.Context, request *UploadScriptRecordingRequest, runtime *dara.RuntimeOptions) (_result *UploadScriptRecordingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadScriptRecording"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadScriptRecordingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revoke script review.
//
// Description:
//
// ***
//
// @param request - WithdrawScriptReviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawScriptReviewResponse
func (client *Client) WithdrawScriptReviewWithContext(ctx context.Context, request *WithdrawScriptReviewRequest, runtime *dara.RuntimeOptions) (_result *WithdrawScriptReviewResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WithdrawScriptReview"),
		Version:     dara.String("2019-12-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &WithdrawScriptReviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
