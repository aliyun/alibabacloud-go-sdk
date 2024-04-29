// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChangeResourceGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// dbs
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the resource group to which you want to move the resource.
	//
	// example:
	//
	// rg-aekz4kee6******
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// dbs1jyajqk******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Set the value to backupplan.
	//
	// example:
	//
	// backupplan
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetClientToken(v string) *ChangeResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionCode(v string) *ChangeResourceGroupRequest {
	s.RegionCode = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the resource was successfully moved. Valid values:
	//
	// 	- **true**: The resource was successfully moved.
	//
	// 	- **false**: The resource failed to be moved.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// Request.Forbidden
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// RAM DENY
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The additional information.
	//
	// example:
	//
	// The resource group is forbidden to operate
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04EBD9F5-F06F-5302-8499-005C72*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetCode(v string) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetData(v string) *ChangeResourceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetErrCode(v string) *ChangeResourceGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetErrMessage(v string) *ChangeResourceGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetSuccess(v string) *ChangeResourceGroupResponseBody {
	s.Success = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CreateDownloadRequest struct {
	// The ID of the backup set. You can call the [DescribeBackups](~~26273~~) operation to query the ID of the backup set.
	//
	// > This parameter is required if the BakSetType parameter is set to full.
	//
	// example:
	//
	// 146005****
	BakSetId *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	// The size of the full backup set. Unit: bytes. You can call the [DescribeBackups](~~26273~~) operation to query the size of the full backup set.
	//
	// example:
	//
	// 216****
	BakSetSize *string `json:"BakSetSize,omitempty" xml:"BakSetSize,omitempty"`
	// The type of the download task. Valid values:
	//
	// 	- **full**: downloads a full backup set.
	//
	// 	- **pitr**: downloads a backup set at a specific point in time.
	//
	// example:
	//
	// full
	BakSetType *string `json:"BakSetType,omitempty" xml:"BakSetType,omitempty"`
	// The point in time at which the backup set is downloaded. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > This parameter is required if the BakSetType parameter is set to pitr.
	//
	// example:
	//
	// 1661331864000
	DownloadPointInTime *string `json:"DownloadPointInTime,omitempty" xml:"DownloadPointInTime,omitempty"`
	// The format to which the downloaded backup set is converted. Valid values:
	//
	// 	- **CSV**
	//
	// 	- **SQL**
	//
	// 	- **Parquet**
	//
	// > This parameter is required.
	//
	// example:
	//
	// CSV
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-wz994c1t1****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](~~26231~~) operation to query the region ID of the instance.
	//
	// example:
	//
	// cn-beijing
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The name of the OSS bucket that is used to store the backup set.
	//
	// 	- This parameter is required if the TargetType parameter is set to OSS.
	//
	// 	- Make sure that your account is granted the **AliyunDBSDefaultRole*	- permission. For more information, see [Use RAM for resource authorization](~~26307~~). You can also grant permissions based on the operation instructions in the Resource Access Management (RAM) console.
	//
	// example:
	//
	// test123
	TargetBucket *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	// The region in which the OSS bucket resides.
	//
	// > This parameter is required if the TargetType parameter is set to OSS.
	//
	// example:
	//
	// cn-beijing
	TargetOssRegion *string `json:"TargetOssRegion,omitempty" xml:"TargetOssRegion,omitempty"`
	// The destination path to which the backup set is downloaded.
	//
	// > This parameter is required if the TargetType parameter is set to OSS.
	//
	// example:
	//
	// test_db/path
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the destination to which the backup set is downloaded. Valid values:
	//
	// 	- **OSS**
	//
	// 	- **URL**
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateDownloadRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadRequest) GoString() string {
	return s.String()
}

func (s *CreateDownloadRequest) SetBakSetId(v string) *CreateDownloadRequest {
	s.BakSetId = &v
	return s
}

func (s *CreateDownloadRequest) SetBakSetSize(v string) *CreateDownloadRequest {
	s.BakSetSize = &v
	return s
}

func (s *CreateDownloadRequest) SetBakSetType(v string) *CreateDownloadRequest {
	s.BakSetType = &v
	return s
}

func (s *CreateDownloadRequest) SetDownloadPointInTime(v string) *CreateDownloadRequest {
	s.DownloadPointInTime = &v
	return s
}

func (s *CreateDownloadRequest) SetFormatType(v string) *CreateDownloadRequest {
	s.FormatType = &v
	return s
}

func (s *CreateDownloadRequest) SetInstanceName(v string) *CreateDownloadRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateDownloadRequest) SetRegionCode(v string) *CreateDownloadRequest {
	s.RegionCode = &v
	return s
}

func (s *CreateDownloadRequest) SetTargetBucket(v string) *CreateDownloadRequest {
	s.TargetBucket = &v
	return s
}

func (s *CreateDownloadRequest) SetTargetOssRegion(v string) *CreateDownloadRequest {
	s.TargetOssRegion = &v
	return s
}

func (s *CreateDownloadRequest) SetTargetPath(v string) *CreateDownloadRequest {
	s.TargetPath = &v
	return s
}

func (s *CreateDownloadRequest) SetTargetType(v string) *CreateDownloadRequest {
	s.TargetType = &v
	return s
}

type CreateDownloadResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// DBS.ParamIsInValid
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the download task.
	Data *CreateDownloadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// DBS.ParamIsInValid
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// formatType can not be empty
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// formatType can not be empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A08F908D-2C35-583F-93C1-ED80753F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDownloadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDownloadResponseBody) SetCode(v string) *CreateDownloadResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDownloadResponseBody) SetData(v *CreateDownloadResponseBodyData) *CreateDownloadResponseBody {
	s.Data = v
	return s
}

func (s *CreateDownloadResponseBody) SetErrCode(v string) *CreateDownloadResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDownloadResponseBody) SetErrMessage(v string) *CreateDownloadResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateDownloadResponseBody) SetMessage(v string) *CreateDownloadResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDownloadResponseBody) SetRequestId(v string) *CreateDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDownloadResponseBody) SetSuccess(v string) *CreateDownloadResponseBody {
	s.Success = &v
	return s
}

type CreateDownloadResponseBodyData struct {
	// The point in time of the backup set if the task is used to download a backup set at a specific point in time. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1661373070000
	BackupSetTime *int64 `json:"BackupSetTime,omitempty" xml:"BackupSetTime,omitempty"`
	// The ID of the full backup set.
	//
	// example:
	//
	// 146005****
	BakSetId *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	// The database and table information that is returned if databases and tables are filtered by the download task.
	//
	// example:
	//
	// testdb
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The state of the download task. Valid values:
	//
	// 	- initializing: The download task was being initialized.
	//
	// 	- queuing: The download task was queuing.
	//
	// 	- running: The download task was running.
	//
	// 	- failed: The download task failed.
	//
	// 	- finished: The download task was complete.
	//
	// 	- expired: The download task expired.
	//
	// > If the TargetType parameter is set to URL, the download task expires in three days after the task is complete.
	//
	// example:
	//
	// initializing
	DownloadStatus *string `json:"DownloadStatus,omitempty" xml:"DownloadStatus,omitempty"`
	// The size of the downloaded data. Unit: bytes.
	//
	// example:
	//
	// 0
	ExportDataSize *int64 `json:"ExportDataSize,omitempty" xml:"ExportDataSize,omitempty"`
	// The format to which the downloaded data is converted.
	//
	// example:
	//
	// CSV
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The time when the download task was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1661940917570
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The size of the processed data. Unit: bytes.
	//
	// example:
	//
	// 0
	ImportDataSize *int64 `json:"ImportDataSize,omitempty" xml:"ImportDataSize,omitempty"`
	// The number of tables that have been downloaded and the total number of tables to be downloaded.
	//
	// > If the task is in the preparation stage, 0/0 is returned.
	//
	// example:
	//
	// 0/0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-beijing
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The destination path to which the backup set is downloaded.
	//
	// >  This parameter is returned if the value of **TargetType is OSS**.
	//
	// example:
	//
	// test_db/path
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the destination to which the backup set is downloaded.
	//
	// example:
	//
	// URL
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the download task.
	//
	// example:
	//
	// dt-qxnsfq5s****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDownloadResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDownloadResponseBodyData) SetBackupSetTime(v int64) *CreateDownloadResponseBodyData {
	s.BackupSetTime = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetBakSetId(v string) *CreateDownloadResponseBodyData {
	s.BakSetId = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetDbList(v string) *CreateDownloadResponseBodyData {
	s.DbList = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetDownloadStatus(v string) *CreateDownloadResponseBodyData {
	s.DownloadStatus = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetExportDataSize(v int64) *CreateDownloadResponseBodyData {
	s.ExportDataSize = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetFormat(v string) *CreateDownloadResponseBodyData {
	s.Format = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetGmtCreate(v int64) *CreateDownloadResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetImportDataSize(v int64) *CreateDownloadResponseBodyData {
	s.ImportDataSize = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetProgress(v string) *CreateDownloadResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetRegionCode(v string) *CreateDownloadResponseBodyData {
	s.RegionCode = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetTargetPath(v string) *CreateDownloadResponseBodyData {
	s.TargetPath = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetTargetType(v string) *CreateDownloadResponseBodyData {
	s.TargetType = &v
	return s
}

func (s *CreateDownloadResponseBodyData) SetTaskId(v string) *CreateDownloadResponseBodyData {
	s.TaskId = &v
	return s
}

type CreateDownloadResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDownloadResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadResponse) GoString() string {
	return s.String()
}

func (s *CreateDownloadResponse) SetHeaders(v map[string]*string) *CreateDownloadResponse {
	s.Headers = v
	return s
}

func (s *CreateDownloadResponse) SetStatusCode(v int32) *CreateDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDownloadResponse) SetBody(v *CreateDownloadResponseBody) *CreateDownloadResponse {
	s.Body = v
	return s
}

type DeleteSandboxInstanceRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](~~437215~~) operation to query the ID of the backup schedule.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](~~193091~~) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	//
	// example:
	//
	// 1hxxxx8xxxxxa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the sandbox instance. You can call the [DescribeSandboxInstances](~~437257~~) operation to query the ID of the sandbox instance.
	//
	// example:
	//
	// 1jxxxxnxxx1xc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteSandboxInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSandboxInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSandboxInstanceRequest) SetBackupPlanId(v string) *DeleteSandboxInstanceRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DeleteSandboxInstanceRequest) SetInstanceId(v string) *DeleteSandboxInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSandboxInstanceRequest) SetZoneId(v string) *DeleteSandboxInstanceRequest {
	s.ZoneId = &v
	return s
}

type DeleteSandboxInstanceResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// operation forbidden due to sandbox is creating.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSandboxInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSandboxInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSandboxInstanceResponseBody) SetCode(v string) *DeleteSandboxInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetData(v string) *DeleteSandboxInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetErrCode(v string) *DeleteSandboxInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetErrMessage(v string) *DeleteSandboxInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetMessage(v string) *DeleteSandboxInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetRequestId(v string) *DeleteSandboxInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSandboxInstanceResponseBody) SetSuccess(v string) *DeleteSandboxInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteSandboxInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSandboxInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSandboxInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSandboxInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSandboxInstanceResponse) SetHeaders(v map[string]*string) *DeleteSandboxInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSandboxInstanceResponse) SetStatusCode(v int32) *DeleteSandboxInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSandboxInstanceResponse) SetBody(v *DeleteSandboxInstanceResponseBody) *DeleteSandboxInstanceResponse {
	s.Body = v
	return s
}

type DescribeBackupDataListRequest struct {
	// example:
	//
	// 213064****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// Snapshot
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// example:
	//
	// DBInstance
	BackupScale *string `json:"BackupScale,omitempty" xml:"BackupScale,omitempty"`
	// example:
	//
	// OK
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// test****
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// 2024-04-17T17:00:32Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// false
	InstanceIsDeleted *bool `json:"InstanceIsDeleted,omitempty" xml:"InstanceIsDeleted,omitempty"`
	// example:
	//
	// pc-2ze3nrr64c5******
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	InstanceRegion *string `json:"InstanceRegion,omitempty" xml:"InstanceRegion,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// example:
	//
	// LEVEL_1
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	// example:
	//
	// 2024-04-17T17:00:16Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupDataListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDataListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListRequest) SetBackupId(v string) *DescribeBackupDataListRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupMethod(v string) *DescribeBackupDataListRequest {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupMode(v string) *DescribeBackupDataListRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupScale(v string) *DescribeBackupDataListRequest {
	s.BackupScale = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupStatus(v string) *DescribeBackupDataListRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetBackupType(v string) *DescribeBackupDataListRequest {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetDataSourceId(v string) *DescribeBackupDataListRequest {
	s.DataSourceId = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetEndTime(v string) *DescribeBackupDataListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetInstanceIsDeleted(v bool) *DescribeBackupDataListRequest {
	s.InstanceIsDeleted = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetInstanceName(v string) *DescribeBackupDataListRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetInstanceRegion(v string) *DescribeBackupDataListRequest {
	s.InstanceRegion = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetPageNumber(v int32) *DescribeBackupDataListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetPageSize(v int32) *DescribeBackupDataListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetRegionCode(v string) *DescribeBackupDataListRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetSceneType(v string) *DescribeBackupDataListRequest {
	s.SceneType = &v
	return s
}

func (s *DescribeBackupDataListRequest) SetStartTime(v string) *DescribeBackupDataListRequest {
	s.StartTime = &v
	return s
}

type DescribeBackupDataListResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeBackupDataListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Request.Forbidden
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 95A5FFD0-7F46-5A7D-9DFE-6A376B4E2A28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupDataListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDataListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponseBody) SetCode(v string) *DescribeBackupDataListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetData(v *DescribeBackupDataListResponseBodyData) *DescribeBackupDataListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetErrCode(v string) *DescribeBackupDataListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetErrMessage(v string) *DescribeBackupDataListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetMessage(v string) *DescribeBackupDataListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetRequestId(v string) *DescribeBackupDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupDataListResponseBody) SetSuccess(v string) *DescribeBackupDataListResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupDataListResponseBodyData struct {
	Content []*DescribeBackupDataListResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// dbtest
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// example:
	//
	// 1
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeBackupDataListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDataListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponseBodyData) SetContent(v []*DescribeBackupDataListResponseBodyDataContent) *DescribeBackupDataListResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetExtra(v string) *DescribeBackupDataListResponseBodyData {
	s.Extra = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetPageNumber(v int64) *DescribeBackupDataListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetPageSize(v int64) *DescribeBackupDataListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetTotalElements(v int64) *DescribeBackupDataListResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyData) SetTotalPages(v int64) *DescribeBackupDataListResponseBodyData {
	s.TotalPages = &v
	return s
}

type DescribeBackupDataListResponseBodyDataContent struct {
	// example:
	//
	// http://oss.com/****
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// example:
	//
	// 2024-04-17T17:00:32Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// example:
	//
	// 213088****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// http://oss.com/****
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	// example:
	//
	// logic
	BackupLocation *string `json:"BackupLocation,omitempty" xml:"BackupLocation,omitempty"`
	// example:
	//
	// Snapshot
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// example:
	//
	// logic_backup
	BackupName *string `json:"BackupName,omitempty" xml:"BackupName,omitempty"`
	// example:
	//
	// DBInstance
	BackupScale *string `json:"BackupScale,omitempty" xml:"BackupScale,omitempty"`
	// example:
	//
	// 25669140480
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// example:
	//
	// 2024-04-17T17:00:16Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// example:
	//
	// OK
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// 84a4c16431f969712e6895992719****
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// example:
	//
	// 1713373221
	ConsistentTime *int64 `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// example:
	//
	// psk2
	Encryption *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	// example:
	//
	// polardb_mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 2024-04-19T05:00:49Z
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	// example:
	//
	// DELAY
	ExpectExpireType *string `json:"ExpectExpireType,omitempty" xml:"ExpectExpireType,omitempty"`
	// example:
	//
	// pc-2ze3nrr64c5******
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 1
	IsAvail       *int32                                                      `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	PolarSnapshot *DescribeBackupDataListResponseBodyDataContentPolarSnapshot `json:"PolarSnapshot,omitempty" xml:"PolarSnapshot,omitempty" type:"Struct"`
	// example:
	//
	// 0
	SupportDeletion *int32 `json:"SupportDeletion,omitempty" xml:"SupportDeletion,omitempty"`
}

func (s DescribeBackupDataListResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDataListResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupDownloadURL(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupEndTime(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupId(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupIntranetDownloadURL(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupLocation(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupLocation = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupMethod(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupMode(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupName(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupName = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupScale(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupScale = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupSize(v int64) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupStartTime(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupStatus(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetBackupType(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetChecksum(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.Checksum = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetConsistentTime(v int64) *DescribeBackupDataListResponseBodyDataContent {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetEncryption(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.Encryption = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetEngine(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.Engine = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetEngineVersion(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.EngineVersion = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetExpectExpireTime(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.ExpectExpireTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetExpectExpireType(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.ExpectExpireType = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetInstanceName(v string) *DescribeBackupDataListResponseBodyDataContent {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetIsAvail(v int32) *DescribeBackupDataListResponseBodyDataContent {
	s.IsAvail = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetPolarSnapshot(v *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) *DescribeBackupDataListResponseBodyDataContent {
	s.PolarSnapshot = v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContent) SetSupportDeletion(v int32) *DescribeBackupDataListResponseBodyDataContent {
	s.SupportDeletion = &v
	return s
}

type DescribeBackupDataListResponseBodyDataContentPolarSnapshot struct {
	// example:
	//
	// abc****
	DumpId *int64 `json:"DumpId,omitempty" xml:"DumpId,omitempty"`
	// example:
	//
	// 25669140589
	DumpSize *int64 `json:"DumpSize,omitempty" xml:"DumpSize,omitempty"`
	// example:
	//
	// 2024-04-19T05:00:49Z
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	// example:
	//
	// DELAY
	ExpectExpireType *string `json:"expectExpireType,omitempty" xml:"expectExpireType,omitempty"`
}

func (s DescribeBackupDataListResponseBodyDataContentPolarSnapshot) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDataListResponseBodyDataContentPolarSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) SetDumpId(v int64) *DescribeBackupDataListResponseBodyDataContentPolarSnapshot {
	s.DumpId = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) SetDumpSize(v int64) *DescribeBackupDataListResponseBodyDataContentPolarSnapshot {
	s.DumpSize = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) SetExpectExpireTime(v string) *DescribeBackupDataListResponseBodyDataContentPolarSnapshot {
	s.ExpectExpireTime = &v
	return s
}

func (s *DescribeBackupDataListResponseBodyDataContentPolarSnapshot) SetExpectExpireType(v string) *DescribeBackupDataListResponseBodyDataContentPolarSnapshot {
	s.ExpectExpireType = &v
	return s
}

type DescribeBackupDataListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupDataListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupDataListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupDataListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupDataListResponse) SetHeaders(v map[string]*string) *DescribeBackupDataListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupDataListResponse) SetStatusCode(v int32) *DescribeBackupDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupDataListResponse) SetBody(v *DescribeBackupDataListResponseBody) *DescribeBackupDataListResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	// example:
	//
	// pc-2ze3nrr64c5******
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetInstanceName(v string) *DescribeBackupPolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetRegionCode(v string) *DescribeBackupPolicyRequest {
	s.RegionCode = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	// example:
	//
	// Success
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeBackupPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// instanceName can not be empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 54A63B3B-AA10-1CC3-A6BB-6CCE98D19628
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetCode(v string) *DescribeBackupPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetData(v *DescribeBackupPolicyResponseBodyData) *DescribeBackupPolicyResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetErrCode(v string) *DescribeBackupPolicyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetErrMessage(v string) *DescribeBackupPolicyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetMessage(v string) *DescribeBackupPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSuccess(v string) *DescribeBackupPolicyResponseBody {
	s.Success = &v
	return s
}

type DescribeBackupPolicyResponseBodyData struct {
	AdvanceDataPolicies []*DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies `json:"AdvanceDataPolicies,omitempty" xml:"AdvanceDataPolicies,omitempty" type:"Repeated"`
	AdvanceLogPolicies  []*DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies  `json:"AdvanceLogPolicies,omitempty" xml:"AdvanceLogPolicies,omitempty" type:"Repeated"`
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// example:
	//
	// 0
	BackupPriority *int32 `json:"BackupPriority,omitempty" xml:"BackupPriority,omitempty"`
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// example:
	//
	// LATEST
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// 1
	EnableBackup *int32 `json:"EnableBackup,omitempty" xml:"EnableBackup,omitempty"`
	// example:
	//
	// 0
	EnableIncBackup *int32 `json:"EnableIncBackup,omitempty" xml:"EnableIncBackup,omitempty"`
	// example:
	//
	// 1
	EnableLogBackup *int32 `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	// example:
	//
	// 120
	HighFrequencyBakInterval *int32 `json:"HighFrequencyBakInterval,omitempty" xml:"HighFrequencyBakInterval,omitempty"`
	// example:
	//
	// Enable
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// example:
	//
	// -1
	IncBackupInterval *int32 `json:"IncBackupInterval,omitempty" xml:"IncBackupInterval,omitempty"`
	// example:
	//
	// 30
	LocalLogRetentionSpace *int32 `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// example:
	//
	// 10
	LogBackupLocalRetentionNumber *string `json:"LogBackupLocalRetentionNumber,omitempty" xml:"LogBackupLocalRetentionNumber,omitempty"`
	// example:
	//
	// 7
	LogBackupRetention *int32 `json:"LogBackupRetention,omitempty" xml:"LogBackupRetention,omitempty"`
	// example:
	//
	// 1010101
	PreferredBackupDate *string `json:"PreferredBackupDate,omitempty" xml:"PreferredBackupDate,omitempty"`
	// example:
	//
	// 23:00Z-24:00Z
	PreferredBackupWindow *string `json:"PreferredBackupWindow,omitempty" xml:"PreferredBackupWindow,omitempty"`
	// example:
	//
	// 23:00Z
	PreferredBackupWindowBegin *string `json:"PreferredBackupWindowBegin,omitempty" xml:"PreferredBackupWindowBegin,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyData) SetAdvanceDataPolicies(v []*DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) *DescribeBackupPolicyResponseBodyData {
	s.AdvanceDataPolicies = v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetAdvanceLogPolicies(v []*DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) *DescribeBackupPolicyResponseBodyData {
	s.AdvanceLogPolicies = v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupMethod(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupPriority(v int32) *DescribeBackupPolicyResponseBodyData {
	s.BackupPriority = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupRetentionPeriod(v int32) *DescribeBackupPolicyResponseBodyData {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetBackupRetentionPolicyOnClusterDeletion(v string) *DescribeBackupPolicyResponseBodyData {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetCategory(v string) *DescribeBackupPolicyResponseBodyData {
	s.Category = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetEnableBackup(v int32) *DescribeBackupPolicyResponseBodyData {
	s.EnableBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetEnableIncBackup(v int32) *DescribeBackupPolicyResponseBodyData {
	s.EnableIncBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetEnableLogBackup(v int32) *DescribeBackupPolicyResponseBodyData {
	s.EnableLogBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetHighFrequencyBakInterval(v int32) *DescribeBackupPolicyResponseBodyData {
	s.HighFrequencyBakInterval = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetHighSpaceUsageProtection(v string) *DescribeBackupPolicyResponseBodyData {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetIncBackupInterval(v int32) *DescribeBackupPolicyResponseBodyData {
	s.IncBackupInterval = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLocalLogRetentionSpace(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLogBackupLocalRetentionNumber(v string) *DescribeBackupPolicyResponseBodyData {
	s.LogBackupLocalRetentionNumber = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetLogBackupRetention(v int32) *DescribeBackupPolicyResponseBodyData {
	s.LogBackupRetention = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetPreferredBackupDate(v string) *DescribeBackupPolicyResponseBodyData {
	s.PreferredBackupDate = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetPreferredBackupWindow(v string) *DescribeBackupPolicyResponseBodyData {
	s.PreferredBackupWindow = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyData) SetPreferredBackupWindowBegin(v string) *DescribeBackupPolicyResponseBodyData {
	s.PreferredBackupWindowBegin = &v
	return s
}

type DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies struct {
	// example:
	//
	// true
	AutoCreated *bool `json:"AutoCreated,omitempty" xml:"AutoCreated,omitempty"`
	// example:
	//
	// F
	BakType *string `json:"BakType,omitempty" xml:"BakType,omitempty"`
	// example:
	//
	// cn-beijing
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// example:
	//
	// level1
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// example:
	//
	// copy
	DumpAction *string `json:"DumpAction,omitempty" xml:"DumpAction,omitempty"`
	// example:
	//
	// dayOfWeek
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// example:
	//
	// crontab
	FilterType *string `json:"FilterType,omitempty" xml:"FilterType,omitempty"`
	// example:
	//
	// 1,2,3,4,5,6,7
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// example:
	//
	// delay
	RetentionType *string `json:"RetentionType,omitempty" xml:"RetentionType,omitempty"`
	// example:
	//
	// 7
	RetentionValue *string `json:"RetentionValue,omitempty" xml:"RetentionValue,omitempty"`
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// example:
	//
	// db
	SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	// example:
	//
	// 71930ac2e9f15e41615e10627c******
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetAutoCreated(v bool) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.AutoCreated = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetBakType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.BakType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetDestRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetDestType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.DestType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetDumpAction(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.DumpAction = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetFilterKey(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.FilterKey = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetFilterType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.FilterType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetFilterValue(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.FilterValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetRetentionType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.RetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetRetentionValue(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.RetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetSrcRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetSrcType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.SrcType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies) SetStrategyId(v string) *DescribeBackupPolicyResponseBodyDataAdvanceDataPolicies {
	s.StrategyId = &v
	return s
}

type DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies struct {
	// example:
	//
	// cn-shanghai
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// example:
	//
	// level1
	DestType *string `json:"DestType,omitempty" xml:"DestType,omitempty"`
	// example:
	//
	// 1
	EnableLogBackup *bool `json:"EnableLogBackup,omitempty" xml:"EnableLogBackup,omitempty"`
	// example:
	//
	// delay
	LogRetentionType *string `json:"LogRetentionType,omitempty" xml:"LogRetentionType,omitempty"`
	// example:
	//
	// 3
	LogRetentionValue *string `json:"LogRetentionValue,omitempty" xml:"LogRetentionValue,omitempty"`
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// example:
	//
	// level1
	SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	// example:
	//
	// 6s67c7i3y8f8p72808p******
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetDestRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.DestRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetDestType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.DestType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetEnableLogBackup(v bool) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.EnableLogBackup = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetLogRetentionType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.LogRetentionType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetLogRetentionValue(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.LogRetentionValue = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetSrcRegion(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetSrcType(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.SrcType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies) SetStrategyId(v string) *DescribeBackupPolicyResponseBodyDataAdvanceLogPolicies {
	s.StrategyId = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetStatusCode(v int32) *DescribeBackupPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeDBTablesRecoveryBackupSetRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeDBTablesRecoveryBackupSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBTablesRecoveryBackupSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryBackupSetRequest) SetInstanceId(v string) *DescribeDBTablesRecoveryBackupSetRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetRequest) SetRegionCode(v string) *DescribeDBTablesRecoveryBackupSetRequest {
	s.RegionCode = &v
	return s
}

type DescribeDBTablesRecoveryBackupSetResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBTablesRecoveryBackupSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBTablesRecoveryBackupSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetCode(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetData(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetErrCode(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetErrMessage(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetMessage(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetRequestId(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetSuccess(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.Success = &v
	return s
}

type DescribeDBTablesRecoveryBackupSetResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBTablesRecoveryBackupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBTablesRecoveryBackupSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBTablesRecoveryBackupSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) SetHeaders(v map[string]*string) *DescribeDBTablesRecoveryBackupSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) SetStatusCode(v int32) *DescribeDBTablesRecoveryBackupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) SetBody(v *DescribeDBTablesRecoveryBackupSetResponseBody) *DescribeDBTablesRecoveryBackupSetResponse {
	s.Body = v
	return s
}

type DescribeDBTablesRecoveryStateRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeDBTablesRecoveryStateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBTablesRecoveryStateRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryStateRequest) SetInstanceId(v string) *DescribeDBTablesRecoveryStateRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateRequest) SetRegionCode(v string) *DescribeDBTablesRecoveryStateRequest {
	s.RegionCode = &v
	return s
}

type DescribeDBTablesRecoveryStateResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBTablesRecoveryStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBTablesRecoveryStateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetCode(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetData(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetErrCode(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetErrMessage(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetMessage(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetRequestId(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetSuccess(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.Success = &v
	return s
}

type DescribeDBTablesRecoveryStateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBTablesRecoveryStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBTablesRecoveryStateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBTablesRecoveryStateResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryStateResponse) SetHeaders(v map[string]*string) *DescribeDBTablesRecoveryStateResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponse) SetStatusCode(v int32) *DescribeDBTablesRecoveryStateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponse) SetBody(v *DescribeDBTablesRecoveryStateResponseBody) *DescribeDBTablesRecoveryStateResponse {
	s.Body = v
	return s
}

type DescribeDBTablesRecoveryTimeRangeRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeDBTablesRecoveryTimeRangeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBTablesRecoveryTimeRangeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryTimeRangeRequest) SetInstanceId(v string) *DescribeDBTablesRecoveryTimeRangeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeRequest) SetRegionCode(v string) *DescribeDBTablesRecoveryTimeRangeRequest {
	s.RegionCode = &v
	return s
}

type DescribeDBTablesRecoveryTimeRangeResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBTablesRecoveryTimeRangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBTablesRecoveryTimeRangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetCode(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetData(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetErrCode(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetErrMessage(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetMessage(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetRequestId(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetSuccess(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.Success = &v
	return s
}

type DescribeDBTablesRecoveryTimeRangeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBTablesRecoveryTimeRangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBTablesRecoveryTimeRangeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBTablesRecoveryTimeRangeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) SetHeaders(v map[string]*string) *DescribeDBTablesRecoveryTimeRangeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) SetStatusCode(v int32) *DescribeDBTablesRecoveryTimeRangeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponse) SetBody(v *DescribeDBTablesRecoveryTimeRangeResponseBody) *DescribeDBTablesRecoveryTimeRangeResponse {
	s.Body = v
	return s
}

type DescribeDownloadBackupSetStorageInfoRequest struct {
	// The ID of the backup set.
	//
	// example:
	//
	// 30****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The validity period of the URL that is used as the download destination. Take note of the following items:
	//
	// 	- Default value: 7200. This means that the URL is valid for 2 hours by default.
	//
	// 	- Valid values: 300 to 86400. Unit: seconds. This means that you can specify a validity period in the range of 5 minutes to 1 day.
	//
	// 	- Before you specify this parameter, convert the validity period to seconds. For example, if you want to set the validity period of the URL to 5 minutes, enter 300.
	//
	// example:
	//
	// 300
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	//
	// > The **BackupSetId*	- parameter is required if you specify the **InstanceName*	- parameter.
	//
	// example:
	//
	// rm-uf6qqf569n435****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](~~26231~~) operation to query the region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The ID of the download task.
	//
	// 	- The **BackupSetId*	- and **InstanceName*	- parameters are required if you do not specify the **TaskId*	- parameter.
	//
	// 	- You can go to the instance details page in the Alibaba Cloud Management Console and click **Backup and Restoration*	- in the left-side navigation pane. On the **Backup Download*	- tab, view the task ID.
	//
	// example:
	//
	// dt-s0ugzak9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDownloadBackupSetStorageInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadBackupSetStorageInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetBackupSetId(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetDuration(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.Duration = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetInstanceName(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetRegionCode(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoRequest) SetTaskId(v string) *DescribeDownloadBackupSetStorageInfoRequest {
	s.TaskId = &v
	return s
}

type DescribeDownloadBackupSetStorageInfoResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// DBS.ParamIsInValid
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeDownloadBackupSetStorageInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// DBS.ParamIsInValid
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Argument: regionCode Must not be empty
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Argument: regionCode Must not be empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 44B8C2F5-919D-5D29-BCD5-DEB03467****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDownloadBackupSetStorageInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadBackupSetStorageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetCode(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetData(v *DescribeDownloadBackupSetStorageInfoResponseBodyData) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetErrCode(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetErrMessage(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetMessage(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetRequestId(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBody) SetSuccess(v string) *DescribeDownloadBackupSetStorageInfoResponseBody {
	s.Success = &v
	return s
}

type DescribeDownloadBackupSetStorageInfoResponseBodyData struct {
	// The validity period of the URL.
	//
	// > This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1661329050
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The private download URL of the backup set.
	//
	// example:
	//
	// http://dbs-137383785969****-cn-hangzhou-1iv12nblw****.oss-cn-hangzhou-internal.aliyuncs.com/dt-u7u4bufa****/dbs_target_file_path/test_123
	PrivateUrl *string `json:"PrivateUrl,omitempty" xml:"PrivateUrl,omitempty"`
	// The public download URL of the backup set.
	//
	// example:
	//
	// http://dbs-137383785969****-cn-hangzhou-1iv12nblw****.oss-cn-hangzhou.aliyuncs.com/dt-u7u4bufa****/dbs_target_file_path/test_456
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
}

func (s DescribeDownloadBackupSetStorageInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadBackupSetStorageInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) SetExpirationTime(v int64) *DescribeDownloadBackupSetStorageInfoResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) SetPrivateUrl(v string) *DescribeDownloadBackupSetStorageInfoResponseBodyData {
	s.PrivateUrl = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponseBodyData) SetPublicUrl(v string) *DescribeDownloadBackupSetStorageInfoResponseBodyData {
	s.PublicUrl = &v
	return s
}

type DescribeDownloadBackupSetStorageInfoResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadBackupSetStorageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadBackupSetStorageInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadBackupSetStorageInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) SetHeaders(v map[string]*string) *DescribeDownloadBackupSetStorageInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) SetStatusCode(v int32) *DescribeDownloadBackupSetStorageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) SetBody(v *DescribeDownloadBackupSetStorageInfoResponseBody) *DescribeDownloadBackupSetStorageInfoResponse {
	s.Body = v
	return s
}

type DescribeDownloadSupportRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// rm-bp1a48p922r4b****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](~~26231~~) operation to query the region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s DescribeDownloadSupportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadSupportRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSupportRequest) SetInstanceName(v string) *DescribeDownloadSupportRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeDownloadSupportRequest) SetRegionCode(v string) *DescribeDownloadSupportRequest {
	s.RegionCode = &v
	return s
}

type DescribeDownloadSupportResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// DBS.ParamIsInValid
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the advanced download feature is supported. Valid values:
	//
	// 	- **true**: The advanced download feature is supported.
	//
	// 	- **false**: The advanced download feature is not supported.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// DBS.ParamIsInValid
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Argument: regionCode Must not be empty
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Argument: regionCode Must not be empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F1A186F7-7B34-5C11-A903-EE23876B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDownloadSupportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadSupportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSupportResponseBody) SetCode(v string) *DescribeDownloadSupportResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetData(v string) *DescribeDownloadSupportResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetErrCode(v string) *DescribeDownloadSupportResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetErrMessage(v string) *DescribeDownloadSupportResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetMessage(v string) *DescribeDownloadSupportResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetRequestId(v string) *DescribeDownloadSupportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadSupportResponseBody) SetSuccess(v string) *DescribeDownloadSupportResponseBody {
	s.Success = &v
	return s
}

type DescribeDownloadSupportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadSupportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadSupportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadSupportResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSupportResponse) SetHeaders(v map[string]*string) *DescribeDownloadSupportResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadSupportResponse) SetStatusCode(v int32) *DescribeDownloadSupportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadSupportResponse) SetBody(v *DescribeDownloadSupportResponseBody) *DescribeDownloadSupportResponse {
	s.Body = v
	return s
}

type DescribeDownloadTaskRequest struct {
	// The ID of the backup set generated when you create a download task. You can call the [DescribeBackups](~~26273~~) operation to query the ID.
	//
	// example:
	//
	// 216****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the Database Backup (DBS) data source. Specify the parameter in the format of *ds-${Instance ID}\_${regionId}*.
	//
	// example:
	//
	// ds-rm-2ze8g2am97624****_cn-hangzhou
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The end of the time range to query. Specify this parameter as a timestamp of the LONG type. Unit: milliseconds.
	//
	// example:
	//
	// 1661941556000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// > This parameter is required.
	//
	// example:
	//
	// rm-bp1imnmcjxdz7****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The column based on which the entries are sorted. By default, the entries are sorted by the time when the download task was created. Set the value to **gmt_create**.
	//
	// example:
	//
	// gmt_create
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The order in which you want to sort the entries. Valid values:
	//
	// 	- **asc**: the ascending order.
	//
	// 	- **desc**: the descending order. This is the default value.
	//
	// example:
	//
	// desc
	OrderDirect *string `json:"OrderDirect,omitempty" xml:"OrderDirect,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 50
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](~~26231~~) operation to query the region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The beginning of the time range to query. Specify this parameter as a timestamp of the LONG type. Unit: milliseconds.
	//
	// example:
	//
	// 1661941554000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the download task. Valid values:
	//
	// 	- **Initializing**: The download task is being initialized.
	//
	// 	- **queueing**: The download task is queuing.
	//
	// 	- **running**: The download task is running.
	//
	// 	- **failed**: The download task fails.
	//
	// 	- **finished**: The download task is complete.
	//
	// 	- **expired**: The download task expires.
	//
	// example:
	//
	// queueing
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The type of the download task. Valid values:
	//
	// 	- **full**: downloads a full backup set.
	//
	// 	- **pitr**: downloads a backup set at a specific point in time.
	//
	// example:
	//
	// full
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskRequest) SetBackupSetId(v string) *DescribeDownloadTaskRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetCurrentPage(v string) *DescribeDownloadTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetDatasourceId(v string) *DescribeDownloadTaskRequest {
	s.DatasourceId = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetEndTime(v string) *DescribeDownloadTaskRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetInstanceName(v string) *DescribeDownloadTaskRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetOrderColumn(v string) *DescribeDownloadTaskRequest {
	s.OrderColumn = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetOrderDirect(v string) *DescribeDownloadTaskRequest {
	s.OrderDirect = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetPageSize(v string) *DescribeDownloadTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetRegionCode(v string) *DescribeDownloadTaskRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetStartTime(v string) *DescribeDownloadTaskRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetState(v string) *DescribeDownloadTaskRequest {
	s.State = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetTaskType(v string) *DescribeDownloadTaskRequest {
	s.TaskType = &v
	return s
}

type DescribeDownloadTaskResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// DBS.InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the tasks.
	Data *DescribeDownloadTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// DBS.InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// instanceName can not be empty
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// instanceName can not be empty
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5D285EB9-A443-592D-9F3D-A888FAC3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDownloadTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBody) SetCode(v string) *DescribeDownloadTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetData(v *DescribeDownloadTaskResponseBodyData) *DescribeDownloadTaskResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetErrCode(v string) *DescribeDownloadTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetErrMessage(v string) *DescribeDownloadTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetMessage(v string) *DescribeDownloadTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetRequestId(v string) *DescribeDownloadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetSuccess(v string) *DescribeDownloadTaskResponseBody {
	s.Success = &v
	return s
}

type DescribeDownloadTaskResponseBodyData struct {
	// The details of the task.
	Content *DescribeDownloadTaskResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The extra description of the download tasks.
	//
	// example:
	//
	// dbtest
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The page number of the returned page. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of full backup tasks.
	//
	// example:
	//
	// 1
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 2
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeDownloadTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBodyData) SetContent(v *DescribeDownloadTaskResponseBodyDataContent) *DescribeDownloadTaskResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeDownloadTaskResponseBodyData) SetExtra(v string) *DescribeDownloadTaskResponseBodyData {
	s.Extra = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyData) SetPageNumber(v int64) *DescribeDownloadTaskResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyData) SetPageSize(v int64) *DescribeDownloadTaskResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyData) SetTotalElements(v int64) *DescribeDownloadTaskResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyData) SetTotalPages(v int64) *DescribeDownloadTaskResponseBodyData {
	s.TotalPages = &v
	return s
}

type DescribeDownloadTaskResponseBodyDataContent struct {
	List []*DescribeDownloadTaskResponseBodyDataContentList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeDownloadTaskResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBodyDataContent) SetList(v []*DescribeDownloadTaskResponseBodyDataContentList) *DescribeDownloadTaskResponseBodyDataContent {
	s.List = v
	return s
}

type DescribeDownloadTaskResponseBodyDataContentList struct {
	// The point in time of the backup set if the task is used to download a backup set at a specific point in time. The value is a timestamp of the LONG type. Unit: millisecond.
	//
	// example:
	//
	// 1663162216000
	BackupSetTime *string `json:"BackupSetTime,omitempty" xml:"BackupSetTime,omitempty"`
	// The ID of the full backup set.
	//
	// example:
	//
	// 148261****
	BakSetId *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	// The details of the databases.
	//
	// example:
	//
	// [dbtest]
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The status of the download task. Valid values:
	//
	// 	- **Initializing**: The download task is being initialized.
	//
	// 	- **queuing**: The download task is queuing.
	//
	// 	- **running**: The download task is running.
	//
	// 	- **failed**: The download task fails.
	//
	// 	- **finished**: The download task is complete.
	//
	// 	- **expired**: The download task expires.
	//
	// example:
	//
	// queueing
	DownloadStatus *string `json:"DownloadStatus,omitempty" xml:"DownloadStatus,omitempty"`
	// The amount of output data. Unit: bytes.
	//
	// example:
	//
	// 0
	ExportDataSize *string `json:"ExportDataSize,omitempty" xml:"ExportDataSize,omitempty"`
	// The format to which the downloaded backup set is converted. Valid values:
	//
	// 	- **csv**
	//
	// 	- **SQL**
	//
	// 	- **Parquet**
	//
	// example:
	//
	// csv
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The time when the download task was created. The value is a timestamp.
	//
	// example:
	//
	// 1663321957000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The amount of data that is processed. Unit: bytes.
	//
	// example:
	//
	// 0
	ImportDataSize *string `json:"ImportDataSize,omitempty" xml:"ImportDataSize,omitempty"`
	// The number of tables that have been downloaded and the total number of tables to be downloaded.
	//
	// example:
	//
	// 0/0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The destination path to which the data is downloaded if the value of **TargetType is OSS**.
	//
	// example:
	//
	// test_db/path
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the method in which the backup set is downloaded. Valid values:
	//
	// 	- **OSS**
	//
	// 	- **URL**
	//
	// example:
	//
	// URL
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The download task ID.
	//
	// example:
	//
	// dt-qxntlvgu****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDownloadTaskResponseBodyDataContentList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskResponseBodyDataContentList) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetBackupSetTime(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.BackupSetTime = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetBakSetId(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.BakSetId = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetDbList(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.DbList = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetDownloadStatus(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.DownloadStatus = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetExportDataSize(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.ExportDataSize = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetFormat(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.Format = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetGmtCreate(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetImportDataSize(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.ImportDataSize = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetProgress(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.Progress = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetRegionCode(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.RegionCode = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetTargetPath(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.TargetPath = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetTargetType(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.TargetType = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyDataContentList) SetTaskId(v string) *DescribeDownloadTaskResponseBodyDataContentList {
	s.TaskId = &v
	return s
}

type DescribeDownloadTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponse) SetHeaders(v map[string]*string) *DescribeDownloadTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadTaskResponse) SetStatusCode(v int32) *DescribeDownloadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadTaskResponse) SetBody(v *DescribeDownloadTaskResponseBody) *DescribeDownloadTaskResponse {
	s.Body = v
	return s
}

type DescribeSandboxBackupSetsRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](~~437215~~) operation to query the ID of the backup schedule.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](~~193091~~) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	//
	// example:
	//
	// 1hxxxx8xxxxxa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the backup set. If this parameter is specified, only the snapshot of the specified backup set is returned. If this parameter is not specified, all the snapshots of the backup schedule are returned.
	//
	// example:
	//
	// 1xxxx2xxxxx1e
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- 30: This is the default value.
	//
	// 	- 50\.
	//
	// 	- 100\.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSandboxBackupSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxBackupSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxBackupSetsRequest) SetBackupPlanId(v string) *DescribeSandboxBackupSetsRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) SetBackupSetId(v string) *DescribeSandboxBackupSetsRequest {
	s.BackupSetId = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) SetPageNumber(v string) *DescribeSandboxBackupSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSandboxBackupSetsRequest) SetPageSize(v string) *DescribeSandboxBackupSetsRequest {
	s.PageSize = &v
	return s
}

type DescribeSandboxBackupSetsResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data. The following parameters are contained:
	//
	// 	- **backupSetTime**: the point in time when the snapshot was created. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// 	- **backupSetId**: the ID of the backup set.
	//
	// 	- **backupSetType**: the type of the snapshot. A value of **Full*	- indicates that the snapshot is a full backup snapshot. A value of **Inc*	- indicates that the snapshot is an incremental backup snapshot.
	//
	// 	- **backupPlanId**: the ID of the backup schedule.
	//
	// example:
	//
	// "Data": {     "number": 2,     "size": 2,     "content": [       {         "backupSetTime": "2021-08-28T23:12:31Z",         "backupSetId": "Inc_1hxxxx8xxxxxa_20210801064200_mysql-bin.000134",         "backupSetType": "Inc",         "backupPlanId": "1hxxxx8xxxxxa"       },       {         "backupSetTime": "2021-08-28T22:42:28Z",         "backupSetId": "1hxxxx8xxxxxa_20210829064228",         "backupSetType": "FULL",         "backupPlanId": "1hxxxx8xxxxxa"       }     ],     "totalElements": 2   },
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSandboxBackupSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxBackupSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxBackupSetsResponseBody) SetCode(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetData(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetErrCode(v string) *DescribeSandboxBackupSetsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetErrMessage(v string) *DescribeSandboxBackupSetsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetMessage(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetRequestId(v string) *DescribeSandboxBackupSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponseBody) SetSuccess(v string) *DescribeSandboxBackupSetsResponseBody {
	s.Success = &v
	return s
}

type DescribeSandboxBackupSetsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSandboxBackupSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSandboxBackupSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxBackupSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxBackupSetsResponse) SetHeaders(v map[string]*string) *DescribeSandboxBackupSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxBackupSetsResponse) SetStatusCode(v int32) *DescribeSandboxBackupSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxBackupSetsResponse) SetBody(v *DescribeSandboxBackupSetsResponseBody) *DescribeSandboxBackupSetsResponse {
	s.Body = v
	return s
}

type DescribeSandboxInstancesRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](~~437215~~) operation to obtain the ID of the backup schedule.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](~~193091~~) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	//
	// example:
	//
	// 1hxxxx8xxxxxa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the sandbox instance. You can call the [CreateSandboxInstance](~~437252~~) operation to obtain the ID of the sandbox instance.
	//
	// example:
	//
	// 1jxxxxnxxx1xc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- 30\. This is the default value.
	//
	// 	- 50
	//
	// 	- 100
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSandboxInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxInstancesRequest) SetBackupPlanId(v string) *DescribeSandboxInstancesRequest {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) SetInstanceId(v string) *DescribeSandboxInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) SetPageNumber(v string) *DescribeSandboxInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSandboxInstancesRequest) SetPageSize(v string) *DescribeSandboxInstancesRequest {
	s.PageSize = &v
	return s
}

type DescribeSandboxInstancesResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// 	- **connectionString**: the connection string of the sandbox instance, in the format of IP address:Port number. This parameter indicates the endpoint of the sandbox instance if the value of the SandboxType parameter is **Sandbox**. This parameter indicates the Network File System (NFS) mount address if the value of the SandboxType parameter is **NFS**.
	//
	// 	- **restoreSeconds**: the time required to create the sandbox instance. Unit: seconds.
	//
	// 	- **restoreTime**: the point in time to which the sandbox instance is restored. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// 	- **instanceId**: the ID of the sandbox instance.
	//
	// 	- **backupSetId**: the ID of the backup set.
	//
	// 	- **createTime**: the point in time when the sandbox instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// 	- **backupPlanId**: the ID of the backup schedule.
	//
	// 	- **vpcId**: the ID of the virtual private cloud (VPC).
	//
	// 	- **vpcSwitchId**: the ID of the VSwitch.
	//
	// 	- **sandboxSpecification**: the specifications of the sandbox instance.
	//
	// 	- **status**: the status of the sandbox instance. Valid values: **running**, **check_pass**, and **stop**.
	//
	// example:
	//
	// {     "number": 1,     "size": 1,     "content": [       {         "connectionString": "172.26.178.229:3306",         "restoreSeconds": 15,         "restoreTime": "2021-08-11T07:26:24Z",         "instanceId": "1jxxxxx9xxxms",         "backupSetId": "1hxxxx8xxxxxa_20210811152624",         "createTime": "2021-08-12T07:40:29Z",         "backupPlanId": "1hxxxx8xxxxxa",         "vpcId": "vpc-bp1dxxxxxjy0xxxxx1xxp",         "sandboxSpecification": "MYSQL_1C_1M_SD",         "status": "running",         "vpcSwitchId": "vsw-bp1bxxxxxumxxxxxwxx2w"       }     ],     "totalElements": 1   }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSandboxInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxInstancesResponseBody) SetCode(v string) *DescribeSandboxInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetData(v string) *DescribeSandboxInstancesResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetErrCode(v string) *DescribeSandboxInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetErrMessage(v string) *DescribeSandboxInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetMessage(v string) *DescribeSandboxInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetRequestId(v string) *DescribeSandboxInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetSuccess(v string) *DescribeSandboxInstancesResponseBody {
	s.Success = &v
	return s
}

type DescribeSandboxInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSandboxInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSandboxInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxInstancesResponse) SetHeaders(v map[string]*string) *DescribeSandboxInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxInstancesResponse) SetStatusCode(v int32) *DescribeSandboxInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxInstancesResponse) SetBody(v *DescribeSandboxInstancesResponseBody) *DescribeSandboxInstancesResponse {
	s.Body = v
	return s
}

type DescribeSandboxRecoveryTimeRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](~~437215~~) operation to obtain the ID of the backup schedule. If you set this parameter to the backup schedule ID obtained by calling the DescribeBackupPlanList operation, the dbs prefix must be removed. Otherwise, the call fails.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](~~193091~~) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	//
	// example:
	//
	// 1jyjal15l****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
}

func (s DescribeSandboxRecoveryTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeRequest) SetBackupPlanId(v string) *DescribeSandboxRecoveryTimeRequest {
	s.BackupPlanId = &v
	return s
}

type DescribeSandboxRecoveryTimeResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribeSandboxRecoveryTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSandboxRecoveryTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetCode(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetData(v *DescribeSandboxRecoveryTimeResponseBodyData) *DescribeSandboxRecoveryTimeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetErrCode(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetErrMessage(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetMessage(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetRequestId(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBody) SetSuccess(v string) *DescribeSandboxRecoveryTimeResponseBody {
	s.Success = &v
	return s
}

type DescribeSandboxRecoveryTimeResponseBodyData struct {
	// The backup schedule of the sandbox instance.
	//
	// example:
	//
	// 1hxxxx8xxxxxa
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The beginning of the time range during which the sandbox instance can be restored. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-08-01T12:01:01Z
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" xml:"RecoveryBeginTime,omitempty"`
	// The end of the time range during which the sandbox instance can be restored. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-08-02T12:01:01Z
	RecoveryEndTime *string `json:"RecoveryEndTime,omitempty" xml:"RecoveryEndTime,omitempty"`
}

func (s DescribeSandboxRecoveryTimeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) SetBackupPlanId(v string) *DescribeSandboxRecoveryTimeResponseBodyData {
	s.BackupPlanId = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) SetRecoveryBeginTime(v string) *DescribeSandboxRecoveryTimeResponseBodyData {
	s.RecoveryBeginTime = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponseBodyData) SetRecoveryEndTime(v string) *DescribeSandboxRecoveryTimeResponseBodyData {
	s.RecoveryEndTime = &v
	return s
}

type DescribeSandboxRecoveryTimeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSandboxRecoveryTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSandboxRecoveryTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeResponse) SetHeaders(v map[string]*string) *DescribeSandboxRecoveryTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponse) SetStatusCode(v int32) *DescribeSandboxRecoveryTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponse) SetBody(v *DescribeSandboxRecoveryTimeResponseBody) *DescribeSandboxRecoveryTimeResponse {
	s.Body = v
	return s
}

type ModifyDBTablesRecoveryStateRequest struct {
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	Retention  *string `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s ModifyDBTablesRecoveryStateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBTablesRecoveryStateRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBTablesRecoveryStateRequest) SetCategory(v string) *ModifyDBTablesRecoveryStateRequest {
	s.Category = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateRequest) SetInstanceId(v string) *ModifyDBTablesRecoveryStateRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateRequest) SetRegionCode(v string) *ModifyDBTablesRecoveryStateRequest {
	s.RegionCode = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateRequest) SetRetention(v string) *ModifyDBTablesRecoveryStateRequest {
	s.Retention = &v
	return s
}

type ModifyDBTablesRecoveryStateResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDBTablesRecoveryStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBTablesRecoveryStateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetCode(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetData(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetErrCode(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetErrMessage(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetMessage(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetRequestId(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetSuccess(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.Success = &v
	return s
}

type ModifyDBTablesRecoveryStateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBTablesRecoveryStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBTablesRecoveryStateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBTablesRecoveryStateResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBTablesRecoveryStateResponse) SetHeaders(v map[string]*string) *ModifyDBTablesRecoveryStateResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponse) SetStatusCode(v int32) *ModifyDBTablesRecoveryStateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponse) SetBody(v *ModifyDBTablesRecoveryStateResponseBody) *ModifyDBTablesRecoveryStateResponse {
	s.Body = v
	return s
}

type SupportDBTableRecoveryRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s SupportDBTableRecoveryRequest) String() string {
	return tea.Prettify(s)
}

func (s SupportDBTableRecoveryRequest) GoString() string {
	return s.String()
}

func (s *SupportDBTableRecoveryRequest) SetInstanceId(v string) *SupportDBTableRecoveryRequest {
	s.InstanceId = &v
	return s
}

func (s *SupportDBTableRecoveryRequest) SetRegionCode(v string) *SupportDBTableRecoveryRequest {
	s.RegionCode = &v
	return s
}

type SupportDBTableRecoveryResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SupportDBTableRecoveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SupportDBTableRecoveryResponseBody) GoString() string {
	return s.String()
}

func (s *SupportDBTableRecoveryResponseBody) SetCode(v string) *SupportDBTableRecoveryResponseBody {
	s.Code = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetData(v string) *SupportDBTableRecoveryResponseBody {
	s.Data = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetErrCode(v string) *SupportDBTableRecoveryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetErrMessage(v string) *SupportDBTableRecoveryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetMessage(v string) *SupportDBTableRecoveryResponseBody {
	s.Message = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetRequestId(v string) *SupportDBTableRecoveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetSuccess(v string) *SupportDBTableRecoveryResponseBody {
	s.Success = &v
	return s
}

type SupportDBTableRecoveryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SupportDBTableRecoveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SupportDBTableRecoveryResponse) String() string {
	return tea.Prettify(s)
}

func (s SupportDBTableRecoveryResponse) GoString() string {
	return s.String()
}

func (s *SupportDBTableRecoveryResponse) SetHeaders(v map[string]*string) *SupportDBTableRecoveryResponse {
	s.Headers = v
	return s
}

func (s *SupportDBTableRecoveryResponse) SetStatusCode(v int32) *SupportDBTableRecoveryResponse {
	s.StatusCode = &v
	return s
}

func (s *SupportDBTableRecoveryResponse) SetBody(v *SupportDBTableRecoveryResponseBody) *SupportDBTableRecoveryResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":            tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-chengdu":            tea.String("dbs-api.cn-chengdu.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-huhehaote":          tea.String("dbs-api.cn-huhehaote.aliyuncs.com"),
		"cn-hangzhou":           tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen":           tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"ap-southeast-1":        tea.String("dbs-api.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("dbs-api.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        tea.String("dbs-api.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        tea.String("dbs-api.ap-southeast-5.aliyuncs.com"),
		"ap-northeast-1":        tea.String("dbs-api.ap-northeast-1.aliyuncs.com"),
		"us-west-1":             tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"us-east-1":             tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"eu-central-1":          tea.String("dbs-api.eu-central-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("dbs-api.cn-hangzhou.aliyuncs.com"),
		"ap-south-1":            tea.String("dbs-api.ap-south-1.aliyuncs.com"),
		"eu-west-1":             tea.String("dbs-api.eu-west-1.aliyuncs.com"),
		"me-east-1":             tea.String("dbs-api.me-east-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dbs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// ### [](#)Supported database engines
//
// 	- ApsaraDB RDS for MySQL
//
// 	- ApsaraDB RDS for PostgreSQL
//
// 	- PolarDB for MySQL
//
// ### [](#)References
//
// For the instances that meet your business requirements, you can create an advanced download task by point in time or backup set. You can set the download destination to a URL or directly upload the downloaded backup set to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving.
//
// 	- [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~)
//
// 	- [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](~~96774~~)
//
// 	- [Download the backup files of a PolarDB for MySQL cluster](~~2627635~~)
//
// @param request - CreateDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDownloadResponse
func (client *Client) CreateDownloadWithOptions(request *CreateDownloadRequest, runtime *util.RuntimeOptions) (_result *CreateDownloadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BakSetId)) {
		query["BakSetId"] = request.BakSetId
	}

	if !tea.BoolValue(util.IsUnset(request.BakSetSize)) {
		query["BakSetSize"] = request.BakSetSize
	}

	if !tea.BoolValue(util.IsUnset(request.BakSetType)) {
		query["BakSetType"] = request.BakSetType
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadPointInTime)) {
		query["DownloadPointInTime"] = request.DownloadPointInTime
	}

	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		query["FormatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	if !tea.BoolValue(util.IsUnset(request.TargetBucket)) {
		query["TargetBucket"] = request.TargetBucket
	}

	if !tea.BoolValue(util.IsUnset(request.TargetOssRegion)) {
		query["TargetOssRegion"] = request.TargetOssRegion
	}

	if !tea.BoolValue(util.IsUnset(request.TargetPath)) {
		query["TargetPath"] = request.TargetPath
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDownload"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDownloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// ### [](#)Supported database engines
//
// 	- ApsaraDB RDS for MySQL
//
// 	- ApsaraDB RDS for PostgreSQL
//
// 	- PolarDB for MySQL
//
// ### [](#)References
//
// For the instances that meet your business requirements, you can create an advanced download task by point in time or backup set. You can set the download destination to a URL or directly upload the downloaded backup set to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving.
//
// 	- [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~)
//
// 	- [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](~~96774~~)
//
// 	- [Download the backup files of a PolarDB for MySQL cluster](~~2627635~~)
//
// @param request - CreateDownloadRequest
//
// @return CreateDownloadResponse
func (client *Client) CreateDownload(request *CreateDownloadRequest) (_result *CreateDownloadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDownloadResponse{}
	_body, _err := client.CreateDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DeleteSandboxInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSandboxInstanceResponse
func (client *Client) DeleteSandboxInstanceWithOptions(request *DeleteSandboxInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteSandboxInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPlanId)) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSandboxInstance"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSandboxInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DeleteSandboxInstanceRequest
//
// @return DeleteSandboxInstanceResponse
func (client *Client) DeleteSandboxInstance(request *DeleteSandboxInstanceRequest) (_result *DeleteSandboxInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSandboxInstanceResponse{}
	_body, _err := client.DeleteSandboxInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupDataListWithOptions(request *DescribeBackupDataListRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupDataListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMethod)) {
		query["BackupMethod"] = request.BackupMethod
	}

	if !tea.BoolValue(util.IsUnset(request.BackupMode)) {
		query["BackupMode"] = request.BackupMode
	}

	if !tea.BoolValue(util.IsUnset(request.BackupScale)) {
		query["BackupScale"] = request.BackupScale
	}

	if !tea.BoolValue(util.IsUnset(request.BackupStatus)) {
		query["BackupStatus"] = request.BackupStatus
	}

	if !tea.BoolValue(util.IsUnset(request.BackupType)) {
		query["BackupType"] = request.BackupType
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIsDeleted)) {
		query["InstanceIsDeleted"] = request.InstanceIsDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegion)) {
		query["InstanceRegion"] = request.InstanceRegion
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		query["SceneType"] = request.SceneType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupDataList"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupDataListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupDataList(request *DescribeBackupDataListRequest) (_result *DescribeBackupDataListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupDataListResponse{}
	_body, _err := client.DescribeBackupDataListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBackupPolicy"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBTablesRecoveryBackupSetWithOptions(request *DescribeDBTablesRecoveryBackupSetRequest, runtime *util.RuntimeOptions) (_result *DescribeDBTablesRecoveryBackupSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBTablesRecoveryBackupSet"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBTablesRecoveryBackupSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBTablesRecoveryBackupSet(request *DescribeDBTablesRecoveryBackupSetRequest) (_result *DescribeDBTablesRecoveryBackupSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBTablesRecoveryBackupSetResponse{}
	_body, _err := client.DescribeDBTablesRecoveryBackupSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBTablesRecoveryStateWithOptions(request *DescribeDBTablesRecoveryStateRequest, runtime *util.RuntimeOptions) (_result *DescribeDBTablesRecoveryStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBTablesRecoveryState"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBTablesRecoveryStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBTablesRecoveryState(request *DescribeDBTablesRecoveryStateRequest) (_result *DescribeDBTablesRecoveryStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBTablesRecoveryStateResponse{}
	_body, _err := client.DescribeDBTablesRecoveryStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBTablesRecoveryTimeRangeWithOptions(request *DescribeDBTablesRecoveryTimeRangeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBTablesRecoveryTimeRangeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDBTablesRecoveryTimeRange"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDBTablesRecoveryTimeRangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBTablesRecoveryTimeRange(request *DescribeDBTablesRecoveryTimeRangeRequest) (_result *DescribeDBTablesRecoveryTimeRangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBTablesRecoveryTimeRangeResponse{}
	_body, _err := client.DescribeDBTablesRecoveryTimeRangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// ### [](#)Supported database engines
//
// 	- ApsaraDB RDS for MySQL
//
// 	- ApsaraDB RDS for PostgreSQL
//
// 	- PolarDB for MySQL
//
// ### [](#)References
//
// 	- [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~)
//
// 	- [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](~~96774~~)
//
// 	- [Download the backup files of a PolarDB for MySQL cluster](~~2627635~~)
//
// @param request - DescribeDownloadBackupSetStorageInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadBackupSetStorageInfoResponse
func (client *Client) DescribeDownloadBackupSetStorageInfoWithOptions(request *DescribeDownloadBackupSetStorageInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadBackupSetStorageInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupSetId)) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadBackupSetStorageInfo"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadBackupSetStorageInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// ### [](#)Supported database engines
//
// 	- ApsaraDB RDS for MySQL
//
// 	- ApsaraDB RDS for PostgreSQL
//
// 	- PolarDB for MySQL
//
// ### [](#)References
//
// 	- [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~)
//
// 	- [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](~~96774~~)
//
// 	- [Download the backup files of a PolarDB for MySQL cluster](~~2627635~~)
//
// @param request - DescribeDownloadBackupSetStorageInfoRequest
//
// @return DescribeDownloadBackupSetStorageInfoResponse
func (client *Client) DescribeDownloadBackupSetStorageInfo(request *DescribeDownloadBackupSetStorageInfoRequest) (_result *DescribeDownloadBackupSetStorageInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadBackupSetStorageInfoResponse{}
	_body, _err := client.DescribeDownloadBackupSetStorageInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// ### [](#)Supported database engines
//
// 	- ApsaraDB RDS for MySQL
//
// 	- ApsaraDB RDS for PostgreSQL
//
// 	- PolarDB for MySQL
//
// ### [](#)References
//
// You can create an advanced download task by point in time or backup set. You can set the download destination to a URL or directly upload the downloaded backup set to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving.
//
// 	- [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~)
//
// 	- [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](~~96774~~)
//
// 	- [Download the backup files of a PolarDB for MySQL cluster](~~2627635~~)
//
// @param request - DescribeDownloadSupportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadSupportResponse
func (client *Client) DescribeDownloadSupportWithOptions(request *DescribeDownloadSupportRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadSupportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadSupport"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadSupportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// ### [](#)Supported database engines
//
// 	- ApsaraDB RDS for MySQL
//
// 	- ApsaraDB RDS for PostgreSQL
//
// 	- PolarDB for MySQL
//
// ### [](#)References
//
// You can create an advanced download task by point in time or backup set. You can set the download destination to a URL or directly upload the downloaded backup set to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving.
//
// 	- [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~)
//
// 	- [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](~~96774~~)
//
// 	- [Download the backup files of a PolarDB for MySQL cluster](~~2627635~~)
//
// @param request - DescribeDownloadSupportRequest
//
// @return DescribeDownloadSupportResponse
func (client *Client) DescribeDownloadSupport(request *DescribeDownloadSupportRequest) (_result *DescribeDownloadSupportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadSupportResponse{}
	_body, _err := client.DescribeDownloadSupportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// ### [](#)Supported database engines
//
// 	- ApsaraDB RDS for MySQL
//
// 	- ApsaraDB RDS for PostgreSQL
//
// 	- PolarDB for MySQL
//
// ### [](#)References
//
// 	- [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~)
//
// 	- [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](~~96774~~)
//
// 	- [Download the backup files of a PolarDB for MySQL cluster](~~2627635~~)
//
// @param request - DescribeDownloadTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDownloadTaskResponse
func (client *Client) DescribeDownloadTaskWithOptions(request *DescribeDownloadTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupSetId)) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DatasourceId)) {
		query["DatasourceId"] = request.DatasourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.OrderColumn)) {
		query["OrderColumn"] = request.OrderColumn
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDirect)) {
		query["OrderDirect"] = request.OrderDirect
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadTask"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// ### [](#)Supported database engines
//
// 	- ApsaraDB RDS for MySQL
//
// 	- ApsaraDB RDS for PostgreSQL
//
// 	- PolarDB for MySQL
//
// ### [](#)References
//
// 	- [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~)
//
// 	- [Download the backup files of an ApsaraDB RDS for PostgreSQL instance](~~96774~~)
//
// 	- [Download the backup files of a PolarDB for MySQL cluster](~~2627635~~)
//
// @param request - DescribeDownloadTaskRequest
//
// @return DescribeDownloadTaskResponse
func (client *Client) DescribeDownloadTask(request *DescribeDownloadTaskRequest) (_result *DescribeDownloadTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadTaskResponse{}
	_body, _err := client.DescribeDownloadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DescribeSandboxBackupSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSandboxBackupSetsResponse
func (client *Client) DescribeSandboxBackupSetsWithOptions(request *DescribeSandboxBackupSetsRequest, runtime *util.RuntimeOptions) (_result *DescribeSandboxBackupSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPlanId)) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupSetId)) {
		query["BackupSetId"] = request.BackupSetId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSandboxBackupSets"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSandboxBackupSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DescribeSandboxBackupSetsRequest
//
// @return DescribeSandboxBackupSetsResponse
func (client *Client) DescribeSandboxBackupSets(request *DescribeSandboxBackupSetsRequest) (_result *DescribeSandboxBackupSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSandboxBackupSetsResponse{}
	_body, _err := client.DescribeSandboxBackupSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DescribeSandboxInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSandboxInstancesResponse
func (client *Client) DescribeSandboxInstancesWithOptions(request *DescribeSandboxInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeSandboxInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPlanId)) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSandboxInstances"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSandboxInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DescribeSandboxInstancesRequest
//
// @return DescribeSandboxInstancesResponse
func (client *Client) DescribeSandboxInstances(request *DescribeSandboxInstancesRequest) (_result *DescribeSandboxInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSandboxInstancesResponse{}
	_body, _err := client.DescribeSandboxInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DescribeSandboxRecoveryTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSandboxRecoveryTimeResponse
func (client *Client) DescribeSandboxRecoveryTimeWithOptions(request *DescribeSandboxRecoveryTimeRequest, runtime *util.RuntimeOptions) (_result *DescribeSandboxRecoveryTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupPlanId)) {
		query["BackupPlanId"] = request.BackupPlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSandboxRecoveryTime"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSandboxRecoveryTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
//
// @param request - DescribeSandboxRecoveryTimeRequest
//
// @return DescribeSandboxRecoveryTimeResponse
func (client *Client) DescribeSandboxRecoveryTime(request *DescribeSandboxRecoveryTimeRequest) (_result *DescribeSandboxRecoveryTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSandboxRecoveryTimeResponse{}
	_body, _err := client.DescribeSandboxRecoveryTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBTablesRecoveryStateWithOptions(request *ModifyDBTablesRecoveryStateRequest, runtime *util.RuntimeOptions) (_result *ModifyDBTablesRecoveryStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	if !tea.BoolValue(util.IsUnset(request.Retention)) {
		query["Retention"] = request.Retention
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDBTablesRecoveryState"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDBTablesRecoveryStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBTablesRecoveryState(request *ModifyDBTablesRecoveryStateRequest) (_result *ModifyDBTablesRecoveryStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBTablesRecoveryStateResponse{}
	_body, _err := client.ModifyDBTablesRecoveryStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SupportDBTableRecoveryWithOptions(request *SupportDBTableRecoveryRequest, runtime *util.RuntimeOptions) (_result *SupportDBTableRecoveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionCode)) {
		query["RegionCode"] = request.RegionCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SupportDBTableRecovery"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SupportDBTableRecoveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SupportDBTableRecovery(request *SupportDBTableRecoveryRequest) (_result *SupportDBTableRecoveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SupportDBTableRecoveryResponse{}
	_body, _err := client.SupportDBTableRecoveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
