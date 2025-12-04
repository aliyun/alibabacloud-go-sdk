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
		"ap-northeast-2-pop":          dara.String("actiontrail.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("actiontrail.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("actiontrail.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("actiontrail.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("actiontrail.aliyuncs.com"),
		"cn-edge-1":                   dara.String("actiontrail.aliyuncs.com"),
		"cn-fujian":                   dara.String("actiontrail.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("actiontrail.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("actiontrail.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("actiontrail.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("actiontrail.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("actiontrail.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("actiontrail.aliyuncs.com"),
		"cn-wuhan":                    dara.String("actiontrail.aliyuncs.com"),
		"cn-yushanfang":               dara.String("actiontrail.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("actiontrail.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("actiontrail.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("actiontrail.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("actiontrail.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("actiontrail.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("actiontrail"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建高级查询历史记录
//
// @param request - CreateAdvancedQueryHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdvancedQueryHistoryResponse
func (client *Client) CreateAdvancedQueryHistoryWithOptions(request *CreateAdvancedQueryHistoryRequest, runtime *dara.RuntimeOptions) (_result *CreateAdvancedQueryHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QuerySql) {
		query["QuerySql"] = request.QuerySql
	}

	if !dara.IsNil(request.SimpleQuery) {
		query["SimpleQuery"] = request.SimpleQuery
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAdvancedQueryHistory"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAdvancedQueryHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建高级查询历史记录
//
// @param request - CreateAdvancedQueryHistoryRequest
//
// @return CreateAdvancedQueryHistoryResponse
func (client *Client) CreateAdvancedQueryHistory(request *CreateAdvancedQueryHistoryRequest) (_result *CreateAdvancedQueryHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAdvancedQueryHistoryResponse{}
	_body, _err := client.CreateAdvancedQueryHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建高级查询模板
//
// @param request - CreateAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdvancedQueryTemplateResponse
func (client *Client) CreateAdvancedQueryTemplateWithOptions(request *CreateAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAdvancedQueryTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SimpleQuery) {
		query["SimpleQuery"] = request.SimpleQuery
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateSql) {
		query["TemplateSql"] = request.TemplateSql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建高级查询模板
//
// @param request - CreateAdvancedQueryTemplateRequest
//
// @return CreateAdvancedQueryTemplateResponse
func (client *Client) CreateAdvancedQueryTemplate(request *CreateAdvancedQueryTemplateRequest) (_result *CreateAdvancedQueryTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAdvancedQueryTemplateResponse{}
	_body, _err := client.CreateAdvancedQueryTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data backfill task.
//
// Description:
//
// Limits
//
//   - Make sure that you have created a single-account trail to deliver events to Simple Log Service by calling the [CreateTrail](https://help.aliyun.com/document_detail/212313.html) operation.
//
//   - Only one data backfill task can run at a time within an Alibaba Cloud account.
//
// This topic provides an example on how to create a data backfill task for a trail named `trail-name`.
//
// @param request - CreateDeliveryHistoryJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeliveryHistoryJobResponse
func (client *Client) CreateDeliveryHistoryJobWithOptions(request *CreateDeliveryHistoryJobRequest, runtime *dara.RuntimeOptions) (_result *CreateDeliveryHistoryJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeliveryHistoryJob"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeliveryHistoryJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data backfill task.
//
// Description:
//
// Limits
//
//   - Make sure that you have created a single-account trail to deliver events to Simple Log Service by calling the [CreateTrail](https://help.aliyun.com/document_detail/212313.html) operation.
//
//   - Only one data backfill task can run at a time within an Alibaba Cloud account.
//
// This topic provides an example on how to create a data backfill task for a trail named `trail-name`.
//
// @param request - CreateDeliveryHistoryJobRequest
//
// @return CreateDeliveryHistoryJobResponse
func (client *Client) CreateDeliveryHistoryJob(request *CreateDeliveryHistoryJobRequest) (_result *CreateDeliveryHistoryJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDeliveryHistoryJobResponse{}
	_body, _err := client.CreateDeliveryHistoryJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a trail. By default, ActionTrail allows you to query events generated within your Alibaba Cloud account in the last 90 days. To query and analyze events generated more than 90 days ago, create a trail to deliver events to Object Storage Service (OSS), Simple Log Service, or MaxCompute.
//
// Description:
//
// *Operation description**
//
// >By default, a trail that is created by calling an operation is in the Disabled state. You must call the StartLogging operation to enable the trail. This way, ActionTrail can deliver events to the destination cloud service.
//
// **Prerequisites**
//
// Before you create a trail, make sure that at least one of the following storage configurations is complete:
//
// - Deliver events to OSS
//
//   - OSS is activated and a bucket is created.
//
// - Deliver events to Simple Log Service
//
//   - Simple Log Service is activated and a project is created.
//
//     >When a trail is created, ActionTrail automatically creates a Logstore named `actiontrail_<Trail name>` in the project. You cannot write data other than the audit data to the Logstore. This ensures the accuracy of the audit data.
//
// - Deliver events to MaxCompute
//
//   - MaxCompute is activated.
//
// >When a trail is created, ActionTrail automatically creates a project named `actiontrail_<Account ID>` on the Projects page. You cannot write data other than the audit data to the project. This ensures the accuracy of the audit data.
//
// **Usage Notes**
//
// This topic provides an example on how to create a single-account trail named `trail-test` to deliver events to an OSS bucket named `audit-log`.
//
// @param request - CreateTrailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrailResponse
func (client *Client) CreateTrailWithOptions(request *CreateTrailRequest, runtime *dara.RuntimeOptions) (_result *CreateTrailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventRW) {
		query["EventRW"] = request.EventRW
	}

	if !dara.IsNil(request.IsOrganizationTrail) {
		query["IsOrganizationTrail"] = request.IsOrganizationTrail
	}

	if !dara.IsNil(request.MaxComputeProjectArn) {
		query["MaxComputeProjectArn"] = request.MaxComputeProjectArn
	}

	if !dara.IsNil(request.MaxComputeWriteRoleArn) {
		query["MaxComputeWriteRoleArn"] = request.MaxComputeWriteRoleArn
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssKeyPrefix) {
		query["OssKeyPrefix"] = request.OssKeyPrefix
	}

	if !dara.IsNil(request.OssWriteRoleArn) {
		query["OssWriteRoleArn"] = request.OssWriteRoleArn
	}

	if !dara.IsNil(request.SlsProjectArn) {
		query["SlsProjectArn"] = request.SlsProjectArn
	}

	if !dara.IsNil(request.SlsWriteRoleArn) {
		query["SlsWriteRoleArn"] = request.SlsWriteRoleArn
	}

	if !dara.IsNil(request.TrailRegion) {
		query["TrailRegion"] = request.TrailRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrail"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trail. By default, ActionTrail allows you to query events generated within your Alibaba Cloud account in the last 90 days. To query and analyze events generated more than 90 days ago, create a trail to deliver events to Object Storage Service (OSS), Simple Log Service, or MaxCompute.
//
// Description:
//
// *Operation description**
//
// >By default, a trail that is created by calling an operation is in the Disabled state. You must call the StartLogging operation to enable the trail. This way, ActionTrail can deliver events to the destination cloud service.
//
// **Prerequisites**
//
// Before you create a trail, make sure that at least one of the following storage configurations is complete:
//
// - Deliver events to OSS
//
//   - OSS is activated and a bucket is created.
//
// - Deliver events to Simple Log Service
//
//   - Simple Log Service is activated and a project is created.
//
//     >When a trail is created, ActionTrail automatically creates a Logstore named `actiontrail_<Trail name>` in the project. You cannot write data other than the audit data to the Logstore. This ensures the accuracy of the audit data.
//
// - Deliver events to MaxCompute
//
//   - MaxCompute is activated.
//
// >When a trail is created, ActionTrail automatically creates a project named `actiontrail_<Account ID>` on the Projects page. You cannot write data other than the audit data to the project. This ensures the accuracy of the audit data.
//
// **Usage Notes**
//
// This topic provides an example on how to create a single-account trail named `trail-test` to deliver events to an OSS bucket named `audit-log`.
//
// @param request - CreateTrailRequest
//
// @return CreateTrailResponse
func (client *Client) CreateTrail(request *CreateTrailRequest) (_result *CreateTrailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTrailResponse{}
	_body, _err := client.CreateTrailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除高级查询历史记录
//
// @param request - DeleteAdvancedQueryHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdvancedQueryHistoryResponse
func (client *Client) DeleteAdvancedQueryHistoryWithOptions(request *DeleteAdvancedQueryHistoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdvancedQueryHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QueryId) {
		query["QueryId"] = request.QueryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAdvancedQueryHistory"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAdvancedQueryHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除高级查询历史记录
//
// @param request - DeleteAdvancedQueryHistoryRequest
//
// @return DeleteAdvancedQueryHistoryResponse
func (client *Client) DeleteAdvancedQueryHistory(request *DeleteAdvancedQueryHistoryRequest) (_result *DeleteAdvancedQueryHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAdvancedQueryHistoryResponse{}
	_body, _err := client.DeleteAdvancedQueryHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除高级查询模板
//
// @param request - DeleteAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdvancedQueryTemplateResponse
func (client *Client) DeleteAdvancedQueryTemplateWithOptions(request *DeleteAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdvancedQueryTemplateResponse, _err error) {
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
		Action:      dara.String("DeleteAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除高级查询模板
//
// @param request - DeleteAdvancedQueryTemplateRequest
//
// @return DeleteAdvancedQueryTemplateResponse
func (client *Client) DeleteAdvancedQueryTemplate(request *DeleteAdvancedQueryTemplateRequest) (_result *DeleteAdvancedQueryTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAdvancedQueryTemplateResponse{}
	_body, _err := client.DeleteAdvancedQueryTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据事件选择器
//
// @param request - DeleteDataEventSelectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataEventSelectorResponse
func (client *Client) DeleteDataEventSelectorWithOptions(request *DeleteDataEventSelectorRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataEventSelectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataEventSelector"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataEventSelectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据事件选择器
//
// @param request - DeleteDataEventSelectorRequest
//
// @return DeleteDataEventSelectorResponse
func (client *Client) DeleteDataEventSelector(request *DeleteDataEventSelectorRequest) (_result *DeleteDataEventSelectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataEventSelectorResponse{}
	_body, _err := client.DeleteDataEventSelectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data backfill task.
//
// Description:
//
// This topic describes how to delete a data backfill task whose ID is `16602`.
//
// @param request - DeleteDeliveryHistoryJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeliveryHistoryJobResponse
func (client *Client) DeleteDeliveryHistoryJobWithOptions(request *DeleteDeliveryHistoryJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeliveryHistoryJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeliveryHistoryJob"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeliveryHistoryJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data backfill task.
//
// Description:
//
// This topic describes how to delete a data backfill task whose ID is `16602`.
//
// @param request - DeleteDeliveryHistoryJobRequest
//
// @return DeleteDeliveryHistoryJobResponse
func (client *Client) DeleteDeliveryHistoryJob(request *DeleteDeliveryHistoryJobRequest) (_result *DeleteDeliveryHistoryJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDeliveryHistoryJobResponse{}
	_body, _err := client.DeleteDeliveryHistoryJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a trail.
//
// Description:
//
// This topic describes how to delete a sample trail named `trail-test`.
//
// @param request - DeleteTrailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrailResponse
func (client *Client) DeleteTrailWithOptions(request *DeleteTrailRequest, runtime *dara.RuntimeOptions) (_result *DeleteTrailResponse, _err error) {
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
		Action:      dara.String("DeleteTrail"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a trail.
//
// Description:
//
// This topic describes how to delete a sample trail named `trail-test`.
//
// @param request - DeleteTrailRequest
//
// @return DeleteTrailResponse
func (client *Client) DeleteTrail(request *DeleteTrailRequest) (_result *DeleteTrailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTrailResponse{}
	_body, _err := client.DeleteTrailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询高级查询历史记录
//
// @param request - DescribeAdvancedQueryHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvancedQueryHistoryResponse
func (client *Client) DescribeAdvancedQueryHistoryWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeAdvancedQueryHistoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvancedQueryHistory"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvancedQueryHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询高级查询历史记录
//
// @return DescribeAdvancedQueryHistoryResponse
func (client *Client) DescribeAdvancedQueryHistory() (_result *DescribeAdvancedQueryHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdvancedQueryHistoryResponse{}
	_body, _err := client.DescribeAdvancedQueryHistoryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询高级查询模板
//
// @param request - DescribeAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAdvancedQueryTemplateResponse
func (client *Client) DescribeAdvancedQueryTemplateWithOptions(request *DescribeAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeAdvancedQueryTemplateResponse, _err error) {
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

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询高级查询模板
//
// @param request - DescribeAdvancedQueryTemplateRequest
//
// @return DescribeAdvancedQueryTemplateResponse
func (client *Client) DescribeAdvancedQueryTemplate(request *DescribeAdvancedQueryTemplateRequest) (_result *DescribeAdvancedQueryTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAdvancedQueryTemplateResponse{}
	_body, _err := client.DescribeAdvancedQueryTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud regions that are supported by ActionTrail.
//
// Description:
//
// For more information, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud regions that are supported by ActionTrail.
//
// Description:
//
// For more information, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举资源生命周期事件
//
// @param request - DescribeResourceLifeCycleEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceLifeCycleEventsResponse
func (client *Client) DescribeResourceLifeCycleEventsWithOptions(request *DescribeResourceLifeCycleEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceLifeCycleEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceLifeCycleEvents"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceLifeCycleEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举资源生命周期事件
//
// @param request - DescribeResourceLifeCycleEventsRequest
//
// @return DescribeResourceLifeCycleEventsResponse
func (client *Client) DescribeResourceLifeCycleEvents(request *DescribeResourceLifeCycleEventsRequest) (_result *DescribeResourceLifeCycleEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceLifeCycleEventsResponse{}
	_body, _err := client.DescribeResourceLifeCycleEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询所有场景
//
// @param request - DescribeScenesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScenesResponse
func (client *Client) DescribeScenesWithOptions(request *DescribeScenesRequest, runtime *dara.RuntimeOptions) (_result *DescribeScenesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SearchCode) {
		query["SearchCode"] = request.SearchCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScenes"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有场景
//
// @param request - DescribeScenesRequest
//
// @return DescribeScenesResponse
func (client *Client) DescribeScenes(request *DescribeScenesRequest) (_result *DescribeScenesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeScenesResponse{}
	_body, _err := client.DescribeScenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举所有模版
//
// @param request - DescribeSearchTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSearchTemplatesResponse
func (client *Client) DescribeSearchTemplatesWithOptions(request *DescribeSearchTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSearchTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSearchTemplates"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSearchTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举所有模版
//
// @param request - DescribeSearchTemplatesRequest
//
// @return DescribeSearchTemplatesResponse
func (client *Client) DescribeSearchTemplates(request *DescribeSearchTemplatesRequest) (_result *DescribeSearchTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSearchTemplatesResponse{}
	_body, _err := client.DescribeSearchTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries created trails.
//
// Description:
//
// This topic shows you how to query the information about the single-account trails within an Alibaba Cloud account. In this example, the information about a trail named `test-4` is returned.
//
// @param request - DescribeTrailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrailsResponse
func (client *Client) DescribeTrailsWithOptions(request *DescribeTrailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeOrganizationTrail) {
		query["IncludeOrganizationTrail"] = request.IncludeOrganizationTrail
	}

	if !dara.IsNil(request.IncludeShadowTrails) {
		query["IncludeShadowTrails"] = request.IncludeShadowTrails
	}

	if !dara.IsNil(request.NameList) {
		query["NameList"] = request.NameList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrails"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries created trails.
//
// Description:
//
// This topic shows you how to query the information about the single-account trails within an Alibaba Cloud account. In this example, the information about a trail named `test-4` is returned.
//
// @param request - DescribeTrailsRequest
//
// @return DescribeTrailsResponse
func (client *Client) DescribeTrails(request *DescribeTrailsRequest) (_result *DescribeTrailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTrailsResponse{}
	_body, _err := client.DescribeTrailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户告警量
//
// @param request - DescribeUserAlertCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAlertCountResponse
func (client *Client) DescribeUserAlertCountWithOptions(request *DescribeUserAlertCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAlertCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserAlertCount"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserAlertCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户告警量
//
// @param request - DescribeUserAlertCountRequest
//
// @return DescribeUserAlertCountResponse
func (client *Client) DescribeUserAlertCount(request *DescribeUserAlertCountRequest) (_result *DescribeUserAlertCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserAlertCountResponse{}
	_body, _err := client.DescribeUserAlertCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户日志量
//
// @param request - DescribeUserLogCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserLogCountResponse
func (client *Client) DescribeUserLogCountWithOptions(request *DescribeUserLogCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserLogCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserLogCount"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserLogCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户日志量
//
// @param request - DescribeUserLogCountRequest
//
// @return DescribeUserLogCountResponse
func (client *Client) DescribeUserLogCount(request *DescribeUserLogCountRequest) (_result *DescribeUserLogCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserLogCountResponse{}
	_body, _err := client.DescribeUserLogCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Enables the Insights feature
//
// @param request - EnableInsightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInsightResponse
func (client *Client) EnableInsightWithOptions(request *EnableInsightRequest, runtime *dara.RuntimeOptions) (_result *EnableInsightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InsightType) {
		query["InsightType"] = request.InsightType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableInsight"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInsightResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enables the Insights feature
//
// @param request - EnableInsightRequest
//
// @return EnableInsightResponse
func (client *Client) EnableInsight(request *EnableInsightRequest) (_result *EnableInsightResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableInsightResponse{}
	_body, _err := client.EnableInsightWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the most recent events that are generated when a specified AccessKey pair is called to access Alibaba Cloud services.
//
// Description:
//
// You can call this operation to query only the information about the most recent events that are generated within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. For more information about supported events, see [Alibaba Cloud services and events that are supported by the AccessKey pair audit feature](https://help.aliyun.com/document_detail/419214.html). Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedEventsResponse
func (client *Client) GetAccessKeyLastUsedEventsWithOptions(request *GetAccessKeyLastUsedEventsRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedEventsResponse, _err error) {
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

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedEvents"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the most recent events that are generated when a specified AccessKey pair is called to access Alibaba Cloud services.
//
// Description:
//
// You can call this operation to query only the information about the most recent events that are generated within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. For more information about supported events, see [Alibaba Cloud services and events that are supported by the AccessKey pair audit feature](https://help.aliyun.com/document_detail/419214.html). Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedEventsRequest
//
// @return GetAccessKeyLastUsedEventsResponse
func (client *Client) GetAccessKeyLastUsedEvents(request *GetAccessKeyLastUsedEventsRequest) (_result *GetAccessKeyLastUsedEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedEventsResponse{}
	_body, _err := client.GetAccessKeyLastUsedEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the most recent call of a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about the most recent call of a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedInfoResponse
func (client *Client) GetAccessKeyLastUsedInfoWithOptions(request *GetAccessKeyLastUsedInfoRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedInfoResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedInfo"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the most recent call of a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about the most recent call of a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedInfoRequest
//
// @return GetAccessKeyLastUsedInfoResponse
func (client *Client) GetAccessKeyLastUsedInfo(request *GetAccessKeyLastUsedInfoRequest) (_result *GetAccessKeyLastUsedInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedInfoResponse{}
	_body, _err := client.GetAccessKeyLastUsedInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the IP addresses that are most recently used when an AccessKey pair is called to access Alibaba Cloud services.
//
// Description:
//
// You can call this operation to query only the information about the IP addresses that are most recently used within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedIpsResponse
func (client *Client) GetAccessKeyLastUsedIpsWithOptions(request *GetAccessKeyLastUsedIpsRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedIpsResponse, _err error) {
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

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedIps"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the IP addresses that are most recently used when an AccessKey pair is called to access Alibaba Cloud services.
//
// Description:
//
// You can call this operation to query only the information about the IP addresses that are most recently used within 400 days after February 1, 2022 when a specified AccessKey pair is called to access Alibaba Cloud services. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedIpsRequest
//
// @return GetAccessKeyLastUsedIpsResponse
func (client *Client) GetAccessKeyLastUsedIps(request *GetAccessKeyLastUsedIpsRequest) (_result *GetAccessKeyLastUsedIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedIpsResponse{}
	_body, _err := client.GetAccessKeyLastUsedIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the Alibaba Cloud services that are most recently accessed by using a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about Alibaba Cloud services that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedProductsResponse
func (client *Client) GetAccessKeyLastUsedProductsWithOptions(request *GetAccessKeyLastUsedProductsRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedProductsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedProducts"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the Alibaba Cloud services that are most recently accessed by using a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about Alibaba Cloud services that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedProductsRequest
//
// @return GetAccessKeyLastUsedProductsResponse
func (client *Client) GetAccessKeyLastUsedProducts(request *GetAccessKeyLastUsedProductsRequest) (_result *GetAccessKeyLastUsedProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedProductsResponse{}
	_body, _err := client.GetAccessKeyLastUsedProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the resources that are most recently accessed by using a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about resources that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedResourcesResponse
func (client *Client) GetAccessKeyLastUsedResourcesWithOptions(request *GetAccessKeyLastUsedResourcesRequest, runtime *dara.RuntimeOptions) (_result *GetAccessKeyLastUsedResourcesResponse, _err error) {
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

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessKeyLastUsedResources"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessKeyLastUsedResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the resources that are most recently accessed by using a specified AccessKey pair.
//
// Description:
//
// You can call this operation to query only the information about resources that are most recently accessed by using a specified AccessKey pair within 400 days after February 1, 2022. Data is updated at 1-hour intervals, which can cause query latency. We recommend that you do not change an AccessKey pair unless required.
//
// @param request - GetAccessKeyLastUsedResourcesRequest
//
// @return GetAccessKeyLastUsedResourcesResponse
func (client *Client) GetAccessKeyLastUsedResources(request *GetAccessKeyLastUsedResourcesRequest) (_result *GetAccessKeyLastUsedResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedResourcesResponse{}
	_body, _err := client.GetAccessKeyLastUsedResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个高级查询模板
//
// @param request - GetAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdvancedQueryTemplateResponse
func (client *Client) GetAdvancedQueryTemplateWithOptions(request *GetAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetAdvancedQueryTemplateResponse, _err error) {
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
		Action:      dara.String("GetAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个高级查询模板
//
// @param request - GetAdvancedQueryTemplateRequest
//
// @return GetAdvancedQueryTemplateResponse
func (client *Client) GetAdvancedQueryTemplate(request *GetAdvancedQueryTemplateRequest) (_result *GetAdvancedQueryTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAdvancedQueryTemplateResponse{}
	_body, _err := client.GetAdvancedQueryTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询事件选择器
//
// @param request - GetDataEventSelectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataEventSelectorResponse
func (client *Client) GetDataEventSelectorWithOptions(request *GetDataEventSelectorRequest, runtime *dara.RuntimeOptions) (_result *GetDataEventSelectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataEventSelector"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataEventSelectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询事件选择器
//
// @param request - GetDataEventSelectorRequest
//
// @return GetDataEventSelectorResponse
func (client *Client) GetDataEventSelector(request *GetDataEventSelectorRequest) (_result *GetDataEventSelectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataEventSelectorResponse{}
	_body, _err := client.GetDataEventSelectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data backfill task.
//
// Description:
//
// This topic provides an example on how to query the details of a data backfill task whose ID is `16602`. The return result shows that historical events for a trail named `trail-name` are delivered to Simple Log Service and the task is complete.
//
// @param request - GetDeliveryHistoryJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeliveryHistoryJobResponse
func (client *Client) GetDeliveryHistoryJobWithOptions(request *GetDeliveryHistoryJobRequest, runtime *dara.RuntimeOptions) (_result *GetDeliveryHistoryJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeliveryHistoryJob"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeliveryHistoryJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data backfill task.
//
// Description:
//
// This topic provides an example on how to query the details of a data backfill task whose ID is `16602`. The return result shows that historical events for a trail named `trail-name` are delivered to Simple Log Service and the task is complete.
//
// @param request - GetDeliveryHistoryJobRequest
//
// @return GetDeliveryHistoryJobResponse
func (client *Client) GetDeliveryHistoryJob(request *GetDeliveryHistoryJobRequest) (_result *GetDeliveryHistoryJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeliveryHistoryJobResponse{}
	_body, _err := client.GetDeliveryHistoryJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the region where global events are stored.
//
// Description:
//
// By default, global events are stored in the Singapore region.
//
// To obtain the permissions to call the API operation, you must submit a ticket.
//
// @param request - GetGlobalEventsStorageRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGlobalEventsStorageRegionResponse
func (client *Client) GetGlobalEventsStorageRegionWithOptions(runtime *dara.RuntimeOptions) (_result *GetGlobalEventsStorageRegionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetGlobalEventsStorageRegion"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGlobalEventsStorageRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the region where global events are stored.
//
// Description:
//
// By default, global events are stored in the Singapore region.
//
// To obtain the permissions to call the API operation, you must submit a ticket.
//
// @return GetGlobalEventsStorageRegionResponse
func (client *Client) GetGlobalEventsStorageRegion() (_result *GetGlobalEventsStorageRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGlobalEventsStorageRegionResponse{}
	_body, _err := client.GetGlobalEventsStorageRegionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 操作审计成熟度查询接口
//
// @param request - GetGovernanceMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGovernanceMetricsResponse
func (client *Client) GetGovernanceMetricsWithOptions(runtime *dara.RuntimeOptions) (_result *GetGovernanceMetricsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetGovernanceMetrics"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGovernanceMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 操作审计成熟度查询接口
//
// @return GetGovernanceMetricsResponse
func (client *Client) GetGovernanceMetrics() (_result *GetGovernanceMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGovernanceMetricsResponse{}
	_body, _err := client.GetGovernanceMetricsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a trail.
//
// Description:
//
// This topic describes how to query the status of a sample single-account trail named `trail-test`.
//
// @param request - GetTrailStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrailStatusResponse
func (client *Client) GetTrailStatusWithOptions(request *GetTrailStatusRequest, runtime *dara.RuntimeOptions) (_result *GetTrailStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsOrganizationTrail) {
		query["IsOrganizationTrail"] = request.IsOrganizationTrail
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrailStatus"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrailStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a trail.
//
// Description:
//
// This topic describes how to query the status of a sample single-account trail named `trail-test`.
//
// @param request - GetTrailStatusRequest
//
// @return GetTrailStatusResponse
func (client *Client) GetTrailStatus(request *GetTrailStatusRequest) (_result *GetTrailStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTrailStatusResponse{}
	_body, _err := client.GetTrailStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询事件选择器
//
// @param request - ListDataEventSelectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataEventSelectorsResponse
func (client *Client) ListDataEventSelectorsWithOptions(request *ListDataEventSelectorsRequest, runtime *dara.RuntimeOptions) (_result *ListDataEventSelectorsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataEventSelectors"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataEventSelectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询事件选择器
//
// @param request - ListDataEventSelectorsRequest
//
// @return ListDataEventSelectorsResponse
func (client *Client) ListDataEventSelectors(request *ListDataEventSelectorsRequest) (_result *ListDataEventSelectorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataEventSelectorsResponse{}
	_body, _err := client.ListDataEventSelectorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据事件支持的服务与事件名称
//
// @param request - ListDataEventServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataEventServicesResponse
func (client *Client) ListDataEventServicesWithOptions(request *ListDataEventServicesRequest, runtime *dara.RuntimeOptions) (_result *ListDataEventServicesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataEventServices"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataEventServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据事件支持的服务与事件名称
//
// @param request - ListDataEventServicesRequest
//
// @return ListDataEventServicesResponse
func (client *Client) ListDataEventServices(request *ListDataEventServicesRequest) (_result *ListDataEventServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataEventServicesResponse{}
	_body, _err := client.ListDataEventServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data backfill tasks.
//
// Description:
//
// This topic provides an example on how to query a list of data backfill tasks. The returned result shows that a data backfill task with the ID `16602` is used to deliver historical events for a trail named `trail-name` to Simple Log Service.
//
// @param request - ListDeliveryHistoryJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeliveryHistoryJobsResponse
func (client *Client) ListDeliveryHistoryJobsWithOptions(request *ListDeliveryHistoryJobsRequest, runtime *dara.RuntimeOptions) (_result *ListDeliveryHistoryJobsResponse, _err error) {
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
		Action:      dara.String("ListDeliveryHistoryJobs"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeliveryHistoryJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data backfill tasks.
//
// Description:
//
// This topic provides an example on how to query a list of data backfill tasks. The returned result shows that a data backfill task with the ID `16602` is used to deliver historical events for a trail named `trail-name` to Simple Log Service.
//
// @param request - ListDeliveryHistoryJobsRequest
//
// @return ListDeliveryHistoryJobsResponse
func (client *Client) ListDeliveryHistoryJobs(request *ListDeliveryHistoryJobsRequest) (_result *ListDeliveryHistoryJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeliveryHistoryJobsResponse{}
	_body, _err := client.ListDeliveryHistoryJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries event details.
//
// Description:
//
// When you call this operation to query event details, you can query the event details at most twice per second.
//
// > Do not frequently call this operation. You can create a trail to deliver events to Log Service. Then, you can query event details in near real time by using the real-time log consumption feature of Log Service. For more information, see [Create a single-account trail](https://help.aliyun.com/document_detail/28810.html), [Create a multi-account trail](https://help.aliyun.com/document_detail/160661.html), and [Overview](https://help.aliyun.com/document_detail/28997.html).
//
// @param request - LookupEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LookupEventsResponse
func (client *Client) LookupEventsWithOptions(request *LookupEventsRequest, runtime *dara.RuntimeOptions) (_result *LookupEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LookupAttribute) {
		query["LookupAttribute"] = request.LookupAttribute
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LookupEvents"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LookupEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries event details.
//
// Description:
//
// When you call this operation to query event details, you can query the event details at most twice per second.
//
// > Do not frequently call this operation. You can create a trail to deliver events to Log Service. Then, you can query event details in near real time by using the real-time log consumption feature of Log Service. For more information, see [Create a single-account trail](https://help.aliyun.com/document_detail/28810.html), [Create a multi-account trail](https://help.aliyun.com/document_detail/160661.html), and [Overview](https://help.aliyun.com/document_detail/28997.html).
//
// @param request - LookupEventsRequest
//
// @return LookupEventsResponse
func (client *Client) LookupEvents(request *LookupEventsRequest) (_result *LookupEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LookupEventsResponse{}
	_body, _err := client.LookupEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建事件选择器
//
// @param request - PutDataEventSelectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutDataEventSelectorResponse
func (client *Client) PutDataEventSelectorWithOptions(request *PutDataEventSelectorRequest, runtime *dara.RuntimeOptions) (_result *PutDataEventSelectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventSelectors) {
		query["EventSelectors"] = request.EventSelectors
	}

	if !dara.IsNil(request.IsTrailAllRegion) {
		query["IsTrailAllRegion"] = request.IsTrailAllRegion
	}

	if !dara.IsNil(request.TrailName) {
		query["TrailName"] = request.TrailName
	}

	if !dara.IsNil(request.TrailRegionIds) {
		query["TrailRegionIds"] = request.TrailRegionIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutDataEventSelector"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutDataEventSelectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建事件选择器
//
// @param request - PutDataEventSelectorRequest
//
// @return PutDataEventSelectorResponse
func (client *Client) PutDataEventSelector(request *PutDataEventSelectorRequest) (_result *PutDataEventSelectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutDataEventSelectorResponse{}
	_body, _err := client.PutDataEventSelectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a trail to deliver events to an Object Storage Service (OSS) bucket or a Simple Log Service Logstore.
//
// Description:
//
// This topic describes how to enable logging for a sample trail named `trail-test`.
//
// @param request - StartLoggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLoggingResponse
func (client *Client) StartLoggingWithOptions(request *StartLoggingRequest, runtime *dara.RuntimeOptions) (_result *StartLoggingResponse, _err error) {
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
		Action:      dara.String("StartLogging"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLoggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a trail to deliver events to an Object Storage Service (OSS) bucket or a Simple Log Service Logstore.
//
// Description:
//
// This topic describes how to enable logging for a sample trail named `trail-test`.
//
// @param request - StartLoggingRequest
//
// @return StartLoggingResponse
func (client *Client) StartLogging(request *StartLoggingRequest) (_result *StartLoggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartLoggingResponse{}
	_body, _err := client.StartLoggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a trail to stop the delivery of events to an Object Storage Service (OSS) bucket or a  Simple Log Service Logstore.
//
// Description:
//
// This topic describes how to disable logging for a sample trail named `trail-test`.
//
// @param request - StopLoggingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLoggingResponse
func (client *Client) StopLoggingWithOptions(request *StopLoggingRequest, runtime *dara.RuntimeOptions) (_result *StopLoggingResponse, _err error) {
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
		Action:      dara.String("StopLogging"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLoggingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a trail to stop the delivery of events to an Object Storage Service (OSS) bucket or a  Simple Log Service Logstore.
//
// Description:
//
// This topic describes how to disable logging for a sample trail named `trail-test`.
//
// @param request - StopLoggingRequest
//
// @return StopLoggingResponse
func (client *Client) StopLogging(request *StopLoggingRequest) (_result *StopLoggingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopLoggingResponse{}
	_body, _err := client.StopLoggingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新高级查询模板
//
// @param request - UpdateAdvancedQueryTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdvancedQueryTemplateResponse
func (client *Client) UpdateAdvancedQueryTemplateWithOptions(request *UpdateAdvancedQueryTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdvancedQueryTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SimpleQuery) {
		query["SimpleQuery"] = request.SimpleQuery
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateSql) {
		query["TemplateSql"] = request.TemplateSql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAdvancedQueryTemplate"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAdvancedQueryTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新高级查询模板
//
// @param request - UpdateAdvancedQueryTemplateRequest
//
// @return UpdateAdvancedQueryTemplateResponse
func (client *Client) UpdateAdvancedQueryTemplate(request *UpdateAdvancedQueryTemplateRequest) (_result *UpdateAdvancedQueryTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAdvancedQueryTemplateResponse{}
	_body, _err := client.UpdateAdvancedQueryTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies the region where you want to store global events.
//
// Description:
//
// By default, global events are stored in the Singapore region.
//
//   - To obtain the permissions to call the API operation, you must submit a ticket.
//
//   - Only the China (Hangzhou) region (cn-hangzhou) and the Singapore region (ap-southeast-1) are supported.
//
// @param request - UpdateGlobalEventsStorageRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGlobalEventsStorageRegionResponse
func (client *Client) UpdateGlobalEventsStorageRegionWithOptions(request *UpdateGlobalEventsStorageRegionRequest, runtime *dara.RuntimeOptions) (_result *UpdateGlobalEventsStorageRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.StorageRegion) {
		query["StorageRegion"] = request.StorageRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGlobalEventsStorageRegion"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGlobalEventsStorageRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the region where you want to store global events.
//
// Description:
//
// By default, global events are stored in the Singapore region.
//
//   - To obtain the permissions to call the API operation, you must submit a ticket.
//
//   - Only the China (Hangzhou) region (cn-hangzhou) and the Singapore region (ap-southeast-1) are supported.
//
// @param request - UpdateGlobalEventsStorageRegionRequest
//
// @return UpdateGlobalEventsStorageRegionResponse
func (client *Client) UpdateGlobalEventsStorageRegion(request *UpdateGlobalEventsStorageRegionRequest) (_result *UpdateGlobalEventsStorageRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateGlobalEventsStorageRegionResponse{}
	_body, _err := client.UpdateGlobalEventsStorageRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of a trail.
//
// Description:
//
// This topic shows you how to change the destination Object Storage Service (OSS) bucket of a sample trail named `trail-test` to `audit-log`.
//
// @param request - UpdateTrailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTrailResponse
func (client *Client) UpdateTrailWithOptions(request *UpdateTrailRequest, runtime *dara.RuntimeOptions) (_result *UpdateTrailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventRW) {
		query["EventRW"] = request.EventRW
	}

	if !dara.IsNil(request.MaxComputeProjectArn) {
		query["MaxComputeProjectArn"] = request.MaxComputeProjectArn
	}

	if !dara.IsNil(request.MaxComputeWriteRoleArn) {
		query["MaxComputeWriteRoleArn"] = request.MaxComputeWriteRoleArn
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OssBucketName) {
		query["OssBucketName"] = request.OssBucketName
	}

	if !dara.IsNil(request.OssKeyPrefix) {
		query["OssKeyPrefix"] = request.OssKeyPrefix
	}

	if !dara.IsNil(request.OssWriteRoleArn) {
		query["OssWriteRoleArn"] = request.OssWriteRoleArn
	}

	if !dara.IsNil(request.SlsProjectArn) {
		query["SlsProjectArn"] = request.SlsProjectArn
	}

	if !dara.IsNil(request.SlsWriteRoleArn) {
		query["SlsWriteRoleArn"] = request.SlsWriteRoleArn
	}

	if !dara.IsNil(request.TrailRegion) {
		query["TrailRegion"] = request.TrailRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTrail"),
		Version:     dara.String("2020-07-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTrailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of a trail.
//
// Description:
//
// This topic shows you how to change the destination Object Storage Service (OSS) bucket of a sample trail named `trail-test` to `audit-log`.
//
// @param request - UpdateTrailRequest
//
// @return UpdateTrailResponse
func (client *Client) UpdateTrail(request *UpdateTrailRequest) (_result *UpdateTrailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTrailResponse{}
	_body, _err := client.UpdateTrailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
