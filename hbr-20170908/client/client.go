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
		"ap-northeast-2-pop":          dara.String("hbr.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("hbr.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("hbr.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("hbr.aliyuncs.com"),
		"cn-edge-1":                   dara.String("hbr.aliyuncs.com"),
		"cn-fujian":                   dara.String("hbr.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("hbr.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("hbr.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("hbr.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("hbr.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("hbr.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("hbr.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("hbr.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("hbr.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("hbr.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("hbr.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("hbr.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("hbr.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("hbr.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("hbr.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("hbr.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("hbr.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("hbr.aliyuncs.com"),
		"cn-wuhan":                    dara.String("hbr.aliyuncs.com"),
		"cn-yushanfang":               dara.String("hbr.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("hbr.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("hbr.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("hbr.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("hbr.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("hbr.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("hbr.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("hbr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Registers a Container Service for Kubernetes (ACK) cluster.
//
// @param request - AddContainerClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddContainerClusterResponse
func (client *Client) AddContainerClusterWithOptions(request *AddContainerClusterRequest, runtime *dara.RuntimeOptions) (_result *AddContainerClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddContainerCluster"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddContainerClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a Container Service for Kubernetes (ACK) cluster.
//
// @param request - AddContainerClusterRequest
//
// @return AddContainerClusterResponse
func (client *Client) AddContainerCluster(request *AddContainerClusterRequest) (_result *AddContainerClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddContainerClusterResponse{}
	_body, _err := client.AddContainerClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a backup job.
//
// @param request - CancelBackupJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelBackupJobResponse
func (client *Client) CancelBackupJobWithOptions(request *CancelBackupJobRequest, runtime *dara.RuntimeOptions) (_result *CancelBackupJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelBackupJob"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelBackupJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a backup job.
//
// @param request - CancelBackupJobRequest
//
// @return CancelBackupJobResponse
func (client *Client) CancelBackupJob(request *CancelBackupJobRequest) (_result *CancelBackupJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelBackupJobResponse{}
	_body, _err := client.CancelBackupJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a restore job.
//
// @param request - CancelRestoreJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRestoreJobResponse
func (client *Client) CancelRestoreJobWithOptions(request *CancelRestoreJobRequest, runtime *dara.RuntimeOptions) (_result *CancelRestoreJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.RestoreId) {
		query["RestoreId"] = request.RestoreId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelRestoreJob"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelRestoreJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a restore job.
//
// @param request - CancelRestoreJobRequest
//
// @return CancelRestoreJobResponse
func (client *Client) CancelRestoreJob(request *CancelRestoreJobRequest) (_result *CancelRestoreJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelRestoreJobResponse{}
	_body, _err := client.CancelRestoreJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which an instance belongs.
//
// Description:
//
//	  In the Cloud Backup console, you can use resource groups to manage resources such as backup vaults, Cloud Backup clients, and SAP HANA instances.
//
//		- A resource is a cloud service entity that you create on Alibaba Cloud, such as an Elastic Compute Service (ECS) instance, a backup vault, or an SAP HANA instance.
//
//		- You can sort resources owned by your Alibaba Cloud account into various resource groups. Resource groups facilitate resource management among multiple projects or applications within your Alibaba Cloud account and simplify permission management.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which an instance belongs.
//
// Description:
//
//	  In the Cloud Backup console, you can use resource groups to manage resources such as backup vaults, Cloud Backup clients, and SAP HANA instances.
//
//		- A resource is a cloud service entity that you create on Alibaba Cloud, such as an Elastic Compute Service (ECS) instance, a backup vault, or an SAP HANA instance.
//
//		- You can sort resources owned by your Alibaba Cloud account into various resource groups. Resource groups facilitate resource management among multiple projects or applications within your Alibaba Cloud account and simplify permission management.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the user has permissions to access the current resource or page.
//
// @param request - CheckRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckRoleResponse
func (client *Client) CheckRoleWithOptions(request *CheckRoleRequest, runtime *dara.RuntimeOptions) (_result *CheckRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckRoleType) {
		query["CheckRoleType"] = request.CheckRoleType
	}

	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckRole"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the user has permissions to access the current resource or page.
//
// @param request - CheckRoleRequest
//
// @return CheckRoleResponse
func (client *Client) CheckRole(request *CheckRoleRequest) (_result *CheckRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckRoleResponse{}
	_body, _err := client.CheckRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a backup job.
//
// @param tmpReq - CreateBackupJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupJobResponse
func (client *Client) CreateBackupJobWithOptions(tmpReq *CreateBackupJobRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBackupJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Detail) {
		request.DetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Detail, dara.String("Detail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ContainerClusterId) {
		query["ContainerClusterId"] = request.ContainerClusterId
	}

	if !dara.IsNil(request.ContainerResources) {
		query["ContainerResources"] = request.ContainerResources
	}

	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.DetailShrink) {
		query["Detail"] = request.DetailShrink
	}

	if !dara.IsNil(request.Exclude) {
		query["Exclude"] = request.Exclude
	}

	if !dara.IsNil(request.Include) {
		query["Include"] = request.Include
	}

	if !dara.IsNil(request.InitiatedByAck) {
		query["InitiatedByAck"] = request.InitiatedByAck
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobName) {
		query["JobName"] = request.JobName
	}

	if !dara.IsNil(request.Options) {
		query["Options"] = request.Options
	}

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SpeedLimit) {
		query["SpeedLimit"] = request.SpeedLimit
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackupJob"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup job.
//
// @param request - CreateBackupJobRequest
//
// @return CreateBackupJobResponse
func (client *Client) CreateBackupJob(request *CreateBackupJobRequest) (_result *CreateBackupJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackupJobResponse{}
	_body, _err := client.CreateBackupJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a backup plan.
//
// Description:
//
// - A backup plan associates data sources with backup policies and other necessary information for backups. After the execution of a backup plan, it generates a backup task that records the progress and results of the backup. If the backup task is successful, a backup snapshot is created. You can use the backup snapshot to create a recovery task.
//
// - A backup plan supports only one type of data source.
//
// - A backup plan supports only a single fixed interval backup cycle strategy.
//
// - A backup plan can back up to only one backup vault.
//
// @param tmpReq - CreateBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupPlanResponse
func (client *Client) CreateBackupPlanWithOptions(tmpReq *CreateBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateBackupPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestDataSourceDetail) {
		request.DestDataSourceDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestDataSourceDetail, dara.String("DestDataSourceDetail"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Detail) {
		request.DetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Detail, dara.String("Detail"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OtsDetail) {
		request.OtsDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OtsDetail, dara.String("OtsDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.Bucket) {
		query["Bucket"] = request.Bucket
	}

	if !dara.IsNil(request.ChangeListPath) {
		query["ChangeListPath"] = request.ChangeListPath
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CreateTime) {
		query["CreateTime"] = request.CreateTime
	}

	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.DestDataSourceDetailShrink) {
		query["DestDataSourceDetail"] = request.DestDataSourceDetailShrink
	}

	if !dara.IsNil(request.DestDataSourceId) {
		query["DestDataSourceId"] = request.DestDataSourceId
	}

	if !dara.IsNil(request.DestSourceType) {
		query["DestSourceType"] = request.DestSourceType
	}

	if !dara.IsNil(request.DetailShrink) {
		query["Detail"] = request.DetailShrink
	}

	if !dara.IsNil(request.Disabled) {
		query["Disabled"] = request.Disabled
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.KeepLatestSnapshots) {
		query["KeepLatestSnapshots"] = request.KeepLatestSnapshots
	}

	if !dara.IsNil(request.PlanName) {
		query["PlanName"] = request.PlanName
	}

	if !dara.IsNil(request.Prefix) {
		query["Prefix"] = request.Prefix
	}

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	if !dara.IsNil(request.Schedule) {
		query["Schedule"] = request.Schedule
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.UdmRegionId) {
		query["UdmRegionId"] = request.UdmRegionId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.Exclude) {
		body["Exclude"] = request.Exclude
	}

	if !dara.IsNil(request.Include) {
		body["Include"] = request.Include
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.OtsDetailShrink) {
		body["OtsDetail"] = request.OtsDetailShrink
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	if !dara.IsNil(request.Rule) {
		body["Rule"] = request.Rule
	}

	if !dara.IsNil(request.SpeedLimit) {
		body["SpeedLimit"] = request.SpeedLimit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a backup plan.
//
// Description:
//
// - A backup plan associates data sources with backup policies and other necessary information for backups. After the execution of a backup plan, it generates a backup task that records the progress and results of the backup. If the backup task is successful, a backup snapshot is created. You can use the backup snapshot to create a recovery task.
//
// - A backup plan supports only one type of data source.
//
// - A backup plan supports only a single fixed interval backup cycle strategy.
//
// - A backup plan can back up to only one backup vault.
//
// @param request - CreateBackupPlanRequest
//
// @return CreateBackupPlanResponse
func (client *Client) CreateBackupPlan(request *CreateBackupPlanRequest) (_result *CreateBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CreateBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs one or more Cloud Backup clients on specified instances.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and pricing of Cloud Backup. For more information, see [Billing methods and billable items](https://help.aliyun.com/document_detail/89062.html).
//
// @param request - CreateClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientsResponse
func (client *Client) CreateClientsWithOptions(request *CreateClientsRequest, runtime *dara.RuntimeOptions) (_result *CreateClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertSetting) {
		query["AlertSetting"] = request.AlertSetting
	}

	if !dara.IsNil(request.ClientInfo) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.UseHttps) {
		query["UseHttps"] = request.UseHttps
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClients"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs one or more Cloud Backup clients on specified instances.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and pricing of Cloud Backup. For more information, see [Billing methods and billable items](https://help.aliyun.com/document_detail/89062.html).
//
// @param request - CreateClientsRequest
//
// @return CreateClientsResponse
func (client *Client) CreateClients(request *CreateClientsRequest) (_result *CreateClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClientsResponse{}
	_body, _err := client.CreateClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a backup plan for an SAP HANA instance.
//
// Description:
//
//	  A backup plan defines the data source, backup policy, and other configurations. After you execute a backup plan, a backup job is generated to record the backup progress and the backup result. If a backup job is completed, a backup snapshot is generated. You can use a backup snapshot to create a restore job.
//
//		- You can specify only one type of data source in a backup plan.
//
//		- You can specify only one interval as a backup cycle in a backup plan.
//
//		- Each backup plan allows you to back up data to only one backup vault.
//
// @param request - CreateHanaBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHanaBackupPlanResponse
func (client *Client) CreateHanaBackupPlanWithOptions(request *CreateHanaBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateHanaBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPrefix) {
		query["BackupPrefix"] = request.BackupPrefix
	}

	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.PlanName) {
		query["PlanName"] = request.PlanName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Schedule) {
		query["Schedule"] = request.Schedule
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHanaBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup plan for an SAP HANA instance.
//
// Description:
//
//	  A backup plan defines the data source, backup policy, and other configurations. After you execute a backup plan, a backup job is generated to record the backup progress and the backup result. If a backup job is completed, a backup snapshot is generated. You can use a backup snapshot to create a restore job.
//
//		- You can specify only one type of data source in a backup plan.
//
//		- You can specify only one interval as a backup cycle in a backup plan.
//
//		- Each backup plan allows you to back up data to only one backup vault.
//
// @param request - CreateHanaBackupPlanRequest
//
// @return CreateHanaBackupPlanResponse
func (client *Client) CreateHanaBackupPlan(request *CreateHanaBackupPlanRequest) (_result *CreateHanaBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHanaBackupPlanResponse{}
	_body, _err := client.CreateHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers an SAP HANA instance.
//
// Description:
//
// To register an SAP HANA instance, you must configure the SAP HANA connection information. After the SAP HANA instance is registered, Cloud Backup installs a backup client on the node of the SAP HANA instance.
//
// @param request - CreateHanaInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHanaInstanceResponse
func (client *Client) CreateHanaInstanceWithOptions(request *CreateHanaInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateHanaInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertSetting) {
		query["AlertSetting"] = request.AlertSetting
	}

	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !dara.IsNil(request.HanaName) {
		query["HanaName"] = request.HanaName
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceNumber) {
		query["InstanceNumber"] = request.InstanceNumber
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.UseSsl) {
		query["UseSsl"] = request.UseSsl
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.ValidateCertificate) {
		query["ValidateCertificate"] = request.ValidateCertificate
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHanaInstance"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHanaInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers an SAP HANA instance.
//
// Description:
//
// To register an SAP HANA instance, you must configure the SAP HANA connection information. After the SAP HANA instance is registered, Cloud Backup installs a backup client on the node of the SAP HANA instance.
//
// @param request - CreateHanaInstanceRequest
//
// @return CreateHanaInstanceResponse
func (client *Client) CreateHanaInstance(request *CreateHanaInstanceRequest) (_result *CreateHanaInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHanaInstanceResponse{}
	_body, _err := client.CreateHanaInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a restore job for an SAP HANA database.
//
// Description:
//
// If you call this operation to restore a database, the database is restored to a specified state. Proceed with caution. For more information, see [Restore databases to an SAP HANA instance](https://help.aliyun.com/document_detail/101178.html).
//
// @param request - CreateHanaRestoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHanaRestoreResponse
func (client *Client) CreateHanaRestoreWithOptions(request *CreateHanaRestoreRequest, runtime *dara.RuntimeOptions) (_result *CreateHanaRestoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupPrefix) {
		query["BackupPrefix"] = request.BackupPrefix
	}

	if !dara.IsNil(request.CheckAccess) {
		query["CheckAccess"] = request.CheckAccess
	}

	if !dara.IsNil(request.ClearLog) {
		query["ClearLog"] = request.ClearLog
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.LogPosition) {
		query["LogPosition"] = request.LogPosition
	}

	if !dara.IsNil(request.MasterClientId) {
		query["MasterClientId"] = request.MasterClientId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.RecoveryPointInTime) {
		query["RecoveryPointInTime"] = request.RecoveryPointInTime
	}

	if !dara.IsNil(request.SidAdmin) {
		query["SidAdmin"] = request.SidAdmin
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceClusterId) {
		query["SourceClusterId"] = request.SourceClusterId
	}

	if !dara.IsNil(request.SystemCopy) {
		query["SystemCopy"] = request.SystemCopy
	}

	if !dara.IsNil(request.UseCatalog) {
		query["UseCatalog"] = request.UseCatalog
	}

	if !dara.IsNil(request.UseDelta) {
		query["UseDelta"] = request.UseDelta
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	if !dara.IsNil(request.VolumeId) {
		query["VolumeId"] = request.VolumeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHanaRestore"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHanaRestoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a restore job for an SAP HANA database.
//
// Description:
//
// If you call this operation to restore a database, the database is restored to a specified state. Proceed with caution. For more information, see [Restore databases to an SAP HANA instance](https://help.aliyun.com/document_detail/101178.html).
//
// @param request - CreateHanaRestoreRequest
//
// @return CreateHanaRestoreResponse
func (client *Client) CreateHanaRestore(request *CreateHanaRestoreRequest) (_result *CreateHanaRestoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHanaRestoreResponse{}
	_body, _err := client.CreateHanaRestoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds one or more data sources to a backup policy.
//
// Description:
//
//	  You can bind data sources to only one policy in each request.
//
//		- Elastic Compute Service (ECS) instances can be bound to only one policy.
//
// @param tmpReq - CreatePolicyBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyBindingsResponse
func (client *Client) CreatePolicyBindingsWithOptions(tmpReq *CreatePolicyBindingsRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyBindingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePolicyBindingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PolicyBindingList) {
		request.PolicyBindingListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PolicyBindingList, dara.String("PolicyBindingList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PolicyBindingListShrink) {
		query["PolicyBindingList"] = request.PolicyBindingListShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicyBindings"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds one or more data sources to a backup policy.
//
// Description:
//
//	  You can bind data sources to only one policy in each request.
//
//		- Elastic Compute Service (ECS) instances can be bound to only one policy.
//
// @param request - CreatePolicyBindingsRequest
//
// @return CreatePolicyBindingsResponse
func (client *Client) CreatePolicyBindings(request *CreatePolicyBindingsRequest) (_result *CreatePolicyBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicyBindingsResponse{}
	_body, _err := client.CreatePolicyBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a backup policy.
//
// Description:
//
// A backup policy records the information required for backup. After you execute a backup policy, a backup job is generated to record the backup progress and the backup result. If a backup job is completed, a backup snapshot is generated. You can use a backup snapshot to create a restore job.
//
//   - A backup policy supports multiple data sources. The data sources can be only Elastic Compute Service (ECS) instances.
//
//   - You can specify only one interval as a backup cycle in a backup policy.
//
//   - Each backup policy allows you to back up data to only one backup vault.
//
// @param tmpReq - CreatePolicyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyV2Response
func (client *Client) CreatePolicyV2WithOptions(tmpReq *CreatePolicyV2Request, runtime *dara.RuntimeOptions) (_result *CreatePolicyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePolicyV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyDescription) {
		body["PolicyDescription"] = request.PolicyDescription
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.RulesShrink) {
		body["Rules"] = request.RulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicyV2"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup policy.
//
// Description:
//
// A backup policy records the information required for backup. After you execute a backup policy, a backup job is generated to record the backup progress and the backup result. If a backup job is completed, a backup snapshot is generated. You can use a backup snapshot to create a restore job.
//
//   - A backup policy supports multiple data sources. The data sources can be only Elastic Compute Service (ECS) instances.
//
//   - You can specify only one interval as a backup cycle in a backup policy.
//
//   - Each backup policy allows you to back up data to only one backup vault.
//
// @param request - CreatePolicyV2Request
//
// @return CreatePolicyV2Response
func (client *Client) CreatePolicyV2(request *CreatePolicyV2Request) (_result *CreatePolicyV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicyV2Response{}
	_body, _err := client.CreatePolicyV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a mirror vault.
//
// Description:
//
// After a backup vault is created, the backup vault is in the INITIALIZING state, and the system automatically runs an initialization task to initialize the backup vault. After the initialization task is completed, the backup vault is in the CREATED state.Call this operation in the region where the mirror vault resides, which is specified by the VaultRegionId parameter.
//
// @param request - CreateReplicationVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReplicationVaultResponse
func (client *Client) CreateReplicationVaultWithOptions(request *CreateReplicationVaultRequest, runtime *dara.RuntimeOptions) (_result *CreateReplicationVaultResponse, _err error) {
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

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.KmsKeyId) {
		query["KmsKeyId"] = request.KmsKeyId
	}

	if !dara.IsNil(request.RedundancyType) {
		query["RedundancyType"] = request.RedundancyType
	}

	if !dara.IsNil(request.ReplicationSourceRegionId) {
		query["ReplicationSourceRegionId"] = request.ReplicationSourceRegionId
	}

	if !dara.IsNil(request.ReplicationSourceVaultId) {
		query["ReplicationSourceVaultId"] = request.ReplicationSourceVaultId
	}

	if !dara.IsNil(request.VaultName) {
		query["VaultName"] = request.VaultName
	}

	if !dara.IsNil(request.VaultRegionId) {
		query["VaultRegionId"] = request.VaultRegionId
	}

	if !dara.IsNil(request.VaultStorageClass) {
		query["VaultStorageClass"] = request.VaultStorageClass
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReplicationVault"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReplicationVaultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mirror vault.
//
// Description:
//
// After a backup vault is created, the backup vault is in the INITIALIZING state, and the system automatically runs an initialization task to initialize the backup vault. After the initialization task is completed, the backup vault is in the CREATED state.Call this operation in the region where the mirror vault resides, which is specified by the VaultRegionId parameter.
//
// @param request - CreateReplicationVaultRequest
//
// @return CreateReplicationVaultResponse
func (client *Client) CreateReplicationVault(request *CreateReplicationVaultRequest) (_result *CreateReplicationVaultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateReplicationVaultResponse{}
	_body, _err := client.CreateReplicationVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a restore job.
//
// Description:
//
// - Create a restore job based on the selected snapshot and the restore destination.
//
// - Currently, the data source type must match the restore destination data source type.
//
// @param tmpReq - CreateRestoreJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRestoreJobResponse
func (client *Client) CreateRestoreJobWithOptions(tmpReq *CreateRestoreJobRequest, runtime *dara.RuntimeOptions) (_result *CreateRestoreJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRestoreJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FailbackDetail) {
		request.FailbackDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FailbackDetail, dara.String("FailbackDetail"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OtsDetail) {
		request.OtsDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OtsDetail, dara.String("OtsDetail"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UdmDetail) {
		request.UdmDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UdmDetail, dara.String("UdmDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.FailbackDetailShrink) {
		query["FailbackDetail"] = request.FailbackDetailShrink
	}

	if !dara.IsNil(request.InitiatedByAck) {
		query["InitiatedByAck"] = request.InitiatedByAck
	}

	if !dara.IsNil(request.Options) {
		query["Options"] = request.Options
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	if !dara.IsNil(request.SnapshotHash) {
		query["SnapshotHash"] = request.SnapshotHash
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.TargetBucket) {
		query["TargetBucket"] = request.TargetBucket
	}

	if !dara.IsNil(request.TargetContainer) {
		query["TargetContainer"] = request.TargetContainer
	}

	if !dara.IsNil(request.TargetContainerClusterId) {
		query["TargetContainerClusterId"] = request.TargetContainerClusterId
	}

	if !dara.IsNil(request.TargetCreateTime) {
		query["TargetCreateTime"] = request.TargetCreateTime
	}

	if !dara.IsNil(request.TargetFileSystemId) {
		query["TargetFileSystemId"] = request.TargetFileSystemId
	}

	if !dara.IsNil(request.TargetInstanceName) {
		query["TargetInstanceName"] = request.TargetInstanceName
	}

	if !dara.IsNil(request.TargetPrefix) {
		query["TargetPrefix"] = request.TargetPrefix
	}

	if !dara.IsNil(request.TargetTableName) {
		query["TargetTableName"] = request.TargetTableName
	}

	if !dara.IsNil(request.TargetTime) {
		query["TargetTime"] = request.TargetTime
	}

	if !dara.IsNil(request.UdmDetailShrink) {
		query["UdmDetail"] = request.UdmDetailShrink
	}

	if !dara.IsNil(request.UdmRegionId) {
		query["UdmRegionId"] = request.UdmRegionId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Exclude) {
		body["Exclude"] = request.Exclude
	}

	if !dara.IsNil(request.Include) {
		body["Include"] = request.Include
	}

	if !dara.IsNil(request.OtsDetailShrink) {
		body["OtsDetail"] = request.OtsDetailShrink
	}

	if !dara.IsNil(request.TargetInstanceId) {
		body["TargetInstanceId"] = request.TargetInstanceId
	}

	if !dara.IsNil(request.TargetPath) {
		body["TargetPath"] = request.TargetPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRestoreJob"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRestoreJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a restore job.
//
// Description:
//
// - Create a restore job based on the selected snapshot and the restore destination.
//
// - Currently, the data source type must match the restore destination data source type.
//
// @param request - CreateRestoreJobRequest
//
// @return CreateRestoreJobResponse
func (client *Client) CreateRestoreJob(request *CreateRestoreJobRequest) (_result *CreateRestoreJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRestoreJobResponse{}
	_body, _err := client.CreateRestoreJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates the parameters and signature required for a file upload URL.
//
// Description:
//
// 1.  You can directly upload a file to Object Storage Service (OSS) by using a form based on the returned value of this operation.
//
// 2.  For more information about how to upload a file to OSS by using a form, see OSS documentation.
//
// 3.  The system periodically deletes files that are uploaded to OSS.
//
// @param request - CreateTempFileUploadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTempFileUploadUrlResponse
func (client *Client) CreateTempFileUploadUrlWithOptions(request *CreateTempFileUploadUrlRequest, runtime *dara.RuntimeOptions) (_result *CreateTempFileUploadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTempFileUploadUrl"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTempFileUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates the parameters and signature required for a file upload URL.
//
// Description:
//
// 1.  You can directly upload a file to Object Storage Service (OSS) by using a form based on the returned value of this operation.
//
// 2.  For more information about how to upload a file to OSS by using a form, see OSS documentation.
//
// 3.  The system periodically deletes files that are uploaded to OSS.
//
// @param request - CreateTempFileUploadUrlRequest
//
// @return CreateTempFileUploadUrlResponse
func (client *Client) CreateTempFileUploadUrl(request *CreateTempFileUploadUrlRequest) (_result *CreateTempFileUploadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTempFileUploadUrlResponse{}
	_body, _err := client.CreateTempFileUploadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a backup vault.
//
// Description:
//
//	  Each Alibaba Cloud account can create up to 100 backup vaults.
//
//		- After a backup vault is created, the backup vault is in the INITIALIZING state, and the system automatically runs an initialization task to initialize the backup vault. After the initialization task is completed, the backup vault is in the CREATED state. A backup job can use a backup vault to store backup data only if the backup vault is in the CREATED state.
//
//	    **
//
//	    **Note*	- Before you call this operation, make sure that you fully understand the billing of Cloud Backup.
//
// @param request - CreateVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVaultResponse
func (client *Client) CreateVaultWithOptions(request *CreateVaultRequest, runtime *dara.RuntimeOptions) (_result *CreateVaultResponse, _err error) {
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

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.KmsKeyId) {
		query["KmsKeyId"] = request.KmsKeyId
	}

	if !dara.IsNil(request.VaultName) {
		query["VaultName"] = request.VaultName
	}

	if !dara.IsNil(request.VaultRegionId) {
		query["VaultRegionId"] = request.VaultRegionId
	}

	if !dara.IsNil(request.VaultStorageClass) {
		query["VaultStorageClass"] = request.VaultStorageClass
	}

	if !dara.IsNil(request.VaultType) {
		query["VaultType"] = request.VaultType
	}

	if !dara.IsNil(request.WormEnabled) {
		query["WormEnabled"] = request.WormEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVault"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVaultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backup vault.
//
// Description:
//
//	  Each Alibaba Cloud account can create up to 100 backup vaults.
//
//		- After a backup vault is created, the backup vault is in the INITIALIZING state, and the system automatically runs an initialization task to initialize the backup vault. After the initialization task is completed, the backup vault is in the CREATED state. A backup job can use a backup vault to store backup data only if the backup vault is in the CREATED state.
//
//	    **
//
//	    **Note*	- Before you call this operation, make sure that you fully understand the billing of Cloud Backup.
//
// @param request - CreateVaultRequest
//
// @return CreateVaultResponse
func (client *Client) CreateVault(request *CreateVaultRequest) (_result *CreateVaultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVaultResponse{}
	_body, _err := client.CreateVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the Elastic Compute Service (ECS) instance that is used for restoration only in ECS Backup Essential Edition.
//
// @param tmpReq - DeleteAirEcsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAirEcsInstanceResponse
func (client *Client) DeleteAirEcsInstanceWithOptions(tmpReq *DeleteAirEcsInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteAirEcsInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteAirEcsInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UninstallClientSourceTypes) {
		request.UninstallClientSourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UninstallClientSourceTypes, dara.String("UninstallClientSourceTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !dara.IsNil(request.UninstallClientSourceTypesShrink) {
		query["UninstallClientSourceTypes"] = request.UninstallClientSourceTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAirEcsInstance"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAirEcsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the Elastic Compute Service (ECS) instance that is used for restoration only in ECS Backup Essential Edition.
//
// @param request - DeleteAirEcsInstanceRequest
//
// @return DeleteAirEcsInstanceResponse
func (client *Client) DeleteAirEcsInstance(request *DeleteAirEcsInstanceRequest) (_result *DeleteAirEcsInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAirEcsInstanceResponse{}
	_body, _err := client.DeleteAirEcsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Cloud Backup client.
//
// Description:
//
//	  You cannot delete the active Cloud Backup clients that receive heartbeat packets within 1 hour. You can call the UninstallBackupClients operation to uninstall a Cloud Backup client. Then, the client becomes inactive.
//
//		- When you perform this operation, resources that are associated with the client are also deleted, including:
//
//	    	- Backup plans
//
//	    	- Backup jobs
//
//	    	- Snapshots
//
// @param request - DeleteBackupClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupClientResponse
func (client *Client) DeleteBackupClientWithOptions(request *DeleteBackupClientRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackupClient"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Cloud Backup client.
//
// Description:
//
//	  You cannot delete the active Cloud Backup clients that receive heartbeat packets within 1 hour. You can call the UninstallBackupClients operation to uninstall a Cloud Backup client. Then, the client becomes inactive.
//
//		- When you perform this operation, resources that are associated with the client are also deleted, including:
//
//	    	- Backup plans
//
//	    	- Backup jobs
//
//	    	- Snapshots
//
// @param request - DeleteBackupClientRequest
//
// @return DeleteBackupClientResponse
func (client *Client) DeleteBackupClient(request *DeleteBackupClientRequest) (_result *DeleteBackupClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackupClientResponse{}
	_body, _err := client.DeleteBackupClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the resources that are related to one or more HBR clients.
//
// Description:
//
// This operation deletes only the resources that are related to HBR clients. The resources include backup plans, backup jobs, and backup snapshots. The operation does not delete HBR clients.
//
// @param tmpReq - DeleteBackupClientResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupClientResourceResponse
func (client *Client) DeleteBackupClientResourceWithOptions(tmpReq *DeleteBackupClientResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupClientResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteBackupClientResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClientIds) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, dara.String("ClientIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIdsShrink) {
		query["ClientIds"] = request.ClientIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackupClientResource"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupClientResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the resources that are related to one or more HBR clients.
//
// Description:
//
// This operation deletes only the resources that are related to HBR clients. The resources include backup plans, backup jobs, and backup snapshots. The operation does not delete HBR clients.
//
// @param request - DeleteBackupClientResourceRequest
//
// @return DeleteBackupClientResourceResponse
func (client *Client) DeleteBackupClientResource(request *DeleteBackupClientResourceRequest) (_result *DeleteBackupClientResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackupClientResourceResponse{}
	_body, _err := client.DeleteBackupClientResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a backup plan.
//
// Description:
//
//	  If you delete a backup plan, the backup jobs are also deleted.
//
//		- If you delete a backup plan, the created snapshot files are not deleted.
//
// @param request - DeleteBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupPlanResponse
func (client *Client) DeleteBackupPlanWithOptions(request *DeleteBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.RequireNoRunningJobs) {
		query["RequireNoRunningJobs"] = request.RequireNoRunningJobs
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a backup plan.
//
// Description:
//
//	  If you delete a backup plan, the backup jobs are also deleted.
//
//		- If you delete a backup plan, the created snapshot files are not deleted.
//
// @param request - DeleteBackupPlanRequest
//
// @return DeleteBackupPlanResponse
func (client *Client) DeleteBackupPlan(request *DeleteBackupPlanRequest) (_result *DeleteBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackupPlanResponse{}
	_body, _err := client.DeleteBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientResponse
func (client *Client) DeleteClientWithOptions(request *DeleteClientRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClient"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteClientRequest
//
// @return DeleteClientResponse
func (client *Client) DeleteClient(request *DeleteClientRequest) (_result *DeleteClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClientResponse{}
	_body, _err := client.DeleteClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an SAP HANA backup plan.
//
// @param request - DeleteHanaBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHanaBackupPlanResponse
func (client *Client) DeleteHanaBackupPlanWithOptions(request *DeleteHanaBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *DeleteHanaBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHanaBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an SAP HANA backup plan.
//
// @param request - DeleteHanaBackupPlanRequest
//
// @return DeleteHanaBackupPlanResponse
func (client *Client) DeleteHanaBackupPlan(request *DeleteHanaBackupPlanRequest) (_result *DeleteHanaBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHanaBackupPlanResponse{}
	_body, _err := client.DeleteHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an SAP HANA instance.
//
// Description:
//
// If you delete an SAP HANA instance, the existing backup data is also deleted and the running backup and restore jobs fail to be completed. Before you delete the SAP HANA instance, make sure that you no longer need the backup data of the instance and no backup or restore jobs are running for the instance. To delete an SAP HANA instance, you must specify the security identifier (SID) of the instance. The SID is three characters in length and starts with a letter. For more information, see [How to find sid user and instance number of HANA db?](https://answers.sap.com/questions/555192/how-to-find-sid-user-and-instance-number-of-hana-d.html?)
//
// @param request - DeleteHanaInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHanaInstanceResponse
func (client *Client) DeleteHanaInstanceWithOptions(request *DeleteHanaInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteHanaInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHanaInstance"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHanaInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an SAP HANA instance.
//
// Description:
//
// If you delete an SAP HANA instance, the existing backup data is also deleted and the running backup and restore jobs fail to be completed. Before you delete the SAP HANA instance, make sure that you no longer need the backup data of the instance and no backup or restore jobs are running for the instance. To delete an SAP HANA instance, you must specify the security identifier (SID) of the instance. The SID is three characters in length and starts with a letter. For more information, see [How to find sid user and instance number of HANA db?](https://answers.sap.com/questions/555192/how-to-find-sid-user-and-instance-number-of-hana-d.html?)
//
// @param request - DeleteHanaInstanceRequest
//
// @return DeleteHanaInstanceResponse
func (client *Client) DeleteHanaInstance(request *DeleteHanaInstanceRequest) (_result *DeleteHanaInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHanaInstanceResponse{}
	_body, _err := client.DeleteHanaInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates one or more data sources from a backup policy. After you disassociate the data sources from the backup policy, the backup policy no longer protects the data sources. Proceed with caution.
//
// @param tmpReq - DeletePolicyBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyBindingResponse
func (client *Client) DeletePolicyBindingWithOptions(tmpReq *DeletePolicyBindingRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeletePolicyBindingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSourceIds) {
		request.DataSourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSourceIds, dara.String("DataSourceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceIdsShrink) {
		body["DataSourceIds"] = request.DataSourceIdsShrink
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicyBinding"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates one or more data sources from a backup policy. After you disassociate the data sources from the backup policy, the backup policy no longer protects the data sources. Proceed with caution.
//
// @param request - DeletePolicyBindingRequest
//
// @return DeletePolicyBindingResponse
func (client *Client) DeletePolicyBinding(request *DeletePolicyBindingRequest) (_result *DeletePolicyBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyBindingResponse{}
	_body, _err := client.DeletePolicyBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a backup policy.
//
// Description:
//
// If you delete a backup policy, the backup policy is disassociated with all data sources. Proceed with caution.
//
// @param request - DeletePolicyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyV2Response
func (client *Client) DeletePolicyV2WithOptions(request *DeletePolicyV2Request, runtime *dara.RuntimeOptions) (_result *DeletePolicyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicyV2"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a backup policy.
//
// Description:
//
// If you delete a backup policy, the backup policy is disassociated with all data sources. Proceed with caution.
//
// @param request - DeletePolicyV2Request
//
// @return DeletePolicyV2Response
func (client *Client) DeletePolicyV2(request *DeletePolicyV2Request) (_result *DeletePolicyV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyV2Response{}
	_body, _err := client.DeletePolicyV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a backup snapshot.
//
// @param request - DeleteSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnapshot"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a backup snapshot.
//
// @param request - DeleteSnapshotRequest
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a protected disk.
//
// @param request - DeleteUdmDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdmDiskResponse
func (client *Client) DeleteUdmDiskWithOptions(request *DeleteUdmDiskRequest, runtime *dara.RuntimeOptions) (_result *DeleteUdmDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUdmDisk"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUdmDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a protected disk.
//
// @param request - DeleteUdmDiskRequest
//
// @return DeleteUdmDiskResponse
func (client *Client) DeleteUdmDisk(request *DeleteUdmDiskRequest) (_result *DeleteUdmDiskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUdmDiskResponse{}
	_body, _err := client.DeleteUdmDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops protection for Elastic Compute Service (ECS) instance backup.
//
// @param request - DeleteUdmEcsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUdmEcsInstanceResponse
func (client *Client) DeleteUdmEcsInstanceWithOptions(request *DeleteUdmEcsInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteUdmEcsInstanceResponse, _err error) {
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
		Action:      dara.String("DeleteUdmEcsInstance"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUdmEcsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops protection for Elastic Compute Service (ECS) instance backup.
//
// @param request - DeleteUdmEcsInstanceRequest
//
// @return DeleteUdmEcsInstanceResponse
func (client *Client) DeleteUdmEcsInstance(request *DeleteUdmEcsInstanceRequest) (_result *DeleteUdmEcsInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUdmEcsInstanceResponse{}
	_body, _err := client.DeleteUdmEcsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a backup vault.
//
// Description:
//
//	  You cannot delete a backup vault within 2 hours after the backup vault is created or a backup vault that is in the INITIALIZING state.
//
//		- After you delete a backup vault, all resources that are associated with the backup vault are deleted. The resources include the Cloud Backup client of the old version, backup plans, backup jobs, snapshots, and restore jobs.
//
// @param request - DeleteVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVaultResponse
func (client *Client) DeleteVaultWithOptions(request *DeleteVaultRequest, runtime *dara.RuntimeOptions) (_result *DeleteVaultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVault"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVaultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a backup vault.
//
// Description:
//
//	  You cannot delete a backup vault within 2 hours after the backup vault is created or a backup vault that is in the INITIALIZING state.
//
//		- After you delete a backup vault, all resources that are associated with the backup vault are deleted. The resources include the Cloud Backup client of the old version, backup plans, backup jobs, snapshots, and restore jobs.
//
// @param request - DeleteVaultRequest
//
// @return DeleteVaultResponse
func (client *Client) DeleteVault(request *DeleteVaultRequest) (_result *DeleteVaultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVaultResponse{}
	_body, _err := client.DeleteVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more HBR clients that meet the specified conditions.
//
// @param tmpReq - DescribeBackupClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupClientsResponse
func (client *Client) DescribeBackupClientsWithOptions(tmpReq *DescribeBackupClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeBackupClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClientIds) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, dara.String("ClientIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientIdsShrink) {
		body["ClientIds"] = request.ClientIdsShrink
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		body["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupClients"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more HBR clients that meet the specified conditions.
//
// @param request - DescribeBackupClientsRequest
//
// @return DescribeBackupClientsResponse
func (client *Client) DescribeBackupClients(request *DescribeBackupClientsRequest) (_result *DescribeBackupClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupClientsResponse{}
	_body, _err := client.DescribeBackupClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more backup jobs that meet the specified conditions.
//
// @param request - DescribeBackupJobs2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupJobs2Response
func (client *Client) DescribeBackupJobs2WithOptions(request *DescribeBackupJobs2Request, runtime *dara.RuntimeOptions) (_result *DescribeBackupJobs2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortDirection) {
		query["SortDirection"] = request.SortDirection
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupJobs2"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupJobs2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more backup jobs that meet the specified conditions.
//
// @param request - DescribeBackupJobs2Request
//
// @return DescribeBackupJobs2Response
func (client *Client) DescribeBackupJobs2(request *DescribeBackupJobs2Request) (_result *DescribeBackupJobs2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupJobs2Response{}
	_body, _err := client.DescribeBackupJobs2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more backup plans that meet the specified conditions.
//
// @param request - DescribeBackupPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPlansResponse
func (client *Client) DescribeBackupPlansWithOptions(request *DescribeBackupPlansRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPlans"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more backup plans that meet the specified conditions.
//
// @param request - DescribeBackupPlansRequest
//
// @return DescribeBackupPlansResponse
func (client *Client) DescribeBackupPlans(request *DescribeBackupPlansRequest) (_result *DescribeBackupPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackupPlansResponse{}
	_body, _err := client.DescribeBackupPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more Cloud Backup clients that meet the specified conditions.
//
// Description:
//
// This operation is applicable only to SAP HANA backup. For Cloud Backup clients of other data sources, call the DescribeBackupClients operation.
//
// @param request - DescribeClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientsResponse
func (client *Client) DescribeClientsWithOptions(request *DescribeClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClients"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more Cloud Backup clients that meet the specified conditions.
//
// Description:
//
// This operation is applicable only to SAP HANA backup. For Cloud Backup clients of other data sources, call the DescribeBackupClients operation.
//
// @param request - DescribeClientsRequest
//
// @return DescribeClientsResponse
func (client *Client) DescribeClients(request *DescribeClientsRequest) (_result *DescribeClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClientsResponse{}
	_body, _err := client.DescribeClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more container clusters that meet the specified conditions.
//
// Description:
//
// You can call this operation to query only Container Service for Kubernetes (ACK) clusters.
//
// @param request - DescribeContainerClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeContainerClusterResponse
func (client *Client) DescribeContainerClusterWithOptions(request *DescribeContainerClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeContainerClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
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
		Action:      dara.String("DescribeContainerCluster"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeContainerClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more container clusters that meet the specified conditions.
//
// Description:
//
// You can call this operation to query only Container Service for Kubernetes (ACK) clusters.
//
// @param request - DescribeContainerClusterRequest
//
// @return DescribeContainerClusterResponse
func (client *Client) DescribeContainerCluster(request *DescribeContainerClusterRequest) (_result *DescribeContainerClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeContainerClusterResponse{}
	_body, _err := client.DescribeContainerClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the accounts used in cross-account backup.
//
// @param request - DescribeCrossAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrossAccountsResponse
func (client *Client) DescribeCrossAccountsWithOptions(request *DescribeCrossAccountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCrossAccountsResponse, _err error) {
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
		Action:      dara.String("DescribeCrossAccounts"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCrossAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the accounts used in cross-account backup.
//
// @param request - DescribeCrossAccountsRequest
//
// @return DescribeCrossAccountsResponse
func (client *Client) DescribeCrossAccounts(request *DescribeCrossAccountsRequest) (_result *DescribeCrossAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCrossAccountsResponse{}
	_body, _err := client.DescribeCrossAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more SAP HANA backup plans that meet the specified conditions.
//
// @param request - DescribeHanaBackupPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHanaBackupPlansResponse
func (client *Client) DescribeHanaBackupPlansWithOptions(request *DescribeHanaBackupPlansRequest, runtime *dara.RuntimeOptions) (_result *DescribeHanaBackupPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHanaBackupPlans"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHanaBackupPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more SAP HANA backup plans that meet the specified conditions.
//
// @param request - DescribeHanaBackupPlansRequest
//
// @return DescribeHanaBackupPlansResponse
func (client *Client) DescribeHanaBackupPlans(request *DescribeHanaBackupPlansRequest) (_result *DescribeHanaBackupPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHanaBackupPlansResponse{}
	_body, _err := client.DescribeHanaBackupPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup parameters of an SAP HANA database.
//
// Description:
//
// If you want to query the backup retention period of a database, you can call the DescribeHanaRetentionSetting operation.
//
// @param request - DescribeHanaBackupSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHanaBackupSettingResponse
func (client *Client) DescribeHanaBackupSettingWithOptions(request *DescribeHanaBackupSettingRequest, runtime *dara.RuntimeOptions) (_result *DescribeHanaBackupSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHanaBackupSetting"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHanaBackupSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup parameters of an SAP HANA database.
//
// Description:
//
// If you want to query the backup retention period of a database, you can call the DescribeHanaRetentionSetting operation.
//
// @param request - DescribeHanaBackupSettingRequest
//
// @return DescribeHanaBackupSettingResponse
func (client *Client) DescribeHanaBackupSetting(request *DescribeHanaBackupSettingRequest) (_result *DescribeHanaBackupSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHanaBackupSettingResponse{}
	_body, _err := client.DescribeHanaBackupSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more SAP HANA backups that meet the specified conditions.
//
// Description:
//
// After you call the DescribeHanaBackupsAsync operation to query the SAP HANA backups that meet the specified conditions, call the DescribeTask operation to query the final result.
//
// @param request - DescribeHanaBackupsAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHanaBackupsAsyncResponse
func (client *Client) DescribeHanaBackupsAsyncWithOptions(request *DescribeHanaBackupsAsyncRequest, runtime *dara.RuntimeOptions) (_result *DescribeHanaBackupsAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.IncludeDifferential) {
		query["IncludeDifferential"] = request.IncludeDifferential
	}

	if !dara.IsNil(request.IncludeIncremental) {
		query["IncludeIncremental"] = request.IncludeIncremental
	}

	if !dara.IsNil(request.IncludeLog) {
		query["IncludeLog"] = request.IncludeLog
	}

	if !dara.IsNil(request.LogPosition) {
		query["LogPosition"] = request.LogPosition
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RecoveryPointInTime) {
		query["RecoveryPointInTime"] = request.RecoveryPointInTime
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceClusterId) {
		query["SourceClusterId"] = request.SourceClusterId
	}

	if !dara.IsNil(request.SystemCopy) {
		query["SystemCopy"] = request.SystemCopy
	}

	if !dara.IsNil(request.UseBackint) {
		query["UseBackint"] = request.UseBackint
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	if !dara.IsNil(request.VolumeId) {
		query["VolumeId"] = request.VolumeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHanaBackupsAsync"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHanaBackupsAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more SAP HANA backups that meet the specified conditions.
//
// Description:
//
// After you call the DescribeHanaBackupsAsync operation to query the SAP HANA backups that meet the specified conditions, call the DescribeTask operation to query the final result.
//
// @param request - DescribeHanaBackupsAsyncRequest
//
// @return DescribeHanaBackupsAsyncResponse
func (client *Client) DescribeHanaBackupsAsync(request *DescribeHanaBackupsAsyncRequest) (_result *DescribeHanaBackupsAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHanaBackupsAsyncResponse{}
	_body, _err := client.DescribeHanaBackupsAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about SAP HANA databases.
//
// Description:
//
// After you register an SAP HANA instance and install a Cloud Backup client on the instance, you can call this operation to query the information about SAP HANA databases. You can call the StartHanaDatabaseAsync operation to start a database and call the StopHanaDatabaseAsync operation to stop a database.
//
// @param request - DescribeHanaDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHanaDatabasesResponse
func (client *Client) DescribeHanaDatabasesWithOptions(request *DescribeHanaDatabasesRequest, runtime *dara.RuntimeOptions) (_result *DescribeHanaDatabasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHanaDatabases"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHanaDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about SAP HANA databases.
//
// Description:
//
// After you register an SAP HANA instance and install a Cloud Backup client on the instance, you can call this operation to query the information about SAP HANA databases. You can call the StartHanaDatabaseAsync operation to start a database and call the StopHanaDatabaseAsync operation to stop a database.
//
// @param request - DescribeHanaDatabasesRequest
//
// @return DescribeHanaDatabasesResponse
func (client *Client) DescribeHanaDatabases(request *DescribeHanaDatabasesRequest) (_result *DescribeHanaDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHanaDatabasesResponse{}
	_body, _err := client.DescribeHanaDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more SAP HANA instances that meet the specified conditions.
//
// @param request - DescribeHanaInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHanaInstancesResponse
func (client *Client) DescribeHanaInstancesWithOptions(request *DescribeHanaInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeHanaInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHanaInstances"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHanaInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more SAP HANA instances that meet the specified conditions.
//
// @param request - DescribeHanaInstancesRequest
//
// @return DescribeHanaInstancesResponse
func (client *Client) DescribeHanaInstances(request *DescribeHanaInstancesRequest) (_result *DescribeHanaInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHanaInstancesResponse{}
	_body, _err := client.DescribeHanaInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more SAP HANA restore jobs that meet the specified conditions.
//
// @param request - DescribeHanaRestoresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHanaRestoresResponse
func (client *Client) DescribeHanaRestoresWithOptions(request *DescribeHanaRestoresRequest, runtime *dara.RuntimeOptions) (_result *DescribeHanaRestoresResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RestoreId) {
		query["RestoreId"] = request.RestoreId
	}

	if !dara.IsNil(request.RestoreStatus) {
		query["RestoreStatus"] = request.RestoreStatus
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHanaRestores"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHanaRestoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more SAP HANA restore jobs that meet the specified conditions.
//
// @param request - DescribeHanaRestoresRequest
//
// @return DescribeHanaRestoresResponse
func (client *Client) DescribeHanaRestores(request *DescribeHanaRestoresRequest) (_result *DescribeHanaRestoresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHanaRestoresResponse{}
	_body, _err := client.DescribeHanaRestoresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup retention period of an SAP HANA database.
//
// Description:
//
//	  If you want to query the backup parameters of a database, you can call the DescribeHanaBackupSetting operation.
//
//		- Cloud Backup deletes the expired catalogs and data that are related to Backint and file backup. The deleted catalogs and data cannot be restored. We recommend that you set the retention period based on your business requirements.
//
// @param request - DescribeHanaRetentionSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHanaRetentionSettingResponse
func (client *Client) DescribeHanaRetentionSettingWithOptions(request *DescribeHanaRetentionSettingRequest, runtime *dara.RuntimeOptions) (_result *DescribeHanaRetentionSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHanaRetentionSetting"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHanaRetentionSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup retention period of an SAP HANA database.
//
// Description:
//
//	  If you want to query the backup parameters of a database, you can call the DescribeHanaBackupSetting operation.
//
//		- Cloud Backup deletes the expired catalogs and data that are related to Backint and file backup. The deleted catalogs and data cannot be restored. We recommend that you set the retention period based on your business requirements.
//
// @param request - DescribeHanaRetentionSettingRequest
//
// @return DescribeHanaRetentionSettingResponse
func (client *Client) DescribeHanaRetentionSetting(request *DescribeHanaRetentionSettingRequest) (_result *DescribeHanaRetentionSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHanaRetentionSettingResponse{}
	_body, _err := client.DescribeHanaRetentionSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about Tablestore instances that are backed up.
//
// @param request - DescribeOtsTableSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOtsTableSnapshotsResponse
func (client *Client) DescribeOtsTableSnapshotsWithOptions(request *DescribeOtsTableSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeOtsTableSnapshotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		body["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.OtsInstances) {
		bodyFlat["OtsInstances"] = request.OtsInstances
	}

	if !dara.IsNil(request.SnapshotIds) {
		bodyFlat["SnapshotIds"] = request.SnapshotIds
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOtsTableSnapshots"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOtsTableSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about Tablestore instances that are backed up.
//
// @param request - DescribeOtsTableSnapshotsRequest
//
// @return DescribeOtsTableSnapshotsResponse
func (client *Client) DescribeOtsTableSnapshots(request *DescribeOtsTableSnapshotsRequest) (_result *DescribeOtsTableSnapshotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOtsTableSnapshotsResponse{}
	_body, _err := client.DescribeOtsTableSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more backup policies.
//
// @param request - DescribePoliciesV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePoliciesV2Response
func (client *Client) DescribePoliciesV2WithOptions(request *DescribePoliciesV2Request, runtime *dara.RuntimeOptions) (_result *DescribePoliciesV2Response, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePoliciesV2"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePoliciesV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more backup policies.
//
// @param request - DescribePoliciesV2Request
//
// @return DescribePoliciesV2Response
func (client *Client) DescribePoliciesV2(request *DescribePoliciesV2Request) (_result *DescribePoliciesV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePoliciesV2Response{}
	_body, _err := client.DescribePoliciesV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query one or more data sources bound to a policy, or query one or more policies bound to a data source.
//
// @param tmpReq - DescribePolicyBindingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolicyBindingsResponse
func (client *Client) DescribePolicyBindingsWithOptions(tmpReq *DescribePolicyBindingsRequest, runtime *dara.RuntimeOptions) (_result *DescribePolicyBindingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePolicyBindingsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataSourceIds) {
		request.DataSourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSourceIds, dara.String("DataSourceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceIdsShrink) {
		body["DataSourceIds"] = request.DataSourceIdsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolicyBindings"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolicyBindingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query one or more data sources bound to a policy, or query one or more policies bound to a data source.
//
// @param request - DescribePolicyBindingsRequest
//
// @return DescribePolicyBindingsResponse
func (client *Client) DescribePolicyBindings(request *DescribePolicyBindingsRequest) (_result *DescribePolicyBindingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePolicyBindingsResponse{}
	_body, _err := client.DescribePolicyBindingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tables of a restorable Tablestore instance.
//
// @param request - DescribeRecoverableOtsInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecoverableOtsInstancesResponse
func (client *Client) DescribeRecoverableOtsInstancesWithOptions(request *DescribeRecoverableOtsInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRecoverableOtsInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRecoverableOtsInstances"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRecoverableOtsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tables of a restorable Tablestore instance.
//
// @param request - DescribeRecoverableOtsInstancesRequest
//
// @return DescribeRecoverableOtsInstancesResponse
func (client *Client) DescribeRecoverableOtsInstances(request *DescribeRecoverableOtsInstancesRequest) (_result *DescribeRecoverableOtsInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRecoverableOtsInstancesResponse{}
	_body, _err := client.DescribeRecoverableOtsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2017-09-08"),
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
// Queries available regions.
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more restore jobs that meet the specified conditions.
//
// @param request - DescribeRestoreJobs2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreJobs2Response
func (client *Client) DescribeRestoreJobs2WithOptions(request *DescribeRestoreJobs2Request, runtime *dara.RuntimeOptions) (_result *DescribeRestoreJobs2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreJobs2"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreJobs2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more restore jobs that meet the specified conditions.
//
// @param request - DescribeRestoreJobs2Request
//
// @return DescribeRestoreJobs2Response
func (client *Client) DescribeRestoreJobs2(request *DescribeRestoreJobs2Request) (_result *DescribeRestoreJobs2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRestoreJobs2Response{}
	_body, _err := client.DescribeRestoreJobs2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an asynchronous job.
//
// @param request - DescribeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskResponse
func (client *Client) DescribeTaskWithOptions(request *DescribeTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTask"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an asynchronous job.
//
// @param request - DescribeTaskRequest
//
// @return DescribeTaskResponse
func (client *Client) DescribeTask(request *DescribeTaskRequest) (_result *DescribeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTaskResponse{}
	_body, _err := client.DescribeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backup snapshots of an Elastic Compute Service (ECS) instance.
//
// @param tmpReq - DescribeUdmSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUdmSnapshotsResponse
func (client *Client) DescribeUdmSnapshotsWithOptions(tmpReq *DescribeUdmSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUdmSnapshotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeUdmSnapshotsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SnapshotIds) {
		request.SnapshotIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SnapshotIds, dara.String("SnapshotIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UdmRegionId) {
		query["UdmRegionId"] = request.UdmRegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SnapshotIdsShrink) {
		body["SnapshotIds"] = request.SnapshotIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUdmSnapshots"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUdmSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup snapshots of an Elastic Compute Service (ECS) instance.
//
// @param request - DescribeUdmSnapshotsRequest
//
// @return DescribeUdmSnapshotsResponse
func (client *Client) DescribeUdmSnapshots(request *DescribeUdmSnapshotsRequest) (_result *DescribeUdmSnapshotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUdmSnapshotsResponse{}
	_body, _err := client.DescribeUdmSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions that support cross-region replication.
//
// @param request - DescribeVaultReplicationRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVaultReplicationRegionsResponse
func (client *Client) DescribeVaultReplicationRegionsWithOptions(request *DescribeVaultReplicationRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVaultReplicationRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVaultReplicationRegions"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVaultReplicationRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions that support cross-region replication.
//
// @param request - DescribeVaultReplicationRegionsRequest
//
// @return DescribeVaultReplicationRegionsResponse
func (client *Client) DescribeVaultReplicationRegions(request *DescribeVaultReplicationRegionsRequest) (_result *DescribeVaultReplicationRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVaultReplicationRegionsResponse{}
	_body, _err := client.DescribeVaultReplicationRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more backup vaults that meet the specified conditions.
//
// @param request - DescribeVaultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVaultsResponse
func (client *Client) DescribeVaultsWithOptions(request *DescribeVaultsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVaultsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	if !dara.IsNil(request.VaultName) {
		query["VaultName"] = request.VaultName
	}

	if !dara.IsNil(request.VaultRegionId) {
		query["VaultRegionId"] = request.VaultRegionId
	}

	if !dara.IsNil(request.VaultType) {
		query["VaultType"] = request.VaultType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVaults"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVaultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more backup vaults that meet the specified conditions.
//
// @param request - DescribeVaultsRequest
//
// @return DescribeVaultsResponse
func (client *Client) DescribeVaults(request *DescribeVaultsRequest) (_result *DescribeVaultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVaultsResponse{}
	_body, _err := client.DescribeVaultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an internal mount target created by Cloud Backup.
//
// Description:
//
//	  If the request is successful, the mount target is deleted.
//
//		- After you create a backup plan for an Apsara File Storage NAS file system, HBR automatically creates a mount target for the file system. You can call this operation to delete the mount target. In the **Status*	- column of the mount target of the NAS file system, the following information is displayed: **This mount target is created by an Alibaba Cloud internal service and cannot be operated. Service name: HBR**.
//
// @param request - DetachNasFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachNasFileSystemResponse
func (client *Client) DetachNasFileSystemWithOptions(request *DetachNasFileSystemRequest, runtime *dara.RuntimeOptions) (_result *DetachNasFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTime) {
		query["CreateTime"] = request.CreateTime
	}

	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachNasFileSystem"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachNasFileSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an internal mount target created by Cloud Backup.
//
// Description:
//
//	  If the request is successful, the mount target is deleted.
//
//		- After you create a backup plan for an Apsara File Storage NAS file system, HBR automatically creates a mount target for the file system. You can call this operation to delete the mount target. In the **Status*	- column of the mount target of the NAS file system, the following information is displayed: **This mount target is created by an Alibaba Cloud internal service and cannot be operated. Service name: HBR**.
//
// @param request - DetachNasFileSystemRequest
//
// @return DetachNasFileSystemResponse
func (client *Client) DetachNasFileSystem(request *DetachNasFileSystemRequest) (_result *DetachNasFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachNasFileSystemResponse{}
	_body, _err := client.DetachNasFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a backup plan.
//
// Description:
//
// After you call this operation, the backup plan is suspended. In the DescribeBackupPlans operation, the Disabled parameter is set to true.
//
// @param request - DisableBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableBackupPlanResponse
func (client *Client) DisableBackupPlanWithOptions(request *DisableBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *DisableBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a backup plan.
//
// Description:
//
// After you call this operation, the backup plan is suspended. In the DescribeBackupPlans operation, the Disabled parameter is set to true.
//
// @param request - DisableBackupPlanRequest
//
// @return DisableBackupPlanResponse
func (client *Client) DisableBackupPlan(request *DisableBackupPlanRequest) (_result *DisableBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableBackupPlanResponse{}
	_body, _err := client.DisableBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an SAP HANA backup plan.
//
// Description:
//
// To enable the backup plan again, call the EnableHanaBackupPlan operation.
//
// @param request - DisableHanaBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableHanaBackupPlanResponse
func (client *Client) DisableHanaBackupPlanWithOptions(request *DisableHanaBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *DisableHanaBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableHanaBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an SAP HANA backup plan.
//
// Description:
//
// To enable the backup plan again, call the EnableHanaBackupPlan operation.
//
// @param request - DisableHanaBackupPlanRequest
//
// @return DisableHanaBackupPlanResponse
func (client *Client) DisableHanaBackupPlan(request *DisableHanaBackupPlanRequest) (_result *DisableHanaBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableHanaBackupPlanResponse{}
	_body, _err := client.DisableHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a backup plan.
//
// Description:
//
// After you call this operation, the backup plan is restarted (Disabled is set to false in the DescribeBackupPlans operation). Cloud Backup continues to perform backups based on the policy specified in the backup plan.
//
// @param request - EnableBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableBackupPlanResponse
func (client *Client) EnableBackupPlanWithOptions(request *EnableBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *EnableBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a backup plan.
//
// Description:
//
// After you call this operation, the backup plan is restarted (Disabled is set to false in the DescribeBackupPlans operation). Cloud Backup continues to perform backups based on the policy specified in the backup plan.
//
// @param request - EnableBackupPlanRequest
//
// @return EnableBackupPlanResponse
func (client *Client) EnableBackupPlan(request *EnableBackupPlanRequest) (_result *EnableBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableBackupPlanResponse{}
	_body, _err := client.EnableBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an SAP HANA backup plan.
//
// Description:
//
// To disable the backup plan again, call the DisableHanaBackupPlan operation.
//
// @param request - EnableHanaBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHanaBackupPlanResponse
func (client *Client) EnableHanaBackupPlanWithOptions(request *EnableHanaBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *EnableHanaBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableHanaBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an SAP HANA backup plan.
//
// Description:
//
// To disable the backup plan again, call the DisableHanaBackupPlan operation.
//
// @param request - EnableHanaBackupPlanRequest
//
// @return EnableHanaBackupPlanResponse
func (client *Client) EnableHanaBackupPlan(request *EnableHanaBackupPlanRequest) (_result *EnableHanaBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableHanaBackupPlanResponse{}
	_body, _err := client.EnableHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a backup plan.
//
// @param request - ExecuteBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteBackupPlanResponse
func (client *Client) ExecuteBackupPlanWithOptions(request *ExecuteBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *ExecuteBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a backup plan.
//
// @param request - ExecuteBackupPlanRequest
//
// @return ExecuteBackupPlanResponse
func (client *Client) ExecuteBackupPlan(request *ExecuteBackupPlanRequest) (_result *ExecuteBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteBackupPlanResponse{}
	_body, _err := client.ExecuteBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Execute a policy for one or all bound data sources.
//
// @param request - ExecutePolicyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecutePolicyV2Response
func (client *Client) ExecutePolicyV2WithOptions(request *ExecutePolicyV2Request, runtime *dara.RuntimeOptions) (_result *ExecutePolicyV2Response, _err error) {
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

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecutePolicyV2"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecutePolicyV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Execute a policy for one or all bound data sources.
//
// @param request - ExecutePolicyV2Request
//
// @return ExecutePolicyV2Response
func (client *Client) ExecutePolicyV2(request *ExecutePolicyV2Request) (_result *ExecutePolicyV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecutePolicyV2Response{}
	_body, _err := client.ExecutePolicyV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a Resource Access Management (RAM) policy.
//
// @param request - GenerateRamPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateRamPolicyResponse
func (client *Client) GenerateRamPolicyWithOptions(request *GenerateRamPolicyRequest, runtime *dara.RuntimeOptions) (_result *GenerateRamPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.RequireBasePolicy) {
		query["RequireBasePolicy"] = request.RequireBasePolicy
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateRamPolicy"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateRamPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a Resource Access Management (RAM) policy.
//
// @param request - GenerateRamPolicyRequest
//
// @return GenerateRamPolicyResponse
func (client *Client) GenerateRamPolicy(request *GenerateRamPolicyRequest) (_result *GenerateRamPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateRamPolicyResponse{}
	_body, _err := client.GenerateRamPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains download links of files such as job reports.
//
// @param request - GetTempFileDownloadLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTempFileDownloadLinkResponse
func (client *Client) GetTempFileDownloadLinkWithOptions(request *GetTempFileDownloadLinkRequest, runtime *dara.RuntimeOptions) (_result *GetTempFileDownloadLinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TempFileKey) {
		query["TempFileKey"] = request.TempFileKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTempFileDownloadLink"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTempFileDownloadLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains download links of files such as job reports.
//
// @param request - GetTempFileDownloadLinkRequest
//
// @return GetTempFileDownloadLinkResponse
func (client *Client) GetTempFileDownloadLink(request *GetTempFileDownloadLinkRequest) (_result *GetTempFileDownloadLinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTempFileDownloadLinkResponse{}
	_body, _err := client.GetTempFileDownloadLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an HBR client on one or more Elastic Compute Service (ECS) instances.
//
// Description:
//
//	  This operation creates an asynchronous job at the backend and calls Cloud Assistant to install an HBR client on an ECS instance.
//
//		- You can call the [DescribeTask](https://help.aliyun.com/document_detail/431265.html) operation to query the execution result of an asynchronous job.
//
//		- The timeout period of an asynchronous job is 15 minutes. We recommend that you call the DescribeTask operation to run the first query 60 seconds after you call the InstallBackupClients operation to install HBR clients. Then, run the next queries at an interval of 30 seconds.
//
// @param tmpReq - InstallBackupClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallBackupClientsResponse
func (client *Client) InstallBackupClientsWithOptions(tmpReq *InstallBackupClientsRequest, runtime *dara.RuntimeOptions) (_result *InstallBackupClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InstallBackupClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallBackupClients"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallBackupClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an HBR client on one or more Elastic Compute Service (ECS) instances.
//
// Description:
//
//	  This operation creates an asynchronous job at the backend and calls Cloud Assistant to install an HBR client on an ECS instance.
//
//		- You can call the [DescribeTask](https://help.aliyun.com/document_detail/431265.html) operation to query the execution result of an asynchronous job.
//
//		- The timeout period of an asynchronous job is 15 minutes. We recommend that you call the DescribeTask operation to run the first query 60 seconds after you call the InstallBackupClients operation to install HBR clients. Then, run the next queries at an interval of 30 seconds.
//
// @param request - InstallBackupClientsRequest
//
// @return InstallBackupClientsResponse
func (client *Client) InstallBackupClients(request *InstallBackupClientsRequest) (_result *InstallBackupClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InstallBackupClientsResponse{}
	_body, _err := client.InstallBackupClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询已保护的资源列表
//
// @param request - ListProtectedResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProtectedResourcesResponse
func (client *Client) ListProtectedResourcesWithOptions(request *ListProtectedResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListProtectedResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreatedByProduct) {
		query["CreatedByProduct"] = request.CreatedByProduct
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Skip) {
		query["Skip"] = request.Skip
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProtectedResources"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProtectedResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询已保护的资源列表
//
// @param request - ListProtectedResourcesRequest
//
// @return ListProtectedResourcesResponse
func (client *Client) ListProtectedResources(request *ListProtectedResourcesRequest) (_result *ListProtectedResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProtectedResourcesResponse{}
	_body, _err := client.ListProtectedResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates Cloud Backup.
//
// @param request - OpenHbrServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenHbrServiceResponse
func (client *Client) OpenHbrServiceWithOptions(runtime *dara.RuntimeOptions) (_result *OpenHbrServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("OpenHbrService"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenHbrServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates Cloud Backup.
//
// @return OpenHbrServiceResponse
func (client *Client) OpenHbrService() (_result *OpenHbrServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenHbrServiceResponse{}
	_body, _err := client.OpenHbrServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more backup snapshots that meet the specified conditions.
//
// @param tmpReq - SearchHistoricalSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchHistoricalSnapshotsResponse
func (client *Client) SearchHistoricalSnapshotsWithOptions(tmpReq *SearchHistoricalSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *SearchHistoricalSnapshotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchHistoricalSnapshotsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("Query"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.QueryShrink) {
		query["Query"] = request.QueryShrink
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchHistoricalSnapshots"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchHistoricalSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more backup snapshots that meet the specified conditions.
//
// @param request - SearchHistoricalSnapshotsRequest
//
// @return SearchHistoricalSnapshotsResponse
func (client *Client) SearchHistoricalSnapshots(request *SearchHistoricalSnapshotsRequest) (_result *SearchHistoricalSnapshotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchHistoricalSnapshotsResponse{}
	_body, _err := client.SearchHistoricalSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an SAP HANA database.
//
// Description:
//
// To stop the database again, call the StopHanaDatabaseAsync operation.
//
// @param request - StartHanaDatabaseAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartHanaDatabaseAsyncResponse
func (client *Client) StartHanaDatabaseAsyncWithOptions(request *StartHanaDatabaseAsyncRequest, runtime *dara.RuntimeOptions) (_result *StartHanaDatabaseAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartHanaDatabaseAsync"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartHanaDatabaseAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an SAP HANA database.
//
// Description:
//
// To stop the database again, call the StopHanaDatabaseAsync operation.
//
// @param request - StartHanaDatabaseAsyncRequest
//
// @return StartHanaDatabaseAsyncResponse
func (client *Client) StartHanaDatabaseAsync(request *StartHanaDatabaseAsyncRequest) (_result *StartHanaDatabaseAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartHanaDatabaseAsyncResponse{}
	_body, _err := client.StartHanaDatabaseAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an SAP HANA database.
//
// Description:
//
// To start the database again, call the StartHanaDatabaseAsync operation.
//
// @param request - StopHanaDatabaseAsyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopHanaDatabaseAsyncResponse
func (client *Client) StopHanaDatabaseAsyncWithOptions(request *StopHanaDatabaseAsyncRequest, runtime *dara.RuntimeOptions) (_result *StopHanaDatabaseAsyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopHanaDatabaseAsync"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopHanaDatabaseAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an SAP HANA database.
//
// Description:
//
// To start the database again, call the StartHanaDatabaseAsync operation.
//
// @param request - StopHanaDatabaseAsyncRequest
//
// @return StopHanaDatabaseAsyncResponse
func (client *Client) StopHanaDatabaseAsync(request *StopHanaDatabaseAsyncRequest) (_result *StopHanaDatabaseAsyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopHanaDatabaseAsyncResponse{}
	_body, _err := client.StopHanaDatabaseAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls a Cloud Backup client from one or more Elastic Compute Service (ECS) instance.
//
// Description:
//
//	  This operation creates an asynchronous job at the backend and calls Cloud Assistant to uninstall a backup client from an ECS instance.
//
//		- You can call the DescribeTask operation to query the execution result of an asynchronous job.
//
//		- The timeout period of an asynchronous job is 15 minutes. We recommend that you call the DescribeTask operation to run the first query 30 seconds after you call the UninstallBackupClients operation to uninstall backup clients. Then, run the next queries at an interval of 30 seconds.
//
// @param tmpReq - UninstallBackupClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallBackupClientsResponse
func (client *Client) UninstallBackupClientsWithOptions(tmpReq *UninstallBackupClientsRequest, runtime *dara.RuntimeOptions) (_result *UninstallBackupClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UninstallBackupClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClientIds) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, dara.String("ClientIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIdsShrink) {
		query["ClientIds"] = request.ClientIdsShrink
	}

	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallBackupClients"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallBackupClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls a Cloud Backup client from one or more Elastic Compute Service (ECS) instance.
//
// Description:
//
//	  This operation creates an asynchronous job at the backend and calls Cloud Assistant to uninstall a backup client from an ECS instance.
//
//		- You can call the DescribeTask operation to query the execution result of an asynchronous job.
//
//		- The timeout period of an asynchronous job is 15 minutes. We recommend that you call the DescribeTask operation to run the first query 30 seconds after you call the UninstallBackupClients operation to uninstall backup clients. Then, run the next queries at an interval of 30 seconds.
//
// @param request - UninstallBackupClientsRequest
//
// @return UninstallBackupClientsResponse
func (client *Client) UninstallBackupClients(request *UninstallBackupClientsRequest) (_result *UninstallBackupClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UninstallBackupClientsResponse{}
	_body, _err := client.UninstallBackupClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls an HBR client.
//
// Description:
//
// If you call this operation, the specified HBR client is uninstalled. To reinstall the HBR client, call the CreateClients operation.
//
// @param request - UninstallClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallClientResponse
func (client *Client) UninstallClientWithOptions(request *UninstallClientRequest, runtime *dara.RuntimeOptions) (_result *UninstallClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallClient"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls an HBR client.
//
// Description:
//
// If you call this operation, the specified HBR client is uninstalled. To reinstall the HBR client, call the CreateClients operation.
//
// @param request - UninstallClientRequest
//
// @return UninstallClientResponse
func (client *Client) UninstallClient(request *UninstallClientRequest) (_result *UninstallClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UninstallClientResponse{}
	_body, _err := client.UninstallClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a backup plan.
//
// @param tmpReq - UpdateBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBackupPlanResponse
func (client *Client) UpdateBackupPlanWithOptions(tmpReq *UpdateBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBackupPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Detail) {
		request.DetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Detail, dara.String("Detail"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OtsDetail) {
		request.OtsDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OtsDetail, dara.String("OtsDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeListPath) {
		query["ChangeListPath"] = request.ChangeListPath
	}

	if !dara.IsNil(request.DetailShrink) {
		query["Detail"] = request.DetailShrink
	}

	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.KeepLatestSnapshots) {
		query["KeepLatestSnapshots"] = request.KeepLatestSnapshots
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.PlanName) {
		query["PlanName"] = request.PlanName
	}

	if !dara.IsNil(request.Prefix) {
		query["Prefix"] = request.Prefix
	}

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	if !dara.IsNil(request.Schedule) {
		query["Schedule"] = request.Schedule
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SpeedLimit) {
		query["SpeedLimit"] = request.SpeedLimit
	}

	if !dara.IsNil(request.UpdatePaths) {
		query["UpdatePaths"] = request.UpdatePaths
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Exclude) {
		body["Exclude"] = request.Exclude
	}

	if !dara.IsNil(request.Include) {
		body["Include"] = request.Include
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.OtsDetailShrink) {
		body["OtsDetail"] = request.OtsDetailShrink
	}

	if !dara.IsNil(request.Rule) {
		body["Rule"] = request.Rule
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a backup plan.
//
// @param request - UpdateBackupPlanRequest
//
// @return UpdateBackupPlanResponse
func (client *Client) UpdateBackupPlan(request *UpdateBackupPlanRequest) (_result *UpdateBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBackupPlanResponse{}
	_body, _err := client.UpdateBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of an HBR client.
//
// Description:
//
// You can call this operation to update the configurations of both the old and new HBR clients.
//
// @param request - UpdateClientSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientSettingsResponse
func (client *Client) UpdateClientSettingsWithOptions(request *UpdateClientSettingsRequest, runtime *dara.RuntimeOptions) (_result *UpdateClientSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertOnPartialComplete) {
		query["AlertOnPartialComplete"] = request.AlertOnPartialComplete
	}

	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.DataNetworkType) {
		query["DataNetworkType"] = request.DataNetworkType
	}

	if !dara.IsNil(request.DataProxySetting) {
		query["DataProxySetting"] = request.DataProxySetting
	}

	if !dara.IsNil(request.MaxCpuCore) {
		query["MaxCpuCore"] = request.MaxCpuCore
	}

	if !dara.IsNil(request.MaxMemory) {
		query["MaxMemory"] = request.MaxMemory
	}

	if !dara.IsNil(request.MaxWorker) {
		query["MaxWorker"] = request.MaxWorker
	}

	if !dara.IsNil(request.ProxyHost) {
		query["ProxyHost"] = request.ProxyHost
	}

	if !dara.IsNil(request.ProxyPassword) {
		query["ProxyPassword"] = request.ProxyPassword
	}

	if !dara.IsNil(request.ProxyPort) {
		query["ProxyPort"] = request.ProxyPort
	}

	if !dara.IsNil(request.ProxyUser) {
		query["ProxyUser"] = request.ProxyUser
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.UseHttps) {
		query["UseHttps"] = request.UseHttps
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClientSettings"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClientSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of an HBR client.
//
// Description:
//
// You can call this operation to update the configurations of both the old and new HBR clients.
//
// @param request - UpdateClientSettingsRequest
//
// @return UpdateClientSettingsResponse
func (client *Client) UpdateClientSettings(request *UpdateClientSettingsRequest) (_result *UpdateClientSettingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClientSettingsResponse{}
	_body, _err := client.UpdateClientSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update container cluster information, including the container cluster name, network type, etc.
//
// @param request - UpdateContainerClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContainerClusterResponse
func (client *Client) UpdateContainerClusterWithOptions(request *UpdateContainerClusterRequest, runtime *dara.RuntimeOptions) (_result *UpdateContainerClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.RenewToken) {
		query["RenewToken"] = request.RenewToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContainerCluster"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContainerClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update container cluster information, including the container cluster name, network type, etc.
//
// @param request - UpdateContainerClusterRequest
//
// @return UpdateContainerClusterResponse
func (client *Client) UpdateContainerCluster(request *UpdateContainerClusterRequest) (_result *UpdateContainerClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateContainerClusterResponse{}
	_body, _err := client.UpdateContainerClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an SAP HANA backup plan.
//
// Description:
//
//	  A backup plan defines the data source, backup policy, and other configurations. After you execute a backup plan, a backup job is generated to record the backup progress and the backup result. If a backup job is completed, a backup snapshot is generated. You can use a backup snapshot to create a restore job.
//
//		- You can specify only one type of data source in a backup plan.
//
//		- You can specify only one interval as a backup cycle in a backup plan.
//
//		- Each backup plan allows you to back up data to only one backup vault.
//
// @param request - UpdateHanaBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHanaBackupPlanResponse
func (client *Client) UpdateHanaBackupPlanWithOptions(request *UpdateHanaBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *UpdateHanaBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPrefix) {
		query["BackupPrefix"] = request.BackupPrefix
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.PlanId) {
		query["PlanId"] = request.PlanId
	}

	if !dara.IsNil(request.PlanName) {
		query["PlanName"] = request.PlanName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Schedule) {
		query["Schedule"] = request.Schedule
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHanaBackupPlan"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHanaBackupPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an SAP HANA backup plan.
//
// Description:
//
//	  A backup plan defines the data source, backup policy, and other configurations. After you execute a backup plan, a backup job is generated to record the backup progress and the backup result. If a backup job is completed, a backup snapshot is generated. You can use a backup snapshot to create a restore job.
//
//		- You can specify only one type of data source in a backup plan.
//
//		- You can specify only one interval as a backup cycle in a backup plan.
//
//		- Each backup plan allows you to back up data to only one backup vault.
//
// @param request - UpdateHanaBackupPlanRequest
//
// @return UpdateHanaBackupPlanResponse
func (client *Client) UpdateHanaBackupPlan(request *UpdateHanaBackupPlanRequest) (_result *UpdateHanaBackupPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHanaBackupPlanResponse{}
	_body, _err := client.UpdateHanaBackupPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the backup parameters of an SAP HANA database.
//
// Description:
//
// You can call the UpdateHanaRetentionSetting operation to update the backup retention period of a database.
//
// @param request - UpdateHanaBackupSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHanaBackupSettingResponse
func (client *Client) UpdateHanaBackupSettingWithOptions(request *UpdateHanaBackupSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateHanaBackupSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogBackupParameterFile) {
		query["CatalogBackupParameterFile"] = request.CatalogBackupParameterFile
	}

	if !dara.IsNil(request.CatalogBackupUsingBackint) {
		query["CatalogBackupUsingBackint"] = request.CatalogBackupUsingBackint
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataBackupParameterFile) {
		query["DataBackupParameterFile"] = request.DataBackupParameterFile
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.EnableAutoLogBackup) {
		query["EnableAutoLogBackup"] = request.EnableAutoLogBackup
	}

	if !dara.IsNil(request.LogBackupParameterFile) {
		query["LogBackupParameterFile"] = request.LogBackupParameterFile
	}

	if !dara.IsNil(request.LogBackupTimeout) {
		query["LogBackupTimeout"] = request.LogBackupTimeout
	}

	if !dara.IsNil(request.LogBackupUsingBackint) {
		query["LogBackupUsingBackint"] = request.LogBackupUsingBackint
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHanaBackupSetting"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHanaBackupSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the backup parameters of an SAP HANA database.
//
// Description:
//
// You can call the UpdateHanaRetentionSetting operation to update the backup retention period of a database.
//
// @param request - UpdateHanaBackupSettingRequest
//
// @return UpdateHanaBackupSettingResponse
func (client *Client) UpdateHanaBackupSetting(request *UpdateHanaBackupSettingRequest) (_result *UpdateHanaBackupSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHanaBackupSettingResponse{}
	_body, _err := client.UpdateHanaBackupSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an SAP HANA instance.
//
// @param request - UpdateHanaInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHanaInstanceResponse
func (client *Client) UpdateHanaInstanceWithOptions(request *UpdateHanaInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateHanaInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlertSetting) {
		query["AlertSetting"] = request.AlertSetting
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.HanaName) {
		query["HanaName"] = request.HanaName
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceNumber) {
		query["InstanceNumber"] = request.InstanceNumber
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.UseSsl) {
		query["UseSsl"] = request.UseSsl
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.ValidateCertificate) {
		query["ValidateCertificate"] = request.ValidateCertificate
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHanaInstance"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHanaInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an SAP HANA instance.
//
// @param request - UpdateHanaInstanceRequest
//
// @return UpdateHanaInstanceResponse
func (client *Client) UpdateHanaInstance(request *UpdateHanaInstanceRequest) (_result *UpdateHanaInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHanaInstanceResponse{}
	_body, _err := client.UpdateHanaInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the backup retention period of an SAP HANA database.
//
// Description:
//
//	  If you want to update the backup parameters of a database, you can call the UpdateHanaBackupSetting operation.
//
//		- Cloud Backup deletes the expired catalogs and data that are related to Backint and file backup. The deleted catalogs and data cannot be restored. We recommend that you set the retention period based on your business requirements.
//
// @param request - UpdateHanaRetentionSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHanaRetentionSettingResponse
func (client *Client) UpdateHanaRetentionSettingWithOptions(request *UpdateHanaRetentionSettingRequest, runtime *dara.RuntimeOptions) (_result *UpdateHanaRetentionSettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Disabled) {
		query["Disabled"] = request.Disabled
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !dara.IsNil(request.Schedule) {
		query["Schedule"] = request.Schedule
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHanaRetentionSetting"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHanaRetentionSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the backup retention period of an SAP HANA database.
//
// Description:
//
//	  If you want to update the backup parameters of a database, you can call the UpdateHanaBackupSetting operation.
//
//		- Cloud Backup deletes the expired catalogs and data that are related to Backint and file backup. The deleted catalogs and data cannot be restored. We recommend that you set the retention period based on your business requirements.
//
// @param request - UpdateHanaRetentionSettingRequest
//
// @return UpdateHanaRetentionSettingResponse
func (client *Client) UpdateHanaRetentionSetting(request *UpdateHanaRetentionSettingRequest) (_result *UpdateHanaRetentionSettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHanaRetentionSettingResponse{}
	_body, _err := client.UpdateHanaRetentionSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the association between a backup policy and a data source.
//
// @param tmpReq - UpdatePolicyBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyBindingResponse
func (client *Client) UpdatePolicyBindingWithOptions(tmpReq *UpdatePolicyBindingRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolicyBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePolicyBindingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdvancedOptions) {
		request.AdvancedOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdvancedOptions, dara.String("AdvancedOptions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedOptionsShrink) {
		query["AdvancedOptions"] = request.AdvancedOptionsShrink
	}

	if !dara.IsNil(request.Disabled) {
		query["Disabled"] = request.Disabled
	}

	if !dara.IsNil(request.Exclude) {
		query["Exclude"] = request.Exclude
	}

	if !dara.IsNil(request.Include) {
		query["Include"] = request.Include
	}

	if !dara.IsNil(request.PolicyBindingDescription) {
		query["PolicyBindingDescription"] = request.PolicyBindingDescription
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SpeedLimit) {
		query["SpeedLimit"] = request.SpeedLimit
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolicyBinding"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolicyBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the association between a backup policy and a data source.
//
// @param request - UpdatePolicyBindingRequest
//
// @return UpdatePolicyBindingResponse
func (client *Client) UpdatePolicyBinding(request *UpdatePolicyBindingRequest) (_result *UpdatePolicyBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePolicyBindingResponse{}
	_body, _err := client.UpdatePolicyBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a backup policy.
//
// Description:
//
// If you modify a backup policy, the modification takes effect on all data sources that are bound to the backup policy. Proceed with caution.
//
// @param tmpReq - UpdatePolicyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyV2Response
func (client *Client) UpdatePolicyV2WithOptions(tmpReq *UpdatePolicyV2Request, runtime *dara.RuntimeOptions) (_result *UpdatePolicyV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePolicyV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Rules) {
		request.RulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rules, dara.String("Rules"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyDescription) {
		body["PolicyDescription"] = request.PolicyDescription
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyName) {
		body["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.RulesShrink) {
		body["Rules"] = request.RulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolicyV2"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolicyV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a backup policy.
//
// Description:
//
// If you modify a backup policy, the modification takes effect on all data sources that are bound to the backup policy. Proceed with caution.
//
// @param request - UpdatePolicyV2Request
//
// @return UpdatePolicyV2Response
func (client *Client) UpdatePolicyV2(request *UpdatePolicyV2Request) (_result *UpdatePolicyV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePolicyV2Response{}
	_body, _err := client.UpdatePolicyV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration information about the backup vault.
//
// @param request - UpdateVaultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVaultResponse
func (client *Client) UpdateVaultWithOptions(request *UpdateVaultRequest, runtime *dara.RuntimeOptions) (_result *UpdateVaultResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	if !dara.IsNil(request.VaultName) {
		query["VaultName"] = request.VaultName
	}

	if !dara.IsNil(request.WormEnabled) {
		query["WormEnabled"] = request.WormEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVault"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVaultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration information about the backup vault.
//
// @param request - UpdateVaultRequest
//
// @return UpdateVaultResponse
func (client *Client) UpdateVault(request *UpdateVaultRequest) (_result *UpdateVaultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVaultResponse{}
	_body, _err := client.UpdateVaultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades an HBR client on one or more Elastic Compute Service (ECS) instances.
//
// Description:
//
//	  This operation creates an asynchronous job at the backend and calls Cloud Assistant to upgrade an HBR client that is installed on an ECS instance.
//
//		- You can call the DescribeTask operation to query the execution result of an asynchronous job.
//
//		- The timeout period of an asynchronous job is 15 minutes.
//
// @param tmpReq - UpgradeBackupClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeBackupClientsResponse
func (client *Client) UpgradeBackupClientsWithOptions(tmpReq *UpgradeBackupClientsRequest, runtime *dara.RuntimeOptions) (_result *UpgradeBackupClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpgradeBackupClientsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ClientIds) {
		request.ClientIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClientIds, dara.String("ClientIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIdsShrink) {
		query["ClientIds"] = request.ClientIdsShrink
	}

	if !dara.IsNil(request.CrossAccountRoleName) {
		query["CrossAccountRoleName"] = request.CrossAccountRoleName
	}

	if !dara.IsNil(request.CrossAccountType) {
		query["CrossAccountType"] = request.CrossAccountType
	}

	if !dara.IsNil(request.CrossAccountUserId) {
		query["CrossAccountUserId"] = request.CrossAccountUserId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeBackupClients"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeBackupClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades an HBR client on one or more Elastic Compute Service (ECS) instances.
//
// Description:
//
//	  This operation creates an asynchronous job at the backend and calls Cloud Assistant to upgrade an HBR client that is installed on an ECS instance.
//
//		- You can call the DescribeTask operation to query the execution result of an asynchronous job.
//
//		- The timeout period of an asynchronous job is 15 minutes.
//
// @param request - UpgradeBackupClientsRequest
//
// @return UpgradeBackupClientsResponse
func (client *Client) UpgradeBackupClients(request *UpgradeBackupClientsRequest) (_result *UpgradeBackupClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeBackupClientsResponse{}
	_body, _err := client.UpgradeBackupClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the Cloud Backup client.
//
// Description:
//
// You can call this operation to upgrade a Cloud Backup client to the latest version. After the Cloud Backup client is upgraded, the version of the client cannot be rolled back.
//
// @param request - UpgradeClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeClientResponse
func (client *Client) UpgradeClientWithOptions(request *UpgradeClientRequest, runtime *dara.RuntimeOptions) (_result *UpgradeClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VaultId) {
		query["VaultId"] = request.VaultId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeClient"),
		Version:     dara.String("2017-09-08"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the Cloud Backup client.
//
// Description:
//
// You can call this operation to upgrade a Cloud Backup client to the latest version. After the Cloud Backup client is upgraded, the version of the client cannot be rolled back.
//
// @param request - UpgradeClientRequest
//
// @return UpgradeClientResponse
func (client *Client) UpgradeClient(request *UpgradeClientRequest) (_result *UpgradeClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeClientResponse{}
	_body, _err := client.UpgradeClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
