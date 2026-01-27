// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Moves a resource from one resource group to another.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
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

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置备份集信息
//
// @param request - ConfigureBackupSetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureBackupSetInfoResponse
func (client *Client) ConfigureBackupSetInfoWithContext(ctx context.Context, request *ConfigureBackupSetInfoRequest, runtime *dara.RuntimeOptions) (_result *ConfigureBackupSetInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupMethod) {
		query["BackupMethod"] = request.BackupMethod
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.BackupSetName) {
		query["BackupSetName"] = request.BackupSetName
	}

	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.CheckSum) {
		query["CheckSum"] = request.CheckSum
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.ExtraMeta) {
		query["ExtraMeta"] = request.ExtraMeta
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UploadStatus) {
		query["UploadStatus"] = request.UploadStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureBackupSetInfo"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureBackupSetInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an advanced backup policy for a PolarDB instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # PolarDB for MySQL
//
// >  This API operation is available only to specific customers. If you want to call this API operation, join the Database Backup (DBS) DingTalk group (ID 35585947) for customer consultation to request permissions.
//
// @param request - CreateAdvancedPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdvancedPolicyResponse
func (client *Client) CreateAdvancedPolicyWithContext(ctx context.Context, request *CreateAdvancedPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateAdvancedPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAdvancedPolicy"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAdvancedPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an advanced download task for an ApsaraDB RDS for MySQL instance, an ApsaraDB RDS for PostgreSQL instance, or a PolarDB for MySQL cluster.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - ApsaraDB RDS for MySQL
//
//   - ApsaraDB RDS for PostgreSQL
//
//   - PolarDB for MySQL
//
// ### [](#)References
//
// For the instances that meet your business requirements, you can create an advanced download task by point in time or backup set. You can set the download destination to a URL or directly upload the downloaded backup set to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving.
//
//   - [Download the backup files of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/98819.html)
//
//   - [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96774.html)
//
//   - [Download the backup files of a PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/2627635.html)
//
// @param request - CreateDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDownloadResponse
func (client *Client) CreateDownloadWithContext(ctx context.Context, request *CreateDownloadRequest, runtime *dara.RuntimeOptions) (_result *CreateDownloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdminDatabase) {
		query["AdminDatabase"] = request.AdminDatabase
	}

	if !dara.IsNil(request.BakSetId) {
		query["BakSetId"] = request.BakSetId
	}

	if !dara.IsNil(request.BakSetSize) {
		query["BakSetSize"] = request.BakSetSize
	}

	if !dara.IsNil(request.BakSetType) {
		query["BakSetType"] = request.BakSetType
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.DownloadPointInTime) {
		query["DownloadPointInTime"] = request.DownloadPointInTime
	}

	if !dara.IsNil(request.FormatType) {
		query["FormatType"] = request.FormatType
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IsCluster) {
		query["IsCluster"] = request.IsCluster
	}

	if !dara.IsNil(request.IsPhysical) {
		query["IsPhysical"] = request.IsPhysical
	}

	if !dara.IsNil(request.PrimaryKeyTypeOnly) {
		query["PrimaryKeyTypeOnly"] = request.PrimaryKeyTypeOnly
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.TargetBucket) {
		query["TargetBucket"] = request.TargetBucket
	}

	if !dara.IsNil(request.TargetOssRegion) {
		query["TargetOssRegion"] = request.TargetOssRegion
	}

	if !dara.IsNil(request.TargetPath) {
		query["TargetPath"] = request.TargetPath
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.UseZstd) {
		query["UseZstd"] = request.UseZstd
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDownload"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDownloadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a sandbox instance.
//
// Description:
//
// This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DeleteSandboxInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSandboxInstanceResponse
func (client *Client) DeleteSandboxInstanceWithContext(ctx context.Context, request *DeleteSandboxInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteSandboxInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSandboxInstance"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSandboxInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the backup data of a PolarDB for MySQL cluster.
//
// Description:
//
// ### [](#)Supported database engine
//
// # PolarDB for MySQL
//
// >  This API operation is available only to specific customers. If you want to call this API operation, join the Database Backup (DBS) DingTalk group (ID: 35585947) for customer consultation to request permissions.
//
// ### [](#)References
//
// [Back up data of PolarDB for MySQL](https://help.aliyun.com/document_detail/88172.html)
//
// @param request - DescribeBackupDataListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupDataListResponse
func (client *Client) DescribeBackupDataListWithContext(ctx context.Context, request *DescribeBackupDataListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupDataListResponse, _err error) {
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

	if !dara.IsNil(request.BackupMethod) {
		query["BackupMethod"] = request.BackupMethod
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.BackupScale) {
		query["BackupScale"] = request.BackupScale
	}

	if !dara.IsNil(request.BackupStatus) {
		query["BackupStatus"] = request.BackupStatus
	}

	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceIsDeleted) {
		query["InstanceIsDeleted"] = request.InstanceIsDeleted
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceRegion) {
		query["InstanceRegion"] = request.InstanceRegion
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupDataList"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupDataListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the backup policy of a PolarDB for MySQL instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # PolarDB for MySQL
//
// >  The API operation is available only to specific customers. If you want to call this API operation, request permissions by joining the Database Backup (DBS) DingTalk group (ID 35585947) for customer consultation.
//
// ### [](#)References
//
// [Topics related to backup policies of PolarDB for MySQL instances](https://help.aliyun.com/document_detail/280422.html)
//
// @param request - DescribeBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPolicyResponse
func (client *Client) DescribeBackupPolicyWithContext(ctx context.Context, request *DescribeBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询存储介质访问详情
//
// @param request - DescribeBakDataSourceStorageAccessInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBakDataSourceStorageAccessInfoResponse
func (client *Client) DescribeBakDataSourceStorageAccessInfoWithContext(ctx context.Context, request *DescribeBakDataSourceStorageAccessInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeBakDataSourceStorageAccessInfoResponse, _err error) {
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

	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.DurationSeconds) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBakDataSourceStorageAccessInfo"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBakDataSourceStorageAccessInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据dbs实例id获取费用列表
//
// @param request - DescribeCostInfoByDbsInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCostInfoByDbsInstanceResponse
func (client *Client) DescribeCostInfoByDbsInstanceWithContext(ctx context.Context, request *DescribeCostInfoByDbsInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeCostInfoByDbsInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCostInfoByDbsInstance"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCostInfoByDbsInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查库表恢复可用的备份集
//
// @param request - DescribeDBTablesRecoveryBackupSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBTablesRecoveryBackupSetResponse
func (client *Client) DescribeDBTablesRecoveryBackupSetWithContext(ctx context.Context, request *DescribeDBTablesRecoveryBackupSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBTablesRecoveryBackupSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBTablesRecoveryBackupSet"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBTablesRecoveryBackupSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查库表恢复状态
//
// @param request - DescribeDBTablesRecoveryStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBTablesRecoveryStateResponse
func (client *Client) DescribeDBTablesRecoveryStateWithContext(ctx context.Context, request *DescribeDBTablesRecoveryStateRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBTablesRecoveryStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBTablesRecoveryState"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBTablesRecoveryStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查库表恢复可恢复的时间范围
//
// @param request - DescribeDBTablesRecoveryTimeRangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBTablesRecoveryTimeRangeResponse
func (client *Client) DescribeDBTablesRecoveryTimeRangeWithContext(ctx context.Context, request *DescribeDBTablesRecoveryTimeRangeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBTablesRecoveryTimeRangeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBTablesRecoveryTimeRange"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBTablesRecoveryTimeRangeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage information of a downloaded backup set.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - ApsaraDB RDS for MySQL that uses Enterprise SSDs (ESSDs)
//
//   - RDS PostgreSQL
//
//   - PolarDB for MySQL
//
// ### [](#)References
//
//   - [Download the backup files of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/98819.html)
//
//   - [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96774.html)
//
//   - [Download the backup files of a PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/2627635.html)
//
// @param request - DescribeDownloadBackupSetStorageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadBackupSetStorageInfoResponse
func (client *Client) DescribeDownloadBackupSetStorageInfoWithContext(ctx context.Context, request *DescribeDownloadBackupSetStorageInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadBackupSetStorageInfoResponse, _err error) {
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

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDownloadBackupSetStorageInfo"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDownloadBackupSetStorageInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether an instance supports the advanced download feature.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - ApsaraDB RDS for MySQL
//
//   - ApsaraDB RDS for PostgreSQL
//
//   - PolarDB for MySQL
//
// ### [](#)References
//
// You can create an advanced download task by point in time or backup set. You can set the download destination to a URL or directly upload the downloaded backup set to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving.
//
//   - [Download the backup files of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/98819.html)
//
//   - [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96774.html)
//
//   - [Download the backup files of a PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/2627635.html)
//
// @param request - DescribeDownloadSupportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadSupportResponse
func (client *Client) DescribeDownloadSupportWithContext(ctx context.Context, request *DescribeDownloadSupportRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadSupportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDownloadSupport"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDownloadSupportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the advanced download tasks for an ApsaraDB RDS for MySQL instance, an ApsaraDB RDS for PostgreSQL instance, or a PolarDB for MySQL cluster.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - ApsaraDB RDS for MySQL
//
//   - ApsaraDB RDS for PostgreSQL
//
//   - PolarDB for MySQL
//
// ### [](#)References
//
//   - [Download the backup files of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/98819.html)
//
//   - [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96774.html)
//
//   - [Download the backup files of a PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/2627635.html)
//
// @param request - DescribeDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadTaskResponse
func (client *Client) DescribeDownloadTaskWithContext(ctx context.Context, request *DescribeDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeDownloadTaskResponse, _err error) {
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

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DatasourceId) {
		query["DatasourceId"] = request.DatasourceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.OrderColumn) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !dara.IsNil(request.OrderDirect) {
		query["OrderDirect"] = request.OrderDirect
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDownloadTask"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDownloadTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the snapshots of an instance.
//
// Description:
//
// Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/203154.html) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](https://help.aliyun.com/document_detail/185577.html). This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DescribeSandboxBackupSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSandboxBackupSetsResponse
func (client *Client) DescribeSandboxBackupSetsWithContext(ctx context.Context, request *DescribeSandboxBackupSetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSandboxBackupSetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
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
		Action:      dara.String("DescribeSandboxBackupSets"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSandboxBackupSetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries sandbox instances that are created within an account.
//
// Description:
//
// This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DescribeSandboxInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSandboxInstancesResponse
func (client *Client) DescribeSandboxInstancesWithContext(ctx context.Context, request *DescribeSandboxInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSandboxInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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
		Action:      dara.String("DescribeSandboxInstances"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSandboxInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the recoverable time range of a sandbox instance.
//
// Description:
//
// Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/203154.html) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](https://help.aliyun.com/document_detail/185577.html). This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DescribeSandboxRecoveryTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSandboxRecoveryTimeResponse
func (client *Client) DescribeSandboxRecoveryTimeWithContext(ctx context.Context, request *DescribeSandboxRecoveryTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSandboxRecoveryTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSandboxRecoveryTime"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSandboxRecoveryTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the backup policy of a PolarDB instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # PolarDB for MySQL
//
// >  This API operation is available only to specific customers. If you want to call this API operation, join the Database Backup (DBS) DingTalk group (ID 35585947) for customer consultation to request permissions.
//
// @param tmpReq - ModifyBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPolicyResponse
func (client *Client) ModifyBackupPolicyWithContext(ctx context.Context, tmpReq *ModifyBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyBackupPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AdvanceDataPolicies) {
		request.AdvanceDataPoliciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdvanceDataPolicies, dara.String("AdvanceDataPolicies"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AdvanceIncPolicies) {
		request.AdvanceIncPoliciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdvanceIncPolicies, dara.String("AdvanceIncPolicies"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AdvanceLogPolicies) {
		request.AdvanceLogPoliciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdvanceLogPolicies, dara.String("AdvanceLogPolicies"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdvanceDataPoliciesShrink) {
		query["AdvanceDataPolicies"] = request.AdvanceDataPoliciesShrink
	}

	if !dara.IsNil(request.AdvanceIncPoliciesShrink) {
		query["AdvanceIncPolicies"] = request.AdvanceIncPoliciesShrink
	}

	if !dara.IsNil(request.AdvanceLogPoliciesShrink) {
		query["AdvanceLogPolicies"] = request.AdvanceLogPoliciesShrink
	}

	if !dara.IsNil(request.BackupMethod) {
		query["BackupMethod"] = request.BackupMethod
	}

	if !dara.IsNil(request.BackupPriority) {
		query["BackupPriority"] = request.BackupPriority
	}

	if !dara.IsNil(request.BackupRetentionPolicyOnClusterDeletion) {
		query["BackupRetentionPolicyOnClusterDeletion"] = request.BackupRetentionPolicyOnClusterDeletion
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.EnableIncBackup) {
		query["EnableIncBackup"] = request.EnableIncBackup
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PreferredBackupWindowBegin) {
		query["PreferredBackupWindowBegin"] = request.PreferredBackupWindowBegin
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPolicy"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改库表恢复状态
//
// @param request - ModifyDBTablesRecoveryStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBTablesRecoveryStateResponse
func (client *Client) ModifyDBTablesRecoveryStateWithContext(ctx context.Context, request *ModifyDBTablesRecoveryStateRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBTablesRecoveryStateResponse, _err error) {
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

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBTablesRecoveryState"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBTablesRecoveryStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重试高级下载任务
//
// @param request - RetryDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryDownloadTaskResponse
func (client *Client) RetryDownloadTaskWithContext(ctx context.Context, request *RetryDownloadTaskRequest, runtime *dara.RuntimeOptions) (_result *RetryDownloadTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryDownloadTask"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryDownloadTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询是否支持库表恢复
//
// @param request - SupportDBTableRecoveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SupportDBTableRecoveryResponse
func (client *Client) SupportDBTableRecoveryWithContext(ctx context.Context, request *SupportDBTableRecoveryRequest, runtime *dara.RuntimeOptions) (_result *SupportDBTableRecoveryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SupportDBTableRecovery"),
		Version:     dara.String("2021-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SupportDBTableRecoveryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
