// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Deprecated: OpenAPI AddClientToBlackList is deprecated
//
// Summary:
//
// Adds a client to the blacklist in Cloud Parallel File Storage (CPFS) to evict its write requests, which serves as an I/O fencing mechanism.
//
// Description:
//
// This operation supports only CPFS parallel file systems.
//
// @param request - AddClientToBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddClientToBlackListResponse
func (client *Client) AddClientToBlackListWithContext(ctx context.Context, request *AddClientToBlackListRequest, runtime *dara.RuntimeOptions) (_result *AddClientToBlackListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIP) {
		query["ClientIP"] = request.ClientIP
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddClientToBlackList"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddClientToBlackListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies an automatic snapshot policy to one or more file systems.
//
// Description:
//
// *Before you begin**
//
// Only Advanced Extreme NAS file systems support this operation. Before you call this operation, complete the following preparations:
//
// - An Advanced Extreme NAS file system is created. When you call the [CreateFileSystem](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-createfilesystem) operation to create a file system, set the following parameters:
//
//   - FileSystemType: extreme
//
//   - StorageType: advance
//
// - An automatic snapshot policy is created. For more information, see [CreateAutoSnapshotPolicy](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-createautosnapshotpolicy).
//
// - The AliyunNASManageENIRole authorization is completed in the console. Before creating an Extreme NAS file system, associate the AliyunNASManageENIRole service-linked role with File Storage NAS to allow File Storage NAS to manage network interface controllers (NICs). For more information, see [Manage the service-linked role for File Storage NAS](https://help.aliyun.com/document_detail/208530.html).
//
// **Operation description**
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Each file system can have only one automatic snapshot policy applied.
//
// - An automatic snapshot policy can be applied to multiple file systems.
//
// - If the target file system already has an automatic snapshot policy applied, calling the ApplyAutoSnapshotPolicy operation replaces the current automatic snapshot policy of the file system.
//
// <props="intl">
//
// - This feature is in free public preview.
//
// - Each file system can have only one automatic snapshot policy applied.
//
// - An automatic snapshot policy can be applied to multiple file systems.
//
// - If the target file system already has an automatic snapshot policy applied, calling the ApplyAutoSnapshotPolicy operation replaces the current automatic snapshot policy of the file system.
//
// @param request - ApplyAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyAutoSnapshotPolicyResponse
func (client *Client) ApplyAutoSnapshotPolicyWithContext(ctx context.Context, request *ApplyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ApplyAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyId) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !dara.IsNil(request.FileSystemIds) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures automatic updates for a specified data flow.
//
// Description:
//
// - This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// - Only CPFS 2.2.0 and later support data flows. You can view the version information on the file system details page in the console.
//
// - You can add automatic update configurations only for data flows in the `Running` state.
//
// - You can add up to five automatic update configurations for a data flow.
//
// - It generally takes 2 to 5 minutes to create an automatic update configuration. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/336901.html) to query the data flow status.
//
// - Automatic updates depend on EventBridge to collect object modification events from the source OSS storage. You must first [activate EventBridge](https://help.aliyun.com/document_detail/182246.html).
//
//	> The event buses and event rules that CPFS creates in EventBridge contain the description `Create for cpfs auto refresh`. Do not modify or delete these event buses or event rules. Otherwise, automatic updates cannot work properly.
//
// - Automatic updates apply to prefixes specified by the RefreshPath parameter. When CPFS configures automatic synchronization for a prefix in a data flow, an event bus is created on the user side, and an event rule is created for the prefix of the source OSS bucket. When an object modification occurs within the prefix of the source OSS bucket, an OSS event is generated in EventBridge and processed by the CPFS data flow.
//
// - After you configure automatic updates (AutoRefresh), when data changes in the source storage, the changed metadata is automatically synchronized to the CPFS file system. The changed data is loaded on demand when a user accesses the file, or you can start a data stream task to load the data.
//
// - The automatic update interval (AutoRefreshInterval) specifies the interval at which CPFS checks whether data updates exist within the prefix of the source OSS bucket. If data updates exist, an automatic synchronization task is started. When the frequency of object modification events in the source OSS bucket exceeds the processing capability of the CPFS data flow, automatic update tasks accumulate, metadata updates are delayed, and the data flow status changes to Misconfigured. To resolve this issue, upgrade the data flow specifications or reduce the modification frequency in OSS.
//
// @param request - ApplyDataFlowAutoRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyDataFlowAutoRefreshResponse
func (client *Client) ApplyDataFlowAutoRefreshWithContext(ctx context.Context, request *ApplyDataFlowAutoRefreshRequest, runtime *dara.RuntimeOptions) (_result *ApplyDataFlowAutoRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRefreshInterval) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !dara.IsNil(request.AutoRefreshPolicy) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !dara.IsNil(request.AutoRefreshs) {
		query["AutoRefreshs"] = request.AutoRefreshs
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyDataFlowAutoRefresh"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Virtual Storage Channel (VSC) device with a file system.
//
// Description:
//
// - Only CPFS for Lingjun supports this feature.
//
// - Batch operations are supported. In batch mode, only one VscId can be associated with multiple FileSystemId values. This means the ResourceIds.VscId values must be the same.
//
// @param request - AttachVscToFilesystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachVscToFilesystemsResponse
func (client *Client) AttachVscToFilesystemsWithContext(ctx context.Context, request *AttachVscToFilesystemsRequest, runtime *dara.RuntimeOptions) (_result *AttachVscToFilesystemsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.RoleChain) {
		query["RoleChain"] = request.RoleChain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachVscToFilesystems"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachVscToFilesystemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the automatic snapshot policy that is created for a file system.
//
// Description:
//
// <props="china">
//
// -  This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// -  Only Advanced Extreme NAS supports this feature.
//
// -  Before you begin: You have applied an automatic snapshot policy to the target Extreme NAS file system by calling [ApplyAutoSnapshotPolicy](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-applyautosnapshotpolicy).
//
// <props="intl">
//
// -  This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// -  Only Advanced Extreme NAS supports this feature.
//
// -  Before you begin: You have applied an automatic snapshot policy to the target Extreme NAS file system by calling [ApplyAutoSnapshotPolicy](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-applyautosnapshotpolicy).
//
// @param request - CancelAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAutoSnapshotPolicyResponse
func (client *Client) CancelAutoSnapshotPolicyWithContext(ctx context.Context, request *CancelAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CancelAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemIds) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the automatic update configuration for a specified data flow.
//
// Description:
//
// - This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// - Data flows are supported only by CPFS 2.2.0 and later. You can view the version information on the file system details page in the console.
//
// - You can cancel the automatic update configuration only for data flows in the `Running` or `Stopped` state.
//
// - Canceling the automatic update configuration typically takes 2 to 5 minutes. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html) to query the status of the cancel automatic update task.
//
// @param request - CancelDataFlowAutoRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDataFlowAutoRefreshResponse
func (client *Client) CancelDataFlowAutoRefreshWithContext(ctx context.Context, request *CancelDataFlowAutoRefreshRequest, runtime *dara.RuntimeOptions) (_result *CancelDataFlowAutoRefreshResponse, _err error) {
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

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RefreshPath) {
		query["RefreshPath"] = request.RefreshPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDataFlowAutoRefresh"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a data flow streaming subtask.
//
// Description:
//
// *Prerequisite resource creation process**
//
// Before you cancel a data flow streaming subtask, create the following resources in order:
//
// 1. Create a data flow: Call [CreateDataFlow](https://www.alibabacloud.com/help/en/cpfs/bmcpfs/developer-reference/api-nas-2017-06-26-createdataflow-bmcpfs) to create a data flow.
//
// 2. Create a data flow task: Call [CreateDataFlowTask](https://www.alibabacloud.com/help/en/cpfs/bmcpfs/developer-reference/api-nas-2017-06-26-createdataflowtask-bmcpfs) to create a streaming import (StreamImport) or streaming export (StreamExport) task.
//
// 3. Create a data flow streaming subtask: Call [CreateDataFlowSubTask](https://www.alibabacloud.com/help/en/cpfs/bmcpfs/developer-reference/api-nas-2017-06-26-createdataflowsubtask-bmcpfs) to create a streaming subtask.
//
// > CPFS for Lingjun file systems must be created in the console.
//
// **Operation description**
//
// - Only CPFS for Lingjun 2.6.0 and later are supported. You can view the version information on the file system details page in the console.
//
// - Prerequisites: A CPFS for Lingjun file system is created, the DataFlow is in the Running state, the DataFlowTask is in the Executing state, and the DataFlowSubTask is in the CREATED or RUNNING state.
//
// - You can cancel a data flow streaming subtask only when it is in the CREATED or RUNNING state.
//
// - Data flow streaming tasks are executed asynchronously. You can call [DescribeDataFlowSubTasks](https://www.alibabacloud.com/help/en/cpfs/bmcpfs/developer-reference/api-nas-2017-06-26-describedataflowtasks-bmcpfs) to query the execution status of a streaming task.
//
// @param request - CancelDataFlowSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDataFlowSubTaskResponse
func (client *Client) CancelDataFlowSubTaskWithContext(ctx context.Context, request *CancelDataFlowSubTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelDataFlowSubTaskResponse, _err error) {
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

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DataFlowSubTaskId) {
		query["DataFlowSubTaskId"] = request.DataFlowSubTaskId
	}

	if !dara.IsNil(request.DataFlowTaskId) {
		query["DataFlowTaskId"] = request.DataFlowTaskId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDataFlowSubTask"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDataFlowSubTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a data flow batch task or streaming task that is in the Pending or Executing state.
//
// Description:
//
// *Prerequisites**
//
// Before you call this operation, create resources in the following order and make sure that the resources at each layer are in the required state:
//
// 1. Create a file system: Call [CreateFileSystem](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createfilesystem-cpfs) to create a CPFS file system.
//
//   - State requirement: The file system is in the Running state.
//
// 2. Create a mount target: Call [CreateMountTarget](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createmounttarget-cpfs) to create a mount target.
//
//   - State requirement: The mount target is in the Active state.
//
// 3. Create a fileset: Call [CreateFileset](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createfileset-cpfs) to create a fileset.
//
//   - State requirement: The fileset is created.
//
// 4. Create a data flow: Call [CreateDataFlow](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createdataflow-cpfs) to create a data flow.
//
//   - State requirement: The data flow is in the Running state.
//
// 5. Create a data flow task: Call [CreateDataFlowTask](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createdataflowtask-cpfs) to create a data flow task.
//
// **Operation description**
//
// - Only CPFS 2.2.0 and later and CPFS for Lingjun 2.4.0 and later support data flow tasks. You can view the version information on the file system details page in the console.
//
// - Only data flow tasks in the `Pending` state can be canceled.
//
// - Canceling a data flow task typically takes 5 to 10 minutes. You can call [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2402275.html) to query the task execution status.
//
// - If a streaming data flow task contains running streaming subtasks, the streaming task cannot be canceled. Otherwise, the InvalidStatus.ResourceMismatch error is returned.
//
// @param request - CancelDataFlowTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDataFlowTaskResponse
func (client *Client) CancelDataFlowTaskWithContext(ctx context.Context, request *CancelDataFlowTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelDataFlowTaskResponse, _err error) {
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

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDataFlowTask"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDataFlowTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the directory quota of a file system.
//
// Description:
//
// Only General-purpose NAS NFS file systems support the directory quota feature.
//
// @param request - CancelDirQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDirQuotaResponse
func (client *Client) CancelDirQuotaWithContext(ctx context.Context, request *CancelDirQuotaRequest, runtime *dara.RuntimeOptions) (_result *CancelDirQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDirQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDirQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the quota configured for a fileset.
//
// Description:
//
// Only CPFS for Lingjun 2.7.0 and later support quota cancellation.
//
// **Before you begin**
//
// - This operation applies only to CPFS for Lingjun file systems of version 2.7.0 or later.
//
// - The file system ID must start with `bmcpfs-`.
//
// - Before you call this operation, make sure that a quota has been configured for the target fileset by calling the [SetFilesetQuota](https://www.alibabacloud.com/help/en/cpfs/bmcpfs/developer-reference/api-nas-2017-06-26-setfilesetquota-bmcpfs) operation.
//
// @param request - CancelFilesetQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelFilesetQuotaResponse
func (client *Client) CancelFilesetQuotaWithContext(ctx context.Context, request *CancelFilesetQuotaRequest, runtime *dara.RuntimeOptions) (_result *CancelFilesetQuotaResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelFilesetQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelFilesetQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a data retrieval task that is in the active state.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// @param request - CancelLifecycleRetrieveJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelLifecycleRetrieveJobResponse
func (client *Client) CancelLifecycleRetrieveJobWithContext(ctx context.Context, request *CancelLifecycleRetrieveJobRequest, runtime *dara.RuntimeOptions) (_result *CancelLifecycleRetrieveJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelLifecycleRetrieveJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a running task in the recycle bin.
//
// Description:
//
// ### Operation description
//
// - Only General-purpose NAS file systems support this feature.
//
// - You can cancel only tasks that are in the Running state. The cancel operation does not take effect on tasks that have reached a final state (PartialSuccess, Success, Fail, or Cancelled).
//
// - When you cancel a running permanent deletion task, files that have already been permanently deleted cannot be recovered.
//
// - When you cancel a running file restoration task, restored files can be queried in the file system, and unrestored files can be queried in the recycle bin.
//
// @param request - CancelRecycleBinJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRecycleBinJobResponse
func (client *Client) CancelRecycleBinJobWithContext(ctx context.Context, request *CancelRecycleBinJobRequest, runtime *dara.RuntimeOptions) (_result *CancelRecycleBinJobResponse, _err error) {
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
		Action:      dara.String("CancelRecycleBinJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelRecycleBinJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which a file system instance belongs.
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
		Version:     dara.String("2017-06-26"),
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
// Creates a permission group.
//
// Description:
//
// - A maximum of 20 permission groups can be created per Alibaba Cloud account in a single region.
//
// - A maximum of 300 rules can be added to a permission group.
//
// - Only permission groups of the VPC network type can be created.
//
// @param request - CreateAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessGroupResponse
func (client *Client) CreateAccessGroupWithContext(ctx context.Context, request *CreateAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.AccessGroupType) {
		query["AccessGroupType"] = request.AccessGroupType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessGroup"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access point.
//
// Description:
//
// - When you invoke the CreateAccessPoint operation to create an access point, some resources are generated asynchronously. After the CreateAccessPoint operation succeeds, execute the [DescribeAccessPoints](https://help.aliyun.com/document_detail/2712239.html) or [DescribeAccessPoint](https://help.aliyun.com/document_detail/2712240.html) operation to query the access point status. Mount the file system only after the access point status becomes **Active**. Otherwise, the mount may fail.
//
// - Only General-purpose NAS NFS file systems support this feature.
//
// - If you enable the RAM policy (EnabledRam), configure the corresponding RAM permissions. For more information, see [Manage access points](https://help.aliyun.com/document_detail/2545998.html).
//
// @param request - CreateAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessPointResponse
func (client *Client) CreateAccessPointWithContext(ctx context.Context, request *CreateAccessPointRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroup) {
		query["AccessGroup"] = request.AccessGroup
	}

	if !dara.IsNil(request.AccessPointName) {
		query["AccessPointName"] = request.AccessPointName
	}

	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
	}

	if !dara.IsNil(request.EnabledRam) {
		query["EnabledRam"] = request.EnabledRam
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.OwnerGroupId) {
		query["OwnerGroupId"] = request.OwnerGroupId
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.Permission) {
		query["Permission"] = request.Permission
	}

	if !dara.IsNil(request.PosixGroupId) {
		query["PosixGroupId"] = request.PosixGroupId
	}

	if !dara.IsNil(request.PosixSecondaryGroupIds) {
		query["PosixSecondaryGroupIds"] = request.PosixSecondaryGroupIds
	}

	if !dara.IsNil(request.PosixUserId) {
		query["PosixUserId"] = request.PosixUserId
	}

	if !dara.IsNil(request.RootDirectory) {
		query["RootDirectory"] = request.RootDirectory
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VswId) {
		query["VswId"] = request.VswId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a permission rule for a permission group.
//
// Description:
//
// You can add up to 300 rules to a permission group.
//
// @param request - CreateAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessRuleResponse
func (client *Client) CreateAccessRuleWithContext(ctx context.Context, request *CreateAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.Ipv6SourceCidrIp) {
		query["Ipv6SourceCidrIp"] = request.Ipv6SourceCidrIp
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RWAccessType) {
		query["RWAccessType"] = request.RWAccessType
	}

	if !dara.IsNil(request.SourceCidrIp) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !dara.IsNil(request.UserAccessType) {
		query["UserAccessType"] = request.UserAccessType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessRule"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Agentic space.
//
// Description:
//
// *Before you begin**
//
// Complete the following preparations before calling this operation:
//
// 1. **An Agentic file system is created**: When calling [CreateFileSystem](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createfilesystem-cpfs) to create a file system, set the following parameters:
//
//   - `StorageType`: Set to `Agentic`.
//
//   - `ProtocolType`: Set to `NFS`.
//
//   - **Supported regions**: The Agentic storage type is available only in specific regions and zones. Check the console for details.
//
// **Usage notes**
//
// - This operation applies only to Agentic file systems.
//
// - Complete call chain example:
//
//  1. Call CreateFileSystem to create an Agentic file system (StorageType=Agentic, ProtocolType=NFS).
//
//  2. After the file system is created (status is Running), call CreateAgenticSpace to create an Agentic space.
//
//  3. To modify the Agentic space configuration, call the [ModifyAgenticSpace](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-modifyagenticspace) operation.
//
// @param request - CreateAgenticSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgenticSpaceResponse
func (client *Client) CreateAgenticSpaceWithContext(ctx context.Context, request *CreateAgenticSpaceRequest, runtime *dara.RuntimeOptions) (_result *CreateAgenticSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Azone) {
		query["Azone"] = request.Azone
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemPath) {
		query["FileSystemPath"] = request.FileSystemPath
	}

	if !dara.IsNil(request.Quota) {
		query["Quota"] = request.Quota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgenticSpace"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgenticSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an automatic snapshot policy.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// - You can create up to 100 automatic snapshot policies per region for each Alibaba Cloud account.
//
// - If a file system contains a large amount of data and the time required to create an automatic snapshot exceeds the interval between two scheduled time points, the next time point is automatically skipped. For example, you set 09:00, 10:00, 11:00, and 12:00 as the automatic snapshot time points. Because the file system contains a large amount of data, the snapshot creation starts at 09:00 and is completed at 10:20, which takes 80 minutes. The system skips the 10:00 time point and continues to create an automatic snapshot at 11:00.
//
// - Each file system supports up to 128 automatic snapshots. After the snapshot quota is reached, the system automatically deletes the earliest automatic snapshots. Manual snapshots are not affected.
//
// - When you modify the retention period of an automatic snapshot policy, the modification takes effect only on new snapshots. Existing snapshots retain their original retention period.
//
// - If an automatic snapshot is being created for a file system, you must wait until the automatic snapshot is completed before you can manually create a snapshot.
//
// - Automatic snapshot policies cannot be executed on file systems that are not in the Normal state.
//
// - Automatic snapshots follow a unified naming format: `auto_yyyyMMdd_X`. In this format, `auto` indicates an automatic snapshot, which distinguishes it from manual snapshots. `yyyyMMdd` indicates the date when the snapshot is created, where `y` indicates the year, `M` indicates the month, and `d` indicates the day. `X` indicates the sequence number of the automatic snapshot created on that day. For example, `auto_20201018_1` indicates the first automatic snapshot created on October 18, 2020.
//
// - A created automatic snapshot policy can be applied to any file system by calling ApplyAutoSnapshotPolicy, and the policy content can be modified by calling ModifyAutoSnapshotPolicy.
//
// <props="intl">
//
// - This feature is in free public preview.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// - You can create up to 100 automatic snapshot policies per region for each Alibaba Cloud account.
//
// - If a file system contains a large amount of data and the time required to create an automatic snapshot exceeds the interval between two scheduled time points, the next time point is automatically skipped. For example, you set 09:00, 10:00, 11:00, and 12:00 as the automatic snapshot time points. Because the file system contains a large amount of data, the snapshot creation starts at 09:00 and is completed at 10:20, which takes 80 minutes. The system skips the 10:00 time point and continues to create an automatic snapshot at 11:00.
//
// - Each file system supports up to 128 automatic snapshots. After the snapshot quota is reached, the system automatically deletes the earliest automatic snapshots. Manual snapshots are not affected.
//
// - When you modify the retention period of an automatic snapshot policy, the modification takes effect only on new snapshots. Existing snapshots retain their original retention period.
//
// - If an automatic snapshot is being created for a file system, you must wait until the automatic snapshot is completed before you can manually create a snapshot.
//
// - Automatic snapshot policies cannot be executed on file systems that are not in the Normal state.
//
// - Automatic snapshots follow a unified naming format: `auto_yyyyMMdd_X`. In this format, `auto` indicates an automatic snapshot, which distinguishes it from manual snapshots. `yyyyMMdd` indicates the date when the snapshot is created, where `y` indicates the year, `M` indicates the month, and `d` indicates the day. `X` indicates the sequence number of the automatic snapshot created on that day. For example, `auto_20201018_1` indicates the first automatic snapshot created on October 18, 2020.
//
// - A created automatic snapshot policy can be applied to any file system by calling ApplyAutoSnapshotPolicy, and the policy content can be modified by calling ModifyAutoSnapshotPolicy.
//
// @param request - CreateAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoSnapshotPolicyResponse
func (client *Client) CreateAutoSnapshotPolicyWithContext(ctx context.Context, request *CreateAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyName) {
		query["AutoSnapshotPolicyName"] = request.AutoSnapshotPolicyName
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.RepeatWeekdays) {
		query["RepeatWeekdays"] = request.RepeatWeekdays
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !dara.IsNil(request.TimePoints) {
		query["TimePoints"] = request.TimePoints
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access point for a CPFS for Lingjun file system.
//
// Description:
//
// Creates an access point for a CPFS for Lingjun file system.
//
// @param request - CreateCpfsAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCpfsAccessPointResponse
func (client *Client) CreateCpfsAccessPointWithContext(ctx context.Context, request *CreateCpfsAccessPointRequest, runtime *dara.RuntimeOptions) (_result *CreateCpfsAccessPointResponse, _err error) {
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

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RootDirectory) {
		query["RootDirectory"] = request.RootDirectory
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCpfsAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCpfsAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data flow between a CPFS file system (General-purpose or CPFS for Lingjun) and source storage.
//
// Description:
//
// - This operation applies to the following products:
//
// | Product | File system ID format | Minimum version that supports data flows |
//
// |------|----------------|------------------------|
//
// | **General-purpose CPFS*	- | Starts with `cpfs-`, such as cpfs-125487\\*\\*\\*\\	- | 2.2.0 or later |
//
// | **CPFS for Lingjun*	- | Starts with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\	- | 2.4.0 or later |
//
// > General-purpose CPFS and CPFS for Lingjun share the same set of API operations but differ in parameter values and feature support. Refer to the corresponding section based on the product type you use.
//
// - Basic operations
//
//   - You can create a data flow only when the General-purpose CPFS or CPFS for Lingjun file system is in the Running state.
//
//   - A maximum of **10*	- data flows can be created for a single General-purpose CPFS or CPFS for Lingjun file system.
//
//   - Creating a data flow typically takes 2 to 5 minutes. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/336901.html) to check whether the data flow is created.
//
// - Permissions.
//
//	When you create a data flow, Cloud Parallel File Storage obtains the `AliyunServiceRoleForNasOssDataflow` and `AliyunServiceRoleForNasEventNotification` service-linked roles. For more information, see [CPFS service-linked roles](https://help.aliyun.com/document_detail/185138.html).
//
// - Usage notes for General-purpose CPFS.
//
//	   This section applies to General-purpose CPFS whose file system IDs start with `cpfs-`.
//
//	- Billing
//
//	  - Creating a data flow incurs fees based on the data flow bandwidth. For more information, see [Billing of General-purpose CPFS](https://help.aliyun.com/document_detail/111858.html).
//
//	  - When you use auto-refresh, EventBridge collects object modification events from the source OSS storage, which incurs fees. For more information, see [EventBridge billing](https://help.aliyun.com/document_detail/163752.html).
//
//	- Data stream specifications
//
//	    - The data stream bandwidth (Throughput) supports three specifications: 600 MB/s, 1200 MB/s, and 1500 MB/s. The data stream bandwidth refers to the maximum transfer bandwidth that the data stream can achieve during data import or export.
//
//	  - Creating a data flow consumes one vSwitch IP address used by the Cloud Parallel File Storage mount target. Make sure that sufficient vSwitch IP resources are available.
//
//	  - Inventory check: Setting DryRun to true checks whether the resources required to create a data stream of the specified specification are available.
//
//	- Fileset
//
//	  - The destination of a data flow is a Fileset in the General-purpose CPFS file system. A Fileset is a new directory tree structure in the General-purpose CPFS file system. It is a small file system within the parent file system and has an independent inode space and management capability.
//
//	  - The Fileset must already exist when you create a data flow and cannot be nested with other Filesets. Only one data flow can be created on a Fileset, corresponding to one source storage.
//
//	  - The maximum number of files in a Fileset is 1 million. If the number of files imported from an OSS bucket exceeds this limit, creating new files returns a `no space` error.
//
//	 > If data already exists in the Fileset, the existing data in the Fileset is cleared and replaced with data synchronized from OSS after the data flow is created.
//
//	- Automatic synchronization
//
//	  - After automatic synchronization (AutoRefresh) is configured, when data in the source storage changes, the changed metadata is automatically synchronized to the Cloud Parallel File Storage file system. The changed data is loaded on demand when a user accesses the file, or loaded by starting a data stream node.
//
//	  - Automatic synchronization relies on EventBridge to collect object modification events from the source OSS storage. You must first [activate EventBridge](https://help.aliyun.com/document_detail/182246.html).
//
//	  - The scope of automatic synchronization is a prefix, specified by the RefreshPath parameter. A maximum of 5 automatic synchronization folders can be configured for a data stream.
//
//	  - The automatic synchronization interval (AutoRefreshInterval) specifies the synchronization interval that Cloud Parallel File Storage sets for checking whether data updates exist in the specified prefix of the source OSS bucket. If data updates exist, an automatic synchronization node is started. When the frequency of object modification events in the source OSS exceeds the processing capacity of the Cloud Parallel File Storage data stream, automatic synchronization nodes accumulate, metadata updates are delayed, and the data stream status changes to `Misconfigured`. You can increase the data stream specification or reduce the OSS modification frequency to resolve this issue.
//
//	  - When automatic synchronization is configured for a prefix in a Cloud Parallel File Storage data stream, an event bus is created on the user side, and an event rule is created for the prefix of the source OSS bucket. When an object modification occurs in the prefix of the source OSS bucket, an OSS event is generated in EventBridge and processed by the Cloud Parallel File Storage data stream.
//
//	   > The event bus and event rules created by Cloud Parallel File Storage in EventBridge have the description `Create for cpfs auto refresh`. Do not modify or delete the event bus or event rules. Otherwise, automatic synchronization does not work Normal.
//
//	- Source storage
//
//	   - Only OSS is supported as source storage. The data stream source storage (SourceStorage) must be an OSS bucket.
//
//	   - Cloud Parallel File Storage data streams support both encrypted and non-encrypted access to OSS. If you select encrypted (SSL) access to OSS, confirm that the encryption in transit of the OSS bucket supports encrypted access.
//
//	   - If data streams of multiple Cloud Parallel File Storage file systems, or multiple data streams of the same Cloud Parallel File Storage file system, use the same OSS bucket as source storage, enable versioning for the OSS bucket to prevent data conflicts when multiple Cloud Parallel File Storage file systems export data to the same source.
//
//	   - Cross-region OSS data streams are not supported. The OSS bucket must be in the same region as the Cloud Parallel File Storage file system.
//
//	     > Before creating a data flow, set a tag (key: cpfs-dataflow, value: true) on the source OSS bucket so that the Cloud Parallel File Storage data stream can access the data in the bucket. Do not delete or modify this tag while the data stream is in use. Otherwise, the Cloud Parallel File Storage data stream cannot access the data in the bucket.
//
// - Usage notes for CPFS for Lingjun.
//
//	   This section applies to CPFS for Lingjun whose file system IDs start with `bmcpfs-`.
//
//	- Source storage
//
//	    - Only OSS is supported as source storage. The source storage (SourceStorage) of a data flow must be an OSS bucket.
//
//	    - CPFS for Lingjun data flows support both encrypted and non-encrypted access to OSS. If you select encrypted (SSL) access to OSS, make sure that the encryption in transit of the OSS bucket supports encrypted access.
//
//	    - If data flows of multiple CPFS for Lingjun file systems, or multiple data flows of the same CPFS for Lingjun file system, use the same OSS bucket as source storage, enable versioning for the OSS bucket to prevent data conflicts when multiple CPFS for Lingjun file systems export data to the same source.
//
//	    - Cross-region OSS data flows are not supported. The OSS bucket must be in the same region as the CPFS for Lingjun file system.
//
//	   - CPFS for Lingjun 2.6.0 or later supports creating data flows with cross-account OSS.
//
//	   - The account ID parameter is required only when you use cross-account OSS.
//
//	   - When you use cross-account OSS, authorize the accounts first. For more information, see [Cross-account data flow authorization](https://help.aliyun.com/document_detail/2713462.html).
//
//	     > Before creating a data flow, set a tag (key: cpfs-dataflow, value: true) on the source OSS bucket so that the CPFS for Lingjun data flow can access the data in the bucket. Do not delete or modify this tag while the data flow is in use. Otherwise, the CPFS for Lingjun data flow cannot access the data in the bucket.
//
//	- Data flow restrictions on file systems.
//
//	  - In the file system path associated with a data flow, you cannot rename a non-empty directory. Otherwise, a Permission Denied or directory not empty error is returned.
//
//	   - Use special characters in directory and file names with caution. The following characters are supported: uppercase and lowercase letters, digits, exclamation marks (!), hyphens (-), underscores (_), periods (.), asterisks (*), and parentheses (()).
//
//	  - Excessively long paths are not supported. The maximum path length supported by data flows is 1023 characters.
//
//	- Data flow import restrictions.
//
//	  - After a symlink file is imported to CPFS for Lingjun, it is converted to a regular file that contains data, and the symlink information is lost.
//
//	  - If the OSS bucket has multiple versions, only the latest version is copied.
//
//	  - File names or subdirectory names longer than 255 bytes are not supported.
//
//	- Data flow export restrictions
//
//	  - After a symlink file is synchronized to OSS, the file that the symlink points to is not synchronized. Instead, the symlink becomes a regular empty object with no data.
//
//	  - A hardlink file is synchronized to OSS as a regular file.
//
//	  - Socket, Device, and Pipe files are converted to regular empty objects with no data when exported to an OSS bucket.
//
//	  - Directory paths longer than 1023 characters are not supported.
//
// @param request - CreateDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataFlowResponse
func (client *Client) CreateDataFlowWithContext(ctx context.Context, request *CreateDataFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateDataFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRefreshInterval) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !dara.IsNil(request.AutoRefreshPolicy) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !dara.IsNil(request.AutoRefreshs) {
		query["AutoRefreshs"] = request.AutoRefreshs
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemPath) {
		query["FileSystemPath"] = request.FileSystemPath
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	if !dara.IsNil(request.SourceSecurityType) {
		query["SourceSecurityType"] = request.SourceSecurityType
	}

	if !dara.IsNil(request.SourceStorage) {
		query["SourceStorage"] = request.SourceStorage
	}

	if !dara.IsNil(request.SourceStoragePath) {
		query["SourceStoragePath"] = request.SourceStoragePath
	}

	if !dara.IsNil(request.Throughput) {
		query["Throughput"] = request.Throughput
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a subtask for a data flow streaming task.
//
// Description:
//
// *Before you begin**
//
// Before you call this operation, create resources in the following order and make sure that the resources at each level are in the required state:
//
// 1. Create a file system: Create a CPFS for Lingjun file system in the console.
//
//   - State requirement: The file system is in the Running state.
//
// 2. Create a data flow: Call [CreateDataFlow](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createdataflow-cpfs) to create a data flow.
//
//   - State requirement: The data flow is in the Running state.
//
// 3. Create a data flow task: Call [CreateDataFlowTask](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createdataflowtask-cpfs) to create a streaming import (StreamImport) or streaming export (StreamExport) task.
//
//   - State requirement: The data flow task is in the Executing state.
//
// **Operation description**
//
// - Only CPFS for Lingjun 2.6.0 and later support this operation. You can view the version information on the file system details page in the console.
//
// - You can create subtasks only for data flow streaming tasks that are in the Executing state.
//
// - Data flow streaming tasks are executed asynchronously. You can call [DescribeDataFlowSubTasks](https://www.alibabacloud.com/help/en/cpfs/bmcpfs/developer-reference/api-nas-2017-06-26-describedataflowsubtasks-bmcpfs) to query the execution status of a streaming task.
//
// - When the data flow task type is a streaming task, the Executing state only indicates that you can create streaming import or streaming export tasks. It does not indicate that the import or export task is running.
//
// @param request - CreateDataFlowSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataFlowSubTaskResponse
func (client *Client) CreateDataFlowSubTaskWithContext(ctx context.Context, request *CreateDataFlowSubTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDataFlowSubTaskResponse, _err error) {
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

	if !dara.IsNil(request.Condition) {
		query["Condition"] = request.Condition
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DataFlowTaskId) {
		query["DataFlowTaskId"] = request.DataFlowTaskId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.DstFilePath) {
		query["DstFilePath"] = request.DstFilePath
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.SrcFilePath) {
		query["SrcFilePath"] = request.SrcFilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataFlowSubTask"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataFlowSubTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data flow task.
//
// Description:
//
// - Notes on CPFS
//
//   - Only CPFS 2.2.0 and later support data flow. You can view the version information on the file system details page in the console.
//
//   - Data flow tasks are executed asynchronously. You can call [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2402275.html) to query the task execution status. The time required for a task depends on the amount of data to be imported or exported. For large amounts of data, split the data into multiple tasks.
//
//   - You can create a data flow task only on a data flow in the Running state.
//
//   - When you manually run a data flow task, the automatic data update task of the data flow is interrupted and waits.
//
//   - When you create an export task, make sure that the total length of the absolute path of the file to be exported from CPFS does not exceed 1,023 characters.
//
// - Notes on CPFS for Lingjun
//
//   - Only CPFS for Lingjun 2.4.0 and later support data flow. You can view the version information on the file system details page in the console.
//
//   - Data flow tasks are executed asynchronously. You can call [DescribeDataFlowTasks](https://help.aliyun.com/document_detail/2402275.html) to query the task execution status. The time required for a task depends on the amount of data to be imported or exported. For large amounts of data, split the data into multiple tasks.
//
//   - You can create a data flow task only on a data flow in the Running state.
//
//   - When you create an export task, make sure that the total length of the absolute path of the file to be exported from CPFS for Lingjun does not exceed 1,023 characters.
//
//   - CPFS for Lingjun supports two task types: batch tasks and streaming tasks. For more information, see [Task type overview](https://help.aliyun.com/document_detail/2845429.html).
//
// @param request - CreateDataFlowTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataFlowTaskResponse
func (client *Client) CreateDataFlowTaskWithContext(ctx context.Context, request *CreateDataFlowTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDataFlowTaskResponse, _err error) {
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

	if !dara.IsNil(request.ConflictPolicy) {
		query["ConflictPolicy"] = request.ConflictPolicy
	}

	if !dara.IsNil(request.CreateDirIfNotExist) {
		query["CreateDirIfNotExist"] = request.CreateDirIfNotExist
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.Directory) {
		query["Directory"] = request.Directory
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.DstDirectory) {
		query["DstDirectory"] = request.DstDirectory
	}

	if !dara.IsNil(request.EntryList) {
		query["EntryList"] = request.EntryList
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Includes) {
		query["Includes"] = request.Includes
	}

	if !dara.IsNil(request.SrcTaskId) {
		query["SrcTaskId"] = request.SrcTaskId
	}

	if !dara.IsNil(request.TaskAction) {
		query["TaskAction"] = request.TaskAction
	}

	if !dara.IsNil(request.TransferFileListPath) {
		query["TransferFileListPath"] = request.TransferFileListPath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataFlowTask"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataFlowTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a directory in a file system.
//
// Description:
//
// Only General-purpose NAS NFS protocol file systems support this feature.
//
// @param request - CreateDirRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDirResponse
func (client *Client) CreateDirWithContext(ctx context.Context, request *CreateDirRequest, runtime *dara.RuntimeOptions) (_result *CreateDirResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.OwnerGroupId) {
		query["OwnerGroupId"] = request.OwnerGroupId
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.Permission) {
		query["Permission"] = request.Permission
	}

	if !dara.IsNil(request.Recursion) {
		query["Recursion"] = request.Recursion
	}

	if !dara.IsNil(request.RootDirectory) {
		query["RootDirectory"] = request.RootDirectory
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDir"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDirResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a directory or file.
//
// Description:
//
// - This operation is available only to specific users.
//
// - Only General-purpose NAS file systems that use the SMB Protocol Type and have access control enabled are supported. Alibaba Cloud Key Management Service (KMS) must also be activated.
//
// ## Prerequisites
//
// Before calling this API to create a file, complete the following four steps:
//
// 1. **Create an SMB General-purpose NAS file system**: Call the [CreateFileSystem](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-createfilesystem) operation with ProtocolType set to SMB to create a General-purpose NAS file system that supports the SMB protocol.
//
// 2. **Activate KMS**: Activate KMS by using one of the following methods:
//
//   - Call the Kms.OpenKmsService API operation.
//
//   - Go to the [console activation page](https://common-buy.aliyun.com/?commodityCode=kms#/open).
//
// 3. **Enable SMB ACL**: Call the [EnableSmbAcl](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-enablesmbacl) operation to enable the SMB ACL feature. Recommended configuration:
//
//   - AuthMethod=NTLM
//
//   - AuthCenter=EDS
//
//   - This configuration does not require an AD domain Keytab and is easier to use.
//
// 4. **Call CreateFile**: After completing the preceding steps, call this operation to create a directory or file.
//
// @param request - CreateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileResponse
func (client *Client) CreateFileWithContext(ctx context.Context, request *CreateFileRequest, runtime *dara.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.OwnerAccessInheritable) {
		query["OwnerAccessInheritable"] = request.OwnerAccessInheritable
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFile"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file system.
//
// Description:
//
// <props="china">
//
// - Before you call this operation, make sure that you fully understand the billing and pricing of File Storage NAS. For more information, see [Billing](https://help.aliyun.com/document_detail/178365.html) and [Pricing](https://www.aliyun.com/price/product?#/nas/).
//
// - To create a file system instance, complete real-name verification. For more information, see [Real-name verification](https://help.aliyun.com/document_detail/48263.html).
//
// - When you invoke this operation, the NAS service-linked role required for the operation is automatically created. For more information, see [Manage the service-linked role for File Storage NAS](https://help.aliyun.com/document_detail/208530.html).
//
// <props="intl">
//
// - Before you call this operation, make sure that you fully understand the billing and pricing of File Storage NAS. For more information, see [Billing](https://help.aliyun.com/document_detail/178365.html) and [Pricing](https://www.alibabacloud.com/zh/product/nas/pricing?_p_lc=1&spm=a2796.7960336.8215766810.64.414bb91aiZQiRG&spm=a2c4g.11186623.help-menu-27516.d_5_1_4_2_0.572a1880hEiKCk).
//
// - To create a file system instance, complete real-name verification. For more information, see [Real-name verification](https://help.aliyun.com/document_detail/48263.html).
//
// - When you invoke this operation, the NAS service-linked role required for the operation is automatically created. For more information, see [Manage the service-linked role for File Storage NAS](https://help.aliyun.com/document_detail/208530.html).
//
// @param request - CreateFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileSystemResponse
func (client *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateFileSystemResponse, _err error) {
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

	if !dara.IsNil(request.Capacity) {
		query["Capacity"] = request.Capacity
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EncryptType) {
		query["EncryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.KmsKeyId) {
		query["KmsKeyId"] = request.KmsKeyId
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.RedundancyType) {
		query["RedundancyType"] = request.RedundancyType
	}

	if !dara.IsNil(request.RedundancyVSwitchIds) {
		query["RedundancyVSwitchIds"] = request.RedundancyVSwitchIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
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
		Action:      dara.String("CreateFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a fileset.
//
// Description:
//
// - Usage notes for CPFS.
//
//   - Only CPFS 2.2.0 and later support fileset creation. You can view the version information on the file system details page in the console.
//
//   - A maximum of 10 filesets can be created for a single CPFS file system.
//
//   - The parent directory must be an existing directory.
//
//   - The maximum depth of a fileset path is 8 levels, with the root directory / as level 0. For example, a fileset path of /test/aaa/ccc/ has a depth of 3 levels.
//
//   - Nesting filesets within a fileset is not supported. If a parent directory is already specified as a fileset, its subdirectories cannot be specified as filesets.
//
//   - A fileset supports a maximum of 1,000,000 files. If the number of files exceeds this limit, adding new files returns a `no space` error message.
//
// - Usage notes for CPFS for Lingjun
//
//   - Only CPFS for Lingjun 2.7.0 and later support fileset creation. You can view the version information on the file system details page in the console.
//
//   - A maximum of 500 filesets can be created for a single CPFS file system.
//
//   - The fileset path must be a new path and cannot be an existing path. Renaming a fileset path is not supported, and the path cannot be a symbolic link.
//
//   - The maximum depth of a fileset path is 8 levels, with the root directory / as level 0. For example, a fileset path of /test/aaa/ccc/ has a depth of 3 levels.
//
//   - If the fileset path contains multiple directory levels, the parent directory must be an existing directory.
//
//   - Nesting filesets within a fileset is not supported. If a parent directory is already specified as a fileset, its subdirectories cannot be specified as filesets. Only one quota is supported for each fileset path.
//
//   - The minimum capacity quota for a fileset is 10 GiB, and the capacity can be increased in increments of 1 GiB.
//
//   - A fileset supports a maximum of 10,000,000,000 files or directories, with a minimum of 10,000 and an increment of 1.
//
//   - When you modify a directory quota, the quota capacity or file count must be greater than the used capacity or file count.
//
//   - Quota statistics have a 15-minute delay. The current actual usage takes effect after 15 minutes.
//
// @param request - CreateFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFilesetResponse
func (client *Client) CreateFilesetWithContext(ctx context.Context, request *CreateFilesetRequest, runtime *dara.RuntimeOptions) (_result *CreateFilesetResponse, _err error) {
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

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemPath) {
		query["FileSystemPath"] = request.FileSystemPath
	}

	if !dara.IsNil(request.Quota) {
		query["Quota"] = request.Quota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFileset"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFilesetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateLDAPConfig is deprecated
//
// Summary:
//
// Adds an LDAP configuration.
//
// @param request - CreateLDAPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLDAPConfigResponse
func (client *Client) CreateLDAPConfigWithContext(ctx context.Context, request *CreateLDAPConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateLDAPConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindDN) {
		query["BindDN"] = request.BindDN
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.SearchBase) {
		query["SearchBase"] = request.SearchBase
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLDAPConfig"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLDAPConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lifecycle management policy.
//
// Description:
//
// - Only General-purpose NAS file systems and CPFS for Lingjun support creating lifecycle management policies.
//
// - Each CPFS for Lingjun file system supports up to 10 Auto-type and 100 OnDemand-type lifecycle management policies.
//
// - Up to 20 lifecycle management policies for General-purpose NAS can be created in each region.
//
// @param request - CreateLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLifecyclePolicyResponse
func (client *Client) CreateLifecyclePolicyWithContext(ctx context.Context, request *CreateLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateLifecyclePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteRules) {
		query["DeleteRules"] = request.DeleteRules
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyName) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	if !dara.IsNil(request.LifecyclePolicyType) {
		query["LifecyclePolicyType"] = request.LifecyclePolicyType
	}

	if !dara.IsNil(request.LifecycleRuleName) {
		query["LifecycleRuleName"] = request.LifecycleRuleName
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.Paths) {
		query["Paths"] = request.Paths
	}

	if !dara.IsNil(request.RetrieveRules) {
		query["RetrieveRules"] = request.RetrieveRules
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.TransitRules) {
		query["TransitRules"] = request.TransitRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLifecyclePolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLifecyclePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data retrieval task.
//
// Description:
//
// - Only General-purpose NAS file systems support this feature.
//
// - Each Alibaba Cloud account can have up to 20 running data retrieval tasks in the same region.
//
// @param request - CreateLifecycleRetrieveJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLifecycleRetrieveJobResponse
func (client *Client) CreateLifecycleRetrieveJobWithContext(ctx context.Context, request *CreateLifecycleRetrieveJobRequest, runtime *dara.RuntimeOptions) (_result *CreateLifecycleRetrieveJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Paths) {
		query["Paths"] = request.Paths
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLifecycleRetrieveJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delivers system logs of a General-purpose NAS file system to log analysis.
//
// @param request - CreateLogAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogAnalysisResponse
func (client *Client) CreateLogAnalysisWithContext(ctx context.Context, request *CreateLogAnalysisRequest, runtime *dara.RuntimeOptions) (_result *CreateLogAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogAnalysis"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mount target.
//
// Description:
//
// - When you call the CreateMountTarget operation to create a mount target, some resources are generated asynchronously. After the CreateMountTarget operation succeeds, call the DescribeMountTargets operation to query the mount target status. Mount the file system only after the mount target status changes to **Active**. Otherwise, the mount operation may fail.
//
// - Calling this operation triggers the automatic creation of the service-linked role required for the operation. For more information, see [Manage the service-linked role for NAS](https://help.aliyun.com/document_detail/208530.html).
//
// @param request - CreateMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMountTargetResponse
func (client *Client) CreateMountTargetWithContext(ctx context.Context, request *CreateMountTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnableIpv6) {
		query["EnableIpv6"] = request.EnableIpv6
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
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
		Action:      dara.String("CreateMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an export directory for a protocol service.
//
// Description:
//
// -  This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// -  Before you begin
//
//	The target CPFS file system must be in the Running state, and a protocol service must have been created.
//
// - Other information
//
//   - The VPC CIDR block of the protocol service export must not overlap with the VPC CIDR block of the file system.
//
//   - The VPC CIDR blocks of multiple exports on the same protocol service must not overlap with each other.
//
//   - You can create up to 10 export directories for a single protocol service.
//
//   - Creating a protocol service export directory consumes IP addresses from the specified vSwitch (up to 32 IP addresses). Make sure that the target vSwitch has sufficient IP address resources.
//
// @param request - CreateProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProtocolMountTargetResponse
func (client *Client) CreateProtocolMountTargetWithContext(ctx context.Context, request *CreateProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateProtocolMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
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
		Action:      dara.String("CreateProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProtocolMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a protocol service for a Cloud Parallel File Storage (CPFS) file system. The creation process takes approximately 5 to 10 minutes.
//
// Description:
//
// - This operation applies only to CPFS file systems.
//
// -  Only CPFS 2.3.0 and later support protocol services. You can call the [DescribeFileSystems](https://help.aliyun.com/document_detail/163314.html) operation to query the version of a file system.
//
// - Protocol service specifications.
//
//	Protocol services include two Protocol Types: General and Cache. Compared with the General type, the Cache type provides hot spot data caching. When the cache is hit, the bandwidth of the Cache type protocol service can exceed the bandwidth of the CPFS file system and reach the maximum bandwidth configured for the protocol service.
//
//
//
//	  -   General: Provides NFS protocol access and [folder-level mount targets](https://help.aliyun.com/document_detail/427175.html) for CPFS. You do not need to configure a POSIX client cluster management. This feature is free of charge.
//
//	  -  Cache: Provides server-side in-memory caching based on the LRU policy on top of the General type. When data is cached in memory, CPFS can provide higher internal network bandwidth. The Cache type protocol service is available in two specifications: Cache L1 and Cache L2, which differ in internal network bandwidth and memory cache size.
//
//	   >  The Cache type protocol service is a paid service and is in invitational preview. For information about the billing of the Cache type protocol service, see [Billable items](https://help.aliyun.com/document_detail/111858.html). If you have any feedback or questions, join the DingTalk user group (group ID: 31045006299) to communicate with CPFS engineers.
//
// - Protocol type.
//
//	Only NFSv3 is supported.
//
// - Before you begin.
//
//	The CPFS file system must be created and in the Running state.
//
// - Other information.
//
//   - Only one protocol service can be created for each CPFS file system.
//
//   - Creating a protocol service consumes IP addresses on the specified vSwitch (up to 32 IP addresses). Make sure that the vSwitch has sufficient IP address resources.
//
// @param request - CreateProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProtocolServiceResponse
func (client *Client) CreateProtocolServiceWithContext(ctx context.Context, request *CreateProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateProtocolServiceResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolSpec) {
		query["ProtocolSpec"] = request.ProtocolSpec
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.Throughput) {
		query["Throughput"] = request.Throughput
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
		Action:      dara.String("CreateProtocolService"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProtocolServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to permanently delete a file or directory from the recycle bin.
//
// Description:
//
// - Only General-purpose NAS file systems support this feature.
//
// - The recycle bin must be enabled.
//
// - Mount the file system by using NFS or SMB (you must use an ECS instance in the same VPC to mount the file system), and then delete the target file from the mount directory to move it to the recycle bin.
//
// - When you permanently delete a directory, the contents in the directory are recursively deleted.
//
// - Only one permanent deletion task can be run at a time for a single file system. While a file or directory is being permanently deleted, you cannot initiate a new restoration or cleanup task.
//
// @param request - CreateRecycleBinDeleteJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecycleBinDeleteJobResponse
func (client *Client) CreateRecycleBinDeleteJobWithContext(ctx context.Context, request *CreateRecycleBinDeleteJobRequest, runtime *dara.RuntimeOptions) (_result *CreateRecycleBinDeleteJobResponse, _err error) {
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
		Action:      dara.String("CreateRecycleBinDeleteJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecycleBinDeleteJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a file from the recycle bin.
//
// Description:
//
// - Only General-purpose NAS file systems support this feature.
//
// - A single file system can run only one file restoration or cleanup task at a time. While a file is being restored, you cannot initiate a new file restoration or cleanup task.
//
// - A single restoration task can restore only one file or directory. Restoring a specified directory recursively restores all files in the directory.
//
// - After a file is restored, data reorganization is performed. Read request performance slightly decreases during data reorganization.
//
// @param request - CreateRecycleBinRestoreJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecycleBinRestoreJobResponse
func (client *Client) CreateRecycleBinRestoreJobWithContext(ctx context.Context, request *CreateRecycleBinRestoreJobRequest, runtime *dara.RuntimeOptions) (_result *CreateRecycleBinRestoreJobResponse, _err error) {
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
		Action:      dara.String("CreateRecycleBinRestoreJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecycleBinRestoreJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a snapshot.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.40a34bbe4vbvKY) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support snapshot creation.
//
// - A maximum of 128 snapshots can be created for a single file system.
//
// - The instance on which the file system is mounted must be in the Normal state. Otherwise, snapshots cannot be created.
//
// - If a snapshot is being created, you cannot create another snapshot for the same file system.
//
// - If the file system reaches its expire and release time while a snapshot is being created, the file system is released and the snapshot in the Creating state is also deleted.
//
// - Creating a snapshot may slightly degrade the performance of the file system and temporarily slow down I/O performance. Avoid creating snapshots during peak hours.
//
// - A snapshot only captures a backup of data at a specific point in time. Incremental data generated by file system operations during snapshot creation is not subject to synchronization to the snapshot.
//
// - Manually created snapshots are retained until they are deleted 15 days after your account has an overdue payment and services are suspended. Periodically delete snapshots that are no longer needed to avoid ongoing snapshot size charges.
//
// <props="intl">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support snapshot creation.
//
// - A maximum of 128 snapshots can be created for a single file system.
//
// - The instance on which the file system is mounted must be in the Normal state. Otherwise, snapshots cannot be created.
//
// - If a snapshot is being created, you cannot create another snapshot for the same file system.
//
// - If the file system reaches its expire and release time while a snapshot is being created, the file system is released and the snapshot in the Creating state is also deleted.
//
// - Creating a snapshot may slightly degrade the performance of the file system and temporarily slow down I/O performance. Avoid creating snapshots during peak hours.
//
// - A snapshot only captures a backup of data at a specific point in time. Incremental data generated by file system operations during snapshot creation is not subject to synchronization to the snapshot.
//
// - Manually created snapshots are retained until they are deleted 15 days after your account has an overdue payment and services are suspended. Periodically delete snapshots that are no longer needed to avoid ongoing snapshot size charges.
//
// @param request - CreateSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithContext(ctx context.Context, request *CreateSnapshotRequest, runtime *dara.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
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

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !dara.IsNil(request.SnapshotName) {
		query["SnapshotName"] = request.SnapshotName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSnapshot"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an existing permission group.
//
// Description:
//
// The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
//
// @param request - DeleteAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessGroupResponse
func (client *Client) DeleteAccessGroupWithContext(ctx context.Context, request *DeleteAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessGroup"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access point.
//
// Description:
//
// - Only NFS protocol file systems of General-purpose NAS support this feature.
//
// - After you delete an access point, all I/O operations that are accessing the folder through the access point are immediately interrupted. Proceed with caution.
//
// @param request - DeleteAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessPointResponse
func (client *Client) DeleteAccessPointWithContext(ctx context.Context, request *DeleteAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a permission rule from a permission group.
//
// Description:
//
// Rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be deleted.
//
// @param request - DeleteAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessRuleResponse
func (client *Client) DeleteAccessRuleWithContext(ctx context.Context, request *DeleteAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.AccessRuleId) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessRule"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Agentic space.
//
// Description:
//
// Applicable to agentic file systems.
//
// @param request - DeleteAgenticSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgenticSpaceResponse
func (client *Client) DeleteAgenticSpaceWithContext(ctx context.Context, request *DeleteAgenticSpaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteAgenticSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgenticSpace"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgenticSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an automatic snapshot policy.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// - If the automatic snapshot policy has been applied to a file system, the file system no longer executes the policy after the policy is deleted.
//
// <props="intl">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// - If the automatic snapshot policy has been applied to a file system, the file system no longer executes the policy after the policy is deleted.
//
// @param request - DeleteAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoSnapshotPolicyResponse
func (client *Client) DeleteAutoSnapshotPolicyWithContext(ctx context.Context, request *DeleteAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyId) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a CPFS for Lingjun access point.
//
// Description:
//
// Deletes a CPFS for Lingjun access point.
//
// @param request - DeleteCpfsAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCpfsAccessPointResponse
func (client *Client) DeleteCpfsAccessPointWithContext(ctx context.Context, request *DeleteCpfsAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteCpfsAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCpfsAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCpfsAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a dataflow.
//
// Description:
//
// - Only CPFS 2.2.0 and later and CPFS for Lingjun 2.4.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
// - You can delete only dataflows in the `Running` or `Stopped` state.
//
// - After a dataflow is deleted, the related resources are released and cannot be recovered. To use a dataflow again, create a new one.
//
// @param request - DeleteDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataFlowResponse
func (client *Client) DeleteDataFlowWithContext(ctx context.Context, request *DeleteDataFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataFlowResponse, _err error) {
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

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file system.
//
// Description:
//
// - You can delete a file system only when the number of mount targets on the file system is 0.
//
// - You can delete a file system only when no lifecycle policy is created for the file system.
//
// - After a file system is deleted, the data cannot be recovered. Proceed with caution.
//
// @param request - DeleteFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileSystemResponse
func (client *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a fileset.
//
// Description:
//
// - Only CPFS 2.2.0 and later and CPFS for Lingjun 2.7.0 and later support fileset deletion. After a fileset is deleted, all data in the directory is permanently deleted and cannot be recovered. Proceed with caution.
//
// - If deletion protection is enabled, you must disable deletion protection before you can delete the fileset.
//
// - All filesets on the target file system must be in the CREATED desired state before you can perform the delete operation.
//
// - Deleting a CPFS general-purpose fileset immediately releases disk space. Deleting a CPFS for Lingjun fileset gradually releases disk space. Deleted data cannot be recovered. Proceed with caution.
//
// @param request - DeleteFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFilesetResponse
func (client *Client) DeleteFilesetWithContext(ctx context.Context, request *DeleteFilesetRequest, runtime *dara.RuntimeOptions) (_result *DeleteFilesetResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFileset"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFilesetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteLDAPConfig is deprecated
//
// Summary:
//
// Deletes LDAP settings.
//
// Description:
//
// # Description
//
// This operation supports only Cloud Parallel File Storage (CPFS) parallel file systems.
//
// @param request - DeleteLDAPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLDAPConfigResponse
func (client *Client) DeleteLDAPConfigWithContext(ctx context.Context, request *DeleteLDAPConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLDAPConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLDAPConfig"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLDAPConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a lifecycle management policy.
//
// Description:
//
// Only General-purpose NAS file systems and CPFS for Lingjun support this feature.
//
// @param request - DeleteLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLifecyclePolicyResponse
func (client *Client) DeleteLifecyclePolicyWithContext(ctx context.Context, request *DeleteLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteLifecyclePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
	}

	if !dara.IsNil(request.LifecyclePolicyName) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLifecyclePolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLifecyclePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops log analysis for a General-purpose NAS file system.
//
// @param request - DeleteLogAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogAnalysisResponse
func (client *Client) DeleteLogAnalysisWithContext(ctx context.Context, request *DeleteLogAnalysisRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogAnalysis"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an existing mount target.
//
// Description:
//
// After a mount target is deleted, it cannot be recovered. Proceed with caution.
//
// @param request - DeleteMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMountTargetResponse
func (client *Client) DeleteMountTargetWithContext(ctx context.Context, request *DeleteMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetDomain) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an export directory of a protocol service.
//
// Description:
//
// This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// @param request - DeleteProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProtocolMountTargetResponse
func (client *Client) DeleteProtocolMountTargetWithContext(ctx context.Context, request *DeleteProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteProtocolMountTargetResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ExportId) {
		query["ExportId"] = request.ExportId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProtocolMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a protocol service from a Cloud Parallel File Storage (CPFS) file system.
//
// Description:
//
// - This operation applies only to CPFS file systems.
//
// - When you delete a protocol service, all export directories in the protocol service are also deleted.
//
// @param request - DeleteProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProtocolServiceResponse
func (client *Client) DeleteProtocolServiceWithContext(ctx context.Context, request *DeleteProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteProtocolServiceResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProtocolService"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProtocolServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified snapshot or cancels an ongoing snapshot creation task.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// <props="intl">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// @param request - DeleteSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshotWithContext(ctx context.Context, request *DeleteSnapshotRequest, runtime *dara.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnapshot"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries permission group information.
//
// @param request - DescribeAccessGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessGroupsResponse
func (client *Client) DescribeAccessGroupsWithContext(ctx context.Context, request *DescribeAccessGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UseUTCDateTime) {
		query["UseUTCDateTime"] = request.UseUTCDateTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessGroups"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an access point.
//
// Description:
//
// Only General-purpose NAS NFS file systems support this feature.
//
// @param request - DescribeAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessPointResponse
func (client *Client) DescribeAccessPointWithContext(ctx context.Context, request *DescribeAccessPointRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeAccessPoints is deprecated, please use NAS::2017-06-26::ListAccessPoints instead.
//
// Summary:
//
// Queries access point information.
//
// Description:
//
// Only General-purpose NAS NFS file systems support this feature.
//
// @param request - DescribeAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessPointsResponse
func (client *Client) DescribeAccessPointsWithContext(ctx context.Context, request *DescribeAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroup) {
		query["AccessGroup"] = request.AccessGroup
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessPoints"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the description of a permission rule.
//
// @param request - DescribeAccessRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessRulesResponse
func (client *Client) DescribeAccessRulesWithContext(ctx context.Context, request *DescribeAccessRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.AccessRuleId) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
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
		Action:      dara.String("DescribeAccessRules"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Agentic spaces.
//
// Description:
//
// Applies to Agentic file systems.
//
// ## Before you begin
//
// Before you call this operation, create resources in the following order:
//
// 1. **Create an Agentic file system**: When you call [CreateFileSystem](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createfilesystem-cpfs) to create a file system, set the following parameters:
//
//   - `StorageType`: Set to `Agentic`.
//
//   - `ProtocolType`: Set to `NFS`.
//
//   - **Supported regions**: The Agentic storage type is available only in specific regions and zones. For more information, see the console.
//
// 2. **Query the AgenticSpace list**: After the file system is created and its status changes to Running, call the DescribeAgenticSpaces operation to query the AgenticSpace list under the file system.
//
//   - If Agentic spaces are created under the file system, the complete list is returned.
//
//   - If no Agentic spaces are created under the file system, an empty list is returned.
//
// @param request - DescribeAgenticSpacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAgenticSpacesResponse
func (client *Client) DescribeAgenticSpacesWithContext(ctx context.Context, request *DescribeAgenticSpacesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAgenticSpacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAgenticSpaces"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAgenticSpacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the automatic snapshot policies that have been created.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// <props="intl">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// @param request - DescribeAutoSnapshotPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoSnapshotPoliciesResponse
func (client *Client) DescribeAutoSnapshotPoliciesWithContext(ctx context.Context, request *DescribeAutoSnapshotPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoSnapshotPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyId) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
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
		Action:      dara.String("DescribeAutoSnapshotPolicies"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoSnapshotPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of automatic snapshot tasks.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// <props="intl">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// @param request - DescribeAutoSnapshotTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoSnapshotTasksResponse
func (client *Client) DescribeAutoSnapshotTasksWithContext(ctx context.Context, request *DescribeAutoSnapshotTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeAutoSnapshotTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyIds) {
		query["AutoSnapshotPolicyIds"] = request.AutoSnapshotPolicyIds
	}

	if !dara.IsNil(request.FileSystemIds) {
		query["FileSystemIds"] = request.FileSystemIds
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
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
		Action:      dara.String("DescribeAutoSnapshotTasks"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAutoSnapshotTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeBlackListClients is deprecated
//
// Summary:
//
// Retrieves the status of blacklisted clients in the Cloud Parallel File Storage (CPFS) service.
//
// Description:
//
// # Description
//
// This operation applies only to CPFS parallel file systems.
//
// @param request - DescribeBlackListClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBlackListClientsResponse
func (client *Client) DescribeBlackListClientsWithContext(ctx context.Context, request *DescribeBlackListClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBlackListClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIP) {
		query["ClientIP"] = request.ClientIP
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBlackListClients"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBlackListClientsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of clients mounted to an access point of a CPFS for Lingjun file system.
//
// Description:
//
// Queries the list of clients mounted to an access point of a CPFS for Lingjun file system.
//
// @param request - DescribeCpfsAccessPointMountedClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCpfsAccessPointMountedClientsResponse
func (client *Client) DescribeCpfsAccessPointMountedClientsWithContext(ctx context.Context, request *DescribeCpfsAccessPointMountedClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCpfsAccessPointMountedClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
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
		Action:      dara.String("DescribeCpfsAccessPointMountedClients"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCpfsAccessPointMountedClientsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access points of CPFS for Lingjun.
//
// Description:
//
// Queries access points of CPFS for Lingjun.
//
// @param request - DescribeCpfsAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCpfsAccessPointsResponse
func (client *Client) DescribeCpfsAccessPointsWithContext(ctx context.Context, request *DescribeCpfsAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCpfsAccessPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCpfsAccessPoints"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCpfsAccessPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the batch subtasks of a data flow.
//
// Description:
//
// Supported only by CPFS for Lingjun 2.6.0 and later. You can view the version information on the file system details page in the console.
//
// **Before you begin**
//
// Before you call this operation to query data flow subtasks, you must create resources and start tasks in the following order:
//
// 1. **Create a bmcpfs file system**: You must create a CPFS for Lingjun file system in the console and wait until its status changes to Running.
//
// 2. **Create a data flow rule**: Call the [CreateDataFlow](https://www.alibabacloud.com/help/en/cpfs/bmcpfs/developer-reference/api-nas-2017-06-26-createdataflow-bmcpfs) operation on the file system to establish a data flow channel.
//
// 3. **Start a data flow task**: Call the [CreateDataFlowTask](https://www.alibabacloud.com/help/en/cpfs/bmcpfs/developer-reference/api-nas-2017-06-26-createdataflowtask-bmcpfs) operation to create and start a data flow task. The task status must be Running.
//
// 4. **Query the subtask list**: Subtasks are automatically generated during the execution of a data flow task. Call the DescribeDataFlowSubTasks operation to query the list of these subtasks.
//
// @param request - DescribeDataFlowSubTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataFlowSubTasksResponse
func (client *Client) DescribeDataFlowSubTasksWithContext(ctx context.Context, request *DescribeDataFlowSubTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataFlowSubTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataFlowSubTasks"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataFlowSubTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of data flow tasks.
//
// Description:
//
// Only CPFS 2.2.0 and later and CPFS for Lingjun 2.4.0 and later support querying data flow tasks. You can view the version information on the file system details page in the console.
//
// @param request - DescribeDataFlowTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataFlowTasksResponse
func (client *Client) DescribeDataFlowTasksWithContext(ctx context.Context, request *DescribeDataFlowTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataFlowTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WithReports) {
		query["WithReports"] = request.WithReports
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataFlowTasks"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataFlowTasksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data flow information of a CPFS file system.
//
// Description:
//
// - Only CPFS 2.2.0 and later and CPFS for Lingjun 2.4.0 and later support data flows. You can view the version information on the file system details page in the console.
//
// - For filter keys (Filters), FsetIds, DataFlowIds, SourceStorage, ThroughputList, and Status require exact match. FileSystemPath, Description, and SourceStoragePath support fuzzy match.
//
// - Combined queries are supported.
//
// @param request - DescribeDataFlowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataFlowsResponse
func (client *Client) DescribeDataFlowsWithContext(ctx context.Context, request *DescribeDataFlowsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDataFlowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDataFlows"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDataFlowsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of each directory quota in a file system.
//
// Description:
//
// Only General-purpose NAS NFS file systems support the directory quota feature.
//
// @param request - DescribeDirQuotasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDirQuotasResponse
func (client *Client) DescribeDirQuotasWithContext(ctx context.Context, request *DescribeDirQuotasRequest, runtime *dara.RuntimeOptions) (_result *DescribeDirQuotasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDirQuotas"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDirQuotasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeFileSystemStatistics is deprecated, please use NAS::2017-06-26::DescribeResourceStatistics instead.
//
// Summary:
//
// Lists the statistics information of file systems under the current account.
//
// @param request - DescribeFileSystemStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileSystemStatisticsResponse
func (client *Client) DescribeFileSystemStatisticsWithContext(ctx context.Context, request *DescribeFileSystemStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileSystemStatisticsResponse, _err error) {
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
		Action:      dara.String("DescribeFileSystemStatistics"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFileSystemStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries file system information.
//
// @param request - DescribeFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileSystemsResponse
func (client *Client) DescribeFileSystemsWithContext(ctx context.Context, request *DescribeFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileSystemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
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

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
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
		Action:      dara.String("DescribeFileSystems"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFileSystemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of created filesets.
//
// Description:
//
// - Only CPFS 2.2.0 and later and CPFS for Lingjun 2.7.0 and later support filesets. You can view the version information on the file system details page in the console.
//
// - For filter keys (Filters), FsetIds requires an exact match, while FileSystemPath and Description support fuzzy match.
//
// - Combination queries are supported.
//
// @param request - DescribeFilesetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFilesetsResponse
func (client *Client) DescribeFilesetsWithContext(ctx context.Context, request *DescribeFilesetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFilesetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderByField) {
		query["OrderByField"] = request.OrderByField
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFilesets"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFilesetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of HpnZones associated with file systems. When compute nodes are located in an associated HpnZone, access to the file system achieves optimal performance.
//
// Description:
//
// - Only CPFS for Lingjun file systems are supported.
//
// - This is a batch operation that allows you to query up to 20 file systems at a time.
//
// @param tmpReq - DescribeFilesystemsAssociatedHpnZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFilesystemsAssociatedHpnZonesResponse
func (client *Client) DescribeFilesystemsAssociatedHpnZonesWithContext(ctx context.Context, tmpReq *DescribeFilesystemsAssociatedHpnZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeFilesystemsAssociatedHpnZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeFilesystemsAssociatedHpnZonesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filesystems) {
		request.FilesystemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filesystems, dara.String("Filesystems"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilesystemsShrink) {
		query["Filesystems"] = request.FilesystemsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFilesystemsAssociatedHpnZones"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFilesystemsAssociatedHpnZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual storage channel information associated with file systems.
//
// Description:
//
// - Only CPFS for Lingjun supports this feature.
//
// - Batch execution is supported. In batch execution mode, only one VscId can be associated with multiple FileSystemIds. This means the ResourceIds.VscId values must be the same.
//
// - To query the association results between file systems and virtual storage channels, call AttachVscToFilesystems to complete the association first. You can call ListVscs of Lingjun Intelligent Computing eflo-controller/2022-12-15 to query VscIds. If no virtual storage channel has been created, call CreateVsc to create one.
//
// @param request - DescribeFilesystemsVscAttachInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFilesystemsVscAttachInfoResponse
func (client *Client) DescribeFilesystemsVscAttachInfoWithContext(ctx context.Context, request *DescribeFilesystemsVscAttachInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeFilesystemsVscAttachInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.RoleChain) {
		query["RoleChain"] = request.RoleChain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFilesystemsVscAttachInfo"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFilesystemsVscAttachInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of lifecycle management policies.
//
// Description:
//
// Only General-purpose NAS file systems and CPFS for Lingjun support this feature.
//
// @param request - DescribeLifecyclePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLifecyclePoliciesResponse
func (client *Client) DescribeLifecyclePoliciesWithContext(ctx context.Context, request *DescribeLifecyclePoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeLifecyclePoliciesResponse, _err error) {
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
		Action:      dara.String("DescribeLifecyclePolicies"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLifecyclePoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution logs of a lifecycle policy. You can query the logs generated within the last 90 days or up to 1,000 log entries. Only CPFS for Lingjun is supported.
//
// Description:
//
// Queries the execution logs of a lifecycle policy. You can query the logs generated within the last 90 days or up to 1,000 log entries. Only CPFS for Lingjun is supported.
//
// ## Before you begin
//
// - Make sure that the file system specified by FileSystemId is a CPFS for Lingjun file system.
//
// - Before you call this operation, call [DescribeLifecyclePolicies](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-describelifecyclepolicies) to query and obtain a valid LifecyclePolicyId.
//
// @param request - DescribeLifecyclePolicyLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLifecyclePolicyLogsResponse
func (client *Client) DescribeLifecyclePolicyLogsWithContext(ctx context.Context, request *DescribeLifecyclePolicyLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeLifecyclePolicyLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
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
		Action:      dara.String("DescribeLifecyclePolicyLogs"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLifecyclePolicyLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the log analysis configurations in log analysis.
//
// @param request - DescribeLogAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogAnalysisResponse
func (client *Client) DescribeLogAnalysisWithContext(ctx context.Context, request *DescribeLogAnalysisRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
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
		Action:      dara.String("DescribeLogAnalysis"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries mount target information.
//
// @param request - DescribeMountTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMountTargetsResponse
func (client *Client) DescribeMountTargetsWithContext(ctx context.Context, request *DescribeMountTargetsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMountTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DualStackMountTargetDomain) {
		query["DualStackMountTargetDomain"] = request.DualStackMountTargetDomain
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetDomain) {
		query["MountTargetDomain"] = request.MountTargetDomain
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
		Action:      dara.String("DescribeMountTargets"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMountTargetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of mounted clients.
//
// Description:
//
// - Only General-purpose NAS file systems support this operation.
//
// - The client list displays client IP addresses that have read or write access to the file system within the last minute. Some mounted client IP addresses that have not accessed the file system may not appear in this list.
//
// @param request - DescribeMountedClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMountedClientsResponse
func (client *Client) DescribeMountedClientsWithContext(ctx context.Context, request *DescribeMountedClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMountedClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIP) {
		query["ClientIP"] = request.ClientIP
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetDomain) {
		query["MountTargetDomain"] = request.MountTargetDomain
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
		Action:      dara.String("DescribeMountedClients"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMountedClientsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the NFS ACL feature is enabled for a specified file system.
//
// Description:
//
// Only General-purpose NAS NFS file systems support this feature.
//
// @param request - DescribeNfsAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNfsAclResponse
func (client *Client) DescribeNfsAclWithContext(ctx context.Context, request *DescribeNfsAclRequest, runtime *dara.RuntimeOptions) (_result *DescribeNfsAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNfsAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNfsAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the export directories of a protocol service.
//
// Description:
//
// This operation is applicable only to Cloud Parallel File Storage (CPFS) file systems.
//
// @param request - DescribeProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtocolMountTargetResponse
func (client *Client) DescribeProtocolMountTargetWithContext(ctx context.Context, request *DescribeProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *DescribeProtocolMountTargetResponse, _err error) {
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

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProtocolServiceIds) {
		query["ProtocolServiceIds"] = request.ProtocolServiceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProtocolMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries protocol service information.
//
// Description:
//
// This operation is applicable only to Cloud Parallel File Storage (CPFS) file systems.
//
// @param request - DescribeProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProtocolServiceResponse
func (client *Client) DescribeProtocolServiceWithContext(ctx context.Context, request *DescribeProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeProtocolServiceResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProtocolServiceIds) {
		query["ProtocolServiceIds"] = request.ProtocolServiceIds
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProtocolService"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProtocolServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available Alibaba Cloud regions.
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
	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
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
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2017-06-26"),
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
// Queries the SMB AD ACL feature.
//
// Description:
//
// - Before you begin: Before calling this API, ensure that the KMS service is activated in the target region. You can activate it by calling the [Kms.OpenKmsService](https://www.alibabacloud.com/help/en/kms/key-management-service/developer-reference/api-kms-2016-01-20-openkmsservice) API or through the console. This restriction also applies to unencrypted file systems where EncryptType=0.
//
// @param request - DescribeSmbAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmbAclResponse
func (client *Client) DescribeSmbAclWithContext(ctx context.Context, request *DescribeSmbAclRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmbAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSmbAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSmbAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about one or more snapshots of a specified file system.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// -   Only Advanced Extreme NAS file systems support this feature.
//
// <props="intl">
//
// -  This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// @param request - DescribeSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnapshotsResponse
func (client *Client) DescribeSnapshotsWithContext(ctx context.Context, request *DescribeSnapshotsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SnapshotIds) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	if !dara.IsNil(request.SnapshotName) {
		query["SnapshotName"] = request.SnapshotName
	}

	if !dara.IsNil(request.SnapshotType) {
		query["SnapshotType"] = request.SnapshotType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnapshots"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of storage plans.
//
// @param request - DescribeStoragePackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStoragePackagesResponse
func (client *Client) DescribeStoragePackagesWithContext(ctx context.Context, request *DescribeStoragePackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeStoragePackagesResponse, _err error) {
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

	if !dara.IsNil(request.UseUTCDateTime) {
		query["UseUTCDateTime"] = request.UseUTCDateTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStoragePackages"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStoragePackagesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all zones in a region and the file system types supported by each zone.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithContext(ctx context.Context, request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates a Virtual Storage Channel (VSC) device from a file system.
//
// Description:
//
// - Only CPFS for Lingjun supports this feature.
//
// - Batch operations are supported. In batch mode, only one VscId can be associated with multiple FileSystemId values. This means the ResourceIds.VscId values must be the same.
//
// @param request - DetachVscFromFilesystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachVscFromFilesystemsResponse
func (client *Client) DetachVscFromFilesystemsWithContext(ctx context.Context, request *DetachVscFromFilesystemsRequest, runtime *dara.RuntimeOptions) (_result *DetachVscFromFilesystemsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.RoleChain) {
		query["RoleChain"] = request.RoleChain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachVscFromFilesystems"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachVscFromFilesystemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the recycle bin feature for a General-purpose NAS file system and deletes all data in the recycle bin.
//
// Description:
//
// - Only General-purpose NAS file systems support this feature.
//
// - Disabling the recycle bin permanently deletes all files in the recycle bin.
//
// - After a disabled recycle bin is re-enabled, the recycle bin is empty. Files deleted before the recycle bin was re-enabled cannot be recovered.
//
// @param request - DisableAndCleanRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAndCleanRecycleBinResponse
func (client *Client) DisableAndCleanRecycleBinWithContext(ctx context.Context, request *DisableAndCleanRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *DisableAndCleanRecycleBinResponse, _err error) {
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
		Action:      dara.String("DisableAndCleanRecycleBin"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableAndCleanRecycleBinResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables EBS Lens.
//
// Description:
//
// - This operation is currently in invitational preview. To use this feature, submit a ticket.
//
// @param request - DisableDataInsightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDataInsightResponse
func (client *Client) DisableDataInsightWithContext(ctx context.Context, request *DisableDataInsightRequest, runtime *dara.RuntimeOptions) (_result *DisableDataInsightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableDataInsight"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableDataInsightResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the NFS ACL feature.
//
// Description:
//
// Only General-purpose NAS NFS file systems support this feature.
//
// @param request - DisableNfsAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableNfsAclResponse
func (client *Client) DisableNfsAclWithContext(ctx context.Context, request *DisableNfsAclRequest, runtime *dara.RuntimeOptions) (_result *DisableNfsAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableNfsAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableNfsAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the SMB AD ACL feature.
//
// Description:
//
// - Before you begin: Before calling this API, make sure that Key Management Service (KMS) is activated in the target region. You can activate KMS by calling the [Kms.OpenKmsService](https://www.alibabacloud.com/help/en/kms/key-management-service/developer-reference/api-kms-2016-01-20-openkmsservice) API or by using the console. The file system protocol type must be SMB. You can call [DescribeFileSystems](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-describefilesystems) to verify the file system protocol type.
//
// @param request - DisableSmbAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSmbAclResponse
func (client *Client) DisableSmbAclWithContext(ctx context.Context, request *DisableSmbAclRequest, runtime *dara.RuntimeOptions) (_result *DisableSmbAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSmbAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSmbAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables EBS Lens.
//
// Description:
//
// - This feature is currently in invitational preview. To use this feature, submit a ticket.
//
// @param request - EnableDataInsightRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDataInsightResponse
func (client *Client) EnableDataInsightWithContext(ctx context.Context, request *EnableDataInsightRequest, runtime *dara.RuntimeOptions) (_result *EnableDataInsightResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableDataInsight"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableDataInsightResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the NFS ACL feature.
//
// Description:
//
// Only General-purpose NAS NFS file systems support this feature.
//
// @param request - EnableNfsAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableNfsAclResponse
func (client *Client) EnableNfsAclWithContext(ctx context.Context, request *EnableNfsAclRequest, runtime *dara.RuntimeOptions) (_result *EnableNfsAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableNfsAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableNfsAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the recycle bin feature.
//
// Description:
//
// Only General-purpose NAS file systems support this operation.
//
// @param request - EnableRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRecycleBinResponse
func (client *Client) EnableRecycleBinWithContext(ctx context.Context, request *EnableRecycleBinRequest, runtime *dara.RuntimeOptions) (_result *EnableRecycleBinResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ReservedDays) {
		query["ReservedDays"] = request.ReservedDays
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableRecycleBin"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableRecycleBinResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the SMB AD ACL feature.
//
// Description:
//
// - By connecting the mount target of an SMB file system to an AD domain, you can implement user identity authentication management and file-level access permissions control in the AD domain. Before you mount and use an SMB file system as an AD domain user, you must register a service for the SMB file system in the AD domain, generate a keytab file, and upload the file to the NAS console to enable the SMB ACL feature. [Generate a keytab file](https://help.aliyun.com/document_detail/154930.html).
//
// - Before you begin: Before invoking this operation, confirm that Key Management Service (KMS) is activated in the target area. You can invoke the [Kms.OpenKmsService](https://www.alibabacloud.com/help/en/kms/key-management-service/developer-reference/api-kms-2016-01-20-openkmsservice) operation or activate KMS in the console.
//
// @param request - EnableSmbAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSmbAclResponse
func (client *Client) EnableSmbAclWithContext(ctx context.Context, request *EnableSmbAclRequest, runtime *dara.RuntimeOptions) (_result *EnableSmbAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Keytab) {
		query["Keytab"] = request.Keytab
	}

	if !dara.IsNil(request.KeytabMd5) {
		query["KeytabMd5"] = request.KeytabMd5
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSmbAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSmbAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an Agentic space.
//
// Description:
//
// Applies to agentic-type file systems.
//
// @param request - GetAgenticSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgenticSpaceResponse
func (client *Client) GetAgenticSpaceWithContext(ctx context.Context, request *GetAgenticSpaceRequest, runtime *dara.RuntimeOptions) (_result *GetAgenticSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgenticSpace"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgenticSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a specified directory contains Infrequent Access (IA) storage class or Archive storage class files, or queries whether a specified file is an IA storage class or Archive storage class file.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// @param request - GetDirectoryOrFilePropertiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDirectoryOrFilePropertiesResponse
func (client *Client) GetDirectoryOrFilePropertiesWithContext(ctx context.Context, request *GetDirectoryOrFilePropertiesRequest, runtime *dara.RuntimeOptions) (_result *GetDirectoryOrFilePropertiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDirectoryOrFileProperties"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDirectoryOrFilePropertiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a created Fileset.
//
// Description:
//
// Only CPFS 2.2.0 and later and CPFS for Lingjun 2.7.0 and later support Filesets. You can view the version information on the file system details page in the console.
//
// @param request - GetFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFilesetResponse
func (client *Client) GetFilesetWithContext(ctx context.Context, request *GetFilesetRequest, runtime *dara.RuntimeOptions) (_result *GetFilesetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileset"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFilesetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the export directory information of a protocol service.
//
// Description:
//
// Queries the export directory information of a protocol service.
//
// **Before you begin**
//
// Before you call this operation, make sure that each dependent resource in the following three-layer resource model already exists:
//
// 1. **File system (FileSystem)**: Call [DescribeFileSystems](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-describefilesystems) to obtain the FileSystemId.
//
// 2. **Protocol service (ProtocolService)**: Call [DescribeProtocolService](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describeprotocolservice-cpfs) to obtain the ProtocolServiceId, or call [CreateProtocolService](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createprotocolservice-cpfs) to create one.
//
// 3. **Export directory (ProtocolMountTarget)**: Call [CreateProtocolMountTarget](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createprotocolmounttarget-cpfs) to create an export directory and obtain the ExportId, or call [DescribeProtocolMountTarget](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describeprotocolmounttarget-cpfs) to query the ExportId.
//
// @param request - GetProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProtocolMountTargetResponse
func (client *Client) GetProtocolMountTargetWithContext(ctx context.Context, request *GetProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *GetProtocolMountTargetResponse, _err error) {
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

	if !dara.IsNil(request.ExportId) {
		query["ExportId"] = request.ExportId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProtocolMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the recycle bin configuration of a specified General-purpose NAS file system.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// @param request - GetRecycleBinAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRecycleBinAttributeResponse
func (client *Client) GetRecycleBinAttributeWithContext(ctx context.Context, request *GetRecycleBinAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetRecycleBinAttributeResponse, _err error) {
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
		Action:      dara.String("GetRecycleBinAttribute"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRecycleBinAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access point information.
//
// Description:
//
// Only General-purpose NAS NFS file systems are supported.
//
// @param request - ListAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessPointsResponse
func (client *Client) ListAccessPointsWithContext(ctx context.Context, request *ListAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *ListAccessPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessPoints"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the directory structure of EBS Lens.
//
// Description:
//
// - This feature is currently in invitational preview. To use this feature, submit a ticket.
//
// @param request - ListDataInsightDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataInsightDirectoriesResponse
func (client *Client) ListDataInsightDirectoriesWithContext(ctx context.Context, request *ListDataInsightDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *ListDataInsightDirectoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ParentDir) {
		query["ParentDir"] = request.ParentDir
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataInsightDirectories"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataInsightDirectoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of IA storage class files and subdirectories that contain IA storage class files under a specified directory in a General-purpose NAS file system.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// @param request - ListDirectoriesAndFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDirectoriesAndFilesResponse
func (client *Client) ListDirectoriesAndFilesWithContext(ctx context.Context, request *ListDirectoriesAndFilesRequest, runtime *dara.RuntimeOptions) (_result *ListDirectoriesAndFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DirectoryOnly) {
		query["DirectoryOnly"] = request.DirectoryOnly
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDirectoriesAndFiles"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDirectoriesAndFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of created data retrieval tasks.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// @param request - ListLifecycleRetrieveJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLifecycleRetrieveJobsResponse
func (client *Client) ListLifecycleRetrieveJobsWithContext(ctx context.Context, request *ListLifecycleRetrieveJobsRequest, runtime *dara.RuntimeOptions) (_result *ListLifecycleRetrieveJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLifecycleRetrieveJobs"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLifecycleRetrieveJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries directories on which delete operations were recently performed.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// **Before you begin**
//
// Before calling this operation, make sure that the recycle bin feature is enabled for the file system. If the recycle bin is not enabled, this operation cannot return the list of recently deleted directories.
//
// **Enable the recycle bin**
//
// To enable the recycle bin feature, call the [EnableRecycleBin](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-enablerecyclebin) operation. After the recycle bin is enabled, deleted directories in the file system are moved to the recycle bin. You can call this operation to query the list of recently deleted directories.
//
// @param request - ListRecentlyRecycledDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentlyRecycledDirectoriesResponse
func (client *Client) ListRecentlyRecycledDirectoriesWithContext(ctx context.Context, request *ListRecentlyRecycledDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *ListRecentlyRecycledDirectoriesResponse, _err error) {
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
		Action:      dara.String("ListRecentlyRecycledDirectories"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecentlyRecycledDirectoriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about one or more tasks in the recycle bin.
//
// Description:
//
// - Only General-purpose NAS file systems support this feature.
//
// - You can query up to 50 most recently executed tasks.
//
// @param request - ListRecycleBinJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecycleBinJobsResponse
func (client *Client) ListRecycleBinJobsWithContext(ctx context.Context, request *ListRecycleBinJobsRequest, runtime *dara.RuntimeOptions) (_result *ListRecycleBinJobsResponse, _err error) {
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
		Action:      dara.String("ListRecycleBinJobs"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecycleBinJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries deleted files or directories.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// @param request - ListRecycledDirectoriesAndFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecycledDirectoriesAndFilesResponse
func (client *Client) ListRecycledDirectoriesAndFilesWithContext(ctx context.Context, request *ListRecycledDirectoriesAndFilesRequest, runtime *dara.RuntimeOptions) (_result *ListRecycledDirectoriesAndFilesResponse, _err error) {
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
		Action:      dara.String("ListRecycledDirectoriesAndFiles"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecycledDirectoriesAndFilesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
		Version:     dara.String("2017-06-26"),
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
// Modifies a permission group.
//
// Description:
//
// The default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
//
// @param request - ModifyAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessGroupResponse
func (client *Client) ModifyAccessGroupWithContext(ctx context.Context, request *ModifyAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccessGroup"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies access point information.
//
// Description:
//
// Only General-purpose NAS NFS file systems support this feature.
//
// @param request - ModifyAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessPointResponse
func (client *Client) ModifyAccessPointWithContext(ctx context.Context, request *ModifyAccessPointRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroup) {
		query["AccessGroup"] = request.AccessGroup
	}

	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.AccessPointName) {
		query["AccessPointName"] = request.AccessPointName
	}

	if !dara.IsNil(request.EnabledRam) {
		query["EnabledRam"] = request.EnabledRam
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a permission rule in a permission group.
//
// Description:
//
// Rules in the default permission group (DEFAULT_VPC_GROUP_NAME) cannot be modified.
//
// @param request - ModifyAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessRuleResponse
func (client *Client) ModifyAccessRuleWithContext(ctx context.Context, request *ModifyAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.AccessRuleId) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !dara.IsNil(request.FileSystemType) {
		query["FileSystemType"] = request.FileSystemType
	}

	if !dara.IsNil(request.Ipv6SourceCidrIp) {
		query["Ipv6SourceCidrIp"] = request.Ipv6SourceCidrIp
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RWAccessType) {
		query["RWAccessType"] = request.RWAccessType
	}

	if !dara.IsNil(request.SourceCidrIp) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	if !dara.IsNil(request.UserAccessType) {
		query["UserAccessType"] = request.UserAccessType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccessRule"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an Agentic space.
//
// Description:
//
// *Prerequisites**
//
// Before you call this operation, complete the following preparations:
//
// 1. **Obtain an AgenticSpaceId**: Call the [CreateAgenticSpace](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-createagenticspace) operation to create an Agentic space and obtain the AgenticSpaceId.
//
// 2. **Create an Agentic file system**: When you call the [CreateFileSystem](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createfilesystem-cpfs) operation to create a file system, set the following parameters:
//
//   - `StorageType`: Set to `Agentic`.
//
//   - `ProtocolType`: Set to `NFS`.
//
//   - **Supported regions**: The Agentic storage type is available only in specific regions and zones. Check the console for details.
//
// **Usage notes**
//
// - This operation applies only to Agentic file systems.
//
// - Complete call chain example:
//
//  1. Call CreateFileSystem to create an Agentic file system (StorageType=Agentic, ProtocolType=NFS).
//
//  2. After the file system is created (status is Running), call CreateAgenticSpace to create an Agentic space.
//
//  3. Call ModifyAgenticSpace to modify the Agentic space configurations.
//
// @param request - ModifyAgenticSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAgenticSpaceResponse
func (client *Client) ModifyAgenticSpaceWithContext(ctx context.Context, request *ModifyAgenticSpaceRequest, runtime *dara.RuntimeOptions) (_result *ModifyAgenticSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAgenticSpace"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAgenticSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an automatic snapshot policy. After you modify an automatic snapshot policy, file systems to which the policy has been applied immediately execute the modified policy.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// <props="intl">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// @param request - ModifyAutoSnapshotPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoSnapshotPolicyResponse
func (client *Client) ModifyAutoSnapshotPolicyWithContext(ctx context.Context, request *ModifyAutoSnapshotPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyAutoSnapshotPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoSnapshotPolicyId) {
		query["AutoSnapshotPolicyId"] = request.AutoSnapshotPolicyId
	}

	if !dara.IsNil(request.AutoSnapshotPolicyName) {
		query["AutoSnapshotPolicyName"] = request.AutoSnapshotPolicyName
	}

	if !dara.IsNil(request.RepeatWeekdays) {
		query["RepeatWeekdays"] = request.RepeatWeekdays
	}

	if !dara.IsNil(request.RetentionDays) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !dara.IsNil(request.TimePoints) {
		query["TimePoints"] = request.TimePoints
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAutoSnapshotPolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAutoSnapshotPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an access point of a CPFS for Lingjun file system.
//
// Description:
//
// Modifies an access point of a CPFS for Lingjun file system.
//
// @param request - ModifyCpfsAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCpfsAccessPointResponse
func (client *Client) ModifyCpfsAccessPointWithContext(ctx context.Context, request *ModifyCpfsAccessPointRequest, runtime *dara.RuntimeOptions) (_result *ModifyCpfsAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCpfsAccessPoint"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCpfsAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a data flow.
//
// Description:
//
// - CPFS 2.2.0 and later and CPFS for Lingjun 2.4.0 and later support data streams.
//
// - You can modify the properties of a data stream only when the data stream is in the `Running (Normal)` state.
//
// - Modifying a data stream typically takes 2 to 5 minutes. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html) to query the status of the data stream modification.
//
// - CPFS data stream specifications:
//
//   - Data stream bandwidth (Throughput) supports three specifications: 600 MB/s, 1,200 MB/s, and 1,500 MB/s. The data stream bandwidth specifies the maximum transmission bandwidth that the data stream can achieve during data import or export.
//
//   - Inventory check: If you set DryRun to true, you can check whether the resources required to modify the data stream to the specified specification are sufficient.
//
// - CPFS billing
//
//	Modifying the data stream bandwidth (Throughput) incurs data stream bandwidth fees. Familiarize yourself with the billable methods of CPFS in advance. For details, see [CPFS billing](https://help.aliyun.com/document_detail/111858.html).
//
// @param request - ModifyDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataFlowResponse
func (client *Client) ModifyDataFlowWithContext(ctx context.Context, request *ModifyDataFlowRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataFlowResponse, _err error) {
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

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Throughput) {
		query["Throughput"] = request.Throughput
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDataFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the automatic update configuration of a dataflow.
//
// Description:
//
// - This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// - Only CPFS 2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
// - You can modify the automatic update configuration of a dataflow only when the dataflow is in the `Running` or `Stopped` state.
//
// - Modifying the automatic update configuration typically takes 2 to 5 minutes. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html) to query the status of the automatic update task.
//
// @param request - ModifyDataFlowAutoRefreshRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataFlowAutoRefreshResponse
func (client *Client) ModifyDataFlowAutoRefreshWithContext(ctx context.Context, request *ModifyDataFlowAutoRefreshRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataFlowAutoRefreshResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRefreshInterval) {
		query["AutoRefreshInterval"] = request.AutoRefreshInterval
	}

	if !dara.IsNil(request.AutoRefreshPolicy) {
		query["AutoRefreshPolicy"] = request.AutoRefreshPolicy
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataFlowAutoRefresh"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDataFlowAutoRefreshResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a file system.
//
// @param tmpReq - ModifyFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFileSystemResponse
func (client *Client) ModifyFileSystemWithContext(ctx context.Context, tmpReq *ModifyFileSystemRequest, runtime *dara.RuntimeOptions) (_result *ModifyFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyFileSystemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AutoUpgradeConfig) {
		request.AutoUpgradeConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AutoUpgradeConfig, dara.String("AutoUpgradeConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Options) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, dara.String("Options"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUpgradeConfigShrink) {
		query["AutoUpgradeConfig"] = request.AutoUpgradeConfigShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.OptionsShrink) {
		query["Options"] = request.OptionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a fileset.
//
// Description:
//
// Only CPFS 2.2.0 and later and CPFS for Lingjun 2.7.0 and later support modifying fileset information.
//
// @param request - ModifyFilesetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFilesetResponse
func (client *Client) ModifyFilesetWithContext(ctx context.Context, request *ModifyFilesetRequest, runtime *dara.RuntimeOptions) (_result *ModifyFilesetResponse, _err error) {
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

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFileset"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFilesetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyLDAPConfig is deprecated
//
// Summary:
//
// Modifies the LDAP configuration.
//
// Description:
//
// This operation supports only Cloud Parallel File Storage (CPFS) parallel file systems.
//
// @param request - ModifyLDAPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLDAPConfigResponse
func (client *Client) ModifyLDAPConfigWithContext(ctx context.Context, request *ModifyLDAPConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyLDAPConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindDN) {
		query["BindDN"] = request.BindDN
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.SearchBase) {
		query["SearchBase"] = request.SearchBase
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLDAPConfig"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLDAPConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a lifecycle management policy.
//
// Description:
//
// This feature is supported only by General-purpose NAS file systems.
//
// @param request - ModifyLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLifecyclePolicyResponse
func (client *Client) ModifyLifecyclePolicyWithContext(ctx context.Context, request *ModifyLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyLifecyclePolicyResponse, _err error) {
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

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
	}

	if !dara.IsNil(request.LifecyclePolicyName) {
		query["LifecyclePolicyName"] = request.LifecyclePolicyName
	}

	if !dara.IsNil(request.LifecycleRuleName) {
		query["LifecycleRuleName"] = request.LifecycleRuleName
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLifecyclePolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLifecyclePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies mount target information.
//
// Description:
//
// Only mount target information of General-purpose NAS or Extreme NAS file systems can be modified.
//
// @param request - ModifyMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMountTargetResponse
func (client *Client) ModifyMountTargetWithContext(ctx context.Context, request *ModifyMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ModifyMountTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.AccessPointAccessOnly) {
		query["AccessPointAccessOnly"] = request.AccessPointAccessOnly
	}

	if !dara.IsNil(request.DualStackMountTargetDomain) {
		query["DualStackMountTargetDomain"] = request.DualStackMountTargetDomain
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.MountTargetDomain) {
		query["MountTargetDomain"] = request.MountTargetDomain
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the export directory parameters of a protocol service. You can modify only the permission group and description. The VPC ID and vSwitch ID cannot be modified. To change them, delete the export directory and create a new one.
//
// Description:
//
// - Before you begin: This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// - Status prerequisite: Before calling this operation, the export directory corresponding to ExportId must be in the AVAILABLE state. Call [DescribeProtocolMountTarget](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describeprotocolmounttarget-cpfs) to query the status and wait until the creation is complete before retrying.
//
// - Dependency chain:
//
//  1. Call [DescribeFileSystems](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describefilesystems-cpfs) (with FileSystemType set to cpfs) to obtain the CPFS file system.
//
//  2. Call [DescribeProtocolService](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describeprotocolservice-cpfs) to verify that the protocol service is in the Running state.
//
//  3. Call [DescribeProtocolMountTarget](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describeprotocolmounttarget-cpfs) to verify that the export directory is in the AVAILABLE state.
//
//  4. If no export directory exists, call [CreateProtocolMountTarget](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createprotocolmounttarget-cpfs) to create one first.
//
// @param request - ModifyProtocolMountTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtocolMountTargetResponse
func (client *Client) ModifyProtocolMountTargetWithContext(ctx context.Context, request *ModifyProtocolMountTargetRequest, runtime *dara.RuntimeOptions) (_result *ModifyProtocolMountTargetResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ExportId) {
		query["ExportId"] = request.ExportId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyProtocolMountTarget"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyProtocolMountTargetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a protocol service. You can modify the description of a protocol service.
//
// Description:
//
// This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// - This operation allows you to modify the description of a protocol service.
//
// - You can obtain the protocol service ID (ProtocolServiceId) in the following ways:
//
//   - From the response of the [CreateProtocolService](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createprotocolservice-cpfs) operation after you create a protocol service.
//
//   - From the response of the [DescribeProtocolService](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describeprotocolservice-cpfs) operation that queries existing protocol services.
//
// @param request - ModifyProtocolServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyProtocolServiceResponse
func (client *Client) ModifyProtocolServiceWithContext(ctx context.Context, request *ModifyProtocolServiceRequest, runtime *dara.RuntimeOptions) (_result *ModifyProtocolServiceResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.ProtocolServiceId) {
		query["ProtocolServiceId"] = request.ProtocolServiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyProtocolService"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyProtocolServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the SMB AD ACL feature.
//
// Description:
//
// Before you begin: Make sure that Key Management Service (KMS) is activated in the target region before you call this operation. Call sequence: Activate KMS → Call EnableSmbAcl to initialize → Call ModifySmbAcl.
//
// @param request - ModifySmbAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmbAclResponse
func (client *Client) ModifySmbAclWithContext(ctx context.Context, request *ModifySmbAclRequest, runtime *dara.RuntimeOptions) (_result *ModifySmbAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableAnonymousAccess) {
		query["EnableAnonymousAccess"] = request.EnableAnonymousAccess
	}

	if !dara.IsNil(request.EncryptData) {
		query["EncryptData"] = request.EncryptData
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.HomeDirPath) {
		query["HomeDirPath"] = request.HomeDirPath
	}

	if !dara.IsNil(request.Keytab) {
		query["Keytab"] = request.Keytab
	}

	if !dara.IsNil(request.KeytabMd5) {
		query["KeytabMd5"] = request.KeytabMd5
	}

	if !dara.IsNil(request.RejectUnencryptedAccess) {
		query["RejectUnencryptedAccess"] = request.RejectUnencryptedAccess
	}

	if !dara.IsNil(request.SuperAdminSid) {
		query["SuperAdminSid"] = request.SuperAdminSid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySmbAcl"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySmbAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RemoveClientFromBlackList is deprecated
//
// Summary:
//
// Removes a client from the blacklist in Cloud Parallel File Storage (CPFS) to resume its write requests.
//
// Description:
//
// This operation supports only CPFS parallel file systems.
//
// @param request - RemoveClientFromBlackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClientFromBlackListResponse
func (client *Client) RemoveClientFromBlackListWithContext(ctx context.Context, request *RemoveClientFromBlackListRequest, runtime *dara.RuntimeOptions) (_result *RemoveClientFromBlackListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientIP) {
		query["ClientIP"] = request.ClientIP
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveClientFromBlackList"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveClientFromBlackListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back a file system to the state of a historical snapshot.
//
// Description:
//
// <props="china">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud201803061139_99860.html?spm=a2c4g.11186623.0.0.5c895ff2YPLrwe) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// - The file system must be in the Normal state.
//
// - The specified SnapshotId parameter must be a historical snapshot created from the same file system specified by FileSystemId.
//
// - Dependency chain:
//
// - (1) Grant the AliyunNASManageENIRole role in the Resource Access Management (RAM) console.
//
// - (2) Call CreateFileSystem (FileSystemType=extreme, StorageType=advance).
//
// - (3) Call DescribeFileSystems and poll until Status=Running.
//
// - (4) Call CreateSnapshot (FileSystemId=ID from the previous step).
//
// - (5) Call DescribeSnapshots and poll until Status=available.
//
// - (6) Call ResetFileSystem (FileSystemId, SnapshotId).
//
// <props="intl">
//
// - This feature is in free public preview. During the public preview, the [File Storage NAS Service-Level Agreement (SLA)](https://www.alibabacloud.com/help/legal/latest/network-attached-storage-service-level-agreement) is not guaranteed.
//
// - Only Advanced Extreme NAS file systems support this feature.
//
// - The file system must be in the Normal state.
//
// - The specified SnapshotId parameter must be a historical snapshot created from the same file system specified by FileSystemId.
//
// - Dependency chain:
//
// - (1) Grant the AliyunNASManageENIRole role in the Resource Access Management (RAM) console.
//
// - (2) Call CreateFileSystem (FileSystemType=extreme, StorageType=advance).
//
// - (3) Call DescribeFileSystems and poll until Status=Running.
//
// - (4) Call CreateSnapshot (FileSystemId=ID from the previous step).
//
// - (5) Call DescribeSnapshots and poll until Status=available.
//
// - (6) Call ResetFileSystem (FileSystemId, SnapshotId).
//
// @param request - ResetFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetFileSystemResponse
func (client *Client) ResetFileSystemWithContext(ctx context.Context, request *ResetFileSystemRequest, runtime *dara.RuntimeOptions) (_result *ResetFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retries a data retrieval task that is in the failed state.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// @param request - RetryLifecycleRetrieveJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryLifecycleRetrieveJobResponse
func (client *Client) RetryLifecycleRetrieveJobWithContext(ctx context.Context, request *RetryLifecycleRetrieveJobRequest, runtime *dara.RuntimeOptions) (_result *RetryLifecycleRetrieveJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryLifecycleRetrieveJob"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryLifecycleRetrieveJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the quota for an Agentic space.
//
// Description:
//
// *Before you begin**
//
// Before you call this operation, complete the following preparations:
//
// 1. **Obtain an AgenticSpaceId**: Call the [CreateAgenticSpace](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-createagenticspace) operation to create an Agentic space and obtain the AgenticSpaceId.
//
// 2. **Create an Agentic file system**: When you call the [CreateFileSystem](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createfilesystem-cpfs) operation to create a file system, set `StorageType=Agentic`.
//
//   - **Supported regions**: The Agentic storage type is available only in specific regions and zones. Check the console for details.
//
// **Usage notes**
//
// - This operation applies only to Agentic file systems.
//
// - Quota statistics have a 15-minute delay. The actual usage takes effect 15 minutes later.
//
// @param request - SetAgenticSpaceQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAgenticSpaceQuotaResponse
func (client *Client) SetAgenticSpaceQuotaWithContext(ctx context.Context, request *SetAgenticSpaceQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetAgenticSpaceQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgenticSpaceId) {
		query["AgenticSpaceId"] = request.AgenticSpaceId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileCountLimit) {
		query["FileCountLimit"] = request.FileCountLimit
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.SizeLimit) {
		query["SizeLimit"] = request.SizeLimit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAgenticSpaceQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAgenticSpaceQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a directory quota for a file system.
//
// Description:
//
// Only General-purpose NAS file systems support the directory quota feature.
//
// @param request - SetDirQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDirQuotaResponse
func (client *Client) SetDirQuotaWithContext(ctx context.Context, request *SetDirQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetDirQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileCountLimit) {
		query["FileCountLimit"] = request.FileCountLimit
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.QuotaType) {
		query["QuotaType"] = request.QuotaType
	}

	if !dara.IsNil(request.SizeLimit) {
		query["SizeLimit"] = request.SizeLimit
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDirQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDirQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a quota for a fileset.
//
// Description:
//
// - Only CPFS for Lingjun 2.7.0 and later support setting quotas for filesets.
//
// - The minimum fileset capacity quota is 10 GiB, and the capacity can be increased in increments of 1 GiB.
//
// - A fileset supports up to 10 billion files or directories. The minimum value is 10,000, and the value can be increased in increments of 1.
//
// - When you modify a directory quota, the quota capacity or file count must be greater than the used capacity or file count.
//
// - You must specify at least one of the capacity limit and the file count limit.
//
// - Quota statistics have a 15-minute delay. The actual usage takes effect 15 minutes later.
//
// @param request - SetFilesetQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetFilesetQuotaResponse
func (client *Client) SetFilesetQuotaWithContext(ctx context.Context, request *SetFilesetQuotaRequest, runtime *dara.RuntimeOptions) (_result *SetFilesetQuotaResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileCountLimit) {
		query["FileCountLimit"] = request.FileCountLimit
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FsetId) {
		query["FsetId"] = request.FsetId
	}

	if !dara.IsNil(request.SizeLimit) {
		query["SizeLimit"] = request.SizeLimit
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetFilesetQuota"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetFilesetQuotaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a data flow.
//
// Description:
//
// - This operation is applicable only to Cloud Parallel File Storage (CPFS) file systems.
//
// - Only CPFS 2.2.0 and later support data flows. You can view the version information on the file system details page in the console.
//
// - You can enable only a data flow that is in the `Stopped` state.
//
// - If DryRun is set to `true`, you can check whether sufficient resources are available to enable a data flow of the specified specifications. If inventory resources are insufficient, the data flow cannot be enabled.
//
// - Enabling a data flow typically takes 2 to 5 minutes. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/2402270.html) to query the data flow status.
//
// @param request - StartDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDataFlowResponse
func (client *Client) StartDataFlowWithContext(ctx context.Context, request *StartDataFlowRequest, runtime *dara.RuntimeOptions) (_result *StartDataFlowResponse, _err error) {
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

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDataFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts the execution of a lifecycle management policy.
//
// Description:
//
// Starts the execution of a lifecycle management policy.
//
// This operation is supported only when LifecyclePolicyType is set to OnDemand in the lifecycle management policy of a CPFS for Lingjun file system.
//
// @param request - StartLifecyclePolicyExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLifecyclePolicyExecutionResponse
func (client *Client) StartLifecyclePolicyExecutionWithContext(ctx context.Context, request *StartLifecyclePolicyExecutionRequest, runtime *dara.RuntimeOptions) (_result *StartLifecyclePolicyExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartLifecyclePolicyExecution"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLifecyclePolicyExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a dataflow.
//
// Description:
//
// - This operation applies only to Cloud Parallel File Storage (CPFS) file systems.
//
// - Only CPFS 2.2.0 and later support dataflows. You can view the version information on the file system details page in the console.
//
// - You can disable only a dataflow in the `Running` state.
//
// - After a dataflow is disabled, you cannot create dataflow tasks for it. If automatic update is configured, data updates on the source are not synchronized to CPFS.
//
// - After a dataflow is disabled, the dataflow bandwidth is no longer billed because resources are reclaimed. However, re-enabling the dataflow may fail due to insufficient inventory.
//
// - Disabling a dataflow typically takes 2 to 5 minutes. You can call [DescribeDataFlows](https://help.aliyun.com/document_detail/2402271.html) to query the dataflow status.
//
// @param request - StopDataFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDataFlowResponse
func (client *Client) StopDataFlowWithContext(ctx context.Context, request *StopDataFlowRequest, runtime *dara.RuntimeOptions) (_result *StopDataFlowResponse, _err error) {
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

	if !dara.IsNil(request.DataFlowId) {
		query["DataFlowId"] = request.DataFlowId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDataFlow"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDataFlowResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the execution of a lifecycle policy.
//
// Description:
//
// Stops the execution of a lifecycle policy.
//
// This operation is supported only when LifecyclePolicyType is set to OnDemand in the lifecycle management policy of a CPFS for Lingjun file system.
//
// @param request - StopLifecyclePolicyExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLifecyclePolicyExecutionResponse
func (client *Client) StopLifecyclePolicyExecutionWithContext(ctx context.Context, request *StopLifecyclePolicyExecutionRequest, runtime *dara.RuntimeOptions) (_result *StopLifecyclePolicyExecutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopLifecyclePolicyExecution"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLifecyclePolicyExecutionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and attaches tags to specified resources. File systems and access points are supported.
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
		Version:     dara.String("2017-06-26"),
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
// Deletes the tags of a specified resource.
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
		Version:     dara.String("2017-06-26"),
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
// Updates the rules of a lifecycle management policy. Only CPFS for Lingjun is supported.
//
// UpdateLifecyclePolicy uses overwrite semantics. If an optional parameter is left empty, the corresponding field is deleted. To append lifecycle rules, first call the DescribeLifecyclePolicies operation to obtain the current lifecycle rule configurations, append the new lifecycle rule configurations, and then call the UpdateLifecyclePolicy operation to update the lifecycle rule configurations.
//
// @param request - UpdateLifecyclePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLifecyclePolicyResponse
func (client *Client) UpdateLifecyclePolicyWithContext(ctx context.Context, request *UpdateLifecyclePolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateLifecyclePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteRules) {
		query["DeleteRules"] = request.DeleteRules
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.LifecyclePolicyId) {
		query["LifecyclePolicyId"] = request.LifecyclePolicyId
	}

	if !dara.IsNil(request.Paths) {
		query["Paths"] = request.Paths
	}

	if !dara.IsNil(request.RetrieveRules) {
		query["RetrieveRules"] = request.RetrieveRules
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.TransitRules) {
		query["TransitRules"] = request.TransitRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLifecyclePolicy"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLifecyclePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the retention period of data in a recycle bin.
//
// Description:
//
// Only General-purpose NAS file systems support this feature.
//
// @param request - UpdateRecycleBinAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecycleBinAttributeResponse
func (client *Client) UpdateRecycleBinAttributeWithContext(ctx context.Context, request *UpdateRecycleBinAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateRecycleBinAttributeResponse, _err error) {
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
		Action:      dara.String("UpdateRecycleBinAttribute"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRecycleBinAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales up the storage capacity of an Extreme NAS file system or a Cloud Parallel File Storage (CPFS) file system.
//
// Description:
//
// - Only Extreme NAS file systems and CPFS file systems can be scaled up.
//
// - General-purpose NAS file systems scale up automatically on demand. You do not need to call this operation.
//
// @param request - UpgradeFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeFileSystemResponse
func (client *Client) UpgradeFileSystemWithContext(ctx context.Context, request *UpgradeFileSystemRequest, runtime *dara.RuntimeOptions) (_result *UpgradeFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Capacity) {
		query["Capacity"] = request.Capacity
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeFileSystem"),
		Version:     dara.String("2017-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
