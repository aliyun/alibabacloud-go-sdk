// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the resource group to which you want to move the resource.
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Set the value to backupplan.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the resource was successfully moved. Valid values:
	//
	// *   **true**: The resource was successfully moved.
	// *   **false**: The resource failed to be moved.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The additional information.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BakSetId *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	// The size of the full backup set. Unit: bytes. You can call the [DescribeBackups](~~26273~~) operation to query the size of the full backup set.
	BakSetSize *string `json:"BakSetSize,omitempty" xml:"BakSetSize,omitempty"`
	// The type of the download task. Valid values:
	//
	// *   **full**: downloads a full backup set.
	// *   **pitr**: downloads a backup set at a specific point in time.
	BakSetType *string `json:"BakSetType,omitempty" xml:"BakSetType,omitempty"`
	// The point in time at which the backup set is downloaded. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > This parameter is required if the BakSetType parameter is set to pitr.
	DownloadPointInTime *string `json:"DownloadPointInTime,omitempty" xml:"DownloadPointInTime,omitempty"`
	// The format to which the downloaded backup set is converted. Valid values:
	//
	// *   **CSV**
	// *   **SQL**
	// *   **Parquet**
	//
	// > This parameter is required.
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The ID of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](~~26231~~) operation to query the region ID of the instance.
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The name of the OSS bucket that is used to store the backup set.
	//
	// *   This parameter is required if the TargetType parameter is set to OSS.
	// *   Make sure that your account is granted the **AliyunDBSDefaultRole** permission. For more information, see [Use RAM for resource authorization](~~26307~~). You can also grant permissions based on the operation instructions in the Resource Access Management (RAM) console.
	TargetBucket *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	// The region in which the OSS bucket resides.
	//
	// > This parameter is required if the TargetType parameter is set to OSS.
	TargetOssRegion *string `json:"TargetOssRegion,omitempty" xml:"TargetOssRegion,omitempty"`
	// The destination path to which the backup set is downloaded.
	//
	// > This parameter is required if the TargetType parameter is set to OSS.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the destination to which the backup set is downloaded. Valid values:
	//
	// *   **OSS**
	// *   **URL**
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateDownloadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
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
	BackupSetTime *int64 `json:"BackupSetTime,omitempty" xml:"BackupSetTime,omitempty"`
	// The ID of the full backup set.
	BakSetId *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	// The database and table information that is returned if databases and tables are filtered by the download task.
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The state of the download task. Valid values:
	//
	// *   initializing: The download task was being initialized.
	// *   queuing: The download task was queuing.
	// *   running: The download task was running.
	// *   failed: The download task failed.
	// *   finished: The download task was complete.
	// *   expired: The download task expired.
	//
	// > If the TargetType parameter is set to URL, the download task expires in three days after the task is complete.
	DownloadStatus *string `json:"DownloadStatus,omitempty" xml:"DownloadStatus,omitempty"`
	// The size of the downloaded data. Unit: bytes.
	ExportDataSize *int64 `json:"ExportDataSize,omitempty" xml:"ExportDataSize,omitempty"`
	// The format to which the downloaded data is converted.
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The time when the download task was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The size of the processed data. Unit: bytes.
	ImportDataSize *int64 `json:"ImportDataSize,omitempty" xml:"ImportDataSize,omitempty"`
	// The number of tables that have been downloaded and the total number of tables to be downloaded.
	//
	// > If the task is in the preparation stage, 0/0 is returned.
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the region in which the instance resides.
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The destination path to which the backup set is downloaded.
	//
	// > This parameter is returned if the TargetType parameter is set to OSS.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the destination to which the backup set is downloaded.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the download task.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateSandboxInstanceRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](~~437215~~) operation to obtain the ID of the backup schedule.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](~~193091~~) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the backup set to be restored, which is the point in time when a snapshot was created. You can call the [DescribeSandboxBackupSets](~~437256~~) operation to obtain the ID.
	//
	// > You need to specify only one of the **BackupSetId** and **RestoreTime** parameters.
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The point in time of the sandbox instance to be restored. You can call the [DescribeSandboxRecoveryTime](~~437258~~) operation to view the recoverable time range. Specify the time in the format of *yyyy-MM-ddTHH:mm:ssZ*. The time must be in UTC.
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The custom name of the sandbox instance.
	SandboxInstanceName *string `json:"SandboxInstanceName,omitempty" xml:"SandboxInstanceName,omitempty"`
	// The password of the privileged account created in the sandbox instance.
	SandboxPassword *string `json:"SandboxPassword,omitempty" xml:"SandboxPassword,omitempty"`
	// The specifications of the sandbox instance. Valid values:
	//
	// *   **MYSQL\_1C\_1M_SD**: 1 CPU core and 1 GB of memory.
	// *   **MYSQL\_1C\_2M_SD**: 1 CPU core and 2 GB of memory.
	// *   **MYSQL\_2C\_4M_SD**: 2 CPU cores and 4 GB of memory.
	// *   **MYSQL\_2C\_8M_SD**: 2 CPU cores and 8 GB of memory.
	// *   **MYSQL\_4C\_8M_SD**: 4 CPU cores and 8 GB of memory.
	// *   **MYSQL\_4C\_16M_SD**: 4 CPU cores and 16 GB of memory.
	// *   **MYSQL\_8C\_16M_SD**: 8 CPU cores and 16 GB of memory.
	// *   **MYSQL\_8C\_32M_SD**: 8 CPU cores and 32 GB of memory.
	//
	// > Different specifications have little impact on the recovery speed. High-specification instances provide better performance after restoration. For more information, see [DBS sandbox fees](~~201466~~).
	SandboxSpecification *string `json:"SandboxSpecification,omitempty" xml:"SandboxSpecification,omitempty"`
	// The type of the sandbox instance. You can call this operation only to create an instance of the **Sandbox** type. After the sandbox instance is created, the MySQL endpoint of the instance is provided.
	SandboxType *string `json:"SandboxType,omitempty" xml:"SandboxType,omitempty"`
	// The privileged account created in the sandbox instance.
	//
	// *   After you specify this parameter, the system creates a privileged account in the sandbox instance. The account is granted the permissions on all databases in the instance.
	//
	// The account of the source database is retained in the sandbox instance.
	//
	// *   If you do not specify this parameter, the database account is the same as that of the source database.
	SandboxUser *string `json:"SandboxUser,omitempty" xml:"SandboxUser,omitempty"`
	// The ID of the virtual private cloud (VPC) that is used to connect to the sandbox instance. If you want to connect to the sandbox instance by using Elastic Compute Service (ECS) instances, you must set this parameter to the VPC in which the ECS instances reside.
	//
	// > You can set this parameter if you want to use it in a recovery drill scenario.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the VSwitch that is used to connect to the sandbox instance.
	VpcSwitchId *string `json:"VpcSwitchId,omitempty" xml:"VpcSwitchId,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateSandboxInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSandboxInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSandboxInstanceRequest) SetBackupPlanId(v string) *CreateSandboxInstanceRequest {
	s.BackupPlanId = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetBackupSetId(v string) *CreateSandboxInstanceRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetRestoreTime(v string) *CreateSandboxInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxInstanceName(v string) *CreateSandboxInstanceRequest {
	s.SandboxInstanceName = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxPassword(v string) *CreateSandboxInstanceRequest {
	s.SandboxPassword = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxSpecification(v string) *CreateSandboxInstanceRequest {
	s.SandboxSpecification = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxType(v string) *CreateSandboxInstanceRequest {
	s.SandboxType = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetSandboxUser(v string) *CreateSandboxInstanceRequest {
	s.SandboxUser = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetVpcId(v string) *CreateSandboxInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetVpcSwitchId(v string) *CreateSandboxInstanceRequest {
	s.VpcSwitchId = &v
	return s
}

func (s *CreateSandboxInstanceRequest) SetZoneId(v string) *CreateSandboxInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateSandboxInstanceResponseBody struct {
	// The error code returned if the request fails.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *CreateSandboxInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSandboxInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSandboxInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSandboxInstanceResponseBody) SetCode(v string) *CreateSandboxInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetData(v *CreateSandboxInstanceResponseBodyData) *CreateSandboxInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetErrCode(v string) *CreateSandboxInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetErrMessage(v string) *CreateSandboxInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetMessage(v string) *CreateSandboxInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetRequestId(v string) *CreateSandboxInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSandboxInstanceResponseBody) SetSuccess(v string) *CreateSandboxInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateSandboxInstanceResponseBodyData struct {
	// The ID of the backup plan.
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the sandbox instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateSandboxInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSandboxInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSandboxInstanceResponseBodyData) SetBackupPlanId(v string) *CreateSandboxInstanceResponseBodyData {
	s.BackupPlanId = &v
	return s
}

func (s *CreateSandboxInstanceResponseBodyData) SetInstanceId(v string) *CreateSandboxInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

type CreateSandboxInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSandboxInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSandboxInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSandboxInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateSandboxInstanceResponse) SetHeaders(v map[string]*string) *CreateSandboxInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateSandboxInstanceResponse) SetStatusCode(v int32) *CreateSandboxInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSandboxInstanceResponse) SetBody(v *CreateSandboxInstanceResponseBody) *CreateSandboxInstanceResponse {
	s.Body = v
	return s
}

type DeleteSandboxInstanceRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](~~437215~~) operation to query the ID of the backup schedule.
	//
	// > If your instance is an ApsaraDB RDS for MySQL instance, you can [configure automatic access to a data source](~~193091~~) to automatically add the instance to DBS and obtain the ID of the backup schedule.
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the sandbox instance. You can call the [DescribeSandboxInstances](~~437257~~) operation to query the ID of the sandbox instance.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSandboxInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBTablesRecoveryBackupSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBTablesRecoveryStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDBTablesRecoveryTimeRangeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The validity period of the URL that is used as the download destination. Take note of the following items:
	//
	// *   Default value: 7200. This means that the URL is valid for 2 hours by default.
	// *   Valid values: 300 to 86400. Unit: seconds. This means that you can specify a validity period in the range of 5 minutes to 1 day.
	// *   Before you specify this parameter, convert the validity period to seconds. For example, if you want to set the validity period of the URL to 5 minutes, enter 300.
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	//
	// > The **BackupSetId** parameter is required if you specify the **InstanceName** parameter.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](~~26231~~) operation to query the region ID of the instance.
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The ID of the download task.
	//
	// *   The **BackupSetId** and **InstanceName** parameters are required if you do not specify the **TaskId** parameter.
	// *   You can go to the instance details page in the Alibaba Cloud Management Console and click **Backup and Restoration** in the left-side navigation pane. On the **Backup Download** tab, view the task ID.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeDownloadBackupSetStorageInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
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
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The private download URL of the backup set.
	PrivateUrl *string `json:"PrivateUrl,omitempty" xml:"PrivateUrl,omitempty"`
	// The public download URL of the backup set.
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDownloadBackupSetStorageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](~~26231~~) operation to query the region ID of the instance.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the advanced download feature is supported. Valid values:
	//
	// *   **true**: The advanced download feature is supported.
	// *   **false**: The advanced download feature is not supported.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDownloadSupportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The page number of the page to return.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the Database Backup (DBS) data source. Specify the parameter in the format of *ds-${Instance ID}\_${regionId}*.
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The end of the time range to query. Specify this parameter as a timestamp of the LONG type. Unit: milliseconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// > This parameter is required.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The column based on which the entries are sorted. By default, the entries are sorted by the time when the download task was created. Set the value to **gmt_create**.
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The order in which you want to sort the entries. Valid values:
	//
	// *   **asc**: the ascending order.
	// *   **desc**: the descending order. This is the default value.
	OrderDirect *string `json:"OrderDirect,omitempty" xml:"OrderDirect,omitempty"`
	// The number of entries to return on each page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](~~26231~~) operation to query the region ID of the instance.
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The beginning of the time range to query. Specify this parameter as a timestamp of the LONG type. Unit: milliseconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the download task. Valid values:
	//
	// *   **Initializing**: The download task is being initialized.
	// *   **queuing**: The download task is queuing.
	// *   **running**: The download task is running.
	// *   **failed**: The download task fails.
	// *   **finished**: The download task is complete.
	// *   **expired**: The download task expires.
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The type of the download task. Valid values:
	//
	// *   **full**: downloads a full backup set.
	// *   **pitr**: downloads a backup set at a specific point in time.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the download task.
	Data *DescribeDownloadTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
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
	// The details of the download task.
	Content *DescribeDownloadTaskResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The extra description of the download tasks.
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The page number of the returned page. The value must be an integer that is greater than 0. Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of full backup tasks.
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of returned pages.
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
	// The point in time of the backup set if the task is used to download a backup set at a specific point in time. The value is a timestamp of the LONG type. Unit: milliseconds.
	BackupSetTime *string `json:"BackupSetTime,omitempty" xml:"BackupSetTime,omitempty"`
	// The ID of the full backup set.
	BakSetId *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	// The databases.
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The state of the download task. Valid values:
	//
	// *   **Initializing**: The download task was being initialized.
	// *   **queuing**: The download task was queuing.
	// *   **running**: The download task was running.
	// *   **failed**: The download task failed.
	// *   **finished**: The download task was complete.
	// *   **expired**: The download task expired.
	DownloadStatus *string `json:"DownloadStatus,omitempty" xml:"DownloadStatus,omitempty"`
	// The amount of output data. Unit: bytes.
	ExportDataSize *string `json:"ExportDataSize,omitempty" xml:"ExportDataSize,omitempty"`
	// The format to which the downloaded backup set is converted. Valid values:
	//
	// *   **csv**
	// *   **SQL**
	// *   **Parquet**
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The time when the download task was created. The value is a timestamp.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The amount of data that is processed. Unit: bytes.
	ImportDataSize *string `json:"ImportDataSize,omitempty" xml:"ImportDataSize,omitempty"`
	// The number of tables that have been downloaded and the total number of tables to be downloaded.
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the region in which the instance resides.
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The destination path to which the data is downloaded if the TargeType parameter is set to OSS.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the method in which the backup set is downloaded. Valid values:
	//
	// *   **OSS**
	// *   **URL**
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The ID of the download task.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the backup set. If this parameter is specified, only the snapshot of the specified backup set is returned. If this parameter is not specified, all the snapshots of the backup schedule are returned.
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   30: This is the default value.
	// *   50\.
	// *   100\.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data. The following parameters are contained:
	//
	// *   **backupSetTime**: the point in time when the snapshot was created. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	// *   **backupSetId**: the ID of the backup set.
	// *   **backupSetType**: the type of the snapshot. A value of **Full** indicates that the snapshot is a full backup snapshot. A value of **Inc** indicates that the snapshot is an incremental backup snapshot.
	// *   **backupPlanId**: the ID of the backup schedule.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSandboxBackupSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the sandbox instance. You can call the [CreateSandboxInstance](~~437252~~) operation to obtain the ID of the sandbox instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// *   30\. This is the default value.
	// *   50
	// *   100
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// *   **connectionString**: the connection string of the sandbox instance, in the format of IP address:Port number. This parameter indicates the endpoint of the sandbox instance if the value of the SandboxType parameter is **Sandbox**. This parameter indicates the Network File System (NFS) mount address if the value of the SandboxType parameter is **NFS**.
	// *   **restoreSeconds**: the time required to create the sandbox instance. Unit: seconds.
	// *   **restoreTime**: the point in time to which the sandbox instance is restored. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	// *   **instanceId**: the ID of the sandbox instance.
	// *   **backupSetId**: the ID of the backup set.
	// *   **createTime**: the point in time when the sandbox instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	// *   **backupPlanId**: the ID of the backup schedule.
	// *   **vpcId**: the ID of the virtual private cloud (VPC).
	// *   **vpcSwitchId**: the ID of the VSwitch.
	// *   **sandboxSpecification**: the specifications of the sandbox instance.
	// *   **status**: the status of the sandbox instance. Valid values: **running**, **check_pass**, and **stop**.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request fails.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSandboxInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribeSandboxRecoveryTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The beginning of the time range during which the sandbox instance can be restored. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" xml:"RecoveryBeginTime,omitempty"`
	// The end of the time range during which the sandbox instance can be restored. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSandboxRecoveryTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDBTablesRecoveryStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SupportDBTableRecoveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

/**
 * For ApsaraDB RDS for MySQL instances that use standard SSDs or enhanced SSDs (ESSDs) and meet your business requirements, you can create an advanced download task by point in time or backup set. You can set the download destination to an URL or directly upload the downloaded data to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving. For more information, see [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~).
 *
 * @param request CreateDownloadRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateDownloadResponse
 */
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

/**
 * For ApsaraDB RDS for MySQL instances that use standard SSDs or enhanced SSDs (ESSDs) and meet your business requirements, you can create an advanced download task by point in time or backup set. You can set the download destination to an URL or directly upload the downloaded data to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving. For more information, see [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~).
 *
 * @param request CreateDownloadRequest
 * @return CreateDownloadResponse
 */
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

/**
 * Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request CreateSandboxInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateSandboxInstanceResponse
 */
func (client *Client) CreateSandboxInstanceWithOptions(request *CreateSandboxInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateSandboxInstanceResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RestoreTime)) {
		query["RestoreTime"] = request.RestoreTime
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxInstanceName)) {
		query["SandboxInstanceName"] = request.SandboxInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxPassword)) {
		query["SandboxPassword"] = request.SandboxPassword
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxSpecification)) {
		query["SandboxSpecification"] = request.SandboxSpecification
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxType)) {
		query["SandboxType"] = request.SandboxType
	}

	if !tea.BoolValue(util.IsUnset(request.SandboxUser)) {
		query["SandboxUser"] = request.SandboxUser
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcSwitchId)) {
		query["VpcSwitchId"] = request.VpcSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSandboxInstance"),
		Version:     tea.String("2021-01-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSandboxInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request CreateSandboxInstanceRequest
 * @return CreateSandboxInstanceResponse
 */
func (client *Client) CreateSandboxInstance(request *CreateSandboxInstanceRequest) (_result *CreateSandboxInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSandboxInstanceResponse{}
	_body, _err := client.CreateSandboxInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request DeleteSandboxInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteSandboxInstanceResponse
 */
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

/**
 * This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request DeleteSandboxInstanceRequest
 * @return DeleteSandboxInstanceResponse
 */
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

/**
 * You can create an advanced download task by point in time or backup set. You can set the download destination to an URL or directly upload the downloaded data to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving. For more information, see [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~).
 *
 * @param request DescribeDownloadBackupSetStorageInfoRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDownloadBackupSetStorageInfoResponse
 */
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

/**
 * You can create an advanced download task by point in time or backup set. You can set the download destination to an URL or directly upload the downloaded data to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving. For more information, see [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~).
 *
 * @param request DescribeDownloadBackupSetStorageInfoRequest
 * @return DescribeDownloadBackupSetStorageInfoResponse
 */
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

/**
 * You can create an advanced download task by point in time or backup set. You can set the download destination to an URL or directly upload the downloaded data to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving. For more information, see [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~).
 *
 * @param request DescribeDownloadSupportRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDownloadSupportResponse
 */
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

/**
 * You can create an advanced download task by point in time or backup set. You can set the download destination to an URL or directly upload the downloaded data to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving. For more information, see [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~).
 *
 * @param request DescribeDownloadSupportRequest
 * @return DescribeDownloadSupportResponse
 */
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

/**
 * You can create an advanced download task by point in time or backup set. You can set the Download Destination parameter to URL or directly upload the downloaded data to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving. For more information, see [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~).
 *
 * @param request DescribeDownloadTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDownloadTaskResponse
 */
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

/**
 * You can create an advanced download task by point in time or backup set. You can set the Download Destination parameter to URL or directly upload the downloaded data to your Object Storage Service (OSS) bucket to facilitate data analysis and offline archiving. For more information, see [Download the backup files of an ApsaraDB RDS for MySQL instance](~~98819~~).
 *
 * @param request DescribeDownloadTaskRequest
 * @return DescribeDownloadTaskResponse
 */
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

/**
 * Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request DescribeSandboxBackupSetsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeSandboxBackupSetsResponse
 */
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

/**
 * Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only for the Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request DescribeSandboxBackupSetsRequest
 * @return DescribeSandboxBackupSetsResponse
 */
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

/**
 * This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request DescribeSandboxInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeSandboxInstancesResponse
 */
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

/**
 * This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request DescribeSandboxInstancesRequest
 * @return DescribeSandboxInstancesResponse
 */
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

/**
 * Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request DescribeSandboxRecoveryTimeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeSandboxRecoveryTimeResponse
 */
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

/**
 * Before you call this operation, you must enable the sandbox feature for the database instance. For more information, see [Use the emergency recovery feature of an ApsaraDB RDS for MySQL instance](~~203154~~) or [Create a sandbox instance for emergency disaster recovery of a self-managed MySQL database](~~185577~~). This operation is available only in Database Backup (DBS) API of the 2021-01-01 version.
 *
 * @param request DescribeSandboxRecoveryTimeRequest
 * @return DescribeSandboxRecoveryTimeResponse
 */
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
