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
		"cn-beijing":  dara.String("schedulerx.cn-beijing.aliyuncs.com"),
		"cn-hangzhou": dara.String("schedulerx.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai": dara.String("schedulerx.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen": dara.String("schedulerx.cn-shenzhen.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("schedulerx2"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Deletes multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//	<groupId>com.aliyun</groupId>
//
//	<artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	<version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchDeleteJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteJobsResponse
func (client *Client) BatchDeleteJobsWithOptions(request *BatchDeleteJobsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteJobsResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JobIdList) {
		body["JobIdList"] = request.JobIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteJobs"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//	<groupId>com.aliyun</groupId>
//
//	<artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	<version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchDeleteJobsRequest
//
// @return BatchDeleteJobsResponse
func (client *Client) BatchDeleteJobs(request *BatchDeleteJobsRequest) (_result *BatchDeleteJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteJobsResponse{}
	_body, _err := client.BatchDeleteJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The additional information that is returned.
//
// @param request - BatchDeleteRouteStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteRouteStrategyResponse
func (client *Client) BatchDeleteRouteStrategyWithOptions(request *BatchDeleteRouteStrategyRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteRouteStrategyResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JobIdList) {
		body["JobIdList"] = request.JobIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteRouteStrategy"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteRouteStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The additional information that is returned.
//
// @param request - BatchDeleteRouteStrategyRequest
//
// @return BatchDeleteRouteStrategyResponse
func (client *Client) BatchDeleteRouteStrategy(request *BatchDeleteRouteStrategyRequest) (_result *BatchDeleteRouteStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteRouteStrategyResponse{}
	_body, _err := client.BatchDeleteRouteStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//	<groupId>com.aliyun</groupId>
//
//	<artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	<version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchDisableJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDisableJobsResponse
func (client *Client) BatchDisableJobsWithOptions(request *BatchDisableJobsRequest, runtime *dara.RuntimeOptions) (_result *BatchDisableJobsResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JobIdList) {
		body["JobIdList"] = request.JobIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDisableJobs"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDisableJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//	<groupId>com.aliyun</groupId>
//
//	<artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	<version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchDisableJobsRequest
//
// @return BatchDisableJobsResponse
func (client *Client) BatchDisableJobs(request *BatchDisableJobsRequest) (_result *BatchDisableJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDisableJobsResponse{}
	_body, _err := client.BatchDisableJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//	<groupId>com.aliyun</groupId>
//
//	<artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	<version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchEnableJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchEnableJobsResponse
func (client *Client) BatchEnableJobsWithOptions(request *BatchEnableJobsRequest, runtime *dara.RuntimeOptions) (_result *BatchEnableJobsResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JobIdList) {
		body["JobIdList"] = request.JobIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchEnableJobs"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchEnableJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables multiple jobs at a time.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//	<groupId>com.aliyun</groupId>
//
//	<artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	<version>1.0.4</version>
//
// </dependency>
//
// ```
//
// @param request - BatchEnableJobsRequest
//
// @return BatchEnableJobsResponse
func (client *Client) BatchEnableJobs(request *BatchEnableJobsRequest) (_result *BatchEnableJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchEnableJobsResponse{}
	_body, _err := client.BatchEnableJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application group. The AppKey is returned.
//
// @param request - CreateAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppGroupResponse
func (client *Client) CreateAppGroupWithOptions(request *CreateAppGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAppGroupResponse, _err error) {
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
		Action:      dara.String("CreateAppGroup"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application group. The AppKey is returned.
//
// @param request - CreateAppGroupRequest
//
// @return CreateAppGroupResponse
func (client *Client) CreateAppGroup(request *CreateAppGroupRequest) (_result *CreateAppGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CreateAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a job and obtains the job ID.
//
// @param request - CreateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithOptions(request *CreateJobRequest, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttemptInterval) {
		body["AttemptInterval"] = request.AttemptInterval
	}

	if !dara.IsNil(request.Calendar) {
		body["Calendar"] = request.Calendar
	}

	if !dara.IsNil(request.ClassName) {
		body["ClassName"] = request.ClassName
	}

	if !dara.IsNil(request.ConsumerSize) {
		body["ConsumerSize"] = request.ConsumerSize
	}

	if !dara.IsNil(request.ContactInfo) {
		body["ContactInfo"] = request.ContactInfo
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.DataOffset) {
		body["DataOffset"] = request.DataOffset
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DispatcherSize) {
		body["DispatcherSize"] = request.DispatcherSize
	}

	if !dara.IsNil(request.ExecuteMode) {
		body["ExecuteMode"] = request.ExecuteMode
	}

	if !dara.IsNil(request.FailEnable) {
		body["FailEnable"] = request.FailEnable
	}

	if !dara.IsNil(request.FailTimes) {
		body["FailTimes"] = request.FailTimes
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.JobType) {
		body["JobType"] = request.JobType
	}

	if !dara.IsNil(request.MaxAttempt) {
		body["MaxAttempt"] = request.MaxAttempt
	}

	if !dara.IsNil(request.MaxConcurrency) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.MissWorkerEnable) {
		body["MissWorkerEnable"] = request.MissWorkerEnable
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Parameters) {
		body["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.QueueSize) {
		body["QueueSize"] = request.QueueSize
	}

	if !dara.IsNil(request.SendChannel) {
		body["SendChannel"] = request.SendChannel
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.SuccessNoticeEnable) {
		body["SuccessNoticeEnable"] = request.SuccessNoticeEnable
	}

	if !dara.IsNil(request.TaskAttemptInterval) {
		body["TaskAttemptInterval"] = request.TaskAttemptInterval
	}

	if !dara.IsNil(request.TaskMaxAttempt) {
		body["TaskMaxAttempt"] = request.TaskMaxAttempt
	}

	if !dara.IsNil(request.TimeExpression) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !dara.IsNil(request.TimeType) {
		body["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.TimeoutEnable) {
		body["TimeoutEnable"] = request.TimeoutEnable
	}

	if !dara.IsNil(request.TimeoutKillEnable) {
		body["TimeoutKillEnable"] = request.TimeoutKillEnable
	}

	if !dara.IsNil(request.Timezone) {
		body["Timezone"] = request.Timezone
	}

	if !dara.IsNil(request.XAttrs) {
		body["XAttrs"] = request.XAttrs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJob"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a job and obtains the job ID.
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a namespace.
//
// @param request - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNamespace"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a namespace.
//
// @param request - CreateNamespaceRequest
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a routing policy.
//
// @param request - CreateRouteStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteStrategyResponse
func (client *Client) CreateRouteStrategyWithOptions(request *CreateRouteStrategyRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteStrategyResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StrategyContent) {
		query["StrategyContent"] = request.StrategyContent
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRouteStrategy"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRouteStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a routing policy.
//
// @param request - CreateRouteStrategyRequest
//
// @return CreateRouteStrategyResponse
func (client *Client) CreateRouteStrategy(request *CreateRouteStrategyRequest) (_result *CreateRouteStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRouteStrategyResponse{}
	_body, _err := client.CreateRouteStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow. By default, the created workflow is disabled. After you update the directed acyclic graph (DAG) of the workflow, you must manually or call the corresponding operation to enable the workflow. You can call this operation only in the professional edition.
//
// @param request - CreateWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflowWithOptions(request *CreateWorkflowRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkflowResponse, _err error) {
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

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MaxConcurrency) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TimeExpression) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !dara.IsNil(request.TimeType) {
		body["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.Timezone) {
		body["Timezone"] = request.Timezone
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkflow"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow. By default, the created workflow is disabled. After you update the directed acyclic graph (DAG) of the workflow, you must manually or call the corresponding operation to enable the workflow. You can call this operation only in the professional edition.
//
// @param request - CreateWorkflowRequest
//
// @return CreateWorkflowResponse
func (client *Client) CreateWorkflow(request *CreateWorkflowRequest) (_result *CreateWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWorkflowResponse{}
	_body, _err := client.CreateWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The additional information that is returned.
//
// @param request - DeleteAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppGroupResponse
func (client *Client) DeleteAppGroupWithOptions(request *DeleteAppGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteJobs) {
		query["DeleteJobs"] = request.DeleteJobs
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppGroup"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The additional information that is returned.
//
// @param request - DeleteAppGroupRequest
//
// @return DeleteAppGroupResponse
func (client *Client) DeleteAppGroup(request *DeleteAppGroupRequest) (_result *DeleteAppGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppGroupResponse{}
	_body, _err := client.DeleteAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified job.
//
// @param request - DeleteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
func (client *Client) DeleteJobWithOptions(request *DeleteJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
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
		Action:      dara.String("DeleteJob"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified job.
//
// @param request - DeleteJobRequest
//
// @return DeleteJobResponse
func (client *Client) DeleteJob(request *DeleteJobRequest) (_result *DeleteJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除命名空间
//
// @param request - DeleteNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNamespace"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除命名空间
//
// @param request - DeleteNamespaceRequest
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a routing policy.
//
// @param request - DeleteRouteStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteStrategyResponse
func (client *Client) DeleteRouteStrategyWithOptions(request *DeleteRouteStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteStrategyResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRouteStrategy"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a routing policy.
//
// @param request - DeleteRouteStrategyRequest
//
// @return DeleteRouteStrategyResponse
func (client *Client) DeleteRouteStrategy(request *DeleteRouteStrategyRequest) (_result *DeleteRouteStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRouteStrategyResponse{}
	_body, _err := client.DeleteRouteStrategyWithOptions(request, runtime)
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
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkflow"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// Returns available regions.
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
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-04-30"),
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
// Returns available regions.
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
// Designates machines.
//
// @param request - DesignateWorkersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DesignateWorkersResponse
func (client *Client) DesignateWorkersWithOptions(request *DesignateWorkersRequest, runtime *dara.RuntimeOptions) (_result *DesignateWorkersResponse, _err error) {
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
		Action:      dara.String("DesignateWorkers"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DesignateWorkersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Designates machines.
//
// @param request - DesignateWorkersRequest
//
// @return DesignateWorkersResponse
func (client *Client) DesignateWorkers(request *DesignateWorkersRequest) (_result *DesignateWorkersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DesignateWorkersResponse{}
	_body, _err := client.DesignateWorkersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a job.
//
// @param request - DisableJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableJobResponse
func (client *Client) DisableJobWithOptions(request *DisableJobRequest, runtime *dara.RuntimeOptions) (_result *DisableJobResponse, _err error) {
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
		Action:      dara.String("DisableJob"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a job.
//
// @param request - DisableJobRequest
//
// @return DisableJobResponse
func (client *Client) DisableJob(request *DisableJobRequest) (_result *DisableJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableJobResponse{}
	_body, _err := client.DisableJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a specified workflow.
//
// @param request - DisableWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableWorkflowResponse
func (client *Client) DisableWorkflowWithOptions(request *DisableWorkflowRequest, runtime *dara.RuntimeOptions) (_result *DisableWorkflowResponse, _err error) {
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
		Action:      dara.String("DisableWorkflow"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a specified workflow.
//
// @param request - DisableWorkflowRequest
//
// @return DisableWorkflowResponse
func (client *Client) DisableWorkflow(request *DisableWorkflowRequest) (_result *DisableWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableWorkflowResponse{}
	_body, _err := client.DisableWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a job.
//
// @param request - EnableJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableJobResponse
func (client *Client) EnableJobWithOptions(request *EnableJobRequest, runtime *dara.RuntimeOptions) (_result *EnableJobResponse, _err error) {
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
		Action:      dara.String("EnableJob"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a job.
//
// @param request - EnableJobRequest
//
// @return EnableJobResponse
func (client *Client) EnableJob(request *EnableJobRequest) (_result *EnableJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableJobResponse{}
	_body, _err := client.EnableJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a specified workflow.
//
// @param request - EnableWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableWorkflowResponse
func (client *Client) EnableWorkflowWithOptions(request *EnableWorkflowRequest, runtime *dara.RuntimeOptions) (_result *EnableWorkflowResponse, _err error) {
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
		Action:      dara.String("EnableWorkflow"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a specified workflow.
//
// @param request - EnableWorkflowRequest
//
// @return EnableWorkflowResponse
func (client *Client) EnableWorkflow(request *EnableWorkflowRequest) (_result *EnableWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableWorkflowResponse{}
	_body, _err := client.EnableWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers a job to immediately run once.
//
// Description:
//
// > The combination of the `JobID` and `ScheduleTime` parameters serves as a unique index. Therefore, after the ExecuteJob operation is called to run a job once, a sleep for one second is required before the ExecuteJob operation is called to run the job again. Otherwise, the job may fail.
//
// @param request - ExecuteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteJobResponse
func (client *Client) ExecuteJobWithOptions(request *ExecuteJobRequest, runtime *dara.RuntimeOptions) (_result *ExecuteJobResponse, _err error) {
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
		Action:      dara.String("ExecuteJob"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Triggers a job to immediately run once.
//
// Description:
//
// > The combination of the `JobID` and `ScheduleTime` parameters serves as a unique index. Therefore, after the ExecuteJob operation is called to run a job once, a sleep for one second is required before the ExecuteJob operation is called to run the job again. Otherwise, the job may fail.
//
// @param request - ExecuteJobRequest
//
// @return ExecuteJobResponse
func (client *Client) ExecuteJob(request *ExecuteJobRequest) (_result *ExecuteJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteJobResponse{}
	_body, _err := client.ExecuteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Immediately triggers a workflow.
//
// @param request - ExecuteWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteWorkflowResponse
func (client *Client) ExecuteWorkflowWithOptions(request *ExecuteWorkflowRequest, runtime *dara.RuntimeOptions) (_result *ExecuteWorkflowResponse, _err error) {
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
		Action:      dara.String("ExecuteWorkflow"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Immediately triggers a workflow.
//
// @param request - ExecuteWorkflowRequest
//
// @return ExecuteWorkflowResponse
func (client *Client) ExecuteWorkflow(request *ExecuteWorkflowRequest) (_result *ExecuteWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteWorkflowResponse{}
	_body, _err := client.ExecuteWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The configuration of the alert. The value is a JSON string. For more information, see \\\\*\\\\*the additional information about response parameters below this table\\\\*\\\\*.
//
// @param request - GetAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppGroupResponse
func (client *Client) GetAppGroupWithOptions(request *GetAppGroupRequest, runtime *dara.RuntimeOptions) (_result *GetAppGroupResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppGroup"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The configuration of the alert. The value is a JSON string. For more information, see \\\\*\\\\*the additional information about response parameters below this table\\\\*\\\\*.
//
// @param request - GetAppGroupRequest
//
// @return GetAppGroupResponse
func (client *Client) GetAppGroup(request *GetAppGroupRequest) (_result *GetAppGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppGroupResponse{}
	_body, _err := client.GetAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a job based on the job ID. In most cases, the obtained information is used to update jobs.
//
// @param request - GetJobInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobInfoResponse
func (client *Client) GetJobInfoWithOptions(request *GetJobInfoRequest, runtime *dara.RuntimeOptions) (_result *GetJobInfoResponse, _err error) {
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
		Action:      dara.String("GetJobInfo"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a job based on the job ID. In most cases, the obtained information is used to update jobs.
//
// @param request - GetJobInfoRequest
//
// @return GetJobInfoResponse
func (client *Client) GetJobInfo(request *GetJobInfoRequest) (_result *GetJobInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobInfoResponse{}
	_body, _err := client.GetJobInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a job instance. You can view the status and progress of the job instance.
//
// @param request - GetJobInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobInstanceResponse
func (client *Client) GetJobInstanceWithOptions(request *GetJobInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetJobInstanceResponse, _err error) {
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
		Action:      dara.String("GetJobInstance"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a job instance. You can view the status and progress of the job instance.
//
// @param request - GetJobInstanceRequest
//
// @return GetJobInstanceResponse
func (client *Client) GetJobInstance(request *GetJobInstanceRequest) (_result *GetJobInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobInstanceResponse{}
	_body, _err := client.GetJobInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the most recent 10 execution instances of a job.
//
// @param request - GetJobInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobInstanceListResponse
func (client *Client) GetJobInstanceListWithOptions(request *GetJobInstanceListRequest, runtime *dara.RuntimeOptions) (_result *GetJobInstanceListResponse, _err error) {
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
		Action:      dara.String("GetJobInstanceList"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the most recent 10 execution instances of a job.
//
// @param request - GetJobInstanceListRequest
//
// @return GetJobInstanceListResponse
func (client *Client) GetJobInstanceList(request *GetJobInstanceListRequest) (_result *GetJobInstanceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetJobInstanceListResponse{}
	_body, _err := client.GetJobInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operational logs of a job. You can call this operation only in the professional edition.
//
// @param request - GetLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogResponse
func (client *Client) GetLogWithOptions(request *GetLogRequest, runtime *dara.RuntimeOptions) (_result *GetLogResponse, _err error) {
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
		Action:      dara.String("GetLog"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operational logs of a job. You can call this operation only in the professional edition.
//
// @param request - GetLogRequest
//
// @return GetLogResponse
func (client *Client) GetLog(request *GetLogRequest) (_result *GetLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLogResponse{}
	_body, _err := client.GetLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询概览数据信息
//
// @param request - GetOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOverviewResponse
func (client *Client) GetOverviewWithOptions(request *GetOverviewRequest, runtime *dara.RuntimeOptions) (_result *GetOverviewResponse, _err error) {
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

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.Operate) {
		query["Operate"] = request.Operate
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOverview"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询概览数据信息
//
// @param request - GetOverviewRequest
//
// @return GetOverviewResponse
func (client *Client) GetOverview(request *GetOverviewRequest) (_result *GetOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOverviewResponse{}
	_body, _err := client.GetOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about a workflow.
//
// @param request - GetWorkFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkFlowResponse
func (client *Client) GetWorkFlowWithOptions(request *GetWorkFlowRequest, runtime *dara.RuntimeOptions) (_result *GetWorkFlowResponse, _err error) {
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
		Action:      dara.String("GetWorkFlow"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about a workflow.
//
// @param request - GetWorkFlowRequest
//
// @return GetWorkFlowResponse
func (client *Client) GetWorkFlow(request *GetWorkFlowRequest) (_result *GetWorkFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkFlowResponse{}
	_body, _err := client.GetWorkFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the list of workers that are connected to an application.
//
// @param request - GetWorkerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkerListResponse
func (client *Client) GetWorkerListWithOptions(request *GetWorkerListRequest, runtime *dara.RuntimeOptions) (_result *GetWorkerListResponse, _err error) {
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
		Action:      dara.String("GetWorkerList"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkerListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of workers that are connected to an application.
//
// @param request - GetWorkerListRequest
//
// @return GetWorkerListResponse
func (client *Client) GetWorkerList(request *GetWorkerListRequest) (_result *GetWorkerListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkerListResponse{}
	_body, _err := client.GetWorkerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified workflow instance, including the state of the workflow instance, the state of each job instance, and the dependencies between job instances. You can call this operation only in the professional edition.
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
		Version:     dara.String("2019-04-30"),
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
// Queries the details of a specified workflow instance, including the state of the workflow instance, the state of each job instance, and the dependencies between job instances. You can call this operation only in the professional edition.
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
// Grants permissions to an application group.
//
// @param request - GrantPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantPermissionResponse
func (client *Client) GrantPermissionWithOptions(request *GrantPermissionRequest, runtime *dara.RuntimeOptions) (_result *GrantPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GrantOption) {
		query["GrantOption"] = request.GrantOption
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantPermission"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions to an application group.
//
// @param request - GrantPermissionRequest
//
// @return GrantPermissionResponse
func (client *Client) GrantPermission(request *GrantPermissionRequest) (_result *GrantPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantPermissionResponse{}
	_body, _err := client.GrantPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of applications.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//	<groupId>com.aliyun</groupId>
//
//	<artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	<version>1.0.5</version>
//
// </dependency>
//
// ```
//
// @param request - ListGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGroupName) {
		query["AppGroupName"] = request.AppGroupName
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGroups"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of applications.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
// ```xml
//
// <dependency>
//
//	<groupId>com.aliyun</groupId>
//
//	<artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	<version>1.0.5</version>
//
// </dependency>
//
// ```
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务脚本历史列表
//
// @param request - ListJobScriptHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobScriptHistoryResponse
func (client *Client) ListJobScriptHistoryWithOptions(request *ListJobScriptHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListJobScriptHistoryResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobScriptHistory"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobScriptHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务脚本历史列表
//
// @param request - ListJobScriptHistoryRequest
//
// @return ListJobScriptHistoryResponse
func (client *Client) ListJobScriptHistory(request *ListJobScriptHistoryRequest) (_result *ListJobScriptHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobScriptHistoryResponse{}
	_body, _err := client.ListJobScriptHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries jobs.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
//	<dependency>
//
//	      <groupId>com.aliyun</groupId>
//
//	      <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	      <version>1.0.5</version>
//
//	</dependency>
//
// @param request - ListJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithOptions(request *ListJobsRequest, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
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
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries jobs.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
//	<dependency>
//
//	      <groupId>com.aliyun</groupId>
//
//	      <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	      <version>1.0.5</version>
//
//	</dependency>
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries namespaces.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
//	<dependency>
//
//	    <groupId>com.aliyun</groupId>
//
//	    <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	    <version>1.0.5</version>
//
//	</dependency>
//
// @param request - ListNamespacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamespacesResponse
func (client *Client) ListNamespacesWithOptions(request *ListNamespacesRequest, runtime *dara.RuntimeOptions) (_result *ListNamespacesResponse, _err error) {
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

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNamespaces"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries namespaces.
//
// Description:
//
// Before you call this operation, you must add the following dependency to the pom.xml file:
//
//	<dependency>
//
//	    <groupId>com.aliyun</groupId>
//
//	    <artifactId>aliyun-java-sdk-schedulerx2</artifactId>
//
//	    <version>1.0.5</version>
//
//	</dependency>
//
// @param request - ListNamespacesRequest
//
// @return ListNamespacesResponse
func (client *Client) ListNamespaces(request *ListNamespacesRequest) (_result *ListNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNamespacesResponse{}
	_body, _err := client.ListNamespacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution history of a workflow. You can call this operation only in the professional edition.
//
// @param request - ListWorkflowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowInstanceResponse
func (client *Client) ListWorkflowInstanceWithOptions(request *ListWorkflowInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowInstanceResponse, _err error) {
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
		Action:      dara.String("ListWorkflowInstance"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution history of a workflow. You can call this operation only in the professional edition.
//
// @param request - ListWorkflowInstanceRequest
//
// @return ListWorkflowInstanceResponse
func (client *Client) ListWorkflowInstance(request *ListWorkflowInstanceRequest) (_result *ListWorkflowInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkflowInstanceResponse{}
	_body, _err := client.ListWorkflowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步任务
//
// @param tmpReq - ManageSchedulerxJobSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ManageSchedulerxJobSyncResponse
func (client *Client) ManageSchedulerxJobSyncWithOptions(tmpReq *ManageSchedulerxJobSyncRequest, runtime *dara.RuntimeOptions) (_result *ManageSchedulerxJobSyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ManageSchedulerxJobSyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobIdList) {
		request.JobIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobIdList, dara.String("JobIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JobIdListShrink) {
		body["JobIdList"] = request.JobIdListShrink
	}

	if !dara.IsNil(request.NamespaceSource) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.OriginalGroupId) {
		body["OriginalGroupId"] = request.OriginalGroupId
	}

	if !dara.IsNil(request.OriginalNamespace) {
		body["OriginalNamespace"] = request.OriginalNamespace
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetGroupId) {
		body["TargetGroupId"] = request.TargetGroupId
	}

	if !dara.IsNil(request.TargetNamespace) {
		body["TargetNamespace"] = request.TargetNamespace
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ManageSchedulerxJobSync"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ManageSchedulerxJobSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 同步任务
//
// @param request - ManageSchedulerxJobSyncRequest
//
// @return ManageSchedulerxJobSyncResponse
func (client *Client) ManageSchedulerxJobSync(request *ManageSchedulerxJobSyncRequest) (_result *ManageSchedulerxJobSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ManageSchedulerxJobSyncResponse{}
	_body, _err := client.ManageSchedulerxJobSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取机器详细信息
//
// @param request - ReadSchedulerxDesignateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadSchedulerxDesignateDetailResponse
func (client *Client) ReadSchedulerxDesignateDetailWithOptions(request *ReadSchedulerxDesignateDetailRequest, runtime *dara.RuntimeOptions) (_result *ReadSchedulerxDesignateDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DesignateType) {
		query["DesignateType"] = request.DesignateType
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadSchedulerxDesignateDetail"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadSchedulerxDesignateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取机器详细信息
//
// @param request - ReadSchedulerxDesignateDetailRequest
//
// @return ReadSchedulerxDesignateDetailResponse
func (client *Client) ReadSchedulerxDesignateDetail(request *ReadSchedulerxDesignateDetailRequest) (_result *ReadSchedulerxDesignateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadSchedulerxDesignateDetailResponse{}
	_body, _err := client.ReadSchedulerxDesignateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取指定机器基本信息
//
// @param request - ReadSchedulerxDesignateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReadSchedulerxDesignateInfoResponse
func (client *Client) ReadSchedulerxDesignateInfoWithOptions(request *ReadSchedulerxDesignateInfoRequest, runtime *dara.RuntimeOptions) (_result *ReadSchedulerxDesignateInfoResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReadSchedulerxDesignateInfo"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReadSchedulerxDesignateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定机器基本信息
//
// @param request - ReadSchedulerxDesignateInfoRequest
//
// @return ReadSchedulerxDesignateInfoResponse
func (client *Client) ReadSchedulerxDesignateInfo(request *ReadSchedulerxDesignateInfoRequest) (_result *ReadSchedulerxDesignateInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReadSchedulerxDesignateInfoResponse{}
	_body, _err := client.ReadSchedulerxDesignateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns a job to obtain the historical data of the job. You can call this operation only in the professional edition.
//
// @param request - RerunJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RerunJobResponse
func (client *Client) RerunJobWithOptions(request *RerunJobRequest, runtime *dara.RuntimeOptions) (_result *RerunJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataTime) {
		body["DataTime"] = request.DataTime
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RerunJob"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RerunJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns a job to obtain the historical data of the job. You can call this operation only in the professional edition.
//
// @param request - RerunJobRequest
//
// @return RerunJobResponse
func (client *Client) RerunJob(request *RerunJobRequest) (_result *RerunJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RerunJobResponse{}
	_body, _err := client.RerunJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns a successful or failed job instance. You can call this operation only in the professional edition.
//
// @param request - RetryJobInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryJobInstanceResponse
func (client *Client) RetryJobInstanceWithOptions(request *RetryJobInstanceRequest, runtime *dara.RuntimeOptions) (_result *RetryJobInstanceResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobInstanceId) {
		query["JobInstanceId"] = request.JobInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryJobInstance"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryJobInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns a successful or failed job instance. You can call this operation only in the professional edition.
//
// @param request - RetryJobInstanceRequest
//
// @return RetryJobInstanceResponse
func (client *Client) RetryJobInstance(request *RetryJobInstanceRequest) (_result *RetryJobInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetryJobInstanceResponse{}
	_body, _err := client.RetryJobInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions that are granted to an Alibaba Cloud Resource Access Management (RAM) user.
//
// @param request - RevokePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokePermissionResponse
func (client *Client) RevokePermissionWithOptions(request *RevokePermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokePermissionResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokePermission"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions that are granted to an Alibaba Cloud Resource Access Management (RAM) user.
//
// @param request - RevokePermissionRequest
//
// @return RevokePermissionResponse
func (client *Client) RevokePermission(request *RevokePermissionRequest) (_result *RevokePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokePermissionResponse{}
	_body, _err := client.RevokePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Forcibly sets the state of a job instance to successful. You can call this operation only in the professional edition.
//
// @param request - SetJobInstanceSuccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetJobInstanceSuccessResponse
func (client *Client) SetJobInstanceSuccessWithOptions(request *SetJobInstanceSuccessRequest, runtime *dara.RuntimeOptions) (_result *SetJobInstanceSuccessResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobInstanceId) {
		query["JobInstanceId"] = request.JobInstanceId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetJobInstanceSuccess"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetJobInstanceSuccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Forcibly sets the state of a job instance to successful. You can call this operation only in the professional edition.
//
// @param request - SetJobInstanceSuccessRequest
//
// @return SetJobInstanceSuccessResponse
func (client *Client) SetJobInstanceSuccess(request *SetJobInstanceSuccessRequest) (_result *SetJobInstanceSuccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetJobInstanceSuccessResponse{}
	_body, _err := client.SetJobInstanceSuccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Forcibly sets the state of a workflow instance to successful. You can call this operation only in the professional edition.
//
// @param request - SetWfInstanceSuccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWfInstanceSuccessResponse
func (client *Client) SetWfInstanceSuccessWithOptions(request *SetWfInstanceSuccessRequest, runtime *dara.RuntimeOptions) (_result *SetWfInstanceSuccessResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		query["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.WfInstanceId) {
		query["WfInstanceId"] = request.WfInstanceId
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetWfInstanceSuccess"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetWfInstanceSuccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Forcibly sets the state of a workflow instance to successful. You can call this operation only in the professional edition.
//
// @param request - SetWfInstanceSuccessRequest
//
// @return SetWfInstanceSuccessResponse
func (client *Client) SetWfInstanceSuccess(request *SetWfInstanceSuccessRequest) (_result *SetWfInstanceSuccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetWfInstanceSuccessResponse{}
	_body, _err := client.SetWfInstanceSuccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a job instance in the running state.
//
// @param request - StopInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
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
		Action:      dara.String("StopInstance"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a job instance in the running state.
//
// @param request - StopInstanceRequest
//
// @return StopInstanceResponse
func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the application group.
//
// @param request - UpdateAppGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAppGroupResponse
func (client *Client) UpdateAppGroupWithOptions(request *UpdateAppGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateAppGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppVersion) {
		query["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MaxConcurrency) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.MonitorConfigJson) {
		query["MonitorConfigJson"] = request.MonitorConfigJson
	}

	if !dara.IsNil(request.MonitorContactsJson) {
		query["MonitorContactsJson"] = request.MonitorContactsJson
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NotificationPolicyName) {
		query["NotificationPolicyName"] = request.NotificationPolicyName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAppGroup"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the application group.
//
// @param request - UpdateAppGroupRequest
//
// @return UpdateAppGroupResponse
func (client *Client) UpdateAppGroup(request *UpdateAppGroupRequest) (_result *UpdateAppGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAppGroupResponse{}
	_body, _err := client.UpdateAppGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration information about a job. By default, you need to call the GetJobInfo operation to obtain the original configuration of the job before you call this operation to modify the configuration as required.
//
// @param request - UpdateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobResponse
func (client *Client) UpdateJobWithOptions(request *UpdateJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AttemptInterval) {
		body["AttemptInterval"] = request.AttemptInterval
	}

	if !dara.IsNil(request.Calendar) {
		body["Calendar"] = request.Calendar
	}

	if !dara.IsNil(request.ClassName) {
		body["ClassName"] = request.ClassName
	}

	if !dara.IsNil(request.ConsumerSize) {
		body["ConsumerSize"] = request.ConsumerSize
	}

	if !dara.IsNil(request.ContactInfo) {
		body["ContactInfo"] = request.ContactInfo
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.DataOffset) {
		body["DataOffset"] = request.DataOffset
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DispatcherSize) {
		body["DispatcherSize"] = request.DispatcherSize
	}

	if !dara.IsNil(request.ExecuteMode) {
		body["ExecuteMode"] = request.ExecuteMode
	}

	if !dara.IsNil(request.FailEnable) {
		body["FailEnable"] = request.FailEnable
	}

	if !dara.IsNil(request.FailTimes) {
		body["FailTimes"] = request.FailTimes
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MaxAttempt) {
		body["MaxAttempt"] = request.MaxAttempt
	}

	if !dara.IsNil(request.MaxConcurrency) {
		body["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.MissWorkerEnable) {
		body["MissWorkerEnable"] = request.MissWorkerEnable
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Parameters) {
		body["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.QueueSize) {
		body["QueueSize"] = request.QueueSize
	}

	if !dara.IsNil(request.SendChannel) {
		body["SendChannel"] = request.SendChannel
	}

	if !dara.IsNil(request.SuccessNoticeEnable) {
		body["SuccessNoticeEnable"] = request.SuccessNoticeEnable
	}

	if !dara.IsNil(request.TaskAttemptInterval) {
		body["TaskAttemptInterval"] = request.TaskAttemptInterval
	}

	if !dara.IsNil(request.TaskDispatchMode) {
		body["TaskDispatchMode"] = request.TaskDispatchMode
	}

	if !dara.IsNil(request.TaskMaxAttempt) {
		body["TaskMaxAttempt"] = request.TaskMaxAttempt
	}

	if !dara.IsNil(request.Template) {
		body["Template"] = request.Template
	}

	if !dara.IsNil(request.TimeExpression) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !dara.IsNil(request.TimeType) {
		body["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.TimeoutEnable) {
		body["TimeoutEnable"] = request.TimeoutEnable
	}

	if !dara.IsNil(request.TimeoutKillEnable) {
		body["TimeoutKillEnable"] = request.TimeoutKillEnable
	}

	if !dara.IsNil(request.Timezone) {
		body["Timezone"] = request.Timezone
	}

	if !dara.IsNil(request.XAttrs) {
		body["XAttrs"] = request.XAttrs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateJob"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration information about a job. By default, you need to call the GetJobInfo operation to obtain the original configuration of the job before you call this operation to modify the configuration as required.
//
// @param request - UpdateJobRequest
//
// @return UpdateJobResponse
func (client *Client) UpdateJob(request *UpdateJobRequest) (_result *UpdateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateJobResponse{}
	_body, _err := client.UpdateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新任务执行脚本
//
// @param request - UpdateJobScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobScriptResponse
func (client *Client) UpdateJobScriptWithOptions(request *UpdateJobScriptRequest, runtime *dara.RuntimeOptions) (_result *UpdateJobScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScriptContent) {
		body["ScriptContent"] = request.ScriptContent
	}

	if !dara.IsNil(request.VersionDescription) {
		body["VersionDescription"] = request.VersionDescription
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateJobScript"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateJobScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新任务执行脚本
//
// @param request - UpdateJobScriptRequest
//
// @return UpdateJobScriptResponse
func (client *Client) UpdateJobScript(request *UpdateJobScriptRequest) (_result *UpdateJobScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateJobScriptResponse{}
	_body, _err := client.UpdateJobScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新命名空间
//
// @param request - UpdateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNamespaceResponse
func (client *Client) UpdateNamespaceWithOptions(request *UpdateNamespaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateNamespaceResponse, _err error) {
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

	if !dara.IsNil(request.NamespaceName) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNamespace"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新命名空间
//
// @param request - UpdateNamespaceRequest
//
// @return UpdateNamespaceResponse
func (client *Client) UpdateNamespace(request *UpdateNamespaceRequest) (_result *UpdateNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.UpdateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about a workflow. You can call this operation only in the professional edition.
//
// @param request - UpdateWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowResponse
func (client *Client) UpdateWorkflowWithOptions(request *UpdateWorkflowRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.TimeExpression) {
		body["TimeExpression"] = request.TimeExpression
	}

	if !dara.IsNil(request.TimeType) {
		body["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflow"),
		Version:     dara.String("2019-04-30"),
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
// Updates the basic information about a workflow. You can call this operation only in the professional edition.
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
// Modifies the nodes and dependencies of a workflow. You can call this operation only in the professional edition.
//
// @param request - UpdateWorkflowDagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowDagResponse
func (client *Client) UpdateWorkflowDagWithOptions(request *UpdateWorkflowDagRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkflowDagResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.DagJson) {
		body["DagJson"] = request.DagJson
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NamespaceSource) {
		body["NamespaceSource"] = request.NamespaceSource
	}

	if !dara.IsNil(request.WorkflowId) {
		body["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkflowDag"),
		Version:     dara.String("2019-04-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkflowDagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the nodes and dependencies of a workflow. You can call this operation only in the professional edition.
//
// @param request - UpdateWorkflowDagRequest
//
// @return UpdateWorkflowDagResponse
func (client *Client) UpdateWorkflowDag(request *UpdateWorkflowDagRequest) (_result *UpdateWorkflowDagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkflowDagResponse{}
	_body, _err := client.UpdateWorkflowDagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
