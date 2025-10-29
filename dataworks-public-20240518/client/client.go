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
		"ap-northeast-1":        dara.String("dataworks.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":            dara.String("dataworks.ap-south-1.aliyuncs.com"),
		"ap-southeast-1":        dara.String("dataworks.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        dara.String("dataworks.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        dara.String("dataworks.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        dara.String("dataworks.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":            dara.String("dataworks.cn-beijing.aliyuncs.com"),
		"cn-chengdu":            dara.String("dataworks.cn-chengdu.aliyuncs.com"),
		"cn-hangzhou":           dara.String("dataworks.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           dara.String("dataworks.cn-hongkong.aliyuncs.com"),
		"cn-huhehaote":          dara.String("dataworks.aliyuncs.com"),
		"cn-qingdao":            dara.String("dataworks.aliyuncs.com"),
		"cn-shanghai":           dara.String("dataworks.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           dara.String("dataworks.cn-shenzhen.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("dataworks.aliyuncs.com"),
		"eu-central-1":          dara.String("dataworks.eu-central-1.aliyuncs.com"),
		"eu-west-1":             dara.String("dataworks.eu-west-1.aliyuncs.com"),
		"me-east-1":             dara.String("dataworks.me-east-1.aliyuncs.com"),
		"us-east-1":             dara.String("dataworks.us-east-1.aliyuncs.com"),
		"us-west-1":             dara.String("dataworks.us-west-1.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("dataworks.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("dataworks.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("dataworks.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("dataworks.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dataworks-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}) (_result map[string]interface{}, _err error) {
	request_ := dara.NewRequest()
	boundary := dara.GetBoundary()
	request_.Protocol = dara.String("HTTPS")
	request_.Method = dara.String("POST")
	request_.Pathname = dara.String("/")
	request_.Headers = map[string]*string{
		"host":       dara.String(dara.ToString(form["host"])),
		"date":       openapiutil.GetDateUTCString(),
		"user-agent": openapiutil.GetUserAgent(dara.String("")),
	}
	request_.Headers["content-type"] = dara.String("multipart/form-data; boundary=" + boundary)
	request_.Body = dara.ToFileForm(form, boundary)
	response_, _err := dara.DoRequest(request_, nil)
	if _err != nil {
		return nil, _err
	}

	_result, _err = _postOSSObject_opResponse(response_)
	if _err != nil {
		return nil, _err
	}

	return _result, nil
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
// Terminates the process for deploying or undeploying an entity. The process is not deleted and can still be queried by calling query operations.
//
// @param request - AbolishPipelineRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbolishPipelineRunResponse
func (client *Client) AbolishPipelineRunWithOptions(request *AbolishPipelineRunRequest, runtime *dara.RuntimeOptions) (_result *AbolishPipelineRunResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbolishPipelineRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbolishPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates the process for deploying or undeploying an entity. The process is not deleted and can still be queried by calling query operations.
//
// @param request - AbolishPipelineRunRequest
//
// @return AbolishPipelineRunResponse
func (client *Client) AbolishPipelineRun(request *AbolishPipelineRunRequest) (_result *AbolishPipelineRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AbolishPipelineRunResponse{}
	_body, _err := client.AbolishPipelineRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从集合中移除实体对象
//
// @param request - AddEntityIntoMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEntityIntoMetaCollectionResponse
func (client *Client) AddEntityIntoMetaCollectionWithOptions(request *AddEntityIntoMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *AddEntityIntoMetaCollectionResponse, _err error) {
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

	if !dara.IsNil(request.MetaCollectionId) {
		query["MetaCollectionId"] = request.MetaCollectionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEntityIntoMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEntityIntoMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从集合中移除实体对象
//
// @param request - AddEntityIntoMetaCollectionRequest
//
// @return AddEntityIntoMetaCollectionResponse
func (client *Client) AddEntityIntoMetaCollection(request *AddEntityIntoMetaCollectionRequest) (_result *AddEntityIntoMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddEntityIntoMetaCollectionResponse{}
	_body, _err := client.AddEntityIntoMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a resource group with a workspace.
//
// Description:
//
// 1.  You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// 2.  Your account must be assigned one of the following roles of the desired workspace:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
//
// @param request - AssociateProjectToResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateProjectToResourceGroupResponse
func (client *Client) AssociateProjectToResourceGroupWithOptions(request *AssociateProjectToResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *AssociateProjectToResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateProjectToResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateProjectToResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a resource group with a workspace.
//
// Description:
//
// 1.  You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// 2.  Your account must be assigned one of the following roles of the desired workspace:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
//
// @param request - AssociateProjectToResourceGroupRequest
//
// @return AssociateProjectToResourceGroupResponse
func (client *Client) AssociateProjectToResourceGroup(request *AssociateProjectToResourceGroupRequest) (_result *AssociateProjectToResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateProjectToResourceGroupResponse{}
	_body, _err := client.AssociateProjectToResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates monitoring rules with a data quality monitoring task.
//
// @param tmpReq - AttachDataQualityRulesToEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachDataQualityRulesToEvaluationTaskResponse
func (client *Client) AttachDataQualityRulesToEvaluationTaskWithOptions(tmpReq *AttachDataQualityRulesToEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *AttachDataQualityRulesToEvaluationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachDataQualityRulesToEvaluationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataQualityRuleIds) {
		request.DataQualityRuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataQualityRuleIds, dara.String("DataQualityRuleIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityEvaluationTaskId) {
		body["DataQualityEvaluationTaskId"] = request.DataQualityEvaluationTaskId
	}

	if !dara.IsNil(request.DataQualityRuleIdsShrink) {
		body["DataQualityRuleIds"] = request.DataQualityRuleIdsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachDataQualityRulesToEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachDataQualityRulesToEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates monitoring rules with a data quality monitoring task.
//
// @param request - AttachDataQualityRulesToEvaluationTaskRequest
//
// @return AttachDataQualityRulesToEvaluationTaskResponse
func (client *Client) AttachDataQualityRulesToEvaluationTask(request *AttachDataQualityRulesToEvaluationTaskRequest) (_result *AttachDataQualityRulesToEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachDataQualityRulesToEvaluationTaskResponse{}
	_body, _err := client.AttachDataQualityRulesToEvaluationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs incremental updates on multiple tasks at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - BatchUpdateTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateTasksResponse
func (client *Client) BatchUpdateTasksWithOptions(tmpReq *BatchUpdateTasksRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tasks) {
		request.TasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tasks, dara.String("Tasks"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.TasksShrink) {
		body["Tasks"] = request.TasksShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs incremental updates on multiple tasks at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - BatchUpdateTasksRequest
//
// @return BatchUpdateTasksResponse
func (client *Client) BatchUpdateTasks(request *BatchUpdateTasksRequest) (_result *BatchUpdateTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUpdateTasksResponse{}
	_body, _err := client.BatchUpdateTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clones an existing data source.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
//
// @param request - CloneDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneDataSourceResponse
func (client *Client) CloneDataSourceWithOptions(request *CloneDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CloneDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloneDataSourceName) {
		query["CloneDataSourceName"] = request.CloneDataSourceName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneDataSource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clones an existing data source.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
//
// @param request - CloneDataSourceRequest
//
// @return CloneDataSourceResponse
func (client *Client) CloneDataSource(request *CloneDataSourceRequest) (_result *CloneDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloneDataSourceResponse{}
	_body, _err := client.CloneDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom monitoring alert rule.
//
// @param tmpReq - CreateAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAlertRuleResponse
func (client *Client) CreateAlertRuleWithOptions(tmpReq *CreateAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAlertRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerCondition) {
		request.TriggerConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerCondition, dara.String("TriggerCondition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.TriggerConditionShrink) {
		query["TriggerCondition"] = request.TriggerConditionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom monitoring alert rule.
//
// @param request - CreateAlertRuleRequest
//
// @return CreateAlertRuleResponse
func (client *Client) CreateAlertRule(request *CreateAlertRuleRequest) (_result *CreateAlertRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAlertRuleResponse{}
	_body, _err := client.CreateAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBusinessResponse
func (client *Client) CreateBusinessWithOptions(request *CreateBusinessRequest, runtime *dara.RuntimeOptions) (_result *CreateBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.UseType) {
		body["UseType"] = request.UseType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateBusinessRequest
//
// @return CreateBusinessResponse
func (client *Client) CreateBusiness(request *CreateBusinessRequest) (_result *CreateBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBusinessResponse{}
	_body, _err := client.CreateBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates components.
//
// @param request - CreateComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateComponentResponse
func (client *Client) CreateComponentWithOptions(request *CreateComponentRequest, runtime *dara.RuntimeOptions) (_result *CreateComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateComponent"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateComponentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates components.
//
// @param request - CreateComponentRequest
//
// @return CreateComponentResponse
func (client *Client) CreateComponent(request *CreateComponentRequest) (_result *CreateComponentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateComponentResponse{}
	_body, _err := client.CreateComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an alert rule for a synchronization task.
//
// @param tmpReq - CreateDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDIAlarmRuleResponse
func (client *Client) CreateDIAlarmRuleWithOptions(tmpReq *CreateDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDIAlarmRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDIAlarmRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotificationSettings) {
		request.NotificationSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationSettings, dara.String("NotificationSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerConditions) {
		request.TriggerConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConditions, dara.String("TriggerConditions"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDIAlarmRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert rule for a synchronization task.
//
// @param request - CreateDIAlarmRuleRequest
//
// @return CreateDIAlarmRuleResponse
func (client *Client) CreateDIAlarmRule(request *CreateDIAlarmRuleRequest) (_result *CreateDIAlarmRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDIAlarmRuleResponse{}
	_body, _err := client.CreateDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new-version synchronization task.
//
// Description:
//
//	  This API operation is available for all DataWorks editions.
//
//		- You can call this API operation to create a synchronization task. When you call this API operation, you must configure parameters such as SourceDataSourceSettings, DestinationDataSourceSettings, MigrationType, TransformationRules, TableMappings, and JobSettings. The SourceDataSourceSettings parameter defines the settings related to the source. The DestinationDataSourceSettings parameter defines the settings related to the destination. The MigrationType parameter defines the synchronization task type. The TransformationRules parameter defines the transformation rules for objects involved in the synchronization task. The TableMappings parameter defines the mappings between rules used to select synchronization objects in the source and transformation rules applied to the selected synchronization objects. The JobSettings parameter defines the settings for the dimension of the synchronization task, including policies for data type mappings between source fields and destination fields and settings for periodic scheduling.
//
// @param tmpReq - CreateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDIJobResponse
func (client *Client) CreateDIJobWithOptions(tmpReq *CreateDIJobRequest, runtime *dara.RuntimeOptions) (_result *CreateDIJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestinationDataSourceSettings) {
		request.DestinationDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationDataSourceSettings, dara.String("DestinationDataSourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.JobSettings) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, dara.String("JobSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSettings) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, dara.String("ResourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceDataSourceSettings) {
		request.SourceDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceDataSourceSettings, dara.String("SourceDataSourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TableMappings) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, dara.String("TableMappings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TransformationRules) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, dara.String("TransformationRules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationDataSourceType) {
		query["DestinationDataSourceType"] = request.DestinationDataSourceType
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.MigrationType) {
		query["MigrationType"] = request.MigrationType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SourceDataSourceType) {
		query["SourceDataSourceType"] = request.SourceDataSourceType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationDataSourceSettingsShrink) {
		body["DestinationDataSourceSettings"] = request.DestinationDataSourceSettingsShrink
	}

	if !dara.IsNil(request.JobSettingsShrink) {
		body["JobSettings"] = request.JobSettingsShrink
	}

	if !dara.IsNil(request.ResourceSettingsShrink) {
		body["ResourceSettings"] = request.ResourceSettingsShrink
	}

	if !dara.IsNil(request.SourceDataSourceSettingsShrink) {
		body["SourceDataSourceSettings"] = request.SourceDataSourceSettingsShrink
	}

	if !dara.IsNil(request.TableMappingsShrink) {
		body["TableMappings"] = request.TableMappingsShrink
	}

	if !dara.IsNil(request.TransformationRulesShrink) {
		body["TransformationRules"] = request.TransformationRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new-version synchronization task.
//
// Description:
//
//	  This API operation is available for all DataWorks editions.
//
//		- You can call this API operation to create a synchronization task. When you call this API operation, you must configure parameters such as SourceDataSourceSettings, DestinationDataSourceSettings, MigrationType, TransformationRules, TableMappings, and JobSettings. The SourceDataSourceSettings parameter defines the settings related to the source. The DestinationDataSourceSettings parameter defines the settings related to the destination. The MigrationType parameter defines the synchronization task type. The TransformationRules parameter defines the transformation rules for objects involved in the synchronization task. The TableMappings parameter defines the mappings between rules used to select synchronization objects in the source and transformation rules applied to the selected synchronization objects. The JobSettings parameter defines the settings for the dimension of the synchronization task, including policies for data type mappings between source fields and destination fields and settings for periodic scheduling.
//
// @param request - CreateDIJobRequest
//
// @return CreateDIJobResponse
func (client *Client) CreateDIJob(request *CreateDIJobRequest) (_result *CreateDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDIJobResponse{}
	_body, _err := client.CreateDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a tag.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - CreateDataAssetTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataAssetTagResponse
func (client *Client) CreateDataAssetTagWithOptions(tmpReq *CreateDataAssetTagRequest, runtime *dara.RuntimeOptions) (_result *CreateDataAssetTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataAssetTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Managers) {
		request.ManagersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Managers, dara.String("Managers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Values) {
		request.ValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Values, dara.String("Values"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.ManagersShrink) {
		query["Managers"] = request.ManagersShrink
	}

	if !dara.IsNil(request.ValueType) {
		query["ValueType"] = request.ValueType
	}

	if !dara.IsNil(request.ValuesShrink) {
		query["Values"] = request.ValuesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataAssetTag"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataAssetTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tag.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param request - CreateDataAssetTagRequest
//
// @return CreateDataAssetTagResponse
func (client *Client) CreateDataAssetTag(request *CreateDataAssetTagRequest) (_result *CreateDataAssetTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataAssetTagResponse{}
	_body, _err := client.CreateDataAssetTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring alert rule in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - CreateDataQualityAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityAlertRuleResponse
func (client *Client) CreateDataQualityAlertRuleWithOptions(tmpReq *CreateDataQualityAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityAlertRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Condition) {
		body["Condition"] = request.Condition
	}

	if !dara.IsNil(request.NotificationShrink) {
		body["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring alert rule in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateDataQualityAlertRuleRequest
//
// @return CreateDataQualityAlertRuleResponse
func (client *Client) CreateDataQualityAlertRule(request *CreateDataQualityAlertRuleRequest) (_result *CreateDataQualityAlertRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataQualityAlertRuleResponse{}
	_body, _err := client.CreateDataQualityAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a monitor in DataWorks Data Quality.
//
// Description:
//
// This API operation is supported in all DataWorks editions.
//
// @param tmpReq - CreateDataQualityEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityEvaluationTaskResponse
func (client *Client) CreateDataQualityEvaluationTaskWithOptions(tmpReq *CreateDataQualityEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityEvaluationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityEvaluationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataQualityRules) {
		request.DataQualityRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataQualityRules, dara.String("DataQualityRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hooks) {
		request.HooksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hooks, dara.String("Hooks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notifications) {
		request.NotificationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notifications, dara.String("Notifications"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityRulesShrink) {
		body["DataQualityRules"] = request.DataQualityRulesShrink
	}

	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HooksShrink) {
		body["Hooks"] = request.HooksShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationsShrink) {
		body["Notifications"] = request.NotificationsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeConf) {
		body["RuntimeConf"] = request.RuntimeConf
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a monitor in DataWorks Data Quality.
//
// Description:
//
// This API operation is supported in all DataWorks editions.
//
// @param request - CreateDataQualityEvaluationTaskRequest
//
// @return CreateDataQualityEvaluationTaskResponse
func (client *Client) CreateDataQualityEvaluationTask(request *CreateDataQualityEvaluationTaskRequest) (_result *CreateDataQualityEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataQualityEvaluationTaskResponse{}
	_body, _err := client.CreateDataQualityEvaluationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a monitor instance.
//
// @param tmpReq - CreateDataQualityEvaluationTaskInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityEvaluationTaskInstanceResponse
func (client *Client) CreateDataQualityEvaluationTaskInstanceWithOptions(tmpReq *CreateDataQualityEvaluationTaskInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityEvaluationTaskInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityEvaluationTaskInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityEvaluationTaskId) {
		body["DataQualityEvaluationTaskId"] = request.DataQualityEvaluationTaskId
	}

	if !dara.IsNil(request.Parameters) {
		body["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityEvaluationTaskInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityEvaluationTaskInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a monitor instance.
//
// @param request - CreateDataQualityEvaluationTaskInstanceRequest
//
// @return CreateDataQualityEvaluationTaskInstanceResponse
func (client *Client) CreateDataQualityEvaluationTaskInstance(request *CreateDataQualityEvaluationTaskInstanceRequest) (_result *CreateDataQualityEvaluationTaskInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataQualityEvaluationTaskInstanceResponse{}
	_body, _err := client.CreateDataQualityEvaluationTaskInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring rule.
//
// @param tmpReq - CreateDataQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityRuleResponse
func (client *Client) CreateDataQualityRuleWithOptions(tmpReq *CreateDataQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckingConfig) {
		request.CheckingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckingConfig, dara.String("CheckingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ErrorHandlers) {
		request.ErrorHandlersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ErrorHandlers, dara.String("ErrorHandlers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SamplingConfig) {
		request.SamplingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SamplingConfig, dara.String("SamplingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckingConfigShrink) {
		body["CheckingConfig"] = request.CheckingConfigShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ErrorHandlersShrink) {
		body["ErrorHandlers"] = request.ErrorHandlersShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SamplingConfigShrink) {
		body["SamplingConfig"] = request.SamplingConfigShrink
	}

	if !dara.IsNil(request.Severity) {
		body["Severity"] = request.Severity
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring rule.
//
// @param request - CreateDataQualityRuleRequest
//
// @return CreateDataQualityRuleResponse
func (client *Client) CreateDataQualityRule(request *CreateDataQualityRuleRequest) (_result *CreateDataQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataQualityRuleResponse{}
	_body, _err := client.CreateDataQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring rule template.
//
// @param tmpReq - CreateDataQualityRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityRuleTemplateResponse
func (client *Client) CreateDataQualityRuleTemplateWithOptions(tmpReq *CreateDataQualityRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityRuleTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityRuleTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckingConfig) {
		request.CheckingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckingConfig, dara.String("CheckingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SamplingConfig) {
		request.SamplingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SamplingConfig, dara.String("SamplingConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckingConfigShrink) {
		body["CheckingConfig"] = request.CheckingConfigShrink
	}

	if !dara.IsNil(request.DirectoryPath) {
		body["DirectoryPath"] = request.DirectoryPath
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SamplingConfigShrink) {
		body["SamplingConfig"] = request.SamplingConfigShrink
	}

	if !dara.IsNil(request.VisibleScope) {
		body["VisibleScope"] = request.VisibleScope
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityRuleTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityRuleTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring rule template.
//
// @param request - CreateDataQualityRuleTemplateRequest
//
// @return CreateDataQualityRuleTemplateResponse
func (client *Client) CreateDataQualityRuleTemplate(request *CreateDataQualityRuleTemplateRequest) (_result *CreateDataQualityRuleTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataQualityRuleTemplateResponse{}
	_body, _err := client.CreateDataQualityRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data quality monitor.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - CreateDataQualityScanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityScanResponse
func (client *Client) CreateDataQualityScanWithOptions(tmpReq *CreateDataQualityScanRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityScanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityScanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ComputeResource) {
		request.ComputeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComputeResource, dara.String("ComputeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hooks) {
		request.HooksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hooks, dara.String("Hooks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ComputeResourceShrink) {
		body["ComputeResource"] = request.ComputeResourceShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HooksShrink) {
		body["Hooks"] = request.HooksShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityScan"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality monitor.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateDataQualityScanRequest
//
// @return CreateDataQualityScanResponse
func (client *Client) CreateDataQualityScan(request *CreateDataQualityScanRequest) (_result *CreateDataQualityScanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataQualityScanResponse{}
	_body, _err := client.CreateDataQualityScanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers a data quality monitoring task and returns the run instance ID.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - CreateDataQualityScanRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityScanRunResponse
func (client *Client) CreateDataQualityScanRunWithOptions(tmpReq *CreateDataQualityScanRunRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityScanRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataQualityScanRunShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityScanId) {
		body["DataQualityScanId"] = request.DataQualityScanId
	}

	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityScanRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityScanRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a data quality monitoring task and returns the run instance ID.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateDataQualityScanRunRequest
//
// @return CreateDataQualityScanRunResponse
func (client *Client) CreateDataQualityScanRun(request *CreateDataQualityScanRunRequest) (_result *CreateDataQualityScanRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataQualityScanRunResponse{}
	_body, _err := client.CreateDataQualityScanRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data quality template.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateDataQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataQualityTemplateResponse
func (client *Client) CreateDataQualityTemplateWithOptions(request *CreateDataQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateDataQualityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataQualityTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataQualityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality template.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateDataQualityTemplateRequest
//
// @return CreateDataQualityTemplateResponse
func (client *Client) CreateDataQualityTemplate(request *CreateDataQualityTemplateRequest) (_result *CreateDataQualityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataQualityTemplateResponse{}
	_body, _err := client.CreateDataQualityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a data source to the development environment or production environment of a workspace.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
//
// @param request - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(request *CreateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionProperties) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !dara.IsNil(request.ConnectionPropertiesMode) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSource"),
		Version:     dara.String("2024-05-18"),
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
// Adds a data source to the development environment or production environment of a workspace.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
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
// Creates a rule for sharing a data source to other workspaces or RAM users.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  If you want to share a data source from Workspace A to Workspace B, you must have the permissions to share the data source in both workspaces. You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Tenant Administrator, Workspace Administrator, and Workspace Owner
//
// @param request - CreateDataSourceSharedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceSharedRuleResponse
func (client *Client) CreateDataSourceSharedRuleWithOptions(request *CreateDataSourceSharedRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceSharedRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.SharedUser) {
		query["SharedUser"] = request.SharedUser
	}

	if !dara.IsNil(request.TargetProjectId) {
		query["TargetProjectId"] = request.TargetProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSourceSharedRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSourceSharedRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a rule for sharing a data source to other workspaces or RAM users.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  If you want to share a data source from Workspace A to Workspace B, you must have the permissions to share the data source in both workspaces. You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Tenant Administrator, Workspace Administrator, and Workspace Owner
//
// @param request - CreateDataSourceSharedRuleRequest
//
// @return CreateDataSourceSharedRuleResponse
func (client *Client) CreateDataSourceSharedRule(request *CreateDataSourceSharedRuleRequest) (_result *CreateDataSourceSharedRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataSourceSharedRuleResponse{}
	_body, _err := client.CreateDataSourceSharedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileResponse
func (client *Client) CreateFileWithOptions(request *CreateFileRequest, runtime *dara.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedSettings) {
		body["AdvancedSettings"] = request.AdvancedSettings
	}

	if !dara.IsNil(request.ApplyScheduleImmediately) {
		body["ApplyScheduleImmediately"] = request.ApplyScheduleImmediately
	}

	if !dara.IsNil(request.AutoParsing) {
		body["AutoParsing"] = request.AutoParsing
	}

	if !dara.IsNil(request.AutoRerunIntervalMillis) {
		body["AutoRerunIntervalMillis"] = request.AutoRerunIntervalMillis
	}

	if !dara.IsNil(request.AutoRerunTimes) {
		body["AutoRerunTimes"] = request.AutoRerunTimes
	}

	if !dara.IsNil(request.ConnectionName) {
		body["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.CreateFolderIfNotExists) {
		body["CreateFolderIfNotExists"] = request.CreateFolderIfNotExists
	}

	if !dara.IsNil(request.CronExpress) {
		body["CronExpress"] = request.CronExpress
	}

	if !dara.IsNil(request.CycleType) {
		body["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.DependentNodeIdList) {
		body["DependentNodeIdList"] = request.DependentNodeIdList
	}

	if !dara.IsNil(request.DependentType) {
		body["DependentType"] = request.DependentType
	}

	if !dara.IsNil(request.EndEffectDate) {
		body["EndEffectDate"] = request.EndEffectDate
	}

	if !dara.IsNil(request.FileDescription) {
		body["FileDescription"] = request.FileDescription
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.IgnoreParentSkipRunningProperty) {
		body["IgnoreParentSkipRunningProperty"] = request.IgnoreParentSkipRunningProperty
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InputList) {
		body["InputList"] = request.InputList
	}

	if !dara.IsNil(request.InputParameters) {
		body["InputParameters"] = request.InputParameters
	}

	if !dara.IsNil(request.OutputParameters) {
		body["OutputParameters"] = request.OutputParameters
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParaValue) {
		body["ParaValue"] = request.ParaValue
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.RerunMode) {
		body["RerunMode"] = request.RerunMode
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceGroupIdentifier) {
		body["ResourceGroupIdentifier"] = request.ResourceGroupIdentifier
	}

	if !dara.IsNil(request.SchedulerType) {
		body["SchedulerType"] = request.SchedulerType
	}

	if !dara.IsNil(request.StartEffectDate) {
		body["StartEffectDate"] = request.StartEffectDate
	}

	if !dara.IsNil(request.StartImmediately) {
		body["StartImmediately"] = request.StartImmediately
	}

	if !dara.IsNil(request.Stop) {
		body["Stop"] = request.Stop
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFileRequest
//
// @return CreateFileResponse
func (client *Client) CreateFile(request *CreateFileRequest) (_result *CreateFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFileResponse{}
	_body, _err := client.CreateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithOptions(request *CreateFolderRequest, runtime *dara.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderPath) {
		body["FolderPath"] = request.FolderPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFolder"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFolderRequest
//
// @return CreateFolderResponse
func (client *Client) CreateFolder(request *CreateFolderRequest) (_result *CreateFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFolderResponse{}
	_body, _err := client.CreateFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user-defined function (UDF) in DataStudio. The information about the UDF is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple UDFs at a time. If you specify multiple UDFs by using FlowSpec, the system creates only the first specified UDF.
//
// @param request - CreateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionResponse
func (client *Client) CreateFunctionWithOptions(request *CreateFunctionRequest, runtime *dara.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user-defined function (UDF) in DataStudio. The information about the UDF is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple UDFs at a time. If you specify multiple UDFs by using FlowSpec, the system creates only the first specified UDF.
//
// @param request - CreateFunctionRequest
//
// @return CreateFunctionResponse
func (client *Client) CreateFunction(request *CreateFunctionRequest) (_result *CreateFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CreateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a lineage between a source entity and a destination entity. Either the source or destination entity must be a custom entity.
//
// @param tmpReq - CreateLineageRelationshipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLineageRelationshipResponse
func (client *Client) CreateLineageRelationshipWithOptions(tmpReq *CreateLineageRelationshipRequest, runtime *dara.RuntimeOptions) (_result *CreateLineageRelationshipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLineageRelationshipShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DstEntity) {
		request.DstEntityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DstEntity, dara.String("DstEntity"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SrcEntity) {
		request.SrcEntityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SrcEntity, dara.String("SrcEntity"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Task) {
		request.TaskShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Task, dara.String("Task"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DstEntityShrink) {
		query["DstEntity"] = request.DstEntityShrink
	}

	if !dara.IsNil(request.SrcEntityShrink) {
		query["SrcEntity"] = request.SrcEntityShrink
	}

	if !dara.IsNil(request.TaskShrink) {
		query["Task"] = request.TaskShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLineageRelationship"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLineageRelationshipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lineage between a source entity and a destination entity. Either the source or destination entity must be a custom entity.
//
// @param request - CreateLineageRelationshipRequest
//
// @return CreateLineageRelationshipResponse
func (client *Client) CreateLineageRelationship(request *CreateLineageRelationshipRequest) (_result *CreateLineageRelationshipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLineageRelationshipResponse{}
	_body, _err := client.CreateLineageRelationshipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a collection in Data Map. Collections include categories, subcategories, data albums, and categories that are created in the data albums.
//
// @param request - CreateMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetaCollectionResponse
func (client *Client) CreateMetaCollectionWithOptions(request *CreateMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *CreateMetaCollectionResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a collection in Data Map. Collections include categories, subcategories, data albums, and categories that are created in the data albums.
//
// @param request - CreateMetaCollectionRequest
//
// @return CreateMetaCollectionResponse
func (client *Client) CreateMetaCollection(request *CreateMetaCollectionRequest) (_result *CreateMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMetaCollectionResponse{}
	_body, _err := client.CreateMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a network and associates the network with a general resource group.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkResponse
func (client *Client) CreateNetworkWithOptions(request *CreateNetworkRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchId) {
		body["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetwork"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network and associates the network with a general resource group.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateNetworkRequest
//
// @return CreateNetworkResponse
func (client *Client) CreateNetwork(request *CreateNetworkRequest) (_result *CreateNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkResponse{}
	_body, _err := client.CreateNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a node in DataStudio. The information about the node is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple nodes at a time. If you specify multiple nodes by using FlowSpec, the system creates only the first specified node.
//
// @param request - CreateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeResponse
func (client *Client) CreateNodeWithOptions(request *CreateNodeRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContainerId) {
		body["ContainerId"] = request.ContainerId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a node in DataStudio. The information about the node is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple nodes at a time. If you specify multiple nodes by using FlowSpec, the system creates only the first specified node.
//
// @param request - CreateNodeRequest
//
// @return CreateNodeResponse
func (client *Client) CreateNode(request *CreateNodeRequest) (_result *CreateNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNodeResponse{}
	_body, _err := client.CreateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a process for deploying or undeploying an entity in Data Studio.
//
// Description:
//
// >  You cannot use this API operation to create a process for multiple entities at a time. If you specify multiple entities in a request, the system creates a process only for the first entity.
//
// @param tmpReq - CreatePipelineRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineRunResponse
func (client *Client) CreatePipelineRunWithOptions(tmpReq *CreatePipelineRunRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePipelineRunShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ObjectIds) {
		request.ObjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectIds, dara.String("ObjectIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ObjectIdsShrink) {
		body["ObjectIds"] = request.ObjectIdsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipelineRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a process for deploying or undeploying an entity in Data Studio.
//
// Description:
//
// >  You cannot use this API operation to create a process for multiple entities at a time. If you specify multiple entities in a request, the system creates a process only for the first entity.
//
// @param request - CreatePipelineRunRequest
//
// @return CreatePipelineRunResponse
func (client *Client) CreatePipelineRun(request *CreatePipelineRunRequest) (_result *CreatePipelineRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePipelineRunResponse{}
	_body, _err := client.CreatePipelineRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// @param tmpReq - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(tmpReq *CreateProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AliyunResourceTags) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, dara.String("AliyunResourceTags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunResourceGroupId) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !dara.IsNil(request.AliyunResourceTagsShrink) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DevEnvironmentEnabled) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !dara.IsNil(request.DevRoleDisabled) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PaiTaskEnabled) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a workspace member and assigns a workspace-level role to the member.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - CreateProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectMemberResponse
func (client *Client) CreateProjectMemberWithOptions(tmpReq *CreateProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("RoleCodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProjectMember"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a workspace member and assigns a workspace-level role to the member.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateProjectMemberRequest
//
// @return CreateProjectMemberResponse
func (client *Client) CreateProjectMember(request *CreateProjectMemberRequest) (_result *CreateProjectMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProjectMemberResponse{}
	_body, _err := client.CreateProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// >  You cannot use this API operation to create multiple file resources at a time. If you specify multiple file resources by using FlowSpec, the system creates only the first specified resource.
//
// Description:
//
// # Private
//
// @param request - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(request *CreateResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceFile) {
		body["ResourceFile"] = request.ResourceFile
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// >  You cannot use this API operation to create multiple file resources at a time. If you specify multiple file resources by using FlowSpec, the system creates only the first specified resource.
//
// Description:
//
// # Private
//
// @param request - CreateResourceRequest
//
// @return CreateResourceResponse
func (client *Client) CreateResource(request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceAdvance(request *CreateResourceAdvanceRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("dataworks-public"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	createResourceReq := &CreateResourceRequest{}
	openapiutil.Convert(request, createResourceReq)
	if !dara.IsNil(request.ResourceFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.ResourceFileObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		createResourceReq.ResourceFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	createResourceResp, _err := client.CreateResourceWithOptions(createResourceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createResourceResp
	return _result, _err
}

// Summary:
//
// Creates a resource file in DataStudio. The following types are supported: JAR, Archive, File, and Python.
//
// @param request - CreateResourceFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceFileResponse
func (client *Client) CreateResourceFileWithOptions(request *CreateResourceFileRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.FileDescription) {
		body["FileDescription"] = request.FileDescription
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.OriginResourceName) {
		body["OriginResourceName"] = request.OriginResourceName
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegisterToCalcEngine) {
		body["RegisterToCalcEngine"] = request.RegisterToCalcEngine
	}

	if !dara.IsNil(request.ResourceFile) {
		body["ResourceFile"] = request.ResourceFile
	}

	if !dara.IsNil(request.StorageURL) {
		body["StorageURL"] = request.StorageURL
	}

	if !dara.IsNil(request.UploadMode) {
		body["UploadMode"] = request.UploadMode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource file in DataStudio. The following types are supported: JAR, Archive, File, and Python.
//
// @param request - CreateResourceFileRequest
//
// @return CreateResourceFileResponse
func (client *Client) CreateResourceFile(request *CreateResourceFileRequest) (_result *CreateResourceFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceFileResponse{}
	_body, _err := client.CreateResourceFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceFileAdvance(request *CreateResourceFileAdvanceRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceFileResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("dataworks-public"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	createResourceFileReq := &CreateResourceFileRequest{}
	openapiutil.Convert(request, createResourceFileReq)
	if !dara.IsNil(request.ResourceFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.ResourceFileObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		createResourceFileReq.ResourceFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	createResourceFileResp, _err := client.CreateResourceFileWithOptions(createResourceFileReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createResourceFileResp
	return _result, _err
}

// Summary:
//
// Creates a serverless resource group.
//
// Description:
//
// 1.  You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// 2.  **Before you call this API operation, you must make sure that you have a good command of the billing details and [pricing](https://help.aliyun.com/document_detail/2680173.html) of serverless resource groups.
//
// @param tmpReq - CreateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithOptions(tmpReq *CreateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateResourceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AliyunResourceTags) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, dara.String("AliyunResourceTags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunResourceGroupId) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !dara.IsNil(request.AliyunResourceTagsShrink) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !dara.IsNil(request.AutoRenewEnabled) {
		body["AutoRenewEnabled"] = request.AutoRenewEnabled
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PaymentDuration) {
		body["PaymentDuration"] = request.PaymentDuration
	}

	if !dara.IsNil(request.PaymentDurationUnit) {
		body["PaymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !dara.IsNil(request.PaymentType) {
		body["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchId) {
		body["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a serverless resource group.
//
// Description:
//
// 1.  You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// 2.  **Before you call this API operation, you must make sure that you have a good command of the billing details and [pricing](https://help.aliyun.com/document_detail/2680173.html) of serverless resource groups.
//
// @param request - CreateResourceGroupRequest
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (_result *CreateResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CreateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a route for a network.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteResponse
func (client *Client) CreateRouteWithOptions(request *CreateRouteRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidr) {
		body["DestinationCidr"] = request.DestinationCidr
	}

	if !dara.IsNil(request.NetworkId) {
		body["NetworkId"] = request.NetworkId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRoute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route for a network.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - CreateRouteRequest
//
// @return CreateRouteResponse
func (client *Client) CreateRoute(request *CreateRouteRequest) (_result *CreateRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRouteResponse{}
	_body, _err := client.CreateRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateUdfFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfFileResponse
func (client *Client) CreateUdfFileWithOptions(request *CreateUdfFileRequest, runtime *dara.RuntimeOptions) (_result *CreateUdfFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		body["ClassName"] = request.ClassName
	}

	if !dara.IsNil(request.CmdDescription) {
		body["CmdDescription"] = request.CmdDescription
	}

	if !dara.IsNil(request.CreateFolderIfNotExists) {
		body["CreateFolderIfNotExists"] = request.CreateFolderIfNotExists
	}

	if !dara.IsNil(request.Example) {
		body["Example"] = request.Example
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FunctionType) {
		body["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.ParameterDescription) {
		body["ParameterDescription"] = request.ParameterDescription
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.Resources) {
		body["Resources"] = request.Resources
	}

	if !dara.IsNil(request.ReturnValue) {
		body["ReturnValue"] = request.ReturnValue
	}

	if !dara.IsNil(request.UdfDescription) {
		body["UdfDescription"] = request.UdfDescription
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUdfFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUdfFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateUdfFileRequest
//
// @return CreateUdfFileResponse
func (client *Client) CreateUdfFile(request *CreateUdfFileRequest) (_result *CreateUdfFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUdfFileResponse{}
	_body, _err := client.CreateUdfFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow in a directory of DataStudio.
//
// Description:
//
// > You cannot use this API operation to create multiple workflows at a time. If you specify multiple workflows by using FlowSpec, the system creates only the first specified workflow. Other specified workflows and the nodes in the workflows are ignored. You can call the CreateNode operation to create a node.
//
// @param request - CreateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowDefinitionResponse
func (client *Client) CreateWorkflowDefinitionWithOptions(request *CreateWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkflowDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow in a directory of DataStudio.
//
// Description:
//
// > You cannot use this API operation to create multiple workflows at a time. If you specify multiple workflows by using FlowSpec, the system creates only the first specified workflow. Other specified workflows and the nodes in the workflows are ignored. You can call the CreateNode operation to create a node.
//
// @param request - CreateWorkflowDefinitionRequest
//
// @return CreateWorkflowDefinitionResponse
func (client *Client) CreateWorkflowDefinition(request *CreateWorkflowDefinitionRequest) (_result *CreateWorkflowDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWorkflowDefinitionResponse{}
	_body, _err := client.CreateWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow instance, such as a data backfill workflow instance, based on configurations.
//
// @param tmpReq - CreateWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowInstancesResponse
func (client *Client) CreateWorkflowInstancesWithOptions(tmpReq *CreateWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DefaultRunProperties) {
		request.DefaultRunPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultRunProperties, dara.String("DefaultRunProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Periods) {
		request.PeriodsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Periods, dara.String("Periods"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoStartEnabled) {
		body["AutoStartEnabled"] = request.AutoStartEnabled
	}

	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DefaultRunPropertiesShrink) {
		body["DefaultRunProperties"] = request.DefaultRunPropertiesShrink
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PeriodsShrink) {
		body["Periods"] = request.PeriodsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TagCreationPolicy) {
		body["TagCreationPolicy"] = request.TagCreationPolicy
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskParameters) {
		body["TaskParameters"] = request.TaskParameters
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkflowParameters) {
		body["WorkflowParameters"] = request.WorkflowParameters
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkflowInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow instance, such as a data backfill workflow instance, based on configurations.
//
// @param request - CreateWorkflowInstancesRequest
//
// @return CreateWorkflowInstancesResponse
func (client *Client) CreateWorkflowInstances(request *CreateWorkflowInstancesRequest) (_result *CreateWorkflowInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWorkflowInstancesResponse{}
	_body, _err := client.CreateWorkflowInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom monitoring alert rule.
//
// @param request - DeleteAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAlertRuleResponse
func (client *Client) DeleteAlertRuleWithOptions(request *DeleteAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAlertRuleResponse, _err error) {
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
		Action:      dara.String("DeleteAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom monitoring alert rule.
//
// @param request - DeleteAlertRuleRequest
//
// @return DeleteAlertRuleResponse
func (client *Client) DeleteAlertRule(request *DeleteAlertRuleRequest) (_result *DeleteAlertRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAlertRuleResponse{}
	_body, _err := client.DeleteAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBusinessResponse
func (client *Client) DeleteBusinessWithOptions(request *DeleteBusinessRequest, runtime *dara.RuntimeOptions) (_result *DeleteBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteBusinessRequest
//
// @return DeleteBusinessResponse
func (client *Client) DeleteBusiness(request *DeleteBusinessRequest) (_result *DeleteBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBusinessResponse{}
	_body, _err := client.DeleteBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a certificate file.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks: Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M.
//
// @param request - DeleteCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCertificateResponse
func (client *Client) DeleteCertificateWithOptions(request *DeleteCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCertificateResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCertificate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a certificate file.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks: Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M.
//
// @param request - DeleteCertificateRequest
//
// @return DeleteCertificateResponse
func (client *Client) DeleteCertificate(request *DeleteCertificateRequest) (_result *DeleteCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCertificateResponse{}
	_body, _err := client.DeleteCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a component.
//
// Description:
//
// >  A UDF that is deployed cannot be deleted. If you want to delete such a UDF, you must first undeploy the UDF.
//
// @param request - DeleteComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteComponentResponse
func (client *Client) DeleteComponentWithOptions(request *DeleteComponentRequest, runtime *dara.RuntimeOptions) (_result *DeleteComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentId) {
		body["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteComponent"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteComponentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a component.
//
// Description:
//
// >  A UDF that is deployed cannot be deleted. If you want to delete such a UDF, you must first undeploy the UDF.
//
// @param request - DeleteComponentRequest
//
// @return DeleteComponentResponse
func (client *Client) DeleteComponent(request *DeleteComponentRequest) (_result *DeleteComponentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteComponentResponse{}
	_body, _err := client.DeleteComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alert rule configured for a synchronization task.
//
// @param request - DeleteDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDIAlarmRuleResponse
func (client *Client) DeleteDIAlarmRuleWithOptions(request *DeleteDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDIAlarmRuleResponse, _err error) {
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
		Action:      dara.String("DeleteDIAlarmRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert rule configured for a synchronization task.
//
// @param request - DeleteDIAlarmRuleRequest
//
// @return DeleteDIAlarmRuleResponse
func (client *Client) DeleteDIAlarmRule(request *DeleteDIAlarmRuleRequest) (_result *DeleteDIAlarmRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDIAlarmRuleResponse{}
	_body, _err := client.DeleteDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a new-version synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDIJobResponse
func (client *Client) DeleteDIJobWithOptions(request *DeleteDIJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteDIJobResponse, _err error) {
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
		Action:      dara.String("DeleteDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a new-version synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteDIJobRequest
//
// @return DeleteDIJobResponse
func (client *Client) DeleteDIJob(request *DeleteDIJobRequest) (_result *DeleteDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDIJobResponse{}
	_body, _err := client.DeleteDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a tag.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - DeleteDataAssetTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataAssetTagResponse
func (client *Client) DeleteDataAssetTagWithOptions(tmpReq *DeleteDataAssetTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataAssetTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDataAssetTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Values) {
		request.ValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Values, dara.String("Values"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.ValuesShrink) {
		query["Values"] = request.ValuesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataAssetTag"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataAssetTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a tag.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param request - DeleteDataAssetTagRequest
//
// @return DeleteDataAssetTagResponse
func (client *Client) DeleteDataAssetTag(request *DeleteDataAssetTagRequest) (_result *DeleteDataAssetTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataAssetTagResponse{}
	_body, _err := client.DeleteDataAssetTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data quality alert rule by ID.
//
// Description:
//
// Subscribe to DataWorks Basic Edition or a higher version to use this API.
//
// @param request - DeleteDataQualityAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityAlertRuleResponse
func (client *Client) DeleteDataQualityAlertRuleWithOptions(request *DeleteDataQualityAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityAlertRuleResponse, _err error) {
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
		Action:      dara.String("DeleteDataQualityAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality alert rule by ID.
//
// Description:
//
// Subscribe to DataWorks Basic Edition or a higher version to use this API.
//
// @param request - DeleteDataQualityAlertRuleRequest
//
// @return DeleteDataQualityAlertRuleResponse
func (client *Client) DeleteDataQualityAlertRule(request *DeleteDataQualityAlertRuleRequest) (_result *DeleteDataQualityAlertRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataQualityAlertRuleResponse{}
	_body, _err := client.DeleteDataQualityAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitoring task.
//
// @param request - DeleteDataQualityEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityEvaluationTaskResponse
func (client *Client) DeleteDataQualityEvaluationTaskWithOptions(request *DeleteDataQualityEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityEvaluationTaskResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitoring task.
//
// @param request - DeleteDataQualityEvaluationTaskRequest
//
// @return DeleteDataQualityEvaluationTaskResponse
func (client *Client) DeleteDataQualityEvaluationTask(request *DeleteDataQualityEvaluationTaskRequest) (_result *DeleteDataQualityEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataQualityEvaluationTaskResponse{}
	_body, _err := client.DeleteDataQualityEvaluationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitoring rule.
//
// @param request - DeleteDataQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityRuleResponse
func (client *Client) DeleteDataQualityRuleWithOptions(request *DeleteDataQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityRuleResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitoring rule.
//
// @param request - DeleteDataQualityRuleRequest
//
// @return DeleteDataQualityRuleResponse
func (client *Client) DeleteDataQualityRule(request *DeleteDataQualityRuleRequest) (_result *DeleteDataQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataQualityRuleResponse{}
	_body, _err := client.DeleteDataQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitoring rule template.
//
// @param request - DeleteDataQualityRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityRuleTemplateResponse
func (client *Client) DeleteDataQualityRuleTemplateWithOptions(request *DeleteDataQualityRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityRuleTemplateResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityRuleTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityRuleTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitoring rule template.
//
// @param request - DeleteDataQualityRuleTemplateRequest
//
// @return DeleteDataQualityRuleTemplateResponse
func (client *Client) DeleteDataQualityRuleTemplate(request *DeleteDataQualityRuleTemplateRequest) (_result *DeleteDataQualityRuleTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataQualityRuleTemplateResponse{}
	_body, _err := client.DeleteDataQualityRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitor.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteDataQualityScanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityScanResponse
func (client *Client) DeleteDataQualityScanWithOptions(request *DeleteDataQualityScanRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityScanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataQualityScan"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality monitor.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteDataQualityScanRequest
//
// @return DeleteDataQualityScanResponse
func (client *Client) DeleteDataQualityScan(request *DeleteDataQualityScanRequest) (_result *DeleteDataQualityScanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataQualityScanResponse{}
	_body, _err := client.DeleteDataQualityScanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data quality rule template by ID.
//
// Description:
//
// ## [](#)Request description
//
//   - **Id**: the unique identifier of the user-defined rule template, in the format `USER_DEFINED:<template_id>`.
//
//   - **ProjectId**: The ID of the DataWorks project to which the rule template belongs.
//
// This API is used to remove data quality rule templates that are no longer needed from the system. Make sure the provided `Id` and `ProjectId` are correct when calling this API operation; otherwise, the deletion may fail or lead to unexpected data loss. Use this function with caution and verify the exact information of the template before performing the operation.
//
// @param request - DeleteDataQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataQualityTemplateResponse
func (client *Client) DeleteDataQualityTemplateWithOptions(request *DeleteDataQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataQualityTemplateResponse, _err error) {
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
		Action:      dara.String("DeleteDataQualityTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataQualityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data quality rule template by ID.
//
// Description:
//
// ## [](#)Request description
//
//   - **Id**: the unique identifier of the user-defined rule template, in the format `USER_DEFINED:<template_id>`.
//
//   - **ProjectId**: The ID of the DataWorks project to which the rule template belongs.
//
// This API is used to remove data quality rule templates that are no longer needed from the system. Make sure the provided `Id` and `ProjectId` are correct when calling this API operation; otherwise, the deletion may fail or lead to unexpected data loss. Use this function with caution and verify the exact information of the template before performing the operation.
//
// @param request - DeleteDataQualityTemplateRequest
//
// @return DeleteDataQualityTemplateResponse
func (client *Client) DeleteDataQualityTemplate(request *DeleteDataQualityTemplateRequest) (_result *DeleteDataQualityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataQualityTemplateResponse{}
	_body, _err := client.DeleteDataQualityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a data source by ID.
//
// Description:
//
// 1.  This API operation is available for all Dataworks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
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
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// Removes a data source by ID.
//
// Description:
//
// 1.  This API operation is available for all Dataworks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
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
// Deletes a sharing rule of a data source by ID.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  If you want to delete a sharing rule of a data source from Workspace A to Workspace B, you must have the permissions to share the data source in Workspace A or Workspace B. You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Tenant Administrator, Workspace Administrator, and Workspace Owner
//
// @param request - DeleteDataSourceSharedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceSharedRuleResponse
func (client *Client) DeleteDataSourceSharedRuleWithOptions(request *DeleteDataSourceSharedRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceSharedRuleResponse, _err error) {
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
		Action:      dara.String("DeleteDataSourceSharedRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceSharedRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a sharing rule of a data source by ID.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  If you want to delete a sharing rule of a data source from Workspace A to Workspace B, you must have the permissions to share the data source in Workspace A or Workspace B. You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Tenant Administrator, Workspace Administrator, and Workspace Owner
//
// @param request - DeleteDataSourceSharedRuleRequest
//
// @return DeleteDataSourceSharedRuleResponse
func (client *Client) DeleteDataSourceSharedRule(request *DeleteDataSourceSharedRuleRequest) (_result *DeleteDataSourceSharedRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSourceSharedRuleResponse{}
	_body, _err := client.DeleteDataSourceSharedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFileRequest
//
// @return DeleteFileResponse
func (client *Client) DeleteFile(request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithOptions(request *DeleteFolderRequest, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFolder"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFolderRequest
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolder(request *DeleteFolderRequest) (_result *DeleteFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFolderResponse{}
	_body, _err := client.DeleteFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function (UDF) in DataStudio.
//
// Description:
//
// >  A UDF that is deployed cannot be deleted. If you want to delete such a UDF, you must first undeploy the UDF.
//
// @param request - DeleteFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunctionWithOptions(request *DeleteFunctionRequest, runtime *dara.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function (UDF) in DataStudio.
//
// Description:
//
// >  A UDF that is deployed cannot be deleted. If you want to delete such a UDF, you must first undeploy the UDF.
//
// @param request - DeleteFunctionRequest
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunction(request *DeleteFunctionRequest) (_result *DeleteFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DeleteFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除血缘关系
//
// @param request - DeleteLineageRelationshipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLineageRelationshipResponse
func (client *Client) DeleteLineageRelationshipWithOptions(request *DeleteLineageRelationshipRequest, runtime *dara.RuntimeOptions) (_result *DeleteLineageRelationshipResponse, _err error) {
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
		Action:      dara.String("DeleteLineageRelationship"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLineageRelationshipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除血缘关系
//
// @param request - DeleteLineageRelationshipRequest
//
// @return DeleteLineageRelationshipResponse
func (client *Client) DeleteLineageRelationship(request *DeleteLineageRelationshipRequest) (_result *DeleteLineageRelationshipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLineageRelationshipResponse{}
	_body, _err := client.DeleteLineageRelationshipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除集合
//
// @param request - DeleteMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetaCollectionResponse
func (client *Client) DeleteMetaCollectionWithOptions(request *DeleteMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetaCollectionResponse, _err error) {
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
		Action:      dara.String("DeleteMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除集合
//
// @param request - DeleteMetaCollectionRequest
//
// @return DeleteMetaCollectionResponse
func (client *Client) DeleteMetaCollection(request *DeleteMetaCollectionRequest) (_result *DeleteMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMetaCollectionResponse{}
	_body, _err := client.DeleteMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates and deletes a network from a general resource group.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkResponse
func (client *Client) DeleteNetworkWithOptions(request *DeleteNetworkRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkResponse, _err error) {
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
		Action:      dara.String("DeleteNetwork"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates and deletes a network from a general resource group.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteNetworkRequest
//
// @return DeleteNetworkResponse
func (client *Client) DeleteNetwork(request *DeleteNetworkRequest) (_result *DeleteNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkResponse{}
	_body, _err := client.DeleteNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a node from DataStudio.
//
// Description:
//
// >  A node that is deployed cannot be deleted. If you want to delete such a node, you must first undeploy the node.
//
// @param request - DeleteNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeResponse
func (client *Client) DeleteNodeWithOptions(request *DeleteNodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteNodeResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a node from DataStudio.
//
// Description:
//
// >  A node that is deployed cannot be deleted. If you want to delete such a node, you must first undeploy the node.
//
// @param request - DeleteNodeRequest
//
// @return DeleteNodeResponse
func (client *Client) DeleteNode(request *DeleteNodeRequest) (_result *DeleteNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNodeResponse{}
	_body, _err := client.DeleteNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a DataWorks workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
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
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a DataWorks workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteProjectRequest
//
// @return DeleteProjectResponse
func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a workspace member and the workspace-level roles that are assigned to the member.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectMemberResponse
func (client *Client) DeleteProjectMemberWithOptions(request *DeleteProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProjectMember"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a workspace member and the workspace-level roles that are assigned to the member.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteProjectMemberRequest
//
// @return DeleteProjectMemberResponse
func (client *Client) DeleteProjectMember(request *DeleteProjectMemberRequest) (_result *DeleteProjectMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProjectMemberResponse{}
	_body, _err := client.DeleteProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a file resource from DataStudio.
//
// Description:
//
// >  A file resource that is deployed cannot be deleted. If you want to delete such a file resource, you must first undeploy the file resource.
//
// @param request - DeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithOptions(request *DeleteResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file resource from DataStudio.
//
// Description:
//
// >  A file resource that is deployed cannot be deleted. If you want to delete such a file resource, you must first undeploy the file resource.
//
// @param request - DeleteResourceRequest
//
// @return DeleteResourceResponse
func (client *Client) DeleteResource(request *DeleteResourceRequest) (_result *DeleteResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a serverless resource group.
//
// Description:
//
// 1.  You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// 2.  **Before you call this API operation, you must make sure that you have a good command of the billing details and [pricing](https://help.aliyun.com/document_detail/2680173.html) of serverless resource groups.
//
// @param request - DeleteResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithOptions(request *DeleteResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
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
		Action:      dara.String("DeleteResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a serverless resource group.
//
// Description:
//
// 1.  You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// 2.  **Before you call this API operation, you must make sure that you have a good command of the billing details and [pricing](https://help.aliyun.com/document_detail/2680173.html) of serverless resource groups.
//
// @param request - DeleteResourceGroupRequest
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (_result *DeleteResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DeleteResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a route from a network resource.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteResponse
func (client *Client) DeleteRouteWithOptions(request *DeleteRouteRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteResponse, _err error) {
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
		Action:      dara.String("DeleteRoute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a route from a network resource.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteRouteRequest
//
// @return DeleteRouteResponse
func (client *Client) DeleteRoute(request *DeleteRouteRequest) (_result *DeleteRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRouteResponse{}
	_body, _err := client.DeleteRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskResponse
func (client *Client) DeleteTaskWithOptions(request *DeleteTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteTaskResponse, _err error) {
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

	if !dara.IsNil(request.ProjectEnv) {
		query["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteTaskRequest
//
// @return DeleteTaskResponse
func (client *Client) DeleteTask(request *DeleteTaskRequest) (_result *DeleteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTaskResponse{}
	_body, _err := client.DeleteTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workflow.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowResponse
func (client *Client) DeleteWorkflowWithOptions(request *DeleteWorkflowRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkflowResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientUniqueCode) {
		body["ClientUniqueCode"] = request.ClientUniqueCode
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkflow"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workflow.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - DeleteWorkflowRequest
//
// @return DeleteWorkflowResponse
func (client *Client) DeleteWorkflow(request *DeleteWorkflowRequest) (_result *DeleteWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWorkflowResponse{}
	_body, _err := client.DeleteWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workflow from DataStudio.
//
// Description:
//
// >  A workflow that is deployed cannot be deleted. If you want to delete such a workflow, you must first undeploy the workflow.
//
// @param request - DeleteWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowDefinitionResponse
func (client *Client) DeleteWorkflowDefinitionWithOptions(request *DeleteWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkflowDefinitionResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workflow from DataStudio.
//
// Description:
//
// >  A workflow that is deployed cannot be deleted. If you want to delete such a workflow, you must first undeploy the workflow.
//
// @param request - DeleteWorkflowDefinitionRequest
//
// @return DeleteWorkflowDefinitionResponse
func (client *Client) DeleteWorkflowDefinition(request *DeleteWorkflowDefinitionRequest) (_result *DeleteWorkflowDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWorkflowDefinitionResponse{}
	_body, _err := client.DeleteWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeployFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployFileResponse
func (client *Client) DeployFileWithOptions(request *DeployFileRequest, runtime *dara.RuntimeOptions) (_result *DeployFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeployFileRequest
//
// @return DeployFileResponse
func (client *Client) DeployFile(request *DeployFileRequest) (_result *DeployFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeployFileResponse{}
	_body, _err := client.DeployFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates monitoring rules from a data quality monitoring task.
//
// @param tmpReq - DetachDataQualityRulesFromEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachDataQualityRulesFromEvaluationTaskResponse
func (client *Client) DetachDataQualityRulesFromEvaluationTaskWithOptions(tmpReq *DetachDataQualityRulesFromEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *DetachDataQualityRulesFromEvaluationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetachDataQualityRulesFromEvaluationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataQualityRuleIds) {
		request.DataQualityRuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataQualityRuleIds, dara.String("DataQualityRuleIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityEvaluationTaskId) {
		body["DataQualityEvaluationTaskId"] = request.DataQualityEvaluationTaskId
	}

	if !dara.IsNil(request.DataQualityRuleIdsShrink) {
		body["DataQualityRuleIds"] = request.DataQualityRuleIdsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachDataQualityRulesFromEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachDataQualityRulesFromEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates monitoring rules from a data quality monitoring task.
//
// @param request - DetachDataQualityRulesFromEvaluationTaskRequest
//
// @return DetachDataQualityRulesFromEvaluationTaskResponse
func (client *Client) DetachDataQualityRulesFromEvaluationTask(request *DetachDataQualityRulesFromEvaluationTaskRequest) (_result *DetachDataQualityRulesFromEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachDataQualityRulesFromEvaluationTaskResponse{}
	_body, _err := client.DetachDataQualityRulesFromEvaluationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a resource group from a workspace.
//
// Description:
//
// 1.  You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// 2.  Your account must be assigned one of the following roles of the desired workspace:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
//
// @param request - DissociateProjectFromResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateProjectFromResourceGroupResponse
func (client *Client) DissociateProjectFromResourceGroupWithOptions(request *DissociateProjectFromResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DissociateProjectFromResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateProjectFromResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateProjectFromResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a resource group from a workspace.
//
// Description:
//
// 1.  You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// 2.  Your account must be assigned one of the following roles of the desired workspace:
//
//   - Tenant Owner, Workspace Administrator, Workspace Owner, and O\\&M
//
// @param request - DissociateProjectFromResourceGroupRequest
//
// @return DissociateProjectFromResourceGroupResponse
func (client *Client) DissociateProjectFromResourceGroup(request *DissociateProjectFromResourceGroupRequest) (_result *DissociateProjectFromResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DissociateProjectFromResourceGroupResponse{}
	_body, _err := client.DissociateProjectFromResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EstablishRelationTableToBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstablishRelationTableToBusinessResponse
func (client *Client) EstablishRelationTableToBusinessWithOptions(request *EstablishRelationTableToBusinessRequest, runtime *dara.RuntimeOptions) (_result *EstablishRelationTableToBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.TableGuid) {
		body["TableGuid"] = request.TableGuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EstablishRelationTableToBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EstablishRelationTableToBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EstablishRelationTableToBusinessRequest
//
// @return EstablishRelationTableToBusinessResponse
func (client *Client) EstablishRelationTableToBusiness(request *EstablishRelationTableToBusinessRequest) (_result *EstablishRelationTableToBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EstablishRelationTableToBusinessResponse{}
	_body, _err := client.EstablishRelationTableToBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a stage in a process.
//
// Description:
//
// >  The stages in a process are sequential. For more information, see the GetDeployment operation. Skipping or repeating a stage is not allowed.
//
// >  The execution of a stage is asynchronous. The response of this operation indicates only whether a stage is triggered but does not indicate whether the execution of the stage is successful. You can call the GetDeployment operation to check whether the execution is successful.
//
// @param request - ExecPipelineRunStageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecPipelineRunStageResponse
func (client *Client) ExecPipelineRunStageWithOptions(request *ExecPipelineRunStageRequest, runtime *dara.RuntimeOptions) (_result *ExecPipelineRunStageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecPipelineRunStage"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecPipelineRunStageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a stage in a process.
//
// Description:
//
// >  The stages in a process are sequential. For more information, see the GetDeployment operation. Skipping or repeating a stage is not allowed.
//
// >  The execution of a stage is asynchronous. The response of this operation indicates only whether a stage is triggered but does not indicate whether the execution of the stage is successful. You can call the GetDeployment operation to check whether the execution is successful.
//
// @param request - ExecPipelineRunStageRequest
//
// @return ExecPipelineRunStageResponse
func (client *Client) ExecPipelineRunStage(request *ExecPipelineRunStageRequest) (_result *ExecPipelineRunStageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecPipelineRunStageResponse{}
	_body, _err := client.ExecPipelineRunStageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a temporary workflow instance based on configurations.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ExecuteAdhocWorkflowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAdhocWorkflowInstanceResponse
func (client *Client) ExecuteAdhocWorkflowInstanceWithOptions(tmpReq *ExecuteAdhocWorkflowInstanceRequest, runtime *dara.RuntimeOptions) (_result *ExecuteAdhocWorkflowInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExecuteAdhocWorkflowInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tasks) {
		request.TasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tasks, dara.String("Tasks"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		body["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TasksShrink) {
		body["Tasks"] = request.TasksShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAdhocWorkflowInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAdhocWorkflowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a temporary workflow instance based on configurations.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ExecuteAdhocWorkflowInstanceRequest
//
// @return ExecuteAdhocWorkflowInstanceResponse
func (client *Client) ExecuteAdhocWorkflowInstance(request *ExecuteAdhocWorkflowInstanceRequest) (_result *ExecuteAdhocWorkflowInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteAdhocWorkflowInstanceResponse{}
	_body, _err := client.ExecuteAdhocWorkflowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a custom alert monitoring rule.
//
// @param request - GetAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertRuleResponse
func (client *Client) GetAlertRuleWithOptions(request *GetAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAlertRuleResponse, _err error) {
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
		Action:      dara.String("GetAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a custom alert monitoring rule.
//
// @param request - GetAlertRuleRequest
//
// @return GetAlertRuleResponse
func (client *Client) GetAlertRule(request *GetAlertRuleRequest) (_result *GetAlertRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAlertRuleResponse{}
	_body, _err := client.GetAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBusinessResponse
func (client *Client) GetBusinessWithOptions(request *GetBusinessRequest, runtime *dara.RuntimeOptions) (_result *GetBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetBusinessRequest
//
// @return GetBusinessResponse
func (client *Client) GetBusiness(request *GetBusinessRequest) (_result *GetBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBusinessResponse{}
	_body, _err := client.GetBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据目录详情
//
// @param request - GetCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogResponse
func (client *Client) GetCatalogWithOptions(request *GetCatalogRequest, runtime *dara.RuntimeOptions) (_result *GetCatalogResponse, _err error) {
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
		Action:      dara.String("GetCatalog"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据目录详情
//
// @param request - GetCatalogRequest
//
// @return GetCatalogResponse
func (client *Client) GetCatalog(request *GetCatalogRequest) (_result *GetCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCatalogResponse{}
	_body, _err := client.GetCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a certificate file.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks: Tenant Owner, Workspace Administrator, Deploy, Develop, Workspace Owner, and O\\&M.
//
// @param request - GetCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCertificateResponse
func (client *Client) GetCertificateWithOptions(request *GetCertificateRequest, runtime *dara.RuntimeOptions) (_result *GetCertificateResponse, _err error) {
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
		Action:      dara.String("GetCertificate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a certificate file.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks: Tenant Owner, Workspace Administrator, Deploy, Develop, Workspace Owner, and O\\&M.
//
// @param request - GetCertificateRequest
//
// @return GetCertificateResponse
func (client *Client) GetCertificate(request *GetCertificateRequest) (_result *GetCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCertificateResponse{}
	_body, _err := client.GetCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取字段详情
//
// @param request - GetColumnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetColumnResponse
func (client *Client) GetColumnWithOptions(request *GetColumnRequest, runtime *dara.RuntimeOptions) (_result *GetColumnResponse, _err error) {
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
		Action:      dara.String("GetColumn"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetColumnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取字段详情
//
// @param request - GetColumnRequest
//
// @return GetColumnResponse
func (client *Client) GetColumn(request *GetColumnRequest) (_result *GetColumnResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetColumnResponse{}
	_body, _err := client.GetColumnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets component information.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// @param request - GetComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetComponentResponse
func (client *Client) GetComponentWithOptions(request *GetComponentRequest, runtime *dara.RuntimeOptions) (_result *GetComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComponentId) {
		query["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetComponent"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetComponentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets component information.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// @param request - GetComponentRequest
//
// @return GetComponentResponse
func (client *Client) GetComponent(request *GetComponentRequest) (_result *GetComponentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetComponentResponse{}
	_body, _err := client.GetComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of asynchronously creating a workflow instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetCreateWorkflowInstancesResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateWorkflowInstancesResultResponse
func (client *Client) GetCreateWorkflowInstancesResultWithOptions(request *GetCreateWorkflowInstancesResultRequest, runtime *dara.RuntimeOptions) (_result *GetCreateWorkflowInstancesResultResponse, _err error) {
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
		Action:      dara.String("GetCreateWorkflowInstancesResult"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCreateWorkflowInstancesResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of asynchronously creating a workflow instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetCreateWorkflowInstancesResultRequest
//
// @return GetCreateWorkflowInstancesResultResponse
func (client *Client) GetCreateWorkflowInstancesResult(request *GetCreateWorkflowInstancesResultRequest) (_result *GetCreateWorkflowInstancesResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCreateWorkflowInstancesResultResponse{}
	_body, _err := client.GetCreateWorkflowInstancesResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIJobResponse
func (client *Client) GetDIJobWithOptions(request *GetDIJobRequest, runtime *dara.RuntimeOptions) (_result *GetDIJobResponse, _err error) {
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
		Action:      dara.String("GetDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDIJobRequest
//
// @return GetDIJobResponse
func (client *Client) GetDIJob(request *GetDIJobRequest) (_result *GetDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDIJobResponse{}
	_body, _err := client.GetDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains logs generated for a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDIJobLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIJobLogResponse
func (client *Client) GetDIJobLogWithOptions(request *GetDIJobLogRequest, runtime *dara.RuntimeOptions) (_result *GetDIJobLogResponse, _err error) {
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
		Action:      dara.String("GetDIJobLog"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDIJobLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains logs generated for a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDIJobLogRequest
//
// @return GetDIJobLogResponse
func (client *Client) GetDIJobLog(request *GetDIJobLogRequest) (_result *GetDIJobLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDIJobLogResponse{}
	_body, _err := client.GetDIJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data quality monitoring and alerting rule by alert rule ID.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityAlertRuleResponse
func (client *Client) GetDataQualityAlertRuleWithOptions(request *GetDataQualityAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityAlertRuleResponse, _err error) {
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
		Action:      dara.String("GetDataQualityAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data quality monitoring and alerting rule by alert rule ID.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityAlertRuleRequest
//
// @return GetDataQualityAlertRuleResponse
func (client *Client) GetDataQualityAlertRule(request *GetDataQualityAlertRuleRequest) (_result *GetDataQualityAlertRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataQualityAlertRuleResponse{}
	_body, _err := client.GetDataQualityAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a monitor.
//
// @param request - GetDataQualityEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityEvaluationTaskResponse
func (client *Client) GetDataQualityEvaluationTaskWithOptions(request *GetDataQualityEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityEvaluationTaskResponse, _err error) {
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
		Action:      dara.String("GetDataQualityEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a monitor.
//
// @param request - GetDataQualityEvaluationTaskRequest
//
// @return GetDataQualityEvaluationTaskResponse
func (client *Client) GetDataQualityEvaluationTask(request *GetDataQualityEvaluationTaskRequest) (_result *GetDataQualityEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataQualityEvaluationTaskResponse{}
	_body, _err := client.GetDataQualityEvaluationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a monitor instance.
//
// @param request - GetDataQualityEvaluationTaskInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityEvaluationTaskInstanceResponse
func (client *Client) GetDataQualityEvaluationTaskInstanceWithOptions(request *GetDataQualityEvaluationTaskInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityEvaluationTaskInstanceResponse, _err error) {
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
		Action:      dara.String("GetDataQualityEvaluationTaskInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityEvaluationTaskInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a monitor instance.
//
// @param request - GetDataQualityEvaluationTaskInstanceRequest
//
// @return GetDataQualityEvaluationTaskInstanceResponse
func (client *Client) GetDataQualityEvaluationTaskInstance(request *GetDataQualityEvaluationTaskInstanceRequest) (_result *GetDataQualityEvaluationTaskInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataQualityEvaluationTaskInstanceResponse{}
	_body, _err := client.GetDataQualityEvaluationTaskInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a data quality monitoring rule.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityRuleResponse
func (client *Client) GetDataQualityRuleWithOptions(request *GetDataQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityRuleResponse, _err error) {
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
		Action:      dara.String("GetDataQualityRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a data quality monitoring rule.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityRuleRequest
//
// @return GetDataQualityRuleResponse
func (client *Client) GetDataQualityRule(request *GetDataQualityRuleRequest) (_result *GetDataQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataQualityRuleResponse{}
	_body, _err := client.GetDataQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a data quality monitoring rule template.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityRuleTemplateResponse
func (client *Client) GetDataQualityRuleTemplateWithOptions(request *GetDataQualityRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityRuleTemplateResponse, _err error) {
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
		Action:      dara.String("GetDataQualityRuleTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityRuleTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a data quality monitoring rule template.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityRuleTemplateRequest
//
// @return GetDataQualityRuleTemplateResponse
func (client *Client) GetDataQualityRuleTemplate(request *GetDataQualityRuleTemplateRequest) (_result *GetDataQualityRuleTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataQualityRuleTemplateResponse{}
	_body, _err := client.GetDataQualityRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets data quality monitoring details.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityScanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityScanResponse
func (client *Client) GetDataQualityScanWithOptions(request *GetDataQualityScanRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityScanResponse, _err error) {
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
		Action:      dara.String("GetDataQualityScan"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets data quality monitoring details.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityScanRequest
//
// @return GetDataQualityScanResponse
func (client *Client) GetDataQualityScan(request *GetDataQualityScanRequest) (_result *GetDataQualityScanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataQualityScanResponse{}
	_body, _err := client.GetDataQualityScanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring run instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityScanRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityScanRunResponse
func (client *Client) GetDataQualityScanRunWithOptions(request *GetDataQualityScanRunRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityScanRunResponse, _err error) {
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
		Action:      dara.String("GetDataQualityScanRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityScanRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data quality monitoring run instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityScanRunRequest
//
// @return GetDataQualityScanRunResponse
func (client *Client) GetDataQualityScanRun(request *GetDataQualityScanRunRequest) (_result *GetDataQualityScanRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataQualityScanRunResponse{}
	_body, _err := client.GetDataQualityScanRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log of a specific task instance that monitors data quality.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityScanRunLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityScanRunLogResponse
func (client *Client) GetDataQualityScanRunLogWithOptions(request *GetDataQualityScanRunLogRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityScanRunLogResponse, _err error) {
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

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataQualityScanRunLog"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityScanRunLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log of a specific task instance that monitors data quality.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityScanRunLogRequest
//
// @return GetDataQualityScanRunLogResponse
func (client *Client) GetDataQualityScanRunLog(request *GetDataQualityScanRunLogRequest) (_result *GetDataQualityScanRunLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataQualityScanRunLogResponse{}
	_body, _err := client.GetDataQualityScanRunLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data quality rule template by ID.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityTemplateResponse
func (client *Client) GetDataQualityTemplateWithOptions(request *GetDataQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetDataQualityTemplateResponse, _err error) {
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
		Action:      dara.String("GetDataQualityTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data quality rule template by ID.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetDataQualityTemplateRequest
//
// @return GetDataQualityTemplateResponse
func (client *Client) GetDataQualityTemplate(request *GetDataQualityTemplateRequest) (_result *GetDataQualityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataQualityTemplateResponse{}
	_body, _err := client.GetDataQualityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a data source by ID.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Deployment, Development, Project Owner, and O\\&M
//
// @param request - GetDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceResponse
func (client *Client) GetDataSourceWithOptions(request *GetDataSourceRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceResponse, _err error) {
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
		Action:      dara.String("GetDataSource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a data source by ID.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Deployment, Development, Project Owner, and O\\&M
//
// @param request - GetDataSourceRequest
//
// @return GetDataSourceResponse
func (client *Client) GetDataSource(request *GetDataSourceRequest) (_result *GetDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataSourceResponse{}
	_body, _err := client.GetDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据库详情
//
// @param request - GetDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseResponse
func (client *Client) GetDatabaseWithOptions(request *GetDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseResponse, _err error) {
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
		Action:      dara.String("GetDatabase"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据库详情
//
// @param request - GetDatabaseRequest
//
// @return GetDatabaseResponse
func (client *Client) GetDatabase(request *GetDatabaseRequest) (_result *GetDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatabaseResponse{}
	_body, _err := client.GetDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetDeploymentPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentPackageResponse
func (client *Client) GetDeploymentPackageWithOptions(request *GetDeploymentPackageRequest, runtime *dara.RuntimeOptions) (_result *GetDeploymentPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		body["DeploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeploymentPackage"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetDeploymentPackageRequest
//
// @return GetDeploymentPackageResponse
func (client *Client) GetDeploymentPackage(request *GetDeploymentPackageRequest) (_result *GetDeploymentPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeploymentPackageResponse{}
	_body, _err := client.GetDeploymentPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a file.
//
// @param request - GetFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileResponse
func (client *Client) GetFileWithOptions(request *GetFileRequest, runtime *dara.RuntimeOptions) (_result *GetFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a file.
//
// @param request - GetFileRequest
//
// @return GetFileResponse
func (client *Client) GetFile(request *GetFileRequest) (_result *GetFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileResponse{}
	_body, _err := client.GetFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a file version.
//
// @param request - GetFileVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileVersionResponse
func (client *Client) GetFileVersionWithOptions(request *GetFileVersionRequest, runtime *dara.RuntimeOptions) (_result *GetFileVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileVersion) {
		body["FileVersion"] = request.FileVersion
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileVersion"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a file version.
//
// @param request - GetFileVersionRequest
//
// @return GetFileVersionResponse
func (client *Client) GetFileVersion(request *GetFileVersionRequest) (_result *GetFileVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileVersionResponse{}
	_body, _err := client.GetFileVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFolderResponse
func (client *Client) GetFolderWithOptions(request *GetFolderRequest, runtime *dara.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.FolderPath) {
		body["FolderPath"] = request.FolderPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFolder"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetFolderRequest
//
// @return GetFolderResponse
func (client *Client) GetFolder(request *GetFolderRequest) (_result *GetFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFolderResponse{}
	_body, _err := client.GetFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a user-defined function (UDF) in DataStudio.
//
// @param request - GetFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionResponse
func (client *Client) GetFunctionWithOptions(request *GetFunctionRequest, runtime *dara.RuntimeOptions) (_result *GetFunctionResponse, _err error) {
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
		Action:      dara.String("GetFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a user-defined function (UDF) in DataStudio.
//
// @param request - GetFunctionRequest
//
// @return GetFunctionResponse
func (client *Client) GetFunction(request *GetFunctionRequest) (_result *GetFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFunctionResponse{}
	_body, _err := client.GetFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data snapshot of an extension point based on the ID of a message in DataWorks OpenEvent when the related extension point event is triggered.
//
// @param request - GetIDEEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIDEEventDetailResponse
func (client *Client) GetIDEEventDetailWithOptions(request *GetIDEEventDetailRequest, runtime *dara.RuntimeOptions) (_result *GetIDEEventDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIDEEventDetail"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIDEEventDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data snapshot of an extension point based on the ID of a message in DataWorks OpenEvent when the related extension point event is triggered.
//
// @param request - GetIDEEventDetailRequest
//
// @return GetIDEEventDetailResponse
func (client *Client) GetIDEEventDetail(request *GetIDEEventDetailRequest) (_result *GetIDEEventDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIDEEventDetailResponse{}
	_body, _err := client.GetIDEEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status information of an asynchronous task. After you call an asynchronous operation, an asynchronous task is generated. You can call the GetJobStatus operation to query the status of the asynchronous task.
//
// @param request - GetJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobStatusResponse
func (client *Client) GetJobStatusWithOptions(request *GetJobStatusRequest, runtime *dara.RuntimeOptions) (_result *GetJobStatusResponse, _err error) {
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
		Action:      dara.String("GetJobStatus"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status information of an asynchronous task. After you call an asynchronous operation, an asynchronous task is generated. You can call the GetJobStatus operation to query the status of the asynchronous task.
//
// @param request - GetJobStatusRequest
//
// @return GetJobStatusResponse
func (client *Client) GetJobStatus(request *GetJobStatusRequest) (_result *GetJobStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobStatusResponse{}
	_body, _err := client.GetJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取血缘关系详情
//
// @param request - GetLineageRelationshipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLineageRelationshipResponse
func (client *Client) GetLineageRelationshipWithOptions(request *GetLineageRelationshipRequest, runtime *dara.RuntimeOptions) (_result *GetLineageRelationshipResponse, _err error) {
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
		Action:      dara.String("GetLineageRelationship"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLineageRelationshipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取血缘关系详情
//
// @param request - GetLineageRelationshipRequest
//
// @return GetLineageRelationshipResponse
func (client *Client) GetLineageRelationship(request *GetLineageRelationshipRequest) (_result *GetLineageRelationshipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLineageRelationshipResponse{}
	_body, _err := client.GetLineageRelationshipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a collection in Data Map. Collections include categories and data albums.
//
// @param request - GetMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaCollectionResponse
func (client *Client) GetMetaCollectionWithOptions(request *GetMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *GetMetaCollectionResponse, _err error) {
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
		Action:      dara.String("GetMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a collection in Data Map. Collections include categories and data albums.
//
// @param request - GetMetaCollectionRequest
//
// @return GetMetaCollectionResponse
func (client *Client) GetMetaCollection(request *GetMetaCollectionRequest) (_result *GetMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaCollectionResponse{}
	_body, _err := client.GetMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a network resource.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkResponse
func (client *Client) GetNetworkWithOptions(request *GetNetworkRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkResponse, _err error) {
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
		Action:      dara.String("GetNetwork"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a network resource.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetNetworkRequest
//
// @return GetNetworkResponse
func (client *Client) GetNetwork(request *GetNetworkRequest) (_result *GetNetworkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNetworkResponse{}
	_body, _err := client.GetNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a node in DataStudio.
//
// @param request - GetNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeResponse
func (client *Client) GetNodeWithOptions(request *GetNodeRequest, runtime *dara.RuntimeOptions) (_result *GetNodeResponse, _err error) {
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
		Action:      dara.String("GetNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a node in DataStudio.
//
// @param request - GetNodeRequest
//
// @return GetNodeResponse
func (client *Client) GetNode(request *GetNodeRequest) (_result *GetNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNodeResponse{}
	_body, _err := client.GetNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves partition details for a data map table. Currently supports MaxCompute and HMS (EMR cluster) types only.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  This operation supports MaxCompute and HMS (EMR cluster) tables only.
//
// @param request - GetPartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPartitionResponse
func (client *Client) GetPartitionWithOptions(request *GetPartitionRequest, runtime *dara.RuntimeOptions) (_result *GetPartitionResponse, _err error) {
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
		Action:      dara.String("GetPartition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves partition details for a data map table. Currently supports MaxCompute and HMS (EMR cluster) types only.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  This operation supports MaxCompute and HMS (EMR cluster) tables only.
//
// @param request - GetPartitionRequest
//
// @return GetPartitionResponse
func (client *Client) GetPartition(request *GetPartitionRequest) (_result *GetPartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPartitionResponse{}
	_body, _err := client.GetPartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a process for deploying or undeploying an entity.
//
// @param request - GetPipelineRunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineRunResponse
func (client *Client) GetPipelineRunWithOptions(request *GetPipelineRunRequest, runtime *dara.RuntimeOptions) (_result *GetPipelineRunResponse, _err error) {
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
		Action:      dara.String("GetPipelineRun"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a process for deploying or undeploying an entity.
//
// @param request - GetPipelineRunRequest
//
// @return GetPipelineRunResponse
func (client *Client) GetPipelineRun(request *GetPipelineRunRequest) (_result *GetPipelineRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPipelineRunResponse{}
	_body, _err := client.GetPipelineRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a DataWorks workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
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
		Action:      dara.String("GetProject"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a DataWorks workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetProjectRequest
//
// @return GetProjectResponse
func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a member in a workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectMemberResponse
func (client *Client) GetProjectMemberWithOptions(request *GetProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *GetProjectMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectMember"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a member in a workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetProjectMemberRequest
//
// @return GetProjectMemberResponse
func (client *Client) GetProjectMember(request *GetProjectMemberRequest) (_result *GetProjectMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectMemberResponse{}
	_body, _err := client.GetProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a role in a DataWorks workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetProjectRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectRoleResponse
func (client *Client) GetProjectRoleWithOptions(request *GetProjectRoleRequest, runtime *dara.RuntimeOptions) (_result *GetProjectRoleResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectRole"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a role in a DataWorks workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetProjectRoleRequest
//
// @return GetProjectRoleResponse
func (client *Client) GetProjectRole(request *GetProjectRoleRequest) (_result *GetProjectRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectRoleResponse{}
	_body, _err := client.GetProjectRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the result of asynchronous workflow instance reruns.
//
// @param request - GetRerunWorkflowInstancesResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRerunWorkflowInstancesResultResponse
func (client *Client) GetRerunWorkflowInstancesResultWithOptions(request *GetRerunWorkflowInstancesResultRequest, runtime *dara.RuntimeOptions) (_result *GetRerunWorkflowInstancesResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRerunWorkflowInstancesResult"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRerunWorkflowInstancesResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the result of asynchronous workflow instance reruns.
//
// @param request - GetRerunWorkflowInstancesResultRequest
//
// @return GetRerunWorkflowInstancesResultResponse
func (client *Client) GetRerunWorkflowInstancesResult(request *GetRerunWorkflowInstancesResultRequest) (_result *GetRerunWorkflowInstancesResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRerunWorkflowInstancesResultResponse{}
	_body, _err := client.GetRerunWorkflowInstancesResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a file resource.
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithOptions(request *GetResourceRequest, runtime *dara.RuntimeOptions) (_result *GetResourceResponse, _err error) {
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
		Action:      dara.String("GetResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a file resource.
//
// @param request - GetResourceRequest
//
// @return GetResourceResponse
func (client *Client) GetResource(request *GetResourceRequest) (_result *GetResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceResponse{}
	_body, _err := client.GetResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a resource group based on its ID.
//
// Description:
//
// You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// @param request - GetResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroupWithOptions(request *GetResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
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
		Action:      dara.String("GetResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a resource group based on its ID.
//
// Description:
//
// You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// @param request - GetResourceGroupRequest
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroup(request *GetResourceGroupRequest) (_result *GetResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.GetResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a route based on its ID.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRouteResponse
func (client *Client) GetRouteWithOptions(request *GetRouteRequest, runtime *dara.RuntimeOptions) (_result *GetRouteResponse, _err error) {
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
		Action:      dara.String("GetRoute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a route based on its ID.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetRouteRequest
//
// @return GetRouteResponse
func (client *Client) GetRoute(request *GetRouteRequest) (_result *GetRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRouteResponse{}
	_body, _err := client.GetRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a schema in Data Map. You can call this API operation to query the information only about MaxCompute and Hologres schemas.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this API operation to query the information only about MaxCompute and Hologres schemas.
//
// @param request - GetSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSchemaResponse
func (client *Client) GetSchemaWithOptions(request *GetSchemaRequest, runtime *dara.RuntimeOptions) (_result *GetSchemaResponse, _err error) {
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
		Action:      dara.String("GetSchema"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a schema in Data Map. You can call this API operation to query the information only about MaxCompute and Hologres schemas.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this API operation to query the information only about MaxCompute and Hologres schemas.
//
// @param request - GetSchemaRequest
//
// @return GetSchemaResponse
func (client *Client) GetSchema(request *GetSchemaRequest) (_result *GetSchemaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSchemaResponse{}
	_body, _err := client.GetSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取表详情
//
// @param request - GetTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableResponse
func (client *Client) GetTableWithOptions(request *GetTableRequest, runtime *dara.RuntimeOptions) (_result *GetTableResponse, _err error) {
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
		Action:      dara.String("GetTable"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取表详情
//
// @param request - GetTableRequest
//
// @return GetTableResponse
func (client *Client) GetTable(request *GetTableRequest) (_result *GetTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableResponse{}
	_body, _err := client.GetTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a task.
//
// @param request - GetTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
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
		Action:      dara.String("GetTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a task.
//
// @param request - GetTaskRequest
//
// @return GetTaskResponse
func (client *Client) GetTask(request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetTaskInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskInstanceResponse
func (client *Client) GetTaskInstanceWithOptions(request *GetTaskInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetTaskInstanceResponse, _err error) {
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
		Action:      dara.String("GetTaskInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetTaskInstanceRequest
//
// @return GetTaskInstanceResponse
func (client *Client) GetTaskInstance(request *GetTaskInstanceRequest) (_result *GetTaskInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskInstanceResponse{}
	_body, _err := client.GetTaskInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the run log generated during a specific run of an instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetTaskInstanceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskInstanceLogResponse
func (client *Client) GetTaskInstanceLogWithOptions(request *GetTaskInstanceLogRequest, runtime *dara.RuntimeOptions) (_result *GetTaskInstanceLogResponse, _err error) {
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
		Action:      dara.String("GetTaskInstanceLog"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskInstanceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the run log generated during a specific run of an instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetTaskInstanceLogRequest
//
// @return GetTaskInstanceLogResponse
func (client *Client) GetTaskInstanceLog(request *GetTaskInstanceLogRequest) (_result *GetTaskInstanceLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskInstanceLogResponse{}
	_body, _err := client.GetTaskInstanceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowResponse
func (client *Client) GetWorkflowWithOptions(request *GetWorkflowRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowResponse, _err error) {
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
		Action:      dara.String("GetWorkflow"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetWorkflowRequest
//
// @return GetWorkflowResponse
func (client *Client) GetWorkflow(request *GetWorkflowRequest) (_result *GetWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkflowResponse{}
	_body, _err := client.GetWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow.
//
// @param request - GetWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowDefinitionResponse
func (client *Client) GetWorkflowDefinitionWithOptions(request *GetWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowDefinitionResponse, _err error) {
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
		Action:      dara.String("GetWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow.
//
// @param request - GetWorkflowDefinitionRequest
//
// @return GetWorkflowDefinitionResponse
func (client *Client) GetWorkflowDefinition(request *GetWorkflowDefinitionRequest) (_result *GetWorkflowDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkflowDefinitionResponse{}
	_body, _err := client.GetWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetWorkflowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowInstanceResponse
func (client *Client) GetWorkflowInstanceWithOptions(request *GetWorkflowInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowInstanceResponse, _err error) {
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
		Action:      dara.String("GetWorkflowInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow instance.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GetWorkflowInstanceRequest
//
// @return GetWorkflowInstanceResponse
func (client *Client) GetWorkflowInstance(request *GetWorkflowInstanceRequest) (_result *GetWorkflowInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkflowInstanceResponse{}
	_body, _err := client.GetWorkflowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Assigns roles to members in a workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - GrantMemberProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantMemberProjectRolesResponse
func (client *Client) GrantMemberProjectRolesWithOptions(tmpReq *GrantMemberProjectRolesRequest, runtime *dara.RuntimeOptions) (_result *GrantMemberProjectRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GrantMemberProjectRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("RoleCodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantMemberProjectRoles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantMemberProjectRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns roles to members in a workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - GrantMemberProjectRolesRequest
//
// @return GrantMemberProjectRolesResponse
func (client *Client) GrantMemberProjectRoles(request *GrantMemberProjectRolesRequest) (_result *GrantMemberProjectRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantMemberProjectRolesResponse{}
	_body, _err := client.GrantMemberProjectRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports a certificate file.
//
// @param request - ImportCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportCertificateResponse
func (client *Client) ImportCertificateWithOptions(request *ImportCertificateRequest, runtime *dara.RuntimeOptions) (_result *ImportCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateFile) {
		query["CertificateFile"] = request.CertificateFile
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportCertificate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a certificate file.
//
// @param request - ImportCertificateRequest
//
// @return ImportCertificateResponse
func (client *Client) ImportCertificate(request *ImportCertificateRequest) (_result *ImportCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportCertificateResponse{}
	_body, _err := client.ImportCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportCertificateAdvance(request *ImportCertificateAdvanceRequest, runtime *dara.RuntimeOptions) (_result *ImportCertificateResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("dataworks-public"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	importCertificateReq := &ImportCertificateRequest{}
	openapiutil.Convert(request, importCertificateReq)
	if !dara.IsNil(request.CertificateFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.CertificateFileObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		importCertificateReq.CertificateFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	importCertificateResp, _err := client.ImportCertificateWithOptions(importCertificateReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = importCertificateResp
	return _result, _err
}

// Summary:
//
// Imports a workflow and its child nodes that are specified by the FlowSpec field to DataStudio.
//
// Description:
//
// >
//
//   - You cannot use this API operation to import multiple workflows at a time. If you specify multiple workflows by using FlowSpec, the system imports only the first specified workflow.
//
//   - ImportWorkflowDefinition is an asynchronous operation. After you send a request, an asynchronous task is generated, and the system returns the ID of the asynchronous task. You can call the GetJobStatus operation to query the status of the asynchronous task.
//
// @param request - ImportWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportWorkflowDefinitionResponse
func (client *Client) ImportWorkflowDefinitionWithOptions(request *ImportWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *ImportWorkflowDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a workflow and its child nodes that are specified by the FlowSpec field to DataStudio.
//
// Description:
//
// >
//
//   - You cannot use this API operation to import multiple workflows at a time. If you specify multiple workflows by using FlowSpec, the system imports only the first specified workflow.
//
//   - ImportWorkflowDefinition is an asynchronous operation. After you send a request, an asynchronous task is generated, and the system returns the ID of the asynchronous task. You can call the GetJobStatus operation to query the status of the asynchronous task.
//
// @param request - ImportWorkflowDefinitionRequest
//
// @return ImportWorkflowDefinitionResponse
func (client *Client) ImportWorkflowDefinition(request *ImportWorkflowDefinitionRequest) (_result *ImportWorkflowDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportWorkflowDefinitionResponse{}
	_body, _err := client.ImportWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of custom monitoring alert rules.
//
// @param tmpReq - ListAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertRulesResponse
func (client *Client) ListAlertRulesWithOptions(tmpReq *ListAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAlertRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAlertRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Receiver) {
		query["Receiver"] = request.Receiver
	}

	if !dara.IsNil(request.TaskIdsShrink) {
		query["TaskIds"] = request.TaskIdsShrink
	}

	if !dara.IsNil(request.TypesShrink) {
		query["Types"] = request.TypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of custom monitoring alert rules.
//
// @param request - ListAlertRulesRequest
//
// @return ListAlertRulesResponse
func (client *Client) ListAlertRules(request *ListAlertRulesRequest) (_result *ListAlertRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAlertRulesResponse{}
	_body, _err := client.ListAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBusinessResponse
func (client *Client) ListBusinessWithOptions(request *ListBusinessRequest, runtime *dara.RuntimeOptions) (_result *ListBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListBusinessRequest
//
// @return ListBusinessResponse
func (client *Client) ListBusiness(request *ListBusinessRequest) (_result *ListBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBusinessResponse{}
	_body, _err := client.ListBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据目录列表
//
// @param tmpReq - ListCatalogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCatalogsResponse
func (client *Client) ListCatalogsWithOptions(tmpReq *ListCatalogsRequest, runtime *dara.RuntimeOptions) (_result *ListCatalogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListCatalogsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCatalogs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCatalogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据目录列表
//
// @param request - ListCatalogsRequest
//
// @return ListCatalogsResponse
func (client *Client) ListCatalogs(request *ListCatalogsRequest) (_result *ListCatalogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCatalogsResponse{}
	_body, _err := client.ListCatalogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of certificate files.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks: Tenant Owner, Workspace Administrator, Deploy, Develop, Visitor, Workspace Owner, O\\&M, Model Designer, Security Administrator, Data Analyst, OpenPlatform Administrator, and Data Governance Administrator.
//
// @param request - ListCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertificatesResponse
func (client *Client) ListCertificatesWithOptions(request *ListCertificatesRequest, runtime *dara.RuntimeOptions) (_result *ListCertificatesResponse, _err error) {
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
		Action:      dara.String("ListCertificates"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of certificate files.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks: Tenant Owner, Workspace Administrator, Deploy, Develop, Visitor, Workspace Owner, O\\&M, Model Designer, Security Administrator, Data Analyst, OpenPlatform Administrator, and Data Governance Administrator.
//
// @param request - ListCertificatesRequest
//
// @return ListCertificatesResponse
func (client *Client) ListCertificates(request *ListCertificatesRequest) (_result *ListCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCertificatesResponse{}
	_body, _err := client.ListCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询字段列表
//
// @param request - ListColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListColumnsResponse
func (client *Client) ListColumnsWithOptions(request *ListColumnsRequest, runtime *dara.RuntimeOptions) (_result *ListColumnsResponse, _err error) {
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
		Action:      dara.String("ListColumns"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询字段列表
//
// @param request - ListColumnsRequest
//
// @return ListColumnsResponse
func (client *Client) ListColumns(request *ListColumnsRequest) (_result *ListColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListColumnsResponse{}
	_body, _err := client.ListColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of components.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListComponentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentsResponse
func (client *Client) ListComponentsWithOptions(request *ListComponentsRequest, runtime *dara.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponents"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of components.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListComponentsRequest
//
// @return ListComponentsResponse
func (client *Client) ListComponents(request *ListComponentsRequest) (_result *ListComponentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComponentsResponse{}
	_body, _err := client.ListComponentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询元数据采集器类型列表
//
// @param request - ListCrawlerTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCrawlerTypesResponse
func (client *Client) ListCrawlerTypesWithOptions(runtime *dara.RuntimeOptions) (_result *ListCrawlerTypesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListCrawlerTypes"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCrawlerTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询元数据采集器类型列表
//
// @return ListCrawlerTypesResponse
func (client *Client) ListCrawlerTypes() (_result *ListCrawlerTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCrawlerTypesResponse{}
	_body, _err := client.ListCrawlerTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Views alert rules configured for a synchronization task.
//
// @param request - ListDIAlarmRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIAlarmRulesResponse
func (client *Client) ListDIAlarmRulesWithOptions(request *ListDIAlarmRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDIAlarmRulesResponse, _err error) {
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
		Action:      dara.String("ListDIAlarmRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIAlarmRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Views alert rules configured for a synchronization task.
//
// @param request - ListDIAlarmRulesRequest
//
// @return ListDIAlarmRulesResponse
func (client *Client) ListDIAlarmRules(request *ListDIAlarmRulesRequest) (_result *ListDIAlarmRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDIAlarmRulesResponse{}
	_body, _err := client.ListDIAlarmRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries events for a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDIJobEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobEventsResponse
func (client *Client) ListDIJobEventsWithOptions(request *ListDIJobEventsRequest, runtime *dara.RuntimeOptions) (_result *ListDIJobEventsResponse, _err error) {
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
		Action:      dara.String("ListDIJobEvents"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIJobEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries events for a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDIJobEventsRequest
//
// @return ListDIJobEventsResponse
func (client *Client) ListDIJobEvents(request *ListDIJobEventsRequest) (_result *ListDIJobEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDIJobEventsResponse{}
	_body, _err := client.ListDIJobEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metrics for a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ListDIJobMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobMetricsResponse
func (client *Client) ListDIJobMetricsWithOptions(tmpReq *ListDIJobMetricsRequest, runtime *dara.RuntimeOptions) (_result *ListDIJobMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDIJobMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MetricName) {
		request.MetricNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricName, dara.String("MetricName"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDIJobMetrics"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIJobMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metrics for a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDIJobMetricsRequest
//
// @return ListDIJobMetricsResponse
func (client *Client) ListDIJobMetrics(request *ListDIJobMetricsRequest) (_result *ListDIJobMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDIJobMetricsResponse{}
	_body, _err := client.ListDIJobMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the running information about a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDIJobRunDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobRunDetailsResponse
func (client *Client) ListDIJobRunDetailsWithOptions(request *ListDIJobRunDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListDIJobRunDetailsResponse, _err error) {
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
		Action:      dara.String("ListDIJobRunDetails"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIJobRunDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the running information about a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDIJobRunDetailsRequest
//
// @return ListDIJobRunDetailsResponse
func (client *Client) ListDIJobRunDetails(request *ListDIJobRunDetailsRequest) (_result *ListDIJobRunDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDIJobRunDetailsResponse{}
	_body, _err := client.ListDIJobRunDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of new-version synchronization tasks in Data Integration. A new-version synchronization task can be a real-time synchronization task used to synchronize full or incremental data in a database, a batch synchronization task used to synchronize full or incremental data in a database, or a real-time synchronization task used to synchronize incremental data in a single table.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDIJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobsResponse
func (client *Client) ListDIJobsWithOptions(request *ListDIJobsRequest, runtime *dara.RuntimeOptions) (_result *ListDIJobsResponse, _err error) {
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
		Action:      dara.String("ListDIJobs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of new-version synchronization tasks in Data Integration. A new-version synchronization task can be a real-time synchronization task used to synchronize full or incremental data in a database, a batch synchronization task used to synchronize full or incremental data in a database, or a real-time synchronization task used to synchronize incremental data in a single table.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDIJobsRequest
//
// @return ListDIJobsResponse
func (client *Client) ListDIJobs(request *ListDIJobsRequest) (_result *ListDIJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDIJobsResponse{}
	_body, _err := client.ListDIJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param request - ListDataAssetTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataAssetTagsResponse
func (client *Client) ListDataAssetTagsWithOptions(request *ListDataAssetTagsRequest, runtime *dara.RuntimeOptions) (_result *ListDataAssetTagsResponse, _err error) {
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
		Action:      dara.String("ListDataAssetTags"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataAssetTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tags.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param request - ListDataAssetTagsRequest
//
// @return ListDataAssetTagsResponse
func (client *Client) ListDataAssetTags(request *ListDataAssetTagsRequest) (_result *ListDataAssetTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataAssetTagsResponse{}
	_body, _err := client.ListDataAssetTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about DataWorks data assets to which tags are added by page.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - ListDataAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataAssetsResponse
func (client *Client) ListDataAssetsWithOptions(tmpReq *ListDataAssetsRequest, runtime *dara.RuntimeOptions) (_result *ListDataAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataAssetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataAssetIds) {
		request.DataAssetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataAssetIds, dara.String("DataAssetIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataAssets"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about DataWorks data assets to which tags are added by page.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param request - ListDataAssetsRequest
//
// @return ListDataAssetsResponse
func (client *Client) ListDataAssets(request *ListDataAssetsRequest) (_result *ListDataAssetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataAssetsResponse{}
	_body, _err := client.ListDataAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of data quality alert rules in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityAlertRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityAlertRulesResponse
func (client *Client) ListDataQualityAlertRulesWithOptions(request *ListDataQualityAlertRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityAlertRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityScanId) {
		query["DataQualityScanId"] = request.DataQualityScanId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataQualityAlertRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data quality alert rules in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityAlertRulesRequest
//
// @return ListDataQualityAlertRulesResponse
func (client *Client) ListDataQualityAlertRules(request *ListDataQualityAlertRulesRequest) (_result *ListDataQualityAlertRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataQualityAlertRulesResponse{}
	_body, _err := client.ListDataQualityAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instances generated by a data quality monitoring task by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityEvaluationTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityEvaluationTaskInstancesResponse
func (client *Client) ListDataQualityEvaluationTaskInstancesWithOptions(request *ListDataQualityEvaluationTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityEvaluationTaskInstancesResponse, _err error) {
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
		Action:      dara.String("ListDataQualityEvaluationTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityEvaluationTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances generated by a data quality monitoring task by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityEvaluationTaskInstancesRequest
//
// @return ListDataQualityEvaluationTaskInstancesResponse
func (client *Client) ListDataQualityEvaluationTaskInstances(request *ListDataQualityEvaluationTaskInstancesRequest) (_result *ListDataQualityEvaluationTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataQualityEvaluationTaskInstancesResponse{}
	_body, _err := client.ListDataQualityEvaluationTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data quality monitoring tasks by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityEvaluationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityEvaluationTasksResponse
func (client *Client) ListDataQualityEvaluationTasksWithOptions(request *ListDataQualityEvaluationTasksRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityEvaluationTasksResponse, _err error) {
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
		Action:      dara.String("ListDataQualityEvaluationTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityEvaluationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data quality monitoring tasks by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityEvaluationTasksRequest
//
// @return ListDataQualityEvaluationTasksResponse
func (client *Client) ListDataQualityEvaluationTasks(request *ListDataQualityEvaluationTasksRequest) (_result *ListDataQualityEvaluationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataQualityEvaluationTasksResponse{}
	_body, _err := client.ListDataQualityEvaluationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityResultsResponse
func (client *Client) ListDataQualityResultsWithOptions(request *ListDataQualityResultsRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityResultsResponse, _err error) {
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
		Action:      dara.String("ListDataQualityResults"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityResultsRequest
//
// @return ListDataQualityResultsResponse
func (client *Client) ListDataQualityResults(request *ListDataQualityResultsRequest) (_result *ListDataQualityResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataQualityResultsResponse{}
	_body, _err := client.ListDataQualityResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data quality monitoring rule templates.
//
// @param request - ListDataQualityRuleTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityRuleTemplatesResponse
func (client *Client) ListDataQualityRuleTemplatesWithOptions(request *ListDataQualityRuleTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityRuleTemplatesResponse, _err error) {
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
		Action:      dara.String("ListDataQualityRuleTemplates"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityRuleTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data quality monitoring rule templates.
//
// @param request - ListDataQualityRuleTemplatesRequest
//
// @return ListDataQualityRuleTemplatesResponse
func (client *Client) ListDataQualityRuleTemplates(request *ListDataQualityRuleTemplatesRequest) (_result *ListDataQualityRuleTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataQualityRuleTemplatesResponse{}
	_body, _err := client.ListDataQualityRuleTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data quality monitoring rules by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityRulesResponse
func (client *Client) ListDataQualityRulesWithOptions(request *ListDataQualityRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityRulesResponse, _err error) {
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
		Action:      dara.String("ListDataQualityRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data quality monitoring rules by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityRulesRequest
//
// @return ListDataQualityRulesResponse
func (client *Client) ListDataQualityRules(request *ListDataQualityRulesRequest) (_result *ListDataQualityRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataQualityRulesResponse{}
	_body, _err := client.ListDataQualityRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution records of data quality scans in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityScanRunsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityScanRunsResponse
func (client *Client) ListDataQualityScanRunsWithOptions(request *ListDataQualityScanRunsRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityScanRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeFrom) {
		query["CreateTimeFrom"] = request.CreateTimeFrom
	}

	if !dara.IsNil(request.CreateTimeTo) {
		query["CreateTimeTo"] = request.CreateTimeTo
	}

	if !dara.IsNil(request.DataQualityScanId) {
		query["DataQualityScanId"] = request.DataQualityScanId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataQualityScanRuns"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityScanRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution records of data quality scans in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityScanRunsRequest
//
// @return ListDataQualityScanRunsResponse
func (client *Client) ListDataQualityScanRuns(request *ListDataQualityScanRunsRequest) (_result *ListDataQualityScanRunsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataQualityScanRunsResponse{}
	_body, _err := client.ListDataQualityScanRunsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of data quality scan tasks in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityScansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityScansResponse
func (client *Client) ListDataQualityScansWithOptions(request *ListDataQualityScansRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityScansResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Table) {
		query["Table"] = request.Table
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataQualityScans"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityScansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data quality scan tasks in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityScansRequest
//
// @return ListDataQualityScansResponse
func (client *Client) ListDataQualityScans(request *ListDataQualityScansRequest) (_result *ListDataQualityScansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataQualityScansResponse{}
	_body, _err := client.ListDataQualityScansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of data quality rule templates in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityTemplatesResponse
func (client *Client) ListDataQualityTemplatesWithOptions(request *ListDataQualityTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListDataQualityTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Catalog) {
		query["Catalog"] = request.Catalog
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataQualityTemplates"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataQualityTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data quality rule templates in a project.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDataQualityTemplatesRequest
//
// @return ListDataQualityTemplatesResponse
func (client *Client) ListDataQualityTemplates(request *ListDataQualityTemplatesRequest) (_result *ListDataQualityTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataQualityTemplatesResponse{}
	_body, _err := client.ListDataQualityTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of sharing rules of a data source.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  If you want to query the sharing rules of a data source that is associated with Workspace A, you must have the permissions to share the data source in Workspace A. You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Tenant Administrator, Workspace Administrator, and Workspace Owner
//
// @param request - ListDataSourceSharedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceSharedRulesResponse
func (client *Client) ListDataSourceSharedRulesWithOptions(request *ListDataSourceSharedRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourceSharedRulesResponse, _err error) {
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
		Action:      dara.String("ListDataSourceSharedRules"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourceSharedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of sharing rules of a data source.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  If you want to query the sharing rules of a data source that is associated with Workspace A, you must have the permissions to share the data source in Workspace A. You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Tenant Administrator, Workspace Administrator, and Workspace Owner
//
// @param request - ListDataSourceSharedRulesRequest
//
// @return ListDataSourceSharedRulesResponse
func (client *Client) ListDataSourceSharedRules(request *ListDataSourceSharedRulesRequest) (_result *ListDataSourceSharedRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSourceSharedRulesResponse{}
	_body, _err := client.ListDataSourceSharedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data sources based on the business information of data sources.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Deploy, Develop, Visitor, Workspace Owner, O\\&M, Model Designer, Security Administrator, Data Analyst, OpenPlatform Administrator, and Data Governance Administrator
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
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSources"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// Queries a list of data sources based on the business information of data sources.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Workspace Administrator, Deploy, Develop, Visitor, Workspace Owner, O\\&M, Model Designer, Security Administrator, Data Analyst, OpenPlatform Administrator, and Data Governance Administrator
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
// 查询数据库列表
//
// @param request - ListDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithOptions(request *ListDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
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
		Action:      dara.String("ListDatabases"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据库列表
//
// @param request - ListDatabasesRequest
//
// @return ListDatabasesResponse
func (client *Client) ListDatabases(request *ListDatabasesRequest) (_result *ListDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabasesResponse{}
	_body, _err := client.ListDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of versions of files to be deployed.
//
// @param tmpReq - ListDeploymentPackageFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentPackageFilesResponse
func (client *Client) ListDeploymentPackageFilesWithOptions(tmpReq *ListDeploymentPackageFilesRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentPackageFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDeploymentPackageFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileIds) {
		request.FileIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileIds, dara.String("FileIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		query["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.ChangeType) {
		query["ChangeType"] = request.ChangeType
	}

	if !dara.IsNil(request.CommitFrom) {
		query["CommitFrom"] = request.CommitFrom
	}

	if !dara.IsNil(request.CommitTo) {
		query["CommitTo"] = request.CommitTo
	}

	if !dara.IsNil(request.CommitUserId) {
		query["CommitUserId"] = request.CommitUserId
	}

	if !dara.IsNil(request.FileIdsShrink) {
		query["FileIds"] = request.FileIdsShrink
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SolutionId) {
		query["SolutionId"] = request.SolutionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentPackageFiles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentPackageFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of versions of files to be deployed.
//
// @param request - ListDeploymentPackageFilesRequest
//
// @return ListDeploymentPackageFilesResponse
func (client *Client) ListDeploymentPackageFiles(request *ListDeploymentPackageFilesRequest) (_result *ListDeploymentPackageFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeploymentPackageFilesResponse{}
	_body, _err := client.ListDeploymentPackageFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of deployment packages.
//
// @param request - ListDeploymentPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentPackagesResponse
func (client *Client) ListDeploymentPackagesWithOptions(request *ListDeploymentPackagesRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentPackagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		body["Creator"] = request.Creator
	}

	if !dara.IsNil(request.EndCreateTime) {
		body["EndCreateTime"] = request.EndCreateTime
	}

	if !dara.IsNil(request.EndExecuteTime) {
		body["EndExecuteTime"] = request.EndExecuteTime
	}

	if !dara.IsNil(request.Executor) {
		body["Executor"] = request.Executor
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeploymentPackages"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of deployment packages.
//
// @param request - ListDeploymentPackagesRequest
//
// @return ListDeploymentPackagesResponse
func (client *Client) ListDeploymentPackages(request *ListDeploymentPackagesRequest) (_result *ListDeploymentPackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeploymentPackagesResponse{}
	_body, _err := client.ListDeploymentPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of descendant instances of an instance by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDownstreamTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownstreamTaskInstancesResponse
func (client *Client) ListDownstreamTaskInstancesWithOptions(request *ListDownstreamTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListDownstreamTaskInstancesResponse, _err error) {
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
		Action:      dara.String("ListDownstreamTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDownstreamTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of descendant instances of an instance by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListDownstreamTaskInstancesRequest
//
// @return ListDownstreamTaskInstancesResponse
func (client *Client) ListDownstreamTaskInstances(request *ListDownstreamTaskInstancesRequest) (_result *ListDownstreamTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDownstreamTaskInstancesResponse{}
	_body, _err := client.ListDownstreamTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of descendant tasks of a task by page.
//
// @param request - ListDownstreamTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDownstreamTasksResponse
func (client *Client) ListDownstreamTasksWithOptions(request *ListDownstreamTasksRequest, runtime *dara.RuntimeOptions) (_result *ListDownstreamTasksResponse, _err error) {
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
		Action:      dara.String("ListDownstreamTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDownstreamTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of descendant tasks of a task by page.
//
// @param request - ListDownstreamTasksRequest
//
// @return ListDownstreamTasksResponse
func (client *Client) ListDownstreamTasks(request *ListDownstreamTasksRequest) (_result *ListDownstreamTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDownstreamTasksResponse{}
	_body, _err := client.ListDownstreamTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of entities in a collection in Data Map. Collections include categories and data albums. Entities can only be tables.
//
// @param request - ListEntitiesInMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEntitiesInMetaCollectionResponse
func (client *Client) ListEntitiesInMetaCollectionWithOptions(request *ListEntitiesInMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *ListEntitiesInMetaCollectionResponse, _err error) {
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
		Action:      dara.String("ListEntitiesInMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEntitiesInMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of entities in a collection in Data Map. Collections include categories and data albums. Entities can only be tables.
//
// @param request - ListEntitiesInMetaCollectionRequest
//
// @return ListEntitiesInMetaCollectionResponse
func (client *Client) ListEntitiesInMetaCollection(request *ListEntitiesInMetaCollectionRequest) (_result *ListEntitiesInMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEntitiesInMetaCollectionResponse{}
	_body, _err := client.ListEntitiesInMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListFileVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileVersionsResponse
func (client *Client) ListFileVersionsWithOptions(request *ListFileVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListFileVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFileVersions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFileVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListFileVersionsRequest
//
// @return ListFileVersionsResponse
func (client *Client) ListFileVersions(request *ListFileVersionsRequest) (_result *ListFileVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFileVersionsResponse{}
	_body, _err := client.ListFileVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithOptions(request *ListFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CommitStatus) {
		body["CommitStatus"] = request.CommitStatus
	}

	if !dara.IsNil(request.ExactFileName) {
		body["ExactFileName"] = request.ExactFileName
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileIdIn) {
		body["FileIdIn"] = request.FileIdIn
	}

	if !dara.IsNil(request.FileTypes) {
		body["FileTypes"] = request.FileTypes
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.LastEditUser) {
		body["LastEditUser"] = request.LastEditUser
	}

	if !dara.IsNil(request.NeedAbsoluteFolderPath) {
		body["NeedAbsoluteFolderPath"] = request.NeedAbsoluteFolderPath
	}

	if !dara.IsNil(request.NeedContent) {
		body["NeedContent"] = request.NeedContent
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.UseType) {
		body["UseType"] = request.UseType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFiles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListFilesRequest
//
// @return ListFilesResponse
func (client *Client) ListFiles(request *ListFilesRequest) (_result *ListFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFilesResponse{}
	_body, _err := client.ListFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListFoldersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFoldersResponse
func (client *Client) ListFoldersWithOptions(request *ListFoldersRequest, runtime *dara.RuntimeOptions) (_result *ListFoldersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentFolderPath) {
		body["ParentFolderPath"] = request.ParentFolderPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFolders"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFoldersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListFoldersRequest
//
// @return ListFoldersResponse
func (client *Client) ListFolders(request *ListFoldersRequest) (_result *ListFoldersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFoldersResponse{}
	_body, _err := client.ListFoldersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of user-defined functions (UDFs) in DataStudio. You can also specify filter conditions to query specific UDFs.
//
// @param request - ListFunctionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionsResponse
func (client *Client) ListFunctionsWithOptions(request *ListFunctionsRequest, runtime *dara.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
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
		Action:      dara.String("ListFunctions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of user-defined functions (UDFs) in DataStudio. You can also specify filter conditions to query specific UDFs.
//
// @param request - ListFunctionsRequest
//
// @return ListFunctionsResponse
func (client *Client) ListFunctions(request *ListFunctionsRequest) (_result *ListFunctionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFunctionsResponse{}
	_body, _err := client.ListFunctionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询血缘关系
//
// @param request - ListLineageRelationshipsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLineageRelationshipsResponse
func (client *Client) ListLineageRelationshipsWithOptions(request *ListLineageRelationshipsRequest, runtime *dara.RuntimeOptions) (_result *ListLineageRelationshipsResponse, _err error) {
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
		Action:      dara.String("ListLineageRelationships"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLineageRelationshipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询血缘关系
//
// @param request - ListLineageRelationshipsRequest
//
// @return ListLineageRelationshipsResponse
func (client *Client) ListLineageRelationships(request *ListLineageRelationshipsRequest) (_result *ListLineageRelationshipsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLineageRelationshipsResponse{}
	_body, _err := client.ListLineageRelationshipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor and descendant entities of an entity in Data Map. You can specify whether to return the lineage between the entities.
//
// @param request - ListLineagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLineagesResponse
func (client *Client) ListLineagesWithOptions(request *ListLineagesRequest, runtime *dara.RuntimeOptions) (_result *ListLineagesResponse, _err error) {
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
		Action:      dara.String("ListLineages"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLineagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor and descendant entities of an entity in Data Map. You can specify whether to return the lineage between the entities.
//
// @param request - ListLineagesRequest
//
// @return ListLineagesResponse
func (client *Client) ListLineages(request *ListLineagesRequest) (_result *ListLineagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLineagesResponse{}
	_body, _err := client.ListLineagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of collections in Data Map. Collections include categories and data albums.
//
// @param request - ListMetaCollectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetaCollectionsResponse
func (client *Client) ListMetaCollectionsWithOptions(request *ListMetaCollectionsRequest, runtime *dara.RuntimeOptions) (_result *ListMetaCollectionsResponse, _err error) {
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
		Action:      dara.String("ListMetaCollections"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetaCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of collections in Data Map. Collections include categories and data albums.
//
// @param request - ListMetaCollectionsRequest
//
// @return ListMetaCollectionsResponse
func (client *Client) ListMetaCollections(request *ListMetaCollectionsRequest) (_result *ListMetaCollectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMetaCollectionsResponse{}
	_body, _err := client.ListMetaCollectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of network resources of a serverless resource group.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListNetworksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworksResponse
func (client *Client) ListNetworksWithOptions(request *ListNetworksRequest, runtime *dara.RuntimeOptions) (_result *ListNetworksResponse, _err error) {
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
		Action:      dara.String("ListNetworks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNetworksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of network resources of a serverless resource group.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListNetworksRequest
//
// @return ListNetworksResponse
func (client *Client) ListNetworks(request *ListNetworksRequest) (_result *ListNetworksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNetworksResponse{}
	_body, _err := client.ListNetworksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of descendant nodes of a node in DataStudio.
//
// @param request - ListNodeDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDependenciesResponse
func (client *Client) ListNodeDependenciesWithOptions(request *ListNodeDependenciesRequest, runtime *dara.RuntimeOptions) (_result *ListNodeDependenciesResponse, _err error) {
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
		Action:      dara.String("ListNodeDependencies"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeDependenciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of descendant nodes of a node in DataStudio.
//
// @param request - ListNodeDependenciesRequest
//
// @return ListNodeDependenciesResponse
func (client *Client) ListNodeDependencies(request *ListNodeDependenciesRequest) (_result *ListNodeDependenciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodeDependenciesResponse{}
	_body, _err := client.ListNodeDependenciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of nodes in DataStudio. You can also specify filter conditions to query specific nodes.
//
// @param request - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
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
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of nodes in DataStudio. You can also specify filter conditions to query specific nodes.
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据表的分区列表
//
// @param request - ListPartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPartitionsResponse
func (client *Client) ListPartitionsWithOptions(request *ListPartitionsRequest, runtime *dara.RuntimeOptions) (_result *ListPartitionsResponse, _err error) {
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
		Action:      dara.String("ListPartitions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据表的分区列表
//
// @param request - ListPartitionsRequest
//
// @return ListPartitionsResponse
func (client *Client) ListPartitions(request *ListPartitionsRequest) (_result *ListPartitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPartitionsResponse{}
	_body, _err := client.ListPartitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about deployment objects by deployment process ID.
//
// @param request - ListPipelineRunItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineRunItemsResponse
func (client *Client) ListPipelineRunItemsWithOptions(request *ListPipelineRunItemsRequest, runtime *dara.RuntimeOptions) (_result *ListPipelineRunItemsResponse, _err error) {
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
		Action:      dara.String("ListPipelineRunItems"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineRunItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about deployment objects by deployment process ID.
//
// @param request - ListPipelineRunItemsRequest
//
// @return ListPipelineRunItemsResponse
func (client *Client) ListPipelineRunItems(request *ListPipelineRunItemsRequest) (_result *ListPipelineRunItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPipelineRunItemsResponse{}
	_body, _err := client.ListPipelineRunItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of processes that are used to deploy or undeploy entities in DataStudio. You can also specify filter conditions to query specific processes.
//
// @param request - ListPipelineRunsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineRunsResponse
func (client *Client) ListPipelineRunsWithOptions(request *ListPipelineRunsRequest, runtime *dara.RuntimeOptions) (_result *ListPipelineRunsResponse, _err error) {
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
		Action:      dara.String("ListPipelineRuns"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of processes that are used to deploy or undeploy entities in DataStudio. You can also specify filter conditions to query specific processes.
//
// @param request - ListPipelineRunsRequest
//
// @return ListPipelineRunsResponse
func (client *Client) ListPipelineRuns(request *ListPipelineRunsRequest) (_result *ListPipelineRunsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.ListPipelineRunsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about members in a workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ListProjectMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembersWithOptions(tmpReq *ListProjectMembersRequest, runtime *dara.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("RoleCodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserIds) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, dara.String("UserIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserIdsShrink) {
		body["UserIds"] = request.UserIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectMembers"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about members in a workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListProjectMembersRequest
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembers(request *ListProjectMembersRequest) (_result *ListProjectMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.ListProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about roles in a DataWorks workspace by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ListProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectRolesResponse
func (client *Client) ListProjectRolesWithOptions(tmpReq *ListProjectRolesRequest, runtime *dara.RuntimeOptions) (_result *ListProjectRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Codes) {
		request.CodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Codes, dara.String("Codes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Names) {
		request.NamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Names, dara.String("Names"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CodesShrink) {
		body["Codes"] = request.CodesShrink
	}

	if !dara.IsNil(request.NamesShrink) {
		body["Names"] = request.NamesShrink
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectRoles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about roles in a DataWorks workspace by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListProjectRolesRequest
//
// @return ListProjectRolesResponse
func (client *Client) ListProjectRoles(request *ListProjectRolesRequest) (_result *ListProjectRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectRolesResponse{}
	_body, _err := client.ListProjectRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of DataWorks workspaces of the tenant to which your account belongs.
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(tmpReq *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AliyunResourceTags) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, dara.String("AliyunResourceTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Names) {
		request.NamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Names, dara.String("Names"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunResourceGroupId) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !dara.IsNil(request.AliyunResourceTagsShrink) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !dara.IsNil(request.DevEnvironmentEnabled) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !dara.IsNil(request.DevRoleDisabled) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.NamesShrink) {
		body["Names"] = request.NamesShrink
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PaiTaskEnabled) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of DataWorks workspaces of the tenant to which your account belongs.
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of resource groups.
//
// @param tmpReq - ListResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithOptions(tmpReq *ListResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListResourceGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AliyunResourceTags) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, dara.String("AliyunResourceTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceGroupTypes) {
		request.ResourceGroupTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceGroupTypes, dara.String("ResourceGroupTypes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Statuses) {
		request.StatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Statuses, dara.String("Statuses"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroups"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of resource groups.
//
// @param request - ListResourceGroupsRequest
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (_result *ListResourceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.ListResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of file resources in DataStudio. You can also specify filter conditions to query specific file resources.
//
// @param request - ListResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
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
		Action:      dara.String("ListResources"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of file resources in DataStudio. You can also specify filter conditions to query specific file resources.
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of routes of a network resource.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoutesResponse
func (client *Client) ListRoutesWithOptions(request *ListRoutesRequest, runtime *dara.RuntimeOptions) (_result *ListRoutesResponse, _err error) {
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
		Action:      dara.String("ListRoutes"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of routes of a network resource.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListRoutesRequest
//
// @return ListRoutesResponse
func (client *Client) ListRoutes(request *ListRoutesRequest) (_result *ListRoutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRoutesResponse{}
	_body, _err := client.ListRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of schemas in a database or a MaxCompute project in Data Map. Only schemas of the MaxCompute and Hologres metadata crawler types are supported.
//
// @param tmpReq - ListSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemasResponse
func (client *Client) ListSchemasWithOptions(tmpReq *ListSchemasRequest, runtime *dara.RuntimeOptions) (_result *ListSchemasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSchemasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Types) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, dara.String("Types"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSchemas"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of schemas in a database or a MaxCompute project in Data Map. Only schemas of the MaxCompute and Hologres metadata crawler types are supported.
//
// @param request - ListSchemasRequest
//
// @return ListSchemasResponse
func (client *Client) ListSchemas(request *ListSchemasRequest) (_result *ListSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSchemasResponse{}
	_body, _err := client.ListSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tables in Data Map. For data source types that do not support schemas, you can call this API operation to query a list of tables in a specific database. For data source types that support schemas, you can call this API operation to query a list of tables in a specific database, MaxCompute project, or schema. Only the basic information about tables is returned. The information about technical metadata and business metadata is not returned.
//
// @param tmpReq - ListTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithOptions(tmpReq *ListTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTablesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableTypes) {
		request.TableTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableTypes, dara.String("TableTypes"), dara.String("simple"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tables in Data Map. For data source types that do not support schemas, you can call this API operation to query a list of tables in a specific database. For data source types that support schemas, you can call this API operation to query a list of tables in a specific database, MaxCompute project, or schema. Only the basic information about tables is returned. The information about technical metadata and business metadata is not returned.
//
// @param request - ListTablesRequest
//
// @return ListTablesResponse
func (client *Client) ListTables(request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of operation logs of an instance by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// You can call this operation to query only the operation logs generated within the previous 31 days.
//
// @param request - ListTaskInstanceOperationLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskInstanceOperationLogsResponse
func (client *Client) ListTaskInstanceOperationLogsWithOptions(request *ListTaskInstanceOperationLogsRequest, runtime *dara.RuntimeOptions) (_result *ListTaskInstanceOperationLogsResponse, _err error) {
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
		Action:      dara.String("ListTaskInstanceOperationLogs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskInstanceOperationLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of operation logs of an instance by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// You can call this operation to query only the operation logs generated within the previous 31 days.
//
// @param request - ListTaskInstanceOperationLogsRequest
//
// @return ListTaskInstanceOperationLogsResponse
func (client *Client) ListTaskInstanceOperationLogs(request *ListTaskInstanceOperationLogsRequest) (_result *ListTaskInstanceOperationLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskInstanceOperationLogsResponse{}
	_body, _err := client.ListTaskInstanceOperationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instances. You can also specify filter conditions to query specific instances.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ListTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskInstancesResponse
func (client *Client) ListTaskInstancesWithOptions(tmpReq *ListTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("TaskIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResource) {
		body["RuntimeResource"] = request.RuntimeResource
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskIdsShrink) {
		body["TaskIds"] = request.TaskIdsShrink
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TriggerRecurrence) {
		body["TriggerRecurrence"] = request.TriggerRecurrence
	}

	if !dara.IsNil(request.TriggerType) {
		body["TriggerType"] = request.TriggerType
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkflowInstanceId) {
		body["WorkflowInstanceId"] = request.WorkflowInstanceId
	}

	if !dara.IsNil(request.WorkflowInstanceType) {
		body["WorkflowInstanceType"] = request.WorkflowInstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances. You can also specify filter conditions to query specific instances.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListTaskInstancesRequest
//
// @return ListTaskInstancesResponse
func (client *Client) ListTaskInstances(request *ListTaskInstancesRequest) (_result *ListTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskInstancesResponse{}
	_body, _err := client.ListTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of operation logs of a task by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// You can call this operation to query only the operation logs generated within the previous 31 days.
//
// @param request - ListTaskOperationLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskOperationLogsResponse
func (client *Client) ListTaskOperationLogsWithOptions(request *ListTaskOperationLogsRequest, runtime *dara.RuntimeOptions) (_result *ListTaskOperationLogsResponse, _err error) {
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
		Action:      dara.String("ListTaskOperationLogs"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskOperationLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of operation logs of a task by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// You can call this operation to query only the operation logs generated within the previous 31 days.
//
// @param request - ListTaskOperationLogsRequest
//
// @return ListTaskOperationLogsResponse
func (client *Client) ListTaskOperationLogs(request *ListTaskOperationLogsRequest) (_result *ListTaskOperationLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskOperationLogsResponse{}
	_body, _err := client.ListTaskOperationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tasks by page. You can also specify filter conditions to query tasks.
//
// @param tmpReq - ListTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithOptions(tmpReq *ListTasksRequest, runtime *dara.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResource) {
		body["RuntimeResource"] = request.RuntimeResource
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TriggerRecurrence) {
		body["TriggerRecurrence"] = request.TriggerRecurrence
	}

	if !dara.IsNil(request.TriggerType) {
		body["TriggerType"] = request.TriggerType
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tasks by page. You can also specify filter conditions to query tasks.
//
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor instances of an instance by page.
//
// @param request - ListUpstreamTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUpstreamTaskInstancesResponse
func (client *Client) ListUpstreamTaskInstancesWithOptions(request *ListUpstreamTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListUpstreamTaskInstancesResponse, _err error) {
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
		Action:      dara.String("ListUpstreamTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUpstreamTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor instances of an instance by page.
//
// @param request - ListUpstreamTaskInstancesRequest
//
// @return ListUpstreamTaskInstancesResponse
func (client *Client) ListUpstreamTaskInstances(request *ListUpstreamTaskInstancesRequest) (_result *ListUpstreamTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUpstreamTaskInstancesResponse{}
	_body, _err := client.ListUpstreamTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor tasks of a task by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListUpstreamTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUpstreamTasksResponse
func (client *Client) ListUpstreamTasksWithOptions(request *ListUpstreamTasksRequest, runtime *dara.RuntimeOptions) (_result *ListUpstreamTasksResponse, _err error) {
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
		Action:      dara.String("ListUpstreamTasks"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUpstreamTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor tasks of a task by page.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListUpstreamTasksRequest
//
// @return ListUpstreamTasksResponse
func (client *Client) ListUpstreamTasks(request *ListUpstreamTasksRequest) (_result *ListUpstreamTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUpstreamTasksResponse{}
	_body, _err := client.ListUpstreamTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of workflows in DataStudio. You can also specify filter conditions to query specific workflows.
//
// @param request - ListWorkflowDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowDefinitionsResponse
func (client *Client) ListWorkflowDefinitionsWithOptions(request *ListWorkflowDefinitionsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowDefinitionsResponse, _err error) {
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
		Action:      dara.String("ListWorkflowDefinitions"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowDefinitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workflows in DataStudio. You can also specify filter conditions to query specific workflows.
//
// @param request - ListWorkflowDefinitionsRequest
//
// @return ListWorkflowDefinitionsResponse
func (client *Client) ListWorkflowDefinitions(request *ListWorkflowDefinitionsRequest) (_result *ListWorkflowDefinitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkflowDefinitionsResponse{}
	_body, _err := client.ListWorkflowDefinitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of workflow instances by page. You can also specify filter conditions to query workflow instances.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ListWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowInstancesResponse
func (client *Client) ListWorkflowInstancesWithOptions(tmpReq *ListWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		body["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workflow instances by page. You can also specify filter conditions to query workflow instances.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListWorkflowInstancesRequest
//
// @return ListWorkflowInstancesResponse
func (client *Client) ListWorkflowInstances(request *ListWorkflowInstancesRequest) (_result *ListWorkflowInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkflowInstancesResponse{}
	_body, _err := client.ListWorkflowInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of workflows by page. You can also specify filter conditions to query workflows.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ListWorkflowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowsResponse
func (client *Client) ListWorkflowsWithOptions(tmpReq *ListWorkflowsRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListWorkflowsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SortBy) {
		body["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.TriggerType) {
		body["TriggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflows"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workflows by page. You can also specify filter conditions to query workflows.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ListWorkflowsRequest
//
// @return ListWorkflowsResponse
func (client *Client) ListWorkflows(request *ListWorkflowsRequest) (_result *ListWorkflowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkflowsResponse{}
	_body, _err := client.ListWorkflowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a user-defined function (UDF) to a path in DataStudio.
//
// @param request - MoveFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveFunctionResponse
func (client *Client) MoveFunctionWithOptions(request *MoveFunctionRequest, runtime *dara.RuntimeOptions) (_result *MoveFunctionResponse, _err error) {
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

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a user-defined function (UDF) to a path in DataStudio.
//
// @param request - MoveFunctionRequest
//
// @return MoveFunctionResponse
func (client *Client) MoveFunction(request *MoveFunctionRequest) (_result *MoveFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveFunctionResponse{}
	_body, _err := client.MoveFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a node to a path in DataStudio.
//
// @param request - MoveNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveNodeResponse
func (client *Client) MoveNodeWithOptions(request *MoveNodeRequest, runtime *dara.RuntimeOptions) (_result *MoveNodeResponse, _err error) {
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

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a node to a path in DataStudio.
//
// @param request - MoveNodeRequest
//
// @return MoveNodeResponse
func (client *Client) MoveNode(request *MoveNodeRequest) (_result *MoveNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveNodeResponse{}
	_body, _err := client.MoveNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a file resource to a path in DataStudio.
//
// @param request - MoveResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceResponse
func (client *Client) MoveResourceWithOptions(request *MoveResourceRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceResponse, _err error) {
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

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a file resource to a path in DataStudio.
//
// @param request - MoveResourceRequest
//
// @return MoveResourceResponse
func (client *Client) MoveResource(request *MoveResourceRequest) (_result *MoveResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceResponse{}
	_body, _err := client.MoveResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a workflow to a path in DataStudio.
//
// @param request - MoveWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveWorkflowDefinitionResponse
func (client *Client) MoveWorkflowDefinitionWithOptions(request *MoveWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *MoveWorkflowDefinitionResponse, _err error) {
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

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a workflow to a path in DataStudio.
//
// @param request - MoveWorkflowDefinitionRequest
//
// @return MoveWorkflowDefinitionResponse
func (client *Client) MoveWorkflowDefinition(request *MoveWorkflowDefinitionRequest) (_result *MoveWorkflowDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveWorkflowDefinitionResponse{}
	_body, _err := client.MoveWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从集合中移除实体对象
//
// @param request - RemoveEntityFromMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveEntityFromMetaCollectionResponse
func (client *Client) RemoveEntityFromMetaCollectionWithOptions(request *RemoveEntityFromMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *RemoveEntityFromMetaCollectionResponse, _err error) {
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

	if !dara.IsNil(request.MetaCollectionId) {
		query["MetaCollectionId"] = request.MetaCollectionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveEntityFromMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveEntityFromMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从集合中移除实体对象
//
// @param request - RemoveEntityFromMetaCollectionRequest
//
// @return RemoveEntityFromMetaCollectionResponse
func (client *Client) RemoveEntityFromMetaCollection(request *RemoveEntityFromMetaCollectionRequest) (_result *RemoveEntityFromMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveEntityFromMetaCollectionResponse{}
	_body, _err := client.RemoveEntityFromMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes multiple upstream dependencies of an instance at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - RemoveTaskInstanceDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTaskInstanceDependenciesResponse
func (client *Client) RemoveTaskInstanceDependenciesWithOptions(tmpReq *RemoveTaskInstanceDependenciesRequest, runtime *dara.RuntimeOptions) (_result *RemoveTaskInstanceDependenciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RemoveTaskInstanceDependenciesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UpstreamTaskInstanceIds) {
		request.UpstreamTaskInstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UpstreamTaskInstanceIds, dara.String("UpstreamTaskInstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.UpstreamTaskInstanceIdsShrink) {
		body["UpstreamTaskInstanceIds"] = request.UpstreamTaskInstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTaskInstanceDependencies"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTaskInstanceDependenciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes multiple upstream dependencies of an instance at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - RemoveTaskInstanceDependenciesRequest
//
// @return RemoveTaskInstanceDependenciesResponse
func (client *Client) RemoveTaskInstanceDependencies(request *RemoveTaskInstanceDependenciesRequest) (_result *RemoveTaskInstanceDependenciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTaskInstanceDependenciesResponse{}
	_body, _err := client.RemoveTaskInstanceDependenciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a user-defined function (UDF) in DataStudio.
//
// @param request - RenameFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameFunctionResponse
func (client *Client) RenameFunctionWithOptions(request *RenameFunctionRequest, runtime *dara.RuntimeOptions) (_result *RenameFunctionResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a user-defined function (UDF) in DataStudio.
//
// @param request - RenameFunctionRequest
//
// @return RenameFunctionResponse
func (client *Client) RenameFunction(request *RenameFunctionRequest) (_result *RenameFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenameFunctionResponse{}
	_body, _err := client.RenameFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a node in DataStudio.
//
// @param request - RenameNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameNodeResponse
func (client *Client) RenameNodeWithOptions(request *RenameNodeRequest, runtime *dara.RuntimeOptions) (_result *RenameNodeResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a node in DataStudio.
//
// @param request - RenameNodeRequest
//
// @return RenameNodeResponse
func (client *Client) RenameNode(request *RenameNodeRequest) (_result *RenameNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenameNodeResponse{}
	_body, _err := client.RenameNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a file resource in DataStudio.
//
// @param request - RenameResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameResourceResponse
func (client *Client) RenameResourceWithOptions(request *RenameResourceRequest, runtime *dara.RuntimeOptions) (_result *RenameResourceResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a file resource in DataStudio.
//
// @param request - RenameResourceRequest
//
// @return RenameResourceResponse
func (client *Client) RenameResource(request *RenameResourceRequest) (_result *RenameResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenameResourceResponse{}
	_body, _err := client.RenameResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a workflow in DataStudio.
//
// @param request - RenameWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameWorkflowDefinitionResponse
func (client *Client) RenameWorkflowDefinitionWithOptions(request *RenameWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *RenameWorkflowDefinitionResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenameWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenameWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a workflow in DataStudio.
//
// @param request - RenameWorkflowDefinitionRequest
//
// @return RenameWorkflowDefinitionResponse
func (client *Client) RenameWorkflowDefinition(request *RenameWorkflowDefinitionRequest) (_result *RenameWorkflowDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenameWorkflowDefinitionResponse{}
	_body, _err := client.RenameWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns multiple instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - RerunTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RerunTaskInstancesResponse
func (client *Client) RerunTaskInstancesWithOptions(tmpReq *RerunTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *RerunTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RerunTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RerunTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RerunTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns multiple instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - RerunTaskInstancesRequest
//
// @return RerunTaskInstancesResponse
func (client *Client) RerunTaskInstances(request *RerunTaskInstancesRequest) (_result *RerunTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RerunTaskInstancesResponse{}
	_body, _err := client.RerunTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns workflow instances.
//
// @param tmpReq - RerunWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RerunWorkflowInstancesResponse
func (client *Client) RerunWorkflowInstancesWithOptions(tmpReq *RerunWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *RerunWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RerunWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.EndTriggerTime) {
		body["EndTriggerTime"] = request.EndTriggerTime
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.FilterShrink) {
		body["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StartTriggerTime) {
		body["StartTriggerTime"] = request.StartTriggerTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RerunWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RerunWorkflowInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns workflow instances.
//
// @param request - RerunWorkflowInstancesRequest
//
// @return RerunWorkflowInstancesResponse
func (client *Client) RerunWorkflowInstances(request *RerunWorkflowInstancesRequest) (_result *RerunWorkflowInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RerunWorkflowInstancesResponse{}
	_body, _err := client.RerunWorkflowInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes multiple suspended instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - ResumeTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeTaskInstancesResponse
func (client *Client) ResumeTaskInstancesWithOptions(tmpReq *ResumeTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *ResumeTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ResumeTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes multiple suspended instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - ResumeTaskInstancesRequest
//
// @return ResumeTaskInstancesResponse
func (client *Client) ResumeTaskInstances(request *ResumeTaskInstancesRequest) (_result *ResumeTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeTaskInstancesResponse{}
	_body, _err := client.ResumeTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes roles that are assigned to a member in a workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - RevokeMemberProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeMemberProjectRolesResponse
func (client *Client) RevokeMemberProjectRolesWithOptions(tmpReq *RevokeMemberProjectRolesRequest, runtime *dara.RuntimeOptions) (_result *RevokeMemberProjectRolesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RevokeMemberProjectRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RoleCodes) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, dara.String("RoleCodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCodesShrink) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeMemberProjectRoles"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeMemberProjectRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes roles that are assigned to a member in a workspace.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - RevokeMemberProjectRolesRequest
//
// @return RevokeMemberProjectRolesResponse
func (client *Client) RevokeMemberProjectRoles(request *RevokeMemberProjectRolesRequest) (_result *RevokeMemberProjectRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeMemberProjectRolesResponse{}
	_body, _err := client.RevokeMemberProjectRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the statuses of multiple instances to successful at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - SetSuccessTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSuccessTaskInstancesResponse
func (client *Client) SetSuccessTaskInstancesWithOptions(tmpReq *SetSuccessTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *SetSuccessTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SetSuccessTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSuccessTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSuccessTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the statuses of multiple instances to successful at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - SetSuccessTaskInstancesRequest
//
// @return SetSuccessTaskInstancesResponse
func (client *Client) SetSuccessTaskInstances(request *SetSuccessTaskInstancesRequest) (_result *SetSuccessTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSuccessTaskInstancesResponse{}
	_body, _err := client.SetSuccessTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a new-version synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - StartDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDIJobResponse
func (client *Client) StartDIJobWithOptions(tmpReq *StartDIJobRequest, runtime *dara.RuntimeOptions) (_result *StartDIJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RealtimeStartSettings) {
		request.RealtimeStartSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RealtimeStartSettings, dara.String("RealtimeStartSettings"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a new-version synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - StartDIJobRequest
//
// @return StartDIJobResponse
func (client *Client) StartDIJob(request *StartDIJobRequest) (_result *StartDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDIJobResponse{}
	_body, _err := client.StartDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts multiple workflow instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - StartWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartWorkflowInstancesResponse
func (client *Client) StartWorkflowInstancesWithOptions(tmpReq *StartWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *StartWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartWorkflowInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts multiple workflow instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - StartWorkflowInstancesRequest
//
// @return StartWorkflowInstancesResponse
func (client *Client) StartWorkflowInstances(request *StartWorkflowInstancesRequest) (_result *StartWorkflowInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartWorkflowInstancesResponse{}
	_body, _err := client.StartWorkflowInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - StopDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDIJobResponse
func (client *Client) StopDIJobWithOptions(request *StopDIJobRequest, runtime *dara.RuntimeOptions) (_result *StopDIJobResponse, _err error) {
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
		Action:      dara.String("StopDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - StopDIJobRequest
//
// @return StopDIJobResponse
func (client *Client) StopDIJob(request *StopDIJobRequest) (_result *StopDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDIJobResponse{}
	_body, _err := client.StopDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops multiple instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - StopTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTaskInstancesResponse
func (client *Client) StopTaskInstancesWithOptions(tmpReq *StopTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *StopTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops multiple instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - StopTaskInstancesRequest
//
// @return StopTaskInstancesResponse
func (client *Client) StopTaskInstances(request *StopTaskInstancesRequest) (_result *StopTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopTaskInstancesResponse{}
	_body, _err := client.StopTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops multiple workflow instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - StopWorkflowInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopWorkflowInstancesResponse
func (client *Client) StopWorkflowInstancesWithOptions(tmpReq *StopWorkflowInstancesRequest, runtime *dara.RuntimeOptions) (_result *StopWorkflowInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopWorkflowInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopWorkflowInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopWorkflowInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops multiple workflow instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - StopWorkflowInstancesRequest
//
// @return StopWorkflowInstancesResponse
func (client *Client) StopWorkflowInstances(request *StopWorkflowInstancesRequest) (_result *StopWorkflowInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopWorkflowInstancesResponse{}
	_body, _err := client.StopWorkflowInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitFileResponse
func (client *Client) SubmitFileWithOptions(request *SubmitFileRequest, runtime *dara.RuntimeOptions) (_result *SubmitFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.SkipAllDeployFileExtensions) {
		body["SkipAllDeployFileExtensions"] = request.SkipAllDeployFileExtensions
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitFileRequest
//
// @return SubmitFileResponse
func (client *Client) SubmitFile(request *SubmitFileRequest) (_result *SubmitFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitFileResponse{}
	_body, _err := client.SubmitFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends multiple instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - SuspendTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendTaskInstancesResponse
func (client *Client) SuspendTaskInstancesWithOptions(tmpReq *SuspendTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *SuspendTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SuspendTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends multiple instances at a time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - SuspendTaskInstancesRequest
//
// @return SuspendTaskInstancesResponse
func (client *Client) SuspendTaskInstances(request *SuspendTaskInstancesRequest) (_result *SuspendTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendTaskInstancesResponse{}
	_body, _err := client.SuspendTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to data assets.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - TagDataAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagDataAssetsResponse
func (client *Client) TagDataAssetsWithOptions(tmpReq *TagDataAssetsRequest, runtime *dara.RuntimeOptions) (_result *TagDataAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TagDataAssetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataAssetIds) {
		request.DataAssetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataAssetIds, dara.String("DataAssetIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoTraceEnabled) {
		query["AutoTraceEnabled"] = request.AutoTraceEnabled
	}

	if !dara.IsNil(request.DataAssetIdsShrink) {
		query["DataAssetIds"] = request.DataAssetIdsShrink
	}

	if !dara.IsNil(request.DataAssetType) {
		query["DataAssetType"] = request.DataAssetType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagDataAssets"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagDataAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to data assets.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param request - TagDataAssetsRequest
//
// @return TagDataAssetsResponse
func (client *Client) TagDataAssets(request *TagDataAssetsRequest) (_result *TagDataAssetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagDataAssetsResponse{}
	_body, _err := client.TagDataAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tests the connectivity between a data source and a resource group.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  Your account must be assigned one of the following roles of the desired workspace: Tenant Owner, Workspace Administrator, Deploy, Develop, Workspace Owner, and O\\&M
//
// @param request - TestDataSourceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestDataSourceConnectivityResponse
func (client *Client) TestDataSourceConnectivityWithOptions(request *TestDataSourceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *TestDataSourceConnectivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestDataSourceConnectivity"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestDataSourceConnectivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the connectivity between a data source and a resource group.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  Your account must be assigned one of the following roles of the desired workspace: Tenant Owner, Workspace Administrator, Deploy, Develop, Workspace Owner, and O\\&M
//
// @param request - TestDataSourceConnectivityRequest
//
// @return TestDataSourceConnectivityResponse
func (client *Client) TestDataSourceConnectivity(request *TestDataSourceConnectivityRequest) (_result *TestDataSourceConnectivityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestDataSourceConnectivityResponse{}
	_body, _err := client.TestDataSourceConnectivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers a task to run by using an HTTP Trigger node at a specified time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - TriggerSchedulerTaskInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerSchedulerTaskInstanceResponse
func (client *Client) TriggerSchedulerTaskInstanceWithOptions(request *TriggerSchedulerTaskInstanceRequest, runtime *dara.RuntimeOptions) (_result *TriggerSchedulerTaskInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TriggerTime) {
		body["TriggerTime"] = request.TriggerTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TriggerSchedulerTaskInstance"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TriggerSchedulerTaskInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a task to run by using an HTTP Trigger node at a specified time.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - TriggerSchedulerTaskInstanceRequest
//
// @return TriggerSchedulerTaskInstanceResponse
func (client *Client) TriggerSchedulerTaskInstance(request *TriggerSchedulerTaskInstanceRequest) (_result *TriggerSchedulerTaskInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TriggerSchedulerTaskInstanceResponse{}
	_body, _err := client.TriggerSchedulerTaskInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from data assets.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - UnTagDataAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagDataAssetsResponse
func (client *Client) UnTagDataAssetsWithOptions(tmpReq *UnTagDataAssetsRequest, runtime *dara.RuntimeOptions) (_result *UnTagDataAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UnTagDataAssetsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataAssetIds) {
		request.DataAssetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataAssetIds, dara.String("DataAssetIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataAssetIdsShrink) {
		query["DataAssetIds"] = request.DataAssetIdsShrink
	}

	if !dara.IsNil(request.DataAssetType) {
		query["DataAssetType"] = request.DataAssetType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnTagDataAssets"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnTagDataAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from data assets.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param request - UnTagDataAssetsRequest
//
// @return UnTagDataAssetsResponse
func (client *Client) UnTagDataAssets(request *UnTagDataAssetsRequest) (_result *UnTagDataAssetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnTagDataAssetsResponse{}
	_body, _err := client.UnTagDataAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a custom alert monitoring rule.
//
// @param tmpReq - UpdateAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAlertRuleResponse
func (client *Client) UpdateAlertRuleWithOptions(tmpReq *UpdateAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateAlertRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerCondition) {
		request.TriggerConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerCondition, dara.String("TriggerCondition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationShrink) {
		query["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.TriggerConditionShrink) {
		query["TriggerCondition"] = request.TriggerConditionShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a custom alert monitoring rule.
//
// @param request - UpdateAlertRuleRequest
//
// @return UpdateAlertRuleResponse
func (client *Client) UpdateAlertRule(request *UpdateAlertRuleRequest) (_result *UpdateAlertRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAlertRuleResponse{}
	_body, _err := client.UpdateAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a workflow.
//
// @param request - UpdateBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBusinessResponse
func (client *Client) UpdateBusinessWithOptions(request *UpdateBusinessRequest, runtime *dara.RuntimeOptions) (_result *UpdateBusinessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBusiness"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a workflow.
//
// @param request - UpdateBusinessRequest
//
// @return UpdateBusinessResponse
func (client *Client) UpdateBusiness(request *UpdateBusinessRequest) (_result *UpdateBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBusinessResponse{}
	_body, _err := client.UpdateBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新字段业务元数据
//
// @param request - UpdateColumnBusinessMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateColumnBusinessMetadataResponse
func (client *Client) UpdateColumnBusinessMetadataWithOptions(request *UpdateColumnBusinessMetadataRequest, runtime *dara.RuntimeOptions) (_result *UpdateColumnBusinessMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateColumnBusinessMetadata"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateColumnBusinessMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新字段业务元数据
//
// @param request - UpdateColumnBusinessMetadataRequest
//
// @return UpdateColumnBusinessMetadataResponse
func (client *Client) UpdateColumnBusinessMetadata(request *UpdateColumnBusinessMetadataRequest) (_result *UpdateColumnBusinessMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateColumnBusinessMetadataResponse{}
	_body, _err := client.UpdateColumnBusinessMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates components.
//
// Description:
//
// This operation is currently in beta. To join the beta testing, please submit a request. You can call this operation after we add you to the beta program.
//
// @param request - UpdateComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateComponentResponse
func (client *Client) UpdateComponentWithOptions(request *UpdateComponentRequest, runtime *dara.RuntimeOptions) (_result *UpdateComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentId) {
		body["ComponentId"] = request.ComponentId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateComponent"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateComponentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates components.
//
// Description:
//
// This operation is currently in beta. To join the beta testing, please submit a request. You can call this operation after we add you to the beta program.
//
// @param request - UpdateComponentRequest
//
// @return UpdateComponentResponse
func (client *Client) UpdateComponent(request *UpdateComponentRequest) (_result *UpdateComponentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateComponentResponse{}
	_body, _err := client.UpdateComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an alert rule configured for a synchronization task.
//
// @param tmpReq - UpdateDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIAlarmRuleResponse
func (client *Client) UpdateDIAlarmRuleWithOptions(tmpReq *UpdateDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDIAlarmRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDIAlarmRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotificationSettings) {
		request.NotificationSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationSettings, dara.String("NotificationSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerConditions) {
		request.TriggerConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConditions, dara.String("TriggerConditions"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDIAlarmRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an alert rule configured for a synchronization task.
//
// @param request - UpdateDIAlarmRuleRequest
//
// @return UpdateDIAlarmRuleResponse
func (client *Client) UpdateDIAlarmRule(request *UpdateDIAlarmRuleRequest) (_result *UpdateDIAlarmRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDIAlarmRuleResponse{}
	_body, _err := client.UpdateDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - UpdateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIJobResponse
func (client *Client) UpdateDIJobWithOptions(tmpReq *UpdateDIJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateDIJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobSettings) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, dara.String("JobSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSettings) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, dara.String("ResourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TableMappings) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, dara.String("TableMappings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TransformationRules) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, dara.String("TransformationRules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DIJobId) {
		query["DIJobId"] = request.DIJobId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.JobSettingsShrink) {
		body["JobSettings"] = request.JobSettingsShrink
	}

	if !dara.IsNil(request.ResourceSettingsShrink) {
		body["ResourceSettings"] = request.ResourceSettingsShrink
	}

	if !dara.IsNil(request.TableMappingsShrink) {
		body["TableMappings"] = request.TableMappingsShrink
	}

	if !dara.IsNil(request.TransformationRulesShrink) {
		body["TransformationRules"] = request.TransformationRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDIJob"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a synchronization task.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateDIJobRequest
//
// @return UpdateDIJobResponse
func (client *Client) UpdateDIJob(request *UpdateDIJobRequest) (_result *UpdateDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDIJobResponse{}
	_body, _err := client.UpdateDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a tag.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param tmpReq - UpdateDataAssetTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataAssetTagResponse
func (client *Client) UpdateDataAssetTagWithOptions(tmpReq *UpdateDataAssetTagRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataAssetTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataAssetTagShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Managers) {
		request.ManagersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Managers, dara.String("Managers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Values) {
		request.ValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Values, dara.String("Values"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.ManagersShrink) {
		query["Managers"] = request.ManagersShrink
	}

	if !dara.IsNil(request.ValuesShrink) {
		query["Values"] = request.ValuesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataAssetTag"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataAssetTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a tag.
//
// Description:
//
// This API operation is available only for DataWorks Enterprise Edition or a more advanced edition.
//
// @param request - UpdateDataAssetTagRequest
//
// @return UpdateDataAssetTagResponse
func (client *Client) UpdateDataAssetTag(request *UpdateDataAssetTagRequest) (_result *UpdateDataAssetTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataAssetTagResponse{}
	_body, _err := client.UpdateDataAssetTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a specified data quality monitoring alert rule.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - UpdateDataQualityAlertRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityAlertRuleResponse
func (client *Client) UpdateDataQualityAlertRuleWithOptions(tmpReq *UpdateDataQualityAlertRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityAlertRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityAlertRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Notification) {
		request.NotificationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notification, dara.String("Notification"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Condition) {
		body["Condition"] = request.Condition
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.NotificationShrink) {
		body["Notification"] = request.NotificationShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityAlertRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified data quality monitoring alert rule.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateDataQualityAlertRuleRequest
//
// @return UpdateDataQualityAlertRuleResponse
func (client *Client) UpdateDataQualityAlertRule(request *UpdateDataQualityAlertRuleRequest) (_result *UpdateDataQualityAlertRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataQualityAlertRuleResponse{}
	_body, _err := client.UpdateDataQualityAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a monitor.
//
// Description:
//
// This API operation is supported in all DataWorks editions.
//
// @param tmpReq - UpdateDataQualityEvaluationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityEvaluationTaskResponse
func (client *Client) UpdateDataQualityEvaluationTaskWithOptions(tmpReq *UpdateDataQualityEvaluationTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityEvaluationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityEvaluationTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataQualityRules) {
		request.DataQualityRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataQualityRules, dara.String("DataQualityRules"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hooks) {
		request.HooksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hooks, dara.String("Hooks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Notifications) {
		request.NotificationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Notifications, dara.String("Notifications"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Target) {
		request.TargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Target, dara.String("Target"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataQualityRulesShrink) {
		body["DataQualityRules"] = request.DataQualityRulesShrink
	}

	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HooksShrink) {
		body["Hooks"] = request.HooksShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NotificationsShrink) {
		body["Notifications"] = request.NotificationsShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeConf) {
		body["RuntimeConf"] = request.RuntimeConf
	}

	if !dara.IsNil(request.TargetShrink) {
		body["Target"] = request.TargetShrink
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityEvaluationTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a monitor.
//
// Description:
//
// This API operation is supported in all DataWorks editions.
//
// @param request - UpdateDataQualityEvaluationTaskRequest
//
// @return UpdateDataQualityEvaluationTaskResponse
func (client *Client) UpdateDataQualityEvaluationTask(request *UpdateDataQualityEvaluationTaskRequest) (_result *UpdateDataQualityEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataQualityEvaluationTaskResponse{}
	_body, _err := client.UpdateDataQualityEvaluationTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data quality monitoring rule.
//
// @param tmpReq - UpdateDataQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityRuleResponse
func (client *Client) UpdateDataQualityRuleWithOptions(tmpReq *UpdateDataQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckingConfig) {
		request.CheckingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckingConfig, dara.String("CheckingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ErrorHandlers) {
		request.ErrorHandlersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ErrorHandlers, dara.String("ErrorHandlers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SamplingConfig) {
		request.SamplingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SamplingConfig, dara.String("SamplingConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckingConfigShrink) {
		body["CheckingConfig"] = request.CheckingConfigShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ErrorHandlersShrink) {
		body["ErrorHandlers"] = request.ErrorHandlersShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SamplingConfigShrink) {
		body["SamplingConfig"] = request.SamplingConfigShrink
	}

	if !dara.IsNil(request.Severity) {
		body["Severity"] = request.Severity
	}

	if !dara.IsNil(request.TemplateCode) {
		body["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityRule"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data quality monitoring rule.
//
// @param request - UpdateDataQualityRuleRequest
//
// @return UpdateDataQualityRuleResponse
func (client *Client) UpdateDataQualityRule(request *UpdateDataQualityRuleRequest) (_result *UpdateDataQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataQualityRuleResponse{}
	_body, _err := client.UpdateDataQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data quality monitoring rule template.
//
// @param tmpReq - UpdateDataQualityRuleTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityRuleTemplateResponse
func (client *Client) UpdateDataQualityRuleTemplateWithOptions(tmpReq *UpdateDataQualityRuleTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityRuleTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityRuleTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CheckingConfig) {
		request.CheckingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CheckingConfig, dara.String("CheckingConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SamplingConfig) {
		request.SamplingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SamplingConfig, dara.String("SamplingConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckingConfigShrink) {
		body["CheckingConfig"] = request.CheckingConfigShrink
	}

	if !dara.IsNil(request.Code) {
		body["Code"] = request.Code
	}

	if !dara.IsNil(request.DirectoryPath) {
		body["DirectoryPath"] = request.DirectoryPath
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SamplingConfigShrink) {
		body["SamplingConfig"] = request.SamplingConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityRuleTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityRuleTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data quality monitoring rule template.
//
// @param request - UpdateDataQualityRuleTemplateRequest
//
// @return UpdateDataQualityRuleTemplateResponse
func (client *Client) UpdateDataQualityRuleTemplate(request *UpdateDataQualityRuleTemplateRequest) (_result *UpdateDataQualityRuleTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataQualityRuleTemplateResponse{}
	_body, _err := client.UpdateDataQualityRuleTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data quality monitor.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - UpdateDataQualityScanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityScanResponse
func (client *Client) UpdateDataQualityScanWithOptions(tmpReq *UpdateDataQualityScanRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityScanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataQualityScanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ComputeResource) {
		request.ComputeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ComputeResource, dara.String("ComputeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Hooks) {
		request.HooksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hooks, dara.String("Hooks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ComputeResourceShrink) {
		body["ComputeResource"] = request.ComputeResourceShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HooksShrink) {
		body["Hooks"] = request.HooksShrink
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityScan"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data quality monitor.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateDataQualityScanRequest
//
// @return UpdateDataQualityScanResponse
func (client *Client) UpdateDataQualityScan(request *UpdateDataQualityScanRequest) (_result *UpdateDataQualityScanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataQualityScanResponse{}
	_body, _err := client.UpdateDataQualityScanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Updates a data quality rule template in a project
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateDataQualityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataQualityTemplateResponse
func (client *Client) UpdateDataQualityTemplateWithOptions(request *UpdateDataQualityTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataQualityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataQualityTemplate"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataQualityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Updates a data quality rule template in a project
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateDataQualityTemplateRequest
//
// @return UpdateDataQualityTemplateResponse
func (client *Client) UpdateDataQualityTemplate(request *UpdateDataQualityTemplateRequest) (_result *UpdateDataQualityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataQualityTemplateResponse{}
	_body, _err := client.UpdateDataQualityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a data source by ID.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Tenant Administrator, Workspace Administrator, Workspace Owner, and O\\&M
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionProperties) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !dara.IsNil(request.ConnectionPropertiesMode) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSource"),
		Version:     dara.String("2024-05-18"),
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
// Modifies a data source by ID.
//
// Description:
//
// 1.  This API operation is available for all DataWorks editions.
//
// 2.  You can call this operation only if you are assigned one of the following roles in DataWorks:
//
//   - Tenant Owner, Tenant Administrator, Workspace Administrator, Workspace Owner, and O\\&M
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

// @param request - UpdateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileResponse
func (client *Client) UpdateFileWithOptions(request *UpdateFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedSettings) {
		body["AdvancedSettings"] = request.AdvancedSettings
	}

	if !dara.IsNil(request.ApplyScheduleImmediately) {
		body["ApplyScheduleImmediately"] = request.ApplyScheduleImmediately
	}

	if !dara.IsNil(request.AutoParsing) {
		body["AutoParsing"] = request.AutoParsing
	}

	if !dara.IsNil(request.AutoRerunIntervalMillis) {
		body["AutoRerunIntervalMillis"] = request.AutoRerunIntervalMillis
	}

	if !dara.IsNil(request.AutoRerunTimes) {
		body["AutoRerunTimes"] = request.AutoRerunTimes
	}

	if !dara.IsNil(request.ConnectionName) {
		body["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.CronExpress) {
		body["CronExpress"] = request.CronExpress
	}

	if !dara.IsNil(request.CycleType) {
		body["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.DependentNodeIdList) {
		body["DependentNodeIdList"] = request.DependentNodeIdList
	}

	if !dara.IsNil(request.DependentType) {
		body["DependentType"] = request.DependentType
	}

	if !dara.IsNil(request.EndEffectDate) {
		body["EndEffectDate"] = request.EndEffectDate
	}

	if !dara.IsNil(request.FileDescription) {
		body["FileDescription"] = request.FileDescription
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.IgnoreParentSkipRunningProperty) {
		body["IgnoreParentSkipRunningProperty"] = request.IgnoreParentSkipRunningProperty
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InputList) {
		body["InputList"] = request.InputList
	}

	if !dara.IsNil(request.InputParameters) {
		body["InputParameters"] = request.InputParameters
	}

	if !dara.IsNil(request.OutputList) {
		body["OutputList"] = request.OutputList
	}

	if !dara.IsNil(request.OutputParameters) {
		body["OutputParameters"] = request.OutputParameters
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParaValue) {
		body["ParaValue"] = request.ParaValue
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.RerunMode) {
		body["RerunMode"] = request.RerunMode
	}

	if !dara.IsNil(request.ResourceGroupIdentifier) {
		body["ResourceGroupIdentifier"] = request.ResourceGroupIdentifier
	}

	if !dara.IsNil(request.SchedulerType) {
		body["SchedulerType"] = request.SchedulerType
	}

	if !dara.IsNil(request.StartEffectDate) {
		body["StartEffectDate"] = request.StartEffectDate
	}

	if !dara.IsNil(request.StartImmediately) {
		body["StartImmediately"] = request.StartImmediately
	}

	if !dara.IsNil(request.Stop) {
		body["Stop"] = request.Stop
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateFileRequest
//
// @return UpdateFileResponse
func (client *Client) UpdateFile(request *UpdateFileRequest) (_result *UpdateFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFileResponse{}
	_body, _err := client.UpdateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolderWithOptions(request *UpdateFolderRequest, runtime *dara.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.FolderName) {
		body["FolderName"] = request.FolderName
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFolder"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateFolderRequest
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolder(request *UpdateFolderRequest) (_result *UpdateFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFolderResponse{}
	_body, _err := client.UpdateFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about a user-defined function (UDF) in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunctionWithOptions(request *UpdateFunctionRequest, runtime *dara.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFunction"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a user-defined function (UDF) in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateFunctionRequest
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunction(request *UpdateFunctionRequest) (_result *UpdateFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.UpdateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 回调扩展点消息的检查结果
//
// @param request - UpdateIDEEventResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIDEEventResultResponse
func (client *Client) UpdateIDEEventResultWithOptions(request *UpdateIDEEventResultRequest, runtime *dara.RuntimeOptions) (_result *UpdateIDEEventResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckResult) {
		body["CheckResult"] = request.CheckResult
	}

	if !dara.IsNil(request.CheckResultTip) {
		body["CheckResultTip"] = request.CheckResultTip
	}

	if !dara.IsNil(request.ExtensionCode) {
		body["ExtensionCode"] = request.ExtensionCode
	}

	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIDEEventResult"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIDEEventResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回调扩展点消息的检查结果
//
// @param request - UpdateIDEEventResultRequest
//
// @return UpdateIDEEventResultResponse
func (client *Client) UpdateIDEEventResult(request *UpdateIDEEventResultRequest) (_result *UpdateIDEEventResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIDEEventResultResponse{}
	_body, _err := client.UpdateIDEEventResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a collection in Data Map, including the collection name, description, and administrator. Collections include categories and data albums. If you want to update the information about a data album, the account that you use must be attached the AliyunDataWorksFullAccess policy, or you are the data album creator or administrator.
//
// @param tmpReq - UpdateMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaCollectionResponse
func (client *Client) UpdateMetaCollectionWithOptions(tmpReq *UpdateMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetaCollectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMetaCollectionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Administrators) {
		request.AdministratorsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Administrators, dara.String("Administrators"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdministratorsShrink) {
		query["Administrators"] = request.AdministratorsShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetaCollection"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a collection in Data Map, including the collection name, description, and administrator. Collections include categories and data albums. If you want to update the information about a data album, the account that you use must be attached the AliyunDataWorksFullAccess policy, or you are the data album creator or administrator.
//
// @param request - UpdateMetaCollectionRequest
//
// @return UpdateMetaCollectionResponse
func (client *Client) UpdateMetaCollection(request *UpdateMetaCollectionRequest) (_result *UpdateMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMetaCollectionResponse{}
	_body, _err := client.UpdateMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about a node in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeResponse
func (client *Client) UpdateNodeWithOptions(request *UpdateNodeRequest, runtime *dara.RuntimeOptions) (_result *UpdateNodeResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNode"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a node in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateNodeRequest
//
// @return UpdateNodeResponse
func (client *Client) UpdateNode(request *UpdateNodeRequest) (_result *UpdateNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNodeResponse{}
	_body, _err := client.UpdateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a DataWorks workspace.
//
// @param request - UpdateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DevEnvironmentEnabled) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !dara.IsNil(request.DevRoleDisabled) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.PaiTaskEnabled) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProject"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a DataWorks workspace.
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about a file resource in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithOptions(request *UpdateResourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceFile) {
		body["ResourceFile"] = request.ResourceFile
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResource"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a file resource in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateResourceRequest
//
// @return UpdateResourceResponse
func (client *Client) UpdateResource(request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceAdvance(request *UpdateResourceAdvanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("dataworks-public"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	updateResourceReq := &UpdateResourceRequest{}
	openapiutil.Convert(request, updateResourceReq)
	if !dara.IsNil(request.ResourceFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.ResourceFileObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		updateResourceReq.ResourceFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	updateResourceResp, _err := client.UpdateResourceWithOptions(updateResourceReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = updateResourceResp
	return _result, _err
}

// Summary:
//
// Updates basic information about a resource group.
//
// Description:
//
// You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// @param request - UpdateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroupWithOptions(request *UpdateResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunResourceGroupId) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResourceGroup"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates basic information about a resource group.
//
// Description:
//
// You can use this API operation only in DataWorks Basic Edition or an advanced edition.
//
// @param request - UpdateResourceGroupRequest
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroup(request *UpdateResourceGroupRequest) (_result *UpdateResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.UpdateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a route.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRouteResponse
func (client *Client) UpdateRouteWithOptions(request *UpdateRouteRequest, runtime *dara.RuntimeOptions) (_result *UpdateRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidr) {
		body["DestinationCidr"] = request.DestinationCidr
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRoute"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a route.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateRouteRequest
//
// @return UpdateRouteResponse
func (client *Client) UpdateRoute(request *UpdateRouteRequest) (_result *UpdateRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRouteResponse{}
	_body, _err := client.UpdateRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the business metadata of a table in Data Map. Currently, only the usage notes of a table can be updated.
//
// @param request - UpdateTableBusinessMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableBusinessMetadataResponse
func (client *Client) UpdateTableBusinessMetadataWithOptions(request *UpdateTableBusinessMetadataRequest, runtime *dara.RuntimeOptions) (_result *UpdateTableBusinessMetadataResponse, _err error) {
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

	if !dara.IsNil(request.Readme) {
		body["Readme"] = request.Readme
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTableBusinessMetadata"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableBusinessMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the business metadata of a table in Data Map. Currently, only the usage notes of a table can be updated.
//
// @param request - UpdateTableBusinessMetadataRequest
//
// @return UpdateTableBusinessMetadataResponse
func (client *Client) UpdateTableBusinessMetadata(request *UpdateTableBusinessMetadataRequest) (_result *UpdateTableBusinessMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTableBusinessMetadataResponse{}
	_body, _err := client.UpdateTableBusinessMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update a task. The changes are synchronized to Data Studio, which creates a new saved version.
//
// @param tmpReq - UpdateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskResponse
func (client *Client) UpdateTaskWithOptions(tmpReq *UpdateTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSource) {
		request.DataSourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSource, dara.String("DataSource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Dependencies) {
		request.DependenciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dependencies, dara.String("Dependencies"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Inputs) {
		request.InputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Inputs, dara.String("Inputs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Outputs) {
		request.OutputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Outputs, dara.String("Outputs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimeResource) {
		request.RuntimeResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeResource, dara.String("RuntimeResource"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Script) {
		request.ScriptShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Script, dara.String("Script"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientUniqueCode) {
		body["ClientUniqueCode"] = request.ClientUniqueCode
	}

	if !dara.IsNil(request.DataSourceShrink) {
		body["DataSource"] = request.DataSourceShrink
	}

	if !dara.IsNil(request.DependenciesShrink) {
		body["Dependencies"] = request.DependenciesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InputsShrink) {
		body["Inputs"] = request.InputsShrink
	}

	if !dara.IsNil(request.InstanceMode) {
		body["InstanceMode"] = request.InstanceMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputsShrink) {
		body["Outputs"] = request.OutputsShrink
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.RerunInterval) {
		body["RerunInterval"] = request.RerunInterval
	}

	if !dara.IsNil(request.RerunMode) {
		body["RerunMode"] = request.RerunMode
	}

	if !dara.IsNil(request.RerunTimes) {
		body["RerunTimes"] = request.RerunTimes
	}

	if !dara.IsNil(request.RuntimeResourceShrink) {
		body["RuntimeResource"] = request.RuntimeResourceShrink
	}

	if !dara.IsNil(request.ScriptShrink) {
		body["Script"] = request.ScriptShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTask"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a task. The changes are synchronized to Data Studio, which creates a new saved version.
//
// @param request - UpdateTaskRequest
//
// @return UpdateTaskResponse
func (client *Client) UpdateTask(request *UpdateTaskRequest) (_result *UpdateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskResponse{}
	_body, _err := client.UpdateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies properties configured for multiple instances at a time. The properties include the priority, resource group for scheduling, and data source.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - UpdateTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskInstancesResponse
func (client *Client) UpdateTaskInstancesWithOptions(tmpReq *UpdateTaskInstancesRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTaskInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskInstances) {
		request.TaskInstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskInstances, dara.String("TaskInstances"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.TaskInstancesShrink) {
		body["TaskInstances"] = request.TaskInstancesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskInstances"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies properties configured for multiple instances at a time. The properties include the priority, resource group for scheduling, and data source.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateTaskInstancesRequest
//
// @return UpdateTaskInstancesResponse
func (client *Client) UpdateTaskInstances(request *UpdateTaskInstancesRequest) (_result *UpdateTaskInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskInstancesResponse{}
	_body, _err := client.UpdateTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateUdfFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfFileResponse
func (client *Client) UpdateUdfFileWithOptions(request *UpdateUdfFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateUdfFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		body["ClassName"] = request.ClassName
	}

	if !dara.IsNil(request.CmdDescription) {
		body["CmdDescription"] = request.CmdDescription
	}

	if !dara.IsNil(request.Example) {
		body["Example"] = request.Example
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FunctionType) {
		body["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.ParameterDescription) {
		body["ParameterDescription"] = request.ParameterDescription
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.Resources) {
		body["Resources"] = request.Resources
	}

	if !dara.IsNil(request.ReturnValue) {
		body["ReturnValue"] = request.ReturnValue
	}

	if !dara.IsNil(request.UdfDescription) {
		body["UdfDescription"] = request.UdfDescription
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUdfFile"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUdfFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateUdfFileRequest
//
// @return UpdateUdfFileResponse
func (client *Client) UpdateUdfFile(request *UpdateUdfFileRequest) (_result *UpdateUdfFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUdfFileResponse{}
	_body, _err := client.UpdateUdfFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a specified workflow in full update mode.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param tmpReq - UpdateWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflowWithOptions(tmpReq *UpdateWorkflowRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateWorkflowShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Dependencies) {
		request.DependenciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dependencies, dara.String("Dependencies"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Outputs) {
		request.OutputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Outputs, dara.String("Outputs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tasks) {
		request.TasksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tasks, dara.String("Tasks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Trigger) {
		request.TriggerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Trigger, dara.String("Trigger"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientUniqueCode) {
		body["ClientUniqueCode"] = request.ClientUniqueCode
	}

	if !dara.IsNil(request.DependenciesShrink) {
		body["Dependencies"] = request.DependenciesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceMode) {
		body["InstanceMode"] = request.InstanceMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputsShrink) {
		body["Outputs"] = request.OutputsShrink
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Parameters) {
		body["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TasksShrink) {
		body["Tasks"] = request.TasksShrink
	}

	if !dara.IsNil(request.TriggerShrink) {
		body["Trigger"] = request.TriggerShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflow"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified workflow in full update mode.
//
// Description:
//
// This API operation is available for all DataWorks editions.
//
// @param request - UpdateWorkflowRequest
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflow(request *UpdateWorkflowRequest) (_result *UpdateWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkflowResponse{}
	_body, _err := client.UpdateWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about a workflow in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple workflows at a time. If you specify multiple workflows in the FlowSpec filed, only the first workflow is created. Other specified workflows and the nodes in the workflows are ignored. You can call the UpdateNode operation to update a node.
//
// @param request - UpdateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowDefinitionResponse
func (client *Client) UpdateWorkflowDefinitionWithOptions(request *UpdateWorkflowDefinitionRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowDefinitionResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Spec) {
		body["Spec"] = request.Spec
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflowDefinition"),
		Version:     dara.String("2024-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a workflow in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple workflows at a time. If you specify multiple workflows in the FlowSpec filed, only the first workflow is created. Other specified workflows and the nodes in the workflows are ignored. You can call the UpdateNode operation to update a node.
//
// @param request - UpdateWorkflowDefinitionRequest
//
// @return UpdateWorkflowDefinitionResponse
func (client *Client) UpdateWorkflowDefinition(request *UpdateWorkflowDefinitionRequest) (_result *UpdateWorkflowDefinitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkflowDefinitionResponse{}
	_body, _err := client.UpdateWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func _postOSSObject_opResponse(response_ *dara.Response) (_result map[string]interface{}, _err error) {
	var respMap map[string]interface{}
	bodyStr, _err := dara.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		respMap = dara.ParseXml(bodyStr, nil)
		err := dara.ToMap(respMap["Error"])
		_err = &openapi.ClientError{
			Code:    dara.String(dara.ToString(err["Code"])),
			Message: dara.String(dara.ToString(err["Message"])),
			Data: map[string]interface{}{
				"httpCode":  dara.IntValue(response_.StatusCode),
				"requestId": dara.ToString(err["RequestId"]),
				"hostId":    dara.ToString(err["HostId"]),
			},
		}
		return _result, _err
	}

	respMap = dara.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = dara.Convert(dara.ToMap(respMap), &_result)

	return _result, _err
}
