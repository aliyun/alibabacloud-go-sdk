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
	client.EndpointMap = map[string]*string{
		"us-west-1":             dara.String("metrics.us-west-1.aliyuncs.com"),
		"us-southeast-1":        dara.String("metrics.us-southeast-1.aliyuncs.com"),
		"us-east-1":             dara.String("metrics.us-east-1.aliyuncs.com"),
		"na-south-1":            dara.String("metrics.na-south-1.aliyuncs.com"),
		"me-east-1":             dara.String("metrics.me-east-1.aliyuncs.com"),
		"me-central-1":          dara.String("metrics.me-central-1.aliyuncs.com"),
		"eu-west-2":             dara.String("metrics.eu-west-2.aliyuncs.com"),
		"eu-west-1":             dara.String("metrics.eu-west-1.aliyuncs.com"),
		"eu-central-1":          dara.String("metrics.eu-central-1.aliyuncs.com"),
		"cn-zhongwei":           dara.String("metrics.cn-zhongwei.aliyuncs.com"),
		"cn-zhengzhou-jva":      dara.String("metrics.cn-zhengzhou-jva.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("metrics.cn-zhangjiakou.aliyuncs.com"),
		"cn-wulanchabu-gic-1":   dara.String("metrics.cn-wulanchabu-gic-1.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("metrics.cn-wulanchabu.aliyuncs.com"),
		"cn-wuhan-lr":           dara.String("metrics.cn-wuhan-lr.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("metrics.cn-shenzhen-finance-1.aliyuncs.com"),
		"cn-shenzhen":           dara.String("metrics.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("metrics.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shanghai":           dara.String("metrics.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":            dara.String("metrics.cn-qingdao.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("metrics.cn-north-2-gov-1.aliyuncs.com"),
		"cn-nanjing":            dara.String("metrics.cn-nanjing.aliyuncs.com"),
		"cn-huhehaote":          dara.String("metrics.cn-huhehaote.aliyuncs.com"),
		"cn-hongkong":           dara.String("metrics.cn-hongkong.aliyuncs.com"),
		"cn-heyuan-acdr-1":      dara.String("metrics.cn-heyuan-acdr-1.aliyuncs.com"),
		"cn-heyuan":             dara.String("metrics.cn-heyuan.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("metrics.cn-hangzhou-finance.aliyuncs.com"),
		"cn-hangzhou":           dara.String("metrics.cn-hangzhou.aliyuncs.com"),
		"cn-guangzhou":          dara.String("metrics.cn-guangzhou.aliyuncs.com"),
		"cn-fuzhou":             dara.String("metrics.cn-fuzhou.aliyuncs.com"),
		"cn-chengdu":            dara.String("metrics.cn-chengdu.aliyuncs.com"),
		"cn-beijing-finance-1":  dara.String("metrics.cn-beijing-finance-1.aliyuncs.com"),
		"cn-beijing":            dara.String("metrics.cn-beijing.aliyuncs.com"),
		"ap-southeast-8":        dara.String("metrics.ap-southeast-8.aliyuncs.com"),
		"ap-southeast-7":        dara.String("metrics.ap-southeast-7.aliyuncs.com"),
		"ap-southeast-6":        dara.String("metrics.ap-southeast-6.aliyuncs.com"),
		"ap-southeast-5":        dara.String("metrics.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-3":        dara.String("metrics.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-2":        dara.String("metrics.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-1":        dara.String("metrics.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            dara.String("metrics.ap-south-1.aliyuncs.com"),
		"ap-northeast-2":        dara.String("metrics.ap-northeast-2.aliyuncs.com"),
		"ap-northeast-1":        dara.String("metrics.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// The AddTags operation attaches tags to specified application groups.
//
// Description:
//
// This topic provides an example of how to attach a tag to the application group `7301****`. In this example, the tag key is `key1` and the tag value is `value1`.
//
// @param request - AddTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTagsResponse
func (client *Client) AddTagsWithOptions(request *AddTagsRequest, runtime *dara.RuntimeOptions) (_result *AddTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTags"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The AddTags operation attaches tags to specified application groups.
//
// Description:
//
// This topic provides an example of how to attach a tag to the application group `7301****`. In this example, the tag key is `key1` and the tag value is `value1`.
//
// @param request - AddTagsRequest
//
// @return AddTagsResponse
func (client *Client) AddTags(request *AddTagsRequest) (_result *AddTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTagsResponse{}
	_body, _err := client.AddTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies an alert template to an application group to generate an alert rule.
//
// Description:
//
// In this example, the `700****` alert template is applied to the `123456` application group. For the generated alert rule, the ID is `applyTemplate8ab74c6b-9f27-47ab-8841-de01dc08****`, and the name is `test123`.
//
// @param request - ApplyMetricRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyMetricRuleTemplateResponse
func (client *Client) ApplyMetricRuleTemplateWithOptions(request *ApplyMetricRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *ApplyMetricRuleTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppendMode) {
		query["AppendMode"] = request.AppendMode
	}

	if !dara.IsNil(request.ApplyMode) {
		query["ApplyMode"] = request.ApplyMode
	}

	if !dara.IsNil(request.EnableEndTime) {
		query["EnableEndTime"] = request.EnableEndTime
	}

	if !dara.IsNil(request.EnableStartTime) {
		query["EnableStartTime"] = request.EnableStartTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.NotifyLevel) {
		query["NotifyLevel"] = request.NotifyLevel
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyMetricRuleTemplate"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyMetricRuleTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies an alert template to an application group to generate an alert rule.
//
// Description:
//
// In this example, the `700****` alert template is applied to the `123456` application group. For the generated alert rule, the ID is `applyTemplate8ab74c6b-9f27-47ab-8841-de01dc08****`, and the name is `test123`.
//
// @param request - ApplyMetricRuleTemplateRequest
//
// @return ApplyMetricRuleTemplateResponse
func (client *Client) ApplyMetricRuleTemplate(request *ApplyMetricRuleTemplateRequest) (_result *ApplyMetricRuleTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyMetricRuleTemplateResponse{}
	_body, _err := client.ApplyMetricRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the BatchCreateInstantSiteMonitor operation to create a batch of site monitoring tasks.
//
// Description:
//
// This topic provides an example of how to create a site monitoring task named `HangZhou_ECS1`. The task uses the `HTTP` protocol to monitor `https://www.aliyun.com`. The response shows that the task is successfully created with the name `HangZhou_ECS1` and the ID `679fbe4f-b80b-4706-91b2-5427b43e****`.
//
// @param request - BatchCreateInstantSiteMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateInstantSiteMonitorResponse
func (client *Client) BatchCreateInstantSiteMonitorWithOptions(request *BatchCreateInstantSiteMonitorRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateInstantSiteMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskList) {
		query["TaskList"] = request.TaskList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateInstantSiteMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateInstantSiteMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the BatchCreateInstantSiteMonitor operation to create a batch of site monitoring tasks.
//
// Description:
//
// This topic provides an example of how to create a site monitoring task named `HangZhou_ECS1`. The task uses the `HTTP` protocol to monitor `https://www.aliyun.com`. The response shows that the task is successfully created with the name `HangZhou_ECS1` and the ID `679fbe4f-b80b-4706-91b2-5427b43e****`.
//
// @param request - BatchCreateInstantSiteMonitorRequest
//
// @return BatchCreateInstantSiteMonitorResponse
func (client *Client) BatchCreateInstantSiteMonitor(request *BatchCreateInstantSiteMonitorRequest) (_result *BatchCreateInstantSiteMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCreateInstantSiteMonitorResponse{}
	_body, _err := client.BatchCreateInstantSiteMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the BatchExport operation to export monitoring data defined in the Cursor operation.
//
// Description:
//
// ### Prerequisites
//
// Make sure that you have called the [Cursor](https://help.aliyun.com/document_detail/2330730.html) operation to obtain the initial `Cursor`.
//
// ### Usage notes
//
// This topic provides an example to show how to export the monitoring data of an initial `Cursor` of the metric `cpu_idle` of the cloud service `acs_ecs_dashboard`. A maximum of 1,000 data entries are returned per call.
//
// @param tmpReq - BatchExportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchExportResponse
func (client *Client) BatchExportWithOptions(tmpReq *BatchExportRequest, runtime *dara.RuntimeOptions) (_result *BatchExportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchExportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Measurements) {
		request.MeasurementsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Measurements, dara.String("Measurements"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Cursor) {
		body["Cursor"] = request.Cursor
	}

	if !dara.IsNil(request.Length) {
		body["Length"] = request.Length
	}

	if !dara.IsNil(request.MeasurementsShrink) {
		body["Measurements"] = request.MeasurementsShrink
	}

	if !dara.IsNil(request.Metric) {
		body["Metric"] = request.Metric
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchExport"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the BatchExport operation to export monitoring data defined in the Cursor operation.
//
// Description:
//
// ### Prerequisites
//
// Make sure that you have called the [Cursor](https://help.aliyun.com/document_detail/2330730.html) operation to obtain the initial `Cursor`.
//
// ### Usage notes
//
// This topic provides an example to show how to export the monitoring data of an initial `Cursor` of the metric `cpu_idle` of the cloud service `acs_ecs_dashboard`. A maximum of 1,000 data entries are returned per call.
//
// @param request - BatchExportRequest
//
// @return BatchExportResponse
func (client *Client) BatchExport(request *BatchExportRequest) (_result *BatchExportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchExportResponse{}
	_body, _err := client.BatchExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the CreateDynamicTagGroup operation to automatically create application groups using tags.
//
// Description:
//
// This operation supports the following Alibaba Cloud services: Elastic Compute Service (ECS), ApsaraDB RDS, and Server Load Balancer (SLB).
//
// This topic provides an example of how to automatically create an application group for resources that have the `ecs_instance` tag key. The alert contact group for the application group is `ECS_Group`.
//
// @param request - CreateDynamicTagGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDynamicTagGroupResponse
func (client *Client) CreateDynamicTagGroupWithOptions(request *CreateDynamicTagGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDynamicTagGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupList) {
		query["ContactGroupList"] = request.ContactGroupList
	}

	if !dara.IsNil(request.EnableInstallAgent) {
		query["EnableInstallAgent"] = request.EnableInstallAgent
	}

	if !dara.IsNil(request.EnableSubscribeEvent) {
		query["EnableSubscribeEvent"] = request.EnableSubscribeEvent
	}

	if !dara.IsNil(request.MatchExpress) {
		query["MatchExpress"] = request.MatchExpress
	}

	if !dara.IsNil(request.MatchExpressFilterRelation) {
		query["MatchExpressFilterRelation"] = request.MatchExpressFilterRelation
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	if !dara.IsNil(request.TagRegionId) {
		query["TagRegionId"] = request.TagRegionId
	}

	if !dara.IsNil(request.TemplateIdList) {
		query["TemplateIdList"] = request.TemplateIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDynamicTagGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDynamicTagGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the CreateDynamicTagGroup operation to automatically create application groups using tags.
//
// Description:
//
// This operation supports the following Alibaba Cloud services: Elastic Compute Service (ECS), ApsaraDB RDS, and Server Load Balancer (SLB).
//
// This topic provides an example of how to automatically create an application group for resources that have the `ecs_instance` tag key. The alert contact group for the application group is `ECS_Group`.
//
// @param request - CreateDynamicTagGroupRequest
//
// @return CreateDynamicTagGroupResponse
func (client *Client) CreateDynamicTagGroup(request *CreateDynamicTagGroupRequest) (_result *CreateDynamicTagGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDynamicTagGroupResponse{}
	_body, _err := client.CreateDynamicTagGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates alert rules in batches for a specified application group by calling the CreateGroupMetricRules operation.
//
// Description:
//
// This topic provides an example on how to create an alert rule for the `cpu_total` metric of Elastic Compute Service (ECS) in the application group `123456`. The alert rule ID is `456789`, the alert rule name is `ECS_Rule1`, the alert severity is `Critical`, the statistical method is `Average`, the comparison operator is `GreaterThanOrEqualToThreshold`, the threshold is `90`, and the retry count is `3`. The response shows that the alert rule `ECS_Rule1` is created.
//
// @param request - CreateGroupMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupMetricRulesResponse
func (client *Client) CreateGroupMetricRulesWithOptions(request *CreateGroupMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupMetricRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupMetricRules) {
		query["GroupMetricRules"] = request.GroupMetricRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroupMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupMetricRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates alert rules in batches for a specified application group by calling the CreateGroupMetricRules operation.
//
// Description:
//
// This topic provides an example on how to create an alert rule for the `cpu_total` metric of Elastic Compute Service (ECS) in the application group `123456`. The alert rule ID is `456789`, the alert rule name is `ECS_Rule1`, the alert severity is `Critical`, the statistical method is `Average`, the comparison operator is `GreaterThanOrEqualToThreshold`, the threshold is `90`, and the retry count is `3`. The response shows that the alert rule `ECS_Rule1` is created.
//
// @param request - CreateGroupMetricRulesRequest
//
// @return CreateGroupMetricRulesResponse
func (client *Client) CreateGroupMetricRules(request *CreateGroupMetricRulesRequest) (_result *CreateGroupMetricRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGroupMetricRulesResponse{}
	_body, _err := client.CreateGroupMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a process monitoring task for an application group.
//
// @param request - CreateGroupMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupMonitoringAgentProcessResponse
func (client *Client) CreateGroupMonitoringAgentProcessWithOptions(request *CreateGroupMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *CreateGroupMonitoringAgentProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MatchExpress) {
		query["MatchExpress"] = request.MatchExpress
	}

	if !dara.IsNil(request.MatchExpressFilterRelation) {
		query["MatchExpressFilterRelation"] = request.MatchExpressFilterRelation
	}

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGroupMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a process monitoring task for an application group.
//
// @param request - CreateGroupMonitoringAgentProcessRequest
//
// @return CreateGroupMonitoringAgentProcessResponse
func (client *Client) CreateGroupMonitoringAgentProcess(request *CreateGroupMonitoringAgentProcessRequest) (_result *CreateGroupMonitoringAgentProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CreateGroupMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an availability monitoring task by calling the CreateHostAvailability operation.
//
// Description:
//
// This topic provides an example of how to create an availability monitoring task named `task1` with the detection type set to `HTTP` in application group `123456`. Alert notifications are sent by email and DingTalk chatbot.
//
// @param request - CreateHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHostAvailabilityResponse
func (client *Client) CreateHostAvailabilityWithOptions(request *CreateHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *CreateHostAvailabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfigEscalationList) {
		query["AlertConfigEscalationList"] = request.AlertConfigEscalationList
	}

	if !dara.IsNil(request.AlertConfigTargetList) {
		query["AlertConfigTargetList"] = request.AlertConfigTargetList
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceList) {
		query["InstanceList"] = request.InstanceList
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskScope) {
		query["TaskScope"] = request.TaskScope
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.TaskOption) {
		query["TaskOption"] = request.TaskOption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHostAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an availability monitoring task by calling the CreateHostAvailability operation.
//
// Description:
//
// This topic provides an example of how to create an availability monitoring task named `task1` with the detection type set to `HTTP` in application group `123456`. Alert notifications are sent by email and DingTalk chatbot.
//
// @param request - CreateHostAvailabilityRequest
//
// @return CreateHostAvailabilityResponse
func (client *Client) CreateHostAvailability(request *CreateHostAvailabilityRequest) (_result *CreateHostAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHostAvailabilityResponse{}
	_body, _err := client.CreateHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the CreateHybridMonitorNamespace operation to create a metric repository.
//
// Description:
//
// ## Before you begin
//
// Make sure that you have activated Hybrid Cloud Monitoring. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// ## Operation description
//
// This topic provides an example on how to create a metric repository named `aliyun` with a data storage duration of `cms.s1.3xlarge`. The response indicates that the metric repository is created.
//
// @param request - CreateHybridMonitorNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridMonitorNamespaceResponse
func (client *Client) CreateHybridMonitorNamespaceWithOptions(request *CreateHybridMonitorNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateHybridMonitorNamespaceResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceRegion) {
		query["NamespaceRegion"] = request.NamespaceRegion
	}

	if !dara.IsNil(request.NamespaceType) {
		query["NamespaceType"] = request.NamespaceType
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHybridMonitorNamespace"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHybridMonitorNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateHybridMonitorNamespace operation to create a metric repository.
//
// Description:
//
// ## Before you begin
//
// Make sure that you have activated Hybrid Cloud Monitoring. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// ## Operation description
//
// This topic provides an example on how to create a metric repository named `aliyun` with a data storage duration of `cms.s1.3xlarge`. The response indicates that the metric repository is created.
//
// @param request - CreateHybridMonitorNamespaceRequest
//
// @return CreateHybridMonitorNamespaceResponse
func (client *Client) CreateHybridMonitorNamespace(request *CreateHybridMonitorNamespaceRequest) (_result *CreateHybridMonitorNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHybridMonitorNamespaceResponse{}
	_body, _err := client.CreateHybridMonitorNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Logstore group for Hybrid Cloud Monitoring.
//
// Description:
//
// ### Before you begin
//
// Make sure that you have activated Simple Log Service (SLS) and created a project and a Logstore. For more information, see [Quick Start](https://help.aliyun.com/document_detail/54604.html).
//
// ### Operation description
//
// This topic provides an example on how to create a Logstore group named `Logstore_test`. The region is `ap-southeast-1`, the project is `aliyun-project`, and the Logstore is `Logstore-ECS`. The response shows that the Logstore group is created.
//
// @param request - CreateHybridMonitorSLSGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridMonitorSLSGroupResponse
func (client *Client) CreateHybridMonitorSLSGroupWithOptions(request *CreateHybridMonitorSLSGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateHybridMonitorSLSGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SLSGroupConfig) {
		query["SLSGroupConfig"] = request.SLSGroupConfig
	}

	if !dara.IsNil(request.SLSGroupDescription) {
		query["SLSGroupDescription"] = request.SLSGroupDescription
	}

	if !dara.IsNil(request.SLSGroupName) {
		query["SLSGroupName"] = request.SLSGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHybridMonitorSLSGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHybridMonitorSLSGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Logstore group for Hybrid Cloud Monitoring.
//
// Description:
//
// ### Before you begin
//
// Make sure that you have activated Simple Log Service (SLS) and created a project and a Logstore. For more information, see [Quick Start](https://help.aliyun.com/document_detail/54604.html).
//
// ### Operation description
//
// This topic provides an example on how to create a Logstore group named `Logstore_test`. The region is `ap-southeast-1`, the project is `aliyun-project`, and the Logstore is `Logstore-ECS`. The response shows that the Logstore group is created.
//
// @param request - CreateHybridMonitorSLSGroupRequest
//
// @return CreateHybridMonitorSLSGroupResponse
func (client *Client) CreateHybridMonitorSLSGroup(request *CreateHybridMonitorSLSGroupRequest) (_result *CreateHybridMonitorSLSGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHybridMonitorSLSGroupResponse{}
	_body, _err := client.CreateHybridMonitorSLSGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the CreateHybridMonitorTask operation to create a data import task for an Alibaba Cloud service or a metric for Simple Log Service (SLS) logs.
//
// Description:
//
// ## Before you begin
//
// - Make sure that you have activated Hybrid Cloud Monitoring. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// - If you want to create a metric for SLS logs, make sure that you have activated SLS and created a project and a Logstore. For more information, see [Quick Start](https://help.aliyun.com/document_detail/54604.html).
//
// ## Operation description
//
// This topic provides an example of how to create a data import task named `aliyun_task` for an Alibaba Cloud service to import the `cpu_total` metric of Elastic Compute Service (ECS) into the `aliyun` metric repository. The response shows that the data import task is created.
//
// @param request - CreateHybridMonitorTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridMonitorTaskResponse
func (client *Client) CreateHybridMonitorTaskWithOptions(request *CreateHybridMonitorTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateHybridMonitorTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachLabels) {
		query["AttachLabels"] = request.AttachLabels
	}

	if !dara.IsNil(request.CloudAccessId) {
		query["CloudAccessId"] = request.CloudAccessId
	}

	if !dara.IsNil(request.CollectInterval) {
		query["CollectInterval"] = request.CollectInterval
	}

	if !dara.IsNil(request.CollectTargetType) {
		query["CollectTargetType"] = request.CollectTargetType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.SLSProcessConfig) {
		query["SLSProcessConfig"] = request.SLSProcessConfig
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	if !dara.IsNil(request.TargetUserIdList) {
		query["TargetUserIdList"] = request.TargetUserIdList
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.YARMConfig) {
		query["YARMConfig"] = request.YARMConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHybridMonitorTask"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHybridMonitorTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateHybridMonitorTask operation to create a data import task for an Alibaba Cloud service or a metric for Simple Log Service (SLS) logs.
//
// Description:
//
// ## Before you begin
//
// - Make sure that you have activated Hybrid Cloud Monitoring. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// - If you want to create a metric for SLS logs, make sure that you have activated SLS and created a project and a Logstore. For more information, see [Quick Start](https://help.aliyun.com/document_detail/54604.html).
//
// ## Operation description
//
// This topic provides an example of how to create a data import task named `aliyun_task` for an Alibaba Cloud service to import the `cpu_total` metric of Elastic Compute Service (ECS) into the `aliyun` metric repository. The response shows that the data import task is created.
//
// @param request - CreateHybridMonitorTaskRequest
//
// @return CreateHybridMonitorTaskResponse
func (client *Client) CreateHybridMonitorTask(request *CreateHybridMonitorTaskRequest) (_result *CreateHybridMonitorTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHybridMonitorTaskResponse{}
	_body, _err := client.CreateHybridMonitorTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the CreateInstantSiteMonitor operation to create a one-time detection task.
//
// Description:
//
// Only Alibaba Cloud accounts that have Network Analysis and Monitoring activated can create one-time detection tasks.
//
// This topic provides an example of how to create a one-time detection task. The example creates a task named `task1` that detects the address `http://www.aliyun.com`. The detection type is `HTTP`, and the number of detection points is `1`.
//
// @param request - CreateInstantSiteMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstantSiteMonitorResponse
func (client *Client) CreateInstantSiteMonitorWithOptions(request *CreateInstantSiteMonitorRequest, runtime *dara.RuntimeOptions) (_result *CreateInstantSiteMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AgentGroup) {
		query["AgentGroup"] = request.AgentGroup
	}

	if !dara.IsNil(request.IspCities) {
		query["IspCities"] = request.IspCities
	}

	if !dara.IsNil(request.OptionsJson) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !dara.IsNil(request.RandomIspCity) {
		query["RandomIspCity"] = request.RandomIspCity
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstantSiteMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstantSiteMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the CreateInstantSiteMonitor operation to create a one-time detection task.
//
// Description:
//
// Only Alibaba Cloud accounts that have Network Analysis and Monitoring activated can create one-time detection tasks.
//
// This topic provides an example of how to create a one-time detection task. The example creates a task named `task1` that detects the address `http://www.aliyun.com`. The detection type is `HTTP`, and the number of detection points is `1`.
//
// @param request - CreateInstantSiteMonitorRequest
//
// @return CreateInstantSiteMonitorResponse
func (client *Client) CreateInstantSiteMonitor(request *CreateInstantSiteMonitorRequest) (_result *CreateInstantSiteMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstantSiteMonitorResponse{}
	_body, _err := client.CreateInstantSiteMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a blacklist policy.
//
// Description:
//
// ### Background information
//
// - CloudMonitor blocks alert notifications based on the blacklist policies that take effect. To block alert notifications when the value of a metric that belongs to a cloud service reaches the threshold that you specified, add the metric to a blacklist policy.
//
// - CloudMonitor allows you to create blacklist policies only based on threshold metrics. You cannot create blacklist policies based on system events. For more information about the cloud services and the thresholds of the metrics that are supported by CloudMonitor, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
//
// @param request - CreateMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetricRuleBlackListResponse
func (client *Client) CreateMetricRuleBlackListWithOptions(request *CreateMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *CreateMetricRuleBlackListResponse, _err error) {
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

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EnableEndTime) {
		query["EnableEndTime"] = request.EnableEndTime
	}

	if !dara.IsNil(request.EnableStartTime) {
		query["EnableStartTime"] = request.EnableStartTime
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	if !dara.IsNil(request.ScopeValue) {
		query["ScopeValue"] = request.ScopeValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetricRuleBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a blacklist policy.
//
// Description:
//
// ### Background information
//
// - CloudMonitor blocks alert notifications based on the blacklist policies that take effect. To block alert notifications when the value of a metric that belongs to a cloud service reaches the threshold that you specified, add the metric to a blacklist policy.
//
// - CloudMonitor allows you to create blacklist policies only based on threshold metrics. You cannot create blacklist policies based on system events. For more information about the cloud services and the thresholds of the metrics that are supported by CloudMonitor, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
//
// @param request - CreateMetricRuleBlackListRequest
//
// @return CreateMetricRuleBlackListResponse
func (client *Client) CreateMetricRuleBlackList(request *CreateMetricRuleBlackListRequest) (_result *CreateMetricRuleBlackListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMetricRuleBlackListResponse{}
	_body, _err := client.CreateMetricRuleBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the CreateMetricRuleResources operation to create a resource associated with an alert rule.
//
// @param request - CreateMetricRuleResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetricRuleResourcesResponse
func (client *Client) CreateMetricRuleResourcesWithOptions(request *CreateMetricRuleResourcesRequest, runtime *dara.RuntimeOptions) (_result *CreateMetricRuleResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetricRuleResources"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetricRuleResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateMetricRuleResources operation to create a resource associated with an alert rule.
//
// @param request - CreateMetricRuleResourcesRequest
//
// @return CreateMetricRuleResourcesResponse
func (client *Client) CreateMetricRuleResources(request *CreateMetricRuleResourcesRequest) (_result *CreateMetricRuleResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMetricRuleResourcesResponse{}
	_body, _err := client.CreateMetricRuleResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an alert template.
//
// @param request - CreateMetricRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetricRuleTemplateResponse
func (client *Client) CreateMetricRuleTemplateWithOptions(request *CreateMetricRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateMetricRuleTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertTemplates) {
		query["AlertTemplates"] = request.AlertTemplates
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetricRuleTemplate"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetricRuleTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert template.
//
// @param request - CreateMetricRuleTemplateRequest
//
// @return CreateMetricRuleTemplateResponse
func (client *Client) CreateMetricRuleTemplate(request *CreateMetricRuleTemplateRequest) (_result *CreateMetricRuleTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMetricRuleTemplateResponse{}
	_body, _err := client.CreateMetricRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task to monitor a process.
//
// @param request - CreateMonitorAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorAgentProcessResponse
func (client *Client) CreateMonitorAgentProcessWithOptions(request *CreateMonitorAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorAgentProcessResponse, _err error) {
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

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	if !dara.IsNil(request.ProcessUser) {
		query["ProcessUser"] = request.ProcessUser
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorAgentProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to monitor a process.
//
// @param request - CreateMonitorAgentProcessRequest
//
// @return CreateMonitorAgentProcessResponse
func (client *Client) CreateMonitorAgentProcess(request *CreateMonitorAgentProcessRequest) (_result *CreateMonitorAgentProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMonitorAgentProcessResponse{}
	_body, _err := client.CreateMonitorAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application group.
//
// Description:
//
// In this example, an application group named `ECS_Group` is created.
//
// @param request - CreateMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupResponse
func (client *Client) CreateMonitorGroupWithOptions(request *CreateMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application group.
//
// Description:
//
// In this example, an application group named `ECS_Group` is created.
//
// @param request - CreateMonitorGroupRequest
//
// @return CreateMonitorGroupResponse
func (client *Client) CreateMonitorGroup(request *CreateMonitorGroupRequest) (_result *CreateMonitorGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.CreateMonitorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application group by using a resource group.
//
// Description:
//
// This topic provides an example on how to create an application group by using the resource group `CloudMonitor` and the alert contact group `ECS_Group`. The region ID of the resource group is `cn-hangzhou`.
//
// @param request - CreateMonitorGroupByResourceGroupIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupByResourceGroupIdResponse
func (client *Client) CreateMonitorGroupByResourceGroupIdWithOptions(request *CreateMonitorGroupByResourceGroupIdRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupByResourceGroupIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupList) {
		query["ContactGroupList"] = request.ContactGroupList
	}

	if !dara.IsNil(request.EnableInstallAgent) {
		query["EnableInstallAgent"] = request.EnableInstallAgent
	}

	if !dara.IsNil(request.EnableSubscribeEvent) {
		query["EnableSubscribeEvent"] = request.EnableSubscribeEvent
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroupByResourceGroupId"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupByResourceGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application group by using a resource group.
//
// Description:
//
// This topic provides an example on how to create an application group by using the resource group `CloudMonitor` and the alert contact group `ECS_Group`. The region ID of the resource group is `cn-hangzhou`.
//
// @param request - CreateMonitorGroupByResourceGroupIdRequest
//
// @return CreateMonitorGroupByResourceGroupIdResponse
func (client *Client) CreateMonitorGroupByResourceGroupId(request *CreateMonitorGroupByResourceGroupIdRequest) (_result *CreateMonitorGroupByResourceGroupIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMonitorGroupByResourceGroupIdResponse{}
	_body, _err := client.CreateMonitorGroupByResourceGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds resources to an application group.
//
// Description:
//
// You can add a maximum of 1,000 instances to an application group at a time. You can add a maximum of 3,000 instances of an Alibaba Cloud service to an application group. The total number of instances that you can add to an application group is unlimited.
//
// In this example, an Elastic Compute Service (ECS) instance in the `China (Hangzhou)` region is added to the `3607****` application group. The instance ID is `i-2ze26xj5wwy12****` and the instance name is `test-instance-ecs`.
//
// @param request - CreateMonitorGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupInstancesResponse
func (client *Client) CreateMonitorGroupInstancesWithOptions(request *CreateMonitorGroupInstancesRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroupInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds resources to an application group.
//
// Description:
//
// You can add a maximum of 1,000 instances to an application group at a time. You can add a maximum of 3,000 instances of an Alibaba Cloud service to an application group. The total number of instances that you can add to an application group is unlimited.
//
// In this example, an Elastic Compute Service (ECS) instance in the `China (Hangzhou)` region is added to the `3607****` application group. The instance ID is `i-2ze26xj5wwy12****` and the instance name is `test-instance-ecs`.
//
// @param request - CreateMonitorGroupInstancesRequest
//
// @return CreateMonitorGroupInstancesResponse
func (client *Client) CreateMonitorGroupInstances(request *CreateMonitorGroupInstancesRequest) (_result *CreateMonitorGroupInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMonitorGroupInstancesResponse{}
	_body, _err := client.CreateMonitorGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the CreateMonitorGroupNotifyPolicy operation to create a pause alert notification policy for an application group.
//
// Description:
//
// During the effective period of the policy, no alert notifications are sent for any alerts triggered within the application group.
//
// This topic provides an example on how to create a pause alert notification policy named `PauseNotify` for the application group `7301****`. The application group pauses alert notifications during the period from `1622949300000` to `1623208500000` (UTC+8 `2021-06-06 11:15:00` to `2021-06-09 11:15:00`).
//
// @param request - CreateMonitorGroupNotifyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupNotifyPolicyResponse
func (client *Client) CreateMonitorGroupNotifyPolicyWithOptions(request *CreateMonitorGroupNotifyPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupNotifyPolicyResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroupNotifyPolicy"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateMonitorGroupNotifyPolicy operation to create a pause alert notification policy for an application group.
//
// Description:
//
// During the effective period of the policy, no alert notifications are sent for any alerts triggered within the application group.
//
// This topic provides an example on how to create a pause alert notification policy named `PauseNotify` for the application group `7301****`. The application group pauses alert notifications during the period from `1622949300000` to `1623208500000` (UTC+8 `2021-06-06 11:15:00` to `2021-06-09 11:15:00`).
//
// @param request - CreateMonitorGroupNotifyPolicyRequest
//
// @return CreateMonitorGroupNotifyPolicyResponse
func (client *Client) CreateMonitorGroupNotifyPolicy(request *CreateMonitorGroupNotifyPolicyRequest) (_result *CreateMonitorGroupNotifyPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.CreateMonitorGroupNotifyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task to monitor a process.
//
// @param request - CreateMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitoringAgentProcessResponse
func (client *Client) CreateMonitoringAgentProcessWithOptions(request *CreateMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitoringAgentProcessResponse, _err error) {
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

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	if !dara.IsNil(request.ProcessUser) {
		query["ProcessUser"] = request.ProcessUser
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitoringAgentProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to monitor a process.
//
// @param request - CreateMonitoringAgentProcessRequest
//
// @return CreateMonitoringAgentProcessResponse
func (client *Client) CreateMonitoringAgentProcess(request *CreateMonitoringAgentProcessRequest) (_result *CreateMonitoringAgentProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMonitoringAgentProcessResponse{}
	_body, _err := client.CreateMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a site monitoring task.
//
// Description:
//
// This topic describes how to create a site monitoring task. The example creates a task named `HanZhou_ECS1` to monitor the URL `https://www.aliyun.com` over `HTTPS`.
//
// @param request - CreateSiteMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSiteMonitorResponse
func (client *Client) CreateSiteMonitorWithOptions(request *CreateSiteMonitorRequest, runtime *dara.RuntimeOptions) (_result *CreateSiteMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AgentGroup) {
		query["AgentGroup"] = request.AgentGroup
	}

	if !dara.IsNil(request.AlertIds) {
		query["AlertIds"] = request.AlertIds
	}

	if !dara.IsNil(request.CustomSchedule) {
		query["CustomSchedule"] = request.CustomSchedule
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspCities) {
		query["IspCities"] = request.IspCities
	}

	if !dara.IsNil(request.OptionsJson) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.VpcConfig) {
		query["VpcConfig"] = request.VpcConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSiteMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSiteMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a site monitoring task.
//
// Description:
//
// This topic describes how to create a site monitoring task. The example creates a task named `HanZhou_ECS1` to monitor the URL `https://www.aliyun.com` over `HTTPS`.
//
// @param request - CreateSiteMonitorRequest
//
// @return CreateSiteMonitorResponse
func (client *Client) CreateSiteMonitor(request *CreateSiteMonitorRequest) (_result *CreateSiteMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSiteMonitorResponse{}
	_body, _err := client.CreateSiteMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the Cursor operation to define the scope of monitoring data to be exported, and returns the Cursor value used for the initial call to the BatchExport operation.
//
// Description:
//
// ### Prerequisites
//
// Make sure that Enterprise CloudMonitor is activated. For more information, see [Activate Enterprise CloudMonitor](https://help.aliyun.com/document_detail/250773.html).
//
// ### Background information
//
// First, call this operation to obtain the initial Cursor. Then, call the [BatchExport](https://help.aliyun.com/document_detail/2329847.html) operation to export monitoring data.
//
// ### Usage notes
//
// This topic provides an example to describe how to define the scope to export data of the `cpu_idle` metric of the `acs_ecs_dashboard` cloud service every 60 seconds in the time range from `1641627000000` (2022-01-08 15:30:00) to `1641645000000` (2022-01-08 20:30:00). The returned result shows the `Cursor` information.
//
// @param tmpReq - CursorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CursorResponse
func (client *Client) CursorWithOptions(tmpReq *CursorRequest, runtime *dara.RuntimeOptions) (_result *CursorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CursorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Matchers) {
		request.MatchersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Matchers, dara.String("Matchers"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MatchersShrink) {
		body["Matchers"] = request.MatchersShrink
	}

	if !dara.IsNil(request.Metric) {
		body["Metric"] = request.Metric
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Period) {
		body["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Cursor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CursorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the Cursor operation to define the scope of monitoring data to be exported, and returns the Cursor value used for the initial call to the BatchExport operation.
//
// Description:
//
// ### Prerequisites
//
// Make sure that Enterprise CloudMonitor is activated. For more information, see [Activate Enterprise CloudMonitor](https://help.aliyun.com/document_detail/250773.html).
//
// ### Background information
//
// First, call this operation to obtain the initial Cursor. Then, call the [BatchExport](https://help.aliyun.com/document_detail/2329847.html) operation to export monitoring data.
//
// ### Usage notes
//
// This topic provides an example to describe how to define the scope to export data of the `cpu_idle` metric of the `acs_ecs_dashboard` cloud service every 60 seconds in the time range from `1641627000000` (2022-01-08 15:30:00) to `1641645000000` (2022-01-08 20:30:00). The returned result shows the `Cursor` information.
//
// @param request - CursorRequest
//
// @return CursorResponse
func (client *Client) Cursor(request *CursorRequest) (_result *CursorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CursorResponse{}
	_body, _err := client.CursorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alert contact.
//
// @param request - DeleteContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactResponse
func (client *Client) DeleteContactWithOptions(request *DeleteContactRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContact"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert contact.
//
// @param request - DeleteContactRequest
//
// @return DeleteContactResponse
func (client *Client) DeleteContact(request *DeleteContactRequest) (_result *DeleteContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContactResponse{}
	_body, _err := client.DeleteContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DeleteContactGroup operation to delete an alert contact group.
//
// @param request - DeleteContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContactGroupResponse
func (client *Client) DeleteContactGroupWithOptions(request *DeleteContactGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteContactGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteContactGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteContactGroup operation to delete an alert contact group.
//
// @param request - DeleteContactGroupRequest
//
// @return DeleteContactGroupResponse
func (client *Client) DeleteContactGroup(request *DeleteContactGroupRequest) (_result *DeleteContactGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteContactGroupResponse{}
	_body, _err := client.DeleteContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the reported monitoring data of a metric.
//
// @param request - DeleteCustomMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomMetricResponse
func (client *Client) DeleteCustomMetricWithOptions(request *DeleteCustomMetricRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Md5) {
		query["Md5"] = request.Md5
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.UUID) {
		query["UUID"] = request.UUID
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomMetric"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the reported monitoring data of a metric.
//
// @param request - DeleteCustomMetricRequest
//
// @return DeleteCustomMetricResponse
func (client *Client) DeleteCustomMetric(request *DeleteCustomMetricRequest) (_result *DeleteCustomMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomMetricResponse{}
	_body, _err := client.DeleteCustomMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a tag rule.
//
// @param request - DeleteDynamicTagGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDynamicTagGroupResponse
func (client *Client) DeleteDynamicTagGroupWithOptions(request *DeleteDynamicTagGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDynamicTagGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicTagRuleId) {
		query["DynamicTagRuleId"] = request.DynamicTagRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDynamicTagGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDynamicTagGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a tag rule.
//
// @param request - DeleteDynamicTagGroupRequest
//
// @return DeleteDynamicTagGroupResponse
func (client *Client) DeleteDynamicTagGroup(request *DeleteDynamicTagGroupRequest) (_result *DeleteDynamicTagGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDynamicTagGroupResponse{}
	_body, _err := client.DeleteDynamicTagGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the push channels of an event-triggered alert rule.
//
// @param request - DeleteEventRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventRuleTargetsResponse
func (client *Client) DeleteEventRuleTargetsWithOptions(request *DeleteEventRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventRuleTargetsResponse, _err error) {
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

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventRuleTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the push channels of an event-triggered alert rule.
//
// @param request - DeleteEventRuleTargetsRequest
//
// @return DeleteEventRuleTargetsResponse
func (client *Client) DeleteEventRuleTargets(request *DeleteEventRuleTargetsRequest) (_result *DeleteEventRuleTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEventRuleTargetsResponse{}
	_body, _err := client.DeleteEventRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes one or more event-triggered alert rules.
//
// @param request - DeleteEventRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventRulesResponse
func (client *Client) DeleteEventRulesWithOptions(request *DeleteEventRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteEventRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleNames) {
		query["RuleNames"] = request.RuleNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEventRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEventRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more event-triggered alert rules.
//
// @param request - DeleteEventRulesRequest
//
// @return DeleteEventRulesResponse
func (client *Client) DeleteEventRules(request *DeleteEventRulesRequest) (_result *DeleteEventRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEventRulesResponse{}
	_body, _err := client.DeleteEventRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DeleteExporterOutput operation to delete a monitoring data export configuration.
//
// @param request - DeleteExporterOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExporterOutputResponse
func (client *Client) DeleteExporterOutputWithOptions(request *DeleteExporterOutputRequest, runtime *dara.RuntimeOptions) (_result *DeleteExporterOutputResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestName) {
		query["DestName"] = request.DestName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExporterOutput"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExporterOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteExporterOutput operation to delete a monitoring data export configuration.
//
// @param request - DeleteExporterOutputRequest
//
// @return DeleteExporterOutputResponse
func (client *Client) DeleteExporterOutput(request *DeleteExporterOutputRequest) (_result *DeleteExporterOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExporterOutputResponse{}
	_body, _err := client.DeleteExporterOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the DeleteExporterRule operation to delete export rules.
//
// @param request - DeleteExporterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExporterRuleResponse
func (client *Client) DeleteExporterRuleWithOptions(request *DeleteExporterRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteExporterRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExporterRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExporterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the DeleteExporterRule operation to delete export rules.
//
// @param request - DeleteExporterRuleRequest
//
// @return DeleteExporterRuleResponse
func (client *Client) DeleteExporterRule(request *DeleteExporterRuleRequest) (_result *DeleteExporterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExporterRuleResponse{}
	_body, _err := client.DeleteExporterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a group process monitoring task by calling the DeleteGroupMonitoringAgentProcess operation.
//
// @param request - DeleteGroupMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupMonitoringAgentProcessResponse
func (client *Client) DeleteGroupMonitoringAgentProcessWithOptions(request *DeleteGroupMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *DeleteGroupMonitoringAgentProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGroupMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a group process monitoring task by calling the DeleteGroupMonitoringAgentProcess operation.
//
// @param request - DeleteGroupMonitoringAgentProcessRequest
//
// @return DeleteGroupMonitoringAgentProcessResponse
func (client *Client) DeleteGroupMonitoringAgentProcess(request *DeleteGroupMonitoringAgentProcessRequest) (_result *DeleteGroupMonitoringAgentProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteGroupMonitoringAgentProcessResponse{}
	_body, _err := client.DeleteGroupMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes one or more availability monitoring jobs.
//
// @param request - DeleteHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHostAvailabilityResponse
func (client *Client) DeleteHostAvailabilityWithOptions(request *DeleteHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *DeleteHostAvailabilityResponse, _err error) {
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
		Action:      dara.String("DeleteHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHostAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more availability monitoring jobs.
//
// @param request - DeleteHostAvailabilityRequest
//
// @return DeleteHostAvailabilityResponse
func (client *Client) DeleteHostAvailability(request *DeleteHostAvailabilityRequest) (_result *DeleteHostAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHostAvailabilityResponse{}
	_body, _err := client.DeleteHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a namespace.
//
// Description:
//
// > If a metric import task is created for metrics in a namespace, you cannot delete the namespace unless you delete the task first.
//
// This topic provides an example on how to delete a namespace named `aliyun`. The response shows that the namespace is deleted.
//
// @param request - DeleteHybridMonitorNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHybridMonitorNamespaceResponse
func (client *Client) DeleteHybridMonitorNamespaceWithOptions(request *DeleteHybridMonitorNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteHybridMonitorNamespaceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHybridMonitorNamespace"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHybridMonitorNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a namespace.
//
// Description:
//
// > If a metric import task is created for metrics in a namespace, you cannot delete the namespace unless you delete the task first.
//
// This topic provides an example on how to delete a namespace named `aliyun`. The response shows that the namespace is deleted.
//
// @param request - DeleteHybridMonitorNamespaceRequest
//
// @return DeleteHybridMonitorNamespaceResponse
func (client *Client) DeleteHybridMonitorNamespace(request *DeleteHybridMonitorNamespaceRequest) (_result *DeleteHybridMonitorNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHybridMonitorNamespaceResponse{}
	_body, _err := client.DeleteHybridMonitorNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Logstore group.
//
// Description:
//
// This topic provides an example on how to delete a Logstore group named `Logstore_test`. The response shows that the Logstore group is deleted.
//
// @param request - DeleteHybridMonitorSLSGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHybridMonitorSLSGroupResponse
func (client *Client) DeleteHybridMonitorSLSGroupWithOptions(request *DeleteHybridMonitorSLSGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteHybridMonitorSLSGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SLSGroupName) {
		query["SLSGroupName"] = request.SLSGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHybridMonitorSLSGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHybridMonitorSLSGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Logstore group.
//
// Description:
//
// This topic provides an example on how to delete a Logstore group named `Logstore_test`. The response shows that the Logstore group is deleted.
//
// @param request - DeleteHybridMonitorSLSGroupRequest
//
// @return DeleteHybridMonitorSLSGroupResponse
func (client *Client) DeleteHybridMonitorSLSGroup(request *DeleteHybridMonitorSLSGroupRequest) (_result *DeleteHybridMonitorSLSGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHybridMonitorSLSGroupResponse{}
	_body, _err := client.DeleteHybridMonitorSLSGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a metric import task for Alibaba Cloud services or a metric for the logs that are imported from Log Service.
//
// Description:
//
// This topic provides an example on how to delete a metric import task whose ID is `36****`. The returned result indicates that the metric import task is deleted.
//
// @param request - DeleteHybridMonitorTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHybridMonitorTaskResponse
func (client *Client) DeleteHybridMonitorTaskWithOptions(request *DeleteHybridMonitorTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteHybridMonitorTaskResponse, _err error) {
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

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHybridMonitorTask"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHybridMonitorTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a metric import task for Alibaba Cloud services or a metric for the logs that are imported from Log Service.
//
// Description:
//
// This topic provides an example on how to delete a metric import task whose ID is `36****`. The returned result indicates that the metric import task is deleted.
//
// @param request - DeleteHybridMonitorTaskRequest
//
// @return DeleteHybridMonitorTaskResponse
func (client *Client) DeleteHybridMonitorTask(request *DeleteHybridMonitorTaskRequest) (_result *DeleteHybridMonitorTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHybridMonitorTaskResponse{}
	_body, _err := client.DeleteHybridMonitorTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a log monitoring metric.
//
// @param request - DeleteLogMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogMonitorResponse
func (client *Client) DeleteLogMonitorWithOptions(request *DeleteLogMonitorRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogId) {
		query["LogId"] = request.LogId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a log monitoring metric.
//
// @param request - DeleteLogMonitorRequest
//
// @return DeleteLogMonitorResponse
func (client *Client) DeleteLogMonitor(request *DeleteLogMonitorRequest) (_result *DeleteLogMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLogMonitorResponse{}
	_body, _err := client.DeleteLogMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DeleteMetricRuleBlackList operation to delete alert blacklist policies.
//
// @param request - DeleteMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRuleBlackListResponse
func (client *Client) DeleteMetricRuleBlackListWithOptions(request *DeleteMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRuleBlackListResponse, _err error) {
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
		Action:      dara.String("DeleteMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRuleBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteMetricRuleBlackList operation to delete alert blacklist policies.
//
// @param request - DeleteMetricRuleBlackListRequest
//
// @return DeleteMetricRuleBlackListResponse
func (client *Client) DeleteMetricRuleBlackList(request *DeleteMetricRuleBlackListRequest) (_result *DeleteMetricRuleBlackListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMetricRuleBlackListResponse{}
	_body, _err := client.DeleteMetricRuleBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates resources from an alert rule.
//
// @param request - DeleteMetricRuleResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRuleResourcesResponse
func (client *Client) DeleteMetricRuleResourcesWithOptions(request *DeleteMetricRuleResourcesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRuleResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetricRuleResources"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRuleResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates resources from an alert rule.
//
// @param request - DeleteMetricRuleResourcesRequest
//
// @return DeleteMetricRuleResourcesResponse
func (client *Client) DeleteMetricRuleResources(request *DeleteMetricRuleResourcesRequest) (_result *DeleteMetricRuleResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMetricRuleResourcesResponse{}
	_body, _err := client.DeleteMetricRuleResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the push channels of an alert rule.
//
// @param request - DeleteMetricRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRuleTargetsResponse
func (client *Client) DeleteMetricRuleTargetsWithOptions(request *DeleteMetricRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRuleTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TargetIds) {
		query["TargetIds"] = request.TargetIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetricRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRuleTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the push channels of an alert rule.
//
// @param request - DeleteMetricRuleTargetsRequest
//
// @return DeleteMetricRuleTargetsResponse
func (client *Client) DeleteMetricRuleTargets(request *DeleteMetricRuleTargetsRequest) (_result *DeleteMetricRuleTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMetricRuleTargetsResponse{}
	_body, _err := client.DeleteMetricRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alert rule template.
//
// @param request - DeleteMetricRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRuleTemplateResponse
func (client *Client) DeleteMetricRuleTemplateWithOptions(request *DeleteMetricRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRuleTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetricRuleTemplate"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRuleTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert rule template.
//
// @param request - DeleteMetricRuleTemplateRequest
//
// @return DeleteMetricRuleTemplateResponse
func (client *Client) DeleteMetricRuleTemplate(request *DeleteMetricRuleTemplateRequest) (_result *DeleteMetricRuleTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMetricRuleTemplateResponse{}
	_body, _err := client.DeleteMetricRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The DeleteMetricRules operation deletes one or more alert rules.
//
// @param request - DeleteMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetricRulesResponse
func (client *Client) DeleteMetricRulesWithOptions(request *DeleteMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetricRulesResponse, _err error) {
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
		Action:      dara.String("DeleteMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetricRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The DeleteMetricRules operation deletes one or more alert rules.
//
// @param request - DeleteMetricRulesRequest
//
// @return DeleteMetricRulesResponse
func (client *Client) DeleteMetricRules(request *DeleteMetricRulesRequest) (_result *DeleteMetricRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMetricRulesResponse{}
	_body, _err := client.DeleteMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application group.
//
// @param request - DeleteMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupResponse
func (client *Client) DeleteMonitorGroupWithOptions(request *DeleteMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application group.
//
// @param request - DeleteMonitorGroupRequest
//
// @return DeleteMonitorGroupResponse
func (client *Client) DeleteMonitorGroup(request *DeleteMonitorGroupRequest) (_result *DeleteMonitorGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMonitorGroupResponse{}
	_body, _err := client.DeleteMonitorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a rule that is used to dynamically add the instances of a service to an application group.
//
// @param request - DeleteMonitorGroupDynamicRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupDynamicRuleResponse
func (client *Client) DeleteMonitorGroupDynamicRuleWithOptions(request *DeleteMonitorGroupDynamicRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupDynamicRuleResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroupDynamicRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupDynamicRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a rule that is used to dynamically add the instances of a service to an application group.
//
// @param request - DeleteMonitorGroupDynamicRuleRequest
//
// @return DeleteMonitorGroupDynamicRuleResponse
func (client *Client) DeleteMonitorGroupDynamicRule(request *DeleteMonitorGroupDynamicRuleRequest) (_result *DeleteMonitorGroupDynamicRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMonitorGroupDynamicRuleResponse{}
	_body, _err := client.DeleteMonitorGroupDynamicRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes instances from an application group.
//
// @param request - DeleteMonitorGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupInstancesResponse
func (client *Client) DeleteMonitorGroupInstancesWithOptions(request *DeleteMonitorGroupInstancesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupInstancesResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroupInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes instances from an application group.
//
// @param request - DeleteMonitorGroupInstancesRequest
//
// @return DeleteMonitorGroupInstancesResponse
func (client *Client) DeleteMonitorGroupInstances(request *DeleteMonitorGroupInstancesRequest) (_result *DeleteMonitorGroupInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMonitorGroupInstancesResponse{}
	_body, _err := client.DeleteMonitorGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a policy that is used to pause alert notifications for an application group.
//
// @param request - DeleteMonitorGroupNotifyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupNotifyPolicyResponse
func (client *Client) DeleteMonitorGroupNotifyPolicyWithOptions(request *DeleteMonitorGroupNotifyPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupNotifyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroupNotifyPolicy"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a policy that is used to pause alert notifications for an application group.
//
// @param request - DeleteMonitorGroupNotifyPolicyRequest
//
// @return DeleteMonitorGroupNotifyPolicyResponse
func (client *Client) DeleteMonitorGroupNotifyPolicy(request *DeleteMonitorGroupNotifyPolicyRequest) (_result *DeleteMonitorGroupNotifyPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMonitorGroupNotifyPolicyResponse{}
	_body, _err := client.DeleteMonitorGroupNotifyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DeleteMonitoringAgentProcess operation to delete the specified process monitoring from a specified host.
//
// @param request - DeleteMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitoringAgentProcessResponse
func (client *Client) DeleteMonitoringAgentProcessWithOptions(request *DeleteMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitoringAgentProcessResponse, _err error) {
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

	if !dara.IsNil(request.ProcessId) {
		query["ProcessId"] = request.ProcessId
	}

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitoringAgentProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteMonitoringAgentProcess operation to delete the specified process monitoring from a specified host.
//
// @param request - DeleteMonitoringAgentProcessRequest
//
// @return DeleteMonitoringAgentProcessResponse
func (client *Client) DeleteMonitoringAgentProcess(request *DeleteMonitoringAgentProcessRequest) (_result *DeleteMonitoringAgentProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMonitoringAgentProcessResponse{}
	_body, _err := client.DeleteMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes site monitoring tasks by calling the DeleteSiteMonitors operation.
//
// @param request - DeleteSiteMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSiteMonitorsResponse
func (client *Client) DeleteSiteMonitorsWithOptions(request *DeleteSiteMonitorsRequest, runtime *dara.RuntimeOptions) (_result *DeleteSiteMonitorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsDeleteAlarms) {
		query["IsDeleteAlarms"] = request.IsDeleteAlarms
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSiteMonitors"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSiteMonitorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes site monitoring tasks by calling the DeleteSiteMonitors operation.
//
// @param request - DeleteSiteMonitorsRequest
//
// @return DeleteSiteMonitorsResponse
func (client *Client) DeleteSiteMonitors(request *DeleteSiteMonitorsRequest) (_result *DeleteSiteMonitorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSiteMonitorsResponse{}
	_body, _err := client.DeleteSiteMonitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries initiative alert rules.
//
// @param request - DescribeActiveMetricRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveMetricRuleListResponse
func (client *Client) DescribeActiveMetricRuleListWithOptions(request *DescribeActiveMetricRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveMetricRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveMetricRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveMetricRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries initiative alert rules.
//
// @param request - DescribeActiveMetricRuleListRequest
//
// @return DescribeActiveMetricRuleListResponse
func (client *Client) DescribeActiveMetricRuleList(request *DescribeActiveMetricRuleListRequest) (_result *DescribeActiveMetricRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeActiveMetricRuleListResponse{}
	_body, _err := client.DescribeActiveMetricRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeAlertHistoryList is deprecated, please use Cms::2019-01-01::DescribeAlertLogList instead.
//
// Summary:
//
// Queries historical alerts.
//
// Description:
//
// This API operation is no longer maintained. We recommend that you call the [DescribeAlertLogList](https://help.aliyun.com/document_detail/201087.html) operation.
//
// @param request - DescribeAlertHistoryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertHistoryListResponse
func (client *Client) DescribeAlertHistoryListWithOptions(request *DescribeAlertHistoryListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertHistoryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ascending) {
		query["Ascending"] = request.Ascending
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertHistoryList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertHistoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeAlertHistoryList is deprecated, please use Cms::2019-01-01::DescribeAlertLogList instead.
//
// Summary:
//
// Queries historical alerts.
//
// Description:
//
// This API operation is no longer maintained. We recommend that you call the [DescribeAlertLogList](https://help.aliyun.com/document_detail/201087.html) operation.
//
// @param request - DescribeAlertHistoryListRequest
//
// @return DescribeAlertHistoryListResponse
// Deprecated
func (client *Client) DescribeAlertHistoryList(request *DescribeAlertHistoryListRequest) (_result *DescribeAlertHistoryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertHistoryListResponse{}
	_body, _err := client.DescribeAlertHistoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics of alert logs.
//
// Description:
//
// Queries the statistics of alert logs.
//
// This topic provides an example on how to query the statistics of alert logs for Elastic Compute Service (ECS) based on the `product` dimension.
//
// @param request - DescribeAlertLogCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertLogCountResponse
func (client *Client) DescribeAlertLogCountWithOptions(request *DescribeAlertLogCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertLogCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroup) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LastMin) {
		query["LastMin"] = request.LastMin
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SendStatus) {
		query["SendStatus"] = request.SendStatus
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertLogCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertLogCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of alert logs.
//
// Description:
//
// Queries the statistics of alert logs.
//
// This topic provides an example on how to query the statistics of alert logs for Elastic Compute Service (ECS) based on the `product` dimension.
//
// @param request - DescribeAlertLogCountRequest
//
// @return DescribeAlertLogCountResponse
func (client *Client) DescribeAlertLogCount(request *DescribeAlertLogCountRequest) (_result *DescribeAlertLogCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertLogCountResponse{}
	_body, _err := client.DescribeAlertLogCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of alert logs that are generated during each interval within a period of time.
//
// Description:
//
// This topic provides an example on how to query the number of alert logs for Elastic Compute Service (ECS) based on the `product` dimension.
//
// @param request - DescribeAlertLogHistogramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertLogHistogramResponse
func (client *Client) DescribeAlertLogHistogramWithOptions(request *DescribeAlertLogHistogramRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertLogHistogramResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroup) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LastMin) {
		query["LastMin"] = request.LastMin
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SendStatus) {
		query["SendStatus"] = request.SendStatus
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertLogHistogram"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertLogHistogramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of alert logs that are generated during each interval within a period of time.
//
// Description:
//
// This topic provides an example on how to query the number of alert logs for Elastic Compute Service (ECS) based on the `product` dimension.
//
// @param request - DescribeAlertLogHistogramRequest
//
// @return DescribeAlertLogHistogramResponse
func (client *Client) DescribeAlertLogHistogram(request *DescribeAlertLogHistogramRequest) (_result *DescribeAlertLogHistogramResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertLogHistogramResponse{}
	_body, _err := client.DescribeAlertLogHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alert history by calling the DescribeAlertLogList operation.
//
// Description:
//
// This operation can query the alert history only within the last year.
//
// This topic provides an example to show how to query the alert history of Elastic Compute Service (ECS) from the cloud service `product` dimension.
//
// @param request - DescribeAlertLogListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertLogListResponse
func (client *Client) DescribeAlertLogListWithOptions(request *DescribeAlertLogListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertLogListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroup) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LastMin) {
		query["LastMin"] = request.LastMin
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SendStatus) {
		query["SendStatus"] = request.SendStatus
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAlertLogList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertLogListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert history by calling the DescribeAlertLogList operation.
//
// Description:
//
// This operation can query the alert history only within the last year.
//
// This topic provides an example to show how to query the alert history of Elastic Compute Service (ECS) from the cloud service `product` dimension.
//
// @param request - DescribeAlertLogListRequest
//
// @return DescribeAlertLogListResponse
func (client *Client) DescribeAlertLogList(request *DescribeAlertLogListRequest) (_result *DescribeAlertLogListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertLogListResponse{}
	_body, _err := client.DescribeAlertLogListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources for which active alerts are triggered based on an alert rule.
//
// @param request - DescribeAlertingMetricRuleResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAlertingMetricRuleResourcesResponse
func (client *Client) DescribeAlertingMetricRuleResourcesWithOptions(request *DescribeAlertingMetricRuleResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAlertingMetricRuleResourcesResponse, _err error) {
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
		Action:      dara.String("DescribeAlertingMetricRuleResources"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAlertingMetricRuleResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources for which active alerts are triggered based on an alert rule.
//
// @param request - DescribeAlertingMetricRuleResourcesRequest
//
// @return DescribeAlertingMetricRuleResourcesResponse
func (client *Client) DescribeAlertingMetricRuleResources(request *DescribeAlertingMetricRuleResourcesRequest) (_result *DescribeAlertingMetricRuleResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAlertingMetricRuleResourcesResponse{}
	_body, _err := client.DescribeAlertingMetricRuleResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DescribeContactGroupList interface to query the list of alarm contact groups.
//
// @param request - DescribeContactGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContactGroupListResponse
func (client *Client) DescribeContactGroupListWithOptions(request *DescribeContactGroupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeContactGroupListResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeContactGroupList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContactGroupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeContactGroupList interface to query the list of alarm contact groups.
//
// @param request - DescribeContactGroupListRequest
//
// @return DescribeContactGroupListResponse
func (client *Client) DescribeContactGroupList(request *DescribeContactGroupListRequest) (_result *DescribeContactGroupListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeContactGroupListResponse{}
	_body, _err := client.DescribeContactGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DescribeContactList operation to query a list of alert contacts.
//
// @param request - DescribeContactListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContactListResponse
func (client *Client) DescribeContactListWithOptions(request *DescribeContactListRequest, runtime *dara.RuntimeOptions) (_result *DescribeContactListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChanelType) {
		query["ChanelType"] = request.ChanelType
	}

	if !dara.IsNil(request.ChanelValue) {
		query["ChanelValue"] = request.ChanelValue
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
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
		Action:      dara.String("DescribeContactList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContactListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeContactList operation to query a list of alert contacts.
//
// @param request - DescribeContactListRequest
//
// @return DescribeContactListResponse
func (client *Client) DescribeContactList(request *DescribeContactListRequest) (_result *DescribeContactListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeContactListResponse{}
	_body, _err := client.DescribeContactListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the alert contacts in an alert contact group.
//
// @param request - DescribeContactListByContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContactListByContactGroupResponse
func (client *Client) DescribeContactListByContactGroupWithOptions(request *DescribeContactListByContactGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeContactListByContactGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeContactListByContactGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContactListByContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the alert contacts in an alert contact group.
//
// @param request - DescribeContactListByContactGroupRequest
//
// @return DescribeContactListByContactGroupResponse
func (client *Client) DescribeContactListByContactGroup(request *DescribeContactListByContactGroupRequest) (_result *DescribeContactListByContactGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeContactListByContactGroupResponse{}
	_body, _err := client.DescribeContactListByContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a custom event.
//
// @param request - DescribeCustomEventAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomEventAttributeResponse
func (client *Client) DescribeCustomEventAttributeWithOptions(request *DescribeCustomEventAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomEventAttributeResponse, _err error) {
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

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
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

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomEventAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomEventAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a custom event.
//
// @param request - DescribeCustomEventAttributeRequest
//
// @return DescribeCustomEventAttributeResponse
func (client *Client) DescribeCustomEventAttribute(request *DescribeCustomEventAttributeRequest) (_result *DescribeCustomEventAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomEventAttributeResponse{}
	_body, _err := client.DescribeCustomEventAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of times that a custom event occurred within a period of time.
//
// Description:
//
// > This operation queries the number of times that a custom event occurred for each service.
//
// @param request - DescribeCustomEventCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomEventCountResponse
func (client *Client) DescribeCustomEventCountWithOptions(request *DescribeCustomEventCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomEventCountResponse, _err error) {
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

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomEventCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomEventCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of times that a custom event occurred within a period of time.
//
// Description:
//
// > This operation queries the number of times that a custom event occurred for each service.
//
// @param request - DescribeCustomEventCountRequest
//
// @return DescribeCustomEventCountResponse
func (client *Client) DescribeCustomEventCount(request *DescribeCustomEventCountRequest) (_result *DescribeCustomEventCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomEventCountResponse{}
	_body, _err := client.DescribeCustomEventCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of times that a custom event occurred during each interval within a period of time.
//
// @param request - DescribeCustomEventHistogramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomEventHistogramResponse
func (client *Client) DescribeCustomEventHistogramWithOptions(request *DescribeCustomEventHistogramRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomEventHistogramResponse, _err error) {
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

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomEventHistogram"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomEventHistogramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of times that a custom event occurred during each interval within a period of time.
//
// @param request - DescribeCustomEventHistogramRequest
//
// @return DescribeCustomEventHistogramResponse
func (client *Client) DescribeCustomEventHistogram(request *DescribeCustomEventHistogramRequest) (_result *DescribeCustomEventHistogramResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomEventHistogramResponse{}
	_body, _err := client.DescribeCustomEventHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the reported monitoring data.
//
// Description:
//
// > You can call the DescribeMetricList operation to query the metrics of cloud services. For more information, see [DescribeMetricList](https://help.aliyun.com/document_detail/51936.html).
//
// @param request - DescribeCustomMetricListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomMetricListResponse
func (client *Client) DescribeCustomMetricListWithOptions(request *DescribeCustomMetricListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomMetricListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Md5) {
		query["Md5"] = request.Md5
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
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
		Action:      dara.String("DescribeCustomMetricList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomMetricListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the reported monitoring data.
//
// Description:
//
// > You can call the DescribeMetricList operation to query the metrics of cloud services. For more information, see [DescribeMetricList](https://help.aliyun.com/document_detail/51936.html).
//
// @param request - DescribeCustomMetricListRequest
//
// @return DescribeCustomMetricListResponse
func (client *Client) DescribeCustomMetricList(request *DescribeCustomMetricListRequest) (_result *DescribeCustomMetricListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomMetricListResponse{}
	_body, _err := client.DescribeCustomMetricListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DescribeDynamicTagRuleList operation to query the rules for dynamic tags.
//
// Description:
//
// This topic provides an example of how to query the rules for the tag key `tagkey1`. The response shows that two rules are returned. The rule IDs are `1536df65-a719-429d-8813-73cc40d7****` and `56e8cebb-b3d7-4a91-9880-78a8c84f****`.
//
// @param request - DescribeDynamicTagRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDynamicTagRuleListResponse
func (client *Client) DescribeDynamicTagRuleListWithOptions(request *DescribeDynamicTagRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDynamicTagRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicTagRuleId) {
		query["DynamicTagRuleId"] = request.DynamicTagRuleId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	if !dara.IsNil(request.TagRegionId) {
		query["TagRegionId"] = request.TagRegionId
	}

	if !dara.IsNil(request.TagValue) {
		query["TagValue"] = request.TagValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDynamicTagRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDynamicTagRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeDynamicTagRuleList operation to query the rules for dynamic tags.
//
// Description:
//
// This topic provides an example of how to query the rules for the tag key `tagkey1`. The response shows that two rules are returned. The rule IDs are `1536df65-a719-429d-8813-73cc40d7****` and `56e8cebb-b3d7-4a91-9880-78a8c84f****`.
//
// @param request - DescribeDynamicTagRuleListRequest
//
// @return DescribeDynamicTagRuleListResponse
func (client *Client) DescribeDynamicTagRuleList(request *DescribeDynamicTagRuleListRequest) (_result *DescribeDynamicTagRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDynamicTagRuleListResponse{}
	_body, _err := client.DescribeDynamicTagRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified event-triggered alert rule by calling the DescribeEventRuleAttribute operation.
//
// Description:
//
// This topic provides an example on how to query the details of the event-triggered alert rule `testRule`.
//
// @param request - DescribeEventRuleAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventRuleAttributeResponse
func (client *Client) DescribeEventRuleAttributeWithOptions(request *DescribeEventRuleAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventRuleAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventRuleAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventRuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified event-triggered alert rule by calling the DescribeEventRuleAttribute operation.
//
// Description:
//
// This topic provides an example on how to query the details of the event-triggered alert rule `testRule`.
//
// @param request - DescribeEventRuleAttributeRequest
//
// @return DescribeEventRuleAttributeResponse
func (client *Client) DescribeEventRuleAttribute(request *DescribeEventRuleAttributeRequest) (_result *DescribeEventRuleAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventRuleAttributeResponse{}
	_body, _err := client.DescribeEventRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries event-triggered alert rules.
//
// @param request - DescribeEventRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventRuleListResponse
func (client *Client) DescribeEventRuleListWithOptions(request *DescribeEventRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	if !dara.IsNil(request.NamePrefix) {
		query["NamePrefix"] = request.NamePrefix
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
		Action:      dara.String("DescribeEventRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries event-triggered alert rules.
//
// @param request - DescribeEventRuleListRequest
//
// @return DescribeEventRuleListResponse
func (client *Client) DescribeEventRuleList(request *DescribeEventRuleListRequest) (_result *DescribeEventRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventRuleListResponse{}
	_body, _err := client.DescribeEventRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of event-triggered alert rules.
//
// Description:
//
// This topic provides an example to query the details of the event-triggered alert rule `testRule`.
//
// @param request - DescribeEventRuleTargetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventRuleTargetListResponse
func (client *Client) DescribeEventRuleTargetListWithOptions(request *DescribeEventRuleTargetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventRuleTargetListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEventRuleTargetList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventRuleTargetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of event-triggered alert rules.
//
// Description:
//
// This topic provides an example to query the details of the event-triggered alert rule `testRule`.
//
// @param request - DescribeEventRuleTargetListRequest
//
// @return DescribeEventRuleTargetListResponse
func (client *Client) DescribeEventRuleTargetList(request *DescribeEventRuleTargetListRequest) (_result *DescribeEventRuleTargetListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventRuleTargetListResponse{}
	_body, _err := client.DescribeEventRuleTargetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of monitoring data exports by calling the DescribeExporterOutputList operation.
//
// @param request - DescribeExporterOutputListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExporterOutputListResponse
func (client *Client) DescribeExporterOutputListWithOptions(request *DescribeExporterOutputListRequest, runtime *dara.RuntimeOptions) (_result *DescribeExporterOutputListResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExporterOutputList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExporterOutputListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of monitoring data exports by calling the DescribeExporterOutputList operation.
//
// @param request - DescribeExporterOutputListRequest
//
// @return DescribeExporterOutputListResponse
func (client *Client) DescribeExporterOutputList(request *DescribeExporterOutputListRequest) (_result *DescribeExporterOutputListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExporterOutputListResponse{}
	_body, _err := client.DescribeExporterOutputListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of data export rules by calling the DescribeExporterRuleList operation.
//
// @param request - DescribeExporterRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExporterRuleListResponse
func (client *Client) DescribeExporterRuleListWithOptions(request *DescribeExporterRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeExporterRuleListResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExporterRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExporterRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data export rules by calling the DescribeExporterRuleList operation.
//
// @param request - DescribeExporterRuleListRequest
//
// @return DescribeExporterRuleListResponse
func (client *Client) DescribeExporterRuleList(request *DescribeExporterRuleListRequest) (_result *DescribeExporterRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExporterRuleListResponse{}
	_body, _err := client.DescribeExporterRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the process monitoring tasks for an application group.
//
// Description:
//
// You can create a process monitoring task to monitor all or the specified Elastic Compute Service (ECS) instances in an application group and configure alert rules for the process monitoring task.
//
// @param request - DescribeGroupMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupMonitoringAgentProcessResponse
func (client *Client) DescribeGroupMonitoringAgentProcessWithOptions(request *DescribeGroupMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupMonitoringAgentProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProcessName) {
		query["ProcessName"] = request.ProcessName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the process monitoring tasks for an application group.
//
// Description:
//
// You can create a process monitoring task to monitor all or the specified Elastic Compute Service (ECS) instances in an application group and configure alert rules for the process monitoring task.
//
// @param request - DescribeGroupMonitoringAgentProcessRequest
//
// @return DescribeGroupMonitoringAgentProcessResponse
func (client *Client) DescribeGroupMonitoringAgentProcess(request *DescribeGroupMonitoringAgentProcessRequest) (_result *DescribeGroupMonitoringAgentProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupMonitoringAgentProcessResponse{}
	_body, _err := client.DescribeGroupMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries availability monitoring tasks.
//
// Description:
//
// This topic provides an example to show how to query all the availability monitoring tasks of your Alibaba Cloud account. The sample responses indicate that the account has one availability monitoring task named `ecs_instance`.
//
// @param request - DescribeHostAvailabilityListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHostAvailabilityListResponse
func (client *Client) DescribeHostAvailabilityListWithOptions(request *DescribeHostAvailabilityListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHostAvailabilityListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHostAvailabilityList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHostAvailabilityListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries availability monitoring tasks.
//
// Description:
//
// This topic provides an example to show how to query all the availability monitoring tasks of your Alibaba Cloud account. The sample responses indicate that the account has one availability monitoring task named `ecs_instance`.
//
// @param request - DescribeHostAvailabilityListRequest
//
// @return DescribeHostAvailabilityListResponse
func (client *Client) DescribeHostAvailabilityList(request *DescribeHostAvailabilityListRequest) (_result *DescribeHostAvailabilityListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHostAvailabilityListResponse{}
	_body, _err := client.DescribeHostAvailabilityListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring data in a namespace.
//
// Description:
//
// ## [](#)Prerequisites
//
// Hybrid Cloud Monitoring is activated. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// ## [](#)Limits
//
// The size of monitoring data that is returned in each call cannot exceed 1.5 MB. If the returned data reaches the upper limit, the query fails. You must reset the query conditions.
//
// ## [](#)Operation description
//
// This topic provides an example to show how to query the monitoring data of the `AliyunEcs_cpu_total` metric in the `default-aliyun` namespace from `1653804865` (14:14:25 on May 29, 2022) to `1653805225` (14:20:25 on May 29, 2022).
//
// @param request - DescribeHybridMonitorDataListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridMonitorDataListResponse
func (client *Client) DescribeHybridMonitorDataListWithOptions(request *DescribeHybridMonitorDataListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridMonitorDataListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PromSQL) {
		query["PromSQL"] = request.PromSQL
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridMonitorDataList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridMonitorDataListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data in a namespace.
//
// Description:
//
// ## [](#)Prerequisites
//
// Hybrid Cloud Monitoring is activated. For more information, see [Activate Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/250773.html).
//
// ## [](#)Limits
//
// The size of monitoring data that is returned in each call cannot exceed 1.5 MB. If the returned data reaches the upper limit, the query fails. You must reset the query conditions.
//
// ## [](#)Operation description
//
// This topic provides an example to show how to query the monitoring data of the `AliyunEcs_cpu_total` metric in the `default-aliyun` namespace from `1653804865` (14:14:25 on May 29, 2022) to `1653805225` (14:20:25 on May 29, 2022).
//
// @param request - DescribeHybridMonitorDataListRequest
//
// @return DescribeHybridMonitorDataListResponse
func (client *Client) DescribeHybridMonitorDataList(request *DescribeHybridMonitorDataListRequest) (_result *DescribeHybridMonitorDataListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridMonitorDataListResponse{}
	_body, _err := client.DescribeHybridMonitorDataListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of namespaces and the details of their data sources.
//
// Description:
//
// This example shows how to query all namespaces that belong to the current account. The response indicates that only one namespace, `aliyun-test`, exists.
//
// @param request - DescribeHybridMonitorNamespaceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridMonitorNamespaceListResponse
func (client *Client) DescribeHybridMonitorNamespaceListWithOptions(request *DescribeHybridMonitorNamespaceListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridMonitorNamespaceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ShowTaskStatistic) {
		query["ShowTaskStatistic"] = request.ShowTaskStatistic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridMonitorNamespaceList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridMonitorNamespaceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of namespaces and the details of their data sources.
//
// Description:
//
// This example shows how to query all namespaces that belong to the current account. The response indicates that only one namespace, `aliyun-test`, exists.
//
// @param request - DescribeHybridMonitorNamespaceListRequest
//
// @return DescribeHybridMonitorNamespaceListResponse
func (client *Client) DescribeHybridMonitorNamespaceList(request *DescribeHybridMonitorNamespaceListRequest) (_result *DescribeHybridMonitorNamespaceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridMonitorNamespaceListResponse{}
	_body, _err := client.DescribeHybridMonitorNamespaceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of Logstore groups by calling the DescribeHybridMonitorSLSGroup operation.
//
// Description:
//
// This topic provides an example of how to query all Logstore groups in the current account. The response shows that the current account has two Logstore groups: `Logstore_test` and `Logstore_aliyun`.
//
// @param request - DescribeHybridMonitorSLSGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridMonitorSLSGroupResponse
func (client *Client) DescribeHybridMonitorSLSGroupWithOptions(request *DescribeHybridMonitorSLSGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridMonitorSLSGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SLSGroupName) {
		query["SLSGroupName"] = request.SLSGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridMonitorSLSGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridMonitorSLSGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of Logstore groups by calling the DescribeHybridMonitorSLSGroup operation.
//
// Description:
//
// This topic provides an example of how to query all Logstore groups in the current account. The response shows that the current account has two Logstore groups: `Logstore_test` and `Logstore_aliyun`.
//
// @param request - DescribeHybridMonitorSLSGroupRequest
//
// @return DescribeHybridMonitorSLSGroupResponse
func (client *Client) DescribeHybridMonitorSLSGroup(request *DescribeHybridMonitorSLSGroupRequest) (_result *DescribeHybridMonitorSLSGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridMonitorSLSGroupResponse{}
	_body, _err := client.DescribeHybridMonitorSLSGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metric import tasks.
//
// Description:
//
// This topic provides an example on how to query all metric import tasks that belong to the current Alibaba Cloud account. The returned result indicates that the current account has only one metric import task. The metric import task is named `aliyun_task`.
//
// @param request - DescribeHybridMonitorTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridMonitorTaskListResponse
func (client *Client) DescribeHybridMonitorTaskListWithOptions(request *DescribeHybridMonitorTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridMonitorTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IncludeAliyunTask) {
		query["IncludeAliyunTask"] = request.IncludeAliyunTask
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridMonitorTaskList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridMonitorTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metric import tasks.
//
// Description:
//
// This topic provides an example on how to query all metric import tasks that belong to the current Alibaba Cloud account. The returned result indicates that the current account has only one metric import task. The metric import task is named `aliyun_task`.
//
// @param request - DescribeHybridMonitorTaskListRequest
//
// @return DescribeHybridMonitorTaskListResponse
func (client *Client) DescribeHybridMonitorTaskList(request *DescribeHybridMonitorTaskListRequest) (_result *DescribeHybridMonitorTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridMonitorTaskListResponse{}
	_body, _err := client.DescribeHybridMonitorTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a Log Monitoring task.
//
// @param request - DescribeLogMonitorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogMonitorAttributeResponse
func (client *Client) DescribeLogMonitorAttributeWithOptions(request *DescribeLogMonitorAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogMonitorAttributeResponse, _err error) {
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
		Action:      dara.String("DescribeLogMonitorAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogMonitorAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a Log Monitoring task.
//
// @param request - DescribeLogMonitorAttributeRequest
//
// @return DescribeLogMonitorAttributeResponse
func (client *Client) DescribeLogMonitorAttribute(request *DescribeLogMonitorAttributeRequest) (_result *DescribeLogMonitorAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogMonitorAttributeResponse{}
	_body, _err := client.DescribeLogMonitorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries log monitoring metrics.
//
// @param request - DescribeLogMonitorListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogMonitorListResponse
func (client *Client) DescribeLogMonitorListWithOptions(request *DescribeLogMonitorListRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogMonitorListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchValue) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogMonitorList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogMonitorListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries log monitoring metrics.
//
// @param request - DescribeLogMonitorListRequest
//
// @return DescribeLogMonitorListResponse
func (client *Client) DescribeLogMonitorList(request *DescribeLogMonitorListRequest) (_result *DescribeLogMonitorListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogMonitorListResponse{}
	_body, _err := client.DescribeLogMonitorListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of a metric for a cloud service.
//
// Description:
//
// ### [](#)Limits
//
// - The total free quota is 1 million calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If the free quota is used up and CloudMonitor Basic (pay-as-you-go) is not activated, these API operations can no longer be called as expected. If you have activated CloudMonitor Basic (pay-as-you-go), these API operations can still be called even if the free quota is used up. If the free quota is used up, you are automatically charged for the excess usage based on the pay-as-you-go billing method. For more information about how to activate CloudMonitor Basic (pay-as-you-go), see [Enable the pay-as-you-go billing method](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
// - Each API operation can be called up to 10 times per second. An Alibaba Cloud account and the Resource Access Management (RAM) users within the account share the quota.
//
// ### [](#)Description
//
// > Different from [DescribeMetricList](https://help.aliyun.com/document_detail/51936.html), the DescribeMetricData operation provides statistical features. You can set the Dimension parameter to `{"instanceId": "i-abcdefgh12****"}` to aggregate all data of your Alibaba Cloud account.
//
// This topic provides an example on how to query the monitoring data of the `cpu_idle` metric for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`.
//
// @param request - DescribeMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricDataResponse
func (client *Client) DescribeMetricDataWithOptions(request *DescribeMetricDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Express) {
		query["Express"] = request.Express
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricData"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of a metric for a cloud service.
//
// Description:
//
// ### [](#)Limits
//
// - The total free quota is 1 million calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If the free quota is used up and CloudMonitor Basic (pay-as-you-go) is not activated, these API operations can no longer be called as expected. If you have activated CloudMonitor Basic (pay-as-you-go), these API operations can still be called even if the free quota is used up. If the free quota is used up, you are automatically charged for the excess usage based on the pay-as-you-go billing method. For more information about how to activate CloudMonitor Basic (pay-as-you-go), see [Enable the pay-as-you-go billing method](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
// - Each API operation can be called up to 10 times per second. An Alibaba Cloud account and the Resource Access Management (RAM) users within the account share the quota.
//
// ### [](#)Description
//
// > Different from [DescribeMetricList](https://help.aliyun.com/document_detail/51936.html), the DescribeMetricData operation provides statistical features. You can set the Dimension parameter to `{"instanceId": "i-abcdefgh12****"}` to aggregate all data of your Alibaba Cloud account.
//
// This topic provides an example on how to query the monitoring data of the `cpu_idle` metric for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`.
//
// @param request - DescribeMetricDataRequest
//
// @return DescribeMetricDataResponse
func (client *Client) DescribeMetricData(request *DescribeMetricDataRequest) (_result *DescribeMetricDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricDataResponse{}
	_body, _err := client.DescribeMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the latest monitoring data of a metric.
//
// Description:
//
// ### [](#)Limits
//
// - The total free quota is 1 million calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If the free quota is used up and CloudMonitor Basic (pay-as-you-go) is not activated, these API operations can no longer be called as expected. If you have activated CloudMonitor Basic (pay-as-you-go), these API operations can still be called even if the free quota is used up. After the free quota is used up, you are charged for the excess usage based on the pay-as-you-go billing method. For more information about how to activate CloudMonitor Basic (pay-as-you-go), see [Enable the pay-as-you-go billing method](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
// - Each API operation can be called up to 50 times per second. An Alibaba Cloud account and the Resource Access Management (RAM) users within the account share the quota.
//
// > If `Throttling.User` or `Request was denied due to user flow control` is returned when you call an API operation, the API operation is throttled. For more information about how to handle the issue, see [How do I handle the throttling of a query API?](https://help.aliyun.com/document_detail/2615031.html)
//
// ### [](#)Precautions
//
// The storage duration of the monitoring data of each cloud service is related to the `Period` parameter (statistical period). A larger value of the `Period` parameter indicates that the monitoring data is distributed in a larger time range and the storage duration of the monitoring data is longer. The following list describes the specific relationships:
//
// - The storage duration is 7 days if the value of the `Period` parameter is less than 60 seconds.
//
// - The storage duration is 31 days if the value of the `Period` parameter is 60 seconds.
//
// - The storage duration is 91 days if the value of the `Period` parameter is greater than or equal to 300 seconds.
//
// ### [](#)Operation description
//
// This topic provides an example on how to query the latest monitoring data of the `CPUUtilization` metric for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`. The returned result indicates that the monitoring data for the instance `i-abcdefgh12****` of the account `123456789876****` is queried at an interval of 60 seconds. The maximum, minimum, and average values of the metric are 100, 93.1, and 99.52.
//
// @param request - DescribeMetricLastRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricLastResponse
func (client *Client) DescribeMetricLastWithOptions(request *DescribeMetricLastRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricLastResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Express) {
		query["Express"] = request.Express
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricLast"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricLastResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latest monitoring data of a metric.
//
// Description:
//
// ### [](#)Limits
//
// - The total free quota is 1 million calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If the free quota is used up and CloudMonitor Basic (pay-as-you-go) is not activated, these API operations can no longer be called as expected. If you have activated CloudMonitor Basic (pay-as-you-go), these API operations can still be called even if the free quota is used up. After the free quota is used up, you are charged for the excess usage based on the pay-as-you-go billing method. For more information about how to activate CloudMonitor Basic (pay-as-you-go), see [Enable the pay-as-you-go billing method](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
// - Each API operation can be called up to 50 times per second. An Alibaba Cloud account and the Resource Access Management (RAM) users within the account share the quota.
//
// > If `Throttling.User` or `Request was denied due to user flow control` is returned when you call an API operation, the API operation is throttled. For more information about how to handle the issue, see [How do I handle the throttling of a query API?](https://help.aliyun.com/document_detail/2615031.html)
//
// ### [](#)Precautions
//
// The storage duration of the monitoring data of each cloud service is related to the `Period` parameter (statistical period). A larger value of the `Period` parameter indicates that the monitoring data is distributed in a larger time range and the storage duration of the monitoring data is longer. The following list describes the specific relationships:
//
// - The storage duration is 7 days if the value of the `Period` parameter is less than 60 seconds.
//
// - The storage duration is 31 days if the value of the `Period` parameter is 60 seconds.
//
// - The storage duration is 91 days if the value of the `Period` parameter is greater than or equal to 300 seconds.
//
// ### [](#)Operation description
//
// This topic provides an example on how to query the latest monitoring data of the `CPUUtilization` metric for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`. The returned result indicates that the monitoring data for the instance `i-abcdefgh12****` of the account `123456789876****` is queried at an interval of 60 seconds. The maximum, minimum, and average values of the metric are 100, 93.1, and 99.52.
//
// @param request - DescribeMetricLastRequest
//
// @return DescribeMetricLastResponse
func (client *Client) DescribeMetricLast(request *DescribeMetricLastRequest) (_result *DescribeMetricLastResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricLastResponse{}
	_body, _err := client.DescribeMetricLastWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The DescribeMetricList operation queries the monitoring data of a specific metric for a cloud service.
//
// Description:
//
// ### Limits
//
// - You have a free quota of 1 million total API calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If you use up the free quota and have not enabled the pay-as-you-go billing method for CloudMonitor Basic, you can no longer use these API operations. If you have enabled the pay-as-you-go billing method, you can continue to use the API operations after the free quota is used up. API calls that exceed the free quota are automatically charged on a pay-as-you-go basis. For more information, see [Enable pay-as-you-go](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
// - You can call each API operation up to 50 times per second. This limit is shared between an Alibaba Cloud account and its RAM users.
//
// > If you receive the `Throttling.User` or `Request was denied due to user flow control` error message when you call an API operation, the API call is throttled. For more information, see [How do I resolve an API call throttling issue?](https://help.aliyun.com/document_detail/2615031.html).
//
// ### Notes
//
// The storage duration of monitoring data for a cloud service depends on the `Period` (statistical period). A larger `Period` value indicates that the monitoring data is less granular and is stored for a longer period. The relationship is as follows:
//
// - If the value of `Period` is less than 60 seconds, the storage duration is 7 days.
//
// - If the value of `Period` is 60 seconds, the storage duration is 31 days.
//
// - If the value of `Period` is 300 seconds or greater, the storage duration is 91 days.
//
// ### Usage notes
//
// This topic provides an example of how to query the monitoring data of the `cpu_idle` metric for the `acs_ecs_dashboard` cloud service. The response shows the data for the instance `i-abcdefgh12****`, which belongs to the Alibaba Cloud account `120886317861****`. At a 60 second interval, the maximum, minimum, and average values of the metric are 100, 93.1, and 99.52, respectively.
//
// @param request - DescribeMetricListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricListResponse
func (client *Client) DescribeMetricListWithOptions(request *DescribeMetricListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Express) {
		query["Express"] = request.Express
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The DescribeMetricList operation queries the monitoring data of a specific metric for a cloud service.
//
// Description:
//
// ### Limits
//
// - You have a free quota of 1 million total API calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If you use up the free quota and have not enabled the pay-as-you-go billing method for CloudMonitor Basic, you can no longer use these API operations. If you have enabled the pay-as-you-go billing method, you can continue to use the API operations after the free quota is used up. API calls that exceed the free quota are automatically charged on a pay-as-you-go basis. For more information, see [Enable pay-as-you-go](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
// - You can call each API operation up to 50 times per second. This limit is shared between an Alibaba Cloud account and its RAM users.
//
// > If you receive the `Throttling.User` or `Request was denied due to user flow control` error message when you call an API operation, the API call is throttled. For more information, see [How do I resolve an API call throttling issue?](https://help.aliyun.com/document_detail/2615031.html).
//
// ### Notes
//
// The storage duration of monitoring data for a cloud service depends on the `Period` (statistical period). A larger `Period` value indicates that the monitoring data is less granular and is stored for a longer period. The relationship is as follows:
//
// - If the value of `Period` is less than 60 seconds, the storage duration is 7 days.
//
// - If the value of `Period` is 60 seconds, the storage duration is 31 days.
//
// - If the value of `Period` is 300 seconds or greater, the storage duration is 91 days.
//
// ### Usage notes
//
// This topic provides an example of how to query the monitoring data of the `cpu_idle` metric for the `acs_ecs_dashboard` cloud service. The response shows the data for the instance `i-abcdefgh12****`, which belongs to the Alibaba Cloud account `120886317861****`. At a 60 second interval, the maximum, minimum, and average values of the metric are 100, 93.1, and 99.52, respectively.
//
// @param request - DescribeMetricListRequest
//
// @return DescribeMetricListResponse
func (client *Client) DescribeMetricList(request *DescribeMetricListRequest) (_result *DescribeMetricListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricListResponse{}
	_body, _err := client.DescribeMetricListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metrics that are supported in CloudMonitor.
//
// Description:
//
// Use this operation together with DescribeMetricList and DescribeMetricLast. For more information, see [DescribeMetricList](https://help.aliyun.com/document_detail/51936.html) and [DescribeMetricLast](https://help.aliyun.com/document_detail/51939.html).
//
// @param request - DescribeMetricMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricMetaListResponse
func (client *Client) DescribeMetricMetaListWithOptions(request *DescribeMetricMetaListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricMetaListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
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
		Action:      dara.String("DescribeMetricMetaList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricMetaListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metrics that are supported in CloudMonitor.
//
// Description:
//
// Use this operation together with DescribeMetricList and DescribeMetricLast. For more information, see [DescribeMetricList](https://help.aliyun.com/document_detail/51936.html) and [DescribeMetricLast](https://help.aliyun.com/document_detail/51939.html).
//
// @param request - DescribeMetricMetaListRequest
//
// @return DescribeMetricMetaListResponse
func (client *Client) DescribeMetricMetaList(request *DescribeMetricMetaListRequest) (_result *DescribeMetricMetaListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricMetaListResponse{}
	_body, _err := client.DescribeMetricMetaListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeMetricRuleBlackList operation to query blacklist policies.
//
// @param request - DescribeMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleBlackListResponse
func (client *Client) DescribeMetricRuleBlackListWithOptions(request *DescribeMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleBlackListResponse, _err error) {
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

	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
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

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeMetricRuleBlackList operation to query blacklist policies.
//
// @param request - DescribeMetricRuleBlackListRequest
//
// @return DescribeMetricRuleBlackListResponse
func (client *Client) DescribeMetricRuleBlackList(request *DescribeMetricRuleBlackListRequest) (_result *DescribeMetricRuleBlackListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricRuleBlackListResponse{}
	_body, _err := client.DescribeMetricRuleBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of alert rules in each state.
//
// @param request - DescribeMetricRuleCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleCountResponse
func (client *Client) DescribeMetricRuleCountWithOptions(request *DescribeMetricRuleCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleCountResponse, _err error) {
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
		Action:      dara.String("DescribeMetricRuleCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of alert rules in each state.
//
// @param request - DescribeMetricRuleCountRequest
//
// @return DescribeMetricRuleCountResponse
func (client *Client) DescribeMetricRuleCount(request *DescribeMetricRuleCountRequest) (_result *DescribeMetricRuleCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricRuleCountResponse{}
	_body, _err := client.DescribeMetricRuleCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all alert rules in the alert rule list.
//
// Description:
//
// This topic provides an example of how to query all alert rules in the alert rule list of the current Alibaba Cloud account. The response shows that the alert rule list contains only one alert rule. The alert rule is named `Rule_01` and has an ID of `applyTemplate344cfd42-0f32-4fd6-805a-88d7908a****`.
//
// @param request - DescribeMetricRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleListResponse
func (client *Client) DescribeMetricRuleListWithOptions(request *DescribeMetricRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertState) {
		query["AlertState"] = request.AlertState
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EnableState) {
		query["EnableState"] = request.EnableState
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleIds) {
		query["RuleIds"] = request.RuleIds
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all alert rules in the alert rule list.
//
// Description:
//
// This topic provides an example of how to query all alert rules in the alert rule list of the current Alibaba Cloud account. The response shows that the alert rule list contains only one alert rule. The alert rule is named `Rule_01` and has an ID of `applyTemplate344cfd42-0f32-4fd6-805a-88d7908a****`.
//
// @param request - DescribeMetricRuleListRequest
//
// @return DescribeMetricRuleListResponse
func (client *Client) DescribeMetricRuleList(request *DescribeMetricRuleListRequest) (_result *DescribeMetricRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricRuleListResponse{}
	_body, _err := client.DescribeMetricRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the notification channels of a specified alert rule by calling the DescribeMetricRuleTargets operation.
//
// Description:
//
// ## Usage notes
//
// This topic provides an example to query the target resources associated with the alert rule `ae06917_75a8c43178ab66****`.
//
// @param request - DescribeMetricRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleTargetsResponse
func (client *Client) DescribeMetricRuleTargetsWithOptions(request *DescribeMetricRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the notification channels of a specified alert rule by calling the DescribeMetricRuleTargets operation.
//
// Description:
//
// ## Usage notes
//
// This topic provides an example to query the target resources associated with the alert rule `ae06917_75a8c43178ab66****`.
//
// @param request - DescribeMetricRuleTargetsRequest
//
// @return DescribeMetricRuleTargetsResponse
func (client *Client) DescribeMetricRuleTargets(request *DescribeMetricRuleTargetsRequest) (_result *DescribeMetricRuleTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricRuleTargetsResponse{}
	_body, _err := client.DescribeMetricRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an alert template by calling the DescribeMetricRuleTemplateAttribute operation.
//
// Description:
//
// This topic provides an example on how to query the details of the alert template `70****`.
//
// @param request - DescribeMetricRuleTemplateAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleTemplateAttributeResponse
func (client *Client) DescribeMetricRuleTemplateAttributeWithOptions(request *DescribeMetricRuleTemplateAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleTemplateAttributeResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleTemplateAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleTemplateAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an alert template by calling the DescribeMetricRuleTemplateAttribute operation.
//
// Description:
//
// This topic provides an example on how to query the details of the alert template `70****`.
//
// @param request - DescribeMetricRuleTemplateAttributeRequest
//
// @return DescribeMetricRuleTemplateAttributeResponse
func (client *Client) DescribeMetricRuleTemplateAttribute(request *DescribeMetricRuleTemplateAttributeRequest) (_result *DescribeMetricRuleTemplateAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricRuleTemplateAttributeResponse{}
	_body, _err := client.DescribeMetricRuleTemplateAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries alert templates.
//
// Description:
//
// This topic provides an example on how to query alert templates. In this example, the following alert templates are returned in the response: `ECS_Template1` and `ECS_Template2`.
//
// @param request - DescribeMetricRuleTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricRuleTemplateListResponse
func (client *Client) DescribeMetricRuleTemplateListWithOptions(request *DescribeMetricRuleTemplateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricRuleTemplateListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.History) {
		query["History"] = request.History
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricRuleTemplateList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricRuleTemplateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert templates.
//
// Description:
//
// This topic provides an example on how to query alert templates. In this example, the following alert templates are returned in the response: `ECS_Template1` and `ECS_Template2`.
//
// @param request - DescribeMetricRuleTemplateListRequest
//
// @return DescribeMetricRuleTemplateListResponse
func (client *Client) DescribeMetricRuleTemplateList(request *DescribeMetricRuleTemplateListRequest) (_result *DescribeMetricRuleTemplateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricRuleTemplateListResponse{}
	_body, _err := client.DescribeMetricRuleTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the latest monitoring data of a metric for a cloud service. The data can be sorted by a specified order.
//
// Description:
//
// ### [](#)Limits
//
// - The total free quota is 1 million calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If the free quota is used up and CloudMonitor Basic (pay-as-you-go) is not activated, these API operations can no longer be called as expected. If you have activated CloudMonitor Basic (pay-as-you-go), these API operations can still be called even if the free quota is used up. After the free quota is used up, you are charged for the excess usage based on the pay-as-you-go billing method. For more information about how to activate CloudMonitor Basic (pay-as-you-go), see [Enable the pay-as-you-go billing method](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
// - Each API operation can be called up to 10 times per second. An Alibaba Cloud account and the Resource Access Management (RAM) users within the account share the quota.
//
// ### [](#)Precautions
//
// The storage duration of the monitoring data of each cloud service is related to the `Period` parameter (statistical period). A larger value of the `Period` parameter indicates that the monitoring data is distributed in a larger time range and the storage duration of the monitoring data is longer. The following list describes the specific relationships:
//
// - The storage duration is 7 days if the value of the `Period` parameter is less than 60 seconds.
//
// - The storage duration is 31 days if the value of the `Period` parameter is 60 seconds.
//
// - The storage duration is 91 days if the value of the `Period` is greater than or equal to 300 seconds.
//
// ### [](#)Operation description
//
// This topic provides an example on how to query the monitoring data of the `cpu_idle` metric in the last 60 seconds for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`. The monitoring data is sorted in descending order based on the `Average` field.
//
// @param request - DescribeMetricTopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetricTopResponse
func (client *Client) DescribeMetricTopWithOptions(request *DescribeMetricTopRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetricTopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Express) {
		query["Express"] = request.Express
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.OrderDesc) {
		query["OrderDesc"] = request.OrderDesc
	}

	if !dara.IsNil(request.Orderby) {
		query["Orderby"] = request.Orderby
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetricTop"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetricTopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latest monitoring data of a metric for a cloud service. The data can be sorted by a specified order.
//
// Description:
//
// ### [](#)Limits
//
// - The total free quota is 1 million calls per month for the DescribeMetricLast, DescribeMetricList, DescribeMetricData, and DescribeMetricTop operations. If the free quota is used up and CloudMonitor Basic (pay-as-you-go) is not activated, these API operations can no longer be called as expected. If you have activated CloudMonitor Basic (pay-as-you-go), these API operations can still be called even if the free quota is used up. After the free quota is used up, you are charged for the excess usage based on the pay-as-you-go billing method. For more information about how to activate CloudMonitor Basic (pay-as-you-go), see [Enable the pay-as-you-go billing method](https://common-buy.aliyun.com/?spm=a2c4g.11186623.0.0.6c8f3481IbSHgG\\&commodityCode=cms_basic_public_cn\\&from_biz_channel=help_bill).
//
// - Each API operation can be called up to 10 times per second. An Alibaba Cloud account and the Resource Access Management (RAM) users within the account share the quota.
//
// ### [](#)Precautions
//
// The storage duration of the monitoring data of each cloud service is related to the `Period` parameter (statistical period). A larger value of the `Period` parameter indicates that the monitoring data is distributed in a larger time range and the storage duration of the monitoring data is longer. The following list describes the specific relationships:
//
// - The storage duration is 7 days if the value of the `Period` parameter is less than 60 seconds.
//
// - The storage duration is 31 days if the value of the `Period` parameter is 60 seconds.
//
// - The storage duration is 91 days if the value of the `Period` is greater than or equal to 300 seconds.
//
// ### [](#)Operation description
//
// This topic provides an example on how to query the monitoring data of the `cpu_idle` metric in the last 60 seconds for Elastic Compute Service (ECS). The namespace of ECS is `acs_ecs_dashboard`. The monitoring data is sorted in descending order based on the `Average` field.
//
// @param request - DescribeMetricTopRequest
//
// @return DescribeMetricTopResponse
func (client *Client) DescribeMetricTop(request *DescribeMetricTopRequest) (_result *DescribeMetricTopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMetricTopResponse{}
	_body, _err := client.DescribeMetricTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cloud services to which the resources in an application group belong and the number of resources that belong to each cloud service in the application group.
//
// @param request - DescribeMonitorGroupCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupCategoriesResponse
func (client *Client) DescribeMonitorGroupCategoriesWithOptions(request *DescribeMonitorGroupCategoriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupCategoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroupCategories"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud services to which the resources in an application group belong and the number of resources that belong to each cloud service in the application group.
//
// @param request - DescribeMonitorGroupCategoriesRequest
//
// @return DescribeMonitorGroupCategoriesResponse
func (client *Client) DescribeMonitorGroupCategories(request *DescribeMonitorGroupCategoriesRequest) (_result *DescribeMonitorGroupCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorGroupCategoriesResponse{}
	_body, _err := client.DescribeMonitorGroupCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the dynamic rules of an application group.
//
// @param request - DescribeMonitorGroupDynamicRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupDynamicRulesResponse
func (client *Client) DescribeMonitorGroupDynamicRulesWithOptions(request *DescribeMonitorGroupDynamicRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupDynamicRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroupDynamicRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupDynamicRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the dynamic rules of an application group.
//
// @param request - DescribeMonitorGroupDynamicRulesRequest
//
// @return DescribeMonitorGroupDynamicRulesResponse
func (client *Client) DescribeMonitorGroupDynamicRules(request *DescribeMonitorGroupDynamicRulesRequest) (_result *DescribeMonitorGroupDynamicRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorGroupDynamicRulesResponse{}
	_body, _err := client.DescribeMonitorGroupDynamicRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources in an application group.
//
// @param request - DescribeMonitorGroupInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupInstanceAttributeResponse
func (client *Client) DescribeMonitorGroupInstanceAttributeWithOptions(request *DescribeMonitorGroupInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupInstanceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
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

	if !dara.IsNil(request.Total) {
		query["Total"] = request.Total
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroupInstanceAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources in an application group.
//
// @param request - DescribeMonitorGroupInstanceAttributeRequest
//
// @return DescribeMonitorGroupInstanceAttributeResponse
func (client *Client) DescribeMonitorGroupInstanceAttribute(request *DescribeMonitorGroupInstanceAttributeRequest) (_result *DescribeMonitorGroupInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorGroupInstanceAttributeResponse{}
	_body, _err := client.DescribeMonitorGroupInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources in an application group.
//
// @param request - DescribeMonitorGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupInstancesResponse
func (client *Client) DescribeMonitorGroupInstancesWithOptions(request *DescribeMonitorGroupInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupInstancesResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroupInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources in an application group.
//
// @param request - DescribeMonitorGroupInstancesRequest
//
// @return DescribeMonitorGroupInstancesResponse
func (client *Client) DescribeMonitorGroupInstances(request *DescribeMonitorGroupInstancesRequest) (_result *DescribeMonitorGroupInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorGroupInstancesResponse{}
	_body, _err := client.DescribeMonitorGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of alert notification pause policies for an application group by calling the DescribeMonitorGroupNotifyPolicyList operation.
//
// @param request - DescribeMonitorGroupNotifyPolicyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupNotifyPolicyListResponse
func (client *Client) DescribeMonitorGroupNotifyPolicyListWithOptions(request *DescribeMonitorGroupNotifyPolicyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupNotifyPolicyListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
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
		Action:      dara.String("DescribeMonitorGroupNotifyPolicyList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupNotifyPolicyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of alert notification pause policies for an application group by calling the DescribeMonitorGroupNotifyPolicyList operation.
//
// @param request - DescribeMonitorGroupNotifyPolicyListRequest
//
// @return DescribeMonitorGroupNotifyPolicyListResponse
func (client *Client) DescribeMonitorGroupNotifyPolicyList(request *DescribeMonitorGroupNotifyPolicyListRequest) (_result *DescribeMonitorGroupNotifyPolicyListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorGroupNotifyPolicyListResponse{}
	_body, _err := client.DescribeMonitorGroupNotifyPolicyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeMonitorGroups operation to query a list of application groups.
//
// Description:
//
// This topic provides an example of how to query a list of application groups. The response shows that two application groups are returned: `testGroup124` and `test123`.
//
// @param request - DescribeMonitorGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorGroupsResponse
func (client *Client) DescribeMonitorGroupsWithOptions(request *DescribeMonitorGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicTagRuleId) {
		query["DynamicTagRuleId"] = request.DynamicTagRuleId
	}

	if !dara.IsNil(request.GroupFounderTagKey) {
		query["GroupFounderTagKey"] = request.GroupFounderTagKey
	}

	if !dara.IsNil(request.GroupFounderTagValue) {
		query["GroupFounderTagValue"] = request.GroupFounderTagValue
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IncludeTemplateHistory) {
		query["IncludeTemplateHistory"] = request.IncludeTemplateHistory
	}

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

	if !dara.IsNil(request.SelectContactGroups) {
		query["SelectContactGroups"] = request.SelectContactGroups
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Types) {
		query["Types"] = request.Types
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorGroups"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeMonitorGroups operation to query a list of application groups.
//
// Description:
//
// This topic provides an example of how to query a list of application groups. The response shows that two application groups are returned: `testGroup124` and `test123`.
//
// @param request - DescribeMonitorGroupsRequest
//
// @return DescribeMonitorGroupsResponse
func (client *Client) DescribeMonitorGroups(request *DescribeMonitorGroupsRequest) (_result *DescribeMonitorGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorGroupsResponse{}
	_body, _err := client.DescribeMonitorGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource quotas of CloudMonitor.
//
// @param request - DescribeMonitorResourceQuotaAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorResourceQuotaAttributeResponse
func (client *Client) DescribeMonitorResourceQuotaAttributeWithOptions(request *DescribeMonitorResourceQuotaAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitorResourceQuotaAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ShowUsed) {
		query["ShowUsed"] = request.ShowUsed
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitorResourceQuotaAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitorResourceQuotaAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource quotas of CloudMonitor.
//
// @param request - DescribeMonitorResourceQuotaAttributeRequest
//
// @return DescribeMonitorResourceQuotaAttributeResponse
func (client *Client) DescribeMonitorResourceQuotaAttribute(request *DescribeMonitorResourceQuotaAttributeRequest) (_result *DescribeMonitorResourceQuotaAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitorResourceQuotaAttributeResponse{}
	_body, _err := client.DescribeMonitorResourceQuotaAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the AccessKey ID and AccessKey secret that are required to install the CloudMonitor agent on a third-party host.
//
// @param request - DescribeMonitoringAgentAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentAccessKeyResponse
func (client *Client) DescribeMonitoringAgentAccessKeyWithOptions(request *DescribeMonitoringAgentAccessKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentAccessKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringAgentAccessKey"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AccessKey ID and AccessKey secret that are required to install the CloudMonitor agent on a third-party host.
//
// @param request - DescribeMonitoringAgentAccessKeyRequest
//
// @return DescribeMonitoringAgentAccessKeyResponse
func (client *Client) DescribeMonitoringAgentAccessKey(request *DescribeMonitoringAgentAccessKeyRequest) (_result *DescribeMonitoringAgentAccessKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitoringAgentAccessKeyResponse{}
	_body, _err := client.DescribeMonitoringAgentAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of the CloudMonitor agent.
//
// @param request - DescribeMonitoringAgentConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentConfigResponse
func (client *Client) DescribeMonitoringAgentConfigWithOptions(request *DescribeMonitoringAgentConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringAgentConfig"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of the CloudMonitor agent.
//
// @param request - DescribeMonitoringAgentConfigRequest
//
// @return DescribeMonitoringAgentConfigResponse
func (client *Client) DescribeMonitoringAgentConfig(request *DescribeMonitoringAgentConfigRequest) (_result *DescribeMonitoringAgentConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitoringAgentConfigResponse{}
	_body, _err := client.DescribeMonitoringAgentConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeMonitoringAgentHosts operation to query a list of all hosts, regardless of whether the CloudMonitor agent is installed.
//
// @param request - DescribeMonitoringAgentHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentHostsResponse
func (client *Client) DescribeMonitoringAgentHostsWithOptions(request *DescribeMonitoringAgentHostsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentHostsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunHost) {
		query["AliyunHost"] = request.AliyunHost
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InstanceRegionId) {
		query["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SerialNumbers) {
		query["SerialNumbers"] = request.SerialNumbers
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.SysomStatus) {
		query["SysomStatus"] = request.SysomStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringAgentHosts"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeMonitoringAgentHosts operation to query a list of all hosts, regardless of whether the CloudMonitor agent is installed.
//
// @param request - DescribeMonitoringAgentHostsRequest
//
// @return DescribeMonitoringAgentHostsResponse
func (client *Client) DescribeMonitoringAgentHosts(request *DescribeMonitoringAgentHostsRequest) (_result *DescribeMonitoringAgentHostsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitoringAgentHostsResponse{}
	_body, _err := client.DescribeMonitoringAgentHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DescribeMonitoringAgentProcesses operation to query the list of processes for a specified resource.
//
// Description:
//
// > Before you call this operation, you must call the CreateMonitoringAgentProcess operation to create a process monitoring task. For more information, see [CreateMonitoringAgentProcess](https://help.aliyun.com/document_detail/2513212.html).
//
// This topic provides an example of how to query the list of processes for the resource i-hp3hl3cx1pbahzy8\\*\\*\\*\\*. The response shows the details of the Nginx and HTTP processes.
//
// @param request - DescribeMonitoringAgentProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentProcessesResponse
func (client *Client) DescribeMonitoringAgentProcessesWithOptions(request *DescribeMonitoringAgentProcessesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentProcessesResponse, _err error) {
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
		Action:      dara.String("DescribeMonitoringAgentProcesses"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeMonitoringAgentProcesses operation to query the list of processes for a specified resource.
//
// Description:
//
// > Before you call this operation, you must call the CreateMonitoringAgentProcess operation to create a process monitoring task. For more information, see [CreateMonitoringAgentProcess](https://help.aliyun.com/document_detail/2513212.html).
//
// This topic provides an example of how to query the list of processes for the resource i-hp3hl3cx1pbahzy8\\*\\*\\*\\*. The response shows the details of the Nginx and HTTP processes.
//
// @param request - DescribeMonitoringAgentProcessesRequest
//
// @return DescribeMonitoringAgentProcessesResponse
func (client *Client) DescribeMonitoringAgentProcesses(request *DescribeMonitoringAgentProcessesRequest) (_result *DescribeMonitoringAgentProcessesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitoringAgentProcessesResponse{}
	_body, _err := client.DescribeMonitoringAgentProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the running status of the CloudMonitor agent by calling the DescribeMonitoringAgentStatuses operation.
//
// Description:
//
// This topic describes how to query the status of the CloudMonitor agent that is installed on the `i-hp3dunahluwajv6f****` instance. The result indicates that the CloudMonitor agent is in the `running` state.
//
// @param request - DescribeMonitoringAgentStatusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringAgentStatusesResponse
func (client *Client) DescribeMonitoringAgentStatusesWithOptions(request *DescribeMonitoringAgentStatusesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringAgentStatusesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostAvailabilityTaskId) {
		query["HostAvailabilityTaskId"] = request.HostAvailabilityTaskId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringAgentStatuses"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringAgentStatusesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the running status of the CloudMonitor agent by calling the DescribeMonitoringAgentStatuses operation.
//
// Description:
//
// This topic describes how to query the status of the CloudMonitor agent that is installed on the `i-hp3dunahluwajv6f****` instance. The result indicates that the CloudMonitor agent is in the `running` state.
//
// @param request - DescribeMonitoringAgentStatusesRequest
//
// @return DescribeMonitoringAgentStatusesResponse
func (client *Client) DescribeMonitoringAgentStatuses(request *DescribeMonitoringAgentStatusesRequest) (_result *DescribeMonitoringAgentStatusesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitoringAgentStatusesResponse{}
	_body, _err := client.DescribeMonitoringAgentStatusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the global configurations of the CloudMonitor agent.
//
// @param request - DescribeMonitoringConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitoringConfigResponse
func (client *Client) DescribeMonitoringConfigWithOptions(request *DescribeMonitoringConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeMonitoringConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMonitoringConfig"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMonitoringConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global configurations of the CloudMonitor agent.
//
// @param request - DescribeMonitoringConfigRequest
//
// @return DescribeMonitoringConfigResponse
func (client *Client) DescribeMonitoringConfig(request *DescribeMonitoringConfigRequest) (_result *DescribeMonitoringConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMonitoringConfigResponse{}
	_body, _err := client.DescribeMonitoringConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DescribeProductResourceTagKeyList operation to query all tag keys of cloud resources in a specified region.
//
// @param request - DescribeProductResourceTagKeyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductResourceTagKeyListResponse
func (client *Client) DescribeProductResourceTagKeyListWithOptions(request *DescribeProductResourceTagKeyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductResourceTagKeyListResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProductResourceTagKeyList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductResourceTagKeyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeProductResourceTagKeyList operation to query all tag keys of cloud resources in a specified region.
//
// @param request - DescribeProductResourceTagKeyListRequest
//
// @return DescribeProductResourceTagKeyListResponse
func (client *Client) DescribeProductResourceTagKeyList(request *DescribeProductResourceTagKeyListRequest) (_result *DescribeProductResourceTagKeyListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProductResourceTagKeyListResponse{}
	_body, _err := client.DescribeProductResourceTagKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cloud services for which the initiative alert feature is enabled.
//
// @param request - DescribeProductsOfActiveMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductsOfActiveMetricRuleResponse
func (client *Client) DescribeProductsOfActiveMetricRuleWithOptions(request *DescribeProductsOfActiveMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductsOfActiveMetricRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProductsOfActiveMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductsOfActiveMetricRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud services for which the initiative alert feature is enabled.
//
// @param request - DescribeProductsOfActiveMetricRuleRequest
//
// @return DescribeProductsOfActiveMetricRuleResponse
func (client *Client) DescribeProductsOfActiveMetricRule(request *DescribeProductsOfActiveMetricRuleRequest) (_result *DescribeProductsOfActiveMetricRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProductsOfActiveMetricRuleResponse{}
	_body, _err := client.DescribeProductsOfActiveMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DescribeProjectMeta operation to query the list of cloud services that support time series monitoring metrics in CloudMonitor.
//
// Description:
//
// Obtains the information about the connected cloud services, including the description, namespace, and tags of each service.
//
// @param request - DescribeProjectMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectMetaResponse
func (client *Client) DescribeProjectMetaWithOptions(request *DescribeProjectMetaRequest, runtime *dara.RuntimeOptions) (_result *DescribeProjectMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
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
		Action:      dara.String("DescribeProjectMeta"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeProjectMeta operation to query the list of cloud services that support time series monitoring metrics in CloudMonitor.
//
// Description:
//
// Obtains the information about the connected cloud services, including the description, namespace, and tags of each service.
//
// @param request - DescribeProjectMetaRequest
//
// @return DescribeProjectMetaResponse
func (client *Client) DescribeProjectMeta(request *DescribeProjectMetaRequest) (_result *DescribeProjectMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProjectMetaResponse{}
	_body, _err := client.DescribeProjectMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a site monitoring task by calling the DescribeSiteMonitorAttribute operation.
//
// Description:
//
// This topic provides an example on how to query the details of the site monitoring task `cc641dff-c19d-45f3-ad0a-818a0c4f****`. The response shows that the task name is `test123`, the monitored address is `https://aliyun.com`, and the carrier is `Alibaba`.
//
// @param request - DescribeSiteMonitorAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorAttributeResponse
func (client *Client) DescribeSiteMonitorAttributeWithOptions(request *DescribeSiteMonitorAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeAlert) {
		query["IncludeAlert"] = request.IncludeAlert
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a site monitoring task by calling the DescribeSiteMonitorAttribute operation.
//
// Description:
//
// This topic provides an example on how to query the details of the site monitoring task `cc641dff-c19d-45f3-ad0a-818a0c4f****`. The response shows that the task name is `test123`, the monitored address is `https://aliyun.com`, and the carrier is `Alibaba`.
//
// @param request - DescribeSiteMonitorAttributeRequest
//
// @return DescribeSiteMonitorAttributeResponse
func (client *Client) DescribeSiteMonitorAttribute(request *DescribeSiteMonitorAttributeRequest) (_result *DescribeSiteMonitorAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteMonitorAttributeResponse{}
	_body, _err := client.DescribeSiteMonitorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the fine-grained monitoring data of a site monitoring task.
//
// @param request - DescribeSiteMonitorDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorDataResponse
func (client *Client) DescribeSiteMonitorDataWithOptions(request *DescribeSiteMonitorDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorDataResponse, _err error) {
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

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorData"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fine-grained monitoring data of a site monitoring task.
//
// @param request - DescribeSiteMonitorDataRequest
//
// @return DescribeSiteMonitorDataResponse
func (client *Client) DescribeSiteMonitorData(request *DescribeSiteMonitorDataRequest) (_result *DescribeSiteMonitorDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteMonitorDataResponse{}
	_body, _err := client.DescribeSiteMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of all carrier detection points by calling the DescribeSiteMonitorISPCityList operation.
//
// Description:
//
// 本文将提供一个示例，查询运营商“联通”在“贵阳市”的探测点详情。
//
// @param request - DescribeSiteMonitorISPCityListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorISPCityListResponse
func (client *Client) DescribeSiteMonitorISPCityListWithOptions(request *DescribeSiteMonitorISPCityListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorISPCityListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.IPV4) {
		query["IPV4"] = request.IPV4
	}

	if !dara.IsNil(request.IPV6) {
		query["IPV6"] = request.IPV6
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.ViewAll) {
		query["ViewAll"] = request.ViewAll
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorISPCityList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorISPCityListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of all carrier detection points by calling the DescribeSiteMonitorISPCityList operation.
//
// Description:
//
// 本文将提供一个示例，查询运营商“联通”在“贵阳市”的探测点详情。
//
// @param request - DescribeSiteMonitorISPCityListRequest
//
// @return DescribeSiteMonitorISPCityListResponse
func (client *Client) DescribeSiteMonitorISPCityList(request *DescribeSiteMonitorISPCityListRequest) (_result *DescribeSiteMonitorISPCityListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteMonitorISPCityListResponse{}
	_body, _err := client.DescribeSiteMonitorISPCityListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of site monitoring tasks by calling the DescribeSiteMonitorList operation.
//
// Description:
//
// This topic provides an example of how to query the list of site monitoring tasks for the current account. The response shows that the current account has one site monitoring task named `HanZhou_ECS2`.
//
// @param request - DescribeSiteMonitorListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorListResponse
func (client *Client) DescribeSiteMonitorListWithOptions(request *DescribeSiteMonitorListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentGroup) {
		query["AgentGroup"] = request.AgentGroup
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskState) {
		query["TaskState"] = request.TaskState
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of site monitoring tasks by calling the DescribeSiteMonitorList operation.
//
// Description:
//
// This topic provides an example of how to query the list of site monitoring tasks for the current account. The response shows that the current account has one site monitoring task named `HanZhou_ECS2`.
//
// @param request - DescribeSiteMonitorListRequest
//
// @return DescribeSiteMonitorListResponse
func (client *Client) DescribeSiteMonitorList(request *DescribeSiteMonitorListRequest) (_result *DescribeSiteMonitorListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteMonitorListResponse{}
	_body, _err := client.DescribeSiteMonitorListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DescribeSiteMonitorLog operation to query the detection logs for a one-time detection task.
//
// Description:
//
// Only Alibaba Cloud accounts with Network Analysis and Monitoring activated can create one-time detection tasks.
//
// This topic provides an example of how to query the detection logs for the one-time detection task `afa5c3ce-f944-4363-9edb-ce919a29****`.
//
// @param request - DescribeSiteMonitorLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorLogResponse
func (client *Client) DescribeSiteMonitorLogWithOptions(request *DescribeSiteMonitorLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Browser) {
		query["Browser"] = request.Browser
	}

	if !dara.IsNil(request.BrowserInfo) {
		query["BrowserInfo"] = request.BrowserInfo
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Device) {
		query["Device"] = request.Device
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorLog"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeSiteMonitorLog operation to query the detection logs for a one-time detection task.
//
// Description:
//
// Only Alibaba Cloud accounts with Network Analysis and Monitoring activated can create one-time detection tasks.
//
// This topic provides an example of how to query the detection logs for the one-time detection task `afa5c3ce-f944-4363-9edb-ce919a29****`.
//
// @param request - DescribeSiteMonitorLogRequest
//
// @return DescribeSiteMonitorLogResponse
func (client *Client) DescribeSiteMonitorLog(request *DescribeSiteMonitorLogRequest) (_result *DescribeSiteMonitorLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteMonitorLogResponse{}
	_body, _err := client.DescribeSiteMonitorLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the DescribeSiteMonitorQuota operation to query the quota and version of site monitoring.
//
// @param request - DescribeSiteMonitorQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorQuotaResponse
func (client *Client) DescribeSiteMonitorQuotaWithOptions(request *DescribeSiteMonitorQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorQuota"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeSiteMonitorQuota operation to query the quota and version of site monitoring.
//
// @param request - DescribeSiteMonitorQuotaRequest
//
// @return DescribeSiteMonitorQuotaResponse
func (client *Client) DescribeSiteMonitorQuota(request *DescribeSiteMonitorQuotaRequest) (_result *DescribeSiteMonitorQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteMonitorQuotaResponse{}
	_body, _err := client.DescribeSiteMonitorQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the average statistics of a specified metric in a specified site monitoring task by calling the DescribeSiteMonitorStatistics operation.
//
// Description:
//
// This topic provides an example on how to query the average statistics of the `Availability` metric in the site monitoring task whose ID is `ef4cdc8b-9dc7-43e7-810e-f950e56c****`. The returned result shows that the availability is `100`.
//
// @param request - DescribeSiteMonitorStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSiteMonitorStatisticsResponse
func (client *Client) DescribeSiteMonitorStatisticsWithOptions(request *DescribeSiteMonitorStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSiteMonitorStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSiteMonitorStatistics"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSiteMonitorStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the average statistics of a specified metric in a specified site monitoring task by calling the DescribeSiteMonitorStatistics operation.
//
// Description:
//
// This topic provides an example on how to query the average statistics of the `Availability` metric in the site monitoring task whose ID is `ef4cdc8b-9dc7-43e7-810e-f950e56c****`. The returned result shows that the availability is `100`.
//
// @param request - DescribeSiteMonitorStatisticsRequest
//
// @return DescribeSiteMonitorStatisticsResponse
func (client *Client) DescribeSiteMonitorStatistics(request *DescribeSiteMonitorStatisticsRequest) (_result *DescribeSiteMonitorStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSiteMonitorStatisticsResponse{}
	_body, _err := client.DescribeSiteMonitorStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of detection points.
//
// Description:
//
// This topic provides an example of how to call the DescribeSyntheticProbeList operation to query detection point details for the China Unicom carrier in Beijing.
//
// @param request - DescribeSyntheticProbeListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSyntheticProbeListResponse
func (client *Client) DescribeSyntheticProbeListWithOptions(request *DescribeSyntheticProbeListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSyntheticProbeListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.IdcProbe) {
		query["IdcProbe"] = request.IdcProbe
	}

	if !dara.IsNil(request.Ipv4) {
		query["Ipv4"] = request.Ipv4
	}

	if !dara.IsNil(request.Ipv6) {
		query["Ipv6"] = request.Ipv6
	}

	if !dara.IsNil(request.Isp) {
		query["Isp"] = request.Isp
	}

	if !dara.IsNil(request.LmProbe) {
		query["LmProbe"] = request.LmProbe
	}

	if !dara.IsNil(request.MbProbe) {
		query["MbProbe"] = request.MbProbe
	}

	if !dara.IsNil(request.ViewAll) {
		query["ViewAll"] = request.ViewAll
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSyntheticProbeList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSyntheticProbeListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of detection points.
//
// Description:
//
// This topic provides an example of how to call the DescribeSyntheticProbeList operation to query detection point details for the China Unicom carrier in Beijing.
//
// @param request - DescribeSyntheticProbeListRequest
//
// @return DescribeSyntheticProbeListResponse
func (client *Client) DescribeSyntheticProbeList(request *DescribeSyntheticProbeListRequest) (_result *DescribeSyntheticProbeListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSyntheticProbeListResponse{}
	_body, _err := client.DescribeSyntheticProbeListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a system event.
//
// @param request - DescribeSystemEventAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemEventAttributeResponse
func (client *Client) DescribeSystemEventAttributeWithOptions(request *DescribeSystemEventAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemEventAttributeResponse, _err error) {
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

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
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

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemEventAttribute"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemEventAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a system event.
//
// @param request - DescribeSystemEventAttributeRequest
//
// @return DescribeSystemEventAttributeResponse
func (client *Client) DescribeSystemEventAttribute(request *DescribeSystemEventAttributeRequest) (_result *DescribeSystemEventAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSystemEventAttributeResponse{}
	_body, _err := client.DescribeSystemEventAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of events that occurred for each Alibaba Cloud service under the current account.
//
// Description:
//
// ### Background information
//
// You can call [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) to obtain the Alibaba Cloud services and their system events supported by CloudMonitor.
//
// ### Operation description
//
// This topic provides an example of how to query the number of events that occurred for Elastic Compute Service (`ECS`) under the current account. The response shows that a total of 3 events occurred.
//
// @param request - DescribeSystemEventCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemEventCountResponse
func (client *Client) DescribeSystemEventCountWithOptions(request *DescribeSystemEventCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemEventCountResponse, _err error) {
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

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemEventCount"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemEventCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of events that occurred for each Alibaba Cloud service under the current account.
//
// Description:
//
// ### Background information
//
// You can call [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html) to obtain the Alibaba Cloud services and their system events supported by CloudMonitor.
//
// ### Operation description
//
// This topic provides an example of how to query the number of events that occurred for Elastic Compute Service (`ECS`) under the current account. The response shows that a total of 3 events occurred.
//
// @param request - DescribeSystemEventCountRequest
//
// @return DescribeSystemEventCountResponse
func (client *Client) DescribeSystemEventCount(request *DescribeSystemEventCountRequest) (_result *DescribeSystemEventCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSystemEventCountResponse{}
	_body, _err := client.DescribeSystemEventCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of times a system event occurred during each interval within a period of time.
//
// @param request - DescribeSystemEventHistogramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemEventHistogramResponse
func (client *Client) DescribeSystemEventHistogramWithOptions(request *DescribeSystemEventHistogramRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemEventHistogramResponse, _err error) {
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

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.SearchKeywords) {
		query["SearchKeywords"] = request.SearchKeywords
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemEventHistogram"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemEventHistogramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of times a system event occurred during each interval within a period of time.
//
// @param request - DescribeSystemEventHistogramRequest
//
// @return DescribeSystemEventHistogramResponse
func (client *Client) DescribeSystemEventHistogram(request *DescribeSystemEventHistogramRequest) (_result *DescribeSystemEventHistogramResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSystemEventHistogramResponse{}
	_body, _err := client.DescribeSystemEventHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the meta information of system events by calling the DescribeSystemEventMetaList operation.
//
// @param request - DescribeSystemEventMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemEventMetaListResponse
func (client *Client) DescribeSystemEventMetaListWithOptions(request *DescribeSystemEventMetaListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemEventMetaListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemEventMetaList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemEventMetaListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the meta information of system events by calling the DescribeSystemEventMetaList operation.
//
// @param request - DescribeSystemEventMetaListRequest
//
// @return DescribeSystemEventMetaListResponse
func (client *Client) DescribeSystemEventMetaList(request *DescribeSystemEventMetaListRequest) (_result *DescribeSystemEventMetaListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSystemEventMetaListResponse{}
	_body, _err := client.DescribeSystemEventMetaListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tag keys by calling the DescribeTagKeyList operation.
//
// @param request - DescribeTagKeyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagKeyListResponse
func (client *Client) DescribeTagKeyListWithOptions(request *DescribeTagKeyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagKeyListResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagKeyList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagKeyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tag keys by calling the DescribeTagKeyList operation.
//
// @param request - DescribeTagKeyListRequest
//
// @return DescribeTagKeyListResponse
func (client *Client) DescribeTagKeyList(request *DescribeTagKeyListRequest) (_result *DescribeTagKeyListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagKeyListResponse{}
	_body, _err := client.DescribeTagKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag values for a specified tag key.
//
// @param request - DescribeTagValueListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagValueListResponse
func (client *Client) DescribeTagValueListWithOptions(request *DescribeTagValueListRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagValueListResponse, _err error) {
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

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagValueList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagValueListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values for a specified tag key.
//
// @param request - DescribeTagValueListRequest
//
// @return DescribeTagValueListResponse
func (client *Client) DescribeTagValueList(request *DescribeTagValueListRequest) (_result *DescribeTagValueListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagValueListResponse{}
	_body, _err := client.DescribeTagValueListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DescribeUnhealthyHostAvailability operation to query a list of unhealthy servers.
//
// @param request - DescribeUnhealthyHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUnhealthyHostAvailabilityResponse
func (client *Client) DescribeUnhealthyHostAvailabilityWithOptions(request *DescribeUnhealthyHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *DescribeUnhealthyHostAvailabilityResponse, _err error) {
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
		Action:      dara.String("DescribeUnhealthyHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUnhealthyHostAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeUnhealthyHostAvailability operation to query a list of unhealthy servers.
//
// @param request - DescribeUnhealthyHostAvailabilityRequest
//
// @return DescribeUnhealthyHostAvailabilityResponse
func (client *Client) DescribeUnhealthyHostAvailability(request *DescribeUnhealthyHostAvailabilityRequest) (_result *DescribeUnhealthyHostAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUnhealthyHostAvailabilityResponse{}
	_body, _err := client.DescribeUnhealthyHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the initiative alert feature for a cloud service.
//
// @param request - DisableActiveMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableActiveMetricRuleResponse
func (client *Client) DisableActiveMetricRuleWithOptions(request *DisableActiveMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *DisableActiveMetricRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableActiveMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableActiveMetricRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the initiative alert feature for a cloud service.
//
// @param request - DisableActiveMetricRuleRequest
//
// @return DisableActiveMetricRuleResponse
func (client *Client) DisableActiveMetricRule(request *DisableActiveMetricRuleRequest) (_result *DisableActiveMetricRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableActiveMetricRuleResponse{}
	_body, _err := client.DisableActiveMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables one or more event-triggered alert rules.
//
// @param request - DisableEventRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableEventRulesResponse
func (client *Client) DisableEventRulesWithOptions(request *DisableEventRulesRequest, runtime *dara.RuntimeOptions) (_result *DisableEventRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleNames) {
		query["RuleNames"] = request.RuleNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableEventRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableEventRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables one or more event-triggered alert rules.
//
// @param request - DisableEventRulesRequest
//
// @return DisableEventRulesResponse
func (client *Client) DisableEventRules(request *DisableEventRulesRequest) (_result *DisableEventRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableEventRulesResponse{}
	_body, _err := client.DisableEventRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables one or more specified availability monitoring jobs.
//
// @param request - DisableHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableHostAvailabilityResponse
func (client *Client) DisableHostAvailabilityWithOptions(request *DisableHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *DisableHostAvailabilityResponse, _err error) {
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
		Action:      dara.String("DisableHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableHostAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables one or more specified availability monitoring jobs.
//
// @param request - DisableHostAvailabilityRequest
//
// @return DisableHostAvailabilityResponse
func (client *Client) DisableHostAvailability(request *DisableHostAvailabilityRequest) (_result *DisableHostAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableHostAvailabilityResponse{}
	_body, _err := client.DisableHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DisableMetricRules operation to disable alert rules.
//
// @param request - DisableMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableMetricRulesResponse
func (client *Client) DisableMetricRulesWithOptions(request *DisableMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *DisableMetricRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableMetricRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DisableMetricRules operation to disable alert rules.
//
// @param request - DisableMetricRulesRequest
//
// @return DisableMetricRulesResponse
func (client *Client) DisableMetricRules(request *DisableMetricRulesRequest) (_result *DisableMetricRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableMetricRulesResponse{}
	_body, _err := client.DisableMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables one or more site monitoring tasks by calling the DisableSiteMonitors operation.
//
// @param request - DisableSiteMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSiteMonitorsResponse
func (client *Client) DisableSiteMonitorsWithOptions(request *DisableSiteMonitorsRequest, runtime *dara.RuntimeOptions) (_result *DisableSiteMonitorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSiteMonitors"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSiteMonitorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables one or more site monitoring tasks by calling the DisableSiteMonitors operation.
//
// @param request - DisableSiteMonitorsRequest
//
// @return DisableSiteMonitorsResponse
func (client *Client) DisableSiteMonitors(request *DisableSiteMonitorsRequest) (_result *DisableSiteMonitorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableSiteMonitorsResponse{}
	_body, _err := client.DisableSiteMonitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the initiative alert feature for a cloud service.
//
// @param request - EnableActiveMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableActiveMetricRuleResponse
func (client *Client) EnableActiveMetricRuleWithOptions(request *EnableActiveMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *EnableActiveMetricRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableActiveMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableActiveMetricRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the initiative alert feature for a cloud service.
//
// @param request - EnableActiveMetricRuleRequest
//
// @return EnableActiveMetricRuleResponse
func (client *Client) EnableActiveMetricRule(request *EnableActiveMetricRuleRequest) (_result *EnableActiveMetricRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableActiveMetricRuleResponse{}
	_body, _err := client.EnableActiveMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The EnableEventRules operation enables one or more event rules.
//
// @param request - EnableEventRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableEventRulesResponse
func (client *Client) EnableEventRulesWithOptions(request *EnableEventRulesRequest, runtime *dara.RuntimeOptions) (_result *EnableEventRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleNames) {
		query["RuleNames"] = request.RuleNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableEventRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableEventRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The EnableEventRules operation enables one or more event rules.
//
// @param request - EnableEventRulesRequest
//
// @return EnableEventRulesResponse
func (client *Client) EnableEventRules(request *EnableEventRulesRequest) (_result *EnableEventRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableEventRulesResponse{}
	_body, _err := client.EnableEventRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a specified availability monitoring task.
//
// @param request - EnableHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHostAvailabilityResponse
func (client *Client) EnableHostAvailabilityWithOptions(request *EnableHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *EnableHostAvailabilityResponse, _err error) {
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
		Action:      dara.String("EnableHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableHostAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a specified availability monitoring task.
//
// @param request - EnableHostAvailabilityRequest
//
// @return EnableHostAvailabilityResponse
func (client *Client) EnableHostAvailability(request *EnableHostAvailabilityRequest) (_result *EnableHostAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableHostAvailabilityResponse{}
	_body, _err := client.EnableHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables multiple blacklist policies at a time.
//
// @param request - EnableMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableMetricRuleBlackListResponse
func (client *Client) EnableMetricRuleBlackListWithOptions(request *EnableMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *EnableMetricRuleBlackListResponse, _err error) {
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

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableMetricRuleBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables multiple blacklist policies at a time.
//
// @param request - EnableMetricRuleBlackListRequest
//
// @return EnableMetricRuleBlackListResponse
func (client *Client) EnableMetricRuleBlackList(request *EnableMetricRuleBlackListRequest) (_result *EnableMetricRuleBlackListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableMetricRuleBlackListResponse{}
	_body, _err := client.EnableMetricRuleBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables alert rules.
//
// @param request - EnableMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableMetricRulesResponse
func (client *Client) EnableMetricRulesWithOptions(request *EnableMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *EnableMetricRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableMetricRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables alert rules.
//
// @param request - EnableMetricRulesRequest
//
// @return EnableMetricRulesResponse
func (client *Client) EnableMetricRules(request *EnableMetricRulesRequest) (_result *EnableMetricRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableMetricRulesResponse{}
	_body, _err := client.EnableMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables site monitoring tasks.
//
// @param request - EnableSiteMonitorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSiteMonitorsResponse
func (client *Client) EnableSiteMonitorsWithOptions(request *EnableSiteMonitorsRequest, runtime *dara.RuntimeOptions) (_result *EnableSiteMonitorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskIds) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSiteMonitors"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSiteMonitorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables site monitoring tasks.
//
// @param request - EnableSiteMonitorsRequest
//
// @return EnableSiteMonitorsResponse
func (client *Client) EnableSiteMonitors(request *EnableSiteMonitorsRequest) (_result *EnableSiteMonitorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableSiteMonitorsResponse{}
	_body, _err := client.EnableSiteMonitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the InstallMonitoringAgent operation to install the CloudMonitor agent on specified Alibaba Cloud hosts.
//
// Description:
//
// ## Prerequisites
//
// Ensure that Cloud Assistant is installed on the Alibaba Cloud host. For more information, see [Cloud Assistant overview](https://help.aliyun.com/document_detail/64601.html).
//
// ## Background information
//
// This API applies only to Alibaba Cloud hosts, which are Elastic Compute Service (ECS) instances. The success rate for installing the CloudMonitor agent using this API is approximately 95%. If the installation fails, you must manually install the agent. For more information about how to install the CloudMonitor agent, see [Install and uninstall the C++ agent](https://help.aliyun.com/document_detail/183482.html).
//
// ## Example
//
// This example shows how to forcibly install the latest version of the CloudMonitor agent on the ECS instance `i-m5e0k0bexac8tykr****`.
//
// @param request - InstallMonitoringAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallMonitoringAgentResponse
func (client *Client) InstallMonitoringAgentWithOptions(request *InstallMonitoringAgentRequest, runtime *dara.RuntimeOptions) (_result *InstallMonitoringAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstallCommand) {
		query["InstallCommand"] = request.InstallCommand
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallMonitoringAgent"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallMonitoringAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the InstallMonitoringAgent operation to install the CloudMonitor agent on specified Alibaba Cloud hosts.
//
// Description:
//
// ## Prerequisites
//
// Ensure that Cloud Assistant is installed on the Alibaba Cloud host. For more information, see [Cloud Assistant overview](https://help.aliyun.com/document_detail/64601.html).
//
// ## Background information
//
// This API applies only to Alibaba Cloud hosts, which are Elastic Compute Service (ECS) instances. The success rate for installing the CloudMonitor agent using this API is approximately 95%. If the installation fails, you must manually install the agent. For more information about how to install the CloudMonitor agent, see [Install and uninstall the C++ agent](https://help.aliyun.com/document_detail/183482.html).
//
// ## Example
//
// This example shows how to forcibly install the latest version of the CloudMonitor agent on the ECS instance `i-m5e0k0bexac8tykr****`.
//
// @param request - InstallMonitoringAgentRequest
//
// @return InstallMonitoringAgentResponse
func (client *Client) InstallMonitoringAgent(request *InstallMonitoringAgentRequest) (_result *InstallMonitoringAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstallMonitoringAgentResponse{}
	_body, _err := client.InstallMonitoringAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ModifyGroupMonitoringAgentProcess operation to modify the process monitoring settings for an application group.
//
// @param request - ModifyGroupMonitoringAgentProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGroupMonitoringAgentProcessResponse
func (client *Client) ModifyGroupMonitoringAgentProcessWithOptions(request *ModifyGroupMonitoringAgentProcessRequest, runtime *dara.RuntimeOptions) (_result *ModifyGroupMonitoringAgentProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MatchExpressFilterRelation) {
		query["MatchExpressFilterRelation"] = request.MatchExpressFilterRelation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyGroupMonitoringAgentProcess"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGroupMonitoringAgentProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyGroupMonitoringAgentProcess operation to modify the process monitoring settings for an application group.
//
// @param request - ModifyGroupMonitoringAgentProcessRequest
//
// @return ModifyGroupMonitoringAgentProcessResponse
func (client *Client) ModifyGroupMonitoringAgentProcess(request *ModifyGroupMonitoringAgentProcessRequest) (_result *ModifyGroupMonitoringAgentProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGroupMonitoringAgentProcessResponse{}
	_body, _err := client.ModifyGroupMonitoringAgentProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an availability monitoring task.
//
// Description:
//
// This topic provides an example on how to change the name of an availability monitoring task named `12345` in an application group named `123456` to `task2`.
//
// @param request - ModifyHostAvailabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostAvailabilityResponse
func (client *Client) ModifyHostAvailabilityWithOptions(request *ModifyHostAvailabilityRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostAvailabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertConfigEscalationList) {
		query["AlertConfigEscalationList"] = request.AlertConfigEscalationList
	}

	if !dara.IsNil(request.AlertConfigTargetList) {
		query["AlertConfigTargetList"] = request.AlertConfigTargetList
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceList) {
		query["InstanceList"] = request.InstanceList
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskScope) {
		query["TaskScope"] = request.TaskScope
	}

	if !dara.IsNil(request.AlertConfig) {
		query["AlertConfig"] = request.AlertConfig
	}

	if !dara.IsNil(request.TaskOption) {
		query["TaskOption"] = request.TaskOption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostAvailability"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostAvailabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an availability monitoring task.
//
// Description:
//
// This topic provides an example on how to change the name of an availability monitoring task named `12345` in an application group named `123456` to `task2`.
//
// @param request - ModifyHostAvailabilityRequest
//
// @return ModifyHostAvailabilityResponse
func (client *Client) ModifyHostAvailability(request *ModifyHostAvailabilityRequest) (_result *ModifyHostAvailabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostAvailabilityResponse{}
	_body, _err := client.ModifyHostAvailabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the ModifyHostInfo operation to modify the display information of a non-Alibaba Cloud host.
//
// Description:
//
// ***
//
// @param request - ModifyHostInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostInfoResponse
func (client *Client) ModifyHostInfoWithOptions(request *ModifyHostInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyHostInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostInfo"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyHostInfo operation to modify the display information of a non-Alibaba Cloud host.
//
// Description:
//
// ***
//
// @param request - ModifyHostInfoRequest
//
// @return ModifyHostInfoResponse
func (client *Client) ModifyHostInfo(request *ModifyHostInfoRequest) (_result *ModifyHostInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHostInfoResponse{}
	_body, _err := client.ModifyHostInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the ModifyHybridMonitorNamespace operation to modify a metric store.
//
// Description:
//
// This topic provides an example on how to change the data storage duration of the metric store `aliyun` to `cms.s1.2xlarge`. The response shows that the metric store is modified.
//
// @param request - ModifyHybridMonitorNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridMonitorNamespaceResponse
func (client *Client) ModifyHybridMonitorNamespaceWithOptions(request *ModifyHybridMonitorNamespaceRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridMonitorNamespaceResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridMonitorNamespace"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridMonitorNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyHybridMonitorNamespace operation to modify a metric store.
//
// Description:
//
// This topic provides an example on how to change the data storage duration of the metric store `aliyun` to `cms.s1.2xlarge`. The response shows that the metric store is modified.
//
// @param request - ModifyHybridMonitorNamespaceRequest
//
// @return ModifyHybridMonitorNamespaceResponse
func (client *Client) ModifyHybridMonitorNamespace(request *ModifyHybridMonitorNamespaceRequest) (_result *ModifyHybridMonitorNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridMonitorNamespaceResponse{}
	_body, _err := client.ModifyHybridMonitorNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a Logstore group by calling the ModifyHybridMonitorSLSGroup operation.
//
// Description:
//
// This topic provides an example on how to change the Logstore in the `aliyun-project` log project in the `ap-southeast-1` region of the Logstore group `Logstore_test` to `Logstore-aliyun-all`. The response shows that the Logstore group is modified.
//
// @param request - ModifyHybridMonitorSLSGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridMonitorSLSGroupResponse
func (client *Client) ModifyHybridMonitorSLSGroupWithOptions(request *ModifyHybridMonitorSLSGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridMonitorSLSGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SLSGroupConfig) {
		query["SLSGroupConfig"] = request.SLSGroupConfig
	}

	if !dara.IsNil(request.SLSGroupDescription) {
		query["SLSGroupDescription"] = request.SLSGroupDescription
	}

	if !dara.IsNil(request.SLSGroupName) {
		query["SLSGroupName"] = request.SLSGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridMonitorSLSGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridMonitorSLSGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a Logstore group by calling the ModifyHybridMonitorSLSGroup operation.
//
// Description:
//
// This topic provides an example on how to change the Logstore in the `aliyun-project` log project in the `ap-southeast-1` region of the Logstore group `Logstore_test` to `Logstore-aliyun-all`. The response shows that the Logstore group is modified.
//
// @param request - ModifyHybridMonitorSLSGroupRequest
//
// @return ModifyHybridMonitorSLSGroupResponse
func (client *Client) ModifyHybridMonitorSLSGroup(request *ModifyHybridMonitorSLSGroupRequest) (_result *ModifyHybridMonitorSLSGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridMonitorSLSGroupResponse{}
	_body, _err := client.ModifyHybridMonitorSLSGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the metrics of a Simple Log Service (SLS) log monitoring task by calling the ModifyHybridMonitorTask operation.
//
// Description:
//
// This topic provides an example on how to change the collection interval of the SLS log monitoring task `36****` to `15` seconds. The response shows that the metric is modified.
//
// @param request - ModifyHybridMonitorTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridMonitorTaskResponse
func (client *Client) ModifyHybridMonitorTaskWithOptions(request *ModifyHybridMonitorTaskRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridMonitorTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachLabels) {
		query["AttachLabels"] = request.AttachLabels
	}

	if !dara.IsNil(request.CollectInterval) {
		query["CollectInterval"] = request.CollectInterval
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SLSProcessConfig) {
		query["SLSProcessConfig"] = request.SLSProcessConfig
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridMonitorTask"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridMonitorTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the metrics of a Simple Log Service (SLS) log monitoring task by calling the ModifyHybridMonitorTask operation.
//
// Description:
//
// This topic provides an example on how to change the collection interval of the SLS log monitoring task `36****` to `15` seconds. The response shows that the metric is modified.
//
// @param request - ModifyHybridMonitorTaskRequest
//
// @return ModifyHybridMonitorTaskResponse
func (client *Client) ModifyHybridMonitorTask(request *ModifyHybridMonitorTaskRequest) (_result *ModifyHybridMonitorTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridMonitorTaskResponse{}
	_body, _err := client.ModifyHybridMonitorTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a blacklist policy.
//
// @param request - ModifyMetricRuleBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMetricRuleBlackListResponse
func (client *Client) ModifyMetricRuleBlackListWithOptions(request *ModifyMetricRuleBlackListRequest, runtime *dara.RuntimeOptions) (_result *ModifyMetricRuleBlackListResponse, _err error) {
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

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EnableEndTime) {
		query["EnableEndTime"] = request.EnableEndTime
	}

	if !dara.IsNil(request.EnableStartTime) {
		query["EnableStartTime"] = request.EnableStartTime
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	if !dara.IsNil(request.Metrics) {
		query["Metrics"] = request.Metrics
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ScopeType) {
		query["ScopeType"] = request.ScopeType
	}

	if !dara.IsNil(request.ScopeValue) {
		query["ScopeValue"] = request.ScopeValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMetricRuleBlackList"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMetricRuleBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a blacklist policy.
//
// @param request - ModifyMetricRuleBlackListRequest
//
// @return ModifyMetricRuleBlackListResponse
func (client *Client) ModifyMetricRuleBlackList(request *ModifyMetricRuleBlackListRequest) (_result *ModifyMetricRuleBlackListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMetricRuleBlackListResponse{}
	_body, _err := client.ModifyMetricRuleBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an alert template.
//
// Description:
//
// This topic provides an example on how to modify an alert template whose version is `1` and ID is `123456`. The alert level is changed to `Critical`. The statistical method is changed to `Average`. The alert threshold comparator is changed to `GreaterThanOrEqualToThreshold`. The alert threshold is changed to `90`. The number of alert retries is changed to `3`. The response shows that the alert template is modified.
//
// @param request - ModifyMetricRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMetricRuleTemplateResponse
func (client *Client) ModifyMetricRuleTemplateWithOptions(request *ModifyMetricRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyMetricRuleTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertTemplates) {
		query["AlertTemplates"] = request.AlertTemplates
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RestVersion) {
		query["RestVersion"] = request.RestVersion
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMetricRuleTemplate"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMetricRuleTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an alert template.
//
// Description:
//
// This topic provides an example on how to modify an alert template whose version is `1` and ID is `123456`. The alert level is changed to `Critical`. The statistical method is changed to `Average`. The alert threshold comparator is changed to `GreaterThanOrEqualToThreshold`. The alert threshold is changed to `90`. The number of alert retries is changed to `3`. The response shows that the alert template is modified.
//
// @param request - ModifyMetricRuleTemplateRequest
//
// @return ModifyMetricRuleTemplateResponse
func (client *Client) ModifyMetricRuleTemplate(request *ModifyMetricRuleTemplateRequest) (_result *ModifyMetricRuleTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMetricRuleTemplateResponse{}
	_body, _err := client.ModifyMetricRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ModifyMonitorGroup operation to modify an application group.
//
// @param request - ModifyMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMonitorGroupResponse
func (client *Client) ModifyMonitorGroupWithOptions(request *ModifyMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyMonitorGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMonitorGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMonitorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyMonitorGroup operation to modify an application group.
//
// @param request - ModifyMonitorGroupRequest
//
// @return ModifyMonitorGroupResponse
func (client *Client) ModifyMonitorGroup(request *ModifyMonitorGroupRequest) (_result *ModifyMonitorGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMonitorGroupResponse{}
	_body, _err := client.ModifyMonitorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ModifyMonitorGroupInstances operation modifies the resources in an application group.
//
// @param request - ModifyMonitorGroupInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMonitorGroupInstancesResponse
func (client *Client) ModifyMonitorGroupInstancesWithOptions(request *ModifyMonitorGroupInstancesRequest, runtime *dara.RuntimeOptions) (_result *ModifyMonitorGroupInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMonitorGroupInstances"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMonitorGroupInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ModifyMonitorGroupInstances operation modifies the resources in an application group.
//
// @param request - ModifyMonitorGroupInstancesRequest
//
// @return ModifyMonitorGroupInstancesResponse
func (client *Client) ModifyMonitorGroupInstances(request *ModifyMonitorGroupInstancesRequest) (_result *ModifyMonitorGroupInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMonitorGroupInstancesResponse{}
	_body, _err := client.ModifyMonitorGroupInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a site monitoring task by calling the ModifySiteMonitor operation.
//
// @param request - ModifySiteMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySiteMonitorResponse
func (client *Client) ModifySiteMonitorWithOptions(request *ModifySiteMonitorRequest, runtime *dara.RuntimeOptions) (_result *ModifySiteMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AlertIds) {
		query["AlertIds"] = request.AlertIds
	}

	if !dara.IsNil(request.CustomSchedule) {
		query["CustomSchedule"] = request.CustomSchedule
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IntervalUnit) {
		query["IntervalUnit"] = request.IntervalUnit
	}

	if !dara.IsNil(request.IspCities) {
		query["IspCities"] = request.IspCities
	}

	if !dara.IsNil(request.OptionsJson) {
		query["OptionsJson"] = request.OptionsJson
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySiteMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySiteMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a site monitoring task by calling the ModifySiteMonitor operation.
//
// @param request - ModifySiteMonitorRequest
//
// @return ModifySiteMonitorResponse
func (client *Client) ModifySiteMonitor(request *ModifySiteMonitorRequest) (_result *ModifySiteMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySiteMonitorResponse{}
	_body, _err := client.ModifySiteMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert contact by calling the PutContact operation.
//
// @param request - PutContactRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutContactResponse
func (client *Client) PutContactWithOptions(request *PutContactRequest, runtime *dara.RuntimeOptions) (_result *PutContactResponse, _err error) {
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

	if !dara.IsNil(request.Describe) {
		query["Describe"] = request.Describe
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Channels) {
		query["Channels"] = request.Channels
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutContact"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert contact by calling the PutContact operation.
//
// @param request - PutContactRequest
//
// @return PutContactResponse
func (client *Client) PutContact(request *PutContactRequest) (_result *PutContactResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutContactResponse{}
	_body, _err := client.PutContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert contact group by calling the PutContactGroup operation.
//
// @param request - PutContactGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutContactGroupResponse
func (client *Client) PutContactGroupWithOptions(request *PutContactGroupRequest, runtime *dara.RuntimeOptions) (_result *PutContactGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !dara.IsNil(request.ContactNames) {
		query["ContactNames"] = request.ContactNames
	}

	if !dara.IsNil(request.Describe) {
		query["Describe"] = request.Describe
	}

	if !dara.IsNil(request.EnableSubscribed) {
		query["EnableSubscribed"] = request.EnableSubscribed
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutContactGroup"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert contact group by calling the PutContactGroup operation.
//
// @param request - PutContactGroupRequest
//
// @return PutContactGroupResponse
func (client *Client) PutContactGroup(request *PutContactGroupRequest) (_result *PutContactGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutContactGroupResponse{}
	_body, _err := client.PutContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reports custom events.
//
// @param request - PutCustomEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCustomEventResponse
func (client *Client) PutCustomEventWithOptions(request *PutCustomEventRequest, runtime *dara.RuntimeOptions) (_result *PutCustomEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventInfo) {
		query["EventInfo"] = request.EventInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutCustomEvent"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutCustomEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports custom events.
//
// @param request - PutCustomEventRequest
//
// @return PutCustomEventResponse
func (client *Client) PutCustomEvent(request *PutCustomEventRequest) (_result *PutCustomEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutCustomEventResponse{}
	_body, _err := client.PutCustomEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the PutCustomEventRule operation to create an alert rule for a custom event.
//
// Description:
//
// Before you call this operation, you must call the PutCustomEvent operation to report the monitoring data of the custom event. For more information, see [PutCustomEvent](https://help.aliyun.com/document_detail/115012.html).
//
// @param request - PutCustomEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCustomEventRuleResponse
func (client *Client) PutCustomEventRuleWithOptions(request *PutCustomEventRuleRequest, runtime *dara.RuntimeOptions) (_result *PutCustomEventRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.EffectiveInterval) {
		query["EffectiveInterval"] = request.EffectiveInterval
	}

	if !dara.IsNil(request.EmailSubject) {
		query["EmailSubject"] = request.EmailSubject
	}

	if !dara.IsNil(request.EventName) {
		query["EventName"] = request.EventName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutCustomEventRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutCustomEventRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the PutCustomEventRule operation to create an alert rule for a custom event.
//
// Description:
//
// Before you call this operation, you must call the PutCustomEvent operation to report the monitoring data of the custom event. For more information, see [PutCustomEvent](https://help.aliyun.com/document_detail/115012.html).
//
// @param request - PutCustomEventRuleRequest
//
// @return PutCustomEventRuleResponse
func (client *Client) PutCustomEventRule(request *PutCustomEventRuleRequest) (_result *PutCustomEventRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutCustomEventRuleResponse{}
	_body, _err := client.PutCustomEventRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the PutCustomMetric operation to report custom monitoring data.
//
// Description:
//
// > We recommend that you use the [PutHybridMonitorMetricData](https://help.aliyun.com/document_detail/383455.html) operation in Enterprise CloudMonitor.
//
// @param request - PutCustomMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCustomMetricResponse
func (client *Client) PutCustomMetricWithOptions(request *PutCustomMetricRequest, runtime *dara.RuntimeOptions) (_result *PutCustomMetricResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MetricList) {
		query["MetricList"] = request.MetricList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutCustomMetric"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutCustomMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the PutCustomMetric operation to report custom monitoring data.
//
// Description:
//
// > We recommend that you use the [PutHybridMonitorMetricData](https://help.aliyun.com/document_detail/383455.html) operation in Enterprise CloudMonitor.
//
// @param request - PutCustomMetricRequest
//
// @return PutCustomMetricResponse
func (client *Client) PutCustomMetric(request *PutCustomMetricRequest) (_result *PutCustomMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutCustomMetricResponse{}
	_body, _err := client.PutCustomMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom monitoring alert rule by calling the PutCustomMetricRule operation.
//
// Description:
//
// 调用本接口前，请先调用PutCustomMetric接口上报自定义监控数据，详情请参见 [PutCustomMetric](https://help.aliyun.com/document_detail/115004.html)。
//
// @param request - PutCustomMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutCustomMetricRuleResponse
func (client *Client) PutCustomMetricRuleWithOptions(request *PutCustomMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *PutCustomMetricRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComparisonOperator) {
		query["ComparisonOperator"] = request.ComparisonOperator
	}

	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.EffectiveInterval) {
		query["EffectiveInterval"] = request.EffectiveInterval
	}

	if !dara.IsNil(request.EmailSubject) {
		query["EmailSubject"] = request.EmailSubject
	}

	if !dara.IsNil(request.EvaluationCount) {
		query["EvaluationCount"] = request.EvaluationCount
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.Statistics) {
		query["Statistics"] = request.Statistics
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutCustomMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutCustomMetricRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom monitoring alert rule by calling the PutCustomMetricRule operation.
//
// Description:
//
// 调用本接口前，请先调用PutCustomMetric接口上报自定义监控数据，详情请参见 [PutCustomMetric](https://help.aliyun.com/document_detail/115004.html)。
//
// @param request - PutCustomMetricRuleRequest
//
// @return PutCustomMetricRuleResponse
func (client *Client) PutCustomMetricRule(request *PutCustomMetricRuleRequest) (_result *PutCustomMetricRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutCustomMetricRuleResponse{}
	_body, _err := client.PutCustomMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies an event-based alert rule by calling the PutEventRule operation.
//
// Description:
//
// If the Event-triggered Alert Rule name does not exist, a new alert rule is created. If the Event-triggered Alert Rule name already exists, the existing alert rule is modified.
//
// This topic provides an example of how to create an event-based alert rule named `myRuleName` for the Alibaba Cloud service `ecs`.
//
// @param request - PutEventRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEventRuleResponse
func (client *Client) PutEventRuleWithOptions(request *PutEventRuleRequest, runtime *dara.RuntimeOptions) (_result *PutEventRuleResponse, _err error) {
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

	if !dara.IsNil(request.EventPattern) {
		query["EventPattern"] = request.EventPattern
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutEventRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutEventRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an event-based alert rule by calling the PutEventRule operation.
//
// Description:
//
// If the Event-triggered Alert Rule name does not exist, a new alert rule is created. If the Event-triggered Alert Rule name already exists, the existing alert rule is modified.
//
// This topic provides an example of how to create an event-based alert rule named `myRuleName` for the Alibaba Cloud service `ecs`.
//
// @param request - PutEventRuleRequest
//
// @return PutEventRuleResponse
func (client *Client) PutEventRule(request *PutEventRuleRequest) (_result *PutEventRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutEventRuleResponse{}
	_body, _err := client.PutEventRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or modifies the notification channels of an event-triggered alert rule by calling the PutEventRuleTargets operation.
//
// @param request - PutEventRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutEventRuleTargetsResponse
func (client *Client) PutEventRuleTargetsWithOptions(request *PutEventRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *PutEventRuleTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContactParameters) {
		query["ContactParameters"] = request.ContactParameters
	}

	if !dara.IsNil(request.FcParameters) {
		query["FcParameters"] = request.FcParameters
	}

	if !dara.IsNil(request.MnsParameters) {
		query["MnsParameters"] = request.MnsParameters
	}

	if !dara.IsNil(request.OpenApiParameters) {
		query["OpenApiParameters"] = request.OpenApiParameters
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SlsParameters) {
		query["SlsParameters"] = request.SlsParameters
	}

	if !dara.IsNil(request.WebhookParameters) {
		query["WebhookParameters"] = request.WebhookParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutEventRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutEventRuleTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies the notification channels of an event-triggered alert rule by calling the PutEventRuleTargets operation.
//
// @param request - PutEventRuleTargetsRequest
//
// @return PutEventRuleTargetsResponse
func (client *Client) PutEventRuleTargets(request *PutEventRuleTargetsRequest) (_result *PutEventRuleTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutEventRuleTargetsResponse{}
	_body, _err := client.PutEventRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies a monitoring data export configuration by calling the PutExporterOutput operation.
//
// Description:
//
// > Only Log Service (SLS) is supported. More products will be supported in the future.
//
// @param request - PutExporterOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutExporterOutputResponse
func (client *Client) PutExporterOutputWithOptions(request *PutExporterOutputRequest, runtime *dara.RuntimeOptions) (_result *PutExporterOutputResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigJson) {
		query["ConfigJson"] = request.ConfigJson
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.DestName) {
		query["DestName"] = request.DestName
	}

	if !dara.IsNil(request.DestType) {
		query["DestType"] = request.DestType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutExporterOutput"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutExporterOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies a monitoring data export configuration by calling the PutExporterOutput operation.
//
// Description:
//
// > Only Log Service (SLS) is supported. More products will be supported in the future.
//
// @param request - PutExporterOutputRequest
//
// @return PutExporterOutputResponse
func (client *Client) PutExporterOutput(request *PutExporterOutputRequest) (_result *PutExporterOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutExporterOutputResponse{}
	_body, _err := client.PutExporterOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Invokes the PutExporterRule operation to create or modify export rules.
//
// @param request - PutExporterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutExporterRuleResponse
func (client *Client) PutExporterRuleWithOptions(request *PutExporterRuleRequest, runtime *dara.RuntimeOptions) (_result *PutExporterRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Describe) {
		query["Describe"] = request.Describe
	}

	if !dara.IsNil(request.DstNames) {
		query["DstNames"] = request.DstNames
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.TargetWindows) {
		query["TargetWindows"] = request.TargetWindows
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutExporterRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutExporterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the PutExporterRule operation to create or modify export rules.
//
// @param request - PutExporterRuleRequest
//
// @return PutExporterRuleResponse
func (client *Client) PutExporterRule(request *PutExporterRuleRequest) (_result *PutExporterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutExporterRuleResponse{}
	_body, _err := client.PutExporterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule in a specified application group.
//
// Description:
//
// This topic provides an example on how to create an alert rule for the `cpu_total` metric of Elastic Compute Service (ECS) in the application group `17285****`. The alert rule ID is `123456`, the alert rule name is `Rule_test`, the alert severity is `Critical`, the statistical method is `Average`, the comparison operator is `GreaterThanOrEqualToThreshold`, the threshold is `90`, and the retry count is `3`. The response shows that the alert rule is created. The alert rule ID is `123456`.
//
// @param request - PutGroupMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutGroupMetricRuleResponse
func (client *Client) PutGroupMetricRuleWithOptions(request *PutGroupMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *PutGroupMetricRuleResponse, _err error) {
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

	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.Dimensions) {
		query["Dimensions"] = request.Dimensions
	}

	if !dara.IsNil(request.EffectiveInterval) {
		query["EffectiveInterval"] = request.EffectiveInterval
	}

	if !dara.IsNil(request.EmailSubject) {
		query["EmailSubject"] = request.EmailSubject
	}

	if !dara.IsNil(request.ExtraDimensionJson) {
		query["ExtraDimensionJson"] = request.ExtraDimensionJson
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NoDataPolicy) {
		query["NoDataPolicy"] = request.NoDataPolicy
	}

	if !dara.IsNil(request.NoEffectiveInterval) {
		query["NoEffectiveInterval"] = request.NoEffectiveInterval
	}

	if !dara.IsNil(request.Options) {
		query["Options"] = request.Options
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	if !dara.IsNil(request.Escalations) {
		query["Escalations"] = request.Escalations
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutGroupMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutGroupMetricRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule in a specified application group.
//
// Description:
//
// This topic provides an example on how to create an alert rule for the `cpu_total` metric of Elastic Compute Service (ECS) in the application group `17285****`. The alert rule ID is `123456`, the alert rule name is `Rule_test`, the alert severity is `Critical`, the statistical method is `Average`, the comparison operator is `GreaterThanOrEqualToThreshold`, the threshold is `90`, and the retry count is `3`. The response shows that the alert rule is created. The alert rule ID is `123456`.
//
// @param request - PutGroupMetricRuleRequest
//
// @return PutGroupMetricRuleResponse
func (client *Client) PutGroupMetricRule(request *PutGroupMetricRuleRequest) (_result *PutGroupMetricRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutGroupMetricRuleResponse{}
	_body, _err := client.PutGroupMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reports the specified monitoring data to a specified metric namespace of Enterprise CloudMonitor by calling the PutHybridMonitorMetricData operation.
//
// Description:
//
// ## Prerequisites
//
// Make sure that you have activated Enterprise CloudMonitor. For more information, see [Activate Enterprise CloudMonitor](https://help.aliyun.com/document_detail/250773.html).
//
// ## Limits
//
// The size of the monitoring data that you can report at a time cannot exceed 1 MB.
//
// ## Usage notes
//
// This topic provides an example to show how to report the monitoring data of the `CPU_Usage` metric to the `default-aliyun` metric namespace of Enterprise CloudMonitor.
//
// @param request - PutHybridMonitorMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutHybridMonitorMetricDataResponse
func (client *Client) PutHybridMonitorMetricDataWithOptions(request *PutHybridMonitorMetricDataRequest, runtime *dara.RuntimeOptions) (_result *PutHybridMonitorMetricDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MetricList) {
		query["MetricList"] = request.MetricList
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutHybridMonitorMetricData"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutHybridMonitorMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports the specified monitoring data to a specified metric namespace of Enterprise CloudMonitor by calling the PutHybridMonitorMetricData operation.
//
// Description:
//
// ## Prerequisites
//
// Make sure that you have activated Enterprise CloudMonitor. For more information, see [Activate Enterprise CloudMonitor](https://help.aliyun.com/document_detail/250773.html).
//
// ## Limits
//
// The size of the monitoring data that you can report at a time cannot exceed 1 MB.
//
// ## Usage notes
//
// This topic provides an example to show how to report the monitoring data of the `CPU_Usage` metric to the `default-aliyun` metric namespace of Enterprise CloudMonitor.
//
// @param request - PutHybridMonitorMetricDataRequest
//
// @return PutHybridMonitorMetricDataResponse
func (client *Client) PutHybridMonitorMetricData(request *PutHybridMonitorMetricDataRequest) (_result *PutHybridMonitorMetricDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutHybridMonitorMetricDataResponse{}
	_body, _err := client.PutHybridMonitorMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies a log monitoring metric by calling the PutLogMonitor operation.
//
// Description:
//
// This topic provides an example of how to create a log monitoring metric named `cpu_total`. The response shows that the log monitoring metric is created. The log monitoring metric ID is `16****`.
//
// @param request - PutLogMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutLogMonitorResponse
func (client *Client) PutLogMonitorWithOptions(request *PutLogMonitorRequest, runtime *dara.RuntimeOptions) (_result *PutLogMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Aggregates) {
		query["Aggregates"] = request.Aggregates
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Groupbys) {
		query["Groupbys"] = request.Groupbys
	}

	if !dara.IsNil(request.LogId) {
		query["LogId"] = request.LogId
	}

	if !dara.IsNil(request.MetricExpress) {
		query["MetricExpress"] = request.MetricExpress
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.SlsLogstore) {
		query["SlsLogstore"] = request.SlsLogstore
	}

	if !dara.IsNil(request.SlsProject) {
		query["SlsProject"] = request.SlsProject
	}

	if !dara.IsNil(request.SlsRegionId) {
		query["SlsRegionId"] = request.SlsRegionId
	}

	if !dara.IsNil(request.Tumblingwindows) {
		query["Tumblingwindows"] = request.Tumblingwindows
	}

	if !dara.IsNil(request.Unit) {
		query["Unit"] = request.Unit
	}

	if !dara.IsNil(request.ValueFilter) {
		query["ValueFilter"] = request.ValueFilter
	}

	if !dara.IsNil(request.ValueFilterRelation) {
		query["ValueFilterRelation"] = request.ValueFilterRelation
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutLogMonitor"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutLogMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies a log monitoring metric by calling the PutLogMonitor operation.
//
// Description:
//
// This topic provides an example of how to create a log monitoring metric named `cpu_total`. The response shows that the log monitoring metric is created. The log monitoring metric ID is `16****`.
//
// @param request - PutLogMonitorRequest
//
// @return PutLogMonitorResponse
func (client *Client) PutLogMonitor(request *PutLogMonitorRequest) (_result *PutLogMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutLogMonitorResponse{}
	_body, _err := client.PutLogMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or modifies push channels for a specified alert rule by calling the PutMetricRuleTargets operation.
//
// Description:
//
// ## Usage notes
//
// This topic provides an example on how to associate the resource `acs:mns:ap-southeast-1:120886317861****:/queues/test/message` with the alert rule `ae06917_75a8c43178ab66****`. The alert trigger target ID is `1`. The response shows that the resource is associated.
//
// @param request - PutMetricRuleTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMetricRuleTargetsResponse
func (client *Client) PutMetricRuleTargetsWithOptions(request *PutMetricRuleTargetsRequest, runtime *dara.RuntimeOptions) (_result *PutMetricRuleTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.Targets) {
		query["Targets"] = request.Targets
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMetricRuleTargets"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMetricRuleTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies push channels for a specified alert rule by calling the PutMetricRuleTargets operation.
//
// Description:
//
// ## Usage notes
//
// This topic provides an example on how to associate the resource `acs:mns:ap-southeast-1:120886317861****:/queues/test/message` with the alert rule `ae06917_75a8c43178ab66****`. The alert trigger target ID is `1`. The response shows that the resource is associated.
//
// @param request - PutMetricRuleTargetsRequest
//
// @return PutMetricRuleTargetsResponse
func (client *Client) PutMetricRuleTargets(request *PutMetricRuleTargetsRequest) (_result *PutMetricRuleTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutMetricRuleTargetsResponse{}
	_body, _err := client.PutMetricRuleTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule to dynamically add instances that meet the rule to an application group.
//
// @param request - PutMonitorGroupDynamicRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMonitorGroupDynamicRuleResponse
func (client *Client) PutMonitorGroupDynamicRuleWithOptions(request *PutMonitorGroupDynamicRuleRequest, runtime *dara.RuntimeOptions) (_result *PutMonitorGroupDynamicRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupRules) {
		query["GroupRules"] = request.GroupRules
	}

	if !dara.IsNil(request.IsAsync) {
		query["IsAsync"] = request.IsAsync
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMonitorGroupDynamicRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMonitorGroupDynamicRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or modifies an alert rule to dynamically add instances that meet the rule to an application group.
//
// @param request - PutMonitorGroupDynamicRuleRequest
//
// @return PutMonitorGroupDynamicRuleResponse
func (client *Client) PutMonitorGroupDynamicRule(request *PutMonitorGroupDynamicRuleRequest) (_result *PutMonitorGroupDynamicRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutMonitorGroupDynamicRuleResponse{}
	_body, _err := client.PutMonitorGroupDynamicRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures global settings for the CloudMonitor agent.
//
// @param request - PutMonitoringConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMonitoringConfigResponse
func (client *Client) PutMonitoringConfigWithOptions(request *PutMonitoringConfigRequest, runtime *dara.RuntimeOptions) (_result *PutMonitoringConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoInstall) {
		query["AutoInstall"] = request.AutoInstall
	}

	if !dara.IsNil(request.EnableInstallAgentNewECS) {
		query["EnableInstallAgentNewECS"] = request.EnableInstallAgentNewECS
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMonitoringConfig"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMonitoringConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures global settings for the CloudMonitor agent.
//
// @param request - PutMonitoringConfigRequest
//
// @return PutMonitoringConfigResponse
func (client *Client) PutMonitoringConfig(request *PutMonitoringConfigRequest) (_result *PutMonitoringConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutMonitoringConfigResponse{}
	_body, _err := client.PutMonitoringConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a threshold alert rule.
//
// Description:
//
// This topic provides an example on how to configure a threshold alert rule for the `cpu_total` metric of the Elastic Computing Service (ECS) instance `i-uf6j91r34rnwawoo****` in the `acs_ecs_dashboard` namespace. The alert contact group is `ECS_Group`, the alert rule name is `test123`, the alert rule ID is `a151cd6023eacee2f0978e03863cc1697c89508****`, the statistical method for the Critical level is `Average`, the comparison operator for the Critical level is `GreaterThanOrEqualToThreshold`, the threshold for the Critical level is `90`, and the retry count for the Critical level is `3`.
//
// > As of August 15, 2024, Statistics validation is increased. The statistical value must match the Statistics of the corresponding metric. For information about how to obtain the value of this parameter, see [Alibaba Cloud service monitoring metrics](https://www.alibabacloud.com/help/en/cms/support/appendix-1-metrics).
//
// @param tmpReq - PutResourceMetricRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutResourceMetricRuleResponse
func (client *Client) PutResourceMetricRuleWithOptions(tmpReq *PutResourceMetricRuleRequest, runtime *dara.RuntimeOptions) (_result *PutResourceMetricRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PutResourceMetricRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CompositeExpression) {
		request.CompositeExpressionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CompositeExpression, dara.String("CompositeExpression"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Prometheus) {
		request.PrometheusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Prometheus, dara.String("Prometheus"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CompositeExpressionShrink) {
		query["CompositeExpression"] = request.CompositeExpressionShrink
	}

	if !dara.IsNil(request.ContactGroups) {
		query["ContactGroups"] = request.ContactGroups
	}

	if !dara.IsNil(request.EffectiveInterval) {
		query["EffectiveInterval"] = request.EffectiveInterval
	}

	if !dara.IsNil(request.EmailSubject) {
		query["EmailSubject"] = request.EmailSubject
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NoDataPolicy) {
		query["NoDataPolicy"] = request.NoDataPolicy
	}

	if !dara.IsNil(request.NoEffectiveInterval) {
		query["NoEffectiveInterval"] = request.NoEffectiveInterval
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrometheusShrink) {
		query["Prometheus"] = request.PrometheusShrink
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.SendOK) {
		query["SendOK"] = request.SendOK
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.Webhook) {
		query["Webhook"] = request.Webhook
	}

	if !dara.IsNil(request.Escalations) {
		query["Escalations"] = request.Escalations
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutResourceMetricRule"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutResourceMetricRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a threshold alert rule.
//
// Description:
//
// This topic provides an example on how to configure a threshold alert rule for the `cpu_total` metric of the Elastic Computing Service (ECS) instance `i-uf6j91r34rnwawoo****` in the `acs_ecs_dashboard` namespace. The alert contact group is `ECS_Group`, the alert rule name is `test123`, the alert rule ID is `a151cd6023eacee2f0978e03863cc1697c89508****`, the statistical method for the Critical level is `Average`, the comparison operator for the Critical level is `GreaterThanOrEqualToThreshold`, the threshold for the Critical level is `90`, and the retry count for the Critical level is `3`.
//
// > As of August 15, 2024, Statistics validation is increased. The statistical value must match the Statistics of the corresponding metric. For information about how to obtain the value of this parameter, see [Alibaba Cloud service monitoring metrics](https://www.alibabacloud.com/help/en/cms/support/appendix-1-metrics).
//
// @param request - PutResourceMetricRuleRequest
//
// @return PutResourceMetricRuleResponse
func (client *Client) PutResourceMetricRule(request *PutResourceMetricRuleRequest) (_result *PutResourceMetricRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutResourceMetricRuleResponse{}
	_body, _err := client.PutResourceMetricRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates multiple threshold alert rules for a specified metric of a specified resource by calling the PutResourceMetricRules operation.
//
// Description:
//
// 本文将提供一个示例，为云服务器ECS `acs_ecs_dashboard`的实例`i-uf6j91r34rnwawoo****`中的监控项`cpu_total`设置阈值报警规则。该报警规则的报警联系组为`ECS_Group`、报警规则名称为`test123`、报警规则ID为`a151cd6023eacee2f0978e03863cc1697c89508****`、Critical级别的统计方法为`Average`、Critical级别的比较符为`GreaterThanOrEqualToThreshold`、Critical级别的阈值为`90`和Critical级别的报警重试次数为`3`。
//
// > 2024-08-15 增加Statistics校验，统计值只能填对应指标的Statistics。关于如何获取该参数的取值，请参见[云产品监控项](https://www.alibabacloud.com/help/en/cms/support/appendix-1-metrics)。
//
// @param request - PutResourceMetricRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutResourceMetricRulesResponse
func (client *Client) PutResourceMetricRulesWithOptions(request *PutResourceMetricRulesRequest, runtime *dara.RuntimeOptions) (_result *PutResourceMetricRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutResourceMetricRules"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutResourceMetricRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple threshold alert rules for a specified metric of a specified resource by calling the PutResourceMetricRules operation.
//
// Description:
//
// 本文将提供一个示例，为云服务器ECS `acs_ecs_dashboard`的实例`i-uf6j91r34rnwawoo****`中的监控项`cpu_total`设置阈值报警规则。该报警规则的报警联系组为`ECS_Group`、报警规则名称为`test123`、报警规则ID为`a151cd6023eacee2f0978e03863cc1697c89508****`、Critical级别的统计方法为`Average`、Critical级别的比较符为`GreaterThanOrEqualToThreshold`、Critical级别的阈值为`90`和Critical级别的报警重试次数为`3`。
//
// > 2024-08-15 增加Statistics校验，统计值只能填对应指标的Statistics。关于如何获取该参数的取值，请参见[云产品监控项](https://www.alibabacloud.com/help/en/cms/support/appendix-1-metrics)。
//
// @param request - PutResourceMetricRulesRequest
//
// @return PutResourceMetricRulesResponse
func (client *Client) PutResourceMetricRules(request *PutResourceMetricRulesRequest) (_result *PutResourceMetricRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutResourceMetricRulesResponse{}
	_body, _err := client.PutResourceMetricRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The RemoveTags operation removes one or more tags.
//
// @param request - RemoveTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTagsResponse
func (client *Client) RemoveTagsWithOptions(request *RemoveTagsRequest, runtime *dara.RuntimeOptions) (_result *RemoveTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupIds) {
		query["GroupIds"] = request.GroupIds
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTags"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The RemoveTags operation removes one or more tags.
//
// @param request - RemoveTagsRequest
//
// @return RemoveTagsResponse
func (client *Client) RemoveTags(request *RemoveTagsRequest) (_result *RemoveTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTagsResponse{}
	_body, _err := client.RemoveTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Debugs a system event of an Alibaba Cloud service.
//
// Description:
//
// Tests whether a system event can be triggered as expected. Simulates a system event and verifies the response when the event triggers an alert.
//
// @param request - SendDryRunSystemEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendDryRunSystemEventResponse
func (client *Client) SendDryRunSystemEventWithOptions(request *SendDryRunSystemEventRequest, runtime *dara.RuntimeOptions) (_result *SendDryRunSystemEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventContent) {
		query["EventContent"] = request.EventContent
	}

	if !dara.IsNil(request.EventName) {
		query["EventName"] = request.EventName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendDryRunSystemEvent"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendDryRunSystemEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Debugs a system event of an Alibaba Cloud service.
//
// Description:
//
// Tests whether a system event can be triggered as expected. Simulates a system event and verifies the response when the event triggers an alert.
//
// @param request - SendDryRunSystemEventRequest
//
// @return SendDryRunSystemEventResponse
func (client *Client) SendDryRunSystemEvent(request *SendDryRunSystemEventRequest) (_result *SendDryRunSystemEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendDryRunSystemEventResponse{}
	_body, _err := client.SendDryRunSystemEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls the CloudMonitor agent from a third-party host.
//
// Description:
//
// > This API operation is not applicable to Elastic Compute Service (ECS) instances. To uninstall the agent from an ECS instance, see [Install and uninstall the CloudMonitor agent](https://help.aliyun.com/document_detail/183482.html).
//
// @param request - UninstallMonitoringAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallMonitoringAgentResponse
func (client *Client) UninstallMonitoringAgentWithOptions(request *UninstallMonitoringAgentRequest, runtime *dara.RuntimeOptions) (_result *UninstallMonitoringAgentResponse, _err error) {
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
		Action:      dara.String("UninstallMonitoringAgent"),
		Version:     dara.String("2019-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallMonitoringAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls the CloudMonitor agent from a third-party host.
//
// Description:
//
// > This API operation is not applicable to Elastic Compute Service (ECS) instances. To uninstall the agent from an ECS instance, see [Install and uninstall the CloudMonitor agent](https://help.aliyun.com/document_detail/183482.html).
//
// @param request - UninstallMonitoringAgentRequest
//
// @return UninstallMonitoringAgentResponse
func (client *Client) UninstallMonitoringAgent(request *UninstallMonitoringAgentRequest) (_result *UninstallMonitoringAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UninstallMonitoringAgentResponse{}
	_body, _err := client.UninstallMonitoringAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
