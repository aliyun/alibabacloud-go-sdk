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
	client.EndpointMap = map[string]*string{
		"cn-shanghai": dara.String("das.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("das"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a database instance to Database Autonomy Service (DAS).
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call DAS, you must set the region to cn-shanghai.
//
// @param request - AddHDMInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddHDMInstanceResponse
func (client *Client) AddHDMInstanceWithOptions(request *AddHDMInstanceRequest, runtime *dara.RuntimeOptions) (_result *AddHDMInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.FlushAccount) {
		query["FlushAccount"] = request.FlushAccount
	}

	if !dara.IsNil(request.InstanceAlias) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !dara.IsNil(request.InstanceArea) {
		query["InstanceArea"] = request.InstanceArea
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.Context) {
		query["__context"] = request.Context
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddHDMInstance"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddHDMInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a database instance to Database Autonomy Service (DAS).
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call DAS, you must set the region to cn-shanghai.
//
// @param request - AddHDMInstanceRequest
//
// @return AddHDMInstanceResponse
func (client *Client) AddHDMInstance(request *AddHDMInstanceRequest) (_result *AddHDMInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddHDMInstanceResponse{}
	_body, _err := client.AddHDMInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cache analysis task.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - You can call this operation to analyze the data structures of ApsaraDB for Redis and the following self-developed data structures of Tair: TairString, TairHash, TairGIS, TairBloom, TairDoc, TairCpc, and TairZset. Other self-developed Tair data structures are not supported.
//
//   - If the specifications of the database instance that you want to analyze are changed, the backup file generated before the specification change cannot be analyzed.
//
//   - Tair ESSD/SSD-based instances are not supported.
//
// @param request - CreateCacheAnalysisJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCacheAnalysisJobResponse
func (client *Client) CreateCacheAnalysisJobWithOptions(request *CreateCacheAnalysisJobRequest, runtime *dara.RuntimeOptions) (_result *CreateCacheAnalysisJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Separators) {
		query["Separators"] = request.Separators
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCacheAnalysisJob"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCacheAnalysisJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cache analysis task.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - You can call this operation to analyze the data structures of ApsaraDB for Redis and the following self-developed data structures of Tair: TairString, TairHash, TairGIS, TairBloom, TairDoc, TairCpc, and TairZset. Other self-developed Tair data structures are not supported.
//
//   - If the specifications of the database instance that you want to analyze are changed, the backup file generated before the specification change cannot be analyzed.
//
//   - Tair ESSD/SSD-based instances are not supported.
//
// @param request - CreateCacheAnalysisJobRequest
//
// @return CreateCacheAnalysisJobResponse
func (client *Client) CreateCacheAnalysisJob(request *CreateCacheAnalysisJobRequest) (_result *CreateCacheAnalysisJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCacheAnalysisJobResponse{}
	_body, _err := client.CreateCacheAnalysisJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates stress testing tasks.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to effectively handle traffic spikes. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html). Before you call this API operation, make sure that your database instances meet the following requirements:
//
//   - The source database instance is an ApsaraDB RDS for MySQL High-availability Edition or Enterprise Edition instance, or a PolarDB for MySQL Cluster Edition cluster.
//
//   - The destination database instance is an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster.
//
//   - The source and destination database instances are connected to DAS. For information about how to connect database instances to DAS, see [Connect an Alibaba Cloud database instance to DAS](https://help.aliyun.com/document_detail/65405.html).
//
//   - DAS Enterprise Edition is enabled for the source and destination database instances. For more information, see [Overview](https://help.aliyun.com/document_detail/190912.html).
//
// @param request - CreateCloudBenchTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudBenchTasksResponse
func (client *Client) CreateCloudBenchTasksWithOptions(request *CreateCloudBenchTasksRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudBenchTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupTime) {
		query["BackupTime"] = request.BackupTime
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DstConnectionString) {
		query["DstConnectionString"] = request.DstConnectionString
	}

	if !dara.IsNil(request.DstInstanceId) {
		query["DstInstanceId"] = request.DstInstanceId
	}

	if !dara.IsNil(request.DstPort) {
		query["DstPort"] = request.DstPort
	}

	if !dara.IsNil(request.DstSuperAccount) {
		query["DstSuperAccount"] = request.DstSuperAccount
	}

	if !dara.IsNil(request.DstSuperPassword) {
		query["DstSuperPassword"] = request.DstSuperPassword
	}

	if !dara.IsNil(request.DstType) {
		query["DstType"] = request.DstType
	}

	if !dara.IsNil(request.DtsJobClass) {
		query["DtsJobClass"] = request.DtsJobClass
	}

	if !dara.IsNil(request.DtsJobId) {
		query["DtsJobId"] = request.DtsJobId
	}

	if !dara.IsNil(request.EndState) {
		query["EndState"] = request.EndState
	}

	if !dara.IsNil(request.GatewayVpcId) {
		query["GatewayVpcId"] = request.GatewayVpcId
	}

	if !dara.IsNil(request.GatewayVpcIp) {
		query["GatewayVpcIp"] = request.GatewayVpcIp
	}

	if !dara.IsNil(request.Rate) {
		query["Rate"] = request.Rate
	}

	if !dara.IsNil(request.RequestDuration) {
		query["RequestDuration"] = request.RequestDuration
	}

	if !dara.IsNil(request.RequestEndTime) {
		query["RequestEndTime"] = request.RequestEndTime
	}

	if !dara.IsNil(request.RequestStartTime) {
		query["RequestStartTime"] = request.RequestStartTime
	}

	if !dara.IsNil(request.SmartPressureTime) {
		query["SmartPressureTime"] = request.SmartPressureTime
	}

	if !dara.IsNil(request.SrcInstanceId) {
		query["SrcInstanceId"] = request.SrcInstanceId
	}

	if !dara.IsNil(request.SrcPublicIp) {
		query["SrcPublicIp"] = request.SrcPublicIp
	}

	if !dara.IsNil(request.SrcSuperAccount) {
		query["SrcSuperAccount"] = request.SrcSuperAccount
	}

	if !dara.IsNil(request.SrcSuperPassword) {
		query["SrcSuperPassword"] = request.SrcSuperPassword
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.WorkDir) {
		query["WorkDir"] = request.WorkDir
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudBenchTasks"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudBenchTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates stress testing tasks.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to effectively handle traffic spikes. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html). Before you call this API operation, make sure that your database instances meet the following requirements:
//
//   - The source database instance is an ApsaraDB RDS for MySQL High-availability Edition or Enterprise Edition instance, or a PolarDB for MySQL Cluster Edition cluster.
//
//   - The destination database instance is an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster.
//
//   - The source and destination database instances are connected to DAS. For information about how to connect database instances to DAS, see [Connect an Alibaba Cloud database instance to DAS](https://help.aliyun.com/document_detail/65405.html).
//
//   - DAS Enterprise Edition is enabled for the source and destination database instances. For more information, see [Overview](https://help.aliyun.com/document_detail/190912.html).
//
// @param request - CreateCloudBenchTasksRequest
//
// @return CreateCloudBenchTasksResponse
func (client *Client) CreateCloudBenchTasks(request *CreateCloudBenchTasksRequest) (_result *CreateCloudBenchTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudBenchTasksResponse{}
	_body, _err := client.CreateCloudBenchTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a diagnostic report.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than 4.3.3. We recommend that you use the latest version.
//
//   - The version of Database Autonomy Service (DAS) SDK must be 1.0.3 or later.
//
//   - If you use an SDK to call DAS, you must set the region to cn-shanghai.
//
//   - This operation supports the following database engines:
//
//   - RDS MySQL
//
//   - PolarDB for MySQL
//
//   - Redis
//
// @param request - CreateDiagnosticReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiagnosticReportResponse
func (client *Client) CreateDiagnosticReportWithOptions(request *CreateDiagnosticReportRequest, runtime *dara.RuntimeOptions) (_result *CreateDiagnosticReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDiagnosticReport"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDiagnosticReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a diagnostic report.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than 4.3.3. We recommend that you use the latest version.
//
//   - The version of Database Autonomy Service (DAS) SDK must be 1.0.3 or later.
//
//   - If you use an SDK to call DAS, you must set the region to cn-shanghai.
//
//   - This operation supports the following database engines:
//
//   - RDS MySQL
//
//   - PolarDB for MySQL
//
//   - Redis
//
// @param request - CreateDiagnosticReportRequest
//
// @return CreateDiagnosticReportResponse
func (client *Client) CreateDiagnosticReport(request *CreateDiagnosticReportRequest) (_result *CreateDiagnosticReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDiagnosticReportResponse{}
	_body, _err := client.CreateDiagnosticReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task that terminates sessions.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
//
//		- If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - CreateKillInstanceSessionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKillInstanceSessionTaskResponse
func (client *Client) CreateKillInstanceSessionTaskWithOptions(request *CreateKillInstanceSessionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateKillInstanceSessionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbUser) {
		query["DbUser"] = request.DbUser
	}

	if !dara.IsNil(request.DbUserPassword) {
		query["DbUserPassword"] = request.DbUserPassword
	}

	if !dara.IsNil(request.IgnoredUsers) {
		query["IgnoredUsers"] = request.IgnoredUsers
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.KillAllSessions) {
		query["KillAllSessions"] = request.KillAllSessions
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.SessionIds) {
		query["SessionIds"] = request.SessionIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKillInstanceSessionTask"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKillInstanceSessionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task that terminates sessions.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
//
//		- If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - CreateKillInstanceSessionTaskRequest
//
// @return CreateKillInstanceSessionTaskResponse
func (client *Client) CreateKillInstanceSessionTask(request *CreateKillInstanceSessionTaskRequest) (_result *CreateKillInstanceSessionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateKillInstanceSessionTaskResponse{}
	_body, _err := client.CreateKillInstanceSessionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建结束会话的任务
//
// @param request - CreateKillInstanceSessionTaskWithMaintainUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKillInstanceSessionTaskWithMaintainUserResponse
func (client *Client) CreateKillInstanceSessionTaskWithMaintainUserWithOptions(request *CreateKillInstanceSessionTaskWithMaintainUserRequest, runtime *dara.RuntimeOptions) (_result *CreateKillInstanceSessionTaskWithMaintainUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IgnoredUsers) {
		query["IgnoredUsers"] = request.IgnoredUsers
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.KillAllSessions) {
		query["KillAllSessions"] = request.KillAllSessions
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.SessionIds) {
		query["SessionIds"] = request.SessionIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKillInstanceSessionTaskWithMaintainUser"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKillInstanceSessionTaskWithMaintainUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建结束会话的任务
//
// @param request - CreateKillInstanceSessionTaskWithMaintainUserRequest
//
// @return CreateKillInstanceSessionTaskWithMaintainUserResponse
func (client *Client) CreateKillInstanceSessionTaskWithMaintainUser(request *CreateKillInstanceSessionTaskWithMaintainUserRequest) (_result *CreateKillInstanceSessionTaskWithMaintainUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateKillInstanceSessionTaskWithMaintainUserResponse{}
	_body, _err := client.CreateKillInstanceSessionTaskWithMaintainUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建最近死锁分析任务
//
// @param request - CreateLatestDeadLockAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLatestDeadLockAnalysisResponse
func (client *Client) CreateLatestDeadLockAnalysisWithOptions(request *CreateLatestDeadLockAnalysisRequest, runtime *dara.RuntimeOptions) (_result *CreateLatestDeadLockAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLatestDeadLockAnalysis"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLatestDeadLockAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建最近死锁分析任务
//
// @param request - CreateLatestDeadLockAnalysisRequest
//
// @return CreateLatestDeadLockAnalysisResponse
func (client *Client) CreateLatestDeadLockAnalysis(request *CreateLatestDeadLockAnalysisRequest) (_result *CreateLatestDeadLockAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLatestDeadLockAnalysisResponse{}
	_body, _err := client.CreateLatestDeadLockAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a tag to a SQL template.
//
// Description:
//
//	  If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - CreateQueryOptimizeTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueryOptimizeTagResponse
func (client *Client) CreateQueryOptimizeTagWithOptions(request *CreateQueryOptimizeTagRequest, runtime *dara.RuntimeOptions) (_result *CreateQueryOptimizeTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comments) {
		query["Comments"] = request.Comments
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SqlIds) {
		query["SqlIds"] = request.SqlIds
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQueryOptimizeTag"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQueryOptimizeTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a tag to a SQL template.
//
// Description:
//
//	  If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - CreateQueryOptimizeTagRequest
//
// @return CreateQueryOptimizeTagResponse
func (client *Client) CreateQueryOptimizeTag(request *CreateQueryOptimizeTagRequest) (_result *CreateQueryOptimizeTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQueryOptimizeTagResponse{}
	_body, _err := client.CreateQueryOptimizeTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates an SQL statement diagnostics request.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call Database Autonomy Service (DAS), you must set the region to cn-shanghai.
//
//   - This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - ApsaraDB RDS for PostgreSQL
//
//   - ApsaraDB RDS for SQL Server
//
//   - PolarDB for MySQL
//
//   - PolarDB for PostgreSQL (compatible with Oracle)
//
//   - ApsaraDB for MongoDB
//
// >  The minor engine version of ApsaraDB RDS for PostgreSQL instances must be 20221230 or later. For more information about how to check and update the minor engine version of an ApsaraDB RDS for PostgreSQL instance, see [Update the minor engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146895.html).
//
// @param request - CreateRequestDiagnosisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRequestDiagnosisResponse
func (client *Client) CreateRequestDiagnosisWithOptions(request *CreateRequestDiagnosisRequest, runtime *dara.RuntimeOptions) (_result *CreateRequestDiagnosisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Sql) {
		query["Sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRequestDiagnosis"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRequestDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an SQL statement diagnostics request.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call Database Autonomy Service (DAS), you must set the region to cn-shanghai.
//
//   - This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - ApsaraDB RDS for PostgreSQL
//
//   - ApsaraDB RDS for SQL Server
//
//   - PolarDB for MySQL
//
//   - PolarDB for PostgreSQL (compatible with Oracle)
//
//   - ApsaraDB for MongoDB
//
// >  The minor engine version of ApsaraDB RDS for PostgreSQL instances must be 20221230 or later. For more information about how to check and update the minor engine version of an ApsaraDB RDS for PostgreSQL instance, see [Update the minor engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146895.html).
//
// @param request - CreateRequestDiagnosisRequest
//
// @return CreateRequestDiagnosisResponse
func (client *Client) CreateRequestDiagnosis(request *CreateRequestDiagnosisRequest) (_result *CreateRequestDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRequestDiagnosisResponse{}
	_body, _err := client.CreateRequestDiagnosisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an offline task for Database Autonomy Service (DAS) Enterprise Edition.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - You can create an offline task only for database instances for which DAS Enterprise Edition V2 or V3 is enabled. For more information about the databases and regions that are supported by various versions of DAS Enterprise Edition, see [Editions and supported features](https://help.aliyun.com/document_detail/156204.html).
//
// @param request - CreateSqlLogTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSqlLogTaskResponse
func (client *Client) CreateSqlLogTaskWithOptions(request *CreateSqlLogTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSqlLogTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSqlLogTask"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSqlLogTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an offline task for Database Autonomy Service (DAS) Enterprise Edition.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - You can create an offline task only for database instances for which DAS Enterprise Edition V2 or V3 is enabled. For more information about the databases and regions that are supported by various versions of DAS Enterprise Edition, see [Editions and supported features](https://help.aliyun.com/document_detail/156204.html).
//
// @param request - CreateSqlLogTaskRequest
//
// @return CreateSqlLogTaskResponse
func (client *Client) CreateSqlLogTask(request *CreateSqlLogTaskRequest) (_result *CreateSqlLogTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSqlLogTaskResponse{}
	_body, _err := client.CreateSqlLogTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a storage analysis task to query the usage details of one or more databases and tables.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for MySQL instances, PolarDB for MySQL clusters, and ApsaraDB for MongoDB instances.
//
//		- For ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters, this operation works the same as the storage analysis feature of the previous version. Tasks generated by this operation cannot be viewed on the Storage Analysis page of the new version in the Database Autonomy Service (DAS) console. If you want to view the tasks and results, call the related API operation to obtain data and save data to your computer.
//
//		- If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - CreateStorageAnalysisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStorageAnalysisTaskResponse
func (client *Client) CreateStorageAnalysisTaskWithOptions(request *CreateStorageAnalysisTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateStorageAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStorageAnalysisTask"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStorageAnalysisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a storage analysis task to query the usage details of one or more databases and tables.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for MySQL instances, PolarDB for MySQL clusters, and ApsaraDB for MongoDB instances.
//
//		- For ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters, this operation works the same as the storage analysis feature of the previous version. Tasks generated by this operation cannot be viewed on the Storage Analysis page of the new version in the Database Autonomy Service (DAS) console. If you want to view the tasks and results, call the related API operation to obtain data and save data to your computer.
//
//		- If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - CreateStorageAnalysisTaskRequest
//
// @return CreateStorageAnalysisTaskResponse
func (client *Client) CreateStorageAnalysisTask(request *CreateStorageAnalysisTaskRequest) (_result *CreateStorageAnalysisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStorageAnalysisTaskResponse{}
	_body, _err := client.CreateStorageAnalysisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a stress testing task.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to handle traffic spikes in an effective manner. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - DeleteCloudBenchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudBenchTaskResponse
func (client *Client) DeleteCloudBenchTaskWithOptions(request *DeleteCloudBenchTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudBenchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudBenchTask"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudBenchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a stress testing task.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to handle traffic spikes in an effective manner. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - DeleteCloudBenchTaskRequest
//
// @return DeleteCloudBenchTaskResponse
func (client *Client) DeleteCloudBenchTask(request *DeleteCloudBenchTaskRequest) (_result *DeleteCloudBenchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudBenchTaskResponse{}
	_body, _err := client.DeleteCloudBenchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the metadata of a stopped DBGateway.
//
// Description:
//
//	  This operation is used to delete the metadata of a DBGateway that is released in a stress testing task created by calling the [CreateCloudBenchTasks](https://help.aliyun.com/document_detail/230665.html) operation.
//
//		- If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// >  If the heartbeat is lost between a DBGateway and the access point for more than 20 seconds, the DBGateway is considered stopped.
//
// @param request - DeleteStopGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStopGatewayResponse
func (client *Client) DeleteStopGatewayWithOptions(request *DeleteStopGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteStopGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStopGateway"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStopGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the metadata of a stopped DBGateway.
//
// Description:
//
//	  This operation is used to delete the metadata of a DBGateway that is released in a stress testing task created by calling the [CreateCloudBenchTasks](https://help.aliyun.com/document_detail/230665.html) operation.
//
//		- If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// >  If the heartbeat is lost between a DBGateway and the access point for more than 20 seconds, the DBGateway is considered stopped.
//
// @param request - DeleteStopGatewayRequest
//
// @return DeleteStopGatewayResponse
func (client *Client) DeleteStopGateway(request *DeleteStopGatewayRequest) (_result *DeleteStopGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStopGatewayResponse{}
	_body, _err := client.DeleteStopGatewayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of the auto scaling feature for an instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoScalingConfigResponse
func (client *Client) DescribeAutoScalingConfigWithOptions(request *DescribeAutoScalingConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoScalingConfigResponse, _err error) {
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
		Action:      dara.String("DescribeAutoScalingConfig"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoScalingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of the auto scaling feature for an instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeAutoScalingConfigRequest
//
// @return DescribeAutoScalingConfigResponse
func (client *Client) DescribeAutoScalingConfig(request *DescribeAutoScalingConfigRequest) (_result *DescribeAutoScalingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoScalingConfigResponse{}
	_body, _err := client.DescribeAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the auto scaling history of an instance.
//
// Description:
//
//	  You can call this operation to query the history information about the automatic performance scaling only of ApsaraDB RDS for MySQL High-availability Edition instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeAutoScalingHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoScalingHistoryResponse
func (client *Client) DescribeAutoScalingHistoryWithOptions(request *DescribeAutoScalingHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoScalingHistoryResponse, _err error) {
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
		Action:      dara.String("DescribeAutoScalingHistory"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoScalingHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the auto scaling history of an instance.
//
// Description:
//
//	  You can call this operation to query the history information about the automatic performance scaling only of ApsaraDB RDS for MySQL High-availability Edition instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeAutoScalingHistoryRequest
//
// @return DescribeAutoScalingHistoryResponse
func (client *Client) DescribeAutoScalingHistory(request *DescribeAutoScalingHistoryRequest) (_result *DescribeAutoScalingHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAutoScalingHistoryResponse{}
	_body, _err := client.DescribeAutoScalingHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a cache analysis task.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable only to ApsaraDB for Redis.
//
// >  You can call this operation to query the top 500 keys in a cache analysis task.
//
// @param request - DescribeCacheAnalysisJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCacheAnalysisJobResponse
func (client *Client) DescribeCacheAnalysisJobWithOptions(request *DescribeCacheAnalysisJobRequest, runtime *dara.RuntimeOptions) (_result *DescribeCacheAnalysisJobResponse, _err error) {
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

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCacheAnalysisJob"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCacheAnalysisJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a cache analysis task.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable only to ApsaraDB for Redis.
//
// >  You can call this operation to query the top 500 keys in a cache analysis task.
//
// @param request - DescribeCacheAnalysisJobRequest
//
// @return DescribeCacheAnalysisJobResponse
func (client *Client) DescribeCacheAnalysisJob(request *DescribeCacheAnalysisJobRequest) (_result *DescribeCacheAnalysisJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCacheAnalysisJobResponse{}
	_body, _err := client.DescribeCacheAnalysisJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of cache analysis tasks.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable only to ApsaraDB for Redis.
//
// @param request - DescribeCacheAnalysisJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCacheAnalysisJobsResponse
func (client *Client) DescribeCacheAnalysisJobsWithOptions(request *DescribeCacheAnalysisJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCacheAnalysisJobsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCacheAnalysisJobs"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCacheAnalysisJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of cache analysis tasks.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable only to ApsaraDB for Redis.
//
// @param request - DescribeCacheAnalysisJobsRequest
//
// @return DescribeCacheAnalysisJobsResponse
func (client *Client) DescribeCacheAnalysisJobs(request *DescribeCacheAnalysisJobsRequest) (_result *DescribeCacheAnalysisJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCacheAnalysisJobsResponse{}
	_body, _err := client.DescribeCacheAnalysisJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries stress testing tasks.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to effectively handle traffic spikes. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - DescribeCloudBenchTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudBenchTasksResponse
func (client *Client) DescribeCloudBenchTasksWithOptions(request *DescribeCloudBenchTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudBenchTasksResponse, _err error) {
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

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudBenchTasks"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudBenchTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries stress testing tasks.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to effectively handle traffic spikes. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - DescribeCloudBenchTasksRequest
//
// @return DescribeCloudBenchTasksResponse
func (client *Client) DescribeCloudBenchTasks(request *DescribeCloudBenchTasksRequest) (_result *DescribeCloudBenchTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudBenchTasksResponse{}
	_body, _err := client.DescribeCloudBenchTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a stress testing task.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether you need to scale up your database instance to handle workloads during peak hours. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - DescribeCloudbenchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudbenchTaskResponse
func (client *Client) DescribeCloudbenchTaskWithOptions(request *DescribeCloudbenchTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudbenchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudbenchTask"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudbenchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a stress testing task.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether you need to scale up your database instance to handle workloads during peak hours. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - DescribeCloudbenchTaskRequest
//
// @return DescribeCloudbenchTaskResponse
func (client *Client) DescribeCloudbenchTask(request *DescribeCloudbenchTaskRequest) (_result *DescribeCloudbenchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudbenchTaskResponse{}
	_body, _err := client.DescribeCloudbenchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a stress testing task.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to effectively handle traffic spikes. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - DescribeCloudbenchTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudbenchTaskConfigResponse
func (client *Client) DescribeCloudbenchTaskConfigWithOptions(request *DescribeCloudbenchTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudbenchTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudbenchTaskConfig"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudbenchTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a stress testing task.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to effectively handle traffic spikes. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - DescribeCloudbenchTaskConfigRequest
//
// @return DescribeCloudbenchTaskConfigResponse
func (client *Client) DescribeCloudbenchTaskConfig(request *DescribeCloudbenchTaskConfigRequest) (_result *DescribeCloudbenchTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudbenchTaskConfigResponse{}
	_body, _err := client.DescribeCloudbenchTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries diagnostics reports.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable to the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
//   - ApsaraDB for Redis
//
// @param request - DescribeDiagnosticReportListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosticReportListResponse
func (client *Client) DescribeDiagnosticReportListWithOptions(request *DescribeDiagnosticReportListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosticReportListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosticReportList"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDiagnosticReportListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries diagnostics reports.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable to the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
//   - ApsaraDB for Redis
//
// @param request - DescribeDiagnosticReportListRequest
//
// @return DescribeDiagnosticReportListResponse
func (client *Client) DescribeDiagnosticReportList(request *DescribeDiagnosticReportListRequest) (_result *DescribeDiagnosticReportListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDiagnosticReportListResponse{}
	_body, _err := client.DescribeDiagnosticReportListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例错误日志
//
// @param request - DescribeErrorLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeErrorLogRecordsResponse
func (client *Client) DescribeErrorLogRecordsWithOptions(request *DescribeErrorLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeErrorLogRecordsResponse, _err error) {
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

	if !dara.IsNil(request.Filters) {
		body["Filters"] = request.Filters
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Role) {
		body["Role"] = request.Role
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeErrorLogRecords"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeErrorLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例错误日志
//
// @param request - DescribeErrorLogRecordsRequest
//
// @return DescribeErrorLogRecordsResponse
func (client *Client) DescribeErrorLogRecords(request *DescribeErrorLogRecordsRequest) (_result *DescribeErrorLogRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeErrorLogRecordsResponse{}
	_body, _err := client.DescribeErrorLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hot keys and the large keys in the memory in real time.
//
// Description:
//
// This operation sorts list, hash, set, and zset keys based on the number of elements contained in these keys. The top three keys that contain the most elements are considered large keys. If the number of queries per second (QPS) of a key is greater than 3,000, the key is considered a hot key.
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than 4.3.3. We recommend that you use the latest version.
//
//   - The version of Database Autonomy Service (DAS) SDK must be 1.0.2 or later.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is available only for ApsaraDB for Redis instances that meet the following requirements:
//
//   - The instance is a Community Edition instance that uses a major version of 5.0 or later or a performance-enhanced instance of the Enhanced Edition (Tair).
//
//   - The ApsaraDB for Redis instance is updated to the latest minor version.
//
// @param request - DescribeHotBigKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHotBigKeysResponse
func (client *Client) DescribeHotBigKeysWithOptions(request *DescribeHotBigKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeHotBigKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHotBigKeys"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHotBigKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hot keys and the large keys in the memory in real time.
//
// Description:
//
// This operation sorts list, hash, set, and zset keys based on the number of elements contained in these keys. The top three keys that contain the most elements are considered large keys. If the number of queries per second (QPS) of a key is greater than 3,000, the key is considered a hot key.
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than 4.3.3. We recommend that you use the latest version.
//
//   - The version of Database Autonomy Service (DAS) SDK must be 1.0.2 or later.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is available only for ApsaraDB for Redis instances that meet the following requirements:
//
//   - The instance is a Community Edition instance that uses a major version of 5.0 or later or a performance-enhanced instance of the Enhanced Edition (Tair).
//
//   - The ApsaraDB for Redis instance is updated to the latest minor version.
//
// @param request - DescribeHotBigKeysRequest
//
// @return DescribeHotBigKeysResponse
func (client *Client) DescribeHotBigKeys(request *DescribeHotBigKeysRequest) (_result *DescribeHotBigKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHotBigKeysResponse{}
	_body, _err := client.DescribeHotBigKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hot keys of an ApsaraDB for Redis instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V4.3.3. We recommend that you use the latest version.
//
//   - The version of your Database Autonomy Service (DAS) SDK must be V1.0.2 or later.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable only to ApsaraDB for Redis instances that meet the following requirements:
//
//   - The ApsaraDB for Redis instance is a Community Edition instance that uses a major version of 4.0 or later or a performance-enhanced instance of the Enhanced Edition (Tair).
//
//   - The ApsaraDB for Redis instance is updated to the latest minor version.
//
// @param request - DescribeHotKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHotKeysResponse
func (client *Client) DescribeHotKeysWithOptions(request *DescribeHotKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeHotKeysResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHotKeys"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHotKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hot keys of an ApsaraDB for Redis instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V4.3.3. We recommend that you use the latest version.
//
//   - The version of your Database Autonomy Service (DAS) SDK must be V1.0.2 or later.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable only to ApsaraDB for Redis instances that meet the following requirements:
//
//   - The ApsaraDB for Redis instance is a Community Edition instance that uses a major version of 4.0 or later or a performance-enhanced instance of the Enhanced Edition (Tair).
//
//   - The ApsaraDB for Redis instance is updated to the latest minor version.
//
// @param request - DescribeHotKeysRequest
//
// @return DescribeHotKeysResponse
func (client *Client) DescribeHotKeys(request *DescribeHotKeysRequest) (_result *DescribeHotKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHotKeysResponse{}
	_body, _err := client.DescribeHotKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Database Autonomy Service (DAS) Enterprise Edition V1 or V2 is enabled for a database instance.
//
// Description:
//
//	  For more information about the database instances that support DAS Enterprise Edition, see [Overview of DAS Enterprise Edition](https://help.aliyun.com/document_detail/190912.html).
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation is applicable only to DAS Enterprise Edition V1 and V2.
//
// >  We recommend that you call the [DescribeSqlLogConfig](https://help.aliyun.com/document_detail/2778837.html) operation to query the DAS Enterprise Edition configurations of a database instance.
//
// @param request - DescribeInstanceDasProRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDasProResponse
func (client *Client) DescribeInstanceDasProWithOptions(request *DescribeInstanceDasProRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDasProResponse, _err error) {
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
		Action:      dara.String("DescribeInstanceDasPro"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDasProResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Database Autonomy Service (DAS) Enterprise Edition V1 or V2 is enabled for a database instance.
//
// Description:
//
//	  For more information about the database instances that support DAS Enterprise Edition, see [Overview of DAS Enterprise Edition](https://help.aliyun.com/document_detail/190912.html).
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation is applicable only to DAS Enterprise Edition V1 and V2.
//
// >  We recommend that you call the [DescribeSqlLogConfig](https://help.aliyun.com/document_detail/2778837.html) operation to query the DAS Enterprise Edition configurations of a database instance.
//
// @param request - DescribeInstanceDasProRequest
//
// @return DescribeInstanceDasProResponse
func (client *Client) DescribeInstanceDasPro(request *DescribeInstanceDasProRequest) (_result *DescribeInstanceDasProResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceDasProResponse{}
	_body, _err := client.DescribeInstanceDasProWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取执行计划
//
// @param request - DescribeQueryExplainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQueryExplainResponse
func (client *Client) DescribeQueryExplainWithOptions(request *DescribeQueryExplainRequest, runtime *dara.RuntimeOptions) (_result *DescribeQueryExplainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		body["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Schema) {
		body["Schema"] = request.Schema
	}

	if !dara.IsNil(request.Sql) {
		body["Sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQueryExplain"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQueryExplainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取执行计划
//
// @param request - DescribeQueryExplainRequest
//
// @return DescribeQueryExplainResponse
func (client *Client) DescribeQueryExplain(request *DescribeQueryExplainRequest) (_result *DescribeQueryExplainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeQueryExplainResponse{}
	_body, _err := client.DescribeQueryExplainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeSlowLogHistogramAsync
//
// @param request - DescribeSlowLogHistogramAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogHistogramAsyncResponse
func (client *Client) DescribeSlowLogHistogramAsyncWithOptions(request *DescribeSlowLogHistogramAsyncRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogHistogramAsyncResponse, _err error) {
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

	if !dara.IsNil(request.Filters) {
		body["Filters"] = request.Filters
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlowLogHistogramAsync"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowLogHistogramAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeSlowLogHistogramAsync
//
// @param request - DescribeSlowLogHistogramAsyncRequest
//
// @return DescribeSlowLogHistogramAsyncResponse
func (client *Client) DescribeSlowLogHistogramAsync(request *DescribeSlowLogHistogramAsyncRequest) (_result *DescribeSlowLogHistogramAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlowLogHistogramAsyncResponse{}
	_body, _err := client.DescribeSlowLogHistogramAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看慢日志明细接口
//
// @param request - DescribeSlowLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Asc) {
		query["Asc"] = request.Asc
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filters) {
		body["Filters"] = request.Filters
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlowLogRecords"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看慢日志明细接口
//
// @param request - DescribeSlowLogRecordsRequest
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 慢日志统计信息
//
// @param request - DescribeSlowLogStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogStatisticResponse
func (client *Client) DescribeSlowLogStatisticWithOptions(request *DescribeSlowLogStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Asc) {
		body["Asc"] = request.Asc
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Filters) {
		body["Filters"] = request.Filters
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderBy) {
		body["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlowLogStatistic"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowLogStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 慢日志统计信息
//
// @param request - DescribeSlowLogStatisticRequest
//
// @return DescribeSlowLogStatisticResponse
func (client *Client) DescribeSlowLogStatistic(request *DescribeSlowLogStatisticRequest) (_result *DescribeSlowLogStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlowLogStatisticResponse{}
	_body, _err := client.DescribeSlowLogStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of Database Autonomy Service (DAS) Enterprise Edition that is enabled for a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlLogConfigResponse
func (client *Client) DescribeSqlLogConfigWithOptions(request *DescribeSqlLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSqlLogConfig"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSqlLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of Database Autonomy Service (DAS) Enterprise Edition that is enabled for a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogConfigRequest
//
// @return DescribeSqlLogConfigResponse
func (client *Client) DescribeSqlLogConfig(request *DescribeSqlLogConfigRequest) (_result *DescribeSqlLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSqlLogConfigResponse{}
	_body, _err := client.DescribeSqlLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log details of a database instance for which Database Autonomy Service (DAS) Enterprise Edition is enabled.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlLogRecordsResponse
func (client *Client) DescribeSqlLogRecordsWithOptions(request *DescribeSqlLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlLogRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSqlLogRecords"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSqlLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log details of a database instance for which Database Autonomy Service (DAS) Enterprise Edition is enabled.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogRecordsRequest
//
// @return DescribeSqlLogRecordsResponse
func (client *Client) DescribeSqlLogRecords(request *DescribeSqlLogRecordsRequest) (_result *DescribeSqlLogRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSqlLogRecordsResponse{}
	_body, _err := client.DescribeSqlLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics of Database Autonomy Service (DAS) Enterprise Edition.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlLogStatisticResponse
func (client *Client) DescribeSqlLogStatisticWithOptions(request *DescribeSqlLogStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlLogStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSqlLogStatistic"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSqlLogStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of Database Autonomy Service (DAS) Enterprise Edition.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogStatisticRequest
//
// @return DescribeSqlLogStatisticResponse
func (client *Client) DescribeSqlLogStatistic(request *DescribeSqlLogStatisticRequest) (_result *DescribeSqlLogStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSqlLogStatisticResponse{}
	_body, _err := client.DescribeSqlLogStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an offline task in Database Autonomy Service (DAS) Enterprise Edition.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlLogTaskResponse
func (client *Client) DescribeSqlLogTaskWithOptions(request *DescribeSqlLogTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlLogTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSqlLogTask"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSqlLogTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an offline task in Database Autonomy Service (DAS) Enterprise Edition.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogTaskRequest
//
// @return DescribeSqlLogTaskResponse
func (client *Client) DescribeSqlLogTask(request *DescribeSqlLogTaskRequest) (_result *DescribeSqlLogTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSqlLogTaskResponse{}
	_body, _err := client.DescribeSqlLogTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the audit log tasks of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlLogTasksResponse
func (client *Client) DescribeSqlLogTasksWithOptions(request *DescribeSqlLogTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlLogTasksResponse, _err error) {
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

	if !dara.IsNil(request.Filters) {
		body["Filters"] = request.Filters
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSqlLogTasks"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSqlLogTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audit log tasks of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - DescribeSqlLogTasksRequest
//
// @return DescribeSqlLogTasksResponse
func (client *Client) DescribeSqlLogTasks(request *DescribeSqlLogTasksRequest) (_result *DescribeSqlLogTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSqlLogTasksResponse{}
	_body, _err := client.DescribeSqlLogTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 100 large keys over a period of time.
//
// Description:
//
// The list, hash, set, and zset keys are sorted based on the number of elements in these keys. The top three keys that have the most elements are considered large keys.
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than 4.3.3. We recommend that you use the latest version.
//
//   - The version of Database Autonomy Service (DAS) SDK must be 1.0.2 or later.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is available only for an ApsaraDB for Redis instance of one of the following versions:
//
//   - The instance is ApsaraDB for Redis Community Edition instances that use a major version of 5.0 or later or a performance-enhanced instance of the ApsaraDB for Redis Enhanced Edition (Tair).
//
//   - The ApsaraDB for Redis instance is updated to the latest minor version.
//
// @param request - DescribeTopBigKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTopBigKeysResponse
func (client *Client) DescribeTopBigKeysWithOptions(request *DescribeTopBigKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeTopBigKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTopBigKeys"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTopBigKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 100 large keys over a period of time.
//
// Description:
//
// The list, hash, set, and zset keys are sorted based on the number of elements in these keys. The top three keys that have the most elements are considered large keys.
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than 4.3.3. We recommend that you use the latest version.
//
//   - The version of Database Autonomy Service (DAS) SDK must be 1.0.2 or later.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is available only for an ApsaraDB for Redis instance of one of the following versions:
//
//   - The instance is ApsaraDB for Redis Community Edition instances that use a major version of 5.0 or later or a performance-enhanced instance of the ApsaraDB for Redis Enhanced Edition (Tair).
//
//   - The ApsaraDB for Redis instance is updated to the latest minor version.
//
// @param request - DescribeTopBigKeysRequest
//
// @return DescribeTopBigKeysResponse
func (client *Client) DescribeTopBigKeys(request *DescribeTopBigKeysRequest) (_result *DescribeTopBigKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTopBigKeysResponse{}
	_body, _err := client.DescribeTopBigKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 100 hotkeys over a period of time.
//
// Description:
//
// If the number of queries per second (QPS) of a key is greater than 3,000, the key is considered a hot key.
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than 4.3.3. We recommend that you use the latest version.
//
//   - The version of Database Autonomy Service (DAS) SDK must be 1.0.2 or later.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is available only for an ApsaraDB for Redis instance of one of the following versions:
//
//   - The instance is a Community Edition instance that uses a major version of 4.0 or later or a performance-enhanced instance of the Enhanced Edition (Tair).
//
//   - The ApsaraDB for Redis instance is updated to the latest minor version.
//
// @param request - DescribeTopHotKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTopHotKeysResponse
func (client *Client) DescribeTopHotKeysWithOptions(request *DescribeTopHotKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeTopHotKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTopHotKeys"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTopHotKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 100 hotkeys over a period of time.
//
// Description:
//
// If the number of queries per second (QPS) of a key is greater than 3,000, the key is considered a hot key.
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than 4.3.3. We recommend that you use the latest version.
//
//   - The version of Database Autonomy Service (DAS) SDK must be 1.0.2 or later.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is available only for an ApsaraDB for Redis instance of one of the following versions:
//
//   - The instance is a Community Edition instance that uses a major version of 4.0 or later or a performance-enhanced instance of the Enhanced Edition (Tair).
//
//   - The ApsaraDB for Redis instance is updated to the latest minor version.
//
// @param request - DescribeTopHotKeysRequest
//
// @return DescribeTopHotKeysResponse
func (client *Client) DescribeTopHotKeys(request *DescribeTopHotKeysRequest) (_result *DescribeTopHotKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTopHotKeysResponse{}
	_body, _err := client.DescribeTopHotKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables all throttling rules that are in effect.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - DisableAllSqlConcurrencyControlRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAllSqlConcurrencyControlRulesResponse
func (client *Client) DisableAllSqlConcurrencyControlRulesWithOptions(request *DisableAllSqlConcurrencyControlRulesRequest, runtime *dara.RuntimeOptions) (_result *DisableAllSqlConcurrencyControlRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAllSqlConcurrencyControlRules"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAllSqlConcurrencyControlRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables all throttling rules that are in effect.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - DisableAllSqlConcurrencyControlRulesRequest
//
// @return DisableAllSqlConcurrencyControlRulesResponse
func (client *Client) DisableAllSqlConcurrencyControlRules(request *DisableAllSqlConcurrencyControlRulesRequest) (_result *DisableAllSqlConcurrencyControlRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableAllSqlConcurrencyControlRulesResponse{}
	_body, _err := client.DisableAllSqlConcurrencyControlRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the automatic tablespace fragment recycling feature for database instances at a time.
//
// Description:
//
// If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// @param request - DisableAutoResourceOptimizeRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAutoResourceOptimizeRulesResponse
func (client *Client) DisableAutoResourceOptimizeRulesWithOptions(request *DisableAutoResourceOptimizeRulesRequest, runtime *dara.RuntimeOptions) (_result *DisableAutoResourceOptimizeRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAutoResourceOptimizeRules"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAutoResourceOptimizeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the automatic tablespace fragment recycling feature for database instances at a time.
//
// Description:
//
// If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// @param request - DisableAutoResourceOptimizeRulesRequest
//
// @return DisableAutoResourceOptimizeRulesResponse
func (client *Client) DisableAutoResourceOptimizeRules(request *DisableAutoResourceOptimizeRulesRequest) (_result *DisableAutoResourceOptimizeRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableAutoResourceOptimizeRulesResponse{}
	_body, _err := client.DisableAutoResourceOptimizeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the automatic SQL throttling feature for multiple database instances at a time.
//
// Description:
//
// If you use an SDK to call operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// @param request - DisableAutoThrottleRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAutoThrottleRulesResponse
func (client *Client) DisableAutoThrottleRulesWithOptions(request *DisableAutoThrottleRulesRequest, runtime *dara.RuntimeOptions) (_result *DisableAutoThrottleRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableAutoThrottleRules"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAutoThrottleRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the automatic SQL throttling feature for multiple database instances at a time.
//
// Description:
//
// If you use an SDK to call operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// @param request - DisableAutoThrottleRulesRequest
//
// @return DisableAutoThrottleRulesResponse
func (client *Client) DisableAutoThrottleRules(request *DisableAutoThrottleRulesRequest) (_result *DisableAutoThrottleRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableAutoThrottleRulesResponse{}
	_body, _err := client.DisableAutoThrottleRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deactivates Database Autonomy Service (DAS) Professional Edition.
//
// Description:
//
//	  For more information about the database instances that support DAS Enterprise Edition, see [Overview](https://help.aliyun.com/document_detail/190912.html).
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation is applicable only to DAS Enterprise Edition V1.
//
// >  We recommend that you call the [ModifySqlLogConfig](https://help.aliyun.com/document_detail/2778835.html) operation to enable or disable DAS Enterprise Edition for a database instance. For more information about the databases and regions supported by each version of DAS Enterprise Edition, see [Editions and supported features](https://help.aliyun.com/document_detail/156204.html).
//
// @param request - DisableDasProRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDasProResponse
func (client *Client) DisableDasProWithOptions(request *DisableDasProRequest, runtime *dara.RuntimeOptions) (_result *DisableDasProResponse, _err error) {
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

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDasPro"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDasProResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deactivates Database Autonomy Service (DAS) Professional Edition.
//
// Description:
//
//	  For more information about the database instances that support DAS Enterprise Edition, see [Overview](https://help.aliyun.com/document_detail/190912.html).
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation is applicable only to DAS Enterprise Edition V1.
//
// >  We recommend that you call the [ModifySqlLogConfig](https://help.aliyun.com/document_detail/2778835.html) operation to enable or disable DAS Enterprise Edition for a database instance. For more information about the databases and regions supported by each version of DAS Enterprise Edition, see [Editions and supported features](https://help.aliyun.com/document_detail/156204.html).
//
// @param request - DisableDasProRequest
//
// @return DisableDasProResponse
func (client *Client) DisableDasPro(request *DisableDasProRequest) (_result *DisableDasProResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableDasProResponse{}
	_body, _err := client.DisableDasProWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the auto scaling feature for a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable only to ApsaraDB for Redis instances.
//
// @param request - DisableInstanceDasConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInstanceDasConfigResponse
func (client *Client) DisableInstanceDasConfigWithOptions(request *DisableInstanceDasConfigRequest, runtime *dara.RuntimeOptions) (_result *DisableInstanceDasConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ScaleType) {
		query["ScaleType"] = request.ScaleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableInstanceDasConfig"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableInstanceDasConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the auto scaling feature for a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable only to ApsaraDB for Redis instances.
//
// @param request - DisableInstanceDasConfigRequest
//
// @return DisableInstanceDasConfigResponse
func (client *Client) DisableInstanceDasConfig(request *DisableInstanceDasConfigRequest) (_result *DisableInstanceDasConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableInstanceDasConfigResponse{}
	_body, _err := client.DisableInstanceDasConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a throttling rule.
//
// Description:
//
// This operation is applicable to the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - DisableSqlConcurrencyControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSqlConcurrencyControlResponse
func (client *Client) DisableSqlConcurrencyControlWithOptions(request *DisableSqlConcurrencyControlRequest, runtime *dara.RuntimeOptions) (_result *DisableSqlConcurrencyControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ItemId) {
		query["ItemId"] = request.ItemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSqlConcurrencyControl"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSqlConcurrencyControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a throttling rule.
//
// Description:
//
// This operation is applicable to the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - DisableSqlConcurrencyControlRequest
//
// @return DisableSqlConcurrencyControlResponse
func (client *Client) DisableSqlConcurrencyControl(request *DisableSqlConcurrencyControlRequest) (_result *DisableSqlConcurrencyControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableSqlConcurrencyControlResponse{}
	_body, _err := client.DisableSqlConcurrencyControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates Database Autonomy Service (DAS) Professional Edition.
//
// Description:
//
//	  If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation is applicable only to DAS Enterprise Edition V1.
//
// >  We recommend that you call the [ModifySqlLogConfig](https://help.aliyun.com/document_detail/2778835.html) operation to activate or deactivate DAS Enterprise Edition for a database instance. For more information about the databases and regions supported by each version of DAS Enterprise Edition, see [DAS editions and supported features](https://help.aliyun.com/document_detail/156204.html).
//
// @param request - EnableDasProRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDasProResponse
func (client *Client) EnableDasProWithOptions(request *EnableDasProRequest, runtime *dara.RuntimeOptions) (_result *EnableDasProResponse, _err error) {
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

	if !dara.IsNil(request.SqlRetention) {
		query["SqlRetention"] = request.SqlRetention
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDasPro"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDasProResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates Database Autonomy Service (DAS) Professional Edition.
//
// Description:
//
//	  If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation is applicable only to DAS Enterprise Edition V1.
//
// >  We recommend that you call the [ModifySqlLogConfig](https://help.aliyun.com/document_detail/2778835.html) operation to activate or deactivate DAS Enterprise Edition for a database instance. For more information about the databases and regions supported by each version of DAS Enterprise Edition, see [DAS editions and supported features](https://help.aliyun.com/document_detail/156204.html).
//
// @param request - EnableDasProRequest
//
// @return EnableDasProResponse
func (client *Client) EnableDasPro(request *EnableDasProRequest) (_result *EnableDasProResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableDasProResponse{}
	_body, _err := client.EnableDasProWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables SQL throttling to control the numbers of database access requests and concurrent SQL statements.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - EnableSqlConcurrencyControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSqlConcurrencyControlResponse
func (client *Client) EnableSqlConcurrencyControlWithOptions(request *EnableSqlConcurrencyControlRequest, runtime *dara.RuntimeOptions) (_result *EnableSqlConcurrencyControlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConcurrencyControlTime) {
		query["ConcurrencyControlTime"] = request.ConcurrencyControlTime
	}

	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxConcurrency) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.SqlKeywords) {
		query["SqlKeywords"] = request.SqlKeywords
	}

	if !dara.IsNil(request.SqlType) {
		query["SqlType"] = request.SqlType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSqlConcurrencyControl"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSqlConcurrencyControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables SQL throttling to control the numbers of database access requests and concurrent SQL statements.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - EnableSqlConcurrencyControlRequest
//
// @return EnableSqlConcurrencyControlResponse
func (client *Client) EnableSqlConcurrencyControl(request *EnableSqlConcurrencyControlRequest) (_result *EnableSqlConcurrencyControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableSqlConcurrencyControlResponse{}
	_body, _err := client.EnableSqlConcurrencyControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously queries the IDs of SQL statements that generate a MySQL error code in the SQL Explorer results of a database instance.
//
// Description:
//
// >  GetAsyncErrorRequestListByCode is an asynchronous operation. After a request is sent, the complete results are not returned immediately. If the value of the **isFinish*	- parameter is **false*	- in the response, wait for 1 second and then send a request again. If the value of the **isFinish*	- parameter is **true**, the complete results are returned.
//
//   - This API operation supports only ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters for which Database Autonomy Service (DAS) Enterprise Edition is enabled. For more information, see [Enable and manage DAS Economy Edition and DAS Enterprise Edition](https://help.aliyun.com/document_detail/163298.html).
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - When you call this operation, the value of the SqlId parameter changes due to the optimization of the SQL template algorithm starting from September 1, 2024. For more information, see [[Notice\\] Optimization of the SQL template algorithm](~~2845725~~).
//
// @param request - GetAsyncErrorRequestListByCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncErrorRequestListByCodeResponse
func (client *Client) GetAsyncErrorRequestListByCodeWithOptions(request *GetAsyncErrorRequestListByCodeRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncErrorRequestListByCodeResponse, _err error) {
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

	if !dara.IsNil(request.ErrorCode) {
		query["ErrorCode"] = request.ErrorCode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncErrorRequestListByCode"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncErrorRequestListByCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously queries the IDs of SQL statements that generate a MySQL error code in the SQL Explorer results of a database instance.
//
// Description:
//
// >  GetAsyncErrorRequestListByCode is an asynchronous operation. After a request is sent, the complete results are not returned immediately. If the value of the **isFinish*	- parameter is **false*	- in the response, wait for 1 second and then send a request again. If the value of the **isFinish*	- parameter is **true**, the complete results are returned.
//
//   - This API operation supports only ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters for which Database Autonomy Service (DAS) Enterprise Edition is enabled. For more information, see [Enable and manage DAS Economy Edition and DAS Enterprise Edition](https://help.aliyun.com/document_detail/163298.html).
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - When you call this operation, the value of the SqlId parameter changes due to the optimization of the SQL template algorithm starting from September 1, 2024. For more information, see [[Notice\\] Optimization of the SQL template algorithm](~~2845725~~).
//
// @param request - GetAsyncErrorRequestListByCodeRequest
//
// @return GetAsyncErrorRequestListByCodeResponse
func (client *Client) GetAsyncErrorRequestListByCode(request *GetAsyncErrorRequestListByCodeRequest) (_result *GetAsyncErrorRequestListByCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAsyncErrorRequestListByCodeResponse{}
	_body, _err := client.GetAsyncErrorRequestListByCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously queries the MySQL error codes in SQL Explorer data and the number of SQL queries corresponding to each error code.
//
// Description:
//
// >  GetAsyncErrorRequestStatByCode is an asynchronous operation After a request is sent, the complete results are not returned immediately. If the value of **isFinish*	- is **false*	- in the response, wait for 1 second and then send a request again. If the value of **isFinish*	- is **true**, the complete results are returned.
//
//   - This API operation supports only ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters for which Database Autonomy Service (DAS) Enterprise Edition is enabled. For more information, see [Purchase DAS Enterprise Edition](https://help.aliyun.com/document_detail/163298.html).
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetAsyncErrorRequestStatByCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncErrorRequestStatByCodeResponse
func (client *Client) GetAsyncErrorRequestStatByCodeWithOptions(request *GetAsyncErrorRequestStatByCodeRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncErrorRequestStatByCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncErrorRequestStatByCode"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncErrorRequestStatByCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously queries the MySQL error codes in SQL Explorer data and the number of SQL queries corresponding to each error code.
//
// Description:
//
// >  GetAsyncErrorRequestStatByCode is an asynchronous operation After a request is sent, the complete results are not returned immediately. If the value of **isFinish*	- is **false*	- in the response, wait for 1 second and then send a request again. If the value of **isFinish*	- is **true**, the complete results are returned.
//
//   - This API operation supports only ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters for which Database Autonomy Service (DAS) Enterprise Edition is enabled. For more information, see [Purchase DAS Enterprise Edition](https://help.aliyun.com/document_detail/163298.html).
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetAsyncErrorRequestStatByCodeRequest
//
// @return GetAsyncErrorRequestStatByCodeResponse
func (client *Client) GetAsyncErrorRequestStatByCode(request *GetAsyncErrorRequestStatByCodeRequest) (_result *GetAsyncErrorRequestStatByCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAsyncErrorRequestStatByCodeResponse{}
	_body, _err := client.GetAsyncErrorRequestStatByCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously obtains the number of failed executions of SQL templates based on SQL Explorer data.
//
// Description:
//
// >  GetAsyncErrorRequestStatResult is an asynchronous operation. After a request is sent, the complete results are not returned immediately. If the value of **isFinish*	- is **false*	- in the response, wait for 1 second and then send a request again. If the value of **isFinish*	- is **true**, the complete results are returned.
//
//   - This API operation supports only ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters for which Database Autonomy Service (DAS) Enterprise Edition is enabled. For more information, see [Purchase DAS Enterprise Edition](https://help.aliyun.com/document_detail/163298.html).
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetAsyncErrorRequestStatResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncErrorRequestStatResultResponse
func (client *Client) GetAsyncErrorRequestStatResultWithOptions(request *GetAsyncErrorRequestStatResultRequest, runtime *dara.RuntimeOptions) (_result *GetAsyncErrorRequestStatResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.SqlIdList) {
		query["SqlIdList"] = request.SqlIdList
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncErrorRequestStatResult"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncErrorRequestStatResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously obtains the number of failed executions of SQL templates based on SQL Explorer data.
//
// Description:
//
// >  GetAsyncErrorRequestStatResult is an asynchronous operation. After a request is sent, the complete results are not returned immediately. If the value of **isFinish*	- is **false*	- in the response, wait for 1 second and then send a request again. If the value of **isFinish*	- is **true**, the complete results are returned.
//
//   - This API operation supports only ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters for which Database Autonomy Service (DAS) Enterprise Edition is enabled. For more information, see [Purchase DAS Enterprise Edition](https://help.aliyun.com/document_detail/163298.html).
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetAsyncErrorRequestStatResultRequest
//
// @return GetAsyncErrorRequestStatResultResponse
func (client *Client) GetAsyncErrorRequestStatResult(request *GetAsyncErrorRequestStatResultRequest) (_result *GetAsyncErrorRequestStatResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAsyncErrorRequestStatResultResponse{}
	_body, _err := client.GetAsyncErrorRequestStatResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage of auto-increment table IDs.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call DAS, you must set the region to cn-shanghai.
//
// @param request - GetAutoIncrementUsageStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoIncrementUsageStatisticResponse
func (client *Client) GetAutoIncrementUsageStatisticWithOptions(request *GetAutoIncrementUsageStatisticRequest, runtime *dara.RuntimeOptions) (_result *GetAutoIncrementUsageStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbNames) {
		query["DbNames"] = request.DbNames
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RatioFilter) {
		query["RatioFilter"] = request.RatioFilter
	}

	if !dara.IsNil(request.RealTime) {
		query["RealTime"] = request.RealTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoIncrementUsageStatistic"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoIncrementUsageStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage of auto-increment table IDs.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call DAS, you must set the region to cn-shanghai.
//
// @param request - GetAutoIncrementUsageStatisticRequest
//
// @return GetAutoIncrementUsageStatisticResponse
func (client *Client) GetAutoIncrementUsageStatistic(request *GetAutoIncrementUsageStatisticRequest) (_result *GetAutoIncrementUsageStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoIncrementUsageStatisticResponse{}
	_body, _err := client.GetAutoIncrementUsageStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the automatic fragment recycling rules of database instances.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - The database instance is an ApsaraDB RDS for MySQL instance of High-availability Edition.
//
//   - The database instance has four or more cores, and **innodb_file_per_table*	- is set to **ON**.
//
// @param request - GetAutoResourceOptimizeRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoResourceOptimizeRulesResponse
func (client *Client) GetAutoResourceOptimizeRulesWithOptions(request *GetAutoResourceOptimizeRulesRequest, runtime *dara.RuntimeOptions) (_result *GetAutoResourceOptimizeRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoResourceOptimizeRules"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoResourceOptimizeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the automatic fragment recycling rules of database instances.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - The database instance is an ApsaraDB RDS for MySQL instance of High-availability Edition.
//
//   - The database instance has four or more cores, and **innodb_file_per_table*	- is set to **ON**.
//
// @param request - GetAutoResourceOptimizeRulesRequest
//
// @return GetAutoResourceOptimizeRulesResponse
func (client *Client) GetAutoResourceOptimizeRules(request *GetAutoResourceOptimizeRulesRequest) (_result *GetAutoResourceOptimizeRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoResourceOptimizeRulesResponse{}
	_body, _err := client.GetAutoResourceOptimizeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the automatic SQL throttling rules of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - The database instance that you want to manage must be of one of the following types:
//
//   - ApsaraDB RDS for MySQL High-availability Edition or Enterprise Edition that runs MySQL 5.6, MySQL 5.7, or MySQL 8.0
//
//   - PolarDB for MySQL Cluster Edition that runs MySQL 5.6, MySQL 5.7, or MySQL 8.0
//
// @param request - GetAutoThrottleRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoThrottleRulesResponse
func (client *Client) GetAutoThrottleRulesWithOptions(request *GetAutoThrottleRulesRequest, runtime *dara.RuntimeOptions) (_result *GetAutoThrottleRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoThrottleRules"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoThrottleRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the automatic SQL throttling rules of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - The database instance that you want to manage must be of one of the following types:
//
//   - ApsaraDB RDS for MySQL High-availability Edition or Enterprise Edition that runs MySQL 5.6, MySQL 5.7, or MySQL 8.0
//
//   - PolarDB for MySQL Cluster Edition that runs MySQL 5.6, MySQL 5.7, or MySQL 8.0
//
// @param request - GetAutoThrottleRulesRequest
//
// @return GetAutoThrottleRulesResponse
func (client *Client) GetAutoThrottleRules(request *GetAutoThrottleRulesRequest) (_result *GetAutoThrottleRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoThrottleRulesResponse{}
	_body, _err := client.GetAutoThrottleRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of notification events of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - After your instance is connected to DAS, notification events such as snapshot capture are triggered if DAS detects changes to database monitoring metrics during anomaly detection.
//
// >  You can query the details of notification events only if the autonomy center is enabled. For more information, see [Autonomy center](https://help.aliyun.com/document_detail/152139.html).
//
// @param request - GetAutonomousNotifyEventContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutonomousNotifyEventContentResponse
func (client *Client) GetAutonomousNotifyEventContentWithOptions(request *GetAutonomousNotifyEventContentRequest, runtime *dara.RuntimeOptions) (_result *GetAutonomousNotifyEventContentResponse, _err error) {
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

	if !dara.IsNil(request.SpanId) {
		query["SpanId"] = request.SpanId
	}

	if !dara.IsNil(request.Context) {
		query["__context"] = request.Context
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutonomousNotifyEventContent"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutonomousNotifyEventContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of notification events of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - After your instance is connected to DAS, notification events such as snapshot capture are triggered if DAS detects changes to database monitoring metrics during anomaly detection.
//
// >  You can query the details of notification events only if the autonomy center is enabled. For more information, see [Autonomy center](https://help.aliyun.com/document_detail/152139.html).
//
// @param request - GetAutonomousNotifyEventContentRequest
//
// @return GetAutonomousNotifyEventContentResponse
func (client *Client) GetAutonomousNotifyEventContent(request *GetAutonomousNotifyEventContentRequest) (_result *GetAutonomousNotifyEventContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutonomousNotifyEventContentResponse{}
	_body, _err := client.GetAutonomousNotifyEventContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the notification events of one or more urgency levels within a period.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - After your instance is connected to DAS, notification events such as snapshot capture are triggered if DAS detects changes to database monitoring metrics during anomaly detection.
//
// >  You can query the details of notification events only if the autonomy center is enabled. For more information, see [Autonomy center](https://help.aliyun.com/document_detail/152139.html).
//
// @param request - GetAutonomousNotifyEventsInRangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutonomousNotifyEventsInRangeResponse
func (client *Client) GetAutonomousNotifyEventsInRangeWithOptions(request *GetAutonomousNotifyEventsInRangeRequest, runtime *dara.RuntimeOptions) (_result *GetAutonomousNotifyEventsInRangeResponse, _err error) {
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

	if !dara.IsNil(request.EventContext) {
		query["EventContext"] = request.EventContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MinLevel) {
		query["MinLevel"] = request.MinLevel
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageOffset) {
		query["PageOffset"] = request.PageOffset
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Context) {
		query["__context"] = request.Context
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutonomousNotifyEventsInRange"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutonomousNotifyEventsInRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the notification events of one or more urgency levels within a period.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - After your instance is connected to DAS, notification events such as snapshot capture are triggered if DAS detects changes to database monitoring metrics during anomaly detection.
//
// >  You can query the details of notification events only if the autonomy center is enabled. For more information, see [Autonomy center](https://help.aliyun.com/document_detail/152139.html).
//
// @param request - GetAutonomousNotifyEventsInRangeRequest
//
// @return GetAutonomousNotifyEventsInRangeResponse
func (client *Client) GetAutonomousNotifyEventsInRange(request *GetAutonomousNotifyEventsInRangeRequest) (_result *GetAutonomousNotifyEventsInRangeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutonomousNotifyEventsInRangeResponse{}
	_body, _err := client.GetAutonomousNotifyEventsInRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the blocking data of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for SQL Server instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetBlockingDetailListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBlockingDetailListResponse
func (client *Client) GetBlockingDetailListWithOptions(request *GetBlockingDetailListRequest, runtime *dara.RuntimeOptions) (_result *GetBlockingDetailListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbNameList) {
		query["DbNameList"] = request.DbNameList
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryHash) {
		query["QueryHash"] = request.QueryHash
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBlockingDetailList"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBlockingDetailListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the blocking data of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for SQL Server instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetBlockingDetailListRequest
//
// @return GetBlockingDetailListResponse
func (client *Client) GetBlockingDetailList(request *GetBlockingDetailListRequest) (_result *GetBlockingDetailListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBlockingDetailListResponse{}
	_body, _err := client.GetBlockingDetailListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the diagnosis of network connectivity when a user accesses a specific database instance by specifying an IP address.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - The database instance that you want to manage is connected to DAS.
//
// @param request - GetDBInstanceConnectivityDiagnosisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDBInstanceConnectivityDiagnosisResponse
func (client *Client) GetDBInstanceConnectivityDiagnosisWithOptions(request *GetDBInstanceConnectivityDiagnosisRequest, runtime *dara.RuntimeOptions) (_result *GetDBInstanceConnectivityDiagnosisResponse, _err error) {
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

	if !dara.IsNil(request.SrcIp) {
		query["SrcIp"] = request.SrcIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDBInstanceConnectivityDiagnosis"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDBInstanceConnectivityDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the diagnosis of network connectivity when a user accesses a specific database instance by specifying an IP address.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - The database instance that you want to manage is connected to DAS.
//
// @param request - GetDBInstanceConnectivityDiagnosisRequest
//
// @return GetDBInstanceConnectivityDiagnosisResponse
func (client *Client) GetDBInstanceConnectivityDiagnosis(request *GetDBInstanceConnectivityDiagnosisRequest) (_result *GetDBInstanceConnectivityDiagnosisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDBInstanceConnectivityDiagnosisResponse{}
	_body, _err := client.GetDBInstanceConnectivityDiagnosisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DAS大模型能力异步逻辑接口
//
// @param request - GetDasAgentSSERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDasAgentSSEResponse
func (client *Client) GetDasAgentSSEWithSSE(request *GetDasAgentSSERequest, runtime *dara.RuntimeOptions, _yield chan *GetDasAgentSSEResponse, _yieldErr chan error) {
	defer close(_yield)
	client.getDasAgentSSEWithSSE_opYieldFunc(_yield, _yieldErr, request, runtime)
	return
}

// Summary:
//
// # DAS大模型能力异步逻辑接口
//
// @param request - GetDasAgentSSERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDasAgentSSEResponse
func (client *Client) GetDasAgentSSEWithOptions(request *GetDasAgentSSERequest, runtime *dara.RuntimeOptions) (_result *GetDasAgentSSEResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDasAgentSSE"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDasAgentSSEResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DAS大模型能力异步逻辑接口
//
// @param request - GetDasAgentSSERequest
//
// @return GetDasAgentSSEResponse
func (client *Client) GetDasAgentSSE(request *GetDasAgentSSERequest) (_result *GetDasAgentSSEResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDasAgentSSEResponse{}
	_body, _err := client.GetDasAgentSSEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the storage usage of a database instance for which Database Autonomy Service (DAS) Enterprise Edition V1 or V2 is enabled.
//
// Description:
//
//	  For information about the database instances that support this operation, see [Overview of DAS Enterprise Edition](https://help.aliyun.com/document_detail/190912.html).
//
//		- If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation is applicable only to DAS Enterprise Edition V1 and V2.
//
// >  We recommend that you call the [DescribeSqlLogStatistic](https://help.aliyun.com/document_detail/2778836.html) operation to query the data statistics of a database instance for which DAS Enterprise Edition is enabled.
//
// @param request - GetDasProServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDasProServiceUsageResponse
func (client *Client) GetDasProServiceUsageWithOptions(request *GetDasProServiceUsageRequest, runtime *dara.RuntimeOptions) (_result *GetDasProServiceUsageResponse, _err error) {
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

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDasProServiceUsage"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDasProServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage usage of a database instance for which Database Autonomy Service (DAS) Enterprise Edition V1 or V2 is enabled.
//
// Description:
//
//	  For information about the database instances that support this operation, see [Overview of DAS Enterprise Edition](https://help.aliyun.com/document_detail/190912.html).
//
//		- If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation is applicable only to DAS Enterprise Edition V1 and V2.
//
// >  We recommend that you call the [DescribeSqlLogStatistic](https://help.aliyun.com/document_detail/2778836.html) operation to query the data statistics of a database instance for which DAS Enterprise Edition is enabled.
//
// @param request - GetDasProServiceUsageRequest
//
// @return GetDasProServiceUsageResponse
func (client *Client) GetDasProServiceUsage(request *GetDasProServiceUsageRequest) (_result *GetDasProServiceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDasProServiceUsageResponse{}
	_body, _err := client.GetDasProServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hot data of audit logs.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable to PolarDB for MySQL, ApsaraDB RDS for MySQL, ApsaraDB RDS for PostgreSQL, and ApsaraDB RDS for SQL Server.
//
// >  The beginning of the time range to query can be up to seven days earlier than the current time. The interval between the start time and the end time cannot exceed one day. This operation can return a maximum of 10,000 entries.
//
// @param request - GetDasSQLLogHotDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDasSQLLogHotDataResponse
func (client *Client) GetDasSQLLogHotDataWithOptions(request *GetDasSQLLogHotDataRequest, runtime *dara.RuntimeOptions) (_result *GetDasSQLLogHotDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ChildDBInstanceIDs) {
		body["ChildDBInstanceIDs"] = request.ChildDBInstanceIDs
	}

	if !dara.IsNil(request.DBName) {
		body["DBName"] = request.DBName
	}

	if !dara.IsNil(request.End) {
		body["End"] = request.End
	}

	if !dara.IsNil(request.Fail) {
		body["Fail"] = request.Fail
	}

	if !dara.IsNil(request.HostAddress) {
		body["HostAddress"] = request.HostAddress
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogicalOperator) {
		body["LogicalOperator"] = request.LogicalOperator
	}

	if !dara.IsNil(request.MaxLatancy) {
		body["MaxLatancy"] = request.MaxLatancy
	}

	if !dara.IsNil(request.MaxRecordsPerPage) {
		body["MaxRecordsPerPage"] = request.MaxRecordsPerPage
	}

	if !dara.IsNil(request.MaxRows) {
		body["MaxRows"] = request.MaxRows
	}

	if !dara.IsNil(request.MaxScanRows) {
		body["MaxScanRows"] = request.MaxScanRows
	}

	if !dara.IsNil(request.MaxSpillCnt) {
		body["MaxSpillCnt"] = request.MaxSpillCnt
	}

	if !dara.IsNil(request.MinLatancy) {
		body["MinLatancy"] = request.MinLatancy
	}

	if !dara.IsNil(request.MinRows) {
		body["MinRows"] = request.MinRows
	}

	if !dara.IsNil(request.MinScanRows) {
		body["MinScanRows"] = request.MinScanRows
	}

	if !dara.IsNil(request.MinSpillCnt) {
		body["MinSpillCnt"] = request.MinSpillCnt
	}

	if !dara.IsNil(request.PageNumbers) {
		body["PageNumbers"] = request.PageNumbers
	}

	if !dara.IsNil(request.QueryKeyword) {
		body["QueryKeyword"] = request.QueryKeyword
	}

	if !dara.IsNil(request.Role) {
		body["Role"] = request.Role
	}

	if !dara.IsNil(request.SortKey) {
		body["SortKey"] = request.SortKey
	}

	if !dara.IsNil(request.SortMethod) {
		body["SortMethod"] = request.SortMethod
	}

	if !dara.IsNil(request.SqlType) {
		body["SqlType"] = request.SqlType
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	if !dara.IsNil(request.State) {
		body["State"] = request.State
	}

	if !dara.IsNil(request.ThreadID) {
		body["ThreadID"] = request.ThreadID
	}

	if !dara.IsNil(request.TraceId) {
		body["TraceId"] = request.TraceId
	}

	if !dara.IsNil(request.TransactionId) {
		body["TransactionId"] = request.TransactionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDasSQLLogHotData"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDasSQLLogHotDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hot data of audit logs.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - This operation is applicable to PolarDB for MySQL, ApsaraDB RDS for MySQL, ApsaraDB RDS for PostgreSQL, and ApsaraDB RDS for SQL Server.
//
// >  The beginning of the time range to query can be up to seven days earlier than the current time. The interval between the start time and the end time cannot exceed one day. This operation can return a maximum of 10,000 entries.
//
// @param request - GetDasSQLLogHotDataRequest
//
// @return GetDasSQLLogHotDataResponse
func (client *Client) GetDasSQLLogHotData(request *GetDasSQLLogHotDataRequest) (_result *GetDasSQLLogHotDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDasSQLLogHotDataResponse{}
	_body, _err := client.GetDasSQLLogHotDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询单个死锁详情
//
// @param request - GetDeadLockDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeadLockDetailResponse
func (client *Client) GetDeadLockDetailWithOptions(request *GetDeadLockDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDeadLockDetailResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TextId) {
		query["TextId"] = request.TextId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeadLockDetail"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeadLockDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询单个死锁详情
//
// @param request - GetDeadLockDetailRequest
//
// @return GetDeadLockDetailResponse
func (client *Client) GetDeadLockDetail(request *GetDeadLockDetailRequest) (_result *GetDeadLockDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeadLockDetailResponse{}
	_body, _err := client.GetDeadLockDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the deadlock details of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for SQL Server instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetDeadLockDetailListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeadLockDetailListResponse
func (client *Client) GetDeadLockDetailListWithOptions(request *GetDeadLockDetailListRequest, runtime *dara.RuntimeOptions) (_result *GetDeadLockDetailListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbNameList) {
		query["DbNameList"] = request.DbNameList
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeadLockDetailList"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeadLockDetailListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deadlock details of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for SQL Server instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetDeadLockDetailListRequest
//
// @return GetDeadLockDetailListResponse
func (client *Client) GetDeadLockDetailList(request *GetDeadLockDetailListRequest) (_result *GetDeadLockDetailListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeadLockDetailListResponse{}
	_body, _err := client.GetDeadLockDetailListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取历史死锁记录
//
// @param request - GetDeadLockHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeadLockHistoryResponse
func (client *Client) GetDeadLockHistoryWithOptions(request *GetDeadLockHistoryRequest, runtime *dara.RuntimeOptions) (_result *GetDeadLockHistoryResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeadLockHistory"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeadLockHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取历史死锁记录
//
// @param request - GetDeadLockHistoryRequest
//
// @return GetDeadLockHistoryResponse
func (client *Client) GetDeadLockHistory(request *GetDeadLockHistoryRequest) (_result *GetDeadLockHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeadLockHistoryResponse{}
	_body, _err := client.GetDeadLockHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询时间范围内基于错误日志分析的死锁数量
//
// @param request - GetDeadlockHistogramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeadlockHistogramResponse
func (client *Client) GetDeadlockHistogramWithOptions(request *GetDeadlockHistogramRequest, runtime *dara.RuntimeOptions) (_result *GetDeadlockHistogramResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeadlockHistogram"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeadlockHistogramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询时间范围内基于错误日志分析的死锁数量
//
// @param request - GetDeadlockHistogramRequest
//
// @return GetDeadlockHistogramResponse
func (client *Client) GetDeadlockHistogram(request *GetDeadlockHistogramRequest) (_result *GetDeadlockHistogramResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeadlockHistogramResponse{}
	_body, _err := client.GetDeadlockHistogramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetEndpointSwitchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEndpointSwitchTaskResponse
func (client *Client) GetEndpointSwitchTaskWithOptions(request *GetEndpointSwitchTaskRequest, runtime *dara.RuntimeOptions) (_result *GetEndpointSwitchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Context) {
		query["__context"] = request.Context
	}

	if !dara.IsNil(request.AccessKey) {
		query["accessKey"] = request.AccessKey
	}

	if !dara.IsNil(request.Signature) {
		query["signature"] = request.Signature
	}

	if !dara.IsNil(request.SkipAuth) {
		query["skipAuth"] = request.SkipAuth
	}

	if !dara.IsNil(request.Timestamp) {
		query["timestamp"] = request.Timestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEndpointSwitchTask"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEndpointSwitchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetEndpointSwitchTaskRequest
//
// @return GetEndpointSwitchTaskResponse
func (client *Client) GetEndpointSwitchTask(request *GetEndpointSwitchTaskRequest) (_result *GetEndpointSwitchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEndpointSwitchTaskResponse{}
	_body, _err := client.GetEndpointSwitchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously queries information about failed SQL queries in SQL Explorer data. You can query up to 20 failed SQL queries within the specific time range.
//
// Description:
//
// >  GetErrorRequestSample is an asynchronous operation. After a request is sent, the complete results are not returned immediately. If the value of **isFinish*	- is **false*	- in the response, wait for 1 second and then send a request again. If the value of **isFinish*	- is **true**, the complete results are returned.
//
//   - This API operation supports only ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters for which Database Autonomy Service (DAS) Enterprise Edition is enabled. For more information, see [Purchase DAS Enterprise Edition](https://help.aliyun.com/document_detail/163298.html).
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetErrorRequestSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErrorRequestSampleResponse
func (client *Client) GetErrorRequestSampleWithOptions(request *GetErrorRequestSampleRequest, runtime *dara.RuntimeOptions) (_result *GetErrorRequestSampleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.SqlId) {
		query["SqlId"] = request.SqlId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetErrorRequestSample"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetErrorRequestSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously queries information about failed SQL queries in SQL Explorer data. You can query up to 20 failed SQL queries within the specific time range.
//
// Description:
//
// >  GetErrorRequestSample is an asynchronous operation. After a request is sent, the complete results are not returned immediately. If the value of **isFinish*	- is **false*	- in the response, wait for 1 second and then send a request again. If the value of **isFinish*	- is **true**, the complete results are returned.
//
//   - This API operation supports only ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters for which Database Autonomy Service (DAS) Enterprise Edition is enabled. For more information, see [Purchase DAS Enterprise Edition](https://help.aliyun.com/document_detail/163298.html).
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetErrorRequestSampleRequest
//
// @return GetErrorRequestSampleResponse
func (client *Client) GetErrorRequestSample(request *GetErrorRequestSampleRequest) (_result *GetErrorRequestSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetErrorRequestSampleResponse{}
	_body, _err := client.GetErrorRequestSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the event subscription settings of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - The database instance that you want to manage is connected to DAS.
//
// @param request - GetEventSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEventSubscriptionResponse
func (client *Client) GetEventSubscriptionWithOptions(request *GetEventSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *GetEventSubscriptionResponse, _err error) {
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
		Action:      dara.String("GetEventSubscription"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEventSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the event subscription settings of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - The database instance that you want to manage is connected to DAS.
//
// @param request - GetEventSubscriptionRequest
//
// @return GetEventSubscriptionResponse
func (client *Client) GetEventSubscription(request *GetEventSubscriptionRequest) (_result *GetEventSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEventSubscriptionResponse{}
	_body, _err := client.GetEventSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Collects the full request statistics in the SQL Explorer results of a database instance by access source.
//
// Description:
//
// The SQL Explorer feature allows you to check the health status of SQL statements and troubleshoot performance issues. For more information, see [SQL Explorer](https://help.aliyun.com/document_detail/204096.html).
//
//   - For more information about database instances that support this feature, see [Overview](https://help.aliyun.com/document_detail/190912.html).
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// @param request - GetFullRequestOriginStatByInstanceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFullRequestOriginStatByInstanceIdResponse
func (client *Client) GetFullRequestOriginStatByInstanceIdWithOptions(request *GetFullRequestOriginStatByInstanceIdRequest, runtime *dara.RuntimeOptions) (_result *GetFullRequestOriginStatByInstanceIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Asc) {
		query["Asc"] = request.Asc
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.SqlType) {
		query["SqlType"] = request.SqlType
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFullRequestOriginStatByInstanceId"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFullRequestOriginStatByInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Collects the full request statistics in the SQL Explorer results of a database instance by access source.
//
// Description:
//
// The SQL Explorer feature allows you to check the health status of SQL statements and troubleshoot performance issues. For more information, see [SQL Explorer](https://help.aliyun.com/document_detail/204096.html).
//
//   - For more information about database instances that support this feature, see [Overview](https://help.aliyun.com/document_detail/190912.html).
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// @param request - GetFullRequestOriginStatByInstanceIdRequest
//
// @return GetFullRequestOriginStatByInstanceIdResponse
func (client *Client) GetFullRequestOriginStatByInstanceId(request *GetFullRequestOriginStatByInstanceIdRequest) (_result *GetFullRequestOriginStatByInstanceIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFullRequestOriginStatByInstanceIdResponse{}
	_body, _err := client.GetFullRequestOriginStatByInstanceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries sample SQL statements in the SQL Explorer data of a database instance by SQL ID. You can query up to 20 sample SQL statements.
//
// Description:
//
// The SQL Explorer feature allows you to check the health status of SQL statements and troubleshoot performance issues. For more information, see [SQL Explorer](https://help.aliyun.com/document_detail/204096.html).
//
//   - For more information about the database engines that support SQL Explorer, see [SQL Explorer](https://help.aliyun.com/document_detail/204096.html).
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetFullRequestSampleByInstanceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFullRequestSampleByInstanceIdResponse
func (client *Client) GetFullRequestSampleByInstanceIdWithOptions(request *GetFullRequestSampleByInstanceIdRequest, runtime *dara.RuntimeOptions) (_result *GetFullRequestSampleByInstanceIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.End) {
		body["End"] = request.End
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SqlId) {
		body["SqlId"] = request.SqlId
	}

	if !dara.IsNil(request.Start) {
		body["Start"] = request.Start
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFullRequestSampleByInstanceId"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFullRequestSampleByInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries sample SQL statements in the SQL Explorer data of a database instance by SQL ID. You can query up to 20 sample SQL statements.
//
// Description:
//
// The SQL Explorer feature allows you to check the health status of SQL statements and troubleshoot performance issues. For more information, see [SQL Explorer](https://help.aliyun.com/document_detail/204096.html).
//
//   - For more information about the database engines that support SQL Explorer, see [SQL Explorer](https://help.aliyun.com/document_detail/204096.html).
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetFullRequestSampleByInstanceIdRequest
//
// @return GetFullRequestSampleByInstanceIdResponse
func (client *Client) GetFullRequestSampleByInstanceId(request *GetFullRequestSampleByInstanceIdRequest) (_result *GetFullRequestSampleByInstanceIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFullRequestSampleByInstanceIdResponse{}
	_body, _err := client.GetFullRequestSampleByInstanceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously collects the full request statistics in the SQL Explorer results of a database instance by SQL ID.
//
// Description:
//
// >  GetFullRequestStatResultByInstanceId is an asynchronous operation. After a request is sent, the complete results are not returned immediately. If the value of the isFinish parameter is **false*	- in the response, wait for 1 second and then send a request again. If the value of the isFinish parameter is **true**, the complete results are returned.
//
// The SQL Explorer feature allows you to check the health status of SQL statements and troubleshoot performance issues. For more information, see [SQL Explorer](https://help.aliyun.com/document_detail/204096.html).
//
//   - For more information about database instances that support this feature, see [Overview of DAS Enterprise Edition](https://help.aliyun.com/document_detail/190912.html).
//
//   - If you use an SDK to call the API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - When you call this operation, the value of the SqlId parameter changes due to the optimization of the SQL template algorithm starting from September 1, 2024. For more information, see [[Notice\\] Optimization of the SQL template algorithm](~~2845725~~).
//
// @param request - GetFullRequestStatResultByInstanceIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFullRequestStatResultByInstanceIdResponse
func (client *Client) GetFullRequestStatResultByInstanceIdWithOptions(request *GetFullRequestStatResultByInstanceIdRequest, runtime *dara.RuntimeOptions) (_result *GetFullRequestStatResultByInstanceIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Asc) {
		query["Asc"] = request.Asc
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.End) {
		query["End"] = request.End
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OriginHost) {
		query["OriginHost"] = request.OriginHost
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.SqlId) {
		query["SqlId"] = request.SqlId
	}

	if !dara.IsNil(request.SqlType) {
		query["SqlType"] = request.SqlType
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFullRequestStatResultByInstanceId"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFullRequestStatResultByInstanceIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously collects the full request statistics in the SQL Explorer results of a database instance by SQL ID.
//
// Description:
//
// >  GetFullRequestStatResultByInstanceId is an asynchronous operation. After a request is sent, the complete results are not returned immediately. If the value of the isFinish parameter is **false*	- in the response, wait for 1 second and then send a request again. If the value of the isFinish parameter is **true**, the complete results are returned.
//
// The SQL Explorer feature allows you to check the health status of SQL statements and troubleshoot performance issues. For more information, see [SQL Explorer](https://help.aliyun.com/document_detail/204096.html).
//
//   - For more information about database instances that support this feature, see [Overview of DAS Enterprise Edition](https://help.aliyun.com/document_detail/190912.html).
//
//   - If you use an SDK to call the API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - When you call this operation, the value of the SqlId parameter changes due to the optimization of the SQL template algorithm starting from September 1, 2024. For more information, see [[Notice\\] Optimization of the SQL template algorithm](~~2845725~~).
//
// @param request - GetFullRequestStatResultByInstanceIdRequest
//
// @return GetFullRequestStatResultByInstanceIdResponse
func (client *Client) GetFullRequestStatResultByInstanceId(request *GetFullRequestStatResultByInstanceIdRequest) (_result *GetFullRequestStatResultByInstanceIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFullRequestStatResultByInstanceIdResponse{}
	_body, _err := client.GetFullRequestStatResultByInstanceIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetHDMAliyunResourceSyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHDMAliyunResourceSyncResultResponse
func (client *Client) GetHDMAliyunResourceSyncResultWithOptions(request *GetHDMAliyunResourceSyncResultRequest, runtime *dara.RuntimeOptions) (_result *GetHDMAliyunResourceSyncResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Context) {
		query["__context"] = request.Context
	}

	if !dara.IsNil(request.AccessKey) {
		query["accessKey"] = request.AccessKey
	}

	if !dara.IsNil(request.Signature) {
		query["signature"] = request.Signature
	}

	if !dara.IsNil(request.SkipAuth) {
		query["skipAuth"] = request.SkipAuth
	}

	if !dara.IsNil(request.Timestamp) {
		query["timestamp"] = request.Timestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHDMAliyunResourceSyncResult"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHDMAliyunResourceSyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetHDMAliyunResourceSyncResultRequest
//
// @return GetHDMAliyunResourceSyncResultResponse
func (client *Client) GetHDMAliyunResourceSyncResult(request *GetHDMAliyunResourceSyncResultRequest) (_result *GetHDMAliyunResourceSyncResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHDMAliyunResourceSyncResultResponse{}
	_body, _err := client.GetHDMAliyunResourceSyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetHDMLastAliyunResourceSyncResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHDMLastAliyunResourceSyncResultResponse
func (client *Client) GetHDMLastAliyunResourceSyncResultWithOptions(request *GetHDMLastAliyunResourceSyncResultRequest, runtime *dara.RuntimeOptions) (_result *GetHDMLastAliyunResourceSyncResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Context) {
		query["__context"] = request.Context
	}

	if !dara.IsNil(request.AccessKey) {
		query["accessKey"] = request.AccessKey
	}

	if !dara.IsNil(request.Signature) {
		query["signature"] = request.Signature
	}

	if !dara.IsNil(request.SkipAuth) {
		query["skipAuth"] = request.SkipAuth
	}

	if !dara.IsNil(request.Timestamp) {
		query["timestamp"] = request.Timestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHDMLastAliyunResourceSyncResult"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHDMLastAliyunResourceSyncResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetHDMLastAliyunResourceSyncResultRequest
//
// @return GetHDMLastAliyunResourceSyncResultResponse
func (client *Client) GetHDMLastAliyunResourceSyncResult(request *GetHDMLastAliyunResourceSyncResultRequest) (_result *GetHDMLastAliyunResourceSyncResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHDMLastAliyunResourceSyncResultResponse{}
	_body, _err := client.GetHDMLastAliyunResourceSyncResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of an inspection that is performed on a database instance by using the inspection and scoring feature.
//
// Description:
//
// Database Autonomy Service (DAS) provides the inspection and scoring feature. This feature allows you to inspect and score the health status of your instance on a regular basis. This helps you obtain information about the status of your databases. For more information, see [Inspection and scoring](https://help.aliyun.com/document_detail/205659.html).
//
// Before you call this operation, take note of the following items:
//
//   - This operation is applicable only to ApsaraDB RDS for MySQL databases, self-managed MySQL databases hosted on Elastic Compute Service (ECS) instances, self-managed MySQL databases in data centers, ApsaraDB for Redis databases, and PolarDB for MySQL databases.
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V4.3.3. We recommend that you use the latest version.
//
//   - The version of DAS SDK must be V1.0.3 or later.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetInstanceInspectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceInspectionsResponse
func (client *Client) GetInstanceInspectionsWithOptions(request *GetInstanceInspectionsRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceInspectionsResponse, _err error) {
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

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceArea) {
		query["InstanceArea"] = request.InstanceArea
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SearchMap) {
		query["SearchMap"] = request.SearchMap
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceInspections"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceInspectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of an inspection that is performed on a database instance by using the inspection and scoring feature.
//
// Description:
//
// Database Autonomy Service (DAS) provides the inspection and scoring feature. This feature allows you to inspect and score the health status of your instance on a regular basis. This helps you obtain information about the status of your databases. For more information, see [Inspection and scoring](https://help.aliyun.com/document_detail/205659.html).
//
// Before you call this operation, take note of the following items:
//
//   - This operation is applicable only to ApsaraDB RDS for MySQL databases, self-managed MySQL databases hosted on Elastic Compute Service (ECS) instances, self-managed MySQL databases in data centers, ApsaraDB for Redis databases, and PolarDB for MySQL databases.
//
//   - If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V4.3.3. We recommend that you use the latest version.
//
//   - The version of DAS SDK must be V1.0.3 or later.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetInstanceInspectionsRequest
//
// @return GetInstanceInspectionsResponse
func (client *Client) GetInstanceInspections(request *GetInstanceInspectionsRequest) (_result *GetInstanceInspectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceInspectionsResponse{}
	_body, _err := client.GetInstanceInspectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of all missing indexes of an instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for SQL Server instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetInstanceMissingIndexListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceMissingIndexListResponse
func (client *Client) GetInstanceMissingIndexListWithOptions(request *GetInstanceMissingIndexListRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceMissingIndexListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvgTotalUserCost) {
		query["AvgTotalUserCost"] = request.AvgTotalUserCost
	}

	if !dara.IsNil(request.AvgUserImpact) {
		query["AvgUserImpact"] = request.AvgUserImpact
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IndexCount) {
		query["IndexCount"] = request.IndexCount
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ObjectName) {
		query["ObjectName"] = request.ObjectName
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReservedPages) {
		query["ReservedPages"] = request.ReservedPages
	}

	if !dara.IsNil(request.ReservedSize) {
		query["ReservedSize"] = request.ReservedSize
	}

	if !dara.IsNil(request.RowCount) {
		query["RowCount"] = request.RowCount
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UniqueCompiles) {
		query["UniqueCompiles"] = request.UniqueCompiles
	}

	if !dara.IsNil(request.UserScans) {
		query["UserScans"] = request.UserScans
	}

	if !dara.IsNil(request.UserSeeks) {
		query["UserSeeks"] = request.UserSeeks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceMissingIndexList"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceMissingIndexListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of all missing indexes of an instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for SQL Server instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetInstanceMissingIndexListRequest
//
// @return GetInstanceMissingIndexListResponse
func (client *Client) GetInstanceMissingIndexList(request *GetInstanceMissingIndexListRequest) (_result *GetInstanceMissingIndexListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceMissingIndexListResponse{}
	_body, _err := client.GetInstanceMissingIndexListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries statistics on automatic SQL optimization events within a period of time, such as the total number of optimization events and the maximum improvement.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this API operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - The database engine is ApsaraDB RDS for MySQL or PolarDB for MySQL.
//
// @param request - GetInstanceSqlOptimizeStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceSqlOptimizeStatisticResponse
func (client *Client) GetInstanceSqlOptimizeStatisticWithOptions(request *GetInstanceSqlOptimizeStatisticRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceSqlOptimizeStatisticResponse, _err error) {
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

	if !dara.IsNil(request.FilterEnable) {
		query["FilterEnable"] = request.FilterEnable
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	if !dara.IsNil(request.UseMerging) {
		query["UseMerging"] = request.UseMerging
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceSqlOptimizeStatistic"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceSqlOptimizeStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries statistics on automatic SQL optimization events within a period of time, such as the total number of optimization events and the maximum improvement.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this API operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - The database engine is ApsaraDB RDS for MySQL or PolarDB for MySQL.
//
// @param request - GetInstanceSqlOptimizeStatisticRequest
//
// @return GetInstanceSqlOptimizeStatisticResponse
func (client *Client) GetInstanceSqlOptimizeStatistic(request *GetInstanceSqlOptimizeStatisticRequest) (_result *GetInstanceSqlOptimizeStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceSqlOptimizeStatisticResponse{}
	_body, _err := client.GetInstanceSqlOptimizeStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the results of a task that terminates sessions.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
//
//		- If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetKillInstanceSessionTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKillInstanceSessionTaskResultResponse
func (client *Client) GetKillInstanceSessionTaskResultWithOptions(request *GetKillInstanceSessionTaskResultRequest, runtime *dara.RuntimeOptions) (_result *GetKillInstanceSessionTaskResultResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKillInstanceSessionTaskResult"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKillInstanceSessionTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a task that terminates sessions.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
//
//		- If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetKillInstanceSessionTaskResultRequest
//
// @return GetKillInstanceSessionTaskResultResponse
func (client *Client) GetKillInstanceSessionTaskResult(request *GetKillInstanceSessionTaskResultRequest) (_result *GetKillInstanceSessionTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetKillInstanceSessionTaskResultResponse{}
	_body, _err := client.GetKillInstanceSessionTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current sessions of an ApsaraDB for MongoDB (MongoDB) instance.
//
// Description:
//
//	  This operation is applicable only to MongoDB instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call API operations of DAS, you must set the region to cn-shanghai.
//
// @param request - GetMongoDBCurrentOpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMongoDBCurrentOpResponse
func (client *Client) GetMongoDBCurrentOpWithOptions(request *GetMongoDBCurrentOpRequest, runtime *dara.RuntimeOptions) (_result *GetMongoDBCurrentOpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterDoc) {
		query["FilterDoc"] = request.FilterDoc
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMongoDBCurrentOp"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMongoDBCurrentOpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current sessions of an ApsaraDB for MongoDB (MongoDB) instance.
//
// Description:
//
//	  This operation is applicable only to MongoDB instances.
//
//		- If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call API operations of DAS, you must set the region to cn-shanghai.
//
// @param request - GetMongoDBCurrentOpRequest
//
// @return GetMongoDBCurrentOpResponse
func (client *Client) GetMongoDBCurrentOp(request *GetMongoDBCurrentOpRequest) (_result *GetMongoDBCurrentOpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMongoDBCurrentOpResponse{}
	_body, _err := client.GetMongoDBCurrentOpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously queries the sessions of an instance and collects statistics on the sessions based on dimensions.
//
// Description:
//
// >  GetMySQLAllSessionAsync is an asynchronous operation. After a request is sent, the system does not return complete results but returns a request ID. You need to use the request ID to initiate requests until the value of the **isFinish*	- field in the returned results is **true**, the complete results are returned. This indicates that to obtain complete data, you must call this operation at least twice.
//
//   - This operation is applicable only to ApsaraDB RDS for MySQL instances, PolarDB for MySQL clusters, and PolarDB-X 2.0 instances.
//
//   - If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetMySQLAllSessionAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMySQLAllSessionAsyncResponse
func (client *Client) GetMySQLAllSessionAsyncWithOptions(request *GetMySQLAllSessionAsyncRequest, runtime *dara.RuntimeOptions) (_result *GetMySQLAllSessionAsyncResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ResultId) {
		query["ResultId"] = request.ResultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMySQLAllSessionAsync"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMySQLAllSessionAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously queries the sessions of an instance and collects statistics on the sessions based on dimensions.
//
// Description:
//
// >  GetMySQLAllSessionAsync is an asynchronous operation. After a request is sent, the system does not return complete results but returns a request ID. You need to use the request ID to initiate requests until the value of the **isFinish*	- field in the returned results is **true**, the complete results are returned. This indicates that to obtain complete data, you must call this operation at least twice.
//
//   - This operation is applicable only to ApsaraDB RDS for MySQL instances, PolarDB for MySQL clusters, and PolarDB-X 2.0 instances.
//
//   - If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetMySQLAllSessionAsyncRequest
//
// @return GetMySQLAllSessionAsyncResponse
func (client *Client) GetMySQLAllSessionAsync(request *GetMySQLAllSessionAsyncRequest) (_result *GetMySQLAllSessionAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMySQLAllSessionAsyncResponse{}
	_body, _err := client.GetMySQLAllSessionAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries access frequency statistics and hot data on partitions of a PolarDB-X 2.0 instance.
//
// Description:
//
// We recommend that you do not call this operation. The data is returned in a special format and is complex to parse. You can use the [heatmap](https://help.aliyun.com/document_detail/470302.html) feature of Database Autonomy Service (DAS) to query the data.
//
// @param request - GetPartitionsHeatmapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPartitionsHeatmapResponse
func (client *Client) GetPartitionsHeatmapWithOptions(request *GetPartitionsHeatmapRequest, runtime *dara.RuntimeOptions) (_result *GetPartitionsHeatmapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TimeRange) {
		query["TimeRange"] = request.TimeRange
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPartitionsHeatmap"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPartitionsHeatmapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access frequency statistics and hot data on partitions of a PolarDB-X 2.0 instance.
//
// Description:
//
// We recommend that you do not call this operation. The data is returned in a special format and is complex to parse. You can use the [heatmap](https://help.aliyun.com/document_detail/470302.html) feature of Database Autonomy Service (DAS) to query the data.
//
// @param request - GetPartitionsHeatmapRequest
//
// @return GetPartitionsHeatmapResponse
func (client *Client) GetPartitionsHeatmap(request *GetPartitionsHeatmapRequest) (_result *GetPartitionsHeatmapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPartitionsHeatmapResponse{}
	_body, _err := client.GetPartitionsHeatmapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trend of a metric for the new version of the performance insight feature of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - An ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster is connected to DAS.
//
//   - The new version of the performance insight feature is enabled for the database instance. For more information, see [Performance insight (new version)](https://help.aliyun.com/document_detail/469117.html).
//
// @param request - GetPfsMetricTrendsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPfsMetricTrendsResponse
func (client *Client) GetPfsMetricTrendsWithOptions(request *GetPfsMetricTrendsRequest, runtime *dara.RuntimeOptions) (_result *GetPfsMetricTrendsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Metric) {
		body["Metric"] = request.Metric
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPfsMetricTrends"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPfsMetricTrendsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trend of a metric for the new version of the performance insight feature of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - An ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster is connected to DAS.
//
//   - The new version of the performance insight feature is enabled for the database instance. For more information, see [Performance insight (new version)](https://help.aliyun.com/document_detail/469117.html).
//
// @param request - GetPfsMetricTrendsRequest
//
// @return GetPfsMetricTrendsResponse
func (client *Client) GetPfsMetricTrends(request *GetPfsMetricTrendsRequest) (_result *GetPfsMetricTrendsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPfsMetricTrendsResponse{}
	_body, _err := client.GetPfsMetricTrendsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SQL sample data for the new version of the performance insight feature of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this API operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - An ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster is connected to DAS.
//
//   - The new version of the performance insight feature is enabled for the database instance. For more information, see [Performance insight (new version)](https://help.aliyun.com/document_detail/469117.html).
//
// @param request - GetPfsSqlSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPfsSqlSampleResponse
func (client *Client) GetPfsSqlSampleWithOptions(request *GetPfsSqlSampleRequest, runtime *dara.RuntimeOptions) (_result *GetPfsSqlSampleResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.SqlId) {
		body["SqlId"] = request.SqlId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPfsSqlSample"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPfsSqlSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SQL sample data for the new version of the performance insight feature of a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this API operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - An ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster is connected to DAS.
//
//   - The new version of the performance insight feature is enabled for the database instance. For more information, see [Performance insight (new version)](https://help.aliyun.com/document_detail/469117.html).
//
// @param request - GetPfsSqlSampleRequest
//
// @return GetPfsSqlSampleResponse
func (client *Client) GetPfsSqlSample(request *GetPfsSqlSampleRequest) (_result *GetPfsSqlSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPfsSqlSampleResponse{}
	_body, _err := client.GetPfsSqlSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the full request data generated by the new version of the performance insight feature of a database instance based on the SQL ID.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this API operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - An ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster is connected to DAS.
//
//   - The new version of the performance insight feature is enabled for the database instance. For more information, see [Performance insight (new version)](https://help.aliyun.com/document_detail/469117.html).
//
// @param request - GetPfsSqlSummariesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPfsSqlSummariesResponse
func (client *Client) GetPfsSqlSummariesWithOptions(request *GetPfsSqlSummariesRequest, runtime *dara.RuntimeOptions) (_result *GetPfsSqlSummariesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Asc) {
		body["Asc"] = request.Asc
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Keywords) {
		body["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderBy) {
		body["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SqlId) {
		body["SqlId"] = request.SqlId
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPfsSqlSummaries"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPfsSqlSummariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the full request data generated by the new version of the performance insight feature of a database instance based on the SQL ID.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this API operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - An ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster is connected to DAS.
//
//   - The new version of the performance insight feature is enabled for the database instance. For more information, see [Performance insight (new version)](https://help.aliyun.com/document_detail/469117.html).
//
// @param request - GetPfsSqlSummariesRequest
//
// @return GetPfsSqlSummariesResponse
func (client *Client) GetPfsSqlSummaries(request *GetPfsSqlSummariesRequest) (_result *GetPfsSqlSummariesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPfsSqlSummariesResponse{}
	_body, _err := client.GetPfsSqlSummariesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about SQL templates based on query governance data.
//
// Description:
//
//	  If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeDataStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryOptimizeDataStatsResponse
func (client *Client) GetQueryOptimizeDataStatsWithOptions(request *GetQueryOptimizeDataStatsRequest, runtime *dara.RuntimeOptions) (_result *GetQueryOptimizeDataStatsResponse, _err error) {
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
		Action:      dara.String("GetQueryOptimizeDataStats"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryOptimizeDataStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about SQL templates based on query governance data.
//
// Description:
//
//	  If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeDataStatsRequest
//
// @return GetQueryOptimizeDataStatsResponse
func (client *Client) GetQueryOptimizeDataStats(request *GetQueryOptimizeDataStatsRequest) (_result *GetQueryOptimizeDataStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueryOptimizeDataStatsResponse{}
	_body, _err := client.GetQueryOptimizeDataStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the best-performing and worst-performing instances based on query governance data.
//
// Description:
//
//	  If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeDataTopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryOptimizeDataTopResponse
func (client *Client) GetQueryOptimizeDataTopWithOptions(request *GetQueryOptimizeDataTopRequest, runtime *dara.RuntimeOptions) (_result *GetQueryOptimizeDataTopResponse, _err error) {
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
		Action:      dara.String("GetQueryOptimizeDataTop"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryOptimizeDataTopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the best-performing and worst-performing instances based on query governance data.
//
// Description:
//
//	  If you use an Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeDataTopRequest
//
// @return GetQueryOptimizeDataTopResponse
func (client *Client) GetQueryOptimizeDataTop(request *GetQueryOptimizeDataTopRequest) (_result *GetQueryOptimizeDataTopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueryOptimizeDataTopResponse{}
	_body, _err := client.GetQueryOptimizeDataTopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries query governance trend data.
//
// Description:
//
//	  If you use Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeDataTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryOptimizeDataTrendResponse
func (client *Client) GetQueryOptimizeDataTrendWithOptions(request *GetQueryOptimizeDataTrendRequest, runtime *dara.RuntimeOptions) (_result *GetQueryOptimizeDataTrendResponse, _err error) {
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
		Action:      dara.String("GetQueryOptimizeDataTrend"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryOptimizeDataTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries query governance trend data.
//
// Description:
//
//	  If you use Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeDataTrendRequest
//
// @return GetQueryOptimizeDataTrendResponse
func (client *Client) GetQueryOptimizeDataTrend(request *GetQueryOptimizeDataTrendRequest) (_result *GetQueryOptimizeDataTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueryOptimizeDataTrendResponse{}
	_body, _err := client.GetQueryOptimizeDataTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the failed SQL statements under a SQL template.
//
// Description:
//
//	  If you use Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeExecErrorSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryOptimizeExecErrorSampleResponse
func (client *Client) GetQueryOptimizeExecErrorSampleWithOptions(request *GetQueryOptimizeExecErrorSampleRequest, runtime *dara.RuntimeOptions) (_result *GetQueryOptimizeExecErrorSampleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SqlId) {
		query["SqlId"] = request.SqlId
	}

	if !dara.IsNil(request.Time) {
		query["Time"] = request.Time
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueryOptimizeExecErrorSample"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryOptimizeExecErrorSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the failed SQL statements under a SQL template.
//
// Description:
//
//	  If you use Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeExecErrorSampleRequest
//
// @return GetQueryOptimizeExecErrorSampleResponse
func (client *Client) GetQueryOptimizeExecErrorSample(request *GetQueryOptimizeExecErrorSampleRequest) (_result *GetQueryOptimizeExecErrorSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueryOptimizeExecErrorSampleResponse{}
	_body, _err := client.GetQueryOptimizeExecErrorSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries SQL templates that failed to be executed.
//
// Description:
//
//	  If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeExecErrorStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryOptimizeExecErrorStatsResponse
func (client *Client) GetQueryOptimizeExecErrorStatsWithOptions(request *GetQueryOptimizeExecErrorStatsRequest, runtime *dara.RuntimeOptions) (_result *GetQueryOptimizeExecErrorStatsResponse, _err error) {
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
		Action:      dara.String("GetQueryOptimizeExecErrorStats"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryOptimizeExecErrorStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries SQL templates that failed to be executed.
//
// Description:
//
//	  If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeExecErrorStatsRequest
//
// @return GetQueryOptimizeExecErrorStatsResponse
func (client *Client) GetQueryOptimizeExecErrorStats(request *GetQueryOptimizeExecErrorStatsRequest) (_result *GetQueryOptimizeExecErrorStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueryOptimizeExecErrorStatsResponse{}
	_body, _err := client.GetQueryOptimizeExecErrorStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags added by the query governance feature to specified database instances.
//
// Description:
//
//	  If you use Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeRuleListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryOptimizeRuleListResponse
func (client *Client) GetQueryOptimizeRuleListWithOptions(request *GetQueryOptimizeRuleListRequest, runtime *dara.RuntimeOptions) (_result *GetQueryOptimizeRuleListResponse, _err error) {
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
		Action:      dara.String("GetQueryOptimizeRuleList"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryOptimizeRuleListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags added by the query governance feature to specified database instances.
//
// Description:
//
//	  If you use Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V2.1.8. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V2.1.8 or later.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeRuleListRequest
//
// @return GetQueryOptimizeRuleListResponse
func (client *Client) GetQueryOptimizeRuleList(request *GetQueryOptimizeRuleListRequest) (_result *GetQueryOptimizeRuleListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueryOptimizeRuleListResponse{}
	_body, _err := client.GetQueryOptimizeRuleListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a share URL provided by the query governance feature.
//
// Description:
//
//	  If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeShareUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryOptimizeShareUrlResponse
func (client *Client) GetQueryOptimizeShareUrlWithOptions(request *GetQueryOptimizeShareUrlRequest, runtime *dara.RuntimeOptions) (_result *GetQueryOptimizeShareUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Asc) {
		query["Asc"] = request.Asc
	}

	if !dara.IsNil(request.DbNames) {
		query["DbNames"] = request.DbNames
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Keywords) {
		query["Keywords"] = request.Keywords
	}

	if !dara.IsNil(request.LogicalOperator) {
		query["LogicalOperator"] = request.LogicalOperator
	}

	if !dara.IsNil(request.OnlyOptimizedSql) {
		query["OnlyOptimizedSql"] = request.OnlyOptimizedSql
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Rules) {
		query["Rules"] = request.Rules
	}

	if !dara.IsNil(request.SqlIds) {
		query["SqlIds"] = request.SqlIds
	}

	if !dara.IsNil(request.TagNames) {
		query["TagNames"] = request.TagNames
	}

	if !dara.IsNil(request.Time) {
		query["Time"] = request.Time
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueryOptimizeShareUrl"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryOptimizeShareUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a share URL provided by the query governance feature.
//
// Description:
//
//	  If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call API operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeShareUrlRequest
//
// @return GetQueryOptimizeShareUrlResponse
func (client *Client) GetQueryOptimizeShareUrl(request *GetQueryOptimizeShareUrlRequest) (_result *GetQueryOptimizeShareUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueryOptimizeShareUrlResponse{}
	_body, _err := client.GetQueryOptimizeShareUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries suggestions provided by query governance for optimizing an SQL template.
//
// Description:
//
//	  If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryOptimizeSolutionResponse
func (client *Client) GetQueryOptimizeSolutionWithOptions(request *GetQueryOptimizeSolutionRequest, runtime *dara.RuntimeOptions) (_result *GetQueryOptimizeSolutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RuleIds) {
		query["RuleIds"] = request.RuleIds
	}

	if !dara.IsNil(request.SqlId) {
		query["SqlId"] = request.SqlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueryOptimizeSolution"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryOptimizeSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries suggestions provided by query governance for optimizing an SQL template.
//
// Description:
//
//	  If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeSolutionRequest
//
// @return GetQueryOptimizeSolutionResponse
func (client *Client) GetQueryOptimizeSolution(request *GetQueryOptimizeSolutionRequest) (_result *GetQueryOptimizeSolutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueryOptimizeSolutionResponse{}
	_body, _err := client.GetQueryOptimizeSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of a SQL statement.
//
// Description:
//
//	  If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueryOptimizeTagResponse
func (client *Client) GetQueryOptimizeTagWithOptions(request *GetQueryOptimizeTagRequest, runtime *dara.RuntimeOptions) (_result *GetQueryOptimizeTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SqlId) {
		query["SqlId"] = request.SqlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQueryOptimizeTag"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQueryOptimizeTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of a SQL statement.
//
// Description:
//
//	  If you use Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//		- This operation supports the following database engines:
//
//	    	- ApsaraDB RDS for MySQL
//
//	    	- PolarDB for MySQL
//
//	    	- ApsaraDB RDS for PostgreSQL
//
// @param request - GetQueryOptimizeTagRequest
//
// @return GetQueryOptimizeTagResponse
func (client *Client) GetQueryOptimizeTag(request *GetQueryOptimizeTagRequest) (_result *GetQueryOptimizeTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQueryOptimizeTagResponse{}
	_body, _err := client.GetQueryOptimizeTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current session on an ApsaraDB for Redis instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB for Redis instances.
//
//		- If you use an SDK to call operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// >  This operation cannot be used to query sessions generated in direct connection mode on ApsaraDB for Redis cluster instances.
//
// @param request - GetRedisAllSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRedisAllSessionResponse
func (client *Client) GetRedisAllSessionWithOptions(request *GetRedisAllSessionRequest, runtime *dara.RuntimeOptions) (_result *GetRedisAllSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRedisAllSession"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRedisAllSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current session on an ApsaraDB for Redis instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB for Redis instances.
//
//		- If you use an SDK to call operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
// >  This operation cannot be used to query sessions generated in direct connection mode on ApsaraDB for Redis cluster instances.
//
// @param request - GetRedisAllSessionRequest
//
// @return GetRedisAllSessionResponse
func (client *Client) GetRedisAllSession(request *GetRedisAllSessionRequest) (_result *GetRedisAllSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRedisAllSessionResponse{}
	_body, _err := client.GetRedisAllSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries SQL diagnostics records by pages.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - ApsaraDB RDS for PostgreSQL
//
//   - ApsaraDB RDS for SQL Server
//
//   - PolarDB for MySQL
//
//   - PolarDB for PostgreSQL (Compatible with Oracle)
//
//   - ApsaraDB for MongoDB
//
// >  The minor engine version of the Apsara RDS for PostgreSQL instance must be 20220130 or later. For more information about how to check and update the minor engine version of an ApsaraDB RDS for PostgreSQL instance, see [Update the minor engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146895.html).
//
// @param request - GetRequestDiagnosisPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRequestDiagnosisPageResponse
func (client *Client) GetRequestDiagnosisPageWithOptions(request *GetRequestDiagnosisPageRequest, runtime *dara.RuntimeOptions) (_result *GetRequestDiagnosisPageResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRequestDiagnosisPage"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRequestDiagnosisPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries SQL diagnostics records by pages.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - ApsaraDB RDS for PostgreSQL
//
//   - ApsaraDB RDS for SQL Server
//
//   - PolarDB for MySQL
//
//   - PolarDB for PostgreSQL (Compatible with Oracle)
//
//   - ApsaraDB for MongoDB
//
// >  The minor engine version of the Apsara RDS for PostgreSQL instance must be 20220130 or later. For more information about how to check and update the minor engine version of an ApsaraDB RDS for PostgreSQL instance, see [Update the minor engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146895.html).
//
// @param request - GetRequestDiagnosisPageRequest
//
// @return GetRequestDiagnosisPageResponse
func (client *Client) GetRequestDiagnosisPage(request *GetRequestDiagnosisPageRequest) (_result *GetRequestDiagnosisPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRequestDiagnosisPageResponse{}
	_body, _err := client.GetRequestDiagnosisPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the results of an SQL diagnostics task.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call the API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - You cannot call this operation to query the diagnostic result of the automatic SQL optimization feature.
//
//   - This operation is applicable to the following database engines:
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - PolarDB for MySQL
//
//   - PolarDB for PostgreSQL (Compatible with Oracle)
//
//   - ApsaraDB for MongoDB
//
// >  If your instance is an ApsaraDB RDS for PostgreSQL instance, make sure that the minor engine version of your instance is 20220130 or later. For more information about how to check and update the minor engine version of an ApsaraDB RDS for PostgreSQL instance, see [Update the minor engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146895.html).
//
// @param request - GetRequestDiagnosisResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRequestDiagnosisResultResponse
func (client *Client) GetRequestDiagnosisResultWithOptions(request *GetRequestDiagnosisResultRequest, runtime *dara.RuntimeOptions) (_result *GetRequestDiagnosisResultResponse, _err error) {
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

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SqlId) {
		query["SqlId"] = request.SqlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRequestDiagnosisResult"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRequestDiagnosisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of an SQL diagnostics task.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call the API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - You cannot call this operation to query the diagnostic result of the automatic SQL optimization feature.
//
//   - This operation is applicable to the following database engines:
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - PolarDB for MySQL
//
//   - PolarDB for PostgreSQL (Compatible with Oracle)
//
//   - ApsaraDB for MongoDB
//
// >  If your instance is an ApsaraDB RDS for PostgreSQL instance, make sure that the minor engine version of your instance is 20220130 or later. For more information about how to check and update the minor engine version of an ApsaraDB RDS for PostgreSQL instance, see [Update the minor engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146895.html).
//
// @param request - GetRequestDiagnosisResultRequest
//
// @return GetRequestDiagnosisResultResponse
func (client *Client) GetRequestDiagnosisResult(request *GetRequestDiagnosisResultRequest) (_result *GetRequestDiagnosisResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRequestDiagnosisResultResponse{}
	_body, _err := client.GetRequestDiagnosisResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the throttling rules that are in effect.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - GetRunningSqlConcurrencyControlRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRunningSqlConcurrencyControlRulesResponse
func (client *Client) GetRunningSqlConcurrencyControlRulesWithOptions(request *GetRunningSqlConcurrencyControlRulesRequest, runtime *dara.RuntimeOptions) (_result *GetRunningSqlConcurrencyControlRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRunningSqlConcurrencyControlRules"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRunningSqlConcurrencyControlRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the throttling rules that are in effect.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - GetRunningSqlConcurrencyControlRulesRequest
//
// @return GetRunningSqlConcurrencyControlRulesResponse
func (client *Client) GetRunningSqlConcurrencyControlRules(request *GetRunningSqlConcurrencyControlRulesRequest) (_result *GetRunningSqlConcurrencyControlRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRunningSqlConcurrencyControlRulesResponse{}
	_body, _err := client.GetRunningSqlConcurrencyControlRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a throttling keyword string based on an SQL statement.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - GetSqlConcurrencyControlKeywordsFromSqlTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSqlConcurrencyControlKeywordsFromSqlTextResponse
func (client *Client) GetSqlConcurrencyControlKeywordsFromSqlTextWithOptions(request *GetSqlConcurrencyControlKeywordsFromSqlTextRequest, runtime *dara.RuntimeOptions) (_result *GetSqlConcurrencyControlKeywordsFromSqlTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SqlText) {
		query["SqlText"] = request.SqlText
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSqlConcurrencyControlKeywordsFromSqlText"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSqlConcurrencyControlKeywordsFromSqlTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a throttling keyword string based on an SQL statement.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - GetSqlConcurrencyControlKeywordsFromSqlTextRequest
//
// @return GetSqlConcurrencyControlKeywordsFromSqlTextResponse
func (client *Client) GetSqlConcurrencyControlKeywordsFromSqlText(request *GetSqlConcurrencyControlKeywordsFromSqlTextRequest) (_result *GetSqlConcurrencyControlKeywordsFromSqlTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSqlConcurrencyControlKeywordsFromSqlTextResponse{}
	_body, _err := client.GetSqlConcurrencyControlKeywordsFromSqlTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the throttling rules that are being executed or have been triggered.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - GetSqlConcurrencyControlRulesHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSqlConcurrencyControlRulesHistoryResponse
func (client *Client) GetSqlConcurrencyControlRulesHistoryWithOptions(request *GetSqlConcurrencyControlRulesHistoryRequest, runtime *dara.RuntimeOptions) (_result *GetSqlConcurrencyControlRulesHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSqlConcurrencyControlRulesHistory"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSqlConcurrencyControlRulesHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the throttling rules that are being executed or have been triggered.
//
// Description:
//
// This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL
//
//   - PolarDB for MySQL
//
// @param request - GetSqlConcurrencyControlRulesHistoryRequest
//
// @return GetSqlConcurrencyControlRulesHistoryResponse
func (client *Client) GetSqlConcurrencyControlRulesHistory(request *GetSqlConcurrencyControlRulesHistoryRequest) (_result *GetSqlConcurrencyControlRulesHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSqlConcurrencyControlRulesHistoryResponse{}
	_body, _err := client.GetSqlConcurrencyControlRulesHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries optimization suggestions that are generated by the SQL diagnostics feature of Database Autonomy Service (DAS).
//
// Description:
//
// The SQL diagnostics feature provides optimization suggestions for instances based on diagnostics results. You can use the optimization suggestions to optimize instance indexes. For more information, see [Automatic SQL optimization](https://help.aliyun.com/document_detail/167895.html).
//
// >  You can call this operation to query only the optimization suggestions that are automatically generated by the SQL diagnostics feature.
//
// Before you call this operation, take note of the following items:
//
//   - This operation is applicable to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetSqlOptimizeAdviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSqlOptimizeAdviceResponse
func (client *Client) GetSqlOptimizeAdviceWithOptions(request *GetSqlOptimizeAdviceRequest, runtime *dara.RuntimeOptions) (_result *GetSqlOptimizeAdviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.EndDt) {
		query["EndDt"] = request.EndDt
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartDt) {
		query["StartDt"] = request.StartDt
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSqlOptimizeAdvice"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSqlOptimizeAdviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries optimization suggestions that are generated by the SQL diagnostics feature of Database Autonomy Service (DAS).
//
// Description:
//
// The SQL diagnostics feature provides optimization suggestions for instances based on diagnostics results. You can use the optimization suggestions to optimize instance indexes. For more information, see [Automatic SQL optimization](https://help.aliyun.com/document_detail/167895.html).
//
// >  You can call this operation to query only the optimization suggestions that are automatically generated by the SQL diagnostics feature.
//
// Before you call this operation, take note of the following items:
//
//   - This operation is applicable to ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters.
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetSqlOptimizeAdviceRequest
//
// @return GetSqlOptimizeAdviceResponse
func (client *Client) GetSqlOptimizeAdvice(request *GetSqlOptimizeAdviceRequest) (_result *GetSqlOptimizeAdviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSqlOptimizeAdviceResponse{}
	_body, _err := client.GetSqlOptimizeAdviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status and results of a storage analysis task.
//
// Description:
//
// >  The physical file size indicates the actual size of an obtained file. Only specific deployment modes of database instances support the display of physical file sizes. The statistics on tables are obtained from `information_schema.tables`. Statistics in MySQL are not updated in real time. Therefore, the statistics may be different from the physical file sizes. If you want to obtain the latest data, you can execute the `ANALYZE TABLE` statement on the relevant tables during off-peak hours.
//
//   - This operation is applicable only to ApsaraDB RDS for MySQL instances, PolarDB for MySQL clusters, and ApsaraDB for MongoDB instances.
//
//   - For ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters, this operation works the same as the storage analysis feature of the previous version. Tasks generated by this operation cannot be viewed on the Storage Analysis page of the new version in the Database Autonomy Service (DAS) console. If you want to view the tasks and results, call the related API operation to obtain data and save data to your computer.
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetStorageAnalysisResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStorageAnalysisResultResponse
func (client *Client) GetStorageAnalysisResultWithOptions(request *GetStorageAnalysisResultRequest, runtime *dara.RuntimeOptions) (_result *GetStorageAnalysisResultResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStorageAnalysisResult"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStorageAnalysisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and results of a storage analysis task.
//
// Description:
//
// >  The physical file size indicates the actual size of an obtained file. Only specific deployment modes of database instances support the display of physical file sizes. The statistics on tables are obtained from `information_schema.tables`. Statistics in MySQL are not updated in real time. Therefore, the statistics may be different from the physical file sizes. If you want to obtain the latest data, you can execute the `ANALYZE TABLE` statement on the relevant tables during off-peak hours.
//
//   - This operation is applicable only to ApsaraDB RDS for MySQL instances, PolarDB for MySQL clusters, and ApsaraDB for MongoDB instances.
//
//   - For ApsaraDB RDS for MySQL instances and PolarDB for MySQL clusters, this operation works the same as the storage analysis feature of the previous version. Tasks generated by this operation cannot be viewed on the Storage Analysis page of the new version in the Database Autonomy Service (DAS) console. If you want to view the tasks and results, call the related API operation to obtain data and save data to your computer.
//
//   - If you use an Alibaba Cloud SDK or DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - GetStorageAnalysisResultRequest
//
// @return GetStorageAnalysisResultResponse
func (client *Client) GetStorageAnalysisResult(request *GetStorageAnalysisResultRequest) (_result *GetStorageAnalysisResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStorageAnalysisResultResponse{}
	_body, _err := client.GetStorageAnalysisResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates all sessions on an instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB for Redis.
//
//		- If you use Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V4.3.3. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V1.0.2 or later.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - KillInstanceAllSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillInstanceAllSessionResponse
func (client *Client) KillInstanceAllSessionWithOptions(request *KillInstanceAllSessionRequest, runtime *dara.RuntimeOptions) (_result *KillInstanceAllSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KillInstanceAllSession"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KillInstanceAllSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates all sessions on an instance.
//
// Description:
//
//	  This operation is applicable only to ApsaraDB for Redis.
//
//		- If you use Alibaba Cloud SDK, make sure that the aliyun-sdk-core version is later than V4.3.3. We recommend that you use the latest version.
//
//		- The version of your Database Autonomy Service (DAS) SDK must be V1.0.2 or later.
//
//		- If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - KillInstanceAllSessionRequest
//
// @return KillInstanceAllSessionResponse
func (client *Client) KillInstanceAllSession(request *KillInstanceAllSessionRequest) (_result *KillInstanceAllSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &KillInstanceAllSessionResponse{}
	_body, _err := client.KillInstanceAllSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the auto scaling configurations of an instance.
//
// Description:
//
// You can call this operation to modify the following auto scaling configurations of an instance: **auto scaling for specifications**, **automatic storage expansion**, **automatic bandwidth adjustment**, and **auto scaling for resources**.
//
//   - You can modify the configurations of the **auto scaling feature for specifications*	- for the following types of database instances:
//
//   - PolarDB for MySQL Cluster Edition instances. For more information about the feature and the billing rules, see [Automatic performance scaling](https://help.aliyun.com/document_detail/169686.html).
//
//   - ApsaraDB RDS for MySQL High-availability Edition instances that use standard SSDs or enhanced SSDs (ESSDs). For more information about the feature and the billing rules, see [Automatic performance scaling](https://help.aliyun.com/document_detail/169686.html).
//
//   - You can modify the configurations of the **automatic storage expansion*	- feature for the following types of database instances:
//
//   - ApsaraDB RDS for MySQL High-availability Edition instances that use standard SSDs or ESSDs. For more information about the feature and the billing rules, see [Automatic space expansion](https://help.aliyun.com/document_detail/173345.html).
//
//   - You can modify the configurations of the **automatic bandwidth adjustment*	- feature for the following types of database instances:
//
//   - ApsaraDB for Redis Classic (Local Disk-based) Edition instances. For more information about the feature and the billing rules, see [Automatic bandwidth adjustment](https://help.aliyun.com/document_detail/216312.html).
//
//   - You can modify the configurations of the **auto scaling feature for resources*	- for the following types of database instances:
//
//   - General-purpose ApsaraDB RDS for MySQL Enterprise Edition instances. For more information about the feature and the billing rules, see [Automatic performance scaling](https://help.aliyun.com/document_detail/169686.html).
//
//   - If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - ModifyAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoScalingConfigResponse
func (client *Client) ModifyAutoScalingConfigWithOptions(request *ModifyAutoScalingConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyAutoScalingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.Shard) {
		query["Shard"] = request.Shard
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAutoScalingConfig"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAutoScalingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the auto scaling configurations of an instance.
//
// Description:
//
// You can call this operation to modify the following auto scaling configurations of an instance: **auto scaling for specifications**, **automatic storage expansion**, **automatic bandwidth adjustment**, and **auto scaling for resources**.
//
//   - You can modify the configurations of the **auto scaling feature for specifications*	- for the following types of database instances:
//
//   - PolarDB for MySQL Cluster Edition instances. For more information about the feature and the billing rules, see [Automatic performance scaling](https://help.aliyun.com/document_detail/169686.html).
//
//   - ApsaraDB RDS for MySQL High-availability Edition instances that use standard SSDs or enhanced SSDs (ESSDs). For more information about the feature and the billing rules, see [Automatic performance scaling](https://help.aliyun.com/document_detail/169686.html).
//
//   - You can modify the configurations of the **automatic storage expansion*	- feature for the following types of database instances:
//
//   - ApsaraDB RDS for MySQL High-availability Edition instances that use standard SSDs or ESSDs. For more information about the feature and the billing rules, see [Automatic space expansion](https://help.aliyun.com/document_detail/173345.html).
//
//   - You can modify the configurations of the **automatic bandwidth adjustment*	- feature for the following types of database instances:
//
//   - ApsaraDB for Redis Classic (Local Disk-based) Edition instances. For more information about the feature and the billing rules, see [Automatic bandwidth adjustment](https://help.aliyun.com/document_detail/216312.html).
//
//   - You can modify the configurations of the **auto scaling feature for resources*	- for the following types of database instances:
//
//   - General-purpose ApsaraDB RDS for MySQL Enterprise Edition instances. For more information about the feature and the billing rules, see [Automatic performance scaling](https://help.aliyun.com/document_detail/169686.html).
//
//   - If you use an Alibaba Cloud SDK or Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
// @param request - ModifyAutoScalingConfigRequest
//
// @return ModifyAutoScalingConfigResponse
func (client *Client) ModifyAutoScalingConfig(request *ModifyAutoScalingConfigRequest) (_result *ModifyAutoScalingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAutoScalingConfigResponse{}
	_body, _err := client.ModifyAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or configures Database Autonomy Service (DAS) Enterprise Edition for a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - By default, the latest version of DAS Enterprise Edition that supports the database instance is enabled. For information about the databases and regions that are supported by different versions of DAS Enterprise Edition, see [Editions and supported features](https://help.aliyun.com/document_detail/156204.html).
//
// @param request - ModifySqlLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySqlLogConfigResponse
func (client *Client) ModifySqlLogConfigWithOptions(request *ModifySqlLogConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifySqlLogConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableAudit) {
		query["EnableAudit"] = request.EnableAudit
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		body["Enable"] = request.Enable
	}

	if !dara.IsNil(request.HotRetention) {
		body["HotRetention"] = request.HotRetention
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequestEnable) {
		body["RequestEnable"] = request.RequestEnable
	}

	if !dara.IsNil(request.Retention) {
		body["Retention"] = request.Retention
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySqlLogConfig"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySqlLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or configures Database Autonomy Service (DAS) Enterprise Edition for a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a DAS SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call operations of DAS, you must set the region ID to cn-shanghai.
//
//   - By default, the latest version of DAS Enterprise Edition that supports the database instance is enabled. For information about the databases and regions that are supported by different versions of DAS Enterprise Edition, see [Editions and supported features](https://help.aliyun.com/document_detail/156204.html).
//
// @param request - ModifySqlLogConfigRequest
//
// @return ModifySqlLogConfigResponse
func (client *Client) ModifySqlLogConfig(request *ModifySqlLogConfigRequest) (_result *ModifySqlLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySqlLogConfigResponse{}
	_body, _err := client.ModifySqlLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs a stress testing task.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to effectively handle traffic spikes. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - RunCloudBenchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCloudBenchTaskResponse
func (client *Client) RunCloudBenchTaskWithOptions(request *RunCloudBenchTaskRequest, runtime *dara.RuntimeOptions) (_result *RunCloudBenchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCloudBenchTask"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCloudBenchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a stress testing task.
//
// Description:
//
// Database Autonomy Service (DAS) provides the intelligent stress testing feature. This feature helps you check whether your instance needs to be scaled up to effectively handle traffic spikes. For more information, see [Intelligent stress testing](https://help.aliyun.com/document_detail/155068.html).
//
// @param request - RunCloudBenchTaskRequest
//
// @return RunCloudBenchTaskResponse
func (client *Client) RunCloudBenchTask(request *RunCloudBenchTaskRequest) (_result *RunCloudBenchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunCloudBenchTaskResponse{}
	_body, _err := client.RunCloudBenchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the event subscription settings for a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - Make sure that the database instance that you want to manage is connected to DAS.
//
// @param request - SetEventSubscriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetEventSubscriptionResponse
func (client *Client) SetEventSubscriptionWithOptions(request *SetEventSubscriptionRequest, runtime *dara.RuntimeOptions) (_result *SetEventSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.ChannelType) {
		query["ChannelType"] = request.ChannelType
	}

	if !dara.IsNil(request.ContactGroupName) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.DispatchRule) {
		query["DispatchRule"] = request.DispatchRule
	}

	if !dara.IsNil(request.EventContext) {
		query["EventContext"] = request.EventContext
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.MinInterval) {
		query["MinInterval"] = request.MinInterval
	}

	if !dara.IsNil(request.Severity) {
		query["Severity"] = request.Severity
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetEventSubscription"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetEventSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the event subscription settings for a database instance.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an Alibaba Cloud SDK or a Database Autonomy Service (DAS) SDK to call this operation, we recommend that you use the latest version of the SDK.
//
//   - If you use an SDK to call the API operations of DAS, you must set the region ID to cn-shanghai.
//
//   - Make sure that the database instance that you want to manage is connected to DAS.
//
// @param request - SetEventSubscriptionRequest
//
// @return SetEventSubscriptionResponse
func (client *Client) SetEventSubscription(request *SetEventSubscriptionRequest) (_result *SetEventSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetEventSubscriptionResponse{}
	_body, _err := client.SetEventSubscriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously configures parameters related to the automatic fragment recycling feature for multiple database instances at a time.
//
// Description:
//
// >  Asynchronous calls do not immediately return the complete results. To obtain the complete results, you must use the value of **ResultId*	- returned in the response to re-initiate the call until the value of **isFinish*	- is **true**.***	- In this case, you must call this operation at least twice.
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call the API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - The database instances must be an ApsaraDB RDS for MySQL High-availability Edition instance.
//
//   - DAS Enterprise Edition must be enabled for the database instance. You can call the call [DescribeInstanceDasPro](https://help.aliyun.com/document_detail/413866.html) operation to query whether DAS Enterprise Edition is enabled.
//
//   - The database instance has four or more CPU cores, and **innodb_file_per_table*	- is set to **ON**.
//
// @param request - UpdateAutoResourceOptimizeRulesAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoResourceOptimizeRulesAsyncResponse
func (client *Client) UpdateAutoResourceOptimizeRulesAsyncWithOptions(request *UpdateAutoResourceOptimizeRulesAsyncRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoResourceOptimizeRulesAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.ResultId) {
		query["ResultId"] = request.ResultId
	}

	if !dara.IsNil(request.TableFragmentationRatio) {
		query["TableFragmentationRatio"] = request.TableFragmentationRatio
	}

	if !dara.IsNil(request.TableSpaceSize) {
		query["TableSpaceSize"] = request.TableSpaceSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAutoResourceOptimizeRulesAsync"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAutoResourceOptimizeRulesAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously configures parameters related to the automatic fragment recycling feature for multiple database instances at a time.
//
// Description:
//
// >  Asynchronous calls do not immediately return the complete results. To obtain the complete results, you must use the value of **ResultId*	- returned in the response to re-initiate the call until the value of **isFinish*	- is **true**.***	- In this case, you must call this operation at least twice.
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call the API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - The database instances must be an ApsaraDB RDS for MySQL High-availability Edition instance.
//
//   - DAS Enterprise Edition must be enabled for the database instance. You can call the call [DescribeInstanceDasPro](https://help.aliyun.com/document_detail/413866.html) operation to query whether DAS Enterprise Edition is enabled.
//
//   - The database instance has four or more CPU cores, and **innodb_file_per_table*	- is set to **ON**.
//
// @param request - UpdateAutoResourceOptimizeRulesAsyncRequest
//
// @return UpdateAutoResourceOptimizeRulesAsyncResponse
func (client *Client) UpdateAutoResourceOptimizeRulesAsync(request *UpdateAutoResourceOptimizeRulesAsyncRequest) (_result *UpdateAutoResourceOptimizeRulesAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAutoResourceOptimizeRulesAsyncResponse{}
	_body, _err := client.UpdateAutoResourceOptimizeRulesAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables, modifies, or disables the automatic SQL optimization feature for multiple database instances at a time.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - DAS Enterprise Edition must be enabled for the database instance that you want to manage. To enable DAS Enterprise Edition for a database instance, you can call the [EnableDasPro](https://help.aliyun.com/document_detail/411645.html) operation.
//
//   - The autonomy service must be enabled for the database instance. For more information, see [Autonomy center](https://help.aliyun.com/document_detail/152139.html).
//
//   - This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL High-availability Edition or Enterprise Edition
//
//   - PolarDB for MySQL Cluster Edition
//
// @param request - UpdateAutoSqlOptimizeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoSqlOptimizeStatusResponse
func (client *Client) UpdateAutoSqlOptimizeStatusWithOptions(request *UpdateAutoSqlOptimizeStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoSqlOptimizeStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Instances) {
		query["Instances"] = request.Instances
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAutoSqlOptimizeStatus"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAutoSqlOptimizeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables, modifies, or disables the automatic SQL optimization feature for multiple database instances at a time.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - DAS Enterprise Edition must be enabled for the database instance that you want to manage. To enable DAS Enterprise Edition for a database instance, you can call the [EnableDasPro](https://help.aliyun.com/document_detail/411645.html) operation.
//
//   - The autonomy service must be enabled for the database instance. For more information, see [Autonomy center](https://help.aliyun.com/document_detail/152139.html).
//
//   - This operation supports the following database engines:
//
//   - ApsaraDB RDS for MySQL High-availability Edition or Enterprise Edition
//
//   - PolarDB for MySQL Cluster Edition
//
// @param request - UpdateAutoSqlOptimizeStatusRequest
//
// @return UpdateAutoSqlOptimizeStatusResponse
func (client *Client) UpdateAutoSqlOptimizeStatus(request *UpdateAutoSqlOptimizeStatusRequest) (_result *UpdateAutoSqlOptimizeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAutoSqlOptimizeStatusResponse{}
	_body, _err := client.UpdateAutoSqlOptimizeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously configures parameters related to the automatic SQL throttling feature for multiple database instances at a time.
//
// Description:
//
// >  Asynchronous calls do not immediately return the complete results. To obtain the complete results, you must use the value of **ResultId*	- returned in the response to re-initiate the call until the value of **isFinish*	- is **true**.***	- In this case, you must call this operation at least twice.
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - The autonomy service must be enabled for the database instance that you want to manage. For more information, see [Autonomy center](https://help.aliyun.com/document_detail/152139.html).
//
//   - The database instance that you want to manage must be of one of the following types:
//
//   - ApsaraDB RDS for MySQL High-availability Edition or Enterprise Edition that runs MySQL 5.6, MySQL 5.7, or MySQL 8.0
//
//   - PolarDB for MySQL Cluster Edition that runs MySQL 5.6, MySQL 5.7, or MySQL 8.0
//
// @param request - UpdateAutoThrottleRulesAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoThrottleRulesAsyncResponse
func (client *Client) UpdateAutoThrottleRulesAsyncWithOptions(request *UpdateAutoThrottleRulesAsyncRequest, runtime *dara.RuntimeOptions) (_result *UpdateAutoThrottleRulesAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbnormalDuration) {
		query["AbnormalDuration"] = request.AbnormalDuration
	}

	if !dara.IsNil(request.ActiveSessions) {
		query["ActiveSessions"] = request.ActiveSessions
	}

	if !dara.IsNil(request.AllowThrottleEndTime) {
		query["AllowThrottleEndTime"] = request.AllowThrottleEndTime
	}

	if !dara.IsNil(request.AllowThrottleStartTime) {
		query["AllowThrottleStartTime"] = request.AllowThrottleStartTime
	}

	if !dara.IsNil(request.AutoKillSession) {
		query["AutoKillSession"] = request.AutoKillSession
	}

	if !dara.IsNil(request.ConsoleContext) {
		query["ConsoleContext"] = request.ConsoleContext
	}

	if !dara.IsNil(request.CpuSessionRelation) {
		query["CpuSessionRelation"] = request.CpuSessionRelation
	}

	if !dara.IsNil(request.CpuUsage) {
		query["CpuUsage"] = request.CpuUsage
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.MaxThrottleTime) {
		query["MaxThrottleTime"] = request.MaxThrottleTime
	}

	if !dara.IsNil(request.ResultId) {
		query["ResultId"] = request.ResultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAutoThrottleRulesAsync"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAutoThrottleRulesAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously configures parameters related to the automatic SQL throttling feature for multiple database instances at a time.
//
// Description:
//
// >  Asynchronous calls do not immediately return the complete results. To obtain the complete results, you must use the value of **ResultId*	- returned in the response to re-initiate the call until the value of **isFinish*	- is **true**.***	- In this case, you must call this operation at least twice.
//
// Before you call this operation, take note of the following items:
//
//   - If you use an SDK to call API operations of Database Autonomy Service (DAS), you must set the region ID to cn-shanghai.
//
//   - The autonomy service must be enabled for the database instance that you want to manage. For more information, see [Autonomy center](https://help.aliyun.com/document_detail/152139.html).
//
//   - The database instance that you want to manage must be of one of the following types:
//
//   - ApsaraDB RDS for MySQL High-availability Edition or Enterprise Edition that runs MySQL 5.6, MySQL 5.7, or MySQL 8.0
//
//   - PolarDB for MySQL Cluster Edition that runs MySQL 5.6, MySQL 5.7, or MySQL 8.0
//
// @param request - UpdateAutoThrottleRulesAsyncRequest
//
// @return UpdateAutoThrottleRulesAsyncResponse
func (client *Client) UpdateAutoThrottleRulesAsync(request *UpdateAutoThrottleRulesAsyncRequest) (_result *UpdateAutoThrottleRulesAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAutoThrottleRulesAsyncResponse{}
	_body, _err := client.UpdateAutoThrottleRulesAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) getDasAgentSSEWithSSE_opYieldFunc(_yield chan *GetDasAgentSSEResponse, _yieldErr chan error, request *GetDasAgentSSERequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDasAgentSSE"),
		Version:     dara.String("2020-01-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
