// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// This API is used to configure a DBS backup plan.
//
// @param request - ConfigureBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigureBackupPlanResponse
func (client *Client) ConfigureBackupPlanWithContext(ctx context.Context, request *ConfigureBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *ConfigureBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoStartBackup) {
		query["AutoStartBackup"] = request.AutoStartBackup
	}

	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupLogIntervalSeconds) {
		query["BackupLogIntervalSeconds"] = request.BackupLogIntervalSeconds
	}

	if !dara.IsNil(request.BackupObjects) {
		query["BackupObjects"] = request.BackupObjects
	}

	if !dara.IsNil(request.BackupPeriod) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupPlanName) {
		query["BackupPlanName"] = request.BackupPlanName
	}

	if !dara.IsNil(request.BackupRateLimit) {
		query["BackupRateLimit"] = request.BackupRateLimit
	}

	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.BackupSpeedLimit) {
		query["BackupSpeedLimit"] = request.BackupSpeedLimit
	}

	if !dara.IsNil(request.BackupStartTime) {
		query["BackupStartTime"] = request.BackupStartTime
	}

	if !dara.IsNil(request.BackupStorageEncryptMethod) {
		query["BackupStorageEncryptMethod"] = request.BackupStorageEncryptMethod
	}

	if !dara.IsNil(request.BackupStorageType) {
		query["BackupStorageType"] = request.BackupStorageType
	}

	if !dara.IsNil(request.BackupStrategyType) {
		query["BackupStrategyType"] = request.BackupStrategyType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossAliyunId) {
		query["CrossAliyunId"] = request.CrossAliyunId
	}

	if !dara.IsNil(request.CrossRoleName) {
		query["CrossRoleName"] = request.CrossRoleName
	}

	if !dara.IsNil(request.DuplicationArchivePeriod) {
		query["DuplicationArchivePeriod"] = request.DuplicationArchivePeriod
	}

	if !dara.IsNil(request.DuplicationInfrequentAccessPeriod) {
		query["DuplicationInfrequentAccessPeriod"] = request.DuplicationInfrequentAccessPeriod
	}

	if !dara.IsNil(request.EnableBackupLog) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !dara.IsNil(request.EnableMysqlPhysicalBackupBinlog) {
		query["EnableMysqlPhysicalBackupBinlog"] = request.EnableMysqlPhysicalBackupBinlog
	}

	if !dara.IsNil(request.EnableSourceEndpointSsl) {
		query["EnableSourceEndpointSsl"] = request.EnableSourceEndpointSsl
	}

	if !dara.IsNil(request.OSSBucketName) {
		query["OSSBucketName"] = request.OSSBucketName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceEndpointDatabaseName) {
		query["SourceEndpointDatabaseName"] = request.SourceEndpointDatabaseName
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		query["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointInstanceType) {
		query["SourceEndpointInstanceType"] = request.SourceEndpointInstanceType
	}

	if !dara.IsNil(request.SourceEndpointOracleHome) {
		query["SourceEndpointOracleHome"] = request.SourceEndpointOracleHome
	}

	if !dara.IsNil(request.SourceEndpointOracleSID) {
		query["SourceEndpointOracleSID"] = request.SourceEndpointOracleSID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		query["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		query["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	if !dara.IsNil(request.SslCaPem) {
		query["SslCaPem"] = request.SslCaPem
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigureBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigureBackupPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates, configures, and starts a backup plan.
//
// Description:
//
// Before you call this operation, ensure that you understand the [billing methods and pricing](https://help.aliyun.com/document_detail/70005.html) of Database Backup (DBS).
//
// @param request - CreateAndStartBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndStartBackupPlanResponse
func (client *Client) CreateAndStartBackupPlanWithContext(ctx context.Context, request *CreateAndStartBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateAndStartBackupPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupLogIntervalSeconds) {
		query["BackupLogIntervalSeconds"] = request.BackupLogIntervalSeconds
	}

	if !dara.IsNil(request.BackupMethod) {
		query["BackupMethod"] = request.BackupMethod
	}

	if !dara.IsNil(request.BackupObjects) {
		query["BackupObjects"] = request.BackupObjects
	}

	if !dara.IsNil(request.BackupPeriod) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupPlanName) {
		query["BackupPlanName"] = request.BackupPlanName
	}

	if !dara.IsNil(request.BackupRateLimit) {
		query["BackupRateLimit"] = request.BackupRateLimit
	}

	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.BackupSpeedLimit) {
		query["BackupSpeedLimit"] = request.BackupSpeedLimit
	}

	if !dara.IsNil(request.BackupStartTime) {
		query["BackupStartTime"] = request.BackupStartTime
	}

	if !dara.IsNil(request.BackupStorageType) {
		query["BackupStorageType"] = request.BackupStorageType
	}

	if !dara.IsNil(request.BackupStrategyType) {
		query["BackupStrategyType"] = request.BackupStrategyType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossAliyunId) {
		query["CrossAliyunId"] = request.CrossAliyunId
	}

	if !dara.IsNil(request.CrossRoleName) {
		query["CrossRoleName"] = request.CrossRoleName
	}

	if !dara.IsNil(request.DatabaseRegion) {
		query["DatabaseRegion"] = request.DatabaseRegion
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.DuplicationArchivePeriod) {
		query["DuplicationArchivePeriod"] = request.DuplicationArchivePeriod
	}

	if !dara.IsNil(request.DuplicationInfrequentAccessPeriod) {
		query["DuplicationInfrequentAccessPeriod"] = request.DuplicationInfrequentAccessPeriod
	}

	if !dara.IsNil(request.EnableBackupLog) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !dara.IsNil(request.FromApp) {
		query["FromApp"] = request.FromApp
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OSSBucketName) {
		query["OSSBucketName"] = request.OSSBucketName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceEndpointDatabaseName) {
		query["SourceEndpointDatabaseName"] = request.SourceEndpointDatabaseName
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		query["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointInstanceType) {
		query["SourceEndpointInstanceType"] = request.SourceEndpointInstanceType
	}

	if !dara.IsNil(request.SourceEndpointOracleSID) {
		query["SourceEndpointOracleSID"] = request.SourceEndpointOracleSID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		query["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		query["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	if !dara.IsNil(request.StorageRegion) {
		query["StorageRegion"] = request.StorageRegion
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAndStartBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndStartBackupPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Database Backup Service (DBS) backup plan.
//
// Description:
//
// To perform this operation in the console, see [Purchase a backup plan](https://help.aliyun.com/document_detail/65909.html).
//
// @param request - CreateBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupPlanResponse
func (client *Client) CreateBackupPlanWithContext(ctx context.Context, request *CreateBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupPlanResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DatabaseRegion) {
		query["DatabaseRegion"] = request.DatabaseRegion
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.FromApp) {
		query["FromApp"] = request.FromApp
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StorageRegion) {
		query["StorageRegion"] = request.StorageRegion
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation creates a task to download a full backup set.
//
// @param request - CreateFullBackupSetDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFullBackupSetDownloadResponse
func (client *Client) CreateFullBackupSetDownloadWithContext(ctx context.Context, request *CreateFullBackupSetDownloadRequest, runtime *dara.RuntimeOptions) (_result *CreateFullBackupSetDownloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupSetDataFormat) {
		query["BackupSetDataFormat"] = request.BackupSetDataFormat
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFullBackupSetDownload"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFullBackupSetDownloadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to obtain a database list by using a backup gateway.
//
// Description:
//
// This API operation returns a task ID. You can call the [GetDBListFromAgent](https://help.aliyun.com/document_detail/2869852.html) operation to query the task result based on the task ID.
//
// @param request - CreateGetDBListFromAgentTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGetDBListFromAgentTaskResponse
func (client *Client) CreateGetDBListFromAgentTaskWithContext(ctx context.Context, request *CreateGetDBListFromAgentTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateGetDBListFromAgentTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DatabaseType) {
		query["DatabaseType"] = request.DatabaseType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGetDBListFromAgentTask"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGetDBListFromAgentTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and starts an incremental backup set download task.
//
// @param request - CreateIncrementBackupSetDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIncrementBackupSetDownloadResponse
func (client *Client) CreateIncrementBackupSetDownloadWithContext(ctx context.Context, request *CreateIncrementBackupSetDownloadRequest, runtime *dara.RuntimeOptions) (_result *CreateIncrementBackupSetDownloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupSetDataFormat) {
		query["BackupSetDataFormat"] = request.BackupSetDataFormat
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.BackupSetName) {
		query["BackupSetName"] = request.BackupSetName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIncrementBackupSetDownload"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIncrementBackupSetDownloadResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface creates DBS restore jobs.
//
// Description:
//
// ### Related operations
//
// - [Recover databases](https://help.aliyun.com/document_detail/85543.html)
//
// - [Tutorials for various database restore configurations](https://help.aliyun.com/document_detail/197144.html)
//
// @param request - CreateRestoreTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRestoreTaskResponse
func (client *Client) CreateRestoreTaskWithContext(ctx context.Context, request *CreateRestoreTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRestoreTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoOpenDatabase) {
		query["AutoOpenDatabase"] = request.AutoOpenDatabase
	}

	if !dara.IsNil(request.AutoShutdownDatabase) {
		query["AutoShutdownDatabase"] = request.AutoShutdownDatabase
	}

	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossAliyunId) {
		query["CrossAliyunId"] = request.CrossAliyunId
	}

	if !dara.IsNil(request.CrossRoleName) {
		query["CrossRoleName"] = request.CrossRoleName
	}

	if !dara.IsNil(request.DestDatabaseInstanceClass) {
		query["DestDatabaseInstanceClass"] = request.DestDatabaseInstanceClass
	}

	if !dara.IsNil(request.DestDatabaseInstanceDatabaseVersion) {
		query["DestDatabaseInstanceDatabaseVersion"] = request.DestDatabaseInstanceDatabaseVersion
	}

	if !dara.IsNil(request.DestDatabaseInstanceRegion) {
		query["DestDatabaseInstanceRegion"] = request.DestDatabaseInstanceRegion
	}

	if !dara.IsNil(request.DestDatabaseInstanceStorageSize) {
		query["DestDatabaseInstanceStorageSize"] = request.DestDatabaseInstanceStorageSize
	}

	if !dara.IsNil(request.DestDatabaseInstanceType) {
		query["DestDatabaseInstanceType"] = request.DestDatabaseInstanceType
	}

	if !dara.IsNil(request.DestDatabaseInstanceVSwitch) {
		query["DestDatabaseInstanceVSwitch"] = request.DestDatabaseInstanceVSwitch
	}

	if !dara.IsNil(request.DestDatabaseInstanceVpc) {
		query["DestDatabaseInstanceVpc"] = request.DestDatabaseInstanceVpc
	}

	if !dara.IsNil(request.DestinationEndpointDatabaseName) {
		query["DestinationEndpointDatabaseName"] = request.DestinationEndpointDatabaseName
	}

	if !dara.IsNil(request.DestinationEndpointIP) {
		query["DestinationEndpointIP"] = request.DestinationEndpointIP
	}

	if !dara.IsNil(request.DestinationEndpointInstanceID) {
		query["DestinationEndpointInstanceID"] = request.DestinationEndpointInstanceID
	}

	if !dara.IsNil(request.DestinationEndpointInstanceType) {
		query["DestinationEndpointInstanceType"] = request.DestinationEndpointInstanceType
	}

	if !dara.IsNil(request.DestinationEndpointOracleSID) {
		query["DestinationEndpointOracleSID"] = request.DestinationEndpointOracleSID
	}

	if !dara.IsNil(request.DestinationEndpointPassword) {
		query["DestinationEndpointPassword"] = request.DestinationEndpointPassword
	}

	if !dara.IsNil(request.DestinationEndpointPort) {
		query["DestinationEndpointPort"] = request.DestinationEndpointPort
	}

	if !dara.IsNil(request.DestinationEndpointRegion) {
		query["DestinationEndpointRegion"] = request.DestinationEndpointRegion
	}

	if !dara.IsNil(request.DestinationEndpointUserName) {
		query["DestinationEndpointUserName"] = request.DestinationEndpointUserName
	}

	if !dara.IsNil(request.DuplicateConflict) {
		query["DuplicateConflict"] = request.DuplicateConflict
	}

	if !dara.IsNil(request.EnableDestinationEndpointSsl) {
		query["EnableDestinationEndpointSsl"] = request.EnableDestinationEndpointSsl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RestoreDestinationMode) {
		query["RestoreDestinationMode"] = request.RestoreDestinationMode
	}

	if !dara.IsNil(request.RestoreDir) {
		query["RestoreDir"] = request.RestoreDir
	}

	if !dara.IsNil(request.RestoreHome) {
		query["RestoreHome"] = request.RestoreHome
	}

	if !dara.IsNil(request.RestoreObjects) {
		query["RestoreObjects"] = request.RestoreObjects
	}

	if !dara.IsNil(request.RestoreTaskName) {
		query["RestoreTaskName"] = request.RestoreTaskName
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.SslCaPem) {
		query["SslCaPem"] = request.SslCaPem
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRestoreTask"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRestoreTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of backup gateways in Database Backup Service (DBS).
//
// @param request - DescribeBackupGatewayListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupGatewayListResponse
func (client *Client) DescribeBackupGatewayListWithContext(ctx context.Context, request *DescribeBackupGatewayListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupGatewayListResponse, _err error) {
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

	if !dara.IsNil(request.Identifier) {
		query["Identifier"] = request.Identifier
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupGatewayList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupGatewayListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation queries the billing information of a backup plan.
//
// @param request - DescribeBackupPlanBillingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPlanBillingResponse
func (client *Client) DescribeBackupPlanBillingWithContext(ctx context.Context, request *DescribeBackupPlanBillingRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPlanBillingResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ShowStorageType) {
		query["ShowStorageType"] = request.ShowStorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPlanBilling"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPlanBillingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation lets you view a DBS backup plan.
//
// Description:
//
// Before you use this operation, ensure that Object Storage Service (OSS) is enabled. For more information, see [Object Storage Service](https://help.aliyun.com/document_detail/31817.html).
//
// @param request - DescribeBackupPlanListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPlanListResponse
func (client *Client) DescribeBackupPlanListWithContext(ctx context.Context, request *DescribeBackupPlanListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupPlanListResponse, _err error) {
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

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupPlanName) {
		query["BackupPlanName"] = request.BackupPlanName
	}

	if !dara.IsNil(request.BackupPlanStatus) {
		query["BackupPlanStatus"] = request.BackupPlanStatus
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShowBackupStrategyInfo) {
		query["ShowBackupStrategyInfo"] = request.ShowBackupStrategyInfo
	}

	if !dara.IsNil(request.ShowRecoverTimeRange) {
		query["ShowRecoverTimeRange"] = request.ShowRecoverTimeRange
	}

	if !dara.IsNil(request.ShowStorageStrategyInfo) {
		query["ShowStorageStrategyInfo"] = request.ShowStorageStrategyInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPlanList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPlanListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of download tasks for backup sets in Database Backup Service (DBS).
//
// @param request - DescribeBackupSetDownloadTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSetDownloadTaskListResponse
func (client *Client) DescribeBackupSetDownloadTaskListWithContext(ctx context.Context, request *DescribeBackupSetDownloadTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupSetDownloadTaskListResponse, _err error) {
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

	if !dara.IsNil(request.BackupSetDownloadTaskId) {
		query["BackupSetDownloadTaskId"] = request.BackupSetDownloadTaskId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupSetDownloadTaskList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupSetDownloadTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Data Lake Analytics (DLA) service for a backup schedule.
//
// @param request - DescribeDLAServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDLAServiceResponse
func (client *Client) DescribeDLAServiceWithContext(ctx context.Context, request *DescribeDLAServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDLAServiceResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDLAService"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDLAServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to list full backup jobs in Database Backup Service (DBS).
//
// @param request - DescribeFullBackupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFullBackupListResponse
func (client *Client) DescribeFullBackupListWithContext(ctx context.Context, request *DescribeFullBackupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeFullBackupListResponse, _err error) {
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

	if !dara.IsNil(request.BackupSetStatus) {
		query["BackupSetStatus"] = request.BackupSetStatus
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ShowProgress) {
		query["ShowProgress"] = request.ShowProgress
	}

	if !dara.IsNil(request.ShowStorageType) {
		query["ShowStorageType"] = request.ShowStorageType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFullBackupList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFullBackupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation queries the list of incremental backup tasks for DBS.
//
// @param request - DescribeIncrementBackupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIncrementBackupListResponse
func (client *Client) DescribeIncrementBackupListWithContext(ctx context.Context, request *DescribeIncrementBackupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeIncrementBackupListResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ShowStorageType) {
		query["ShowStorageType"] = request.ShowStorageType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIncrementBackupList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIncrementBackupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the error code of a Database Backup Service (DBS) job.
//
// @param request - DescribeJobErrorCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeJobErrorCodeResponse
func (client *Client) DescribeJobErrorCodeWithContext(ctx context.Context, request *DescribeJobErrorCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeJobErrorCodeResponse, _err error) {
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

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeJobErrorCode"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeJobErrorCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the CIDR blocks of nodes on which Database Backup (DBS) is running.
//
// @param request - DescribeNodeCidrListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeCidrListResponse
func (client *Client) DescribeNodeCidrListWithContext(ctx context.Context, request *DescribeNodeCidrListRequest, runtime *dara.RuntimeOptions) (_result *DescribeNodeCidrListResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNodeCidrList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNodeCidrListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation queries the precheck progress for a backup plan or a restore job.
//
// @param request - DescribePreCheckProgressListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreCheckProgressListResponse
func (client *Client) DescribePreCheckProgressListWithContext(ctx context.Context, request *DescribePreCheckProgressListRequest, runtime *dara.RuntimeOptions) (_result *DescribePreCheckProgressListResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RestoreTaskId) {
		query["RestoreTaskId"] = request.RestoreTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePreCheckProgressList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePreCheckProgressListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions where DBS is available.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation returns the time ranges available for restoring data from a backup plan.
//
// @param request - DescribeRestoreRangeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreRangeInfoResponse
func (client *Client) DescribeRestoreRangeInfoWithContext(ctx context.Context, request *DescribeRestoreRangeInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreRangeInfoResponse, _err error) {
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

	if !dara.IsNil(request.BeginTimestampForRestore) {
		query["BeginTimestampForRestore"] = request.BeginTimestampForRestore
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestampForRestore) {
		query["EndTimestampForRestore"] = request.EndTimestampForRestore
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecentlyRestore) {
		query["RecentlyRestore"] = request.RecentlyRestore
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreRangeInfo"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreRangeInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries restore jobs in Database Backup Service (DBS).
//
// @param request - DescribeRestoreTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreTaskListResponse
func (client *Client) DescribeRestoreTaskListWithContext(ctx context.Context, request *DescribeRestoreTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreTaskListResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RestoreTaskId) {
		query["RestoreTaskId"] = request.RestoreTaskId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRestoreTaskList"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disable incremental backup for a backup plan.
//
// Description:
//
// ## Impact
//
// After you disable incremental backup, the backup plan no longer performs incremental backups.
//
// @param request - DisableBackupLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableBackupLogResponse
func (client *Client) DisableBackupLogWithContext(ctx context.Context, request *DisableBackupLogRequest, runtime *dara.RuntimeOptions) (_result *DisableBackupLogResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DisableMysqlPhysicalBackupBinlogOnly) {
		query["DisableMysqlPhysicalBackupBinlogOnly"] = request.DisableMysqlPhysicalBackupBinlogOnly
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableBackupLog"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableBackupLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation enables incremental backup for a backup plan.
//
// Description:
//
// ## Impact
//
// Enabling incremental backup incurs no additional charge. However, this operation generates backup traffic and consumes storage space. The fees for the traffic and storage are the same as those for a full backup. You can use the free quota from a backup plan or a storage plan to cover these costs.
//
// @param request - EnableBackupLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableBackupLogResponse
func (client *Client) EnableBackupLogWithContext(ctx context.Context, request *EnableBackupLogRequest, runtime *dara.RuntimeOptions) (_result *EnableBackupLogResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.EnableMysqlPhysicalBackupBinlog) {
		query["EnableMysqlPhysicalBackupBinlog"] = request.EnableMysqlPhysicalBackupBinlog
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableBackupLog"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableBackupLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of a task that is used to query a database list by using a backup gateway based on the ID of the task.
//
// Description:
//
// 您需要调用 [CreateGetDBListFromAgentTask](https://help.aliyun.com/document_detail/2869847.html) 接口创建一个异步任务获取 TaskId（异步任务 ID）。将 TaskId 传入 GetDBListFromAgent 接口后，即可获取物理备份库表数据。
//
// @param request - GetDBListFromAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDBListFromAgentResponse
func (client *Client) GetDBListFromAgentWithContext(ctx context.Context, request *GetDBListFromAgentRequest, runtime *dara.RuntimeOptions) (_result *GetDBListFromAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDBListFromAgent"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDBListFromAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the objects included in a Database Backup Service (DBS) backup plan.
//
// @param request - ModifyBackupObjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupObjectsResponse
func (client *Client) ModifyBackupObjectsWithContext(ctx context.Context, request *ModifyBackupObjectsRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupObjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupObjects) {
		query["BackupObjects"] = request.BackupObjects
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupObjects"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupObjectsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a backup plan.
//
// @param request - ModifyBackupPlanNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPlanNameResponse
func (client *Client) ModifyBackupPlanNameWithContext(ctx context.Context, request *ModifyBackupPlanNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupPlanNameResponse, _err error) {
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

	if !dara.IsNil(request.BackupPlanName) {
		query["BackupPlanName"] = request.BackupPlanName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPlanName"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupPlanNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables, configures, or disables the automatic download feature.
//
// Description:
//
// 使用本接口前请先确认备份数据是否存储在 DBS 的内置 OSS 上，您可通过调用 [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) 接口查看 BackupStorageType 取值情况。
//
// @param request - ModifyBackupSetDownloadRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupSetDownloadRulesResponse
func (client *Client) ModifyBackupSetDownloadRulesWithContext(ctx context.Context, request *ModifyBackupSetDownloadRulesRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupSetDownloadRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupSetDownloadDir) {
		query["BackupSetDownloadDir"] = request.BackupSetDownloadDir
	}

	if !dara.IsNil(request.BackupSetDownloadTargetType) {
		query["BackupSetDownloadTargetType"] = request.BackupSetDownloadTargetType
	}

	if !dara.IsNil(request.BackupSetDownloadTargetTypeLocation) {
		query["BackupSetDownloadTargetTypeLocation"] = request.BackupSetDownloadTargetTypeLocation
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FullDataFormat) {
		query["FullDataFormat"] = request.FullDataFormat
	}

	if !dara.IsNil(request.IncrementDataFormat) {
		query["IncrementDataFormat"] = request.IncrementDataFormat
	}

	if !dara.IsNil(request.OpenAutoDownload) {
		query["OpenAutoDownload"] = request.OpenAutoDownload
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupSetDownloadRules"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupSetDownloadRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation modifies a Database Backup source endpoint.
//
// @param request - ModifyBackupSourceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupSourceEndpointResponse
func (client *Client) ModifyBackupSourceEndpointWithContext(ctx context.Context, request *ModifyBackupSourceEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupSourceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupGatewayId) {
		query["BackupGatewayId"] = request.BackupGatewayId
	}

	if !dara.IsNil(request.BackupObjects) {
		query["BackupObjects"] = request.BackupObjects
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CrossAliyunId) {
		query["CrossAliyunId"] = request.CrossAliyunId
	}

	if !dara.IsNil(request.CrossRoleName) {
		query["CrossRoleName"] = request.CrossRoleName
	}

	if !dara.IsNil(request.EnableSourceEndpointSsl) {
		query["EnableSourceEndpointSsl"] = request.EnableSourceEndpointSsl
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SourceEndpointDatabaseName) {
		query["SourceEndpointDatabaseName"] = request.SourceEndpointDatabaseName
	}

	if !dara.IsNil(request.SourceEndpointIP) {
		query["SourceEndpointIP"] = request.SourceEndpointIP
	}

	if !dara.IsNil(request.SourceEndpointInstanceID) {
		query["SourceEndpointInstanceID"] = request.SourceEndpointInstanceID
	}

	if !dara.IsNil(request.SourceEndpointInstanceType) {
		query["SourceEndpointInstanceType"] = request.SourceEndpointInstanceType
	}

	if !dara.IsNil(request.SourceEndpointOracleHome) {
		query["SourceEndpointOracleHome"] = request.SourceEndpointOracleHome
	}

	if !dara.IsNil(request.SourceEndpointOracleSID) {
		query["SourceEndpointOracleSID"] = request.SourceEndpointOracleSID
	}

	if !dara.IsNil(request.SourceEndpointPassword) {
		query["SourceEndpointPassword"] = request.SourceEndpointPassword
	}

	if !dara.IsNil(request.SourceEndpointPort) {
		query["SourceEndpointPort"] = request.SourceEndpointPort
	}

	if !dara.IsNil(request.SourceEndpointRegion) {
		query["SourceEndpointRegion"] = request.SourceEndpointRegion
	}

	if !dara.IsNil(request.SourceEndpointUserName) {
		query["SourceEndpointUserName"] = request.SourceEndpointUserName
	}

	if !dara.IsNil(request.SslCaPem) {
		query["SslCaPem"] = request.SslCaPem
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupSourceEndpoint"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupSourceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the backup time of a backup schedule.
//
// @param request - ModifyBackupStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupStrategyResponse
func (client *Client) ModifyBackupStrategyWithContext(ctx context.Context, request *ModifyBackupStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupLogIntervalSeconds) {
		query["BackupLogIntervalSeconds"] = request.BackupLogIntervalSeconds
	}

	if !dara.IsNil(request.BackupPeriod) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !dara.IsNil(request.BackupPlanId) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !dara.IsNil(request.BackupStartTime) {
		query["BackupStartTime"] = request.BackupStartTime
	}

	if !dara.IsNil(request.BackupStrategyType) {
		query["BackupStrategyType"] = request.BackupStrategyType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupStrategy"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the lifecycle of stored data in a backup plan.
//
// @param request - ModifyStorageStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyStorageStrategyResponse
func (client *Client) ModifyStorageStrategyWithContext(ctx context.Context, request *ModifyStorageStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyStorageStrategyResponse, _err error) {
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

	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.BackupStorageEncryptMethod) {
		query["BackupStorageEncryptMethod"] = request.BackupStorageEncryptMethod
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DuplicationArchivePeriod) {
		query["DuplicationArchivePeriod"] = request.DuplicationArchivePeriod
	}

	if !dara.IsNil(request.DuplicationInfrequentAccessPeriod) {
		query["DuplicationInfrequentAccessPeriod"] = request.DuplicationInfrequentAccessPeriod
	}

	if !dara.IsNil(request.IncrementBackupRetentionPeriod) {
		query["IncrementBackupRetentionPeriod"] = request.IncrementBackupRetentionPeriod
	}

	if !dara.IsNil(request.IncrementDuplicationArchivePeriod) {
		query["IncrementDuplicationArchivePeriod"] = request.IncrementDuplicationArchivePeriod
	}

	if !dara.IsNil(request.IncrementDuplicationInfrequentAccessPeriod) {
		query["IncrementDuplicationInfrequentAccessPeriod"] = request.IncrementDuplicationInfrequentAccessPeriod
	}

	if !dara.IsNil(request.LogBackupRetentionPeriod) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !dara.IsNil(request.LogDuplicationArchivePeriod) {
		query["LogDuplicationArchivePeriod"] = request.LogDuplicationArchivePeriod
	}

	if !dara.IsNil(request.LogDuplicationInfrequentAccessPeriod) {
		query["LogDuplicationInfrequentAccessPeriod"] = request.LogDuplicationInfrequentAccessPeriod
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyStorageStrategy"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyStorageStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation releases a pay-as-you-go backup plan.
//
// Description:
//
// ## Impact
//
// After you release a backup plan, the service for the backup instance is stopped and you are no longer charged for the instance.
//
// @param request - ReleaseBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseBackupPlanResponse
func (client *Client) ReleaseBackupPlanWithContext(ctx context.Context, request *ReleaseBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *ReleaseBackupPlanResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseBackupPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a backup schedule.
//
// @param request - RenewBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewBackupPlanResponse
func (client *Client) RenewBackupPlanWithContext(ctx context.Context, request *RenewBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *RenewBackupPlanResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewBackupPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation starts a DBS backup plan.
//
// @param request - StartBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartBackupPlanResponse
func (client *Client) StartBackupPlanWithContext(ctx context.Context, request *StartBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *StartBackupPlanResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartBackupPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a DBS restore job.
//
// @param request - StartRestoreTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRestoreTaskResponse
func (client *Client) StartRestoreTaskWithContext(ctx context.Context, request *StartRestoreTaskRequest, runtime *dara.RuntimeOptions) (_result *StartRestoreTaskResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RestoreTaskId) {
		query["RestoreTaskId"] = request.RestoreTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRestoreTask"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRestoreTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation pauses a DBS backup plan.
//
// @param request - StopBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopBackupPlanResponse
func (client *Client) StopBackupPlanWithContext(ctx context.Context, request *StopBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *StopBackupPlanResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.StopMethod) {
		query["StopMethod"] = request.StopMethod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopBackupPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades a backup schedule.
//
// @param request - UpgradeBackupPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeBackupPlanResponse
func (client *Client) UpgradeBackupPlanWithContext(ctx context.Context, request *UpgradeBackupPlanRequest, runtime *dara.RuntimeOptions) (_result *UpgradeBackupPlanResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeBackupPlan"),
		Version:     dara.String("2019-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeBackupPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
