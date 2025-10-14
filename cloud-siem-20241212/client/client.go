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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cloud-siem"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 检查升级项
//
// @param request - CheckUpgradeItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUpgradeItemResponse
func (client *Client) CheckUpgradeItemWithOptions(request *CheckUpgradeItemRequest, runtime *dara.RuntimeOptions) (_result *CheckUpgradeItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.UpgradeItemId) {
		body["UpgradeItemId"] = request.UpgradeItemId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckUpgradeItem"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckUpgradeItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查升级项
//
// @param request - CheckUpgradeItemRequest
//
// @return CheckUpgradeItemResponse
func (client *Client) CheckUpgradeItem(request *CheckUpgradeItemRequest) (_result *CheckUpgradeItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckUpgradeItemResponse{}
	_body, _err := client.CheckUpgradeItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据源
//
// @param request - CreateDataIngestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataIngestionResponse
func (client *Client) CreateDataIngestionWithOptions(request *CreateDataIngestionRequest, runtime *dara.RuntimeOptions) (_result *CreateDataIngestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CapacityCount) {
		body["CapacityCount"] = request.CapacityCount
	}

	if !dara.IsNil(request.DataIngestionMode) {
		body["DataIngestionMode"] = request.DataIngestionMode
	}

	if !dara.IsNil(request.DataIngestionStateCode) {
		body["DataIngestionStateCode"] = request.DataIngestionStateCode
	}

	if !dara.IsNil(request.DataIngestionType) {
		body["DataIngestionType"] = request.DataIngestionType
	}

	if !dara.IsNil(request.DataSourceEditable) {
		body["DataSourceEditable"] = request.DataSourceEditable
	}

	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationRuleEditable) {
		body["NormalizationRuleEditable"] = request.NormalizationRuleEditable
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.ScanDataSourceId) {
		body["ScanDataSourceId"] = request.ScanDataSourceId
	}

	if !dara.IsNil(request.StreamJobId) {
		body["StreamJobId"] = request.StreamJobId
	}

	if !dara.IsNil(request.UpdateTime) {
		body["UpdateTime"] = request.UpdateTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataIngestion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataIngestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据源
//
// @param request - CreateDataIngestionRequest
//
// @return CreateDataIngestionResponse
func (client *Client) CreateDataIngestion(request *CreateDataIngestionRequest) (_result *CreateDataIngestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataIngestionResponse{}
	_body, _err := client.CreateDataIngestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSetResponse
func (client *Client) CreateDataSetWithOptions(request *CreateDataSetRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSetDescription) {
		body["DataSetDescription"] = request.DataSetDescription
	}

	if !dara.IsNil(request.DataSetFieldKeyName) {
		body["DataSetFieldKeyName"] = request.DataSetFieldKeyName
	}

	if !dara.IsNil(request.DataSetFileName) {
		body["DataSetFileName"] = request.DataSetFileName
	}

	if !dara.IsNil(request.DataSetName) {
		body["DataSetName"] = request.DataSetName
	}

	if !dara.IsNil(request.DataSetStatus) {
		body["DataSetStatus"] = request.DataSetStatus
	}

	if !dara.IsNil(request.DataSetType) {
		body["DataSetType"] = request.DataSetType
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.IpWhitelistRecognizers) {
		bodyFlat["IpWhitelistRecognizers"] = request.IpWhitelistRecognizers
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSet"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateDataSetRequest
//
// @return CreateDataSetResponse
func (client *Client) CreateDataSet(request *CreateDataSetRequest) (_result *CreateDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataSetResponse{}
	_body, _err := client.CreateDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据源
//
// @param tmpReq - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(tmpReq *CreateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataSourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSourceIds) {
		request.DataSourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSourceIds, dara.String("DataSourceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DataSourceReferences) {
		request.DataSourceReferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSourceReferences, dara.String("DataSourceReferences"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceFrom) {
		body["DataSourceFrom"] = request.DataSourceFrom
	}

	if !dara.IsNil(request.DataSourceIdsShrink) {
		body["DataSourceIds"] = request.DataSourceIdsShrink
	}

	if !dara.IsNil(request.DataSourceName) {
		body["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.DataSourceRecognizeEnabled) {
		body["DataSourceRecognizeEnabled"] = request.DataSourceRecognizeEnabled
	}

	if !dara.IsNil(request.DataSourceRecognizer) {
		body["DataSourceRecognizer"] = request.DataSourceRecognizer
	}

	if !dara.IsNil(request.DataSourceReferencesShrink) {
		body["DataSourceReferences"] = request.DataSourceReferencesShrink
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceStores) {
		bodyFlat["DataSourceStores"] = request.DataSourceStores
	}

	if !dara.IsNil(request.DataSourceTemplateId) {
		body["DataSourceTemplateId"] = request.DataSourceTemplateId
	}

	if !dara.IsNil(request.DataSourceType) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogProjectName) {
		body["LogProjectName"] = request.LogProjectName
	}

	if !dara.IsNil(request.LogRegionId) {
		body["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogStoreName) {
		body["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.UpdateTime) {
		body["UpdateTime"] = request.UpdateTime
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSource"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据源
//
// @param request - CreateDataSourceRequest
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建检测规则
//
// @param request - CreateDetectionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDetectionRuleResponse
func (client *Client) CreateDetectionRuleWithOptions(request *CreateDetectionRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDetectionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertAttCk) {
		body["AlertAttCk"] = request.AlertAttCk
	}

	if !dara.IsNil(request.AlertLevel) {
		body["AlertLevel"] = request.AlertLevel
	}

	if !dara.IsNil(request.AlertSchemaId) {
		body["AlertSchemaId"] = request.AlertSchemaId
	}

	if !dara.IsNil(request.AlertTacticId) {
		body["AlertTacticId"] = request.AlertTacticId
	}

	if !dara.IsNil(request.AlertThresholdCount) {
		body["AlertThresholdCount"] = request.AlertThresholdCount
	}

	if !dara.IsNil(request.AlertThresholdGroup) {
		body["AlertThresholdGroup"] = request.AlertThresholdGroup
	}

	if !dara.IsNil(request.AlertThresholdPeriod) {
		body["AlertThresholdPeriod"] = request.AlertThresholdPeriod
	}

	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.DetectionExpressionContent) {
		body["DetectionExpressionContent"] = request.DetectionExpressionContent
	}

	if !dara.IsNil(request.DetectionExpressionType) {
		body["DetectionExpressionType"] = request.DetectionExpressionType
	}

	if !dara.IsNil(request.DetectionRuleDescription) {
		body["DetectionRuleDescription"] = request.DetectionRuleDescription
	}

	if !dara.IsNil(request.DetectionRuleName) {
		body["DetectionRuleName"] = request.DetectionRuleName
	}

	if !dara.IsNil(request.DetectionRuleStatus) {
		body["DetectionRuleStatus"] = request.DetectionRuleStatus
	}

	if !dara.IsNil(request.DetectionRuleType) {
		body["DetectionRuleType"] = request.DetectionRuleType
	}

	if !dara.IsNil(request.EntityMappings) {
		body["EntityMappings"] = request.EntityMappings
	}

	if !dara.IsNil(request.IncidentAggregationExpression) {
		body["IncidentAggregationExpression"] = request.IncidentAggregationExpression
	}

	if !dara.IsNil(request.IncidentAggregationType) {
		body["IncidentAggregationType"] = request.IncidentAggregationType
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogCategoryId) {
		body["LogCategoryId"] = request.LogCategoryId
	}

	if !dara.IsNil(request.LogSchemaId) {
		body["LogSchemaId"] = request.LogSchemaId
	}

	if !dara.IsNil(request.PlaybookParameters) {
		body["PlaybookParameters"] = request.PlaybookParameters
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.ScheduleBeginTime) {
		body["ScheduleBeginTime"] = request.ScheduleBeginTime
	}

	if !dara.IsNil(request.ScheduleExpression) {
		body["ScheduleExpression"] = request.ScheduleExpression
	}

	if !dara.IsNil(request.ScheduleMaxRetries) {
		body["ScheduleMaxRetries"] = request.ScheduleMaxRetries
	}

	if !dara.IsNil(request.ScheduleMaxTimeout) {
		body["ScheduleMaxTimeout"] = request.ScheduleMaxTimeout
	}

	if !dara.IsNil(request.ScheduleType) {
		body["ScheduleType"] = request.ScheduleType
	}

	if !dara.IsNil(request.ScheduleWindow) {
		body["ScheduleWindow"] = request.ScheduleWindow
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDetectionRule"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDetectionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建检测规则
//
// @param request - CreateDetectionRuleRequest
//
// @return CreateDetectionRuleResponse
func (client *Client) CreateDetectionRule(request *CreateDetectionRuleRequest) (_result *CreateDetectionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDetectionRuleResponse{}
	_body, _err := client.CreateDetectionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建导出任务
//
// @param request - CreateExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExportTaskResponse
func (client *Client) CreateExportTaskWithOptions(request *CreateExportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExportTaskParameter) {
		body["ExportTaskParameter"] = request.ExportTaskParameter
	}

	if !dara.IsNil(request.ExportTaskType) {
		body["ExportTaskType"] = request.ExportTaskType
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExportTask"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建导出任务
//
// @param request - CreateExportTaskRequest
//
// @return CreateExportTaskResponse
func (client *Client) CreateExportTask(request *CreateExportTaskRequest) (_result *CreateExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExportTaskResponse{}
	_body, _err := client.CreateExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建LogStore
//
// @param request - CreateLogStoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogStoreResponse
func (client *Client) CreateLogStoreWithOptions(request *CreateLogStoreRequest, runtime *dara.RuntimeOptions) (_result *CreateLogStoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogProjectName) {
		body["LogProjectName"] = request.LogProjectName
	}

	if !dara.IsNil(request.LogRegionId) {
		body["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogStoreName) {
		body["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogStore"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建LogStore
//
// @param request - CreateLogStoreRequest
//
// @return CreateLogStoreResponse
func (client *Client) CreateLogStore(request *CreateLogStoreRequest) (_result *CreateLogStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLogStoreResponse{}
	_body, _err := client.CreateLogStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建标准化规则
//
// @param tmpReq - CreateNormalizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNormalizationRuleResponse
func (client *Client) CreateNormalizationRuleWithOptions(tmpReq *CreateNormalizationRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateNormalizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateNormalizationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NormalizationRuleIds) {
		request.NormalizationRuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NormalizationRuleIds, dara.String("NormalizationRuleIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtendContentPacked) {
		body["ExtendContentPacked"] = request.ExtendContentPacked
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationCategoryId) {
		body["NormalizationCategoryId"] = request.NormalizationCategoryId
	}

	if !dara.IsNil(request.NormalizationRuleDescription) {
		body["NormalizationRuleDescription"] = request.NormalizationRuleDescription
	}

	if !dara.IsNil(request.NormalizationRuleExpression) {
		body["NormalizationRuleExpression"] = request.NormalizationRuleExpression
	}

	if !dara.IsNil(request.NormalizationRuleFormat) {
		body["NormalizationRuleFormat"] = request.NormalizationRuleFormat
	}

	if !dara.IsNil(request.NormalizationRuleIdsShrink) {
		body["NormalizationRuleIds"] = request.NormalizationRuleIdsShrink
	}

	if !dara.IsNil(request.NormalizationRuleMode) {
		body["NormalizationRuleMode"] = request.NormalizationRuleMode
	}

	if !dara.IsNil(request.NormalizationRuleName) {
		body["NormalizationRuleName"] = request.NormalizationRuleName
	}

	if !dara.IsNil(request.NormalizationRuleType) {
		body["NormalizationRuleType"] = request.NormalizationRuleType
	}

	if !dara.IsNil(request.NormalizationRuleVersion) {
		body["NormalizationRuleVersion"] = request.NormalizationRuleVersion
	}

	if !dara.IsNil(request.NormalizationSchemaId) {
		body["NormalizationSchemaId"] = request.NormalizationSchemaId
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorId) {
		body["VendorId"] = request.VendorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNormalizationRule"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNormalizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建标准化规则
//
// @param request - CreateNormalizationRuleRequest
//
// @return CreateNormalizationRuleResponse
func (client *Client) CreateNormalizationRule(request *CreateNormalizationRuleRequest) (_result *CreateNormalizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNormalizationRuleResponse{}
	_body, _err := client.CreateNormalizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建产品
//
// @param request - CreateProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProductResponse
func (client *Client) CreateProductWithOptions(request *CreateProductRequest, runtime *dara.RuntimeOptions) (_result *CreateProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductName) {
		body["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorName) {
		body["VendorName"] = request.VendorName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProduct"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建产品
//
// @param request - CreateProductRequest
//
// @return CreateProductResponse
func (client *Client) CreateProduct(request *CreateProductRequest) (_result *CreateProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProductResponse{}
	_body, _err := client.CreateProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建厂商
//
// @param request - CreateVendorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVendorResponse
func (client *Client) CreateVendorWithOptions(request *CreateVendorRequest, runtime *dara.RuntimeOptions) (_result *CreateVendorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorName) {
		body["VendorName"] = request.VendorName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVendor"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVendorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建厂商
//
// @param request - CreateVendorRequest
//
// @return CreateVendorResponse
func (client *Client) CreateVendor(request *CreateVendorRequest) (_result *CreateVendorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVendorResponse{}
	_body, _err := client.CreateVendorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据接入
//
// @param request - DeleteDataIngestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataIngestionResponse
func (client *Client) DeleteDataIngestionWithOptions(request *DeleteDataIngestionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataIngestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataIngestionId) {
		body["DataIngestionId"] = request.DataIngestionId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataIngestion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataIngestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据接入
//
// @param request - DeleteDataIngestionRequest
//
// @return DeleteDataIngestionResponse
func (client *Client) DeleteDataIngestion(request *DeleteDataIngestionRequest) (_result *DeleteDataIngestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataIngestionResponse{}
	_body, _err := client.DeleteDataIngestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - DeleteDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSetResponse
func (client *Client) DeleteDataSetWithOptions(request *DeleteDataSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSetId) {
		body["DataSetId"] = request.DataSetId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSet"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - DeleteDataSetRequest
//
// @return DeleteDataSetResponse
func (client *Client) DeleteDataSet(request *DeleteDataSetRequest) (_result *DeleteDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSetResponse{}
	_body, _err := client.DeleteDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据集记录
//
// @param request - DeleteDataSetRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSetRecordResponse
func (client *Client) DeleteDataSetRecordWithOptions(request *DeleteDataSetRecordRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSetRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSetId) {
		body["DataSetId"] = request.DataSetId
	}

	if !dara.IsNil(request.DataSetRecordIds) {
		body["DataSetRecordIds"] = request.DataSetRecordIds
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSetRecord"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSetRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集记录
//
// @param request - DeleteDataSetRecordRequest
//
// @return DeleteDataSetRecordResponse
func (client *Client) DeleteDataSetRecord(request *DeleteDataSetRecordRequest) (_result *DeleteDataSetRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSetRecordResponse{}
	_body, _err := client.DeleteDataSetRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据源
//
// @param request - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSource"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据源
//
// @param request - DeleteDataSourceRequest
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除检测规则
//
// @param request - DeleteDetectionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDetectionRuleResponse
func (client *Client) DeleteDetectionRuleWithOptions(request *DeleteDetectionRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDetectionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DetectionRuleId) {
		body["DetectionRuleId"] = request.DetectionRuleId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDetectionRule"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDetectionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除检测规则
//
// @param request - DeleteDetectionRuleRequest
//
// @return DeleteDetectionRuleResponse
func (client *Client) DeleteDetectionRule(request *DeleteDetectionRuleRequest) (_result *DeleteDetectionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDetectionRuleResponse{}
	_body, _err := client.DeleteDetectionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除LogStore
//
// @param request - DeleteLogStoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogStoreResponse
func (client *Client) DeleteLogStoreWithOptions(request *DeleteLogStoreRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogStoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogProjectName) {
		body["LogProjectName"] = request.LogProjectName
	}

	if !dara.IsNil(request.LogRegionId) {
		body["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogStoreName) {
		body["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogStore"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除LogStore
//
// @param request - DeleteLogStoreRequest
//
// @return DeleteLogStoreResponse
func (client *Client) DeleteLogStore(request *DeleteLogStoreRequest) (_result *DeleteLogStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLogStoreResponse{}
	_body, _err := client.DeleteLogStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除标准化规则
//
// @param request - DeleteNormalizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNormalizationRuleResponse
func (client *Client) DeleteNormalizationRuleWithOptions(request *DeleteNormalizationRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteNormalizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNormalizationRule"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNormalizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标准化规则
//
// @param request - DeleteNormalizationRuleRequest
//
// @return DeleteNormalizationRuleResponse
func (client *Client) DeleteNormalizationRule(request *DeleteNormalizationRuleRequest) (_result *DeleteNormalizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNormalizationRuleResponse{}
	_body, _err := client.DeleteNormalizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除标准化规则版本
//
// @param request - DeleteNormalizationRuleVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNormalizationRuleVersionResponse
func (client *Client) DeleteNormalizationRuleVersionWithOptions(request *DeleteNormalizationRuleVersionRequest, runtime *dara.RuntimeOptions) (_result *DeleteNormalizationRuleVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.NormalizationRuleVersion) {
		body["NormalizationRuleVersion"] = request.NormalizationRuleVersion
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNormalizationRuleVersion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNormalizationRuleVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标准化规则版本
//
// @param request - DeleteNormalizationRuleVersionRequest
//
// @return DeleteNormalizationRuleVersionResponse
func (client *Client) DeleteNormalizationRuleVersion(request *DeleteNormalizationRuleVersionRequest) (_result *DeleteNormalizationRuleVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNormalizationRuleVersionResponse{}
	_body, _err := client.DeleteNormalizationRuleVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除产品
//
// @param request - DeleteProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProductResponse
func (client *Client) DeleteProductWithOptions(request *DeleteProductRequest, runtime *dara.RuntimeOptions) (_result *DeleteProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProduct"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除产品
//
// @param request - DeleteProductRequest
//
// @return DeleteProductResponse
func (client *Client) DeleteProduct(request *DeleteProductRequest) (_result *DeleteProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProductResponse{}
	_body, _err := client.DeleteProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除厂商
//
// @param request - DeleteVendorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVendorResponse
func (client *Client) DeleteVendorWithOptions(request *DeleteVendorRequest, runtime *dara.RuntimeOptions) (_result *DeleteVendorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorId) {
		body["VendorId"] = request.VendorId
	}

	if !dara.IsNil(request.VendorName) {
		body["VendorName"] = request.VendorName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVendor"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVendorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除厂商
//
// @param request - DeleteVendorRequest
//
// @return DeleteVendorResponse
func (client *Client) DeleteVendor(request *DeleteVendorRequest) (_result *DeleteVendorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVendorResponse{}
	_body, _err := client.DeleteVendorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止数据接入
//
// @param request - DisableDataIngestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDataIngestionResponse
func (client *Client) DisableDataIngestionWithOptions(request *DisableDataIngestionRequest, runtime *dara.RuntimeOptions) (_result *DisableDataIngestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataIngestionId) {
		body["DataIngestionId"] = request.DataIngestionId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDataIngestion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDataIngestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止数据接入
//
// @param request - DisableDataIngestionRequest
//
// @return DisableDataIngestionResponse
func (client *Client) DisableDataIngestion(request *DisableDataIngestionRequest) (_result *DisableDataIngestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDataIngestionResponse{}
	_body, _err := client.DisableDataIngestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动数据接入
//
// @param request - EnableDataIngestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDataIngestionResponse
func (client *Client) EnableDataIngestionWithOptions(request *EnableDataIngestionRequest, runtime *dara.RuntimeOptions) (_result *EnableDataIngestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataIngestionId) {
		body["DataIngestionId"] = request.DataIngestionId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDataIngestion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDataIngestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动数据接入
//
// @param request - EnableDataIngestionRequest
//
// @return EnableDataIngestionResponse
func (client *Client) EnableDataIngestion(request *EnableDataIngestionRequest) (_result *EnableDataIngestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDataIngestionResponse{}
	_body, _err := client.EnableDataIngestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看LogStore
//
// @param request - ExecuteLogQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteLogQueryResponse
func (client *Client) ExecuteLogQueryWithOptions(request *ExecuteLogQueryRequest, runtime *dara.RuntimeOptions) (_result *ExecuteLogQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExtendContentPacked) {
		body["ExtendContentPacked"] = request.ExtendContentPacked
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogProjectName) {
		body["LogProjectName"] = request.LogProjectName
	}

	if !dara.IsNil(request.LogQuery) {
		body["LogQuery"] = request.LogQuery
	}

	if !dara.IsNil(request.LogRegionId) {
		body["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogStoreName) {
		body["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.NormalizationSchemaId) {
		body["NormalizationSchemaId"] = request.NormalizationSchemaId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteLogQuery"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteLogQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看LogStore
//
// @param request - ExecuteLogQueryRequest
//
// @return ExecuteLogQueryResponse
func (client *Client) ExecuteLogQuery(request *ExecuteLogQueryRequest) (_result *ExecuteLogQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteLogQueryResponse{}
	_body, _err := client.ExecuteLogQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行升级
//
// @param request - ExecuteUpgradeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteUpgradeResponse
func (client *Client) ExecuteUpgradeWithOptions(request *ExecuteUpgradeRequest, runtime *dara.RuntimeOptions) (_result *ExecuteUpgradeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteUpgrade"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteUpgradeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行升级
//
// @param request - ExecuteUpgradeRequest
//
// @return ExecuteUpgradeResponse
func (client *Client) ExecuteUpgrade(request *ExecuteUpgradeRequest) (_result *ExecuteUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteUpgradeResponse{}
	_body, _err := client.ExecuteUpgradeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据批量接入
//
// @param request - GetDataBatchIngestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataBatchIngestionResponse
func (client *Client) GetDataBatchIngestionWithOptions(request *GetDataBatchIngestionRequest, runtime *dara.RuntimeOptions) (_result *GetDataBatchIngestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataBatchIngestion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataBatchIngestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据批量接入
//
// @param request - GetDataBatchIngestionRequest
//
// @return GetDataBatchIngestionResponse
func (client *Client) GetDataBatchIngestion(request *GetDataBatchIngestionRequest) (_result *GetDataBatchIngestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataBatchIngestionResponse{}
	_body, _err := client.GetDataBatchIngestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日志管理页面里用户数据存储的详情。
//
// @param request - GetDataStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataStorageResponse
func (client *Client) GetDataStorageWithOptions(request *GetDataStorageRequest, runtime *dara.RuntimeOptions) (_result *GetDataStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataStorage"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志管理页面里用户数据存储的详情。
//
// @param request - GetDataStorageRequest
//
// @return GetDataStorageResponse
func (client *Client) GetDataStorage(request *GetDataStorageRequest) (_result *GetDataStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataStorageResponse{}
	_body, _err := client.GetDataStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新检测规则
//
// @param request - GetDetectionStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetectionStatisticResponse
func (client *Client) GetDetectionStatisticWithOptions(request *GetDetectionStatisticRequest, runtime *dara.RuntimeOptions) (_result *GetDetectionStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDetectionStatistic"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDetectionStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新检测规则
//
// @param request - GetDetectionStatisticRequest
//
// @return GetDetectionStatisticResponse
func (client *Client) GetDetectionStatistic(request *GetDetectionStatisticRequest) (_result *GetDetectionStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDetectionStatisticResponse{}
	_body, _err := client.GetDetectionStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取导出任务进度
//
// @param request - GetExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExportTaskResponse
func (client *Client) GetExportTaskWithOptions(request *GetExportTaskRequest, runtime *dara.RuntimeOptions) (_result *GetExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExportId) {
		body["ExportId"] = request.ExportId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExportTask"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取导出任务进度
//
// @param request - GetExportTaskRequest
//
// @return GetExportTaskResponse
func (client *Client) GetExportTask(request *GetExportTaskRequest) (_result *GetExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExportTaskResponse{}
	_body, _err := client.GetExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取事件列表
//
// @param request - GetIncidentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIncidentResponse
func (client *Client) GetIncidentWithOptions(request *GetIncidentRequest, runtime *dara.RuntimeOptions) (_result *GetIncidentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IncidentUuid) {
		body["IncidentUuid"] = request.IncidentUuid
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIncident"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIncidentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取事件列表
//
// @param request - GetIncidentRequest
//
// @return GetIncidentResponse
func (client *Client) GetIncident(request *GetIncidentRequest) (_result *GetIncidentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIncidentResponse{}
	_body, _err := client.GetIncidentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看LogStore
//
// @param request - GetLogTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogTicketResponse
func (client *Client) GetLogTicketWithOptions(request *GetLogTicketRequest, runtime *dara.RuntimeOptions) (_result *GetLogTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLogTicket"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLogTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看LogStore
//
// @param request - GetLogTicketRequest
//
// @return GetLogTicketResponse
func (client *Client) GetLogTicket(request *GetLogTicketRequest) (_result *GetLogTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLogTicketResponse{}
	_body, _err := client.GetLogTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标准化规则
//
// @param request - GetNormalizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNormalizationRuleResponse
func (client *Client) GetNormalizationRuleWithOptions(request *GetNormalizationRuleRequest, runtime *dara.RuntimeOptions) (_result *GetNormalizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNormalizationRule"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNormalizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标准化规则
//
// @param request - GetNormalizationRuleRequest
//
// @return GetNormalizationRuleResponse
func (client *Client) GetNormalizationRule(request *GetNormalizationRuleRequest) (_result *GetNormalizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNormalizationRuleResponse{}
	_body, _err := client.GetNormalizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标准化规则指定版本信息
//
// @param request - GetNormalizationRuleVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNormalizationRuleVersionResponse
func (client *Client) GetNormalizationRuleVersionWithOptions(request *GetNormalizationRuleVersionRequest, runtime *dara.RuntimeOptions) (_result *GetNormalizationRuleVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.NormalizationRuleVersion) {
		body["NormalizationRuleVersion"] = request.NormalizationRuleVersion
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNormalizationRuleVersion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNormalizationRuleVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标准化规则指定版本信息
//
// @param request - GetNormalizationRuleVersionRequest
//
// @return GetNormalizationRuleVersionResponse
func (client *Client) GetNormalizationRuleVersion(request *GetNormalizationRuleVersionRequest) (_result *GetNormalizationRuleVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNormalizationRuleVersionResponse{}
	_body, _err := client.GetNormalizationRuleVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Schema信息以及字段
//
// @param request - GetNormalizationSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNormalizationSchemaResponse
func (client *Client) GetNormalizationSchemaWithOptions(request *GetNormalizationSchemaRequest, runtime *dara.RuntimeOptions) (_result *GetNormalizationSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationSchemaId) {
		body["NormalizationSchemaId"] = request.NormalizationSchemaId
	}

	if !dara.IsNil(request.NormalizationSchemaType) {
		body["NormalizationSchemaType"] = request.NormalizationSchemaType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNormalizationSchema"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNormalizationSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Schema信息以及字段
//
// @param request - GetNormalizationSchemaRequest
//
// @return GetNormalizationSchemaResponse
func (client *Client) GetNormalizationSchema(request *GetNormalizationSchemaRequest) (_result *GetNormalizationSchemaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNormalizationSchemaResponse{}
	_body, _err := client.GetNormalizationSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户配置信息
//
// @param request - GetUserConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserConfigResponse
func (client *Client) GetUserConfigWithOptions(request *GetUserConfigRequest, runtime *dara.RuntimeOptions) (_result *GetUserConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserConfig"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户配置信息
//
// @param request - GetUserConfigRequest
//
// @return GetUserConfigResponse
func (client *Client) GetUserConfig(request *GetUserConfigRequest) (_result *GetUserConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserConfigResponse{}
	_body, _err := client.GetUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询接入模板
//
// @param request - ListDataIngestionTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataIngestionTemplatesResponse
func (client *Client) ListDataIngestionTemplatesWithOptions(request *ListDataIngestionTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListDataIngestionTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataIngestionTemplateStatus) {
		body["DataIngestionTemplateStatus"] = request.DataIngestionTemplateStatus
	}

	if !dara.IsNil(request.DataSourceTemplateIds) {
		body["DataSourceTemplateIds"] = request.DataSourceTemplateIds
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataIngestionTemplates"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataIngestionTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询接入模板
//
// @param request - ListDataIngestionTemplatesRequest
//
// @return ListDataIngestionTemplatesResponse
func (client *Client) ListDataIngestionTemplates(request *ListDataIngestionTemplatesRequest) (_result *ListDataIngestionTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataIngestionTemplatesResponse{}
	_body, _err := client.ListDataIngestionTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据接入任务列表
//
// @param tmpReq - ListDataIngestionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataIngestionsResponse
func (client *Client) ListDataIngestionsWithOptions(tmpReq *ListDataIngestionsRequest, runtime *dara.RuntimeOptions) (_result *ListDataIngestionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataIngestionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataIngestionIds) {
		request.DataIngestionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataIngestionIds, dara.String("DataIngestionIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DataIngestionTemplateIds) {
		request.DataIngestionTemplateIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataIngestionTemplateIds, dara.String("DataIngestionTemplateIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataIngestionIdsShrink) {
		body["DataIngestionIds"] = request.DataIngestionIdsShrink
	}

	if !dara.IsNil(request.DataIngestionStatus) {
		body["DataIngestionStatus"] = request.DataIngestionStatus
	}

	if !dara.IsNil(request.DataIngestionTemplateIdsShrink) {
		body["DataIngestionTemplateIds"] = request.DataIngestionTemplateIdsShrink
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataIngestions"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataIngestionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据接入任务列表
//
// @param request - ListDataIngestionsRequest
//
// @return ListDataIngestionsResponse
func (client *Client) ListDataIngestions(request *ListDataIngestionsRequest) (_result *ListDataIngestionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataIngestionsResponse{}
	_body, _err := client.ListDataIngestionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集记录列表
//
// @param request - ListDataSetRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSetRecordsResponse
func (client *Client) ListDataSetRecordsWithOptions(request *ListDataSetRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListDataSetRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSetId) {
		body["DataSetId"] = request.DataSetId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSetRecords"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSetRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集记录列表
//
// @param request - ListDataSetRecordsRequest
//
// @return ListDataSetRecordsResponse
func (client *Client) ListDataSetRecords(request *ListDataSetRecordsRequest) (_result *ListDataSetRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSetRecordsResponse{}
	_body, _err := client.ListDataSetRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集列表
//
// @param tmpReq - ListDataSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSetsResponse
func (client *Client) ListDataSetsWithOptions(tmpReq *ListDataSetsRequest, runtime *dara.RuntimeOptions) (_result *ListDataSetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataSetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSetIds) {
		request.DataSetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSetIds, dara.String("DataSetIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSetId) {
		body["DataSetId"] = request.DataSetId
	}

	if !dara.IsNil(request.DataSetIdsShrink) {
		body["DataSetIds"] = request.DataSetIdsShrink
	}

	if !dara.IsNil(request.DataSetName) {
		body["DataSetName"] = request.DataSetName
	}

	if !dara.IsNil(request.DataSetStatus) {
		body["DataSetStatus"] = request.DataSetStatus
	}

	if !dara.IsNil(request.DataSetType) {
		body["DataSetType"] = request.DataSetType
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderDirection) {
		body["OrderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.OrderFieldName) {
		body["OrderFieldName"] = request.OrderFieldName
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSets"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集列表
//
// @param request - ListDataSetsRequest
//
// @return ListDataSetsResponse
func (client *Client) ListDataSets(request *ListDataSetsRequest) (_result *ListDataSetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSetsResponse{}
	_body, _err := client.ListDataSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据源模板
//
// @param tmpReq - ListDataSourceTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceTemplatesResponse
func (client *Client) ListDataSourceTemplatesWithOptions(tmpReq *ListDataSourceTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataSourceTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSourceTemplateIds) {
		request.DataSourceTemplateIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSourceTemplateIds, dara.String("DataSourceTemplateIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceTemplateIdsShrink) {
		body["DataSourceTemplateIds"] = request.DataSourceTemplateIdsShrink
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSourceTemplates"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据源模板
//
// @param request - ListDataSourceTemplatesRequest
//
// @return ListDataSourceTemplatesResponse
func (client *Client) ListDataSourceTemplates(request *ListDataSourceTemplatesRequest) (_result *ListDataSourceTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSourceTemplatesResponse{}
	_body, _err := client.ListDataSourceTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取厂商列表
//
// @param tmpReq - ListDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithOptions(tmpReq *ListDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSourceIds) {
		request.DataSourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSourceIds, dara.String("DataSourceIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.DataSourceTemplateIds) {
		request.DataSourceTemplateIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSourceTemplateIds, dara.String("DataSourceTemplateIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.LogUserIds) {
		request.LogUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LogUserIds, dara.String("LogUserIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceFrom) {
		body["DataSourceFrom"] = request.DataSourceFrom
	}

	if !dara.IsNil(request.DataSourceIdsShrink) {
		body["DataSourceIds"] = request.DataSourceIdsShrink
	}

	if !dara.IsNil(request.DataSourceName) {
		body["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.DataSourceStatus) {
		body["DataSourceStatus"] = request.DataSourceStatus
	}

	if !dara.IsNil(request.DataSourceStoreStatus) {
		body["DataSourceStoreStatus"] = request.DataSourceStoreStatus
	}

	if !dara.IsNil(request.DataSourceTemplateIdsShrink) {
		body["DataSourceTemplateIds"] = request.DataSourceTemplateIdsShrink
	}

	if !dara.IsNil(request.DataSourceType) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogProjectName) {
		body["LogProjectName"] = request.LogProjectName
	}

	if !dara.IsNil(request.LogRegionId) {
		body["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogStoreName) {
		body["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.LogUserIdsShrink) {
		body["LogUserIds"] = request.LogUserIdsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSources"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取厂商列表
//
// @param request - ListDataSourcesRequest
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSources(request *ListDataSourcesRequest) (_result *ListDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取检测规则列表
//
// @param tmpReq - ListDetectionRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDetectionRulesResponse
func (client *Client) ListDetectionRulesWithOptions(tmpReq *ListDetectionRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDetectionRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDetectionRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DetectionRuleIds) {
		request.DetectionRuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DetectionRuleIds, dara.String("DetectionRuleIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertAttCk) {
		body["AlertAttCk"] = request.AlertAttCk
	}

	if !dara.IsNil(request.AlertLevel) {
		body["AlertLevel"] = request.AlertLevel
	}

	if !dara.IsNil(request.AlertTacticId) {
		body["AlertTacticId"] = request.AlertTacticId
	}

	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.DetectionExpressionType) {
		body["DetectionExpressionType"] = request.DetectionExpressionType
	}

	if !dara.IsNil(request.DetectionRuleId) {
		body["DetectionRuleId"] = request.DetectionRuleId
	}

	if !dara.IsNil(request.DetectionRuleIdsShrink) {
		body["DetectionRuleIds"] = request.DetectionRuleIdsShrink
	}

	if !dara.IsNil(request.DetectionRuleName) {
		body["DetectionRuleName"] = request.DetectionRuleName
	}

	if !dara.IsNil(request.DetectionRuleStatus) {
		body["DetectionRuleStatus"] = request.DetectionRuleStatus
	}

	if !dara.IsNil(request.DetectionRuleType) {
		body["DetectionRuleType"] = request.DetectionRuleType
	}

	if !dara.IsNil(request.IncidentAggregationType) {
		body["IncidentAggregationType"] = request.IncidentAggregationType
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogCategoryId) {
		body["LogCategoryId"] = request.LogCategoryId
	}

	if !dara.IsNil(request.LogSchemaId) {
		body["LogSchemaId"] = request.LogSchemaId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderDirection) {
		body["OrderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.OrderFieldName) {
		body["OrderFieldName"] = request.OrderFieldName
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDetectionRules"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDetectionRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取检测规则列表
//
// @param request - ListDetectionRulesRequest
//
// @return ListDetectionRulesResponse
func (client *Client) ListDetectionRules(request *ListDetectionRulesRequest) (_result *ListDetectionRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDetectionRulesResponse{}
	_body, _err := client.ListDetectionRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取事件列表
//
// @param tmpReq - ListIncidentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIncidentsResponse
func (client *Client) ListIncidentsWithOptions(tmpReq *ListIncidentsRequest, runtime *dara.RuntimeOptions) (_result *ListIncidentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListIncidentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IncidentUuids) {
		request.IncidentUuidsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IncidentUuids, dara.String("IncidentUuids"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IncidentName) {
		query["IncidentName"] = request.IncidentName
	}

	if !dara.IsNil(request.IncidentUuidsShrink) {
		query["IncidentUuids"] = request.IncidentUuidsShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertUuid) {
		body["AlertUuid"] = request.AlertUuid
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IncidentStatus) {
		body["IncidentStatus"] = request.IncidentStatus
	}

	if !dara.IsNil(request.IncidentTags) {
		body["IncidentTags"] = request.IncidentTags
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderDirection) {
		body["OrderDirection"] = request.OrderDirection
	}

	if !dara.IsNil(request.OrderFieldName) {
		body["OrderFieldName"] = request.OrderFieldName
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RelateAssetId) {
		body["RelateAssetId"] = request.RelateAssetId
	}

	if !dara.IsNil(request.RelateEntityId) {
		body["RelateEntityId"] = request.RelateEntityId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.RoleType) {
		body["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.ThreatLevel) {
		body["ThreatLevel"] = request.ThreatLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIncidents"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIncidentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取事件列表
//
// @param request - ListIncidentsRequest
//
// @return ListIncidentsResponse
func (client *Client) ListIncidents(request *ListIncidentsRequest) (_result *ListIncidentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIncidentsResponse{}
	_body, _err := client.ListIncidentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日志Project列表
//
// @param request - ListLogProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogProjectsResponse
func (client *Client) ListLogProjectsWithOptions(request *ListLogProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListLogProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogRegionId) {
		body["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogProjects"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志Project列表
//
// @param request - ListLogProjectsRequest
//
// @return ListLogProjectsResponse
func (client *Client) ListLogProjects(request *ListLogProjectsRequest) (_result *ListLogProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLogProjectsResponse{}
	_body, _err := client.ListLogProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取所有的区域
//
// @param request - ListLogRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogRegionsResponse
func (client *Client) ListLogRegionsWithOptions(request *ListLogRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListLogRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogRegions"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取所有的区域
//
// @param request - ListLogRegionsRequest
//
// @return ListLogRegionsResponse
func (client *Client) ListLogRegions(request *ListLogRegionsRequest) (_result *ListLogRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLogRegionsResponse{}
	_body, _err := client.ListLogRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日志store列表
//
// @param request - ListLogStoresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogStoresResponse
func (client *Client) ListLogStoresWithOptions(request *ListLogStoresRequest, runtime *dara.RuntimeOptions) (_result *ListLogStoresResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogProjectName) {
		body["LogProjectName"] = request.LogProjectName
	}

	if !dara.IsNil(request.LogRegionId) {
		body["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogStores"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogStoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志store列表
//
// @param request - ListLogStoresRequest
//
// @return ListLogStoresResponse
func (client *Client) ListLogStores(request *ListLogStoresRequest) (_result *ListLogStoresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLogStoresResponse{}
	_body, _err := client.ListLogStoresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标准化目录
//
// @param request - ListNormalizationCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNormalizationCategoriesResponse
func (client *Client) ListNormalizationCategoriesWithOptions(request *ListNormalizationCategoriesRequest, runtime *dara.RuntimeOptions) (_result *ListNormalizationCategoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NormalizationCategoryType) {
		body["NormalizationCategoryType"] = request.NormalizationCategoryType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNormalizationCategories"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNormalizationCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标准化目录
//
// @param request - ListNormalizationCategoriesRequest
//
// @return ListNormalizationCategoriesResponse
func (client *Client) ListNormalizationCategories(request *ListNormalizationCategoriesRequest) (_result *ListNormalizationCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNormalizationCategoriesResponse{}
	_body, _err := client.ListNormalizationCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标准化日志所有字段
//
// @param request - ListNormalizationFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNormalizationFieldsResponse
func (client *Client) ListNormalizationFieldsWithOptions(request *ListNormalizationFieldsRequest, runtime *dara.RuntimeOptions) (_result *ListNormalizationFieldsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
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

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNormalizationFields"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNormalizationFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标准化日志所有字段
//
// @param request - ListNormalizationFieldsRequest
//
// @return ListNormalizationFieldsResponse
func (client *Client) ListNormalizationFields(request *ListNormalizationFieldsRequest) (_result *ListNormalizationFieldsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNormalizationFieldsResponse{}
	_body, _err := client.ListNormalizationFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取规则的安全能力
//
// @param tmpReq - ListNormalizationRuleCapacitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNormalizationRuleCapacitiesResponse
func (client *Client) ListNormalizationRuleCapacitiesWithOptions(tmpReq *ListNormalizationRuleCapacitiesRequest, runtime *dara.RuntimeOptions) (_result *ListNormalizationRuleCapacitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNormalizationRuleCapacitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NormalizationRuleIds) {
		request.NormalizationRuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NormalizationRuleIds, dara.String("NormalizationRuleIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.NormalizationRuleIdsShrink) {
		body["NormalizationRuleIds"] = request.NormalizationRuleIdsShrink
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNormalizationRuleCapacities"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNormalizationRuleCapacitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取规则的安全能力
//
// @param request - ListNormalizationRuleCapacitiesRequest
//
// @return ListNormalizationRuleCapacitiesResponse
func (client *Client) ListNormalizationRuleCapacities(request *ListNormalizationRuleCapacitiesRequest) (_result *ListNormalizationRuleCapacitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNormalizationRuleCapacitiesResponse{}
	_body, _err := client.ListNormalizationRuleCapacitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标准化规则版本列表
//
// @param request - ListNormalizationRuleVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNormalizationRuleVersionsResponse
func (client *Client) ListNormalizationRuleVersionsWithOptions(request *ListNormalizationRuleVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListNormalizationRuleVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNormalizationRuleVersions"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNormalizationRuleVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标准化规则版本列表
//
// @param request - ListNormalizationRuleVersionsRequest
//
// @return ListNormalizationRuleVersionsResponse
func (client *Client) ListNormalizationRuleVersions(request *ListNormalizationRuleVersionsRequest) (_result *ListNormalizationRuleVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNormalizationRuleVersionsResponse{}
	_body, _err := client.ListNormalizationRuleVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标准化规则列表
//
// @param tmpReq - ListNormalizationRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNormalizationRulesResponse
func (client *Client) ListNormalizationRulesWithOptions(tmpReq *ListNormalizationRulesRequest, runtime *dara.RuntimeOptions) (_result *ListNormalizationRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListNormalizationRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NormalizationRuleIds) {
		request.NormalizationRuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NormalizationRuleIds, dara.String("NormalizationRuleIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NormalizationCategoryId) {
		body["NormalizationCategoryId"] = request.NormalizationCategoryId
	}

	if !dara.IsNil(request.NormalizationRuleIdsShrink) {
		body["NormalizationRuleIds"] = request.NormalizationRuleIdsShrink
	}

	if !dara.IsNil(request.NormalizationRuleName) {
		body["NormalizationRuleName"] = request.NormalizationRuleName
	}

	if !dara.IsNil(request.NormalizationRuleType) {
		body["NormalizationRuleType"] = request.NormalizationRuleType
	}

	if !dara.IsNil(request.NormalizationSchemaId) {
		body["NormalizationSchemaId"] = request.NormalizationSchemaId
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.OrderType) {
		body["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorId) {
		body["VendorId"] = request.VendorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNormalizationRules"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNormalizationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标准化规则列表
//
// @param request - ListNormalizationRulesRequest
//
// @return ListNormalizationRulesResponse
func (client *Client) ListNormalizationRules(request *ListNormalizationRulesRequest) (_result *ListNormalizationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNormalizationRulesResponse{}
	_body, _err := client.ListNormalizationRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标准化类目
//
// @param request - ListNormalizationSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNormalizationSchemasResponse
func (client *Client) ListNormalizationSchemasWithOptions(request *ListNormalizationSchemasRequest, runtime *dara.RuntimeOptions) (_result *ListNormalizationSchemasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NormalizationCategoryId) {
		body["NormalizationCategoryId"] = request.NormalizationCategoryId
	}

	if !dara.IsNil(request.NormalizationSchemaType) {
		body["NormalizationSchemaType"] = request.NormalizationSchemaType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNormalizationSchemas"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNormalizationSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标准化类目
//
// @param request - ListNormalizationSchemasRequest
//
// @return ListNormalizationSchemasResponse
func (client *Client) ListNormalizationSchemas(request *ListNormalizationSchemasRequest) (_result *ListNormalizationSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNormalizationSchemasResponse{}
	_body, _err := client.ListNormalizationSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取产品列表
//
// @param tmpReq - ListProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithOptions(tmpReq *ListProductsRequest, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProductsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ProductIds) {
		request.ProductIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProductIds, dara.String("ProductIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProductIdsShrink) {
		body["ProductIds"] = request.ProductIdsShrink
	}

	if !dara.IsNil(request.ProductName) {
		body["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorId) {
		body["VendorId"] = request.VendorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProducts"),
		Version:     dara.String("2024-12-12"),
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
// 获取产品列表
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
// 获取接入流量统计
//
// @param tmpReq - ListTrafficStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficStatisticsResponse
func (client *Client) ListTrafficStatisticsWithOptions(tmpReq *ListTrafficStatisticsRequest, runtime *dara.RuntimeOptions) (_result *ListTrafficStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTrafficStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LogUserIds) {
		request.LogUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LogUserIds, dara.String("LogUserIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogUserIdsShrink) {
		body["LogUserIds"] = request.LogUserIdsShrink
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionTag) {
		body["RegionTag"] = request.RegionTag
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.TrafficStatisticPeriod) {
		body["TrafficStatisticPeriod"] = request.TrafficStatisticPeriod
	}

	if !dara.IsNil(request.TrafficStatisticPeriodType) {
		body["TrafficStatisticPeriodType"] = request.TrafficStatisticPeriodType
	}

	if !dara.IsNil(request.TrafficStatisticType) {
		body["TrafficStatisticType"] = request.TrafficStatisticType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrafficStatistics"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrafficStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取接入流量统计
//
// @param request - ListTrafficStatisticsRequest
//
// @return ListTrafficStatisticsResponse
func (client *Client) ListTrafficStatistics(request *ListTrafficStatisticsRequest) (_result *ListTrafficStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTrafficStatisticsResponse{}
	_body, _err := client.ListTrafficStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取升级项列表
//
// @param request - ListUpgradeItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUpgradeItemsResponse
func (client *Client) ListUpgradeItemsWithOptions(request *ListUpgradeItemsRequest, runtime *dara.RuntimeOptions) (_result *ListUpgradeItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUpgradeItems"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUpgradeItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取升级项列表
//
// @param request - ListUpgradeItemsRequest
//
// @return ListUpgradeItemsResponse
func (client *Client) ListUpgradeItems(request *ListUpgradeItemsRequest) (_result *ListUpgradeItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUpgradeItemsResponse{}
	_body, _err := client.ListUpgradeItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取厂商列表
//
// @param tmpReq - ListVendorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVendorsResponse
func (client *Client) ListVendorsWithOptions(tmpReq *ListVendorsRequest, runtime *dara.RuntimeOptions) (_result *ListVendorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListVendorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.VendorIds) {
		request.VendorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VendorIds, dara.String("VendorIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorIdsShrink) {
		body["VendorIds"] = request.VendorIdsShrink
	}

	if !dara.IsNil(request.VendorName) {
		body["VendorName"] = request.VendorName
	}

	if !dara.IsNil(request.VendorType) {
		body["VendorType"] = request.VendorType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVendors"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVendorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取厂商列表
//
// @param request - ListVendorsRequest
//
// @return ListVendorsResponse
func (client *Client) ListVendors(request *ListVendorsRequest) (_result *ListVendorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVendorsResponse{}
	_body, _err := client.ListVendorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据存储的清空操作，该动作会删除已有的数据，重新初始化物理存储。
//
// @param request - ResetDataStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetDataStorageResponse
func (client *Client) ResetDataStorageWithOptions(request *ResetDataStorageRequest, runtime *dara.RuntimeOptions) (_result *ResetDataStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetDataStorage"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetDataStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据存储的清空操作，该动作会删除已有的数据，重新初始化物理存储。
//
// @param request - ResetDataStorageRequest
//
// @return ResetDataStorageResponse
func (client *Client) ResetDataStorage(request *ResetDataStorageRequest) (_result *ResetDataStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetDataStorageResponse{}
	_body, _err := client.ResetDataStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置标准化规则默认版本
//
// @param request - SetDefaultNormalizationRuleVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultNormalizationRuleVersionResponse
func (client *Client) SetDefaultNormalizationRuleVersionWithOptions(request *SetDefaultNormalizationRuleVersionRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultNormalizationRuleVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.NormalizationRuleVersion) {
		body["NormalizationRuleVersion"] = request.NormalizationRuleVersion
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultNormalizationRuleVersion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultNormalizationRuleVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置标准化规则默认版本
//
// @param request - SetDefaultNormalizationRuleVersionRequest
//
// @return SetDefaultNormalizationRuleVersionResponse
func (client *Client) SetDefaultNormalizationRuleVersion(request *SetDefaultNormalizationRuleVersionRequest) (_result *SetDefaultNormalizationRuleVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDefaultNormalizationRuleVersionResponse{}
	_body, _err := client.SetDefaultNormalizationRuleVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据批量接入
//
// @param tmpReq - UpdateDataBatchIngestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataBatchIngestionResponse
func (client *Client) UpdateDataBatchIngestionWithOptions(tmpReq *UpdateDataBatchIngestionRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataBatchIngestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataBatchIngestionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataIngestionIds) {
		request.DataIngestionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataIngestionIds, dara.String("DataIngestionIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.LogUserIds) {
		request.LogUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LogUserIds, dara.String("LogUserIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoScanNew) {
		body["AutoScanNew"] = request.AutoScanNew
	}

	if !dara.IsNil(request.DataBatchIngestionMode) {
		body["DataBatchIngestionMode"] = request.DataBatchIngestionMode
	}

	if !dara.IsNil(request.DataIngestionIdsShrink) {
		body["DataIngestionIds"] = request.DataIngestionIdsShrink
	}

	if !dara.IsNil(request.DataSourceRecognizeEnabled) {
		body["DataSourceRecognizeEnabled"] = request.DataSourceRecognizeEnabled
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogUserIdsShrink) {
		body["LogUserIds"] = request.LogUserIdsShrink
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataBatchIngestion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataBatchIngestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据批量接入
//
// @param request - UpdateDataBatchIngestionRequest
//
// @return UpdateDataBatchIngestionResponse
func (client *Client) UpdateDataBatchIngestion(request *UpdateDataBatchIngestionRequest) (_result *UpdateDataBatchIngestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataBatchIngestionResponse{}
	_body, _err := client.UpdateDataBatchIngestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据接入信息
//
// @param request - UpdateDataIngestionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataIngestionResponse
func (client *Client) UpdateDataIngestionWithOptions(request *UpdateDataIngestionRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataIngestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataIngestionId) {
		body["DataIngestionId"] = request.DataIngestionId
	}

	if !dara.IsNil(request.DataIngestionMode) {
		body["DataIngestionMode"] = request.DataIngestionMode
	}

	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataIngestion"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataIngestionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据接入信息
//
// @param request - UpdateDataIngestionRequest
//
// @return UpdateDataIngestionResponse
func (client *Client) UpdateDataIngestion(request *UpdateDataIngestionRequest) (_result *UpdateDataIngestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataIngestionResponse{}
	_body, _err := client.UpdateDataIngestionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新接入模板
//
// @param request - UpdateDataIngestionTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataIngestionTemplateResponse
func (client *Client) UpdateDataIngestionTemplateWithOptions(request *UpdateDataIngestionTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataIngestionTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataIngestionStatus) {
		body["DataIngestionStatus"] = request.DataIngestionStatus
	}

	if !dara.IsNil(request.DataIngestionTemplateId) {
		body["DataIngestionTemplateId"] = request.DataIngestionTemplateId
	}

	if !dara.IsNil(request.DataIngestionTemplateName) {
		body["DataIngestionTemplateName"] = request.DataIngestionTemplateName
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataIngestionTemplate"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataIngestionTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新接入模板
//
// @param request - UpdateDataIngestionTemplateRequest
//
// @return UpdateDataIngestionTemplateResponse
func (client *Client) UpdateDataIngestionTemplate(request *UpdateDataIngestionTemplateRequest) (_result *UpdateDataIngestionTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataIngestionTemplateResponse{}
	_body, _err := client.UpdateDataIngestionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据集
//
// @param request - UpdateDataSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSetResponse
func (client *Client) UpdateDataSetWithOptions(request *UpdateDataSetRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSetDescription) {
		body["DataSetDescription"] = request.DataSetDescription
	}

	if !dara.IsNil(request.DataSetFileName) {
		body["DataSetFileName"] = request.DataSetFileName
	}

	if !dara.IsNil(request.DataSetId) {
		body["DataSetId"] = request.DataSetId
	}

	if !dara.IsNil(request.DataSetName) {
		body["DataSetName"] = request.DataSetName
	}

	if !dara.IsNil(request.DataSetStatus) {
		body["DataSetStatus"] = request.DataSetStatus
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.IpWhitelistRecognizers) {
		bodyFlat["IpWhitelistRecognizers"] = request.IpWhitelistRecognizers
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSet"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据集
//
// @param request - UpdateDataSetRequest
//
// @return UpdateDataSetResponse
func (client *Client) UpdateDataSet(request *UpdateDataSetRequest) (_result *UpdateDataSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataSetResponse{}
	_body, _err := client.UpdateDataSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据集记录
//
// @param request - UpdateDataSetRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSetRecordResponse
func (client *Client) UpdateDataSetRecordWithOptions(request *UpdateDataSetRecordRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSetRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSetFileName) {
		body["DataSetFileName"] = request.DataSetFileName
	}

	if !dara.IsNil(request.DataSetId) {
		body["DataSetId"] = request.DataSetId
	}

	if !dara.IsNil(request.DataSetRecords) {
		body["DataSetRecords"] = request.DataSetRecords
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSetRecord"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSetRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据集记录
//
// @param request - UpdateDataSetRecordRequest
//
// @return UpdateDataSetRecordResponse
func (client *Client) UpdateDataSetRecord(request *UpdateDataSetRecordRequest) (_result *UpdateDataSetRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataSetRecordResponse{}
	_body, _err := client.UpdateDataSetRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据源
//
// @param request - UpdateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSourceWithOptions(request *UpdateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceFrom) {
		body["DataSourceFrom"] = request.DataSourceFrom
	}

	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.DataSourceName) {
		body["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.DataSourceRecognizeEnabled) {
		body["DataSourceRecognizeEnabled"] = request.DataSourceRecognizeEnabled
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceStores) {
		bodyFlat["DataSourceStores"] = request.DataSourceStores
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogProjectName) {
		body["LogProjectName"] = request.LogProjectName
	}

	if !dara.IsNil(request.LogRegionId) {
		body["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogStoreName) {
		body["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSource"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据源
//
// @param request - UpdateDataSourceRequest
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSource(request *UpdateDataSourceRequest) (_result *UpdateDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.UpdateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改数据源模板
//
// @param tmpReq - UpdateDataSourceTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceTemplateResponse
func (client *Client) UpdateDataSourceTemplateWithOptions(tmpReq *UpdateDataSourceTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataSourceTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LogUserIds) {
		request.LogUserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LogUserIds, dara.String("LogUserIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceRecognizeEnabled) {
		query["DataSourceRecognizeEnabled"] = request.DataSourceRecognizeEnabled
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoScanNew) {
		body["AutoScanNew"] = request.AutoScanNew
	}

	if !dara.IsNil(request.DataSourceTemplateId) {
		body["DataSourceTemplateId"] = request.DataSourceTemplateId
	}

	if !dara.IsNil(request.DataSourceTemplateName) {
		body["DataSourceTemplateName"] = request.DataSourceTemplateName
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogProjectPattern) {
		body["LogProjectPattern"] = request.LogProjectPattern
	}

	if !dara.IsNil(request.LogRegionIds) {
		body["LogRegionIds"] = request.LogRegionIds
	}

	if !dara.IsNil(request.LogStorePattern) {
		body["LogStorePattern"] = request.LogStorePattern
	}

	if !dara.IsNil(request.LogUserIdsShrink) {
		body["LogUserIds"] = request.LogUserIdsShrink
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSourceTemplate"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改数据源模板
//
// @param request - UpdateDataSourceTemplateRequest
//
// @return UpdateDataSourceTemplateResponse
func (client *Client) UpdateDataSourceTemplate(request *UpdateDataSourceTemplateRequest) (_result *UpdateDataSourceTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataSourceTemplateResponse{}
	_body, _err := client.UpdateDataSourceTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取日志管理页面里用户数据存储的详情。
//
// @param request - UpdateDataStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataStorageResponse
func (client *Client) UpdateDataStorageWithOptions(request *UpdateDataStorageRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataStorageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataStorageRegionId) {
		body["DataStorageRegionId"] = request.DataStorageRegionId
	}

	if !dara.IsNil(request.DeliveryStatus) {
		body["DeliveryStatus"] = request.DeliveryStatus
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataStorage"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取日志管理页面里用户数据存储的详情。
//
// @param request - UpdateDataStorageRequest
//
// @return UpdateDataStorageResponse
func (client *Client) UpdateDataStorage(request *UpdateDataStorageRequest) (_result *UpdateDataStorageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataStorageResponse{}
	_body, _err := client.UpdateDataStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 操作日志投递.
//
// @param request - UpdateDataStorageDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataStorageDeliveryResponse
func (client *Client) UpdateDataStorageDeliveryWithOptions(request *UpdateDataStorageDeliveryRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataStorageDeliveryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogCode) {
		body["LogCode"] = request.LogCode
	}

	if !dara.IsNil(request.LogDeliveryStatus) {
		body["LogDeliveryStatus"] = request.LogDeliveryStatus
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataStorageDelivery"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataStorageDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 操作日志投递.
//
// @param request - UpdateDataStorageDeliveryRequest
//
// @return UpdateDataStorageDeliveryResponse
func (client *Client) UpdateDataStorageDelivery(request *UpdateDataStorageDeliveryRequest) (_result *UpdateDataStorageDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataStorageDeliveryResponse{}
	_body, _err := client.UpdateDataStorageDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据存储中日志的数据保存天数。
//
// @param request - UpdateDataStorageTtlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataStorageTtlResponse
func (client *Client) UpdateDataStorageTtlWithOptions(request *UpdateDataStorageTtlRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataStorageTtlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogStoreColdTtl) {
		body["LogStoreColdTtl"] = request.LogStoreColdTtl
	}

	if !dara.IsNil(request.LogStoreHotTtl) {
		body["LogStoreHotTtl"] = request.LogStoreHotTtl
	}

	if !dara.IsNil(request.LogStoreName) {
		body["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.LogStoreTtl) {
		body["LogStoreTtl"] = request.LogStoreTtl
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataStorageTtl"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataStorageTtlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据存储中日志的数据保存天数。
//
// @param request - UpdateDataStorageTtlRequest
//
// @return UpdateDataStorageTtlResponse
func (client *Client) UpdateDataStorageTtl(request *UpdateDataStorageTtlRequest) (_result *UpdateDataStorageTtlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataStorageTtlResponse{}
	_body, _err := client.UpdateDataStorageTtlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新检测规则
//
// @param request - UpdateDetectionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDetectionRuleResponse
func (client *Client) UpdateDetectionRuleWithOptions(request *UpdateDetectionRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDetectionRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertAttCk) {
		body["AlertAttCk"] = request.AlertAttCk
	}

	if !dara.IsNil(request.AlertLevel) {
		body["AlertLevel"] = request.AlertLevel
	}

	if !dara.IsNil(request.AlertSchemaId) {
		body["AlertSchemaId"] = request.AlertSchemaId
	}

	if !dara.IsNil(request.AlertTacticId) {
		body["AlertTacticId"] = request.AlertTacticId
	}

	if !dara.IsNil(request.AlertThresholdCount) {
		body["AlertThresholdCount"] = request.AlertThresholdCount
	}

	if !dara.IsNil(request.AlertThresholdGroup) {
		body["AlertThresholdGroup"] = request.AlertThresholdGroup
	}

	if !dara.IsNil(request.AlertThresholdPeriod) {
		body["AlertThresholdPeriod"] = request.AlertThresholdPeriod
	}

	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.DetectionExpressionContent) {
		body["DetectionExpressionContent"] = request.DetectionExpressionContent
	}

	if !dara.IsNil(request.DetectionExpressionType) {
		body["DetectionExpressionType"] = request.DetectionExpressionType
	}

	if !dara.IsNil(request.DetectionRuleDescription) {
		body["DetectionRuleDescription"] = request.DetectionRuleDescription
	}

	if !dara.IsNil(request.DetectionRuleId) {
		body["DetectionRuleId"] = request.DetectionRuleId
	}

	if !dara.IsNil(request.DetectionRuleName) {
		body["DetectionRuleName"] = request.DetectionRuleName
	}

	if !dara.IsNil(request.DetectionRuleStatus) {
		body["DetectionRuleStatus"] = request.DetectionRuleStatus
	}

	if !dara.IsNil(request.DetectionRuleType) {
		body["DetectionRuleType"] = request.DetectionRuleType
	}

	if !dara.IsNil(request.EntityMappings) {
		body["EntityMappings"] = request.EntityMappings
	}

	if !dara.IsNil(request.IncidentAggregationExpression) {
		body["IncidentAggregationExpression"] = request.IncidentAggregationExpression
	}

	if !dara.IsNil(request.IncidentAggregationType) {
		body["IncidentAggregationType"] = request.IncidentAggregationType
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogCategoryId) {
		body["LogCategoryId"] = request.LogCategoryId
	}

	if !dara.IsNil(request.LogSchemaId) {
		body["LogSchemaId"] = request.LogSchemaId
	}

	if !dara.IsNil(request.PlaybookParameters) {
		body["PlaybookParameters"] = request.PlaybookParameters
	}

	if !dara.IsNil(request.PlaybookUuid) {
		body["PlaybookUuid"] = request.PlaybookUuid
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScheduleBeginTime) {
		body["ScheduleBeginTime"] = request.ScheduleBeginTime
	}

	if !dara.IsNil(request.ScheduleExpression) {
		body["ScheduleExpression"] = request.ScheduleExpression
	}

	if !dara.IsNil(request.ScheduleMaxRetries) {
		body["ScheduleMaxRetries"] = request.ScheduleMaxRetries
	}

	if !dara.IsNil(request.ScheduleMaxTimeout) {
		body["ScheduleMaxTimeout"] = request.ScheduleMaxTimeout
	}

	if !dara.IsNil(request.ScheduleType) {
		body["ScheduleType"] = request.ScheduleType
	}

	if !dara.IsNil(request.ScheduleWindow) {
		body["ScheduleWindow"] = request.ScheduleWindow
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDetectionRule"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDetectionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新检测规则
//
// @param request - UpdateDetectionRuleRequest
//
// @return UpdateDetectionRuleResponse
func (client *Client) UpdateDetectionRule(request *UpdateDetectionRuleRequest) (_result *UpdateDetectionRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDetectionRuleResponse{}
	_body, _err := client.UpdateDetectionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新标准化规则
//
// @param tmpReq - UpdateNormalizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNormalizationRuleResponse
func (client *Client) UpdateNormalizationRuleWithOptions(tmpReq *UpdateNormalizationRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateNormalizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateNormalizationRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NormalizationRuleIds) {
		request.NormalizationRuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NormalizationRuleIds, dara.String("NormalizationRuleIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtendContentPacked) {
		body["ExtendContentPacked"] = request.ExtendContentPacked
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationRuleDescription) {
		body["NormalizationRuleDescription"] = request.NormalizationRuleDescription
	}

	if !dara.IsNil(request.NormalizationRuleExpression) {
		body["NormalizationRuleExpression"] = request.NormalizationRuleExpression
	}

	if !dara.IsNil(request.NormalizationRuleFormat) {
		body["NormalizationRuleFormat"] = request.NormalizationRuleFormat
	}

	if !dara.IsNil(request.NormalizationRuleId) {
		body["NormalizationRuleId"] = request.NormalizationRuleId
	}

	if !dara.IsNil(request.NormalizationRuleIdsShrink) {
		body["NormalizationRuleIds"] = request.NormalizationRuleIdsShrink
	}

	if !dara.IsNil(request.NormalizationRuleMode) {
		body["NormalizationRuleMode"] = request.NormalizationRuleMode
	}

	if !dara.IsNil(request.NormalizationRuleName) {
		body["NormalizationRuleName"] = request.NormalizationRuleName
	}

	if !dara.IsNil(request.NormalizationRuleType) {
		body["NormalizationRuleType"] = request.NormalizationRuleType
	}

	if !dara.IsNil(request.NormalizationSchemaId) {
		body["NormalizationSchemaId"] = request.NormalizationSchemaId
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorId) {
		body["VendorId"] = request.VendorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNormalizationRule"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNormalizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新标准化规则
//
// @param request - UpdateNormalizationRuleRequest
//
// @return UpdateNormalizationRuleResponse
func (client *Client) UpdateNormalizationRule(request *UpdateNormalizationRuleRequest) (_result *UpdateNormalizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNormalizationRuleResponse{}
	_body, _err := client.UpdateNormalizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新产品品
//
// @param request - UpdateProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProductResponse
func (client *Client) UpdateProductWithOptions(request *UpdateProductRequest, runtime *dara.RuntimeOptions) (_result *UpdateProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductName) {
		body["ProductName"] = request.ProductName
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorName) {
		body["VendorName"] = request.VendorName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProduct"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新产品品
//
// @param request - UpdateProductRequest
//
// @return UpdateProductResponse
func (client *Client) UpdateProduct(request *UpdateProductRequest) (_result *UpdateProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProductResponse{}
	_body, _err := client.UpdateProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新厂商
//
// @param request - UpdateVendorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVendorResponse
func (client *Client) UpdateVendorWithOptions(request *UpdateVendorRequest, runtime *dara.RuntimeOptions) (_result *UpdateVendorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	if !dara.IsNil(request.VendorId) {
		body["VendorId"] = request.VendorId
	}

	if !dara.IsNil(request.VendorName) {
		body["VendorName"] = request.VendorName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVendor"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVendorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新厂商
//
// @param request - UpdateVendorRequest
//
// @return UpdateVendorResponse
func (client *Client) UpdateVendor(request *UpdateVendorRequest) (_result *UpdateVendorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVendorResponse{}
	_body, _err := client.UpdateVendorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验LogStore
//
// @param request - ValidateLogStoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateLogStoreResponse
func (client *Client) ValidateLogStoreWithOptions(request *ValidateLogStoreRequest, runtime *dara.RuntimeOptions) (_result *ValidateLogStoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.LogProjectName) {
		body["LogProjectName"] = request.LogProjectName
	}

	if !dara.IsNil(request.LogRegionId) {
		body["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogStoreName) {
		body["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.LogUserId) {
		body["LogUserId"] = request.LogUserId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateLogStore"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateLogStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验LogStore
//
// @param request - ValidateLogStoreRequest
//
// @return ValidateLogStoreResponse
func (client *Client) ValidateLogStore(request *ValidateLogStoreRequest) (_result *ValidateLogStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateLogStoreResponse{}
	_body, _err := client.ValidateLogStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验规则和数据
//
// @param request - ValidateNormalizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateNormalizationRuleResponse
func (client *Client) ValidateNormalizationRuleWithOptions(request *ValidateNormalizationRuleRequest, runtime *dara.RuntimeOptions) (_result *ValidateNormalizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.Lang) {
		body["Lang"] = request.Lang
	}

	if !dara.IsNil(request.NormalizationCategoryId) {
		body["NormalizationCategoryId"] = request.NormalizationCategoryId
	}

	if !dara.IsNil(request.NormalizationSchemaId) {
		body["NormalizationSchemaId"] = request.NormalizationSchemaId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleFor) {
		body["RoleFor"] = request.RoleFor
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateNormalizationRule"),
		Version:     dara.String("2024-12-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateNormalizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验规则和数据
//
// @param request - ValidateNormalizationRuleRequest
//
// @return ValidateNormalizationRuleResponse
func (client *Client) ValidateNormalizationRule(request *ValidateNormalizationRuleRequest) (_result *ValidateNormalizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateNormalizationRuleResponse{}
	_body, _err := client.ValidateNormalizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
