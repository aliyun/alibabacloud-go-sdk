// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a client rule policy.
//
// Description:
//
// Creates a client rule policy. The policy type is specified by PolicyType.
//
// @param tmpReq - CreateClientRuleStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientRuleStrategyResponse
func (client *Client) CreateClientRuleStrategyWithContext(ctx context.Context, tmpReq *CreateClientRuleStrategyRequest, runtime *dara.RuntimeOptions) (_result *CreateClientRuleStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateClientRuleStrategyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExcludeUserNameList) {
		request.ExcludeUserNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeUserNameList, dara.String("ExcludeUserNameList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IncludeGroupIdList) {
		request.IncludeGroupIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IncludeGroupIdList, dara.String("IncludeGroupIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PortOnList) {
		request.PortOnListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PortOnList, dara.String("PortOnList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SoftWareOnList) {
		request.SoftWareOnListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SoftWareOnList, dara.String("SoftWareOnList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.AllGroup) {
		query["AllGroup"] = request.AllGroup
	}

	if !dara.IsNil(request.ApprovalId) {
		query["ApprovalId"] = request.ApprovalId
	}

	if !dara.IsNil(request.Cmdline) {
		query["Cmdline"] = request.Cmdline
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeUserNameListShrink) {
		query["ExcludeUserNameList"] = request.ExcludeUserNameListShrink
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.InSwitch) {
		query["InSwitch"] = request.InSwitch
	}

	if !dara.IsNil(request.IncludeGroupIdListShrink) {
		query["IncludeGroupIdList"] = request.IncludeGroupIdListShrink
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.MainButtonTextCh) {
		query["MainButtonTextCh"] = request.MainButtonTextCh
	}

	if !dara.IsNil(request.MainButtonTextEn) {
		query["MainButtonTextEn"] = request.MainButtonTextEn
	}

	if !dara.IsNil(request.MatchMode) {
		query["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.MatchTargetIds) {
		query["MatchTargetIds"] = request.MatchTargetIds
	}

	if !dara.IsNil(request.Md5List) {
		query["Md5List"] = request.Md5List
	}

	if !dara.IsNil(request.MinorButtonTextCh) {
		query["MinorButtonTextCh"] = request.MinorButtonTextCh
	}

	if !dara.IsNil(request.MinorButtonTextEn) {
		query["MinorButtonTextEn"] = request.MinorButtonTextEn
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NewFilePath) {
		query["NewFilePath"] = request.NewFilePath
	}

	if !dara.IsNil(request.OutSwitch) {
		query["OutSwitch"] = request.OutSwitch
	}

	if !dara.IsNil(request.ParentCmdline) {
		query["ParentCmdline"] = request.ParentCmdline
	}

	if !dara.IsNil(request.ParentProcPath) {
		query["ParentProcPath"] = request.ParentProcPath
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PortOnListShrink) {
		query["PortOnList"] = request.PortOnListShrink
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ProcPath) {
		query["ProcPath"] = request.ProcPath
	}

	if !dara.IsNil(request.RegistryContent) {
		query["RegistryContent"] = request.RegistryContent
	}

	if !dara.IsNil(request.RegistryKey) {
		query["RegistryKey"] = request.RegistryKey
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.SelectionKey) {
		query["SelectionKey"] = request.SelectionKey
	}

	if !dara.IsNil(request.SoftWareOnListShrink) {
		query["SoftWareOnList"] = request.SoftWareOnListShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StrategySourceType) {
		query["StrategySourceType"] = request.StrategySourceType
	}

	if !dara.IsNil(request.SupportsApproval) {
		query["SupportsApproval"] = request.SupportsApproval
	}

	if !dara.IsNil(request.UsContent) {
		query["UsContent"] = request.UsContent
	}

	if !dara.IsNil(request.UsTitle) {
		query["UsTitle"] = request.UsTitle
	}

	if !dara.IsNil(request.ZhContent) {
		query["ZhContent"] = request.ZhContent
	}

	if !dara.IsNil(request.ZhTitle) {
		query["ZhTitle"] = request.ZhTitle
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClientRuleStrategy"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClientRuleStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a general asset selection configuration and generates a SelectionKey for use in What to do next.
//
// Description:
//
// This operation is used for the initialization of a general asset selection configuration and generates a SelectionKey. The addition or selection of specific assets is completed in What to do next. When you invoke this operation, it only generates the configuration object and does not save the configuration record.
//
// @param request - CreateSelectionConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSelectionConfigResponse
func (client *Client) CreateSelectionConfigWithContext(ctx context.Context, request *CreateSelectionConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSelectionConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSelectionConfig"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSelectionConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a node to query alerting events triggered by the same rule or Alarm Metric.
//
// @param request - CreateSimilarSecurityEventsQueryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSimilarSecurityEventsQueryTaskResponse
func (client *Client) CreateSimilarSecurityEventsQueryTaskWithContext(ctx context.Context, request *CreateSimilarSecurityEventsQueryTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSimilarSecurityEventsQueryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityEventId) {
		query["SecurityEventId"] = request.SecurityEventId
	}

	if !dara.IsNil(request.SimilarEventScenarioCode) {
		query["SimilarEventScenarioCode"] = request.SimilarEventScenarioCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSimilarSecurityEventsQueryTask"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSimilarSecurityEventsQueryTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes client rule policies based on a list of policy IDs.
//
// @param tmpReq - DeleteClientRuleStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientRuleStrategyResponse
func (client *Client) DeleteClientRuleStrategyWithContext(ctx context.Context, tmpReq *DeleteClientRuleStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientRuleStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteClientRuleStrategyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IdList) {
		request.IdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IdList, dara.String("IdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IdListShrink) {
		query["IdList"] = request.IdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClientRuleStrategy"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientRuleStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes alert whitelist configurations in batches.
//
// @param request - DeleteSecurityEventMarkMissListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecurityEventMarkMissListResponse
func (client *Client) DeleteSecurityEventMarkMissListWithContext(ctx context.Context, request *DeleteSecurityEventMarkMissListRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecurityEventMarkMissListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecurityEventMarkMissList"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecurityEventMarkMissListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the effective scope mode, target list, and scope record information based on the asset scope ID.
//
// @param request - DescribeAssetScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetScopeResponse
func (client *Client) DescribeAssetScopeWithContext(ctx context.Context, request *DescribeAssetScopeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAssetScopeResponse, _err error) {
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
		Action:      dara.String("DescribeAssetScope"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAssetScopeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries security alert types.
//
// @param request - DescribeNsasSuspEventTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNsasSuspEventTypeResponse
func (client *Client) DescribeNsasSuspEventTypeWithContext(ctx context.Context, request *DescribeNsasSuspEventTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeNsasSuspEventTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContainerFieldName) {
		query["ContainerFieldName"] = request.ContainerFieldName
	}

	if !dara.IsNil(request.ContainerFieldValue) {
		query["ContainerFieldValue"] = request.ContainerFieldValue
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.Uuids) {
		query["Uuids"] = request.Uuids
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNsasSuspEventType"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNsasSuspEventTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the automatic whitelist rules for alerts.
//
// @param request - DescribeSecurityEventMarkMissListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityEventMarkMissListResponse
func (client *Client) DescribeSecurityEventMarkMissListWithContext(ctx context.Context, request *DescribeSecurityEventMarkMissListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityEventMarkMissListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DisposalWay) {
		query["DisposalWay"] = request.DisposalWay
	}

	if !dara.IsNil(request.EventName) {
		query["EventName"] = request.EventName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityEventMarkMissList"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityEventMarkMissListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the processing status of security alerts based on the task ID or alert event ID.
//
// @param request - DescribeSecurityEventOperationStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityEventOperationStatusResponse
func (client *Client) DescribeSecurityEventOperationStatusWithContext(ctx context.Context, request *DescribeSecurityEventOperationStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityEventOperationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityEventIds) {
		query["SecurityEventIds"] = request.SecurityEventIds
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityEventOperationStatus"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityEventOperationStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the handling operations, operability, and whitelist rule configurations for a specified endpoint alert in its current state.
//
// Description:
//
// This operation queries the handling operations that can be displayed for a specified endpoint alert. When initiating a subsequent response, you must select the operation items returned in this response where UserCanOperate is true, and submit parameters as required by the response operation. The query result does not guarantee that the subsequent response will succeed.
//
// The actual returned results are affected by feature switches, canary releases, dynamic allowlists, alert templates, alert statuses, and filter conditions. The list may be empty. The examples are used to illustrate field formats or rule combinations and do not represent that a specific alert will definitely return these values. MappingMarkFields and its subfields are generic compatibility structures. The current query pipeline does not populate their elements, and you should not infer that this operation supports the corresponding capabilities based on them.
//
// @param request - DescribeSecurityEventOperationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityEventOperationsResponse
func (client *Client) DescribeSecurityEventOperationsWithContext(ctx context.Context, request *DescribeSecurityEventOperationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityEventId) {
		query["SecurityEventId"] = request.SecurityEventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityEventOperations"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the handling scenarios for alerts triggered by the same rule or type.
//
// @param request - DescribeSimilarEventScenariosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSimilarEventScenariosResponse
func (client *Client) DescribeSimilarEventScenariosWithContext(ctx context.Context, request *DescribeSimilarEventScenariosRequest, runtime *dara.RuntimeOptions) (_result *DescribeSimilarEventScenariosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityEventId) {
		query["SecurityEventId"] = request.SecurityEventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSimilarEventScenarios"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSimilarEventScenariosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries aggregated alerts triggered by the same rule or type.
//
// @param request - DescribeSimilarSecurityEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSimilarSecurityEventsResponse
func (client *Client) DescribeSimilarSecurityEventsWithContext(ctx context.Context, request *DescribeSimilarSecurityEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSimilarSecurityEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSimilarSecurityEvents"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSimilarSecurityEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries endpoint security alerts and supports filtering by alert, endpoint, and processing status.
//
// @param request - DescribeSuspEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSuspEventsResponse
func (client *Client) DescribeSuspEventsWithContext(ctx context.Context, request *DescribeSuspEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSuspEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmEventName) {
		query["AlarmEventName"] = request.AlarmEventName
	}

	if !dara.IsNil(request.AlarmUniqueInfo) {
		query["AlarmUniqueInfo"] = request.AlarmUniqueInfo
	}

	if !dara.IsNil(request.AssetsTypeList) {
		query["AssetsTypeList"] = request.AssetsTypeList
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ContainerFieldName) {
		query["ContainerFieldName"] = request.ContainerFieldName
	}

	if !dara.IsNil(request.ContainerFieldValue) {
		query["ContainerFieldValue"] = request.ContainerFieldValue
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Dealed) {
		query["Dealed"] = request.Dealed
	}

	if !dara.IsNil(request.Department) {
		query["Department"] = request.Department
	}

	if !dara.IsNil(request.EventNames) {
		query["EventNames"] = request.EventNames
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InnerIp) {
		query["InnerIp"] = request.InnerIp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Levels) {
		query["Levels"] = request.Levels
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OperateErrorCodeList) {
		query["OperateErrorCodeList"] = request.OperateErrorCodeList
	}

	if !dara.IsNil(request.OperateTimeEnd) {
		query["OperateTimeEnd"] = request.OperateTimeEnd
	}

	if !dara.IsNil(request.OperateTimeStart) {
		query["OperateTimeStart"] = request.OperateTimeStart
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentEventTypes) {
		query["ParentEventTypes"] = request.ParentEventTypes
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SortColumn) {
		query["SortColumn"] = request.SortColumn
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.SrcIp) {
		query["SrcIp"] = request.SrcIp
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StrictMode) {
		query["StrictMode"] = request.StrictMode
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.TimeEnd) {
		query["TimeEnd"] = request.TimeEnd
	}

	if !dara.IsNil(request.TimeStart) {
		query["TimeStart"] = request.TimeStart
	}

	if !dara.IsNil(request.UniqueInfo) {
		query["UniqueInfo"] = request.UniqueInfo
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.Uuids) {
		query["Uuids"] = request.Uuids
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TacticId) {
		body["TacticId"] = request.TacticId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSuspEvents"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a client rule policy.
//
// Description:
//
// Queries the detailed information of a client rule policy based on the policy ID.
//
// @param request - GetClientRuleStrategyDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientRuleStrategyDetailResponse
func (client *Client) GetClientRuleStrategyDetailWithContext(ctx context.Context, request *GetClientRuleStrategyDetailRequest, runtime *dara.RuntimeOptions) (_result *GetClientRuleStrategyDetailResponse, _err error) {
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
		Action:      dara.String("GetClientRuleStrategyDetail"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientRuleStrategyDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of system rules and custom rules in proactive defense policies.
//
// Description:
//
// Counts the number of proactive defense policies (PolicyType=2) for the current account and returns the number of system rule policies and custom rule policies.
//
// @param request - GetClientRuleStrategyStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientRuleStrategyStatisticsResponse
func (client *Client) GetClientRuleStrategyStatisticsWithContext(ctx context.Context, request *GetClientRuleStrategyStatisticsRequest, runtime *dara.RuntimeOptions) (_result *GetClientRuleStrategyStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClientRuleStrategyStatistics"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientRuleStrategyStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Handles security alert events.
//
// @param request - HandleSecurityEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandleSecurityEventsResponse
func (client *Client) HandleSecurityEventsWithContext(ctx context.Context, request *HandleSecurityEventsRequest, runtime *dara.RuntimeOptions) (_result *HandleSecurityEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MarkBatch) {
		query["MarkBatch"] = request.MarkBatch
	}

	if !dara.IsNil(request.MarkMissParam) {
		query["MarkMissParam"] = request.MarkMissParam
	}

	if !dara.IsNil(request.OperationCode) {
		query["OperationCode"] = request.OperationCode
	}

	if !dara.IsNil(request.OperationParams) {
		query["OperationParams"] = request.OperationParams
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SecurityEventIds) {
		query["SecurityEventIds"] = request.SecurityEventIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HandleSecurityEvents"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs batch operations on similar security alerts filtered by a query task.
//
// Description:
//
// First, call the CreateSimilarSecurityEventsQueryTask operation to create a similar alert query task. Wait until the data extraction for the task is complete and the task reaches the DataRetrieved state before you call this operation. The processing scope covers the security alerts filtered under the TaskId. This scope is determined by the similar scenario and alert configurations selected when the query task was created. The scope does not equal all alerts of the same type, nor is it fixed to alerts with the same IP address. After this operation is successfully called, the processing continues to run asynchronously in the background.
//
// @param request - HandleSimilarSecurityEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandleSimilarSecurityEventsResponse
func (client *Client) HandleSimilarSecurityEventsWithContext(ctx context.Context, request *HandleSimilarSecurityEventsRequest, runtime *dara.RuntimeOptions) (_result *HandleSimilarSecurityEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MarkMissParam) {
		query["MarkMissParam"] = request.MarkMissParam
	}

	if !dara.IsNil(request.OperationCode) {
		query["OperationCode"] = request.OperationCode
	}

	if !dara.IsNil(request.OperationParams) {
		query["OperationParams"] = request.OperationParams
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HandleSimilarSecurityEvents"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HandleSimilarSecurityEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the client rule policy list by using a paged query with paging.
//
// Description:
//
// Performs a paged query to retrieve client rule policies based on the policy type (PolicyType) with paging. You can call GetClientRuleStrategyDetail to query the details.
//
// @param request - ListClientRuleStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientRuleStrategyResponse
func (client *Client) ListClientRuleStrategyWithContext(ctx context.Context, request *ListClientRuleStrategyRequest, runtime *dara.RuntimeOptions) (_result *ListClientRuleStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClientRuleStrategy"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClientRuleStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries system defense rules and the switch status of rules under a specified SelectionKey by paging in a paged query.
//
// Description:
//
// Queries the system rule catalog and the switch status of rules under a specified SelectionKey. You can use the returned SwitchId and the same SelectionKey to call the OperateTargetConfig operation to modify the switch status.
//
// @param tmpReq - ListSystemConfigRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSystemConfigRulesResponse
func (client *Client) ListSystemConfigRulesWithContext(ctx context.Context, tmpReq *ListSystemConfigRulesRequest, runtime *dara.RuntimeOptions) (_result *ListSystemConfigRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSystemConfigRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AggregationIds) {
		request.AggregationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AggregationIds, dara.String("AggregationIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuleTypes) {
		request.RuleTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuleTypes, dara.String("RuleTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AggregationIdsShrink) {
		query["AggregationIds"] = request.AggregationIdsShrink
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleTypesShrink) {
		query["RuleTypes"] = request.RuleTypesShrink
	}

	if !dara.IsNil(request.SelectionKey) {
		query["SelectionKey"] = request.SelectionKey
	}

	if !dara.IsNil(request.SystemType) {
		query["SystemType"] = request.SystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSystemConfigRules"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSystemConfigRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the switch status of system rules in batches by SelectionKey.
//
// Description:
//
// Updates the override configuration of rule switches for each SwitchId under the same SelectionKey. Obtain the rule SwitchId by calling ListSystemConfigRules first. Do not use this operation to replace UpdateClientRuleStrategy for updating the overall policy switch Status, and do not use software process names as system rule SwitchIds.
//
// @param tmpReq - OperateTargetConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateTargetConfigResponse
func (client *Client) OperateTargetConfigWithContext(ctx context.Context, tmpReq *OperateTargetConfigRequest, runtime *dara.RuntimeOptions) (_result *OperateTargetConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &OperateTargetConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SwitchIdList) {
		request.SwitchIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SwitchIdList, dara.String("SwitchIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SelectionKey) {
		query["SelectionKey"] = request.SelectionKey
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SwitchIdListShrink) {
		query["SwitchIdList"] = request.SwitchIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateTargetConfig"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateTargetConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the effective scope mode and target list of a specified asset scope record.
//
// @param request - UpdateAssetScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAssetScopeResponse
func (client *Client) UpdateAssetScopeWithContext(ctx context.Context, request *UpdateAssetScopeRequest, runtime *dara.RuntimeOptions) (_result *UpdateAssetScopeResponse, _err error) {
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

	if !dara.IsNil(request.MatchMode) {
		query["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.MatchTargetIds) {
		query["MatchTargetIds"] = request.MatchTargetIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAssetScope"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAssetScopeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a client rule policy.
//
// Description:
//
// Updates a client rule policy based on the policy ID.
//
// @param tmpReq - UpdateClientRuleStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientRuleStrategyResponse
func (client *Client) UpdateClientRuleStrategyWithContext(ctx context.Context, tmpReq *UpdateClientRuleStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateClientRuleStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateClientRuleStrategyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExcludeUserNameList) {
		request.ExcludeUserNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeUserNameList, dara.String("ExcludeUserNameList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IncludeGroupIdList) {
		request.IncludeGroupIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IncludeGroupIdList, dara.String("IncludeGroupIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PortOnList) {
		request.PortOnListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PortOnList, dara.String("PortOnList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SoftWareOnList) {
		request.SoftWareOnListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SoftWareOnList, dara.String("SoftWareOnList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.AllGroup) {
		query["AllGroup"] = request.AllGroup
	}

	if !dara.IsNil(request.ApprovalId) {
		query["ApprovalId"] = request.ApprovalId
	}

	if !dara.IsNil(request.Cmdline) {
		query["Cmdline"] = request.Cmdline
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExcludeUserNameListShrink) {
		query["ExcludeUserNameList"] = request.ExcludeUserNameListShrink
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InSwitch) {
		query["InSwitch"] = request.InSwitch
	}

	if !dara.IsNil(request.IncludeGroupIdListShrink) {
		query["IncludeGroupIdList"] = request.IncludeGroupIdListShrink
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.MainButtonTextCh) {
		query["MainButtonTextCh"] = request.MainButtonTextCh
	}

	if !dara.IsNil(request.MainButtonTextEn) {
		query["MainButtonTextEn"] = request.MainButtonTextEn
	}

	if !dara.IsNil(request.MatchMode) {
		query["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.MatchTargetIds) {
		query["MatchTargetIds"] = request.MatchTargetIds
	}

	if !dara.IsNil(request.Md5List) {
		query["Md5List"] = request.Md5List
	}

	if !dara.IsNil(request.MinorButtonTextCh) {
		query["MinorButtonTextCh"] = request.MinorButtonTextCh
	}

	if !dara.IsNil(request.MinorButtonTextEn) {
		query["MinorButtonTextEn"] = request.MinorButtonTextEn
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NewFilePath) {
		query["NewFilePath"] = request.NewFilePath
	}

	if !dara.IsNil(request.OutSwitch) {
		query["OutSwitch"] = request.OutSwitch
	}

	if !dara.IsNil(request.ParentCmdline) {
		query["ParentCmdline"] = request.ParentCmdline
	}

	if !dara.IsNil(request.ParentProcPath) {
		query["ParentProcPath"] = request.ParentProcPath
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PortOnListShrink) {
		query["PortOnList"] = request.PortOnListShrink
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ProcPath) {
		query["ProcPath"] = request.ProcPath
	}

	if !dara.IsNil(request.RegistryContent) {
		query["RegistryContent"] = request.RegistryContent
	}

	if !dara.IsNil(request.RegistryKey) {
		query["RegistryKey"] = request.RegistryKey
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.SoftWareOnListShrink) {
		query["SoftWareOnList"] = request.SoftWareOnListShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StrategySourceType) {
		query["StrategySourceType"] = request.StrategySourceType
	}

	if !dara.IsNil(request.SupportsApproval) {
		query["SupportsApproval"] = request.SupportsApproval
	}

	if !dara.IsNil(request.UsContent) {
		query["UsContent"] = request.UsContent
	}

	if !dara.IsNil(request.UsTitle) {
		query["UsTitle"] = request.UsTitle
	}

	if !dara.IsNil(request.ZhContent) {
		query["ZhContent"] = request.ZhContent
	}

	if !dara.IsNil(request.ZhTitle) {
		query["ZhTitle"] = request.ZhTitle
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClientRuleStrategy"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClientRuleStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
