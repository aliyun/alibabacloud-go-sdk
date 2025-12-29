// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// @param request - AllocateDBInstanceSrvNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateDBInstanceSrvNetworkAddressResponse
func (client *Client) AllocateDBInstanceSrvNetworkAddressWithContext(ctx context.Context, request *AllocateDBInstanceSrvNetworkAddressRequest, runtime *dara.RuntimeOptions) (_result *AllocateDBInstanceSrvNetworkAddressResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SrvConnectionType) {
		query["SrvConnectionType"] = request.SrvConnectionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateDBInstanceSrvNetworkAddress"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateDBInstanceSrvNetworkAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for an internal endpoint for a shard or Configserver node in an ApsaraDB for MongoDB sharded cluster instance.
//
// Description:
//
// This operation is applicable only to sharded cluster instances. For more information, see [Apply for an endpoint for a shard or Configserver node](https://help.aliyun.com/document_detail/134037.html).
//
// >  The allocated endpoints can be used only for internal access. To gain Internet access, you must call the [AllocatePublicNetworkAddress](https://help.aliyun.com/document_detail/67602.html) operation to apply for public endpoints.
//
// @param request - AllocateNodePrivateNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateNodePrivateNetworkAddressResponse
func (client *Client) AllocateNodePrivateNetworkAddressWithContext(ctx context.Context, request *AllocateNodePrivateNetworkAddressRequest, runtime *dara.RuntimeOptions) (_result *AllocateNodePrivateNetworkAddressResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateNodePrivateNetworkAddress"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateNodePrivateNetworkAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Allocates a public endpoint to an instance.
//
// @param request - AllocatePublicNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocatePublicNetworkAddressResponse
func (client *Client) AllocatePublicNetworkAddressWithContext(ctx context.Context, request *AllocatePublicNetworkAddressRequest, runtime *dara.RuntimeOptions) (_result *AllocatePublicNetworkAddressResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("AllocatePublicNetworkAddress"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocatePublicNetworkAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ids) {
		query["Ids"] = request.Ids
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("CancelActiveOperationTasks"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// You can call this operation to check whether KMS keys are authorized to ApsaraDB for MongoDB instances.
//
// Description:
//
// Before you enable Transparent Data Encryption (TDE) by calling the [ModifyDBInstanceTDE](https://help.aliyun.com/document_detail/131267.html) operation, you can call this operation to check whether KMS keys are authorized to ApsaraDB for MongoDB instances.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TargetRegionId) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCloudResourceAuthorized"),
		Version:     dara.String("2015-12-01"),
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
// Queries whether the data of an ApsaraDB for MongoDB instance can be restored.
//
// Description:
//
// This operation is applicable to replica set instances and sharded cluster instances.
//
// >  After you call this operation to confirm that the data of the instance can be restored, you can call the [CreateDBInstance](https://help.aliyun.com/document_detail/61763.html) operation to restore data to a new instance.
//
// @param request - CheckRecoveryConditionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckRecoveryConditionResponse
func (client *Client) CheckRecoveryConditionWithContext(ctx context.Context, request *CheckRecoveryConditionRequest, runtime *dara.RuntimeOptions) (_result *CheckRecoveryConditionResponse, _err error) {
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

	if !dara.IsNil(request.DatabaseNames) {
		query["DatabaseNames"] = request.DatabaseNames
	}

	if !dara.IsNil(request.DestRegion) {
		query["DestRegion"] = request.DestRegion
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	if !dara.IsNil(request.SourceDBInstance) {
		query["SourceDBInstance"] = request.SourceDBInstance
	}

	if !dara.IsNil(request.SourceDBInstance) {
		query["SourceDBInstance"] = request.SourceDBInstance
	}

	if !dara.IsNil(request.SrcRegion) {
		query["SrcRegion"] = request.SrcRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckRecoveryCondition"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckRecoveryConditionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a service-linked role is created.
//
// @param request - CheckServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceLinkedRoleResponse
func (client *Client) CheckServiceLinkedRoleWithContext(ctx context.Context, request *CheckServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckServiceLinkedRole"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account that is granted read-only permissions for shard nodes in an ApsaraDB for MongoDB sharded cluster instance that uses cloud disks.
//
// Description:
//
//	  You can create an account for shard nodes only in an ApsaraDB for MongoDB sharded cluster instance that uses cloud disks.
//
//		- The account is granted read-only permissions.
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
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("CreateAccount"),
		Version:     dara.String("2015-12-01"),
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
// Creates a backup set for an ApsaraDB for MongoDB instance.
//
// Description:
//
// When you call this operation, the instance must be in the Running state.
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
	if !dara.IsNil(request.BackupMethod) {
		query["BackupMethod"] = request.BackupMethod
	}

	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("CreateBackup"),
		Version:     dara.String("2015-12-01"),
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
// Creates or clones an ApsaraDB for MongoDB replica set instance.
//
// Description:
//
// Make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/product/apsaradb-for-mongodb/pricing) of ApsaraDB for MongoDB before you call this operation.
//
// For more information about the instance types of ApsaraDB for MongoDB instances, see [Instance types](https://www.alibabacloud.com/help/en/mongodb/product-overview/instance-types-1).
//
// To create sharded cluster instances, you can call the [CreateShardingDBInstance](~~CreateShardingDBInstance~~) operation.
//
// @param request - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstanceWithContext(ctx context.Context, request *CreateDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceStorage) {
		query["DBInstanceStorage"] = request.DBInstanceStorage
	}

	if !dara.IsNil(request.DatabaseNames) {
		query["DatabaseNames"] = request.DatabaseNames
	}

	if !dara.IsNil(request.Encrypted) {
		query["Encrypted"] = request.Encrypted
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.GlobalSecurityGroupIds) {
		query["GlobalSecurityGroupIds"] = request.GlobalSecurityGroupIds
	}

	if !dara.IsNil(request.HiddenZoneId) {
		query["HiddenZoneId"] = request.HiddenZoneId
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.ProvisionedIops) {
		query["ProvisionedIops"] = request.ProvisionedIops
	}

	if !dara.IsNil(request.ReadonlyReplicas) {
		query["ReadonlyReplicas"] = request.ReadonlyReplicas
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicationFactor) {
		query["ReplicationFactor"] = request.ReplicationFactor
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	if !dara.IsNil(request.SecondaryZoneId) {
		query["SecondaryZoneId"] = request.SecondaryZoneId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.SrcDBInstanceId) {
		query["SrcDBInstanceId"] = request.SrcDBInstanceId
	}

	if !dara.IsNil(request.SrcRegion) {
		query["SrcRegion"] = request.SrcRegion
	}

	if !dara.IsNil(request.StorageEngine) {
		query["StorageEngine"] = request.StorageEngine
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
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
		Action:      dara.String("CreateDBInstance"),
		Version:     dara.String("2015-12-01"),
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
// Creates a global IP whitelist template.
//
// @param request - CreateGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGlobalSecurityIPGroupResponse
func (client *Client) CreateGlobalSecurityIPGroupWithContext(ctx context.Context, request *CreateGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateGlobalSecurityIPGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GIpList) {
		query["GIpList"] = request.GIpList
	}

	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGlobalSecurityIPGroup"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a shard or mongos node to an ApsaraDB for MongoDB instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing) of ApsaraDB for MongoDB.
//
// This operation applies only to sharded cluster instances.
//
// @param request - CreateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeResponse
func (client *Client) CreateNodeWithContext(ctx context.Context, request *CreateNodeRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeResponse, _err error) {
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

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeClass) {
		query["NodeClass"] = request.NodeClass
	}

	if !dara.IsNil(request.NodeStorage) {
		query["NodeStorage"] = request.NodeStorage
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReadonlyReplicas) {
		query["ReadonlyReplicas"] = request.ReadonlyReplicas
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ShardDirect) {
		query["ShardDirect"] = request.ShardDirect
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNode"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch adds mongos or shard nodes for a sharded cluster instance.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/product/apsaradb-for-mongodb/pricing) of ApsaraDB for MongoDB.
//
// This operation is applicable only to sharded cluster instances.
//
// @param request - CreateNodeBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeBatchResponse
func (client *Client) CreateNodeBatchWithContext(ctx context.Context, request *CreateNodeBatchRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeBatchResponse, _err error) {
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

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.FromApp) {
		query["FromApp"] = request.FromApp
	}

	if !dara.IsNil(request.NodesInfo) {
		query["NodesInfo"] = request.NodesInfo
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ShardDirect) {
		query["ShardDirect"] = request.ShardDirect
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNodeBatch"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodeBatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateNodeRoleTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeRoleTagResponse
func (client *Client) CreateNodeRoleTagWithContext(ctx context.Context, request *CreateNodeRoleTagRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeRoleTagResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ShardList) {
		query["ShardList"] = request.ShardList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNodeRoleTag"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodeRoleTagResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or clones an ApsaraDB for MongoDB sharded cluster instance.
//
// Description:
//
//	  Make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/product/apsaradb-for-mongodb/pricing) of ApsaraDB for MongoDB before you call this operation.
//
//		- For more information about the instance types of ApsaraDB for MongoDB, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
//
//		- To create standalone instances and replica set instances, you can call the [CreateDBInstance](https://help.aliyun.com/document_detail/61763.html) operation.
//
// @param request - CreateShardingDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateShardingDBInstanceResponse
func (client *Client) CreateShardingDBInstanceWithContext(ctx context.Context, request *CreateShardingDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateShardingDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountPassword) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigServer) {
		query["ConfigServer"] = request.ConfigServer
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DestRegion) {
		query["DestRegion"] = request.DestRegion
	}

	if !dara.IsNil(request.Encrypted) {
		query["Encrypted"] = request.Encrypted
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.GlobalSecurityGroupIds) {
		query["GlobalSecurityGroupIds"] = request.GlobalSecurityGroupIds
	}

	if !dara.IsNil(request.HiddenZoneId) {
		query["HiddenZoneId"] = request.HiddenZoneId
	}

	if !dara.IsNil(request.Mongos) {
		query["Mongos"] = request.Mongos
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.ProvisionedIops) {
		query["ProvisionedIops"] = request.ProvisionedIops
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicaSet) {
		query["ReplicaSet"] = request.ReplicaSet
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	if !dara.IsNil(request.SecondaryZoneId) {
		query["SecondaryZoneId"] = request.SecondaryZoneId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.SrcDBInstanceId) {
		query["SrcDBInstanceId"] = request.SrcDBInstanceId
	}

	if !dara.IsNil(request.SrcRegion) {
		query["SrcRegion"] = request.SrcRegion
	}

	if !dara.IsNil(request.StorageEngine) {
		query["StorageEngine"] = request.StorageEngine
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
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
		Action:      dara.String("CreateShardingDBInstance"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateShardingDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除MongoDB备份集
//
// @param request - DeleteBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupResponse
func (client *Client) DeleteBackupWithContext(ctx context.Context, request *DeleteBackupRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DeleteBackup"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases an ApsaraDB for MongoDB instance.
//
// Description:
//
// Before you call this operation, make sure that the instance meets the following requirements
//
//   - The instance is in the Running state.
//
//   - The billing method of the instance is pay-as-you-go.
//
// > After an instance is released, all data in the instance is cleared and cannot be recovered. Proceed with caution.
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DeleteDBInstance"),
		Version:     dara.String("2015-12-01"),
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
// Deletes a global IP whitelist template.
//
// @param request - DeleteGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGlobalSecurityIPGroupResponse
func (client *Client) DeleteGlobalSecurityIPGroupWithContext(ctx context.Context, request *DeleteGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteGlobalSecurityIPGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGlobalSecurityIPGroup"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a shard or mongos node from an ApsaraDB for MongoDB sharded cluster instance.
//
// Description:
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The instance is in the Running state.
//
//   - The instance is a sharded cluster instance.
//
//   - The billing method of the instance is pay-as-you-go.
//
//   - The number of the shard or mongos nodes in the instance is greater than two.
//
// @param request - DeleteNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeResponse
func (client *Client) DeleteNodeWithContext(ctx context.Context, request *DeleteNodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteNodeResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DeleteNode"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the database accounts of an ApsaraDB for MongoDB instance.
//
// Description:
//
// >  This operation can be used to query only the information of the root account.
//
// @param request - DescribeAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountsResponse
func (client *Client) DescribeAccountsWithContext(ctx context.Context, request *DescribeAccountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeAccounts"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeActiveOperationMaintenanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationMaintenanceConfigResponse
func (client *Client) DescribeActiveOperationMaintenanceConfigWithContext(ctx context.Context, request *DescribeActiveOperationMaintenanceConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationMaintenanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeActiveOperationMaintenanceConfig"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationMaintenanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about  tasks of an ApsaraDB for MongoDB instance.
//
// @param request - DescribeActiveOperationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTaskResponse
func (client *Client) DescribeActiveOperationTaskWithContext(ctx context.Context, request *DescribeActiveOperationTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsHistory) {
		query["IsHistory"] = request.IsHistory
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

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveOperationTask"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of operation and maintenance tasks on an ApsaraDB for MongoDB instance.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("DescribeActiveOperationTaskCount"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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

// @param request - DescribeActiveOperationTaskRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTaskRegionResponse
func (client *Client) DescribeActiveOperationTaskRegionWithContext(ctx context.Context, request *DescribeActiveOperationTaskRegionRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTaskRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsHistory) {
		query["IsHistory"] = request.IsHistory
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeActiveOperationTaskRegion"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTaskRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types of Operation and Maintenance tasks and the number of tasks of each type for an ApsaraDB for MongoDB instance.
//
// Description:
//
// This operation is no longer updated and will be unavailable.
//
// @param request - DescribeActiveOperationTaskTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActiveOperationTaskTypeResponse
func (client *Client) DescribeActiveOperationTaskTypeWithContext(ctx context.Context, request *DescribeActiveOperationTaskTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeActiveOperationTaskTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsHistory) {
		query["IsHistory"] = request.IsHistory
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("DescribeActiveOperationTaskType"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActiveOperationTaskTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of operation and maintenance tasks initiated for an ApsaraDB for MongoDB instance.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowCancel) {
		query["AllowCancel"] = request.AllowCancel
	}

	if !dara.IsNil(request.AllowChange) {
		query["AllowChange"] = request.AllowChange
	}

	if !dara.IsNil(request.ChangeLevel) {
		query["ChangeLevel"] = request.ChangeLevel
	}

	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.InsName) {
		query["InsName"] = request.InsName
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

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      dara.String("DescribeActiveOperationTasks"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Queries the types of entries in the audit log collected for an ApsaraDB for MongoDB instance.
//
// Description:
//
//	  The instance must be in the running state when you call this operation.
//
//		- This operation is applicable only to **general-purpose local-disk*	- and **dedicated local-disk*	- instances.
//
//		- You can call this operation up to 30 times per minute. To call this operation at a higher frequency, use a Logstore. For more information, see [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
//
// @param request - DescribeAuditLogFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditLogFilterResponse
func (client *Client) DescribeAuditLogFilterWithContext(ctx context.Context, request *DescribeAuditLogFilterRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditLogFilterResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuditLogFilter"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditLogFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the audit log feature is enabled for an ApsaraDB for MongoDB instance.
//
// Description:
//
//	  The instance must be in the running state when you call this operation.
//
//		- This operation is applicable only to **general-purpose local-disk*	- and **dedicated local-disk*	- instances.
//
//		- You can call this operation up to 30 times per minute. To call this operation at a higher frequency, use a Logstore. For more information, see [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
//
// @param request - DescribeAuditPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditPolicyResponse
func (client *Client) DescribeAuditPolicyWithContext(ctx context.Context, request *DescribeAuditPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditPolicyResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeAuditPolicy"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audit logs of an ApsaraDB for MongoDB instance.
//
// Description:
//
//	  When you call this operation, ensure that the audit log feature of the instance is enabled. Otherwise, the operation returns an empty audit log.
//
//		- This operation is applicable only to **general-purpose local-disk*	- and **dedicated local-disk*	- instances.
//
//		- You can call this operation up to 30 times per minute. To call this operation at a higher frequency, use a Logstore. For more information, see [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
//
// @param request - DescribeAuditRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuditRecordsResponse
func (client *Client) DescribeAuditRecordsWithContext(ctx context.Context, request *DescribeAuditRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuditRecordsResponse, _err error) {
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

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Form) {
		query["Form"] = request.Form
	}

	if !dara.IsNil(request.LogicalOperator) {
		query["LogicalOperator"] = request.LogicalOperator
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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

	if !dara.IsNil(request.QueryKeywords) {
		query["QueryKeywords"] = request.QueryKeywords
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

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuditRecords"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuditRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of the zones that are supported by an ApsaraDB for MongoDB instance.
//
// Description:
//
// Queries the zones in which an ApsaraDB for MongoDB instance can be deployed under specified purchase conditions. The region ID is required in the purchase condition.
//
// @param request - DescribeAvailabilityZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailabilityZonesResponse
func (client *Client) DescribeAvailabilityZonesWithContext(ctx context.Context, request *DescribeAvailabilityZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailabilityZonesResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.ExcludeSecondaryZoneId) {
		query["ExcludeSecondaryZoneId"] = request.ExcludeSecondaryZoneId
	}

	if !dara.IsNil(request.ExcludeZoneId) {
		query["ExcludeZoneId"] = request.ExcludeZoneId
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.MongoType) {
		query["MongoType"] = request.MongoType
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

	if !dara.IsNil(request.ReplicationFactor) {
		query["ReplicationFactor"] = request.ReplicationFactor
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StorageSupport) {
		query["StorageSupport"] = request.StorageSupport
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailabilityZones"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailabilityZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the engine versions to which an ApsaraDB for MongoDB instance can be upgraded.
//
// @param request - DescribeAvailableEngineVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableEngineVersionResponse
func (client *Client) DescribeAvailableEngineVersionWithContext(ctx context.Context, request *DescribeAvailableEngineVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableEngineVersionResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeAvailableEngineVersion"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableEngineVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available resources in the specified zone.
//
// @param request - DescribeAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableResourceResponse
func (client *Client) DescribeAvailableResourceWithContext(ctx context.Context, request *DescribeAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
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

	if !dara.IsNil(request.ReplicationFactor) {
		query["ReplicationFactor"] = request.ReplicationFactor
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableResource"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases at a specified time or the databases in a specified backup set before you restore a database for an ApsaraDB for MongoDB instance.
//
// Description:
//
// You can call the [CreateDBInstance](https://help.aliyun.com/document_detail/61763.html) operation to restore a database for an ApsaraDB for MongoDB instance. For more information, see [Restore one database of an ApsaraDB for MongoDB instance](https://help.aliyun.com/document_detail/112274.html).
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The instance was created after March 26, 2019.
//
//   - The instance is located in the China (Qingdao), China (Beijing), China (Zhangjiakou), China (Hohhot), China (Hangzhou), China (Shanghai), China (Shenzhen), or Singapore region. Other regions are not supported.
//
//   - The instance is a replica set instance.
//
//   - The instance runs MongoDB 3.4, MongoDB 4.0, or MongoDB 4.2. In addition, the instance uses local disks to store data.
//
//   - The storage engine of the instance is WiredTiger.
//
// @param request - DescribeBackupDBsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupDBsResponse
func (client *Client) DescribeBackupDBsWithContext(ctx context.Context, request *DescribeBackupDBsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupDBsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.SourceDBInstance) {
		query["SourceDBInstance"] = request.SourceDBInstance
	}

	if !dara.IsNil(request.SourceDBInstance) {
		query["SourceDBInstance"] = request.SourceDBInstance
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupDBs"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupDBsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup policy of an ApsaraDB for MongoDB instance.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SrcRegion) {
		query["SrcRegion"] = request.SrcRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2015-12-01"),
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
// Queries the storage used for backup in an ApsaraDB for MongoDB replica set or sharded cluster instance that uses cloud disks. Note that you are charged only for the backup-used storage of each shard in a sharded cluster instance. You can call this operation only to query the storage used by a single shard in the instance for backup.
//
// @param request - DescribeBackupStorageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupStorageResponse
func (client *Client) DescribeBackupStorageWithContext(ctx context.Context, request *DescribeBackupStorageRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupStorageResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
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
		Action:      dara.String("DescribeBackupStorage"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupStorageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries backup tasks running in an ApsaraDB for MongoDB replica set or sharded cluster instance that uses cloud disks.
//
// @param request - DescribeBackupTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupTasksResponse
func (client *Client) DescribeBackupTasksWithContext(ctx context.Context, request *DescribeBackupTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupJobId) {
		query["BackupJobId"] = request.BackupJobId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupTasks"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup sets of an ApsaraDB for MongoDB instance.
//
// @param request - DescribeBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupsResponse
func (client *Client) DescribeBackupsWithContext(ctx context.Context, request *DescribeBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
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

	if !dara.IsNil(request.BackupJobId) {
		query["BackupJobId"] = request.BackupJobId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DestRegion) {
		query["DestRegion"] = request.DestRegion
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SrcRegion) {
		query["SrcRegion"] = request.SrcRegion
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackups"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup sets of an ApsaraDB for MongoDB sharded cluster instance that uses cloud disks.
//
// Description:
//
// For a sharded cluster instance that is created before October 19, 2023 and uses cloud disks, you must call the [TransferClusterBackup](https://help.aliyun.com/document_detail/2587931.html) operation to switch the instance from the shard backup mode to the cluster backup mode before you call the DescribeClusterBackups operation.
//
// By default, cloud disk-based sharded cluster instances that are created after October 19, 2023 are in the cluster backup mode.
//
// @param request - DescribeClusterBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterBackupsResponse
func (client *Client) DescribeClusterBackupsWithContext(ctx context.Context, request *DescribeClusterBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterBackupsResponse, _err error) {
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

	if !dara.IsNil(request.BackupJobId) {
		query["BackupJobId"] = request.BackupJobId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DestRegion) {
		query["DestRegion"] = request.DestRegion
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IsOnlyGetClusterBackUp) {
		query["IsOnlyGetClusterBackUp"] = request.IsOnlyGetClusterBackUp
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SrcRegion) {
		query["SrcRegion"] = request.SrcRegion
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterBackups"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterBackupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time range to which you can restore the data of an ApsaraDB for MongoDB sharded cluster instance that uses cloud disks.
//
// Description:
//
//	  The instance is an ApsaraDB for MongoDB sharded cluster instance that runs MongoDB 4.4 or later and uses enhanced SSDs (ESSDs) to store data.
//
//		- You can call the TransferClusterBackup operation only for instances that are created before October 19, 2023 to switch the instances to the cluster backup mode. The DescribeClusterRecoverTime operation is applicable only to instances that are switched to the cluster backup mode or instances that are created on or after October 19, 2023.
//
// @param request - DescribeClusterRecoverTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterRecoverTimeResponse
func (client *Client) DescribeClusterRecoverTimeWithContext(ctx context.Context, request *DescribeClusterRecoverTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterRecoverTimeResponse, _err error) {
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

	if !dara.IsNil(request.DestRegion) {
		query["DestRegion"] = request.DestRegion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SrcRegion) {
		query["SrcRegion"] = request.SrcRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClusterRecoverTime"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClusterRecoverTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an ApsaraDB for MongoDB instance.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.IsDelete) {
		query["IsDelete"] = request.IsDelete
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceAttribute"),
		Version:     dara.String("2015-12-01"),
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
// Queries the details of a key for an ApsaraDB for MongoDB instance.
//
// Description:
//
// When you call the DescribeDBInstanceEncryptionKey operation, the instance must have transparent data encryption (TDE) enabled in BYOK mode. You can call the [ModifyDBInstanceTDE](https://help.aliyun.com/document_detail/131267.html) operation to enable TDE.
//
// @param request - DescribeDBInstanceEncryptionKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceEncryptionKeyResponse
func (client *Client) DescribeDBInstanceEncryptionKeyWithContext(ctx context.Context, request *DescribeDBInstanceEncryptionKeyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceEncryptionKeyResponse, _err error) {
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

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeDBInstanceEncryptionKey"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceEncryptionKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the collection frequency of monitoring data for an ApsaraDB for MongoDB instance.
//
// @param request - DescribeDBInstanceMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceMonitorResponse
func (client *Client) DescribeDBInstanceMonitorWithContext(ctx context.Context, request *DescribeDBInstanceMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceMonitorResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeDBInstanceMonitor"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the performance data of an ApsaraDB for MongoDB instance.
//
// @param request - DescribeDBInstancePerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancePerformanceResponse
func (client *Client) DescribeDBInstancePerformanceWithContext(ctx context.Context, request *DescribeDBInstancePerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancePerformanceResponse, _err error) {
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

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReplicaSetRole) {
		query["ReplicaSetRole"] = request.ReplicaSetRole
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.SearchId) {
		query["SearchId"] = request.SearchId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstancePerformance"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Secure Sockets Layer (SSL) settings of an ApsaraDB for MongoDB instance.
//
// Description:
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The instance is in the Running state.
//
//   - The instance is a replica set instance.
//
//   - The instance runs MongoDB 3.4 or later.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeDBInstanceSSL"),
		Version:     dara.String("2015-12-01"),
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
// 查看规格信息详情
//
// @param request - DescribeDBInstanceSpecInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceSpecInfoResponse
func (client *Client) DescribeDBInstanceSpecInfoWithContext(ctx context.Context, request *DescribeDBInstanceSpecInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceSpecInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceSpecInfo"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceSpecInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the primary/secondary switching logs of an ApsaraDB for MongoDB instance.
//
// Description:
//
// Before you call this operation, make sure that the ApsaraDB for MongoDB instance meets the following requirements:
//
//   - The instance is a replica set or sharded cluster instance.
//
//   - The instance uses local physical disks to store data.
//
// @param request - DescribeDBInstanceSwitchLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceSwitchLogResponse
func (client *Client) DescribeDBInstanceSwitchLogWithContext(ctx context.Context, request *DescribeDBInstanceSwitchLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceSwitchLogResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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
		Action:      dara.String("DescribeDBInstanceSwitchLog"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceSwitchLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Transparent Data Encryption (TDE) is enabled for an ApsaraDB for MongoDB instance.
//
// Description:
//
// >  For more information about TDE, see [TDE](https://help.aliyun.com/document_detail/131048.html).
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The instance is a replica set or sharded cluster instance.
//
//   - The storage engine of the instance is WiredTiger.
//
//   - The database engine version of the instance is 4.0 or 4.2. If the database engine version is earlier than 4.0, you can call the [UpgradeDBInstanceEngineVersion](https://help.aliyun.com/document_detail/67608.html) operation to upgrade the database engine.
//
// @param request - DescribeDBInstanceTDEInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceTDEInfoResponse
func (client *Client) DescribeDBInstanceTDEInfoWithContext(ctx context.Context, request *DescribeDBInstanceTDEInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceTDEInfoResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeDBInstanceTDEInfo"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceTDEInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ApsaraDB for MongoDB instances.
//
// Description:
//
// The list of replica set and standalone instances is displayed when the **DBInstanceType*	- parameter uses the default value **replicate**. To query a list of sharded cluster instances, you must set the **DBInstanceType*	- parameter to **sharding**.
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
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ConnectionDomain) {
		query["ConnectionDomain"] = request.ConnectionDomain
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceStatus) {
		query["DBInstanceStatus"] = request.DBInstanceStatus
	}

	if !dara.IsNil(request.DBInstanceType) {
		query["DBInstanceType"] = request.DBInstanceType
	}

	if !dara.IsNil(request.DBNodeType) {
		query["DBNodeType"] = request.DBNodeType
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.Expired) {
		query["Expired"] = request.Expired
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
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

	if !dara.IsNil(request.ReplicationFactor) {
		query["ReplicationFactor"] = request.ReplicationFactor
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
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
		Action:      dara.String("DescribeDBInstances"),
		Version:     dara.String("2015-12-01"),
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
// Queries the overview information of one or more ApsaraDB for MongoDB instances.
//
// Description:
//
//	  If you do not specify an instance when you call this operation, the overview information of all instances in a specific region within this account is returned.
//
//		- Paged query is disabled for this operation.
//
// @param request - DescribeDBInstancesOverviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesOverviewResponse
func (client *Client) DescribeDBInstancesOverviewWithContext(ctx context.Context, request *DescribeDBInstancesOverviewRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.InstanceClass) {
		query["InstanceClass"] = request.InstanceClass
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ShowTags) {
		query["ShowTags"] = request.ShowTags
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
		Action:      dara.String("DescribeDBInstancesOverview"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancesOverviewResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries entries in error logs of an ApsaraDB for MongoDB instance.
//
// Description:
//
//	  This operation is applicable only to **general-purpose local-disk*	- and **dedicated local-disk*	- instances.
//
//		- You can call this operation up to 30 times per minute. To call this operation at a higher frequency, use a Logstore. For more information, see [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
//
// @param request - DescribeErrorLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeErrorLogRecordsResponse
func (client *Client) DescribeErrorLogRecordsWithContext(ctx context.Context, request *DescribeErrorLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeErrorLogRecordsResponse, _err error) {
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

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LogicalOperator) {
		query["LogicalOperator"] = request.LogicalOperator
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
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

	if !dara.IsNil(request.QueryKeywords) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeErrorLogRecords"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeErrorLogRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global IP whitelist template of an ApsaraDB for MongoDB instance.
//
// @param request - DescribeGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalSecurityIPGroupResponse
func (client *Client) DescribeGlobalSecurityIPGroupWithContext(ctx context.Context, request *DescribeGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalSecurityIPGroupResponse, _err error) {
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
		Action:      dara.String("DescribeGlobalSecurityIPGroup"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global IP whitelist templates associated with an ApsaraDB for MongoDB instance.
//
// @param request - DescribeGlobalSecurityIPGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGlobalSecurityIPGroupRelationResponse
func (client *Client) DescribeGlobalSecurityIPGroupRelationWithContext(ctx context.Context, request *DescribeGlobalSecurityIPGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *DescribeGlobalSecurityIPGroupRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
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
		Action:      dara.String("DescribeGlobalSecurityIPGroupRelation"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tasks in the task center.
//
// @param request - DescribeHistoryTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryTasksResponse
func (client *Client) DescribeHistoryTasksWithContext(ctx context.Context, request *DescribeHistoryTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromExecTime) {
		query["FromExecTime"] = request.FromExecTime
	}

	if !dara.IsNil(request.FromStartTime) {
		query["FromStartTime"] = request.FromStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.ToExecTime) {
		query["ToExecTime"] = request.ToExecTime
	}

	if !dara.IsNil(request.ToStartTime) {
		query["ToStartTime"] = request.ToStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryTasks"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the overview of a task in the task center.
//
// @param request - DescribeHistoryTasksStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryTasksStatResponse
func (client *Client) DescribeHistoryTasksStatWithContext(ctx context.Context, request *DescribeHistoryTasksStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryTasksStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromExecTime) {
		query["FromExecTime"] = request.FromExecTime
	}

	if !dara.IsNil(request.FromStartTime) {
		query["FromStartTime"] = request.FromStartTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.ToExecTime) {
		query["ToExecTime"] = request.ToExecTime
	}

	if !dara.IsNil(request.ToStartTime) {
		query["ToStartTime"] = request.ToStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryTasksStat"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryTasksStatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query whether auto-renewal is enabled for an ApsaraDB for MongoDB instance.
//
// Description:
//
// This operation is applicable to subscription instances.
//
// @param request - DescribeInstanceAutoRenewalAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceAutoRenewalAttributeResponse
func (client *Client) DescribeInstanceAutoRenewalAttributeWithContext(ctx context.Context, request *DescribeInstanceAutoRenewalAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceAutoRenewalAttributeResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceType) {
		query["DBInstanceType"] = request.DBInstanceType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAutoRenewalAttribute"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time required to restore the data of an ApsaraDB for MongoDB replica set instance that uses cloud disks.
//
// @param request - DescribeInstanceRecoverTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceRecoverTimeResponse
func (client *Client) DescribeInstanceRecoverTimeWithContext(ctx context.Context, request *DescribeInstanceRecoverTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceRecoverTimeResponse, _err error) {
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

	if !dara.IsNil(request.DestRegion) {
		query["DestRegion"] = request.DestRegion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SrcRegion) {
		query["SrcRegion"] = request.SrcRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceRecoverTime"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceRecoverTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the release notes of the minor versions of an ApsaraDB for MongoDB instance.
//
// @param request - DescribeKernelReleaseNotesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKernelReleaseNotesResponse
func (client *Client) DescribeKernelReleaseNotesWithContext(ctx context.Context, request *DescribeKernelReleaseNotesRequest, runtime *dara.RuntimeOptions) (_result *DescribeKernelReleaseNotesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KernelVersion) {
		query["KernelVersion"] = request.KernelVersion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeKernelReleaseNotes"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKernelReleaseNotesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Key Management Service (KMS) keys that are available for disk encryption.
//
// Description:
//
// Queried keys are available only for disk encryption.
//
// @param request - DescribeKmsKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKmsKeysResponse
func (client *Client) DescribeKmsKeysWithContext(ctx context.Context, request *DescribeKmsKeysRequest, runtime *dara.RuntimeOptions) (_result *DescribeKmsKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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
		Action:      dara.String("DescribeKmsKeys"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKmsKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logging configurations of an ApsaraDB for MongoDB instance.
//
// Description:
//
// This operation is applicable only to **general-purpose local-disk*	- and **dedicated local-disk*	- instances.
//
// This operation depends on the audit log feature of ApsaraDB for MongoDB. You can enable the audit log feature based on your business requirements. For more information, see [Enable the audit log feature](https://help.aliyun.com/document_detail/59903.html).
//
//   - Starting from January 6, 2022, the official edition of the audit log feature has been launched in all regions, and new applications for the free trial edition have ended. For more information, see [Notice on official launch of the pay-as-you-go audit log feature and no more application for the free trial edition](https://help.aliyun.com/document_detail/377480.html)
//
//   - You are charged for the official edition of the audit log feature based on the storage capacity that is consumed by audit logs and the retention period of the audit logs. For more information, see [Pricing of ApsaraDB for MongoDB instances](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing).
//
// @param request - DescribeMongoDBLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMongoDBLogConfigResponse
func (client *Client) DescribeMongoDBLogConfigWithContext(ctx context.Context, request *DescribeMongoDBLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeMongoDBLogConfigResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeMongoDBLogConfig"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMongoDBLogConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameter modification records of an ApsaraDB for MongoDB instance.
//
// @param request - DescribeParameterModificationHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterModificationHistoryResponse
func (client *Client) DescribeParameterModificationHistoryWithContext(ctx context.Context, request *DescribeParameterModificationHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterModificationHistoryResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeParameterModificationHistory"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterModificationHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of default parameter templates for ApsaraDB for MongoDB instances.
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
	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
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

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeParameterTemplates"),
		Version:     dara.String("2015-12-01"),
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
// Queries the parameter settings of an ApsaraDB for MongoDB instance.
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
	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ExtraParam) {
		query["ExtraParam"] = request.ExtraParam
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeParameters"),
		Version:     dara.String("2015-12-01"),
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
// Queries the pricing information of an ApsaraDB for MongoDB instance.
//
// @param request - DescribePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceResponse
func (client *Client) DescribePriceWithContext(ctx context.Context, request *DescribePriceRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstances) {
		query["DBInstances"] = request.DBInstances
	}

	if !dara.IsNil(request.OrderParamOut) {
		query["OrderParamOut"] = request.OrderParamOut
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("DescribePrice"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRdsVSwitchsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsVSwitchsResponse
func (client *Client) DescribeRdsVSwitchsWithContext(ctx context.Context, request *DescribeRdsVSwitchsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsVSwitchsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      dara.String("DescribeRdsVSwitchs"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdsVpcs"),
		Version:     dara.String("2015-12-01"),
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
// Queries all regions and zones supported for an ApsaraDB for MongoDB instance.
//
// Description:
//
// >  To query available regions and zones in which an ApsaraDB for MongoDB instance can be created, call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/149719.html) operation.
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
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
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
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2015-12-01"),
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
// Queries the monthly renewal price of an ApsaraDB for MongoDB instance.
//
// Description:
//
// This operation is applicable to subscription instances.
//
// @param request - DescribeRenewalPriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRenewalPriceResponse
func (client *Client) DescribeRenewalPriceWithContext(ctx context.Context, request *DescribeRenewalPriceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRenewalPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeRenewalPrice"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRenewalPriceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the role and connection information of an ApsaraDB for MongoDB instance.
//
// Description:
//
// This operation is applicable to replica set instances and standalone instances, but not to sharded cluster instances.
//
// @param request - DescribeReplicaSetRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeReplicaSetRoleResponse
func (client *Client) DescribeReplicaSetRoleWithContext(ctx context.Context, request *DescribeReplicaSetRoleRequest, runtime *dara.RuntimeOptions) (_result *DescribeReplicaSetRoleResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeReplicaSetRole"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeReplicaSetRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries ApsaraDB for MongoDB instances whose backups are restored within seven days.
//
// @param request - DescribeRestoreDBInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRestoreDBInstanceListResponse
func (client *Client) DescribeRestoreDBInstanceListWithContext(ctx context.Context, request *DescribeRestoreDBInstanceListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRestoreDBInstanceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreationTimeAfter) {
		query["CreationTimeAfter"] = request.CreationTimeAfter
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("DescribeRestoreDBInstanceList"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRestoreDBInstanceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRoleTagStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoleTagStatusResponse
func (client *Client) DescribeRoleTagStatusWithContext(ctx context.Context, request *DescribeRoleTagStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeRoleTagStatusResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeRoleTagStatus"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoleTagStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the role and zone of each node in an ApsaraDB for MongoDB instance.
//
// Description:
//
// > For more information, see [View the zone of a node](https://help.aliyun.com/document_detail/123825.html).
//
// This operation is applicable to replica set instances and sharded cluster instances, but cannot be performed on standalone instances.
//
// @param request - DescribeRoleZoneInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRoleZoneInfoResponse
func (client *Client) DescribeRoleZoneInfoWithContext(ctx context.Context, request *DescribeRoleZoneInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeRoleZoneInfoResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeRoleZoneInfo"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRoleZoneInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries entries in operational logs of an ApsaraDB for MongoDB instance.
//
// Description:
//
//	  This operation is applicable only to **general-purpose local-disk*	- and **dedicated local-disk*	- instances.
//
//		- You can call this operation up to 30 times per minute. To call this operation at a higher frequency, use a Logstore. For more information, see [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
//
// @param request - DescribeRunningLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRunningLogRecordsResponse
func (client *Client) DescribeRunningLogRecordsWithContext(ctx context.Context, request *DescribeRunningLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRunningLogRecordsResponse, _err error) {
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

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LogicalOperator) {
		query["LogicalOperator"] = request.LogicalOperator
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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

	if !dara.IsNil(request.QueryKeywords) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRunningLogRecords"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRunningLogRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query ECS security groups that are bound to an ApsaraDB for MongoDB instance.
//
// @param request - DescribeSecurityGroupConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityGroupConfigurationResponse
func (client *Client) DescribeSecurityGroupConfigurationWithContext(ctx context.Context, request *DescribeSecurityGroupConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityGroupConfigurationResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeSecurityGroupConfiguration"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityGroupConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the IP whitelists of an ApsaraDB for MongoDB instance.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ShowHDMIps) {
		query["ShowHDMIps"] = request.ShowHDMIps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityIps"),
		Version:     dara.String("2015-12-01"),
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
// Queries connection information about an ApsaraDB for MongoDB sharded cluster instance.
//
// Description:
//
// This operation is applicable only to sharded cluster instances.
//
// @param request - DescribeShardingNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeShardingNetworkAddressResponse
func (client *Client) DescribeShardingNetworkAddressWithContext(ctx context.Context, request *DescribeShardingNetworkAddressRequest, runtime *dara.RuntimeOptions) (_result *DescribeShardingNetworkAddressResponse, _err error) {
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

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("DescribeShardingNetworkAddress"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeShardingNetworkAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of entries in slow query logs of an ApsaraDB for MongoDB instance.
//
// Description:
//
//	  This operation is applicable only to **general-purpose local-disk*	- and **dedicated local-disk*	- instances.
//
//		- You can call this operation up to 30 times per minute. To call this operation at a higher frequency, use a Logstore. For more information, see [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.LogicalOperator) {
		query["LogicalOperator"] = request.LogicalOperator
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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

	if !dara.IsNil(request.QueryKeywords) {
		query["QueryKeywords"] = request.QueryKeywords
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("DescribeSlowLogRecords"),
		Version:     dara.String("2015-12-01"),
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
// Queries all tags in a specified region.
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
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2015-12-01"),
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
// Queries the list of custom keys for an ApsaraDB for MongoDB instance.
//
// Description:
//
// You can use the custom key obtained by calling the DescribeUserEncryptionKeyList operation to enable TDE. For more information, see [ModifyDBInstanceTDE](https://help.aliyun.com/document_detail/131267.html).
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleARN) {
		query["RoleARN"] = request.RoleARN
	}

	if !dara.IsNil(request.TargetRegionId) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserEncryptionKeyList"),
		Version:     dara.String("2015-12-01"),
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

// @param request - DescribeVpcsForMongoDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcsForMongoDBResponse
func (client *Client) DescribeVpcsForMongoDBWithContext(ctx context.Context, request *DescribeVpcsForMongoDBRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcsForMongoDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcsForMongoDB"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcsForMongoDBResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Destroys an ApsaraDB for MongoDB instance.
//
// Description:
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The instance is a replica set instance or a sharded cluster instance that uses local disks.
//
//   - The billing method of the instance is subscription.
//
//   - The instance has expired and is in the **Locking*	- state.
//
// **
//
// **Warning*	- Data cannot be restored after the instance is destroyed. Proceed with caution.
//
// @param request - DestroyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DestroyInstanceResponse
func (client *Client) DestroyInstanceWithContext(ctx context.Context, request *DestroyInstanceRequest, runtime *dara.RuntimeOptions) (_result *DestroyInstanceResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("DestroyInstance"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DestroyInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether sufficient resources are available in a region in which you want to create or upgrade an ApsaraDB for MongoDB instance.
//
// Description:
//
// This operation is applicable to replica set instances and sharded cluster instances. You can call this operation to check whether resources are sufficient for creating an instance, upgrading a replica set or sharded cluster instance, or upgrading a single node of the sharded cluster instance.
//
// > You can call this operation a maximum of 200 times per minute.
//
// @param request - EvaluateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateResourceResponse
func (client *Client) EvaluateResourceWithContext(ctx context.Context, request *EvaluateResourceRequest, runtime *dara.RuntimeOptions) (_result *EvaluateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReadonlyReplicas) {
		query["ReadonlyReplicas"] = request.ReadonlyReplicas
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicationFactor) {
		query["ReplicationFactor"] = request.ReplicationFactor
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ShardsInfo) {
		query["ShardsInfo"] = request.ShardsInfo
	}

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EvaluateResource"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EvaluateResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the relationship between ApsaraDB for MongoDB instances and tags.
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2015-12-01"),
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
// Migrates an ApsaraDB for MongoDB instance to a specific zone.
//
// Description:
//
//	  This operation is available only for replica set instances that run MongoDB 4.2 or earlier and sharded cluster instances.
//
//		- If you have applied for a public endpoint for the ApsaraDB for MongoDB instance, you must call the [ReleasePublicNetworkAddress](https://help.aliyun.com/document_detail/67604.html) operation to release the public endpoint before you call the MigrateAvailableZone operation.
//
//		- Transparent data encryption (TDE) is disabled for the ApsaraDB for MongoDB instance.
//
//		- The source zone and the destination zone belong to the same region.
//
//		- A vSwitch is created in the destination zone. This prerequisite must be met if the instance resides in a virtual private cloud (VPC). For more information about how to create a vSwitch, see [Work with vSwitches](https://help.aliyun.com/document_detail/65387.html).
//
// @param request - MigrateAvailableZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateAvailableZoneResponse
func (client *Client) MigrateAvailableZoneWithContext(ctx context.Context, request *MigrateAvailableZoneRequest, runtime *dara.RuntimeOptions) (_result *MigrateAvailableZoneResponse, _err error) {
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

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.HiddenZoneId) {
		query["HiddenZoneId"] = request.HiddenZoneId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecondaryZoneId) {
		query["SecondaryZoneId"] = request.SecondaryZoneId
	}

	if !dara.IsNil(request.Vswitch) {
		query["Vswitch"] = request.Vswitch
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateAvailableZone"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateAvailableZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to migrate an ApsaraDB for MongoDB instance to another zone.
//
// Description:
//
// This operation is applicable only to replica set instances, but not to standalone instances or sharded cluster instances.
//
// >  If you have applied for a public endpoint of the instance, you must first call the [ReleasePublicNetworkAddress](https://help.aliyun.com/document_detail/67604.html) operation to release the public endpoint.
//
// @param request - MigrateToOtherZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateToOtherZoneResponse
func (client *Client) MigrateToOtherZoneWithContext(ctx context.Context, request *MigrateToOtherZoneRequest, runtime *dara.RuntimeOptions) (_result *MigrateToOtherZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      dara.String("MigrateToOtherZone"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateToOtherZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of the root account in an ApsaraDB for MongoDB instance.
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

	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyAccountDescription"),
		Version:     dara.String("2015-12-01"),
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

// @param request - ModifyActiveOperationMaintenanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActiveOperationMaintenanceConfigResponse
func (client *Client) ModifyActiveOperationMaintenanceConfigWithContext(ctx context.Context, request *ModifyActiveOperationMaintenanceConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyActiveOperationMaintenanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CycleTime) {
		query["CycleTime"] = request.CycleTime
	}

	if !dara.IsNil(request.CycleType) {
		query["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.MaintainEndTime) {
		query["MaintainEndTime"] = request.MaintainEndTime
	}

	if !dara.IsNil(request.MaintainStartTime) {
		query["MaintainStartTime"] = request.MaintainStartTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationMaintenanceConfig"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyActiveOperationMaintenanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the switching time of scheduled O\\\\\\&M tasks for an ApsaraDB for MongoDB instance.
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationTasks"),
		Version:     dara.String("2015-12-01"),
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
// Queries the types of logs collected by the audit log feature of an ApsaraDB for MongoDB instance.
//
// Description:
//
//	  The instance must be in the running state when you call this operation.
//
//		- This operation is applicable only to **general-purpose local-disk*	- or **dedicated local-disk*	- instances.
//
//		- You can call this operation up to 30 times per minute. To call this operation at a higher frequency, use a Logstore. For more information, see [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
//
// @param request - ModifyAuditLogFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAuditLogFilterResponse
func (client *Client) ModifyAuditLogFilterWithContext(ctx context.Context, request *ModifyAuditLogFilterRequest, runtime *dara.RuntimeOptions) (_result *ModifyAuditLogFilterResponse, _err error) {
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

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleType) {
		query["RoleType"] = request.RoleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAuditLogFilter"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAuditLogFilterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the audit log feature or configures the log storage duration for an ApsaraDB for MongoDB instance.
//
// Description:
//
//	  This operation is applicable only to **general-purpose local-disk*	- and **dedicated local-disk*	- instances.
//
//		- You can call this operation up to 30 times per minute. To call this operation at a higher frequency, use a Logstore. For more information, see [Manage a Logstore](https://help.aliyun.com/document_detail/48990.html).
//
// @param request - ModifyAuditPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAuditPolicyResponse
func (client *Client) ModifyAuditPolicyWithContext(ctx context.Context, request *ModifyAuditPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAuditPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditLogSwitchSource) {
		query["AuditLogSwitchSource"] = request.AuditLogSwitchSource
	}

	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.StoragePeriod) {
		query["StoragePeriod"] = request.StoragePeriod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAuditPolicy"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAuditPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改MongoDB备份集的过期时间
//
// @param request - ModifyBackupExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupExpireTimeResponse
func (client *Client) ModifyBackupExpireTimeWithContext(ctx context.Context, request *ModifyBackupExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupExpireTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupExpireTime) {
		query["BackupExpireTime"] = request.BackupExpireTime
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyBackupExpireTime"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupExpireTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a backup policy for an ApsaraDB for MongoDB instance.
//
// Description:
//
// The cross-region backup feature is suitable only for replica set or sharded cluster instances that use cloud disks.
//
// @param request - ModifyBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupPolicyResponse
func (client *Client) ModifyBackupPolicyWithContext(ctx context.Context, request *ModifyBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupInterval) {
		query["BackupInterval"] = request.BackupInterval
	}

	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.BackupRetentionPolicyOnClusterDeletion) {
		query["BackupRetentionPolicyOnClusterDeletion"] = request.BackupRetentionPolicyOnClusterDeletion
	}

	if !dara.IsNil(request.CrossBackupPeriod) {
		query["CrossBackupPeriod"] = request.CrossBackupPeriod
	}

	if !dara.IsNil(request.CrossBackupType) {
		query["CrossBackupType"] = request.CrossBackupType
	}

	if !dara.IsNil(request.CrossLogRetentionType) {
		query["CrossLogRetentionType"] = request.CrossLogRetentionType
	}

	if !dara.IsNil(request.CrossLogRetentionValue) {
		query["CrossLogRetentionValue"] = request.CrossLogRetentionValue
	}

	if !dara.IsNil(request.CrossRetentionType) {
		query["CrossRetentionType"] = request.CrossRetentionType
	}

	if !dara.IsNil(request.CrossRetentionValue) {
		query["CrossRetentionValue"] = request.CrossRetentionValue
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DestRegion) {
		query["DestRegion"] = request.DestRegion
	}

	if !dara.IsNil(request.EnableBackupLog) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !dara.IsNil(request.EnableCrossLogBackup) {
		query["EnableCrossLogBackup"] = request.EnableCrossLogBackup
	}

	if !dara.IsNil(request.HighFrequencyBackupRetention) {
		query["HighFrequencyBackupRetention"] = request.HighFrequencyBackupRetention
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.LogBackupRetentionPeriod) {
		query["LogBackupRetentionPeriod"] = request.LogBackupRetentionPeriod
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PreferredBackupPeriod) {
		query["PreferredBackupPeriod"] = request.PreferredBackupPeriod
	}

	if !dara.IsNil(request.PreferredBackupTime) {
		query["PreferredBackupTime"] = request.PreferredBackupTime
	}

	if !dara.IsNil(request.PreserveOneEachHour) {
		query["PreserveOneEachHour"] = request.PreserveOneEachHour
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SnapshotBackupType) {
		query["SnapshotBackupType"] = request.SnapshotBackupType
	}

	if !dara.IsNil(request.SrcRegion) {
		query["SrcRegion"] = request.SrcRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupPolicy"),
		Version:     dara.String("2015-12-01"),
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

// @param request - ModifyDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceAttributeResponse
func (client *Client) ModifyDBInstanceAttributeWithContext(ctx context.Context, request *ModifyDBInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceReleaseProtection) {
		query["DBInstanceReleaseProtection"] = request.DBInstanceReleaseProtection
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyDBInstanceAttribute"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例配置
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyDBInstanceConfig"),
		Version:     dara.String("2015-12-01"),
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
// Modifies the endpoint that is used to connect to an ApsaraDB for MongoDB instance.
//
// Description:
//
// You can modify the connection strings and ports of the following instances:
//
//   - You can modify the connection strings of instances that use local or cloud disks.
//
//   - You can only modify the ports of instances that use cloud disks.
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
	if !dara.IsNil(request.CurrentConnectionString) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NewConnectionString) {
		query["NewConnectionString"] = request.NewConnectionString
	}

	if !dara.IsNil(request.NewPort) {
		query["NewPort"] = request.NewPort
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyDBInstanceConnectionString"),
		Version:     dara.String("2015-12-01"),
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
// Modifies the name of an ApsaraDB for MongoDB instance.
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyDBInstanceDescription"),
		Version:     dara.String("2015-12-01"),
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
// Modifies the disk type of an ApsaraDB for MongoDB instance.
//
// @param request - ModifyDBInstanceDiskTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceDiskTypeResponse
func (client *Client) ModifyDBInstanceDiskTypeWithContext(ctx context.Context, request *ModifyDBInstanceDiskTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceDiskTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DbInstanceStorageType) {
		query["DbInstanceStorageType"] = request.DbInstanceStorageType
	}

	if !dara.IsNil(request.ExtraParam) {
		query["ExtraParam"] = request.ExtraParam
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.ProvisionedIops) {
		query["ProvisionedIops"] = request.ProvisionedIops
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceDiskType"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceDiskTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window of an ApsaraDB for MongoDB instance.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.MaintainEndTime) {
		query["MaintainEndTime"] = request.MaintainEndTime
	}

	if !dara.IsNil(request.MaintainStartTime) {
		query["MaintainStartTime"] = request.MaintainStartTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyDBInstanceMaintainTime"),
		Version:     dara.String("2015-12-01"),
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
// You can call this operation to set the monitoring granularity for an ApsaraDB for MongoDB instance.
//
// Description:
//
// >  This operation is applicable only to the ApsaraDB for MongoDB console of the previous version due to the change in the feature of adjusting collection intervals of monitoring data.
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - A replica set or sharded cluster instance is used.
//
//   - MongoDB 3.4 (the latest minor version) or MongoDB 4.0 is selected.
//
// @param request - ModifyDBInstanceMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceMonitorResponse
func (client *Client) ModifyDBInstanceMonitorWithContext(ctx context.Context, request *ModifyDBInstanceMonitorRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceMonitorResponse, _err error) {
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

	if !dara.IsNil(request.Granularity) {
		query["Granularity"] = request.Granularity
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyDBInstanceMonitor"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extends the retention period of the classic network endpoint of an ApsaraDB for MongoDB instance.
//
// Description:
//
// Before you call this operation, make sure that the instance meets the following requirements:
//
//   - The instance is in the Running state.
//
//   - The network of the instance is in hybrid access mode.
//
// >  This operation is supported by replica set instances and sharded cluster instances. This operation is not supported by standalone instances.
//
// @param request - ModifyDBInstanceNetExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceNetExpireTimeResponse
func (client *Client) ModifyDBInstanceNetExpireTimeWithContext(ctx context.Context, request *ModifyDBInstanceNetExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceNetExpireTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassicExpendExpiredDays) {
		query["ClassicExpendExpiredDays"] = request.ClassicExpendExpiredDays
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifyDBInstanceNetExpireTime"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceNetExpireTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the network type of an ApsaraDB for MongoDB instance.
//
// Description:
//
// This operation is applicable to replica set instances and sharded cluster instances, but not standalone instances. You can call this operation to change the network of an instance from a classic network to a VPC.
//
// @param request - ModifyDBInstanceNetworkTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceNetworkTypeResponse
func (client *Client) ModifyDBInstanceNetworkTypeWithContext(ctx context.Context, request *ModifyDBInstanceNetworkTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassicExpiredDays) {
		query["ClassicExpiredDays"] = request.ClassicExpiredDays
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RetainClassic) {
		query["RetainClassic"] = request.RetainClassic
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
		Action:      dara.String("ModifyDBInstanceNetworkType"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the SSL settings of an ApsaraDB for MongoDB instance.
//
// Description:
//
// ## Usage
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The instance is in the running state.
//
//   - The instance is a replica set instance.
//
//   - The engine version of the instance is 3.4 or 4.0.
//
// >  When you enable or disable SSL encryption or update the SSL certificate, the instance restarts. We recommend that you call this operation during off-peak hours.
//
// @param request - ModifyDBInstanceSSLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceSSLResponse
func (client *Client) ModifyDBInstanceSSLWithContext(ctx context.Context, request *ModifyDBInstanceSSLRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceSSLResponse, _err error) {
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

	if !dara.IsNil(request.ForceEncryption) {
		query["ForceEncryption"] = request.ForceEncryption
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SSLAction) {
		query["SSLAction"] = request.SSLAction
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceSSL"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the specifications or storage space of an ApsaraDB for MongoDB standalone, replica set, or serverless instance. Serverless instances are available only on the China site (aliyun.com).
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing) of ApsaraDB for MongoDB.
//
// This operation applies only to standalone and replica set instances. To modify the specifications of sharded cluster instances, you can call the [ModifyNodeSpec](https://help.aliyun.com/document_detail/61911.html), [CreateNode](https://help.aliyun.com/document_detail/61922.html), [DeleteNode](https://help.aliyun.com/document_detail/61816.html), or [ModifyNodeSpecBatch](https://help.aliyun.com/document_detail/61923.html) operation.
//
// @param request - ModifyDBInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceSpecResponse
func (client *Client) ModifyDBInstanceSpecWithContext(ctx context.Context, request *ModifyDBInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceStorage) {
		query["DBInstanceStorage"] = request.DBInstanceStorage
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.ExtraParam) {
		query["ExtraParam"] = request.ExtraParam
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReadonlyReplicas) {
		query["ReadonlyReplicas"] = request.ReadonlyReplicas
	}

	if !dara.IsNil(request.ReplicationFactor) {
		query["ReplicationFactor"] = request.ReplicationFactor
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SearchNodeClass) {
		query["SearchNodeClass"] = request.SearchNodeClass
	}

	if !dara.IsNil(request.SearchNodeCount) {
		query["SearchNodeCount"] = request.SearchNodeCount
	}

	if !dara.IsNil(request.SearchNodeStorage) {
		query["SearchNodeStorage"] = request.SearchNodeStorage
	}

	if !dara.IsNil(request.TargetHiddenZoneId) {
		query["TargetHiddenZoneId"] = request.TargetHiddenZoneId
	}

	if !dara.IsNil(request.TargetSecondaryZoneId) {
		query["TargetSecondaryZoneId"] = request.TargetSecondaryZoneId
	}

	if !dara.IsNil(request.TargetVswitchId) {
		query["TargetVswitchId"] = request.TargetVswitchId
	}

	if !dara.IsNil(request.TargetZoneId) {
		query["TargetZoneId"] = request.TargetZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceSpec"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the transparent data encryption (TDE) status of an ApsaraDB for MongoDB instance.
//
// Description:
//
// TDE allows you to perform real-time I/O encryption and decryption on data files. Data is encrypted before it is written to a disk and is decrypted when it is read from the disk to the memory. For more information, see [Configure TDE](https://help.aliyun.com/document_detail/131048.html).
//
// >  TDE cannot be disabled after it is enabled.
//
// Before you call this operation, make sure that the ApsaraDB for MongoDB instance meets the following requirements:
//
//   - A replica set or sharded cluster instance is used.
//
//   - The storage engine of the instance is WiredTiger.
//
//   - The instance uses local disks to store data.
//
//   - The database engine version of the instance is 4.0 or 4.2. If the database engine version is earlier than 4.0, you can call the [UpgradeDBInstanceEngineVersion](https://help.aliyun.com/document_detail/67608.html) operation to upgrade the database engine.
//
// @param request - ModifyDBInstanceTDERequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceTDEResponse
func (client *Client) ModifyDBInstanceTDEWithContext(ctx context.Context, request *ModifyDBInstanceTDERequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceTDEResponse, _err error) {
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

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.EncryptorName) {
		query["EncryptorName"] = request.EncryptorName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleARN) {
		query["RoleARN"] = request.RoleARN
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	if !dara.IsNil(request.TDEStatus) {
		query["TDEStatus"] = request.TDEStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceTDE"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceTDEResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the global IP whitelist template associated with an ApsaraDB for MongoDB instance.
//
// @param request - ModifyGlobalSecurityIPGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalSecurityIPGroupResponse
func (client *Client) ModifyGlobalSecurityIPGroupWithContext(ctx context.Context, request *ModifyGlobalSecurityIPGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalSecurityIPGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GIpList) {
		query["GIpList"] = request.GIpList
	}

	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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
		Action:      dara.String("ModifyGlobalSecurityIPGroup"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalSecurityIPGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a global IP whitelist template associated with an ApsaraDB for MongoDB instance.
//
// @param request - ModifyGlobalSecurityIPGroupNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalSecurityIPGroupNameResponse
func (client *Client) ModifyGlobalSecurityIPGroupNameWithContext(ctx context.Context, request *ModifyGlobalSecurityIPGroupNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalSecurityIPGroupNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GlobalIgName) {
		query["GlobalIgName"] = request.GlobalIgName
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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
		Action:      dara.String("ModifyGlobalSecurityIPGroupName"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalSecurityIPGroupNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the mapping between a global whitelist template and an ApsaraDB for MongoDB instance.
//
// @param request - ModifyGlobalSecurityIPGroupRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGlobalSecurityIPGroupRelationResponse
func (client *Client) ModifyGlobalSecurityIPGroupRelationWithContext(ctx context.Context, request *ModifyGlobalSecurityIPGroupRelationRequest, runtime *dara.RuntimeOptions) (_result *ModifyGlobalSecurityIPGroupRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.GlobalSecurityGroupId) {
		query["GlobalSecurityGroupId"] = request.GlobalSecurityGroupId
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
		Action:      dara.String("ModifyGlobalSecurityIPGroupRelation"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyGlobalSecurityIPGroupRelationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables auto-renewal for an ApsaraDB for MongoDB instance.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing).
//
// This operation is applicable to subscription instances.
//
// >  When auto-renewal is enabled, your payment will be collected nine days before the expiration date of ApsaraDB for MongoDB. Ensure that your account has sufficient balance.
//
// @param request - ModifyInstanceAutoRenewalAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAutoRenewalAttributeResponse
func (client *Client) ModifyInstanceAutoRenewalAttributeWithContext(ctx context.Context, request *ModifyInstanceAutoRenewalAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAutoRenewalAttributeResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
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
		Action:      dara.String("ModifyInstanceAutoRenewalAttribute"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAutoRenewalAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables password-free access over Virtual Private Cloud (VPC) for an ApsaraDB for MongoDB instance.
//
// Description:
//
// Before you call this operation, make sure that the ApsaraDB for MongoDB instance meets the following requirements:
//
//   - The instance is a replica set or sharded cluster instance.
//
//   - The database engine version of the instance is 4.0 (with the minor version of mongodb_20190408_3.0.11 or later) or 4.2. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/62010.html) operation to view the database engine version of the instance. If necessary, you can call the [UpgradeDBInstanceEngineVersion](https://help.aliyun.com/document_detail/67608.html) operation to upgrade the database engine version of the instance.
//
//   - The network type of the instance must be VPC. If the network type of the instance is classic network, you must call the [ModifyDBInstanceNetworkType](https://help.aliyun.com/document_detail/62138.html) operation to change the network type to VPC.
//
//   - You can only disable but not enable password-free access over VPC.
//
// @param request - ModifyInstanceVpcAuthModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceVpcAuthModeResponse
func (client *Client) ModifyInstanceVpcAuthModeWithContext(ctx context.Context, request *ModifyInstanceVpcAuthModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceVpcAuthModeResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VpcAuthMode) {
		query["VpcAuthMode"] = request.VpcAuthMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceVpcAuthMode"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceVpcAuthModeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the specifications and storage capacity of a node of an ApsaraDB for MongoDB sharded cluster instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing) of ApsaraDB for MongoDB.
//
// > This operation is applicable only to sharded cluster instances.
//
// @param request - ModifyNodeSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodeSpecResponse
func (client *Client) ModifyNodeSpecWithContext(ctx context.Context, request *ModifyNodeSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyNodeSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.FromApp) {
		query["FromApp"] = request.FromApp
	}

	if !dara.IsNil(request.NodeClass) {
		query["NodeClass"] = request.NodeClass
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeStorage) {
		query["NodeStorage"] = request.NodeStorage
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReadonlyReplicas) {
		query["ReadonlyReplicas"] = request.ReadonlyReplicas
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.TargetHiddenZoneId) {
		query["TargetHiddenZoneId"] = request.TargetHiddenZoneId
	}

	if !dara.IsNil(request.TargetSecondaryZoneId) {
		query["TargetSecondaryZoneId"] = request.TargetSecondaryZoneId
	}

	if !dara.IsNil(request.TargetVswitchId) {
		query["TargetVswitchId"] = request.TargetVswitchId
	}

	if !dara.IsNil(request.TargetZoneId) {
		query["TargetZoneId"] = request.TargetZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNodeSpec"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNodeSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the configurations of mongos or shard nodes in an ApsaraDB for MongoDB sharded cluster instance.
//
// Description:
//
// Make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing) of ApsaraDB for MongoDB before you call this operation.
//
// This operation is applicable only to sharded cluster instances.
//
// When you upgrade or downgrade the configurations of multiple sharded cluster instances in batches, the specifications of the instances are limited. For example, if you want to expand the storage capacity of the instances, the storage capacity of the instances after expansion must be greater than the current capacity. When the specifications of multiple sharded cluster instances are different, limits are defined based on the specifications of a random sharded cluster instance. In this case, you may be unable to upgrade or downgrade the configurations of the instances. In this case, we recommend that you call the ModifyNodeSpec operation to individually change the configurations of each sharded cluster instance.
//
// @param request - ModifyNodeSpecBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodeSpecBatchResponse
func (client *Client) ModifyNodeSpecBatchWithContext(ctx context.Context, request *ModifyNodeSpecBatchRequest, runtime *dara.RuntimeOptions) (_result *ModifyNodeSpecBatchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.NodesInfo) {
		query["NodesInfo"] = request.NodesInfo
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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

	if !dara.IsNil(request.TargetHiddenZoneId) {
		query["TargetHiddenZoneId"] = request.TargetHiddenZoneId
	}

	if !dara.IsNil(request.TargetSecondaryZoneId) {
		query["TargetSecondaryZoneId"] = request.TargetSecondaryZoneId
	}

	if !dara.IsNil(request.TargetVswitchId) {
		query["TargetVswitchId"] = request.TargetVswitchId
	}

	if !dara.IsNil(request.TargetZoneId) {
		query["TargetZoneId"] = request.TargetZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNodeSpecBatch"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNodeSpecBatchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameters of an ApsaraDB for MongoDB instance.
//
// Description:
//
//	  The instance must be in the Running state when you call this operation.
//
//		- If you call this operation to modify specific instance parameters and the modification for part of the parameters can take effect only after an instance restart, the instance is automatically restarted after this operation is called. You can call the [DescribeParameterTemplates](https://help.aliyun.com/document_detail/67618.html) operation to query the parameters that take effect only after the instance is restarted.
//
// @param request - ModifyParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParametersResponse
func (client *Client) ModifyParametersWithContext(ctx context.Context, request *ModifyParametersRequest, runtime *dara.RuntimeOptions) (_result *ModifyParametersResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
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

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyParameters"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyParametersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves an ApsaraDB for MongoDB instance to a specified resource group.
//
// Description:
//
// Resource Management allows you to build an organizational structure for resources based on your business requirements. You can use resource directories, folders, accounts, and resource groups to hierarchically organize and manage resources. For more information, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html)
//
// @param request - ModifyResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyResourceGroupResponse
func (client *Client) ModifyResourceGroupWithContext(ctx context.Context, request *ModifyResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyResourceGroupResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("ModifyResourceGroup"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyResourceGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to modify an ECS Security group that is bound to an ApsaraDB for MongoDB instance.
//
// Description:
//
// >  For a sharded cluster instance, the bound ECS security group takes effect only for mongos nodes.
//
// @param request - ModifySecurityGroupConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityGroupConfigurationResponse
func (client *Client) ModifySecurityGroupConfigurationWithContext(ctx context.Context, request *ModifySecurityGroupConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityGroupConfigurationResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityGroupConfiguration"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecurityGroupConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP address whitelist of an ApsaraDB for MongoDB instance.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityIpGroupAttribute) {
		query["SecurityIpGroupAttribute"] = request.SecurityIpGroupAttribute
	}

	if !dara.IsNil(request.SecurityIpGroupName) {
		query["SecurityIpGroupName"] = request.SecurityIpGroupName
	}

	if !dara.IsNil(request.SecurityIps) {
		query["SecurityIps"] = request.SecurityIps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityIps"),
		Version:     dara.String("2015-12-01"),
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
// 修改MongoDB实例的SRV连接地址
//
// @param request - ModifySrvNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySrvNetworkAddressResponse
func (client *Client) ModifySrvNetworkAddressWithContext(ctx context.Context, request *ModifySrvNetworkAddressRequest, runtime *dara.RuntimeOptions) (_result *ModifySrvNetworkAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionType) {
		query["ConnectionType"] = request.ConnectionType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NewConnectionString) {
		query["NewConnectionString"] = request.NewConnectionString
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ModifySrvNetworkAddress"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySrvNetworkAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a task in the task center for an ApsaraDB for MongoDB instance.
//
// Description:
//
// The actions performed by this operation for a task vary based on the current state of the task. The supported actions for a task can be obtained from the value of the actionInfo parameter in the DescribeHistoryTasks operation.
//
// @param request - ModifyTaskInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTaskInfoResponse
func (client *Client) ModifyTaskInfoWithContext(ctx context.Context, request *ModifyTaskInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyTaskInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionParams) {
		query["ActionParams"] = request.ActionParams
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

	if !dara.IsNil(request.StepName) {
		query["StepName"] = request.StepName
	}

	if !dara.IsNil(request.TaskAction) {
		query["TaskAction"] = request.TaskAction
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTaskInfo"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the internal endpoint of a shard or Configserver node in a sharded cluster instance.
//
// Description:
//
//	  This operation can be used to release the internal endpoint of a shard or Configserver node in a sharded cluster instance. For more information, see [Release the endpoint of a shard or Configserver node](https://help.aliyun.com/document_detail/134067.html).
//
//		- To release the public endpoint of a shard or Configserver node in a sharded cluster instance, you can call the [ReleasePublicNetworkAddress](https://help.aliyun.com/document_detail/67604.html) operation.
//
// @param request - ReleaseNodePrivateNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseNodePrivateNetworkAddressResponse
func (client *Client) ReleaseNodePrivateNetworkAddressWithContext(ctx context.Context, request *ReleaseNodePrivateNetworkAddressRequest, runtime *dara.RuntimeOptions) (_result *ReleaseNodePrivateNetworkAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionType) {
		query["ConnectionType"] = request.ConnectionType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ReleaseNodePrivateNetworkAddress"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseNodePrivateNetworkAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an ApsaraDB for MongoDB instance.
//
// @param request - ReleasePublicNetworkAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleasePublicNetworkAddressResponse
func (client *Client) ReleasePublicNetworkAddressWithContext(ctx context.Context, request *ReleasePublicNetworkAddressRequest, runtime *dara.RuntimeOptions) (_result *ReleasePublicNetworkAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionType) {
		query["ConnectionType"] = request.ConnectionType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ReleasePublicNetworkAddress"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleasePublicNetworkAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually renews an ApsaraDB for MongoDB subscription instance.
//
// Description:
//
// Make sure that you fully understand the billing methods and pricing of ApsaraDB for MongoDB before you call this operation. For more information about the pricing of ApsaraDB for MongoDB, visit the [pricing tab of the product buy page](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing).
//
// This operation is only applicable to instances that use the subscription billing method.
//
// @param request - RenewDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewDBInstanceResponse
func (client *Client) RenewDBInstanceWithContext(ctx context.Context, request *RenewDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
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
		Action:      dara.String("RenewDBInstance"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of the root account in an ApsaraDB for MongoDB instance.
//
// Description:
//
// >  This operation can be used to reset only the password of the root account of an instance.
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

	if !dara.IsNil(request.CharacterType) {
		query["CharacterType"] = request.CharacterType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ResetAccountPassword"),
		Version:     dara.String("2015-12-01"),
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
// Restarts an ApsaraDB for MongoDB instance.
//
// Description:
//
// This operation can also be used to restart an instance, or restart a shard or mongos node in a sharded cluster instance.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDBInstance"),
		Version:     dara.String("2015-12-01"),
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
// Restarts a node in an ApsaraDB for MongoDB instance.
//
// Description:
//
// You can call this operation to restart a node in a replica set instance or a child instance in a sharded cluster instance.
//
// >  When you call this operation, the instance must meet the following requirements:
//
//   - The instance is in the Running state.
//
//   - The instance is a replica set or sharded cluster instance of the standard edition.
//
// @param request - RestartNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartNodeResponse
func (client *Client) RestartNodeWithContext(ctx context.Context, request *RestartNodeRequest, runtime *dara.RuntimeOptions) (_result *RestartNodeResponse, _err error) {
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

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartNode"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the primary and secondary nodes for an ApsaraDB for MongoDB instance.
//
// Description:
//
// The instance must be running when you call this operation.
//
// >
//
//   - This operation is applicable to replica set instances and sharded cluster instances, but cannot be performed on standalone instances.
//
//   - On replica set instances, the switch is performed between instances. On sharded cluster instances, the switch is performed between shards.
//
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleIds) {
		query["RoleIds"] = request.RoleIds
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchDBInstanceHA"),
		Version:     dara.String("2015-12-01"),
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
// Binds tags to ApsaraDB for MongoDB instances.
//
// Description:
//
// If you have a large number of instances, you can create multiple tags, bind the tags to the instances, and filter the instances by tag.
//
//   - A tag consists of a key and a value. Each key must be unique in a region for an Alibaba Cloud account. Different keys can be mapped to the same value.
//
//   - If the tag that you specify does not exist, this tag is automatically created and bound to the specified instance.
//
//   - If a tag that has the same key is already bound to the instance, the new tag overwrites the existing tag.
//
//   - You can bind up to 20 tags to each instance.
//
//   - You can bind tags to up to 50 instances each time you call the operation.
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
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2015-12-01"),
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
// Switches the backup mode of an ApsaraDB for MongoDB sharded cluster instance to the cluster backup mode. After the instance is switched to the cluster backup mode, the instance supports high-frequency backup.
//
// Description:
//
//	  The instance is an ApsaraDB for MongoDB sharded cluster instance that runs MongoDB 4.4 or later and uses enhanced SSDs (ESSDs) to store data.
//
//		- You can call the TransferClusterBackup operation only for instances that are created before October 19, 2023 to switch the instances to the cluster backup mode. Cloud disk-based sharded cluster instances that are created on or after October 19, 2023 are set to the cluster backup mode by default.
//
// @param request - TransferClusterBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferClusterBackupResponse
func (client *Client) TransferClusterBackupWithContext(ctx context.Context, request *TransferClusterBackupRequest, runtime *dara.RuntimeOptions) (_result *TransferClusterBackupResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("TransferClusterBackup"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferClusterBackupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of an instance from pay-as-you-go to subscription or from subscription to pay-as-you-go.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/product/apsaradb-for-mongodb/pricing) of ApsaraDB for MongoDB
//
// Before you call this API operation, make sure that the ApsaraDB for MongoDB instance meets the following requirements:
//
//   - The instance is in the Running state.
//
//   - Your instance has no unpaid billing method change orders.
//
//   - The instance type is available for purchase. For more information about unavailable instance types, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
//
// > To change the billing method of an instance whose instance type is no longer available to purchase, call the [ModifyDBInstanceSpec](https://help.aliyun.com/document_detail/61816.html) or [ModifyNodeSpec](https://help.aliyun.com/document_detail/61923.html) operation to change the instance type first.
//
// @param request - TransformInstanceChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformInstanceChargeTypeResponse
func (client *Client) TransformInstanceChargeTypeWithContext(ctx context.Context, request *TransformInstanceChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *TransformInstanceChargeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
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
		Action:      dara.String("TransformInstanceChargeType"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransformInstanceChargeTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of an ApsaraDB for MongoDB instance from pay-as-you-go to subscription.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/zh/product/apsaradb-for-mongodb/pricing).
//
// A subscription instance cannot be changed to a pay-as-you-go instance. To avoid wasting resources, proceed with caution.
//
// Before you call this API operation, make sure that the ApsaraDB for MongoDB instance meets the following requirements:
//
//   - The instance is in the running state.
//
//   - The billing method of the instance is pay-as-you-go.
//
//   - The instance has no unpaid subscription orders.
//
//   - The instance type is available for purchase. For more information about unavailable instance types, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
//
// >  To change the billing method of an instance whose instance type is no longer available to subscription, call the [ModifyDBInstanceSpec](https://help.aliyun.com/document_detail/61816.html) or [ModifyNodeSpec](https://help.aliyun.com/document_detail/61923.html) operation to first change the instance type.
//
// @param request - TransformToPrePaidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformToPrePaidResponse
func (client *Client) TransformToPrePaidWithContext(ctx context.Context, request *TransformToPrePaidRequest, runtime *dara.RuntimeOptions) (_result *TransformToPrePaidResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.CouponNo) {
		query["CouponNo"] = request.CouponNo
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
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
		Action:      dara.String("TransformToPrePaid"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransformToPrePaidResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a tag if the tag is not added to another instance.
//
// Description:
//
// >
//
//   - You can remove up to 20 tags at a time.
//
//   - If you remove a tag from all instances, the tag is automatically deleted.
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2015-12-01"),
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
// Upgrades the database version of an ApsaraDB for MongoDB instance.
//
// Description:
//
// The instance must be in the running state when you call this operation.
//
// > 	- The available database versions depend on the storage engine used by the instance. For more information, see [Upgrades of MongoDB major versions](https://help.aliyun.com/document_detail/398673.html). You can also call the [DescribeAvailableEngineVersion](https://help.aliyun.com/document_detail/141355.html) operation to query the available database versions.
//
// > 	- You cannot downgrade the MongoDB version of an instance after you upgrade it.
//
// > 	- The instance is automatically restarted for two to three times during the upgrade process. Make sure that you upgrade the instance during off-peak hours.
//
// @param request - UpgradeDBInstanceEngineVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceEngineVersionResponse
func (client *Client) UpgradeDBInstanceEngineVersionWithContext(ctx context.Context, request *UpgradeDBInstanceEngineVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBInstanceEngineVersionResponse, _err error) {
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceEngineVersion"),
		Version:     dara.String("2015-12-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBInstanceEngineVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the minor version of an ApsaraDB for MongoDB instance.
//
// Description:
//
// When you call the UpgradeDBInstanceKernelVersion operation, the instance must be in the Running state.
//
// > 	- The UpgradeDBInstanceKernelVersion operation is applicable to replica set and sharded cluster instances, but not to standalone instances.
//
// > 	- The instance will be restarted once during the upgrade. Call this operation during off-peak hours.
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SwitchMode) {
		query["SwitchMode"] = request.SwitchMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceKernelVersion"),
		Version:     dara.String("2015-12-01"),
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
