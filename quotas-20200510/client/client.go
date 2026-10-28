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
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("quotas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates an alert rule for a quota item.
//
// Description:
//
// The quota alerting feature has been upgraded and this API operation will be deprecated. Use the CloudMonitor API operations [PutResourceMetricRules](https://help.aliyun.com/document_detail/2513316.html) and [PutMetricRuleTargets](https://help.aliyun.com/document_detail/2513302.html) to create alert rules for the new version of quota alerting.
//
// For more information about how to use API operations to manage the new version of quota alerting, see [Manage new quota alerts by using API operations](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - CreateQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaAlarmResponse
func (client *Client) CreateQuotaAlarmWithOptions(request *CreateQuotaAlarmRequest, runtime *dara.RuntimeOptions) (_result *CreateQuotaAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert rule for a quota item.
//
// Description:
//
// The quota alerting feature has been upgraded and this API operation will be deprecated. Use the CloudMonitor API operations [PutResourceMetricRules](https://help.aliyun.com/document_detail/2513316.html) and [PutMetricRuleTargets](https://help.aliyun.com/document_detail/2513302.html) to create alert rules for the new version of quota alerting.
//
// For more information about how to use API operations to manage the new version of quota alerting, see [Manage new quota alerts by using API operations](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - CreateQuotaAlarmRequest
//
// @return CreateQuotaAlarmResponse
func (client *Client) CreateQuotaAlarm(request *CreateQuotaAlarmRequest) (_result *CreateQuotaAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQuotaAlarmResponse{}
	_body, _err := client.CreateQuotaAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a quota increase request.
//
// Description:
//
// > After you create a quota increase request, you can call the [ListQuotaApplications](https://help.aliyun.com/document_detail/440568.html) operation to view the approval result. If the value of `Status` is `Agree`, the request is approved.
//
// @param request - CreateQuotaApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaApplicationResponse
func (client *Client) CreateQuotaApplicationWithOptions(request *CreateQuotaApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateQuotaApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a quota increase request.
//
// Description:
//
// > After you create a quota increase request, you can call the [ListQuotaApplications](https://help.aliyun.com/document_detail/440568.html) operation to view the approval result. If the value of `Status` is `Agree`, the request is approved.
//
// @param request - CreateQuotaApplicationRequest
//
// @return CreateQuotaApplicationResponse
func (client *Client) CreateQuotaApplication(request *CreateQuotaApplicationRequest) (_result *CreateQuotaApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQuotaApplicationResponse{}
	_body, _err := client.CreateQuotaApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates batch quota increase applications for members in a resource folder. When you add a quota item to a quota template, only new members in the resource folder autocommit quota increase applications, and existing members remain unchanged. To increase quotas for existing members by using a quota template, you can select the quota template to commit batch quota increase applications for multiple members in the resource folder. Only the administrator account of the resource folder can create batch quota applications.
//
// Description:
//
// ### Limits
//
// A single application can include a maximum of 10 quota items.
//
// @param request - CreateQuotaApplicationsForTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQuotaApplicationsForTemplateResponse
func (client *Client) CreateQuotaApplicationsForTemplateWithOptions(request *CreateQuotaApplicationsForTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateQuotaApplicationsForTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates batch quota increase applications for members in a resource folder. When you add a quota item to a quota template, only new members in the resource folder autocommit quota increase applications, and existing members remain unchanged. To increase quotas for existing members by using a quota template, you can select the quota template to commit batch quota increase applications for multiple members in the resource folder. Only the administrator account of the resource folder can create batch quota applications.
//
// Description:
//
// ### Limits
//
// A single application can include a maximum of 10 quota items.
//
// @param request - CreateQuotaApplicationsForTemplateRequest
//
// @return CreateQuotaApplicationsForTemplateResponse
func (client *Client) CreateQuotaApplicationsForTemplate(request *CreateQuotaApplicationsForTemplateRequest) (_result *CreateQuotaApplicationsForTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQuotaApplicationsForTemplateResponse{}
	_body, _err := client.CreateQuotaApplicationsForTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a quota template. After the management account creates a quota template, the template automatically submits quota increase applications for new members added to the resource directory. Existing members remain unaffected. You can use a quota template to apply for increases to multiple quota items at a time, improving the efficiency and automation of quota management across the organization. Only the management account of a resource directory can create a quota template.
//
// Description:
//
// ### Before you begin
//
// Make sure that the status of the quota template is enabled (that is, `ServiceStatus` is `1`).
//
// Call the [GetQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450407.html) operation to query the status of the quota template. If the `ServiceStatus` of the quota template is `0` (not configured) or `-1` (disabled), call the [ModifyQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450406.html) operation to change the status to `1` (enabled).
//
// ### Usage notes
//
// After you create a quota template, you can call the [ListQuotaApplicationsForTemplate](https://help.aliyun.com/document_detail/2584864.html) operation to view the approval result. If `Status` is displayed as `Agree`, the application is approved.
//
// @param request - CreateTemplateQuotaItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateQuotaItemResponse
func (client *Client) CreateTemplateQuotaItemWithOptions(request *CreateTemplateQuotaItemRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateQuotaItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a quota template. After the management account creates a quota template, the template automatically submits quota increase applications for new members added to the resource directory. Existing members remain unaffected. You can use a quota template to apply for increases to multiple quota items at a time, improving the efficiency and automation of quota management across the organization. Only the management account of a resource directory can create a quota template.
//
// Description:
//
// ### Before you begin
//
// Make sure that the status of the quota template is enabled (that is, `ServiceStatus` is `1`).
//
// Call the [GetQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450407.html) operation to query the status of the quota template. If the `ServiceStatus` of the quota template is `0` (not configured) or `-1` (disabled), call the [ModifyQuotaTemplateServiceStatus](https://help.aliyun.com/document_detail/450406.html) operation to change the status to `1` (enabled).
//
// ### Usage notes
//
// After you create a quota template, you can call the [ListQuotaApplicationsForTemplate](https://help.aliyun.com/document_detail/2584864.html) operation to view the approval result. If `Status` is displayed as `Agree`, the application is approved.
//
// @param request - CreateTemplateQuotaItemRequest
//
// @return CreateTemplateQuotaItemResponse
func (client *Client) CreateTemplateQuotaItem(request *CreateTemplateQuotaItemRequest) (_result *CreateTemplateQuotaItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTemplateQuotaItemResponse{}
	_body, _err := client.CreateTemplateQuotaItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a quota usage alert.
//
// Description:
//
// The quota alerting feature has been upgraded, and this API operation will be deprecated soon. This API operation can only delete legacy quota usage alerts. To delete new-version quota usage alerts, use the CloudMonitor API operations [DeleteMetricRules](https://help.aliyun.com/document_detail/2513295.html) and [DeleteMetricRuleTargets](https://help.aliyun.com/document_detail/2513294.html).
//
// For more information about managing new-version quota alerts by using API operations, see [Manage new-version quota alerts by using API operations](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - DeleteQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQuotaAlarmResponse
func (client *Client) DeleteQuotaAlarmWithOptions(request *DeleteQuotaAlarmRequest, runtime *dara.RuntimeOptions) (_result *DeleteQuotaAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quota usage alert.
//
// Description:
//
// The quota alerting feature has been upgraded, and this API operation will be deprecated soon. This API operation can only delete legacy quota usage alerts. To delete new-version quota usage alerts, use the CloudMonitor API operations [DeleteMetricRules](https://help.aliyun.com/document_detail/2513295.html) and [DeleteMetricRuleTargets](https://help.aliyun.com/document_detail/2513294.html).
//
// For more information about managing new-version quota alerts by using API operations, see [Manage new-version quota alerts by using API operations](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - DeleteQuotaAlarmRequest
//
// @return DeleteQuotaAlarmResponse
func (client *Client) DeleteQuotaAlarm(request *DeleteQuotaAlarmRequest) (_result *DeleteQuotaAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQuotaAlarmResponse{}
	_body, _err := client.DeleteQuotaAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a quota template. After the management account deletes a quota template, the template no longer automatically submits quota increase requests for new members added to the resource directory. Only the management account of the resource directory can delete quota templates.
//
// @param request - DeleteTemplateQuotaItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateQuotaItemResponse
func (client *Client) DeleteTemplateQuotaItemWithOptions(request *DeleteTemplateQuotaItemRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateQuotaItemResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quota template. After the management account deletes a quota template, the template no longer automatically submits quota increase requests for new members added to the resource directory. Only the management account of the resource directory can delete quota templates.
//
// @param request - DeleteTemplateQuotaItemRequest
//
// @return DeleteTemplateQuotaItemResponse
func (client *Client) DeleteTemplateQuotaItem(request *DeleteTemplateQuotaItemRequest) (_result *DeleteTemplateQuotaItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTemplateQuotaItemResponse{}
	_body, _err := client.DeleteTemplateQuotaItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quota details of a specified cloud service.
//
// @param request - GetProductQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductQuotaResponse
func (client *Client) GetProductQuotaWithOptions(request *GetProductQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetProductQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota details of a specified cloud service.
//
// @param request - GetProductQuotaRequest
//
// @return GetProductQuotaResponse
func (client *Client) GetProductQuota(request *GetProductQuotaRequest) (_result *GetProductQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProductQuotaResponse{}
	_body, _err := client.GetProductQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the quota dimensions supported by the target cloud service.
//
// @param request - GetProductQuotaDimensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductQuotaDimensionResponse
func (client *Client) GetProductQuotaDimensionWithOptions(request *GetProductQuotaDimensionRequest, runtime *dara.RuntimeOptions) (_result *GetProductQuotaDimensionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the quota dimensions supported by the target cloud service.
//
// @param request - GetProductQuotaDimensionRequest
//
// @return GetProductQuotaDimensionResponse
func (client *Client) GetProductQuotaDimension(request *GetProductQuotaDimensionRequest) (_result *GetProductQuotaDimensionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProductQuotaDimensionResponse{}
	_body, _err := client.GetProductQuotaDimensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified quota alert.
//
// Description:
//
// The quota alert feature has been upgraded. This API operation will be deprecated soon. This API operation can only query the details of legacy quota usage alerts. To query the details of new-version quota usage alerts, use the CloudMonitor API operation [DescribeMetricRuleList](https://help.aliyun.com/document_detail/2513291.html). For more information about how to use API operations to manage new-version quota alerts, see [Use API operations to manage new-version quota alerts](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - GetQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaAlarmResponse
func (client *Client) GetQuotaAlarmWithOptions(request *GetQuotaAlarmRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified quota alert.
//
// Description:
//
// The quota alert feature has been upgraded. This API operation will be deprecated soon. This API operation can only query the details of legacy quota usage alerts. To query the details of new-version quota usage alerts, use the CloudMonitor API operation [DescribeMetricRuleList](https://help.aliyun.com/document_detail/2513291.html). For more information about how to use API operations to manage new-version quota alerts, see [Use API operations to manage new-version quota alerts](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - GetQuotaAlarmRequest
//
// @return GetQuotaAlarmResponse
func (client *Client) GetQuotaAlarm(request *GetQuotaAlarmRequest) (_result *GetQuotaAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQuotaAlarmResponse{}
	_body, _err := client.GetQuotaAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a quota increase application.
//
// @param request - GetQuotaApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaApplicationResponse
func (client *Client) GetQuotaApplicationWithOptions(request *GetQuotaApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a quota increase application.
//
// @param request - GetQuotaApplicationRequest
//
// @return GetQuotaApplicationResponse
func (client *Client) GetQuotaApplication(request *GetQuotaApplicationRequest) (_result *GetQuotaApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQuotaApplicationResponse{}
	_body, _err := client.GetQuotaApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the approval information of a quota application, including the average approval time, whether approval reminders are supported, and the approval reminder interval.
//
// Description:
//
// ### Before you begin
//
// Make sure that you have created a quota increase application. For more information, see [CreateQuotaApplication](https://help.aliyun.com/document_detail/440566.html).
//
// @param request - GetQuotaApplicationApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaApplicationApprovalResponse
func (client *Client) GetQuotaApplicationApprovalWithOptions(request *GetQuotaApplicationApprovalRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaApplicationApprovalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the approval information of a quota application, including the average approval time, whether approval reminders are supported, and the approval reminder interval.
//
// Description:
//
// ### Before you begin
//
// Make sure that you have created a quota increase application. For more information, see [CreateQuotaApplication](https://help.aliyun.com/document_detail/440566.html).
//
// @param request - GetQuotaApplicationApprovalRequest
//
// @return GetQuotaApplicationApprovalResponse
func (client *Client) GetQuotaApplicationApproval(request *GetQuotaApplicationApprovalRequest) (_result *GetQuotaApplicationApprovalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQuotaApplicationApprovalResponse{}
	_body, _err := client.GetQuotaApplicationApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the quota template.
//
// Description:
//
// The default status of a quota template is unset (that is, `ServiceStatus` is `0`). You can use the features of the quota template only when its status is enabled (that is, `ServiceStatus` is `1`).
//
// @param request - GetQuotaTemplateServiceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaTemplateServiceStatusResponse
func (client *Client) GetQuotaTemplateServiceStatusWithOptions(request *GetQuotaTemplateServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaTemplateServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the quota template.
//
// Description:
//
// The default status of a quota template is unset (that is, `ServiceStatus` is `0`). You can use the features of the quota template only when its status is enabled (that is, `ServiceStatus` is `1`).
//
// @param request - GetQuotaTemplateServiceStatusRequest
//
// @return GetQuotaTemplateServiceStatusResponse
func (client *Client) GetQuotaTemplateServiceStatus(request *GetQuotaTemplateServiceStatusRequest) (_result *GetQuotaTemplateServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQuotaTemplateServiceStatusResponse{}
	_body, _err := client.GetQuotaTemplateServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取配额使用率概览
//
// @param request - GetQuotaUtilizationOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaUtilizationOverviewResponse
func (client *Client) GetQuotaUtilizationOverviewWithOptions(request *GetQuotaUtilizationOverviewRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaUtilizationOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuotaUtilizationOverview"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaUtilizationOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取配额使用率概览
//
// @param request - GetQuotaUtilizationOverviewRequest
//
// @return GetQuotaUtilizationOverviewResponse
func (client *Client) GetQuotaUtilizationOverview(request *GetQuotaUtilizationOverviewRequest) (_result *GetQuotaUtilizationOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQuotaUtilizationOverviewResponse{}
	_body, _err := client.GetQuotaUtilizationOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the history of quota alerts.
//
// Description:
//
// The quota alert feature has been upgraded and this API operation will be deprecated. This API operation can only query the history of legacy quota usage alerts. To query the history of new quota usage alerts, use the CloudMonitor API operations [DescribeAlertLogCount](https://help.aliyun.com/document_detail/2513275.html) and [DescribeAlertLogList](https://help.aliyun.com/document_detail/2513276.html). For more information about how to use API operations to manage new quota alerts, see [Use API operations to manage new quota alerts](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - ListAlarmHistoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlarmHistoriesResponse
func (client *Client) ListAlarmHistoriesWithOptions(request *ListAlarmHistoriesRequest, runtime *dara.RuntimeOptions) (_result *ListAlarmHistoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the history of quota alerts.
//
// Description:
//
// The quota alert feature has been upgraded and this API operation will be deprecated. This API operation can only query the history of legacy quota usage alerts. To query the history of new quota usage alerts, use the CloudMonitor API operations [DescribeAlertLogCount](https://help.aliyun.com/document_detail/2513275.html) and [DescribeAlertLogList](https://help.aliyun.com/document_detail/2513276.html). For more information about how to use API operations to manage new quota alerts, see [Use API operations to manage new quota alerts](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - ListAlarmHistoriesRequest
//
// @return ListAlarmHistoriesResponse
func (client *Client) ListAlarmHistories(request *ListAlarmHistoriesRequest) (_result *ListAlarmHistoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAlarmHistoriesResponse{}
	_body, _err := client.ListAlarmHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of quotas associated with a specific quota. Some quotas are interdependent and should be increased simultaneously.
//
// @param request - ListDependentQuotasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDependentQuotasResponse
func (client *Client) ListDependentQuotasWithOptions(request *ListDependentQuotasRequest, runtime *dara.RuntimeOptions) (_result *ListDependentQuotasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of quotas associated with a specific quota. Some quotas are interdependent and should be increased simultaneously.
//
// @param request - ListDependentQuotasRequest
//
// @return ListDependentQuotasResponse
func (client *Client) ListDependentQuotas(request *ListDependentQuotasRequest) (_result *ListDependentQuotasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDependentQuotasResponse{}
	_body, _err := client.ListDependentQuotasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation queries the dimension groups of a specified Alibaba Cloud service.
//
// @param request - ListProductDimensionGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductDimensionGroupsResponse
func (client *Client) ListProductDimensionGroupsWithOptions(request *ListProductDimensionGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListProductDimensionGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation queries the dimension groups of a specified Alibaba Cloud service.
//
// @param request - ListProductDimensionGroupsRequest
//
// @return ListProductDimensionGroupsResponse
func (client *Client) ListProductDimensionGroups(request *ListProductDimensionGroupsRequest) (_result *ListProductDimensionGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProductDimensionGroupsResponse{}
	_body, _err := client.ListProductDimensionGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quota dimensions supported by the specified cloud service.
//
// @param request - ListProductQuotaDimensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductQuotaDimensionsResponse
func (client *Client) ListProductQuotaDimensionsWithOptions(request *ListProductQuotaDimensionsRequest, runtime *dara.RuntimeOptions) (_result *ListProductQuotaDimensionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota dimensions supported by the specified cloud service.
//
// @param request - ListProductQuotaDimensionsRequest
//
// @return ListProductQuotaDimensionsResponse
func (client *Client) ListProductQuotaDimensions(request *ListProductQuotaDimensionsRequest) (_result *ListProductQuotaDimensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProductQuotaDimensionsResponse{}
	_body, _err := client.ListProductQuotaDimensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quotas of a specified cloud service.
//
// @param request - ListProductQuotasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductQuotasResponse
func (client *Client) ListProductQuotasWithOptions(request *ListProductQuotasRequest, runtime *dara.RuntimeOptions) (_result *ListProductQuotasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quotas of a specified cloud service.
//
// @param request - ListProductQuotasRequest
//
// @return ListProductQuotasResponse
func (client *Client) ListProductQuotas(request *ListProductQuotasRequest) (_result *ListProductQuotasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProductQuotasResponse{}
	_body, _err := client.ListProductQuotasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Alibaba Cloud services supported by Quota Center.
//
// Description:
//
// The query results are the same as those in [Alibaba Cloud services supported by Quota Center](https://help.aliyun.com/document_detail/182368.html).
//
// @param request - ListProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithOptions(request *ListProductsRequest, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Alibaba Cloud services supported by Quota Center.
//
// Description:
//
// The query results are the same as those in [Alibaba Cloud services supported by Quota Center](https://help.aliyun.com/document_detail/182368.html).
//
// @param request - ListProductsRequest
//
// @return ListProductsResponse
func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries quota usage alerts.
//
// Description:
//
// The quota alert feature has been upgraded. This API operation is scheduled for deprecation and queries only earlier-version quota usage alerts. To query new-version quota usage alerts, call the CloudMonitor API operation [DescribeMetricRuleList](https://help.aliyun.com/document_detail/2513291.html). For more information about how to manage new-version quota alerts using API operations, see [Use API operations to manage new-version quota alerts](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - ListQuotaAlarmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaAlarmsResponse
func (client *Client) ListQuotaAlarmsWithOptions(request *ListQuotaAlarmsRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaAlarmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quota usage alerts.
//
// Description:
//
// The quota alert feature has been upgraded. This API operation is scheduled for deprecation and queries only earlier-version quota usage alerts. To query new-version quota usage alerts, call the CloudMonitor API operation [DescribeMetricRuleList](https://help.aliyun.com/document_detail/2513291.html). For more information about how to manage new-version quota alerts using API operations, see [Use API operations to manage new-version quota alerts](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - ListQuotaAlarmsRequest
//
// @return ListQuotaAlarmsResponse
func (client *Client) ListQuotaAlarms(request *ListQuotaAlarmsRequest) (_result *ListQuotaAlarmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQuotaAlarmsResponse{}
	_body, _err := client.ListQuotaAlarmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListQuotaApplicationTemplatesWithOptions(request *ListQuotaApplicationTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaApplicationTemplatesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListQuotaApplicationTemplatesResponse
func (client *Client) ListQuotaApplicationTemplates(request *ListQuotaApplicationTemplatesRequest) (_result *ListQuotaApplicationTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQuotaApplicationTemplatesResponse{}
	_body, _err := client.ListQuotaApplicationTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of quota increase applications.
//
// @param request - ListQuotaApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationsResponse
func (client *Client) ListQuotaApplicationsWithOptions(request *ListQuotaApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of quota increase applications.
//
// @param request - ListQuotaApplicationsRequest
//
// @return ListQuotaApplicationsResponse
func (client *Client) ListQuotaApplications(request *ListQuotaApplicationsRequest) (_result *ListQuotaApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQuotaApplicationsResponse{}
	_body, _err := client.ListQuotaApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of quota increase applications for a member in a resource directory.
//
// @param request - ListQuotaApplicationsDetailForTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationsDetailForTemplateResponse
func (client *Client) ListQuotaApplicationsDetailForTemplateWithOptions(request *ListQuotaApplicationsDetailForTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaApplicationsDetailForTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of quota increase applications for a member in a resource directory.
//
// @param request - ListQuotaApplicationsDetailForTemplateRequest
//
// @return ListQuotaApplicationsDetailForTemplateResponse
func (client *Client) ListQuotaApplicationsDetailForTemplate(request *ListQuotaApplicationsDetailForTemplateRequest) (_result *ListQuotaApplicationsDetailForTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQuotaApplicationsDetailForTemplateResponse{}
	_body, _err := client.ListQuotaApplicationsDetailForTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the history of batch requests for member quotas submitted using a quota template.
//
// @param request - ListQuotaApplicationsForTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaApplicationsForTemplateResponse
func (client *Client) ListQuotaApplicationsForTemplateWithOptions(request *ListQuotaApplicationsForTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaApplicationsForTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the history of batch requests for member quotas submitted using a quota template.
//
// @param request - ListQuotaApplicationsForTemplateRequest
//
// @return ListQuotaApplicationsForTemplateResponse
func (client *Client) ListQuotaApplicationsForTemplate(request *ListQuotaApplicationsForTemplateRequest) (_result *ListQuotaApplicationsForTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQuotaApplicationsForTemplateResponse{}
	_body, _err := client.ListQuotaApplicationsForTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举配额使用率报告
//
// @param request - ListQuotaUtilizationReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaUtilizationReportResponse
func (client *Client) ListQuotaUtilizationReportWithOptions(request *ListQuotaUtilizationReportRequest, runtime *dara.RuntimeOptions) (_result *ListQuotaUtilizationReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Adjustable) {
		body["Adjustable"] = request.Adjustable
	}

	if !dara.IsNil(request.KeyWords) {
		body["KeyWords"] = request.KeyWords
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

	if !dara.IsNil(request.QuotaCategory) {
		body["QuotaCategory"] = request.QuotaCategory
	}

	if !dara.IsNil(request.QuotaRegionId) {
		body["QuotaRegionId"] = request.QuotaRegionId
	}

	if !dara.IsNil(request.UtilizationStatuses) {
		body["UtilizationStatuses"] = request.UtilizationStatuses
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuotaUtilizationReport"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaUtilizationReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举配额使用率报告
//
// @param request - ListQuotaUtilizationReportRequest
//
// @return ListQuotaUtilizationReportResponse
func (client *Client) ListQuotaUtilizationReport(request *ListQuotaUtilizationReportRequest) (_result *ListQuotaUtilizationReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQuotaUtilizationReportResponse{}
	_body, _err := client.ListQuotaUtilizationReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyQuotaTemplateServiceStatusWithOptions(request *ModifyQuotaTemplateServiceStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyQuotaTemplateServiceStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyQuotaTemplateServiceStatusResponse
func (client *Client) ModifyQuotaTemplateServiceStatus(request *ModifyQuotaTemplateServiceStatusRequest) (_result *ModifyQuotaTemplateServiceStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyQuotaTemplateServiceStatusResponse{}
	_body, _err := client.ModifyQuotaTemplateServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Only the administrator account of a resource directory can modify a quota template. Modifications to a quota template apply only to new members that are added to the resource directory and do not affect existing members.
//
// @param request - ModifyTemplateQuotaItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateQuotaItemResponse
func (client *Client) ModifyTemplateQuotaItemWithOptions(request *ModifyTemplateQuotaItemRequest, runtime *dara.RuntimeOptions) (_result *ModifyTemplateQuotaItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Only the administrator account of a resource directory can modify a quota template. Modifications to a quota template apply only to new members that are added to the resource directory and do not affect existing members.
//
// @param request - ModifyTemplateQuotaItemRequest
//
// @return ModifyTemplateQuotaItemResponse
func (client *Client) ModifyTemplateQuotaItem(request *ModifyTemplateQuotaItemRequest) (_result *ModifyTemplateQuotaItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTemplateQuotaItemResponse{}
	_body, _err := client.ModifyTemplateQuotaItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a reminder for a quota application that is pending approval.
//
// Description:
//
// > This operation applies only to quota applications that support reminders. In the GetQuotaApplicationApproval response, the `SupportReminder` parameter for these applications is `true`.
//
// @param request - RemindQuotaApplicationApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemindQuotaApplicationApprovalResponse
func (client *Client) RemindQuotaApplicationApprovalWithOptions(request *RemindQuotaApplicationApprovalRequest, runtime *dara.RuntimeOptions) (_result *RemindQuotaApplicationApprovalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a reminder for a quota application that is pending approval.
//
// Description:
//
// > This operation applies only to quota applications that support reminders. In the GetQuotaApplicationApproval response, the `SupportReminder` parameter for these applications is `true`.
//
// @param request - RemindQuotaApplicationApprovalRequest
//
// @return RemindQuotaApplicationApprovalResponse
func (client *Client) RemindQuotaApplicationApproval(request *RemindQuotaApplicationApprovalRequest) (_result *RemindQuotaApplicationApprovalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemindQuotaApplicationApprovalResponse{}
	_body, _err := client.RemindQuotaApplicationApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发起配额使用率报告查询
//
// @param request - StartQuotaUtilizationReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartQuotaUtilizationReportResponse
func (client *Client) StartQuotaUtilizationReportWithOptions(request *StartQuotaUtilizationReportRequest, runtime *dara.RuntimeOptions) (_result *StartQuotaUtilizationReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("StartQuotaUtilizationReport"),
		Version:     dara.String("2020-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartQuotaUtilizationReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起配额使用率报告查询
//
// @param request - StartQuotaUtilizationReportRequest
//
// @return StartQuotaUtilizationReportResponse
func (client *Client) StartQuotaUtilizationReport(request *StartQuotaUtilizationReportRequest) (_result *StartQuotaUtilizationReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartQuotaUtilizationReportResponse{}
	_body, _err := client.StartQuotaUtilizationReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies quota alert information.
//
// Description:
//
// This API operation is deprecated because the quota alert feature has been upgraded. To modify information about new quota usage alerts, use the CloudMonitor API operations [PutResourceMetricRules](https://help.aliyun.com/document_detail/2513316.html) and [PutMetricRuleTargets](https://help.aliyun.com/document_detail/2513302.html). For more information about using API operations to manage new quota alerts, see [Use API operations to manage new quota alerts](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - UpdateQuotaAlarmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaAlarmResponse
func (client *Client) UpdateQuotaAlarmWithOptions(request *UpdateQuotaAlarmRequest, runtime *dara.RuntimeOptions) (_result *UpdateQuotaAlarmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies quota alert information.
//
// Description:
//
// This API operation is deprecated because the quota alert feature has been upgraded. To modify information about new quota usage alerts, use the CloudMonitor API operations [PutResourceMetricRules](https://help.aliyun.com/document_detail/2513316.html) and [PutMetricRuleTargets](https://help.aliyun.com/document_detail/2513302.html). For more information about using API operations to manage new quota alerts, see [Use API operations to manage new quota alerts](https://help.aliyun.com/document_detail/2863234.html).
//
// @param request - UpdateQuotaAlarmRequest
//
// @return UpdateQuotaAlarmResponse
func (client *Client) UpdateQuotaAlarm(request *UpdateQuotaAlarmRequest) (_result *UpdateQuotaAlarmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQuotaAlarmResponse{}
	_body, _err := client.UpdateQuotaAlarmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
