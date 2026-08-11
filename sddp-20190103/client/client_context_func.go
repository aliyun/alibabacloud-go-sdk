// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Modifies the general alert configuration parameters.
//
// Description:
//
// Creates or restores configurations based on the codes of common alert configuration items, allowing you to manage these configurations.
//
// # Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigResponse
func (client *Client) CreateConfigWithContext(ctx context.Context, request *CreateConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateConfigResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConfig"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scan authorization for a database, project, or Object Storage Service (OSS) bucket.
//
// Description:
//
// This operation is used to authorize scanning of data assets to further protect the data security of your data assets.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - CreateDataLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLimitResponse
func (client *Client) CreateDataLimitWithContext(ctx context.Context, request *CreateDataLimitRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLimitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.AutoScan) {
		query["AutoScan"] = request.AutoScan
	}

	if !dara.IsNil(request.CertificatePermission) {
		query["CertificatePermission"] = request.CertificatePermission
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.EventStatus) {
		query["EventStatus"] = request.EventStatus
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.InstantlyScan) {
		query["InstantlyScan"] = request.InstantlyScan
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogStoreDay) {
		query["LogStoreDay"] = request.LogStoreDay
	}

	if !dara.IsNil(request.OcrStatus) {
		query["OcrStatus"] = request.OcrStatus
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SamplingSize) {
		query["SamplingSize"] = request.SamplingSize
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLimit"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLimitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom sensitive data detection rule by calling CreateRule.
//
// @param request - CreateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRuleResponse
func (client *Client) CreateRuleWithContext(ctx context.Context, request *CreateRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentCategory) {
		query["ContentCategory"] = request.ContentCategory
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.ModelRuleIds) {
		query["ModelRuleIds"] = request.ModelRuleIds
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.StatExpress) {
		query["StatExpress"] = request.StatExpress
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SupportForm) {
		query["SupportForm"] = request.SupportForm
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TemplateRuleIds) {
		query["TemplateRuleIds"] = request.TemplateRuleIds
	}

	if !dara.IsNil(request.WarnLevel) {
		query["WarnLevel"] = request.WarnLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRule"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateScanTask operation to create a custom scan task for detecting sensitive data in assets that have been successfully authorized for detection.
//
// Description:
//
// This operation is applicable to users who want to create custom scan tasks for authorized assets. It allows users to flexibly control the interval between scan tasks and the runtime of each scan task.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - CreateScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScanTaskResponse
func (client *Client) CreateScanTaskWithContext(ctx context.Context, request *CreateScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateScanTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataLimitId) {
		query["DataLimitId"] = request.DataLimitId
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.IntervalDay) {
		query["IntervalDay"] = request.IntervalDay
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.OssScanPath) {
		query["OssScanPath"] = request.OssScanPath
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RunHour) {
		query["RunHour"] = request.RunHour
	}

	if !dara.IsNil(request.RunMinute) {
		query["RunMinute"] = request.RunMinute
	}

	if !dara.IsNil(request.ScanRange) {
		query["ScanRange"] = request.ScanRange
	}

	if !dara.IsNil(request.ScanRangeContent) {
		query["ScanRangeContent"] = request.ScanRangeContent
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskUserName) {
		query["TaskUserName"] = request.TaskUserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScanTask"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScanTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call CreateSlrRole to create a service-linked role for Data Security Center (DSC). This role authorizes DSC to access your cloud resources.
//
// Description:
//
// This operation allows DSC to access the resources of Alibaba Cloud services such as OSS, RDS, and MaxCompute. After you call this operation, the system automatically creates a service-linked role. The role is named AliyunServiceRoleForSDDP, and its access policy is AliyunServiceRolePolicyForSDDP.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 for each user. If you exceed the limit, API calls are throttled, which can affect your business. Call this operation at a reasonable rate.
//
// @param request - CreateSlrRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSlrRoleResponse
func (client *Client) CreateSlrRoleWithContext(ctx context.Context, request *CreateSlrRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateSlrRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSlrRole"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSlrRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes authorized data assets such as databases, instances, or buckets from connection authorization.
//
// Description:
//
// This operation is used to revoke authorization for data assets that a user has authorized, facilitating authorization management.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DeleteDataLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLimitResponse
func (client *Client) DeleteDataLimitWithContext(ctx context.Context, request *DeleteDataLimitRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLimitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLimit"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLimitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom sensitive data detection rule.
//
// @param request - DeleteRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRuleResponse
func (client *Client) DeleteRuleWithContext(ctx context.Context, request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRule"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of audit alert logs.
//
// Description:
//
// This operation is used to query the list of data audit alert logs, which facilitates alerting search and alerting handling.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Invoke this operation at an appropriate frequency.
//
// @param request - DescribeAuditLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditLogsResponse
func (client *Client) DescribeAuditLogsWithContext(ctx context.Context, request *DescribeAuditLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsyncRequestId) {
		query["AsyncRequestId"] = request.AsyncRequestId
	}

	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientUa) {
		query["ClientUa"] = request.ClientUa
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.EffectRowRange) {
		query["EffectRowRange"] = request.EffectRowRange
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecuteTimeRange) {
		query["ExecuteTimeRange"] = request.ExecuteTimeRange
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IpType) {
		query["IpType"] = request.IpType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LoadWhiteList) {
		query["LoadWhiteList"] = request.LoadWhiteList
	}

	if !dara.IsNil(request.LogQueryOpJson) {
		query["LogQueryOpJson"] = request.LogQueryOpJson
	}

	if !dara.IsNil(request.LogSource) {
		query["LogSource"] = request.LogSource
	}

	if !dara.IsNil(request.MemberAccount) {
		query["MemberAccount"] = request.MemberAccount
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.OssObjectKey) {
		query["OssObjectKey"] = request.OssObjectKey
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RuleAggQuery) {
		query["RuleAggQuery"] = request.RuleAggQuery
	}

	if !dara.IsNil(request.RuleCategory) {
		query["RuleCategory"] = request.RuleCategory
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SqlText) {
		query["SqlText"] = request.SqlText
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuditLogs"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of industry templates.
//
// @param request - DescribeCategoryTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCategoryTemplateListResponse
func (client *Client) DescribeCategoryTemplateListWithContext(ctx context.Context, request *DescribeCategoryTemplateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCategoryTemplateListResponse, _err error) {
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

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UsageScenario) {
		query["UsageScenario"] = request.UsageScenario
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCategoryTemplateList"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCategoryTemplateListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a paginated list of rules in a data classification template.
//
// Description:
//
// Retrieves the rules in a data classification template to help you review the rule details.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user on this operation is 10 calls per second. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeCategoryTemplateRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCategoryTemplateRuleListResponse
func (client *Client) DescribeCategoryTemplateRuleListWithContext(ctx context.Context, request *DescribeCategoryTemplateRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCategoryTemplateRuleListResponse, _err error) {
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

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCategoryTemplateRuleList"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCategoryTemplateRuleListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries column data in data asset tables, such as MaxCompute and ApsaraDB RDS tables, that are authorized for connection by Data Security Center.
//
// Description:
//
// This operation is used to view column data in sensitive data asset tables, which helps you accurately analyze sensitive data.
//
// ## Notes
//
// The DescribeColumns operation has been revised to DescribeColumnsV2. Use the newer version DescribeColumnsV2 when developing applications.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnsResponse
func (client *Client) DescribeColumnsWithContext(ctx context.Context, request *DescribeColumnsRequest, runtime *dara.RuntimeOptions) (_result *DescribeColumnsResponse, _err error) {
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

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ModelTagId) {
		query["ModelTagId"] = request.ModelTagId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SensLevelName) {
		query["SensLevelName"] = request.SensLevelName
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.TableId) {
		query["TableId"] = request.TableId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateRuleId) {
		query["TemplateRuleId"] = request.TemplateRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeColumns"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries column data in data asset tables, such as MaxCompute and ApsaraDB RDS tables, that are connected to and authorized by Data Security Center.
//
// @param request - DescribeColumnsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnsV2Response
func (client *Client) DescribeColumnsV2WithContext(ctx context.Context, request *DescribeColumnsV2Request, runtime *dara.RuntimeOptions) (_result *DescribeColumnsV2Response, _err error) {
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

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SensLevelName) {
		query["SensLevelName"] = request.SensLevelName
	}

	if !dara.IsNil(request.TableId) {
		query["TableId"] = request.TableId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeColumnsV2"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColumnsV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries common configuration items for anomaly alerts.
//
// Description:
//
// # Usage notes
//
// Queries common configuration items for anomaly alerts, which you can use to create or restore alert configurations.
//
// # QPS limit
//
// The maximum number of queries per second (QPS) per user is 10. If this limit is exceeded, API calls are throttled. This may impact your business. Call this operation only as needed.
//
// @param request - DescribeConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfigsResponse
func (client *Client) DescribeConfigsWithContext(ctx context.Context, request *DescribeConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConfigs"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConfigsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for data assets on the Overview page of Data Security Center (DSC).
//
// Description:
//
// This operation is typically used to query data assets of different types on the overview page of DSC.
//
// ## Usage notes
//
// This operation is deprecated and no longer maintained.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeDataAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataAssetsResponse
func (client *Client) DescribeDataAssetsWithContext(ctx context.Context, request *DescribeDataAssetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataAssetsResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RangeId) {
		query["RangeId"] = request.RangeId
	}

	if !dara.IsNil(request.RiskLevels) {
		query["RiskLevels"] = request.RiskLevels
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataAssets"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataAssetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an authorized data asset, such as a MaxCompute project, ApsaraDB RDS database, or OSS bucket.
//
// @param request - DescribeDataLimitDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataLimitDetailResponse
func (client *Client) DescribeDataLimitDetailWithContext(ctx context.Context, request *DescribeDataLimitDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataLimitDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataLimitDetail"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataLimitDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeDataLimitSet operation to query the authorization list for unstructured assets or the list of regions supported by Data Security Center.
//
// Description:
//
// Use this operation to retrieve a list of authorized product assets. This list helps you search for and aggregate resources.
//
// ## Notes
//
// In the future, this operation will be used only to retrieve the list of regions that a product supports. Other features will no longer be maintained.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. API calls that exceed this limit are throttled. Throttling can affect your business. We recommend that you call this operation a reasonable number of times.
//
// @param request - DescribeDataLimitSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataLimitSetResponse
func (client *Client) DescribeDataLimitSetWithContext(ctx context.Context, request *DescribeDataLimitSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataLimitSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.RegionType) {
		query["RegionType"] = request.RegionType
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataLimitSet"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataLimitSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data assets, including authorized instances, databases, and buckets.
//
// @param request - DescribeDataLimitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataLimitsResponse
func (client *Client) DescribeDataLimitsWithContext(ctx context.Context, request *DescribeDataLimitsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataLimitsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.CheckStatus) {
		query["CheckStatus"] = request.CheckStatus
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DatamaskStatus) {
		query["DatamaskStatus"] = request.DatamaskStatus
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberAccount) {
		query["MemberAccount"] = request.MemberAccount
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataLimits"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataLimitsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution history of data masking tasks.
//
// Description:
//
// This operation is used to retrieve the execution status of static data masking tasks, allowing you to search for task statuses and view task progress.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeDataMaskingRunHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataMaskingRunHistoryResponse
func (client *Client) DescribeDataMaskingRunHistoryWithContext(ctx context.Context, request *DescribeDataMaskingRunHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataMaskingRunHistoryResponse, _err error) {
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

	if !dara.IsNil(request.DstType) {
		query["DstType"] = request.DstType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MainProcessId) {
		query["MainProcessId"] = request.MainProcessId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SrcTableName) {
		query["SrcTableName"] = request.SrcTableName
	}

	if !dara.IsNil(request.SrcType) {
		query["SrcType"] = request.SrcType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataMaskingRunHistory"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataMaskingRunHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of static data masking tasks.
//
// Description:
//
// This operation is used to retrieve the list of static data masking tasks for search and task management purposes.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - DescribeDataMaskingTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataMaskingTasksResponse
func (client *Client) DescribeDataMaskingTasksWithContext(ctx context.Context, request *DescribeDataMaskingTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataMaskingTasksResponse, _err error) {
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

	if !dara.IsNil(request.DstType) {
		query["DstType"] = request.DstType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataMaskingTasks"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataMaskingTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the column detection results of a data table.
//
// Description:
//
// ## Before you begin
//
// The DescribeDataObjectColumnDetail operation has been revised to DescribeDataObjectColumnDetailV2. Use the newer version DescribeDataObjectColumnDetailV2 when developing applications.
//
// @param request - DescribeDataObjectColumnDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataObjectColumnDetailResponse
func (client *Client) DescribeDataObjectColumnDetailWithContext(ctx context.Context, request *DescribeDataObjectColumnDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataObjectColumnDetailResponse, _err error) {
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

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataObjectColumnDetail"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataObjectColumnDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the column detection results of a data table.
//
// @param request - DescribeDataObjectColumnDetailV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataObjectColumnDetailV2Response
func (client *Client) DescribeDataObjectColumnDetailV2WithContext(ctx context.Context, request *DescribeDataObjectColumnDetailV2Request, runtime *dara.RuntimeOptions) (_result *DescribeDataObjectColumnDetailV2Response, _err error) {
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

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataObjectColumnDetailV2"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataObjectColumnDetailV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data detection results of data tables and files.
//
// Description:
//
// Queries the detection results of data tables and files. This allows you to query data detection results of assets from a global perspective.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - DescribeDataObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataObjectsResponse
func (client *Client) DescribeDataObjectsWithContext(ctx context.Context, request *DescribeDataObjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.APIVersion) {
		query["APIVersion"] = request.APIVersion
	}

	if !dara.IsNil(request.Bucket) {
		query["Bucket"] = request.Bucket
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Cursor) {
		query["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.CursorDirection) {
		query["CursorDirection"] = request.CursorDirection
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.FileCategoryCode) {
		query["FileCategoryCode"] = request.FileCategoryCode
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogStore) {
		query["LogStore"] = request.LogStore
	}

	if !dara.IsNil(request.LogStoreFlag) {
		query["LogStoreFlag"] = request.LogStoreFlag
	}

	if !dara.IsNil(request.MemberAccount) {
		query["MemberAccount"] = request.MemberAccount
	}

	if !dara.IsNil(request.ModelIds) {
		query["ModelIds"] = request.ModelIds
	}

	if !dara.IsNil(request.ModelTagIds) {
		query["ModelTagIds"] = request.ModelTagIds
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentCategoryIds) {
		query["ParentCategoryIds"] = request.ParentCategoryIds
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductIds) {
		query["ProductIds"] = request.ProductIds
	}

	if !dara.IsNil(request.Project) {
		query["Project"] = request.Project
	}

	if !dara.IsNil(request.QueryName) {
		query["QueryName"] = request.QueryName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RiskLevelIdList) {
		query["RiskLevelIdList"] = request.RiskLevelIdList
	}

	if !dara.IsNil(request.RiskLevels) {
		query["RiskLevels"] = request.RiskLevels
	}

	if !dara.IsNil(request.RuleIds) {
		query["RuleIds"] = request.RuleIds
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataObjects"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataObjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of OSS file types that can be detected.
//
// @param request - DescribeDocTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocTypesResponse
func (client *Client) DescribeDocTypesWithContext(ctx context.Context, request *DescribeDocTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDocTypes"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDocTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a single anomalous activity, including the time when the anomalous activity occurred, the anomaly description, and the handling status.
//
// @param request - DescribeEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventDetailResponse
func (client *Client) DescribeEventDetailWithContext(ctx context.Context, request *DescribeEventDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventDetailResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventDetail"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries anomalous activity types.
//
// @param request - DescribeEventTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventTypesResponse
func (client *Client) DescribeEventTypesWithContext(ctx context.Context, request *DescribeEventTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ParentTypeId) {
		query["ParentTypeId"] = request.ParentTypeId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventTypes"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists anomalous events.
//
// Description:
//
// This operation queries alerts for data breach risks to help you find and handle them.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 for each user. If you exceed the limit, your API calls are throttled. This may affect your business. Plan your API calls accordingly.
//
// @param request - DescribeEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
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

	if !dara.IsNil(request.DealUserId) {
		query["DealUserId"] = request.DealUserId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SubTypeCode) {
		query["SubTypeCode"] = request.SubTypeCode
	}

	if !dara.IsNil(request.TargetProductCode) {
		query["TargetProductCode"] = request.TargetProductCode
	}

	if !dara.IsNil(request.TypeCode) {
		query["TypeCode"] = request.TypeCode
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.WarnLevel) {
		query["WarnLevel"] = request.WarnLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEvents"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the completion status of a detection task by task ID. You can obtain the task ID from the ID field in the response of the CreateScanTask or ScanOssObjectV1 operation.
//
// Description:
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation appropriately.
//
// @param request - DescribeIdentifyTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIdentifyTaskStatusResponse
func (client *Client) DescribeIdentifyTaskStatusWithContext(ctx context.Context, request *DescribeIdentifyTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeIdentifyTaskStatusResponse, _err error) {
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
		Action:      dara.String("DescribeIdentifyTaskStatus"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIdentifyTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data assets.
//
// Description:
//
// Queries the list of authorized or unauthorized data assets based on the AuthStatus parameter to help you understand the authorization status of your data assets.
//
// This operation is no longer used in the new console.
//
// ## QPS limit
//
// Each user can call this operation up to 10 times per second. If this limit is exceeded, API calls are throttled, which may affect your business.
//
// @param request - DescribeInstanceSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSourcesResponse
func (client *Client) DescribeInstanceSourcesWithContext(ctx context.Context, request *DescribeInstanceSourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.AuthStatus) {
		query["AuthStatus"] = request.AuthStatus
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SearchType) {
		query["SearchType"] = request.SearchType
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSources"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of data asset instances for authorized MaxCompute, ApsaraDB RDS, and OSS connections.
//
// Description:
//
// When you call the DescribeInstances operation, you can set parameters such as search keywords and risk levels of data asset instances to retrieve a list of data asset instances that meet the specified conditions.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the number of calls exceeds the limit, throttling is triggered, which may affect your business. Call this operation as needed.
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithContext(ctx context.Context, request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
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

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a single OSS storage object that is authorized for connection in Data Security Center.
//
// Description:
//
// This operation is used to query the details of an OSS storage object, which helps you accurately locate sensitive asset information in OSS.
//
// ## Before you begin
//
// The DescribeOssObjectDetail operation has been revised to DescribeOssObjectDetailV2. Use the newer version DescribeOssObjectDetailV2 when developing applications.
//
// ## Rate limit
//
// The single-user queries per second (QPS) limit for this operation is 10. If the number of calls exceeds the limit, throttling is triggered, which may affect your business. Call this operation as appropriate.
//
// @param request - DescribeOssObjectDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssObjectDetailResponse
func (client *Client) DescribeOssObjectDetailWithContext(ctx context.Context, request *DescribeOssObjectDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssObjectDetailResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssObjectDetail"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssObjectDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a single storage object in OSS that is authorized for connection to Data Security Center.
//
// Description:
//
// This operation is used to query the details of an OSS storage object, which helps you accurately locate sensitive asset information in OSS.
//
// @param request - DescribeOssObjectDetailV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssObjectDetailV2Response
func (client *Client) DescribeOssObjectDetailV2WithContext(ctx context.Context, request *DescribeOssObjectDetailV2Request, runtime *dara.RuntimeOptions) (_result *DescribeOssObjectDetailV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssObjectDetailV2"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssObjectDetailV2Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists authorized OSS objects.
//
// @param request - DescribeOssObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssObjectsResponse
func (client *Client) DescribeOssObjectsWithContext(ctx context.Context, request *DescribeOssObjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssObjectsResponse, _err error) {
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

	if !dara.IsNil(request.FileCategoryCode) {
		query["FileCategoryCode"] = request.FileCategoryCode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LastScanTimeEnd) {
		query["LastScanTimeEnd"] = request.LastScanTimeEnd
	}

	if !dara.IsNil(request.LastScanTimeStart) {
		query["LastScanTimeStart"] = request.LastScanTimeStart
	}

	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOssObjects"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssObjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about MaxCompute packages authorized for scanning, including package names, owner accounts, and risk levels.
//
// Description:
//
// This API is typically used to query a list of MaxCompute packages. This helps you search for packages and obtain an overview of sensitive information.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, your API calls are throttled. This can affect your business. Make sure to call this API within the limit.
//
// @param request - DescribePackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackagesResponse
func (client *Client) DescribePackagesWithContext(ctx context.Context, request *DescribePackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribePackagesResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePackages"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of authorized or unauthorized assets.
//
// @param request - DescribeParentInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParentInstanceResponse
func (client *Client) DescribeParentInstanceWithContext(ctx context.Context, request *DescribeParentInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeParentInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthStatus) {
		query["AuthStatus"] = request.AuthStatus
	}

	if !dara.IsNil(request.CheckStatus) {
		query["CheckStatus"] = request.CheckStatus
	}

	if !dara.IsNil(request.ClusterStatus) {
		query["ClusterStatus"] = request.ClusterStatus
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MemberAccount) {
		query["MemberAccount"] = request.MemberAccount
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParentInstance"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParentInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeRiskLevels operation to retrieve a list of risk levels for sensitive data.
//
// Description:
//
// You can use this operation to retrieve a list of risk levels for sensitive data that are defined in the current template. This lets you view the number of rules that reference each risk level and the maximum risk level in the template.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for a single user. If you exceed this limit, your API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeRiskLevelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRiskLevelsResponse
func (client *Client) DescribeRiskLevelsWithContext(ctx context.Context, request *DescribeRiskLevelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskLevelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRiskLevels"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRiskLevelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of sensitive data detection rules.
//
// @param request - DescribeRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRulesResponse
func (client *Client) DescribeRulesWithContext(ctx context.Context, request *DescribeRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ContentCategory) {
		query["ContentCategory"] = request.ContentCategory
	}

	if !dara.IsNil(request.CooperationChannel) {
		query["CooperationChannel"] = request.CooperationChannel
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.CustomType) {
		query["CustomType"] = request.CustomType
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.KeywordCompatible) {
		query["KeywordCompatible"] = request.KeywordCompatible
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.Simplify) {
		query["Simplify"] = request.Simplify
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SupportForm) {
		query["SupportForm"] = request.SupportForm
	}

	if !dara.IsNil(request.WarnLevel) {
		query["WarnLevel"] = request.WarnLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRules"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tables in data assets, such as MaxCompute and RDS, that Data Security Center is authorized to access.
//
// Description:
//
// You can call the DescribeTables operation to retrieve information about specific data asset tables. You can specify parameters such as search keywords and risk levels.
//
// ## QPS limits
//
// Each Alibaba Cloud account can call this operation up to 10 times per second. If you exceed this limit, throttling is triggered, which may affect your business. We recommend that you call this operation at a sustainable rate.
//
// @param request - DescribeTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTablesResponse
func (client *Client) DescribeTablesWithContext(ctx context.Context, request *DescribeTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PackageId) {
		query["PackageId"] = request.PackageId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTables"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all rules in an industry-specific template.
//
// @param request - DescribeTemplateAllRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateAllRulesResponse
func (client *Client) DescribeTemplateAllRulesWithContext(ctx context.Context, request *DescribeTemplateAllRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplateAllRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplateAllRules"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplateAllRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a user account by calling DescribeUserStatus.
//
// Description:
//
// Queries information about the current logon account. This helps you gain a comprehensive understanding of the effectiveness of Data Security Center (DSC).
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - DescribeUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserStatusResponse
func (client *Client) DescribeUserStatusWithContext(ctx context.Context, request *DescribeUserStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserStatus"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DisableUserConfig operation to disable a user configuration. After the user configuration is disabled, you can call the CreateConfig operation with the same request parameter Code to restore the anomaly alert general configuration for the configuration item.
//
// Description:
//
// This operation is used to disable a user configuration based on the code of a configuration item in the anomaly alert general configuration module, allowing you to promptly modify the effective status of the user configuration.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation as needed.
//
// @param request - DisableUserConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableUserConfigResponse
func (client *Client) DisableUserConfigWithContext(ctx context.Context, request *DisableUserConfigRequest, runtime *dara.RuntimeOptions) (_result *DisableUserConfigResponse, _err error) {
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

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableUserConfig"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableUserConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ExecDatamask operation to dynamically mask data.
//
// @param request - ExecDatamaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecDatamaskResponse
func (client *Client) ExecDatamaskWithContext(ctx context.Context, request *ExecDatamaskRequest, runtime *dara.RuntimeOptions) (_result *ExecDatamaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecDatamask"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecDatamaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a data masking task.
//
// @param request - ManualTriggerMaskingProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManualTriggerMaskingProcessResponse
func (client *Client) ManualTriggerMaskingProcessWithContext(ctx context.Context, request *ManualTriggerMaskingProcessRequest, runtime *dara.RuntimeOptions) (_result *ManualTriggerMaskingProcessResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManualTriggerMaskingProcess"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManualTriggerMaskingProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use the MaskOssImage operation to mask images stored as objects.
//
// Description:
//
// *Prerequisites**
//
// To use this operation, you must have an image masking quota. Each call deducts one unit from your quota.
//
// **QPS limit**
//
// The QPS limit for a single user is 10. If you exceed this limit, API calls are throttled, which can affect your business. To prevent service disruptions, operate within this limit.
//
// **Usage notes**
//
// After masking is complete, the system stores the masked image in the aliyun_dsc_desensitization folder within the source bucket.
//
// For example, an image at exampledir/test.png in a bucket is saved as aliyun_dsc_desensitization/exampledir/test.png after masking.
//
// For more information, see https\\://help.aliyun.com/zh/dsc/data-security-center/user-guide/picture-desensitization
//
// @param request - MaskOssImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MaskOssImageResponse
func (client *Client) MaskOssImageWithContext(ctx context.Context, request *MaskOssImageRequest, runtime *dara.RuntimeOptions) (_result *MaskOssImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.IsAlwaysUpload) {
		query["IsAlwaysUpload"] = request.IsAlwaysUpload
	}

	if !dara.IsNil(request.IsCoverObject) {
		query["IsCoverObject"] = request.IsCoverObject
	}

	if !dara.IsNil(request.IsSupportRestore) {
		query["IsSupportRestore"] = request.IsSupportRestore
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaskRuleIdList) {
		query["MaskRuleIdList"] = request.MaskRuleIdList
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MaskOssImage"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MaskOssImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration items of a Data Security Center (DSC) connection authorization.
//
// @param request - ModifyDataLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataLimitResponse
func (client *Client) ModifyDataLimitWithContext(ctx context.Context, request *ModifyDataLimitRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataLimitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.AutoScan) {
		query["AutoScan"] = request.AutoScan
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogStoreDay) {
		query["LogStoreDay"] = request.LogStoreDay
	}

	if !dara.IsNil(request.ModifyPassword) {
		query["ModifyPassword"] = request.ModifyPassword
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SamplingSize) {
		query["SamplingSize"] = request.SamplingSize
	}

	if !dara.IsNil(request.SecurityGroupIdList) {
		query["SecurityGroupIdList"] = request.SecurityGroupIdList
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.VSwitchIdList) {
		query["VSwitchIdList"] = request.VSwitchIdList
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataLimit"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDataLimitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the risk level definition rules for sensitive data, including the default risk level for unrecognized data and the risk level for data classified as "sensitive".
//
// Description:
//
// This operation is used to modify the risk level definition rules for sensitive data, which helps you plan risk levels.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If this limit is exceeded, API calls are throttled, which may affect your business. Call this operation at a reasonable frequency.
//
// @param request - ModifyDefaultLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefaultLevelResponse
func (client *Client) ModifyDefaultLevelWithContext(ctx context.Context, request *ModifyDefaultLevelRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefaultLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultId) {
		query["DefaultId"] = request.DefaultId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SensitiveIds) {
		query["SensitiveIds"] = request.SensitiveIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefaultLevel"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefaultLevelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Handles anomalous activities.
//
// Description:
//
// This API operation is typically used to handle alerts for data breach threats, helping you protect your data assets promptly.
//
// ## QPS limits
//
// This API operation has a queries per second (QPS) limit of 10 for each user. If you exceed the limit, API calls are throttled, which can affect your business. We recommend calling the API operation at a reasonable rate.
//
// @param request - ModifyEventStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEventStatusResponse
func (client *Client) ModifyEventStatusWithContext(ctx context.Context, request *ModifyEventStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyEventStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Backed) {
		query["Backed"] = request.Backed
	}

	if !dara.IsNil(request.DealReason) {
		query["DealReason"] = request.DealReason
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEventStatus"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEventStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes this operation to enable the feature for detecting anomalous activity events of specified child classes.
//
// @param request - ModifyEventTypeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEventTypeStatusResponse
func (client *Client) ModifyEventTypeStatusWithContext(ctx context.Context, request *ModifyEventTypeStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyEventTypeStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SubTypeIds) {
		query["SubTypeIds"] = request.SubTypeIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEventTypeStatus"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEventTypeStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyReportTaskStatus operation to enable or disable report tasks.
//
// Description:
//
// After you activate Data Security Center (DSC), report tasks are enabled by default. If you disable report tasks, Report Center, Cloud-native Data Audit Overview, and Data Security Lab will not generate new statistical data. Existing data is not affected.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - ModifyReportTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyReportTaskStatusResponse
func (client *Client) ModifyReportTaskStatusWithContext(ctx context.Context, request *ModifyReportTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyReportTaskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ReportTaskStatus) {
		query["ReportTaskStatus"] = request.ReportTaskStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyReportTaskStatus"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyReportTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a custom sensitive data detection rule in Data Security Center (DSC).
//
// Description:
//
// When calling this operation, you must specify the rule name, rule ID, and rule content parameters.
//
// ## QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation at an appropriate frequency.
//
// @param request - ModifyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRuleResponse
func (client *Client) ModifyRuleWithContext(ctx context.Context, request *ModifyRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MatchType) {
		query["MatchType"] = request.MatchType
	}

	if !dara.IsNil(request.ModelRuleIds) {
		query["ModelRuleIds"] = request.ModelRuleIds
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.SupportForm) {
		query["SupportForm"] = request.SupportForm
	}

	if !dara.IsNil(request.TemplateRuleIds) {
		query["TemplateRuleIds"] = request.TemplateRuleIds
	}

	if !dara.IsNil(request.WarnLevel) {
		query["WarnLevel"] = request.WarnLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRule"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the detection feature of a sensitive data detection rule.
//
// @param request - ModifyRuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRuleStatusResponse
func (client *Client) ModifyRuleStatusWithContext(ctx context.Context, request *ModifyRuleStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyRuleStatusResponse, _err error) {
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

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRuleStatus"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRuleStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores an image that has been masked by calling the RestoreOssImage operation.
//
// Description:
//
// For files that were masked by calling MaskOssImage with IsAlwaysUpload set to true, you can call RestoreOssImage to retrieve the original image.
//
// For example, the image aliyun_dsc_desensitization/exampledir/test.png in the bucket is stored as aliyun_dsc_original/exampledir/test.png after restoration.
//
// @param request - RestoreOssImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreOssImageResponse
func (client *Client) RestoreOssImageWithContext(ctx context.Context, request *RestoreOssImageRequest, runtime *dara.RuntimeOptions) (_result *RestoreOssImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bucket) {
		query["Bucket"] = request.Bucket
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ObjectKey) {
		query["ObjectKey"] = request.ObjectKey
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.TargetObjectKey) {
		query["TargetObjectKey"] = request.TargetObjectKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreOssImage"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreOssImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ScanOssObjectV1 operation creates a scan task to detect sensitive data in a specified object.
//
// Description:
//
// ### Prerequisites
//
// You must authorize and connect to the specified bucket before you call this operation. If the bucket is not authorized, the API call returns the bucket_not_authorized error code.
//
// ### QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed the limit, API calls are throttled. This may affect your business. We recommend that you call the operation at a reasonable rate.
//
// ### Usage notes
//
// You can use the returned task ID to call the DescribeIdentifyTaskStatus operation to check the running status of the task.
//
// After the task is complete, call the DescribeOssObjectDetailV2 operation and provide the BucketName, ServiceRegionId, and ObjectKey to view the sensitive data detection results for the object.
//
// @param tmpReq - ScanOssObjectV1Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScanOssObjectV1Response
func (client *Client) ScanOssObjectV1WithContext(ctx context.Context, tmpReq *ScanOssObjectV1Request, runtime *dara.RuntimeOptions) (_result *ScanOssObjectV1Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ScanOssObjectV1ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectKeyList) {
		request.ObjectKeyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectKeyList, dara.String("ObjectKeyList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BucketName) {
		query["BucketName"] = request.BucketName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ObjectKeyListShrink) {
		query["ObjectKeyList"] = request.ObjectKeyListShrink
	}

	if !dara.IsNil(request.ServiceRegionId) {
		query["ServiceRegionId"] = request.ServiceRegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScanOssObjectV1"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ScanOssObjectV1Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the StopMaskingProcess operation to stop a data masking task. You can call the ManualTriggerMaskingProcess operation to restart a stopped task using its unique resource ID.
//
// Description:
//
// This operation stops a running data masking task. For example, you can call this operation if you no longer need to mask data for a previously configured task.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, API calls are throttled. This may affect your business. Ensure that you call this operation within the specified limit.
//
// @param request - StopMaskingProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMaskingProcessResponse
func (client *Client) StopMaskingProcessWithContext(ctx context.Context, request *StopMaskingProcessRequest, runtime *dara.RuntimeOptions) (_result *StopMaskingProcessResponse, _err error) {
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

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopMaskingProcess"),
		Version:     dara.String("2019-01-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopMaskingProcessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
