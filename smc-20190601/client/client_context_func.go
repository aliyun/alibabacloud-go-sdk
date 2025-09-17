// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Associates multiple migration sources with a workgroup.
//
// Description:
//
// A migration source can be associated with only one workgroup.
//
// @param request - AssociateSourceServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateSourceServersResponse
func (client *Client) AssociateSourceServersWithContext(ctx context.Context, request *AssociateSourceServersRequest, runtime *dara.RuntimeOptions) (_result *AssociateSourceServersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.WorkgroupId) {
		query["WorkgroupId"] = request.WorkgroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateSourceServers"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateSourceServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an activation code.
//
// Description:
//
// If you want to import the information of migration sources by using an activation code, you can call this operation to create one.
//
// @param request - CreateAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessTokenResponse
func (client *Client) CreateAccessTokenWithContext(ctx context.Context, request *CreateAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["Count"] = request.Count
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.TimeToLiveInDays) {
		query["TimeToLiveInDays"] = request.TimeToLiveInDays
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessToken"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Server Migration Center (SMC) allows you to migrate Alibaba Cloud ECS instances across zones in the same region. You can also change the instance type (vCPU and memory) within the same instance family to meet your business requirements. You can use this API to create a cross-zone migration job.
//
// Description:
//
// For more information about the limits and impacts of cross-zone migration, see [Cross-zone ECS instance migration](https://help.aliyun.com/document_detail/476797.html).
//
// @param request - CreateCrossZoneMigrationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCrossZoneMigrationJobResponse
func (client *Client) CreateCrossZoneMigrationJobWithContext(ctx context.Context, request *CreateCrossZoneMigrationJobRequest, runtime *dara.RuntimeOptions) (_result *CreateCrossZoneMigrationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Disk) {
		query["Disk"] = request.Disk
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

	if !dara.IsNil(request.TargetInstanceType) {
		query["TargetInstanceType"] = request.TargetInstanceType
	}

	if !dara.IsNil(request.TargetVSwitchId) {
		query["TargetVSwitchId"] = request.TargetVSwitchId
	}

	if !dara.IsNil(request.TargetZoneId) {
		query["TargetZoneId"] = request.TargetZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCrossZoneMigrationJob"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCrossZoneMigrationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a migration job for a source server.
//
// Description:
//
// ## Usage notes
//
//   - You can create migration jobs only for source servers that are in the Available state.
//
//   - Each source server can be associated with only one migration job that is in the Ready, Running, Stopped, Waiting, InError, or Expired state.
//
//   - You can create a maximum of 1,000 migration jobs within each Alibaba Cloud account.
//
//   - If you migrate a source server to an image, you must specify the ImageName, SystemDiskSize, and DataDisk parameters.
//
//   - If you use a virtual private cloud (VPC) to migrate data, the VSwitchId parameter is required and the VpcId parameter is optional.
//
//   - Server Migration Center (SMC) allows you to migrate source servers to Docker container images. This allows you to migrate containerized applications in a cost-effective way.
//
// @param request - CreateReplicationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReplicationJobResponse
func (client *Client) CreateReplicationJobWithContext(ctx context.Context, request *CreateReplicationJobRequest, runtime *dara.RuntimeOptions) (_result *CreateReplicationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ContainerNamespace) {
		query["ContainerNamespace"] = request.ContainerNamespace
	}

	if !dara.IsNil(request.ContainerRepository) {
		query["ContainerRepository"] = request.ContainerRepository
	}

	if !dara.IsNil(request.ContainerTag) {
		query["ContainerTag"] = request.ContainerTag
	}

	if !dara.IsNil(request.DataDisk) {
		query["DataDisk"] = request.DataDisk
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Disks) {
		query["Disks"] = request.Disks
	}

	if !dara.IsNil(request.Frequency) {
		query["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceRamRole) {
		query["InstanceRamRole"] = request.InstanceRamRole
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.LaunchTemplateId) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !dara.IsNil(request.LaunchTemplateVersion) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !dara.IsNil(request.LicenseType) {
		query["LicenseType"] = request.LicenseType
	}

	if !dara.IsNil(request.MaxNumberOfImageToKeep) {
		query["MaxNumberOfImageToKeep"] = request.MaxNumberOfImageToKeep
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetMode) {
		query["NetMode"] = request.NetMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplicationParameters) {
		query["ReplicationParameters"] = request.ReplicationParameters
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.RunOnce) {
		query["RunOnce"] = request.RunOnce
	}

	if !dara.IsNil(request.ScheduledStartTime) {
		query["ScheduledStartTime"] = request.ScheduledStartTime
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SystemDiskPart) {
		query["SystemDiskPart"] = request.SystemDiskPart
	}

	if !dara.IsNil(request.SystemDiskSize) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ValidTime) {
		query["ValidTime"] = request.ValidTime
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReplicationJob"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReplicationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workgroup. You can create a workgroup to manage the lifecycles of multiple migration tasks at a time. This is suitable for scenarios in which multiple servers are migrated.
//
// Description:
//
//	  You can create up to 50 workgroups within an Alibaba Cloud account.
//
//		- A workgroup can be associated with a maximum of 50 migration sources.
//
//		- A migration source can be associated with only one workgroup.
//
// @param request - CreateWorkgroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkgroupResponse
func (client *Client) CreateWorkgroupWithContext(ctx context.Context, request *CreateWorkgroupRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkgroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkgroup"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkgroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an incremental migration job that periodically runs. After you call this operation to stop an incremental migration job, the migration job is complete.
//
// Description:
//
// ## Usage notes
//
//   - The incremental migration job must be in the Waiting state.
//
//   - After you call this operation, the incremental migration job no longer periodically runs. In the meantime, Server Migration Center (SMC) determines whether to perform a full data migration for the last time based on the value of the `SyncData` parameter. If you set the SyncData parameter to `false`, SMC releases intermediate resources without data migration before the migration job is complete. If you set the SyncData parameter to `true`, SMC performs a full data migration and releases intermediate resources before the migration job is complete.
//
// @param request - CutOverReplicationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CutOverReplicationJobResponse
func (client *Client) CutOverReplicationJobWithContext(ctx context.Context, request *CutOverReplicationJobRequest, runtime *dara.RuntimeOptions) (_result *CutOverReplicationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.SyncData) {
		query["SyncData"] = request.SyncData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CutOverReplicationJob"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CutOverReplicationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an activation code.
//
// Description:
//
// You can call this operation to delete an activation code if you no longer need to import the information about migration sources by using the activation code or if the activation code has expired.
//
// @param request - DeleteAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessTokenResponse
func (client *Client) DeleteAccessTokenWithContext(ctx context.Context, request *DeleteAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessTokenId) {
		query["AccessTokenId"] = request.AccessTokenId
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
		Action:      dara.String("DeleteAccessToken"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a migration job.
//
// Description:
//
// ## Usage notes
//
//   - Deleted migration jobs cannot be restored.
//
//   - After a migration job is deleted, associated resources, such as the intermediate instance, are automatically released.
//
// @param request - DeleteReplicationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReplicationJobResponse
func (client *Client) DeleteReplicationJobWithContext(ctx context.Context, request *DeleteReplicationJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteReplicationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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
		Action:      dara.String("DeleteReplicationJob"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteReplicationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a migration source.
//
// Description:
//
// ## Usage notes
//
//   - If a migration job is created for the migration source and the migration job is in the Running state, the migration source cannot be deleted.
//
//   - If a migration job is created for the migration source but the migration job is not in the Running state, you can set the `Force` parameter to true to delete the migration source.
//
// @param request - DeleteSourceServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSourceServerResponse
func (client *Client) DeleteSourceServerWithContext(ctx context.Context, request *DeleteSourceServerRequest, runtime *dara.RuntimeOptions) (_result *DeleteSourceServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSourceServer"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSourceServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workgroup.
//
// Description:
//
// To delete a workgroup, you must delete or dissociate the migration source that is associated with the workgroup. For more information, see [DeleteSourceServer](https://help.aliyun.com/document_detail/2402124.html).
//
// @param request - DeleteWorkgroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkgroupResponse
func (client *Client) DeleteWorkgroupWithContext(ctx context.Context, request *DeleteWorkgroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkgroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.WorkgroupId) {
		query["WorkgroupId"] = request.WorkgroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkgroup"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkgroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of migration jobs.
//
// Description:
//
// ## Usage notes
//
//   - You can specify multiple request parameters to be queried. Specified parameters are evaluated by using the AND operator. Only the specified parameters are used as filter conditions.
//
//   - Server Migration Center (SMC) allows you to migrate source servers to Docker container images. You can use SMC to migrate containerized applications in a cost-effective way. For more information, see [Terms](https://help.aliyun.com/document_detail/60744.html).
//
// @param request - DescribeReplicationJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeReplicationJobsResponse
func (client *Client) DescribeReplicationJobsWithContext(ctx context.Context, request *DescribeReplicationJobsRequest, runtime *dara.RuntimeOptions) (_result *DescribeReplicationJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessStatus) {
		query["BusinessStatus"] = request.BusinessStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeReplicationJobs"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeReplicationJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more source servers.
//
// Description:
//
// ## [](#)Usage notes
//
// You can specify multiple request parameters to filter instances. Specified parameters have logical AND relations. Only the specified parameters are used as filter conditions.
//
// @param request - DescribeSourceServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSourceServersResponse
func (client *Client) DescribeSourceServersWithContext(ctx context.Context, request *DescribeSourceServersRequest, runtime *dara.RuntimeOptions) (_result *DescribeSourceServersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	if !dara.IsNil(request.RelatedJobType) {
		query["RelatedJobType"] = request.RelatedJobType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.WorkgroupId) {
		query["WorkgroupId"] = request.WorkgroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSourceServers"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSourceServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about workgroups. After you create a workgroup, you can query the information about the workgroup, such as the name, description, and alert information.
//
// @param request - DescribeWorkgroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWorkgroupsResponse
func (client *Client) DescribeWorkgroupsWithContext(ctx context.Context, request *DescribeWorkgroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeWorkgroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.WorkgroupId) {
		query["WorkgroupId"] = request.WorkgroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWorkgroups"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWorkgroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an activation code.
//
// Description:
//
// To prevent an activation code from being leaked, you can call this operation to disable the activation code. Disabled activation codes can no longer be used to import the information about migration sources. However, migration sources whose information has been imported are not affected.
//
// @param request - DisableAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAccessTokenResponse
func (client *Client) DisableAccessTokenWithContext(ctx context.Context, request *DisableAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *DisableAccessTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessTokenId) {
		query["AccessTokenId"] = request.AccessTokenId
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
		Action:      dara.String("DisableAccessToken"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAccessTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates migration sources from a workgroup. If you do not need to use a workgroup to migrate migration sources, you can disassociate the migration sources from the workgroup.
//
// @param request - DisassociateSourceServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateSourceServersResponse
func (client *Client) DisassociateSourceServersWithContext(ctx context.Context, request *DisassociateSourceServersRequest, runtime *dara.RuntimeOptions) (_result *DisassociateSourceServersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.WorkgroupId) {
		query["WorkgroupId"] = request.WorkgroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateSourceServers"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateSourceServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries activation codes and the usage details of the activation codes.
//
// Description:
//
// You can call this operation to query activation codes and the usage details of the activation codes. An activation code can be in the activated, unactivated, or expired state.
//
// @param request - ListAccessTokensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessTokensResponse
func (client *Client) ListAccessTokensWithContext(ctx context.Context, request *ListAccessTokensRequest, runtime *dara.RuntimeOptions) (_result *ListAccessTokensResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessTokenId) {
		query["AccessTokenId"] = request.AccessTokenId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessTokens"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessTokensResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to Server Migration Center (SMC) resources. SMC resources include migration sources and migration jobs.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
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
		Version:     dara.String("2019-06-01"),
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
// Modifies the parameters of a migration job.
//
// Description:
//
// ## Usage notes
//
// Before you modify the parameters of a migration job, take note of the following information:
//
//   - The `Name` and `Description` parameters can be modified during the lifecycle of the migration job.
//
//   - The `Frequency` and `MaxNumberOfImageToKeep` parameters can be modified only before the migration job is executed or when the migration job is in the `Waiting` state.
//
//   - Other parameters can be modified only before the migration job is executed.
//
// @param request - ModifyReplicationJobAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyReplicationJobAttributeResponse
func (client *Client) ModifyReplicationJobAttributeWithContext(ctx context.Context, request *ModifyReplicationJobAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyReplicationJobAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContainerNamespace) {
		query["ContainerNamespace"] = request.ContainerNamespace
	}

	if !dara.IsNil(request.ContainerRepository) {
		query["ContainerRepository"] = request.ContainerRepository
	}

	if !dara.IsNil(request.ContainerTag) {
		query["ContainerTag"] = request.ContainerTag
	}

	if !dara.IsNil(request.DataDisk) {
		query["DataDisk"] = request.DataDisk
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Frequency) {
		query["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceRamRole) {
		query["InstanceRamRole"] = request.InstanceRamRole
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.LaunchTemplateId) {
		query["LaunchTemplateId"] = request.LaunchTemplateId
	}

	if !dara.IsNil(request.LaunchTemplateVersion) {
		query["LaunchTemplateVersion"] = request.LaunchTemplateVersion
	}

	if !dara.IsNil(request.MaxNumberOfImageToKeep) {
		query["MaxNumberOfImageToKeep"] = request.MaxNumberOfImageToKeep
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetMode) {
		query["NetMode"] = request.NetMode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReplicationParameters) {
		query["ReplicationParameters"] = request.ReplicationParameters
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ScheduledStartTime) {
		query["ScheduledStartTime"] = request.ScheduledStartTime
	}

	if !dara.IsNil(request.SystemDiskPart) {
		query["SystemDiskPart"] = request.SystemDiskPart
	}

	if !dara.IsNil(request.SystemDiskSize) {
		query["SystemDiskSize"] = request.SystemDiskSize
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ValidTime) {
		query["ValidTime"] = request.ValidTime
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyReplicationJobAttribute"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyReplicationJobAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a migration source.
//
// Description:
//
// ## Usage notes
//
// You can call this operation regardless of the status of the migration source.
//
// @param request - ModifySourceServerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySourceServerAttributeResponse
func (client *Client) ModifySourceServerAttributeWithContext(ctx context.Context, request *ModifySourceServerAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifySourceServerAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySourceServerAttribute"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySourceServerAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a workgroup.
//
// @param request - ModifyWorkgroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyWorkgroupAttributeResponse
func (client *Client) ModifyWorkgroupAttributeWithContext(ctx context.Context, request *ModifyWorkgroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyWorkgroupAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.WorkgroupId) {
		query["WorkgroupId"] = request.WorkgroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyWorkgroupAttribute"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyWorkgroupAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a migration job.
//
// Description:
//
// ## Usage notes
//
// This operation can only be used to start the migration jobs that are in the Ready, Stopped, or InError state.
//
// @param request - StartReplicationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartReplicationJobResponse
func (client *Client) StartReplicationJobWithContext(ctx context.Context, request *StartReplicationJobRequest, runtime *dara.RuntimeOptions) (_result *StartReplicationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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
		Action:      dara.String("StartReplicationJob"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartReplicationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses a migration job.
//
// Description:
//
// ## Usage notes
//
// You can call this operation to pause only a migration job whose primary status is Running and business status is Syncing.
//
// @param request - StopReplicationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopReplicationJobResponse
func (client *Client) StopReplicationJobWithContext(ctx context.Context, request *StopReplicationJobRequest, runtime *dara.RuntimeOptions) (_result *StopReplicationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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
		Action:      dara.String("StopReplicationJob"),
		Version:     dara.String("2019-06-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopReplicationJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates tags and adds them to Server Migration Center (SMC) resources. The SMC resources include migration sources and jobs.
//
// Description:
//
// ## Usage notes
//
// Up to 20 tags can be added to each SMC resource.
//
// Before you add tags to an SMC resource, Alibaba Cloud checks the number of the tags that have been added to the resource. If the maximum number is reached, an error message is returned.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
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
		Version:     dara.String("2019-06-01"),
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
// Removes tags that are added to Server Migration Center (SMC) resources. The SMC resources include migration sources and jobs.
//
// Description:
//
// You can call this operation to remove tags that are added to one or more SMC resources and delete the tags if the tags are no longer used.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
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
		Version:     dara.String("2019-06-01"),
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
