// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - AlignStoragePrimaryAzoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlignStoragePrimaryAzoneResponse
func (client *Client) AlignStoragePrimaryAzoneWithContext(ctx context.Context, request *AlignStoragePrimaryAzoneRequest, runtime *dara.RuntimeOptions) (_result *AlignStoragePrimaryAzoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StorageInstanceName) {
		query["StorageInstanceName"] = request.StorageInstanceName
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlignStoragePrimaryAzone"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AlignStoragePrimaryAzoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the AllocateColdDataVolume operation.
//
// @param request - AllocateColdDataVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateColdDataVolumeResponse
func (client *Client) AllocateColdDataVolumeWithContext(ctx context.Context, request *AllocateColdDataVolumeRequest, runtime *dara.RuntimeOptions) (_result *AllocateColdDataVolumeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateColdDataVolume"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateColdDataVolumeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the AllocateInstancePublicConnection operation to create a public IP address.
//
// @param request - AllocateInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateInstancePublicConnectionResponse
func (client *Client) AllocateInstancePublicConnectionWithContext(ctx context.Context, request *AllocateInstancePublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.InstanceClusterName) {
		query["InstanceClusterName"] = request.InstanceClusterName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateInstancePublicConnection"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a public domain name for a Mem0 instance.
//
// Description:
//
// This operation is used to confirm that no active connections exist before a rollback task, to ensure operation safety.
//
// @param request - AllocateMem0PublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateMem0PublicConnectionResponse
func (client *Client) AllocateMem0PublicConnectionWithContext(ctx context.Context, request *AllocateMem0PublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *AllocateMem0PublicConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateMem0PublicConnection"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateMem0PublicConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Mounts a columnar instance to a specified primary database instance.
//
// @param request - AttachColumnarInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachColumnarInstanceResponse
func (client *Client) AttachColumnarInstanceWithContext(ctx context.Context, request *AttachColumnarInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachColumnarInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachColumnarInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachColumnarInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels active O&M event tasks by calling the CancelActiveOperationTasks operation.
//
// @param request - CancelActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelActiveOperationTasksResponse
func (client *Client) CancelActiveOperationTasksWithContext(ctx context.Context, request *CancelActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *CancelActiveOperationTasksResponse, _err error) {
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
		Action:      dara.String("CancelActiveOperationTasks"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelActiveOperationTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the resource group of an instance.
//
// Description:
//
// Note:
//
// - The **endpoint*	- differs from other operations. Use **polardbx.aliyuncs.com*	- for Chinese mainland regions and Singapore. For other regions, use **polardbx.{region id}.aliyunc.com**.
//
// - When testing this API operation, if a service unavailable error is returned, verify that the **endpoint*	- is correct. You can switch the **service address*	- to **Dubai*	- or **India*	- to change the **endpoint*	- to **polardbx.aliyuncs.com**.
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
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Version:     dara.String("2020-02-02"),
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
// Checks whether a PolarDB-X instance is authorized to use Key Management Service (KMS).
//
// @param request - CheckCloudResourceAuthorizedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCloudResourceAuthorizedResponse
func (client *Client) CheckCloudResourceAuthorizedWithContext(ctx context.Context, request *CheckCloudResourceAuthorizedRequest, runtime *dara.RuntimeOptions) (_result *CheckCloudResourceAuthorizedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCloudResourceAuthorized"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCloudResourceAuthorizedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a service health check.
//
// @param request - CheckHealthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckHealthResponse
func (client *Client) CheckHealthWithContext(ctx context.Context, request *CheckHealthRequest, runtime *dara.RuntimeOptions) (_result *CheckHealthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("CheckHealth"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckHealthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether SQL audit logs of a specified database instance have been successfully connected to Simple Log Service (SLS).
//
// After the call, the system returns the connection status between the SQL audit feature and SLS for the current instance, the project and Logstore configuration information, and whether synchronization is normal.
//
// Description:
//
// > 	- The SQL audit and analysis feature of PolarDB-X 2.0 is free of charge. However, Simple Log Service charges fees for storage space, read traffic, number of requests, data transformation, and data delivery. For more information about the SQL audit feature, see [Enable SQL audit and analysis](https://help.aliyun.com/document_detail/184619.html).
//
// @param request - CheckSqlAuditSlsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSqlAuditSlsStatusResponse
func (client *Client) CheckSqlAuditSlsStatusWithContext(ctx context.Context, request *CheckSqlAuditSlsStatusRequest, runtime *dara.RuntimeOptions) (_result *CheckSqlAuditSlsStatusResponse, _err error) {
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
		Action:      dara.String("CheckSqlAuditSlsStatus"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSqlAuditSlsStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Closes the database engine migration process for a specified instance. After you start a data migration task from another database (such as a self-managed MySQL database or an ApsaraDB RDS instance) to PolarDB-X, you can call this operation to safely stop the migration process if you need to terminate or clean up the migration state.
//
// @param request - CloseEngineMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseEngineMigrationResponse
func (client *Client) CloseEngineMigrationWithContext(ctx context.Context, request *CloseEngineMigrationRequest, runtime *dara.RuntimeOptions) (_result *CloseEngineMigrationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContinueEnableBinlog) {
		query["ContinueEnableBinlog"] = request.ContinueEnableBinlog
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseEngineMigration"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseEngineMigrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Confirms that no active connections exist before performing a rollback switchover.
//
// Description:
//
// Confirms that no active connections exist before a rollback task to ensure operation safety.
//
// @param request - ConfirmNoConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmNoConnectionResponse
func (client *Client) ConfirmNoConnectionWithContext(ctx context.Context, request *ConfirmNoConnectionRequest, runtime *dara.RuntimeOptions) (_result *ConfirmNoConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmNoConnection"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmNoConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account by calling the CreateAccount operation.
//
// @param request - CreateAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountResponse
func (client *Client) CreateAccountWithContext(ctx context.Context, request *CreateAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateBackup operation to create a backup.
//
// @param request - CreateBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackupResponse
func (client *Client) CreateBackupWithContext(ctx context.Context, request *CreateBackupRequest, runtime *dara.RuntimeOptions) (_result *CreateBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackup"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom endpoint for a database instance.
//
// Description:
//
// <props="china">For more information about instance accounts, see [Account management](https://help.aliyun.com/document_detail/172163.html)..
//
// @param request - CreateCustomEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomEndpointResponse
func (client *Client) CreateCustomEndpointWithContext(ctx context.Context, request *CreateCustomEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeAutoEnter) {
		query["NodeAutoEnter"] = request.NodeAutoEnter
	}

	if !dara.IsNil(request.NodeIds) {
		query["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.NodeRole) {
		query["NodeRole"] = request.NodeRole
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomEndpoint"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database by calling the CreateDB operation.
//
// @param request - CreateDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBResponse
func (client *Client) CreateDBWithContext(ctx context.Context, request *CreateDBRequest, runtime *dara.RuntimeOptions) (_result *CreateDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.Charset) {
		query["Charset"] = request.Charset
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbDescription) {
		query["DbDescription"] = request.DbDescription
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	if !dara.IsNil(request.StoragePoolName) {
		query["StoragePoolName"] = request.StoragePoolName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDB"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateDBInstance operation to create a PolarDB-X instance.
//
// Description:
//
// ***
//
// @param tmpReq - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstanceWithContext(ctx context.Context, tmpReq *CreateDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDBInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExtraParams) {
		request.ExtraParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtraParams, dara.String("ExtraParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.CNNodeCount) {
		query["CNNodeCount"] = request.CNNodeCount
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CnClass) {
		query["CnClass"] = request.CnClass
	}

	if !dara.IsNil(request.DBNodeClass) {
		query["DBNodeClass"] = request.DBNodeClass
	}

	if !dara.IsNil(request.DBNodeCount) {
		query["DBNodeCount"] = request.DBNodeCount
	}

	if !dara.IsNil(request.DNNodeCount) {
		query["DNNodeCount"] = request.DNNodeCount
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DnClass) {
		query["DnClass"] = request.DnClass
	}

	if !dara.IsNil(request.DnStorageSpace) {
		query["DnStorageSpace"] = request.DnStorageSpace
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.ExtraParamsShrink) {
		query["ExtraParams"] = request.ExtraParamsShrink
	}

	if !dara.IsNil(request.IsColumnarReadDBInstance) {
		query["IsColumnarReadDBInstance"] = request.IsColumnarReadDBInstance
	}

	if !dara.IsNil(request.IsReadDBInstance) {
		query["IsReadDBInstance"] = request.IsReadDBInstance
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.OriginMinorVersion) {
		query["OriginMinorVersion"] = request.OriginMinorVersion
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrimaryDBInstanceName) {
		query["PrimaryDBInstanceName"] = request.PrimaryDBInstanceName
	}

	if !dara.IsNil(request.PrimaryZone) {
		query["PrimaryZone"] = request.PrimaryZone
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecondaryZone) {
		query["SecondaryZone"] = request.SecondaryZone
	}

	if !dara.IsNil(request.Series) {
		query["Series"] = request.Series
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.TertiaryZone) {
		query["TertiaryZone"] = request.TertiaryZone
	}

	if !dara.IsNil(request.TopologyType) {
		query["TopologyType"] = request.TopologyType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data import task to import external data files, such as SQL scripts and CSV files, into a target database instance.
//
// Description:
//
// Creates a data import task that imports SQL or CSV files stored in OSS or ECS, or directly provided files, into a target database instance. By specifying the instance ID, database name, engine type, data source (such as an OSS path), and import type, the system performs data write operations asynchronously or synchronously. This operation is applicable to scenarios such as data migration, initialization, and data backfill. A task ID is returned for subsequent status queries and management.
//
// @param request - CreateDataImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataImportTaskResponse
func (client *Client) CreateDataImportTaskWithContext(ctx context.Context, request *CreateDataImportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDataImportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DstDb) {
		query["DstDb"] = request.DstDb
	}

	if !dara.IsNil(request.DstPassword) {
		query["DstPassword"] = request.DstPassword
	}

	if !dara.IsNil(request.DstResId) {
		query["DstResId"] = request.DstResId
	}

	if !dara.IsNil(request.DstUserName) {
		query["DstUserName"] = request.DstUserName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	if !dara.IsNil(request.SrcDb) {
		query["SrcDb"] = request.SrcDb
	}

	if !dara.IsNil(request.SrcPassword) {
		query["SrcPassword"] = request.SrcPassword
	}

	if !dara.IsNil(request.SrcResId) {
		query["SrcResId"] = request.SrcResId
	}

	if !dara.IsNil(request.SrcUserName) {
		query["SrcUserName"] = request.SrcUserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataImportTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AI gateway consumer for a PolarDB-X instance.
//
// Description:
//
// ***.
//
// @param request - CreateGatewayConsumerForPolarDBXRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayConsumerForPolarDBXResponse
func (client *Client) CreateGatewayConsumerForPolarDBXWithContext(ctx context.Context, request *CreateGatewayConsumerForPolarDBXRequest, runtime *dara.RuntimeOptions) (_result *CreateGatewayConsumerForPolarDBXResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGatewayConsumerForPolarDBX"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayConsumerForPolarDBXResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Global Database Network (GDN) instance.
//
// @param request - CreateGdnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGdnInstanceResponse
func (client *Client) CreateGdnInstanceWithContext(ctx context.Context, request *CreateGdnInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateGdnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GdnMode) {
		query["GdnMode"] = request.GdnMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RplConflictStrategy) {
		query["RplConflictStrategy"] = request.RplConflictStrategy
	}

	if !dara.IsNil(request.RplDmlStrategy) {
		query["RplDmlStrategy"] = request.RplDmlStrategy
	}

	if !dara.IsNil(request.RplSyncDdl) {
		query["RplSyncDdl"] = request.RplSyncDdl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGdnInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGdnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建GDN从实例
//
// Description:
//
// <props="china">更多关于实例账号的信息，请参见[账号管理](https://help.aliyun.com/document_detail/172163.html)。
//
// @param request - CreateGdnStandbyMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGdnStandbyMemberResponse
func (client *Client) CreateGdnStandbyMemberWithContext(ctx context.Context, request *CreateGdnStandbyMemberRequest, runtime *dara.RuntimeOptions) (_result *CreateGdnStandbyMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.CNNodeCount) {
		query["CNNodeCount"] = request.CNNodeCount
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloneInstanceName) {
		query["CloneInstanceName"] = request.CloneInstanceName
	}

	if !dara.IsNil(request.CnClass) {
		query["CnClass"] = request.CnClass
	}

	if !dara.IsNil(request.DNNodeCount) {
		query["DNNodeCount"] = request.DNNodeCount
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DnClass) {
		query["DnClass"] = request.DnClass
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrimaryZone) {
		query["PrimaryZone"] = request.PrimaryZone
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecondaryZone) {
		query["SecondaryZone"] = request.SecondaryZone
	}

	if !dara.IsNil(request.Series) {
		query["Series"] = request.Series
	}

	if !dara.IsNil(request.SourceInstanceRegion) {
		query["SourceInstanceRegion"] = request.SourceInstanceRegion
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.TertiaryZone) {
		query["TertiaryZone"] = request.TertiaryZone
	}

	if !dara.IsNil(request.TopologyType) {
		query["TopologyType"] = request.TopologyType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGdnStandbyMember"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGdnStandbyMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enable Memory Engine
//
// Description:
//
// <props="china">For more information about instance accounts, see [Account management](https://help.aliyun.com/document_detail/172163.html).
//
// @param request - CreateMem0Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMem0Response
func (client *Client) CreateMem0WithContext(ctx context.Context, request *CreateMem0Request, runtime *dara.RuntimeOptions) (_result *CreateMem0Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMem0"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMem0Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Supabase instance.
//
// Description:
//
// ***
//
// @param request - CreatePolardbxSupabaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolardbxSupabaseInstanceResponse
func (client *Client) CreatePolardbxSupabaseInstanceWithContext(ctx context.Context, request *CreatePolardbxSupabaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreatePolardbxSupabaseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DashboardPassword) {
		query["DashboardPassword"] = request.DashboardPassword
	}

	if !dara.IsNil(request.DbInstanceDescription) {
		query["DbInstanceDescription"] = request.DbInstanceDescription
	}

	if !dara.IsNil(request.DbPassword) {
		query["DbPassword"] = request.DbPassword
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TenantMode) {
		query["TenantMode"] = request.TenantMode
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolardbxSupabaseInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolardbxSupabaseInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a health check on the replication task during data migration.
//
// Description:
//
// During the data synchronization phase, proactively initiates a diagnostic task for the replication task to check for issues such as latency, replication interruption, or data inconsistency.
//
// @param request - CreateRplInspectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRplInspectionTaskResponse
func (client *Client) CreateRplInspectionTaskWithContext(ctx context.Context, request *CreateRplInspectionTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRplInspectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DstDb) {
		query["DstDb"] = request.DstDb
	}

	if !dara.IsNil(request.DstPassword) {
		query["DstPassword"] = request.DstPassword
	}

	if !dara.IsNil(request.DstResId) {
		query["DstResId"] = request.DstResId
	}

	if !dara.IsNil(request.DstUserName) {
		query["DstUserName"] = request.DstUserName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	if !dara.IsNil(request.SrcPassword) {
		query["SrcPassword"] = request.SrcPassword
	}

	if !dara.IsNil(request.SrcUserName) {
		query["SrcUserName"] = request.SrcUserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRplInspectionTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRplInspectionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a SQL statement assessment task. You submit SQL statements through this operation, and the system performs static analysis and risk assessment to identify execution risks and compatibility issues on the target database instance.
//
// @param request - CreateSQLEvaluateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSQLEvaluateTaskResponse
func (client *Client) CreateSQLEvaluateTaskWithContext(ctx context.Context, request *CreateSQLEvaluateTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSQLEvaluateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DstDb) {
		query["DstDb"] = request.DstDb
	}

	if !dara.IsNil(request.DstPassword) {
		query["DstPassword"] = request.DstPassword
	}

	if !dara.IsNil(request.DstResId) {
		query["DstResId"] = request.DstResId
	}

	if !dara.IsNil(request.DstUserName) {
		query["DstUserName"] = request.DstUserName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskDesc) {
		query["SlinkTaskDesc"] = request.SlinkTaskDesc
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	if !dara.IsNil(request.SrcDb) {
		query["SrcDb"] = request.SrcDb
	}

	if !dara.IsNil(request.SrcPassword) {
		query["SrcPassword"] = request.SrcPassword
	}

	if !dara.IsNil(request.SrcResId) {
		query["SrcResId"] = request.SrcResId
	}

	if !dara.IsNil(request.SrcResType) {
		query["SrcResType"] = request.SrcResType
	}

	if !dara.IsNil(request.SrcUserName) {
		query["SrcUserName"] = request.SrcUserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSQLEvaluateTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSQLEvaluateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource storage pool.
//
// @param request - CreateStoragePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStoragePoolResponse
func (client *Client) CreateStoragePoolWithContext(ctx context.Context, request *CreateStoragePoolRequest, runtime *dara.RuntimeOptions) (_result *CreateStoragePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StoragePoolDNList) {
		query["StoragePoolDNList"] = request.StoragePoolDNList
	}

	if !dara.IsNil(request.StoragePoolName) {
		query["StoragePoolName"] = request.StoragePoolName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStoragePool"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStoragePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database schema import task. This operation allows you to import SQL script files or text content that contains DDL statements into a target database instance, and automatically performs schema operations such as creating tables, indexes, views, and stored procedures.
//
// @param request - CreateStructureImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStructureImportTaskResponse
func (client *Client) CreateStructureImportTaskWithContext(ctx context.Context, request *CreateStructureImportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateStructureImportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStructureImportTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStructureImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom endpoint.
//
// Description:
//
// ***
//
// @param request - CreateSubCNInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSubCNInstanceResponse
func (client *Client) CreateSubCNInstanceWithContext(ctx context.Context, request *CreateSubCNInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateSubCNInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.IsAutoCreate) {
		query["IsAutoCreate"] = request.IsAutoCreate
	}

	if !dara.IsNil(request.ReadType) {
		query["ReadType"] = request.ReadType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSubCNInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSubCNInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the public endpoint for a Supabase instance.
//
// Description:
//
// ***
//
// @param request - CreateSupabaseNetTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSupabaseNetTypeResponse
func (client *Client) CreateSupabaseNetTypeWithContext(ctx context.Context, request *CreateSupabaseNetTypeRequest, runtime *dara.RuntimeOptions) (_result *CreateSupabaseNetTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSupabaseNetType"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSupabaseNetTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a privileged user. Each instance can have only one privileged user.
//
// @param request - CreateSuperAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSuperAccountResponse
func (client *Client) CreateSuperAccountWithContext(ctx context.Context, request *CreateSuperAccountRequest, runtime *dara.RuntimeOptions) (_result *CreateSuperAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSuperAccount"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSuperAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a conversion task to change a specific status or configuration of an instance.
//
// @param request - CreateTransformOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransformOperationResponse
func (client *Client) CreateTransformOperationWithContext(ctx context.Context, request *CreateTransformOperationRequest, runtime *dara.RuntimeOptions) (_result *CreateTransformOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.Operation) {
		query["Operation"] = request.Operation
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTransformOperation"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTransformOperationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an account by calling the DeleteAccount operation.
//
// @param request - DeleteAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountResponse
func (client *Client) DeleteAccountWithContext(ctx context.Context, request *DeleteAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom endpoint domain name.
//
// Description:
//
// Deletes the custom endpoint of a specified database instance and disables access through the domain name.
//
// @param request - DeleteCustomEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomEndpointResponse
func (client *Client) DeleteCustomEndpointWithContext(ctx context.Context, request *DeleteCustomEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomEndpointId) {
		query["CustomEndpointId"] = request.CustomEndpointId
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomEndpoint"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a database by calling the DeleteDB operation.
//
// @param request - DeleteDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBResponse
func (client *Client) DeleteDBWithContext(ctx context.Context, request *DeleteDBRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDB"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DeleteDBInstance operation to delete an instance.
//
// @param request - DeleteDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstanceResponse
func (client *Client) DeleteDBInstanceWithContext(ctx context.Context, request *DeleteDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an evaluation import task.
//
// Description:
//
// Deletes a created evaluation task and performs subsequent data import operations.
//
// @param request - DeleteEvaluateAndImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEvaluateAndImportTaskResponse
func (client *Client) DeleteEvaluateAndImportTaskWithContext(ctx context.Context, request *DeleteEvaluateAndImportTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteEvaluateAndImportTaskResponse, _err error) {
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

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEvaluateAndImportTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEvaluateAndImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a global database network (GDN) instance.
//
// @param request - DeleteGdnInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGdnInstanceResponse
func (client *Client) DeleteGdnInstanceWithContext(ctx context.Context, request *DeleteGdnInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteGdnInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GdnInstanceName) {
		query["GdnInstanceName"] = request.GdnInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGdnInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGdnInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the memory engine.
//
// Description:
//
// Deletes the custom endpoint of a specified database instance and disables the access entry for the domain name.
//
// @param request - DeleteMem0Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMem0Response
func (client *Client) DeleteMem0WithContext(ctx context.Context, request *DeleteMem0Request, runtime *dara.RuntimeOptions) (_result *DeleteMem0Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMem0"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMem0Response{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Supabase instance.
//
// Description:
//
// ***
//
// @param request - DeletePolardbxSupabaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolardbxSupabaseInstanceResponse
func (client *Client) DeletePolardbxSupabaseInstanceWithContext(ctx context.Context, request *DeletePolardbxSupabaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeletePolardbxSupabaseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolardbxSupabaseInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolardbxSupabaseInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义地址
//
// Description:
//
// ***
//
// @param request - DeleteSubCNInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSubCNInstanceResponse
func (client *Client) DeleteSubCNInstanceWithContext(ctx context.Context, request *DeleteSubCNInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteSubCNInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.InstanceClusterName) {
		query["InstanceClusterName"] = request.InstanceClusterName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSubCNInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSubCNInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of a Supabase instance.
//
// Description:
//
// ***
//
// @param request - DeleteSupabaseNetTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSupabaseNetTypeResponse
func (client *Client) DeleteSupabaseNetTypeWithContext(ctx context.Context, request *DeleteSupabaseNetTypeRequest, runtime *dara.RuntimeOptions) (_result *DeleteSupabaseNetTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSupabaseNetType"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSupabaseNetTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of accounts by calling the DescribeAccountList operation.
//
// @param request - DescribeAccountListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountListResponse
func (client *Client) DescribeAccountListWithContext(ctx context.Context, request *DescribeAccountListRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeActiveOperationMaintainConf operation to display the O&M window configuration.
//
// @param request - DescribeActiveOperationMaintainConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationMaintainConfResponse
func (client *Client) DescribeActiveOperationMaintainConfWithContext(ctx context.Context, request *DescribeActiveOperationMaintainConfRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationMaintainConfResponse, _err error) {
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
		Action:      dara.String("DescribeActiveOperationMaintainConf"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationMaintainConfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the total number of O&M events.
//
// @param request - DescribeActiveOperationTaskCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTaskCountResponse
func (client *Client) DescribeActiveOperationTaskCountWithContext(ctx context.Context, request *DescribeActiveOperationTaskCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTaskCountResponse, _err error) {
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
		Action:      dara.String("DescribeActiveOperationTaskCount"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTaskCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries O&M events by calling the DescribeActiveOperationTasks operation.
//
// @param request - DescribeActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTasksResponse
func (client *Client) DescribeActiveOperationTasksWithContext(ctx context.Context, request *DescribeActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTasksResponse, _err error) {
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
		Action:      dara.String("DescribeActiveOperationTasks"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists cold storage tables.
//
// @param request - DescribeArchiveTableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeArchiveTableListResponse
func (client *Client) DescribeArchiveTableListWithContext(ctx context.Context, request *DescribeArchiveTableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeArchiveTableListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeArchiveTableList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeArchiveTableListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of available regions that support cross-region operations, typically used for multi-region deployment or data replication scenarios.
//
// @param request - DescribeAvailableCrossRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableCrossRegionsResponse
func (client *Client) DescribeAvailableCrossRegionsWithContext(ctx context.Context, request *DescribeAvailableCrossRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableCrossRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableCrossRegions"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableCrossRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeBackupPolicy operation to query the backup settings of an instance.
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
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2020-02-02"),
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
// Queries the details of a backup set by calling the DescribeBackupSet operation.
//
// @param request - DescribeBackupSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSetResponse
func (client *Client) DescribeBackupSetWithContext(ctx context.Context, request *DescribeBackupSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupSetResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DestCrossRegion) {
		query["DestCrossRegion"] = request.DestCrossRegion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupSet"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeBackupSetList operation to query the list of backup sets.
//
// @param request - DescribeBackupSetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupSetListResponse
func (client *Client) DescribeBackupSetListWithContext(ctx context.Context, request *DescribeBackupSetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupSetListResponse, _err error) {
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
		Action:      dara.String("DescribeBackupSetList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupSetListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeBinaryLogList operation to query binlog logs.
//
// Description:
//
// - Binlog files are retained for 15 days by default.
//
// - The returned log list includes all logs whose record end time is after the query start time and whose record start time is before the query end time.
//
// - When the DownloadLink is not NULL, you can use this URL to download the backup file. This URL is valid for 2 days after it is generated. Download the file before the URL expires.
//
// @param request - DescribeBinaryLogListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBinaryLogListResponse
func (client *Client) DescribeBinaryLogListWithContext(ctx context.Context, request *DescribeBinaryLogListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBinaryLogListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeBinaryLogList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBinaryLogListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of instance specifications.
//
// Description:
//
// - Binary log files are retained for 15 days by default.
//
// - The returned log list includes all logs whose log record end time is after the query start time and whose log record start time is before the query end time.
//
// - If DownloadLink is not NULL, you can use this URL to download the backup file. This URL is valid for 2 days after it is generated. Download the file before the URL expires.
//
// @param request - DescribeCdcClassListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdcClassListResponse
func (client *Client) DescribeCdcClassListWithContext(ctx context.Context, request *DescribeCdcClassListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdcClassListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdcClassList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdcClassListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries CDC information.
//
// @param request - DescribeCdcInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdcInfoResponse
func (client *Client) DescribeCdcInfoWithContext(ctx context.Context, request *DescribeCdcInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdcInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdcInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdcInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of CDC versions.
//
// Description:
//
// - Binary log files are retained for 15 days by default.
//
// - The returned log list includes all logs whose log record end time is after the query start time and whose log record start time is before the query end time.
//
// - When DownloadLink is not NULL, you can download the backup file from this URL. The URL is valid for 2 days after it is generated. Download the file before the URL expires.
//
// @param request - DescribeCdcVersionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdcVersionListResponse
func (client *Client) DescribeCdcVersionListWithContext(ctx context.Context, request *DescribeCdcVersionListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdcVersionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdcVersionList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdcVersionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the DescribeCharacterSet operation to query the character set types supported by databases in a target instance.
//
// @param request - DescribeCharacterSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCharacterSetResponse
func (client *Client) DescribeCharacterSetWithContext(ctx context.Context, request *DescribeCharacterSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeCharacterSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCharacterSet"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCharacterSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The cold storage basic information.
//
// @param request - DescribeColdDataBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColdDataBasicInfoResponse
func (client *Client) DescribeColdDataBasicInfoWithContext(ctx context.Context, request *DescribeColdDataBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeColdDataBasicInfoResponse, _err error) {
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
		Action:      dara.String("DescribeColdDataBasicInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColdDataBasicInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of column store specifications for an instance.
//
// @param request - DescribeColumnarClassListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnarClassListResponse
func (client *Client) DescribeColumnarClassListWithContext(ctx context.Context, request *DescribeColumnarClassListRequest, runtime *dara.RuntimeOptions) (_result *DescribeColumnarClassListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeColumnarClassList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColumnarClassListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries column store information.
//
// @param request - DescribeColumnarInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnarInfoResponse
func (client *Client) DescribeColumnarInfoWithContext(ctx context.Context, request *DescribeColumnarInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeColumnarInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeColumnarInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColumnarInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries column store version information.
//
// Description:
//
// - Binary log files are retained for 15 days by default.
//
// - The returned log list includes all logs whose log record end time is after the query start time and whose log record start time is before the query end time.
//
// - When DownloadLink is not NULL, you can download the backup file from this URL. This URL is valid for 2 days after it is generated. Download the file before the expiration time.
//
// @param request - DescribeColumnarVersionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeColumnarVersionListResponse
func (client *Client) DescribeColumnarVersionListWithContext(ctx context.Context, request *DescribeColumnarVersionListRequest, runtime *dara.RuntimeOptions) (_result *DescribeColumnarVersionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeColumnarVersionList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeColumnarVersionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the property information of a specified component, including property names and types.
//
// @param request - DescribeComponentPropetiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComponentPropetiesResponse
func (client *Client) DescribeComponentPropetiesWithContext(ctx context.Context, request *DescribeComponentPropetiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeComponentPropetiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.ComponentName) {
		query["ComponentName"] = request.ComponentName
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
		Action:      dara.String("DescribeComponentPropeties"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComponentPropetiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of custom endpoints defined by the user.
//
// Description:
//
// Queries the list of custom endpoints configured by the user. You can use this operation to manage and view the settings of private connections or VPC endpoint services.
//
// @param request - DescribeCustomEndpointListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomEndpointListResponse
func (client *Client) DescribeCustomEndpointListWithContext(ctx context.Context, request *DescribeCustomEndpointListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomEndpointListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckDeleteCN) {
		query["CheckDeleteCN"] = request.CheckDeleteCN
	}

	if !dara.IsNil(request.CustomEndpointIds) {
		query["CustomEndpointIds"] = request.CustomEndpointIds
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomEndpointList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomEndpointListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeDBInstanceAttribute operation to retrieve instance attributes.
//
// @param request - DescribeDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceAttributeResponse
func (client *Client) DescribeDBInstanceAttributeWithContext(ctx context.Context, request *DescribeDBInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceAttribute"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeDBInstanceConfig operation to retrieve the configuration parameters of an instance.
//
// @param request - DescribeDBInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceConfigResponse
func (client *Client) DescribeDBInstanceConfigWithContext(ctx context.Context, request *DescribeDBInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceConfig"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询endpoint列表
//
// Description:
//
// 该接口用于获取用户已配置的自定义终端节点（Endpoint）列表，便于管理和查看私有连接或VPC终端服务的设置。
//
// @param request - DescribeDBInstanceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceEndpointResponse
func (client *Client) DescribeDBInstanceEndpointWithContext(ctx context.Context, request *DescribeDBInstanceEndpointRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceEndpoint"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the high availability (HA) information of an instance.
//
// @param request - DescribeDBInstanceHARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceHAResponse
func (client *Client) DescribeDBInstanceHAWithContext(ctx context.Context, request *DescribeDBInstanceHARequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceHAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceHA"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceHAResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Views SSL information.
//
// @param request - DescribeDBInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceSSLResponse
func (client *Client) DescribeDBInstanceSSLWithContext(ctx context.Context, request *DescribeDBInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceSSLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceSSL"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeDBInstanceTDE operation to retrieve the details of Transparent Data Encryption (TDE) for an instance.
//
// @param request - DescribeDBInstanceTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceTDEResponse
func (client *Client) DescribeDBInstanceTDEWithContext(ctx context.Context, request *DescribeDBInstanceTDERequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceTDEResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceTDE"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceTDEResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeDBInstanceTopology operation to retrieve the topology information of an instance.
//
// @param request - DescribeDBInstanceTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceTopologyResponse
func (client *Client) DescribeDBInstanceTopologyWithContext(ctx context.Context, request *DescribeDBInstanceTopologyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceTopologyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MinuteSimple) {
		query["MinuteSimple"] = request.MinuteSimple
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
		Action:      dara.String("DescribeDBInstanceTopology"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceTopologyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the basic information about an instance by using the endpoint of the instance.
//
// @param request - DescribeDBInstanceViaEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceViaEndpointResponse
func (client *Client) DescribeDBInstanceViaEndpointWithContext(ctx context.Context, request *DescribeDBInstanceViaEndpointRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceViaEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Endpoint) {
		query["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceViaEndpoint"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceViaEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeDBInstances operation to query a list of instances.
//
// @param request - DescribeDBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesResponse
func (client *Client) DescribeDBInstancesWithContext(ctx context.Context, request *DescribeDBInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbVersion) {
		query["DbVersion"] = request.DbVersion
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MustHasCdc) {
		query["MustHasCdc"] = request.MustHasCdc
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Series) {
		query["Series"] = request.Series
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstances"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the performance data of an instance by calling the DescribeDBNodePerformance operation.
//
// Description:
//
// Note:
//
// - The **endpoint*	- differs from other API operations. Use **polardbx.aliyuncs.com*	- for Chinese regions and Singapore. For other regions, use **polardbx.{region id}.aliyunc.com**.
//
// - When debugging this API operation, if a service not active error is returned, confirm that the **endpoint*	- is correct. You can switch the **service address*	- to **Dubai*	- or **India*	- and change the **endpoint*	- to **polardbx.aliyuncs.com**.
//
// @param request - DescribeDBNodePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBNodePerformanceResponse
func (client *Client) DescribeDBNodePerformanceWithContext(ctx context.Context, request *DescribeDBNodePerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBNodePerformanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DBNodeIds) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !dara.IsNil(request.DBNodeRole) {
		query["DBNodeRole"] = request.DBNodeRole
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
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
		Action:      dara.String("DescribeDBNodePerformance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution details of a data import task.
//
// @param request - DescribeDataImportTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataImportTaskInfoResponse
func (client *Client) DescribeDataImportTaskInfoWithContext(ctx context.Context, request *DescribeDataImportTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataImportTaskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FailPageNumber) {
		query["FailPageNumber"] = request.FailPageNumber
	}

	if !dara.IsNil(request.FailPageSize) {
		query["FailPageSize"] = request.FailPageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	if !dara.IsNil(request.SuccessPageNumber) {
		query["SuccessPageNumber"] = request.SuccessPageNumber
	}

	if !dara.IsNil(request.SuccessPageSize) {
		query["SuccessPageSize"] = request.SuccessPageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataImportTaskInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataImportTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of databases by calling the DescribeDbList operation.
//
// @param request - DescribeDbListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDbListResponse
func (client *Client) DescribeDbListWithContext(ctx context.Context, request *DescribeDbListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDbListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDbList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDbListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeDistributeTableList operation to retrieve the list of database tables.
//
// @param request - DescribeDistributeTableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDistributeTableListResponse
func (client *Client) DescribeDistributeTableListWithContext(ctx context.Context, request *DescribeDistributeTableListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDistributeTableListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDistributeTableList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDistributeTableListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the currently enabled cross-zone configurations.
//
// @param request - DescribeEnabledCrossRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEnabledCrossRegionsResponse
func (client *Client) DescribeEnabledCrossRegionsWithContext(ctx context.Context, request *DescribeEnabledCrossRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEnabledCrossRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEnabledCrossRegions"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEnabledCrossRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of PolarDB-X assessment import tasks. (Single).
//
// @param request - DescribeEvaluateAndImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEvaluateAndImportTaskResponse
func (client *Client) DescribeEvaluateAndImportTaskWithContext(ctx context.Context, request *DescribeEvaluateAndImportTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeEvaluateAndImportTaskResponse, _err error) {
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

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEvaluateAndImportTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEvaluateAndImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of assessment import tasks for PolarDB-X.
//
// Description:
//
// Creates a data import task. You can use this operation to import SQL or CSV files stored in OSS or ECS, or directly provided files, into a destination database instance. By specifying the instance ID, database name, engine type, data source (such as an OSS path), and import type, the system performs data write operations asynchronously or synchronously. This operation is applicable to scenarios such as data migration, initialization, and data backfill. A task ID is returned for subsequent status queries and management.
//
// @param request - DescribeEvaluateAndImportTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEvaluateAndImportTasksResponse
func (client *Client) DescribeEvaluateAndImportTasksWithContext(ctx context.Context, request *DescribeEvaluateAndImportTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeEvaluateAndImportTasksResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEvaluateAndImportTasks"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEvaluateAndImportTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves historical events.
//
// @param request - DescribeEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventsResponse
func (client *Client) DescribeEventsWithContext(ctx context.Context, request *DescribeEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeEventsResponse, _err error) {
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
		Action:      dara.String("DescribeEvents"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of global database network (GDN) instances.
//
// @param request - DescribeGdnInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGdnInstancesResponse
func (client *Client) DescribeGdnInstancesWithContext(ctx context.Context, request *DescribeGdnInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGdnInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterType) {
		query["FilterType"] = request.FilterType
	}

	if !dara.IsNil(request.FilterValue) {
		query["FilterValue"] = request.FilterValue
	}

	if !dara.IsNil(request.GDNId) {
		query["GDNId"] = request.GDNId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGdnInstances"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGdnInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View Memory Engine Information
//
// Description:
//
// > 	- The SQL audit and analysis feature of PolarDB-X 2.0 is free of charge. However, Log Service charges fees for storage space, read traffic, the number of requests, data transformation, and data shipping. For more information about the SQL audit feature, see [Enable SQL audit and analysis](https://help.aliyun.com/document_detail/184619.html).
//
// @param request - DescribeMem0InfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMem0InfoResponse
func (client *Client) DescribeMem0InfoWithContext(ctx context.Context, request *DescribeMem0InfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeMem0InfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMem0Info"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMem0InfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the whitelist of a memory engine.
//
// Description:
//
// > 	- The SQL audit and analysis feature of PolarDB-X 2.0 is free of charge. However, Log Service charges fees for storage space, read traffic, number of requests, data processing, and data shipping. For more information about the SQL audit feature, see [Enable SQL Audit and Analysis](https://help.aliyun.com/document_detail/184619.html).
//
// @param request - DescribeMem0SecurityIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMem0SecurityIpsResponse
func (client *Client) DescribeMem0SecurityIpsWithContext(ctx context.Context, request *DescribeMem0SecurityIpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMem0SecurityIpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMem0SecurityIps"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMem0SecurityIpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exposes instance topology information and download links for commercial backup sets to support offline restoration of cloud instances.
//
// @param request - DescribeOpenBackupSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpenBackupSetResponse
func (client *Client) DescribeOpenBackupSetWithContext(ctx context.Context, request *DescribeOpenBackupSetRequest, runtime *dara.RuntimeOptions) (_result *DescribeOpenBackupSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOpenBackupSet"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOpenBackupSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of database parameter templates created by users or supported by the system, including basic information, associated engine types, and modification times of each parameter template.
//
// @param request - DescribeParameterGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterGroupsResponse
func (client *Client) DescribeParameterGroupsWithContext(ctx context.Context, request *DescribeParameterGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterGroupsResponse, _err error) {
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
		Action:      dara.String("DescribeParameterGroups"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeParameterTemplates operation to retrieve the parameter template list for an instance.
//
// @param request - DescribeParameterTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterTemplatesResponse
func (client *Client) DescribeParameterTemplatesWithContext(ctx context.Context, request *DescribeParameterTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.ParamLevel) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterTemplates"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeParameters operation to retrieve instance parameters.
//
// @param request - DescribeParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParametersResponse
func (client *Client) DescribeParametersWithContext(ctx context.Context, request *DescribeParametersRequest, runtime *dara.RuntimeOptions) (_result *DescribeParametersResponse, _err error) {
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

	if !dara.IsNil(request.ParamLevel) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameters"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParametersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of all data nodes (Data Nodes) in a PolarDB-X instance, including node status, storage usage, and network addresses, to facilitate operations management and monitoring.
//
// @param request - DescribePolarxDataNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePolarxDataNodesResponse
func (client *Client) DescribePolarxDataNodesWithContext(ctx context.Context, request *DescribePolarxDataNodesRequest, runtime *dara.RuntimeOptions) (_result *DescribePolarxDataNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePolarxDataNodes"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePolarxDataNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of VPCs available for PolarDB-X.
//
// Description:
//
// Queries the list of Virtual Private Clouds (VPCs) available under your account for database instances. You can use this operation to select an appropriate network environment when creating or managing database instances.
//
// @param request - DescribeRdsVpcsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsVpcsResponse
func (client *Client) DescribeRdsVpcsWithContext(ctx context.Context, request *DescribeRdsVpcsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsVpcsResponse, _err error) {
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

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdsVpcs"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of available vSwitches in a virtual private cloud (VPC) for an ApsaraDB RDS instance.
//
// Description:
//
// Queries the list of available virtual private clouds (VPCs) under your account for an instance, so that you can select an appropriate network environment when creating or managing a database instance.
//
// @param request - DescribeRdsVswitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsVswitchesResponse
func (client *Client) DescribeRdsVswitchesWithContext(ctx context.Context, request *DescribeRdsVswitchesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsVswitchesResponse, _err error) {
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdsVswitches"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsVswitchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a replication lag inspection task for an instance.
//
// Description:
//
// During the data synchronization phase, proactively initiates a diagnostic task for the replication task to check for exceptions such as latency, interruption, or data inconsistency.
//
// @param request - DescribeRplInspectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRplInspectionTaskResponse
func (client *Client) DescribeRplInspectionTaskWithContext(ctx context.Context, request *DescribeRplInspectionTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeRplInspectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FailPageNumber) {
		query["FailPageNumber"] = request.FailPageNumber
	}

	if !dara.IsNil(request.FailPageSize) {
		query["FailPageSize"] = request.FailPageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	if !dara.IsNil(request.SuccessPageNumber) {
		query["SuccessPageNumber"] = request.SuccessPageNumber
	}

	if !dara.IsNil(request.SuccessPageSize) {
		query["SuccessPageSize"] = request.SuccessPageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRplInspectionTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRplInspectionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Display the ScaleOut migration task progress.
//
// @param request - DescribeScaleOutMigrateTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScaleOutMigrateTaskListResponse
func (client *Client) DescribeScaleOutMigrateTaskListWithContext(ctx context.Context, request *DescribeScaleOutMigrateTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeScaleOutMigrateTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScaleOutMigrateTaskList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScaleOutMigrateTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeSecurityIps operation to view the IP whitelist of an instance.
//
// @param request - DescribeSecurityIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityIpsResponse
func (client *Client) DescribeSecurityIpsWithContext(ctx context.Context, request *DescribeSecurityIpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityIpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityIps"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityIpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage usage details of an instance, including the total capacity, used space, remaining space, and other information.
//
// @param request - DescribeShowStorageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeShowStorageInfoResponse
func (client *Client) DescribeShowStorageInfoWithContext(ctx context.Context, request *DescribeShowStorageInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeShowStorageInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeShowStorageInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeShowStorageInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of slow SQL statements on compute nodes and storage nodes of a PolarDB-X instance.
//
// @param request - DescribeSlowLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogRecordsResponse
func (client *Client) DescribeSlowLogRecordsWithContext(ctx context.Context, request *DescribeSlowLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.DBNodeIds) {
		query["DBNodeIds"] = request.DBNodeIds
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeSlowLogRecords"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries SQL audit information.
//
// Description:
//
// > 	- The SQL audit and analysis feature of PolarDB-X 2.0 is free of charge. However, Simple Log Service charges fees for storage space, read traffic, number of requests, data transformation, and data delivery. For more information about the SQL audit feature, see [Enable SQL audit and analysis](https://help.aliyun.com/document_detail/184619.html).
//
// @param request - DescribeSqlAuditInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlAuditInfoResponse
func (client *Client) DescribeSqlAuditInfoWithContext(ctx context.Context, request *DescribeSqlAuditInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlAuditInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditAccountName) {
		query["AuditAccountName"] = request.AuditAccountName
	}

	if !dara.IsNil(request.AuditAccountPassword) {
		query["AuditAccountPassword"] = request.AuditAccountPassword
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSqlAuditInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSqlAuditInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of SQL flashback tasks by calling the DescribeSqlFlashbackTaskList operation.
//
// @param request - DescribeSqlFlashbackTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSqlFlashbackTaskListResponse
func (client *Client) DescribeSqlFlashbackTaskListWithContext(ctx context.Context, request *DescribeSqlFlashbackTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSqlFlashbackTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PolardbxInstanceId) {
		query["PolardbxInstanceId"] = request.PolardbxInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSqlFlashbackTaskList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSqlFlashbackTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a storage pool, including capacity, usage, and status.
//
// @param request - DescribeStoragePoolInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStoragePoolInfoResponse
func (client *Client) DescribeStoragePoolInfoWithContext(ctx context.Context, request *DescribeStoragePoolInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeStoragePoolInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStoragePoolInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStoragePoolInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a database schema import task.
//
// Description:
//
// The CreateDataImportTask operation creates a data import task. You can use this operation to import SQL or CSV files stored in OSS or ECS, or directly provided, into a destination database instance. Specify the instance ID, database name, engine type, data source (such as an OSS path), and import type. The system performs the data write operation asynchronously or synchronously. This operation is applicable to scenarios such as data migration, initialization, and data backfill. A task ID is returned for subsequent status queries and management.
//
// @param request - DescribeStructureImportTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStructureImportTaskInfoResponse
func (client *Client) DescribeStructureImportTaskInfoWithContext(ctx context.Context, request *DescribeStructureImportTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeStructureImportTaskInfoResponse, _err error) {
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

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStructureImportTaskInfo"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStructureImportTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Supabase API key.
//
// Description:
//
// - Binary log files are retained for 15 days by default.
//
// - The returned log list includes all logs whose log record end time is later than the query start time and whose log record start time is earlier than the query end time.
//
// - If DownloadLink is not NULL, you can use this URL to download the backup file. The URL is valid for 2 days after it is generated. Download the file before the URL expires.
//
// @param request - DescribeSupabaseApiKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupabaseApiKeyResponse
func (client *Client) DescribeSupabaseApiKeyWithContext(ctx context.Context, request *DescribeSupabaseApiKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupabaseApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupabaseApiKey"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupabaseApiKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Supabase instance.
//
// Description:
//
// ***
//
// @param request - DescribeSupabaseInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupabaseInstanceAttributeResponse
func (client *Client) DescribeSupabaseInstanceAttributeWithContext(ctx context.Context, request *DescribeSupabaseInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupabaseInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupabaseInstanceAttribute"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupabaseInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Supabase instances.
//
// Description:
//
// Queries the list of custom endpoints configured by the user for managing and viewing private connection or VPC endpoint service settings.
//
// @param request - DescribeSupabaseInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupabaseInstancesResponse
func (client *Client) DescribeSupabaseInstancesWithContext(ctx context.Context, request *DescribeSupabaseInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupabaseInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupabaseInstances"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupabaseInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP whitelist of a Supabase instance.
//
// Description:
//
// - Binary log files are retained for 15 days by default.
//
// - The returned log list includes all logs whose log record end time is later than the specified query start time and whose log record start time is earlier than the specified query end time.
//
// - If DownloadLink is not NULL, you can use this URL to download the backup file. The URL is valid for 2 days after it is generated. Download the file before the URL expires.
//
// @param request - DescribeSupabaseIpWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupabaseIpWhitelistResponse
func (client *Client) DescribeSupabaseIpWhitelistWithContext(ctx context.Context, request *DescribeSupabaseIpWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupabaseIpWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupabaseIpWhitelist"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupabaseIpWhitelistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves tag information.
//
// @param request - DescribeTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
func (client *Client) DescribeTagsWithContext(ctx context.Context, request *DescribeTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the task list of an instance.
//
// @param request - DescribeTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithContext(ctx context.Context, request *DescribeTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTasks"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current status of a task that upgrades or converts a standard instance to an Enterprise instance.
//
// @param request - DescribeTransformStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransformStatusResponse
func (client *Client) DescribeTransformStatusWithContext(ctx context.Context, request *DescribeTransformStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransformStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.QueryReport) {
		query["QueryReport"] = request.QueryReport
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTransformStatus"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTransformStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the DescribeUserEncryptionKeyList operation to retrieve the list of Transparent Data Encryption (TDE) keys for a user.
//
// @param request - DescribeUserEncryptionKeyListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserEncryptionKeyListResponse
func (client *Client) DescribeUserEncryptionKeyListWithContext(ctx context.Context, request *DescribeUserEncryptionKeyListRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserEncryptionKeyList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables three-role mode.
//
// @param request - DisableRightsSeparationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableRightsSeparationResponse
func (client *Client) DisableRightsSeparationWithContext(ctx context.Context, request *DisableRightsSeparationRequest, runtime *dara.RuntimeOptions) (_result *DisableRightsSeparationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbaAccountName) {
		query["DbaAccountName"] = request.DbaAccountName
	}

	if !dara.IsNil(request.DbaAccountPassword) {
		query["DbaAccountPassword"] = request.DbaAccountPassword
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableRightsSeparation"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableRightsSeparationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the SQL audit feature.
//
// Description:
//
// > After the SQL audit and analysis feature is disabled, data is no longer written to Simple Log Service. However, you can still view historical data, which incurs storage and indexing fees. You can delete all data to deactivate Simple Log Service. After Simple Log Service is deactivated, no further fees are incurred. For more information about deactivating Simple Log Service, see [Deactivate Simple Log Service](https://help.aliyun.com/document_detail/90650.html).
//
// @param request - DisableSqlAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSqlAuditResponse
func (client *Client) DisableSqlAuditWithContext(ctx context.Context, request *DisableSqlAuditRequest, runtime *dara.RuntimeOptions) (_result *DisableSqlAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditAccountName) {
		query["AuditAccountName"] = request.AuditAccountName
	}

	if !dara.IsNil(request.AuditAccountPassword) {
		query["AuditAccountPassword"] = request.AuditAccountPassword
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSqlAudit"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSqlAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables three-role mode.
//
// @param request - EnableRightsSeparationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRightsSeparationResponse
func (client *Client) EnableRightsSeparationWithContext(ctx context.Context, request *EnableRightsSeparationRequest, runtime *dara.RuntimeOptions) (_result *EnableRightsSeparationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditAccountDescription) {
		query["AuditAccountDescription"] = request.AuditAccountDescription
	}

	if !dara.IsNil(request.AuditAccountName) {
		query["AuditAccountName"] = request.AuditAccountName
	}

	if !dara.IsNil(request.AuditAccountPassword) {
		query["AuditAccountPassword"] = request.AuditAccountPassword
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountDescription) {
		query["SecurityAccountDescription"] = request.SecurityAccountDescription
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableRightsSeparation"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableRightsSeparationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the SQL audit feature.
//
// Description:
//
// > 	- The PolarDB-X 2.0 SQL audit and analysis feature itself is free of charge. However, Log Service charges fees for storage space, read traffic, number of requests, data transformation, data shipping, and other services. For more information about the SQL audit feature, see [Enable SQL Audit and Analysis](https://help.aliyun.com/document_detail/184619.html).
//
// @param request - EnableSqlAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSqlAuditResponse
func (client *Client) EnableSqlAuditWithContext(ctx context.Context, request *EnableSqlAuditRequest, runtime *dara.RuntimeOptions) (_result *EnableSqlAuditResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditAccountName) {
		query["AuditAccountName"] = request.AuditAccountName
	}

	if !dara.IsNil(request.AuditAccountPassword) {
		query["AuditAccountPassword"] = request.AuditAccountPassword
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ExpireAfterDays) {
		query["ExpireAfterDays"] = request.ExpireAfterDays
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSqlAudit"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSqlAuditResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of labels.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates an instance from one zone to another.
//
// @param request - MigrateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateDBInstanceResponse
func (client *Client) MigrateDBInstanceWithContext(ctx context.Context, request *MigrateDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *MigrateDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.PrimaryZoneId) {
		query["PrimaryZoneId"] = request.PrimaryZoneId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecondaryZoneId) {
		query["SecondaryZoneId"] = request.SecondaryZoneId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	if !dara.IsNil(request.TertiaryZoneId) {
		query["TertiaryZoneId"] = request.TertiaryZoneId
	}

	if !dara.IsNil(request.TopologyType) {
		query["TopologyType"] = request.TopologyType
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateDBInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyAccountDescription operation to modify the description of an account.
//
// @param request - ModifyAccountDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountDescriptionResponse
func (client *Client) ModifyAccountDescriptionWithContext(ctx context.Context, request *ModifyAccountDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountDescription) {
		query["AccountDescription"] = request.AccountDescription
	}

	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountDescription"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the permissions of a standard account.
//
// @param request - ModifyAccountPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountPrivilegeResponse
func (client *Client) ModifyAccountPrivilegeWithContext(ctx context.Context, request *ModifyAccountPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountPrivilegeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPrivilege) {
		query["AccountPrivilege"] = request.AccountPrivilege
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountPrivilege"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountPrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyActiveOperationMaintainConf operation to modify the time configuration of O&M events.
//
// @param request - ModifyActiveOperationMaintainConfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationMaintainConfResponse
func (client *Client) ModifyActiveOperationMaintainConfWithContext(ctx context.Context, request *ModifyActiveOperationMaintainConfRequest, runtime *dara.RuntimeOptions) (_result *ModifyActiveOperationMaintainConfResponse, _err error) {
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
		Action:      dara.String("ModifyActiveOperationMaintainConf"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyActiveOperationMaintainConfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyActiveOperationTasks operation to modify the execution time of O&M events.
//
// @param request - ModifyActiveOperationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationTasksResponse
func (client *Client) ModifyActiveOperationTasksWithContext(ctx context.Context, request *ModifyActiveOperationTasksRequest, runtime *dara.RuntimeOptions) (_result *ModifyActiveOperationTasksResponse, _err error) {
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

	if !dara.IsNil(request.ImmediateStart) {
		query["ImmediateStart"] = request.ImmediateStart
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationTasks"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyActiveOperationTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the CDC configuration.
//
// Description:
//
// ***.
//
// @param request - ModifyCdcClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCdcClassResponse
func (client *Client) ModifyCdcClassWithContext(ctx context.Context, request *ModifyCdcClassRequest, runtime *dara.RuntimeOptions) (_result *ModifyCdcClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CDCNodeCount) {
		query["CDCNodeCount"] = request.CDCNodeCount
	}

	if !dara.IsNil(request.CdcClass) {
		query["CdcClass"] = request.CdcClass
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCdcClass"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCdcClassResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the column store specifications.
//
// Description:
//
// ***.
//
// @param request - ModifyColumnarClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyColumnarClassResponse
func (client *Client) ModifyColumnarClassWithContext(ctx context.Context, request *ModifyColumnarClassRequest, runtime *dara.RuntimeOptions) (_result *ModifyColumnarClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ColumnarClass) {
		query["ColumnarClass"] = request.ColumnarClass
	}

	if !dara.IsNil(request.ColumnarNodeCount) {
		query["ColumnarNodeCount"] = request.ColumnarNodeCount
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyColumnarClass"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyColumnarClassResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a custom endpoint.
//
// Description:
//
// In PolarDB-X, in addition to the default primary endpoint and read-only endpoint, you can create custom endpoints to implement more flexible read/write splitting, load balancing, or business isolation strategies. The ModifyCustomEndpoint operation allows you to dynamically adjust the attributes of these custom endpoints without restarting the instance.
//
// @param request - ModifyCustomEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomEndpointResponse
func (client *Client) ModifyCustomEndpointWithContext(ctx context.Context, request *ModifyCustomEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomEndpointId) {
		query["CustomEndpointId"] = request.CustomEndpointId
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeAutoEnter) {
		query["NodeAutoEnter"] = request.NodeAutoEnter
	}

	if !dara.IsNil(request.NodeIds) {
		query["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.NodeRole) {
		query["NodeRole"] = request.NodeRole
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomEndpoint"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the network configurations of a custom network endpoint, such as the subnet, security group, and public network access settings.
//
// Description:
//
// Deletes the custom endpoint of a specified database instance and disables access through the domain name.
//
// @param request - ModifyCustomEndpointNetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomEndpointNetResponse
func (client *Client) ModifyCustomEndpointNetWithContext(ctx context.Context, request *ModifyCustomEndpointNetRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomEndpointNetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnPrefix) {
		query["ConnPrefix"] = request.ConnPrefix
	}

	if !dara.IsNil(request.CustomEndpointId) {
		query["CustomEndpointId"] = request.CustomEndpointId
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomEndpointNet"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomEndpointNetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyDBInstanceClass operation to modify the specifications of an instance.
//
// @param request - ModifyDBInstanceClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceClassResponse
func (client *Client) ModifyDBInstanceClassWithContext(ctx context.Context, request *ModifyDBInstanceClassRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceClassResponse, _err error) {
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

	if !dara.IsNil(request.CnClass) {
		query["CnClass"] = request.CnClass
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DnClass) {
		query["DnClass"] = request.DnClass
	}

	if !dara.IsNil(request.DnStorageSpace) {
		query["DnStorageSpace"] = request.DnStorageSpace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SpecifiedDNScale) {
		query["SpecifiedDNScale"] = request.SpecifiedDNScale
	}

	if !dara.IsNil(request.SpecifiedDNSpecMapJson) {
		query["SpecifiedDNSpecMapJson"] = request.SpecifiedDNSpecMapJson
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !dara.IsNil(request.TargetDBInstanceClass) {
		query["TargetDBInstanceClass"] = request.TargetDBInstanceClass
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceClass"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceClassResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyDBInstanceConfig operation to modify instance configuration items.
//
// @param request - ModifyDBInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceConfigResponse
func (client *Client) ModifyDBInstanceConfigWithContext(ctx context.Context, request *ModifyDBInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.ConfigValue) {
		query["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConfig"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the connection string of an instance.
//
// @param request - ModifyDBInstanceConnectionStringRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceConnectionStringResponse
func (client *Client) ModifyDBInstanceConnectionStringWithContext(ctx context.Context, request *ModifyDBInstanceConnectionStringRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.NewPort) {
		query["NewPort"] = request.NewPort
	}

	if !dara.IsNil(request.NewPrefix) {
		query["NewPrefix"] = request.NewPrefix
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConnectionString"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyDBInstanceDescription operation to modify the description of an instance.
//
// @param request - ModifyDBInstanceDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceDescriptionResponse
func (client *Client) ModifyDBInstanceDescriptionWithContext(ctx context.Context, request *ModifyDBInstanceDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceDescription"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window of an instance.
//
// Description:
//
// ***.
//
// @param request - ModifyDBInstanceMaintainTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceMaintainTimeResponse
func (client *Client) ModifyDBInstanceMaintainTimeWithContext(ctx context.Context, request *ModifyDBInstanceMaintainTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.MaintainTime) {
		query["MaintainTime"] = request.MaintainTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceMaintainTime"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the virtual IP (VIP) address or network properties bound to a database instance. This operation supports changing the internal IP address, adjusting the vSwitch, migrating across zones, and other operations.
//
// Description:
//
// ***.
//
// @param request - ModifyDBInstanceVipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceVipResponse
func (client *Client) ModifyDBInstanceVipWithContext(ctx context.Context, request *ModifyDBInstanceVipRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceVipResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.InstanceClusterName) {
		query["InstanceClusterName"] = request.InstanceClusterName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceVip"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceVipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyDatabaseDescription operation to modify the description of a database.
//
// @param request - ModifyDatabaseDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseDescriptionResponse
func (client *Client) ModifyDatabaseDescriptionWithContext(ctx context.Context, request *ModifyDatabaseDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DbDescription) {
		query["DbDescription"] = request.DbDescription
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDatabaseDescription"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatabaseDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration or migration parameters of a database engine migration task, such as the source database, destination database, migration objects, or migration mode.
//
// Description:
//
// ***.
//
// @param request - ModifyEngineMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEngineMigrationResponse
func (client *Client) ModifyEngineMigrationWithContext(ctx context.Context, request *ModifyEngineMigrationRequest, runtime *dara.RuntimeOptions) (_result *ModifyEngineMigrationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStrings) {
		query["ConnectionStrings"] = request.ConnectionStrings
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.NewMasterDBInstanceName) {
		query["NewMasterDBInstanceName"] = request.NewMasterDBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceDBInstanceName) {
		query["SourceDBInstanceName"] = request.SourceDBInstanceName
	}

	if !dara.IsNil(request.SwapConnectionString) {
		query["SwapConnectionString"] = request.SwapConnectionString
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEngineMigration"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEngineMigrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the whitelist of a memory engine.
//
// Description:
//
// ***
//
// @param request - ModifyMem0SecurityIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMem0SecurityIpsResponse
func (client *Client) ModifyMem0SecurityIpsWithContext(ctx context.Context, request *ModifyMem0SecurityIpsRequest, runtime *dara.RuntimeOptions) (_result *ModifyMem0SecurityIpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMem0SecurityIps"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMem0SecurityIpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifyParameter operation to modify instance parameters, including compute layer and storage layer parameters.
//
// @param request - ModifyParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParameterResponse
func (client *Client) ModifyParameterWithContext(ctx context.Context, request *ModifyParameterRequest, runtime *dara.RuntimeOptions) (_result *ModifyParameterResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ParamLevel) {
		query["ParamLevel"] = request.ParamLevel
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyParameter"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyParameterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ModifySecurityIps operation to modify the whitelist of an instance.
//
// @param request - ModifySecurityIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityIpsResponse
func (client *Client) ModifySecurityIpsWithContext(ctx context.Context, request *ModifySecurityIpsRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityIpsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityIps"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the Supabase Dashboard password.
//
// Description:
//
// ***
//
// @param request - ModifySupabaseDashboardPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySupabaseDashboardPasswordResponse
func (client *Client) ModifySupabaseDashboardPasswordWithContext(ctx context.Context, request *ModifySupabaseDashboardPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifySupabaseDashboardPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.NewPassword) {
		query["NewPassword"] = request.NewPassword
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySupabaseDashboardPassword"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySupabaseDashboardPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP whitelist of a Supabase instance.
//
// Description:
//
// ***
//
// @param request - ModifySupabaseSecurityIPListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySupabaseSecurityIPListResponse
func (client *Client) ModifySupabaseSecurityIPListWithContext(ctx context.Context, request *ModifySupabaseSecurityIPListRequest, runtime *dara.RuntimeOptions) (_result *ModifySupabaseSecurityIPListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySupabaseSecurityIPList"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySupabaseSecurityIPListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a pre-check and feasibility assessment for a recovery task before you execute SQL flashback recovery.
//
// @param request - PreCheckSqlFlashbackTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreCheckSqlFlashbackTaskResponse
func (client *Client) PreCheckSqlFlashbackTaskWithContext(ctx context.Context, request *PreCheckSqlFlashbackTaskRequest, runtime *dara.RuntimeOptions) (_result *PreCheckSqlFlashbackTaskResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PolardbxInstanceId) {
		query["PolardbxInstanceId"] = request.PolardbxInstanceId
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
		Action:      dara.String("PreCheckSqlFlashbackTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreCheckSqlFlashbackTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes the metadata of an import task.
//
// @param request - RefreshImportMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshImportMetaResponse
func (client *Client) RefreshImportMetaWithContext(ctx context.Context, request *RefreshImportMetaRequest, runtime *dara.RuntimeOptions) (_result *RefreshImportMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshImportMeta"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshImportMetaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the ReleaseColdDataVolume operation.
//
// @param request - ReleaseColdDataVolumeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseColdDataVolumeResponse
func (client *Client) ReleaseColdDataVolumeWithContext(ctx context.Context, request *ReleaseColdDataVolumeRequest, runtime *dara.RuntimeOptions) (_result *ReleaseColdDataVolumeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseColdDataVolume"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseColdDataVolumeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public network connection of an instance by calling the ReleaseInstancePublicConnection operation.
//
// @param request - ReleaseInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstancePublicConnectionResponse
func (client *Client) ReleaseInstancePublicConnectionWithContext(ctx context.Context, request *ReleaseInstancePublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentConnectionString) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstancePublicConnection"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public domain name for a Mem0 instance.
//
// Description:
//
// This operation is used to verify that no active connections exist before a rollback task to ensure operational safety.
//
// @param request - ReleaseMem0PublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseMem0PublicConnectionResponse
func (client *Client) ReleaseMem0PublicConnectionWithContext(ctx context.Context, request *ReleaseMem0PublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *ReleaseMem0PublicConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentConnectionString) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseMem0PublicConnection"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseMem0PublicConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of an account.
//
// @param request - ResetAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordResponse
func (client *Client) ResetAccountPasswordWithContext(ctx context.Context, request *ResetAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountPassword"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of an account.
//
// Description:
//
// ***.
//
// @param request - ResetAccountPasswordRestrictRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordRestrictResponse
func (client *Client) ResetAccountPasswordRestrictWithContext(ctx context.Context, request *ResetAccountPasswordRestrictRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountPasswordRestrictResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityAccountName) {
		query["SecurityAccountName"] = request.SecurityAccountName
	}

	if !dara.IsNil(request.SecurityAccountPassword) {
		query["SecurityAccountPassword"] = request.SecurityAccountPassword
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountPasswordRestrict"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAccountPasswordRestrictResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the API key of the memory engine.
//
// Description:
//
// ***.
//
// @param request - ResetMem0AccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetMem0AccountPasswordResponse
func (client *Client) ResetMem0AccountPasswordWithContext(ctx context.Context, request *ResetMem0AccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetMem0AccountPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.Mem0ApiKey) {
		query["Mem0ApiKey"] = request.Mem0ApiKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetMem0AccountPassword"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetMem0AccountPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an instance by calling the RestartDBInstance operation.
//
// @param request - RestartDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBInstanceResponse
func (client *Client) RestartDBInstanceWithContext(ctx context.Context, request *RestartDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDBInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a data import task.
//
// @param request - RestartDataImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDataImportTaskResponse
func (client *Client) RestartDataImportTaskWithContext(ctx context.Context, request *RestartDataImportTaskRequest, runtime *dara.RuntimeOptions) (_result *RestartDataImportTaskResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDataImportTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDataImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a Supabase instance.
//
// Description:
//
// ***
//
// @param request - RestartSupabaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartSupabaseInstanceResponse
func (client *Client) RestartSupabaseInstanceWithContext(ctx context.Context, request *RestartSupabaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartSupabaseInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartSupabaseInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartSupabaseInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 克隆PolarDB-X实例
//
// Description:
//
// ***
//
// @param request - RestoreDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreDBInstanceResponse
func (client *Client) RestoreDBInstanceWithContext(ctx context.Context, request *RestoreDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestoreDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BackupSetId) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !dara.IsNil(request.BackupSetRegion) {
		query["BackupSetRegion"] = request.BackupSetRegion
	}

	if !dara.IsNil(request.CNNodeCount) {
		query["CNNodeCount"] = request.CNNodeCount
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CloneInstanceName) {
		query["CloneInstanceName"] = request.CloneInstanceName
	}

	if !dara.IsNil(request.CnClass) {
		query["CnClass"] = request.CnClass
	}

	if !dara.IsNil(request.DBNodeClass) {
		query["DBNodeClass"] = request.DBNodeClass
	}

	if !dara.IsNil(request.DBNodeCount) {
		query["DBNodeCount"] = request.DBNodeCount
	}

	if !dara.IsNil(request.DNNodeCount) {
		query["DNNodeCount"] = request.DNNodeCount
	}

	if !dara.IsNil(request.DnClass) {
		query["DnClass"] = request.DnClass
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.GdnRole) {
		query["GdnRole"] = request.GdnRole
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrimaryZone) {
		query["PrimaryZone"] = request.PrimaryZone
	}

	if !dara.IsNil(request.RecoveryTypeCode) {
		query["RecoveryTypeCode"] = request.RecoveryTypeCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.SecondaryZone) {
		query["SecondaryZone"] = request.SecondaryZone
	}

	if !dara.IsNil(request.Series) {
		query["Series"] = request.Series
	}

	if !dara.IsNil(request.SourceInstanceRegion) {
		query["SourceInstanceRegion"] = request.SourceInstanceRegion
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.TertiaryZone) {
		query["TertiaryZone"] = request.TertiaryZone
	}

	if !dara.IsNil(request.TopologyType) {
		query["TopologyType"] = request.TopologyType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreDBInstance"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is used to skip the current step.
//
// @param request - SkipCurrentStepRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SkipCurrentStepResponse
func (client *Client) SkipCurrentStepWithContext(ctx context.Context, request *SkipCurrentStepRequest, runtime *dara.RuntimeOptions) (_result *SkipCurrentStepResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentStep) {
		query["CurrentStep"] = request.CurrentStep
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SkipCurrentStep"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SkipCurrentStepResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a database cutover for a database migration or synchronization task.
//
// @param request - StartSwitchDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSwitchDatabaseResponse
func (client *Client) StartSwitchDatabaseWithContext(ctx context.Context, request *StartSwitchDatabaseRequest, runtime *dara.RuntimeOptions) (_result *StartSwitchDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DstMainConnectString) {
		query["DstMainConnectString"] = request.DstMainConnectString
	}

	if !dara.IsNil(request.DstMainPort) {
		query["DstMainPort"] = request.DstMainPort
	}

	if !dara.IsNil(request.IsModifyEndpoint) {
		query["IsModifyEndpoint"] = request.IsModifyEndpoint
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	if !dara.IsNil(request.SrcMainConnectString) {
		query["SrcMainConnectString"] = request.SrcMainConnectString
	}

	if !dara.IsNil(request.SrcMainPort) {
		query["SrcMainPort"] = request.SrcMainPort
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartSwitchDatabase"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartSwitchDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a data import task.
//
// @param request - StopDataImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDataImportTaskResponse
func (client *Client) StopDataImportTaskWithContext(ctx context.Context, request *StopDataImportTaskRequest, runtime *dara.RuntimeOptions) (_result *StopDataImportTaskResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SlinkTaskId) {
		query["SlinkTaskId"] = request.SlinkTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDataImportTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDataImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a SQL flashback task by calling the SubmitSqlFlashbackTask operation.
//
// @param request - SubmitSqlFlashbackTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSqlFlashbackTaskResponse
func (client *Client) SubmitSqlFlashbackTaskWithContext(ctx context.Context, request *SubmitSqlFlashbackTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitSqlFlashbackTaskResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PolardbxInstanceId) {
		query["PolardbxInstanceId"] = request.PolardbxInstanceId
	}

	if !dara.IsNil(request.RecallRestoreType) {
		query["RecallRestoreType"] = request.RecallRestoreType
	}

	if !dara.IsNil(request.RecallType) {
		query["RecallType"] = request.RecallType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SqlPk) {
		query["SqlPk"] = request.SqlPk
	}

	if !dara.IsNil(request.SqlType) {
		query["SqlType"] = request.SqlType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TraceId) {
		query["TraceId"] = request.TraceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSqlFlashbackTask"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSqlFlashbackTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SwitchDBInstanceHARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchDBInstanceHAResponse
func (client *Client) SwitchDBInstanceHAWithContext(ctx context.Context, request *SwitchDBInstanceHARequest, runtime *dara.RuntimeOptions) (_result *SwitchDBInstanceHAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !dara.IsNil(request.TargetPrimaryAzoneId) {
		query["TargetPrimaryAzoneId"] = request.TargetPrimaryAzoneId
	}

	if !dara.IsNil(request.TargetPrimaryRegionId) {
		query["TargetPrimaryRegionId"] = request.TargetPrimaryRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchDBInstanceHA"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchDBInstanceHAResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a primary/secondary switchover for a global database network (GDN).
//
// @param request - SwitchGdnMemberRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchGdnMemberRoleResponse
func (client *Client) SwitchGdnMemberRoleWithContext(ctx context.Context, request *SwitchGdnMemberRoleRequest, runtime *dara.RuntimeOptions) (_result *SwitchGdnMemberRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DstMainConnectString) {
		query["DstMainConnectString"] = request.DstMainConnectString
	}

	if !dara.IsNil(request.DstMainPort) {
		query["DstMainPort"] = request.DstMainPort
	}

	if !dara.IsNil(request.IsModifyEndpoint) {
		query["IsModifyEndpoint"] = request.IsModifyEndpoint
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SrcMainConnectString) {
		query["SrcMainConnectString"] = request.SrcMainConnectString
	}

	if !dara.IsNil(request.SrcMainPort) {
		query["SrcMainPort"] = request.SrcMainPort
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	if !dara.IsNil(request.TaskTimeout) {
		query["TaskTimeout"] = request.TaskTimeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchGdnMemberRole"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchGdnMemberRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the TagResources operation to add tags to resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from a resource.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the UpdateBackupPolicy operation to modify the backup policy of an instance.
//
// @param request - UpdateBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBackupPolicyResponse
func (client *Client) UpdateBackupPolicyWithContext(ctx context.Context, request *UpdateBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupPeriod) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !dara.IsNil(request.BackupPlanBegin) {
		query["BackupPlanBegin"] = request.BackupPlanBegin
	}

	if !dara.IsNil(request.BackupSetRetention) {
		query["BackupSetRetention"] = request.BackupSetRetention
	}

	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.BackupWay) {
		query["BackupWay"] = request.BackupWay
	}

	if !dara.IsNil(request.ColdDataBackupInterval) {
		query["ColdDataBackupInterval"] = request.ColdDataBackupInterval
	}

	if !dara.IsNil(request.ColdDataBackupRetention) {
		query["ColdDataBackupRetention"] = request.ColdDataBackupRetention
	}

	if !dara.IsNil(request.CrossRegionDataBackupRetention) {
		query["CrossRegionDataBackupRetention"] = request.CrossRegionDataBackupRetention
	}

	if !dara.IsNil(request.CrossRegionFilterValue) {
		query["CrossRegionFilterValue"] = request.CrossRegionFilterValue
	}

	if !dara.IsNil(request.CrossRegionLogBackupRetention) {
		query["CrossRegionLogBackupRetention"] = request.CrossRegionLogBackupRetention
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DestCrossRegion) {
		query["DestCrossRegion"] = request.DestCrossRegion
	}

	if !dara.IsNil(request.ForceCleanOnHighSpaceUsage) {
		query["ForceCleanOnHighSpaceUsage"] = request.ForceCleanOnHighSpaceUsage
	}

	if !dara.IsNil(request.IsCrossRegionDataBackupEnabled) {
		query["IsCrossRegionDataBackupEnabled"] = request.IsCrossRegionDataBackupEnabled
	}

	if !dara.IsNil(request.IsCrossRegionLogBackupEnabled) {
		query["IsCrossRegionLogBackupEnabled"] = request.IsCrossRegionLogBackupEnabled
	}

	if !dara.IsNil(request.IsEnabled) {
		query["IsEnabled"] = request.IsEnabled
	}

	if !dara.IsNil(request.LocalLogRetention) {
		query["LocalLogRetention"] = request.LocalLogRetention
	}

	if !dara.IsNil(request.LocalLogRetentionNumber) {
		query["LocalLogRetentionNumber"] = request.LocalLogRetentionNumber
	}

	if !dara.IsNil(request.LogLocalRetentionSpace) {
		query["LogLocalRetentionSpace"] = request.LogLocalRetentionSpace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoveLogRetention) {
		query["RemoveLogRetention"] = request.RemoveLogRetention
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBackupPolicy"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBackupPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例的管控参数
//
// Description:
//
// <props="china">更多关于实例账号的信息，请参见[账号管理](https://help.aliyun.com/document_detail/172163.html)。
//
// @param request - UpdateCustinsParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustinsParamResponse
func (client *Client) UpdateCustinsParamWithContext(ctx context.Context, request *UpdateCustinsParamRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustinsParamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustinsParam"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustinsParamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the SSL configuration of an instance.
//
// @param request - UpdateDBInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDBInstanceSSLResponse
func (client *Client) UpdateDBInstanceSSLWithContext(ctx context.Context, request *UpdateDBInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *UpdateDBInstanceSSLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertCommonName) {
		query["CertCommonName"] = request.CertCommonName
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.EnableSSL) {
		query["EnableSSL"] = request.EnableSSL
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDBInstanceSSL"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDBInstanceSSLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables Transparent Data Encryption (TDE) for an instance by calling the UpdateDBInstanceTDE operation.
//
// @param request - UpdateDBInstanceTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDBInstanceTDEResponse
func (client *Client) UpdateDBInstanceTDEWithContext(ctx context.Context, request *UpdateDBInstanceTDERequest, runtime *dara.RuntimeOptions) (_result *UpdateDBInstanceTDEResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.TDEStatus) {
		query["TDEStatus"] = request.TDEStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDBInstanceTDE"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDBInstanceTDEResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the UpdatePolarDBXInstanceNode operation to change the number of nodes for an instance, including scaling out and scaling in. This request generates a billable order.
//
// @param request - UpdatePolarDBXInstanceNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolarDBXInstanceNodeResponse
func (client *Client) UpdatePolarDBXInstanceNodeWithContext(ctx context.Context, request *UpdatePolarDBXInstanceNodeRequest, runtime *dara.RuntimeOptions) (_result *UpdatePolarDBXInstanceNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddDNSpec) {
		query["AddDNSpec"] = request.AddDNSpec
	}

	if !dara.IsNil(request.CNNodeCount) {
		query["CNNodeCount"] = request.CNNodeCount
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DNNodeCount) {
		query["DNNodeCount"] = request.DNNodeCount
	}

	if !dara.IsNil(request.DbInstanceNodeCount) {
		query["DbInstanceNodeCount"] = request.DbInstanceNodeCount
	}

	if !dara.IsNil(request.DeleteDNIds) {
		query["DeleteDNIds"] = request.DeleteDNIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StoragePoolName) {
		query["StoragePoolName"] = request.StoragePoolName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolarDBXInstanceNode"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolarDBXInstanceNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the CDC node version of a PolarDB-X instance.
//
// @param request - UpgradeCDCVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeCDCVersionResponse
func (client *Client) UpgradeCDCVersionWithContext(ctx context.Context, request *UpgradeCDCVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeCDCVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdcDbVersion) {
		query["CdcDbVersion"] = request.CdcDbVersion
	}

	if !dara.IsNil(request.CdcMinorVersion) {
		query["CdcMinorVersion"] = request.CdcMinorVersion
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeCDCVersion"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeCDCVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the column store version.
//
// @param request - UpgradeColumnarVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeColumnarVersionResponse
func (client *Client) UpgradeColumnarVersionWithContext(ctx context.Context, request *UpgradeColumnarVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeColumnarVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ColumnarVersion) {
		query["ColumnarVersion"] = request.ColumnarVersion
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeColumnarVersion"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeColumnarVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Invokes the UpgradeDBInstanceKernelVersion operation to update the kernel version of an instance.
//
// @param request - UpgradeDBInstanceKernelVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceKernelVersionResponse
func (client *Client) UpgradeDBInstanceKernelVersionWithContext(ctx context.Context, request *UpgradeDBInstanceKernelVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBInstanceKernelVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.MinorVersion) {
		query["MinorVersion"] = request.MinorVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceKernelVersion"),
		Version:     dara.String("2020-02-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBInstanceKernelVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
