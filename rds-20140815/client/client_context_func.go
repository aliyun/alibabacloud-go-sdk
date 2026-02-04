// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Accepts the default operation for a system event in the Inquiring state and authorizes the system to perform the default operation.
//
// @param request - AcceptRCInquiredSystemEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptRCInquiredSystemEventResponse
func (client *Client) AcceptRCInquiredSystemEventWithContext(ctx context.Context, request *AcceptRCInquiredSystemEventRequest, runtime *dara.RuntimeOptions) (_result *AcceptRCInquiredSystemEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptRCInquiredSystemEvent"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptRCInquiredSystemEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches workloads over from the source PostgreSQL instance to the destination ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Use the cloud migration feature](https://help.aliyun.com/document_detail/365562.html)
//
// @param request - ActivateMigrationTargetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateMigrationTargetInstanceResponse
func (client *Client) ActivateMigrationTargetInstanceWithContext(ctx context.Context, request *ActivateMigrationTargetInstanceRequest, runtime *dara.RuntimeOptions) (_result *ActivateMigrationTargetInstanceResponse, _err error) {
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

	if !dara.IsNil(request.ForceSwitch) {
		query["ForceSwitch"] = request.ForceSwitch
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      dara.String("ActivateMigrationTargetInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateMigrationTargetInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)Usage notes
//
//   - Each tag consists of a tag key and a tag value. The tag key is required, and the tag value is optional.
//
//   - The tag key and tag value cannot start with aliyun.
//
//   - The tag key and tag value are not case-sensitive.
//
//   - The maximum length of a tag key is 64 characters, and the maximum length of a tag value is 128 characters.
//
//   - A maximum of 10 tags can be added to each instance. Each tag that is added to the same instance must have a unique tag key. If you add a new tag to the instance and the key of the new tag is the same as that of an existing tag, the new tag overwrites the existing tag.
//
// @param request - AddTagsToResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTagsToResourceResponse
func (client *Client) AddTagsToResourceWithContext(ctx context.Context, request *AddTagsToResourceRequest, runtime *dara.RuntimeOptions) (_result *AddTagsToResourceResponse, _err error) {
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

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.ProxyId) {
		query["proxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTagsToResource"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTagsToResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Apply for a public endpoint for an ApsaraDB RDS instance
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Apply for a public endpoint for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/26128.html)
//
//   - [Apply for a public endpoint for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/97738.html)
//
//   - [Apply for a public endpoint for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/97736.html)
//
//   - [Apply for a public endpoint for an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97740.html)
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
	if !dara.IsNil(request.BabelfishPort) {
		query["BabelfishPort"] = request.BabelfishPort
	}

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.GeneralGroupName) {
		query["GeneralGroupName"] = request.GeneralGroupName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PGBouncerPort) {
		query["PGBouncerPort"] = request.PGBouncerPort
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
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
		Version:     dara.String("2014-08-15"),
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
// Applies for a read-only routing endpoint for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### [](#)Feature description
//
// If read-only instances are attached to a primary ApsaraDB RDS for SQL Server instance, you can call this operation to apply for a unified read-only routing endpoint for the primary instance. After you apply for a read-only routing endpoint for a primary instance, the existing endpoints of the primary instance and its read-only instances remain valid. In addition, you can still apply for internal and public endpoints.
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The shared proxy feature is enabled for your ApsaraDB RDS for MySQL instance.
//
//   - The instance is in the Running state.
//
//   - Read-only instances are attached to the primary instance.
//
//   - The instance does not have an ongoing Data Transmission Service (DTS) migration task.
//
//   - The instance runs one of the following database versions and RDS editions:
//
//   - SQL Server on RDS Cluster Edition
//
//   - MySQL 5.7 on RDS High-availability Edition (with local disks)
//
//   - MySQL 5.6
//
// @param request - AllocateReadWriteSplittingConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateReadWriteSplittingConnectionResponse
func (client *Client) AllocateReadWriteSplittingConnectionWithContext(ctx context.Context, request *AllocateReadWriteSplittingConnectionRequest, runtime *dara.RuntimeOptions) (_result *AllocateReadWriteSplittingConnectionResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DistributionType) {
		query["DistributionType"] = request.DistributionType
	}

	if !dara.IsNil(request.MaxDelayTime) {
		query["MaxDelayTime"] = request.MaxDelayTime
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AllocateReadWriteSplittingConnection"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateReadWriteSplittingConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an elastic IP address (EIP) with an RDS Custom instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// [Introduction to ApsaraDB RDS Custom](https://help.aliyun.com/document_detail/2864363.html)
//
// ### [](#)Precautions
//
// If the RDS Custom instance uses a public IP address, the public IP address is automatically released after you associate an EIP with the instance.
//
// @param request - AssociateEipAddressWithRCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateEipAddressWithRCInstanceResponse
func (client *Client) AssociateEipAddressWithRCInstanceWithContext(ctx context.Context, request *AssociateEipAddressWithRCInstanceRequest, runtime *dara.RuntimeOptions) (_result *AssociateEipAddressWithRCInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllocationId) {
		query["AllocationId"] = request.AllocationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateEipAddressWithRCInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateEipAddressWithRCInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a pay-as-you-go data disk or a system disk to an RDS Custom instance. The instance and the disk must reside in the same zone.
//
// @param request - AttachRCDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachRCDiskResponse
func (client *Client) AttachRCDiskWithContext(ctx context.Context, request *AttachRCDiskRequest, runtime *dara.RuntimeOptions) (_result *AttachRCDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteWithInstance) {
		query["DeleteWithInstance"] = request.DeleteWithInstance
	}

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachRCDisk"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachRCDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds RDS Custom nodes to a Container Service for Kubernetes (ACK) cluster.
//
// @param tmpReq - AttachRCInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachRCInstancesResponse
func (client *Client) AttachRCInstancesWithContext(ctx context.Context, tmpReq *AttachRCInstancesRequest, runtime *dara.RuntimeOptions) (_result *AttachRCInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachRCInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.KeyPair) {
		query["KeyPair"] = request.KeyPair
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachRCInstances"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachRCInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a whitelist template with an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// @param request - AttachWhitelistTemplateToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachWhitelistTemplateToInstanceResponse
func (client *Client) AttachWhitelistTemplateToInstanceWithContext(ctx context.Context, request *AttachWhitelistTemplateToInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachWhitelistTemplateToInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InsName) {
		query["InsName"] = request.InsName
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachWhitelistTemplateToInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachWhitelistTemplateToInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds rules to the specified security group.
//
// @param tmpReq - AuthorizeRCSecurityGroupPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeRCSecurityGroupPermissionResponse
func (client *Client) AuthorizeRCSecurityGroupPermissionWithContext(ctx context.Context, tmpReq *AuthorizeRCSecurityGroupPermissionRequest, runtime *dara.RuntimeOptions) (_result *AuthorizeRCSecurityGroupPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AuthorizeRCSecurityGroupPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SecurityGroupPermissions) {
		request.SecurityGroupPermissionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityGroupPermissions, dara.String("SecurityGroupPermissions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityGroupPermissionsShrink) {
		query["SecurityGroupPermissions"] = request.SecurityGroupPermissionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeRCSecurityGroupPermission"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeRCSecurityGroupPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries system-assigned read weights.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### [](#)Feature description
//
// When the [read/write splitting](https://help.aliyun.com/document_detail/51073.html) feature is enabled, this operation is used to calculate system-assigned read weights. For more information about custom read weights, see [DescribeDBInstanceNetInfo](https://help.aliyun.com/document_detail/610423.html).
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - If the instance runs MySQL, the instance uses a shared proxy.
//
//   - The instance runs one of the following MySQL versions and RDS editions:
//
//   - MySQL 5.7 on RDS High-availability Edition (with local disks)
//
//   - MySQL 5.6
//
//   - SQL Server on RDS Cluster Edition
//
// @param request - CalculateDBInstanceWeightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CalculateDBInstanceWeightResponse
func (client *Client) CalculateDBInstanceWeightWithContext(ctx context.Context, request *CalculateDBInstanceWeightRequest, runtime *dara.RuntimeOptions) (_result *CalculateDBInstanceWeightResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CalculateDBInstanceWeight"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CalculateDBInstanceWeightResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels O\\&M tasks that are not started.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Scheduled events for ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/104183.html)
//
//   - [Scheduled events for ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/104452.html)
//
//   - [Scheduled events for ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/104451.html)
//
//   - [Scheduled events for ApsaraDB RDS for MariaDB instances](https://help.aliyun.com/document_detail/104454.html)
//
// ### [](#)Usage notes
//
// Tasks cannot be canceled in the following situations:
//
//   - The allowCancel parameter is set to 0.
//
//   - The current time is later than the task start time.
//
//   - The status of the task is not set to 3. The value 3 specifies that the task is waiting to be executed.
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
		Action:      dara.String("CancelActiveOperationTasks"),
		Version:     dara.String("2014-08-15"),
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
// Checks whether the username of the account that you want to create on an instance is available.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - CheckAccountNameAvailableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccountNameAvailableResponse
func (client *Client) CheckAccountNameAvailableWithContext(ctx context.Context, request *CheckAccountNameAvailableRequest, runtime *dara.RuntimeOptions) (_result *CheckAccountNameAvailableResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAccountNameAvailable"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAccountNameAvailableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks permissions that are granted on an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TargetRegionId) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCloudResourceAuthorized"),
		Version:     dara.String("2014-08-15"),
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
// Queries whether an instance can be restored by using a cross-region backup set.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// >  If your ApsaraDB RDS for PostgreSQL instance uses the new architecture and is created after October 10, 2022, this feature is not supported for the RDS instance. For more information, see [[Notice\\] SLR authorization is required to create an ApsaraDB RDS for PostgreSQL instance from October 10, 2022](~~452313~~).
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use the cross-region backup feature of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/120824.html) and [Restore the data of an ApsaraDB RDS for MySQL instance across regions](https://help.aliyun.com/document_detail/120875.html)
//
//   - [Use the cross-region backup feature of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206671.html) and [Restore the data of an ApsaraDB RDS for PostgreSQL across regions](https://help.aliyun.com/document_detail/206662.html)
//
//   - [Use the cross-region backup feature of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/187923.html) and [Restore the data of an ApsaraDB RDS for SQL Server across regions](https://help.aliyun.com/document_detail/187924.html)
//
// @param request - CheckCreateDdrDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCreateDdrDBInstanceResponse
func (client *Client) CheckCreateDdrDBInstanceWithContext(ctx context.Context, request *CheckCreateDdrDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CheckCreateDdrDBInstanceResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceStorage) {
		query["DBInstanceStorage"] = request.DBInstanceStorage
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
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

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	if !dara.IsNil(request.SourceDBInstanceName) {
		query["SourceDBInstanceName"] = request.SourceDBInstanceName
	}

	if !dara.IsNil(request.SourceRegion) {
		query["SourceRegion"] = request.SourceRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCreateDdrDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCreateDdrDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a database name is unique and conforms to the naming conventions on an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// @param request - CheckDBNameAvailableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDBNameAvailableResponse
func (client *Client) CheckDBNameAvailableWithContext(ctx context.Context, request *CheckDBNameAvailableRequest, runtime *dara.RuntimeOptions) (_result *CheckDBNameAvailableResponse, _err error) {
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

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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
		Action:      dara.String("CheckDBNameAvailable"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDBNameAvailableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the CheckInstanceExist operation to query whether an ApsaraDB RDS instance exists.
//
// @param request - CheckInstanceExistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInstanceExistResponse
func (client *Client) CheckInstanceExistWithContext(ctx context.Context, request *CheckInstanceExistRequest, runtime *dara.RuntimeOptions) (_result *CheckInstanceExistResponse, _err error) {
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
		Action:      dara.String("CheckInstanceExist"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckInstanceExistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看是否已创建服务关联角色（SLR）和是否开租
//
// @param request - CheckRdsCustomInitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckRdsCustomInitResponse
func (client *Client) CheckRdsCustomInitWithContext(ctx context.Context, request *CheckRdsCustomInitRequest, runtime *dara.RuntimeOptions) (_result *CheckRdsCustomInitResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceLinkedRole) {
		query["ServiceLinkedRole"] = request.ServiceLinkedRole
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckRdsCustomInit"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckRdsCustomInitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a service-linked role (SLR) is created.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
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

	if !dara.IsNil(request.ServiceLinkedRole) {
		query["ServiceLinkedRole"] = request.ServiceLinkedRole
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckServiceLinkedRole"),
		Version:     dara.String("2014-08-15"),
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
// Restores the data of an original instance to a new instance. The new instance is called a cloned instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Restore data of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96147.html)
//
//   - [Restore data of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96776.html)
//
//   - [Restore data of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95722.html)
//
//   - [Restore data of an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97151.html)
//
// @param tmpReq - CloneDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneDBInstanceResponse
func (client *Client) CloneDBInstanceWithContext(ctx context.Context, tmpReq *CloneDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CloneDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CloneDBInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServerlessConfig) {
		request.ServerlessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfig, dara.String("ServerlessConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BackupId) {
		query["BackupId"] = request.BackupId
	}

	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.BpeEnabled) {
		query["BpeEnabled"] = request.BpeEnabled
	}

	if !dara.IsNil(request.BurstingEnabled) {
		query["BurstingEnabled"] = request.BurstingEnabled
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CustomExtraInfo) {
		query["CustomExtraInfo"] = request.CustomExtraInfo
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

	if !dara.IsNil(request.DBInstanceStorage) {
		query["DBInstanceStorage"] = request.DBInstanceStorage
	}

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.DbNames) {
		query["DbNames"] = request.DbNames
	}

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.IoAccelerationEnabled) {
		query["IoAccelerationEnabled"] = request.IoAccelerationEnabled
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTable) {
		query["RestoreTable"] = request.RestoreTable
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.ServerlessConfigShrink) {
		query["ServerlessConfig"] = request.ServerlessConfigShrink
	}

	if !dara.IsNil(request.TableMeta) {
		query["TableMeta"] = request.TableMeta
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

	if !dara.IsNil(request.ZoneIdSlave1) {
		query["ZoneIdSlave1"] = request.ZoneIdSlave1
	}

	if !dara.IsNil(request.ZoneIdSlave2) {
		query["ZoneIdSlave2"] = request.ZoneIdSlave2
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replicates a parameter template to the current region or another region.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/130565.html)
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/457176.html)
//
// @param request - CloneParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneParameterGroupResponse
func (client *Client) CloneParameterGroupWithContext(ctx context.Context, request *CloneParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *CloneParameterGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupDesc) {
		query["ParameterGroupDesc"] = request.ParameterGroupDesc
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.ParameterGroupName) {
		query["ParameterGroupName"] = request.ParameterGroupName
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

	if !dara.IsNil(request.TargetRegionId) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloneParameterGroup"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloneParameterGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Marks the notifications of an instance within your Alibaba Cloud account as confirmed.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)Feature description
//
// After you call the [QueryNotify](https://help.aliyun.com/document_detail/610443.html) operation to query notifications for an instance, you can call this operation to mark the notifications as confirmed.
//
// @param tmpReq - ConfirmNotifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmNotifyResponse
func (client *Client) ConfirmNotifyWithContext(ctx context.Context, tmpReq *ConfirmNotifyRequest, runtime *dara.RuntimeOptions) (_result *ConfirmNotifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ConfirmNotifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotifyIdList) {
		request.NotifyIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotifyIdList, dara.String("NotifyIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Confirmor) {
		body["Confirmor"] = request.Confirmor
	}

	if !dara.IsNil(request.NotifyIdListShrink) {
		body["NotifyIdList"] = request.NotifyIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfirmNotify"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfirmNotifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replicates the databases of an instance that runs SQL Server 2008 R2 to another instance. This operation is phased out.
//
// Description:
//
// This operation is phased out.
//
// @param request - CopyDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDatabaseResponse
func (client *Client) CopyDatabaseWithContext(ctx context.Context, request *CopyDatabaseRequest, runtime *dara.RuntimeOptions) (_result *CopyDatabaseResponse, _err error) {
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

	if !dara.IsNil(request.DstDBName) {
		query["DstDBName"] = request.DstDBName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReserveAccount) {
		query["ReserveAccount"] = request.ReserveAccount
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

	if !dara.IsNil(request.SrcDBName) {
		query["SrcDBName"] = request.SrcDBName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyDatabase"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replicates databases between RDS SQL Server instances.
//
// Description:
//
// ### Supported database engines
//
// # RDS SQL Server
//
// ### References
//
// > : Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Replicate databases between ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/95702.html)
//
// @param request - CopyDatabaseBetweenInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDatabaseBetweenInstancesResponse
func (client *Client) CopyDatabaseBetweenInstancesWithContext(ctx context.Context, request *CopyDatabaseBetweenInstancesRequest, runtime *dara.RuntimeOptions) (_result *CopyDatabaseBetweenInstancesResponse, _err error) {
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

	if !dara.IsNil(request.DbNames) {
		query["DbNames"] = request.DbNames
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.SyncUserPrivilege) {
		query["SyncUserPrivilege"] = request.SyncUserPrivilege
	}

	if !dara.IsNil(request.TargetDBInstanceId) {
		query["TargetDBInstanceId"] = request.TargetDBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyDatabaseBetweenInstances"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyDatabaseBetweenInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database account.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// ### References
//
// > : Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Create an account on an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96089.html)
//
//   - [Create an account on an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96753.html)
//
//   - [Create an account on an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95810.html)
//
//   - [Create an account on an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97132.html)
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

	if !dara.IsNil(request.AccountType) {
		query["AccountType"] = request.AccountType
	}

	if !dara.IsNil(request.CheckPolicy) {
		query["CheckPolicy"] = request.CheckPolicy
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
		Version:     dara.String("2014-08-15"),
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
// Creates a backup file for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)Feature description
//
// This operation uses the backup feature of ApsaraDB RDS to create a backup set. You can also call an operation of Database Backup (DBS) to create a backup set. For more information, see [List of operations by function](https://help.aliyun.com/document_detail/2402073.html).
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The instance is in the Running state.
//
//   - The instance does not have ongoing backup tasks.
//
//   - The number of backup sets that can be created for an instance per day cannot exceed 20.
//
// ### [](#)References
//
//   - [Use the data backup feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/378074.html)
//
//   - [Use the data backup feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96772.html)
//
//   - [Use the data backup feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95717.html)
//
//   - [Use the data backup feature for an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97147.html)
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

	if !dara.IsNil(request.BackupStrategy) {
		query["BackupStrategy"] = request.BackupStrategy
	}

	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackup"),
		Version:     dara.String("2014-08-15"),
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
// Creates an assessment task for cloud migration to an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Use the cloud migration feature](https://help.aliyun.com/document_detail/365562.html)
//
// @param request - CreateCloudMigrationPrecheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudMigrationPrecheckTaskResponse
func (client *Client) CreateCloudMigrationPrecheckTaskWithContext(ctx context.Context, request *CreateCloudMigrationPrecheckTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudMigrationPrecheckTaskResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceAccount) {
		query["SourceAccount"] = request.SourceAccount
	}

	if !dara.IsNil(request.SourceCategory) {
		query["SourceCategory"] = request.SourceCategory
	}

	if !dara.IsNil(request.SourceIpAddress) {
		query["SourceIpAddress"] = request.SourceIpAddress
	}

	if !dara.IsNil(request.SourcePassword) {
		query["SourcePassword"] = request.SourcePassword
	}

	if !dara.IsNil(request.SourcePort) {
		query["SourcePort"] = request.SourcePort
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudMigrationPrecheckTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudMigrationPrecheckTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud migration task for an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Use the cloud migration feature](https://help.aliyun.com/document_detail/365562.html)
//
// @param request - CreateCloudMigrationTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudMigrationTaskResponse
func (client *Client) CreateCloudMigrationTaskWithContext(ctx context.Context, request *CreateCloudMigrationTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudMigrationTaskResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceAccount) {
		query["SourceAccount"] = request.SourceAccount
	}

	if !dara.IsNil(request.SourceCategory) {
		query["SourceCategory"] = request.SourceCategory
	}

	if !dara.IsNil(request.SourceIpAddress) {
		query["SourceIpAddress"] = request.SourceIpAddress
	}

	if !dara.IsNil(request.SourcePassword) {
		query["SourcePassword"] = request.SourcePassword
	}

	if !dara.IsNil(request.SourcePort) {
		query["SourcePort"] = request.SourcePort
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudMigrationTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudMigrationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// >  Fees of an instance are changed if the call is successful. Before you call this operation, carefully read the related topics. If an error message appears when you call this operation, you can search for the error message to view the cause of the error.
//
//   - [Create an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/148036.html).
//
//   - [Create a serverless ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/412231.html).
//
//   - [Create an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/148038.html)
//
//   - [Create a serverless ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/607753.html)
//
//   - [Create an ApsaraDB RDS for PostgreSQL instance for which Babelfish is enabled](https://help.aliyun.com/document_detail/428615.html)
//
//   - [Create an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/148037.html)
//
//   - [Create a serverless ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/603465.html)
//
//   - [Create an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/148040.html)
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
	if !dara.IsNil(tmpReq.ServerlessConfig) {
		request.ServerlessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfig, dara.String("ServerlessConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AutoCreateProxy) {
		query["AutoCreateProxy"] = request.AutoCreateProxy
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BabelfishConfig) {
		query["BabelfishConfig"] = request.BabelfishConfig
	}

	if !dara.IsNil(request.BpeEnabled) {
		query["BpeEnabled"] = request.BpeEnabled
	}

	if !dara.IsNil(request.BurstingEnabled) {
		query["BurstingEnabled"] = request.BurstingEnabled
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ColdDataEnabled) {
		query["ColdDataEnabled"] = request.ColdDataEnabled
	}

	if !dara.IsNil(request.ConnectionMode) {
		query["ConnectionMode"] = request.ConnectionMode
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.CreateStrategy) {
		query["CreateStrategy"] = request.CreateStrategy
	}

	if !dara.IsNil(request.CustomExtraInfo) {
		query["CustomExtraInfo"] = request.CustomExtraInfo
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceNetType) {
		query["DBInstanceNetType"] = request.DBInstanceNetType
	}

	if !dara.IsNil(request.DBInstanceStorage) {
		query["DBInstanceStorage"] = request.DBInstanceStorage
	}

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.DBIsIgnoreCase) {
		query["DBIsIgnoreCase"] = request.DBIsIgnoreCase
	}

	if !dara.IsNil(request.DBParamGroupId) {
		query["DBParamGroupId"] = request.DBParamGroupId
	}

	if !dara.IsNil(request.DBTimeZone) {
		query["DBTimeZone"] = request.DBTimeZone
	}

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
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

	if !dara.IsNil(request.ExternalReplication) {
		query["ExternalReplication"] = request.ExternalReplication
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.IoAccelerationEnabled) {
		query["IoAccelerationEnabled"] = request.IoAccelerationEnabled
	}

	if !dara.IsNil(request.OptimizedWrites) {
		query["OptimizedWrites"] = request.OptimizedWrites
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RoleARN) {
		query["RoleARN"] = request.RoleARN
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.ServerlessConfigShrink) {
		query["ServerlessConfig"] = request.ServerlessConfigShrink
	}

	if !dara.IsNil(request.StorageAutoScale) {
		query["StorageAutoScale"] = request.StorageAutoScale
	}

	if !dara.IsNil(request.StorageThreshold) {
		query["StorageThreshold"] = request.StorageThreshold
	}

	if !dara.IsNil(request.StorageUpperBound) {
		query["StorageUpperBound"] = request.StorageUpperBound
	}

	if !dara.IsNil(request.SystemDBCharset) {
		query["SystemDBCharset"] = request.SystemDBCharset
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetDedicatedHostIdForLog) {
		query["TargetDedicatedHostIdForLog"] = request.TargetDedicatedHostIdForLog
	}

	if !dara.IsNil(request.TargetDedicatedHostIdForMaster) {
		query["TargetDedicatedHostIdForMaster"] = request.TargetDedicatedHostIdForMaster
	}

	if !dara.IsNil(request.TargetDedicatedHostIdForSlave) {
		query["TargetDedicatedHostIdForSlave"] = request.TargetDedicatedHostIdForSlave
	}

	if !dara.IsNil(request.TargetMinorVersion) {
		query["TargetMinorVersion"] = request.TargetMinorVersion
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.UserBackupId) {
		query["UserBackupId"] = request.UserBackupId
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.WhitelistTemplateList) {
		query["WhitelistTemplateList"] = request.WhitelistTemplateList
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.ZoneIdSlave1) {
		query["ZoneIdSlave1"] = request.ZoneIdSlave1
	}

	if !dara.IsNil(request.ZoneIdSlave2) {
		query["ZoneIdSlave2"] = request.ZoneIdSlave2
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstance"),
		Version:     dara.String("2014-08-15"),
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
// Creates an endpoint for an instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engines
//
// # MySQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Create a read-only endpoint for a cluster](https://help.aliyun.com/document_detail/464132.html)
//
// @param tmpReq - CreateDBInstanceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceEndpointResponse
func (client *Client) CreateDBInstanceEndpointWithContext(ctx context.Context, tmpReq *CreateDBInstanceEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDBInstanceEndpointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeItems) {
		request.NodeItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeItems, dara.String("NodeItems"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceEndpointDescription) {
		query["DBInstanceEndpointDescription"] = request.DBInstanceEndpointDescription
	}

	if !dara.IsNil(request.DBInstanceEndpointType) {
		query["DBInstanceEndpointType"] = request.DBInstanceEndpointType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeItemsShrink) {
		query["NodeItems"] = request.NodeItemsShrink
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      dara.String("CreateDBInstanceEndpoint"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBInstanceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a public endpoint for an instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engine
//
// # MySQL
//
// ### [](#)Precautions
//
//   - You can create a public endpoint of an endpoint type only when no public endpoint is created for this endpoint type.
//
//   - The node weights and other configurations are the same as those of the internal endpoint of this endpoint type. Each type of endpoint can contain an internal endpoint and a public endpoint.
//
// @param request - CreateDBInstanceEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceEndpointAddressResponse
func (client *Client) CreateDBInstanceEndpointAddressWithContext(ctx context.Context, request *CreateDBInstanceEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceEndpointAddressResponse, _err error) {
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

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceEndpointId) {
		query["DBInstanceEndpointId"] = request.DBInstanceEndpointId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.IpType) {
		query["IpType"] = request.IpType
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstanceEndpointAddress"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBInstanceEndpointAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rebuilds an instance from the recycle bin.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// ### References
//
// > : Fees are generated if the call is successful. Before you call this operation, carefully read the following documentation:
//
//   - [Manage ApsaraDB RDS for MySQL instances in the recycle bin](https://help.aliyun.com/document_detail/96065.html)
//
//   - [Manage ApsaraDB RDS for PostgreSQL instances in the recycle bin](https://help.aliyun.com/document_detail/96752.html)
//
//   - [Manage ApsaraDB RDS for SQL Server instances in the recycle bin](https://help.aliyun.com/document_detail/95669.html)
//
//   - [Manage ApsaraDB RDS for MariaDB instances in the recycle bin](https://help.aliyun.com/document_detail/97131.html)
//
// @param request - CreateDBInstanceForRebuildRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceForRebuildResponse
func (client *Client) CreateDBInstanceForRebuildWithContext(ctx context.Context, request *CreateDBInstanceForRebuildRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceForRebuildResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceNetType) {
		query["DBInstanceNetType"] = request.DBInstanceNetType
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
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

	if !dara.IsNil(request.ZoneIdSlave1) {
		query["ZoneIdSlave1"] = request.ZoneIdSlave1
	}

	if !dara.IsNil(request.ZoneIdSlave2) {
		query["ZoneIdSlave2"] = request.ZoneIdSlave2
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstanceForRebuild"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBInstanceForRebuildResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a security group rule to an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// [Configure security group settings for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/2392322.html)
//
// @param request - CreateDBInstanceSecurityGroupRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceSecurityGroupRuleResponse
func (client *Client) CreateDBInstanceSecurityGroupRuleWithContext(ctx context.Context, request *CreateDBInstanceSecurityGroupRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceSecurityGroupRuleResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceCidrIp) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstanceSecurityGroupRule"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBInstanceSecurityGroupRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a node to an ApsaraDB RDS for MySQL or ApsaraDB RDS for PostgreSQL instance that runs RDS Cluster Edition. An RDS instance that runs RDS Cluster Edition is referred to as an RDS cluster.
//
// Description:
//
// ### Supported database engines
//
// # MySQL
//
// ### References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Add a node to an ApsaraDB RDS for MySQL cluster](https://help.aliyun.com/document_detail/464129.html)
//
// @param tmpReq - CreateDBNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBNodesResponse
func (client *Client) CreateDBNodesWithContext(ctx context.Context, tmpReq *CreateDBNodesRequest, runtime *dara.RuntimeOptions) (_result *CreateDBNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDBNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBNode) {
		request.DBNodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBNode, dara.String("DBNode"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBNodeShrink) {
		query["DBNode"] = request.DBNodeShrink
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
		Action:      dara.String("CreateDBNodes"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates the endpoint that is used to connect to the dedicated proxy of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Configure the dedicated proxy endpoint of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/184921.html)
//
//   - [Configure the dedicated proxy endpoint of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/418274.html)
//
// @param request - CreateDBProxyEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBProxyEndpointAddressResponse
func (client *Client) CreateDBProxyEndpointAddressWithContext(ctx context.Context, request *CreateDBProxyEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateDBProxyEndpointAddressResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBProxyConnectStringNetType) {
		query["DBProxyConnectStringNetType"] = request.DBProxyConnectStringNetType
	}

	if !dara.IsNil(request.DBProxyEndpointId) {
		query["DBProxyEndpointId"] = request.DBProxyEndpointId
	}

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
	}

	if !dara.IsNil(request.DBProxyNewConnectStringPort) {
		query["DBProxyNewConnectStringPort"] = request.DBProxyNewConnectStringPort
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("CreateDBProxyEndpointAddress"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBProxyEndpointAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database in an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Create a database in an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96105.html)
//
//   - [Create a database in an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96758.html)
//
//   - [Create a database in an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95698.html)
//
//   - [Create a database in an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97136.html)
//
// @param request - CreateDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabaseWithContext(ctx context.Context, request *CreateDatabaseRequest, runtime *dara.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CharacterSetName) {
		query["CharacterSetName"] = request.CharacterSetName
	}

	if !dara.IsNil(request.DBDescription) {
		query["DBDescription"] = request.DBDescription
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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
		Action:      dara.String("CreateDatabase"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores data to a new instance across regions.
//
// Description:
//
// >  Before restoration, you can call the CheckCreateDdrDBInstance operation to check whether a cross-region backup set can be used for cross-region restoration.
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/120824.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206671.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/187923.html)
//
// @param request - CreateDdrInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDdrInstanceResponse
func (client *Client) CreateDdrInstanceWithContext(ctx context.Context, request *CreateDdrInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDdrInstanceResponse, _err error) {
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

	if !dara.IsNil(request.BackupSetRegion) {
		query["BackupSetRegion"] = request.BackupSetRegion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionMode) {
		query["ConnectionMode"] = request.ConnectionMode
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceNetType) {
		query["DBInstanceNetType"] = request.DBInstanceNetType
	}

	if !dara.IsNil(request.DBInstanceStorage) {
		query["DBInstanceStorage"] = request.DBInstanceStorage
	}

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
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

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
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

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	if !dara.IsNil(request.RoleARN) {
		query["RoleARN"] = request.RoleARN
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.SourceDBInstanceName) {
		query["SourceDBInstanceName"] = request.SourceDBInstanceName
	}

	if !dara.IsNil(request.SourceRegion) {
		query["SourceRegion"] = request.SourceRegion
	}

	if !dara.IsNil(request.SystemDBCharset) {
		query["SystemDBCharset"] = request.SystemDBCharset
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
		Action:      dara.String("CreateDdrInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDdrInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ApsaraDB RDS global active database cluster.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
// @param request - CreateGADInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGADInstanceResponse
func (client *Client) CreateGADInstanceWithContext(ctx context.Context, request *CreateGADInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateGADInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CentralDBInstanceId) {
		query["CentralDBInstanceId"] = request.CentralDBInstanceId
	}

	if !dara.IsNil(request.CentralRdsDtsAdminAccount) {
		query["CentralRdsDtsAdminAccount"] = request.CentralRdsDtsAdminAccount
	}

	if !dara.IsNil(request.CentralRdsDtsAdminPassword) {
		query["CentralRdsDtsAdminPassword"] = request.CentralRdsDtsAdminPassword
	}

	if !dara.IsNil(request.CentralRegionId) {
		query["CentralRegionId"] = request.CentralRegionId
	}

	if !dara.IsNil(request.DBList) {
		query["DBList"] = request.DBList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UnitNode) {
		query["UnitNode"] = request.UnitNode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGADInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGADInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds nodes to an ApsaraDB RDS global active database cluster.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Add or remove unit nodes](https://help.aliyun.com/document_detail/331851.html)
//
// @param request - CreateGadInstanceMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGadInstanceMemberResponse
func (client *Client) CreateGadInstanceMemberWithContext(ctx context.Context, request *CreateGadInstanceMemberRequest, runtime *dara.RuntimeOptions) (_result *CreateGadInstanceMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CentralDBInstanceId) {
		query["CentralDBInstanceId"] = request.CentralDBInstanceId
	}

	if !dara.IsNil(request.CentralRdsDtsAdminAccount) {
		query["CentralRdsDtsAdminAccount"] = request.CentralRdsDtsAdminAccount
	}

	if !dara.IsNil(request.CentralRdsDtsAdminPassword) {
		query["CentralRdsDtsAdminPassword"] = request.CentralRdsDtsAdminPassword
	}

	if !dara.IsNil(request.CentralRegionId) {
		query["CentralRegionId"] = request.CentralRegionId
	}

	if !dara.IsNil(request.DBList) {
		query["DBList"] = request.DBList
	}

	if !dara.IsNil(request.GadInstanceId) {
		query["GadInstanceId"] = request.GadInstanceId
	}

	if !dara.IsNil(request.UnitNode) {
		query["UnitNode"] = request.UnitNode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGadInstanceMember"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGadInstanceMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据导入任务
//
// @param request - CreateImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImportTaskResponse
func (client *Client) CreateImportTaskWithContext(ctx context.Context, request *CreateImportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateImportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.EstimatedSize) {
		query["EstimatedSize"] = request.EstimatedSize
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourcePlatform) {
		query["SourcePlatform"] = request.SourcePlatform
	}

	if !dara.IsNil(request.StreamPort) {
		query["StreamPort"] = request.StreamPort
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	if !dara.IsNil(request.XtrabackupPath) {
		query["XtrabackupPath"] = request.XtrabackupPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImportTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建全密态规则
//
// @param tmpReq - CreateMaskingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMaskingRulesResponse
func (client *Client) CreateMaskingRulesWithContext(ctx context.Context, tmpReq *CreateMaskingRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateMaskingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMaskingRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RuleConfig) {
		request.RuleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuleConfig, dara.String("RuleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.DefaultAlgo) {
		query["DefaultAlgo"] = request.DefaultAlgo
	}

	if !dara.IsNil(request.MaskingAlgo) {
		query["MaskingAlgo"] = request.MaskingAlgo
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

	if !dara.IsNil(request.RuleConfigShrink) {
		query["RuleConfig"] = request.RuleConfigShrink
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMaskingRules"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMaskingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a migration task to restore backup files from an Object Storage Service (OSS) bucket to an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)Limits
//
// Data migration across Alibaba Cloud accounts is not supported. For example, backup files in an Object Storage Service (OSS) bucket within Alibaba Cloud Account A cannot be migrated to an ApsaraDB RDS for SQL Server instance within Alibaba Cloud Account B.
//
// >  You can migrate backup files from an OSS bucket within Alibaba Cloud Account A to an OSS bucket within Alibaba Cloud Account B. Make sure that the data in the OSS bucket and the RDS instance belong to the same Alibaba Cloud account. Then, you can call this operation to create a migration task. For more information, see [Use Data Online Migration to migrate data between accounts](https://help.aliyun.com/document_detail/342762.html).
//
// ### [](#)References
//
// >  Before you call this operation, carefully read the following topic. Make sure that you fully understand the prerequisites, preparations, and impacts for calling this operation.
//
// [Migrate data from a self-managed SQL Server instance to an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/100019.html)
//
// @param request - CreateMigrateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMigrateTaskResponse
func (client *Client) CreateMigrateTaskWithContext(ctx context.Context, request *CreateMigrateTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMigrateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.CheckDBMode) {
		query["CheckDBMode"] = request.CheckDBMode
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.IsOnlineDB) {
		query["IsOnlineDB"] = request.IsOnlineDB
	}

	if !dara.IsNil(request.MigrateTaskId) {
		query["MigrateTaskId"] = request.MigrateTaskId
	}

	if !dara.IsNil(request.OSSUrls) {
		query["OSSUrls"] = request.OSSUrls
	}

	if !dara.IsNil(request.OssObjectPositions) {
		query["OssObjectPositions"] = request.OssObjectPositions
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
		Action:      dara.String("CreateMigrateTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMigrateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Opens the database that is involved in a backup data migration task of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - SQL Server
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// This operation is used to migrate backup data to the cloud. Before you call this operation, make sure that you understand the descriptions in the following topics:
//
//   - [Migrate the full backup data of a self-managed SQL Server database to an ApsaraDB RDS instance that runs SQL Server 2008 R2](https://help.aliyun.com/document_detail/95737.html)
//
//   - [Migrate full backup data of SQL Server 2012, 2014, 2016, 2017, or 2019 databases](https://help.aliyun.com/document_detail/95738.html)
//
//   - [Migrate incremental backup data to ApsaraDB RDS for SQL Server 2012, 2014, 2016, 2017, or 2019](https://help.aliyun.com/document_detail/95736.html)
//
// @param request - CreateOnlineDatabaseTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOnlineDatabaseTaskResponse
func (client *Client) CreateOnlineDatabaseTaskWithContext(ctx context.Context, request *CreateOnlineDatabaseTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateOnlineDatabaseTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckDBMode) {
		query["CheckDBMode"] = request.CheckDBMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.MigrateTaskId) {
		query["MigrateTaskId"] = request.MigrateTaskId
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
		Action:      dara.String("CreateOnlineDatabaseTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOnlineDatabaseTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a node from an ApsaraDB RDS for MySQL instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engines
//
// # RDS MySQL
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Delete a node from an ApsaraDB RDS for MySQL instance that runs RDS Cluster Edition](https://help.aliyun.com/document_detail/464130.html)
//
// @param tmpReq - CreateOrderForDeleteDBNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderForDeleteDBNodesResponse
func (client *Client) CreateOrderForDeleteDBNodesWithContext(ctx context.Context, tmpReq *CreateOrderForDeleteDBNodesRequest, runtime *dara.RuntimeOptions) (_result *CreateOrderForDeleteDBNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOrderForDeleteDBNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBNodeId) {
		request.DBNodeIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBNodeId, dara.String("DBNodeId"), dara.String("json"))
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

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBNodeIdShrink) {
		query["DBNodeId"] = request.DBNodeIdShrink
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
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
		Action:      dara.String("CreateOrderForDeleteDBNodes"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrderForDeleteDBNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a parameter template for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/130565.html)
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/457176.html)
//
// @param request - CreateParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParameterGroupResponse
func (client *Client) CreateParameterGroupWithContext(ctx context.Context, request *CreateParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateParameterGroupResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupDesc) {
		query["ParameterGroupDesc"] = request.ParameterGroupDesc
	}

	if !dara.IsNil(request.ParameterGroupName) {
		query["ParameterGroupName"] = request.ParameterGroupName
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
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
		Action:      dara.String("CreateParameterGroup"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateParameterGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an extension for a database.
//
// Description:
//
// ### [](#)Supported database engines
//
// # PostgreSQL
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Manage extensions](https://help.aliyun.com/document_detail/2402409.html)
//
// ### [](#)Usage notes
//
// Install only the extensions that are supported by the major engine version of the instance. Otherwise, the installation fails.
//
//   - For more information, see [Supported extensions](https://help.aliyun.com/document_detail/142340.html).
//
//   - You can call an API operation to query the major engine version of the instance. For more information, see [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/610394.html).
//
// @param request - CreatePostgresExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePostgresExtensionsResponse
func (client *Client) CreatePostgresExtensionsWithContext(ctx context.Context, request *CreatePostgresExtensionsRequest, runtime *dara.RuntimeOptions) (_result *CreatePostgresExtensionsResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBNames) {
		query["DBNames"] = request.DBNames
	}

	if !dara.IsNil(request.Extensions) {
		query["Extensions"] = request.Extensions
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

	if !dara.IsNil(request.RiskConfirmed) {
		query["RiskConfirmed"] = request.RiskConfirmed
	}

	if !dara.IsNil(request.SourceDatabase) {
		query["SourceDatabase"] = request.SourceDatabase
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePostgresExtensions"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePostgresExtensionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a deployment set for an RDS Custom instance in a region. Before you call this operation, you must specify parameters such as OnUnableToRedeployFailedInstance, DeploymentSetName, and Strategy.
//
// @param request - CreateRCDeploymentSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRCDeploymentSetResponse
func (client *Client) CreateRCDeploymentSetWithContext(ctx context.Context, request *CreateRCDeploymentSetRequest, runtime *dara.RuntimeOptions) (_result *CreateRCDeploymentSetResponse, _err error) {
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

	if !dara.IsNil(request.DeploymentSetName) {
		query["DeploymentSetName"] = request.DeploymentSetName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupCount) {
		query["GroupCount"] = request.GroupCount
	}

	if !dara.IsNil(request.OnUnableToRedeployFailedInstance) {
		query["OnUnableToRedeployFailedInstance"] = request.OnUnableToRedeployFailedInstance
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Strategy) {
		query["Strategy"] = request.Strategy
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRCDeploymentSet"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRCDeploymentSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data disk for an RDS Custom instance.
//
// Description:
//
//	  The disk can be an ultra disk, an Enterprise SSD (ESSD), an SSD, or a Premium ESSD. By default, Premium ESSD is used.
//
//		- When you set InstanceChargeType to **Prepaid**, the disk billing method is subscription. You must set **InstanceId*	- to the ID of a subscription instance. The expiration time of the disk must be the same as that of the instance to which the disk is attached.
//
//		- When you set **InstanceChargeType*	- to Postpaid, the disk billing method is pay-as-you-go. You do not need to attach the disk to an instance. You can also attach the pay-as-you-go disk to an instance of any billing method based on your business requirements.
//
//		- The type and number of disks that can be attached to an instance vary based on instance specifications.
//
// @param request - CreateRCDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRCDiskResponse
func (client *Client) CreateRCDiskWithContext(ctx context.Context, request *CreateRCDiskRequest, runtime *dara.RuntimeOptions) (_result *CreateRCDiskResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DiskCategory) {
		query["DiskCategory"] = request.DiskCategory
	}

	if !dara.IsNil(request.DiskName) {
		query["DiskName"] = request.DiskName
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PerformanceLevel) {
		query["PerformanceLevel"] = request.PerformanceLevel
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRCDisk"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRCDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom image for an RDS Custom instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS SQL Server
//
// ### [](#)References
//
//   - [Introduction to RDS Custom for MySQL](https://help.aliyun.com/document_detail/2844223.html)
//
//   - [Introduction to RDS Custom for SQL Server](https://help.aliyun.com/document_detail/2864363.html)
//
// ### [](#)Usage
//
//   - Method 1: Create a custom image by using a snapshot generated from the **system disk**. In this case, specify the SnapshotId and ImageName parameters at the same time in the request.
//
//   - Method 2: Create a custom image by using an RDS Custom instance. In this case, specify the InstanceId and ImageName parameters at the same time in the request.
//
// @param request - CreateRCImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRCImageResponse
func (client *Client) CreateRCImageWithContext(ctx context.Context, request *CreateRCImageRequest, runtime *dara.RuntimeOptions) (_result *CreateRCImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRCImage"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRCImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an edge node pool in the Container Service for Kubernetes (ACK) Edge cluster to which the RDS Custom instance belongs.
//
// @param tmpReq - CreateRCNodePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRCNodePoolResponse
func (client *Client) CreateRCNodePoolWithContext(ctx context.Context, tmpReq *CreateRCNodePoolRequest, runtime *dara.RuntimeOptions) (_result *CreateRCNodePoolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRCNodePoolShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DataDisk) {
		request.DataDiskShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataDisk, dara.String("DataDisk"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SystemDisk) {
		request.SystemDiskShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SystemDisk, dara.String("SystemDisk"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CreateMode) {
		query["CreateMode"] = request.CreateMode
	}

	if !dara.IsNil(request.DataDiskShrink) {
		query["DataDisk"] = request.DataDiskShrink
	}

	if !dara.IsNil(request.DeploymentSetId) {
		query["DeploymentSetId"] = request.DeploymentSetId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.InternetMaxBandwidthOut) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !dara.IsNil(request.IoOptimized) {
		query["IoOptimized"] = request.IoOptimized
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.NodePoolName) {
		query["NodePoolName"] = request.NodePoolName
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityEnhancementStrategy) {
		query["SecurityEnhancementStrategy"] = request.SecurityEnhancementStrategy
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.SupportCase) {
		query["SupportCase"] = request.SupportCase
	}

	if !dara.IsNil(request.SystemDiskShrink) {
		query["SystemDisk"] = request.SystemDiskShrink
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
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
		Action:      dara.String("CreateRCNodePool"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRCNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a snapshot for a cloud disk.
//
// Description:
//
// In the following scenarios, you cannot create snapshots for a specific disk:
//
//   - The number of manual snapshots of the disk has reached 256.
//
//   - A snapshot is being created for the disk.
//
//   - The instance to which the disk is attached has never been started.
//
//   - The instance to which the disk is attached is not in the **Stopped*	- or **Running*	- state.
//
// When you create a snapshot, take note of the following items:
//
//   - If a snapshot is being created, the snapshot cannot be used to create a custom image by calling the CreateImage operation.
//
//   - When a snapshot is being created for a disk that is attached to an RDS Custom instance, do not change the instance state.
//
//   - You can create snapshots for a disk that is in the **Expired*	- state. If the release time scheduled for a disk arrives when a snapshot is being created for the disk, the snapshot in the Creating state is deleted when the disk is released.
//
// @param request - CreateRCSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRCSnapshotResponse
func (client *Client) CreateRCSnapshotWithContext(ctx context.Context, request *CreateRCSnapshotRequest, runtime *dara.RuntimeOptions) (_result *CreateRCSnapshotResponse, _err error) {
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

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.InstantAccess) {
		query["InstantAccess"] = request.InstantAccess
	}

	if !dara.IsNil(request.InstantAccessRetentionDays) {
		query["InstantAccessRetentionDays"] = request.InstantAccessRetentionDays
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRCSnapshot"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRCSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a read-only instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Create a read-only ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/56991.html)
//
//   - [Create a read-only ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/108959.html)
//
//   - [Create a read-only ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/99005.html)
//
// @param request - CreateReadOnlyDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReadOnlyDBInstanceResponse
func (client *Client) CreateReadOnlyDBInstanceWithContext(ctx context.Context, request *CreateReadOnlyDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateReadOnlyDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateProxy) {
		query["AutoCreateProxy"] = request.AutoCreateProxy
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BpeEnabled) {
		query["BpeEnabled"] = request.BpeEnabled
	}

	if !dara.IsNil(request.BurstingEnabled) {
		query["BurstingEnabled"] = request.BurstingEnabled
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CustomExtraInfo) {
		query["CustomExtraInfo"] = request.CustomExtraInfo
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

	if !dara.IsNil(request.DBInstanceStorage) {
		query["DBInstanceStorage"] = request.DBInstanceStorage
	}

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.GdnInstanceName) {
		query["GdnInstanceName"] = request.GdnInstanceName
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.InstructionSetArch) {
		query["InstructionSetArch"] = request.InstructionSetArch
	}

	if !dara.IsNil(request.IoAccelerationEnabled) {
		query["IoAccelerationEnabled"] = request.IoAccelerationEnabled
	}

	if !dara.IsNil(request.IsAnalyticReadOnlyIns) {
		query["IsAnalyticReadOnlyIns"] = request.IsAnalyticReadOnlyIns
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
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

	if !dara.IsNil(request.TargetDedicatedHostIdForMaster) {
		query["TargetDedicatedHostIdForMaster"] = request.TargetDedicatedHostIdForMaster
	}

	if !dara.IsNil(request.TddlBizType) {
		query["TddlBizType"] = request.TddlBizType
	}

	if !dara.IsNil(request.TddlRegionConfig) {
		query["TddlRegionConfig"] = request.TddlRegionConfig
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
		Action:      dara.String("CreateReadOnlyDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReadOnlyDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data synchronization task for a disaster recovery (DR) ApsaraDB RDS instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
//   - SQL Server
//
// >  The parameters vary based on database engines.
//
// @param request - CreateReplicationLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReplicationLinkResponse
func (client *Client) CreateReplicationLinkWithContext(ctx context.Context, request *CreateReplicationLinkRequest, runtime *dara.RuntimeOptions) (_result *CreateReplicationLinkResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ReplicatorAccount) {
		query["ReplicatorAccount"] = request.ReplicatorAccount
	}

	if !dara.IsNil(request.ReplicatorPassword) {
		query["ReplicatorPassword"] = request.ReplicatorPassword
	}

	if !dara.IsNil(request.SourceAddress) {
		query["SourceAddress"] = request.SourceAddress
	}

	if !dara.IsNil(request.SourceCategory) {
		query["SourceCategory"] = request.SourceCategory
	}

	if !dara.IsNil(request.SourceInstanceName) {
		query["SourceInstanceName"] = request.SourceInstanceName
	}

	if !dara.IsNil(request.SourceInstanceRegionId) {
		query["SourceInstanceRegionId"] = request.SourceInstanceRegionId
	}

	if !dara.IsNil(request.SourcePort) {
		query["SourcePort"] = request.SourcePort
	}

	if !dara.IsNil(request.TargetAddress) {
		query["TargetAddress"] = request.TargetAddress
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReplicationLink"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReplicationLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a credential for a user who uses the Data API feature.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
// @param request - CreateSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretResponse
func (client *Client) CreateSecretWithContext(ctx context.Context, request *CreateSecretRequest, runtime *dara.RuntimeOptions) (_result *CreateSecretResponse, _err error) {
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

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.DbNames) {
		query["DbNames"] = request.DbNames
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
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

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecret"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role.
//
// Description:
//
// ### Supported database engine
//
// # PostgreSQL
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Service-linked roles](https://help.aliyun.com/document_detail/342840.html)
//
// @param request - CreateServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithContext(ctx context.Context, request *CreateServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.ServiceLinkedRole) {
		query["ServiceLinkedRole"] = request.ServiceLinkedRole
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceLinkedRole"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a temporary instance for an RDS instance that runs SQL Server 2008 R2 and uses local disks.
//
// Description:
//
// ### [](#)Supported database engines
//
// Your RDS instance runs SQL Server 2008 R2 with local disks.
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Restore the data of an ApsaraDB RDS for SQL Server instance by using a temporary RDS instance](https://help.aliyun.com/document_detail/95724.html)
//
// @param request - CreateTempDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTempDBInstanceResponse
func (client *Client) CreateTempDBInstanceWithContext(ctx context.Context, request *CreateTempDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateTempDBInstanceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTempDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTempDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下单领券接口
//
// @param request - CreateYouhuiForOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateYouhuiForOrderResponse
func (client *Client) CreateYouhuiForOrderWithContext(ctx context.Context, request *CreateYouhuiForOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateYouhuiForOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActivityId) {
		query["ActivityId"] = request.ActivityId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PromotionId) {
		query["PromotionId"] = request.PromotionId
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
		Action:      dara.String("CreateYouhuiForOrder"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateYouhuiForOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ApsaraDB RDS for SQL Server instance from an Active Directory (AD) domain.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - SQL Server
//
// @param request - DeleteADSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteADSettingResponse
func (client *Client) DeleteADSettingWithContext(ctx context.Context, request *DeleteADSettingRequest, runtime *dara.RuntimeOptions) (_result *DeleteADSettingResponse, _err error) {
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
		Action:      dara.String("DeleteADSetting"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteADSettingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an account from an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Delete a database account from an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96104.html)
//
//   - [Delete a database account from an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/147649.html)
//
//   - [Delete a database account from an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95694.html)
//
//   - [Delete a database account from an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97135.html)
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
		Action:      dara.String("DeleteAccount"),
		Version:     dara.String("2014-08-15"),
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
// Deletes the data backup files of an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// > Only instances that run RDS High-availability Edition are supported.
//
// ### Description
//
// You can call this operation to delete backup sets of the instance itself. Backup sets of the associated instances such as read-only, disaster recovery, and cloned instances are not deleted.
//
// ### Precautions
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The instance is in the Running state.
//
//   - If the log backup feature is disabled, instances cannot be restored by point in time. You can delete data backup sets that are retained for more than seven days.
//
//   - If the log backup feature is enabled and the log backup retention period is shorter than the data backup retention period, you can delete the data backup files that are retained for a period longer than the log backup retention period.
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
		Version:     dara.String("2014-08-15"),
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
// Deletes the backup files of an ApsaraDB RDS for SQL Server instance. This operation is available only for users that have been added to the whitelist of the instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// >  **This operation is not supported for new users. **Select other methods to [reduce or save backup costs](https://help.aliyun.com/document_detail/95718.html). Users who are added to the whitelist can still use the backup file. Check the availability of the backup file before you delete it. After the backup file is deleted, you cannot retrieve it.
//
// @param request - DeleteBackupFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupFileResponse
func (client *Client) DeleteBackupFileWithContext(ctx context.Context, request *DeleteBackupFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackupFileResponse, _err error) {
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

	if !dara.IsNil(request.BackupTime) {
		query["BackupTime"] = request.BackupTime
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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
		Action:      dara.String("DeleteBackupFile"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Note Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Release an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96057.html)
//
//   - [Release an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96749.html)
//
//   - [Release an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95662.html)
//
//   - [Release an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97128.html)
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReleasedKeepPolicy) {
		query["ReleasedKeepPolicy"] = request.ReleasedKeepPolicy
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
		Version:     dara.String("2014-08-15"),
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
// Deletes an endpoint for an instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engines
//
// # MySQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Delete the read-only endpoint of an ApsaraDB RDS for MySQL cluster](https://help.aliyun.com/document_detail/464133.html)
//
// @param request - DeleteDBInstanceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstanceEndpointResponse
func (client *Client) DeleteDBInstanceEndpointWithContext(ctx context.Context, request *DeleteDBInstanceEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBInstanceEndpointResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceEndpointId) {
		query["DBInstanceEndpointId"] = request.DBInstanceEndpointId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBInstanceEndpoint"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBInstanceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engines
//
// # MySQL
//
// ### [](#)Precautions
//
// You can delete only the public endpoint of each endpoint type from the instance. If you want to delete an internal endpoint of any endpoint type, you can directly delete the type of endpoint.
//
// @param request - DeleteDBInstanceEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstanceEndpointAddressResponse
func (client *Client) DeleteDBInstanceEndpointAddressWithContext(ctx context.Context, request *DeleteDBInstanceEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBInstanceEndpointAddressResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionString) {
		body["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceEndpointId) {
		body["DBInstanceEndpointId"] = request.DBInstanceEndpointId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBInstanceEndpointAddress"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBInstanceEndpointAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a security group rule that is configured for an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// [Configure security group settings for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/2392322.html)
//
// @param request - DeleteDBInstanceSecurityGroupRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstanceSecurityGroupRuleResponse
func (client *Client) DeleteDBInstanceSecurityGroupRuleWithContext(ctx context.Context, request *DeleteDBInstanceSecurityGroupRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBInstanceSecurityGroupRuleResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityGroupRuleIds) {
		query["SecurityGroupRuleIds"] = request.SecurityGroupRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBInstanceSecurityGroupRule"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBInstanceSecurityGroupRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a node from an instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engine
//
// # MySQL
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Delete a node from an ApsaraDB RDS for MySQL cluster](https://help.aliyun.com/document_detail/464130.html)
//
// @param tmpReq - DeleteDBNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBNodesResponse
func (client *Client) DeleteDBNodesWithContext(ctx context.Context, tmpReq *DeleteDBNodesRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDBNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBNodeId) {
		request.DBNodeIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBNodeId, dara.String("DBNodeId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBNodeIdShrink) {
		query["DBNodeId"] = request.DBNodeIdShrink
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
		Action:      dara.String("DeleteDBNodes"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the endpoint that is used to connect to the dedicated proxy of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Configure the dedicated proxy endpoint of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/184921.html)
//
//   - [Configure the dedicated proxy endpoint for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/418274.html)
//
// @param request - DeleteDBProxyEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBProxyEndpointAddressResponse
func (client *Client) DeleteDBProxyEndpointAddressWithContext(ctx context.Context, request *DeleteDBProxyEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBProxyEndpointAddressResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBProxyConnectStringNetType) {
		query["DBProxyConnectStringNetType"] = request.DBProxyConnectStringNetType
	}

	if !dara.IsNil(request.DBProxyEndpointId) {
		query["DBProxyEndpointId"] = request.DBProxyEndpointId
	}

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBProxyEndpointAddress"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBProxyEndpointAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a database from an RDS instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// ### [](#)References
//
// > : Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Delete a database from an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96106.html)
//
//   - [Delete a database from an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96759.html)
//
//   - [Delete a database from an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95699.html)
//
//   - [Delete a database from an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97137.html)
//
// @param request - DeleteDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatabaseResponse
func (client *Client) DeleteDatabaseWithContext(ctx context.Context, request *DeleteDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatabase"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a global active database cluster.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
// ### [Usage notes](#)
//
//   - A global active database cluster cannot be restored after it is deleted. Proceed with caution.
//
//   - If you delete a global active database cluster, the system removes all nodes and Data Transmission Service (DTS) synchronization tasks from the cluster. However, the system does not release the ApsaraDB RDS for MySQL instances that run as nodes in the cluster. If you no longer need the ApsaraDB RDS for MySQL instances, you can call the [DeleteDBInstance](https://help.aliyun.com/document_detail/26229.html) to release the instances one after another.
//
// @param request - DeleteGadInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGadInstanceResponse
func (client *Client) DeleteGadInstanceWithContext(ctx context.Context, request *DeleteGadInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteGadInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GadInstanceName) {
		query["GadInstanceName"] = request.GadInstanceName
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
		Action:      dara.String("DeleteGadInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGadInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除全密态规则
//
// @param request - DeleteMaskingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMaskingRulesResponse
func (client *Client) DeleteMaskingRulesWithContext(ctx context.Context, request *DeleteMaskingRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteMaskingRulesResponse, _err error) {
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

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMaskingRules"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMaskingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a parameter template from an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/130565.html)
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/457176.html)
//
// @param request - DeleteParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParameterGroupResponse
func (client *Client) DeleteParameterGroupWithContext(ctx context.Context, request *DeleteParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteParameterGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
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
		Action:      dara.String("DeleteParameterGroup"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteParameterGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an extension from a database.
//
// Description:
//
// ### Supported database engines
//
// # RDS PostgreSQL
//
// ### References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Manage extensions](https://help.aliyun.com/document_detail/2402409.html)
//
// @param request - DeletePostgresExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePostgresExtensionsResponse
func (client *Client) DeletePostgresExtensionsWithContext(ctx context.Context, request *DeletePostgresExtensionsRequest, runtime *dara.RuntimeOptions) (_result *DeletePostgresExtensionsResponse, _err error) {
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

	if !dara.IsNil(request.DBNames) {
		query["DBNames"] = request.DBNames
	}

	if !dara.IsNil(request.Extensions) {
		query["Extensions"] = request.Extensions
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
		Action:      dara.String("DeletePostgresExtensions"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePostgresExtensionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a RDS Custom node from a Container Service for Kubernetes (ACK) cluster.
//
// @param tmpReq - DeleteRCClusterNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRCClusterNodesResponse
func (client *Client) DeleteRCClusterNodesWithContext(ctx context.Context, tmpReq *DeleteRCClusterNodesRequest, runtime *dara.RuntimeOptions) (_result *DeleteRCClusterNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteRCClusterNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.NodesShrink) {
		query["Nodes"] = request.NodesShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRCClusterNodes"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRCClusterNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a deployment set for an RDS Custom instance. Before you call this operation, you must specify parameters such as RegionId and DeploymentSetId.
//
// @param request - DeleteRCDeploymentSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRCDeploymentSetResponse
func (client *Client) DeleteRCDeploymentSetWithContext(ctx context.Context, request *DeleteRCDeploymentSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteRCDeploymentSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentSetId) {
		query["DeploymentSetId"] = request.DeploymentSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRCDeploymentSet"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRCDeploymentSetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go data disk. The data disk can be a basic disk, an ultra disk, a standard SSD, or an Enterprise SSD (ESSD).
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - Manual snapshots of the disk are retained.
//
//   - The disk must be in the Unattached (Available) state.
//
//   - If no disk with the specified disk ID exists, the request will be ignored.
//
// @param request - DeleteRCDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRCDiskResponse
func (client *Client) DeleteRCDiskWithContext(ctx context.Context, request *DeleteRCDiskRequest, runtime *dara.RuntimeOptions) (_result *DeleteRCDiskResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRCDisk"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRCDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除RDS用户专属主机实例
//
// @param request - DeleteRCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRCInstanceResponse
func (client *Client) DeleteRCInstanceWithContext(ctx context.Context, request *DeleteRCInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteRCInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRCInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRCInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a subscription RDS Custom instance.
//
// Description:
//
// After an instance is released, all physical resources used by the instance are recycled. Relevant data is erased and cannot be restored.
//
// @param tmpReq - DeleteRCInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRCInstancesResponse
func (client *Client) DeleteRCInstancesWithContext(ctx context.Context, tmpReq *DeleteRCInstancesRequest, runtime *dara.RuntimeOptions) (_result *DeleteRCInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteRCInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceId) {
		request.InstanceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceId, dara.String("InstanceId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.InstanceIdShrink) {
		query["InstanceId"] = request.InstanceIdShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TerminateSubscription) {
		query["TerminateSubscription"] = request.TerminateSubscription
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRCInstances"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRCInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the edge node pool of an RDS Custom instance.
//
// @param request - DeleteRCNodePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRCNodePoolResponse
func (client *Client) DeleteRCNodePoolWithContext(ctx context.Context, request *DeleteRCNodePoolRequest, runtime *dara.RuntimeOptions) (_result *DeleteRCNodePoolResponse, _err error) {
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

	if !dara.IsNil(request.NodePoolId) {
		query["NodePoolId"] = request.NodePoolId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRCNodePool"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRCNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cloud disk snapshot.
//
// Description:
//
// Before you call this operation, take note of the following items:
//
//   - If the specified snapshot ID does not exist, the request will be ignored.
//
//   - If the snapshot is used to create custom images, the snapshot cannot be deleted. You must delete the created custom images before you can delete the snapshot.
//
//   - If the snapshot is used to create cloud disks and `Force` is not specified or is set to `false`, the snapshot cannot be directly deleted. If you want to delete the snapshot, set `Force` to true to forcefully delete the snapshot. The cloud disks created from the snapshot cannot be re-initialized after the snapshot is forcefully deleted.
//
// @param request - DeleteRCSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRCSnapshotResponse
func (client *Client) DeleteRCSnapshotWithContext(ctx context.Context, request *DeleteRCSnapshotRequest, runtime *dara.RuntimeOptions) (_result *DeleteRCSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRCSnapshot"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRCSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RCVCluster删除接口
//
// @param request - DeleteRCVClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRCVClusterResponse
func (client *Client) DeleteRCVClusterWithContext(ctx context.Context, request *DeleteRCVClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteRCVClusterResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRCVCluster"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRCVClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the data synchronization link for a disaster recovery (DR) ApsaraDB RDS instance and promotes the DR instance to the primary instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
//   - SQL Server
//
// @param request - DeleteReplicationLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReplicationLinkResponse
func (client *Client) DeleteReplicationLinkWithContext(ctx context.Context, request *DeleteReplicationLinkRequest, runtime *dara.RuntimeOptions) (_result *DeleteReplicationLinkResponse, _err error) {
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

	if !dara.IsNil(request.PromoteToMaster) {
		query["PromoteToMaster"] = request.PromoteToMaster
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteReplicationLink"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteReplicationLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the credential of a user who uses the Data API feature.
//
// @param request - DeleteSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretResponse
func (client *Client) DeleteSecretWithContext(ctx context.Context, request *DeleteSecretRequest, runtime *dara.RuntimeOptions) (_result *DeleteSecretResponse, _err error) {
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

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
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

	if !dara.IsNil(request.SecretArn) {
		query["SecretArn"] = request.SecretArn
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecret"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecretResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a replication slot of an instance.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - PostgreSQL
//
// ### [](#)Precautions
//
// You can delete a replication slot only when the status of the slot is **INACTIVE**. You can call the DescribeSlots operation to query the status of a replication slot.
//
// @param request - DeleteSlotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSlotResponse
func (client *Client) DeleteSlotWithContext(ctx context.Context, request *DeleteSlotRequest, runtime *dara.RuntimeOptions) (_result *DeleteSlotResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SlotName) {
		query["SlotName"] = request.SlotName
	}

	if !dara.IsNil(request.SlotStatus) {
		query["SlotStatus"] = request.SlotStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSlot"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSlotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the backup files of an ApsaraDB RDS for MySQL instance.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - MySQL
//
// ### [](#)Usage notes
//
//   - A full backup file contains the data of a self-managed MySQL instance. You can restore the data of a self-managed MySQL instance from a full backup file to an ApsaraDB RDS for MySQL instance. For more information, see [Migrate the data of a self-managed MySQL 5.7 instance to the cloud](https://help.aliyun.com/document_detail/251779.html).
//
//   - This operation deletes full backup files only from the ApsaraDB RDS console. This operation does not affect the full backup files that are stored as objects in Object Storage Service (OSS) buckets. After you call this operation to delete a full backup file, you can call the ImportUserBackupFile operation to reimport the full backup file.
//
// @param request - DeleteUserBackupFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserBackupFileResponse
func (client *Client) DeleteUserBackupFileWithContext(ctx context.Context, request *DeleteUserBackupFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserBackupFileResponse, _err error) {
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
		Action:      dara.String("DeleteUserBackupFile"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserBackupFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the migration tasks of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
// @param request - DescibeImportsFromDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescibeImportsFromDatabaseResponse
func (client *Client) DescibeImportsFromDatabaseWithContext(ctx context.Context, request *DescibeImportsFromDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DescibeImportsFromDatabaseResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.ImportId) {
		query["ImportId"] = request.ImportId
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescibeImportsFromDatabase"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescibeImportsFromDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the AD domain of an instance. The information includes whether the instance is added to the AD domain, the name of the AD domain, and the account.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - SQL Server
//
// @param request - DescribeADInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeADInfoResponse
func (client *Client) DescribeADInfoWithContext(ctx context.Context, request *DescribeADInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeADInfoResponse, _err error) {
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
		Action:      dara.String("DescribeADInfo"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeADInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询全密态用户权限
//
// @param request - DescribeAccountMaskingPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountMaskingPrivilegeResponse
func (client *Client) DescribeAccountMaskingPrivilegeWithContext(ctx context.Context, request *DescribeAccountMaskingPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountMaskingPrivilegeResponse, _err error) {
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

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountMaskingPrivilege"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountMaskingPrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about the accounts that are created on an ApsaraDB RDS instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
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
		Action:      dara.String("DescribeAccounts"),
		Version:     dara.String("2014-08-15"),
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

// Summary:
//
// Queries whether the historical events feature is enabled.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// @param request - DescribeActionEventPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeActionEventPolicyResponse
func (client *Client) DescribeActionEventPolicyWithContext(ctx context.Context, request *DescribeActionEventPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeActionEventPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeActionEventPolicy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeActionEventPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about scheduled O\\\\\\\\\\\\\\\\\\\\&M tasks for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
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
		Version:     dara.String("2014-08-15"),
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
// Queries whitelist templates at a time by using fuzzy search.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// @param request - DescribeAllWhitelistTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllWhitelistTemplateResponse
func (client *Client) DescribeAllWhitelistTemplateWithContext(ctx context.Context, request *DescribeAllWhitelistTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeAllWhitelistTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FuzzySearch) {
		query["FuzzySearch"] = request.FuzzySearch
	}

	if !dara.IsNil(request.MaxRecordsPerPage) {
		query["MaxRecordsPerPage"] = request.MaxRecordsPerPage
	}

	if !dara.IsNil(request.PageNumbers) {
		query["PageNumbers"] = request.PageNumbers
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

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAllWhitelistTemplate"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAllWhitelistTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of analytic instances that are associated with an ApsaraDB RDS for MySQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # MySQL
//
// ### [](#)References
//
// [Create and view an analytic instance](https://help.aliyun.com/document_detail/2950002.html)
//
// @param request - DescribeAnalyticdbByPrimaryDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAnalyticdbByPrimaryDBInstanceResponse
func (client *Client) DescribeAnalyticdbByPrimaryDBInstanceWithContext(ctx context.Context, request *DescribeAnalyticdbByPrimaryDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAnalyticdbByPrimaryDBInstanceResponse, _err error) {
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
		Action:      dara.String("DescribeAnalyticdbByPrimaryDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAnalyticdbByPrimaryDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specifications that are supported for an instance. The specifications include the instance type and the storage capacity.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeAvailableClassesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableClassesResponse
func (client *Client) DescribeAvailableClassesWithContext(ctx context.Context, request *DescribeAvailableClassesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableClassesResponse, _err error) {
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

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("DescribeAvailableClasses"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableClassesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available destination regions to which the cross-region backup files from a specific source region can be replicated.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use the cross-region backup feature of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/120824.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206671.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/187923.html)
//
// @param request - DescribeAvailableCrossRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableCrossRegionResponse
func (client *Client) DescribeAvailableCrossRegionWithContext(ctx context.Context, request *DescribeAvailableCrossRegionRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableCrossRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeAvailableCrossRegion"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableCrossRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all Enhanced Monitoring metrics that are supported by an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [View the Enhanced Monitoring metrics](https://help.aliyun.com/document_detail/299200.html)
//
// @param request - DescribeAvailableMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableMetricsResponse
func (client *Client) DescribeAvailableMetricsWithContext(ctx context.Context, request *DescribeAvailableMetricsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableMetricsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableMetrics"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time range within which data can be restored from a cross-region backup file.
//
// Description:
//
// >  To view the time range within which you can restore data from a standard backup set, see DescribeBackups.
//
// ### [](#)Supported database engines
//
// # ApsaraDB RDS for MySQL instances with local disks
//
// ### [](#)References
//
// >  Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Use the cross-region backup feature](https://help.aliyun.com/document_detail/120824.html)
//
// @param request - DescribeAvailableRecoveryTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableRecoveryTimeResponse
func (client *Client) DescribeAvailableRecoveryTimeWithContext(ctx context.Context, request *DescribeAvailableRecoveryTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableRecoveryTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossBackupId) {
		query["CrossBackupId"] = request.CrossBackupId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("DescribeAvailableRecoveryTime"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableRecoveryTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available zones for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// > You can call this operation to query the available zones for an instance. The query result may be different from the zones available on the buy page of the ApsaraDB RDS console. The values of some parameters on the buy page vary based on the actual sales policy. The actual information on the [buy page](https://rdsbuy.console.aliyun.com/create/rds/PostgreSQL) prevails.
//
// @param request - DescribeAvailableZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableZonesResponse
func (client *Client) DescribeAvailableZonesWithContext(ctx context.Context, request *DescribeAvailableZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableZonesResponse, _err error) {
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

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DispenseMode) {
		query["DispenseMode"] = request.DispenseMode
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("DescribeAvailableZones"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases that are involved in a backup file.
//
// Description:
//
// > This operation is phased out.
//
// @param request - DescribeBackupDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupDatabaseResponse
func (client *Client) DescribeBackupDatabaseWithContext(ctx context.Context, request *DescribeBackupDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackupDatabaseResponse, _err error) {
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
		Action:      dara.String("DescribeBackupDatabase"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup settings of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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
	if !dara.IsNil(request.BackupPolicyMode) {
		query["BackupPolicyMode"] = request.BackupPolicyMode
	}

	if !dara.IsNil(request.CompressType) {
		query["CompressType"] = request.CompressType
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

	if !dara.IsNil(request.ReleasedKeepPolicy) {
		query["ReleasedKeepPolicy"] = request.ReleasedKeepPolicy
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
		Action:      dara.String("DescribeBackupPolicy"),
		Version:     dara.String("2014-08-15"),
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
// Queries the backup tasks of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
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

	if !dara.IsNil(request.BackupJobStatus) {
		query["BackupJobStatus"] = request.BackupJobStatus
	}

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Flag) {
		query["Flag"] = request.Flag
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
		Action:      dara.String("DescribeBackupTasks"),
		Version:     dara.String("2014-08-15"),
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
// Queries the data backup files of an ApsaraDB RDS instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
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

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.BackupStatus) {
		query["BackupStatus"] = request.BackupStatus
	}

	if !dara.IsNil(request.BackupType) {
		query["BackupType"] = request.BackupType
	}

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
		Action:      dara.String("DescribeBackups"),
		Version:     dara.String("2014-08-15"),
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
// Queries the binary log files of an instance that runs MySQL or MariaDB or the write-ahead logging (WAL) files of an instance that runs PostgreSQL.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - MariaDB
//
// ### Usage notes
//
//   - If the return value of the **DownloadLink*	- parameter is NULL, ApsaraDB RDS does not provide a download URL.
//
//   - If the return value of the **DownloadLink*	- parameter is not NULL, ApsaraDB RDS provides a URL for you to download backup files. The expiration time of the URL is specified by the **LinkExpiredTime*	- parameter. You must download the backup files before the expiration time.
//
//   - If you use a RAM user to download backup files, you must grant permissions to the RAM user. For more information, see [Grant backup file download permissions to a RAM user with read-only permissions](https://help.aliyun.com/document_detail/100043.html).
//
//   - Each log file that is returned by this operation contains the log entries that are generated over the time range that is specified by the StartTime and EndTime parameters.
//
// @param request - DescribeBinlogFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBinlogFilesResponse
func (client *Client) DescribeBinlogFilesWithContext(ctx context.Context, request *DescribeBinlogFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeBinlogFilesResponse, _err error) {
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
		Action:      dara.String("DescribeBinlogFiles"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBinlogFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the character sets that are supported by an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeCharacterSetNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCharacterSetNameResponse
func (client *Client) DescribeCharacterSetNameWithContext(ctx context.Context, request *DescribeCharacterSetNameRequest, runtime *dara.RuntimeOptions) (_result *DescribeCharacterSetNameResponse, _err error) {
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
		Action:      dara.String("DescribeCharacterSetName"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCharacterSetNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about the instance types of an instance by using the code of the instance types.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeClassDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClassDetailsResponse
func (client *Client) DescribeClassDetailsWithContext(ctx context.Context, request *DescribeClassDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeClassDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassCode) {
		query["ClassCode"] = request.ClassCode
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
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
		Action:      dara.String("DescribeClassDetails"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClassDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the details about the assessment report for cloud migration to an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
// @param request - DescribeCloudMigrationPrecheckResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudMigrationPrecheckResultResponse
func (client *Client) DescribeCloudMigrationPrecheckResultWithContext(ctx context.Context, request *DescribeCloudMigrationPrecheckResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudMigrationPrecheckResultResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceIpAddress) {
		query["SourceIpAddress"] = request.SourceIpAddress
	}

	if !dara.IsNil(request.SourcePort) {
		query["SourcePort"] = request.SourcePort
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudMigrationPrecheckResult"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudMigrationPrecheckResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about the cloud migration task of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
// @param request - DescribeCloudMigrationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudMigrationResultResponse
func (client *Client) DescribeCloudMigrationResultWithContext(ctx context.Context, request *DescribeCloudMigrationResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudMigrationResultResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceIpAddress) {
		query["SourceIpAddress"] = request.SourceIpAddress
	}

	if !dara.IsNil(request.SourcePort) {
		query["SourcePort"] = request.SourcePort
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudMigrationResult"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudMigrationResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the character set collations and time zones that are available for use in ApsaraDB RDS for SQL Server.
//
// Description:
//
// ### Supported database engine
//
// # SQL Server
//
// @param request - DescribeCollationTimeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCollationTimeZonesResponse
func (client *Client) DescribeCollationTimeZonesWithContext(ctx context.Context, request *DescribeCollationTimeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCollationTimeZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeCollationTimeZones"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCollationTimeZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of assured serverless.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [Assured serverless](https://help.aliyun.com/document_detail/2928780.html)
//
// @param request - DescribeComputeBurstConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeComputeBurstConfigResponse
func (client *Client) DescribeComputeBurstConfigWithContext(ctx context.Context, request *DescribeComputeBurstConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeComputeBurstConfigResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeComputeBurstConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeComputeBurstConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases and tables whose data is included in a cross-region backup file of an instance.
//
// Description:
//
// ApsaraDB RDS for MySQL instances support cross-region backup and restoration. For more information, see [Back up an ApsaraDB RDS for MySQL instance across regions](https://help.aliyun.com/document_detail/120824.html) and [Restore the data of an ApsaraDB RDS for MySQL instance across regions](https://help.aliyun.com/document_detail/120875.html).
//
// Before you call this operation, make sure that the instance runs one of the following database engines:
//
//   - MySQL. For more information, see [Back up an ApsaraDB RDS for MySQL instance across regions](https://help.aliyun.com/document_detail/120824.html).
//
//   - SQL Server. For more information, see [Back up an ApsaraDB RDS for SQL Server instance across regions](https://help.aliyun.com/document_detail/187923.html).
//
//   - PostgreSQL. For more information, see [Enable cross-region backups for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206671.html).
//
// @param request - DescribeCrossBackupMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrossBackupMetaListResponse
func (client *Client) DescribeCrossBackupMetaListWithContext(ctx context.Context, request *DescribeCrossBackupMetaListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCrossBackupMetaListResponse, _err error) {
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

	if !dara.IsNil(request.GetDbName) {
		query["GetDbName"] = request.GetDbName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Pattern) {
		query["Pattern"] = request.Pattern
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCrossBackupMetaList"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCrossBackupMetaListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instances for which the cross-region backup feature is enabled in a region and the cross-region backup settings of these instances.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Use the cross-region backup feature of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/120824.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206671.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/187923.html)
//
// @param request - DescribeCrossRegionBackupDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrossRegionBackupDBInstanceResponse
func (client *Client) DescribeCrossRegionBackupDBInstanceWithContext(ctx context.Context, request *DescribeCrossRegionBackupDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeCrossRegionBackupDBInstanceResponse, _err error) {
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
		Action:      dara.String("DescribeCrossRegionBackupDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCrossRegionBackupDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cross-region data backup files of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - ApsaraDB RDS for MySQL instances with local disks
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// ### [](#)References
//
// >  Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/120824.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/187923.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206671.html)
//
// >  For more information about how to query cross-region log backup files, see DescribeCrossRegionLogBackupFiles.
//
// @param request - DescribeCrossRegionBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrossRegionBackupsResponse
func (client *Client) DescribeCrossRegionBackupsWithContext(ctx context.Context, request *DescribeCrossRegionBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCrossRegionBackupsResponse, _err error) {
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

	if !dara.IsNil(request.CrossBackupId) {
		query["CrossBackupId"] = request.CrossBackupId
	}

	if !dara.IsNil(request.CrossBackupRegion) {
		query["CrossBackupRegion"] = request.CrossBackupRegion
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCrossRegionBackups"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCrossRegionBackupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cross-region log backup files of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/120824.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/187923.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206671.html)
//
// >  For more information about how to query cross-region data backup files, see [DescribeCrossRegionBackups](https://help.aliyun.com/document_detail/121733.html).
//
// @param request - DescribeCrossRegionLogBackupFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCrossRegionLogBackupFilesResponse
func (client *Client) DescribeCrossRegionLogBackupFilesWithContext(ctx context.Context, request *DescribeCrossRegionLogBackupFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCrossRegionLogBackupFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossBackupRegion) {
		query["CrossBackupRegion"] = request.CrossBackupRegion
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeCrossRegionLogBackupFiles"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCrossRegionLogBackupFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例最新变配订单
//
// @param request - DescribeCurrentModifyOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCurrentModifyOrderResponse
func (client *Client) DescribeCurrentModifyOrderWithContext(ctx context.Context, request *DescribeCurrentModifyOrderRequest, runtime *dara.RuntimeOptions) (_result *DescribeCurrentModifyOrderResponse, _err error) {
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

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
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
		Action:      dara.String("DescribeCurrentModifyOrder"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCurrentModifyOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例资源使用情况
//
// @param request - DescribeCustinsResourceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustinsResourceInfoResponse
func (client *Client) DescribeCustinsResourceInfoWithContext(ctx context.Context, request *DescribeCustinsResourceInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustinsResourceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceIds) {
		query["DBInstanceIds"] = request.DBInstanceIds
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
		Action:      dara.String("DescribeCustinsResourceInfo"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustinsResourceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
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

	if !dara.IsNil(request.Expired) {
		query["Expired"] = request.Expired
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceAttribute"),
		Version:     dara.String("2014-08-15"),
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
// Queries information about the tags that are added to an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeDBInstanceByTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceByTagsResponse
func (client *Client) DescribeDBInstanceByTagsWithContext(ctx context.Context, request *DescribeDBInstanceByTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceByTagsResponse, _err error) {
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

	if !dara.IsNil(request.ProxyId) {
		query["proxyId"] = request.ProxyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceByTags"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceByTagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询RDS实例的列加密（CLS）配置信息
//
// @param request - DescribeDBInstanceCLSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceCLSResponse
func (client *Client) DescribeDBInstanceCLSWithContext(ctx context.Context, request *DescribeDBInstanceCLSRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceCLSResponse, _err error) {
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
		Action:      dara.String("DescribeDBInstanceCLS"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceCLSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例链路诊断信息
//
// @param request - DescribeDBInstanceConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceConnectivityResponse
func (client *Client) DescribeDBInstanceConnectivityWithContext(ctx context.Context, request *DescribeDBInstanceConnectivityRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceConnectivityResponse, _err error) {
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
		Action:      dara.String("DescribeDBInstanceConnectivity"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceConnectivityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeDBInstanceDetail operation to query the details of an instance.
//
// Description:
//
// This operation is phased out.
//
// @param request - DescribeDBInstanceDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceDetailResponse
func (client *Client) DescribeDBInstanceDetailWithContext(ctx context.Context, request *DescribeDBInstanceDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceDetailResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("DescribeDBInstanceDetail"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeDBInstanceEncryptionKey operation to check whether disk encryption is enabled for an instance. You can also query details about the keys that are used for disk encryption. This operation is supported for instances that run MySQL, SQL Server, or PostgreSQL.
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

	if !dara.IsNil(request.TargetRegionId) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceEncryptionKey"),
		Version:     dara.String("2014-08-15"),
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
// Queries the information about the endpoints of an instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engine
//
// # MySQL
//
// @param request - DescribeDBInstanceEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceEndpointsResponse
func (client *Client) DescribeDBInstanceEndpointsWithContext(ctx context.Context, request *DescribeDBInstanceEndpointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceEndpointsResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceEndpointId) {
		query["DBInstanceEndpointId"] = request.DBInstanceEndpointId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceEndpoints"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the high availability mode and data replication mode of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Query the data replication mode of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96055.html)
//
//   - [Query the data replication mode of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/151265.html)
//
//   - [Query the data replication mode of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/415433.html)
//
// @param request - DescribeDBInstanceHAConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceHAConfigResponse
func (client *Client) DescribeDBInstanceHAConfigWithContext(ctx context.Context, request *DescribeDBInstanceHAConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceHAConfigResponse, _err error) {
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
		Action:      dara.String("DescribeDBInstanceHAConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceHAConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelist of an ApsaraDB RDS instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// @param request - DescribeDBInstanceIPArrayListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceIPArrayListResponse
func (client *Client) DescribeDBInstanceIPArrayListWithContext(ctx context.Context, request *DescribeDBInstanceIPArrayListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WhitelistNetworkType) {
		query["WhitelistNetworkType"] = request.WhitelistNetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceIPArrayList"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the internal IP address and hostname of the Elastic Compute Service (ECS) instance on which the ApsaraDB RDS for SQL Server instance runs.
//
// Description:
//
// ### [](#)Supported database engines
//
// # SQL Server
//
// ### [](#)Prerequisites
//
//   - The RDS instance runs RDS Basic Edition, RDS High-availability Edition, or RDS Cluster Edition. If your RDS instance runs RDS High-availability Edition, make sure that the instance runs SQL Server 2012 or later.
//
//   - The RDS instance belongs to a general-purpose or dedicated instance family. The shared instance family is not supported.
//
//   - If the RDS instance runs RDS Basic Edition, the instance is created on or after September 02, 2022. You can view the Creation Time parameter of an instance in the Status section of the Basic Information page in the ApsaraDB RDS console.
//
// ### [](#)References
//
// >  Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Configure a distributed transaction whitelist](https://help.aliyun.com/document_detail/124321.html)
//
//   - [Connect Kingdee K/3 WISE to an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/124188.html)
//
// @param request - DescribeDBInstanceIpHostnameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceIpHostnameResponse
func (client *Client) DescribeDBInstanceIpHostnameWithContext(ctx context.Context, request *DescribeDBInstanceIpHostnameRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceIpHostnameResponse, _err error) {
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
		Action:      dara.String("DescribeDBInstanceIpHostname"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceIpHostnameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Enhanced Monitoring metrics that are displayed for an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [View the Enhanced Monitoring metrics](https://help.aliyun.com/document_detail/299200.html)
//
// @param request - DescribeDBInstanceMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceMetricsResponse
func (client *Client) DescribeDBInstanceMetricsWithContext(ctx context.Context, request *DescribeDBInstanceMetricsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceMetricsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstanceMetrics"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring frequency of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
//   - MariaDB
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
		Action:      dara.String("DescribeDBInstanceMonitor"),
		Version:     dara.String("2014-08-15"),
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
// Queries all endpoints of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeDBInstanceNetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceNetInfoResponse
func (client *Client) DescribeDBInstanceNetInfoWithContext(ctx context.Context, request *DescribeDBInstanceNetInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceNetRWSplitType) {
		query["DBInstanceNetRWSplitType"] = request.DBInstanceNetRWSplitType
	}

	if !dara.IsNil(request.Flag) {
		query["Flag"] = request.Flag
	}

	if !dara.IsNil(request.GeneralGroupName) {
		query["GeneralGroupName"] = request.GeneralGroupName
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
		Action:      dara.String("DescribeDBInstanceNetInfo"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all endpoints of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeDBInstanceNetInfoForChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceNetInfoForChannelResponse
func (client *Client) DescribeDBInstanceNetInfoForChannelWithContext(ctx context.Context, request *DescribeDBInstanceNetInfoForChannelRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceNetInfoForChannelResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceNetRWSplitType) {
		query["DBInstanceNetRWSplitType"] = request.DBInstanceNetRWSplitType
	}

	if !dara.IsNil(request.Flag) {
		query["Flag"] = request.Flag
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
		Action:      dara.String("DescribeDBInstanceNetInfoForChannel"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceNetInfoForChannelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the performance metrics of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
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
		Action:      dara.String("DescribeDBInstancePerformance"),
		Version:     dara.String("2014-08-15"),
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

// Deprecated: OpenAPI DescribeDBInstancePromoteActivity is deprecated
//
// Summary:
//
// The operation is phased out.
//
// @param request - DescribeDBInstancePromoteActivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancePromoteActivityResponse
func (client *Client) DescribeDBInstancePromoteActivityWithContext(ctx context.Context, request *DescribeDBInstancePromoteActivityRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancePromoteActivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUid) {
		query["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.DbInstanceName) {
		query["DbInstanceName"] = request.DbInstanceName
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
		Action:      dara.String("DescribeDBInstancePromoteActivity"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancePromoteActivityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of shared proxies that are enabled on an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### [](#)Feature description
//
// This operation is used to query the shared proxy settings of an instance that runs MySQL or the read/write splitting settings of an instance that runs SQL Server. For more information about how to query the dedicated proxy settings of an ApsaraDB RDS for MySQL instance, see [DescribeDBProxy](https://help.aliyun.com/document_detail/610506.html).
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The shared proxy feature must be enabled for the primary instance.
//
//   - The read/write splitting feature must be enabled for the primary instance.
//
// @param request - DescribeDBInstanceProxyConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceProxyConfigurationResponse
func (client *Client) DescribeDBInstanceProxyConfigurationWithContext(ctx context.Context, request *DescribeDBInstanceProxyConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceProxyConfigurationResponse, _err error) {
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
		Action:      dara.String("DescribeDBInstanceProxyConfiguration"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceProxyConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and configurations of a native replication instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # MySQL
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [ApsaraDB RDS for MySQL native replication instances](https://help.aliyun.com/document_detail/2856487.html)
//
// @param request - DescribeDBInstanceReplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceReplicationResponse
func (client *Client) DescribeDBInstanceReplicationWithContext(ctx context.Context, request *DescribeDBInstanceReplicationRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceReplicationResponse, _err error) {
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
		Action:      dara.String("DescribeDBInstanceReplication"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceReplicationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SSL configurations of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// ### [](#)References
//
//   - [Use the SSL encryption feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96120.html)
//
//   - [Use the SSL encryption feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/229518.html)
//
//   - [Use the SSL encryption feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95715.html)
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
		Version:     dara.String("2014-08-15"),
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
// Queries the security group rules that are configured for an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// [Configure security group settings for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/2392322.html)
//
// @param request - DescribeDBInstanceSecurityGroupRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceSecurityGroupRuleResponse
func (client *Client) DescribeDBInstanceSecurityGroupRuleWithContext(ctx context.Context, request *DescribeDBInstanceSecurityGroupRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceSecurityGroupRuleResponse, _err error) {
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
		Action:      dara.String("DescribeDBInstanceSecurityGroupRule"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceSecurityGroupRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Transparent Data Encryption (TDE) feature for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
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
		Action:      dara.String("DescribeDBInstanceTDE"),
		Version:     dara.String("2014-08-15"),
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
// Queries instances.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionMode) {
		query["ConnectionMode"] = request.ConnectionMode
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
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

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.DedicatedHostId) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.Expired) {
		query["Expired"] = request.Expired
	}

	if !dara.IsNil(request.Filter) {
		query["Filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceLevel) {
		query["InstanceLevel"] = request.InstanceLevel
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.QueryAutoRenewal) {
		query["QueryAutoRenewal"] = request.QueryAutoRenewal
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

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
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

	if !dara.IsNil(request.ProxyId) {
		query["proxyId"] = request.ProxyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstances"),
		Version:     dara.String("2014-08-15"),
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

// Deprecated: OpenAPI DescribeDBInstancesAsCsv is deprecated, please use Rds::2014-08-15::DescribeDBInstances instead.
//
// Summary:
//
// Queries the instances.
//
// Description:
//
// >  The DescribeDBInstancesAsCsv operation is phased out. You can call the DescribeDBInstances operation.
//
// @param request - DescribeDBInstancesAsCsvRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesAsCsvResponse
func (client *Client) DescribeDBInstancesAsCsvWithContext(ctx context.Context, request *DescribeDBInstancesAsCsvRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesAsCsvResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CachedAsync) {
		query["CachedAsync"] = request.CachedAsync
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ExportKey) {
		query["ExportKey"] = request.ExportKey
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
		Action:      dara.String("DescribeDBInstancesAsCsv"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancesAsCsvResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an ApsaraDB RDS instance based on the remaining subscription duration of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// @param request - DescribeDBInstancesByExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesByExpireTimeResponse
func (client *Client) DescribeDBInstancesByExpireTimeWithContext(ctx context.Context, request *DescribeDBInstancesByExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesByExpireTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpirePeriod) {
		query["ExpirePeriod"] = request.ExpirePeriod
	}

	if !dara.IsNil(request.Expired) {
		query["Expired"] = request.Expired
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.ProxyId) {
		query["proxyId"] = request.ProxyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstancesByExpireTime"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancesByExpireTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeDBInstancePerformance operation to query the performance of instances.
//
// Description:
//
// This operation is phased out.
//
// @param request - DescribeDBInstancesByPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesByPerformanceResponse
func (client *Client) DescribeDBInstancesByPerformanceWithContext(ctx context.Context, request *DescribeDBInstancesByPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesByPerformanceResponse, _err error) {
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

	if !dara.IsNil(request.SortKey) {
		query["SortKey"] = request.SortKey
	}

	if !dara.IsNil(request.SortMethod) {
		query["SortMethod"] = request.SortMethod
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.ProxyId) {
		query["proxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstancesByPerformance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancesByPerformanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeDBInstancesForClone is deprecated, please use Rds::2014-08-15::DescribeDBInstances instead.
//
// Summary:
//
// Queries a list of instances.
//
// Description:
//
// This operation is phased out.
//
// @param request - DescribeDBInstancesForCloneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesForCloneResponse
func (client *Client) DescribeDBInstancesForCloneWithContext(ctx context.Context, request *DescribeDBInstancesForCloneRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesForCloneResponse, _err error) {
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

	if !dara.IsNil(request.ConnectionMode) {
		query["ConnectionMode"] = request.ConnectionMode
	}

	if !dara.IsNil(request.CurrentInstanceId) {
		query["CurrentInstanceId"] = request.CurrentInstanceId
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
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

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.Expired) {
		query["Expired"] = request.Expired
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
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

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
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

	if !dara.IsNil(request.ProxyId) {
		query["proxyId"] = request.ProxyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstancesForClone"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancesForCloneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries minor engine versions that are available for an ApsaraDB RDS for MySQL instance or an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)Usage notes
//
// Before you purchase or upgrade an instance that runs MySQL or PostgreSQL, you can call the DescribeDBMiniEngineVersions operation to query the minor engine versions that are available for the instance.
//
// @param request - DescribeDBMiniEngineVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBMiniEngineVersionsResponse
func (client *Client) DescribeDBMiniEngineVersionsWithContext(ctx context.Context, request *DescribeDBMiniEngineVersionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBMiniEngineVersionsResponse, _err error) {
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

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.MinorVersionTag) {
		query["MinorVersionTag"] = request.MinorVersionTag
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBMiniEngineVersions"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBMiniEngineVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of the database proxy feature for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
// @param request - DescribeDBProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBProxyResponse
func (client *Client) DescribeDBProxyWithContext(ctx context.Context, request *DescribeDBProxyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBProxyResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
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
		Action:      dara.String("DescribeDBProxy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBProxyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the database proxy endpoints of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
// @param request - DescribeDBProxyEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBProxyEndpointResponse
func (client *Client) DescribeDBProxyEndpointWithContext(ctx context.Context, request *DescribeDBProxyEndpointRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBProxyEndpointResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBProxyConnectString) {
		query["DBProxyConnectString"] = request.DBProxyConnectString
	}

	if !dara.IsNil(request.DBProxyEndpointId) {
		query["DBProxyEndpointId"] = request.DBProxyEndpointId
	}

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
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
		Action:      dara.String("DescribeDBProxyEndpoint"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBProxyEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the performance data of the database proxy for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// >  Starting October 17, 2023, ApsaraDB RDS provides a dedicated proxy free of charge for each ApsaraDB RDS for MySQL instance on RDS Cluster Edition. For more information, see [[Special offers/Price changes\\] One proxy is provided free of charge for ApsaraDB RDS for MySQL instances on RDS Cluster Edition](~~2555466~~).
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [View the monitoring data of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/194241.html)
//
//   - [View the monitoring data of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/418275.html)
//
// @param request - DescribeDBProxyPerformanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBProxyPerformanceResponse
func (client *Client) DescribeDBProxyPerformanceWithContext(ctx context.Context, request *DescribeDBProxyPerformanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBProxyPerformanceResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
	}

	if !dara.IsNil(request.DBProxyInstanceType) {
		query["DBProxyInstanceType"] = request.DBProxyInstanceType
	}

	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MetricsName) {
		query["MetricsName"] = request.MetricsName
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBProxyPerformance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBProxyPerformanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distributed transaction whitelists of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # SQL Server
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Configures a distributed transaction whitelist for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/124321.html)
//
// @param request - DescribeDTCSecurityIpHostsForSQLServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDTCSecurityIpHostsForSQLServerResponse
func (client *Client) DescribeDTCSecurityIpHostsForSQLServerWithContext(ctx context.Context, request *DescribeDTCSecurityIpHostsForSQLServerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDTCSecurityIpHostsForSQLServerResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDTCSecurityIpHostsForSQLServer"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDTCSecurityIpHostsForSQLServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about the databases on an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatabasesResponse
func (client *Client) DescribeDatabasesWithContext(ctx context.Context, request *DescribeDatabasesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatabasesResponse, _err error) {
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

	if !dara.IsNil(request.DBStatus) {
		query["DBStatus"] = request.DBStatus
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
		Action:      dara.String("DescribeDatabases"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an ApsaraDB MyBase dedicated cluster.
//
// Description:
//
// Dedicated clusters allow you to manage a number of instances in a cluster at a time. You can create multiple dedicated clusters in a single region. Each dedicated cluster consists of multiple hosts. You can create multiple instances on each host. For more information, see [What is ApsaraDB MyBase?](https://help.aliyun.com/document_detail/141455.html)
//
// @param request - DescribeDedicatedHostGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedHostGroupsResponse
func (client *Client) DescribeDedicatedHostGroupsWithContext(ctx context.Context, request *DescribeDedicatedHostGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedHostGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.ImageCategory) {
		query["ImageCategory"] = request.ImageCategory
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
		Action:      dara.String("DescribeDedicatedHostGroups"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDedicatedHostGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the hosts in a dedicated cluster.
//
// Description:
//
// Dedicated clusters allow you to manage a number of instances at a time. You can create multiple dedicated clusters in a single region. Each dedicated cluster consists of multiple hosts. You can create multiple instances on each host. For more information, see [What is ApsaraDB MyBase?](https://help.aliyun.com/document_detail/141455.html)
//
// @param request - DescribeDedicatedHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDedicatedHostsResponse
func (client *Client) DescribeDedicatedHostsWithContext(ctx context.Context, request *DescribeDedicatedHostsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDedicatedHostsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllocationStatus) {
		query["AllocationStatus"] = request.AllocationStatus
	}

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.DedicatedHostId) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !dara.IsNil(request.HostStatus) {
		query["HostStatus"] = request.HostStatus
	}

	if !dara.IsNil(request.HostType) {
		query["HostType"] = request.HostType
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
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

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDedicatedHosts"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDedicatedHostsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data backup files of an ApsaraDB RDS instance that is released.
//
// Description:
//
// ### Supported database engine
//
// # MySQL
//
// > This operation is available only for instances that use local disks.
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
// For more information about how to retain the data backup files of an instance after the instance is released, see [Configure automatic backup](https://help.aliyun.com/document_detail/98818.html).
//
// @param request - DescribeDetachedBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDetachedBackupsResponse
func (client *Client) DescribeDetachedBackupsWithContext(ctx context.Context, request *DescribeDetachedBackupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDetachedBackupsResponse, _err error) {
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

	if !dara.IsNil(request.BackupMode) {
		query["BackupMode"] = request.BackupMode
	}

	if !dara.IsNil(request.BackupStatus) {
		query["BackupStatus"] = request.BackupStatus
	}

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

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("DescribeDetachedBackups"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDetachedBackupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the error logs of an instance over a specific time range.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeErrorLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeErrorLogsResponse
func (client *Client) DescribeErrorLogsWithContext(ctx context.Context, request *DescribeErrorLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeErrorLogsResponse, _err error) {
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
		Action:      dara.String("DescribeErrorLogs"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeErrorLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries historical events of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Historical events of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/129759.html)
//
//   - [Historical events of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/131008.html)
//
//   - [Historical events of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/131013.html)
//
//   - [Historical events of an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/131010.html)
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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
		Action:      dara.String("DescribeEvents"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Queries the information about a global active database cluster or information about all global active database clusters in a region.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
// @param request - DescribeGadInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGadInstancesResponse
func (client *Client) DescribeGadInstancesWithContext(ctx context.Context, request *DescribeGadInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGadInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GadInstanceName) {
		query["GadInstanceName"] = request.GadInstanceName
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
		Action:      dara.String("DescribeGadInstances"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGadInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the availability check method of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// [What is availability detection?](https://help.aliyun.com/document_detail/207467.html)
//
// @param request - DescribeHADiagnoseConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHADiagnoseConfigResponse
func (client *Client) DescribeHADiagnoseConfigWithContext(ctx context.Context, request *DescribeHADiagnoseConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeHADiagnoseConfigResponse, _err error) {
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
		Action:      dara.String("DescribeHADiagnoseConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHADiagnoseConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of the automatic primary/secondary switchover feature for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeHASwitchConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHASwitchConfigResponse
func (client *Client) DescribeHASwitchConfigWithContext(ctx context.Context, request *DescribeHASwitchConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeHASwitchConfigResponse, _err error) {
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
		Action:      dara.String("DescribeHASwitchConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHASwitchConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries historical events in the event center.
//
// @param request - DescribeHistoryEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryEventsResponse
func (client *Client) DescribeHistoryEventsWithContext(ctx context.Context, request *DescribeHistoryEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArchiveStatus) {
		query["ArchiveStatus"] = request.ArchiveStatus
	}

	if !dara.IsNil(request.EventCategory) {
		query["EventCategory"] = request.EventCategory
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventLevel) {
		query["EventLevel"] = request.EventLevel
	}

	if !dara.IsNil(request.EventStatus) {
		query["EventStatus"] = request.EventStatus
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.FromStartTime) {
		query["FromStartTime"] = request.FromStartTime
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.ToStartTime) {
		query["ToStartTime"] = request.ToStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryEvents"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of historical events in the event center.
//
// @param request - DescribeHistoryEventsStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryEventsStatResponse
func (client *Client) DescribeHistoryEventsStatWithContext(ctx context.Context, request *DescribeHistoryEventsStatRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryEventsStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArchiveStatus) {
		query["ArchiveStatus"] = request.ArchiveStatus
	}

	if !dara.IsNil(request.FromStartTime) {
		query["FromStartTime"] = request.FromStartTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ToStartTime) {
		query["ToStartTime"] = request.ToStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryEventsStat"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryEventsStatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical tasks that are created within 30 days.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Tasks of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/474275.html)
//
//   - [Tasks of an ApsaraDB RDS for PostrgreSQL instance](https://help.aliyun.com/document_detail/474537.html)
//
//   - [Tasks of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/614826.html)
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
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
		Version:     dara.String("2014-08-15"),
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
// Collects tasks in the task center.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
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
		Version:     dara.String("2014-08-15"),
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
// 查询主机组弹性策略参数
//
// @param request - DescribeHostGroupElasticStrategyParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHostGroupElasticStrategyParametersResponse
func (client *Client) DescribeHostGroupElasticStrategyParametersWithContext(ctx context.Context, request *DescribeHostGroupElasticStrategyParametersRequest, runtime *dara.RuntimeOptions) (_result *DescribeHostGroupElasticStrategyParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedHostGroupName) {
		query["DedicatedHostGroupName"] = request.DedicatedHostGroupName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHostGroupElasticStrategyParameters"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHostGroupElasticStrategyParametersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the webshell URL that is used to connect to the host of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - SQL Server
//
// ### [](#)Prerequisite
//
//   - The instance meets the following requirements:
//
//   - The instance resides in a region other than the China (Zhangjiakou) region.
//
//   - The instance runs RDS Basic Edition, RDS Cluster Edition, or RDS High-availability Edition. If your instance runs RDS High-availability Edition, make sure that the instance runs SQL Server 2012 or later.
//
//   - The instance belongs to the general-purpose or dedicated instance family. The shared instance family is not supported.
//
//   - The instance resides in a virtual private cloud (VPC). For more information about how to change the network type of an RDS instance, see [Change the network type](https://help.aliyun.com/document_detail/95707.html).
//
//   - If the instance runs RDS High-availability Edition or RDS Cluster Edition, the instance is created on or after January 1, 2021. If the instance runs RDS Basic Edition, the instance is created on or after September 02, 2022. You can view the **Creation Time*	- parameter of an instance in the **Status*	- section of the **Basic Information*	- page in the ApsaraDB RDS console.
//
//   - Your **Alibaba Cloud account*	- is used for logons.
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Create a host account and use the host account for logons](https://help.aliyun.com/document_detail/354862.html)
//
// @param request - DescribeHostWebShellRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHostWebShellResponse
func (client *Client) DescribeHostWebShellWithContext(ctx context.Context, request *DescribeHostWebShellRequest, runtime *dara.RuntimeOptions) (_result *DescribeHostWebShellResponse, _err error) {
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

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionID) {
		query["RegionID"] = request.RegionID
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
		Action:      dara.String("DescribeHostWebShell"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHostWebShellResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询原生复制上云任务详情
//
// @param request - DescribeImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImportTaskResponse
func (client *Client) DescribeImportTaskWithContext(ctx context.Context, request *DescribeImportTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeImportTaskResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImportTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据导入预检查状态
//
// @param request - DescribeImportTaskValidationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImportTaskValidationResponse
func (client *Client) DescribeImportTaskValidationWithContext(ctx context.Context, request *DescribeImportTaskValidationRequest, runtime *dara.RuntimeOptions) (_result *DescribeImportTaskValidationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
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
		Action:      dara.String("DescribeImportTaskValidation"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImportTaskValidationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the automatic renewal status of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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

	if !dara.IsNil(request.ProxyId) {
		query["proxyId"] = request.ProxyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceAutoRenewalAttribute"),
		Version:     dara.String("2014-08-15"),
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
// Queries the cross-region backup settings of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/120824.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206671.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/187923.html)
//
// @param request - DescribeInstanceCrossBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceCrossBackupPolicyResponse
func (client *Client) DescribeInstanceCrossBackupPolicyWithContext(ctx context.Context, request *DescribeInstanceCrossBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceCrossBackupPolicyResponse, _err error) {
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
		Action:      dara.String("DescribeInstanceCrossBackupPolicy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceCrossBackupPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the reserved keywords of an instance. The reserved keywords cannot be used for the usernames of accounts or the names of databases.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// @param request - DescribeInstanceKeywordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceKeywordsResponse
func (client *Client) DescribeInstanceKeywordsWithContext(ctx context.Context, request *DescribeInstanceKeywordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceKeywordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
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
		Action:      dara.String("DescribeInstanceKeywords"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceKeywordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries associated whitelists by instance name.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// @param request - DescribeInstanceLinkedWhitelistTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceLinkedWhitelistTemplateResponse
func (client *Client) DescribeInstanceLinkedWhitelistTemplateWithContext(ctx context.Context, request *DescribeInstanceLinkedWhitelistTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceLinkedWhitelistTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InsName) {
		query["InsName"] = request.InsName
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
		Action:      dara.String("DescribeInstanceLinkedWhitelistTemplate"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceLinkedWhitelistTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the specified resource of Key Management Service (KMS) is associated with an ApsaraDB RDS instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// @param request - DescribeKmsAssociateResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeKmsAssociateResourcesResponse
func (client *Client) DescribeKmsAssociateResourcesWithContext(ctx context.Context, request *DescribeKmsAssociateResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeKmsAssociateResourcesResponse, _err error) {
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

	if !dara.IsNil(request.KmsResourceId) {
		query["KmsResourceId"] = request.KmsResourceId
	}

	if !dara.IsNil(request.KmsResourceRegionId) {
		query["KmsResourceRegionId"] = request.KmsResourceRegionId
	}

	if !dara.IsNil(request.KmsResourceType) {
		query["KmsResourceType"] = request.KmsResourceType
	}

	if !dara.IsNil(request.KmsResourceUser) {
		query["KmsResourceUser"] = request.KmsResourceUser
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
		Action:      dara.String("DescribeKmsAssociateResources"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeKmsAssociateResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time range to which an RDS instance can be restored.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - MariaDB
//
// @param request - DescribeLocalAvailableRecoveryTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLocalAvailableRecoveryTimeResponse
func (client *Client) DescribeLocalAvailableRecoveryTimeWithContext(ctx context.Context, request *DescribeLocalAvailableRecoveryTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLocalAvailableRecoveryTimeResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLocalAvailableRecoveryTime"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLocalAvailableRecoveryTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log backup files of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # SQL Server
//
// >  You can call the DescribeBinlogFiles operation to query the log files of instances that run different database engines.
//
// @param request - DescribeLogBackupFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogBackupFilesResponse
func (client *Client) DescribeLogBackupFilesWithContext(ctx context.Context, request *DescribeLogBackupFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogBackupFilesResponse, _err error) {
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
		Action:      dara.String("DescribeLogBackupFiles"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogBackupFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取RDS营销项目中待升级实例信息
//
// @param request - DescribeMarketingActivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMarketingActivityResponse
func (client *Client) DescribeMarketingActivityWithContext(ctx context.Context, request *DescribeMarketingActivityRequest, runtime *dara.RuntimeOptions) (_result *DescribeMarketingActivityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUid) {
		query["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.Bid) {
		query["Bid"] = request.Bid
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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

	if !dara.IsNil(request.UpgradeCode) {
		query["UpgradeCode"] = request.UpgradeCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMarketingActivity"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMarketingActivityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询全密态规则
//
// @param request - DescribeMaskingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMaskingRulesResponse
func (client *Client) DescribeMaskingRulesWithContext(ctx context.Context, request *DescribeMaskingRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeMaskingRulesResponse, _err error) {
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

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMaskingRules"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMaskingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the databases and tables that can be restored from a specified backup set.
//
// Description:
//
// ### [](#)Supported database engines
//
// # MySQL
//
// > This operation is available for RDS instances that run MySQL 8.0, MySQL 5.7, and MySQL 5.6 on RDS High-availability Edition with local disks.
//
// ### [](#)Description
//
// Before you call the [RestoreTable](https://help.aliyun.com/document_detail/131510.html) operation to restore individual databases or tables of an ApsaraDB RDS for MySQL instance, you can call this operation to query the information about the databases and tables that can be restored. For more information, see [Restore individual databases and tables of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/103175.html).
//
// @param request - DescribeMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMetaListResponse
func (client *Client) DescribeMetaListWithContext(ctx context.Context, request *DescribeMetaListRequest, runtime *dara.RuntimeOptions) (_result *DescribeMetaListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupSetID) {
		query["BackupSetID"] = request.BackupSetID
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.GetDbName) {
		query["GetDbName"] = request.GetDbName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Pattern) {
		query["Pattern"] = request.Pattern
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMetaList"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMetaListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a task that is used to import the backup data of an ApsaraDB RDS for SQL Server instance to an Object Storage Service (OSS) bucket.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - SQL Server
//
// @param request - DescribeMigrateTaskByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrateTaskByIdResponse
func (client *Client) DescribeMigrateTaskByIdWithContext(ctx context.Context, request *DescribeMigrateTaskByIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrateTaskByIdResponse, _err error) {
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

	if !dara.IsNil(request.MigrateTaskId) {
		query["MigrateTaskId"] = request.MigrateTaskId
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
		Action:      dara.String("DescribeMigrateTaskById"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMigrateTaskByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tasks that are created to migrate the backup data of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - SQL Server
//
// ### [](#)Usage notes
//
// This operation allows you to query the migration tasks that are created for the instance over the last week.
//
// ### [](#)Precautions
//
//   - This operation is supported only for migration tasks that are created to migrate full backup files.
//
//   - This operation is not supported for instances that run SQL Server 2017 on RDS Cluster Edition.
//
// @param request - DescribeMigrateTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMigrateTasksResponse
func (client *Client) DescribeMigrateTasksWithContext(ctx context.Context, request *DescribeMigrateTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeMigrateTasksResponse, _err error) {
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMigrateTasks"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMigrateTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about the modifications to the pg_hba.conf file of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # RDS PostgreSQL
//
// @param request - DescribeModifyPGHbaConfigLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModifyPGHbaConfigLogResponse
func (client *Client) DescribeModifyPGHbaConfigLogWithContext(ctx context.Context, request *DescribeModifyPGHbaConfigLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeModifyPGHbaConfigLogResponse, _err error) {
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

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModifyPGHbaConfigLog"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModifyPGHbaConfigLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameter modification logs of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeModifyParameterLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModifyParameterLogResponse
func (client *Client) DescribeModifyParameterLogWithContext(ctx context.Context, request *DescribeModifyParameterLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeModifyParameterLogResponse, _err error) {
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
		Action:      dara.String("DescribeModifyParameterLog"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backup files that are included in a backup data migration task of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - SQL Server
//
// ### [Usage notes](#)
//
// This operation is not supported for instances that run SQL Server 2017 EE or SQL Server 2019 EE.
//
// @param request - DescribeOssDownloadsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOssDownloadsResponse
func (client *Client) DescribeOssDownloadsWithContext(ctx context.Context, request *DescribeOssDownloadsRequest, runtime *dara.RuntimeOptions) (_result *DescribeOssDownloadsResponse, _err error) {
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

	if !dara.IsNil(request.MigrateTaskId) {
		query["MigrateTaskId"] = request.MigrateTaskId
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
		Action:      dara.String("DescribeOssDownloads"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOssDownloadsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of the pg_hba.conf file of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # RDS PostgreSQL
//
// @param request - DescribePGHbaConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePGHbaConfigResponse
func (client *Client) DescribePGHbaConfigWithContext(ctx context.Context, request *DescribePGHbaConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribePGHbaConfigResponse, _err error) {
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
		Action:      dara.String("DescribePGHbaConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePGHbaConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a parameter template.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use a parameter template for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/130565.html)
//
//   - [Use a parameter template for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/457176.html)
//
// @param request - DescribeParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeParameterGroupResponse
func (client *Client) DescribeParameterGroupWithContext(ctx context.Context, request *DescribeParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeParameterGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
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
		Action:      dara.String("DescribeParameterGroup"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeParameterGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameter templates that are available in a region.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/130565.html)
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/457176.html)
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableDetail) {
		query["EnableDetail"] = request.EnableDetail
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
		Action:      dara.String("DescribeParameterGroups"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
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
// Queries parameter templates.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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
		Action:      dara.String("DescribeParameterTemplates"),
		Version:     dara.String("2014-08-15"),
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
// Queries the parameter settings of an instance.
//
// Description:
//
// ### Applicable engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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
		Action:      dara.String("DescribeParameters"),
		Version:     dara.String("2014-08-15"),
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
// Queries extensions that are installed on a database.
//
// Description:
//
// ### Supported database engines
//
// # RDS PostgreSQL
//
// ### References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Manage extensions](https://help.aliyun.com/document_detail/2402409.html)
//
// @param request - DescribePostgresExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePostgresExtensionsResponse
func (client *Client) DescribePostgresExtensionsWithContext(ctx context.Context, request *DescribePostgresExtensionsRequest, runtime *dara.RuntimeOptions) (_result *DescribePostgresExtensionsResponse, _err error) {
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

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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
		Action:      dara.String("DescribePostgresExtensions"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePostgresExtensionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param tmpReq - DescribePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePriceResponse
func (client *Client) DescribePriceWithContext(ctx context.Context, tmpReq *DescribePriceRequest, runtime *dara.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBNode) {
		request.DBNodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBNode, dara.String("DBNode"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ServerlessConfig) {
		request.ServerlessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfig, dara.String("ServerlessConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
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

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.DBNodeShrink) {
		query["DBNode"] = request.DBNodeShrink
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.InstanceUsedType) {
		query["InstanceUsedType"] = request.InstanceUsedType
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

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
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

	if !dara.IsNil(request.ServerlessConfigShrink) {
		query["ServerlessConfig"] = request.ServerlessConfigShrink
	}

	if !dara.IsNil(request.TimeType) {
		query["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePrice"),
		Version:     dara.String("2014-08-15"),
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

// Summary:
//
// 查询RDS快捷售卖配置
//
// @param request - DescribeQuickSaleConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQuickSaleConfigResponse
func (client *Client) DescribeQuickSaleConfigWithContext(ctx context.Context, request *DescribeQuickSaleConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeQuickSaleConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Commodity) {
		query["Commodity"] = request.Commodity
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQuickSaleConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQuickSaleConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询可用区的资源库存
//
// @param request - DescribeRCAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCAvailableResourceResponse
func (client *Client) DescribeRCAvailableResourceWithContext(ctx context.Context, request *DescribeRCAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCAvailableResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cores) {
		query["Cores"] = request.Cores
	}

	if !dara.IsNil(request.DataDiskCategory) {
		query["DataDiskCategory"] = request.DataDiskCategory
	}

	if !dara.IsNil(request.DedicatedHostId) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !dara.IsNil(request.DestinationResource) {
		query["DestinationResource"] = request.DestinationResource
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.IoOptimized) {
		query["IoOptimized"] = request.IoOptimized
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.NetworkCategory) {
		query["NetworkCategory"] = request.NetworkCategory
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SpotDuration) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.SystemDiskCategory) {
		query["SystemDiskCategory"] = request.SystemDiskCategory
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCAvailableResource"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCAvailableResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Cloud Assistant Agent is installed on one or more RDS Custom instances. If Cloud Assistant Agent is installed, the system queries the total number of Cloud Assistant commands that have been run, the number of Cloud Assistant commands that are being run, and the time when Cloud Assistant commands were last run.
//
// Description:
//
//	  Before you run commands on or send files to instances, especially new instances, we recommend that you query the status of Cloud Assistant on the instances by calling this operation and checking the return value of `CloudAssistantStatus`. Run commands on or send files to the instances only when the return value is `true`.
//
//		- During a paged query, when you retrieve the first page of results, set `MaxResults` to specify the maximum number of entries to return in the call. The return value of `NextToken` is a pagination token that can be used in the next call to retrieve a new page of results. When you retrieve a new page of results, set `NextToken` to the `NextToken` value returned in the previous call and set `MaxResults` to specify the maximum number of entries to return in this call.
//
// @param tmpReq - DescribeRCCloudAssistantStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCCloudAssistantStatusResponse
func (client *Client) DescribeRCCloudAssistantStatusWithContext(ctx context.Context, tmpReq *DescribeRCCloudAssistantStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCCloudAssistantStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRCCloudAssistantStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OSType) {
		query["OSType"] = request.OSType
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
		Action:      dara.String("DescribeRCCloudAssistantStatus"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCCloudAssistantStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the kubeconfig file of a Container Service for Kubernetes (ACK) cluster in which an RDS Custom instance resides.
//
// Description:
//
// Kubeconfig files store identity and authentication information that is used by clients to access ACK clusters. To use kubectl to manage an ACK cluster, you must use the kubeconfig file to connect to the ACK cluster. We recommend that you keep kubeconfig files confidential and revoke kubeconfig files that are not in use. This helps prevent data leaks caused by the disclosure of kubeconfig files.
//
// @param request - DescribeRCClusterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCClusterConfigResponse
func (client *Client) DescribeRCClusterConfigWithContext(ctx context.Context, request *DescribeRCClusterConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCClusterConfigResponse, _err error) {
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

	if !dara.IsNil(request.TemporaryDurationMinutes) {
		query["TemporaryDurationMinutes"] = request.TemporaryDurationMinutes
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCClusterConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCClusterConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the RDS custom nodes in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - DescribeRCClusterNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCClusterNodesResponse
func (client *Client) DescribeRCClusterNodesWithContext(ctx context.Context, request *DescribeRCClusterNodesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCClusterNodesResponse, _err error) {
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

	if !dara.IsNil(request.NodePoolId) {
		query["NodePoolId"] = request.NodePoolId
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

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCClusterNodes"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCClusterNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Container Service for Kubernetes (ACK) clusters to which RDS Custom nodes reside in a specific region.
//
// @param request - DescribeRCClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCClustersResponse
func (client *Client) DescribeRCClustersWithContext(ctx context.Context, request *DescribeRCClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCClustersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Profile) {
		query["Profile"] = request.Profile
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCClusters"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCClustersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of one or more deployment sets for RDS Custom instances. Before you call this operation, you must specify parameters such as DeploymentSetIds, Strategy, and DeploymentSetName.
//
// @param request - DescribeRCDeploymentSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCDeploymentSetsResponse
func (client *Client) DescribeRCDeploymentSetsWithContext(ctx context.Context, request *DescribeRCDeploymentSetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCDeploymentSetsResponse, _err error) {
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
		Action:      dara.String("DescribeRCDeploymentSets"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCDeploymentSetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the disk information about an RDS Custom instance.
//
// @param request - DescribeRCDisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCDisksResponse
func (client *Client) DescribeRCDisksWithContext(ctx context.Context, request *DescribeRCDisksRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCDisksResponse, _err error) {
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
		Action:      dara.String("DescribeRCDisks"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCDisksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询RDS用户专属主机实例
//
// @param request - DescribeRCElasticScalingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCElasticScalingResponse
func (client *Client) DescribeRCElasticScalingWithContext(ctx context.Context, request *DescribeRCElasticScalingRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCElasticScalingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SupportCase) {
		query["SupportCase"] = request.SupportCase
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCElasticScaling"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCElasticScalingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom images that can be used to create an RDS Custom instance. Before you call this operation, you must specify parameters such as RegionId.
//
// @param request - DescribeRCImageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCImageListResponse
func (client *Client) DescribeRCImageListWithContext(ctx context.Context, request *DescribeRCImageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCImageListResponse, _err error) {
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
		Action:      dara.String("DescribeRCImageList"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCImageListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an RDS Custom instance.
//
// @param request - DescribeRCInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCInstanceAttributeResponse
func (client *Client) DescribeRCInstanceAttributeWithContext(ctx context.Context, request *DescribeRCInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCInstanceAttributeResponse, _err error) {
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

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.MaxDisksResults) {
		query["MaxDisksResults"] = request.MaxDisksResults
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCInstanceAttribute"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of DDos attacks on an RDS Custom for SQL Server instance and monitors the security status of the instance in real time to assess potential security risks.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// [Introduction to ApsaraDB RDS Custom](https://help.aliyun.com/document_detail/2864363.html)
//
// @param request - DescribeRCInstanceDdosCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCInstanceDdosCountResponse
func (client *Client) DescribeRCInstanceDdosCountWithContext(ctx context.Context, request *DescribeRCInstanceDdosCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCInstanceDdosCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCInstanceDdosCount"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCInstanceDdosCountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the system events of ApsaraDB RDS instances. When you call this operation, you can specify parameters, such as InstanceId and EventType, in the request. By default, non-active system events are queried.
//
// Description:
//
//	  You can query system events that were completed within the last 30 days. No limits apply to the time range for querying uncompleted system events.
//
//		- If you do not specify the EventCycleStatus or InstanceEventCycleStatus parameter, only system events in the Avoidated, Executed, Canceled, or Failed state are included in the query results by default.
//
//		- You can also specify the InstanceEventCycleStatus parameter in the request to query the system events that are in the Scheduled, Executing, or Inquiring state.
//
// @param request - DescribeRCInstanceHistoryEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCInstanceHistoryEventsResponse
func (client *Client) DescribeRCInstanceHistoryEventsWithContext(ctx context.Context, request *DescribeRCInstanceHistoryEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCInstanceHistoryEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventCycleStatus) {
		query["EventCycleStatus"] = request.EventCycleStatus
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.ImpactLevel) {
		query["ImpactLevel"] = request.ImpactLevel
	}

	if !dara.IsNil(request.InstanceEventCycleStatus) {
		query["InstanceEventCycleStatus"] = request.InstanceEventCycleStatus
	}

	if !dara.IsNil(request.InstanceEventType) {
		query["InstanceEventType"] = request.InstanceEventType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
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

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.EventPublishTime) {
		query["EventPublishTime"] = request.EventPublishTime
	}

	if !dara.IsNil(request.NotBefore) {
		query["NotBefore"] = request.NotBefore
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCInstanceHistoryEvents"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCInstanceHistoryEventsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the protection information about an RDS Custom for SQL Server instance and the details of the Anti-DDoS Origin instance to which the RDS Custom instance is added.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// [Introduction to ApsaraDB RDS Custom](https://help.aliyun.com/document_detail/2864363.html)
//
// >  If one or more assets of the current Alibaba Cloud account are added to an [Anti-DDoS Origin instance](https://help.aliyun.com/document_detail/63643.html), you can call the DescribeRCInstanceIpAddress operation to query the DDoS mitigation information and the details of the Anti-DDoS Origin instance. The information and the details include the basic protection threshold and traffic scrubbing threshold for the assets, DDoS mitigation status of the assets, ID of the instance, and the mitigation status of the instance.
//
// @param request - DescribeRCInstanceIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCInstanceIpAddressResponse
func (client *Client) DescribeRCInstanceIpAddressWithContext(ctx context.Context, request *DescribeRCInstanceIpAddressRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCInstanceIpAddressResponse, _err error) {
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

	if !dara.IsNil(request.DdosRegionId) {
		query["DdosRegionId"] = request.DdosRegionId
	}

	if !dara.IsNil(request.DdosStatus) {
		query["DdosStatus"] = request.DdosStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIp) {
		query["InstanceIp"] = request.InstanceIp
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCInstanceIpAddress"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCInstanceIpAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instance families of RDS Custom instances.
//
// @param request - DescribeRCInstanceTypeFamiliesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCInstanceTypeFamiliesResponse
func (client *Client) DescribeRCInstanceTypeFamiliesWithContext(ctx context.Context, request *DescribeRCInstanceTypeFamiliesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCInstanceTypeFamiliesResponse, _err error) {
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
		Action:      dara.String("DescribeRCInstanceTypeFamilies"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCInstanceTypeFamiliesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instance types of RDS Custom instances.
//
// @param tmpReq - DescribeRCInstanceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCInstanceTypesResponse
func (client *Client) DescribeRCInstanceTypesWithContext(ctx context.Context, tmpReq *DescribeRCInstanceTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCInstanceTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRCInstanceTypesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceType) {
		request.InstanceTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceType, dara.String("InstanceType"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceTypeShrink) {
		query["InstanceType"] = request.InstanceTypeShrink
	}

	if !dara.IsNil(request.InstanceTypeFamily) {
		query["InstanceTypeFamily"] = request.InstanceTypeFamily
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCInstanceTypes"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCInstanceTypesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Virtual Network Computing (VNC) logon address of an RDS Custom instance.
//
// Description:
//
// The address returned is valid only for 15 seconds. If you do not use the returned address to establish a connection within 15 seconds, the address expires and you must call the operation again to obtain a new address.
//
// @param request - DescribeRCInstanceVncUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCInstanceVncUrlResponse
func (client *Client) DescribeRCInstanceVncUrlWithContext(ctx context.Context, request *DescribeRCInstanceVncUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCInstanceVncUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCInstanceVncUrl"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCInstanceVncUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an RDS Custom instance.
//
// @param request - DescribeRCInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCInstancesResponse
func (client *Client) DescribeRCInstancesWithContext(ctx context.Context, request *DescribeRCInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCInstancesResponse, _err error) {
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

	if !dara.IsNil(request.HostIp) {
		query["HostIp"] = request.HostIp
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
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

	if !dara.IsNil(request.PublicIp) {
		query["PublicIp"] = request.PublicIp
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCInstances"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询云助手命令执行结果
//
// @param tmpReq - DescribeRCInvocationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCInvocationResultsResponse
func (client *Client) DescribeRCInvocationResultsWithContext(ctx context.Context, tmpReq *DescribeRCInvocationResultsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCInvocationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRCInvocationResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CommandId) {
		query["CommandId"] = request.CommandId
	}

	if !dara.IsNil(request.ContentEncoding) {
		query["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.IncludeHistory) {
		query["IncludeHistory"] = request.IncludeHistory
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InvokeId) {
		query["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.InvokeRecordStatus) {
		query["InvokeRecordStatus"] = request.InvokeRecordStatus
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCInvocationResults"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCInvocationResultsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of a metric for an RDS Custom instance.
//
// @param request - DescribeRCMetricListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCMetricListResponse
func (client *Client) DescribeRCMetricListWithContext(ctx context.Context, request *DescribeRCMetricListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCMetricListResponse, _err error) {
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
		Action:      dara.String("DescribeRCMetricList"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCMetricListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration information about the edge node pool of an RDS Custom instance.
//
// @param request - DescribeRCNodePoolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCNodePoolResponse
func (client *Client) DescribeRCNodePoolWithContext(ctx context.Context, request *DescribeRCNodePoolRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCNodePoolResponse, _err error) {
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
		Action:      dara.String("DescribeRCNodePool"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCNodePoolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更实例规格或系统盘类型之前，查询某一可用区下实例规格或系统盘的库存情况
//
// @param tmpReq - DescribeRCResourcesModificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCResourcesModificationResponse
func (client *Client) DescribeRCResourcesModificationWithContext(ctx context.Context, tmpReq *DescribeRCResourcesModificationRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCResourcesModificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeRCResourcesModificationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Conditionss) {
		request.ConditionssShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Conditionss, dara.String("Conditionss"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConditionssShrink) {
		query["Conditionss"] = request.ConditionssShrink
	}

	if !dara.IsNil(request.Cores) {
		query["Cores"] = request.Cores
	}

	if !dara.IsNil(request.DestinationResource) {
		query["DestinationResource"] = request.DestinationResource
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Memory) {
		query["Memory"] = request.Memory
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

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
		Action:      dara.String("DescribeRCResourcesModification"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCResourcesModificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security groups of RDS Custom instances.
//
// @param request - DescribeRCSecurityGroupListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCSecurityGroupListResponse
func (client *Client) DescribeRCSecurityGroupListWithContext(ctx context.Context, request *DescribeRCSecurityGroupListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCSecurityGroupListResponse, _err error) {
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

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCSecurityGroupList"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCSecurityGroupListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 描述RC安全组规则
//
// @param request - DescribeRCSecurityGroupPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCSecurityGroupPermissionResponse
func (client *Client) DescribeRCSecurityGroupPermissionWithContext(ctx context.Context, request *DescribeRCSecurityGroupPermissionRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCSecurityGroupPermissionResponse, _err error) {
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

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCSecurityGroupPermission"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCSecurityGroupPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of snapshots. The details include the status of the snapshots, the amount of remaining time required to create the snapshots, and the retention period of the automatic snapshots in days.
//
// @param request - DescribeRCSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCSnapshotsResponse
func (client *Client) DescribeRCSnapshotsWithContext(ctx context.Context, request *DescribeRCSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCSnapshotsResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SnapshotIds) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCSnapshots"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCSnapshotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 描述vCluster
//
// @param request - DescribeRCVClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRCVClusterResponse
func (client *Client) DescribeRCVClusterWithContext(ctx context.Context, request *DescribeRCVClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeRCVClusterResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRCVCluster"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRCVClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeRdsResourceSettings is deprecated
//
// Summary:
//
// Obtains the notification settings for instance resources. This operation can still be called but is no longer maintained.
//
// Description:
//
// This operation can still be called but is no longer maintained.
//
// @param request - DescribeRdsResourceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdsResourceSettingsResponse
func (client *Client) DescribeRdsResourceSettingsWithContext(ctx context.Context, request *DescribeRdsResourceSettingsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdsResourceSettingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceNiche) {
		query["ResourceNiche"] = request.ResourceNiche
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
		Action:      dara.String("DescribeRdsResourceSettings"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdsResourceSettingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latency of data replication between a primary instance and its read-only instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// @param request - DescribeReadDBInstanceDelayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeReadDBInstanceDelayResponse
func (client *Client) DescribeReadDBInstanceDelayWithContext(ctx context.Context, request *DescribeReadDBInstanceDelayRequest, runtime *dara.RuntimeOptions) (_result *DescribeReadDBInstanceDelayResponse, _err error) {
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

	if !dara.IsNil(request.ReadInstanceId) {
		query["ReadInstanceId"] = request.ReadInstanceId
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
		Action:      dara.String("DescribeReadDBInstanceDelay"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeReadDBInstanceDelayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeRegionInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionInfosResponse
func (client *Client) DescribeRegionInfosWithContext(ctx context.Context, request *DescribeRegionInfosRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionInfosResponse, _err error) {
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
		Action:      dara.String("DescribeRegionInfos"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionInfosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of all regions and zones for ApsaraDB RDS. The regions that are no longer supported are also queried. Exercise caution when you call this operation.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2014-08-15"),
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
// Queries the renewal fees for a subscription instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
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

	if !dara.IsNil(request.TimeType) {
		query["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRenewalPrice"),
		Version:     dara.String("2014-08-15"),
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
// Queries the operation logs of the data synchronization task for a specified ApsaraDB RDS instance.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - PostgreSQL
//
// @param request - DescribeReplicationLinkLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeReplicationLinkLogsResponse
func (client *Client) DescribeReplicationLinkLogsWithContext(ctx context.Context, request *DescribeReplicationLinkLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeReplicationLinkLogsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeReplicationLinkLogs"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeReplicationLinkLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 概览页资源详情
//
// @param request - DescribeResourceDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceDetailsResponse
func (client *Client) DescribeResourceDetailsWithContext(ctx context.Context, request *DescribeResourceDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceDetailsResponse, _err error) {
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
		Action:      dara.String("DescribeResourceDetails"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the storage usage of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeResourceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceUsageResponse
func (client *Client) DescribeResourceUsageWithContext(ctx context.Context, request *DescribeResourceUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceUsageResponse, _err error) {
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
		Action:      dara.String("DescribeResourceUsage"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the SQL Explorer (SQL Audit) feature for an ApsaraDB RDS instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// @param request - DescribeSQLCollectorPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLCollectorPolicyResponse
func (client *Client) DescribeSQLCollectorPolicyWithContext(ctx context.Context, request *DescribeSQLCollectorPolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLCollectorPolicyResponse, _err error) {
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
		Action:      dara.String("DescribeSQLCollectorPolicy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLCollectorPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the retention period of the log files that are generated by the SQL Explorer feature for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// @param request - DescribeSQLCollectorRetentionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLCollectorRetentionResponse
func (client *Client) DescribeSQLCollectorRetentionWithContext(ctx context.Context, request *DescribeSQLCollectorRetentionRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLCollectorRetentionResponse, _err error) {
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
		Action:      dara.String("DescribeSQLCollectorRetention"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLCollectorRetentionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log files that are generated by the SQL Explorer (SQL Audit) feature for an instance. The DescribeSQLLogFiles operation does not return the log files that are generated by the SQL Explorer feature and manually exported from the ApsaraDB RDS console. The DescribeSQLLogFiles operation returns the SQL Explorer log files that are generated by calling the DescribeSQLLogRecords operation with the request parameter Form set to File.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//     **
//
//     **Note*	- If your instance runs SQL Server, only SQL Server 2008 R2 is supported.
//
// ### [](#)Precautions
//
//   - The DescribeSQLLogFiles operation does not return the log files that are generated by SQL Explorer Trial Edition for an ApsaraDB RDS for MySQL instance.
//
//   - The DescribeSQLLogFiles operation does not return the log files that are generated by the SQL Explorer feature and manually exported from the ApsaraDB RDS console. The DescribeSQLLogFiles operation returns the SQL Explorer log files that are generated by calling the [DescribeSQLLogRecords](https://help.aliyun.com/document_detail/610533.html) operation with the request parameter **Form*	- set to **File**.
//
//   - The exported files are retained for only two days.
//
//     **
//
//     **Note*	- If you have enabled Database Autonomy Service (DAS) Enterprise Edition V2 or V3 and have enabled the SQL Explorer and Audit feature, the exported files are retained for seven days. You can call the [DescribeSqlLogConfig](https://help.aliyun.com/document_detail/2778837.html) operation to query the information about the enabled DAS Enterprise Edition.
//
// @param request - DescribeSQLLogFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLLogFilesResponse
func (client *Client) DescribeSQLLogFilesWithContext(ctx context.Context, request *DescribeSQLLogFilesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLLogFilesResponse, _err error) {
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

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
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
		Action:      dara.String("DescribeSQLLogFiles"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLLogFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs that are generated by the SQL Explorer (SQL Audit) feature for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [Usage notes](#)
//
//   - You can call this operation up to 1,000 times per minute per account. The calls initiated by using both your Alibaba Cloud account and RAM users within your Alibaba Cloud account are counted.
//
//   - This operation cannot be used to query the logs that are generated by SQL Explorer Trial Edition for an ApsaraDB RDS for MySQL instance.
//
//   - When you call this operation and set the **Form*	- parameter to **File*	- to generate an audit file, a maximum of 1 million log entries can be recorded in the audit file, and you cannot filter log entries by keyword.
//
// @param request - DescribeSQLLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLLogRecordsResponse
func (client *Client) DescribeSQLLogRecordsWithContext(ctx context.Context, request *DescribeSQLLogRecordsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLLogRecordsResponse, _err error) {
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

	if !dara.IsNil(request.Database) {
		query["Database"] = request.Database
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Form) {
		query["Form"] = request.Form
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

	if !dara.IsNil(request.SQLId) {
		query["SQLId"] = request.SQLId
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
		Action:      dara.String("DescribeSQLLogRecords"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLLogRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries SQL log reports.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeSQLLogReportListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSQLLogReportListResponse
func (client *Client) DescribeSQLLogReportListWithContext(ctx context.Context, request *DescribeSQLLogReportListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSQLLogReportListResponse, _err error) {
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
		Action:      dara.String("DescribeSQLLogReportList"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSQLLogReportListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the credential of a user who uses the Data API feature.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - MySQL
//
// @param request - DescribeSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecretsResponse
func (client *Client) DescribeSecretsWithContext(ctx context.Context, request *DescribeSecretsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecretsResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
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
		Action:      dara.String("DescribeSecrets"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecretsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries ECS security groups to which an instance is added.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Configure a security group for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/201042.html)
//
//   - [Configure a security group for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206310.html)
//
//   - [Configure a security group for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/2392322.html)
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
		Version:     dara.String("2014-08-15"),
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
// Queries all replication slots of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
// @param request - DescribeSlotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlotsResponse
func (client *Client) DescribeSlotsWithContext(ctx context.Context, request *DescribeSlotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlotsResponse, _err error) {
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
		Action:      dara.String("DescribeSlots"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the slow log details of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### Precautions
//
// The response parameters returned by this operation are updated every minute.
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SQLHASH) {
		query["SQLHASH"] = request.SQLHASH
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlowLogRecords"),
		Version:     dara.String("2014-08-15"),
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
// Queries the statistics on slow query logs.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//     **
//
//     **Note*	- This operation is not supported for RDS instances that run MySQL 5.7 on RDS Basic Edition.
//
//   - SQL Server
//
//     **
//
//     **Note*	- This operation is supported only for RDS instances that run SQL Server 2008 R2.
//
//   - MariaDB
//
// ### [](#)Prerequisites
//
//   - Slow query logs are not collected in real time and may show a latency of 6 to 8 hours.
//
//   - If the return result is empty, check whether the StartTime and EndTime parameters are in UTC. If yes, no slow logs are generated within the specified time range.
//
//   - Starting from September 01, 2024, the template algorithm for slow queries is optimized. When you call the operation, you must change the value of the **SQLHASH*	- parameter. For more information, see [[Notice\\] Optimization of the template algorithm for slow queries](~~2845725~~).
//
// @param request - DescribeSlowLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowLogsResponse
func (client *Client) DescribeSlowLogsWithContext(ctx context.Context, request *DescribeSlowLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlowLogsResponse, _err error) {
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

	if !dara.IsNil(request.SortKey) {
		query["SortKey"] = request.SortKey
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlowLogs"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlowLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the disk of an ApsaraDB RDS for SQL Server instance can be resized online.
//
// Description:
//
// ### Supported database engine
//
// # SQL Server
//
// @param request - DescribeSupportOnlineResizeDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSupportOnlineResizeDiskResponse
func (client *Client) DescribeSupportOnlineResizeDiskWithContext(ctx context.Context, request *DescribeSupportOnlineResizeDiskRequest, runtime *dara.RuntimeOptions) (_result *DescribeSupportOnlineResizeDiskResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSupportOnlineResizeDisk"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSupportOnlineResizeDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [Usage notes](#)
//
//   - If an instance ID is specified, all tags that are added to this instance are queried, and other filter conditions are invalid.
//
//   - If you specify only TagKey, the results that match the specified TagKey are returned. If you specify both TagKey and TagValue, the results that match both the specified TagKey and TagValue are returned.
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.ProxyId) {
		query["proxyId"] = request.ProxyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2014-08-15"),
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
// Queries the tasks in the Waiting or Executing state on an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// >  You can call the [DescribeHistoryTasks](https://help.aliyun.com/document_detail/2627863.html) operation to query the tasks on an ApsaraDB RDS for MySQL or ApsaraDB RDS for PostgreSQL instance
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskAction) {
		query["TaskAction"] = request.TaskAction
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTasks"),
		Version:     dara.String("2014-08-15"),
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
// Queries the check report for a major engine version upgrade of an ApsaraDB RDS for MySQL instance or ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # MySQL
//
// # PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Check report for the major engine version upgrade of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/2794383.html)
//
//   - [Upgrade the major engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/203309.html)
//
//   - [Introduction to the check report of a major engine version upgrade for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/218391.html)
//
// @param request - DescribeUpgradeMajorVersionPrecheckTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpgradeMajorVersionPrecheckTaskResponse
func (client *Client) DescribeUpgradeMajorVersionPrecheckTaskWithContext(ctx context.Context, request *DescribeUpgradeMajorVersionPrecheckTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpgradeMajorVersionPrecheckTaskResponse, _err error) {
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

	if !dara.IsNil(request.TargetMajorVersion) {
		query["TargetMajorVersion"] = request.TargetMajorVersion
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpgradeMajorVersionPrecheckTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpgradeMajorVersionPrecheckTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical tasks for major engine version upgrades of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # PostgreSQL
//
// @param request - DescribeUpgradeMajorVersionTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpgradeMajorVersionTasksResponse
func (client *Client) DescribeUpgradeMajorVersionTasksWithContext(ctx context.Context, request *DescribeUpgradeMajorVersionTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpgradeMajorVersionTasksResponse, _err error) {
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

	if !dara.IsNil(request.TargetMajorVersion) {
		query["TargetMajorVersion"] = request.TargetMajorVersion
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpgradeMajorVersionTasks"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpgradeMajorVersionTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询交换机列表
//
// @param request - DescribeVSwitchListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVSwitchListResponse
func (client *Client) DescribeVSwitchListWithContext(ctx context.Context, request *DescribeVSwitchListRequest, runtime *dara.RuntimeOptions) (_result *DescribeVSwitchListResponse, _err error) {
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

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVSwitchList"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVSwitchListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of vSwitch that are available in a virtual private cloud (VPC).
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - DescribeVSwitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVSwitchesResponse
func (client *Client) DescribeVSwitchesWithContext(ctx context.Context, request *DescribeVSwitchesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
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
		Action:      dara.String("DescribeVSwitches"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the specified IP whitelist.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// @param request - DescribeWhitelistTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWhitelistTemplateResponse
func (client *Client) DescribeWhitelistTemplateWithContext(ctx context.Context, request *DescribeWhitelistTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeWhitelistTemplateResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWhitelistTemplate"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWhitelistTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries associated instances by whitelist template.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// @param request - DescribeWhitelistTemplateLinkedInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWhitelistTemplateLinkedInstanceResponse
func (client *Client) DescribeWhitelistTemplateLinkedInstanceWithContext(ctx context.Context, request *DescribeWhitelistTemplateLinkedInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeWhitelistTemplateLinkedInstanceResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWhitelistTemplateLinkedInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWhitelistTemplateLinkedInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DestroyDBInstance operation to destroy an instance. The instance is in the Locked state.
//
// Description:
//
// The DestroyDBInstance operation is phased out.
//
// @param request - DestroyDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DestroyDBInstanceResponse
func (client *Client) DestroyDBInstanceWithContext(ctx context.Context, request *DestroyDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *DestroyDBInstanceResponse, _err error) {
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
		Action:      dara.String("DestroyDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DestroyDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a unit node from a global active database cluster.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
// ### [Usage notes](#)
//
// This operation can be used to remove only unit nodes.
//
// @param request - DetachGadInstanceMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachGadInstanceMemberResponse
func (client *Client) DetachGadInstanceMemberWithContext(ctx context.Context, request *DetachGadInstanceMemberRequest, runtime *dara.RuntimeOptions) (_result *DetachGadInstanceMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GadInstanceName) {
		query["GadInstanceName"] = request.GadInstanceName
	}

	if !dara.IsNil(request.MemberInstanceName) {
		query["MemberInstanceName"] = request.MemberInstanceName
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
		Action:      dara.String("DetachGadInstanceMember"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachGadInstanceMemberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a pay-as-you-go data disk or a system disk from an RDS Custom instance.
//
// @param request - DetachRCDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachRCDiskResponse
func (client *Client) DetachRCDiskWithContext(ctx context.Context, request *DetachRCDiskRequest, runtime *dara.RuntimeOptions) (_result *DetachRCDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteWithInstance) {
		query["DeleteWithInstance"] = request.DeleteWithInstance
	}

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachRCDisk"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachRCDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a whitelist template from an instance.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// @param request - DetachWhitelistTemplateToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachWhitelistTemplateToInstanceResponse
func (client *Client) DetachWhitelistTemplateToInstanceWithContext(ctx context.Context, request *DetachWhitelistTemplateToInstanceRequest, runtime *dara.RuntimeOptions) (_result *DetachWhitelistTemplateToInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InsName) {
		query["InsName"] = request.InsName
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachWhitelistTemplateToInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachWhitelistTemplateToInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 评估紧急本地扩容磁盘解锁可使用的磁盘空间
//
// @param request - EvaluateLocalExtendDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateLocalExtendDiskResponse
func (client *Client) EvaluateLocalExtendDiskWithContext(ctx context.Context, request *EvaluateLocalExtendDiskRequest, runtime *dara.RuntimeOptions) (_result *EvaluateLocalExtendDiskResponse, _err error) {
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

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EvaluateLocalExtendDisk"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EvaluateLocalExtendDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the topology of an instance.
//
// Description:
//
// ### Supported database engines
//
// # RDS MySQL
//
// @param request - GetDBInstanceTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDBInstanceTopologyResponse
func (client *Client) GetDBInstanceTopologyWithContext(ctx context.Context, request *GetDBInstanceTopologyRequest, runtime *dara.RuntimeOptions) (_result *GetDBInstanceTopologyResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDBInstanceTopology"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDBInstanceTopologyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SSL encryption settings for a dedicated proxy endpoint of an ApsaraDB RDS for MySQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # RDS MySQL
//
// @param request - GetDbProxyInstanceSslRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDbProxyInstanceSslResponse
func (client *Client) GetDbProxyInstanceSslWithContext(ctx context.Context, request *GetDbProxyInstanceSslRequest, runtime *dara.RuntimeOptions) (_result *GetDbProxyInstanceSslResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDbProxyInstanceSsl"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDbProxyInstanceSslResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants an account the permissions on a database of an instance.
//
// Description:
//
// Each account can be granted permissions on one or more databases. Before you call this operation, make sure that the instance is in the Running state.
//
// > This operation is not supported for instances that run SQL Server 2017 on RDS Cluster Edition or run PostgreSQL with local disks.
//
// @param request - GrantAccountPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantAccountPrivilegeResponse
func (client *Client) GrantAccountPrivilegeWithContext(ctx context.Context, request *GrantAccountPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *GrantAccountPrivilegeResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantAccountPrivilege"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grant permissions to a service account.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Grant permissions to the service account of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96102.html)
//
//   - [Grant permissions to the service account of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95693.html)
//
// @param request - GrantOperatorPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantOperatorPermissionResponse
func (client *Client) GrantOperatorPermissionWithContext(ctx context.Context, request *GrantOperatorPermissionRequest, runtime *dara.RuntimeOptions) (_result *GrantOperatorPermissionResponse, _err error) {
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

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Privileges) {
		query["Privileges"] = request.Privileges
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
		Action:      dara.String("GrantOperatorPermission"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantOperatorPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports the backup data of a self-managed instance that runs MySQL 5.7 to an ApsaraDB RDS for MySQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
// ### [](#)Description
//
// A full backup file contains the data of a self-managed MySQL instance. You can restore the data of a self-managed MySQL instance from a full backup file to an ApsaraDB RDS for MySQL instance.
//
// ### [](#)Usage notes
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The self-managed MySQL instance runs MySQL 5.7 and is backed up by using XtraBackup. The name of the backup file ends with `_qp.xb`. For more information, see [Migrate the data of a self-managed MySQL 5.7 instance to the cloud](https://help.aliyun.com/document_detail/251779.html).
//
//   - The full backup file of the self-managed MySQL instance is uploaded to an Object Storage Service (OSS) bucket in the region of the ApsaraDB RDS for MySQL instance. For more information, see [Migrate the data of a self-managed MySQL 5.7 instance to the cloud](https://help.aliyun.com/document_detail/251779.html).
//
// >  This operation is supported only for MySQL 5.7.
//
// @param request - ImportUserBackupFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportUserBackupFileResponse
func (client *Client) ImportUserBackupFileWithContext(ctx context.Context, request *ImportUserBackupFileRequest, runtime *dara.RuntimeOptions) (_result *ImportUserBackupFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupFile) {
		query["BackupFile"] = request.BackupFile
	}

	if !dara.IsNil(request.BucketRegion) {
		query["BucketRegion"] = request.BucketRegion
	}

	if !dara.IsNil(request.BuildReplication) {
		query["BuildReplication"] = request.BuildReplication
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.MasterInfo) {
		query["MasterInfo"] = request.MasterInfo
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
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

	if !dara.IsNil(request.RestoreSize) {
		query["RestoreSize"] = request.RestoreSize
	}

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	if !dara.IsNil(request.SourceInfo) {
		query["SourceInfo"] = request.SourceInfo
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportUserBackupFile"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportUserBackupFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为实例安装云助手Agent
//
// @param tmpReq - InstallRCCloudAssistantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallRCCloudAssistantResponse
func (client *Client) InstallRCCloudAssistantWithContext(ctx context.Context, tmpReq *InstallRCCloudAssistantRequest, runtime *dara.RuntimeOptions) (_result *InstallRCCloudAssistantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InstallRCCloudAssistantShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallRCCloudAssistant"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallRCCloudAssistantResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specification details of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// @param request - ListClassesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClassesResponse
func (client *Client) ListClassesWithContext(ctx context.Context, request *ListClassesRequest, runtime *dara.RuntimeOptions) (_result *ListClassesResponse, _err error) {
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

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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
		Action:      dara.String("ListClasses"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClassesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询上云任务列表
//
// @param request - ListImportTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImportTasksResponse
func (client *Client) ListImportTasksWithContext(ctx context.Context, request *ListImportTasksRequest, runtime *dara.RuntimeOptions) (_result *ListImportTasksResponse, _err error) {
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImportTasks"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImportTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RCVCluster列表接口
//
// @param request - ListRCVClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRCVClustersResponse
func (client *Client) ListRCVClustersWithContext(ctx context.Context, request *ListRCVClustersRequest, runtime *dara.RuntimeOptions) (_result *ListRCVClustersResponse, _err error) {
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
		Action:      dara.String("ListRCVClusters"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRCVClustersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to one or more instances.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
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
		Version:     dara.String("2014-08-15"),
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
// Queries the details about the full backup files that are imported into an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
// ### [](#)Feature description
//
//   - A full backup file contains the data of a self-managed MySQL database. You can restore the data of a self-managed MySQL database from a full backup file to an ApsaraDB RDS for MySQL instance. For more information, see [Migrate the data of a self-managed MySQL 5.7 instance to the cloud](https://help.aliyun.com/document_detail/251779.html).
//
//   - Before you call the [CreateDBInstance](https://help.aliyun.com/document_detail/26228.html) operation to create an ApsaraDB RDS for MySQL instance into which you want to import full backup files, you can call this operation to query the IDs of full backup files.
//
//   - You can call the [ImportUserBackupFile](https://help.aliyun.com/document_detail/260266.html) operation to import a full backup file into an ApsaraDB RDS for MySQL instance.
//
// @param request - ListUserBackupFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserBackupFilesResponse
func (client *Client) ListUserBackupFilesWithContext(ctx context.Context, request *ListUserBackupFilesRequest, runtime *dara.RuntimeOptions) (_result *ListUserBackupFilesResponse, _err error) {
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

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.OssUrl) {
		query["OssUrl"] = request.OssUrl
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
		Action:      dara.String("ListUserBackupFiles"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserBackupFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Locks an account of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Lock an account of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/147649.html)
//
// @param request - LockAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockAccountResponse
func (client *Client) LockAccountWithContext(ctx context.Context, request *LockAccountRequest, runtime *dara.RuntimeOptions) (_result *LockAccountResponse, _err error) {
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
		Action:      dara.String("LockAccount"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LockAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates an instance to a different zone.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Migrate an ApsaraDB RDS for MySQL instance across zones in the same region](https://help.aliyun.com/document_detail/96746.html)
//
//   - [Migrate an ApsaraDB RDS for PostgreSQL instance across zones in the same region](https://help.aliyun.com/document_detail/96746.html)
//
//   - [Migrate an ApsaraDB RDS for SQL Server instance across zones in the same region](https://help.aliyun.com/document_detail/95658.html)
//
// @param request - MigrateConnectionToOtherZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateConnectionToOtherZoneResponse
func (client *Client) MigrateConnectionToOtherZoneWithContext(ctx context.Context, request *MigrateConnectionToOtherZoneRequest, runtime *dara.RuntimeOptions) (_result *MigrateConnectionToOtherZoneResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("MigrateConnectionToOtherZone"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateConnectionToOtherZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates an instance across hosts in a dedicated cluster.
//
// Description:
//
// Dedicated clusters allow you to manage a number of instances at a time. You can create multiple dedicated clusters in a single region. Each dedicated cluster consists of multiple hosts. You can create multiple instances on each host. For more information, see [What is ApsaraDB MyBase?](https://help.aliyun.com/document_detail/141455.html)
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
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
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

	if !dara.IsNil(request.SpecifiedTime) {
		query["SpecifiedTime"] = request.SpecifiedTime
	}

	if !dara.IsNil(request.TargetDedicatedHostIdForMaster) {
		query["TargetDedicatedHostIdForMaster"] = request.TargetDedicatedHostIdForMaster
	}

	if !dara.IsNil(request.TargetDedicatedHostIdForSlave) {
		query["TargetDedicatedHostIdForSlave"] = request.TargetDedicatedHostIdForSlave
	}

	if !dara.IsNil(request.ZoneIdForFollower) {
		query["ZoneIdForFollower"] = request.ZoneIdForFollower
	}

	if !dara.IsNil(request.ZoneIdForLog) {
		query["ZoneIdForLog"] = request.ZoneIdForLog
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateDBInstance"),
		Version:     dara.String("2014-08-15"),
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
// Changes the zone of a node from an ApsaraDB RDS for MySQL instance that runs RDS Cluster Edition.
//
// @param tmpReq - MigrateDBNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateDBNodesResponse
func (client *Client) MigrateDBNodesWithContext(ctx context.Context, tmpReq *MigrateDBNodesRequest, runtime *dara.RuntimeOptions) (_result *MigrateDBNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &MigrateDBNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBNode) {
		request.DBNodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBNode, dara.String("DBNode"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBNodeShrink) {
		query["DBNode"] = request.DBNodeShrink
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
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

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateDBNodes"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateDBNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the whitelist mode of an instance from the standard whitelist mode to the enhanced whitelist mode.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Change the whitelist mode of an ApsaraDB RDS for MySQL instance to the enhanced whitelist mode](https://help.aliyun.com/document_detail/96117.html)
//
//   - [Change the whitelist mode of an ApsaraDB RDS for PostgreSQL instance to the enhanced whitelist mode](https://help.aliyun.com/document_detail/96767.html)
//
// @param request - MigrateSecurityIPModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateSecurityIPModeResponse
func (client *Client) MigrateSecurityIPModeWithContext(ctx context.Context, request *MigrateSecurityIPModeRequest, runtime *dara.RuntimeOptions) (_result *MigrateSecurityIPModeResponse, _err error) {
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
		Action:      dara.String("MigrateSecurityIPMode"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateSecurityIPModeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates an instance across zones in the same region.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Migrate an ApsaraDB RDS for MySQL instance across zones in the same region](https://help.aliyun.com/document_detail/96053.html)
//
//   - [Migrate an ApsaraDB RDS for PostgreSQL instance across zones in the same region](https://help.aliyun.com/document_detail/96746.html)
//
//   - [Migrate an ApsaraDB RDS for SQL Server instance across zones in the same region](https://help.aliyun.com/document_detail/95658.html)
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
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.CustomExtraInfo) {
		query["CustomExtraInfo"] = request.CustomExtraInfo
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

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.IoAccelerationEnabled) {
		query["IoAccelerationEnabled"] = request.IoAccelerationEnabled
	}

	if !dara.IsNil(request.IsModifySpec) {
		query["IsModifySpec"] = request.IsModifySpec
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

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
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

	if !dara.IsNil(request.ZoneIdSlave1) {
		query["ZoneIdSlave1"] = request.ZoneIdSlave1
	}

	if !dara.IsNil(request.ZoneIdSlave2) {
		query["ZoneIdSlave2"] = request.ZoneIdSlave2
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateToOtherZone"),
		Version:     dara.String("2014-08-15"),
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
// Modifies the AD domain information about an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Connect an RDS instance to a self-managed domain](https://help.aliyun.com/document_detail/170734.html)
//
// @param request - ModifyADInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyADInfoResponse
func (client *Client) ModifyADInfoWithContext(ctx context.Context, request *ModifyADInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyADInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ADAccountName) {
		query["ADAccountName"] = request.ADAccountName
	}

	if !dara.IsNil(request.ADDNS) {
		query["ADDNS"] = request.ADDNS
	}

	if !dara.IsNil(request.ADPassword) {
		query["ADPassword"] = request.ADPassword
	}

	if !dara.IsNil(request.ADServerIpAddress) {
		query["ADServerIpAddress"] = request.ADServerIpAddress
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("ModifyADInfo"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyADInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a password policy is applied to an account.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - SQL Server
//
// @param request - ModifyAccountCheckPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountCheckPolicyResponse
func (client *Client) ModifyAccountCheckPolicyWithContext(ctx context.Context, request *ModifyAccountCheckPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountCheckPolicyResponse, _err error) {
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

	if !dara.IsNil(request.CheckPolicy) {
		query["CheckPolicy"] = request.CheckPolicy
	}

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
		Action:      dara.String("ModifyAccountCheckPolicy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountCheckPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a database account.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
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
		Version:     dara.String("2014-08-15"),
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
// 修改全密态用户权限
//
// @param request - ModifyAccountMaskingPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountMaskingPrivilegeResponse
func (client *Client) ModifyAccountMaskingPrivilegeWithContext(ctx context.Context, request *ModifyAccountMaskingPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountMaskingPrivilegeResponse, _err error) {
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

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Privilege) {
		query["Privilege"] = request.Privilege
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

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccountMaskingPrivilege"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountMaskingPrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the password policy for an account of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// SQL Server (This parameter is unavailable for ApsaraDB RDS for SQL Server instances that belong to the shared instance family and run SQL Server 2008 R2.)
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Create a custom password policy for an account of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95640.html)
//
// @param request - ModifyAccountSecurityPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccountSecurityPolicyResponse
func (client *Client) ModifyAccountSecurityPolicyWithContext(ctx context.Context, request *ModifyAccountSecurityPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccountSecurityPolicyResponse, _err error) {
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

	if !dara.IsNil(request.GroupPolicy) {
		query["GroupPolicy"] = request.GroupPolicy
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
		Action:      dara.String("ModifyAccountSecurityPolicy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccountSecurityPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the event history feature of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [View the event history of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/129759.html)
//
//   - [View the event history of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/131008.html)
//
//   - [View the event history of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/131013.html)
//
//   - [View the event history of an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/131010.html)
//
// @param request - ModifyActionEventPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyActionEventPolicyResponse
func (client *Client) ModifyActionEventPolicyWithContext(ctx context.Context, request *ModifyActionEventPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyActionEventPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableEventLog) {
		query["EnableEventLog"] = request.EnableEventLog
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
		Action:      dara.String("ModifyActionEventPolicy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyActionEventPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the switching time of scheduled O\\\\\\&M tasks for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Scheduled events for ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/104183.html)
//
//   - [Scheduled events for ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/104452.html)
//
//   - [Scheduled events for ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/104451.html)
//
//   - [Scheduled events for ApsaraDB RDS for MariaDB instances](https://help.aliyun.com/document_detail/104454.html)
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyActiveOperationTasks"),
		Version:     dara.String("2014-08-15"),
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
// Modifies the backup policy settings of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Configure an automatic backup policy for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/98818.html)
//
//   - [Configure an automatic backup policy for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96772.html)
//
//   - [Configure an automatic backup policy for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95717.html)
//
//   - [Configure an automatic backup policy for an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97147.html)
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
	if !dara.IsNil(request.ArchiveBackupKeepCount) {
		query["ArchiveBackupKeepCount"] = request.ArchiveBackupKeepCount
	}

	if !dara.IsNil(request.ArchiveBackupKeepPolicy) {
		query["ArchiveBackupKeepPolicy"] = request.ArchiveBackupKeepPolicy
	}

	if !dara.IsNil(request.ArchiveBackupRetentionPeriod) {
		query["ArchiveBackupRetentionPeriod"] = request.ArchiveBackupRetentionPeriod
	}

	if !dara.IsNil(request.BackupInterval) {
		query["BackupInterval"] = request.BackupInterval
	}

	if !dara.IsNil(request.BackupLog) {
		query["BackupLog"] = request.BackupLog
	}

	if !dara.IsNil(request.BackupMethod) {
		query["BackupMethod"] = request.BackupMethod
	}

	if !dara.IsNil(request.BackupPolicyMode) {
		query["BackupPolicyMode"] = request.BackupPolicyMode
	}

	if !dara.IsNil(request.BackupPriority) {
		query["BackupPriority"] = request.BackupPriority
	}

	if !dara.IsNil(request.BackupRetentionPeriod) {
		query["BackupRetentionPeriod"] = request.BackupRetentionPeriod
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.CompressType) {
		query["CompressType"] = request.CompressType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EnableBackupLog) {
		query["EnableBackupLog"] = request.EnableBackupLog
	}

	if !dara.IsNil(request.EnableIncrementDataBackup) {
		query["EnableIncrementDataBackup"] = request.EnableIncrementDataBackup
	}

	if !dara.IsNil(request.HighSpaceUsageProtection) {
		query["HighSpaceUsageProtection"] = request.HighSpaceUsageProtection
	}

	if !dara.IsNil(request.LocalLogRetentionHours) {
		query["LocalLogRetentionHours"] = request.LocalLogRetentionHours
	}

	if !dara.IsNil(request.LocalLogRetentionSpace) {
		query["LocalLogRetentionSpace"] = request.LocalLogRetentionSpace
	}

	if !dara.IsNil(request.LogBackupFrequency) {
		query["LogBackupFrequency"] = request.LogBackupFrequency
	}

	if !dara.IsNil(request.LogBackupLocalRetentionNumber) {
		query["LogBackupLocalRetentionNumber"] = request.LogBackupLocalRetentionNumber
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

	if !dara.IsNil(request.ReleasedKeepPolicy) {
		query["ReleasedKeepPolicy"] = request.ReleasedKeepPolicy
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
		Action:      dara.String("ModifyBackupPolicy"),
		Version:     dara.String("2014-08-15"),
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
// Extends the expiration time of backup sets generated by manual backup for a single database, including physical backup sets and full backup sets.
//
// Description:
//
// ### [](#)Supported database engines
//
// # RDS SQL Server
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Manually back up the data of an RDS for SQL Server instance](https://help.aliyun.com/document_detail/95717.html)
//
// @param request - ModifyBackupSetExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackupSetExpireTimeResponse
func (client *Client) ModifyBackupSetExpireTimeWithContext(ctx context.Context, request *ModifyBackupSetExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackupSetExpireTimeResponse, _err error) {
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

	if !dara.IsNil(request.ExpectExpireTime) {
		query["ExpectExpireTime"] = request.ExpectExpireTime
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackupSetExpireTime"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackupSetExpireTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the character set collation and time zone of system databases on an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # RDS SQL Server
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Change the character set collation and the time zone of system databases](https://help.aliyun.com/document_detail/95700.html)
//
// @param request - ModifyCollationTimeZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCollationTimeZoneResponse
func (client *Client) ModifyCollationTimeZoneWithContext(ctx context.Context, request *ModifyCollationTimeZoneRequest, runtime *dara.RuntimeOptions) (_result *ModifyCollationTimeZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Collation) {
		query["Collation"] = request.Collation
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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

	if !dara.IsNil(request.Timezone) {
		query["Timezone"] = request.Timezone
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCollationTimeZone"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCollationTimeZoneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of assured serverless or disables assured serverless.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [Assured serverless](https://help.aliyun.com/document_detail/2928780.html)
//
// @param request - ModifyComputeBurstConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyComputeBurstConfigResponse
func (client *Client) ModifyComputeBurstConfigWithContext(ctx context.Context, request *ModifyComputeBurstConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyComputeBurstConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BurstStatus) {
		query["BurstStatus"] = request.BurstStatus
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CpuEnlargeThreshold) {
		query["CpuEnlargeThreshold"] = request.CpuEnlargeThreshold
	}

	if !dara.IsNil(request.CpuShrinkThreshold) {
		query["CpuShrinkThreshold"] = request.CpuShrinkThreshold
	}

	if !dara.IsNil(request.CrontabJobId) {
		query["CrontabJobId"] = request.CrontabJobId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.MemoryEnlargeThreshold) {
		query["MemoryEnlargeThreshold"] = request.MemoryEnlargeThreshold
	}

	if !dara.IsNil(request.MemoryShrinkThreshold) {
		query["MemoryShrinkThreshold"] = request.MemoryShrinkThreshold
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

	if !dara.IsNil(request.ScaleMaxCpus) {
		query["ScaleMaxCpus"] = request.ScaleMaxCpus
	}

	if !dara.IsNil(request.ScaleMaxMemory) {
		query["ScaleMaxMemory"] = request.ScaleMaxMemory
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyComputeBurstConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyComputeBurstConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例资源
//
// @param request - ModifyCustinsResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustinsResourceResponse
func (client *Client) ModifyCustinsResourceWithContext(ctx context.Context, request *ModifyCustinsResourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustinsResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdjustDeadline) {
		query["AdjustDeadline"] = request.AdjustDeadline
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.IncreaseRatio) {
		query["IncreaseRatio"] = request.IncreaseRatio
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RestoreOriginalSpecification) {
		query["RestoreOriginalSpecification"] = request.RestoreOriginalSpecification
	}

	if !dara.IsNil(request.TargetValue) {
		query["TargetValue"] = request.TargetValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustinsResource"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustinsResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// @param request - ModifyDBDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBDescriptionResponse
func (client *Client) ModifyDBDescriptionWithContext(ctx context.Context, request *ModifyDBDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBDescription) {
		query["DBDescription"] = request.DBDescription
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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
		Action:      dara.String("ModifyDBDescription"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the method that is used to update the minor engine version of an ApsaraDB RDS for MySQL instance or an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Modify automatic update settings for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96059.html)
//
//   - [Modify automatic update settings for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146895.html)
//
// @param request - ModifyDBInstanceAutoUpgradeMinorVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceAutoUpgradeMinorVersionResponse
func (client *Client) ModifyDBInstanceAutoUpgradeMinorVersionWithContext(ctx context.Context, request *ModifyDBInstanceAutoUpgradeMinorVersionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceAutoUpgradeMinorVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUpgradeMinorVersion) {
		query["AutoUpgradeMinorVersion"] = request.AutoUpgradeMinorVersion
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("ModifyDBInstanceAutoUpgradeMinorVersion"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceAutoUpgradeMinorVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置RDS实例开启/修改/关闭列加密状态
//
// @param request - ModifyDBInstanceCLSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceCLSResponse
func (client *Client) ModifyDBInstanceCLSWithContext(ctx context.Context, request *ModifyDBInstanceCLSRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceCLSResponse, _err error) {
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

	if !dara.IsNil(request.EncryptionAlgorithm) {
		query["EncryptionAlgorithm"] = request.EncryptionAlgorithm
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.EncryptionKeyMode) {
		query["EncryptionKeyMode"] = request.EncryptionKeyMode
	}

	if !dara.IsNil(request.EncryptionStatus) {
		query["EncryptionStatus"] = request.EncryptionStatus
	}

	if !dara.IsNil(request.IsRotate) {
		query["IsRotate"] = request.IsRotate
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

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.WhiteListMode) {
		query["WhiteListMode"] = request.WhiteListMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceCLS"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceCLSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration item of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
//   - SQL Server
//
// >  The configuration items that are supported are pgbouncer and clear_errorlog. For more information, see [PgBouncer of ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/2398301.html) and [Error log cleanup of ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/95645.html).
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

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

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConfig"),
		Version:     dara.String("2014-08-15"),
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
// Modifies the endpoint and port of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation:
//
//   - [Change the endpoint and port number of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96163.html)
//
//   - [Change the endpoint and port number of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96788.html)
//
//   - [Change the endpoint and port number of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95740.html)
//
//   - [Change the endpoint and port number of an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97157.html)
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
	if !dara.IsNil(request.BabelfishPort) {
		query["BabelfishPort"] = request.BabelfishPort
	}

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.CurrentConnectionString) {
		query["CurrentConnectionString"] = request.CurrentConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.GeneralGroupName) {
		query["GeneralGroupName"] = request.GeneralGroupName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PGBouncerPort) {
		query["PGBouncerPort"] = request.PGBouncerPort
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RetainVip) {
		query["RetainVip"] = request.RetainVip
	}

	if !dara.IsNil(request.TargetDBInstanceId) {
		query["TargetDBInstanceId"] = request.TargetDBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceConnectionString"),
		Version:     dara.String("2014-08-15"),
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
// Configures the replication latency for a read-only ApsaraDB RDS for MySQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Configure a data replication latency for a read-only ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96056.html)
//
// @param request - ModifyDBInstanceDelayedReplicationTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceDelayedReplicationTimeResponse
func (client *Client) ModifyDBInstanceDelayedReplicationTimeWithContext(ctx context.Context, request *ModifyDBInstanceDelayedReplicationTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceDelayedReplicationTimeResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReadSQLReplicationTime) {
		query["ReadSQLReplicationTime"] = request.ReadSQLReplicationTime
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
		Action:      dara.String("ModifyDBInstanceDelayedReplicationTime"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceDelayedReplicationTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enable or disable the release protection feature for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Enable and disable instance release protection for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/414512.html)
//
//   - [Enable and disable instance release protection for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/471512.html)
//
//   - [Enable and disable instance release protection for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/416209.html)
//
//   - [Enable and disable release protection for an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/414512.html)
//
// @param request - ModifyDBInstanceDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceDeletionProtectionResponse
func (client *Client) ModifyDBInstanceDeletionProtectionWithContext(ctx context.Context, request *ModifyDBInstanceDeletionProtectionRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceDeletionProtectionResponse, _err error) {
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

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
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
		Action:      dara.String("ModifyDBInstanceDeletionProtection"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceDeletionProtectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyDBInstanceDescription operation to modify the name of an instance.
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
		Version:     dara.String("2014-08-15"),
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
// Modifies the weight of an endpoint for an instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engine
//
// # MySQL
//
// @param tmpReq - ModifyDBInstanceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceEndpointResponse
func (client *Client) ModifyDBInstanceEndpointWithContext(ctx context.Context, tmpReq *ModifyDBInstanceEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyDBInstanceEndpointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeItems) {
		request.NodeItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeItems, dara.String("NodeItems"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceEndpointDescription) {
		query["DBInstanceEndpointDescription"] = request.DBInstanceEndpointDescription
	}

	if !dara.IsNil(request.DBInstanceEndpointId) {
		query["DBInstanceEndpointId"] = request.DBInstanceEndpointId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NodeItemsShrink) {
		query["NodeItems"] = request.NodeItemsShrink
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceEndpoint"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about the endpoint of an instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engines
//
// # MySQL
//
// ### [](#)Precautions
//
//   - You can modify the following information about the endpoint of an instance: the public and internal endpoints, the public and internal ports, and the virtual private cloud (VPC), vSwitch, and IP address of the internal endpoint.
//
//   - The VPC and vSwitch must be modified at the same time. If you specify the VPC, vSwitch, and IP address of the internal endpoint, you do not need to specify the endpoint and port. If you specify the endpoint and port, you do not need to specify the VPC, vSwitch, and IP address of the internal endpoint.
//
// @param request - ModifyDBInstanceEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceEndpointAddressResponse
func (client *Client) ModifyDBInstanceEndpointAddressWithContext(ctx context.Context, request *ModifyDBInstanceEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceEndpointAddressResponse, _err error) {
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

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceEndpointId) {
		query["DBInstanceEndpointId"] = request.DBInstanceEndpointId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      dara.String("ModifyDBInstanceEndpointAddress"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceEndpointAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the high availability (HA) and data replication mode of an instance.
//
// @param request - ModifyDBInstanceHAConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceHAConfigResponse
func (client *Client) ModifyDBInstanceHAConfigWithContext(ctx context.Context, request *ModifyDBInstanceHAConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceHAConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.HAMode) {
		query["HAMode"] = request.HAMode
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

	if !dara.IsNil(request.SyncMode) {
		query["SyncMode"] = request.SyncMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceHAConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceHAConfigResponse{}
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
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Set a maintenance window for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96052.html)
//
//   - [Set a maintenance window for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96799.html)
//
//   - [Set a maintenance window for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95657.html)
//
//   - [Set a maintenance window for an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97473.html)
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.MaintainTime) {
		query["MaintainTime"] = request.MaintainTime
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
		Version:     dara.String("2014-08-15"),
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
// Modifies the Enhanced Monitoring metrics that are displayed for an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [View the Enhanced Monitoring metrics](https://help.aliyun.com/document_detail/299200.html)
//
// @param request - ModifyDBInstanceMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceMetricsResponse
func (client *Client) ModifyDBInstanceMetricsWithContext(ctx context.Context, request *ModifyDBInstanceMetricsRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceMetricsResponse, _err error) {
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

	if !dara.IsNil(request.MetricsConfig) {
		query["MetricsConfig"] = request.MetricsConfig
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceMetrics"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceMetricsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a monitoring frequency.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### [](#)Usage notes
//
// If you use the Every 5 Seconds monitoring frequency, you are charged additional fees. Before you call this operation, make sure that you understand the [billing methods and pricing](https://help.aliyun.com/document_detail/45020.html) of ApsaraDB RDS.
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Configure the monitoring frequency for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96112.html)
//
//   - [Configure the monitoring frequency for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95710.html)
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
		Action:      dara.String("ModifyDBInstanceMonitor"),
		Version:     dara.String("2014-08-15"),
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
// Changes the expiration time of the classic network endpoint of an instance in hybrid access mode.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### [](#)References
//
//   - [Configure the hybrid access solution for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96110.html)
//
//   - [Configure the hybrid access solution for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95708.html)
//
// @param request - ModifyDBInstanceNetworkExpireTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceNetworkExpireTimeResponse
func (client *Client) ModifyDBInstanceNetworkExpireTimeWithContext(ctx context.Context, request *ModifyDBInstanceNetworkExpireTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceNetworkExpireTimeResponse, _err error) {
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
		Action:      dara.String("ModifyDBInstanceNetworkExpireTime"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceNetworkExpireTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Changes the network type of an ApsaraDB RDS instance from classic network to VPC
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Change the network type of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96109.html)
//
//   - [Change the network type of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96761.html)
//
//   - [Change the network type of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95707.html)
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

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ReadWriteSplittingClassicExpiredDays) {
		query["ReadWriteSplittingClassicExpiredDays"] = request.ReadWriteSplittingClassicExpiredDays
	}

	if !dara.IsNil(request.ReadWriteSplittingPrivateIpAddress) {
		query["ReadWriteSplittingPrivateIpAddress"] = request.ReadWriteSplittingPrivateIpAddress
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
		Action:      dara.String("ModifyDBInstanceNetworkType"),
		Version:     dara.String("2014-08-15"),
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
// Changes the billing method of an instance from pay-as-you-go to subscription.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// >  Fees of an instance are changed if the call is successful. Before you call this operation, carefully read the following topics:
//
//   - [Change the billing method of an ApsaraDB RDS for MySQL instance from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/96048.html)
//
//   - [Change the billing method of an ApsaraDB RDS for PostgreSQL instance from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/96743.html)
//
//   - [Change the billing method of an ApsaraDB RDS for SQL Server instance from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/95631.html)
//
//   - [Change the billing method of an ApsaraDB RDS for MariaDB instance from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/97120.html)
//
// @param request - ModifyDBInstancePayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstancePayTypeResponse
func (client *Client) ModifyDBInstancePayTypeWithContext(ctx context.Context, request *ModifyDBInstancePayTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstancePayTypeResponse, _err error) {
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

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstancePayType"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstancePayTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the native replication feature of ApsaraDB RDS for MySQL.
//
// Description:
//
// If you want to enable the native replication feature for an ApsaraDB RDS for MySQL instance, the following requirements must be met:
//
//   - The RDS instance runs MySQL 5.7.
//
//   - The RDS instance runs RDS Basic Edition.
//
//   - The RDS instance uses the pay-as-you-go or subscription billing method.
//
//   - The RDS instance runs a minor engine version of 20240930 or later.
//
// For more information, see [Enable the native replication feature](https://help.aliyun.com/document_detail/2856530.html).
//
// @param request - ModifyDBInstanceReplicationSwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceReplicationSwitchResponse
func (client *Client) ModifyDBInstanceReplicationSwitchWithContext(ctx context.Context, request *ModifyDBInstanceReplicationSwitchRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceReplicationSwitchResponse, _err error) {
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

	if !dara.IsNil(request.ExternalReplication) {
		query["ExternalReplication"] = request.ExternalReplication
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceReplicationSwitch"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceReplicationSwitchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the SSL encryption settings of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Use the SSL encryption feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96120.html)
//
//   - [Use the SSL encryption feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/229517.html)
//
//   - [Use the SSL encryption feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95715.html)
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
	if !dara.IsNil(request.ACL) {
		query["ACL"] = request.ACL
	}

	if !dara.IsNil(request.CAType) {
		query["CAType"] = request.CAType
	}

	if !dara.IsNil(request.Certificate) {
		query["Certificate"] = request.Certificate
	}

	if !dara.IsNil(request.ClientCACert) {
		query["ClientCACert"] = request.ClientCACert
	}

	if !dara.IsNil(request.ClientCAEnabled) {
		query["ClientCAEnabled"] = request.ClientCAEnabled
	}

	if !dara.IsNil(request.ClientCertRevocationList) {
		query["ClientCertRevocationList"] = request.ClientCertRevocationList
	}

	if !dara.IsNil(request.ClientCrlEnabled) {
		query["ClientCrlEnabled"] = request.ClientCrlEnabled
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

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

	if !dara.IsNil(request.PassWord) {
		query["PassWord"] = request.PassWord
	}

	if !dara.IsNil(request.ReplicationACL) {
		query["ReplicationACL"] = request.ReplicationACL
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SSLEnabled) {
		query["SSLEnabled"] = request.SSLEnabled
	}

	if !dara.IsNil(request.ServerCert) {
		query["ServerCert"] = request.ServerCert
	}

	if !dara.IsNil(request.ServerKey) {
		query["ServerKey"] = request.ServerKey
	}

	if !dara.IsNil(request.TlsVersion) {
		query["TlsVersion"] = request.TlsVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceSSL"),
		Version:     dara.String("2014-08-15"),
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
// Modifies a security group rule that is configured for an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// [Configure security group settings for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/2392322.html)
//
// @param request - ModifyDBInstanceSecurityGroupRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceSecurityGroupRuleResponse
func (client *Client) ModifyDBInstanceSecurityGroupRuleWithContext(ctx context.Context, request *ModifyDBInstanceSecurityGroupRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceSecurityGroupRuleResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityGroupRuleId) {
		query["SecurityGroupRuleId"] = request.SecurityGroupRuleId
	}

	if !dara.IsNil(request.SourceCidrIp) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceSecurityGroupRule"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceSecurityGroupRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the instance type and storage capacity of an ApsaraDB RDS instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)Billing details
//
// [Fees for specification changes](https://help.aliyun.com/document_detail/57178.html) are generated if the call is successful. Before you call this operation, carefully read the following topics.
//
// ### [](#)References
//
//   - [Change the specifications of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96061.html)
//
//   - [Change the specifications of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96750.html)
//
//   - [Change the specifications of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95665.html)
//
//   - [Change the specifications of an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97129.html)
//
// @param tmpReq - ModifyDBInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceSpecResponse
func (client *Client) ModifyDBInstanceSpecWithContext(ctx context.Context, tmpReq *ModifyDBInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyDBInstanceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServerlessConfiguration) {
		request.ServerlessConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfiguration, dara.String("ServerlessConfiguration"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowMajorVersionUpgrade) {
		query["AllowMajorVersionUpgrade"] = request.AllowMajorVersionUpgrade
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BurstingEnabled) {
		query["BurstingEnabled"] = request.BurstingEnabled
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.ColdDataEnabled) {
		query["ColdDataEnabled"] = request.ColdDataEnabled
	}

	if !dara.IsNil(request.CompressionMode) {
		query["CompressionMode"] = request.CompressionMode
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

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.IoAccelerationEnabled) {
		query["IoAccelerationEnabled"] = request.IoAccelerationEnabled
	}

	if !dara.IsNil(request.OptimizedWrites) {
		query["OptimizedWrites"] = request.OptimizedWrites
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.ReadOnlyDBInstanceClass) {
		query["ReadOnlyDBInstanceClass"] = request.ReadOnlyDBInstanceClass
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

	if !dara.IsNil(request.ServerlessConfigurationShrink) {
		query["ServerlessConfiguration"] = request.ServerlessConfigurationShrink
	}

	if !dara.IsNil(request.SourceBiz) {
		query["SourceBiz"] = request.SourceBiz
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.TargetMinorVersion) {
		query["TargetMinorVersion"] = request.TargetMinorVersion
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.ZoneIdSlave1) {
		query["ZoneIdSlave1"] = request.ZoneIdSlave1
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceSpec"),
		Version:     dara.String("2014-08-15"),
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
// Enables the Transparent Data Encryption (TDE) feature for an ApsaraDB RDS instance and modifies the TDE status for the instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Configure TDE for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96121.html)
//
//   - [Configure TDE for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/465652.html)
//
//   - [Configure TDE for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95716.html)
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
	if !dara.IsNil(request.Certificate) {
		query["Certificate"] = request.Certificate
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.EncryptionKey) {
		query["EncryptionKey"] = request.EncryptionKey
	}

	if !dara.IsNil(request.IsRotate) {
		query["IsRotate"] = request.IsRotate
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PassWord) {
		query["PassWord"] = request.PassWord
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      dara.String("ModifyDBInstanceTDE"),
		Version:     dara.String("2014-08-15"),
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
// 修改实例向量支持状态
//
// @param request - ModifyDBInstanceVectorSupportStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceVectorSupportStatusResponse
func (client *Client) ModifyDBInstanceVectorSupportStatusWithContext(ctx context.Context, request *ModifyDBInstanceVectorSupportStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceVectorSupportStatusResponse, _err error) {
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceVectorSupportStatus"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceVectorSupportStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the specifications, storage type, and storage capacity of an ApsaraDB RDS for MySQL instance that runs RDS Cluster Edition.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - MySQL
//
// ### [](#)References
//
// [Change instance specifications](https://help.aliyun.com/document_detail/2627998.html)
//
// >  Fees of an instance are changed if the call is successful. Before you call this operation, carefully read the related topics.
//
// @param tmpReq - ModifyDBNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBNodeResponse
func (client *Client) ModifyDBNodeWithContext(ctx context.Context, tmpReq *ModifyDBNodeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBNodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyDBNodeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBNode) {
		request.DBNodeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBNode, dara.String("DBNode"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceStorage) {
		query["DBInstanceStorage"] = request.DBInstanceStorage
	}

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.DBNodeShrink) {
		query["DBNode"] = request.DBNodeShrink
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProduceAsync) {
		query["ProduceAsync"] = request.ProduceAsync
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
		Action:      dara.String("ModifyDBNode"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBNodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or modifies the database proxy feature for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// >  Starting October 17, 2023, ApsaraDB RDS for MySQL instances that run RDS Cluster Edition offer one free-of-charge dedicated database proxy for each unit in phases. For more information, see [[Special offers/Price changes\\] One dedicated proxy is provided free of charge for ApsaraDB RDS for MySQL instances on RDS Cluster Edition](~~2555466~~).
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Enable the database proxy feature of ApsaraDB RDS for MySQL](https://help.aliyun.com/document_detail/197456.html)
//
//   - [Enable the database proxy feature of ApsaraDB RDS for PostgreSQL](https://help.aliyun.com/document_detail/418272.html)
//
// @param tmpReq - ModifyDBProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBProxyResponse
func (client *Client) ModifyDBProxyWithContext(ctx context.Context, tmpReq *ModifyDBProxyRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBProxyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyDBProxyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBProxyNodes) {
		request.DBProxyNodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBProxyNodes, dara.String("DBProxyNodes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigDBProxyService) {
		query["ConfigDBProxyService"] = request.ConfigDBProxyService
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
	}

	if !dara.IsNil(request.DBProxyInstanceNum) {
		query["DBProxyInstanceNum"] = request.DBProxyInstanceNum
	}

	if !dara.IsNil(request.DBProxyInstanceType) {
		query["DBProxyInstanceType"] = request.DBProxyInstanceType
	}

	if !dara.IsNil(request.DBProxyNodesShrink) {
		query["DBProxyNodes"] = request.DBProxyNodesShrink
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PersistentConnectionStatus) {
		query["PersistentConnectionStatus"] = request.PersistentConnectionStatus
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
		Action:      dara.String("ModifyDBProxy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBProxyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the connection settings for a database proxy endpoint.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Configure the connection settings for a database proxy endpoint for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/2621331.html)
//
//   - [Configure the connection settings for a database proxy endpoint for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/418273.html)
//
// @param request - ModifyDBProxyEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBProxyEndpointResponse
func (client *Client) ModifyDBProxyEndpointWithContext(ctx context.Context, request *ModifyDBProxyEndpointRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBProxyEndpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CausalConsistReadTimeout) {
		query["CausalConsistReadTimeout"] = request.CausalConsistReadTimeout
	}

	if !dara.IsNil(request.ConfigDBProxyFeatures) {
		query["ConfigDBProxyFeatures"] = request.ConfigDBProxyFeatures
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBProxyEndpointId) {
		query["DBProxyEndpointId"] = request.DBProxyEndpointId
	}

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
	}

	if !dara.IsNil(request.DbEndpointAliases) {
		query["DbEndpointAliases"] = request.DbEndpointAliases
	}

	if !dara.IsNil(request.DbEndpointMinSlaveCount) {
		query["DbEndpointMinSlaveCount"] = request.DbEndpointMinSlaveCount
	}

	if !dara.IsNil(request.DbEndpointOperator) {
		query["DbEndpointOperator"] = request.DbEndpointOperator
	}

	if !dara.IsNil(request.DbEndpointReadWriteMode) {
		query["DbEndpointReadWriteMode"] = request.DbEndpointReadWriteMode
	}

	if !dara.IsNil(request.DbEndpointType) {
		query["DbEndpointType"] = request.DbEndpointType
	}

	if !dara.IsNil(request.EffectiveSpecificTime) {
		query["EffectiveSpecificTime"] = request.EffectiveSpecificTime
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReadOnlyInstanceDistributionType) {
		query["ReadOnlyInstanceDistributionType"] = request.ReadOnlyInstanceDistributionType
	}

	if !dara.IsNil(request.ReadOnlyInstanceMaxDelayTime) {
		query["ReadOnlyInstanceMaxDelayTime"] = request.ReadOnlyInstanceMaxDelayTime
	}

	if !dara.IsNil(request.ReadOnlyInstanceWeight) {
		query["ReadOnlyInstanceWeight"] = request.ReadOnlyInstanceWeight
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
		Action:      dara.String("ModifyDBProxyEndpoint"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBProxyEndpointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the database proxy endpoint of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Configure the database proxy endpoint of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/184921.html)
//
//   - [Configure the database proxy endpoint of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/418274.html)
//
// @param request - ModifyDBProxyEndpointAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBProxyEndpointAddressResponse
func (client *Client) ModifyDBProxyEndpointAddressWithContext(ctx context.Context, request *ModifyDBProxyEndpointAddressRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBProxyEndpointAddressResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBProxyConnectStringNetType) {
		query["DBProxyConnectStringNetType"] = request.DBProxyConnectStringNetType
	}

	if !dara.IsNil(request.DBProxyEndpointId) {
		query["DBProxyEndpointId"] = request.DBProxyEndpointId
	}

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
	}

	if !dara.IsNil(request.DBProxyNewConnectString) {
		query["DBProxyNewConnectString"] = request.DBProxyNewConnectString
	}

	if !dara.IsNil(request.DBProxyNewConnectStringPort) {
		query["DBProxyNewConnectStringPort"] = request.DBProxyNewConnectStringPort
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
		Action:      dara.String("ModifyDBProxyEndpointAddress"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBProxyEndpointAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the configuration of a database proxy for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// >  Starting October 17, 2023, ApsaraDB RDS for MySQL instances that run RDS Cluster Edition offer one free-of-charge dedicated database proxy for each unit in phases. For more information, see [[Special offers/Price changes\\] One dedicated proxy is provided free of charge for ApsaraDB RDS for MySQL instances on RDS Cluster Edition](~~2555466~~).
//
// @param tmpReq - ModifyDBProxyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBProxyInstanceResponse
func (client *Client) ModifyDBProxyInstanceWithContext(ctx context.Context, tmpReq *ModifyDBProxyInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBProxyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyDBProxyInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBProxyNodes) {
		request.DBProxyNodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBProxyNodes, dara.String("DBProxyNodes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MigrateAZ) {
		request.MigrateAZShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MigrateAZ, dara.String("MigrateAZ"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
	}

	if !dara.IsNil(request.DBProxyInstanceNum) {
		query["DBProxyInstanceNum"] = request.DBProxyInstanceNum
	}

	if !dara.IsNil(request.DBProxyInstanceType) {
		query["DBProxyInstanceType"] = request.DBProxyInstanceType
	}

	if !dara.IsNil(request.DBProxyNodesShrink) {
		query["DBProxyNodes"] = request.DBProxyNodesShrink
	}

	if !dara.IsNil(request.EffectiveSpecificTime) {
		query["EffectiveSpecificTime"] = request.EffectiveSpecificTime
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.MigrateAZShrink) {
		query["MigrateAZ"] = request.MigrateAZShrink
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

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBProxyInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBProxyInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a distributed transaction whitelist for an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Configure a distributed transaction whitelist](https://help.aliyun.com/document_detail/124321.html)
//
// @param request - ModifyDTCSecurityIpHostsForSQLServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDTCSecurityIpHostsForSQLServerResponse
func (client *Client) ModifyDTCSecurityIpHostsForSQLServerWithContext(ctx context.Context, request *ModifyDTCSecurityIpHostsForSQLServerRequest, runtime *dara.RuntimeOptions) (_result *ModifyDTCSecurityIpHostsForSQLServerResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityIpHosts) {
		query["SecurityIpHosts"] = request.SecurityIpHosts
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WhiteListGroupName) {
		query["WhiteListGroupName"] = request.WhiteListGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDTCSecurityIpHostsForSQLServer"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDTCSecurityIpHostsForSQLServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures automatic storage expansion for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Configure automatic storage expansion for ApsaraDB RDS for MySQL](https://help.aliyun.com/document_detail/173826.html)
//
//   - [Configure automatic storage expansion for ApsaraDB RDS for PostgreSQL](https://help.aliyun.com/document_detail/432496.html)
//
// @param request - ModifyDasInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDasInstanceConfigResponse
func (client *Client) ModifyDasInstanceConfigWithContext(ctx context.Context, request *ModifyDasInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDasInstanceConfigResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StorageAutoScale) {
		query["StorageAutoScale"] = request.StorageAutoScale
	}

	if !dara.IsNil(request.StorageThreshold) {
		query["StorageThreshold"] = request.StorageThreshold
	}

	if !dara.IsNil(request.StorageUpperBound) {
		query["StorageUpperBound"] = request.StorageUpperBound
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDasInstanceConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDasInstanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the property settings of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - SQL Server
//
// ### [](#)References
//
// You can call this operation to modify the database properties of an ApsaraDB RDS for SQL Server instance and archive data from an instance that uses cloud disks to an Object Storage Service (OSS) bucket. Before you call this operation to archive data to an OSS bucket, you must enable the data archiving feature in the ApsaraDB RDS console. For more information, see [Modify database properties](https://help.aliyun.com/document_detail/2401398.html) and [Archive cloud disk data to an OSS bucket](https://help.aliyun.com/document_detail/2767189.html).
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
// @param request - ModifyDatabaseConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseConfigResponse
func (client *Client) ModifyDatabaseConfigWithContext(ctx context.Context, request *ModifyDatabaseConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatabaseConfigResponse, _err error) {
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

	if !dara.IsNil(request.DatabasePropertyName) {
		query["DatabasePropertyName"] = request.DatabasePropertyName
	}

	if !dara.IsNil(request.DatabasePropertyValue) {
		query["DatabasePropertyValue"] = request.DatabasePropertyValue
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
		Action:      dara.String("ModifyDatabaseConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatabaseConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures SSL encryption for an dedicated proxy endpoint of an ApsaraDB RDS for MySQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # RDS MySQL
//
// ### [](#)References
//
// > : Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation:
//
// [Configure SSL encryption for a proxy endpoint](https://help.aliyun.com/document_detail/188164.html)
//
// @param request - ModifyDbProxyInstanceSslRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDbProxyInstanceSslResponse
func (client *Client) ModifyDbProxyInstanceSslWithContext(ctx context.Context, request *ModifyDbProxyInstanceSslRequest, runtime *dara.RuntimeOptions) (_result *ModifyDbProxyInstanceSslResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.DbProxyConnectString) {
		query["DbProxyConnectString"] = request.DbProxyConnectString
	}

	if !dara.IsNil(request.DbProxyEndpointId) {
		query["DbProxyEndpointId"] = request.DbProxyEndpointId
	}

	if !dara.IsNil(request.DbProxySslEnabled) {
		query["DbProxySslEnabled"] = request.DbProxySslEnabled
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDbProxyInstanceSsl"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDbProxyInstanceSslResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about the events in the event center.
//
// @param request - ModifyEventInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEventInfoResponse
func (client *Client) ModifyEventInfoWithContext(ctx context.Context, request *ModifyEventInfoRequest, runtime *dara.RuntimeOptions) (_result *ModifyEventInfoResponse, _err error) {
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

	if !dara.IsNil(request.EventAction) {
		query["EventAction"] = request.EventAction
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEventInfo"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEventInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Change the availability check method of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [What is availability check?](https://help.aliyun.com/document_detail/207467.html)
//
// @param request - ModifyHADiagnoseConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHADiagnoseConfigResponse
func (client *Client) ModifyHADiagnoseConfigWithContext(ctx context.Context, request *ModifyHADiagnoseConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyHADiagnoseConfigResponse, _err error) {
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

	if !dara.IsNil(request.TcpConnectionType) {
		query["TcpConnectionType"] = request.TcpConnectionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHADiagnoseConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHADiagnoseConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the automatic primary/secondary switchover feature for an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Use the automatic primary/secondary switchover feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96054.html)
//
//   - [Use the automatic primary/secondary switchover feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96747.html)
//
//   - [Use the automatic primary/secondary switchover feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95659.html)
//
//   - [Use the automatic primary/secondary switchover feature for an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97127.html)
//
// @param request - ModifyHASwitchConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHASwitchConfigResponse
func (client *Client) ModifyHASwitchConfigWithContext(ctx context.Context, request *ModifyHASwitchConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyHASwitchConfigResponse, _err error) {
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

	if !dara.IsNil(request.HAConfig) {
		query["HAConfig"] = request.HAConfig
	}

	if !dara.IsNil(request.ManualHATime) {
		query["ManualHATime"] = request.ManualHATime
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
		Action:      dara.String("ModifyHASwitchConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHASwitchConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改上云任务
//
// @param request - ModifyImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyImportTaskResponse
func (client *Client) ModifyImportTaskWithContext(ctx context.Context, request *ModifyImportTaskRequest, runtime *dara.RuntimeOptions) (_result *ModifyImportTaskResponse, _err error) {
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

	if !dara.IsNil(request.Operation) {
		query["Operation"] = request.Operation
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyImportTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the auto-renewal settings of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
//	Notice: Fees are generated if the call is successful. Before you call this operation, carefully read the following topics:
//
//	- [Use the auto-renewal feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96049.html)
//
//	- [Use the auto-renewal feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96740.html)
//
//	- [Use the auto-renewal feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95635.html)
//
//	- [Use the auto-renewal feature for an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97121.html)
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
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
		Version:     dara.String("2014-08-15"),
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
// Modifies the cross-region backup settings of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Use the cross-region backup feature of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/120824.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206671.html)
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/187923.html)
//
// @param request - ModifyInstanceCrossBackupPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceCrossBackupPolicyResponse
func (client *Client) ModifyInstanceCrossBackupPolicyWithContext(ctx context.Context, request *ModifyInstanceCrossBackupPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceCrossBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupEnabled) {
		query["BackupEnabled"] = request.BackupEnabled
	}

	if !dara.IsNil(request.CrossBackupRegion) {
		query["CrossBackupRegion"] = request.CrossBackupRegion
	}

	if !dara.IsNil(request.CrossBackupType) {
		query["CrossBackupType"] = request.CrossBackupType
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.LogBackupEnabled) {
		query["LogBackupEnabled"] = request.LogBackupEnabled
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

	if !dara.IsNil(request.RetentType) {
		query["RetentType"] = request.RetentType
	}

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceCrossBackupPolicy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceCrossBackupPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改全密态规则
//
// @param tmpReq - ModifyMaskingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMaskingRulesResponse
func (client *Client) ModifyMaskingRulesWithContext(ctx context.Context, tmpReq *ModifyMaskingRulesRequest, runtime *dara.RuntimeOptions) (_result *ModifyMaskingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyMaskingRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RuleConfig) {
		request.RuleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuleConfig, dara.String("RuleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceName) {
		query["DBInstanceName"] = request.DBInstanceName
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	if !dara.IsNil(request.DefaultAlgo) {
		query["DefaultAlgo"] = request.DefaultAlgo
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.MaskingAlgo) {
		query["MaskingAlgo"] = request.MaskingAlgo
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

	if !dara.IsNil(request.RuleConfigShrink) {
		query["RuleConfig"] = request.RuleConfigShrink
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMaskingRules"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMaskingRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the pg_hba.conf file of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Connect an ApsaraDB RDS for PostgreSQL instance to a self-managed AD domain](https://help.aliyun.com/document_detail/349288.html)
//
//   - [The pg_hba.conf File](https://www.postgresql.org/docs/11/auth-pg-hba-conf.html)
//
// @param request - ModifyPGHbaConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPGHbaConfigResponse
func (client *Client) ModifyPGHbaConfigWithContext(ctx context.Context, request *ModifyPGHbaConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyPGHbaConfigResponse, _err error) {
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

	if !dara.IsNil(request.HbaItem) {
		query["HbaItem"] = request.HbaItem
	}

	if !dara.IsNil(request.OpsType) {
		query["OpsType"] = request.OpsType
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
		Action:      dara.String("ModifyPGHbaConfig"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPGHbaConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameter values of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Modify the parameters of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96063.html)
//
//   - [Modify the parameters of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96751.html)
//
//   - [Modify the parameters of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95667.html)
//
//   - [Modify the parameters of an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97130.html)
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

	if !dara.IsNil(request.Forcerestart) {
		query["Forcerestart"] = request.Forcerestart
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
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

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyParameter"),
		Version:     dara.String("2014-08-15"),
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
// Modifies the parameter template of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/130565.html)
//
//   - [Use a parameter template to configure the parameters of ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/457176.html)
//
// @param request - ModifyParameterGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParameterGroupResponse
func (client *Client) ModifyParameterGroupWithContext(ctx context.Context, request *ModifyParameterGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyParameterGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ParameterGroupDesc) {
		query["ParameterGroupDesc"] = request.ParameterGroupDesc
	}

	if !dara.IsNil(request.ParameterGroupId) {
		query["ParameterGroupId"] = request.ParameterGroupId
	}

	if !dara.IsNil(request.ParameterGroupName) {
		query["ParameterGroupName"] = request.ParameterGroupName
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
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
		Action:      dara.String("ModifyParameterGroup"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyParameterGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RDS MySQL修改参数定时任务
//
// @param request - ModifyParameterTimedScheduleTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyParameterTimedScheduleTaskResponse
func (client *Client) ModifyParameterTimedScheduleTaskWithContext(ctx context.Context, request *ModifyParameterTimedScheduleTaskRequest, runtime *dara.RuntimeOptions) (_result *ModifyParameterTimedScheduleTaskResponse, _err error) {
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

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyParameterTimedScheduleTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyParameterTimedScheduleTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改块存储属性
//
// @param request - ModifyRCDiskAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCDiskAttributeResponse
func (client *Client) ModifyRCDiskAttributeWithContext(ctx context.Context, request *ModifyRCDiskAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCDiskAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BurstingEnabled) {
		query["BurstingEnabled"] = request.BurstingEnabled
	}

	if !dara.IsNil(request.DeleteWithInstance) {
		query["DeleteWithInstance"] = request.DeleteWithInstance
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.DiskName) {
		query["DiskName"] = request.DiskName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCDiskAttribute"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCDiskAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改RDS用户磁盘付费类型
//
// @param request - ModifyRCDiskChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCDiskChargeTypeResponse
func (client *Client) ModifyRCDiskChargeTypeWithContext(ctx context.Context, request *ModifyRCDiskChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCDiskChargeTypeResponse, _err error) {
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

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCDiskChargeType"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCDiskChargeTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the disk type or performance level (PL) of the cloud disks of an RDS Custom instance.
//
// Description:
//
// >  To minimize the impacts on your business, we recommend that you change specifications during off-peak hours.
//
// Take note of the following items:
//
//   - For a pay-as-you-go Enterprise SSD (ESSD), you can upgrade or downgrade its PL. However, you cannot downgrade the performance level to PL0.
//
//   - The ESSD must be in the In Use (In_Use) or Unattached (Available) state.
//
//   - If the ESSD is attached to an instance, the instance must be in the Running or Stopped state. The instance cannot be in the Expired state or stopped due to overdue payments.
//
//   - The performance level of an ESSD is limited by the capacity of the ESSD. If you cannot upgrade the PL of an ESSD, you can expand the capacity of the ESSD.
//
// @param request - ModifyRCDiskSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCDiskSpecResponse
func (client *Client) ModifyRCDiskSpecWithContext(ctx context.Context, request *ModifyRCDiskSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCDiskSpecResponse, _err error) {
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

	if !dara.IsNil(request.DiskCategory) {
		query["DiskCategory"] = request.DiskCategory
	}

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.PerformanceLevel) {
		query["PerformanceLevel"] = request.PerformanceLevel
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCDiskSpec"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCDiskSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询RDS用户专属主机实例
//
// @param request - ModifyRCElasticScalingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCElasticScalingResponse
func (client *Client) ModifyRCElasticScalingWithContext(ctx context.Context, request *ModifyRCElasticScalingRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCElasticScalingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScalingEnabled) {
		query["ScalingEnabled"] = request.ScalingEnabled
	}

	if !dara.IsNil(request.ScheduledRule) {
		query["ScheduledRule"] = request.ScheduledRule
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCElasticScaling"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCElasticScalingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades the instance type of a subscription RDS Custom instance. The new instance type takes effect for the remaining lifecycle of the instance.
//
// Description:
//
// Before you call this operation, make sure that you are familiar with the billing methods, pricing, and refund rules of RDS Custom.
//
// Before you call this operation, take note of the following items:
//
//   - You cannot change the instance type of an expired instance. You can renew the instance and try again.
//
//   - When you downgrade the instance type of an instance, take note of the following items:
//
//   - The instance must be in the Stopped state.
//
//   - The price difference is refunded to the payment account you used. Vouchers that have been redeemed are not refundable.
//
//   - The operation is asynchronous. Wait 5 to 10 seconds for the instance type change to complete. Then, restart the instance by calling the RebootInstance operation or by using the console for the instance type change to take effect. If you restart only the operating system of the instance, the instance type change does not take effect. If the instance is in the Stopped state, you need only to start the instance. You do not need to restart the instance after it enters the Running state.
//
// @param request - ModifyRCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCInstanceResponse
func (client *Client) ModifyRCInstanceWithContext(ctx context.Context, request *ModifyRCInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCInstanceResponse, _err error) {
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

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RebootTime) {
		query["RebootTime"] = request.RebootTime
	}

	if !dara.IsNil(request.RebootWhenFinished) {
		query["RebootWhenFinished"] = request.RebootWhenFinished
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of an RDS Custom instance, such as the password, hostname, security groups, and whether release protection is enabled.
//
// @param tmpReq - ModifyRCInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCInstanceAttributeResponse
func (client *Client) ModifyRCInstanceAttributeWithContext(ctx context.Context, tmpReq *ModifyRCInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyRCInstanceAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SecurityGroupIds) {
		request.SecurityGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityGroupIds, dara.String("SecurityGroupIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Reboot) {
		query["Reboot"] = request.Reboot
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityGroupIdsShrink) {
		query["SecurityGroupIds"] = request.SecurityGroupIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCInstanceAttribute"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCInstanceAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the billing method of an RDS Custom instance or cloud disks. You can call this operation to change the billing method of instances between pay-as-you-go and subscription.
//
// Description:
//
// ### [](#)Precautions
//
//   - Before you call this operation, make sure that you are familiar with the subscription and pay-as-you-go billing methods and pricing of RDS Custom.
//
//   - The instances must be in the **Running*	- or **Stopped*	- state, and you have no overdue payments for the instances.
//
//   - The disk is in the **In_use*	- state and the billing method of the disk has not been changed within the previous 15 minutes.
//
//   - After you change the billing method, the payment is automatically completed. Make sure that the balance in your account is sufficient. Otherwise, your order becomes invalid and is canceled.
//
// ### [](#)Considerations
//
// For more information, see the following documentation:
//
//   - [Change the billing method of an instance](https://help.aliyun.com/document_detail/2878542.html)
//
//   - [Change the billing method of a disk](https://help.aliyun.com/document_detail/2878547.html)
//
// @param request - ModifyRCInstanceChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCInstanceChargeTypeResponse
func (client *Client) ModifyRCInstanceChargeTypeWithContext(ctx context.Context, request *ModifyRCInstanceChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCInstanceChargeTypeResponse, _err error) {
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

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.IncludeDataDisks) {
		query["IncludeDataDisks"] = request.IncludeDataDisks
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCInstanceChargeType"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCInstanceChargeTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of an RDS Custom instance.
//
// @param request - ModifyRCInstanceDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCInstanceDescriptionResponse
func (client *Client) ModifyRCInstanceDescriptionWithContext(ctx context.Context, request *ModifyRCInstanceDescriptionRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCInstanceDescriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceDescription) {
		query["InstanceDescription"] = request.InstanceDescription
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCInstanceDescription"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCInstanceDescriptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the key pair of an RDS Custom instance.
//
// @param request - ModifyRCInstanceKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCInstanceKeyPairResponse
func (client *Client) ModifyRCInstanceKeyPairWithContext(ctx context.Context, request *ModifyRCInstanceKeyPairRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCInstanceKeyPairResponse, _err error) {
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

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.Reboot) {
		query["Reboot"] = request.Reboot
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCInstanceKeyPair"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCInstanceKeyPairResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the public bandwidth of an RDS Custom for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # Custom for SQL Server
//
// @param request - ModifyRCInstanceNetworkSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCInstanceNetworkSpecResponse
func (client *Client) ModifyRCInstanceNetworkSpecWithContext(ctx context.Context, request *ModifyRCInstanceNetworkSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCInstanceNetworkSpecResponse, _err error) {
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

	if !dara.IsNil(request.InternetMaxBandwidthOut) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !dara.IsNil(request.NetworkChargeType) {
		query["NetworkChargeType"] = request.NetworkChargeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCInstanceNetworkSpec"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCInstanceNetworkSpecResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改RC实例安全组
//
// @param request - ModifyRCInstanceVpcAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCInstanceVpcAttributeResponse
func (client *Client) ModifyRCInstanceVpcAttributeWithContext(ctx context.Context, request *ModifyRCInstanceVpcAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCInstanceVpcAttributeResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCInstanceVpcAttribute"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCInstanceVpcAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改RC安全组规则
//
// @param request - ModifyRCSecurityGroupPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCSecurityGroupPermissionResponse
func (client *Client) ModifyRCSecurityGroupPermissionWithContext(ctx context.Context, request *ModifyRCSecurityGroupPermissionRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCSecurityGroupPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestCidrIp) {
		query["DestCidrIp"] = request.DestCidrIp
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityGroupRuleId) {
		query["SecurityGroupRuleId"] = request.SecurityGroupRuleId
	}

	if !dara.IsNil(request.SourceCidrIp) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !dara.IsNil(request.SourcePortRange) {
		query["SourcePortRange"] = request.SourcePortRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCSecurityGroupPermission"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCSecurityGroupPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改RCVCluster
//
// @param tmpReq - ModifyRCVClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRCVClusterResponse
func (client *Client) ModifyRCVClusterWithContext(ctx context.Context, tmpReq *ModifyRCVClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyRCVClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyRCVClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SupportDiskPerformanceLevel) {
		request.SupportDiskPerformanceLevelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SupportDiskPerformanceLevel, dara.String("SupportDiskPerformanceLevel"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SupportDiskPerformanceLevelShrink) {
		query["SupportDiskPerformanceLevel"] = request.SupportDiskPerformanceLevelShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRCVCluster"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRCVClusterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the latency threshold of the read/write splitting link and the read weights of a primary instance and its read-only instances.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The shared proxy feature is enabled for your ApsaraDB RDS for MySQL instance.
//
//   - The read/write splitting feature is enabled for your ApsaraDB RDS for MySQL instance.
//
//   - The instance must run one of the following database engine versions and RDS editions:
//
//   - MySQL 5.7 on RDS High-availability Edition with local disks
//
//   - MySQL 5.6
//
//   - SQL Server on RDS Cluster Edition
//
// @param request - ModifyReadWriteSplittingConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyReadWriteSplittingConnectionResponse
func (client *Client) ModifyReadWriteSplittingConnectionWithContext(ctx context.Context, request *ModifyReadWriteSplittingConnectionRequest, runtime *dara.RuntimeOptions) (_result *ModifyReadWriteSplittingConnectionResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DistributionType) {
		query["DistributionType"] = request.DistributionType
	}

	if !dara.IsNil(request.MaxDelayTime) {
		query["MaxDelayTime"] = request.MaxDelayTime
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyReadWriteSplittingConnection"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyReadWriteSplittingConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the latency at which a read-only ApsaraDB RDS for MySQL instance replicates data from its primary instance.
//
// Description:
//
// ### Supported database engines
//
// # RDS MySQL
//
// ### References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Set the data replication latency of a read-only ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96056.html)
//
// @param request - ModifyReadonlyInstanceDelayReplicationTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyReadonlyInstanceDelayReplicationTimeResponse
func (client *Client) ModifyReadonlyInstanceDelayReplicationTimeWithContext(ctx context.Context, request *ModifyReadonlyInstanceDelayReplicationTimeRequest, runtime *dara.RuntimeOptions) (_result *ModifyReadonlyInstanceDelayReplicationTimeResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReadSQLReplicationTime) {
		query["ReadSQLReplicationTime"] = request.ReadSQLReplicationTime
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
		Action:      dara.String("ModifyReadonlyInstanceDelayReplicationTime"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyReadonlyInstanceDelayReplicationTimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves an ApsaraDB RDS instance to a specified resource group.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Transfer resources across resource groups](https://help.aliyun.com/document_detail/94487.html)
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
		Action:      dara.String("ModifyResourceGroup"),
		Version:     dara.String("2014-08-15"),
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
// This operation can still be called but is no longer maintained. This operation enables or disables the SQL Explorer (SQL Audit) feature for an instance.
//
// Description:
//
// This operation can still be called but is no longer maintained. We recommend that you call the [ModifySqlLogConfig](https://help.aliyun.com/document_detail/2778835.html) operation instead of this operation.
//
// @param request - ModifySQLCollectorPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySQLCollectorPolicyResponse
func (client *Client) ModifySQLCollectorPolicyWithContext(ctx context.Context, request *ModifySQLCollectorPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifySQLCollectorPolicyResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SQLCollectorStatus) {
		query["SQLCollectorStatus"] = request.SQLCollectorStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySQLCollectorPolicy"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the retention period of the log files that are generated by the SQL Explorer feature for an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
// # RDS MySQL
//
// ### Precautions
//
// After you shorten the log backup retention period, log backup files that are stored longer than the specified log backup retention period are immediately deleted.
//
// ### References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Use the SQL Explorer and Audit feature](https://help.aliyun.com/document_detail/476574.html)
//
// @param request - ModifySQLCollectorRetentionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySQLCollectorRetentionResponse
func (client *Client) ModifySQLCollectorRetentionWithContext(ctx context.Context, request *ModifySQLCollectorRetentionRequest, runtime *dara.RuntimeOptions) (_result *ModifySQLCollectorRetentionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("ModifySQLCollectorRetention"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySQLCollectorRetentionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the ECS security groups to which an instance is added.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Configure a security group for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/201042.html)
//
//   - [Configure a security group for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/206310.html)
//
//   - [Configure a security group for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/2392322.html)
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
		Version:     dara.String("2014-08-15"),
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
// Modifies the IP address whitelist of an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Configure an IP address whitelist for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96118.html)
//
//   - [Configure an IP address whitelist for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/43187.html)
//
//   - [Configure an IP address whitelist for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/43186.html)
//
//   - [Configure an IP address whitelist for an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/90336.html)
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
	if !dara.IsNil(request.DBInstanceIPArrayAttribute) {
		query["DBInstanceIPArrayAttribute"] = request.DBInstanceIPArrayAttribute
	}

	if !dara.IsNil(request.DBInstanceIPArrayName) {
		query["DBInstanceIPArrayName"] = request.DBInstanceIPArrayName
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.FreshWhiteListReadins) {
		query["FreshWhiteListReadins"] = request.FreshWhiteListReadins
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityIPType) {
		query["SecurityIPType"] = request.SecurityIPType
	}

	if !dara.IsNil(request.SecurityIps) {
		query["SecurityIps"] = request.SecurityIps
	}

	if !dara.IsNil(request.WhitelistNetworkType) {
		query["WhitelistNetworkType"] = request.WhitelistNetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityIps"),
		Version:     dara.String("2014-08-15"),
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
// Modifies information about the historical tasks in the task center.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
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
		Version:     dara.String("2014-08-15"),
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
// Edits a whitelist. You can call this operation to create, modify, or delete a whitelist.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
// @param request - ModifyWhitelistTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWhitelistTemplateResponse
func (client *Client) ModifyWhitelistTemplateWithContext(ctx context.Context, request *ModifyWhitelistTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyWhitelistTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpWhitelist) {
		query["IpWhitelist"] = request.IpWhitelist
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWhitelistTemplate"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWhitelistTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除节点创建订单预检查
//
// @param tmpReq - PreCheckCreateOrderForDeleteDBNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreCheckCreateOrderForDeleteDBNodesResponse
func (client *Client) PreCheckCreateOrderForDeleteDBNodesWithContext(ctx context.Context, tmpReq *PreCheckCreateOrderForDeleteDBNodesRequest, runtime *dara.RuntimeOptions) (_result *PreCheckCreateOrderForDeleteDBNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PreCheckCreateOrderForDeleteDBNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DBNodeId) {
		request.DBNodeIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DBNodeId, dara.String("DBNodeId"), dara.String("json"))
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

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBNodeIdShrink) {
		query["DBNodeId"] = request.DBNodeIdShrink
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
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
		Action:      dara.String("PreCheckCreateOrderForDeleteDBNodes"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreCheckCreateOrderForDeleteDBNodesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether DuckDB-based analytical instances can be created for the specified RDS for PostgreSQL primary instance. If DuckDB-based analytical instances cannot be created, this operation returns the failure causes and provides solutions or recommended specification values.
//
// Description:
//
// ### [](#)Supported database engine
//
// # RDS PostgreSQL
//
// ### [](#)References
//
// [DuckDB-based analytical instance](https://help.aliyun.com/document_detail/2977241.html)
//
// @param request - PrecheckDuckDBDependencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrecheckDuckDBDependencyResponse
func (client *Client) PrecheckDuckDBDependencyWithContext(ctx context.Context, request *PrecheckDuckDBDependencyRequest, runtime *dara.RuntimeOptions) (_result *PrecheckDuckDBDependencyResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TargetMode) {
		query["TargetMode"] = request.TargetMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PrecheckDuckDBDependency"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PrecheckDuckDBDependencyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the on-premises logs of an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### Description
//
// The system automatically uploads log backup files to Object Storage Service (OSS) buckets. If the remaining storage of an instance is insufficient, you can call this operation to upload the log backup files of the instance to OSS buckets. After the upload is complete, the system deletes these files from the instance to release storage. This operation is called to upload log backup files from an instance to OSS buckets and then delete these files from the instance. If the instance runs SQL Server, transaction log backup files are compressed before they are uploaded.
//
// ### Precautions
//
//   - When you upload log backup files, the data restoration feature is not affected.
//
//   - This operation is called to release storage. The backup storage usage is not reduced.
//
//   - The OSS buckets to which log backup files are uploaded are provided by the system. You do not need to purchase these OSS buckets. In addition, you cannot access these OSS buckets.
//
// @param request - PurgeDBInstanceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurgeDBInstanceLogResponse
func (client *Client) PurgeDBInstanceLogWithContext(ctx context.Context, request *PurgeDBInstanceLogRequest, runtime *dara.RuntimeOptions) (_result *PurgeDBInstanceLogResponse, _err error) {
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
		Action:      dara.String("PurgeDBInstanceLog"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PurgeDBInstanceLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the notifications of an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### Feature description
//
// The notifications are highlighted at the top of the ApsaraDB RDS console. The notifications include renewal reminders and reminders of instance creation failures.
//
// After you call this operation to query notifications, you can call the [ConfirmNotify](https://help.aliyun.com/document_detail/610444.html) operation to mark the notifications as confirmed, which means that you understand the content of the notifications.
//
// @param request - QueryNotifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryNotifyResponse
func (client *Client) QueryNotifyWithContext(ctx context.Context, request *QueryNotifyRequest, runtime *dara.RuntimeOptions) (_result *QueryNotifyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		body["From"] = request.From
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.To) {
		body["To"] = request.To
	}

	if !dara.IsNil(request.WithConfirmed) {
		body["WithConfirmed"] = request.WithConfirmed
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryNotify"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryNotifyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// rds机器人热点问题
//
// @param request - QueryRecommendByCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecommendByCodeResponse
func (client *Client) QueryRecommendByCodeWithContext(ctx context.Context, request *QueryRecommendByCodeRequest, runtime *dara.RuntimeOptions) (_result *QueryRecommendByCodeResponse, _err error) {
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
		Action:      dara.String("QueryRecommendByCode"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecommendByCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建服务关联角色和开租
//
// @param request - RdsCustomInitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RdsCustomInitResponse
func (client *Client) RdsCustomInitWithContext(ctx context.Context, request *RdsCustomInitRequest, runtime *dara.RuntimeOptions) (_result *RdsCustomInitResponse, _err error) {
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

	if !dara.IsNil(request.ServiceLinkedRole) {
		query["ServiceLinkedRole"] = request.ServiceLinkedRole
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RdsCustomInit"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RdsCustomInitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an RDS Custom instance that is in the Running state.
//
// @param request - RebootRCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootRCInstanceResponse
func (client *Client) RebootRCInstanceWithContext(ctx context.Context, request *RebootRCInstanceRequest, runtime *dara.RuntimeOptions) (_result *RebootRCInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ForceStop) {
		query["ForceStop"] = request.ForceStop
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RebootTime) {
		query["RebootTime"] = request.RebootTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootRCInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootRCInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts multiple RDS Custom instances at a time.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// @param tmpReq - RebootRCInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootRCInstancesResponse
func (client *Client) RebootRCInstancesWithContext(ctx context.Context, tmpReq *RebootRCInstancesRequest, runtime *dara.RuntimeOptions) (_result *RebootRCInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RebootRCInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchOptimization) {
		query["BatchOptimization"] = request.BatchOptimization
	}

	if !dara.IsNil(request.ForceReboot) {
		query["ForceReboot"] = request.ForceReboot
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.RebootTime) {
		query["RebootTime"] = request.RebootTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootRCInstances"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootRCInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rebuilds the secondary instance of a primary instance in a dedicated cluster.
//
// Description:
//
// Dedicated clusters allow you to manage a number of instances at a time. You can create multiple dedicated clusters in a single region. Each dedicated cluster consists of multiple hosts. You can create multiple instances on each host. For more information, see [What is ApsaraDB MyBase?](https://help.aliyun.com/document_detail/141455.html)
//
// @param request - RebuildDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebuildDBInstanceResponse
func (client *Client) RebuildDBInstanceWithContext(ctx context.Context, request *RebuildDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *RebuildDBInstanceResponse, _err error) {
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

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.DedicatedHostId) {
		query["DedicatedHostId"] = request.DedicatedHostId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RebuildNodeType) {
		query["RebuildNodeType"] = request.RebuildNodeType
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
		Action:      dara.String("RebuildDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebuildDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-creates a data synchronization link for a disaster recovery ApsaraDB RDS instance.
//
// Description:
//
// ### [](#)Supported database engine
//
//   - PostgreSQL
//
// @param request - RebuildReplicationLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebuildReplicationLinkResponse
func (client *Client) RebuildReplicationLinkWithContext(ctx context.Context, request *RebuildReplicationLinkRequest, runtime *dara.RuntimeOptions) (_result *RebuildReplicationLinkResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebuildReplicationLink"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebuildReplicationLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches a primary instance to a disaster recovery instance or a disaster recovery instance to a primary instance.
//
// Description:
//
// The operation is phased out.
//
// @param request - ReceiveDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReceiveDBInstanceResponse
func (client *Client) ReceiveDBInstanceWithContext(ctx context.Context, request *ReceiveDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReceiveDBInstanceResponse, _err error) {
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

	if !dara.IsNil(request.GuardDBInstanceId) {
		query["GuardDBInstanceId"] = request.GuardDBInstanceId
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
		Action:      dara.String("ReceiveDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReceiveDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores backup data of an ApsaraDB RDS for SQL Server instance to an existing instance or a new instance.
//
// Description:
//
// ### [](#)Supported database engines
//
// # SQL Server 2012 or later
//
// ### [](#)References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Restore the data of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95722.html)
//
// @param request - RecoveryDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoveryDBInstanceResponse
func (client *Client) RecoveryDBInstanceWithContext(ctx context.Context, request *RecoveryDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *RecoveryDBInstanceResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.DBInstanceStorage) {
		query["DBInstanceStorage"] = request.DBInstanceStorage
	}

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.DbNames) {
		query["DbNames"] = request.DbNames
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.TargetDBInstanceId) {
		query["TargetDBInstanceId"] = request.TargetDBInstanceId
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
		Action:      dara.String("RecoveryDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecoveryDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Redeploys an RDS Custom instance when the instance receives a system event notification.
//
// Description:
//
// RedeployInstance is an **asynchronous*	- operation. It migrates data before it restarts the instance. If the instance is successfully redeployed, it enters the Running state. If the instance fails to be redeployed, it returns to the original physical server and the state before the redeployment.
//
// When you call this operation, take note of the following items:
//
// The instance must be in the Running or Stopped state. After the instance is redeployed, the state of the instance has the following changes:
//
//   - If the instance is in the Running state before redeployment, the instance enters the Stopped state.
//
//   - If the instance is in the Stopped state before redeployment, the instance enters the Starting state.
//
//   - If an instance receives notifications about simulated events that are created by calling the CreateSimulatedSystemEvent operation for the instance, you cannot call this operation to redeploy the instance.
//
// The following table lists the types and states of events that you can handle by calling the RedeployInstance operation.
//
//   - Instance redeployment due to system maintenance: SystemMaintenance.Redeploy. The event state is Inquiring or Scheduled.
//
//   - Instance redeployment due to system failures: SystemFailure.Redeploy. The event state is Inquiring.
//
// @param request - RedeployRCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RedeployRCInstanceResponse
func (client *Client) RedeployRCInstanceWithContext(ctx context.Context, request *RedeployRCInstanceRequest, runtime *dara.RuntimeOptions) (_result *RedeployRCInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceStop) {
		query["ForceStop"] = request.ForceStop
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RedeployRCInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RedeployRCInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
//   - [Release the public endpoint of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/26128.html)
//
//   - [Release the public endpoint of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/97738.html)
//
//   - [Release the public endpoint of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/97736.html)
//
//   - [Release the public endpoint of an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97740.html)
//
// @param request - ReleaseInstanceConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceConnectionResponse
func (client *Client) ReleaseInstanceConnectionWithContext(ctx context.Context, request *ReleaseInstanceConnectionRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceConnectionResponse, _err error) {
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

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
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
		Action:      dara.String("ReleaseInstanceConnection"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstanceConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Release the public endpoint of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/26128.html)
//
//   - [Release the public endpoint of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/97738.html)
//
//   - [Release the public endpoint of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/97736.html)
//
//   - [Release the public endpoint of an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97740.html)
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
		Action:      dara.String("ReleaseInstancePublicConnection"),
		Version:     dara.String("2014-08-15"),
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
// Releases the read/write splitting endpoint of an instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### Prerequisites
//
// Before you call this operation, make sure that the following requirements are met:
//
//   - The shared proxy feature is enabled for your ApsaraDB RDS for MySQL instance.
//
//   - The read/write splitting feature is enabled for the instance.
//
//   - The instance must run one of the following database engine versions and RDS editions:
//
//   - MySQL 5.7 on RDS High-availability Edition (with local disks)
//
//   - MySQL 5.6
//
//   - SQL Server on RDS Cluster Edition
//
// @param request - ReleaseReadWriteSplittingConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseReadWriteSplittingConnectionResponse
func (client *Client) ReleaseReadWriteSplittingConnectionWithContext(ctx context.Context, request *ReleaseReadWriteSplittingConnectionRequest, runtime *dara.RuntimeOptions) (_result *ReleaseReadWriteSplittingConnectionResponse, _err error) {
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

	if !dara.IsNil(request.RWAddressType) {
		query["RWAddressType"] = request.RWAddressType
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
		Action:      dara.String("ReleaseReadWriteSplittingConnection"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseReadWriteSplittingConnectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)Usage notes
//
//   - A maximum of 10 tags can be removed in a single request.
//
//   - If a tag is removed from all instances to which the tag is added, the tag is automatically deleted.
//
//   - If you specify only TagKey, all tags that match the TagKey condition are removed.
//
//   - You must specify at least TagKey or a pair of TagKey and TagValue.
//
// @param request - RemoveTagsFromResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTagsFromResourceResponse
func (client *Client) RemoveTagsFromResourceWithContext(ctx context.Context, request *RemoveTagsFromResourceRequest, runtime *dara.RuntimeOptions) (_result *RemoveTagsFromResourceResponse, _err error) {
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

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.ProxyId) {
		query["proxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTagsFromResource"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTagsFromResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually renews an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// >  Fees of an instance are changed if the call is successful. Before you call this operation, carefully read the following topics:
//
//   - [Manually renew an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96050.html)
//
//   - [Manually renew an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96741.html)
//
//   - [Manually renew an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95637.html)
//
//   - [Manually renew an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97122.html)
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithContext(ctx context.Context, request *RenewInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
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

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
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
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a subscription RDS Custom instance.
//
// @param request - RenewRCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewRCInstanceResponse
func (client *Client) RenewRCInstanceWithContext(ctx context.Context, request *RenewRCInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewRCInstanceResponse, _err error) {
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

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PeriodAlign) {
		query["PeriodAlign"] = request.PeriodAlign
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.TimeType) {
		query["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewRCInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewRCInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reinstalls the operating system (OS) of an RDS Custom instance.
//
// Description:
//
//	  The instance must be in the Stopped state.
//
//		- If you reinstall the system, the data on the original system disk is lost. Exercise caution when you perform this operation.
//
// @param request - ReplaceRCInstanceSystemDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceRCInstanceSystemDiskResponse
func (client *Client) ReplaceRCInstanceSystemDiskWithContext(ctx context.Context, request *ReplaceRCInstanceSystemDiskRequest, runtime *dara.RuntimeOptions) (_result *ReplaceRCInstanceSystemDiskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsLocalDisk) {
		query["IsLocalDisk"] = request.IsLocalDisk
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplaceRCInstanceSystemDisk"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplaceRCInstanceSystemDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the permissions of the privileged account.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Reset of the permissions of privileged accounts](https://help.aliyun.com/document_detail/140724.html)
//
// @param request - ResetAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountResponse
func (client *Client) ResetAccountWithContext(ctx context.Context, request *ResetAccountRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountResponse, _err error) {
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
		Action:      dara.String("ResetAccount"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of a database account.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Reset the password of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96100.html)
//
//   - [Reset the password of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96814.html)
//
//   - [Reset the password of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95691.html)
//
//   - [Reset the password of an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97133.html)
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
		Version:     dara.String("2014-08-15"),
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
// Expand the storage capacity of an RDS Custom instance.
//
// @param request - ResizeRCInstanceDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResizeRCInstanceDiskResponse
func (client *Client) ResizeRCInstanceDiskWithContext(ctx context.Context, request *ResizeRCInstanceDiskRequest, runtime *dara.RuntimeOptions) (_result *ResizeRCInstanceDiskResponse, _err error) {
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

	if !dara.IsNil(request.DiskId) {
		query["DiskId"] = request.DiskId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NewSize) {
		query["NewSize"] = request.NewSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResizeRCInstanceDisk"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResizeRCInstanceDiskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an instance.
//
// Description:
//
// ### Supported database engines
//
//   - RDS MySQL
//
//   - RDS PostgreSQL
//
//   - RDS SQL Server
//
//   - RDS MariaDB
//
// ### References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Restart an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96051.html)
//
//   - [Restart an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96798.html)
//
//   - [Restart an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95656.html)
//
//   - [Restart an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97472.html)
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
		Action:      dara.String("RestartDBInstance"),
		Version:     dara.String("2014-08-15"),
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
// Restores data to an existing instance across regions.
//
// Description:
//
// >  Before restoration, you can call the CheckCreateDdrDBInstance operation to check whether a cross-region backup set can be used for cross-region restoration.
//
// ### [](#)Supported database engines
//
// # MySQL
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Use the cross-region backup feature for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/120824.html)
//
//   - [Restore the data of an ApsaraDB RDS for MySQL instance across regions](https://help.aliyun.com/document_detail/120875.html)
//
// @param request - RestoreDdrTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreDdrTableResponse
func (client *Client) RestoreDdrTableWithContext(ctx context.Context, request *RestoreDdrTableRequest, runtime *dara.RuntimeOptions) (_result *RestoreDdrTableResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.RestoreType) {
		query["RestoreType"] = request.RestoreType
	}

	if !dara.IsNil(request.SourceDBInstanceName) {
		query["SourceDBInstanceName"] = request.SourceDBInstanceName
	}

	if !dara.IsNil(request.SourceRegion) {
		query["SourceRegion"] = request.SourceRegion
	}

	if !dara.IsNil(request.TableMeta) {
		query["TableMeta"] = request.TableMeta
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreDdrTable"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreDdrTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores individual databases or tables of an instance to the original instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Restore individual databases and tables of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/103175.html)
//
//   - [Restore individual databases and tables of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/613672.html)
//
// @param request - RestoreTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreTableResponse
func (client *Client) RestoreTableWithContext(ctx context.Context, request *RestoreTableRequest, runtime *dara.RuntimeOptions) (_result *RestoreTableResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.InstantRecovery) {
		query["InstantRecovery"] = request.InstantRecovery
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

	if !dara.IsNil(request.RestoreTime) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !dara.IsNil(request.TableMeta) {
		query["TableMeta"] = request.TableMeta
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreTable"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the permissions on a database of an ApsaraDB RDS instance from an account.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)Prerequisites
//
//   - The instance is in the Running state.
//
//   - The database is in the Running state.
//
// ### [](#)Precautions
//
//   - The permissions that can be revoked include SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, and TRIGGER.
//
//   - This operation is not supported for instances that run SQL Server 2017 on RDS Cluster Edition and run PostgreSQL.
//
// @param request - RevokeAccountPrivilegeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeAccountPrivilegeResponse
func (client *Client) RevokeAccountPrivilegeWithContext(ctx context.Context, request *RevokeAccountPrivilegeRequest, runtime *dara.RuntimeOptions) (_result *RevokeAccountPrivilegeResponse, _err error) {
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

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
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
		Action:      dara.String("RevokeAccountPrivilege"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeAccountPrivilegeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions from the service account of an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Grant permissions to the service account of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96102.html)
//
//   - [Grant permissions to the service account of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146887.html)
//
//   - [Grant permissions to the service account of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95693.html)
//
// @param request - RevokeOperatorPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeOperatorPermissionResponse
func (client *Client) RevokeOperatorPermissionWithContext(ctx context.Context, request *RevokeOperatorPermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeOperatorPermissionResponse, _err error) {
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
		Action:      dara.String("RevokeOperatorPermission"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeOperatorPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes security group rules with the specified IDs.
//
// @param tmpReq - RevokeRCSecurityGroupPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeRCSecurityGroupPermissionResponse
func (client *Client) RevokeRCSecurityGroupPermissionWithContext(ctx context.Context, tmpReq *RevokeRCSecurityGroupPermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeRCSecurityGroupPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RevokeRCSecurityGroupPermissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SecurityGroupRuleIdList) {
		request.SecurityGroupRuleIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityGroupRuleIdList, dara.String("SecurityGroupRuleIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SecurityGroupRuleIdListShrink) {
		query["SecurityGroupRuleIdList"] = request.SecurityGroupRuleIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeRCSecurityGroupPermission"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeRCSecurityGroupPermissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建并执行云助手命令
//
// @param tmpReq - RunRCCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunRCCommandResponse
func (client *Client) RunRCCommandWithContext(ctx context.Context, tmpReq *RunRCCommandRequest, runtime *dara.RuntimeOptions) (_result *RunRCCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunRCCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceTags) {
		request.ResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTags, dara.String("ResourceTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CommandContent) {
		query["CommandContent"] = request.CommandContent
	}

	if !dara.IsNil(request.ContainerId) {
		query["ContainerId"] = request.ContainerId
	}

	if !dara.IsNil(request.ContainerName) {
		query["ContainerName"] = request.ContainerName
	}

	if !dara.IsNil(request.ContentEncoding) {
		query["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableParameter) {
		query["EnableParameter"] = request.EnableParameter
	}

	if !dara.IsNil(request.Frequency) {
		query["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.KeepCommand) {
		query["KeepCommand"] = request.KeepCommand
	}

	if !dara.IsNil(request.Launcher) {
		query["Launcher"] = request.Launcher
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepeatMode) {
		query["RepeatMode"] = request.RepeatMode
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceTagsShrink) {
		query["ResourceTags"] = request.ResourceTagsShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TerminationMode) {
		query["TerminationMode"] = request.TerminationMode
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.WindowsPasswordName) {
		query["WindowsPasswordName"] = request.WindowsPasswordName
	}

	if !dara.IsNil(request.WorkingDir) {
		query["WorkingDir"] = request.WorkingDir
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunRCCommand"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunRCCommandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates one or more subscription RDS Custom instances. Before you call this operation, you must specify parameters such as ImageId, InstanceType, VSwitchId, and SecurityGroupId.
//
// Description:
//
//	  Before you create RDS Custom instances, you must submit a ticket to add your Alibaba Cloud account to a whitelist.
//
//		- You can create only subscription RDS Custom instances.
//
//		- Subscription RDS Custom instances are supported in the China (Shanghai), China (Shenzhen), China (Beijing), and China (Hangzhou) regions.
//
// @param tmpReq - RunRCInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunRCInstancesResponse
func (client *Client) RunRCInstancesWithContext(ctx context.Context, tmpReq *RunRCInstancesRequest, runtime *dara.RuntimeOptions) (_result *RunRCInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunRCInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateAckEdgeParam) {
		request.CreateAckEdgeParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateAckEdgeParam, dara.String("CreateAckEdgeParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DataDisk) {
		request.DataDiskShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataDisk, dara.String("DataDisk"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SystemDisk) {
		request.SystemDiskShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SystemDisk, dara.String("SystemDisk"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcuType) {
		query["AcuType"] = request.AcuType
	}

	if !dara.IsNil(request.Amount) {
		query["Amount"] = request.Amount
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CreateAckEdgeParamShrink) {
		query["CreateAckEdgeParam"] = request.CreateAckEdgeParamShrink
	}

	if !dara.IsNil(request.CreateExtraParam) {
		query["CreateExtraParam"] = request.CreateExtraParam
	}

	if !dara.IsNil(request.CreateMode) {
		query["CreateMode"] = request.CreateMode
	}

	if !dara.IsNil(request.DataDiskShrink) {
		query["DataDisk"] = request.DataDiskShrink
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.DeploymentSetId) {
		query["DeploymentSetId"] = request.DeploymentSetId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.InternetMaxBandwidthOut) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !dara.IsNil(request.IoOptimized) {
		query["IoOptimized"] = request.IoOptimized
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PasswordInherit) {
		query["PasswordInherit"] = request.PasswordInherit
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ScheduledRule) {
		query["ScheduledRule"] = request.ScheduledRule
	}

	if !dara.IsNil(request.SecurityEnhancementStrategy) {
		query["SecurityEnhancementStrategy"] = request.SecurityEnhancementStrategy
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SpotStrategy) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !dara.IsNil(request.SupportCase) {
		query["SupportCase"] = request.SupportCase
	}

	if !dara.IsNil(request.SystemDiskShrink) {
		query["SystemDisk"] = request.SystemDiskShrink
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.UserDataInBase64) {
		query["UserDataInBase64"] = request.UserDataInBase64
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
		Action:      dara.String("RunRCInstances"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunRCInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### [](#)References
//
// >  Before you call this operation, read the following topics and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Start an ApsaraDB RDS for SQL instance](https://help.aliyun.com/document_detail/462504.html)
//
// @param request - StartDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDBInstanceResponse
func (client *Client) StartDBInstanceWithContext(ctx context.Context, request *StartDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartDBInstanceResponse, _err error) {
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

	if !dara.IsNil(request.DBInstanceTransType) {
		query["DBInstanceTransType"] = request.DBInstanceTransType
	}

	if !dara.IsNil(request.DedicatedHostGroupId) {
		query["DedicatedHostGroupId"] = request.DedicatedHostGroupId
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
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

	if !dara.IsNil(request.SpecifiedTime) {
		query["SpecifiedTime"] = request.SpecifiedTime
	}

	if !dara.IsNil(request.Storage) {
		query["Storage"] = request.Storage
	}

	if !dara.IsNil(request.TargetDBInstanceClass) {
		query["TargetDBInstanceClass"] = request.TargetDBInstanceClass
	}

	if !dara.IsNil(request.TargetDedicatedHostIdForLog) {
		query["TargetDedicatedHostIdForLog"] = request.TargetDedicatedHostIdForLog
	}

	if !dara.IsNil(request.TargetDedicatedHostIdForMaster) {
		query["TargetDedicatedHostIdForMaster"] = request.TargetDedicatedHostIdForMaster
	}

	if !dara.IsNil(request.TargetDedicatedHostIdForSlave) {
		query["TargetDedicatedHostIdForSlave"] = request.TargetDedicatedHostIdForSlave
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
		Action:      dara.String("StartDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts RDS Custom instances that are in the Stopped state. After the operation is successfully called, the instances enter the Starting state.
//
// @param request - StartRCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRCInstanceResponse
func (client *Client) StartRCInstanceWithContext(ctx context.Context, request *StartRCInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartRCInstanceResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRCInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRCInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts multiple RDS Custom instances at a time.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// @param tmpReq - StartRCInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRCInstancesResponse
func (client *Client) StartRCInstancesWithContext(ctx context.Context, tmpReq *StartRCInstancesRequest, runtime *dara.RuntimeOptions) (_result *StartRCInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartRCInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchOptimization) {
		query["BatchOptimization"] = request.BatchOptimization
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRCInstances"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRCInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//	[Suspend an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/462504.html)
//
// @param request - StopDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDBInstanceResponse
func (client *Client) StopDBInstanceWithContext(ctx context.Context, request *StopDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopDBInstanceResponse, _err error) {
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
		Action:      dara.String("StopDBInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDBInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an RDS Custom instance that is in the Running state. After the operation is successfully called, the status of the RDS Custom instance changes from Stopping to Stopped.
//
// @param request - StopRCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRCInstanceResponse
func (client *Client) StopRCInstanceWithContext(ctx context.Context, request *StopRCInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopRCInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceStop) {
		query["ForceStop"] = request.ForceStop
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StoppedMode) {
		query["StoppedMode"] = request.StoppedMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRCInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRCInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops multiple RDS Custom instances at a time.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - RDS MySQL
//
//   - RDS SQL Server
//
// ### [](#)References
//
//   - [Introduction to RDS Custom for MySQL](https://help.aliyun.com/document_detail/2844223.html)
//
//   - [Introduction to RDS Custom for SQL Server](https://help.aliyun.com/document_detail/2864363.html)
//
// @param tmpReq - StopRCInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRCInstancesResponse
func (client *Client) StopRCInstancesWithContext(ctx context.Context, tmpReq *StopRCInstancesRequest, runtime *dara.RuntimeOptions) (_result *StopRCInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StopRCInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchOptimization) {
		query["BatchOptimization"] = request.BatchOptimization
	}

	if !dara.IsNil(request.ForceStop) {
		query["ForceStop"] = request.ForceStop
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StoppedMode) {
		query["StoppedMode"] = request.StoppedMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRCInstances"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRCInstancesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches workloads between primary and secondary ApsaraDB RDS instances.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Switch workloads between primary and secondary ApsaraDB RDS for MySQL instances](https://help.aliyun.com/document_detail/96054.html)
//
//   - [Switch workloads between primary and secondary ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/96747.html)
//
//   - [Switch workloads between primary and secondary ApsaraDB RDS for SQL Server instances](https://help.aliyun.com/document_detail/95659.html)
//
//   - [Switch workloads between primary and secondary ApsaraDB RDS for MariaDB instances](https://help.aliyun.com/document_detail/97127.html)
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

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
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
		Action:      dara.String("SwitchDBInstanceHA"),
		Version:     dara.String("2014-08-15"),
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
// Switches between internal and public endpoints of an instance in the classic network.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - SQL Server
//
// ### Prerequisites
//
//   - The instance is connected by using its internal or public endpoint.
//
//   - The instance is in the Running state.
//
//   - The number of times that you have switched the instance between its internal and public endpoints within the last 24 hours does not reach 20.
//
//   - The instance resides in the classic network.
//
// ### Usage notes
//
// After the endpoint that is used to connect to the instance is changed, you must update the endpoint information in the code of your application and restart the application.
//
// @param request - SwitchDBInstanceNetTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchDBInstanceNetTypeResponse
func (client *Client) SwitchDBInstanceNetTypeWithContext(ctx context.Context, request *SwitchDBInstanceNetTypeRequest, runtime *dara.RuntimeOptions) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
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

	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.ConnectionStringType) {
		query["ConnectionStringType"] = request.ConnectionStringType
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

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
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
		Action:      dara.String("SwitchDBInstanceNetType"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the virtual private cloud (VPC) and vSwitch for an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Change the VPC and vSwitch for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/137567.html)
//
//   - [Change the vSwitch for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146885.html)
//
//   - [Change the VPC and vSwitch for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/347675.html)
//
// @param request - SwitchDBInstanceVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchDBInstanceVpcResponse
func (client *Client) SwitchDBInstanceVpcWithContext(ctx context.Context, request *SwitchDBInstanceVpcRequest, runtime *dara.RuntimeOptions) (_result *SwitchDBInstanceVpcResponse, _err error) {
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

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Action:      dara.String("SwitchDBInstanceVpc"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchDBInstanceVpcResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a zero-downtime workload switchover after the major engine version upgrade of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// Supported database engine
//
//   - PostgreSQL
//
// @param request - SwitchOverMajorVersionUpgradeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchOverMajorVersionUpgradeResponse
func (client *Client) SwitchOverMajorVersionUpgradeWithContext(ctx context.Context, request *SwitchOverMajorVersionUpgradeRequest, runtime *dara.RuntimeOptions) (_result *SwitchOverMajorVersionUpgradeResponse, _err error) {
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

	if !dara.IsNil(request.SwitchoverTimeout) {
		query["SwitchoverTimeout"] = request.SwitchoverTimeout
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchOverMajorVersionUpgrade"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchOverMajorVersionUpgradeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the data synchronization link from a the primary ApsaraDB RDS for SQL Server instance to a disaster recovery (DR) instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// @param request - SwitchReplicationLinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchReplicationLinkResponse
func (client *Client) SwitchReplicationLinkWithContext(ctx context.Context, request *SwitchReplicationLinkRequest, runtime *dara.RuntimeOptions) (_result *SwitchReplicationLinkResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.TargetInstanceName) {
		query["TargetInstanceName"] = request.TargetInstanceName
	}

	if !dara.IsNil(request.TargetInstanceRegionId) {
		query["TargetInstanceRegionId"] = request.TargetInstanceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchReplicationLink"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchReplicationLinkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes a custom key pair to an RDS Custom instance. If you change the key pair that you created for your RDS Custom instance and you want the change to immediately take effect on the RDS Custom instance, you can call this operation to synchronize the new key pair to the RDS Custom instance. For example, you delete a key pair that has the same name as another key pair and recreate the key pair.
//
// @param request - SyncRCKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncRCKeyPairResponse
func (client *Client) SyncRCKeyPairWithContext(ctx context.Context, request *SyncRCKeyPairRequest, runtime *dara.RuntimeOptions) (_result *SyncRCKeyPairResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SyncMode) {
		query["SyncMode"] = request.SyncMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncRCKeyPair"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncRCKeyPairResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds security group rules to an RDS Custom for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// [Introduction to ApsaraDB RDS Custom](https://help.aliyun.com/document_detail/2864363.html)
//
// @param request - SyncRCSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncRCSecurityGroupResponse
func (client *Client) SyncRCSecurityGroupWithContext(ctx context.Context, request *SyncRCSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *SyncRCSecurityGroupResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncRCSecurityGroup"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncRCSecurityGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and adds tags to one or more instances.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Create tags for an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96149.html)
//
//   - [Create tags for an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96777.html)
//
//   - [Create tags for an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/95726.html)
//
//   - [Create tags for an ApsaraDB RDS for MariaDB instance](https://help.aliyun.com/document_detail/97152.html)
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
		Action:      dara.String("TagResources"),
		Version:     dara.String("2014-08-15"),
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
// Terminates an ongoing migration task of an ApsaraDB RDS for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - SQL Server
//
// @param request - TerminateMigrateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateMigrateTaskResponse
func (client *Client) TerminateMigrateTaskWithContext(ctx context.Context, request *TerminateMigrateTaskRequest, runtime *dara.RuntimeOptions) (_result *TerminateMigrateTaskResponse, _err error) {
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

	if !dara.IsNil(request.MigrateTaskId) {
		query["MigrateTaskId"] = request.MigrateTaskId
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
		Action:      dara.String("TerminateMigrateTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminateMigrateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### References
//
// > Fees are generated if the call is successful. Before you call this operation, you must read the following documentation.
//
//   - [Change the billing method of an ApsaraDB RDS for MySQL instance from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/96048.html) or [Change the billing method of an ApsaraDB RDS for MySQL instance from subscription to pay-as-you-go](https://help.aliyun.com/document_detail/161875.html)
//
//   - [Change the billing method of an ApsaraDB RDS for PostgreSQL instance from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/96743.html) or [Change the billing method of an ApsaraDB RDS for PostgreSQL instance from subscription to pay-as-you-go](https://help.aliyun.com/document_detail/162756.html)
//
//   - [Change the billing method of an ApsaraDB RDS for SQL Server instance from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/95631.html) or [Change the billing method of an ApsaraDB RDS for SQL Server instance from subscription to pay-as-you-go](https://help.aliyun.com/document_detail/162755.html)
//
//   - [Change the billing method of an ApsaraDB RDS for MariaDB instance from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/97120.html) or [Change the billing method of an ApsaraDB RDS for MariaDB instance from subscription to pay-as-you-go](https://help.aliyun.com/document_detail/169252.html)
//
// @param request - TransformDBInstancePayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformDBInstancePayTypeResponse
func (client *Client) TransformDBInstancePayTypeWithContext(ctx context.Context, request *TransformDBInstancePayTypeRequest, runtime *dara.RuntimeOptions) (_result *TransformDBInstancePayTypeResponse, _err error) {
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

	if !dara.IsNil(request.AutoUseCoupon) {
		query["AutoUseCoupon"] = request.AutoUseCoupon
	}

	if !dara.IsNil(request.BusinessInfo) {
		query["BusinessInfo"] = request.BusinessInfo
	}

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

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PromotionCode) {
		query["PromotionCode"] = request.PromotionCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransformDBInstancePayType"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransformDBInstancePayTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a elastic IP address (EIP) from an RDS Custom for SQL Server instance.
//
// Description:
//
// ### [](#)Supported database engine
//
// # SQL Server
//
// ### [](#)References
//
// [Introduction to ApsaraDB RDS Custom](https://help.aliyun.com/document_detail/2864363.html)
//
// @param request - UnassociateEipAddressWithRCInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassociateEipAddressWithRCInstanceResponse
func (client *Client) UnassociateEipAddressWithRCInstanceWithContext(ctx context.Context, request *UnassociateEipAddressWithRCInstanceRequest, runtime *dara.RuntimeOptions) (_result *UnassociateEipAddressWithRCInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllocationId) {
		query["AllocationId"] = request.AllocationId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnassociateEipAddressWithRCInstance"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnassociateEipAddressWithRCInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks a database account of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### Supported database engine
//
// # PostgreSQL
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Lock an account of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/147649.html)
//
// @param request - UnlockAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockAccountResponse
func (client *Client) UnlockAccountWithContext(ctx context.Context, request *UnlockAccountRequest, runtime *dara.RuntimeOptions) (_result *UnlockAccountResponse, _err error) {
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
		Action:      dara.String("UnlockAccount"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from one or more instances.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
//   - MariaDB
//
// ### [](#)Usage notes
//
//   - You can remove up to 20 tags at a time.
//
//   - If a tag is removed from an instance and is not added to other instances, the tag is automatically deleted.
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

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2014-08-15"),
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
// Updates the version of an extension on a database.
//
// Description:
//
// ### Supported database engines
//
// # RDS PostgreSQL
//
// ### References
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Manage extensions](https://help.aliyun.com/document_detail/2402409.html)
//
// @param request - UpdatePostgresExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePostgresExtensionsResponse
func (client *Client) UpdatePostgresExtensionsWithContext(ctx context.Context, request *UpdatePostgresExtensionsRequest, runtime *dara.RuntimeOptions) (_result *UpdatePostgresExtensionsResponse, _err error) {
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

	if !dara.IsNil(request.DBNames) {
		query["DBNames"] = request.DBNames
	}

	if !dara.IsNil(request.Extensions) {
		query["Extensions"] = request.Extensions
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
		Action:      dara.String("UpdatePostgresExtensions"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePostgresExtensionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description and retention period of a full backup file.
//
// Description:
//
// ### Supported database engines
//
// # RDS MySQL
//
// ### References
//
// A full backup file contains the data of a self-managed MySQL database. You can restore the data of a self-managed MySQL database from a full backup file to an ApsaraDB RDS for MySQL instance. For more information, see [Migrate the data of a self-managed MySQL 5.7 or MySQL 8.0 instance to an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/251779.html).
//
// > : Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// @param request - UpdateUserBackupFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserBackupFileResponse
func (client *Client) UpdateUserBackupFileWithContext(ctx context.Context, request *UpdateUserBackupFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserBackupFileResponse, _err error) {
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

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
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

	if !dara.IsNil(request.Retention) {
		query["Retention"] = request.Retention
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserBackupFile"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserBackupFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the major engine version of an ApsaraDB RDS for MySQL instance.
//
// Description:
//
// ### Supported database engine
//
// # MySQL
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
// [Upgrade the major engine version of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96058.html)
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
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EffectiveTime) {
		query["EffectiveTime"] = request.EffectiveTime
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceEngineVersion"),
		Version:     dara.String("2014-08-15"),
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
// Updates the minor engine version of an ApsaraDB RDS instance.
//
// Description:
//
// ### Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
//   - SQL Server
//
// ### References
//
// > Before you call this operation, read the following documentation and make sure that you fully understand the prerequisites and impacts of this operation.
//
//   - [Update the minor engine version of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/96059.html)
//
//   - [Update the minor engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/146895.html)
//
//   - [Update the minor engine version of an ApsaraDB RDS for SQL Server instance](https://help.aliyun.com/document_detail/213582.html)
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.TargetMinorVersion) {
		query["TargetMinorVersion"] = request.TargetMinorVersion
	}

	if !dara.IsNil(request.UpgradeTime) {
		query["UpgradeTime"] = request.UpgradeTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceKernelVersion"),
		Version:     dara.String("2014-08-15"),
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

// Summary:
//
// Initiates a task to upgrade the major engine version of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### Supported database engine
//
// # PostgreSQL
//
// ### References
//
// Fees are generated if the call is successful. Before you call this operation, read the following documentation and make sure that you fully understand the billing rules, prerequisites, and impacts of this operation.
//
// [Upgrade the major engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/203309.html)
//
// @param request - UpgradeDBInstanceMajorVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceMajorVersionResponse
func (client *Client) UpgradeDBInstanceMajorVersionWithContext(ctx context.Context, request *UpgradeDBInstanceMajorVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBInstanceMajorVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowDDL) {
		query["AllowDDL"] = request.AllowDDL
	}

	if !dara.IsNil(request.CollectStatMode) {
		query["CollectStatMode"] = request.CollectStatMode
	}

	if !dara.IsNil(request.CustomExtraInfo) {
		query["CustomExtraInfo"] = request.CustomExtraInfo
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

	if !dara.IsNil(request.DBInstanceStorageType) {
		query["DBInstanceStorageType"] = request.DBInstanceStorageType
	}

	if !dara.IsNil(request.InstanceNetworkType) {
		query["InstanceNetworkType"] = request.InstanceNetworkType
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.PrivateIpAddress) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SwitchOver) {
		query["SwitchOver"] = request.SwitchOver
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	if !dara.IsNil(request.TargetMajorVersion) {
		query["TargetMajorVersion"] = request.TargetMajorVersion
	}

	if !dara.IsNil(request.UpgradeMode) {
		query["UpgradeMode"] = request.UpgradeMode
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

	if !dara.IsNil(request.ZoneIdSlave1) {
		query["ZoneIdSlave1"] = request.ZoneIdSlave1
	}

	if !dara.IsNil(request.ZoneIdSlave2) {
		query["ZoneIdSlave2"] = request.ZoneIdSlave2
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceMajorVersion"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBInstanceMajorVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a precheck before the upgrade of the major engine version of an ApsaraDB RDS for PostgreSQL instance.
//
// Description:
//
// ### Supported database engine
//
// # PostgreSQL
//
// ### References
//
// > Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
// [Upgrade the major engine version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/203309.html)
//
// @param request - UpgradeDBInstanceMajorVersionPrecheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceMajorVersionPrecheckResponse
func (client *Client) UpgradeDBInstanceMajorVersionPrecheckWithContext(ctx context.Context, request *UpgradeDBInstanceMajorVersionPrecheckRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBInstanceMajorVersionPrecheckResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TargetMajorVersion) {
		query["TargetMajorVersion"] = request.TargetMajorVersion
	}

	if !dara.IsNil(request.UpgradeMode) {
		query["UpgradeMode"] = request.UpgradeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceMajorVersionPrecheck"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBInstanceMajorVersionPrecheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the database proxy version of an instance.
//
// Description:
//
// ### [](#)Supported database engines
//
//   - MySQL
//
//   - PostgreSQL
//
// ### [](#)References
//
// >  Before you call this operation, carefully read the following documentation. Make sure that you fully understand the prerequisites and impacts for calling this operation.
//
//   - [Upgrade the database proxy version of an ApsaraDB RDS for MySQL instance](https://help.aliyun.com/document_detail/197465.html)
//
//   - [Upgrade the database proxy version of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/418469.html)
//
// @param request - UpgradeDBProxyInstanceKernelVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBProxyInstanceKernelVersionResponse
func (client *Client) UpgradeDBProxyInstanceKernelVersionWithContext(ctx context.Context, request *UpgradeDBProxyInstanceKernelVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBProxyInstanceKernelVersionResponse, _err error) {
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

	if !dara.IsNil(request.DBProxyEngineType) {
		query["DBProxyEngineType"] = request.DBProxyEngineType
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

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.TargetMinorVersion) {
		query["TargetMinorVersion"] = request.TargetMinorVersion
	}

	if !dara.IsNil(request.UpgradeTime) {
		query["UpgradeTime"] = request.UpgradeTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBProxyInstanceKernelVersion"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBProxyInstanceKernelVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预检验数据导入任务参数
//
// @param request - ValidateImportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateImportTaskResponse
func (client *Client) ValidateImportTaskWithContext(ctx context.Context, request *ValidateImportTaskRequest, runtime *dara.RuntimeOptions) (_result *ValidateImportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.EstimatedSize) {
		query["EstimatedSize"] = request.EstimatedSize
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !dara.IsNil(request.SourcePlatform) {
		query["SourcePlatform"] = request.SourcePlatform
	}

	if !dara.IsNil(request.StreamPort) {
		query["StreamPort"] = request.StreamPort
	}

	if !dara.IsNil(request.User) {
		query["User"] = request.User
	}

	if !dara.IsNil(request.XtrabackupPath) {
		query["XtrabackupPath"] = request.XtrabackupPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateImportTask"),
		Version:     dara.String("2014-08-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateImportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
