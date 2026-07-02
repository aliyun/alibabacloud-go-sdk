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
		"cn-hongkong":    dara.String("sddp-api.cn-hongkong.aliyuncs.com"),
		"cn-zhangjiakou": dara.String("sddp.cn-zhangjiakou.aliyuncs.com"),
		"ap-southeast-1": dara.String("sddp.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("sddp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
func (client *Client) CreateConfigWithOptions(request *CreateConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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
// @return CreateConfigResponse
func (client *Client) CreateConfig(request *CreateConfigRequest) (_result *CreateConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConfigResponse{}
	_body, _err := client.CreateConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the CreateDataLimit operation to grant permissions to scan databases, projects, and buckets.
//
// Description:
//
// You can use this operation to grant permissions to scan your data assets. This helps improve the security of your data assets.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, the system throttles your API calls. This may affect your business. Plan your calls accordingly.
//
// @param request - CreateDataLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLimitResponse
func (client *Client) CreateDataLimitWithOptions(request *CreateDataLimitRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLimitResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the CreateDataLimit operation to grant permissions to scan databases, projects, and buckets.
//
// Description:
//
// You can use this operation to grant permissions to scan your data assets. This helps improve the security of your data assets.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, the system throttles your API calls. This may affect your business. Plan your calls accordingly.
//
// @param request - CreateDataLimitRequest
//
// @return CreateDataLimitResponse
func (client *Client) CreateDataLimit(request *CreateDataLimitRequest) (_result *CreateDataLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataLimitResponse{}
	_body, _err := client.CreateDataLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call CreateRule to create a custom sensitive data detection rule.
//
// @param request - CreateRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRuleResponse
func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call CreateRule to create a custom sensitive data detection rule.
//
// @param request - CreateRuleRequest
//
// @return CreateRuleResponse
func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the CreateScanTask operation to create a custom scan task to detect sensitive data in authorized assets.
//
// Description:
//
// This operation creates custom scan tasks for authorized assets. You can control the run interval and runtime of each scan task.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for a single user. If you exceed this limit, API calls are throttled. This may impact your business. Plan your calls accordingly.
//
// @param request - CreateScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScanTaskResponse
func (client *Client) CreateScanTaskWithOptions(request *CreateScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateScanTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the CreateScanTask operation to create a custom scan task to detect sensitive data in authorized assets.
//
// Description:
//
// This operation creates custom scan tasks for authorized assets. You can control the run interval and runtime of each scan task.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for a single user. If you exceed this limit, API calls are throttled. This may impact your business. Plan your calls accordingly.
//
// @param request - CreateScanTaskRequest
//
// @return CreateScanTaskResponse
func (client *Client) CreateScanTask(request *CreateScanTaskRequest) (_result *CreateScanTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScanTaskResponse{}
	_body, _err := client.CreateScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateSlrRoleWithOptions(request *CreateSlrRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateSlrRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateSlrRoleResponse
func (client *Client) CreateSlrRole(request *CreateSlrRoleRequest) (_result *CreateSlrRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSlrRoleResponse{}
	_body, _err := client.CreateSlrRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the scan authorization for a data asset, such as a database, instance, or bucket.
//
// Description:
//
// This operation is typically used to revoke authorization for data assets. This helps you manage data access permissions.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed this limit, API calls are throttled. This may affect your business. We recommend that you call this operation within this limit.
//
// @param request - DeleteDataLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLimitResponse
func (client *Client) DeleteDataLimitWithOptions(request *DeleteDataLimitRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLimitResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the scan authorization for a data asset, such as a database, instance, or bucket.
//
// Description:
//
// This operation is typically used to revoke authorization for data assets. This helps you manage data access permissions.
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed this limit, API calls are throttled. This may affect your business. We recommend that you call this operation within this limit.
//
// @param request - DeleteDataLimitRequest
//
// @return DeleteDataLimitResponse
func (client *Client) DeleteDataLimit(request *DeleteDataLimitRequest) (_result *DeleteDataLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLimitResponse{}
	_body, _err := client.DeleteDataLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRuleResponse
func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeAuditLogsWithOptions(request *DescribeAuditLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditLogsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeAuditLogsResponse
func (client *Client) DescribeAuditLogs(request *DescribeAuditLogsRequest) (_result *DescribeAuditLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuditLogsResponse{}
	_body, _err := client.DescribeAuditLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists industry-specific templates.
//
// @param request - DescribeCategoryTemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCategoryTemplateListResponse
func (client *Client) DescribeCategoryTemplateListWithOptions(request *DescribeCategoryTemplateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCategoryTemplateListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists industry-specific templates.
//
// @param request - DescribeCategoryTemplateListRequest
//
// @return DescribeCategoryTemplateListResponse
func (client *Client) DescribeCategoryTemplateList(request *DescribeCategoryTemplateListRequest) (_result *DescribeCategoryTemplateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCategoryTemplateListResponse{}
	_body, _err := client.DescribeCategoryTemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCategoryTemplateRuleListWithOptions(request *DescribeCategoryTemplateRuleListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCategoryTemplateRuleListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCategoryTemplateRuleListResponse
func (client *Client) DescribeCategoryTemplateRuleList(request *DescribeCategoryTemplateRuleListRequest) (_result *DescribeCategoryTemplateRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCategoryTemplateRuleListResponse{}
	_body, _err := client.DescribeCategoryTemplateRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeColumns API to query column data in data asset tables, such as MaxCompute and RDS, that are authorized to connect to Data Security Center.
//
// Description:
//
// This API is typically used to view column data in sensitive data asset information tables. This helps users accurately analyze sensitive data.
//
// ## Notes
//
// The DescribeColumns API has been revised and replaced by DescribeColumnsV2. Use the newer DescribeColumnsV2 version when developing applications.
//
// ## QPS Limits
//
// The single-user QPS limit for this API is 10 calls per second. If you exceed this limit, API calls will be rate-limited. This may affect your business. You should call the API reasonably.
//
// @param request - DescribeColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnsResponse
func (client *Client) DescribeColumnsWithOptions(request *DescribeColumnsRequest, runtime *dara.RuntimeOptions) (_result *DescribeColumnsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeColumns API to query column data in data asset tables, such as MaxCompute and RDS, that are authorized to connect to Data Security Center.
//
// Description:
//
// This API is typically used to view column data in sensitive data asset information tables. This helps users accurately analyze sensitive data.
//
// ## Notes
//
// The DescribeColumns API has been revised and replaced by DescribeColumnsV2. Use the newer DescribeColumnsV2 version when developing applications.
//
// ## QPS Limits
//
// The single-user QPS limit for this API is 10 calls per second. If you exceed this limit, API calls will be rate-limited. This may affect your business. You should call the API reasonably.
//
// @param request - DescribeColumnsRequest
//
// @return DescribeColumnsResponse
func (client *Client) DescribeColumns(request *DescribeColumnsRequest) (_result *DescribeColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeColumnsResponse{}
	_body, _err := client.DescribeColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The DescribeColumnsV2 operation queries data in the columns of data asset tables, such as those in MaxCompute and RDS, that are authorized in Data Security Center.
//
// @param request - DescribeColumnsV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnsV2Response
func (client *Client) DescribeColumnsV2WithOptions(request *DescribeColumnsV2Request, runtime *dara.RuntimeOptions) (_result *DescribeColumnsV2Response, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The DescribeColumnsV2 operation queries data in the columns of data asset tables, such as those in MaxCompute and RDS, that are authorized in Data Security Center.
//
// @param request - DescribeColumnsV2Request
//
// @return DescribeColumnsV2Response
func (client *Client) DescribeColumnsV2(request *DescribeColumnsV2Request) (_result *DescribeColumnsV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeColumnsV2Response{}
	_body, _err := client.DescribeColumnsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeConfigsWithOptions(request *DescribeConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeConfigsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeConfigsResponse
func (client *Client) DescribeConfigs(request *DescribeConfigsRequest) (_result *DescribeConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeConfigsResponse{}
	_body, _err := client.DescribeConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDataAssetsWithOptions(request *DescribeDataAssetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataAssetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDataAssetsResponse
func (client *Client) DescribeDataAssets(request *DescribeDataAssetsRequest) (_result *DescribeDataAssetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataAssetsResponse{}
	_body, _err := client.DescribeDataAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDataLimitDetailWithOptions(request *DescribeDataLimitDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataLimitDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDataLimitDetailResponse
func (client *Client) DescribeDataLimitDetail(request *DescribeDataLimitDetailRequest) (_result *DescribeDataLimitDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataLimitDetailResponse{}
	_body, _err := client.DescribeDataLimitDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDataLimitSetWithOptions(request *DescribeDataLimitSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataLimitSetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDataLimitSetResponse
func (client *Client) DescribeDataLimitSet(request *DescribeDataLimitSetRequest) (_result *DescribeDataLimitSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataLimitSetResponse{}
	_body, _err := client.DescribeDataLimitSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of data assets for authorized instances, databases, and buckets.
//
// @param request - DescribeDataLimitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataLimitsResponse
func (client *Client) DescribeDataLimitsWithOptions(request *DescribeDataLimitsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataLimitsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data assets for authorized instances, databases, and buckets.
//
// @param request - DescribeDataLimitsRequest
//
// @return DescribeDataLimitsResponse
func (client *Client) DescribeDataLimits(request *DescribeDataLimitsRequest) (_result *DescribeDataLimitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataLimitsResponse{}
	_body, _err := client.DescribeDataLimitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call DescribeDataMaskingRunHistory to query the execution history of data masking tasks.
//
// Description:
//
// This operation retrieves the execution history of static data masking tasks. You can use it to search for task statuses and view task progress.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user on this operation is 10 calls per second. Calls that exceed this limit are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeDataMaskingRunHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataMaskingRunHistoryResponse
func (client *Client) DescribeDataMaskingRunHistoryWithOptions(request *DescribeDataMaskingRunHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataMaskingRunHistoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call DescribeDataMaskingRunHistory to query the execution history of data masking tasks.
//
// Description:
//
// This operation retrieves the execution history of static data masking tasks. You can use it to search for task statuses and view task progress.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user on this operation is 10 calls per second. Calls that exceed this limit are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - DescribeDataMaskingRunHistoryRequest
//
// @return DescribeDataMaskingRunHistoryResponse
func (client *Client) DescribeDataMaskingRunHistory(request *DescribeDataMaskingRunHistoryRequest) (_result *DescribeDataMaskingRunHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataMaskingRunHistoryResponse{}
	_body, _err := client.DescribeDataMaskingRunHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call DescribeDataMaskingTasks to retrieve a list of data masking tasks.
//
// Description:
//
// This operation retrieves a list of static data masking tasks, which you can then search and manage.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, your API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeDataMaskingTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataMaskingTasksResponse
func (client *Client) DescribeDataMaskingTasksWithOptions(request *DescribeDataMaskingTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataMaskingTasksResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call DescribeDataMaskingTasks to retrieve a list of data masking tasks.
//
// Description:
//
// This operation retrieves a list of static data masking tasks, which you can then search and manage.
//
// ## QPS limits
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, your API calls are throttled, which may affect your business. We recommend that you call this operation at a reasonable rate.
//
// @param request - DescribeDataMaskingTasksRequest
//
// @return DescribeDataMaskingTasksResponse
func (client *Client) DescribeDataMaskingTasks(request *DescribeDataMaskingTasksRequest) (_result *DescribeDataMaskingTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataMaskingTasksResponse{}
	_body, _err := client.DescribeDataMaskingTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detection results for columns in a data table.
//
// Description:
//
// ## Notes
//
// The DescribeDataObjectColumnDetail operation has been updated to DescribeDataObjectColumnDetailV2. We recommend that you use the latest version, DescribeDataObjectColumnDetailV2, for application development.
//
// @param request - DescribeDataObjectColumnDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataObjectColumnDetailResponse
func (client *Client) DescribeDataObjectColumnDetailWithOptions(request *DescribeDataObjectColumnDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataObjectColumnDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detection results for columns in a data table.
//
// Description:
//
// ## Notes
//
// The DescribeDataObjectColumnDetail operation has been updated to DescribeDataObjectColumnDetailV2. We recommend that you use the latest version, DescribeDataObjectColumnDetailV2, for application development.
//
// @param request - DescribeDataObjectColumnDetailRequest
//
// @return DescribeDataObjectColumnDetailResponse
func (client *Client) DescribeDataObjectColumnDetail(request *DescribeDataObjectColumnDetailRequest) (_result *DescribeDataObjectColumnDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataObjectColumnDetailResponse{}
	_body, _err := client.DescribeDataObjectColumnDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detection results for the columns of a data table.
//
// @param request - DescribeDataObjectColumnDetailV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataObjectColumnDetailV2Response
func (client *Client) DescribeDataObjectColumnDetailV2WithOptions(request *DescribeDataObjectColumnDetailV2Request, runtime *dara.RuntimeOptions) (_result *DescribeDataObjectColumnDetailV2Response, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detection results for the columns of a data table.
//
// @param request - DescribeDataObjectColumnDetailV2Request
//
// @return DescribeDataObjectColumnDetailV2Response
func (client *Client) DescribeDataObjectColumnDetailV2(request *DescribeDataObjectColumnDetailV2Request) (_result *DescribeDataObjectColumnDetailV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataObjectColumnDetailV2Response{}
	_body, _err := client.DescribeDataObjectColumnDetailV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query data detection results for tables and files.
//
// Description:
//
// This operation queries data detection results for tables and files, to provide a comprehensive view across all your assets.
//
// ## QPS limit
//
// The per-user QPS limit for this operation is 10 requests per second. If you exceed this limit, the system throttles your API calls. To prevent business disruptions, call this operation only when necessary.
//
// @param request - DescribeDataObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataObjectsResponse
func (client *Client) DescribeDataObjectsWithOptions(request *DescribeDataObjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataObjectsResponse, _err error) {
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

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query data detection results for tables and files.
//
// Description:
//
// This operation queries data detection results for tables and files, to provide a comprehensive view across all your assets.
//
// ## QPS limit
//
// The per-user QPS limit for this operation is 10 requests per second. If you exceed this limit, the system throttles your API calls. To prevent business disruptions, call this operation only when necessary.
//
// @param request - DescribeDataObjectsRequest
//
// @return DescribeDataObjectsResponse
func (client *Client) DescribeDataObjects(request *DescribeDataObjectsRequest) (_result *DescribeDataObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDataObjectsResponse{}
	_body, _err := client.DescribeDataObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of file types supported by Object Storage Service (OSS).
//
// @param request - DescribeDocTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocTypesResponse
func (client *Client) DescribeDocTypesWithOptions(request *DescribeDocTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDocTypesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of file types supported by Object Storage Service (OSS).
//
// @param request - DescribeDocTypesRequest
//
// @return DescribeDocTypesResponse
func (client *Client) DescribeDocTypes(request *DescribeDocTypesRequest) (_result *DescribeDocTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDocTypesResponse{}
	_body, _err := client.DescribeDocTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an anomalous event, including its occurrence time, description, and handling status.
//
// @param request - DescribeEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventDetailResponse
func (client *Client) DescribeEventDetailWithOptions(request *DescribeEventDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an anomalous event, including its occurrence time, description, and handling status.
//
// @param request - DescribeEventDetailRequest
//
// @return DescribeEventDetailResponse
func (client *Client) DescribeEventDetail(request *DescribeEventDetailRequest) (_result *DescribeEventDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventDetailResponse{}
	_body, _err := client.DescribeEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeEventTypesWithOptions(request *DescribeEventTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventTypesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeEventTypesResponse
func (client *Client) DescribeEventTypes(request *DescribeEventTypesRequest) (_result *DescribeEventTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventTypesResponse{}
	_body, _err := client.DescribeEventTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeEventsWithOptions(request *DescribeEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeEventsResponse
func (client *Client) DescribeEvents(request *DescribeEventsRequest) (_result *DescribeEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEventsResponse{}
	_body, _err := client.DescribeEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the completion status of a detection task based on the task ID. You can obtain the task ID from the Id field in the return value of a CreateScanTask or ScanOssObjectV1 API call.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 for each user. If you exceed the limit, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - DescribeIdentifyTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIdentifyTaskStatusResponse
func (client *Client) DescribeIdentifyTaskStatusWithOptions(request *DescribeIdentifyTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeIdentifyTaskStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the completion status of a detection task based on the task ID. You can obtain the task ID from the Id field in the return value of a CreateScanTask or ScanOssObjectV1 API call.
//
// Description:
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 for each user. If you exceed the limit, API calls are throttled, which may affect your business. Call this operation at a reasonable rate.
//
// @param request - DescribeIdentifyTaskStatusRequest
//
// @return DescribeIdentifyTaskStatusResponse
func (client *Client) DescribeIdentifyTaskStatus(request *DescribeIdentifyTaskStatusRequest) (_result *DescribeIdentifyTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIdentifyTaskStatusResponse{}
	_body, _err := client.DescribeIdentifyTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeInstanceSourcesWithOptions(request *DescribeInstanceSourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeInstanceSourcesResponse
func (client *Client) DescribeInstanceSources(request *DescribeInstanceSourcesRequest) (_result *DescribeInstanceSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceSourcesResponse{}
	_body, _err := client.DescribeInstanceSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of authorized MaxCompute, RDS, and OSS data asset instances.
//
// Description:
//
// When you call the DescribeInstances operation, you can set parameters such as search keywords and the threat level of data asset instances to retrieve a list of instances that meet your requirements.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user for this operation is 10 calls per second. If you exceed this limit, API calls are throttled. This can affect your business. Plan your calls accordingly.
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of authorized MaxCompute, RDS, and OSS data asset instances.
//
// Description:
//
// When you call the DescribeInstances operation, you can set parameters such as search keywords and the threat level of data asset instances to retrieve a list of instances that meet your requirements.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user for this operation is 10 calls per second. If you exceed this limit, API calls are throttled. This can affect your business. Plan your calls accordingly.
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains detailed information about an authorized OSS object in Data Security Center.
//
// Description:
//
// This operation is typically used to query the details of OSS objects. This information helps you accurately locate sensitive data assets in OSS.
//
// ## Usage notes
//
// The DescribeOssObjectDetail operation has been updated to DescribeOssObjectDetailV2. We recommend that you use the new version, DescribeOssObjectDetailV2, when you develop applications.
//
// ## QPS limit
//
// A single user can make up to 10 queries per second (QPS). If you exceed the limit, API calls are throttled. This may affect your business. We recommend that you plan your calls accordingly.
//
// @param request - DescribeOssObjectDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssObjectDetailResponse
func (client *Client) DescribeOssObjectDetailWithOptions(request *DescribeOssObjectDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssObjectDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains detailed information about an authorized OSS object in Data Security Center.
//
// Description:
//
// This operation is typically used to query the details of OSS objects. This information helps you accurately locate sensitive data assets in OSS.
//
// ## Usage notes
//
// The DescribeOssObjectDetail operation has been updated to DescribeOssObjectDetailV2. We recommend that you use the new version, DescribeOssObjectDetailV2, when you develop applications.
//
// ## QPS limit
//
// A single user can make up to 10 queries per second (QPS). If you exceed the limit, API calls are throttled. This may affect your business. We recommend that you plan your calls accordingly.
//
// @param request - DescribeOssObjectDetailRequest
//
// @return DescribeOssObjectDetailResponse
func (client *Client) DescribeOssObjectDetail(request *DescribeOssObjectDetailRequest) (_result *DescribeOssObjectDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOssObjectDetailResponse{}
	_body, _err := client.DescribeOssObjectDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains detailed information about an authorized OSS object in Data Security Center.
//
// Description:
//
// This operation queries the details of OSS objects. You can use this operation to locate sensitive data assets in OSS.
//
// @param request - DescribeOssObjectDetailV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssObjectDetailV2Response
func (client *Client) DescribeOssObjectDetailV2WithOptions(request *DescribeOssObjectDetailV2Request, runtime *dara.RuntimeOptions) (_result *DescribeOssObjectDetailV2Response, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains detailed information about an authorized OSS object in Data Security Center.
//
// Description:
//
// This operation queries the details of OSS objects. You can use this operation to locate sensitive data assets in OSS.
//
// @param request - DescribeOssObjectDetailV2Request
//
// @return DescribeOssObjectDetailV2Response
func (client *Client) DescribeOssObjectDetailV2(request *DescribeOssObjectDetailV2Request) (_result *DescribeOssObjectDetailV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOssObjectDetailV2Response{}
	_body, _err := client.DescribeOssObjectDetailV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeOssObjectsWithOptions(request *DescribeOssObjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssObjectsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeOssObjectsResponse
func (client *Client) DescribeOssObjects(request *DescribeOssObjectsRequest) (_result *DescribeOssObjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOssObjectsResponse{}
	_body, _err := client.DescribeOssObjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribePackagesWithOptions(request *DescribePackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribePackagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribePackagesResponse
func (client *Client) DescribePackages(request *DescribePackagesRequest) (_result *DescribePackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePackagesResponse{}
	_body, _err := client.DescribePackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists assets and their authorization status.
//
// @param request - DescribeParentInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParentInstanceResponse
func (client *Client) DescribeParentInstanceWithOptions(request *DescribeParentInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeParentInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists assets and their authorization status.
//
// @param request - DescribeParentInstanceRequest
//
// @return DescribeParentInstanceResponse
func (client *Client) DescribeParentInstance(request *DescribeParentInstanceRequest) (_result *DescribeParentInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeParentInstanceResponse{}
	_body, _err := client.DescribeParentInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRiskLevelsWithOptions(request *DescribeRiskLevelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRiskLevelsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRiskLevelsResponse
func (client *Client) DescribeRiskLevels(request *DescribeRiskLevelsRequest) (_result *DescribeRiskLevelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRiskLevelsResponse{}
	_body, _err := client.DescribeRiskLevelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRulesWithOptions(request *DescribeRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRulesResponse
func (client *Client) DescribeRules(request *DescribeRulesRequest) (_result *DescribeRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRulesResponse{}
	_body, _err := client.DescribeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeTablesWithOptions(request *DescribeTablesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTablesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeTablesResponse
func (client *Client) DescribeTables(request *DescribeTablesRequest) (_result *DescribeTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTablesResponse{}
	_body, _err := client.DescribeTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeTemplateAllRulesWithOptions(request *DescribeTemplateAllRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplateAllRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeTemplateAllRulesResponse
func (client *Client) DescribeTemplateAllRules(request *DescribeTemplateAllRulesRequest) (_result *DescribeTemplateAllRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTemplateAllRulesResponse{}
	_body, _err := client.DescribeTemplateAllRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a user account.
//
// Description:
//
// Retrieves information about the current account, such as your usage of Data Security Center (DSC).
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed this limit, API calls are throttled. This may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserStatusResponse
func (client *Client) DescribeUserStatusWithOptions(request *DescribeUserStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a user account.
//
// Description:
//
// Retrieves information about the current account, such as your usage of Data Security Center (DSC).
//
// ## QPS limit
//
// The queries per second (QPS) limit for this operation is 10 calls per second for each user. If you exceed this limit, API calls are throttled. This may affect your business. Call this operation at a reasonable frequency.
//
// @param request - DescribeUserStatusRequest
//
// @return DescribeUserStatusResponse
func (client *Client) DescribeUserStatus(request *DescribeUserStatusRequest) (_result *DescribeUserStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserStatusResponse{}
	_body, _err := client.DescribeUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DisableUserConfig operation to disable a user configuration. After a configuration is disabled, you can call the CreateConfig operation and specify the same Code parameter to restore the general anomaly alert configuration.
//
// Description:
//
// This operation disables a user configuration based on the code of a configuration item in the general anomaly alert configuration module. This lets you promptly change the status of the user configuration.
//
// ## QPS limits
//
// This operation is limited to 10 queries per second (QPS) per user. Calls that exceed this limit are throttled. Throttling may impact your business. Plan your calls accordingly.
//
// @param request - DisableUserConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableUserConfigResponse
func (client *Client) DisableUserConfigWithOptions(request *DisableUserConfigRequest, runtime *dara.RuntimeOptions) (_result *DisableUserConfigResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DisableUserConfig operation to disable a user configuration. After a configuration is disabled, you can call the CreateConfig operation and specify the same Code parameter to restore the general anomaly alert configuration.
//
// Description:
//
// This operation disables a user configuration based on the code of a configuration item in the general anomaly alert configuration module. This lets you promptly change the status of the user configuration.
//
// ## QPS limits
//
// This operation is limited to 10 queries per second (QPS) per user. Calls that exceed this limit are throttled. Throttling may impact your business. Plan your calls accordingly.
//
// @param request - DisableUserConfigRequest
//
// @return DisableUserConfigResponse
func (client *Client) DisableUserConfig(request *DisableUserConfigRequest) (_result *DisableUserConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableUserConfigResponse{}
	_body, _err := client.DisableUserConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecDatamaskWithOptions(request *ExecDatamaskRequest, runtime *dara.RuntimeOptions) (_result *ExecDatamaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecDatamaskResponse
func (client *Client) ExecDatamask(request *ExecDatamaskRequest) (_result *ExecDatamaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecDatamaskResponse{}
	_body, _err := client.ExecDatamaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ManualTriggerMaskingProcessWithOptions(request *ManualTriggerMaskingProcessRequest, runtime *dara.RuntimeOptions) (_result *ManualTriggerMaskingProcessResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ManualTriggerMaskingProcessResponse
func (client *Client) ManualTriggerMaskingProcess(request *ManualTriggerMaskingProcessRequest) (_result *ManualTriggerMaskingProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ManualTriggerMaskingProcessResponse{}
	_body, _err := client.ManualTriggerMaskingProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) MaskOssImageWithOptions(request *MaskOssImageRequest, runtime *dara.RuntimeOptions) (_result *MaskOssImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return MaskOssImageResponse
func (client *Client) MaskOssImage(request *MaskOssImageRequest) (_result *MaskOssImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MaskOssImageResponse{}
	_body, _err := client.MaskOssImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call ModifyDataLimit to modify the configuration items of a connection authorization in Data Security Center (DSC).
//
// @param request - ModifyDataLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataLimitResponse
func (client *Client) ModifyDataLimitWithOptions(request *ModifyDataLimitRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataLimitResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call ModifyDataLimit to modify the configuration items of a connection authorization in Data Security Center (DSC).
//
// @param request - ModifyDataLimitRequest
//
// @return ModifyDataLimitResponse
func (client *Client) ModifyDataLimit(request *ModifyDataLimitRequest) (_result *ModifyDataLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDataLimitResponse{}
	_body, _err := client.ModifyDataLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the rules that define threat levels for sensitive data. This includes the default threat level for unidentified data and the threat levels for data that is classified as sensitive.
//
// Description:
//
// This API modifies the rules that define threat levels for sensitive data to help with threat level planning.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, API calls are throttled. Throttling can impact your business. We recommend that you call this API at a reasonable rate.
//
// @param request - ModifyDefaultLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefaultLevelResponse
func (client *Client) ModifyDefaultLevelWithOptions(request *ModifyDefaultLevelRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefaultLevelResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the rules that define threat levels for sensitive data. This includes the default threat level for unidentified data and the threat levels for data that is classified as sensitive.
//
// Description:
//
// This API modifies the rules that define threat levels for sensitive data to help with threat level planning.
//
// ## QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you exceed this limit, API calls are throttled. Throttling can impact your business. We recommend that you call this API at a reasonable rate.
//
// @param request - ModifyDefaultLevelRequest
//
// @return ModifyDefaultLevelResponse
func (client *Client) ModifyDefaultLevel(request *ModifyDefaultLevelRequest) (_result *ModifyDefaultLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefaultLevelResponse{}
	_body, _err := client.ModifyDefaultLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyEventStatusWithOptions(request *ModifyEventStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyEventStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyEventStatusResponse
func (client *Client) ModifyEventStatus(request *ModifyEventStatusRequest) (_result *ModifyEventStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEventStatusResponse{}
	_body, _err := client.ModifyEventStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation enables anomalous activity detection for subtypes.
//
// @param request - ModifyEventTypeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEventTypeStatusResponse
func (client *Client) ModifyEventTypeStatusWithOptions(request *ModifyEventTypeStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyEventTypeStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation enables anomalous activity detection for subtypes.
//
// @param request - ModifyEventTypeStatusRequest
//
// @return ModifyEventTypeStatusResponse
func (client *Client) ModifyEventTypeStatus(request *ModifyEventTypeStatusRequest) (_result *ModifyEventTypeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEventTypeStatusResponse{}
	_body, _err := client.ModifyEventTypeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyReportTaskStatusWithOptions(request *ModifyReportTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyReportTaskStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyReportTaskStatusResponse
func (client *Client) ModifyReportTaskStatus(request *ModifyReportTaskStatusRequest) (_result *ModifyReportTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyReportTaskStatusResponse{}
	_body, _err := client.ModifyReportTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a custom sensitive data detection rule in Data Security Center (DSC).
//
// Description:
//
// You must specify the rule name, rule ID, and rule content.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for a single user. If the limit is exceeded, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - ModifyRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRuleResponse
func (client *Client) ModifyRuleWithOptions(request *ModifyRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You must specify the rule name, rule ID, and rule content.
//
// ## QPS limits
//
// The queries per second (QPS) limit for this operation is 10 calls per second for a single user. If the limit is exceeded, API calls are throttled. This may affect your business. Plan your calls accordingly.
//
// @param request - ModifyRuleRequest
//
// @return ModifyRuleResponse
func (client *Client) ModifyRule(request *ModifyRuleRequest) (_result *ModifyRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRuleResponse{}
	_body, _err := client.ModifyRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables sensitive data detection rules.
//
// @param request - ModifyRuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRuleStatusResponse
func (client *Client) ModifyRuleStatusWithOptions(request *ModifyRuleStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyRuleStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables sensitive data detection rules.
//
// @param request - ModifyRuleStatusRequest
//
// @return ModifyRuleStatusResponse
func (client *Client) ModifyRuleStatus(request *ModifyRuleStatusRequest) (_result *ModifyRuleStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRuleStatusResponse{}
	_body, _err := client.ModifyRuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the RestoreOssImage operation to restore desensitized images.
//
// Description:
//
// You can use RestoreOssImage to retrieve the original images that were processed by the MaskOssImage operation if the IsAlwaysUpload parameter is set to `true`.
//
// For example, the image `aliyun_dsc_desensitization/exampledir/test.png` in a bucket is restored and saved as `aliyun_dsc_original/exampledir/test.png`.
//
// @param request - RestoreOssImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreOssImageResponse
func (client *Client) RestoreOssImageWithOptions(request *RestoreOssImageRequest, runtime *dara.RuntimeOptions) (_result *RestoreOssImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the RestoreOssImage operation to restore desensitized images.
//
// Description:
//
// You can use RestoreOssImage to retrieve the original images that were processed by the MaskOssImage operation if the IsAlwaysUpload parameter is set to `true`.
//
// For example, the image `aliyun_dsc_desensitization/exampledir/test.png` in a bucket is restored and saved as `aliyun_dsc_original/exampledir/test.png`.
//
// @param request - RestoreOssImageRequest
//
// @return RestoreOssImageResponse
func (client *Client) RestoreOssImage(request *RestoreOssImageRequest) (_result *RestoreOssImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestoreOssImageResponse{}
	_body, _err := client.RestoreOssImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ScanOssObjectV1WithOptions(tmpReq *ScanOssObjectV1Request, runtime *dara.RuntimeOptions) (_result *ScanOssObjectV1Response, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ScanOssObjectV1Request
//
// @return ScanOssObjectV1Response
func (client *Client) ScanOssObjectV1(request *ScanOssObjectV1Request) (_result *ScanOssObjectV1Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ScanOssObjectV1Response{}
	_body, _err := client.ScanOssObjectV1WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopMaskingProcessWithOptions(request *StopMaskingProcessRequest, runtime *dara.RuntimeOptions) (_result *StopMaskingProcessResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StopMaskingProcessResponse
func (client *Client) StopMaskingProcess(request *StopMaskingProcessRequest) (_result *StopMaskingProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopMaskingProcessResponse{}
	_body, _err := client.StopMaskingProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
