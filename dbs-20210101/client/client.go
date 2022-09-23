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

type CreateDownloadRequest struct {
	BakSetId            *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	BakSetSize          *string `json:"BakSetSize,omitempty" xml:"BakSetSize,omitempty"`
	BakSetType          *string `json:"BakSetType,omitempty" xml:"BakSetType,omitempty"`
	DownloadPointInTime *string `json:"DownloadPointInTime,omitempty" xml:"DownloadPointInTime,omitempty"`
	FormatType          *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	InstanceName        *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionCode          *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	TargetBucket        *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	TargetOssRegion     *string `json:"TargetOssRegion,omitempty" xml:"TargetOssRegion,omitempty"`
	TargetOssUid        *string `json:"TargetOssUid,omitempty" xml:"TargetOssUid,omitempty"`
	TargetPath          *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetType          *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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

func (s *CreateDownloadRequest) SetTargetOssUid(v string) *CreateDownloadRequest {
	s.TargetOssUid = &v
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
	Code       *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *CreateDownloadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode    *string                         `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string                         `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                         `json:"Success,omitempty" xml:"Success,omitempty"`
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
	BackupSetTime  *int64  `json:"BackupSetTime,omitempty" xml:"BackupSetTime,omitempty"`
	BakSetId       *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	DbList         *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	DownloadStatus *string `json:"DownloadStatus,omitempty" xml:"DownloadStatus,omitempty"`
	ExportDataSize *int64  `json:"ExportDataSize,omitempty" xml:"ExportDataSize,omitempty"`
	Format         *string `json:"Format,omitempty" xml:"Format,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ImportDataSize *int64  `json:"ImportDataSize,omitempty" xml:"ImportDataSize,omitempty"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RegionCode     *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	TargetPath     *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetType     *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	BackupPlanId         *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupSetId          *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	SandboxInstanceName  *string `json:"SandboxInstanceName,omitempty" xml:"SandboxInstanceName,omitempty"`
	SandboxPassword      *string `json:"SandboxPassword,omitempty" xml:"SandboxPassword,omitempty"`
	SandboxSpecification *string `json:"SandboxSpecification,omitempty" xml:"SandboxSpecification,omitempty"`
	SandboxType          *string `json:"SandboxType,omitempty" xml:"SandboxType,omitempty"`
	SandboxUser          *string `json:"SandboxUser,omitempty" xml:"SandboxUser,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcSwitchId          *string `json:"VpcSwitchId,omitempty" xml:"VpcSwitchId,omitempty"`
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

type CreateSandboxInstanceResponseBody struct {
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *CreateSandboxInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode    *string                                `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string                                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                                `json:"Success,omitempty" xml:"Success,omitempty"`
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
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

type DeleteSandboxInstanceResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
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

type DescribeDownloadBackupSetStorageInfoRequest struct {
	BackupSetId  *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionCode   *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Code       *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *DescribeDownloadBackupSetStorageInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode    *string                                               `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string                                               `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                                               `json:"Success,omitempty" xml:"Success,omitempty"`
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
	ExpirationTime *int64  `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	PrivateUrl     *string `json:"PrivateUrl,omitempty" xml:"PrivateUrl,omitempty"`
	PublicUrl      *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
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
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionCode   *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
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
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
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
	BackupSetId  *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	CurrentPage  *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OrderColumn  *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	OrderDirect  *string `json:"OrderDirect,omitempty" xml:"OrderDirect,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionCode   *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	TaskType     *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
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
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *DescribeDownloadTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode    *string                               `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string                               `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                               `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Content    *DescribeDownloadTaskResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	Extra      *string                                      `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Number     *int64                                       `json:"Number,omitempty" xml:"Number,omitempty"`
	Size       *int64                                       `json:"Size,omitempty" xml:"Size,omitempty"`
	Total      *int64                                       `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalPages *int64                                       `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
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

func (s *DescribeDownloadTaskResponseBodyData) SetNumber(v int64) *DescribeDownloadTaskResponseBodyData {
	s.Number = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyData) SetSize(v int64) *DescribeDownloadTaskResponseBodyData {
	s.Size = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyData) SetTotal(v int64) *DescribeDownloadTaskResponseBodyData {
	s.Total = &v
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
	BackupSetTime  *string `json:"BackupSetTime,omitempty" xml:"BackupSetTime,omitempty"`
	BakSetId       *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	DbList         *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	DownloadStatus *string `json:"DownloadStatus,omitempty" xml:"DownloadStatus,omitempty"`
	ExportDataSize *string `json:"ExportDataSize,omitempty" xml:"ExportDataSize,omitempty"`
	Format         *string `json:"Format,omitempty" xml:"Format,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ImportDataSize *string `json:"ImportDataSize,omitempty" xml:"ImportDataSize,omitempty"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RegionCode     *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	TargetPath     *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	TargetType     *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	BackupSetId  *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
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
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *DescribeSandboxRecoveryTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrCode    *string                                      `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string                                      `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string                                      `json:"Success,omitempty" xml:"Success,omitempty"`
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
	BackupPlanId      *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	RecoveryBeginTime *string `json:"RecoveryBeginTime,omitempty" xml:"RecoveryBeginTime,omitempty"`
	RecoveryEndTime   *string `json:"RecoveryEndTime,omitempty" xml:"RecoveryEndTime,omitempty"`
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

	if !tea.BoolValue(util.IsUnset(request.TargetOssUid)) {
		query["TargetOssUid"] = request.TargetOssUid
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
