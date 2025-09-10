// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a quota alert.
//
// Description:
//
// The quota alerting feature has been upgraded and this API operation will be deprecated. If you want to create a quota alert of the new version, call CloudMonitor API operations. For more information, see [Use API operations to manage new quota alert rules](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - CreateQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaAlarmResponse
func (client *Client) CreateQuotaAlarmWithContext(ctx context.Context, request *CreateQuotaAlarmRequest, runtime *dara.RuntimeOptions) (_result *CreateQuotaAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmName) {
		body["AlarmName"] = request.AlarmName
	}

	if !dara.IsNil(request.OriginalContext) {
		body["OriginalContext"] = request.OriginalContext
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaDimensions) {
		body["QuotaDimensions"] = request.QuotaDimensions
	}

	if !dara.IsNil(request.Threshold) {
		body["Threshold"] = request.Threshold
	}

	if !dara.IsNil(request.ThresholdPercent) {
		body["ThresholdPercent"] = request.ThresholdPercent
	}

	if !dara.IsNil(request.ThresholdType) {
		body["ThresholdType"] = request.ThresholdType
	}

	if !dara.IsNil(request.WebHook) {
		body["WebHook"] = request.WebHook
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQuotaAlarm"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQuotaAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an application to increase a quota.
//
// Description:
//
// In this example, the operation is called to submit an application to increase the value of a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. The quota belongs to Elastic Compute Service (ECS). The expected value of the quota is `804`, the application reason is `Scale Out`, and the ID of the region to which the quota belongs is `cn-hangzhou`.
//
// @param request - CreateQuotaApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaApplicationResponse
func (client *Client) CreateQuotaApplicationWithContext(ctx context.Context, request *CreateQuotaApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateQuotaApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuditMode) {
		body["AuditMode"] = request.AuditMode
	}

	if !dara.IsNil(request.DesireValue) {
		body["DesireValue"] = request.DesireValue
	}

	if !dara.IsNil(request.Dimensions) {
		body["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EffectiveTime) {
		body["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EnvLanguage) {
		body["EnvLanguage"] = request.EnvLanguage
	}

	if !dara.IsNil(request.ExpireTime) {
		body["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.NoticeType) {
		body["NoticeType"] = request.NoticeType
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !dara.IsNil(request.Reason) {
		body["Reason"] = request.Reason
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQuotaApplication"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQuotaApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a quota increase application. After you add a quota item to a quota template, the system automatically submits quota applications only for new members of the resource directory. The quota values for existing members remain unchanged. If you want to increase the quota values of existing members, you can submit a quota application for the members by applying quota templates to the members. Only the management account of a resource directory can create multiple quota applications at a time.
//
// Description:
//
// ### [](#)QPS limit
//
// You can add a maximum of 10 quota items to a quota template at a time.
//
// @param request - CreateQuotaApplicationsForTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaApplicationsForTemplateResponse
func (client *Client) CreateQuotaApplicationsForTemplateWithContext(ctx context.Context, request *CreateQuotaApplicationsForTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateQuotaApplicationsForTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunUids) {
		body["AliyunUids"] = request.AliyunUids
	}

	if !dara.IsNil(request.DesireValue) {
		body["DesireValue"] = request.DesireValue
	}

	if !dara.IsNil(request.Dimensions) {
		body["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EffectiveTime) {
		body["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EnvLanguage) {
		body["EnvLanguage"] = request.EnvLanguage
	}

	if !dara.IsNil(request.ExpireTime) {
		body["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.NoticeType) {
		body["NoticeType"] = request.NoticeType
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !dara.IsNil(request.Reason) {
		body["Reason"] = request.Reason
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQuotaApplicationsForTemplate"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQuotaApplicationsForTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a quota template by using the management account of a resource directory. After you create a quota template, if a member is added to the resource directory, the quota template automatically submits a quota increase request for the member. The quota values for existing members remain unchanged. You can use a quota template to apply for increases on multiple quotas at the same time. This automated approach improves the efficiency of quota management across your organization. Only the management account of a resource directory can create quota templates.
//
// Description:
//
// ### [](#)Usage notes
//
// You must set the `ServiceStatus` parameter to `1`. This ensures that the quota template is enabled.
//
// You can call the [GetQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450407.html) operation to query the status of a quota template. If the value of the `ServiceStatus` parameter in the response is `0` or `-1`, you must call the [ModifyQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450406.html) operation to modify the value to `1`. A value of 0 indicates that the quota template is not configured. A value of -1 indicates that the quota template is disabled. A value of 1 indicates that the quota template is enabled.
//
// ### [](#)
//
// After you create a quota template, you can call the [ListQuotaApplicationsForTemplate](https://help.aliyun.com/document_detail/2584864.html) operation to view the approval result. If the value of the `Status` parameter in the response is `Agree`, the quota template is approved.
//
// @param request - CreateTemplateQuotaItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateQuotaItemResponse
func (client *Client) CreateTemplateQuotaItemWithContext(ctx context.Context, request *CreateTemplateQuotaItemRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateQuotaItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DesireValue) {
		body["DesireValue"] = request.DesireValue
	}

	if !dara.IsNil(request.Dimensions) {
		body["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EffectiveTime) {
		body["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EnvLanguage) {
		body["EnvLanguage"] = request.EnvLanguage
	}

	if !dara.IsNil(request.ExpireTime) {
		body["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.NoticeType) {
		body["NoticeType"] = request.NoticeType
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTemplateQuotaItem"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTemplateQuotaItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quota alert.
//
// Description:
//
//	  The quota alerting feature has been upgraded and this API operation will be deprecated. You can call this operation only to delete a quota alert rule of the old version. If you want to delete a quota alert rule of the new version, call the CloudMonitor API operation [DeleteMetricRules](https://help.aliyun.com/document_detail/2513295.html) or [DeleteMetricRuleTargets](https://help.aliyun.com/document_detail/2513294.html). For more information about how to call API operations to manage quota alert rules of the new version, see [Manage quota alerts of the new version by calling API operations](https://help.aliyun.com/document_detail/2863234.html).
//
//		- In this example, the API operation is called to delete a quota alert rule whose ID is `6b512ab7-da3a-4142-b529-2b2a9294****`.
//
// @param request - DeleteQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQuotaAlarmResponse
func (client *Client) DeleteQuotaAlarmWithContext(ctx context.Context, request *DeleteQuotaAlarmRequest, runtime *dara.RuntimeOptions) (_result *DeleteQuotaAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmId) {
		body["AlarmId"] = request.AlarmId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQuotaAlarm"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQuotaAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quota template by using the management account of a resource directory. After you delete a quota template, if a member is added to the resource directory, the quota template no longer automatically submits a quota increase request for the member. Only the management account of a resource directory can delete quota templates.
//
// @param request - DeleteTemplateQuotaItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateQuotaItemResponse
func (client *Client) DeleteTemplateQuotaItemWithContext(ctx context.Context, request *DeleteTemplateQuotaItemRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateQuotaItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplateQuotaItem"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateQuotaItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a quota of a cloud service.
//
// Description:
//
// In this example, the operation is called to query the details of a quota whose ID is `q_security-groups` and whose name is Maximum Number of Security Groups. This quota belongs to Elastic Compute Service (ECS). The query result shows the details of the quota. The details include the name, ID, description, quota value, used quota, unit, and dimension of the quota. In this example, the quota name is `Maximum Number of Security Groups`. The quota ID is `q_security-groups`. The description is `The maximum number of security groups that can be created for the current account`. The quota value is `801`. The used quota is `26`. The quota unit is `Number of security groups`. The quota dimension is `{"regionId":"cn-hangzhou"}`.
//
// @param request - GetProductQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductQuotaResponse
func (client *Client) GetProductQuotaWithContext(ctx context.Context, request *GetProductQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetProductQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		body["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProductQuota"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProductQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a quota dimension that is supported by an Alibaba Cloud service.
//
// Description:
//
// In this example, the operation is called to query the details of a quota dimension whose key is `regionId`. The quota dimension belongs to Elastic Compute Service (ECS) Quotas by Instance Type whose service code is ecs-spec. The following query results are returned:
//
//   - The values of the quota dimension include `cn-shenzhen`, `cn-beijing`, and `cn-hangzhou`.
//
//   - The name of the quota dimension is `region`.
//
// @param request - GetProductQuotaDimensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductQuotaDimensionResponse
func (client *Client) GetProductQuotaDimensionWithContext(ctx context.Context, request *GetProductQuotaDimensionRequest, runtime *dara.RuntimeOptions) (_result *GetProductQuotaDimensionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DependentDimensions) {
		body["DependentDimensions"] = request.DependentDimensions
	}

	if !dara.IsNil(request.DimensionKey) {
		body["DimensionKey"] = request.DimensionKey
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProductQuotaDimension"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProductQuotaDimensionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// In this example, the operation is called to query the details of a quota alert. The details of the alert are returned. The query results include the alert ID, alert name, alert contact, and time when the quota alert was created.
//
// Description:
//
//	  The quota alerting feature has been upgraded and this API operation will be deprecated. You can call this operation only to query the details about the quota alert rules of the old version. If you want to query the details about the quota alert rules of the new version, call CloudMonitor API operations. For more information, see [Use API operations to manage new quota alert rules](https://help.aliyun.com/document_detail/2863234.html).
//
//		- In this example, the operation is called to query the details of a quota alert rule whose ID is `78d7e436-4b25-4897-84b5-d7b656bb****`. The details of the alert rule are returned. The query result includes the alert ID, alert name, alert contact, and the time when the quota alert rule was created.
//
// @param request - GetQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaAlarmResponse
func (client *Client) GetQuotaAlarmWithContext(ctx context.Context, request *GetQuotaAlarmRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmId) {
		body["AlarmId"] = request.AlarmId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuotaAlarm"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a specified application that is submitted to increase a quota.
//
// Description:
//
// In this example, the operation is called to query the details about an application whose ID is `d314d6ae-867d-484c-9009-3d421a80****`. The query result shows the details about the application. The details include the application ID, application time, expected quota value, and application result.
//
// @param request - GetQuotaApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaApplicationResponse
func (client *Client) GetQuotaApplicationWithContext(ctx context.Context, request *GetQuotaApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuotaApplication"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaApplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about quota application approval, such as the average amount of time required for approval, whether approval reminders are supported, and the interval between two consecutive approval reminders.
//
// Description:
//
// ### [](#)Prerequisites
//
// Make sure that you have created an application for quota increase. For more information, see [CreateQuotaApplication](https://help.aliyun.com/document_detail/440566.html).
//
// @param request - GetQuotaApplicationApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaApplicationApprovalResponse
func (client *Client) GetQuotaApplicationApprovalWithContext(ctx context.Context, request *GetQuotaApplicationApprovalRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaApplicationApprovalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuotaApplicationApproval"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaApplicationApprovalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a quota template.
//
// Description:
//
// By default, the value of `ServiceStatus` is `0`, which indicates that no quota template is specified. If you want to use a quota template, make sure that the quota template is enabled. In this case, the value of `ServiceStatus` is `1`.
//
// @param request - GetQuotaTemplateServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaTemplateServiceStatusResponse
func (client *Client) GetQuotaTemplateServiceStatusWithContext(ctx context.Context, request *GetQuotaTemplateServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaTemplateServiceStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceDirectoryId) {
		body["ResourceDirectoryId"] = request.ResourceDirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuotaTemplateServiceStatus"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaTemplateServiceStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert records.
//
// Description:
//
// The quota alerting feature has been upgraded and this API operation will be deprecated. You can call this operation only to query the historical records of quota alert rules of the old version. If you want to query the historical records of quota alert rules of the new version, call the CloudMonitor API operation [DescribeAlertLogCount](https://help.aliyun.com/document_detail/2513275.html) or [DescribeAlertLogList](https://help.aliyun.com/document_detail/2513276.html). For more information about how to call API operations to manage quota alert rules of the new version, see [Manage quota alerts of the new version by calling API operations](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - ListAlarmHistoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlarmHistoriesResponse
func (client *Client) ListAlarmHistoriesWithContext(ctx context.Context, request *ListAlarmHistoriesRequest, runtime *dara.RuntimeOptions) (_result *ListAlarmHistoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmId) {
		body["AlarmId"] = request.AlarmId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlarmHistories"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlarmHistoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quotas on which a specified quota depends.
//
// Description:
//
// In this example, the operation is called to query the quotas on which a Container Service for Kubernetes (ACK) quota whose ID is `q_i5uzm3` depends. This quota is the maximum number of nodes that can be created in an ACK cluster. The query result indicates that the specified quota depends on the following three quotas:
//
//   - An Elastic Compute Service (ECS) quota whose ID is `q_elastic-network-interfaces`. This quota is the maximum number of ENIs (Secondary ENIs) that can be owned by an Alibaba Cloud account. The quota is available in the following regions: `cn-shenzhen`, `cn-beijing`, and `cn-hangzhou`.
//
//   - A Server Load Balancer (SLB) quota whose ID is `q_fh20b0`. This quota is the number of servers that can be attached to the backend of an SLB instance.
//
//   - An SLB quota whose ID is `q_3mmbsp`. This quota is the number of SLB instances that can be owned by an Alibaba Cloud account.
//
// @param request - ListDependentQuotasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDependentQuotasResponse
func (client *Client) ListDependentQuotasWithContext(ctx context.Context, request *ListDependentQuotasRequest, runtime *dara.RuntimeOptions) (_result *ListDependentQuotasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDependentQuotas"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDependentQuotasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dimension groups of an Alibaba Cloud service.
//
// Description:
//
// This topic provides an example on how to call the ListProductDimensionGroups operation to query the dimension groups of Object Storage Service (OSS). In this example, a dimension group is returned. The group name is `OSS_Group`, the group code is `oss_wf1ngqmd7q`, and the group key is `chargeType`.
//
// @param request - ListProductDimensionGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductDimensionGroupsResponse
func (client *Client) ListProductDimensionGroupsWithContext(ctx context.Context, request *ListProductDimensionGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListProductDimensionGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProductDimensionGroups"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductDimensionGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota dimensions that are supported by the specified Alibaba Cloud service.
//
// Description:
//
// In this example, the operation is called to query the quota dimensions that are supported by Elastic Compute Service (ECS). The query results show all the quota dimensions that are supported by ECS.
//
// @param request - ListProductQuotaDimensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductQuotaDimensionsResponse
func (client *Client) ListProductQuotaDimensionsWithContext(ctx context.Context, request *ListProductQuotaDimensionsRequest, runtime *dara.RuntimeOptions) (_result *ListProductQuotaDimensionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProductQuotaDimensions"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductQuotaDimensionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quotas of a specific Alibaba Cloud service.
//
// Description:
//
// In this example, the operation is called to query the quotas whose instance type is `ecs.g5.2xlarge`. The quotas belong to Elastic Compute Service (ECS) Quotas by Instance Type. The query result includes the name, ID, unit, dimensions, and cycle of each quota.
//
// @param request - ListProductQuotasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductQuotasResponse
func (client *Client) ListProductQuotasWithContext(ctx context.Context, request *ListProductQuotasRequest, runtime *dara.RuntimeOptions) (_result *ListProductQuotasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		body["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.GroupCode) {
		body["GroupCode"] = request.GroupCode
	}

	if !dara.IsNil(request.KeyWord) {
		body["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SortOrder) {
		body["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProductQuotas"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductQuotasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud services that support Quota Center.
//
// Description:
//
// The services in the query result are the same as the services listed in [Alibaba Cloud services that support Quota Center](https://help.aliyun.com/document_detail/182368.html).
//
// @param request - ListProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithContext(ctx context.Context, request *ListProductsRequest, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProducts"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quota alerts.
//
// Description:
//
// The quota alerting feature has been upgraded and this API operation will be deprecated. You can call this operation only to query quota alert rules of the old version. If you want to query quota alert rules of the new version, call the CloudMonitor API operation [DescribeMetricRuleList](https://help.aliyun.com/document_detail/2513291.html). For more information about how to call API operations to manage quota alert rules of the new version, see [Manage quota alerts of the new version by calling API operations](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - ListQuotaAlarmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaAlarmsResponse
func (client *Client) ListQuotaAlarmsWithContext(ctx context.Context, request *ListQuotaAlarmsRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaAlarmsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmName) {
		body["AlarmName"] = request.AlarmName
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaDimensions) {
		body["QuotaDimensions"] = request.QuotaDimensions
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotaAlarms"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaAlarmsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quota templates by using the management account of a resource directory.
//
// @param request - ListQuotaApplicationTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationTemplatesResponse
func (client *Client) ListQuotaApplicationTemplatesWithContext(ctx context.Context, request *ListQuotaApplicationTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaApplicationTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		body["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotaApplicationTemplates"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaApplicationTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an application that is submitted to increase a quota.
//
// Description:
//
// In this example, the operation is called to query the details of an application that is submitted to increase a quota whose ID is `q_i5uzm3` and whose name is Maximum Number of Nodes. This quota belongs to Container Service for Kubernetes (ACK). The query result shows the details of the application. The details include the application ID, application time, requested quota, and application result. In this example, the application ID is `b926571d-cc09-4711-b547-58a615f0****`. The application time is `2021-01-15T09:13:53Z`. The expected quota value is `101`. The application result is `Agree`.
//
// @param request - ListQuotaApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationsResponse
func (client *Client) ListQuotaApplicationsWithContext(ctx context.Context, request *ListQuotaApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		body["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.KeyWord) {
		body["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotaApplications"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaApplicationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a quota increase application for member accounts in a resource directory.
//
// @param request - ListQuotaApplicationsDetailForTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationsDetailForTemplateResponse
func (client *Client) ListQuotaApplicationsDetailForTemplateWithContext(ctx context.Context, request *ListQuotaApplicationsDetailForTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaApplicationsDetailForTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunUid) {
		body["AliyunUid"] = request.AliyunUid
	}

	if !dara.IsNil(request.BatchQuotaApplicationId) {
		body["BatchQuotaApplicationId"] = request.BatchQuotaApplicationId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotaApplicationsDetailForTemplate"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaApplicationsDetailForTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the application records of a quota template that is used to apply for quotas for member accounts.
//
// @param request - ListQuotaApplicationsForTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationsForTemplateResponse
func (client *Client) ListQuotaApplicationsForTemplateWithContext(ctx context.Context, request *ListQuotaApplicationsForTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaApplicationsForTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplyEndTime) {
		body["ApplyEndTime"] = request.ApplyEndTime
	}

	if !dara.IsNil(request.ApplyStartTime) {
		body["ApplyStartTime"] = request.ApplyStartTime
	}

	if !dara.IsNil(request.BatchQuotaApplicationId) {
		body["BatchQuotaApplicationId"] = request.BatchQuotaApplicationId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotaApplicationsForTemplate"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaApplicationsForTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of a quota template. By default, the quota template is not configured. If the management account of a resource directory uses a quota template for the first time, you must enable the quota template. Only the management account of a resource directory can change the status of quota templates.
//
// Description:
//
// ### [](#)Prerequisites
//
// A resource directory is enabled. For more information, see [EnableResourceDirectory](https://help.aliyun.com/document_detail/604185.html).
//
// ### [](#)Usage notes
//
// If the `ServiceStatus` parameter is set to `0` or `-1`, you can call this operation to set the parameter to `1`. Then, you can call the [CreateTemplateQuotaItem](https://help.aliyun.com/document_detail/450615.html) operation to create a quota template.
//
// @param request - ModifyQuotaTemplateServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyQuotaTemplateServiceStatusResponse
func (client *Client) ModifyQuotaTemplateServiceStatusWithContext(ctx context.Context, request *ModifyQuotaTemplateServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyQuotaTemplateServiceStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ServiceStatus) {
		body["ServiceStatus"] = request.ServiceStatus
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyQuotaTemplateServiceStatus"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyQuotaTemplateServiceStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the quota template.
//
// @param request - ModifyTemplateQuotaItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateQuotaItemResponse
func (client *Client) ModifyTemplateQuotaItemWithContext(ctx context.Context, request *ModifyTemplateQuotaItemRequest, runtime *dara.RuntimeOptions) (_result *ModifyTemplateQuotaItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QuotaCategory) {
		query["QuotaCategory"] = request.QuotaCategory
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DesireValue) {
		body["DesireValue"] = request.DesireValue
	}

	if !dara.IsNil(request.Dimensions) {
		body["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EffectiveTime) {
		body["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EnvLanguage) {
		body["EnvLanguage"] = request.EnvLanguage
	}

	if !dara.IsNil(request.ExpireTime) {
		body["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.NoticeType) {
		body["NoticeType"] = request.NoticeType
	}

	if !dara.IsNil(request.ProductCode) {
		body["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.QuotaActionCode) {
		body["QuotaActionCode"] = request.QuotaActionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTemplateQuotaItem"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTemplateQuotaItemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reminds the approver of a quota application to review the application. This operation is applicable to quota applications that support the approval reminding feature.
//
// Description:
//
// >  You can call this operation to enable the approval reminder feature for quota applications that support this feature. To check whether this feature is supported, you can view the value of `SupportReminder` in the GetQuotaApplicationApproval operation. If the value of SupportReminder is `true`, this feature is supported.
//
// @param request - RemindQuotaApplicationApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemindQuotaApplicationApprovalResponse
func (client *Client) RemindQuotaApplicationApprovalWithContext(ctx context.Context, request *RemindQuotaApplicationApprovalRequest, runtime *dara.RuntimeOptions) (_result *RemindQuotaApplicationApprovalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemindQuotaApplicationApproval"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemindQuotaApplicationApprovalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a quota alert rule.
//
// Description:
//
//	  The quota alerting feature has been upgraded and this API operation will be deprecated. If you want to modify the information about a specific quota alert rule of the new version, call the CloudMonitor API operation [PutResourceMetricRules](https://help.aliyun.com/document_detail/2513316.html) or [PutMetricRuleTargets](https://help.aliyun.com/document_detail/2513302.html). For more information about how to call API operations to manage quota alert rules of the new version, see [Manage quota alerts of the new version by calling API operations](https://help.aliyun.com/document_detail/2863234.html).
//
//		- In this example, the API operation is called to modify the information about a quota alert rule whose ID is `a2efa7fc-832f-47bb-8054-39e28012****` and whose name is `rules`. The alert threshold is changed from `150` to `160`.
//
// @param request - UpdateQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaAlarmResponse
func (client *Client) UpdateQuotaAlarmWithContext(ctx context.Context, request *UpdateQuotaAlarmRequest, runtime *dara.RuntimeOptions) (_result *UpdateQuotaAlarmResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmId) {
		body["AlarmId"] = request.AlarmId
	}

	if !dara.IsNil(request.AlarmName) {
		body["AlarmName"] = request.AlarmName
	}

	if !dara.IsNil(request.Threshold) {
		body["Threshold"] = request.Threshold
	}

	if !dara.IsNil(request.ThresholdPercent) {
		body["ThresholdPercent"] = request.ThresholdPercent
	}

	if !dara.IsNil(request.ThresholdType) {
		body["ThresholdType"] = request.ThresholdType
	}

	if !dara.IsNil(request.WebHook) {
		body["WebHook"] = request.WebHook
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQuotaAlarm"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQuotaAlarmResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
